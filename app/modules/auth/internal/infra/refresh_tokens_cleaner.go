package infra

import (
	"context"
	"fmt"
	"log/slog"
	"runtime/debug"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type RefreshTokensCleaner struct {
	repo domain.RefreshTokenRepo
	log  *slog.Logger
	clk  clock.Clock
}

func NewRefreshTokensCleaner(
	repo domain.RefreshTokenRepo,
	log *slog.Logger,
	clk clock.Clock,
) *RefreshTokensCleaner {
	return &RefreshTokensCleaner{repo, log, clk}
}

func (s *RefreshTokensCleaner) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	for {
		select {
		case <-ticker.C:
			s.tick(ctx)
		case <-ctx.Done():
			return
		}
	}
}

func (s *RefreshTokensCleaner) tick(ctx context.Context) {
	defer func() {
		if r := recover(); r != nil {
			s.log.Error(fmt.Sprintf("[RefreshTokensCleaner]: tick panic: %v", r), "stack", string(debug.Stack()))
		}
	}()

	tokens, err := s.repo.Expired(ctx, s.clk.Now())
	if err != nil {
		s.log.Error(err.Error())
	}

	if len(tokens) == 0 {
		return
	}

	ids := make([]int, len(tokens))
	for _, t := range tokens {
		ids = append(ids, t.ID)
	}

	if err = s.repo.RemoveBatch(ctx, ids); err != nil {
		s.log.Error(err.Error())
	}

	s.log.Info(fmt.Sprintf("[RefreshTokensCleaner]: remove %d tokens", len(tokens)))
}

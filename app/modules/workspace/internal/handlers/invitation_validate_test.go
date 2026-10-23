package handlers_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/handlers"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type fakeInvitationRepo struct {
	getByHashFn func(ctx context.Context, hash string) (domain.Invitation, error)
}

func (f *fakeInvitationRepo) Create(ctx context.Context, i domain.Invitation) (domain.Invitation, error) {
	return domain.Invitation{}, nil
}
func (f *fakeInvitationRepo) GetByHash(ctx context.Context, hash string) (domain.Invitation, error) {
	if f.getByHashFn != nil {
		return f.getByHashFn(ctx, hash)
	}
	return domain.Invitation{}, domain.InvitationNotFoundError
}
func (f *fakeInvitationRepo) GetById(ctx context.Context, invitationId int) (domain.Invitation, error) {
	return domain.Invitation{}, nil
}
func (f *fakeInvitationRepo) ListByWorkspaceId(ctx context.Context, workspaceId int) ([]domain.Invitation, error) {
	return nil, nil
}
func (f *fakeInvitationRepo) Delete(ctx context.Context, invitationId int) error {
	return nil
}
func (f *fakeInvitationRepo) Update(ctx context.Context, i domain.Invitation) error {
	return nil
}

type fakeHasher struct{}

func (f fakeHasher) CreateToken() (string, string, error) {
	return "plain", "hashed_plain", nil
}
func (f fakeHasher) Hash(token string) string {
	return "hashed_" + token
}

func TestInvitationValidate_Handle(t *testing.T) {
	now := time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	clk := clock.NewMock(now)
	hasher := fakeHasher{}

	t.Run("valid invitation", func(t *testing.T) {
		repo := &fakeInvitationRepo{
			getByHashFn: func(ctx context.Context, hash string) (domain.Invitation, error) {
				if hash != "hashed_raw_token" {
					return domain.Invitation{}, domain.InvitationNotFoundError
				}
				return domain.Invitation{
					ID:        1,
					Email:     "user@example.com",
					Role:      "member",
					ExpiresAt: now.Add(24 * time.Hour),
				}, nil
			},
		}

		h := handlers.NewInvitationValidate(repo, hasher, clk)
		view, err := h.Handle(context.Background(), contract.InvitationValidate{Hash: "raw_token"})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if view.Email != "user@example.com" || view.Role != "member" {
			t.Fatalf("unexpected view: %+v", view)
		}
	})

	t.Run("expired invitation", func(t *testing.T) {
		repo := &fakeInvitationRepo{
			getByHashFn: func(ctx context.Context, hash string) (domain.Invitation, error) {
				return domain.Invitation{
					ID:        1,
					ExpiresAt: now.Add(-1 * time.Hour),
				}, nil
			},
		}

		h := handlers.NewInvitationValidate(repo, hasher, clk)
		_, err := h.Handle(context.Background(), contract.InvitationValidate{Hash: "raw_token"})
		if !errors.Is(err, domain.InvitationExpiredOrAcceptedError) {
			t.Fatalf("expected InvitationExpiredOrAcceptedError, got: %v", err)
		}
	})

	t.Run("revoked invitation", func(t *testing.T) {
		repo := &fakeInvitationRepo{
			getByHashFn: func(ctx context.Context, hash string) (domain.Invitation, error) {
				return domain.Invitation{
					ID:        1,
					ExpiresAt: now.Add(24 * time.Hour),
					RevokedAt: now.Add(-1 * time.Hour),
				}, nil
			},
		}

		h := handlers.NewInvitationValidate(repo, hasher, clk)
		_, err := h.Handle(context.Background(), contract.InvitationValidate{Hash: "raw_token"})
		if !errors.Is(err, domain.InvitationExpiredOrAcceptedError) {
			t.Fatalf("expected InvitationExpiredOrAcceptedError, got: %v", err)
		}
	})

	t.Run("accepted invitation", func(t *testing.T) {
		repo := &fakeInvitationRepo{
			getByHashFn: func(ctx context.Context, hash string) (domain.Invitation, error) {
				return domain.Invitation{
					ID:         1,
					ExpiresAt:  now.Add(24 * time.Hour),
					AcceptedAt: now.Add(-1 * time.Hour),
				}, nil
			},
		}

		h := handlers.NewInvitationValidate(repo, hasher, clk)
		_, err := h.Handle(context.Background(), contract.InvitationValidate{Hash: "raw_token"})
		if !errors.Is(err, domain.InvitationExpiredOrAcceptedError) {
			t.Fatalf("expected InvitationExpiredOrAcceptedError, got: %v", err)
		}
	})
}

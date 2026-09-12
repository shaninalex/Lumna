// app/modules/identity/internal/infra/reader.go
package infra

import (
	"context"
	"errors"

	"gitlab.com/shaninalex/lumna/app/modules/identity/internal/domain"
)

// Reader реалізує contract.Reader — рівно те, що identity готовий показати
// іншим модулям. Не домен, не record.
type Reader struct {
	repo domain.IdentityRepo
}

func NewReader(r domain.IdentityRepo) *Reader { return &Reader{repo: r} }

func (r *Reader) DisplayNames(ctx context.Context, ids []int) (map[int]string, error) {
	return r.repo.DisplayNames(ctx, ids)
}

func (r *Reader) Exists(ctx context.Context, id int) (bool, error) {
	_, err := r.repo.ByID(ctx, id)
	if errors.Is(err, domain.ErrNotFound) {
		return false, nil
	}
	return err == nil, err
}

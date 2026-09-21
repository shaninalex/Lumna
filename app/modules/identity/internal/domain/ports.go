package domain

import "context"

type IdentityRepo interface {
	Save(ctx context.Context, i Identity) (Identity, error)
	ByID(ctx context.Context, id int) (Identity, error)
	ByEmail(ctx context.Context, email string) (Identity, error)
	DisplayNames(ctx context.Context, ids []int) (map[int]string, error)
	List(ctx context.Context, limit, offset int) ([]Identity, error)
}

type CredentialRepo interface {
	Save(ctx context.Context, c Credential) (Credential, error)
	ByIdentityAndProvider(ctx context.Context, id int, p Provider) (Credential, error)
}

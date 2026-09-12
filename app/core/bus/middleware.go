package bus

import "context"

type Invoke func(ctx context.Context, msg any) (any, error)

type Middleware func(next Invoke) Invoke

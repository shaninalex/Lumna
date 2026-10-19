package infra

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/contract"
)

type EmailQueue struct {
}

func NewEmailQueue() contract.EmailQueue {
	return &EmailQueue{}
}

var _ contract.EmailQueue = (*EmailQueue)(nil)

func (s *EmailQueue) RunEmailQueue(ctx context.Context) {

}

package infra

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/modules/notifications/contract"
)

// EmailDispatcher drains the email_queue table and sends pending emails.
//
// Run it in exactly one process at a time: either `serve --email-queue` or the
// standalone email queue command, never both and never on several replicas.
// Pending rows are picked up without locking, so two dispatchers running at
// once will send the same email twice.
type EmailDispatcher struct {
}

func NewEmailDispatcher() contract.EmailDispatcher {
	return &EmailDispatcher{}
}

var _ contract.EmailDispatcher = (*EmailDispatcher)(nil)

func (s *EmailDispatcher) Run(ctx context.Context) {

}

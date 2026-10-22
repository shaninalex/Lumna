package notifications

import (
	"errors"
	"log/slog"

	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/notifications/contract"
	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/handlers"
	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/infra"
	"gitlab.com/shaninalex/lumna/app/modules/notifications/internal/infra/storage"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gitlab.com/shaninalex/lumna/app/platform/mail"
)

type Deps struct {
	DB       *database.DB
	Log      *slog.Logger
	Clock    clock.Clock
	EventBus *bus.EventBus
	Pusher   contract.Pusher
	Sender   mail.Sender
	BaseUrl  string
}

type Module struct {
	deps Deps

	// workers
	emailDispatcher contract.EmailDispatcher

	// event handlers
	workItemStageChanged *handlers.WorkItemStageChanged
	invitationCreated    *handlers.InvitationCreated
}

func New(d Deps) *Module {
	notificationsRepo := storage.NewNotificationRepo(d.DB)
	emailRepo := storage.NewEmailRepo(d.DB)
	return &Module{
		deps: d,

		// workers
		emailDispatcher: infra.NewEmailDispatcher(emailRepo, d.Log, d.Sender, d.Clock),

		// event handlers
		workItemStageChanged: handlers.NewWorkItemStageChanged(notificationsRepo, d.Clock, d.EventBus, d.Pusher),
		invitationCreated:    handlers.NewInvitationCreated(emailRepo, d.Clock, d.BaseUrl),
	}
}

func (s *Module) Name() string { return "notifications" }

func (s *Module) Register(a *core.App) error {
	return errors.Join(
		bus.Subscribe(a.Events, s.workItemStageChanged.Handle),
		bus.Subscribe(a.Events, s.invitationCreated.Handle),
	)
}

func (m *Module) EmailDispatcher() contract.EmailDispatcher { return m.emailDispatcher }

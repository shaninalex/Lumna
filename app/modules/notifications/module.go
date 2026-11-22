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
	workItemAssigned     *handlers.WorkItemAssigned
	workItemUnassigned   *handlers.WorkItemUnassigned
	workItemMoved        *handlers.WorkItemMoved
	invitationCreated    *handlers.InvitationCreated
	invitationAccepted   *handlers.InvitationAccepted
}

func New(d Deps) *Module {
	notificationsRepo := storage.NewNotificationRepo(d.DB)
	emailRepo := storage.NewEmailRepo(d.DB)
	deliverer := handlers.NewDeliverer(notificationsRepo, d.EventBus, d.Pusher)
	return &Module{
		deps: d,

		// workers
		emailDispatcher: infra.NewEmailDispatcher(emailRepo, d.Log, d.Sender, d.Clock),

		// event handlers
		workItemStageChanged: handlers.NewWorkItemStageChanged(d.Clock, deliverer),
		workItemAssigned:     handlers.NewWorkItemAssigned(d.Clock, deliverer),
		workItemUnassigned:   handlers.NewWorkItemUnassigned(d.Clock, deliverer),
		workItemMoved:        handlers.NewWorkItemMoved(d.Clock, deliverer),
		invitationCreated:    handlers.NewInvitationCreated(emailRepo, d.Clock, d.BaseUrl),
		invitationAccepted:   handlers.NewInvitationAccepted(d.Clock, deliverer),
	}
}

func (s *Module) Name() string { return "notifications" }

func (s *Module) Register(a *core.App) error {
	return errors.Join(
		bus.Subscribe(a.Events, s.workItemStageChanged.Handle),
		bus.Subscribe(a.Events, s.workItemAssigned.Handle),
		bus.Subscribe(a.Events, s.workItemUnassigned.Handle),
		bus.Subscribe(a.Events, s.invitationCreated.Handle),
		bus.Subscribe(a.Events, s.invitationAccepted.Handle),
	)
}

func (m *Module) EmailDispatcher() contract.EmailDispatcher { return m.emailDispatcher }

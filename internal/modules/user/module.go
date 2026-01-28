package module

import (
	ports2 "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/domain"
	userDomainPorts "github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	"github.com/racibaz/go-arch/internal/modules/user/features/signup/v1/application/commands"
	"github.com/racibaz/go-arch/pkg/ddd"
	"github.com/racibaz/go-arch/pkg/logger"
)

// UserModule encapsulates the components related to the User module.
type UserModule struct {
	repository                 userDomainPorts.UserRepository
	registerUserCommandHandler ports2.CommandHandler[commands.RegisterUserCommandV1]
	logger                     logger.Logger
	notifier                   userDomainPorts.NotificationAdapter
}

// NewUserModule initializes a new UserModule with the provided components.
func NewUserModule(
	repository userDomainPorts.UserRepository,
	registerUserCommandHandler ports2.CommandHandler[commands.RegisterUserCommandV1],
	logger logger.Logger,
	notifier userDomainPorts.NotificationAdapter,
) *UserModule {
	return &UserModule{
		repository:                 repository,
		registerUserCommandHandler: registerUserCommandHandler,
		logger:                     logger,
		notifier:                   notifier,
	}
}

func (m UserModule) Repository() userDomainPorts.UserRepository {
	return m.repository
}

func (m UserModule) RegisterUserCommandHandler() ports2.CommandHandler[commands.RegisterUserCommandV1] {
	return m.registerUserCommandHandler
}

func (m UserModule) Notifier() userDomainPorts.NotificationAdapter {
	return m.notifier
}

func (m UserModule) Logger() logger.Logger {
	return m.logger
}

// RegisterNotificationHandlers registers notification handlers for domain events.
func RegisterNotificationHandlers(
	notificationHandlers ddd.EventHandler[ddd.AggregateEvent],
	domainSubscriber ddd.EventSubscriber[ddd.AggregateEvent],
) {
	domainSubscriber.Subscribe(domain.UserRegisteredEvent, notificationHandlers)
}

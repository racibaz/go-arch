package module

import (
	ports2 "github.com/racibaz/go-arch/internal/modules/shared/application/ports"
	"github.com/racibaz/go-arch/internal/modules/user/domain"
	userDomainPorts "github.com/racibaz/go-arch/internal/modules/user/domain/ports"
	loginV1Query "github.com/racibaz/go-arch/internal/modules/user/features/login/v1/application/queries"
	logoutCommands "github.com/racibaz/go-arch/internal/modules/user/features/logout/v1/application/commands"
	meQueryV1 "github.com/racibaz/go-arch/internal/modules/user/features/me/v1/application/query"
	refreshTokenQueryV1 "github.com/racibaz/go-arch/internal/modules/user/features/refreshToken/v1/application/query"
	"github.com/racibaz/go-arch/internal/modules/user/features/signup/v1/application/commands"
	"github.com/racibaz/go-arch/pkg/ddd"
	"github.com/racibaz/go-arch/pkg/logger"
)

// UserModule encapsulates the components related to the User module.
type UserModule struct {
	repository           userDomainPorts.UserRepository
	signupCommandHandler ports2.CommandHandler[commands.RegisterUserCommandV1]
	loginQueryHandler    ports2.QueryHandler[loginV1Query.LoginQueryV1, *loginV1Query.LoginQueryResponse]
	logoutCommandHandler ports2.CommandHandler[logoutCommands.LogoutCommandV1]
	refreshTokenHandler  ports2.QueryHandler[refreshTokenQueryV1.RefreshTokenQueryV1, *refreshTokenQueryV1.RefreshTokenQueryResponseV1]
	meHandler            ports2.QueryHandler[meQueryV1.MeQueryHandlerQuery, *meQueryV1.MeQueryHandlerResponse]
	logger               logger.Logger
	notifier             userDomainPorts.NotificationAdapter
}

// NewUserModule initializes a new UserModule with the provided components.
func NewUserModule(
	repository userDomainPorts.UserRepository,
	registerUserCommandHandler ports2.CommandHandler[commands.RegisterUserCommandV1],
	loginQueryHandler ports2.QueryHandler[loginV1Query.LoginQueryV1, *loginV1Query.LoginQueryResponse],
	logoutCommandHandler ports2.CommandHandler[logoutCommands.LogoutCommandV1],
	refreshTokenHandler ports2.QueryHandler[refreshTokenQueryV1.RefreshTokenQueryV1, *refreshTokenQueryV1.RefreshTokenQueryResponseV1],
	meHandler ports2.QueryHandler[meQueryV1.MeQueryHandlerQuery, *meQueryV1.MeQueryHandlerResponse],
	logger logger.Logger,
	notifier userDomainPorts.NotificationAdapter,
) *UserModule {
	return &UserModule{
		repository:           repository,
		signupCommandHandler: registerUserCommandHandler,
		loginQueryHandler:    loginQueryHandler,
		logoutCommandHandler: logoutCommandHandler,
		refreshTokenHandler:  refreshTokenHandler,
		meHandler:            meHandler,
		logger:               logger,
		notifier:             notifier,
	}
}

func (m UserModule) Repository() userDomainPorts.UserRepository {
	return m.repository
}

func (m UserModule) RegisterUserCommandHandler() ports2.CommandHandler[commands.RegisterUserCommandV1] {
	return m.signupCommandHandler
}

func (m UserModule) LoginQueryHandler() ports2.QueryHandler[loginV1Query.LoginQueryV1, *loginV1Query.LoginQueryResponse] {
	return m.loginQueryHandler
}

func (m UserModule) LogoutCommandHandler() ports2.CommandHandler[logoutCommands.LogoutCommandV1] {
	return m.logoutCommandHandler
}

func (m UserModule) RefreshTokenHandler() ports2.QueryHandler[refreshTokenQueryV1.RefreshTokenQueryV1, *refreshTokenQueryV1.RefreshTokenQueryResponseV1] {
	return m.refreshTokenHandler
}

func (m UserModule) MeHandler() ports2.QueryHandler[meQueryV1.MeQueryHandlerQuery, *meQueryV1.MeQueryHandlerResponse] {
	return m.meHandler
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

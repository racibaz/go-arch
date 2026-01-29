package module

import (
	"sync"
	"testing"

	userQueries "github.com/racibaz/go-arch/internal/modules/user/features/login/v1/application/queries"
	userCommands "github.com/racibaz/go-arch/internal/modules/user/features/signup/v1/application/commands"
	userNotification "github.com/racibaz/go-arch/internal/modules/user/infrastructure/notification/sms"
	gormUserRepo "github.com/racibaz/go-arch/internal/modules/user/infrastructure/persistence/gorm/repositories"
	"github.com/racibaz/go-arch/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestBuildModule(t *testing.T) {
	// Reset the singleton instance for clean testing
	userModuleInstance = nil
	once = sync.Once{}

	module := BuildModule()

	assert.NotNil(t, module)
	assert.NotNil(t, module.Repository())
	assert.NotNil(t, module.RegisterUserCommandHandler())
	assert.NotNil(t, module.LoginQueryHandler())
	assert.NotNil(t, module.Logger())
	assert.NotNil(t, module.Notifier())

	// Further checks can be added to verify the concrete types of the returned components
	// For example:
	assert.IsType(t, &gormUserRepo.GormUserRepository{}, module.Repository())
	assert.IsType(t, &userCommands.RegisterUserHandler{}, module.RegisterUserCommandHandler())
	assert.IsType(t, &userQueries.LoginHandler{}, module.LoginQueryHandler())
	assert.IsType(t, &logger.ZapLogger{}, module.Logger())
	assert.IsType(t, userNotification.TwilioSmsNotificationAdapter{}, module.Notifier())
}

// TODO: Implement integration tests for Routes and GrpcRoutes.
// Directly unit testing these functions (which register routes on a gin.Engine or grpc.Server)
// would require extensive mocking of gin.Engine/grpc.Server and the MapHttpRoute functions,
// which is beyond the scope of a simple test file and more suited for integration tests.

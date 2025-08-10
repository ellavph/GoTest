package container

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"TestGO/internal/application/services"
	"TestGO/internal/domain/repositories"
	interfaceServices "TestGO/internal/interfaces/services"
	sqlRepo "TestGO/internal/infrastructure/database/sql"
	"TestGO/internal/infrastructure/security"
	"TestGO/internal/interfaces/http/handlers"
	"TestGO/internal/interfaces/http/middleware"
)

// Container gerencia todas as dependências da aplicação
type Container struct {
	// Repositories
	UserRepository      repositories.UserRepository
	CompanyRepository   repositories.CompanyRepository
	TestSuiteRepository repositories.TestSuiteRepository

	// Services
	AuthService      interfaceServices.AuthService
	UserService      interfaceServices.UserService
	CompanyService   interfaceServices.CompanyService
	TestSuiteService interfaceServices.TestSuiteService

	// Infrastructure Services
	PasswordService *security.PasswordService
	JWTService      *security.JWTService

	// Handlers
	AuthHandler      *handlers.AuthHandler
	UserHandler      *handlers.UserHandler
	CompanyHandler   *handlers.CompanyHandler
	TestSuiteHandler *handlers.TestSuiteHandler

	// Middleware
	AuthMiddleware *middleware.AuthMiddleware
}

// NewContainer cria uma nova instância do container
func NewContainer(db *pgxpool.Pool, jwtSecret string) *Container {
	// Infrastructure Services
	passwordService := security.NewPasswordService(12)
	jwtService := security.NewJWTService(
		jwtSecret,
		15*time.Minute, // Access token expiry
		7*24*time.Hour, // Refresh token expiry
	)

	// Repositories
	userRepo := sqlRepo.NewUserRepository(db)
	companyRepo := sqlRepo.NewCompanyRepository(db)
	testSuiteRepo := sqlRepo.NewTestSuiteRepository(db)

	// Application Services
	authService := services.NewAuthService(userRepo, passwordService, jwtService)
	userService := services.NewUserService(userRepo, passwordService)
	companyService := services.NewCompanyService(companyRepo)
	testSuiteService := services.NewTestSuiteService(testSuiteRepo)

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	companyHandler := handlers.NewCompanyHandler(companyService)
	testSuiteHandler := handlers.NewTestSuiteHandler(testSuiteService)

	// Middleware
	authMiddleware := middleware.NewAuthMiddleware(authService)

	return &Container{
		// Repositories
		UserRepository:      userRepo,
		CompanyRepository:   companyRepo,
		TestSuiteRepository: testSuiteRepo,

		// Services
		AuthService:      authService,
		UserService:      userService,
		CompanyService:   companyService,
		TestSuiteService: testSuiteService,

		// Infrastructure Services
		PasswordService: passwordService,
		JWTService:      jwtService,

		// Handlers
		AuthHandler:      authHandler,
		UserHandler:      userHandler,
		CompanyHandler:   companyHandler,
		TestSuiteHandler: testSuiteHandler,

		// Middleware
		AuthMiddleware: authMiddleware,
	}
}
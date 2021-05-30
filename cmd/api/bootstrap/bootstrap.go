package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"

	mooc "github.com/LuisCusihuaman/go-hexagonal-http-api/internal"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/creating"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/increasing"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/platform/bus/inmemory"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/platform/server"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func Run() error {
	var cfg config
	err := envconfig.Process("MOOC", &cfg)
	if err != nil {
		return err
	}

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)
	courseRepository := mysql.NewCourseRepository(db, cfg.DbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)

	increasingCourseCounterService := increasing.NewCourseCounterService()
	increaseCoursesHandler := creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseCounterService)

	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)
	eventBus.Subscribe(mooc.CourseCreatedEventType, increaseCoursesHandler)

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `default:"user"`
	DbPass    string        `default:"password"`
	DbHost    string        `default:"localhost"`
	DbPort    uint          `default:"3306"`
	DbName    string        `default:"backoffice"`
	DbTimeout time.Duration `default:"5s"`
}

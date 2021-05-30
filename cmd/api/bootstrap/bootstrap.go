package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	mooc "github.com/LuisCusihuaman/go-hexagonal-http-api/internal"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/creating"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/increasing"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/platform/bus/inmemory"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/platform/server"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	host            = "localhost"
	port            = 8080
	shutdownTimeout = 10 * time.Second

	dbUser    = "codely"
	dbPass    = "codely"
	dbHost    = "localhost"
	dbPort    = "3306"
	dbName    = "codely"
	dbTimeout = 5 * time.Second
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)
	courseRepository := mysql.NewCourseRepository(db, dbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)

	increasingCourseCounterService := increasing.NewCourseCounterService()
	increaseCoursesHandler := creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseCounterService)

	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)
	eventBus.Subscribe(mooc.CourseCreatedEventType, increaseCoursesHandler)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}

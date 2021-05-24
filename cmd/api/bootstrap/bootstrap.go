package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/creating"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/platform/bus/inmemory"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/platform/server"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "codely"
	dbPass = "codely"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	commandBus := inmemory.NewCommandBus()

	courseRepository := mysql.NewCourseRepository(db)

	creatingCourseService := creating.NewCourseService(courseRepository)

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)

	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	srv := server.New(host, port, commandBus)
	return srv.Run()
}

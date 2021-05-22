package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/platform/server"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/platform/storage/mysql"
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

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}

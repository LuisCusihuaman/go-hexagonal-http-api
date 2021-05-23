package mysql

const (
	sqlCourseTable = "courses"
)

type dtoCourse struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Duration string `db:"duration"`
}

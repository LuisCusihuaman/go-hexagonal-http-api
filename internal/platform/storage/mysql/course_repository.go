package mysql

import (
	"context"
	"database/sql"
	"fmt"

	mooc "github.com/LuisCusihuaman/go-hexagonal-http-api/internal"
	"github.com/huandu/go-sqlbuilder"
)

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{db: db}
}
func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID(),
		Name:     course.Name(),
		Duration: course.Duration(),
	}).Build()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}
	return nil
}

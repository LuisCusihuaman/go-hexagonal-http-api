package mysql

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	mooc "github.com/LuisCusihuaman/go-hexagonal-http-api/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestCourseRepository_Save_RepositoryError(t *testing.T) {
	courseID, courseName, courseDuration := "195a17a1-fa2d-46f8-97e9-da292f1e99b5", "Test Course", "2 months"
	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)").
		WithArgs(courseID, courseName, courseDuration).
		WillReturnError(errors.New("something-failed"))

	repo := NewCourseRepository(db, 1*time.Millisecond)
	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet()) // on repo.Save expect same query
	assert.Error(t, err)
}

func Test_CourseRepository_Save_Succeed(t *testing.T) {
	courseID, courseName, courseDuration := "195a17a1-fa2d-46f8-97e9-da292f1e99b5", "Test Course", "2 months"
	course, err := mooc.NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)").
		WithArgs(courseID, courseName, courseDuration).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewCourseRepository(db, 1*time.Millisecond)

	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}

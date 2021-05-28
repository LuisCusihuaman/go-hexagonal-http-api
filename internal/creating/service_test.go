package creating

import (
	"context"
	"errors"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/platform/storage/storagemocks"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/kit/event/eventmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCourseService(t *testing.T) {
	courseID, courseName, courseDuration := "1b9ef50d-74f2-4a96-b99d-0d5ebb178d18", "Test Course", "1 month"

	t.Run("course created successfully", func(t *testing.T) {
		courseRepositoryMock, eventBusMock := new(storagemocks.CourseRepository), new(eventmocks.Bus)
		courseRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)
		eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

		courseService := NewCourseService(courseRepositoryMock, eventBusMock)
		err := courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

		courseRepositoryMock.AssertExpectations(t)
		eventBusMock.AssertExpectations(t)
		assert.NoError(t, err)
	})

	t.Run("course create fails by repository", func(t *testing.T) {
		courseRepositoryMock, eventBusMock := new(storagemocks.CourseRepository), new(eventmocks.Bus)
		courseRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).
			Return(errors.New("something unexpected happend"))

		courseService := NewCourseService(courseRepositoryMock, eventBusMock)
		err := courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

		courseRepositoryMock.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("course create fails by event bus", func(t *testing.T) {
		courseRepositoryMock, eventBusMock := new(storagemocks.CourseRepository), new(eventmocks.Bus)
		courseRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)
		eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).
			Return(errors.New("something unexpected happend"))

		courseService := NewCourseService(courseRepositoryMock, eventBusMock)
		err := courseService.CreateCourse(context.Background(), courseID, courseName, courseDuration)

		eventBusMock.AssertExpectations(t)
		assert.Error(t, err)
	})
}

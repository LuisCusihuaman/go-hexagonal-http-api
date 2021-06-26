package increasing

import (
	"context"
	"errors"
	mooc "github.com/LuisCusihuaman/go-hexagonal-http-api/internal"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/kit/event"
)

type IncreaseCoursesCounterOnCourseCreated struct {
	increasingService CourseCounterService
}

func NewIncreaseCoursesCounterOnCourseCreated(increaserService CourseCounterService) IncreaseCoursesCounterOnCourseCreated {
	return IncreaseCoursesCounterOnCourseCreated{increasingService: increaserService}
}

func (h IncreaseCoursesCounterOnCourseCreated) Handle(_ context.Context, evt event.Event) error {
	courseCreateEvt, ok := evt.(mooc.CourseCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}
	return h.increasingService.Increase(courseCreateEvt.ID())
}

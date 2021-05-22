package courses

import (
	"bytes"
	"encoding/json"
	"github.com/LuisCusihuaman/go-hexagonal-http-api/internal/platform/storage/storagemocks"
	"github.com/gin-gonic/gin"
	"github.com/huandu/go-assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	courseRepository := new(storagemocks.CourseRepository)
	courseRepository.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/courses", CreateHandler(courseRepository))

	t.Run("given a invalid request it returns 400", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:   "b24f2566-9109-4dc3-830a-54ca80eb6110",
			Name: "Demo Course",
		}
		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.Body)

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:       "8a1c5cdc-ba57-445a-994d-aa412d23723f",
			Name:     "Demo Course",
			Duration: "10 months",
		}

		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(res.Body)

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}

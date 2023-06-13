package test_controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/aabdullahgungor/personal-resume-api/internal/controller"
	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/aabdullahgungor/personal-resume-api/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestUniversityController_GetAllUniversities(t *testing.T) {

	t.Run("Error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]model.University{}, errors.New("hata!")).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.GetAllUniversities(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]model.University{}, nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.GetAllUniversities(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)

	})
}

func TestUniversityController_GetUniversityById(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(model.University{}, service.ErrUniversityNotFound).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		params := []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		ctx.Params = params
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.GetUniversityById(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(model.University{
			UniversityName: "Ataturk University",
		}, nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		params := []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		ctx.Params = params
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.GetUniversityById(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)

		var responseData model.University
		json.NewDecoder(w.Body).Decode(&responseData)
		assert.Equal(t, "Ataturk University", responseData.UniversityName)
		t.Log("\nUniversity name is: ", responseData.UniversityName)

	})
}

func TestUniversityController_CreateUniversity(t *testing.T) {
	t.Run("ErrorCreate", func(t *testing.T) {
		university := model.University{
			UniversityName: "Ataturk University",
		}
		jsonValue, _ := json.Marshal(university)
		byteUniversity := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().Create(&university).Return(errors.New("hata")).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteUniversity)
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.CreateUniversity(ctx)

		assert.Equal(t, http.StatusNotAcceptable, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		university := model.University{
			UniversityName: "Ataturk University",
		}
		jsonValue, _ := json.Marshal(university)
		byteUniversity := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().Create(&university).Return(nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteUniversity)
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.CreateUniversity(ctx)

		assert.Equal(t, http.StatusCreated, w.Code)
		t.Log(w.Body.String())

	})
}

func TestUniversityController_EditUniversity(t *testing.T) {
	t.Run("ErrorEdit", func(t *testing.T) {
		university := model.University{
			UniversityName: "Ataturk University",
		}
		jsonValue, _ := json.Marshal(university)
		byteUniversity := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().Edit(&university).Return(errors.New("hata")).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "PUT"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteUniversity)
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.EditUniversity(ctx)

		assert.Equal(t, http.StatusNotAcceptable, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		university := model.University{
			UniversityName: "Ataturk University",
		}
		jsonValue, _ := json.Marshal(university)
		byteUniversity := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().Edit(&university).Return(nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "PUT"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteUniversity)
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.EditUniversity(ctx)

		assert.Equal(t, http.StatusCreated, w.Code)
		t.Log(w.Body.String())
	})
}

func TestUniversityController_DeleteUniversity(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(service.ErrUniversityNotFound).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "DELETE"
		ctx.Request.Header.Set("Content-Type", "application/json")
		params := []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		ctx.Params = params
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.DeleteUniversity(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "DELETE"
		ctx.Request.Header.Set("Content-Type", "application/json")
		params := []gin.Param{
			{
				Key:   "id",
				Value: "1",
			},
		}
		ctx.Params = params
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.DeleteUniversity(ctx)

		assert.Equal(t, http.StatusAccepted, w.Code)
		t.Log(w.Body.String())
	})
}

package test_controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
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

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.GetAllUniversities(ctx)

		req, _ := http.NewRequest("GET", "api/v1/universities", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]model.University{}, nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.GetAllUniversities(ctx)

		req, _ := http.NewRequest("GET", "api/v1/universities", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)

	})
}

func TestUniversityController_GetUniversityById(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(model.University{}, service.ErrUniversityNotFound).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.GetUniversityById(ctx)

		req, _ := http.NewRequest("GET", "api/v1/universities/:id", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(model.University{
			UniversityName: "Ataturk University",
		}, nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.GetUniversityById(ctx)

		req, _ := http.NewRequest("GET", "api/v1/universites/:id", nil)
		r.ServeHTTP(w, req)
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

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.CreateUniversity(ctx)
		req, err := http.NewRequest("POST", "api/v1/universities", byteUniversity)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
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

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.CreateUniversity(ctx)
		req, err := http.NewRequest("POST", "api/v1/universities", byteUniversity)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
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

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.EditUniversity(ctx)
		req, err := http.NewRequest("PUT", "api/v1/universities", byteUniversity)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
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

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.EditUniversity(ctx)
		req, err := http.NewRequest("PUT", "api/v1/universities", byteUniversity)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		t.Log(w.Body.String())
	})
}

func TestUniversityController_DeleteUniversity(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(service.ErrUniversityNotFound).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, _ := gin.CreateTestContext(w)
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.DeleteUniversity(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIUniversityService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, _ := gin.CreateTestContext(w)
		universityTestController := controller.NewUniversityController(mockService)
		universityTestController.DeleteUniversity(ctx)

		assert.Equal(t, http.StatusAccepted, w.Code)
		t.Log(w.Body.String())
	})
}

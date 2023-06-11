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

func TestExperienceController_GetAllExperiences(t *testing.T) {

	t.Run("Error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIExperienceService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]model.Experience{}, errors.New("hata!")).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		experienceTestController := controller.NewExperienceController(mockService)
		experienceTestController.GetAllExperiences(ctx)

		req, _ := http.NewRequest("GET", "api/v1/experiences", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIExperienceService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]model.Experience{}, nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		experienceTestController := controller.NewExperienceController(mockService)
		experienceTestController.GetAllExperiences(ctx)

		req, _ := http.NewRequest("GET", "api/v1/experiences", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)

	})
}

func TestExperienceController_GetExperienceById(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIExperienceService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(model.Experience{}, service.ErrExperienceNotFound).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		experienceTestController := controller.NewExperienceController(mockService)
		experienceTestController.GetExperienceById(ctx)

		req, _ := http.NewRequest("GET", "api/v1/experiences/:id", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIExperienceService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(model.Experience{
			CompanyName: "EMFA",
			Position:    "Sales Support Engineer",
			StartYear:   "2019-08-01T00:00:00Z",
			FinishYear:  "2020-02-14T00:00:00Z",
			PersonalID:  1,
		}, nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		experienceTestController := controller.NewExperienceController(mockService)
		experienceTestController.GetExperienceById(ctx)

		req, _ := http.NewRequest("GET", "api/v1/experiences/:id", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)

		var responseData model.Experience
		json.NewDecoder(w.Body).Decode(&responseData)
		assert.Equal(t, "EMFA", responseData.CompanyName)
		t.Log("\nCompany name is: ", responseData.CompanyName)

	})
}

func TestExperienceController_CreateExperience(t *testing.T) {
	t.Run("ErrorCreate", func(t *testing.T) {
		experience := model.Experience{
			CompanyName: "EMFA",
			Position:    "Sales Support Engineer",
			StartYear:   "2019-08-01T00:00:00Z",
			FinishYear:  "2020-02-14T00:00:00Z",
			PersonalID:  1,
		}
		jsonValue, _ := json.Marshal(experience)
		byteExperience := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIExperienceService(mockCtrl)
		mockService.EXPECT().Create(&experience).Return(errors.New("hata")).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		experienceTestController := controller.NewExperienceController(mockService)
		experienceTestController.CreateExperience(ctx)
		req, err := http.NewRequest("POST", "api/v1/experiences", byteExperience)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotAcceptable, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		experience := model.Experience{
			CompanyName: "EMFA",
			Position:    "Sales Support Engineer",
			StartYear:   "2019-08-01T00:00:00Z",
			FinishYear:  "2020-02-14T00:00:00Z",
			PersonalID:  1,
		}
		jsonValue, _ := json.Marshal(experience)
		byteExperience := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIExperienceService(mockCtrl)
		mockService.EXPECT().Create(&experience).Return(nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		experienceTestController := controller.NewExperienceController(mockService)
		experienceTestController.CreateExperience(ctx)
		req, err := http.NewRequest("POST", "api/v1/experiences", byteExperience)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		t.Log(w.Body.String())

	})
}

func TestExperienceController_EditExperience(t *testing.T) {
	t.Run("ErrorEdit", func(t *testing.T) {
		experience := model.Experience{
			CompanyName: "EMFA",
			Position:    "Sales Support Engineer",
			StartYear:   "2019-08-01T00:00:00Z",
			FinishYear:  "2020-02-14T00:00:00Z",
			PersonalID:  1,
		}
		jsonValue, _ := json.Marshal(experience)
		byteExperience := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIExperienceService(mockCtrl)
		mockService.EXPECT().Edit(&experience).Return(errors.New("hata")).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		experienceTestController := controller.NewExperienceController(mockService)
		experienceTestController.EditExperience(ctx)
		req, err := http.NewRequest("PUT", "api/v1/experiences", byteExperience)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotAcceptable, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		experience := model.Experience{
			CompanyName: "EMFA",
			Position:    "Sales Support Engineer",
			StartYear:   "2019-08-01T00:00:00Z",
			FinishYear:  "2020-02-14T00:00:00Z",
			PersonalID:  1,
		}
		jsonValue, _ := json.Marshal(experience)
		byteExperience := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIExperienceService(mockCtrl)
		mockService.EXPECT().Edit(&experience).Return(nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		experienceTestController := controller.NewExperienceController(mockService)
		experienceTestController.EditExperience(ctx)
		req, err := http.NewRequest("PUT", "api/v1/experiences", byteExperience)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		t.Log(w.Body.String())
	})
}

func TestExperienceController_DeleteExperience(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIExperienceService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(service.ErrExperienceNotFound).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, _ := gin.CreateTestContext(w)
		experienceTestController := controller.NewExperienceController(mockService)
		experienceTestController.DeleteExperience(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIExperienceService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, _ := gin.CreateTestContext(w)
		experienceTestController := controller.NewExperienceController(mockService)
		experienceTestController.DeleteExperience(ctx)

		assert.Equal(t, http.StatusAccepted, w.Code)
		t.Log(w.Body.String())
	})
}

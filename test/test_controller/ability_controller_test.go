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

func TestAbilityController_GetAllAbilities(t *testing.T) {

	t.Run("Error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIAbilityService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]model.Ability{}, errors.New("hata!")).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		abilityTestController := controller.NewAbilityController(mockService)
		abilityTestController.GetAllAbilities(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIAbilityService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]model.Ability{}, nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		abilityTestController := controller.NewAbilityController(mockService)
		abilityTestController.GetAllAbilities(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)

	})
}

func TestAbilityController_GetAbilityById(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIAbilityService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(model.Ability{}, service.ErrAbilityNotFound).AnyTimes()

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
		abilityTestController := controller.NewAbilityController(mockService)
		abilityTestController.GetAbilityById(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIAbilityService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(model.Ability{
			AbilityName: "Go",
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
		abilityTestController := controller.NewAbilityController(mockService)
		abilityTestController.GetAbilityById(ctx)
		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)

		var responseData model.Ability
		json.NewDecoder(w.Body).Decode(&responseData)
		assert.Equal(t, "Go", responseData.AbilityName)
		t.Log("\nAbility name is: ", responseData.AbilityName)

	})
}

func TestAbilityController_CreateAbility(t *testing.T) {
	t.Run("ErrorCreate", func(t *testing.T) {
		ability := model.Ability{
			AbilityName: "Go",
		}
		jsonValue, _ := json.Marshal(ability)
		byteAbility := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIAbilityService(mockCtrl)
		mockService.EXPECT().Create(&ability).Return(errors.New("hata")).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteAbility)
		abilityTestController := controller.NewAbilityController(mockService)
		abilityTestController.CreateAbility(ctx)

		assert.Equal(t, http.StatusNotAcceptable, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		ability := model.Ability{
			AbilityName: "Go",
		}
		jsonValue, _ := json.Marshal(ability)
		byteAbility := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIAbilityService(mockCtrl)
		mockService.EXPECT().Create(&ability).Return(nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteAbility)
		abilityTestController := controller.NewAbilityController(mockService)
		abilityTestController.CreateAbility(ctx)

		assert.Equal(t, http.StatusCreated, w.Code)
		t.Log(w.Body.String())

	})
}

func TestAbilityController_EditAbility(t *testing.T) {
	t.Run("ErrorEdit", func(t *testing.T) {
		ability := model.Ability{
			AbilityName: "Go",
		}
		jsonValue, _ := json.Marshal(ability)
		byteAbility := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIAbilityService(mockCtrl)
		mockService.EXPECT().Edit(&ability).Return(errors.New("hata")).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "PUT"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteAbility)
		abilityTestController := controller.NewAbilityController(mockService)
		abilityTestController.EditAbility(ctx)

		assert.Equal(t, http.StatusNotAcceptable, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		ability := model.Ability{
			AbilityName: "Go",
		}
		jsonValue, _ := json.Marshal(ability)
		byteAbility := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIAbilityService(mockCtrl)
		mockService.EXPECT().Edit(&ability).Return(nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "PUT"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(byteAbility)
		abilityTestController := controller.NewAbilityController(mockService)
		abilityTestController.EditAbility(ctx)

		assert.Equal(t, http.StatusCreated, w.Code)
		t.Log(w.Body.String())
	})
}

func TestAbilityController_DeleteAbility(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIAbilityService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(service.ErrAbilityNotFound).AnyTimes()

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
		abilityTestController := controller.NewAbilityController(mockService)
		abilityTestController.DeleteAbility(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIAbilityService(mockCtrl)
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
		abilityTestController := controller.NewAbilityController(mockService)
		abilityTestController.DeleteAbility(ctx)

		assert.Equal(t, http.StatusAccepted, w.Code)
		t.Log(w.Body.String())
	})
}

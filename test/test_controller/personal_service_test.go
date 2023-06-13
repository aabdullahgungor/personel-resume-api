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

func TestPersonalController_GetAllPersonals(t *testing.T) {

	t.Run("Error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]model.Personal{}, errors.New("hata!")).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.GetAllPersonals(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]model.Personal{}, nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "GET"
		ctx.Request.Header.Set("Content-Type", "application/json")
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.GetAllPersonals(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)

	})
}

func TestPersonalController_GetPersonalById(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(model.Personal{}, service.ErrPersonalNotFound).AnyTimes()

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
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.GetPersonalById(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(model.Personal{
			Name:     "Abdullah",
			Surname:  "Gungor",
			UserName: "abdullahgungor",
			Email:    "abdullahgungor@hotmail.com.tr",
			Password: "123456",
			UserType: "admin",
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
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.GetPersonalById(ctx)

		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)

		var responseData model.Personal
		json.NewDecoder(w.Body).Decode(&responseData)
		assert.Equal(t, "Abdullah", responseData.Name)
		t.Log("\nName is: ", responseData.Name)

	})
}

func TestPersonalController_CreatePersonal(t *testing.T) {
	t.Run("ErrorCreate", func(t *testing.T) {
		personal := model.Personal{
			Name:     "Abdullah",
			Surname:  "Gungor",
			UserName: "abdullahgungor",
			Email:    "abdullahgungor@hotmail.com.tr",
			Password: "123456",
			UserType: "admin",
		}
		jsonValue, _ := json.Marshal(personal)
		bytePersonal := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
		mockService.EXPECT().Create(&personal).Return(errors.New("hata")).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(bytePersonal)
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.CreatePersonal(ctx)

		assert.Equal(t, http.StatusNotAcceptable, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		personal := model.Personal{
			Name:     "Abdullah",
			Surname:  "Gungor",
			UserName: "abdullahgungor",
			Email:    "abdullahgungor@hotmail.com.tr",
			Password: "123456",
			UserType: "admin",
		}
		jsonValue, _ := json.Marshal(personal)
		bytePersonal := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
		mockService.EXPECT().Create(&personal).Return(nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "POST"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(bytePersonal)
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.CreatePersonal(ctx)

		assert.Equal(t, http.StatusCreated, w.Code)
		t.Log(w.Body.String())

	})
}

func TestPersonalController_EditPersonal(t *testing.T) {
	t.Run("ErrorEdit", func(t *testing.T) {
		personal := model.Personal{
			Name:     "Abdullah",
			Surname:  "Gungor",
			UserName: "abdullahgungor",
			Email:    "abdullahgungor@hotmail.com.tr",
			Password: "123456",
			UserType: "admin",
		}
		jsonValue, _ := json.Marshal(personal)
		bytePersonal := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
		mockService.EXPECT().Edit(&personal).Return(errors.New("hata")).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "PUT"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(bytePersonal)
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.EditPersonal(ctx)

		assert.Equal(t, http.StatusNotAcceptable, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		personal := model.Personal{
			Name:     "Abdullah",
			Surname:  "Gungor",
			UserName: "abdullahgungor",
			Email:    "abdullahgungor@hotmail.com.tr",
			Password: "123456",
			UserType: "admin",
		}
		jsonValue, _ := json.Marshal(personal)
		bytePersonal := bytes.NewBuffer(jsonValue)

		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
		mockService.EXPECT().Edit(&personal).Return(nil).AnyTimes()

		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Header: make(http.Header),
			URL:    &url.URL{},
		}
		ctx.Request.Method = "PUT"
		ctx.Request.Header.Set("Content-Type", "application/json")
		ctx.Request.Body = io.NopCloser(bytePersonal)
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.EditPersonal(ctx)

		assert.Equal(t, http.StatusCreated, w.Code)
		t.Log(w.Body.String())
	})
}

func TestPersonalController_DeletePersonal(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(service.ErrPersonalNotFound).AnyTimes()

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
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.DeletePersonal(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		id := "1"
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
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
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.DeletePersonal(ctx)

		assert.Equal(t, http.StatusAccepted, w.Code)
		t.Log(w.Body.String())
	})
}

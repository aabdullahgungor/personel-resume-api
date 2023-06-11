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

func TestPersonalController_GetAllPersonals(t *testing.T) {

	t.Run("Error", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]model.Personal{}, errors.New("hata!")).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.GetAllPersonals(ctx)

		req, _ := http.NewRequest("GET", "api/v1/personals", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
		mockService.EXPECT().GetAll().Return([]model.Personal{}, nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.GetAllPersonals(ctx)

		req, _ := http.NewRequest("GET", "api/v1/personals", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println(w.Code)

	})
}

func TestPersonalController_GetPersonalById(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
		mockService.EXPECT().GetById(id).Return(model.Personal{}, service.ErrPersonalNotFound).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.GetPersonalById(ctx)

		req, _ := http.NewRequest("GET", "api/v1/personals/:id", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNotFound, w.Code)
		fmt.Println(w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		var id string
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

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.GetPersonalById(ctx)

		req, _ := http.NewRequest("GET", "api/v1/personals/:id", nil)
		r.ServeHTTP(w, req)
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

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.CreatePersonal(ctx)
		req, err := http.NewRequest("POST", "api/v1/personals", bytePersonal)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
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

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.CreatePersonal(ctx)
		req, err := http.NewRequest("POST", "api/v1/personals", bytePersonal)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
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

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.EditPersonal(ctx)
		req, err := http.NewRequest("PUT", "api/v1/personals", bytePersonal)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
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

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, r := gin.CreateTestContext(w)
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.EditPersonal(ctx)
		req, err := http.NewRequest("PUT", "api/v1/personals", bytePersonal)
		if err != nil {
			fmt.Println(err)
		}
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
		t.Log(w.Body.String())
	})
}

func TestPersonalController_DeletePersonal(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(service.ErrPersonalNotFound).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, _ := gin.CreateTestContext(w)
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.DeletePersonal(ctx)

		assert.Equal(t, http.StatusNotFound, w.Code)
		t.Log(w.Body.String())
	})

	t.Run("Success", func(t *testing.T) {
		var id string
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockIPersonalService(mockCtrl)
		mockService.EXPECT().Delete(id).Return(nil).AnyTimes()

		w := httptest.NewRecorder()
		gin.SetMode(gin.ReleaseMode)
		ctx, _ := gin.CreateTestContext(w)
		personalTestController := controller.NewPersonalController(mockService)
		personalTestController.DeletePersonal(ctx)

		assert.Equal(t, http.StatusAccepted, w.Code)
		t.Log(w.Body.String())
	})
}

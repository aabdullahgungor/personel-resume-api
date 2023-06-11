package test_service

import (
	"testing"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository/mocks"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDefaultPersonalService_GetAll(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIPersonalRepository(mockCtrl)
	mockRepository.EXPECT().GetAllPersonals().Return([]model.Personal{{
		Name:     "Abdullah",
		Surname:  "Gungor",
		UserName: "abdullahgungor",
		Email:    "abdullahgungor@hotmail.com.tr",
		Password: "123456",
		UserType: "admin",
	},
		{
			Name:     "Ismail",
			Surname:  "Gungor",
			UserName: "ismailgungor",
			Email:    "ismailgungor@hotmail.com",
			Password: "123456",
			UserType: "user",
		}}, nil)

	personalService := service.NewDefaultPersonalService(mockRepository)
	personals, err := personalService.GetAll()

	if assert.Nil(t, err) {
		if len(personals) == 2 {
			t.Log("Personal counts is matching, func run succesfuly")
		} else {
			t.Log("Personal counts not matching, there is a problem in func")
		}
	} else {
		t.Log(err)
	}
}
func TestDefaultPersonalService_GetById(t *testing.T) {
	id_int := 1
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIPersonalRepository(mockCtrl)
	mockRepository.EXPECT().GetPersonalById(id_int).Return(model.Personal{}, repository.ErrPersonalNotFound)

	personalService := service.NewDefaultPersonalService(mockRepository)
	id_str := "1"
	_, err := personalService.GetById(id_str)

	assert.ErrorIs(t, err, repository.ErrPersonalNotFound)
}
func TestDefaultPersonalService_Create(t *testing.T) {

	personal := model.Personal{
		Name:     "Abdullah",
		Surname:  "Gungor",
		UserName: "abdullahgungor",
		Email:    "abdullahgungor@hotmail.com.tr",
		Password: "123456",
		UserType: "admin",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIPersonalRepository(mockCtrl)
	mockRepository.EXPECT().CreatePersonal(&personal).Return(nil).Times(1)

	personalService := service.NewDefaultPersonalService(mockRepository)
	err := personalService.Create(&personal)

	if assert.Nil(t, err) {
		t.Log("Success Create Personal")
	} else {
		t.Log("Personal cannot create")
	}

}
func TestDefaultPersonalService_Edit(t *testing.T) {

	personal := model.Personal{
		Name:     "Abdullah",
		Surname:  "Gungor",
		UserName: "abdullahgungor",
		Email:    "abdullahgungor@hotmail.com.tr",
		Password: "123456",
		UserType: "admin",
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIPersonalRepository(mockCtrl)
	mockRepository.EXPECT().EditPersonal(&personal).Return(nil).Times(1)

	personalService := service.NewDefaultPersonalService(mockRepository)
	err := personalService.Edit(&personal)

	if assert.Nil(t, err) {
		t.Log("Success Update Personal")
	} else {
		t.Log("Personal cannot update")
	}
}
func TestDefaultPersonalService_Delete(t *testing.T) {

	id_int := 1
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIPersonalRepository(mockCtrl)
	mockRepository.EXPECT().DeletePersonal(id_int).Return(nil).Times(1)

	personalService := service.NewDefaultPersonalService(mockRepository)
	id_str := "1"
	err := personalService.Delete(id_str)

	if assert.Nil(t, err) {
		t.Log("Success delete Personal")
	} else {
		t.Log("Personal cannot delete")
	}
}

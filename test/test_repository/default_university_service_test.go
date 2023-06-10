package test_repository

import (
	"testing"

	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository"
	"github.com/aabdullahgungor/personal-resume-api/internal/repository/mocks"
	"github.com/aabdullahgungor/personal-resume-api/internal/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDefaultUniversityService_GetAll(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIUniversityRepository(mockCtrl)
	mockRepository.EXPECT().GetAllUniversities().Return([]model.University{{
		UniversityName: "Ataturk University",
	},
		{
			UniversityName: "Sakarya University",
		}}, nil)

	universityService := service.NewDefaultUniversityService(mockRepository)
	universities, err := universityService.GetAll()

	if assert.Nil(t, err) {
		if len(universities) == 2 {
			t.Log("University counts is matching, func run succesfuly")
		} else {
			t.Log("University counts not matching, there is a problem in func")
		}
	} else {
		t.Log(err)
	}
}

func TestDefaultUniversityService_GetById(t *testing.T) {
	id := "1"
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIUniversityRepository(mockCtrl)
	mockRepository.EXPECT().GetUniversityById(gomock.Eq(id)).Return(model.University{}, repository.ErrUniversityNotFound)

	universityService := service.NewDefaultUniversityService(mockRepository)
	_, err := universityService.GetById(id)

	assert.ErrorIs(t, err, repository.ErrUniversityNotFound)
}

func TestDefaultUniversityService_Create(t *testing.T) {

	university := model.University{UniversityName: "Ataturk University"}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIUniversityRepository(mockCtrl)
	mockRepository.EXPECT().CreateUniversity(&university).Return(nil).Times(1)

	universityService := service.NewDefaultUniversityService(mockRepository)
	err := universityService.Create(&university)

	if assert.Nil(t, err) {
		t.Log("Success Create University")
	} else {
		t.Log("University cannot create")
	}

}

func TestDefaultUniversityService_Edit(t *testing.T) {

	university := model.University{UniversityName: "Sakarya University"}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIUniversityRepository(mockCtrl)
	mockRepository.EXPECT().EditUniversity(&university).Return(nil).Times(1)

	universityService := service.NewDefaultUniversityService(mockRepository)
	err := universityService.Edit(&university)

	if assert.Nil(t, err) {
		t.Log("Success Update University")
	} else {
		t.Log("University cannot update")
	}
}

func TestDefaultUniversityService_Delete(t *testing.T) {

	id := "1"
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIUniversityRepository(mockCtrl)
	mockRepository.EXPECT().DeleteUniversity(gomock.Eq(id)).Return(nil).Times(1)

	universityService := service.NewDefaultUniversityService(mockRepository)
	err := universityService.Delete(id)

	if assert.Nil(t, err) {
		t.Log("Success delete University")
	} else {
		t.Log("University cannot delete")
	}
}

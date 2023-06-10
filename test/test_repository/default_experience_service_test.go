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

func TestDefaultExperienceService_GetAll(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIExperienceRepository(mockCtrl)
	mockRepository.EXPECT().GetAllExperiences().Return([]model.Experience{{
		CompanyName: "EMFA",
		Position:    "Sales Support Engineer",
		StartYear:   "2019-08-01T00:00:00Z",
		FinishYear:  "2020-02-14T00:00:00Z",
		PersonalID:  1,
	},
		{
			CompanyName: "GRUP ARGE",
			Position:    "System Engineer",
			StartYear:   "2020-02-14T00:00:00Z",
			FinishYear:  "2021-09-30T00:00:00Z",
			PersonalID:  1,
		}}, nil)

	experienceService := service.NewDefaultExperienceService(mockRepository)
	experiences, err := experienceService.GetAll()

	if assert.Nil(t, err) {
		if len(experiences) == 2 {
			t.Log("Experience counts is matching, func run succesfuly")
		} else {
			t.Log("Experience counts not matching, there is a problem in func")
		}
	} else {
		t.Log(err)
	}
}

func TestDefaultExperienceService_GetById(t *testing.T) {
	id := "1"
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIExperienceRepository(mockCtrl)
	mockRepository.EXPECT().GetExperienceById(gomock.Eq(id)).Return(model.Experience{}, repository.ErrExperienceNotFound)

	experienceService := service.NewDefaultExperienceService(mockRepository)
	_, err := experienceService.GetById(id)

	assert.ErrorIs(t, err, repository.ErrExperienceNotFound)
}

func TestDefaultExperienceService_Create(t *testing.T) {

	experience := model.Experience{
		CompanyName: "EMFA",
		Position:    "Sales Support Engineer",
		StartYear:   "2019-08-01T00:00:00Z",
		FinishYear:  "2020-02-14T00:00:00Z",
		PersonalID:  1,
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIExperienceRepository(mockCtrl)
	mockRepository.EXPECT().CreateExperience(&experience).Return(nil).Times(1)

	experienceService := service.NewDefaultExperienceService(mockRepository)
	err := experienceService.Create(&experience)

	if assert.Nil(t, err) {
		t.Log("Success Create Experience")
	} else {
		t.Log("Experience cannot create")
	}

}

func TestDefaultExperienceService_Edit(t *testing.T) {

	experience := model.Experience{
		CompanyName: "EMFA",
		Position:    "Sales Support Engineer",
		StartYear:   "2019-08-01T00:00:00Z",
		FinishYear:  "2020-02-14T00:00:00Z",
		PersonalID:  1,
	}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIExperienceRepository(mockCtrl)
	mockRepository.EXPECT().EditExperience(&experience).Return(nil).Times(1)

	experienceService := service.NewDefaultExperienceService(mockRepository)
	err := experienceService.Edit(&experience)

	if assert.Nil(t, err) {
		t.Log("Success Update Experience")
	} else {
		t.Log("Experience cannot update")
	}
}

func TestDefaultExperienceService_Delete(t *testing.T) {

	id := "1"
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIExperienceRepository(mockCtrl)
	mockRepository.EXPECT().DeleteExperience(gomock.Eq(id)).Return(nil).Times(1)

	experienceService := service.NewDefaultExperienceService(mockRepository)
	err := experienceService.Delete(id)

	if assert.Nil(t, err) {
		t.Log("Success delete Experience")
	} else {
		t.Log("Experience cannot delete")
	}
}

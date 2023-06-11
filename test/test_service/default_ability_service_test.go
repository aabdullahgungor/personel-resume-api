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

func TestDefaultAbilityService_GetAll(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIAbilityRepository(mockCtrl)
	mockRepository.EXPECT().GetAllAbilities().Return([]model.Ability{{
		AbilityName: "Go",
	},
		{
			AbilityName: "Ruby",
		}}, nil)

	abilityService := service.NewDefaultAbilityService(mockRepository)
	abilities, err := abilityService.GetAll()

	if assert.Nil(t, err) {
		if len(abilities) == 2 {
			t.Log("Ability counts is matching, func run succesfuly")
		} else {
			t.Log("Ability counts not matching, there is a problem in func")
		}
	} else {
		t.Log(err)
	}
}

func TestDefaultAbilityService_GetById(t *testing.T) {
	id_int := 1
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIAbilityRepository(mockCtrl)
	mockRepository.EXPECT().GetAbilityById(id_int).Return(model.Ability{}, repository.ErrAbilityNotFound)

	abilityService := service.NewDefaultAbilityService(mockRepository)
	id_str := "1"
	_, err := abilityService.GetById(id_str)

	assert.ErrorIs(t, err, repository.ErrAbilityNotFound)
}

func TestDefaultProductService_Create(t *testing.T) {

	ability := model.Ability{AbilityName: "Go"}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIAbilityRepository(mockCtrl)
	mockRepository.EXPECT().CreateAbility(&ability).Return(nil).Times(1)

	abilityService := service.NewDefaultAbilityService(mockRepository)
	err := abilityService.Create(&ability)

	if assert.Nil(t, err) {
		t.Log("Success Create Product")
	} else {
		t.Log("Product cannot create")
	}

}

func TestDefaultProductService_Edit(t *testing.T) {

	ability := model.Ability{AbilityName: "Go"}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIAbilityRepository(mockCtrl)
	mockRepository.EXPECT().EditAbility(&ability).Return(nil).Times(1)

	abilityService := service.NewDefaultAbilityService(mockRepository)
	err := abilityService.Edit(&ability)

	if assert.Nil(t, err) {
		t.Log("Success Update Ability")
	} else {
		t.Log("Ability cannot update")
	}
}

func TestDefaultProductService_Delete(t *testing.T) {

	id_int := 1
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockRepository := mocks.NewMockIAbilityRepository(mockCtrl)
	mockRepository.EXPECT().DeleteAbility(id_int).Return(nil).Times(1)

	abilityService := service.NewDefaultAbilityService(mockRepository)
	id_str := "1"
	err := abilityService.Delete(id_str)

	if assert.Nil(t, err) {
		t.Log("Success delete Ability")
	} else {
		t.Log("Ability cannot delete")
	}
}

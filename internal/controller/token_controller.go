package controller

import (
	"net/http"

	"github.com/aabdullahgungor/personal-resume-api/internal/auth"
	"github.com/aabdullahgungor/personal-resume-api/internal/model"
	"github.com/aabdullahgungor/personal-resume-api/internal/service/interfaces"
	"github.com/gin-gonic/gin"
)

type tokenController struct {
	service interfaces.IPersonalService
}

func NewTokenController(ts interfaces.IPersonalService) *tokenController {
	return &tokenController{
		service: ts,
	}
}

func (ts *tokenController) GenerateToken(context *gin.Context) {
	var request model.Authentication
	//var personal model.Personal
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if email exists and password is correct
	personal, err := ts.service.GetByEmail(request.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		context.Abort()
		return
	}
	credentialError := personal.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(personal.Email, personal.UserName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}

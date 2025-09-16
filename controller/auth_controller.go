package controller

import (
	"github.com/gin-gonic/gin"
	"loginRegister/model"
	"loginRegister/repository"
	"loginRegister/utils"
	"net/http"
)

func Register(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := repository.GetUserByUsername(user.Username); err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
		return
	}

	if _, err := repository.FindUserByPhone(user.Phone); err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "手机号已存在"})
		return
	}

	if _, err := repository.FindUserByEmail(user.Email); err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "邮箱已存在"})
		return
	}

	hashedPwd, err := utils.HashPassword(user.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = hashedPwd

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := repository.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user *model.User
	var err error
	if user, err = repository.GetUserByUsername(input.Username); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	if !utils.CheckPassword(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	token, err := utils.GenerateJWT(user.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

package handlers

import (
	"LK_back/pkg/db"
	"LK_back/pkg/models"
	"LK_back/pkg/schemas"
	"LK_back/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func CreateUserHandler(c *gin.Context) {
	var newUser schemas.SigUpUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if newUser.Password != newUser.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пароли не совпадают"})
		return
	}
	saveUser := models.User{
		ID:           uuid.New().String(),
		NickName:     newUser.NickName,
		HashPassword: newUser.Password,
		FirstName:    newUser.FirstName,
		LastName:     newUser.LastName,
		OtherName:    newUser.OtherName,
		Email:        newUser.Email,
		PhoneNumber:  newUser.PhoneNumber,
	}
	if err := db.DB().Save(&saveUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь уже существует"})
		return
	}
	service.SendSmsService(newUser.PhoneNumber)
	c.JSON(http.StatusCreated, gin.H{"message": saveUser.ID})
}

func Authorization(c *gin.Context) {
	var userLogin schemas.LoginUser
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Некорректный логин или пароль"})
	}

	var user models.User
	if err := db.DB().Where(models.User{NickName: userLogin.NickName}).First(&user).Error; err != nil {
		log.Fatal(err)
		return
	}
	if userLogin.Password != user.HashPassword {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Incorrect password"})
	}
	c.JSON(http.StatusOK, gin.H{"data": user.ID})
}

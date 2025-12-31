package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liubaicai/navi-home/models"
)

func getCurrentUser(c *gin.Context) *models.User {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID == nil {
		return nil
	}
	user, err := models.GetUserByID(userID.(uint))
	if err != nil {
		return nil
	}
	return user
}

func setCurrentUser(c *gin.Context, user *models.User) {
	session := sessions.Default(c)
	if user != nil {
		session.Set("user_id", user.ID)
	} else {
		session.Delete("user_id")
	}
	session.Save()
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liubaicai/navi-home/models"
)

func SignInPage(c *gin.Context) {
	user := getCurrentUser(c)
	if user != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.HTML(http.StatusOK, "sign_in.html", gin.H{
		"Alert": "",
	})
}

func SignIn(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	user, err := models.GetUserByEmail(email)
	if err != nil || !user.CheckPassword(password) {
		c.HTML(http.StatusOK, "sign_in.html", gin.H{
			"Alert": "邮箱或密码错误",
		})
		return
	}

	setCurrentUser(c, user)
	c.Redirect(http.StatusFound, "/")
}

func SignUpPage(c *gin.Context) {
	user := getCurrentUser(c)
	if user != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	c.HTML(http.StatusOK, "sign_up.html", gin.H{
		"Errors":                nil,
		"MinimumPasswordLength": 6,
	})
}

func SignUp(c *gin.Context) {
	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	passwordConfirmation := c.PostForm("password_confirmation")

	var errors []string

	if email == "" {
		errors = append(errors, "邮箱不能为空")
	}
	if username == "" {
		errors = append(errors, "用户名不能为空")
	}
	if password == "" {
		errors = append(errors, "密码不能为空")
	}
	if len(password) < 6 {
		errors = append(errors, "密码至少需要6个字符")
	}
	if password != passwordConfirmation {
		errors = append(errors, "密码确认不匹配")
	}

	// Check if email already exists
	existingUser, _ := models.GetUserByEmail(email)
	if existingUser != nil {
		errors = append(errors, "邮箱已被注册")
	}

	if len(errors) > 0 {
		c.HTML(http.StatusOK, "sign_up.html", gin.H{
			"Errors":                errors,
			"MinimumPasswordLength": 6,
		})
		return
	}

	user, err := models.CreateUser(email, username, password)
	if err != nil {
		c.HTML(http.StatusOK, "sign_up.html", gin.H{
			"Errors":                []string{"注册失败，请重试"},
			"MinimumPasswordLength": 6,
		})
		return
	}

	setCurrentUser(c, user)
	c.Redirect(http.StatusFound, "/")
}

func SignOut(c *gin.Context) {
	setCurrentUser(c, nil)
	c.Redirect(http.StatusFound, "/")
}

func EditUserPage(c *gin.Context) {
	user := getCurrentUser(c)
	if user == nil {
		c.Redirect(http.StatusFound, "/users/sign_in")
		return
	}

	c.HTML(http.StatusOK, "edit_user.html", gin.H{
		"User":   user,
		"Errors": nil,
	})
}

func EditUser(c *gin.Context) {
	user := getCurrentUser(c)
	if user == nil {
		c.Redirect(http.StatusFound, "/users/sign_in")
		return
	}

	email := c.PostForm("email")
	username := c.PostForm("username")
	password := c.PostForm("password")
	passwordConfirmation := c.PostForm("password_confirmation")
	currentPassword := c.PostForm("current_password")

	var errors []string

	if email == "" {
		errors = append(errors, "邮箱不能为空")
	}
	if username == "" {
		errors = append(errors, "用户名不能为空")
	}
	if password != "" && len(password) < 6 {
		errors = append(errors, "密码至少需要6个字符")
	}
	if password != passwordConfirmation {
		errors = append(errors, "密码确认不匹配")
	}
	if currentPassword == "" {
		errors = append(errors, "请输入当前密码")
	}

	// Check if email already exists (for different user)
	existingUser, _ := models.GetUserByEmail(email)
	if existingUser != nil && existingUser.ID != user.ID {
		errors = append(errors, "邮箱已被其他用户使用")
	}

	if len(errors) > 0 {
		c.HTML(http.StatusOK, "edit_user.html", gin.H{
			"User":   user,
			"Errors": errors,
		})
		return
	}

	err := user.Update(email, username, password, currentPassword)
	if err != nil {
		c.HTML(http.StatusOK, "edit_user.html", gin.H{
			"User":   user,
			"Errors": []string{"当前密码错误"},
		})
		return
	}

	c.Redirect(http.StatusFound, "/")
}

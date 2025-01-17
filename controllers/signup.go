package controllers

import (
	"log"
	"net/http"
	"regexp"
	"shvx/models"

	"github.com/gin-gonic/gin"
)

func UserSignUP(c *gin.Context) {
	var userInfo models.UserSignUPInfo
	userName := c.PostForm("user_name")
	if len(userName) == 0 {
		log.Println("Please fill the mendetory field (User name can't be empty)")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "name can not be blank",
			"status":  "failed",
		})
		return
	}
	userInfo.Name = userName
	email := c.PostForm("email")
	if len(email) == 0 {
		log.Println("Please fill the mendetory field (User email can't be empty)")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email can not be blank",
			"status":  "failed",
		})
		return
	}
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// Compile the regular expression pattern.
	re, err := regexp.Compile(emailPattern)
	// If there is an error compiling the regex pattern, log the error and return false and the error.
	if err != nil {
		log.Println("ValidateEmailPattern: Failed while compile regular expression for email pattern with error: ", err)
		return
	}
	// Check if the email address matches the compiled regex pattern.
	if re.MatchString(email) {
		// If the email matches the pattern, log the success and return true with no error.
		log.Println("ValidateEmailPattern: validating email successfull.")
	}
	userInfo.Email = email
	phone := c.PostForm("phone")
	if len(phone) == 0 {
		log.Println("Please fill the mendetory field (User phone can't be empty)")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "phone can not be blank",
			"status":  "failed",
		})
		return
	}
	matchphone, err := regexp.MatchString("^[+]?[0-9]{10,15}$", phone)
	if !matchphone || err != nil {
		log.Println("invalid email please,insert valid email")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "innvalid phone no. can not be acceptable",
			"status":  "failed",
		})
		return
	}
	userInfo.PhoneNumber = phone
	password := c.PostForm("password")
	if len(password) < 8 && len(password) > 15 {
		log.Println("Please fill the mendetory field (User password can't be empty)")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password can not be blank",
			"status":  "failed",
		})
		return
	}
	// Check for at least one letter (a-z or A-Z)
	hasLetter, _ := regexp.MatchString(`[a-zA-Z]`, password)
	if !hasLetter {
		log.Println("must contain latter")
		return
	}
	// Check for at least one digit (0-9)
	hasDigit, _ := regexp.MatchString(`\d`, password)
	if !hasDigit {
		log.Println("must contain digits")
		return
	}
	// Check for at least one special character (@, $, !, %, *, ?, &, #)
	hasSpecial, _ := regexp.MatchString(`[@$!%*?&#]`, password)
	if !hasSpecial {
		log.Println("must contain special character.")
		return
	}
	userInfo.Password = password
	confirmPassword := c.PostForm("confirm_password")
	if len(confirmPassword) == 0 {
		log.Println("Please fill the mendetory field (User password can't be empty)")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "confirmation password can not be blank",
			"status":  "failed",
		})
		return
	}
	if password != confirmPassword {
		log.Println("Please fill the correct password on confirm password field")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password miss match!!!",
			"status":  "failed",
		})
		return
	}
	err = models.UserSignUP(userInfo)
	if err != nil {
		log.Println("Failed to save the information of user while signUp with error :", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to save user details i database.",
			"status":  "failed",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "welcome in SHVX.",
		"status":  "success",
	})
}
func UserSignIn(c *gin.Context) {

	email := c.PostForm("email")
	if len(email) == 0 {
		log.Println("Please fill the mendetory field (User email can't be empty)")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email can not be blank",
			"status":  "failed",
		})
		return
	}
	match, err := regexp.MatchString("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+.[a-zA-Z]{2,}$", email)
	if !match || err != nil {
		log.Println("invalid email please,insert valid email")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid email can not be acceptable",
			"status":  "failed",
		})
		return
	}
	userPassword := c.PostForm("password")
	if len(userPassword) == 0 {
		log.Println("Please fill the mendetory field (User password can't be empty)")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password can not be blank",
			"status":  "failed",
		})
		return
	}
	prevPassword, userName, err := models.GetUserPassword(email)
	if err != nil {
		log.Println("Failed to authenticate the information of user while signIn with error :", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to authenticate user details from database.",
			"status":  "failed",
			"error":   err.Error(),
		})
		return
	}
	if prevPassword == userPassword {
		log.Printf("Hi, %s welcome back in SHVX, you sighIn successfully!!", userName)
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome in SHVX.",
			"status":  "success",
		})
	} else {
		log.Println("I dont have your record in my database, please register first.")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "please register first.",
			"status":  "failed",
		})
	}
}

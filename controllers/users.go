package controllers

import (
	"net/http"
	// "time"
	// "time"
	// "log"

	"github.com/haqim007/dairy_v0.1/helpers"
	"github.com/haqim007/dairy_v0.1/structs"

	"github.com/gin-gonic/gin"
)

//CreateUser to create new data to database
func (idb *InDB) CreateUser(c *gin.Context) {
	var (
			user structs.User
			postUser structs.PostUser
			result gin.H
		)
	c.Bind(&postUser)
	httpstatus := http.StatusBadRequest
	var allIsValid bool = true

	if postUser.Fullname == ""{
		allIsValid = false
		result = gin.H{
			"result": nil,
			"message":"Fullname is required!",
		}
	}
	if postUser.Birthday == ""{
		allIsValid = false
		result = gin.H{
			"result": nil,
			"message":"Birthday is required!",
		}
	}
	if postUser.Email == ""{
		allIsValid = false
		result = gin.H{
			"result": nil,
			"message":"Email is required!",
		}
	}

	if postUser.Username == ""{
		result = gin.H{
			"result": nil,
			"message":"Username is required!",
		}
	}

	if postUser.Password == ""{
		allIsValid = false
		result = gin.H{
			"result": nil,
			"message":"Password is required!",
		}
	}else if helpers.PasswordValidation(postUser.Password) == false {
		allIsValid = false
		result = gin.H{
			"result": nil,
			"message":"Password minimum eight characters, at least one letter, one number and one special character!",
		}
	}

	if allIsValid == true {
		user.Fullname = postUser.Fullname
		user.Birthday = helpers.StringToDate(postUser.Birthday)
		user.Email = postUser.Email
		user.Username = postUser.Username
		user.Password = postUser.Password
		idb.DB.Create(&user)
		result = gin.H{
			"result": user,
			"message": "success",
		}
		httpstatus = http.StatusOK	
	}
	c.JSON(httpstatus, result)
}
//GetUser a
func (idb *InDB) GetUser(c *gin.Context){
	var (
		user structs.User
		login structs.Login
		// userSession structs.UserSession
		result gin.H
	)
	httpstatus := http.StatusBadRequest
	c.Bind(&login)
	allIsValid := true

	if login.Username == ""{
		result = gin.H{
			"access_token": nil,
			"refresh_token":nil,
			"message":"Username is required!",
		}
	}

	if login.Password == ""{
		allIsValid = false
		result = gin.H{
			"access_token": nil,
			"refresh_token":nil,
			"message":"Password is required!",
		}
	}
	
	if allIsValid {
		err := idb.DB.Where("(username = ? OR email = ?) AND password = ?", login.Username,login.Username, login.Password).First(&user).Error
		if err != nil {
			httpstatus = http.StatusUnauthorized
			result = gin.H{
				"message": "User not found",
				"access_token": nil,
				"refresh_token":nil,
				
			}
		} else {
			token, err := helpers.CreateToken(user.ID)
			if err != nil {
				c.JSON(http.StatusUnprocessableEntity, err.Error())
				return
			}

			// idb.DB.Where("userid = ?", user.ID).Delete(&userSession)

			// userSession.UserID = user.ID
			// userSession.Token = token
			// userSession.IssuedAt = time.Now()

			// idb.DB.Create(&userSession)
			
			httpstatus = http.StatusOK
			result = gin.H{
				"message": "success",
				"access_token": token["access_token"],
				"refresh_token": token["refresh_token"],
			}
		}
	}

	c.JSON(httpstatus, result)
}
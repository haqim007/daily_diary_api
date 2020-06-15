package controllers

import (
	"fmt"
	"net/http"

	// "time"
	// "log"

	"github.com/haqim007/dairy_v0.1/helpers"
	"github.com/haqim007/dairy_v0.1/structs"

	"github.com/gin-gonic/gin"
)

//CreateDiary a
func (idb *InDB) CreateDiary(c *gin.Context) {
	var (
		diary     structs.Diary
		postDiary structs.PostDiary
		result    gin.H
	)
	c.Bind(&postDiary)

	httpstatus := http.StatusBadRequest

	var allIsValid bool = true

	if postDiary.Date == "" {
		allIsValid = false
		result = gin.H{
			"result":  nil,
			"message": "Date is required!",
		}
	}
	if postDiary.Content == "" && allIsValid {
		allIsValid = false
		result = gin.H{
			"result":  nil,
			"message": "Content is required!",
		}
	}

	//validate token
	claims, ok, err := helpers.ValidateToken(postDiary.AccessToken)

	if !ok && allIsValid {
		fmt.Println(err)
		allIsValid = false
		httpstatus = http.StatusUnauthorized
		result = gin.H{
			"result":  nil,
			"message": err.Error(),
		}
	}

	if allIsValid {

		UserID := claims["user_id"].(float64)
		diary.UserID = uint(UserID)
		diary.Date = helpers.StringToDate(postDiary.Date)
		err := idb.DB.Where("user_id = ? AND date = ?", UserID, diary.Date).First(&diary).Error

		diary.Content = postDiary.Content
		diary.Date = helpers.StringToDate(postDiary.Date)

		if err != nil {
			fmt.Println(err)
			// 	idb.DB.Create(&diary)
		}
		// else {
		idb.DB.Save(&diary)
		result = gin.H{
			"result":  diary,
			"message": "success",
		}
		// }

		httpstatus = http.StatusOK
	}
	c.JSON(httpstatus, result)
}

//GetDiaries a
func (idb *InDB) GetDiaries(c *gin.Context) {

	type TokenPost struct {
		AccessToken string `json:"access_token" form:"access_token" binding:"required"`
	}

	var (
		diary     []structs.Diary
		tokenPost TokenPost
		count     int
		result    gin.H
	)
	c.Bind(&tokenPost)

	httpstatus := http.StatusBadRequest
	allIsValid := true

	//validate token
	_, ok, err := helpers.ValidateToken(tokenPost.AccessToken)

	if !ok && allIsValid {
		fmt.Println(err)
		allIsValid = false
		httpstatus = http.StatusUnauthorized
		result = gin.H{
			"count":   0,
			"result":  nil,
			"message": err.Error(),
		}
	}

	req := c.Request.URL.Query()
	if _, okQuerter := req["quarter"]; !okQuerter && allIsValid {
		//do something here
		allIsValid = false
		result = gin.H{
			"result":  nil,
			"message": "Quarter params required",
		}
	}
	if _, okYear := req["year"]; !okYear && allIsValid {
		//do something here
		allIsValid = false
		result = gin.H{
			"result":  nil,
			"message": "Year params required",
		}
	}
	if allIsValid {
		quarter := req["quarter"][0]
		year := req["year"][0]
		result = gin.H{
			"result":  req["quarter"][0],
			"message": "success",
		}
		months := make(map[string]int)
		if quarter == "1" {
			months["first"] = 1
			months["last"] = 3
		} else if quarter == "2" {
			months["first"] = 4
			months["last"] = 6
		} else if quarter == "3" {
			months["first"] = 7
			months["last"] = 9
		} else if quarter == "4" {
			months["first"] = 10
			months["last"] = 12
		}
		httpstatus = http.StatusOK
		err := idb.DB.Where("(extract(month from date) BETWEEN ? AND ?) AND extract(year from date) = ?", months["first"], months["last"], year).Find(&diary).Count(&count).Error
		if err != nil {
			httpstatus = http.StatusOK
			result = gin.H{
				"count":   0,
				"result":  nil,
				"message": err.Error(),
			}
		} else {

			httpstatus = http.StatusOK
			result = gin.H{
				"count":   count,
				"result":  diary,
				"message": "success",
			}
		}

	}

	c.JSON(httpstatus, result)
}

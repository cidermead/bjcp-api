package controller

import (
	"fmt"
	// "strconv"
	// "strings"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
	"github.com/cidermead/bjcp-api/include"
	"github.com/cidermead/bjcp-api/model"
)

// var db *gorm.DB
// var err error

// Post struct alias
type Question = model.Question

func GetQuestion(c *gin.Context) {
	db = include.GetDB()
	id := c.Params.ByName("id")

	var question Question

	if err := db.Where("id = ? ", id).First(&question).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		db.Model(&question)
		c.JSON(200, question)
	}
}

func GetRandom(c *gin.Context) {
	var question Question

	db = include.GetDB()
	exam := c.Params.ByName("exam")
	topic := c.Params.ByName("topic")

	query := &Question{
		Active: true,
		Deleted: false,
	}

	if exam != "" {
		query.Exam = exam
	}

	if topic != "" {
		query.Topic = topic
	}

	if err := db.Where(query).Order("random()").First(&question).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		db.Model(&question)
		c.JSON(200, question)
	}
}



/*
func GetQuestions(c *gin.Context) {
	db = include.GetDB()
	var posts []Post
	var data Data
	var count int64

	//Get name from query
	name := c.DefaultQuery("name", "")

	//Get description from query
	description := c.DefaultQuery("description", "")

	// Order By filtering option add
	Sort := c.DefaultQuery("order", "id|desc")
	SortArray := strings.Split(Sort, "|")

	// Define and get offset for pagination
	offset := c.Query("offset")
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}

	// Define and get limit for pagination
	limit := c.Query("limit")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 25
	}

	query := db.Limit(limitInt)
	query = query.Offset(offsetInt)
	query = query.Order(SortArray[0] + " " + SortArray[1])

	// In postgres you shoud use ILIKE to make search case insensitive
	if "name" != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if "description" != "" {
		query = query.Where("description LIKE ?", "%"+description+"%")
	}

	if err := query.Find(&posts).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {

		// We are resetting offset to 0 to return total number.
		// This is a fix for Gorm offset issue
		offsetInt = 0
		query = query.Offset(offsetInt)
		query.Table("posts").Count(&count)

		data.Total = count
		data.Data = posts

		c.JSON(200, data)
	}
}
*/
/*
func CreatePost(c *gin.Context) {
	db = include.GetDB()
	var post Post

	c.BindJSON(&post)

	if err := db.Create(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, post)
	}
}

func UpdatePost(c *gin.Context) {
	db = include.GetDB()
	var post Post
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	c.BindJSON(&post)

	db.Save(&post)
	c.JSON(200, post)
}

func DeletePost(c *gin.Context) {
	db = include.GetDB()
	id := c.Params.ByName("id")
	var post Post

	if err := db.Where("id = ? ", id).Delete(&post).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{"id#" + id: "deleted"})
	}
}
*/

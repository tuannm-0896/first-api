package Controllers

import (
	"first-api/Models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetCourses ... Get all Courses
func GetCourses(c *gin.Context) {
	var Course []Models.Course
	err := Models.GetAllCourses(&Course)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Course)
	}
}

//CreateCourse ... Create Course
func CreateCourse(c *gin.Context) {
	var Course Models.Course
	c.BindJSON(&Course)
	err := Models.CreateCourse(&Course)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Course)
	}
}

//GetCourseByID ... Get the Course by id
func GetCourseByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var Course Models.Course
	err := Models.GetCourseByID(&Course, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Course)
	}
}

//UpdateCourse ... Update the Course information
func UpdateCourse(c *gin.Context) {
	var Course Models.Course
	id := c.Params.ByName("id")
	err := Models.GetCourseByID(&Course, id)
	if err != nil {
		c.JSON(http.StatusNotFound, Course)
	}
	c.BindJSON(&Course)
	err = Models.UpdateCourse(&Course, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, Course)
	}
}

//DeleteCourse ... Delete the Course
func DeleteCourse(c *gin.Context) {
	var Course Models.Course
	id := c.Params.ByName("id")
	err := Models.DeleteCourse(&Course, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

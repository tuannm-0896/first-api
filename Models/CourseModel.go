package Models

import (
	"first-api/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Course struct {
	Id             uint   `json:"id"`
	Title          string `json:"title"`
	CourseContents []CourseContent
}

func (b *Course) TableName() string {
	return "course"
}

//GetAllCourses Fetch all Course data
func GetAllCourses(course *[]Course) (err error) {
	if err = Config.DB.Find(course).Error; err != nil {
		return err
	}
	return nil
}

//CreateCourse ... Insert New data
func CreateCourse(course *Course) (err error) {
	if err = Config.DB.Create(course).Error; err != nil {
		return err
	}
	return nil
}

//GetCourseByID ... Fetch only one Course by Id
func GetCourseByID(course *Course, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(course).Error; err != nil {
		return err
	}
	return nil
}

//UpdateCourse ... Update Course
func UpdateCourse(course *Course, id string) (err error) {
	fmt.Println(course)
	Config.DB.Save(course)
	return nil
}

//DeleteCourse ... Delete Course
func DeleteCourse(course *Course, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(course)
	return nil
}

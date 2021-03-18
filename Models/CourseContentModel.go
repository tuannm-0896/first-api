package Models

import (
	"first-api/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type CourseContent struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CourseID    uint   `json:"course_id"`
}

func (b *CourseContent) TableName() string {
	return "course_content"
}

func GetAllCoursesContent(course_content *[]CourseContent) (err error) {
	fmt.Println("test fmt")
	if err = Config.DB.Find(course_content).Error; err != nil {
		return err
	}
	return nil
}

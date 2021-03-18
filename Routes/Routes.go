package Routes

import (
	"first-api/Controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/first-api")
	{
		grp1.GET("courses", Controllers.GetCourses)
		grp1.POST("course", Controllers.CreateCourse)
		grp1.GET("course/:id", Controllers.GetCourseByID)
		grp1.PUT("course/:id", Controllers.UpdateCourse)
		grp1.DELETE("course/:id", Controllers.DeleteCourse)
	}
	return r
}

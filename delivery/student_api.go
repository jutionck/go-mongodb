package delivery

import (
	"github.com/gin-gonic/gin"
	"go-mongodb/db"
	"go-mongodb/model"
	"go-mongodb/repository"
	"go-mongodb/usecase"
	"net/http"
	"strconv"
)

type StudentApi struct {
	router  *gin.RouterGroup
	useCase usecase.IStudentUseCase
}

func NewStudentApi(router *gin.RouterGroup, resource *db.Resource) *StudentApi {
	userRoute := router.Group("/students")
	studentRepo := repository.NewStudentRepository(resource)
	studentApi := StudentApi{
		router:  userRoute,
		useCase: usecase.NewStudentUseCase(studentRepo),
	}
	studentApi.initRouter()
	return &studentApi

}

func (api *StudentApi) initRouter() {
	api.router.GET("", api.getAllStudentWithPagination)
	api.router.GET("/:name", api.getStudentByName)
	api.router.POST("", api.createStudent)
}

func (api *StudentApi) createStudent(c *gin.Context) {
	var student model.Student
	err := c.BindJSON(&student)
	if err != nil {
		return
	}
	registeredStudent, err := api.useCase.NewRegistration(student)
	c.JSON(http.StatusOK, gin.H{
		"message": registeredStudent,
	})
}

func (api *StudentApi) getStudentByName(c *gin.Context) {
	name := c.Param("name")
	student, err := api.useCase.FindStudentInfoByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": student,
	})
}

func (api *StudentApi) getAllStudentWithPagination(c *gin.Context) {
	skip := c.Query("skip")
	limit := c.Query("limit")
	convertSkip, _ := strconv.Atoi(skip)
	convertLimit, _ := strconv.Atoi(limit)
	student, err := api.useCase.FindAllStudentWithPagination(int64(convertSkip), int64(convertLimit))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": student,
	})
}

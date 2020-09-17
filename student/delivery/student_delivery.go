package delivery

import (
	"github.com/YuukiRen/esGolang/domain"
	"github.com/labstack/echo"
	"net/http"
)

type StudentHandler struct {
	studentUsecase domain.StudentUsecase
}

func NewStudentDelivery(e *echo.Echo, usecase domain.StudentUsecase){
	handler := &StudentHandler{
		usecase,
	}
	e.GET("/students", handler.GetAllStudents)
	e.POST("/student", handler.CreateStudent)
	e.GET("/student/:nim", handler.GetStudentByNIM)
	e.DELETE("/student/:nim", handler.DeleteStudent)
}

func (s *StudentHandler) GetAllStudents(c echo.Context) error{
	ctx := c.Request().Context()

	students, errResp := s.studentUsecase.GetAll(ctx)
	if errResp != nil {
		domain.MakeLogEntry(c).Error(errResp.Message)
		return c.JSON(errResp.Code, errResp)
	}
	domain.MakeLogEntry(c).Info("request done.")
	return c.JSON(http.StatusOK, domain.JSONResponse{
		Data:    students,
		Message: "Search success",
	})
}
func (s *StudentHandler) GetStudentByNIM(c echo.Context) error {
	nimParam:= c.Param("nim")
	ctx := c.Request().Context()
	student, errResp := s.studentUsecase.GetOneByNIM(ctx, nimParam)
	if errResp != nil {
		domain.MakeLogEntry(c).Error(errResp.Message)
		return c.JSON(errResp.Code, errResp)
	}
	domain.MakeLogEntry(c).Info("request done.")
	return c.JSON(http.StatusOK, domain.JSONResponse{
		Data:    student,
		Message: "Search success",
	})
}

func (s *StudentHandler) CreateStudent(c echo.Context) error {
	student := domain.Student{}
	err := c.Bind(&student)
	if err != nil {
		domain.MakeLogEntry(c).Error(err)
		return c.JSON(http.StatusUnprocessableEntity, domain.ErrorResponse{
			Message: "Unable to process data",
		})
	}
	if ok, err := domain.IsRequestValid(student); !ok {
		domain.MakeLogEntry(c).Error(err)
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid input",
		})
	}
	ctx := c.Request().Context()
	student, errResp := s.studentUsecase.Create(ctx, student)
	if errResp != nil{
		domain.MakeLogEntry(c).Error(errResp.Message)
		return c.JSON(errResp.Code, errResp)
	}
	domain.MakeLogEntry(c).Info("request done.")
	return c.JSON(http.StatusOK, domain.JSONResponse{
		Data:    student,
		Message: "Student created successfully",
	})
}
func (s *StudentHandler) DeleteStudent(c echo.Context) error {
	nimParam:= c.Param("nim")
	ctx := c.Request().Context()
	student, errResp := s.studentUsecase.Delete(ctx, nimParam)
	if errResp != nil {
		domain.MakeLogEntry(c).Error(errResp.Message)
		return c.JSON(errResp.Code, errResp)
	}
	domain.MakeLogEntry(c).Info("request done.")
	return c.JSON(http.StatusOK, domain.JSONResponse{
		Data:    student,
		Message: "Delete success",
	})
}

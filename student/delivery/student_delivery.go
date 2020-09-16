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
}

func (s *StudentHandler) GetAllStudents(c echo.Context) error{
	ctx := c.Request().Context()

	students, err := s.studentUsecase.GetAll(ctx)
	if err != nil {
		return c.JSON(err.Code, err.Message)
	}
	return c.JSON(http.StatusOK, domain.JSONResponse{
		Data:    students,
		Message: "Search success",
	})
}
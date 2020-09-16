package usecase

import (
	"context"
	"fmt"
	"github.com/YuukiRen/esGolang/domain"
	"time"
)

type StudentUsecase struct {
	studentRepository domain.StudentRepository
	ctxTimeout time.Duration
}
func NewStudentUsecase(repository domain.StudentRepository, timeout time.Duration) domain.StudentUsecase{
	return &StudentUsecase{
		repository,
		timeout,
	}
}

func (s *StudentUsecase) GetAll(c context.Context) ([]domain.Student, *domain.ErrorResponse) {
	ctx, cancel := context.WithTimeout(c, s.ctxTimeout)
	defer cancel()
	searchResult, err := s.studentRepository.GetAll(ctx)
	if err != nil {
		errMessage := fmt.Sprintf("Fail to get all data in students: %s",
			err.Error())
		return nil, &domain.ErrorResponse{
			500,
			errMessage,
		}
	}
	students, err := ConvertElasticResultToStudents(searchResult)
	if err != nil {
		errMessage := fmt.Sprintf("Fail to convert elastic result to object: %s",
			err.Error())
		return nil, &domain.ErrorResponse{
			500,
			errMessage,
		}
	}
	return students, nil
}



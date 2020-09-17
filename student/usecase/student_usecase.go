package usecase

import (
	"context"
	"fmt"
	"github.com/YuukiRen/esGolang/domain"
	"net/http"
	"time"
)

type StudentUsecase struct {
	studentRepository domain.StudentRepository
	ctxTimeout time.Duration
}

func (s *StudentUsecase) Delete(c context.Context, nim string) (domain.Student, *domain.ErrorResponse) {
	ctx, cancel := context.WithTimeout(c, s.ctxTimeout)
	defer cancel()

	searchResult, err := s.studentRepository.GetOne(ctx, "nim", nim)
	if err != nil{
		errMessage := fmt.Sprintf("[Usecase][DeleteCheck]Failed to check student existence in database: %s",
			err.Error())
		return domain.Student{}, &domain.ErrorResponse{
			http.StatusInternalServerError,
			errMessage,
		}
	}

	students, err := ConvertElasticResultToStudents(searchResult)
	if err != nil{
		errMessage := fmt.Sprintf("[Usecase][ElasticConversion]Fail to convert elastic result to object: %s",
			err.Error())
		return domain.Student{}, &domain.ErrorResponse{
			http.StatusInternalServerError,
			errMessage,
		}
	}
	if len(students) != 1{
		errMessage := fmt.Sprintf("No Student found")
		return domain.Student{}, &domain.ErrorResponse{
			http.StatusBadRequest,
			errMessage,
		}
	}
	err = s.studentRepository.Delete(ctx, "nim", nim)
	if err != nil {
		errMessage := fmt.Sprintf("[Usecase][DeleteStudent]Fail to delete student: %s",
			err.Error())
		return domain.Student{}, &domain.ErrorResponse{
			http.StatusInternalServerError,
			errMessage,
		}
	}
	return students[0], nil
}

func (s *StudentUsecase) GetOneByNIM(c context.Context, nim string) (domain.Student, *domain.ErrorResponse) {
	ctx, cancel := context.WithTimeout(c, s.ctxTimeout)
	defer cancel()

	searchResult, err := s.studentRepository.GetOne(ctx, "nim", nim)
	if err != nil {
		errMessage := fmt.Sprintf("[Usecase][GetOneByNIM]Fail to get student : %s",
			err.Error())
		return domain.Student{}, &domain.ErrorResponse{
			http.StatusInternalServerError,
			errMessage,
		}
	}

	students, err := ConvertElasticResultToStudents(searchResult)
	if err != nil {
		errMessage := fmt.Sprintf("[Usecase][ElasticConversion]Fail to convert elastic result to object: %s",
			err.Error())
		return domain.Student{}, &domain.ErrorResponse{
			http.StatusInternalServerError,
			errMessage,
		}
	}

	if len(students) <= 0{
		return domain.Student{}, &domain.ErrorResponse{
			http.StatusOK,
			"No student found",
		}
	}
	return students[0], nil
}

func (s *StudentUsecase) Create(c context.Context, student domain.Student) (domain.Student, *domain.ErrorResponse) {
	ctx, cancel := context.WithTimeout(c, s.ctxTimeout)
	defer cancel()
	searchResult, err := s.studentRepository.GetOne(ctx, "nim", student.NIM)
	if err != nil{
		errMessage := fmt.Sprintf("[Usecase][CreateCheck]Failed to check student existence in database: %s",
			err.Error())
		return domain.Student{}, &domain.ErrorResponse{
			http.StatusInternalServerError,
			errMessage,
		}
	}
	if found := searchResult.Hits.TotalHits.Value; found > 0{
		errMessage := fmt.Sprintf("Student with nim: %s existed.",
			student.NIM)
		return domain.Student{}, &domain.ErrorResponse{
			http.StatusConflict,
			errMessage,
		}
	}
	err = s.studentRepository.Create(ctx, student)
	if err != nil{
		errMessage := fmt.Sprintf("[Usecase][CreateStudent]Failed to insert student to database: %s",
			err.Error())
		return domain.Student{}, &domain.ErrorResponse{
			http.StatusInternalServerError,
			errMessage,
		}
	}
	return student, nil
}

func (s *StudentUsecase) GetAll(c context.Context) ([]domain.Student, *domain.ErrorResponse) {
	ctx, cancel := context.WithTimeout(c, s.ctxTimeout)
	defer cancel()
	searchResult, err := s.studentRepository.GetAll(ctx)
	if err != nil {
		errMessage := fmt.Sprintf("[Usecase][GetAllStudents]Fail to get all data in students: %s",
			err.Error())
		return nil, &domain.ErrorResponse{
			http.StatusInternalServerError,
			errMessage,
		}
	}
	students, err := ConvertElasticResultToStudents(searchResult)
	if err != nil {
		errMessage := fmt.Sprintf("[Usecase][ElasticConversion]Fail to convert elastic result to object: %s",
			err.Error())
		return nil, &domain.ErrorResponse{
			http.StatusInternalServerError,
			errMessage,
		}
	}
	return students, nil
}


func NewStudentUsecase(repository domain.StudentRepository, timeout time.Duration) domain.StudentUsecase{
	return &StudentUsecase{
		repository,
		timeout,
	}
}



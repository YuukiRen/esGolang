package domain

import (
	"context"
	"github.com/olivere/elastic/v7"
)

type Student struct {
	Name string	`json:"name"  validate:"required"`
	NIM string	`json:"nim"  validate:"required"`
	Age uint64	`json:"age"  validate:"required"`
	GPA float64 `json:"gpa"  validate:"required"`
}

type StudentUsecase interface {
	GetAll(c context.Context) ([]Student, *ErrorResponse)
	GetOneByNIM(c context.Context, nim string) (Student, *ErrorResponse)
	Create(c context.Context, student Student) (Student, *ErrorResponse)
	Delete(c context.Context, nim string) (Student, *ErrorResponse)
}

type StudentRepository interface {
	GetAll(context.Context) (*elastic.SearchResult, error)
	GetOne(context.Context, string, interface{}) (*elastic.SearchResult, error)
	Create(context.Context, Student) error
	Delete(context.Context, string, interface{}) error
}
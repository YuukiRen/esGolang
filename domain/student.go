package domain

import (
	"context"
	"github.com/olivere/elastic/v7"
)

type Student struct {
	Name string	`json:"name"`
	Age uint64	`json:"age"`
	GPA float64 `json:"gpa"`
}

type StudentUsecase interface {
	GetAll(c context.Context) ([]Student, *ErrorResponse)
}

type StudentRepository interface {
	GetAll(context.Context) (*elastic.SearchResult, error)
}
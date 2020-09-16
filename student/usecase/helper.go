package usecase

import (
	"encoding/json"
	"github.com/YuukiRen/esGolang/domain"
	"github.com/olivere/elastic/v7"
)

func ConvertElasticResultToStudents(result *elastic.SearchResult) (students []domain.Student, err error){
	for _, hit := range result.Hits.Hits{
		var student domain.Student
		err:= json.Unmarshal(hit.Source, &student)
		if err != nil{
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}
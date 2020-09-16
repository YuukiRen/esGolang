package elastic_search

import (
	"context"
	"github.com/YuukiRen/esGolang/domain"
	"github.com/olivere/elastic/v7"
)

type elasticStudent struct {
	esClient *elastic.Client
}
func NewElasticStudentRepository(client *elastic.Client) domain.StudentRepository{
	return &elasticStudent{esClient: client}
}


func (s *elasticStudent) GetAll(ctx context.Context) (*elastic.SearchResult, error) {
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchAllQuery())

	LogESQuery(searchSource)

	searchResult, err := s.esClient.Search().
		Index("students").
		SearchSource(searchSource).
		Do(ctx)
	if err != nil{
		return nil, err
	}
	return searchResult, nil
}

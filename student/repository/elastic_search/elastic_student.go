package elastic_search

import (
	"context"
	"encoding/json"
	"github.com/YuukiRen/esGolang/domain"
	"github.com/olivere/elastic/v7"
)

type elasticStudent struct {
	esClient *elastic.Client
}

func (s *elasticStudent) Delete(ctx context.Context, category string, value interface{}) error {
	_, err := elastic.NewDeleteByQueryService(s.esClient).
		Index("students").
		Query(elastic.NewMatchQuery(category,value)).
		Do(ctx)
	//TODO: look for a way to simulate transaction
	if err != nil{
		return err
	}
	return nil
}

func (s *elasticStudent) GetOne(ctx context.Context, category string, value interface{}) (*elastic.SearchResult, error) {
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery(category, value))

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

func (s *elasticStudent) Create(ctx context.Context, student domain.Student) error {
	dataJSON, err := json.Marshal(student)
	if err != nil{
		return err
	}
	js := string(dataJSON)

	_, err = s.esClient.Index().
		Index("students").
		BodyJson(js).
		Do(ctx)
	if err != nil{
		return err
	}
	return nil
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

func NewElasticStudentRepository(client *elastic.Client) domain.StudentRepository{
	return &elasticStudent{esClient: client}
}


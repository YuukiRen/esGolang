package main

import (
	"context"
	"github.com/YuukiRen/esGolang/domain"
	middL "github.com/YuukiRen/esGolang/middleware"
	studentDelivery "github.com/YuukiRen/esGolang/student/delivery"
	studentRepository "github.com/YuukiRen/esGolang/student/repository/elastic_search"
	studentUsecase "github.com/YuukiRen/esGolang/student/usecase"
	"github.com/labstack/echo"
	"github.com/olivere/elastic/v7"
	"log"
	"time"
)


func NewElasticClient() (*elastic.Client, error){
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	if err != nil{
		return nil, err
	}
	log.Println("Elastic client initialized")
	return client, nil
}
func main(){
	esClient, err := NewElasticClient()
	if err != nil {
		log.Fatalf("Fail to initialize Elastic Client: %v", err)
	}
	ctx := context.Background()

	e:= echo.New()
	middLware := middL.InitMiddleware()
	e.Use(middLware.Logging)
	e.Use(middLware.CORS)

	err = domain.CreateIndexIfDoesNotExist(ctx, esClient, "students")
	if err != nil{
		log.Fatalf("Fail to create new index: %s", err.Error())
	}
	sr := studentRepository.NewElasticStudentRepository(esClient)
	su := studentUsecase.NewStudentUsecase(sr, 2*time.Second)
	studentDelivery.NewStudentDelivery(e, su)

	log.Fatal(e.Start(":9090"))
}

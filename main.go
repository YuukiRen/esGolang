package main

import (
	"github.com/olivere/elastic/v7"
	"log"
	"github.com/labstack/echo"
	student_delivery "github.com/YuukiRen/esGolang/student/delivery"
	student_repository "github.com/YuukiRen/esGolang/student/repository/elastic_search"
	student_usecase "github.com/YuukiRen/esGolang/student/usecase"
	"time"
	middL "github.com/YuukiRen/esGolang/middleware"
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

	e:= echo.New()
	middLware := middL.InitMiddleware()
	e.Use(middLware.Logging)
	e.Use(middLware.CORS)
	sr := student_repository.NewElasticStudentRepository(esClient)
	su := student_usecase.NewStudentUsecase(sr, 2*time.Second)
	student_delivery.NewStudentDelivery(e, su)

	log.Fatal(e.Start(":9090"))
}


//func insert(esClient *elastic.Client, ctx context.Context){
//
//	student := Student{
//		Name: "Alvin Reinaldo",
//		Age:  21,
//		GPA:  3.84,
//	}
//
//	dataJSON, err := json.Marshal(student)
//	if err != nil{
//		log.Fatalf("Fail to marshal student object: %v", err)
//	}
//	js := string(dataJSON)
//	idx, err := esClient.Index().
//		Index("students").
//		BodyJson(js).
//		Do(ctx)
//	if err != nil{
//		log.Fatalf("Fail to insert json data to index: %v", err)
//	}
//	log.Println("Insertion Succeed", idx)
//}


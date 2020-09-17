package elastic_search

import (
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func LogESQuery(searchSource *elastic.SearchSource){
	queryStr, err1 := searchSource.Source()
	queryJs, err2 := json.Marshal(queryStr)

	if err1 != nil || err2 != nil {
		fmt.Println("[elastic_search][LogESQuery] err during query marshal=", err1, err2)
	}
	fmt.Println("[elastic_search][LogESQuery] Final ESQuery=\n", string(queryJs))
}
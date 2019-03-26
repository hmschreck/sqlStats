package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"time"
)

func SendToElk(elasticServer string, indexName string, input MySQLProcessList) {
	client, err := elastic.NewClient(elastic.SetURL(elasticServer))
	if err != nil {
		fmt.Println(input)
	}
	ctx := context.Background()
	timestamp := time.Now().Format("2006-01-02")
	fullIndexName := fmt.Sprintf("%s-%s", indexName, timestamp)
	bulkRequest := client.Bulk()
	for _, process := range input.Processes {
		jsonPacket, err := json.Marshal(process)
		if err != nil {
			panic(err)
		}
		jsonString := string(jsonPacket)
		newBulkRequest := elastic.NewBulkIndexRequest().Index(fullIndexName).Type("processlist").Doc(jsonString)
		//_, err = client.Index().Index(fullIndexName).Type("processlist").BodyJson(jsonString).Do(ctx)
		//if err != nil {
		//	panic(err)
		//}
		bulkRequest = bulkRequest.Add(newBulkRequest)
	}
	_, err = bulkRequest.Do(ctx)
	if err != nil {
		panic(err)
	}
}

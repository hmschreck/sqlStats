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
	for _, process := range input.Processes {
		process.DatabaseHost = input.DatabaseHost
		process.Date = input.Date
		jsonPacket, err := json.Marshal(process)
		if err != nil {
			panic(err)
		}
		jsonString := string(jsonPacket)
		_, err = client.Index().Index(fullIndexName).Type("processlist").BodyJson(jsonString).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
}

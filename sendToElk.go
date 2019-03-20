package main

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"time"
)


func SendToElk(elasticServer string, indexName string, jsonPacket string) {
	client, err := elastic.NewClient(elastic.SetURL(elasticServer))
	if err != nil {
		fmt.Println(jsonPacket)
	}
	ctx := context.Background()
	timestamp := time.Now().Format("01-02-15")
	fullIndexName := fmt.Sprintf("%s-%s", indexName, timestamp)
	_, err = client.Index().Index(fullIndexName).Type("processlist").BodyJson(jsonPacket).Do(ctx)
	if err != nil {
		panic(err)
	}
}




package main

import (
	"os"

	"github.com/pyaesone17/poc-module/dynamodb"
	"github.com/pyaesone17/poc-module/framework"
	"github.com/pyaesone17/poc-module/mongodb"
)

// Client implementation

func main() {
	datasource := buildDataAdaptor()
	datasource.Put("xyx", []string{"1", "2"})
	datasource.Find("xyx")
}

func buildDataAdaptor() framework.DataAdaptor {
	var adaptor framework.DataAdaptor
	if os.Getenv("db") == "mongodb" {
		adaptor = mongodb.NewRepository()
	}
	adaptor = dynamodb.NewRepository()

	return framework.NewDataSource(adaptor)
}

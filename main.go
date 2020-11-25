package main

import (
	"github.com/pyaesone17/poc-module/framework/datafactory"
)

// Client implementation

func main() {
	datasource := datafactory.BuildDataAdaptor("dynamodb")
	datasource.Put("xyx", []string{"1", "2"})
	datasource.Find("xyx")
}

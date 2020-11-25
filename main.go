package main

import "github.com/pyaesone17/poc-module/framework"

// Client implementation

func main() {
	datasource := framework.BuildDataAdaptor("dynamodb")
	datasource.Put("xyx", []string{"1", "2"})
	datasource.Find("xyx")
}

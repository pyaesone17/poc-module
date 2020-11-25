package datafactory

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/pyaesone17/poc-dynamodb-adaptor/dynamodbadaptor/repository"
	"github.com/pyaesone17/poc-module/framework"
	"github.com/pyaesone17/poc-mongodb-adaptor/mongodb"
)

type dataAdaptor struct {
	client  framework.DataAdaptor
	version string
}

// BuildDataAdaptor is building adaptor
func BuildDataAdaptor(dbtype string) framework.DataAdaptor {
	var adaptor framework.DataAdaptor
	if dbtype == "mongodb" {
		adaptor = mongodb.NewRepository()
	}

	adaptor = repository.NewRepository()
	return NewDataSource(adaptor)
}

func (dr dataAdaptor) Put(id string, data interface{}) {
	dr.debug()

	// May be we can do like framework level validation here
	dr.client.Put(id, data)
}

func (dr dataAdaptor) Find(id string) interface{} {
	dr.debug()
	// May be we can do like framework level validation here
	return dr.client.Find(id)
}

// NewDataSource is adaptor factory which will read env config and return appropriate adaptor
func NewDataSource(adaptor framework.DataAdaptor) framework.DataAdaptor {
	frameworkAdaptor := &dataAdaptor{client: adaptor, version: "0.0.1"}

	fmt.Println("-------------------------------------")
	fmt.Println("")

	c := color.New(color.FgGreen)
	c.Printf("Starting the framework %s \n", frameworkAdaptor.version)

	fmt.Println("")
	fmt.Println("-------------------------------------")

	return frameworkAdaptor
}

func (dr dataAdaptor) debug() {
	defer fmt.Println("-------------------------------------")

	c := color.New(color.FgCyan)

	c.Printf("Framework data adaptor version: %s \n", dr.version)
	c.Println("Processing framework level validation")
	c.Println("Passed framework level validation")
}

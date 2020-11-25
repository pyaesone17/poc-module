package framework

// DataAdaptor is source
type DataAdaptor interface {
	Put(id string, data interface{})
	Find(id string) interface{}
}

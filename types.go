package litebus

import (
	"fmt"
	"io"
	"reflect"
)

type Handler struct {
	call reflect.Value
	name string
}

type Result struct {
	Name string
	Error error
}

type Results []Result

func (r Results) IncludeError() bool {
	for _,v := range r {
		if v.Error != nil {
			return true
		}
	}
	return false
}

func (r Results) Error() string {
	var msgs string
	for _,v := range r {
		if v.Error != nil {
			msgs += fmt.Sprintf("[%s:%s] ", v.Name, v.Error)
		}
	}
	return msgs
}


type Subscriber interface{
	Subscribe(topic string,callback interface{}) error
}

type Publisher interface{
	Publish(topic string,params ...interface{}) Results
	PublishWithOption(topic string,opt *PublishOption,params ...interface{}) Results
}

type Processor interface{
	Subscriber
	Publisher
	io.Closer
}
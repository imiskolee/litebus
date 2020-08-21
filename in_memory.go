package litebus

import (
	"errors"
	"reflect"
)

var MsgPrefix = "[LiteBus.InMemoryProcessor]"

type InMemoryProcessor struct {
	handlers map[string][]Handler
}

func NewInMemoryProcessor() Processor {
	return &InMemoryProcessor{
		handlers: make(map[string][]Handler),
	}
}

func (p *InMemoryProcessor) Subscribe(topic string,call interface{}) error {
	rv := reflect.ValueOf(call)
	tv := reflect.TypeOf(call)
	if rv.Kind() != reflect.Func {
		return errors.New(MsgPrefix + " Wrong callback,it's must be reflect.Func")
	}
	if tv.NumOut() != 1 {
		return errors.New(MsgPrefix + " Wrong callback signature,it's must return single value as type error")
	}
	if p.handlers[topic] == nil {
		p.handlers[topic] = make([]Handler,0)
	}
	p.handlers[topic] = append(p.handlers[topic],Handler{
		call:rv,
		name : tv.Name(),
	})
	return nil
}

func (p *InMemoryProcessor) Publish(topic string,params ...interface{}) Results {
	return p.publish(topic,&PublishOption{},params...)
}

func (p *InMemoryProcessor) PublishWithOption(topic string,opt *PublishOption,params ...interface{}) Results {
	return p.publish(topic,opt,params...)
}

func (p *InMemoryProcessor) publish(topic string,opt *PublishOption,params ...interface{}) Results {
	handlers := p.handlers[topic]
	var inValue []reflect.Value
	for _,v := range params {
		inValue = append(inValue,reflect.ValueOf(v))
	}
	var results Results
	for _,handler := range handlers {
		out := handler.call.Call(inValue)
		isNil := out[0].IsNil()
		err,ok := out[0].Interface().(error)
		if !isNil && !ok {
			panic(MsgPrefix + " Wrong subscriber " + handler.name + " signature, it's must return a type error")
		}
		results = append(results,Result{
			Error: err,
			Name :handler.name,
		})
		if err != nil && opt.FailFirst {
			return results
		}
	}
	return results
}

func (p *InMemoryProcessor) Close() error {
	//dummy impl,nothing need to do for in memory processor.
	return nil
}




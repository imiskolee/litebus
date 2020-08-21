package litebus

import (
	"errors"
	"fmt"
	"testing"
)

func WrongCallback() (int,int) {
	return 0,0
}
func WrongCallbackZeroOut() {

}

func Callback(echo string) error {
	fmt.Println(echo)
	return nil
}

func CallbackWithError(echo string) error {
	fmt.Println(echo)
	return errors.New("Always error here. ")
}

func TestWrongCallback(t *testing.T) {
	bus := New(InMemory)
	{
		err := bus.Subscribe("a", WrongCallback)
		if err == nil {
			t.Fatal("must have error")
		}
	}
	{
		err := bus.Subscribe("a", WrongCallbackZeroOut)
		if err == nil {
			t.Fatal("must have error")
		}
	}
}


func TestCallback(t *testing.T) {
	bus := New(InMemory)
	{
		err := bus.Subscribe("a", Callback)
		if err != nil {
			t.Fatal("it's need success, bu got " + err.Error())
		}
	}
	{
		err := bus.Subscribe("a", CallbackWithError)
		if err != nil {
			t.Fatal("it's need success, bu got " + err.Error())
		}
	}
}

func TestCallbackRun(t *testing.T) {
	bus := New(InMemory)
	{
		err := bus.Subscribe("a", Callback)
		if err != nil {
			t.Fatal("it's need success, bu got " + err.Error())
		}
	}
	{
		err := bus.Subscribe("a", CallbackWithError)
		if err != nil {
			t.Fatal("it's need success, bu got " + err.Error())
		}
	}

	{
		results := bus.Publish("a","echo message")
		if !results.IncludeError() {
			t.Fatal("must be included error")
		}
	}
}









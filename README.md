# litebus
A lightweight event bus framework inside golang process. 

## Methods

### New

```go
 bus := litebus.New(litebus.InMemory)
```

### Subscribe
```go
func Callback(n1,n2,n3 int) error {
    return nil
}
if err := bus.Subsribe("topic",callback); err != nil {
    panic(err)
}
```





### Publish
publish a message to event bus.
```
results := bus.Publish("topic",1,2,3)
```














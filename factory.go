package litebus

const (
	InMemory = "in_memory"
)

func New(name string) Processor{
	switch name {
	case InMemory:
		return NewInMemoryProcessor()
	}
	panic(MsgPrefix + " Can not support this processor:" + name)
}





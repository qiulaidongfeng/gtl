package stack

type TsStack interface {
	Tspush(x interface{}) error
	Tspop() (interface{}, error)
	Tssize() uint64
	Tsclear()
	Tslook(size uint64) (interface{}, error)
}

type NoTsStack interface {
	Push(x interface{}) error
	Pop() (interface{}, error)
	Size() uint64
	Clear()
	Look(size uint64) (interface{}, error)
}

type Stack interface {
	TsStack
	NoTsStack
}

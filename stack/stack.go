package stack

type TsStack interface {
	Tspush(x interface{})
	Tspop() (interface{}, error)
	Tssize() uint64
	Tsclear()
	Tslook(size uint64) (interface{}, error)
}

type NoTsStack interface {
	Push(x interface{})
	Pop() (interface{}, error)
	Size() uint64
	Clear()
	Look(size uint64) (interface{}, error)
}

type Stack interface {
	TsStack
	NoTsStack
}

// list
package list

type NoTsNode interface {
	Set(x interface{}) error
	Get() (interface{}, error)
	Next() Node
	Prev() Node
}

type TsNode interface {
	Tsset(x interface{}) error
	Tsget() (interface{}, error)
	Tsnext() Node
	Tsprev() Node
}

type Node interface {
	NoTsNode
	TsNode
}

type Simplelist interface {
	Lnsert(x Node) error
	Remove() error
	Get(size uint64) (Node, error)
	Set(size uint64, x Node) error
}

type TsSimplelist interface {
	Tslnsert(x Node) error
	Tsremove() error
	Tsget(size uint64) (Node, error)
	Tsset(size uint64, x Node) error
}

type NoTslist interface {
	Simplelist
	TsSimplelist
	LnsertAt(size uint64, x Node) error
	RemoveAt(size uint64) error
	Len() uint64
	Clear()
	String() string
}

type Tslist interface {
	Simplelist
	TsSimplelist
	TsLnsertAt(size uint64, x Node) error
	TsRemoveAt(size uint64) error
	Tslen() uint64
	Tsclear()
	TsString() string
}

type List interface {
	NoTslist
	Tslist
}

type Splist = Simplelist
type TSplist = TsSimplelist

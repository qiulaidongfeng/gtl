package context

import (
	"context"
	"fmt"
	"sync"
)

func ExampleWithValue_map() {
	m1 := make(map[interface{}]interface{})
	m1["a"] = "b"
	ctx := WithValue(context.Background(), GoMap(m1))
	fmt.Println(ctx.Value("a"))
	//Output:
	//b
}

func ExampleWithValue_syncMap() {
	var m1 sync.Map
	m1.Store("a", "b")
	ctx := WithValue(context.Background(), &m1)
	fmt.Println(ctx.Value("a"))
	//Output:
	//b
}

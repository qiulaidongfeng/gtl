package gunsafe

func ExampleBtos() {
	b := []byte("hello")
	s := Btos(b)
	fmt.Println(s)
	//Output:
	//hello
}

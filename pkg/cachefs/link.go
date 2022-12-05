//go:build !go1.16
// +build !go1.16

package cachefs

//go:linkname readall ioutil.ReadAll
func readall(r Reader) ([]byte, error)

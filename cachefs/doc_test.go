package cachefs

import (
	"net/http"
)

func ExampleHttpCacheFs() {
	http.Handle("/", http.FileServer(NewHttpCacheFs(`E:\前端\html`)))
}

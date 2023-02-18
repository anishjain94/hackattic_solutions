package loadbalancer

// import (
// 	"net/http"
// 	"net/http/httputil"
// 	"net/url"
// )

// func main() {
// 	url, _ := url.Parse("http://localhost:8080")

// 	reverProxy := httputil.NewSingleHostReverseProxy(url)
// 	http.HandleFunc("/", reverProxy.ServeHTTP)

// }

// func (s *ServerPool) NextIndex() int {
// 	// return (int(atomic.AddInt64()))
// }

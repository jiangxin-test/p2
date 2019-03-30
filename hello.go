package hello

import (
	"fmt"

	h "github.com/jiangxin-test/hello/v3"
)

func SayHello(name string) {
	fmt.Println(h.Hello(name))
}

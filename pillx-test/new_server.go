// new_server
package main

import (
	"io"
	"../../pillX"
	"fmt"
)

func helloHandler(rw *pillx.Response, req *pillx.Request) {
	//fmt.Print(string(req.Content))
    io.WriteString(rw, "World")
}


func main() {
	server := &pillx.Server{
		Addr:          ":8080",
		Handler:        nil,
	}
	pillx.HandleFunc(0x0DDC, helloHandler)
	fmt.Println("pillX服务端引擎启动")
	server.ListenAndServe()
}

/**
* @Author:Tristan
* @Date: 2022/1/18 2:56 下午
 */

package main

import (
	"github.com/shanlongpan/grpc-mesh/grpcserver"
	"github.com/shanlongpan/grpc-mesh/httpserver"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go grpcserver.Serve(&wg, "17070")

	wg.Add(1)
	go httpserver.Serve(&wg, "17070", "18080")

	wg.Wait()
}

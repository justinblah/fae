package main

import (
	"fmt"
	"github.com/funkygao/fae/proxy"
	"github.com/funkygao/fae/servant/gen-go/fun/rpc"
	"time"
)

func main() {
	t1 := time.Now()

	client := proxy.Servant(":9001")
	defer client.Transport.Close()

	ctx := rpc.NewContext()
	ctx.Caller = "me"
	for i := 0; i < 10; i++ {
		r, _ := client.Ping(ctx)

		fmt.Println(r, time.Since(t1))
		t1 = time.Now()
	}
}

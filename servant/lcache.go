/*
local cache key:string, value:[]byte.
*/
package servant

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/funkygao/fae/servant/gen-go/fun/rpc"
	"github.com/funkygao/golib/cache"
)

func (this *FunServantImpl) onLcLruEvicted(key cache.Key, value interface{}) {

}

func (this *FunServantImpl) LcSet(ctx *rpc.Context,
	key string, value []byte) (r bool, intError error) {
	this.lc.Set(key, value)
	r = true

	return
}

func (this *FunServantImpl) LcGet(ctx *rpc.Context, key string) (r []byte,
	miss *rpc.TCacheMissed, intError error) {
	result, ok := this.lc.Get(key)
	if !ok {
		miss = rpc.NewTCacheMissed()
		miss.Message = thrift.StringPtr("lcache missed: " + key) // optional
	} else {
		r = result.([]byte)
	}

	return
}

func (this *FunServantImpl) LcDel(ctx *rpc.Context, key string) (intError error) {
	this.lc.Del(key)
	return
}

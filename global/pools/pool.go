package pools

import (
	"github.com/panjf2000/ants/v2"
	"kubernetes-go-demo/config"
)

var (
	pool *ants.Pool
)

func Pool() *ants.Pool {
	return pool
}

func InitPool() {
	poolNum := config.GetConfig().PoolNum
	var err error
	pool, err = ants.NewPool(poolNum)
	if err != nil {
		panic(err)
	}
}

func ClosePool(){
	pool.Release()
}

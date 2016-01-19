package main

import "github.com/lrsec/goroutine_pool"

func initReceingStage(input chan interface{}) (output chan interface{}, err error) {

	handler := func(interface{}) interface{} {

		//TODO read/write connection messages
		return nil
	}

	//TODO 调节 pool 参数
	pool, err := goroutine_pool.NewPool(10000, 1000*10000, 20000, 10*60*1000, 10000, 100, input, handler)
	if err != nil {
		return nil, nil, err
	}

	return input, pool.OutboundChannel, nil
}

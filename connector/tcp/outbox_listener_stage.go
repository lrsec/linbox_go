package main

import (
	"github.com/lrsec/goroutine_pool"
	"net"
)

func initOutboxListenerStage(input interface{}) (output chan interface{}, err error) {

	handler := func(input interface{}) interface{} {

		conn := input.(*net.Conn)

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

func readMsg(conn *net.Conn) (interface{}, error) {

}

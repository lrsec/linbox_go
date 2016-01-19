package main

import (
	log "github.com/cihub/seelog"
	//errors "github.com/lrsec/errors/wrapper"
	"github.com/lrsec/go-util/config"
	"github.com/lrsec/goroutine_pool"
	"net"
)

func main() {
	configs := &Config{}

	err := config.LoadConfig("configs/configs.json", configs)
	if err != nil {
		panic(err.Error())
	}

	// receiving stage
	receInput := make(chan interface{}, 10000)
	receOutput, err := initReceingStage(receInput)
	if err != nil {
		panic(err.Error())
	}

	// mq producer stage
	_, err = initMQProducerStage(receOutput)
	if err != nil {
		panic(err.Error())
	}

	// outbox listener stage
	obListenerInput := make(chan interface{}, 10000)
	obListenerOutput, err := initOutboxListenerStage(obListenerInput)
	if err != nil {
		panic(err.Error())
	}

	// Sending Stage
	_, err = initSendingStage(obListenerOutput)
	if err != nil {
		panic(err.Error())
	}

	// start tcp listener
	ln, err := net.Listen("tcp", configs.ImIp+":"+configs.ImPort)
	if err != nil {
		panic(err.Error())
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Error("Accept tcp connection fail. ", err)
			continue
		}

		receInput <- conn
	}
}

type Config struct {
	ImIp   string `json:"im_ip"`
	ImPort string `json:"im_port"`
}

package bloc_client

import (
	"time"
)

func (bC *blocClient) Run() {
	// periodicly register function
	bC.RegisterFunctionsToServer()
	go func(b *blocClient) {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			go bC.RegisterFunctionsToServer()
		}
	}(bC)

	// function consumer
	go bC.FunctionRunConsumerWithoutLocalObjectStorageImplemention()

	forever := make(chan struct{})
	<-forever
}

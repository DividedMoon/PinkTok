package service

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestErrProcess(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	errChannel := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		println("1 start")
		time.Sleep(2 * time.Second)
		println("1 finished")
		wg.Done()
	}()

	go func() {
		println("2 start")
		time.Sleep(1 * time.Second)
		errChannel <- fmt.Errorf("error")
		time.Sleep(1 * time.Second)
		println("2 finished")
		wg.Done()
	}()

	go func() {
		select {
		case err := <-errChannel:
			println("err:", err)
			cancel()
		case <-ctx.Done():
			println("ctx done")

		}
	}()
	wg.Wait()
}

package main

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestSy( t *testing.T)  {
	
	wg := sync.WaitGroup{}
	
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go printString("for ->> " +  strconv.FormatInt(int64(i),10), wg)
	}
	wg.Wait()
	fmt.Println("完成了")
}
func printString(s string,wg sync.WaitGroup) {
	wg.Done()
	fmt.Println(s)
}




func TestContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	
	go handle(ctx, 1500*time.Millisecond)
	
	select {
	case <- ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <- ctx.Done():
		fmt.Println("handle", ctx.Err())
	
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
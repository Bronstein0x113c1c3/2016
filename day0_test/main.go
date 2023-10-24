package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func testA(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done A")
			wg.Done()
			fmt.Println("A gone!")
			return
		default:
			fmt.Println("1")
		}
	}
}
func testB(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done B")
			wg.Done()
			fmt.Println("B gone!")
			return
		default:
			fmt.Println("2")
		}
	}
}
func testC(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done C")
			wg.Done()
			fmt.Println("C gone!")
			return
		default:
			fmt.Println("3")
		}
	}

}
func main() {
	wg := &sync.WaitGroup{}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	wg.Add(3)
	go testA(ctx, wg)
	go testB(ctx, wg)
	go testC(ctx, wg)
	wg.Wait()
	<-ctx.Done()
	// stop()
}

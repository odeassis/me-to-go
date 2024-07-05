package main

import (
	"context"
	"fmt"
	"time"
)

func printUntilCancel(ctx context.Context) {
	count := 0

	for {
		select {
		case <- ctx.Done():
			fmt.Println("Cancel signal received, exiting...")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("Printing until cancel, number = %d\n", count)
			count ++
		}
	}
}

func main() {
	ctx := context.Background();
	ctx, cancel := context.WithCancel(ctx);

	go printUntilCancel(ctx)

	defer cancel()

	time.Sleep(10 * time.Second)

}
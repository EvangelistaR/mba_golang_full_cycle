package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	timer := 5 * time.Second
	ctx, cancel := context.WithTimeout(ctx, timer)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// Simulate a long-running hotel booking process
	select {
	case <-time.After(3 * time.Second):
		println("Hotel booked successfully!")
	case <-ctx.Done():
		println("Failed to book hotel:", ctx.Err().Error())
	}
}

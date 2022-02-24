package main

import (
	"context"
	"fmt"
)

func ctxWithValue(ctx context.Context, key, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

func makeMagicWithCtx(ctx context.Context) string {
	val := ctx.Value("magic")
	if val == nil {
		return "nil"
	}
	return val.(string)
}

func main() {
	ctx := context.Background()
	ctx = ctxWithValue(ctx, "magic", "Wingardium Leviosa")
	response := makeMagicWithCtx(ctx)
	fmt.Println(response)
}

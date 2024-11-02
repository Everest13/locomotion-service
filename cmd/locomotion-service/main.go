package main

import (
	"context"
	"flag"
	"google.golang.org/grpc/grpclog"
	"locomotion-service/internal/app"
)

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		//todo logger
		return
	}

	flag.Parse()

	if err := a.Run(ctx); err != nil {
		grpclog.Fatal(err)
	}
}

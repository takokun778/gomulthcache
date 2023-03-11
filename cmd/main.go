package main

import (
	"gomulticache/internal/adaptor/controller"
	"gomulticache/internal/domain/model"
	"gomulticache/internal/driver/redis"
	"gomulticache/internal/driver/server"
	"gomulticache/internal/usecase/interactor"
	"os"
	"time"
)

func main() {
	cc, err := redis.NewCache()
	if err != nil {
		panic(err)
	}

	cl := redis.NewClient(os.Getenv("REDIS_URL"))

	rds, err := redis.New[model.Sample]().Of(cc, cl, 60*time.Second)
	if err != nil {
		panic(err)
	}

	sgu := interactor.NewSampleGet(rds)

	sdu := interactor.NewSampleDel(rds)

	ssu := interactor.NewSampleSet(rds)

	sample := controller.NewSample(ssu, sgu, sdu)

	server.NewHTTPServer(sample).Run()
}

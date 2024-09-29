package main

import (
	"context"
	"log"
	"time"

	"github.com/LaughG33k/songApi/iternal/app"
)

func main() {

	tm, canc := context.WithTimeout(context.TODO(), 15*time.Second)
	defer canc()
	app, err := app.NewApp(tm)

	if err != nil {
		log.Fatal(err)
	}

	app.Run()

}

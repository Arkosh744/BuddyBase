package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/Arkosh744/go-buddy-db/internal/database"
	"github.com/Arkosh744/go-buddy-db/internal/database/compute"
	"github.com/Arkosh744/go-buddy-db/internal/database/storage"
	"github.com/Arkosh744/go-buddy-db/internal/database/storage/mem"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	computeLayer, err := compute.NewCompute(compute.NewParser(), compute.NewAnalyzer(), log)
	if err != nil {
		log.Fatal("cannot start compute layer", zap.Error(err))
	}

	storageLayer, err := storage.NewStorage(mem.NewEngine(), log)
	if err != nil {
		log.Fatal("cannot start storage layer", zap.Error(err))
	}

	databaseLayer, err := database.NewDatabase(computeLayer, storageLayer, log)
	if err != nil {
		log.Fatal("cannot start database layer", zap.Error(err))
	}

	log.Info("database app started")

	for {
		fmt.Print("enter a command: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if err = scanner.Err(); err != nil {
			log.Fatal("scanner")
		}
		input := scanner.Text()

		switch input {
		case "stop", "exit", "quit":
			fmt.Println("database app terminated")
			return
		case "":
			fmt.Println("nothing provided")
			continue
		}

		resp, err := databaseLayer.HandleQuery(ctx, input)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(resp)
	}
}

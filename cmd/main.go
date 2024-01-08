package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	parser, err := cmpt.NewParser(logger)
	analyzer, err := cmpt.NewAnalyzer(logger)
	storage := strg.NewStorage()

	compute, err := cmpt.NewCompute(parser, analyzer, logger)

	database, err := db.NewDatabase(compute, storage, logger)
	if err != nil {
		logger.Fatal("cannot start database", zap.Error(err))
	}

	logger.Info("database app started")

	for {
		fmt.Print("enter a command: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		if err = scanner.Err(); err != nil {
			logger.Fatal("scanner")
		}
		input := scanner.Text()

		switch input {
		case "stop":
			logger.Info("database app terminated")
			return
		case "":
			fmt.Println("nothing provided")
			continue
		}

		resp, err := database.HandleQuery(ctx, input)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(resp)
	}
}

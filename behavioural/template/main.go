package main

import (
	"log"
)

func main() {
	Register("executorA", NewExecutorA())
	Register("executorB", NewExecutorB())

	executor, err := GetExecutor("executorB")
	if err != nil {
		log.Fatalln(err)
	}

	executor.Validate()
	executor.Convert()
	executor.Execute()

}

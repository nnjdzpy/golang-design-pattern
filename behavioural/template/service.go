package main

import (
	"fmt"
)

type ExecutorA struct {
}

func NewExecutorA() Executor {
	return &ExecutorA{}
}

func (e *ExecutorA) Validate() {
	fmt.Println("executorA validate")
}

func (e *ExecutorA) Convert() {
	fmt.Println("executorA convert")
}

func (e *ExecutorA) Execute() {
	fmt.Println("executorA execute")
}

type ExecutorB struct {
}

func NewExecutorB() Executor {
	return &ExecutorB{}
}

func (e *ExecutorB) Validate() {
	fmt.Println("executorB validate")
}

func (e *ExecutorB) Convert() {
	fmt.Println("executorB convert")
}

func (e *ExecutorB) Execute() {
	fmt.Println("executorB execute")
}

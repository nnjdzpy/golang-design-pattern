package main

import (
	"errors"
)

type Executor interface {
	Validate()
	Convert()
	Execute()
}

var ExecutorMap map[string]Executor

func Register(name string, executor Executor) {
	if ExecutorMap == nil {
		ExecutorMap = make(map[string]Executor, 0)
	}
	ExecutorMap[name] = executor
	return
}

func GetExecutor(name string) (Executor, error) {
	executor, ok := ExecutorMap[name]
	if !ok {
		return nil, errors.New("can not find")
	}
	return executor, nil
}

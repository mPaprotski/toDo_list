package main

import (
	"todoapp/scanner"
	"todoapp/todo"
)

func main() {
	todoList := todo.NewList()

	scanner := scanner.NewScanner(todoList)

	scanner.Start()

}

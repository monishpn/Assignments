package main

import "testing"

var msg = []string{
	"Hello",
	"World",
	"How are you?",
	"How was your day?",
}

func Test_create(t *testing.T) {
	exp := 3
	i := 0
	for i = 0; i < exp; i++ {
		CreateTask(i, msg[i])
	}

	op := len(m)

	op_msg := m[op-1]

	if op_msg != msg[exp-1] {
		t.Errorf("Expected: %v got: %v", msg[exp-1], op_msg)
	}

	if op != exp {
		t.Errorf("createTask failed, exp %d, got %d", exp, op)
	}

}

func TestDeleteTask(t *testing.T) {
	exp := 3

	for i := 0; i < exp; i++ {
		CreateTask(i, msg[i])
	}
	deleteTask(exp)
	op := len(m)
	op_msg := m[op-1]

	if op_msg != msg[exp-1] {
		t.Errorf("Expected: %v got: %v", msg[exp-1], op_msg)
	}
	if op != exp {
		t.Errorf("deleteTask failed, exp %d, got %d", exp, op)
	}
}

package main

import (
	"testing"
)

var msg = []string{
	"Hello",
	"World",
	"How are you?",
	"How was your day?",
	"Today was a nice day",
	"Yeah it was a nice day",
	"Bring in the Energy",
	"Should have done it today",
}

type test struct {
	step    []byte
	exp_len int
}

var tests = []test{
	{[]byte{'c', 'c', 'c'}, 3},
	{[]byte{'c', 'c', 'd'}, 1},
	{[]byte{'c', 'c', 'd', 'd'}, 0},
	{[]byte{'c', 'c', 'c', 'd'}, 2},
}

func Test(t *testing.T) {
	var i int = 0
	for test_case, T := range tests {
		i = 0

		for _, operation := range T.step {

			if operation == 'c' {

				//If the previous task is 'd' then the i value would have changed(line number:49), so increment it back
				if i != 0 && T.step[i-1] == 'd' {
					i++
				}

				CreateTask(i, msg[i])
				i++
			}

			if operation == 'd' {
				i = i - 1
				deleteTask(i)
				//Better not to increment here, because if the next operation is 'd' then it will delete the same id as of now
			}
		}

		if T.exp_len != len(m) {
			t.Errorf("%d : expected %v, got %v", test_case+1, T.exp_len, len(m))
		}

		if T.exp_len != 0 && msg[T.exp_len-1] != msg[len(m)-1] {
			t.Errorf("%d : expected %v, got %v", test_case+1, msg[T.exp_len-1], msg[len(m)-1])
		}

		for len(m) != 0 {
			deleteTask(len(m) - 1)
		}

	}
}

//func Test_create(t *testing.T) {
//	exp := 3
//	i := 0
//	for i = 0; i < exp; i++ {
//		CreateTask(i, msg[i])
//	}
//
//	op := len(m)
//
//	op_msg := m[op-1]
//
//	if op_msg != msg[exp-1] {
//		t.Errorf("Expected: %v got: %v", msg[exp-1], op_msg)
//	}
//
//	if op != exp {
//		t.Errorf("createTask failed, exp %d, got %d", exp, op)
//	}
//
//}
//
//func TestDeleteTask(t *testing.T) {
//	exp := 3
//
//	for i := 0; i < exp; i++ {
//		CreateTask(i, msg[i])
//	}
//	deleteTask(exp)
//	op := len(m)
//	op_msg := m[op-1]
//
//	if op_msg != msg[exp-1] {
//		t.Errorf("Expected: %v got: %v", msg[exp-1], op_msg)
//	}
//	if op != exp {
//		t.Errorf("deleteTask failed, exp %d, got %d", exp, op)
//	}
//}

package main

import (
	"fmt"
)

var m = make(map[int]string)

func idGen() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

func createTask(getID int, t string) {
	m[getID] = t
}

func viewTask() {

	fmt.Println("\n")
	fmt.Println("-------PENDING TASK--------")

	for key, value := range m {

		fmt.Printf("%d : %s\n", key, value)

	}
}

func deleteTask(id int) {
	if len(m) == 0 {
		fmt.Println("No Task Found")
		return
	}
	delete(m, id)

}

func main() {

	exit := 0

	getID := idGen()

	for exit == 0 {

		fmt.Printf("---ENTER THE ACTION TO DO--- \n " +
			"Press 1 : For Adding a task\n " +
			"Press 2 : View Pending Tasks\n" +
			"Press 3 : Mark task as done\n" +
			"Enter any other key to EXIT\n")

		var choice int
		_, err := fmt.Scanf("%d", &choice)

		if err != nil {
			exit = 1
		} else {

			switch choice {
			case 1:

				id := getID()
				var task string
				fmt.Println("\nEnter the task to do :")
				fmt.Scanf("%s", &task)

				createTask(id, task)
				fmt.Println("Task has been created\n\n")
			case 2:
				viewTask()
				fmt.Println("")
			case 3:
				var id int
				fmt.Println("\nEnter the task to delete :")
				fmt.Scanf("%d", &id)
				deleteTask(id)
				fmt.Println("Task has been deleted\n")
			default:
				exit = 1

			}

		}

	}

	fmt.Println("THANK YOUUU")

}

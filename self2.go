package main
 
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main(

    for {
        fmt.Println("\n")
        fmt.Println("1. Add Task")
        fmt.Println("2. List Tasks")
        fmt.Println("3. Mark Task as Done")
        fmt.Println("4. Delete Task")
        fmt.Println("5. Exit")
        fmt.Print("Enter choice: ")
 
        input, _ := reader.ReadString('\n')
        choice, err := strconv.Atoi(strings.TrimSpace(input))
 
        if err != nil {
            fmt.Println("Invalid choice, please try again.")
            continue
        } else {
            fmt.Printf("\nChoice is : %d\n\n", choice)
            switch choice {
            case 1:
                tasksList = addTask(taskCount, tasksList)
                taskCount++
            case 2:
                listTask(tasksList)
            case 3:
                tasksList = markDone(tasksList)
            case 4:
                tasksList = deleteTask(tasksList)
            case 5:
                exit()
            }
        }
 
    }
)
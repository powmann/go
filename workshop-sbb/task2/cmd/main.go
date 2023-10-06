package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type task struct {
	desc   string
	status bool
}

func main() {

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("-> ")
	// text, _ := reader.ReadString('\n')
	// text = strings.Replace(text, "\n", "", -1)

	// if strings.Compare("hi", text) == 0 {
	// 	fmt.Println("hello, Yourself")
	// }

	var tasks []task
	args := getArgs()
	if len(args) > 0 {
		tasks = readFromFile(args[0])
	}

	for {
		clearScreen()
		fmt.Println("MAIN MENU:")
		fmt.Println("1.add")
		fmt.Println("2.list")
		fmt.Println("3.complete")
		fmt.Println("4.delete")
		fmt.Println("5.save")
		fmt.Println("9.exit")
		fmt.Print("-> ")
		reader := bufio.NewReader(os.Stdin)
		opt, _ := reader.ReadString('\n')
		opt = strings.Replace(opt, "\n", "", -1)

		optint, err := strconv.Atoi(opt) // cast string to int
		if err != nil {
			printErrorMsg("invalid option")
			continue
		}

		switch optint {
		case 1:
			tasks = addTask(tasks)
		case 2:
			listTasks(tasks)
		case 3:
			tasks = completeTask(tasks)
		case 4:
			tasks = removeTask(tasks)
		case 5:
			writeToFile(tasks)
		case 9:
			os.Exit(0)
		default:
			printErrorMsg("invalid option")
		}
	}
}

func completeTask(tasks []task) []task {
	listOpenTasks(tasks)
	if len(tasks) > 0 {
		fmt.Println("which id? ")
		var in string
		fmt.Scanln(&in)
		i, err := strconv.Atoi(in) // cast string to int
		if err == nil {
			// validate input, if i is not in range of tasks, return
			if i < 0 || i > len(tasks)-1 {
				printErrorMsg("invalid id")
				return tasks
			}
		} else {
			printErrorMsg("invalid id")
			return tasks
		}
		tasks[i].status = false
		return tasks
	} else {
		return tasks
	}
}

func printErrorMsg(msg string) {
	// print error message in red
	fmt.Printf("\033[31m%s\033[0m\n", msg)
	fmt.Println("press return to continue")
	fmt.Scanln()
}

func removeTask(tasks []task) []task {
	listTasks(tasks)
	if len(tasks) > 0 {
		fmt.Println("which id? ")
		var in string
		fmt.Scanln(&in)
		i, err := strconv.Atoi(in) // cast string to int
		if err == nil {
			// validate input, if i is not in range of tasks, return
			if i < 0 || i > len(tasks)-1 {
				printErrorMsg("invalid id")
				return tasks
			}
		} else {
			printErrorMsg("invalid id")
			return tasks
		}
		return append(tasks[:i], tasks[i+1:]...) // re-slice and return
	} else {
		return tasks
	}
}

func listTasks(tasks []task) {
	if len(tasks) > 0 {
		fmt.Printf("\nlist ALL tasks:\n")
		for i, task := range tasks {
			fmt.Printf("id: %v task:%v status:%v \n", i, task.desc, task.status)
		}
	} else {
		fmt.Println("no tasks")
	}
	fmt.Printf("\npress return to continue")
	fmt.Scanln()
}

func listOpenTasks(tasks []task) {
	if len(tasks) > 0 {
		fmt.Printf("\nlist of tasks:\n")
		for i, task := range tasks {
			if task.status {
				fmt.Printf("id: %v task:%v \n", i, task.desc)
			}
		}
	} else {
		fmt.Println("no tasks")
	}
	fmt.Printf("\npress return to continue")
	fmt.Scanln()
}

func addTask(tasks []task) []task {
	fmt.Println("task desc:")
	var input string
	fmt.Scanln(&input)

	t := task{
		desc:   input,
		status: true,
	}
	tasks = append(tasks, t)
	return tasks

}

func clearScreen() {
	fmt.Println("\033[H\033[2J")
}

func writeToFile(tasks []task) {
	f, err := os.Create("tasks.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	for _, task := range tasks {
		// write desc and status to file separated by pipe

		f.WriteString(task.desc + "|" + strconv.FormatBool(task.status) + "\n")
	}
}

func readFromFile(filename string) []task {
	var tasks []task
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return tasks
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// split line into desc and status
		s := strings.Split(line, "|")
		status, err := strconv.ParseBool(s[1])
		if err != nil {
			fmt.Println(err)
			continue
		}
		t := task{
			desc:   s[0],
			status: status,
		}
		tasks = append(tasks, t)
	}
	return tasks
}

// retrieve the arguments passed to the program
func getArgs() []string {
	args := os.Args[1:] // ignore the first argument which is the program name
	return args
}

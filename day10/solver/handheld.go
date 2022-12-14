package solver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Zairian/advent-of-code-2022/utils"
)

func signalStrength() int {
	counter := 0
	sum := 0
	x := 1
	duringVal := 0
	queue := loadCommands("./day10/input.txt")

	duringVal, queue = executeCommands(20, &x, queue)
	counter += 20
	sum += counter * duringVal

	for i := 0; i < 5; i++ {
		duringVal, queue = executeCommands(40, &x, queue)
		counter += 40
		sum += counter * duringVal
	}

	return sum
}

func drawCRT() {
	queue := loadCommands("./day10/input.txt")
	x := 1
	duringVal := 0

	for i := 0; i < 6; i++ {
		for y := 0; y < 40; y++ {
			duringVal, queue = executeCommands(1, &x, queue)
			if duringVal-1 == y || duringVal == y || duringVal+1 == y {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// Loads commands from file into queue
func loadCommands(file string) []string {
	var queue []string
	s := utils.GetScanner(file)
	defer utils.CloseOS()

	for s.Scan() {
		queue = enqueue(queue, s.Text())
	}
	return queue
}

// Executes all the commands until the specified cycle and returns the value during the cycle and the updated queue
func executeCommands(cycle int, x *int, q []string) (int, []string) {
	var cmd string
	var duringVal int
	for i := 0; i < cycle; i++ {
		cmd, q = dequeue(q)
		op, val := parseCMD(cmd)
		if i == cycle-1 {
			duringVal = *x
		}
		execute(x, op, val)
	}
	return duringVal, q
}

func enqueue(q []string, e string) []string {
	op, _ := parseCMD(e)
	switch op {
	case "noop":
		q = append(q, e)
	case "addx":
		q = append(q, "")
		q = append(q, e)
	}
	return q
}

func dequeue(q []string) (string, []string) {
	x := q[0]
	q = q[1:]
	return x, q
}

// Returns true when execution is complete
func execute(x *int, op string, val int) {
	switch op {
	case "addx":
		*x += val
	}
}

// Returns the operation (string) and value from a command
func parseCMD(c string) (string, int) {
	cmd := strings.Split(c, " ")
	op := cmd[0]
	if len(cmd) > 1 {
		value, _ := strconv.Atoi(cmd[1])
		return op, value
	}
	return op, 0
}

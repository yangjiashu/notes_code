package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./testhash"
)

var cs *testhash.Map
var hadInit bool = false

func handleHashCommands(tokens []string) {
	switch tokens[1] {
	case "init":
		if hadInit {
			fmt.Println("已经初始化过了")
		} else if len(tokens) > 3 {
			val, err := strconv.Atoi(tokens[2])
			if err != nil {
				fmt.Println("replicas参数必须是整数，代表每个真实节点对应节点虚拟节点的个数")
			} else {
				cs = testhash.New(val, nil)
				cs.AddNodes(tokens[3:]...)
				hadInit = true
			}
		} else {
			fmt.Println("hash init <replicas>...<nodes>")
		}
	case "list":
		if !hadInit {
			fmt.Println("请先初始化系统：hash init <replicas>...<nodes>")
		} else {
			cs.List()
		}
	case "add":
		if !hadInit {
			fmt.Println("请先初始化系统：hash init <replicas>...<nodes>")
		} else if len(tokens) == 3 {
			cs.AddNode(tokens[2])
		} else {
			fmt.Println("USAGE: hash add <nodename>")
		}
	case "delete":
		if !hadInit {
			fmt.Println("请先初始化系统：hash init <replicas>...<nodes>")
		} else if len(tokens) == 3 {
			cs.DeleteNode(tokens[2])
		} else {
			fmt.Println("USAGE: hash delete <nodename>")
		}
	case "request":
		if !hadInit {
			fmt.Println("请先初始化系统：hash init <replicas>...<nodes>")
		} else if len(tokens) == 3 {
			fmt.Printf("处理请求[ %s ]的节点为[ %s ]\n", tokens[2], cs.Get(tokens[2]))
		} else {
			fmt.Println("USAGE: hash request <key>")
		}
	default:
		fmt.Println("Unrecognized hash command:", tokens[1])
	}
}

func main() {
	fmt.Println(`
			Enter following commands to control the system:
			hash list -- View the existing nodes
			hash init <replicas>...<nodes> -- Init the system, replicas represents the number of virtual nodes per reality node.
			hash add <nodename> -- Add a node.
			hash delete <nodename> -- Delete a node
			hash request <key> -- Send a request to the system`)

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "hash" {
			handleHashCommands(tokens)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}

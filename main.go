// i3go project main.go
package main

import (
	"fmt"
	"github.com/proxypoke/i3ipc"
	"os"
	"strings"
)

var (
	command = ""
)

const (
	Help    = "help"
	Version = "version"
	Tree    = "tree"
)

const Usage = `Usage: %s <command> [args]

	commands: 
		%s: shows this message
		%s: shows the version of i3.
		%s: shows the tree structure
`

func parse_cmd() {
	command = os.Args[0]

	var paths = strings.SplitAfter(command, "/")

	command = paths[len(paths)-1]
}

func show_usage() {
	fmt.Fprintf(os.Stderr, Usage, command, Help, Version, Tree)
}

func show_version(ipcsocket *i3ipc.IPCSocket) {
	if version, err := ipcsocket.GetVersion(); err == nil {
		fmt.Println("i3wm " + version.Human_Readable)
	} else {
		fmt.Fprintf(os.Stderr, "error: ", err)
		os.Exit(3)
	}
}

func show_tree(ipcsocket *i3ipc.IPCSocket) {
	if node, err := ipcsocket.GetTree(); err == nil {
		fmt.Println(node)
	} else {
		fmt.Fprintf(os.Stderr, "error: ", err)
		os.Exit(3)
	}
}

func main() {
	parse_cmd()
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "%s: Invalid usage:\n\n", command)
		show_usage()
		os.Exit(1)
	}
	if ipcsocket, err := i3ipc.GetIPCSocket(); err == nil {
		switch fn := os.Args[1]; fn {
		case Version:
			show_version(ipcsocket)
		case Tree:
			show_tree(ipcsocket)
		case Help:
			show_usage()
		}
	} else {
		fmt.Fprintf(os.Stderr, "socket error: ", err)
		os.Exit(2)
	}
}

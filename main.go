package main

import (
	"fmt"
	"os"
	"strings"
)

type Arguments struct {
	src_loc     string
	dst_loc     string
	num_workers int
	block_size  int
}

func main() {

	// capture arguments and store into Args struct
	for _, arg := range os.Args {
		// find index of equals sign in current arg (if avail)
		curr_equals_index := strings.IndexRune(arg, '=')
		if curr_equals_index == -1 { // failure -> couldn't find '=' rune in curr arg
			continue
		}

		// if equals sign avail, switch text before equals sign (to find which argument)
		switch arg[0:curr_equals_index] {
		case "--src":
			fmt.Println("Found src")

		case "--dst":
			fmt.Println("Found dst")

		case "--workers":
			fmt.Println("Found workers")

		case "--block_size":
			fmt.Println("Found block_size")

		default: // invalid argument -> skip this one
			continue
		}

	}

	// extrapolate src file from args

	// extrapolate dst file location from args

	// check that dst file location is valid

	// extrapolate num workers from args

	// extrapolate block size from args

}

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	arguments := Arguments{}
	for _, arg := range os.Args {
		// find index of equals sign in current arg (if avail)
		curr_equals_index := strings.IndexRune(arg, '=')
		if curr_equals_index == -1 { // failure -> couldn't find '=' rune in curr arg
			continue
		} else if len(arg) == (curr_equals_index + 1) {
			continue // failure -> no text after equals
		}
		curr_val_after_equals := arg[(curr_equals_index + 1):]

		// if equals sign avail, switch text before equals sign (to find which argument)
		switch arg[0:curr_equals_index] {
		case "--src":
			arguments.src_loc = curr_val_after_equals // validity not checked yet

		case "--dst":
			arguments.dst_loc = curr_val_after_equals // validity not checked yet

		case "--workers":
			num_workers, worker_parse_err := strconv.Atoi(curr_val_after_equals)
			if worker_parse_err != nil {
				log.Printf("ERROR: Could not parse worker argument (%s)", curr_val_after_equals)
				os.Exit(-1)
			}
			arguments.num_workers = num_workers

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

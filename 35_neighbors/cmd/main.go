package main

import (
	"fmt"
	"go_block_chain/35_neighbors/utils"
)

func main() {
	fmt.Println(utils.FindNeighbors("127.0.0.1", 5000, 0, 3, 5000, 5003))
}

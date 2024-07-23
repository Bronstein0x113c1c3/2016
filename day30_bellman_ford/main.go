package main

import (
	"fmt"
	"log"
	"os"
)

type segment struct {
	start  byte
	end    byte
	weight int
}

const MAX int = 1<<32 - 1

// func res(lowest byte, source byte,
// 	max_vertex byte, prev *[256]byte,
// 	distance *[256]int) {

// }
func res(distance []int, prev []byte) {
	fmt.Printf("Distance: \n")
	for i, _ := range distance {
		fmt.Printf("|%5v", i+1)
	}
	fmt.Println()
	for _, j := range distance {
		fmt.Printf("|%5v", j)
	}
	fmt.Println()
	fmt.Printf("Prev: \n")
	for i, _ := range distance {
		fmt.Printf("|%5v", i+1)
	}
	fmt.Println()
	for _, j := range prev {
		fmt.Printf("|%5v", j)
	}
	fmt.Println()
}
func bellman(lowest byte, source byte,
	max_vertex byte, prev *[256]byte,
	distance *[256]int, segments []segment, log *log.Logger) bool {
	for i := lowest; i < max_vertex; i++ {
		// if i == source {
		// 	continue
		// }
		for _, segment := range segments {
			u, v, w := segment.start, segment.end, segment.weight
			fmt.Println("------------------------------------------------")
			log.Printf("u: %v, v: %v, weight (u,v): %v \n", u, v, w)
			fmt.Printf("distance from %v to %v: %v\n", source, u, distance[u])
			fmt.Printf("distance from %v to %v: %v\n", source, v, distance[v])
			if distance[u] != MAX && distance[u]+w < distance[v] {
				log.Println("Change detected!!!!")
				fmt.Println("------------------------------------------------")
				distance[v] = distance[u] + w
				prev[v] = u
				fmt.Printf("new distance from %v to %v: %v \n", source, v, distance[v])
				fmt.Printf("go through %v\n", u)
				continue
			}
			log.Println("Unchanged!!!!")
		}
		fmt.Println("******************************************************")
		log.Printf("Done the big iteration %v \n", i)
		res((*distance)[lowest:max_vertex+1], (*prev)[lowest:max_vertex+1])
		fmt.Println("******************************************************")
	}
	for _, seg := range segments {
		u, v, w := seg.start, seg.end, seg.weight
		if distance[u] != MAX && distance[u]+w < distance[v] {
			log.Println("Negative cycle detected!")
			return false
		}
	}
	log.Println("found!!!")
	return true
}

func trace(destination byte, source byte, prev [256]byte) {
	if destination == source {
		fmt.Printf("%v", destination)
	} else {
		trace(prev[destination], source, prev)
		fmt.Printf("---->%v", destination)
	}
}
func main() {
	logger := log.Default()
	logger.SetOutput(os.Stdout)
	var prev [256]byte
	var distance [256]int
	//we will create point 0 to 5, source is 0
	source := byte(2)
	lowest := byte(1)
	max := byte(5)
	for i := lowest; i <= max; i++ {
		distance[i] = MAX
	}
	prev[source] = source
	distance[source] = 0
	segments := []segment{
		{1, 2, 6},
		{1, 4, 7},
		{2, 3, 5},
		{2, 5, -4},
		{2, 4, 8},
		{3, 2, -2},
		{4, 3, -3},
		{4, 5, 9},
		{5, 1, 2},
	}

	if bellman(lowest, source, max, &prev, &distance, segments, logger) {
		for i := lowest; i <= max; i++ {
			fmt.Printf("for dest %v: ", i)
			trace(i, source, prev)
			fmt.Println()
		}
	}
}

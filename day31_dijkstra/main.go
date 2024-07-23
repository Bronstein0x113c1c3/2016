package main

import (
	"fmt"
	"log"
	"os"
	"slices"
)

type Point struct {
	name                 string
	distance_from_source uint
}

// type Assoc struct {
// 	end     Point
// 	segment Segment
// }

type Segment struct {
	start  string
	end    string
	weight uint
}

var prev map[string]string

func FindSegment(p *Point, s []Segment) []Segment {
	res := make([]Segment, 0)
	for i := range s {
		if p.name == s[i].start {
			res = append(res, s[i])
		}
	}
	return res
}
func FindPoint(p []*Point, endpoint string) *Point {
	index := slices.IndexFunc(p, func(p *Point) bool {
		return p.name == endpoint
	})
	if index == -1 {
		return nil
	}
	return p[index]
}
func Relaxing(start *Point, end *Point, segment Segment, prev map[string]string, log *log.Logger) {
	if start == nil {
		log.Println("?????? problem with input")
		return
	}
	if end == nil {
		log.Println("it's visited!!!!")
		return
	}
	// fmt.Println("No Problem...")
	if start.distance_from_source+segment.weight < end.distance_from_source {
		log.Println("change detected!!!")
		end.distance_from_source = start.distance_from_source + segment.weight
		fmt.Printf("new source for %v: %v with distance %v, segment %v \n", end.name, start.name, end.distance_from_source, segment.weight)
		prev[end.name] = start.name
		return
	}
	log.Println("unchanged!!!")
}
func GetRes(list_of_point []*Point, prev map[string]string) {
	clone_list := make([]*Point, len(list_of_point))
	copy(clone_list, list_of_point)
	slices.SortStableFunc(clone_list, func(a, b *Point) int {
		x := 0
		if a.name[0] > b.name[0] {
			x = 1
		} else if a.name[0] == b.name[0] {
			x = 0
		} else if a.name[0] < b.name[0] {
			x = -1
		}
		return x
	})
	fmt.Println("the distance table:")
	for i := range clone_list {
		fmt.Printf("%10v|", clone_list[i].name)
	}
	fmt.Println()
	for i := range clone_list {
		if clone_list[i].distance_from_source == 1<<32-1 {
			fmt.Printf("%10v|", "inf")
			continue
		}
		fmt.Printf("%10v|", clone_list[i].distance_from_source)
	}
	fmt.Println()
	fmt.Println("the linking table:")
	for i := range clone_list {
		fmt.Printf("%10v|", clone_list[i].name)
	}
	fmt.Println()
	for i := range clone_list {
		if _, ok := prev[clone_list[i].name]; !ok {
			fmt.Printf("%10v|", "undef")
			continue
		}
		fmt.Printf("%10v|", prev[clone_list[i].name])

	}
	fmt.Println()
}

// func Route(dest A)
func Dijkstra(list_of_point_a []*Point, segments []Segment, prev map[string]string, log *log.Logger) {
	list_of_point := make([]*Point, len(list_of_point_a))
	copy(list_of_point, list_of_point_a)
	for len(list_of_point) > 0 {
		slices.SortStableFunc(list_of_point, func(a, b *Point) int {
			x := 0
			if a.distance_from_source < b.distance_from_source {
				x = 1
			} else if a.distance_from_source == b.distance_from_source {
				x = 0
			} else if a.distance_from_source > b.distance_from_source {
				x = -1
			}
			return x
		})

		// for i := range list_of_point {
		// 	fmt.Println(list_of_point[i])
		// }
		cur_point := list_of_point[len(list_of_point)-1]

		for _, segment := range FindSegment(cur_point, segments) {
			// log.Println(segment)
			first_point := FindPoint(list_of_point, segment.start)
			end_point := FindPoint(list_of_point, segment.end)
			// fmt.Println(first_point, " ", end_point)
			Relaxing(first_point, end_point, segment, prev, log)
		}
		list_of_point = slices.Delete(list_of_point, len(list_of_point)-1, len(list_of_point))
		GetRes(list_of_point_a, prev)
	}
}
func Route(cur string, prev map[string]string) {
	if _, ok := prev[cur]; !ok {
		fmt.Printf("%v", cur)
		return
	}
	Route(prev[cur], prev)
	fmt.Printf("---->%v", cur)
}
func main() {
	log := log.Default()
	log.SetOutput(os.Stdout)
	prev := map[string]string{}
	// prev["A"] = "A"
	list_of_point_a := []*Point{
		{"A", 1<<32 - 1},
		{"B", 1<<32 - 1},
		{"C", 1<<32 - 1},
		{"D", 1<<32 - 1},
		{"E", 1<<32 - 1},
		{"F", 1<<32 - 1},
		{"G", 1<<32 - 1},
	}
	list_of_point_a[1].distance_from_source = 0
	list_of_point := make([]*Point, len(list_of_point_a))
	copy(list_of_point, list_of_point_a)
	segments := []Segment{
		{"A", "C", 1},
		{"A", "D", 2},
		{"C", "A", 1},
		{"C", "D", 1},
		{"C", "B", 2},
		{"C", "E", 3},
		{"B", "C", 2},
		{"B", "F", 3},
		{"D", "A", 2},
		{"D", "C", 1},
		{"D", "G", 1},
		{"G", "D", 1},
		{"G", "F", 1},
		{"E", "C", 3},
		{"E", "F", 2},
	}

	// for len(list_of_point) > 0 {
	// 	slices.SortStableFunc(list_of_point, func(a, b *Point) int {
	// 		x := 0
	// 		if a.distance_from_source < b.distance_from_source {
	// 			x = 1
	// 		} else if a.distance_from_source == b.distance_from_source {
	// 			x = 0
	// 		} else if a.distance_from_source > b.distance_from_source {
	// 			x = -1
	// 		}
	// 		return x
	// 	})

	// 	// for i := range list_of_point {
	// 	// 	fmt.Println(list_of_point[i])
	// 	// }
	// 	cur_point := list_of_point[len(list_of_point)-1]

	// 	for _, segment := range FindSegment(cur_point, segments) {
	// 		// log.Println(segment)
	// 		first_point := FindPoint(list_of_point, segment.start)
	// 		end_point := FindPoint(list_of_point, segment.end)
	// 		// fmt.Println(first_point, " ", end_point)
	// 		Relaxing(first_point, end_point, segment, prev)
	// 	}
	// 	list_of_point = slices.Delete(list_of_point, len(list_of_point)-1, len(list_of_point))
	// 	GetRes(list_of_point_a, prev)
	// }
	Dijkstra(list_of_point_a, segments, prev, log)
	fmt.Println("the route...")
	Route("F", prev)
	fmt.Println()
}

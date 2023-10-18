package order

import (
	"sort"
)

type Node struct {
	Value int32
	Nodes []*Node
}

func Order(cityNodes int32, cityFrom []int32, cityTo []int32, company int32) []int32 {
	city := createNode(cityFrom, cityTo, cityFrom[0])
	output := city.Visit(make(map[int32]bool, 0), []int32{}, company)
	return output
}

func createNode(cityFrom, cityTo []int32, current int32) *Node {
	city := &Node{
		Value: current,
	}

	for i := 0; i < len(cityFrom); i++ {
		if cityFrom[i] == city.Value {
			city.Nodes = append(city.Nodes, createNode(cityFrom, cityTo, cityTo[i]))
		}
	}

	return city
}

func (city *Node) Visit(visited map[int32]bool, output []int32, company int32) []int32 {
	if !visited[city.Value] && company != city.Value {
		output = append(output, city.Value)
		visited[city.Value] = true
	}

	if len(city.Nodes) > 0 {
		sort.Slice(city.Nodes, func(i, j int) bool {
			return city.Nodes[i].Value < city.Nodes[j].Value
		})

		for i := 0; i < len(city.Nodes); i++ {
			if !visited[city.Nodes[i].Value] && company != city.Nodes[i].Value {
				output = append(output, city.Nodes[i].Value)
				visited[city.Nodes[i].Value] = true
			}
		}
		for i := 0; i < len(city.Nodes); i++ {
			output = city.Nodes[i].Visit(visited, output, company)
		}
	}
	return output
}

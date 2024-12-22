package day16

import (
	"fmt"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

type Crawler struct {
	position  [2]int
	direction [2]int
	path      [][2]int
	distance  int
}

func init() {
	registry.Registry["day16/part2"] = solve2
}

func solve2(inputFile string) (int64, error) {
	// Get grid from input
	grid, err := utils.ConstructGrid(inputFile)
	if err != nil {
		return 0, err
	}

	// Get source and destination
	source := [2]int{0, 0}
	dest := [2]int{0, 0}
	for i, row := range grid {
		for j, cell := range row {
			if cell == 'S' {
				source = [2]int{i, j}
			} else if cell == 'E' {
				dest = [2]int{i, j}
			}
		}
	}
	// Set source and dest to '.' in grid
	grid[source[0]][source[1]] = '.'
	grid[dest[0]][dest[1]] = '.'

	// Calculate paths
	paths := calculatePaths(grid, source, dest)

	return int64(paths), nil
}

func calculatePaths(grid [][]rune, source [2]int, destination [2]int) int {
	// A* algorithm to find shortest path(s)
	// Convert source and destination to Node
	src := utils.Node{X: source[0], Y: source[1]}
	dest := utils.Node{X: destination[0], Y: destination[1]}

	// Create open and closed sets
	nodes := map[utils.Node]*GraphMultiNode{} // All nodes
	queue := priorityqueue.NewWith(func(n1, n2 interface{}) int {
		a := nodes[n1.(utils.Node)]
		b := nodes[n2.(utils.Node)]
		diff := a.FCost - b.FCost
		if diff == 0 {
			return a.HCost - b.HCost
		}
		return diff
	}) // Priority queue
	visited := hashset.New() // Visited nodes

	// Initialize
	queue.Enqueue(src)
	nodes[src] = NewGraphMultiNode(0, utils.ManhattenDistance(src, dest))

	// Traverse the grid
	for !queue.Empty() {
		_node, _ := queue.Dequeue()
		node := _node.(utils.Node)
		graphNode := nodes[node]
		visited.Add(node)

		fmt.Print(node)
		if graphNode.Backrefs != nil {
			for _, backref := range graphNode.Backrefs.Values() {
				fmt.Print(" <- ", backref)
			}
		}
		print(" | ")
		if node == dest {
			continue
		}

		// Explore neighbors
		for _, dir := range utils.Directions {
			nb := utils.Node{X: node.X + dir[0], Y: node.Y + dir[1]}

			// If node is invalid, skip
			if nb.Invalid(len(grid), len(grid[0])) || grid[nb.X][nb.Y] != '.' {
				continue
			}

			cost := graphNode.GCost + 1
			// Project current direction backwards to find expected last position
			lastPos := utils.Node{X: node.X - dir[0], Y: node.Y - dir[1]}
			// Check if last position is a backref
			if !graphNode.Backrefs.Contains(lastPos) {
				// Change of direction, increase cost
				cost += 1000
			}

			fmt.Print(nb, cost)

			nbNode, ok := nodes[nb]
			if ok {
				fmt.Print(" O ", nbNode.GCost, graphNode.GCost)
			}
			if !ok {
				fmt.Print(" N ")
				nbNode = NewGraphMultiNode(cost, utils.ManhattenDistance(nb, dest))
				// Add to backrefs
				nbNode.Backrefs = hashset.New()
				nbNode.Backrefs.Add(node)
			} else if nbNode.GCost > cost {
				fmt.Print(" H ", nb)
				// Update old node with new Cost
				nbNode.UpdateGCost(cost + 1)
				// Replace backrefs with new hashset
				nbNode.Backrefs = hashset.New()
				nbNode.Backrefs.Add(node)
			} else if nbNode.GCost == cost {
				fmt.Print(" E ")
				// Equal cost, add to backrefs
				nbNode.Backrefs.Add(node)
			} else {
				continue
			}
			if visited.Contains(nb) {
				continue
			}
			nodes[nb] = nbNode

			// Add node for further processing
			queue.Enqueue(nb)
		}
		println()
	}

	path := constructPath(dest, nodes)

	// Print paths on grid
	for i, row := range grid {
		for j, cell := range row {
			if path.Contains(utils.Node{X: i, Y: j}) {
				fmt.Print("O")
			} else {
				fmt.Print(string(cell))
			}
		}
		println()
	}

	return path.Size()
}

func constructPath(node utils.Node, nodes map[utils.Node]*GraphMultiNode) *hashset.Set {
	path := hashset.New()
	stack := []utils.Node{node}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if path.Contains(current) {
			continue
		}

		path.Add(current)
		for _, backref := range nodes[current].Backrefs.Values() {
			stack = append(stack, backref.(utils.Node))
		}
	}

	return path
}

type GraphMultiNode struct {
	utils.GraphNode
	Backrefs *hashset.Set
}

func NewGraphMultiNode(gCost, fCost int) *GraphMultiNode {
	return &GraphMultiNode{*utils.NewGraphNode(gCost, fCost), hashset.New()}
}

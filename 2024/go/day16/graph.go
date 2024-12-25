package day16

import (
	"github.com/daveydark/adventofcode/2024/internal/utils"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

// Graph State: [position][source] -> [distance, count]
type GraphState = map[utils.Node]map[utils.Node][2]int

func calculatePaths(grid [][]rune, source [2]int, destination [2]int) int {
	// A* algorithm to find shortest path(s)
	// Convert source and destination to Node
	src := utils.Node{X: source[0], Y: source[1]}
	// dest := utils.Node{X: destination[0], Y: destination[1]}

	// Init Structures
	graph := make(GraphState)
	visited := hashset.New()
	pq := priorityqueue.NewWith(func(a, b interface{}) int {
		n1 := a.([2]utils.Node)
		n2 := b.([2]utils.Node)
		return graph[n1[0]][n1[1]][0] - graph[n2[0]][n2[1]][0]
	})

	// Start from source, facing right
	graph[src][utils.NewNode(src.X, src.Y-1)] = [2]int{0, 1}
	pq.Enqueue([2]utils.Node{src, utils.NewNode(src.X, src.Y-1)})

	// Traverse using A* algorithm
	for !pq.Empty() {
		_node, _ := pq.Dequeue()
		node := _node.([2]utils.Node)
		nodeSrc := node[1]
		currNode := node[0]
		currData := graph[currNode][nodeSrc]
		visited.Add(node)

		for _, dir := range utils.Directions {
			nb := utils.NewNode(currNode.X+dir[0], currNode.Y+dir[1])
			adjNode := [2]utils.Node{currNode, nb}

			// Check if node is valid
			if nb.Invalid(len(grid), len(grid[0])) || grid[nb.X][nb.Y] != '.' || visited.Contains(adjNode) {
				continue
			}

			// Calculate new distance
			newDist := currData[0] + 1
			// If we turned, add 1000 to distance
			if dir[0] != nodeSrc.X-currNode.X || dir[1] != nodeSrc.Y-currNode.Y {
				newDist += 1000
			}

			// Get data from graph
			oldData, ok := graph[currNode][nodeSrc]
			if !ok {
				oldData = [2]int{0, 0}
			} else {
				// If we have a shorter path, skip
				if newDist >= oldData[0] {
					continue
				}
			}
			newData := [2]int{newDist, oldData[1] + 1}
			graph[currNode][nodeSrc] = newData
		}
	}

	return 0
}

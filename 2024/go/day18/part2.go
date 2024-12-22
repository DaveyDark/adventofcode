package day18

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/daveydark/adventofcode/2024/internal/registry"
	"github.com/daveydark/adventofcode/2024/internal/utils"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

func init() {
	registry.Registry["day18/part2"] = solve2
}

func solve2(inputFile string) (int64, error) {
	// Make grid
	grid := [GRID_SIZE][GRID_SIZE]bool{}

	// Read input file and create scanner
	file, err := os.Open(inputFile)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)

	// Record all obstacles
	bytes := [][2]int{}
	for scanner.Scan() {
		line := scanner.Text()
		numsStr := strings.Split(line, ",")
		x, _ := strconv.Atoi(numsStr[0])
		y, _ := strconv.Atoi(numsStr[1])
		bytes = append(bytes, [2]int{x, y})
	}

	// Add obstacles to grid upto BYTE_LIMIT
	for i := 0; i < BYTE_LIMIT; i++ {
		grid[bytes[i][1]][bytes[i][0]] = true
	}

	// Calculate shortest path
	path := calculatePath(grid)

	// Add remaining obstacles to grid one by one
	for i := BYTE_LIMIT + 1; i < len(bytes); i++ {
		node := utils.Node{X: bytes[i][1], Y: bytes[i][0]}
		grid[node.X][node.Y] = true
		if path.Contains(node) {
			// Recalculate shortest path
			path = calculatePath(grid)
			// Check if path is still possible
			if path == nil {
				// Since we can only return i64, we will print the result to stdout
				fmt.Println(strconv.Itoa(bytes[i][0]) + "," + strconv.Itoa(bytes[i][1]))
				return 1, nil
			}
		}
	}

	fmt.Println(bytes[len(bytes)-1][0], ",", bytes[len(bytes)-1][1])
	return 1, nil
}

func calculatePath(grid [GRID_SIZE][GRID_SIZE]bool) *hashset.Set {
	// A* algorithm
	// utils.Nodes map to store information about all nodes
	nodes := map[utils.Node]*utils.GraphNode{}
	start := utils.Node{X: 0, Y: 0}                       // Start node
	end := utils.Node{X: GRID_SIZE - 1, Y: GRID_SIZE - 1} // End node
	// Create start node
	nodes[start] = utils.NewGraphNode(0, utils.ManhattenDistance(start, end))
	// Create Priority Queue to provide nodes
	queue := priorityqueue.NewWith(func(n1, n2 interface{}) int {
		a := nodes[n1.(utils.Node)]
		b := nodes[n2.(utils.Node)]
		diff := a.FCost - b.FCost
		if diff == 0 {
			return a.HCost - b.HCost
		}
		return diff
	})
	queue.Enqueue(start)
	visited := hashset.New()

	for !queue.Empty() {
		// Grab next node
		_node, _ := queue.Dequeue()
		node := _node.(utils.Node)
		graphNode := nodes[node]
		visited.Add(node)

		// If we reached destination, stop
		if node == end {
			break
		}

		// Explore neighbors
		for _, dir := range utils.Directions {
			nb := utils.Node{X: node.X + dir[0], Y: node.Y + dir[1]}

			// If the node is invalid, skip it
			if nb.Invalid(GRID_SIZE, GRID_SIZE) || grid[nb.X][nb.Y] || visited.Contains(nb) {
				continue
			}

			nbNode, ok := nodes[nb]
			if !ok {
				// Create new Graphutils.Node
				nbNode = utils.NewGraphNode(graphNode.GCost+1, utils.ManhattenDistance(nb, end))
				nbNode.Backref = &node
				nodes[nb] = nbNode
			} else if nbNode.GCost > graphNode.GCost+1 {
				// Update old node with new Cost
				nbNode.UpdateGCost(graphNode.GCost + 1)
				nbNode.Backref = &node
			} else {
				continue
			}

			// Add node for further processing
			queue.Enqueue(nb)
		}
	}

	// Check if path is impossible
	_, ok := nodes[end]
	if !ok {
		return nil
	}

	// Trace path using backrefs
	path := hashset.New()
	next := &end
	for next != nil {
		path.Add(*next)
		next = nodes[*next].Backref
	}

	return path
}

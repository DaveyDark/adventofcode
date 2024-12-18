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
		node := Node{bytes[i][1], bytes[i][0]}
		grid[node.x][node.y] = true
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

type Node struct {
	x int
	y int
}

func (this Node) invalid() bool {
	return this.x < 0 || this.y < 0 || this.x == GRID_SIZE || this.y == GRID_SIZE
}

type GraphNode struct {
	gCost   int // Distance from start node
	hCost   int // Distance from end node
	fCost   int // gCost + hCost
	backref *Node
}

func (g *GraphNode) updateGCost(gc int) {
	g.gCost = gc
	g.fCost = gc + g.hCost
}

func NewGraphNode(gCost, fCost int) *GraphNode {
	return &GraphNode{gCost, fCost, gCost + fCost, nil}
}

func calculatePath(grid [GRID_SIZE][GRID_SIZE]bool) *hashset.Set {
	// A* algorithm
	// Nodes map to store information about all nodes
	nodes := map[Node]*GraphNode{}
	start := Node{0, 0}                       // Start node
	end := Node{GRID_SIZE - 1, GRID_SIZE - 1} // End node
	// Create start node
	nodes[start] = NewGraphNode(0, distanceBetween(start, end))
	// Create Priority Queue to provide nodes
	queue := priorityqueue.NewWith(func(n1, n2 interface{}) int {
		a := nodes[n1.(Node)]
		b := nodes[n2.(Node)]
		diff := a.fCost - b.fCost
		if diff == 0 {
			return a.hCost - b.hCost
		}
		return diff
	})
	queue.Enqueue(start)
	visited := hashset.New()

	for !queue.Empty() {
		// Grab next node
		_node, _ := queue.Dequeue()
		node := _node.(Node)
		graphNode := nodes[node]
		visited.Add(node)

		// If we reached destination, stop
		if node == end {
			break
		}

		// Explore neighbors
		for _, dir := range utils.Directions {
			nb := Node{node.x + dir[0], node.y + dir[1]}

			// If the node is invalid, skip it
			if nb.invalid() || grid[nb.x][nb.y] || visited.Contains(nb) {
				continue
			}

			nbNode, ok := nodes[nb]
			if !ok {
				// Create new GraphNode
				nbNode = NewGraphNode(graphNode.gCost+1, distanceBetween(nb, end))
				nbNode.backref = &node
				nodes[nb] = nbNode
			} else if nbNode.gCost > graphNode.gCost+1 {
				// Update old node with new Cost
				nbNode.updateGCost(graphNode.gCost + 1)
				nbNode.backref = &node
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
		next = nodes[*next].backref
	}

	return path
}

func distanceBetween(start, end Node) int {
	xDist := start.x - end.x
	yDist := start.y - end.y
	if xDist < 0 {
		xDist = -xDist
	}
	if yDist < 0 {
		yDist = -yDist
	}
	return xDist + yDist
}

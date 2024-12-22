package utils

func ManhattenDistance(start, end Node) int {
	xDist := start.X - end.X
	yDist := start.Y - end.Y
	if xDist < 0 {
		xDist = -xDist
	}
	if yDist < 0 {
		yDist = -yDist
	}
	return xDist + yDist
}

// Graph related structs

type Node struct {
	X int
	Y int
}

func (this Node) Invalid(rows, cols int) bool {
	return this.X < 0 || this.Y < 0 || this.X == rows || this.Y == cols
}

type GraphNode struct {
	GCost   int // Distance from start node
	HCost   int // Distance from end node
	FCost   int // gCost + hCost
	Backref *Node
}

func (g *GraphNode) UpdateGCost(gc int) {
	g.GCost = gc
	g.FCost = gc + g.HCost
}

func NewGraphNode(gCost, fCost int) *GraphNode {
	return &GraphNode{gCost, fCost, gCost + fCost, nil}
}

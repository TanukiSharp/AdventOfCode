package day16

import (
	"aoc/2022/shared"
	"math"
	"sort"
)

type AStar struct {
	openSet  shared.HashSet[*Node]
	cameFrom map[*Node]*Node
	fScore   map[*Node]int64
	gScore   map[*Node]int64
}

func (astar *AStar) reconstructPath(current *Node) []*Node {
	totalPath := []*Node{current}
	var ok bool

	for {
		current, ok = astar.cameFrom[current]

		if ok == false {
			break
		}

		totalPath = append([]*Node{current}, totalPath...)
	}

	return totalPath
}

func getScore(node *Node, scoreMap map[*Node]int64) int64 {
	score, ok := scoreMap[node]
	if ok {
		return score
	}
	return math.MaxInt32
}

// A* finds a path from start to goal.
// h is the heuristic function. h(n) estimates the cost to reach goal from node n.
func (astar *AStar) Run(start, goal *Node, h func(*Node) int64, d func(*Node, *Node) int64) []*Node {
	// The set of discovered nodes that may need to be (re-)expanded.
	// Initially, only the start node is known.
	// This is usually implemented as a min-heap or priority queue rather than a hash-set.
	astar.openSet = shared.HashSet[*Node]{}
	astar.openSet.Add(start)

	// For node n, cameFrom[n] is the node immediately preceding it on the cheapest path from start
	// to n currently known.
	astar.cameFrom = map[*Node]*Node{}

	// For node n, gScore[n] is the cost of the cheapest path from start to n currently known.
	astar.gScore = map[*Node]int64{}
	astar.gScore[start] = 0.0

	// For node n, fScore[n] := gScore[n] + h(n). fScore[n] represents our current best guess as to
	// how cheap a path could be from start to finish if it goes through n.
	astar.fScore = map[*Node]int64{}
	astar.fScore[start] = h(start)

	for len(astar.openSet) > 0 {
		// This operation can occur in O(Log(N)) time if openSet is a min-heap or a priority queue
		current := astar.findNodeWithBestScore(astar.fScore)

		if current == goal {
			return astar.reconstructPath(current)
		}

		astar.openSet.Remove(current)

		for _, neighbor := range current.neighbors {
			astar.next(current, neighbor, h, d)
		}
	}

	// Open set is empty but goal was never reached
	return nil
}

type scoredNode struct {
	node  *Node
	score int64
}

func (astar *AStar) findNodeWithBestScore(scoreMap map[*Node]int64) *Node {
	temp := []*scoredNode{}

	for node := range astar.openSet {
		score, ok := astar.fScore[node]

		if ok {
			temp = append(temp, &scoredNode{node, score})
		}
	}

	sort.Slice(temp, func(i, j int) bool { return temp[i].score < temp[j].score })

	return temp[0].node
}

func (astar *AStar) next(current *Node, nextNode *Node, h func(*Node) int64, d func(*Node, *Node) int64) {
	// d(current,neighbor) is the weight of the edge from current to neighbor
	// tentative_gScore is the distance from start to the neighbor through current

	gScoreCurrent := getScore(current, astar.gScore)

	if gScoreCurrent == math.MaxInt32 {
		return
	}

	gScoreNext := getScore(nextNode, astar.gScore)

	var tentativeGScore int64 = gScoreCurrent + d(current, nextNode)

	if tentativeGScore < gScoreNext {
		// This path to neighbor is better than any previous one. Record it!
		astar.cameFrom[nextNode] = current
		astar.gScore[nextNode] = tentativeGScore
		astar.fScore[nextNode] = tentativeGScore + h(nextNode)
		if astar.openSet.Contains(nextNode) == false {
			astar.openSet.Add(nextNode)
			nextNode.isValveOpened = true
		}
	}
}

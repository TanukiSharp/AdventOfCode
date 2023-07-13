package day16

import (
	"aoc/2022/shared"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Puzzle struct{}

func (*Puzzle) Day() int     { return 16 }
func (*Puzzle) IsTest() bool { return true }

func (*Puzzle) Run(input string) {
	graph := parseInput(input)

	startNode := graph.entryPoint
	clock := NewClock()

	for startNode != nil {
		travel := oneStep(startNode, clock, graph)

		if travel == nil {
			break
		}

		travel.Print()

		lastNodeName := travel.nodes[len(travel.nodes)-1]
		startNode = graph.nodes[lastNodeName]
		clock = travel.clock
		clock.remainingTime += 1

		for _, name := range travel.nodes {
			graph.nodes[name].isValveOpened = true
		}

		// Dirty hack below.
		graph.nodes[lastNodeName].flowRate = 0
		graph.nodes[lastNodeName].isValveOpened = false
	}
}

func oneStep(startNode *Node, clock Clock, graph *Graph) *Travel {
	travelWalk := &TravelWalk{
		travels: []*Travel{},
	}

	walk(startNode, NewTravel(clock), travelWalk)

	if len(travelWalk.travels) == 0 {
		return nil
	}

	if len(travelWalk.travels) == 1 {
		return travelWalk.travels[0]
	}

	sort.Slice(travelWalk.travels, func(i, j int) bool {
		timeRemainDiff := travelWalk.travels[i].clock.remainingTime - travelWalk.travels[j].clock.remainingTime

		if timeRemainDiff > 0 {
			return true
		} else if timeRemainDiff < 0 {
			return false
		}

		return travelWalk.travels[i].clock.currentScore > travelWalk.travels[j].clock.currentScore
	})

	return travelWalk.travels[0]
}

func (t *Travel) Print() {
	fmt.Printf("Score: %d\n", t.clock.currentScore)
	fmt.Printf("Remaining time: %d\n", t.clock.remainingTime)
	fmt.Printf("Path: %s\n", strings.Join(t.nodes, " -> "))
	fmt.Println()
}

// d := func(from, to *Node) int64 {
// 	return 1
// }

// h := func(to *Node) int64 {
// 	if to.isValveOpened {
// 		return 0
// 	}
// 	return (int64)(to.flowRate)
// }

// astar := &AStar{}

// travel := astar.Run(graph.entryPoint, graph.nodes["HH"], h, d)

// for _, node := range travel {
// 	fmt.Printf("%s\n", node.name)
// }

type Clock struct {
	remainingTime int
	currentScore  int
}

func NewClock() Clock {
	return Clock{
		remainingTime: 31,
		currentScore:  0,
	}
}

func (c *Clock) Travel() Clock {
	return Clock{
		remainingTime: shared.Max(0, c.remainingTime-1),
		currentScore:  c.currentScore,
	}
}

func (c *Clock) OpenValve(valveFlowRate int) Clock {
	remainingTime := shared.Max(0, c.remainingTime-1)
	return Clock{
		remainingTime: remainingTime,
		currentScore:  c.currentScore + shared.Max(0, (remainingTime-1))*valveFlowRate,
	}
}

func walk(node *Node, currentTravel *Travel, travelWalk *TravelWalk) {
	if node.isValveOpened {
		return
	}

	node.isValveOpened = true

	if currentTravel.clock.remainingTime <= 1 {
		return
	}

	currentTravel.TravelTo(node.name)

	for _, child := range node.neighbors {
		walk(child, currentTravel.Clone(), travelWalk)
	}

	if node.flowRate > 0 {
		currentTravel.clock = currentTravel.clock.OpenValve(node.flowRate)
		travelWalk.travels = append(travelWalk.travels, currentTravel)
	}
}

type TravelWalk struct {
	travels []*Travel
}

type Travel struct {
	id    *string
	clock Clock
	nodes []string
}

func NewTravel(clock Clock) *Travel {
	return &Travel{
		clock: clock,
		nodes: []string{},
	}
}

func (t *Travel) Clone() *Travel {
	newNodes := make([]string, len(t.nodes))
	copy(newNodes, t.nodes)

	return &Travel{
		clock: t.clock,
		nodes: newNodes,
	}
}

func (t *Travel) TravelTo(node string) *Travel {
	t.clock = t.clock.Travel()
	t.nodes = append(t.nodes, node)
	return t
}

func (t *Travel) GetId() string {
	if t.id == nil {
		*(t.id) = strings.Join(t.nodes, ":")
	}
	return *(t.id)
}

func parseInput(input string) *Graph {
	intermediateNodes := map[string][]string{}

	var rootNode *Node
	nodes := map[string]*Node{}

	for _, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)

		if len(line) == 0 {
			continue
		}

		regex, _ := regexp.Compile(`Valve (\w+) has flow rate=(\d+); tunnel(s?) lead(s?) to valve(s?) (.+)`)
		matches := regex.FindSubmatch(([]byte)(line))

		name := (string)(matches[1])
		flowRate, _ := strconv.Atoi((string)(matches[2]))
		neighbors := strings.Split((string)(matches[6]), ", ")

		node := &Node{
			name:      name,
			flowRate:  flowRate,
			neighbors: map[string]*Node{},
		}

		nodes[name] = node

		if name == "AA" {
			rootNode = node
		}

		for _, neighborName := range neighbors {
			list := intermediateNodes[name]
			if list == nil {
				list = []string{}
			}
			list = append(list, neighborName)
			intermediateNodes[name] = list
		}
	}

	for name, node := range nodes {
		for _, neighborName := range intermediateNodes[name] {
			node.neighbors[neighborName] = nodes[neighborName]
		}
	}

	return &Graph{
		entryPoint: rootNode,
		nodes:      nodes,
	}
}

type Node struct {
	name          string
	flowRate      int
	isValveOpened bool
	neighbors     map[string]*Node
}

type Graph struct {
	entryPoint *Node
	nodes      map[string]*Node
}

func (g *Graph) ResetValves() {
	for _, node := range g.nodes {
		node.isValveOpened = false
	}
}

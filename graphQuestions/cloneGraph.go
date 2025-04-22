package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func main2() {

	//to be fixed
	//adjList := [[2,4],[1,3],[2,4],[1,3]]

}

func cloneGraph(node *Node) *Node {

	if node == nil {
		return nil
	}
	root := &Node{
		node.Val,
		node.Neighbors,
	}
	// we main following properties for items in the queue
	// - the nodes in the queue are copies of the original nodes
	// - but the neighbors of the nodes in the queue are still the original pointers
	queue := make([]*Node, 0)

	curr := root
	queue = append(queue, curr)

	// we maintain a map of visited nodes, same properties as the queue:
	// - the nodes in the queue are copies of the original nodes
	// - but the neighbors of the nodes in the queue are still the original pointers
	visited := make(map[int]*Node)
	visited[node.Val] = root

	for len(queue) > 0 {
		// pop one item
		curr = queue[0]
		queue = queue[1:]

		currNeighbors := make([]*Node, 0)
		// create copies of all its Neighbors (but the Neighbors of Neighbors
		// will have the original pointers for now)
		for _, neighbor := range curr.Neighbors {
			var neighborCopy *Node = nil

			// if we have already cloned this node before, we must use the same node
			if visited[neighbor.Val] != nil {
				neighborCopy = visited[neighbor.Val]
			} else {
				neighborCopy = &Node{
					neighbor.Val,
					neighbor.Neighbors, // these pointers are still the original pointers
				}
				visited[neighbor.Val] = neighborCopy
				// queue it, because we have to clone its neighbors
				queue = append(queue, neighborCopy)
			}
			// let the current node store copies of the neighbors
			currNeighbors = append(currNeighbors, neighborCopy)
		}
		curr.Neighbors = currNeighbors

	}

	return root
}

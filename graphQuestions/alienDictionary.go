package main

import (
	"bytes"
	"fmt"
)

func main7() {
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}

	fmt.Println(alienOrder(words))
}

func alienOrder(words []string) string {
	if len(words) == 0 {
		return ""
	}

	if len(words) == 1 {
		return words[0]
	}

	var buffer bytes.Buffer
	graph := make(map[string]deps)

	buildGraph(words, graph)

	explore(graph, &buffer)

	return buffer.String()
}

// topo sort
// note: we need to pass pointer to buffer, since the WriteString function needs a pointer receiver.
func explore(graph map[string]deps, buffer *bytes.Buffer) {
	var queue []string
	// using queue, doing bfs, first add init values into queue
	for k, v := range graph {
		if len(v) == 0 {
			queue = append(queue, k)
			delete(graph, k)
		}
	}

	for len(queue) > 0 {
		c := string(queue[0])
		queue = queue[1:]
		buffer.WriteString(c)

		for k1, v1 := range graph {
			if _, ok := v1[c]; ok {
				delete(v1, c)
			}

			if len(v1) == 0 {
				queue = append(queue, k1)
				delete(graph, k1)
			}
		}
	}

	if len(graph) > 0 {
		buffer.Reset()
	}
}

// use map[string]bool as set of string
type deps map[string]bool

func buildGraph(words []string, graph map[string]deps) {
	for i := 0; i < len(words)-1; i++ {
		addWords(words[i], words[i+1], graph)
	}
}

func addWords(s1 string, s2 string, graph map[string]deps) {
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		c1 := string(s1[i])
		c2 := string(s2[j])

		if _, ok := graph[c1]; !ok {
			graph[c1] = make(deps)
		}

		if _, ok := graph[c2]; !ok {
			graph[c2] = make(deps)
		}

		if c1 == c2 {
			i++
			j++
		} else {
			graph[c2][c1] = true
			break
		}
	}

	for i < len(s1) {
		c := string(s1[i])
		if _, ok := graph[c]; !ok {
			graph[c] = make(deps)
		}
		i++
	}

	for j < len(s2) {
		c := string(s2[j])
		if _, ok := graph[c]; !ok {
			graph[c] = make(deps)
		}
		j++
	}
}

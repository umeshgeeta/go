/*
 * Leetcode problem: 	https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/description/
 *
 * You are given an integer n. There is an undirected graph with n nodes, numbered from 0 to n - 1.
 * You are given a 2D integer array edges where edges[i] = [ai, bi] denotes that there exists an undirected edge
 * connecting nodes ai and bi.
 * Return the number of pairs of different nodes that are unreachable from each other.
 *
 * Author: Umesh Patil
 */

package main

import "fmt"

func countPairs(n int, edges [][]int) int64 {
	graph := make(map[int][]int)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	visited := make([]bool, n)
	islands := []int{}
	var dfs func(int) int

	dfs = func(node int) int {
		visited[node] = true
		connected := 1

		for _, child := range graph[node] {
			if !visited[child] {
				connected += dfs(child)
			}
		}

		return connected
	}

	for nodeIndex := 0; nodeIndex < n; nodeIndex++ {
		if !visited[nodeIndex] {
			islands = append(islands, dfs(nodeIndex))
		}
	}

	unconnectedPairs := 0
	accumulatedIslandNodes := 0
	for _, anIslandSize := range islands {
		unconnectedPairs += accumulatedIslandNodes * anIslandSize
		accumulatedIslandNodes += anIslandSize
	}

	return int64(unconnectedPairs)
}

func main() {
	var edges [][]int
	var cp int64

	edges = [][]int{{0, 1}, {0, 2}, {1, 2}}
	cp = countPairs(3, edges)
	fmt.Printf("Count of Unreachable Pairs of Nodes in an Undirected Graph: %d\n", cp)

	edges = [][]int{{0, 2}, {0, 5}, {2, 4}, {1, 6}, {5, 4}}
	cp = countPairs(7, edges)
	fmt.Printf("Count of Unreachable Pairs of Nodes in an Undirected Graph: %d\n", cp)

	edges = [][]int{{1, 0}, {3, 1}, {0, 4}, {2, 1}}
	cp = countPairs(5, edges)
	fmt.Printf("Count of Unreachable Pairs of Nodes in an Undirected Graph: %d\n", cp)
}

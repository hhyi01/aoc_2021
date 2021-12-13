package main

import (
	"fmt"
	"math"
	"strings"
)

func day12Part1(filePath string) {
	puzzleInput := readFileLines(filePath)
	adjList := setAdjList(puzzleInput)
	fmt.Println("Adjacency list:", adjList)
	findPaths("start", "end", adjList, "Part 2")
}

func findPaths(source string, dest string, adjList map[string][]string, part string)  {
	var pathsStack [][]string
	var path []string
	path = append(path, source)
	pathsStack = append(pathsStack, path)
	var allPaths []string
	for {
		path = pathsStack[len(pathsStack)-1]
		pathsStack = pathsStack[:len(pathsStack)-1]
		lastNode := path[len(path)-1]
		if lastNode == dest {
			stringPath := stringifyPath(path)
			allPaths = append(allPaths, stringPath)
		}
		adjNodes := adjList[lastNode]
		for _, node := range adjNodes {
			var visited bool
			if part == "Part 1" {
				visited = checkVisited(path, node)
			} else {
				visited = checkVisitedPart2(path, node, source, dest)
			}
			if !visited {
				nextPath := copyPath(path)
				nextPath = append(nextPath, node)
				pathsStack = append(pathsStack, nextPath)
			}
		}
		if len(pathsStack) == 0 {
			break
		}
	}
	fmt.Printf("%s - number of paths through this cave system: %d", part, len(allPaths))
}

func stringifyPath(path []string) string {
	return strings.Join(path[:],",")
}

func checkVisited(path []string, node string) bool {
	for _, v := range path {
		nodeUpper := strings.ToUpper(node)
		if v == node && nodeUpper != node {
			return true
		}
	}
	return false
}

func maxSmallCavesInPath(path []string) int {
	smallCavesCount := make(map[string]int)
	maxSmallCaves := math.MinInt
	for _, node := range path {
		nodeLower := strings.ToLower(node)
		if nodeLower == node {
			_, found := smallCavesCount[node]
			if found {
				smallCavesCount[node] += 1
			} else {
				smallCavesCount[node] = 1
			}
			if smallCavesCount[node] > maxSmallCaves {
				maxSmallCaves = smallCavesCount[node]
			}
		}
	}
	return maxSmallCaves
}

func checkVisitedPart2(path []string, node string, source string, dest string) bool {
	maxSmallCaves := maxSmallCavesInPath(path)
	visited := false
	for _, v := range path {
		nodeUpper := strings.ToUpper(node)
		if v == node {
			if nodeUpper != node {
				// small cave
				if maxSmallCaves > 1 {
					visited = true
				} else {
					if node == source || node == dest {
						visited = true
					}
				}
			}
		}
	}
	return visited
}

func copyPath(path []string) []string {
	var copiedPath []string
	for _, v := range path {
		copiedPath = append(copiedPath, v)
	}
	return copiedPath
}

func setAdjList(puzzleInput []string) map[string][]string {
	adjList := make(map[string][]string)
	for _, line := range puzzleInput {
		edge := strings.Split(line, "-")
		node1 := strings.TrimSpace(edge[0])
		node2 := strings.TrimSpace(edge[1])
		_, foundNode1 := adjList[node1]
		if !foundNode1 {
			adjList[node1] = append(adjList[node1], node2)
		} else {
			nodeInList := findNode(adjList[node1], node2)
			if !nodeInList {
				adjList[node1] = append(adjList[node1], node2)
			}
		}
		_, foundNode2 := adjList[node2]
		if !foundNode2 {
			adjList[node2] = append(adjList[node2], node1)
		} else {
			nodeInList := findNode(adjList[node2], node1)
			if !nodeInList {
				adjList[node2] = append(adjList[node2], node1)
			}
		}
	}
	return adjList
}

func findNode(edges []string, node string) bool {
	for _, v := range edges {
		if v == node {
			return true
		}
	}
	return false
}
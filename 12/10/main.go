package main

import (
	"../utils"
	"fmt"
	"sort"
	"strconv"
)

type Node struct {
	val       int
	count     int
	called    bool
	oneMore   *Node
	twoMore   *Node
	threeMore *Node
}

func (n *Node) dfs() (int, int) {
	if n.oneMore != nil {
		o, t := n.oneMore.dfs()
		return o + 1, t
	} else if n.twoMore != nil {
		o, t := n.twoMore.dfs()
		return o, t
	} else if n.threeMore != nil {
		o, t := n.threeMore.dfs()
		return o, t + 1
	}
	return 0, 0
}

func (n *Node) dfs2(num int) (int, bool) {
	count, found := 0, false
	if n.called {
		return n.count, true
	}
	if n.oneMore != nil {
		count, found = n.oneMore.dfs2(num)
	}
	if n.twoMore != nil {
		newCount, newFound := n.twoMore.dfs2(num)
		found = found || newFound
		if newFound {
			count += newCount
		}
	}
	if n.threeMore != nil {
		newCount, newFound := n.threeMore.dfs2(num)
		found = found || newFound
		if newFound {
			count += newCount
		}
	}

	n.called = true
	n.count = count

	if n.val == num {
		return 1, true
	}
	return count, found
}

func part1(inputs []string) int {
	arr := getArrOfInts(inputs)
	tree := generateTree(arr)
	oneDiff, threeDiff := tree.dfs()
	return oneDiff * threeDiff
}

func part2(inputs []string) {
	arr := getArrOfInts(inputs)
	tree := generateTree(arr)
	fmt.Println(tree.dfs2(arr[len(arr)-1]))
}

func getArrOfInts(inputs []string) []int {
	arr := make([]int, len(inputs)+1)
	max := 0
	for j, input := range inputs {
		i, _ := strconv.Atoi(input)
		if max < i {
			max = i
		}
		arr[j] = i
	}
	arr[len(inputs)] = max + 3
	sort.Ints(arr)
	return arr
}

func generateTree(nums []int) Node {
	numMapping := make(map[int]*Node)
	rootNode := Node{
		val:   0,
		count: -1,
	}
	numMapping[0] = &rootNode
	for _, num := range nums {
		newNode := Node{
			val:   num,
			count: -1,
		}
		numMapping[num] = &newNode
		if n, ok := numMapping[num-1]; ok {
			n.oneMore = &newNode
		}
		if n, ok := numMapping[num-2]; ok {
			n.twoMore = &newNode
		}
		if n, ok := numMapping[num-3]; ok {
			n.threeMore = &newNode
		}

	}
	return rootNode
}

func main() {
	inputs := utils.ReadTextFile("./input.txt")
	invalidValue := part1(inputs)
	fmt.Println(invalidValue)
	part2(inputs)

}

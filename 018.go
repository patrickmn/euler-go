package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	triangleStr = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`
)

type triangleNode struct {
	N  int
	T  *triangle
	L  *triangleNode
	R  *triangleNode
	ri int
	i  int
}

func (n *triangleNode) LinkDown() {
	if n.L != nil || n.R != nil { // Already linked
		return
	}
	lr := len(n.T.Rows[n.ri])
	nr := n.T.Rows[n.ri+1]
	lnr := len(nr)

	if lnr > lr { // Next row will have one more number
		n.L = n.T.GetNode(n.ri+1, n.i, nr[n.i])
		n.R = n.T.GetNode(n.ri+1, n.i+1, nr[n.i+1])
		if n.ri+1 != len(n.T.Rows)-1 { // Not last row
			n.L.LinkDown()
			n.R.LinkDown()
		}
	} else { // Next row will have one less number
		if n.i > 0 {
			n.L = n.T.GetNode(n.ri+1, n.i-1, nr[n.i-1])
			if lnr > 1 {
				n.L.LinkDown()
			}
		}
		if n.i < lnr {
			n.R = n.T.GetNode(n.ri+1, n.i, nr[n.i])
			if lnr > 1 {
				n.R.LinkDown()
			}
		}
	}
}

func (n *triangleNode) SumDownOne() {
	r := n.T.Rows[n.ri]
	lr := len(r)
	nr := n.T.Rows[n.ri+1]
	lnr := len(nr)

	if lnr > lr { // Next row will have one more number
		// TODO
	} else { // Next row will have one less number
		if lr-1 == n.i { // Last node doesn't have a node next to it
			return
		}
		on := n.T.nmap[n.ri][n.i+1]
		nn := n.T.nmap[n.ri+1][n.i]
		if n.N > on.N {
			nn.N += n.N
		} else {
			nn.N += on.N
		}
	}
}

type triangle struct {
	Rows  [][]int
	Nodes []*triangleNode
	nmap  []map[int]*triangleNode
}

func (t *triangle) ParseRows(s string) {
	strRows := strings.Split(s, "\n")
	t.Rows = make([][]int, len(strRows))
	for i, v := range strRows {
		numStrs := strings.Split(v, " ")
		t.Rows[i] = make([]int, len(numStrs))
		for oi, ov := range numStrs {
			t.Rows[i][oi], _ = strconv.Atoi(ov)
		}
	}
}

func (t *triangle) ReverseRows() {
	for i, j := 0, len(t.Rows)-1; i < j; i, j = i+1, j-1 {
		t.Rows[i], t.Rows[j] = t.Rows[j], t.Rows[i]
	}
}

func (t *triangle) GetNode(ri, i, n int) *triangleNode {
	node, found := t.nmap[ri][i]
	if found {
		return node
	}
	node = &triangleNode{
		N:  n,
		T:  t,
		ri: ri,
		i:  i,
	}
	t.nmap[ri][i] = node
	return node
}

func (t *triangle) GenerateNodes() {
	first := t.Rows[0]
	t.Nodes = make([]*triangleNode, len(first))
	t.nmap = make([]map[int]*triangleNode, len(t.Rows))
	for i := range t.nmap {
		t.nmap[i] = make(map[int]*triangleNode)
	}
	for i, v := range first {
		n := t.GetNode(0, i, v)
		t.Nodes[i] = n
		n.LinkDown()
	}
}

func showNode(n *triangleNode) {
	fmt.Printf("\n%d", n.N)
	if n.L != nil {
		fmt.Printf(" L: %d", n.L.N)
		showNode(n.L)
	}
	if n.R != nil {
		fmt.Printf(" R: %d", n.R.N)
		showNode(n.R)
	}
}

func Euler018() Result {
	t := new(triangle)

	// Parse triangle
	t.ParseRows(triangleStr)

	// Turn it upside-down
	t.ReverseRows()

	// Generate L+R nodes
	t.GenerateNodes()

	// Find the largest sum
	for _, v := range t.nmap[:len(t.nmap)-1] {
		for _, ov := range v {
			ov.SumDownOne()
		}
	}
	return t.nmap[len(t.nmap)-1][0].N
}

package main

import "fmt"

type RangeRing struct {
    nodes []*Node
}

// NewRangeRing creates an empty range-based ring.
func NewRangeRing() *RangeRing {
    return &RangeRing{nodes: make([]*Node, 0)}
}

// AddNode adds a node to the ring.
func (rr *RangeRing) AddNode(node *Node) {
    rr.nodes = append(rr.nodes, node)
    fmt.Printf("Added node %d at ring index %d\n", node.ID, len(rr.nodes)-1)
}

// GetNode returns the node responsible for the key based on hash range.
func (rr *RangeRing) GetNode(key string) *Node {
    h := hashKey(key)
    n := uint32(len(rr.nodes))
    rangeSize := uint64(^uint32(0))/uint64(n) + 1
    idx := int(uint64(h) / rangeSize)
    return rr.nodes[idx]
}

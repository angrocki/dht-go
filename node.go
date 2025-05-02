package main

import "fmt"

// NodeID is a simple integer identifier for each node.
type NodeID uint32

// Node simulates a DHT peer with its own storage and inbox.
type Node struct {
    ID    NodeID
    inbox chan Message
    store map[string]string
}

// NewNode creates a node with the given ID.
func NewNode(id NodeID) *Node {
    return &Node{
        ID:    id,
        inbox: make(chan Message, 10),
        store: make(map[string]string),
    }
}

// Run starts the node's event loop. Call in a goroutine.
func (n *Node) Run() {
    for msg := range n.inbox {
        switch msg.Type {
        case STORE:
            n.store[msg.Key] = msg.Value
            fmt.Printf("[Node %d] Stored '%s' -> '%s'\n", n.ID, msg.Key, msg.Value)

        case GET:
            val, ok := n.store[msg.Key]
            if ok {
                msg.ReplyChan <- val
            } else {
                msg.ReplyChan <- ""
            }
        }
    }
}

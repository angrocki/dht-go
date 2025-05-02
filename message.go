package main

// MessageType enumerates the kinds of DHT operations.
type MessageType int

const (
    STORE MessageType = 0
    GET MessageType = 1
)

// Message represents a request or reply in the DHT.
type Message struct {
    Type      MessageType
    Key       string
    Value     string        // used for STORE
    ReplyChan chan string   // used for GET replies
}

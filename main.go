package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"
)

func main() {
    num := flag.Int("nodes", 3, "number of nodes in the DHT ring")
    flag.Parse()

    ring := NewRangeRing()
    for i := 0; i < *num; i++ {
        node := NewNode(NodeID(i))
        ring.AddNode(node)
        go node.Run()
    }

    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Simple DHT CLI (range-based). Commands:")
    fmt.Println("  store <key> <value>")
    fmt.Println("  get <key>")
    fmt.Println("  exit")

    for {
        fmt.Print("> ")
        line, _ := reader.ReadString('\n')
        parts := strings.Fields(strings.TrimSpace(line))
        if len(parts) == 0 {
            continue
        }
        switch parts[0] {
        case "store":
            if len(parts) < 3 {
                fmt.Println("Usage: store <key> <value>")
                continue
            }
            key := parts[1]
            val := strings.Join(parts[2:], " ")
            node := ring.GetNode(key)
            node.inbox <- Message{Type: STORE, Key: key, Value: val}

        case "get":
            if len(parts) != 2 {
                fmt.Println("Usage: get <key>")
                continue
            }
            key := parts[1]
            reply := make(chan string)
            node := ring.GetNode(key)
            node.inbox <- Message{Type: GET, Key: key, ReplyChan: reply}
            val := <-reply
            if val == "" {
                fmt.Printf("Key '%s' not found\n", key)
            } else {
                fmt.Printf("Got '%s' -> '%s'\n", key, val)
            }

        case "exit":
            fmt.Println("Exiting.")
            return

        default:
            fmt.Println("Unknown command; use store, get, or exit.")
        }
    }
}

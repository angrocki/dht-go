# Distributed Hash Tables

## Implementation

We were not able to meet our initial goal of implementing a decentralized DHT. Instead we pivoted to centralized DHT where we kept track of all the nodes running on different goroutines.

Each node in the system was represented by a mock server, implemented using goroutines to simulate concurrency and inter-node communication. We created added a CLI to be able to interact with our system via the terminal.

## Challenges

An interesting challenge we faced was setting up consistent hashing. Although we explored several different methods, we ultimately used range-based partitioning because it was the easiest to implement and reasonable since in our case we were not dynamically adding and removing nodes. However, this choice came with trade-offs in terms of flexibility and scalability. We also encountered difficulties with implementing fault tolerance, particularly in handling node failures and ensuring key replication.

## How to Run

1. Clone the repository on your computer with `go` already installed
2. Inside the repository, run: `go build`
3. Run: `./dht-go`

package main

import (
    "hash/fnv"
)

// hashKey computes a 32-bit FNV-1a hash of the given string.
func hashKey(key string) uint32 {
    h := fnv.New32a()
    h.Write([]byte(key))
    return h.Sum32()
}

package main

import (
	"sync"
)

type User struct {
	Id string
	mu sync.Mutex
}

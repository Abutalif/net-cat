package models

import "sync"

type User struct {
	Name     string
	IsActive bool
}

// reference to create a usecase
type Server struct {
	Users sync.Map
	host  string
	port  string
}

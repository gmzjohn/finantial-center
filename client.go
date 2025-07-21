package main

type Client struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Phone    string `json:"phone"`
}
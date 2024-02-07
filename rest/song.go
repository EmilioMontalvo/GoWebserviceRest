package main

type Song struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Path  string `json:"path"`
	Plays int    `json:"plays"`
}

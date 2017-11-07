package main

type Status struct {
    Status      string `json:"status"`
}

type Object struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Completed  bool   `json:"completed"`
}

type Objects []Object

package main

import (
	"fmt"
	"log"
	"os"
)

type User struct {
	Id int
	Name, Location string
}

type Player struct {
	User
	GameId int
}

type Job struct {
	Command string
	*log.Logger
}

func main() {
	p := Player{}
	p.Id = 1913
	p.Name = "Bob"
	p.Location = "North Pole"
	p.GameId = 8350825382
	fmt.Printf("%+v", p)

	q := Player{
		User{Id: 1914, Name: "Bobert", Location: "South Pole"},
		421949210492,
	}

	fmt.Printf("%+v\n", q)

	job := &Job{"demo", log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Print("starting now...")
}
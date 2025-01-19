package main

import "fmt"

// iterrator
type Movie struct {
	name  string
	genre string
}

type MovieIterator interface {
	hasNext() bool
	next() Movie
}

type MovieCollection interface {
	AddMovie(movie Movie)
	CreateIterator() MovieIterator
}

func main() {

	action := &ActionMovie{}
	action.AddMovie(Movie{name: "piku", genre: "emtional"})
	comedy := &ComedyMovie{}
	comedy.AddMovie(Movie{name: "pk", genre: "funny"})

	actionIterator := action.CreateIterator()
	comedyIterator := comedy.CreateIterator()

	for actionIterator.hasNext() {
		fmt.Println(" this is action movie", actionIterator.next().name)
	}
	for comedyIterator.hasNext() {
		fmt.Println(" this is com movie", comedyIterator.next())
	}

}

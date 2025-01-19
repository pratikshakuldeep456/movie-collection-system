package main

type ComedyMovie struct {
	movies []Movie
}

func (a *ComedyMovie) AddMovie(m Movie) {
	a.movies = append(a.movies, m)
}

type ComedyMovieIterator struct {
	movie []Movie
	index int
}

func (a *ComedyMovie) CreateIterator() MovieIterator {
	return &ComedyMovieIterator{movie: a.movies, index: 0}

}

func (a *ComedyMovieIterator) hasNext() bool {
	return a.index < len(a.movie)
}

func (a *ComedyMovieIterator) next() Movie {
	if a.hasNext() {
		ans := a.movie[a.index]
		a.index++
		return ans
	}
	return Movie{}
}

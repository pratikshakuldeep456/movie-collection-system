package main

type ActionMovie struct {
	movies []Movie
}

func (a *ActionMovie) AddMovie(m Movie) {
	a.movies = append(a.movies, m)
}

type ActionMovieIterator struct {
	movie []Movie
	index int
}

func (a *ActionMovie) CreateIterator() MovieIterator {
	return &ActionMovieIterator{movie: a.movies, index: 0}

}
func (a *ActionMovieIterator) hasNext() bool {
	return a.index < len(a.movie)
}

func (a *ActionMovieIterator) next() Movie {
	if a.hasNext() {
		ans := a.movie[a.index]
		a.index++
		return ans
	}
	return Movie{}
}

package api

import (
	"github.com/moosebot/plex-rek/server/internal/config"
)

// TMDBSearchResponse defines the results of the multi search of TheMovieDb
type TMDBSearchResponse struct {
	Page         int `json:"page"`
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Results      []struct {
		OriginalName     string   `json:"original_name,omitempty"`
		ID               int      `json:"id"`
		MediaType        string   `json:"media_type"`
		Name             string   `json:"name,omitempty"`
		VoteCount        int      `json:"vote_count"`
		VoteAverage      float64  `json:"vote_average"`
		PosterPath       string   `json:"poster_path"`
		FirstAirDate     string   `json:"first_air_date,omitempty"`
		Popularity       float64  `json:"popularity"`
		GenreIds         []int    `json:"genre_ids"`
		OriginalLanguage string   `json:"original_language"`
		BackdropPath     string   `json:"backdrop_path"`
		Overview         string   `json:"overview"`
		OriginCountry    []string `json:"origin_country,omitempty"`
		Video            bool     `json:"video,omitempty"`
		Title            string   `json:"title,omitempty"`
		OriginalTitle    string   `json:"original_title,omitempty"`
		Adult            bool     `json:"adult,omitempty"`
		ReleaseDate      string   `json:"release_date,omitempty"`
	} `json:"results"`
}

// SearchResults defines the entire object returning to the client
type SearchResults struct {
	Movies  []Movie
	TvShows []TvShow
}

// Movie defines the movie object returning to the client
type Movie struct {
	Title             string
	Overview          string
	PosterImagePath   string
	BackdropImagePath string
	ReleaseDate       string
}

// TvShow defines the tvshow object returning to the client
type TvShow struct {
	Name              string
	Overview          string
	PosterImagePath   string
	BackdropImagePath string
	FirstAirDate      string
}

// Search will perform a search on TheMovieDb's entire repository, including movies, tv shows, and people
func Search(query string) SearchResults {
	var config = config.GetConfig()

	var movies = make([]Movie, 0)
	movies = append(movies, Movie{Title: "Animal Kingdom", Overview: "Following the death of his mother, J finds himself living with his estranged family, under the watchful eye of his doting grandmother, Smurf, mother to the Cody boys. J quickly comes to believe that he is a player in this world. But, as he soon discovers, this world is far larger and more menacing than he could ever imagine. J finds himself at the center of a cold-blooded revenge plot that turns the family upside down.", PosterImagePath: config.TheMovieDb.ImageBaseURL + "/zhj8YPQKuRev5N3KoHacsPnF4mB.jpg", BackdropImagePath: config.TheMovieDb.ImageBaseURL + "/1V8C68OvU7IDjZO9GtfwfmHa8X7.jpg", ReleaseDate: "2010-06-03"})

	var tvshows = make([]TvShow, 0)
	tvshows = append(tvshows, TvShow{Name: "Animal Kingdom", Overview: "The series centers on 17-year-old Joshua \"J\" Cody, who moves in with his freewheeling relatives in their Southern California beach town after his mother dies of a heroin overdose. Headed by boot-tough matriarch Janine \"Smurf\" Cody and her right-hand Baz, who runs the business and calls the shots, the clan also consists of Pope, the oldest and most dangerous of the Cody boys; Craig, the tough and fearless middle son; and Deran, the troubled, suspicious \"baby\" of the family.", PosterImagePath: config.TheMovieDb.ImageBaseURL + "/yfeUKQPfhcemXBKKswq1KtxXFsQ.jpg", BackdropImagePath: config.TheMovieDb.ImageBaseURL + "/rcS7cbgRUS6E5Qnolzm05lF02eM.jpg", FirstAirDate: "2016-06-14"})

	return SearchResults{Movies: movies, TvShows: tvshows}
}

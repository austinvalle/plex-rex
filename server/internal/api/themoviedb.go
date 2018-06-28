package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/moosebot/plex-rex/server/internal/config"
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

var httpClient = &http.Client{Timeout: 10 * time.Second}

// Search will perform a search on TheMovieDb's entire repository, including movies, tv shows, and people
func Search(query string) SearchResults {
	var config = config.GetConfig()
	req, err := http.NewRequest("GET", config.TheMovieDb.BaseURL+"/search/multi", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("api_key", config.TheMovieDb.APIKey)
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()

	r, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var searchResponse TMDBSearchResponse
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&searchResponse)

	var searchResults = SearchResults{}
	for index := 0; index < searchResponse.TotalResults; index++ {
		var result = searchResponse.Results[index]
		if result.MediaType == "movie" {
			searchResults.Movies = append(searchResults.Movies, Movie{
				Title:             result.Title,
				Overview:          result.Overview,
				PosterImagePath:   config.TheMovieDb.ImageBaseURL + result.PosterPath,
				BackdropImagePath: config.TheMovieDb.ImageBaseURL + result.BackdropPath,
				ReleaseDate:       result.ReleaseDate,
			})
			continue
		}

		if result.MediaType == "tv" {
			searchResults.TvShows = append(searchResults.TvShows, TvShow{
				Name:              result.Name,
				Overview:          result.Overview,
				PosterImagePath:   config.TheMovieDb.ImageBaseURL + result.PosterPath,
				BackdropImagePath: config.TheMovieDb.ImageBaseURL + result.BackdropPath,
				FirstAirDate:      result.FirstAirDate,
			})
			continue
		}
	}

	return searchResults
}

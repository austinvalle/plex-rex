using System.Collections.Generic;
using PlexRexAPI.Models;

namespace PlexRexAPI.Services
{
    /*
     * TheMovieDB - https://developers.themoviedb.org/3
     */
    public class TmdbApiService : ICallTmdbApi
    {
        /*
         * https://api.themoviedb.org/3/search/multi?query=<query>&api_key=<api-key>
         */
        public MultiSearchResults MultiSearch()
        {
            return new MultiSearchResults
            {
                Movies = new List<Movie>
                {
                    new Movie
                    {
                        Title = "Test"
                    }
                }
            };
        }
    }
}

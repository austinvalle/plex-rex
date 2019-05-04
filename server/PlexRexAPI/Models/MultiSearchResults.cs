using System.Collections.Generic;

namespace PlexRexAPI.Models
{
    public class MultiSearchResults
    {
        public IList<Movie> Movies { get; set; }
        public IList<TvShow> TvShows { get; set; }
    }
}

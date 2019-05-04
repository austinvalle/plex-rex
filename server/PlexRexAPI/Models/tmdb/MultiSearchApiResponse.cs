using System;

namespace PlexRexAPI.Models.tmdb
{
    /*
     * https://api.themoviedb.org/3/search/multi?query=Animal%20Kingdom&api_key=ABC123
     */
    public class MultiSearchApiResponse
    {
        public int Page { get; set; }
        public int Total_results { get; set; }
        public int Total_pages { get; set; }
        public MultiSearchApiResult[] Results { get; set; }
    }

    public class MultiSearchApiResult
    {
        public string Original_name { get; set; }
        public int Id { get; set; }
        public string Media_type { get; set; }
        public string Name { get; set; }
        public int Vote_count { get; set; }
        public float Vote_average { get; set; }
        public string Poster_path { get; set; }
        public DateTime? First_air_date { get; set; }
        public float Popularity { get; set; }
        public int?[] Genre_ids { get; set; }
        public string Original_language { get; set; }
        public string Backdrop_path { get; set; }
        public string Overview { get; set; }
        public string[] Origin_country { get; set; }
        public bool Video { get; set; }
        public string Title { get; set; }
        public string Original_title { get; set; }
        public bool Adult { get; set; }
        public DateTime? Release_date { get; set; }
    }

}

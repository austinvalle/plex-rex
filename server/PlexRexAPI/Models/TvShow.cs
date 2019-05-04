using System;

namespace PlexRexAPI.Models
{
    public class TvShow : Media
    {
        public string Name { get; set; }
        public DateTime? FirstAirDate { get; set; }
    }
}

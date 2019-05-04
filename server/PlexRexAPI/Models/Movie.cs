using System;

namespace PlexRexAPI.Models
{
    public class Movie : Media
    {
        public string Title { get; set; }
        public DateTime? ReleaseDate { get; set; }
    }
}

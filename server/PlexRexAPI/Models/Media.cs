using System;

namespace PlexRexAPI.Models
{
    public abstract class Media
    {
        private readonly Uri TMDB_BASE_URI = new Uri("https://image.tmdb.org/t/p/w500/");
        private Uri _posterImagePath;
        private Uri _backdropImagePath;

        public string Overview { get; set; }

        public Uri PosterImagePath {
            get
            {
                return _posterImagePath; 
            }
            set
            {
                _posterImagePath = new Uri(TMDB_BASE_URI, value);
            }
        }

        public Uri BackdropImagePath
        {
            get
            {
                return _backdropImagePath;
            }
            set
            {
                _backdropImagePath = new Uri(TMDB_BASE_URI, value);
            }
        }
    }
}

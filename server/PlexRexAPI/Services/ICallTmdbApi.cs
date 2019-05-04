using PlexRexAPI.Models;

namespace PlexRexAPI.Services
{
    public interface ICallTmdbApi
    {
        MultiSearchResults MultiSearch();
    }
}

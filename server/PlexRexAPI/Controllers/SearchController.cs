using Microsoft.AspNetCore.Mvc;
using PlexRexAPI.Models;
using PlexRexAPI.Services;

namespace PlexRexAPI.Controllers
{
    [Route("api/search")]
    [ApiController]
    public class SearchController : ControllerBase
    {
        private readonly ICallTmdbApi _tmdbApiService;

        public SearchController(ICallTmdbApi tmdbApiService)
        {
            _tmdbApiService = tmdbApiService;
        }

        [HttpGet]
        public ActionResult<MultiSearchResults> MultiSearch()
        {
            return _tmdbApiService.MultiSearch();
        }
    }
}
using System;
using System.Threading.Tasks;
using System.Net.Http;
using System.Threading;

namespace PitWarner
{
    public interface IHttpService
    {
        Task<HttpResponseMessage> Get(string url, CancellationTokenSource token = null);
        Task<HttpResponseMessage> Post(string url, object postData, CancellationTokenSource token = null);
        Task<HttpResponseMessage> Head(string url, CancellationTokenSource token = null);
    }
}


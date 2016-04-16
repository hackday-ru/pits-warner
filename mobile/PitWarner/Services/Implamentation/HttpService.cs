using System;
using System.Net;
using System.Net.Http;
using System.Diagnostics;

namespace PitWarner
{
    public class HttpService : IHttpService
    {
        public HttpService()
        {
        }

        #region IHttpService implementation

        public async System.Threading.Tasks.Task<System.Net.Http.HttpResponseMessage> Get(string url, System.Threading.CancellationTokenSource token = null)
        {
            try
            {
                var cookieContainer = new CookieContainer();
                using (var handler = new HttpClientHandler { CookieContainer =  cookieContainer})
                {
                    using (var client = new HttpClient(handler))
                    {
                        if (!string.IsNullOrWhiteSpace(Variables.BASE_HOST))
                            client.BaseAddress = new Uri(Variables.BASE_HOST);

                        using (var message = new HttpRequestMessage())
                        {
                            message.Method = HttpMethod.Get;
                            message.RequestUri = new Uri (url);
                            Debug.WriteLine($"PirWarner HTTP -> Отправка запроса {url}");
                            return await client.SendAsync (message).ConfigureAwait (false);
                        }
                    }
                }
            }
            catch (WebException ex)
            {
                Debug.WriteLine("PirWarner -> Проблемы с GET запросом: " + ex);
                return null;
            }
        }

        public System.Threading.Tasks.Task<System.Net.Http.HttpResponseMessage> Post(string url, object postData, System.Threading.CancellationTokenSource token = null)
        {
            throw new NotImplementedException();
        }

        public System.Threading.Tasks.Task<System.Net.Http.HttpResponseMessage> Head(string url, System.Threading.CancellationTokenSource token = null)
        {
            throw new NotImplementedException();
        }

        #endregion
    }
}


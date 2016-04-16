using System;
using Newtonsoft.Json;
using System.Collections.Generic;
using System.Net;

namespace PitWarner
{
    public class ApiService : IApiService
    {
        readonly IHttpService _httpService;

        public ApiService(IHttpService httpService)
        {
            _httpService = httpService;
        }

        #region IApiService implementation

        public async System.Threading.Tasks.Task<System.Collections.Generic.List<PitModel>> GetPits(double lat, double lon, int radius, System.Threading.CancellationTokenSource token)
        {
            try
            {
                var url = Variables.BASE_HOST + string.Format ("/raw?lat={0}&lng={1}&radius={2}", lat, lon, radius);

                var response = await _httpService.Get(url, token);
                if (response != null) {
                    var data = await response.Content.ReadAsStringAsync ();
                    var pits = JsonConvert.DeserializeObject<List<PitModel>> (data);
                    return pits;
                } 

                return new List<PitModel>();
            }
            catch (WebException ex)
            {
                System.Diagnostics.Debug.WriteLine("Запрос не работает");
                return new List<PitModel>();
            }
            catch (JsonReaderException ex)
            {
                System.Diagnostics.Debug.WriteLine("Не могу прочитать json");
                return new List<PitModel>();
            }
            catch (Exception ex)
            {
                System.Diagnostics.Debug.WriteLine("Просто Exception");
                return new List<PitModel>();
            }
        }

        #endregion
    }
}


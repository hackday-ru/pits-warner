using System;
using System.Threading.Tasks;
using System.Collections.Generic;
using System.Threading;

namespace PitWarner
{
    public interface IApiService
    {
        Task<List<PitModel>> GetPits(double lat, double lon, int radius, CancellationTokenSource token);
    }
}


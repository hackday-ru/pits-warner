using System;
using System.Threading.Tasks;
using System.Collections.Generic;
using System.Threading;

namespace PitWarner
{
    public interface IApiService
    {
        Task<List<PitModel>> GetPits(CancellationTokenSource token);
    }
}


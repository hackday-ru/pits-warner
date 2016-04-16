using System;
using System.Collections.Generic;

namespace PitWarner
{
    public interface IGeoService
    {
        List<PitModel> GetNearPits(double lat, double lon);
    }
}


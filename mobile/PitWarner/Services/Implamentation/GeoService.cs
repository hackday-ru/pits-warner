using System;

namespace PitWarner
{
    public class GeoService : IGeoService
    {
        public GeoService()
        {
        }

        #region IGeoService implementation

        public System.Collections.Generic.List<PitModel> GetNearPits(double lat, double lon)
        {
            throw new NotImplementedException();
        }

        #endregion
    }
}


using System;
using System.Collections.Generic;
using System.Linq;

namespace PitWarner
{
    public class GeoCalc
    {
        public static double GetDistanceBetween2Points(PitModel coordinate1, PitModel coordinate2)
        {
            var R = 6372795;
            var dLat = ToRad(coordinate2.Lat - coordinate1.Lat);
            var dLon = ToRad(coordinate2.Lon - coordinate1.Lon);
            var lat1 = ToRad(coordinate1.Lat);
            var lat2 = ToRad(coordinate2.Lat);
            var a = Math.Sin(dLat / 2) * Math.Sin(dLat / 2) + Math.Sin(dLon / 2) * Math.Sin(dLon / 2) * Math.Cos(lat1) * Math.Cos(lat2);
            var c = 2 * Math.Atan2(Math.Sqrt(a), Math.Sqrt(1 - a));
            var d = R * c;
            return d;
        }

        private static double ToRad(double grad)
        {
            return grad * Math.PI / 180;
        }

        public static bool InPoly(IList<PitModel> pits, PitModel pt)
        {
            double maxLat = pits.Max(p => p.Lat);
            double minLat = pits.Min(p => p.Lat);
            double maxLon = pits.Max(p => p.Lon);
            double minLon = pits.Min(p => p.Lon);

            var betweenLat = (minLat <= pt.Lat) && (maxLat >= pt.Lat);
            var betweenLon = (minLon <= pt.Lon) && (maxLon >= pt.Lon);

            return betweenLat && betweenLon;
        }

//        public static Dictionary<int, double> CalculateDistances(List<PitModel> coordinatesFromDB, PitModel currentCoordinate)
//        {
//            return coordinatesFromDB.ToDictionary(coordinateItem => coordinateItem.Id - 1, coordinateItem => GetDistanceBetween2Points(currentCoordinate, coordinateItem));
//        }
    }
}


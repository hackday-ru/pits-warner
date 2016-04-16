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
            var dLat = ToRad(coordinate2.lat - coordinate1.lat);
            var dLon = ToRad(coordinate2.lng - coordinate1.lng);
            var lat1 = ToRad(coordinate1.lat);
            var lat2 = ToRad(coordinate2.lat);
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
            double maxLat = pits.Max(p => p.lat);
            double minLat = pits.Min(p => p.lat);
            double maxLon = pits.Max(p => p.lng);
            double minLon = pits.Min(p => p.lng);

            var betweenLat = (minLat <= pt.lat) && (maxLat >= pt.lat);
            var betweenLon = (minLon <= pt.lng) && (maxLon >= pt.lng);

            return betweenLat && betweenLon;
        }

//        public static decimal MakeRectangle(double lat, double lng)
//        {
//
//
////            decimal oneMetrInDegrees = 1 / 111000;
////            return oneMetrInDegrees;
//        }

//        public static Dictionary<int, double> CalculateDistances(List<PitModel> coordinatesFromDB, PitModel currentCoordinate)
//        {
//            return coordinatesFromDB.ToDictionary(coordinateItem => coordinateItem.Id - 1, coordinateItem => GetDistanceBetween2Points(currentCoordinate, coordinateItem));
//        }
    }
}


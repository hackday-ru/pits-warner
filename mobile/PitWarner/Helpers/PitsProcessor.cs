using System;
using System.Collections.Generic;
using System.Linq;

namespace PitWarner
{
    public class PitsProcessor
    {
        const float L = 10;

        private PitModel[] _points;

        public PitsProcessor(IEnumerable<PitModel> allpoints)
        {
            _points = allpoints.ToArray();
        }

        public PitModel[] GetPitsAhead() {
            return new PitModel[0];
        }

        public static Poligon GetPoly()
        {
//            PitModel a, b, c, d;
            throw new NotImplementedException();

        }

        #region

        static float leftTurn(PitModel a, PitModel b, PitModel c) {
            return a.X * (b.Y - c.Y) - a.Y * (b.X - c.X) + (b.X * c.Y - b.Y * c.X);
        }

        static bool isInside(Poligon pl, PitModel c) {
            for (int i = 1; i < pl.points.Count; i++)
            {
                var a = pl.points[i - 1];
                var b = pl.points[i];
                if (leftTurn(a, b, c) < 0)
                {
                    return false;
                }
            }

            var a2 = pl.points[pl.points.Count() - 1];
            var b2 = pl.points[0];

            if (leftTurn(a2, b2, c) < 0)
            {
                return false;
            }
            return true;
        }

        private PitModel[] getPointsInPoly(Poligon p) {
            float minX = int.MaxValue;
            float maxX = int.MinValue;
            float minY = int.MaxValue;
            float maxY = int.MinValue;
            foreach (var point in _points)
            {
                minX = Math.Min(minX, point.X);
                maxX = Math.Max(maxX, point.X);
                minY = Math.Min(minY, point.Y);
                minY = Math.Min(minY, point.Y);
            }

            List<PitModel> pointsInBBox = new List<PitModel>();

            foreach (var point in _points)
            {
                if (point.X > maxX)
                    continue;
                if (point.X < minX)
                    continue;
                if (point.Y > maxY)
                    continue;
                if (point.Y < minY)
                    continue;

                pointsInBBox.Add(point);
            }

            return pointsInBBox.Where(pnt => isInside(p, pnt)).ToArray(); 

            //var pointsInBox = Allpoints.Where()
        }

        #endregion
    }

    public struct Rect {
        public PitModel p;
        public PitModel q;
    }

    public class Poligon {
        public List<PitModel> points;
    }
}


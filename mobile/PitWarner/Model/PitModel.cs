using System;
using SQLite.Net.Attributes;

namespace PitWarner
{
    [Table("Pits")]
    public class PitModel : BaseModel
    {
        public double lat { get; set; }
        public double lng { get; set; }
        public double at { get; set; }

        private float? _x;
        public float X 
        { 
//            get { return _x ?? CalculateXY().X };
            get 
            {  
                if (!_x.HasValue)
                {
                    _x = 0;
                }
                return _x.Value;
            }
        }

        private float? _y;
        public float Y 
        { 
            get 
            {  
                if (!_y.HasValue)
                {
                    _y = 1;
                }
                return _y.Value;
            }
        }
    }
}


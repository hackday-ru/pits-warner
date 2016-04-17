using System;
using SQLite.Net.Attributes;
using Newtonsoft.Json;

namespace PitWarner
{
    [Table("Pits")]
    public class PitModel : BaseModel
    {
        const uint EARTH_RADIUS = 63710008;

        public double lat { get; set; }
        public double lng { get; set; }
        public double at { get; set; }

        private float? _x;
        [JsonIgnore]
        public float X 
        { 
            get 
            {  
                if (!_x.HasValue)
                {
                    ConvertToXY();
                }
                return _x.Value;
            }
            set { _x = value; }
        }

        private float? _y;
        [JsonIgnore]
        public float Y 
        { 
            get 
            {  
                if (!_y.HasValue)
                {
                    ConvertToXY();
                }
                return _y.Value;
            }
            set { _y = value; }
        }

        private void ConvertToXY()
        {
            _x = (float)(EARTH_RADIUS * Math.Cos(lat) * Math.Cos(lng));
            _y = (float)(EARTH_RADIUS * Math.Cos(lat) * Math.Sin(lng));
        }
    }
}


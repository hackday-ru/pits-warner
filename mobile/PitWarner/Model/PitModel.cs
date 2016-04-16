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
    }
}


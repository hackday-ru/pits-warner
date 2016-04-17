using System;
using System.Collections.Generic;

namespace PitWarner
{
    public interface IDataBaseService
    {
        void SaveData(IList<PitModel> pits);
        List<PitModel> ReadData();
    }
}


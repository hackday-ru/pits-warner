using System;
using System.Threading;

namespace PitWarner
{
    public interface IBackgroundService
    {
        void Start();
        void Stop();
    }
}


using System;
using MvvmCross.Core.ViewModels;
using System.Threading.Tasks;
using System.Windows.Input;
using System.Threading;
using MvvmCross.Platform;
using System.Collections.Generic;
using System.Linq;
using MvvmCross.Plugins.Location;
using System.Diagnostics;
using Chance.MvvmCross.Plugins.UserInteraction;
using MvvmCross.Plugins.Accelerometer;

namespace PitWarner.ViewModels
{
    public class MainViewModel : BaseViewModel
    {
        const int DISTANCE_LIMIT = 100;

        readonly IApiService _apiService;
        readonly IDataBaseService _dbService;
        readonly IMvxAccelerometer _accelerometer;
        private readonly IMvxLocationWatcher _watcher;
        private PitsProcessor _processor;

        MvxAccelerometerReading _accValue;

        List<PitModel> _pits;
        MvxGeoLocation _lastLocation;

        public MainViewModel(IApiService apiService, IDataBaseService dbService, IMvxLocationWatcher watcher, IMvxAccelerometer accelerometer)
        {
            _accelerometer = accelerometer;
            _apiService = apiService;
            _dbService = dbService;
            _watcher = watcher;
        }

        public async override void Start()
        {
            base.Start();

            updatePits();
            _watcher.Start(new MvxLocationOptions(), OnLocation, OnError);

            _accelerometer.Start();
            _accelerometer.ReadingAvailable += (sender, e) => {
                _accValue = e.Value;
                X = _accValue.X;
                Y = _accValue.Y;
                Z = _accValue.Z;
            };

        }

        private double _lat;
        public double Lat
        {
            get
            { 
                return _lat; 
            }
            set
            {
                _lat = value; 
                RaisePropertyChanged(() => Lat);
            }
        }

        private double _lon;
        public double Lon
        {
            get
            { 
                return _lon; 
            }
            set
            { 
                _lon = value; 
                RaisePropertyChanged(() => Lon);
            }
        }

        private double _x;
        public double X
        {
            get
            { 
                return _x; 
            }
            set
            { 
                _x = value; 
                RaisePropertyChanged(() => X);
            }
        }

        private double _y;
        public double Y
        {
            get
            { 
                return _y; 
            }
            set
            { 
                _y = value; 
                RaisePropertyChanged(() => Y);
            }
        }

        private double _z;
        public double Z
        {
            get
            { 
                return _z; 
            }
            set
            { 
                _z = value; 
                RaisePropertyChanged(() => Z);
            }
        }

        private int _countOfDots;
        public int CountOfDots
        {
            get
            { 
                return _countOfDots;
            }
            set
            { 
                if (_countOfDots != value)
                    Mvx.Resolve<IUserInteraction>().Alert(string.Format("Впереди {0} ямки"));

                _countOfDots = value; 
                RaisePropertyChanged(() => CountOfDots);
            }
        }

        #region Location Methods

        private async void OnLocation(MvxGeoLocation currentLocation)
        {
            
            _lastLocation = currentLocation;

            if (_lastLocation != null)
            {
                Lat = _lastLocation.Coordinates.Latitude;
                Lon = _lastLocation.Coordinates.Longitude;
            }

            var newPits = _processor.GetPitsAhead(_lastLocation);
            if (newPits.Length > 0)
            {
                // TODO: make a notification

                Mvx.Resolve<IUserInteraction>().Alert("Впереди яма!");
            }

        }

        private async void updatePits() {
            _pits = _dbService.ReadData();
            if (_pits == null || _pits.Count == 0)
            {
                _pits = await _apiService.GetPits(
                    _lastLocation.Coordinates.Latitude, _lastLocation.Coordinates.Longitude, Variables.Radius, null
                );
                _dbService.SaveData (_pits);
            }
            _processor = new PitsProcessor(_pits);
        }

        private void OnError(MvxLocationError error)
        {
            Mvx.Error("Seen location error {0}", error.Code);
        }

        #endregion

        private MvxCommand _showPits;
        public ICommand ShowPits
        {
            get
            { 
                _showPits = _showPits ?? new MvxCommand(async () => {
                    

                });
                return _showPits;
            }
        }



        void PostProcess(IList<PitModel> pits)
        {
            
        }
    }
}


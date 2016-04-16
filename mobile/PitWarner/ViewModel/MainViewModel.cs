using System;
using MvvmCross.Core.ViewModels;
using System.Threading.Tasks;
using System.Windows.Input;
using System.Threading;
using MvvmCross.Platform;
using System.Collections.Generic;
using System.Linq;
using MvvmCross.Plugins.Location;

namespace PitWarner.ViewModels
{
    public class MainViewModel : BaseViewModel
    {
        const double DISTANCE_LIMIT = 300; //в метрах
        readonly IApiService _apiService;
        private readonly IMvxLocationWatcher _watcher;

        List<PitModel> _pits;
        MvxGeoLocation _lastLocation;

        public MainViewModel(IApiService apiService, IMvxLocationWatcher watcher)
        {
            _apiService = apiService;
            _watcher = watcher;
        }

        public override void Start()
        {
            base.Start();

            _watcher.Start(new MvxLocationOptions(), OnLocation, OnError);


//            var pits = new List<PitModel>
//            { 
//                new PitModel { Lat = 0, Lon = 0 },
//                new PitModel { Lat = 0, Lon = 10 },
//                    new PitModel { Lat = 10, Lon = 10 },
//                    new PitModel { Lat = 10, Lon = 0 },
//            };
//
//            var isContain = GeoCalc.InPoly(pits, new PitModel{ Lat = -1, Lon = 5 });
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

        #region Location Methods

        private void OnLocation(MvxGeoLocation currentLocation)
        {
            _lastLocation = currentLocation;

            if (_lastLocation != null)
            {
                Lat = _lastLocation.Coordinates.Latitude;
                Lon = _lastLocation.Coordinates.Longitude;
            }
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

                    _pits = await _apiService.GetPits(_lastLocation.Coordinates.Latitude, _lastLocation.Coordinates.Longitude, null);

                    if(_pits != null && _pits.Any())
                        PitProcess(_pits);

                });
                return _showPits;
            }
        }

        void PitProcess(List<PitModel> pits)
        {
            if (_lastLocation == null)
                return;

            var nearPoints = new List<PitModel>();

            foreach (var pit in pits)
            {
                var distanceToPit = GeoCalc.GetDistanceBetween2Points(
                    pit,
                    new PitModel { Lat = _lastLocation.Coordinates.Latitude, Lon = _lastLocation.Coordinates.Longitude, At = _lastLocation.Coordinates.Altitude ?? 0 }
                );

                if (distanceToPit < DISTANCE_LIMIT)
                    nearPoints.Add(pit);
            }
        }
    }
}


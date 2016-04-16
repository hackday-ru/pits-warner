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

namespace PitWarner.ViewModels
{
    public class MainViewModel : BaseViewModel
    {


        readonly IApiService _apiService;
        readonly IDataBaseService _dbService;
        private readonly IMvxLocationWatcher _watcher;

        List<PitModel> _pits;
        MvxGeoLocation _lastLocation;

        public MainViewModel(IApiService apiService, IDataBaseService dbService, IMvxLocationWatcher watcher)
        {
            _apiService = apiService;
            _dbService = dbService;
            _watcher = watcher;
        }

        public async override void Start()
        {
            base.Start();

            _watcher.Start(new MvxLocationOptions(), OnLocation, OnError);

//            var pits = await _apiService.GetPits(_lastLocation.Coordinates.Latitude, _lastLocation.Coordinates.Longitude, Variables.Radius, null);

            var pits = await _apiService.GetPits(59.8950, 30.3168, Variables.Radius, null);

            _dbService.SaveData (pits);


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

//                    var data = GeoCalc.MakeRectangle();
//                    Debug.WriteLine(data);

                });
                return _showPits;
            }
        }

        //void PitProcess(List<PitModel> pits)
        //{
        //    if (_lastLocation == null)
        //        return;

        //    var nearPoints = new List<PitModel>();

        //    foreach (var pit in pits)
        //    {
        //        var distanceToPit = GeoCalc.GetDistanceBetween2Points(
        //            pit,
        //            new PitModel { lat = _lastLocation.Coordinates.Latitude, lng = _lastLocation.Coordinates.Longitude, at = _lastLocation.Coordinates.Altitude ?? 0 }
        //        );

        //        if (distanceToPit < DISTANCE_LIMIT)
        //            nearPoints.Add(pit);
        //    }
        //}
    }
}


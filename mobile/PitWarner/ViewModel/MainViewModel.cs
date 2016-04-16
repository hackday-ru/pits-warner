using System;
using MvvmCross.Core.ViewModels;
using System.Threading.Tasks;
using System.Windows.Input;
using System.Threading;
using MvvmCross.Platform;
using System.Collections.Generic;
using System.Linq;

namespace PitWarner.ViewModels
{
    public class MainViewModel : BaseViewModel
    {
        readonly IApiService _apiService;

        List<PitModel> _pits;

        public MainViewModel(IApiService apiService)
        {
            _apiService = apiService;
        }

        private MvxCommand _showPits;
        public ICommand ShowPits
        {
            get
            { 
                _showPits = _showPits ?? new MvxCommand(async () => {

                    _pits = await _apiService.GetPits(null);

                    if(_pits != null && _pits.Any())
                        PitProcess(_pits);

                });
                return _showPits;
            }
        }

        void PitProcess(IList<PitModel> pits)
        {
            
        }
    }
}


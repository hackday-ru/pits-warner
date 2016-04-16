using System;
using MvvmCross.Core.ViewModels;
using System.Threading.Tasks;
using System.Windows.Input;
using System.Threading;

namespace PitWarner.ViewModels
{
    public class MainViewModel : BaseViewModel
    {
        public MainViewModel()
        {
            
        }

        private MvxCommand _cancelCollectData;
        public ICommand CancelCollectData
        {
            get
            { 
                _cancelCollectData = _cancelCollectData ?? new MvxCommand(() => {
                    

                });
                return _cancelCollectData;
            }
        }

        public override void Start()
        {
            base.Start();

        }
    }
}


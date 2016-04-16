using System;

using UIKit;
using MvvmCross.iOS.Views;
using MvvmCross.Binding.BindingContext;
using PitWarner.ViewModels;

namespace PitWarner.iOS
{
    public partial class MainView : MvxViewController
    {
        public MainView()
            : base("MainView", null)
        {
        }

        public override void ViewDidLoad()
        {
            base.ViewDidLoad();

            var set = this.CreateBindingSet<MainView, MainViewModel>();
            set.Bind(StartButton).To(vm => vm.ShowPits);
            set.Apply();
        }

        public override void DidReceiveMemoryWarning()
        {
            base.DidReceiveMemoryWarning();
            // Release any cached data, images, etc that aren't in use.
        }
    }
}



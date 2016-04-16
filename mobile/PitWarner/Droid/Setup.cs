using System;
using MvvmCross.Droid.Platform;
using MvvmCross.Core.ViewModels;
using MvvmCross.Platform.Platform;
using Android.Content;

namespace PitWarner.Droid
{
    public class Setup : MvxAndroidSetup
    {

        public Setup(Context applicationContext) : base (applicationContext)
        {
            
        }
//        protected override void InitializeFirstChance()
//        {
//            Mvx.RegisterSingleton<IZipService>(new DroidZipService());
//            Mvx.RegisterSingleton<IPlatformFileService>(new DroidFileService());
//            Mvx.RegisterSingleton<IPlatformHttpService>(new DroidPlatformHttpService());
//            base.InitializeFirstChance();
//        }

        protected override IMvxApplication CreateApp()
        {
            return new App();
        }
    }
}


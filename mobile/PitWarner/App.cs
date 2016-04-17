using System;
using MvvmCross.Core.ViewModels;
using MvvmCross.Platform.IoC;
using System.Diagnostics;
using MvvmCross.Platform;
using MvvmCross.Plugins.File;

namespace PitWarner
{
    public class App : MvxApplication
    {
        public App()
        {
        }

        public override void Initialize()
        {
            CreatableTypes()
                .EndingWith("Service")
                .AsInterfaces()
                .RegisterAsLazySingleton();

            Debug.WriteLine(Mvx.Resolve<IMvxFileStore>().NativePath(""));


            RegisterAppStart<ViewModels.MainViewModel>();
        }
    }
}


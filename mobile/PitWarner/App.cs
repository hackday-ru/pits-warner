using System;
using MvvmCross.Core.ViewModels;
using MvvmCross.Platform.IoC;

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


            RegisterAppStart<ViewModels.MainViewModel>();
        }
    }
}


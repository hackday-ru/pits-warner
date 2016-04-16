using Android.App;
using Android.Widget;
using Android.OS;
using PitWarner.ViewModels;
using MvvmCross.Droid.Views;
using MvvmCross.Binding.BindingContext;

namespace PitWarner.Droid
{
    [Activity(Label = "PitWarner", MainLauncher = true, Icon = "@mipmap/icon")]
    public class MainActivity : MvxActivity<MainViewModel>
    {
        int count = 1;

        protected override void OnCreate(Bundle savedInstanceState)
        {
            base.OnCreate(savedInstanceState);

            // Set our view from the "main" layout resource
            SetContentView(Resource.Layout.Main);

            var set = this.CreateBindingSet<MainActivity, MainViewModel>();

            // Get our button from the layout resource,
            // and attach an event to it
            Button button = FindViewById<Button>(Resource.Id.myButton);
			
            button.Click += delegate
            {
                button.Text = string.Format("{0} clicks!", count++);
            };
        }
    }
}



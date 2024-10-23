using System;
using System.Windows;
using System.Net.Http;

namespace GW_UI
{
    public partial class App : Application
    {
        private void Application_Startup(object sender, StartupEventArgs e)
        {

        }

        public static readonly HttpClient HttpClient = new HttpClient
        {
            BaseAddress = new Uri("http://demo.localdev.me")
        };
    }
}

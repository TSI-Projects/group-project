using System;
using System.Windows;
using System.Net.Http;
using System.Net.Http.Headers;


namespace GW_UI
{
    public partial class App : Application
    {
        private void Application_Startup(object sender, StartupEventArgs e)
        {
            System.Diagnostics.PresentationTraceSources.DataBindingSource.Switch.Level = System.Diagnostics.SourceLevels.Critical;

        }

        public static readonly HttpClient HttpClient = new HttpClient
        {
            BaseAddress = new Uri("http://demo.localdev.me")
        };

        public static void SetToken(string Token)
        {
            HttpClient.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Bearer", Token);
        }
    }
}

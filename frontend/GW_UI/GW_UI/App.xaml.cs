using System;
using System.Collections.Generic;
using System.Configuration;
using System.Data;
using System.Linq;
using System.Threading.Tasks;
using System.Windows;
using System.Net.Http;

namespace GW_UI
{
    /// <summary>
    /// Interaction logic for App.xaml
    /// </summary>
    public partial class App : Application
    {
        public static HttpClient httpClient;

        private void Application_Startup(object sender, StartupEventArgs e)
        {
            httpClient = new HttpClient(); //почитать
        }
    }
}

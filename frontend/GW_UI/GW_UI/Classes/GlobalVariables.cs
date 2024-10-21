using System;
using System.Net.Http;

namespace GW_UI
{
    public static class GlobalVariables
    {
        public static readonly HttpClient HttpClient = new HttpClient
        {
            BaseAddress = new Uri("http://demo.localdev.me")
        };
    }
}
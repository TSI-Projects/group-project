using GW_UI.Classes;
using System;
using System.Collections.Generic;
using System.Net.Http.Json;
using System.Windows;
using System.Windows.Controls;

namespace GW_UI
{
    public partial class Login : Window
    {
        public Login()
        {
            InitializeComponent();
        }

        private async void LoginButton_Click(object sender, RoutedEventArgs e)
        {
            var users = new Users
            {
                Username = UsernameTextBox.Text,
                Password = PasswordBox.Password
            };

            try
            {
                var response = await App.HttpClient.PostAsJsonAsync("/api/login", users);
                var loginResponse = await response.Content.ReadFromJsonAsync<LoginResponse>();
                if (!loginResponse.Success)
                {
                    throw new Exception(loginResponse.Error.Message);
                }
                App.SetToken(loginResponse.AccessToken);
                MainMenu mainMenu = new MainMenu();
                mainMenu.Show();
                this.Close();
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }
        }

        public void LogoutButton_Click(object sender, RoutedEventArgs e)
        {
            Application.Current.Shutdown();
        }

        private void Hyperlink_Click(object sender, RoutedEventArgs e)
        {
            MessageBox.Show("Please contact support for help with logging in.", "Login Assistance", MessageBoxButton.OK, MessageBoxImage.Information);
        }
    }
}
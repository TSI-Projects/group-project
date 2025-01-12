using System;
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
                if (response.IsSuccessStatusCode)
                {
                    MainMenu mainMenu = new MainMenu();
                    mainMenu.Show();
                    var token = await response.Content.ReadAsStringAsync();
                    App.SetToken(token);
                    this.Close();
                }
                else
                {
                    MessageBox.Show("Wrong Login or Password " + response.ReasonPhrase);
                }
            }
            catch (Exception ex)
            {
                MessageBox.Show("Ошибка при отправке данных: " + ex.Message);
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
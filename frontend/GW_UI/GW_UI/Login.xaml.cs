using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http.Json;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Shapes;

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
    }
}
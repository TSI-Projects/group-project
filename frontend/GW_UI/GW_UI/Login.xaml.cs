using System;
using System.Collections.Generic;
using System.Linq;
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

        private void LoginButton_Click(object sender, RoutedEventArgs e)
        {
            string login = UsernameTextBox.Text;
            string password = PasswordBox.Password;

            // Проверка логина и пароля
            if (login == "11111111" && password == "1")
            {
                // Открываем главное меню при успешной проверке
                MainMenu mainMenu = new MainMenu();
                mainMenu.Show();
                this.Close(); // Закрываем окно логина
            }
            else
            {
                // Выводим сообщение об ошибке при неверном логине или пароле
                MessageBox.Show("Неверный логин или пароль", "Ошибка входа", MessageBoxButton.OK, MessageBoxImage.Error);
            }
        }
    }
}
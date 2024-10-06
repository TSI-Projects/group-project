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
    /// <summary>
    /// Interaction logic for Employees.xaml
    /// </summary>
    public partial class Employees : Window
    {
        public Employees()
        {
            InitializeComponent();
        }

        public void LogoutButton_Click(object sender, RoutedEventArgs e)
        {
            Application.Current.Shutdown();
        }

        private void HomeButton_Click(object sender, RoutedEventArgs e)
        {
            MainMenu mainMenu = new MainMenu();
            mainMenu.Show();
            this.Close();
        }

        private void RemoveText(object sender, RoutedEventArgs e)
        {
            TextBox textBox = sender as TextBox;
            if (textBox != null)
            {
                TextBlock placeholder = (TextBlock)this.FindName($"{textBox.Name.Replace("TextBox", "Placeholder")}");
                if (placeholder != null)
                {
                    placeholder.Visibility = Visibility.Hidden;
                }
            }
        }

        private void AddText(object sender, RoutedEventArgs e)
        {
            TextBox textBox = sender as TextBox;
            if (textBox != null && string.IsNullOrEmpty(textBox.Text))
            {
                TextBlock placeholder = (TextBlock)this.FindName($"{textBox.Name.Replace("TextBox", "Placeholder")}");
                if (placeholder != null)
                {
                    placeholder.Visibility = Visibility.Visible;
                }
            }
        }

        private void AddEmployee_Click(object sender, RoutedEventArgs e)
        {
            // Логика добавления нового сотрудника
            //EmployeesList.Add(new Employee { FirstName = FirstNameTextBox.Text, LastName = LastNameTextBox.Text });
        }

        private void DeleteEmployee_Click(object sender, RoutedEventArgs e)
        {
            // Логика удаления выбранного сотрудника
            //if (EmployeeGrid.SelectedItem != null)
            //{
            //    EmployeesList.Remove((Employee)EmployeeGrid.SelectedItem);
            //}
        }

        private void BackButton_Click(object sender, RoutedEventArgs e)
        {
            Menu menuPage = new Menu();
            menuPage.Show();
            this.Close();
        }
    }
}

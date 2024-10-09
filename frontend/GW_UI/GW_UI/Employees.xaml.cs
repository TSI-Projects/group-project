using System;
using System.Windows;
using System.Windows.Controls;
using System.Collections.ObjectModel;

namespace GW_UI
{
    public partial class Employees : Window
    {
        //private List<Employee> EmployeesList = new List<Employee>();
        private ObservableCollection<Employee> EmployeesList = new ObservableCollection<Employee>();

        public Employees()
        {
            InitializeComponent();
            EmployeeGrid.ItemsSource = EmployeesList; // источник данных для DataGrid
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
            // логика добавления нового сотрудника
            EmployeesList.Add(new Employee { FirstName = FirstNameTextBox.Text, LastName = LastNameTextBox.Text });
        }

        private void DeleteEmployee_Click(object sender, RoutedEventArgs e)
        {
            // логика удаления выбранного сотрудника
            if (EmployeeGrid.SelectedItem != null)
            {
                EmployeesList.Remove((Employee)EmployeeGrid.SelectedItem);
            }
        }

        private void BackButton_Click(object sender, RoutedEventArgs e)
        {
            Menu menuPage = new Menu();
            menuPage.Show();
            this.Close();
        }
    }

    // Класс Employee (создать отдельно файл с классом если заработает)
    public class Employee
    {
        public string FirstName { get; set; }
        public string LastName { get; set; }
    }
}

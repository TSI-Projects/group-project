using System;
using System.Windows;
using System.Windows.Controls;
using System.Collections.ObjectModel;
using System.Collections.Generic;
using System.Net.Http;
using System.Net.Http.Json;
using System.Text.Json.Serialization;

namespace GW_UI
{
    public partial class Employees : Window
    {
        private ObservableCollection<Employee> EmployeesList = new ObservableCollection<Employee>();

        public Employees()
        {
            InitializeComponent();
            EmployeeGrid.ItemsSource = EmployeesList; // источник данных для DataGrid
            this.Loaded += EmployeeWindow_Loaded;
        }

        private async void EmployeeWindow_Loaded(object sender, RoutedEventArgs e)
        {
            var result = await App.HttpClient.GetFromJsonAsync<List<Employee>>("/api/workers");
            //можно оптимизировать, использовать метод вместо фор лупа
            if (result == null)
            {
                return;
            }

            foreach (Employee emp in result)
            {
                EmployeesList.Add(emp);
            }

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

        private void BackButton_Click(object sender, RoutedEventArgs e)
        {
            Menu menuPage = new Menu();
            menuPage.Show();
            Close();
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

        private async void AddEmployee_Click(object sender, RoutedEventArgs e)
        {
            // логика добавления нового сотрудника
            var data = new Employee(FirstNameTextBox.Text, LastNameTextBox.Text);
            var result = await App.HttpClient.PostAsJsonAsync("/api/workers", data);

            EmployeesList.Add(new Employee(FirstNameTextBox.Text, LastNameTextBox.Text));
        }

        private async void DeleteEmployee_Click(object sender, RoutedEventArgs e)
        {
            var worker = (Employee)EmployeeGrid.SelectedItem;
            
           await App.HttpClient.DeleteAsync($"/api/worker/{worker.ID}");
            // логика удаления выбранного сотрудника
            if (EmployeeGrid.SelectedItem != null)
            {
                EmployeesList.Remove((Employee)EmployeeGrid.SelectedItem);
            }
        }
    }
}

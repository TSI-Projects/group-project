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
            var client = new HttpClient(); //Создать глобальную переменную
            client.BaseAddress = new Uri("http://demo.localdev.me");
  
            var result = await client.GetFromJsonAsync<List<Employee>>("/api/workers");
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
            var client = new HttpClient(); //Создать глобальную переменную
            client.BaseAddress = new Uri("http://demo.localdev.me");
            var data = new Employee(FirstNameTextBox.Text, LastNameTextBox.Text);
            var result = await client.PostAsJsonAsync("/api/workers", data);

            EmployeesList.Add(new Employee(FirstNameTextBox.Text, LastNameTextBox.Text));
        }

        private async void DeleteEmployee_Click(object sender, RoutedEventArgs e)
        {
            var worker = (Employee)EmployeeGrid.SelectedItem;
            var client = new HttpClient(); //Создать глобальную переменную
            client.BaseAddress = new Uri("http://demo.localdev.me");
            
           await client.DeleteAsync($"/api/worker/{worker.ID}");
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
            Close();
        }
    }

    //// Класс Employee (создать отдельно файл с классом если заработает)
    //public class Employee
    //{
    //    public Employee(string firstName, string lastName)
    //    {
    //        FirstName = firstName;
    //        LastName = lastName;
    //    }

    //    [JsonPropertyName("id")]
    //    public int ID { get; set; }
    //    [JsonPropertyName("first_name")]
    //    public string FirstName { get; set; }
    //    [JsonPropertyName("last_name")]
    //    public string LastName { get; set; }
    //}
}

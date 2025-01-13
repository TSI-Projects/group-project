using System;
using System.Windows;
using System.Windows.Controls;
using System.Collections.ObjectModel;
using System.Collections.Generic;
using System.Net.Http;
using System.Net.Http.Json;
using System.Text.Json.Serialization;
using GW_UI.Classes;

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
            try
            {
                var result = await App.HttpClient.GetFromJsonAsync<EmployeeResponse>("/api/workers");
                if (!result.Success && result.Error != null)
                {
                    throw new Exception(result.Error.Message);
                }

                foreach (Employee emp in result.Workers)
                {
                    EmployeesList.Add(emp);
                }
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
            try
            {
                var data = new Employee(FirstNameTextBox.Text, LastNameTextBox.Text);
                var result = await App.HttpClient.PostAsJsonAsync("/api/workers", data);
                var body = await result.Content.ReadFromJsonAsync<EmployeeResponse>();

                if (!body.Success && body.Error != null)
                {
                    throw new Exception(body.Error.Message);
                }
                EmployeesList.Add(new Employee(FirstNameTextBox.Text, LastNameTextBox.Text));
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }

        }

        private async void DeleteEmployee_Click(object sender, RoutedEventArgs e)
        {
            try
            {   
                var worker = (Employee)EmployeeGrid.SelectedItem;
                var result = await App.HttpClient.DeleteAsync($"/api/worker/{worker.ID}");
                var body = await result.Content.ReadFromJsonAsync<EmployeeResponse>();
                if (!body.Success && body.Error != null)
                {
                    throw new Exception(body.Error.Message);
                }
                // логика удаления выбранного сотрудника
                if (EmployeeGrid.SelectedItem != null)
                {
                    EmployeesList.Remove((Employee)EmployeeGrid.SelectedItem);
                }
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }
        }
    }
}

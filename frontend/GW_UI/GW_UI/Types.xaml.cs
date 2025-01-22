using System.Windows;
using System.Windows.Controls;
using System.Collections.ObjectModel;
using System.Collections.Generic;
using System.Net.Http.Json;
using System;
using GW_UI.Classes;

namespace GW_UI
{
    public partial class Types : Window
    {
        private ObservableCollection<TypeItem> TypesList = new ObservableCollection<TypeItem>();

        public Types()
        {
            InitializeComponent();
            TypeGrid.ItemsSource = TypesList; // источник данных для DataGrid
            this.Loaded += TypesWindow_Loaded;
        }

        private async void TypesWindow_Loaded(object sender, RoutedEventArgs e)
        {
            try 
            {
                var result = await App.HttpClient.GetFromJsonAsync<TypeResponse>("/api/orders/types");

                if (!result.Success && result.Error != null)
                {
                    throw new Exception(result.Error.Message);
                }

                foreach (TypeItem type in result.Types)
                {
                    TypesList.Add(type);
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

        private async void AddType_Click(object sender, RoutedEventArgs e)
        {
            try
            {
                var data = new TypeItem { TypeName = TypeTextBox.Text };
                var result = await App.HttpClient.PostAsJsonAsync("/api/orders/types", data);
                var body = await result.Content.ReadFromJsonAsync<TypeResponse>();

                if (!body.Success && body.Error != null)
                {
                    throw new Exception(body.Error.Message);
                }
                TypesList.Add(new TypeItem { TypeName = TypeTextBox.Text });
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }
        }

        private async void DeleteType_Click(object sender, RoutedEventArgs e)
        {
            try
            {
                var type = (TypeItem)TypeGrid.SelectedItem;
                if (type == null)
                {
                    MessageBox.Show("Please select a type to delete.");
                    return;
                }

                MessageBoxResult result = MessageBox.Show(
                    "Are you sure you want to delete the selected type?",
                    "Confirmation",
                    MessageBoxButton.YesNo,
                    MessageBoxImage.Warning
                );

                if (result == MessageBoxResult.Yes)
                {
                    var response = await App.HttpClient.DeleteAsync($"/api/orders/type/{type.ID}");
                    var body = await response.Content.ReadFromJsonAsync<EmployeeResponse>();

                    if (!body.Success && body.Error != null)
                    {
                        throw new Exception(body.Error.Message);
                    }

                    TypesList.Remove(type);
                }
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }
        }

        private void BackButton_Click(object sender, RoutedEventArgs e)
        {
            Menu menuPage = new Menu();
            menuPage.Show();
            Close();
        }
    }
}
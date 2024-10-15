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
            var client = new HttpClient();
            client.BaseAddress = new Uri("http://demo.localdev.me");

            var result = await client.GetFromJsonAsync<List<TypeItem>>("/api/orders/types");

            if (result == null)
            {
                return;
            }

            foreach (TypeItem type in result)
            {
                TypesList.Add(type);
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
            // логика добавления нового типа
            var client = new HttpClient();
            client.BaseAddress = new Uri("http://demo.localdev.me");
            var data = new TypeItem { TypeName = TypeTextBox.Text };
            var result = await client.PostAsJsonAsync("/api/orders/types", data);

            TypesList.Add(new TypeItem {TypeName = TypeTextBox.Text});
        }

        private async void DeleteType_Click(object sender, RoutedEventArgs e)
        {
            var type = (TypeItem)TypeGrid.SelectedItem;
            var client = new HttpClient();
            client.BaseAddress = new Uri("http://demo.localdev.me");

            await client.DeleteAsync($"/api/type/{type.ID}");

            if (TypeGrid.SelectedItem != null)
            {
                TypesList.Remove((TypeItem)TypeGrid.SelectedItem);
            }
        }

        private void BackButton_Click(object sender, RoutedEventArgs e)
        {
            Menu menuPage = new Menu();
            menuPage.Show();
            Close();
        }
    }

    // Класс TypeItem
    public class TypeItem
    {
        [JsonPropertyName("id")]
        public int ID { get; set; }

        [JsonPropertyName("full_name")]
        public string TypeName { get; set; }
    }
}

using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Net.Http.Json;
using System.Windows;

namespace GW_UI
{
    public partial class CompletedOrders : Window
    {
        private ObservableCollection<Order> OrdersList = new ObservableCollection<Order>();

        public CompletedOrders()
        {
            InitializeComponent();
            OrdersDataGrid.ItemsSource = OrdersList; // источник данных для DataGrid
            this.Loaded += OrderWindow_Loaded;
        }

        private async void OrderWindow_Loaded(object sender, RoutedEventArgs e)
        {
            try
            {
                var orders = await App.HttpClient.GetFromJsonAsync<List<Order>>("/api/orders/completed");

                if (orders != null)
                {
                    OrdersList.Clear();
                    foreach (var order in orders)
                    {
                        OrdersList.Add(order);
                    }
                }
            }
            catch (Exception ex)
            {
                MessageBox.Show($"Ошибка загрузки данных: {ex.Message}");
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
            Close();
        }

        private void BackButton_Click(object sender, RoutedEventArgs e)
        {
            MainMenu mainMenu = new MainMenu();
            mainMenu.Show();
            Close();
        }
    }
}

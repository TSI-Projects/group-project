using GW_UI.Classes;
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
                var response = await App.HttpClient.GetFromJsonAsync<OrderResponse>("/api/orders/completed");

                if (!response.Success && response.Error != null)
                {
                    throw new Exception(response.Error.Message);
                }
                OrdersList.Clear();
                foreach (var order in response.Orders)
                {
                    OrdersList.Add(order);
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

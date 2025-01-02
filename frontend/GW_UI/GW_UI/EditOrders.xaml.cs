using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Net.Http.Json;
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
    public partial class EditOrders : Window
    {
        private ObservableCollection<Order> OrdersList = new ObservableCollection<Order>();


        public EditOrders()
        {
            InitializeComponent();
            OrdersDataGrid.ItemsSource = OrdersList; // источник данных для DataGrid
            this.Loaded += OrderWindow_Loaded;
        }

        //private async void OrderWindow_Loaded(object sender, RoutedEventArgs e)
        //{
        //    var result = await App.HttpClient.GetFromJsonAsync<List<Order>>("/api/orders");
        //    //var result = await App.HttpClient.GetFromJsonAsync<List<TypeItem>>("/api/orders/types");

        //    //можно оптимизировать, использовать метод вместо фор лупа
        //    if (result == null)
        //    {
        //        return;
        //    }

        //    foreach (Order emp in result)
        //    {
        //        OrdersList.Add(emp);
        //    }

        //}

        private async void OrderWindow_Loaded(object sender, RoutedEventArgs e)
        {
            try
            {
                var orders = await App.HttpClient.GetFromJsonAsync<List<Order>>("/api/orders");
                var types = await App.HttpClient.GetFromJsonAsync<List<TypeItem>>("/api/orders/types");
                var employees = await App.HttpClient.GetFromJsonAsync<List<Employee>>("/api/workers");

                if (orders != null)
                {
                    foreach (var order in orders)
                    {
                        // Устанавливаем связанные объекты
                        order.TypeItem = types.FirstOrDefault(t => t.ID == order.OrderTypeId);
                        order.Employee = employees.FirstOrDefault(r => r.ID == order.WorkerId);
                    }

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



        private void EditOrder_Click(object sender, RoutedEventArgs e)
        {

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

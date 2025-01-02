using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Net.Http.Json;
using System.Windows;
using System.Windows.Controls;

namespace GW_UI
{
    public partial class EditOrders : Window
    {
        private ObservableCollection<Order> OrdersList = new ObservableCollection<Order>();
        private Button currentEditButton;

        public EditOrders()
        {
            InitializeComponent();
            OrdersDataGrid.ItemsSource = OrdersList; // источник данных для DataGrid
            this.Loaded += OrderWindow_Loaded;
        }

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
            if (sender is Button editButton && OrdersDataGrid.SelectedItem is Order selectedOrder)
            {
                // Сохранение ссылки на текущую кнопку
                currentEditButton = editButton;

                // Включить режим редактирования для DataGrid
                OrdersDataGrid.IsReadOnly = false;

                // Изменить стиль и функциональность кнопки
                editButton.Style = (Style)FindResource("SaveButtonStyle");
                editButton.Click -= EditOrder_Click; // Отписка от предыдущего события
                editButton.Click += SaveOrder_Click; // Подписка на новое событие
            }
            else
            {
                MessageBox.Show("Выберите строку для редактирования.");
            }
        }

        private async void SaveOrder_Click(object sender, RoutedEventArgs e)
        {
            if (OrdersDataGrid.SelectedItem is Order selectedOrder)
            {
                try
                {
                    // Отправить обновления в API
                    var response = await App.HttpClient.PutAsJsonAsync($"/api/orders/{selectedOrder.Id}", selectedOrder);
                    if (response.IsSuccessStatusCode)
                    {
                        MessageBox.Show("Изменения успешно сохранены!");
                        OrdersDataGrid.IsReadOnly = true;

                        // Вернуть кнопку в режим "Редактировать"
                        if (currentEditButton != null)
                        {
                            currentEditButton.Style = (Style)FindResource("EditButtonStyle");
                            currentEditButton.Click -= SaveOrder_Click; // Отписка от события сохранения
                            currentEditButton.Click += EditOrder_Click; // Подписка на событие редактирования
                        }
                    }
                    else
                    {
                        MessageBox.Show("Ошибка сохранения изменений: " + response.ReasonPhrase);
                    }
                }
                catch (Exception ex)
                {
                    MessageBox.Show($"Ошибка при сохранении изменений: {ex.Message}");
                }
            }
            else
            {
                MessageBox.Show("Выберите строку для сохранения изменений.");
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

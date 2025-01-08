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

        private void EditOrder_Click(object sender, RoutedEventArgs e)
        {
            Window editPage = new Window
            {
                Title = "Edit Page",
                Content = new EditPage(),
                Height = 800,
                Width = 1200,
                Owner = this,
                WindowStartupLocation = WindowStartupLocation.CenterOwner
            };

            editPage.ShowDialog();
        }

        //private void EditOrder_Click(object sender, RoutedEventArgs e)
        //{
        //    if (sender is Button editButton && OrdersDataGrid.SelectedItem is Order selectedOrder)
        //    {
        //        // Сохранение ссылки на текущую кнопку
        //        currentEditButton = editButton;

        //        // Включить режим редактирования для DataGrid
        //        OrdersDataGrid.IsReadOnly = false;

        //        // Изменить стиль и функциональность кнопки
        //        editButton.Style = (Style)FindResource("SaveButtonStyle");
        //        editButton.Click -= EditOrder_Click; // Отписка от предыдущего события
        //        editButton.Click += SaveOrder_Click; // Подписка на новое событие
        //    }
        //    else
        //    {
        //        MessageBox.Show("Выберите строку для редактирования.");
        //    }
        //}

        private async void SaveOrder_Click(object sender, RoutedEventArgs e)
        {
            if (!(OrdersDataGrid.SelectedItem is Order selectedOrder))
            {
                return;
            }


            OrdersDataGrid.IsReadOnly = false;
            try
            {
                selectedOrder.TotalPrice = double.Parse(AccessCellByColumnName("Total"));
                selectedOrder.Prepayment = double.Parse(AccessCellByColumnName("Prepayment"));
                selectedOrder.Reason = AccessCellByColumnName("Reason");
                //selectedOrder.OrderStatus.IsOutsourced = bool.Parse(AccessCellByColumnName("Outsource"));

                // Отправить обновления в API
                var response = await App.HttpClient.PutAsJsonAsync($"/api/orders", selectedOrder);
                if (!response.IsSuccessStatusCode)
                {
                    MessageBox.Show("Ошибка сохранения изменений: " + response.ReasonPhrase);
                    return;
                }

                // Вернуть кнопку в режим "Редактировать"
                if (currentEditButton != null)
                {
                    currentEditButton.Style = (Style)FindResource("EditButtonStyle");
                    currentEditButton.Click -= SaveOrder_Click; // Отписка от события сохранения
                    currentEditButton.Click += EditOrder_Click; // Подписка на событие редактирования
                }
                MessageBox.Show("Изменения успешно сохранены!");
            }

            catch (Exception ex)
            {
                MessageBox.Show($"Ошибка при сохранении изменений: {ex.Message}");
            }

            OrdersDataGrid.IsReadOnly = true;
        }

        //private void ToggleEditing()
        //{ 
        //    OrdersDataGrid.Row
        //}

        private string AccessCellByColumnName(string columnName)
        {
            var item = OrdersDataGrid.SelectedItem; // Get the item at the specified row index
            var column = OrdersDataGrid.Columns.FirstOrDefault(c => c.Header.ToString() == columnName);

            if (column == null)
            {
                throw new Exception("Column Name is Not Found");
            }

            var cellContent = column.GetCellContent(item); // Get the cell content

            if (cellContent is TextBlock textBlock)
            {
                return textBlock.Text;
            }

            return "";
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

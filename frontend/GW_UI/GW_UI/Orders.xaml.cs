using System;
using System.Collections.Generic;
using System.Linq;
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
using System.Windows.Controls.Primitives;
using System.Net.Http.Json;
using System.Collections.ObjectModel;

namespace GW_UI
{
    public partial class Orders : Window
    {
        private ToggleButton activeLanguageButton;
        public ObservableCollection<TypeItem> OrderTypes { get; set; }
        public ObservableCollection<Employee> Employees { get; set; }


        public Orders()
        {
            InitializeComponent();
            OrderTypes = new ObservableCollection<TypeItem>();
            Employees = new ObservableCollection<Employee>();
            OrderTypeComboBox.ItemsSource = OrderTypes;
            EmployeeNameComboBox.ItemsSource = Employees;
            this.Loaded += OrdersWindow_Loaded; //Сделать отписку
        }

        private async void OrdersWindow_Loaded(object sender, RoutedEventArgs e)
        {
            try
            {
                var orderTypes = await App.HttpClient.GetFromJsonAsync<List<TypeItem>>("/api/orders/types"); //Поправить
                if (orderTypes != null)
                {
                    foreach (TypeItem type in orderTypes)
                    {
                        OrderTypes.Add(type);
                    }
                }
                var employees = await App.HttpClient.GetFromJsonAsync<List<Employee>>("/api/workers"); //Поправить
                if (employees != null)
                {
                    foreach (Employee employee in employees)
                    {
                        Employees.Add(employee);
                    }
                }
            }
            catch (Exception ex)
            {
                MessageBox.Show("Ошибка загрузки типов заказов: " + ex.Message);
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

        private void RemoveText(object sender, RoutedEventArgs e)
        {
            TextBox textBox = sender as TextBox;
            if (textBox != null)
            {
                TextBlock placeholder = (TextBlock)this.FindName($"{textBox.Name.Replace("TextBox", "TextBlock")}");
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
                TextBlock placeholder = (TextBlock)this.FindName($"{textBox.Name.Replace("TextBox", "TextBlock")}");
                if (placeholder != null)
                {
                    placeholder.Visibility = Visibility.Visible;
                }
            }
        }

        private void OnTextChanged(object sender, TextChangedEventArgs e)
        {
            TextBox textBox = sender as TextBox;
            TextBlock placeholder = (TextBlock)this.FindName($"{textBox.Name.Replace("TextBox", "TextBlock")}");
            if (placeholder != null)
            {
                placeholder.Visibility = string.IsNullOrEmpty(textBox.Text) ? Visibility.Visible : Visibility.Hidden;
            }
        }

        private void RuButton_Click(object sender, RoutedEventArgs e)
        {
            HandleLanguageButtonClick(sender as ToggleButton);
        }

        private void LvButton_Click(object sender, RoutedEventArgs e)
        {
            HandleLanguageButtonClick(sender as ToggleButton);
        }

        private void EngButton_Click(object sender, RoutedEventArgs e)
        {
            HandleLanguageButtonClick(sender as ToggleButton);
        }

        private void HandleLanguageButtonClick(ToggleButton clickedButton)
        {
            if (activeLanguageButton != null && activeLanguageButton != clickedButton)
            {
                activeLanguageButton.IsChecked = false; // Деактивировать кнопку
            }
            // Если кнопка активна, снять с нее активацию
            if (activeLanguageButton == clickedButton)
            {
                activeLanguageButton = null;
            }
            else
            {
                activeLanguageButton = clickedButton; // Сделать новую кнопку активной
            }
        }

        private async void AddOrder_Click(object sender, RoutedEventArgs e)
        {
            int orderTypeId = (OrderTypeComboBox.SelectedItem as TypeItem).ID;
            int workerId = (EmployeeNameComboBox.SelectedItem as Employee).ID;
            double.TryParse(ClientPhoneTextBox.Text, out double customerId);
            //string requestDate = RequestDateTextBox.Text; //Вот тут вопросики возникают, в дизайне есть графа для даты, а в backend нет
            string reason = ReasonTextBox.Text;
            string defectDescription = DefectDescriptionTextBox.Text;
            double.TryParse(PrepaymentTextBox.Text, out double prepayment);
            double.TryParse(TotalCostTextBox.Text, out double totalPrice);

            var orderRequest = new Order
            {
                OrderTypeId = orderTypeId,
                WorkerId = workerId,
                CustomerId = customerId,
                Reason = reason,
                Defect = defectDescription,
                TotalPrice = totalPrice,
                Prepayment = prepayment,
                LanguageId = languageId
            };

            try
            {
                var response = await App.HttpClient.PostAsJsonAsync("/api/orders", orderRequest);
                if (response.IsSuccessStatusCode)
                {
                    MessageBox.Show("Заказ успешно добавлен!");
                }
                else
                {
                    MessageBox.Show("Ошибка добавления заказа: " + response.ReasonPhrase);
                }
            }
            catch (Exception ex)
            {
                MessageBox.Show("Ошибка при отправке данных: " + ex.Message);
            }
        }

        private int languageId = 0; //это временный костыль, надо придумать будет как будем языки передавать, если не будет CustomerID, то просто RU LV ENG

        private void RuButton_Checked(object sender, RoutedEventArgs e)
        {

        }

        private void LvButton_Checked(object sender, RoutedEventArgs e)
        {

        }

        private void EngButton_Checked(object sender, RoutedEventArgs e)
        {

        }

        //private void UpdateLanguageSelection(ToggleButton clickedButton)
        //{
        //    if (activeLanguageButton != null && activeLanguageButton != clickedButton)
        //    {
        //        activeLanguageButton.IsChecked = false;
        //    }
        //    activeLanguageButton = clickedButton;
        //    if (clickedButton != null)
        //    {
        //        clickedButton.IsChecked = true;
        //    }
        //}

        private void EmployeeNameComboBox_SelectionChanged(object sender, SelectionChangedEventArgs e)
        {
            if (EmployeeNameComboBox.SelectedItem != null)
            {
                EmployeeNameTextBlock.Text = "";
            }
        }

        private void OrderTypeComboBox_SelectionChanged(object sender, SelectionChangedEventArgs e)
        {
            if (OrderTypeComboBox.SelectedItem != null)
            {
                OrderTypeTextBlock.Text = "";
            }
        }

        private void DatePicker_SelectionChanged(object sender, SelectionChangedEventArgs e)
        {
            if (RequestDatePicker.SelectedDate != null)
            {
                RequestDateTextBlock.Text = "";
            }
        }

    }
}

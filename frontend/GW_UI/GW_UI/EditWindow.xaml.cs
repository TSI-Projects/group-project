using GW_UI.UserControls;
using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Net.Http.Json;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Controls.Primitives;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Effects;
using System.Windows.Media.Imaging;
using System.Windows.Shapes;

namespace GW_UI
{
    public partial class EditWindow : Window
    {
        private ToggleButton activeLanguageButton;
        public ObservableCollection<TypeItem> OrderTypes { get; set; }
        public ObservableCollection<Employee> Employees { get; set; }

        public Order order;
        public Employee emp;
        public Customer customer;

        public enum SelectedLanguage
        {
            RU = 1, LV = 2, ENG = 3
        }

        private SelectedLanguage selectedLanguage = SelectedLanguage.RU;

        public EditWindow(Order order)
        {
            InitializeComponent();
            this.order = order;
            DataContext = this;  // Set DataContext for binding.

            OrderTypes = new ObservableCollection<TypeItem>();
            OrderTypeComboBox.ItemsSource = OrderTypes;

            Employees = new ObservableCollection<Employee>();
            EmployeeNameComboBox.ItemsSource = Employees;


            //выгружаем на страницу
            OutsourceCheck.IsChecked = order.OrderStatus.IsOutsourced;
            CalledBackCheck.IsChecked = order.OrderStatus.CustomerNotifiedAt != null;
            ProductModelTextBox.Text = order.ItemName.ToString();   
            ClientPhoneTextBox.Text = order.Customer.PhoneNumber.ToString();
            ReasonTextBox.Text = order.Reason.ToString();
            DefectDescriptionTextBox.Text = order.Defect.ToString();
            TotalCostTextBox.Text = order.TotalPrice.ToString();
            PrepaymentTextBox.Text = order.Prepayment.ToString();

            switch (order.Customer.LanguageId)
            {
                case 1:
                    RuButton.IsChecked = true;
                    activeLanguageButton = RuButton;
                    selectedLanguage = SelectedLanguage.RU;
                    break;
                case 2:
                    LvButton.IsChecked = true;
                    activeLanguageButton = LvButton;
                    selectedLanguage = SelectedLanguage.LV;
                    break;
                case 3:
                    EngButton.IsChecked = true;
                    activeLanguageButton = EngButton;
                    selectedLanguage = SelectedLanguage.ENG;
                    break;
            }

            this.Loaded += OrdersWindow_Loaded;
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

                OrderTypeComboBox.SelectedValue = order.TypeItem.ID;
                EmployeeNameComboBox.SelectedValue = order.Employee.ID;

            }
            catch (Exception ex)
            {
                MessageBox.Show("Ошибка загрузки типов заказов: " + ex.Message);
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
            selectedLanguage = SelectedLanguage.RU;
            HandleLanguageButtonClick(sender as ToggleButton);
        }

        private void LvButton_Click(object sender, RoutedEventArgs e)
        {
            selectedLanguage = SelectedLanguage.LV;
            HandleLanguageButtonClick(sender as ToggleButton);
        }

        private void EngButton_Click(object sender, RoutedEventArgs e)
        {
            selectedLanguage = SelectedLanguage.ENG;
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

        private void RuButton_Checked(object sender, RoutedEventArgs e)
        {

        }

        private void LvButton_Checked(object sender, RoutedEventArgs e)
        {

        }

        private void EngButton_Checked(object sender, RoutedEventArgs e)
        {

        }

        private void CancelButton_Click(object sender, RoutedEventArgs e)
        {
            EditOrders editOrders = new EditOrders();
            editOrders.Show();
            Close();
        }

        private async void SaveButton_Click(object sender, RoutedEventArgs e)
        {
            try
            {
                if (DoneCheck.IsChecked == true)
                {
                    order.OrderStatus.ReadyAt = DateTime.Now;
                }

                if (CalledBackCheck.IsChecked == true)
                {
                    order.OrderStatus.CustomerNotifiedAt = DateTime.Now;
                }
                else
                {
                    order.OrderStatus.CustomerNotifiedAt = null;
                }

                order.OrderStatus.IsOutsourced = (bool)OutsourceCheck.IsChecked;
                order.WorkerId = (int)EmployeeNameComboBox.SelectedValue;
                order.OrderTypeId = (int)OrderTypeComboBox.SelectedValue;
                order.Customer.PhoneNumber = ClientPhoneTextBox.Text;
                order.Reason = ReasonTextBox.Text;
                order.ItemName = ProductModelTextBox.Text;
                order.Defect = DefectDescriptionTextBox.Text;
                order.TotalPrice = double.Parse(TotalCostTextBox.Text);
                order.Prepayment = double.Parse(PrepaymentTextBox.Text);
                order.Customer.LanguageId = (int)selectedLanguage;

                //загрузить все данные в ордер 

                var response = await App.HttpClient.PutAsJsonAsync($"/api/orders", order);
                if (!response.IsSuccessStatusCode)
                {
                    MessageBox.Show("Ошибка сохранения изменений: " + response.ReasonPhrase);
                    return;
                }

                MessageBox.Show("Изменения успешно сохранены!");
            }

            catch (Exception ex)
            {
                MessageBox.Show($"Ошибка при сохранении изменений: {ex.Message}");
            }

        }

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


        private void TextBox_PreviewTextInput(object sender, TextCompositionEventArgs e)
        {
            // Проверка, что вводимые символы — это только цифры
            e.Handled = !IsTextNumeric(e.Text);
        }

        private static bool IsTextNumeric(string text)
        {
            Regex regex = new Regex("^[0-9]+$");
            return regex.IsMatch(text);
        }
    }
}

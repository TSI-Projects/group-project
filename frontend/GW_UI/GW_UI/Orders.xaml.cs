using System;
using System.Collections.Generic;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Input;
using System.Windows.Controls.Primitives;
using System.Net.Http.Json;
using System.Collections.ObjectModel;
using System.Text.RegularExpressions;
using GW_UI.Classes;

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

        public enum SelectedLanguage
        {
            RU = 1, LV = 2, ENG = 3
        }

        private SelectedLanguage selectedLanguage = SelectedLanguage.RU;

        private async void OrdersWindow_Loaded(object sender, RoutedEventArgs e)
        {
            try
            {
                var orderTypes = await App.HttpClient.GetFromJsonAsync<TypeResponse>("/api/orders/types"); //Поправить
                if (!orderTypes.Success && orderTypes.Error != null)
                {
                    throw new Exception(orderTypes.Error.Message);
                }

                foreach (TypeItem type in orderTypes.Types)
                {
                    OrderTypes.Add(type);
                }
                var employees = await App.HttpClient.GetFromJsonAsync<EmployeeResponse>("/api/workers"); //Поправить
                if (!employees.Success && employees.Error != null)
                {
                    throw new Exception(employees.Error.Message);
                }

                foreach (Employee emp in employees.Workers)
                {
                    Employees.Add(emp);
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

        private async void AddOrder_Click(object sender, RoutedEventArgs e)
        {
            // Проверка на заполненность обязательных полей
            if (OrderTypeComboBox.SelectedValue == null ||
                EmployeeNameComboBox.SelectedValue == null ||
                string.IsNullOrWhiteSpace(ProductModelTextBox.Text) ||
                string.IsNullOrWhiteSpace(ClientPhoneTextBox.Text) ||
                string.IsNullOrWhiteSpace(ReasonTextBox.Text) ||
                string.IsNullOrWhiteSpace(DefectDescriptionTextBox.Text) ||
                string.IsNullOrWhiteSpace(TotalCostTextBox.Text) ||
                string.IsNullOrWhiteSpace(PrepaymentTextBox.Text))
            {
                MessageBox.Show("All fields must be filled in before saving the order.", "Error", MessageBoxButton.OK, MessageBoxImage.Warning);
                return;
            }

            var customer = new Customer
            {
                PhoneNumber = ClientPhoneTextBox.Text,
                LanguageId = (int)selectedLanguage
            };

            DateTime utcDate = DateTime.SpecifyKind(DateTime.Today, DateTimeKind.Utc);

            var orderRequest = new Order
            {
                OrderTypeId = (int)OrderTypeComboBox.SelectedValue,
                WorkerId = (int)EmployeeNameComboBox.SelectedValue,
                ItemName = ProductModelTextBox.Text,
                Customer = customer,
                Reason = ReasonTextBox.Text,
                Defect = DefectDescriptionTextBox.Text,
                TotalPrice = double.Parse(TotalCostTextBox.Text),
                Prepayment = double.Parse(PrepaymentTextBox.Text),
                CreatedAt = utcDate
            };

            try
            {
                var response = await App.HttpClient.PostAsJsonAsync("/api/orders", orderRequest);
                var body = await response.Content.ReadFromJsonAsync<OrderResponse>();

                if (!body.Success && body.Error != null)
                {
                    throw new Exception(body.Error.Message);
                }
                MessageBox.Show("Order added successfully!");
                ClearInputFields();
            }
            catch (Exception ex)
            {
                MessageBox.Show(ex.Message);
            }
        }

        private void ClearInputFields()
        {
            // Очистка полей ввода
            ClientPhoneTextBox.Text = string.Empty;
            ReasonTextBox.Text = string.Empty;
            DefectDescriptionTextBox.Text = string.Empty;
            TotalCostTextBox.Text = string.Empty;
            PrepaymentTextBox.Text = string.Empty;
            ProductModelTextBox.Text = string.Empty;

            // Сброс выбранных значений в ComboBox
            OrderTypeComboBox.SelectedIndex = -1; //clear
            EmployeeNameComboBox.SelectedIndex = -1;

            foreach (var textBox in new TextBox[] { ClientPhoneTextBox, ReasonTextBox, DefectDescriptionTextBox, TotalCostTextBox, PrepaymentTextBox })
            {
                AddText(textBox, null);
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

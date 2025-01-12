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
using System.Text.RegularExpressions;

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
                CreatedAt = utcDate // Передаем DateTime
            };

            try
            {
                var response = await App.HttpClient.PostAsJsonAsync("/api/orders", orderRequest);
                if (response.IsSuccessStatusCode)
                {
                    MessageBox.Show("Заказ успешно добавлен!");
                    ClearInputFields();
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

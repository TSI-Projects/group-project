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
        public ObservableCollection<TypeItem> AvailableOrderTypes { get; set; }

        public Orders()
        {
            InitializeComponent();
            AvailableOrderTypes = new ObservableCollection<TypeItem>();
            OrderTypeComboBox.ItemsSource = AvailableOrderTypes; // Убедитесь, что элемент управления уже инициализирован
            this.Loaded += OrdersWindow_Loaded;
        }

        private async void OrdersWindow_Loaded(object sender, RoutedEventArgs e)
        {
            try
            {
                var orderTypes = await App.HttpClient.GetFromJsonAsync<List<TypeItem>>("/api/orders/types");
                if (orderTypes != null)
                {
                    foreach (TypeItem type in orderTypes)
                    {
                        AvailableOrderTypes.Add(type);
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

        // ---------------------------------------------//

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
                activeLanguageButton.IsChecked = false; // Деактивировать предыдущую кнопку
            }

            // Если текущая кнопка была активна, снять с нее активацию
            if (activeLanguageButton == clickedButton)
            {
                activeLanguageButton = null;
            }
            else
            {
                activeLanguageButton = clickedButton; // Сделать новую кнопку активной
            }
        }

        private void AddOrder_Click(object sender, RoutedEventArgs e)
        {

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
    }
}

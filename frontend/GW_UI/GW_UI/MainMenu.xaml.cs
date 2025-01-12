using System.Windows;

namespace GW_UI
{
    public partial class MainMenu : Window
    {
        public MainMenu()
        {
            InitializeComponent();
        }

        public void LogoutButton_Click(object sender, RoutedEventArgs e)
        {
            Application.Current.Shutdown();
        }

        private void SettingsButton_Click(object sender, RoutedEventArgs e)
        {
            Menu menuPage = new Menu();
            menuPage.Show();
            Close();
        }

        private void CreateOrderButton_Click(object sender, RoutedEventArgs e)
        {
            Orders orderPage = new Orders();
            orderPage.Show();
            Close();
        }

        private void OrderListButton_Click(object sender, RoutedEventArgs e)
        {
            EditOrders editOrderPage = new EditOrders();
            editOrderPage.Show();
            Close();
        }

        private void Button_Click(object sender, RoutedEventArgs e)
        {
            CompletedOrders completedOrders = new CompletedOrders();
            completedOrders.Show();
            Close();
        }
    }
}

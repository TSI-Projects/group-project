using System.Windows;

namespace GW_UI
{
    public partial class Menu : Window
    {
        public Menu()
        {
            InitializeComponent();
        }

        private void BackButton_Click(object sender, RoutedEventArgs e)
        {
            MainMenu mainMenu = new MainMenu();
            mainMenu.Show();
            this.Close();
        }

        public void LogoutButton_Click(object sender, RoutedEventArgs e)
        {
            Application.Current.Shutdown();
        }

        private void EmployeesButton_Click(object sender, RoutedEventArgs e)
        {
            Employees employees = new Employees();
            employees.Show();
            this.Close();
        }

        private void TypesButton_Click(object sender, RoutedEventArgs e)
        {
            Types types = new Types();
            types.Show();
            this.Close();
        }

    }
}

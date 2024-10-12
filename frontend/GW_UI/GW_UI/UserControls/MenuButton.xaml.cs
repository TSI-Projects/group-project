using System.Windows;
using System.Windows.Controls;
using System.Windows.Input;

namespace GW_UI.UserControls
{
    public partial class MenuButton : UserControl
    {
        // Регистрация RoutedEvent для Click
        public static readonly RoutedEvent ClickEvent = EventManager.RegisterRoutedEvent(
            "Click", // Имя события
            RoutingStrategy.Bubble, // Стратегия маршрутизации
            typeof(RoutedEventHandler), // Тип делегата
            typeof(MenuButton)); // Владелец события

        // CLR-событие обертка для RoutedEvent
        public event RoutedEventHandler Click
        {
            add { AddHandler(ClickEvent, value); }
            remove { RemoveHandler(ClickEvent, value); }
        }

        public MenuButton()
        {
            InitializeComponent();
            this.MouseUp += OnMouseUp;
            this.Unloaded += MenuButton_Unloaded;
        }

        private void MenuButton_Unloaded(object sender, RoutedEventArgs e)
        {
            this.MouseUp -= OnMouseUp;
            this.Unloaded -= MenuButton_Unloaded;
        }

        private void OnMouseUp(object sender, MouseButtonEventArgs e)
        {
            if (e.ChangedButton == MouseButton.Left)
            {
                RaiseEvent(new RoutedEventArgs(ClickEvent, this));
            }
        }

        public string Title
        {
            get { return (string)GetValue(TitleProperty); }
            set { SetValue(TitleProperty, value); }
        }

        public static readonly DependencyProperty TitleProperty =
            DependencyProperty.Register("Title", typeof(string), typeof(MenuButton));

        public bool IsActive
        {
            get { return (bool)GetValue(IsActiveProperty); }
            set { SetValue(IsActiveProperty, value); }
        }

        public static readonly DependencyProperty IsActiveProperty =
            DependencyProperty.Register("IsActive", typeof(bool), typeof(MenuButton));

        public MahApps.Metro.IconPacks.PackIconMaterialKind Icon
        {
            get { return (MahApps.Metro.IconPacks.PackIconMaterialKind)GetValue(IconProperty); }
            set { SetValue(IconProperty, value); }
        }

        public static readonly DependencyProperty IconProperty =
            DependencyProperty.Register("Icon", typeof(MahApps.Metro.IconPacks.PackIconMaterialKind), typeof(MenuButton));
    }
}

using System;
using System.Globalization;
using System.Windows.Data;

namespace GW_UI
{
    public class LanguageConverter : IValueConverter
    {
        public object Convert(object value, Type targetType, object parameter, CultureInfo culture)
        {
            int languageId = (int)value;
            switch (languageId)
            {
                case 1:
                    return "Russian";
                case 2:
                    return "Latvian";
                case 3:
                    return "English";
                default:
                    return "Unknown";
            }
        }

        public object ConvertBack(object value, Type targetType, object parameter, CultureInfo culture)
        {
            throw new NotSupportedException("This converter only works for one-way binding.");
        }
    }
}

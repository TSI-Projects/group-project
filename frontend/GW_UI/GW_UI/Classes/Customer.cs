using System.Text.Json.Serialization;

namespace GW_UI
{
    public class Customer
    {
        [JsonPropertyName("id")]
        public int Id { get; set; }

        [JsonPropertyName("language_id")]
        public int LanguageId { get; set; }

        [JsonPropertyName("phone_number")]
        public string PhoneNumber { get; set; }
    }
}

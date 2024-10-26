using System.Text.Json.Serialization;

namespace GW_UI
{
    public class TypeItem
    {
        [JsonPropertyName("id")]
        public int ID { get; set; }

        [JsonPropertyName("full_name")]
        public string TypeName { get; set; }
    }
}

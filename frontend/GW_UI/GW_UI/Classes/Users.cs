using System.Text.Json.Serialization;

namespace GW_UI
{
    internal class Users
    {
        [JsonPropertyName("username")]
        public string Username { get; set; }

        [JsonPropertyName("password")]
        public string Password { get; set; }
    }
}

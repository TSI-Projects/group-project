using System.Text.Json.Serialization;

namespace GW_UI
{
    public class Employee
    {
        public Employee(string firstName, string lastName)
        {
            FirstName = firstName;
            LastName = lastName;
        }

        [JsonPropertyName("id")]
        public int ID { get; set; }

        [JsonPropertyName("first_name")]
        public string FirstName { get; set; }

        [JsonPropertyName("last_name")]
        public string LastName { get; set; }
    }
}

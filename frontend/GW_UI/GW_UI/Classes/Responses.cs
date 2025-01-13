using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.Json.Serialization;
using System.Threading.Tasks;

namespace GW_UI.Classes
{
    internal class OrderResponse
    {
        [JsonPropertyName("order")]
        public Order Order { get; set; }

        [JsonPropertyName("orders")]
        public List<Order> Orders { get; set; }

        [JsonPropertyName("success")]
        public bool Success { get; set; }

        [JsonPropertyName("error")]
        public Error Error { get; set; }
    }
    internal class EmployeeResponse
    {
        [JsonPropertyName("worker")]
        public Employee Worker { get; set; }

        [JsonPropertyName("workers")]
        public List<Employee> Workers { get; set; }

        [JsonPropertyName("success")]
        public bool Success { get; set; }

        [JsonPropertyName("error")]
        public Error Error { get; set; }
    }
    internal class TypeResponse
    {
        [JsonPropertyName("order_type")]
        public TypeItem Type { get; set; }

        [JsonPropertyName("order_types")]
        public List<TypeItem> Types { get; set; }

        [JsonPropertyName("success")]
        public bool Success { get; set; }

        [JsonPropertyName("error")]
        public Error Error { get; set; }
    }
}

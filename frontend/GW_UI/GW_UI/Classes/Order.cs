using System;
using System.Text.Json.Serialization;

namespace GW_UI
{
    public class Order
    {
        [JsonPropertyName("id")]
        public int Id { get; set; }

        [JsonPropertyName("order_status_id")]
        public int OrderStatusId { get; set; }

        [JsonPropertyName("order_type_id")]
        public int OrderTypeId { get; set; }

        [JsonPropertyName("worker_id")]
        public int WorkerId { get; set; }

        [JsonPropertyName("customer")]
        public Customer Customer { get; set; }

        [JsonPropertyName("reason")]
        public string Reason { get; set; }

        [JsonPropertyName("defect")]
        public string Defect { get; set; }

        [JsonPropertyName("total_price")]
        public double TotalPrice { get; set; }

        [JsonPropertyName("prepayment")]
        public double Prepayment { get; set; }

        [JsonPropertyName("created_at")]
        public DateTime CreatedAt { get; set; }

        [JsonPropertyName("item_name")]
        public string ItemName { get; set; }

        // Добавление статуса заказа как свойства класса
        [JsonPropertyName("status")]
        public OrderStatus Status { get; set; }

        public TypeItem TypeItem { get; set; }
        public Employee Employee { get; set; }
    }
}

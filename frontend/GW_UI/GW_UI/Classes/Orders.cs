using System;
using System.Text.Json.Serialization;

namespace GW_UI
{
    //public class Order
    //{
    //    [JsonPropertyName("id")]
    //    public int Id { get; set; }

    //    [JsonPropertyName("order_type_id")]
    //    public int OrderTypeId { get; set; }

    //    [JsonPropertyName("worker_id")]
    //    public int WorkerId { get; set; }

    //    [JsonPropertyName("customer_id")]
    //    public int CustomerId { get; set; }

    //    [JsonPropertyName("reason")]
    //    public string Reason { get; set; }

    //    [JsonPropertyName("defect")]
    //    public string Defect { get; set; }

    //    [JsonPropertyName("total_price")]
    //    public double TotalPrice { get; set; }

    //    [JsonPropertyName("prepayment")]
    //    public double Prepayment { get; set; }
    //}

    public class Order
    {
        [JsonPropertyName("id")]
        public int Id { get; set; }

        [JsonPropertyName("order_type_id")]
        public int OrderTypeId { get; set; }

        [JsonPropertyName("worker_id")]
        public int WorkerId { get; set; }

        [JsonPropertyName("customer_id")]
        public double CustomerId { get; set; }

        [JsonPropertyName("language_id")]
        public int LanguageId { get; set; }

        [JsonPropertyName("reason")]
        public string Reason { get; set; }

        [JsonPropertyName("defect")]
        public string Defect { get; set; }

        [JsonPropertyName("total_price")]
        public double TotalPrice { get; set; }

        [JsonPropertyName("prepayment")]
        public double Prepayment { get; set; }
    }
}

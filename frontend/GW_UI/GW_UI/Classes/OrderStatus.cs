using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.Json.Serialization;
using System.Threading.Tasks;

namespace GW_UI
{
    public class OrderStatus
    {
        [JsonPropertyName("id")]
        public int Id { get; set; }

        [JsonPropertyName("ready_at")]
        public DateTime? ReadyAt { get; set; }

        [JsonPropertyName("returned_at")]
        public DateTime? ReturnedAt { get; set; }

        [JsonPropertyName("customer_notified_at")]
        public DateTime? CustomerNotifiedAt { get; set; }
        public bool IsCustomerNotified => CustomerNotifiedAt != null;

        [JsonPropertyName("is_outsourced")]
        public bool IsOutsourced { get; set; }

        [JsonPropertyName("is_recipient_lost")]
        public bool IsRecipientLost { get; set; }
    }
}

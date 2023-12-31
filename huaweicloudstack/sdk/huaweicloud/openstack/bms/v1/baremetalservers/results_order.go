package baremetalservers

import (
	"github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/sdk/huaweicloud"
)

type OrderResponse struct {
	OrderID string `json:"order_id"`
	JobID   string `json:"job_id"`
}

type OrderResult struct {
	golangsdk.Result
}

func (r OrderResult) ExtractOrderResponse() (*OrderResponse, error) {
	order := new(OrderResponse)
	err := r.ExtractInto(order)
	return order, err
}

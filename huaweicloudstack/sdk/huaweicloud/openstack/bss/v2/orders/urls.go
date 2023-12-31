package orders

import "github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/sdk/huaweicloud"

func unsubscribeURL(sc *golangsdk.ServiceClient) string {
	return sc.ServiceURL("orders/subscriptions/resources/unsubscribe")
}

func getURL(sc *golangsdk.ServiceClient, id string) string {
	return sc.ServiceURL("orders/customer-orders/details", id)
}

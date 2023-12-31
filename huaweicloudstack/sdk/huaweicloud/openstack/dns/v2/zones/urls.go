package zones

import "github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/sdk/huaweicloud"

func baseURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("zones")
}

func zoneURL(c *golangsdk.ServiceClient, zoneID string) string {
	return c.ServiceURL("zones", zoneID)
}

func associateURL(client *golangsdk.ServiceClient, zoneID string) string {
	return client.ServiceURL("zones", zoneID, "associaterouter")
}

func disassociateURL(client *golangsdk.ServiceClient, zoneID string) string {
	return client.ServiceURL("zones", zoneID, "disassociaterouter")
}

package drill

import "github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/sdk/huaweicloud"

func rootURL(c *golangsdk.ServiceClient) string {
	return c.ServiceURL("disaster-recovery-drills")
}

func resourceURL(c *golangsdk.ServiceClient, id string) string {
	return c.ServiceURL("disaster-recovery-drills", id)
}

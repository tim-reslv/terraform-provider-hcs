package networks

import "github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/sdk/huaweicloud"

func rootURL(client *golangsdk.ServiceClient, ns string) string {
	return client.ServiceURL("namespaces", ns, "networks")
}

func resourceURL(client *golangsdk.ServiceClient, ns, name string) string {
	return client.ServiceURL("namespaces", ns, "networks", name)
}

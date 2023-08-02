---
subcategory: "Global Accelerator (GA)"
---

# huaweicloud_ga_health_check

Manages a GA health check resource within HuaweiCloud.

## Example Usage

```hcl
variable "endpoint_group_id" {}

resource "huaweicloud_ga_health_check" "test" {
  endpoint_group_id = var.endpoint_group_id
  enabled           = true
  interval          = 10
  max_retries       = 5
  port              = 8001
  timeout           = 10
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the resource.
  If omitted, the provider-level region will be used. Changing this parameter will create a new resource.

* `endpoint_group_id` - (Required, String, ForceNew) Specifies the endpoint group ID.

  Changing this parameter will create a new resource.

* `enabled` - (Required, Bool) Specifies whether to enable health check.

* `interval` - (Required, Int) Specifies the health check interval, in seconds.

* `max_retries` - (Required, Int) Specifies the maximum number of retries.
  Specifies the number of consecutive health checks when the health check result of an endpoint changes
  from **HEALTHY** to **UNHEALTHY**, or from **UNHEALTHY** to **HEALTHY**.

* `port` - (Required, Int) Specifies the port used for the health check.

* `timeout` - (Required, Int) Specifies the timeout duration of the health check, in seconds.
  It is recommended that you set a value less than that of parameter **interval**.

* `protocol` - (Optional, String) Specifies the health check protocol.
  Only **TCP** supported for now. Defaults to **TCP**.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The resource ID.

* `status` - Indicates the provisioning status. The value can be one of the following:
  + **ACTIVE**: The resource is running.
  + **PENDING**: The status is to be determined.
  + **ERROR**: Failed to create the resource.
  + **DELETING**: The resource is being deleted.

* `created_at` - Indicates when the health check was configured.

* `updated_at` - Indicates when the health check was updated.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minutes.
* `update` - Default is 10 minutes.
* `delete` - Default is 10 minutes.

## Import

The healthcheck can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_ga_health_check.test 30a04052-b553-4292-af4e-a483106653ce
```

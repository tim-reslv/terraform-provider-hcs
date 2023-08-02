---
subcategory: "Web Application Firewall (WAF)"
---

# huaweicloud_waf_policy

Manages a WAF policy resource within HuaweiCloud.

-> **NOTE:** All WAF resources depend on WAF instances, and the WAF instances need to be purchased before they can be
used. The policy resource can be used in Cloud Mode, Dedicated Mode and ELB Mode.

## Example Usage

```hcl
variable "enterprise_project_id" {}

resource "huaweicloud_waf_policy" "policy_1" {
  name                  = "policy_1"
  protection_mode       = "log"
  level                 = 2
  enterprise_project_id = var.enterprise_project_id
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the WAF policy resource. If omitted, the
  provider-level region will be used. Changing this setting will push a new certificate.

* `name` - (Required, String) Specifies the policy name. The maximum length is 256 characters. Only digits, letters,
  underscores(_), and hyphens(-) are allowed.

* `protection_mode` - (Optional, String) Specifies the protective action after a rule is matched. Defaults to `log`.
  Valid values are:
  + `block`: WAF blocks and logs detected attacks.
  + `log`: WAF logs detected attacks only.

* `level` - (Optional, Int) Specifies the protection level. Defaults to `2`. Valid values are:
  + `1`: low
  + `2`: medium
  + `3`: high

* `enterprise_project_id` - (Optional, String, ForceNew) Specifies the enterprise project ID of WAF policy.
  Changing this parameter will create a new resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The policy ID in UUID format.

* `full_detection` - The detection mode in Precise Protection.
  + `true`: full detection, Full detection finishes all threat detections before blocking requests that meet Precise
      Protection specified conditions.
  + `false`: instant detection. Instant detection immediately ends threat detection after blocking a request that
      meets Precise Protection specified conditions.

* `options` - The protection switches. The options object structure is documented below.

The `options` block supports:

* `basic_web_protection` - Indicates whether Basic Web Protection is enabled.

* `general_check` - Indicates whether General Check in Basic Web Protection is enabled.

* `crawler` - Indicates whether the master crawler detection switch in Basic Web Protection is enabled.

* `crawler_engine` - Indicates whether the Search Engine switch in Basic Web Protection is enabled.

* `crawler_scanner` - Indicates whether the Scanner switch in Basic Web Protection is enabled.

* `crawler_script` - Indicates whether the Script Tool switch in Basic Web Protection is enabled.

* `crawler_other` - Indicates whether detection of other crawlers in Basic Web Protection is enabled.

* `webshell` - Indicates whether webshell detection in Basic Web Protection is enabled.

* `cc_attack_protection` - Indicates whether CC Attack Protection is enabled.

* `precise_protection` - Indicates whether Precise Protection is enabled.

* `blacklist` - Indicates whether Blacklist and Whitelist is enabled.

* `data_masking` - Indicates whether Data Masking is enabled.

* `false_alarm_masking` - Indicates whether False Alarm Masking is enabled.

* `web_tamper_protection` - Indicates whether Web Tamper Protection is enabled.

## Import

There are two ways to import WAF policy state.

* Using the `id`, e.g.

```bash
$ terraform import huaweicloud_waf_policy.test <id>
```

* Using `id` and `enterprise_project_id`, separated by a slash, e.g.

```bash
$ terraform import huaweicloud_waf_policy.test <id>/<enterprise_project_id>
```

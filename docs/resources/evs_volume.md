---
subcategory: "Elastic Volume Service (EVS)"
---

# huaweicloud_evs_volume

Manages a volume resource within HuaweiCloud.

## Example Usage

```hcl
resource "huaweicloud_evs_volume" "volume" {
  name              = "volume"
  description       = "my volume"
  volume_type       = "SAS"
  size              = 20
  availability_zone = "cn-north-4a"

  tags = {
    foo = "bar"
    key = "value"
  }
}
```

## Example Usage with KMS encryption

```hcl
resource "huaweicloud_evs_volume" "volume" {
  name              = "volume"
  description       = "my volume"
  volume_type       = "SAS"
  size              = 20
  kms_id            = var.kms_id
  availability_zone = "cn-north-4a"

  tags = {
    foo = "bar"
    key = "value"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) Specifies the region in which to create the disk. If omitted, the
  provider-level region will be used. Changing this creates a new disk.

* `availability_zone` - (Required, String, ForceNew) Specifies the availability zone for the disk. Changing this creates
  a new disk.

* `volume_type` - (Required, String, ForceNew) Specifies the disk type. Currently, the value can be SAS, SSD, GPSSD or
  ESSD.
  + SAS: specifies the high I/O disk type.
  + SSD: specifies the ultra-high I/O disk type.
  + GPSSD: specifies the general purpose SSD disk type.
  + ESSD: Extreme SSD type.

      If the specified disk type is not available in the AZ, the disk will fail to create. Changing this creates a new
      disk.

* `name` - (Optional, String) Specifies the disk name. The value can contain a maximum of 255 bytes.

* `size` - (Optional, Int) Specifies the disk size, in GB. The valid value is range from:
  + System disk: 1 GB to 1024 GB
  + Data disk: 10 GB to 32768 GB

  This parameter is required when:
  + Create an empty disk.
  + Create the disk from a snapshot. The disk size must be greater than or equal to the snapshot size.
  + Create the disk from an image. The disk size must be greater than or equal to the minimum disk capacity required by
  min_disk in the image attributes.

  This parameter is optional when you create the disk from a backup. If this parameter is not specified, the
  disk size is equal to the backup size.

  -> **NOTE:** Shrinking the disk is not supported.

* `description` - (Optional, String) Specifies the disk description. The value can contain a maximum of 255 bytes.

* `image_id` - (Optional, String, ForceNew) Specifies the image ID from which to create the disk. Changing this creates
  a new disk.

* `backup_id` - (Optional, String, ForceNew) Specifies the backup ID from which to create the disk. Changing this
  creates a new disk.

* `snapshot_id` - (Optional, String, ForceNew) Specifies the snapshot ID from which to create the disk. Changing this
  creates a new disk.

* `tags` - (Optional, Map) Specifies the key/value pairs to associate with the disk.

* `multiattach` - (Optional, Bool, ForceNew) Specifies whether the disk is shareable. The default value is false.
  Changing this creates a new disk.

* `kms_id` - (Optional, String, ForceNew) Specifies the Encryption KMS ID to create the disk. Changing this creates a
  new disk.

* `device_type` - (Optional, String, ForceNew) Specifies the device type of disk to create. Valid options are VBD and
  SCSI. Defaults to VBD. Changing this creates a new disk.

* `dedicated_storage_id` - (Optional, String, ForceNew) Specifies the ID of the DSS storage pool accommodating the disk.

* `enterprise_project_id` - (Optional, String, ForceNew) Specifies the enterprise project id of the disk. Changing this
  creates a new disk.

* `cascade` - (Optional, Bool) Specifies the delete mode of snapshot. The default value is false. All snapshot
  associated with the disk will also be deleted when the parameter is set to true.

  -> This parameter is only valid for pay-as-you-go resources, and the snapshots bound to the package period resources
     will be removed while resources unsubscribed.

* `charging_mode` - (Optional, String, ForceNew) Specifies the charging mode of the disk.
  The valid values are as follows:
  + **prePaid**: the yearly/monthly billing mode.
  + **postPaid**: the pay-per-use billing mode.
    Changing this creates a new disk.

* `period_unit` - (Optional, String, ForceNew) Specifies the charging period unit of the disk.
  Valid values are **month** and **year**. This parameter is mandatory if `charging_mode` is set to **prePaid**.
  Changing this creates a new disk.

* `period` - (Optional, Int, ForceNew) Specifies the charging period of the disk.
  If `period_unit` is set to **month**, the value ranges from 1 to 9.
  If `period_unit` is set to **year**, the valid value is 1.
  This parameter is mandatory if `charging_mode` is set to **prePaid**.
  Changing this creates a new disk.

* `auto_renew` - (Optional, String) Specifies whether auto renew is enabled.
  Valid values are **true** and **false**.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - A resource ID in UUID format.
* `attachment` - If a disk is attached to an instance, this attribute will display the Attachment ID, Instance ID, and
  the Device as the Instance sees it. The [object](#attachment_struct) structure is documented below.
* `wwn` - The unique identifier used for mounting the EVS disk.

<a name="attachment_struct"></a>
The `attachment` block supports:

* `id` - The ID of the attachment information.

* `instance_id` - The ID of the server to which the disk is attached.

* `device` - The device name.

* `dedicated_storage_name` - The name of the DSS storage pool accommodating the disk.

## Import

Volumes can be imported using the `id`, e.g.

```
$ terraform import huaweicloud_evs_volume.volume_1 14a80bc7-c12c-4fe0-a38a-cb77eeac9bd6
```

Note that the imported state may not be identical to your resource definition, due to some attributes missing from the
API response, security or some other reason. The missing attributes include: **cascade**, **period_unit**, **period**
and **auto_renew**. It is generally recommended running terraform plan after importing an disk.
You can then decide if changes should be applied to the disk, or the resource definition should be updated to align
with the disk. Also you can ignore changes as below.

```
resource "huaweicloud_evs_volume" "volume_1" {
    ...

  lifecycle {
    ignore_changes = [
      cascade,
    ]
  }
}
```

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 10 minute.
* `update` - Default is 3 minute.
* `delete` - Default is 3 minute.

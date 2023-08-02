---
subcategory: "MapReduce Service (MRS)"
---

# huaweicloud_mapreduce_cluster

Manages a cluster resource within HuaweiCloud MRS.

## Example Usage

### Create an analysis cluster

```hcl
data "huaweicloud_availability_zones" "test" {}

variable "cluster_name" {}
variable "password" {}
variable "vpc_id" {}
variable "subnet_id" {}

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = var.cluster_name
  version            = "MRS 1.9.2"
  type               = "ANALYSIS"
  component_list     = ["Hadoop", "Hive", "Tez"]
  manager_admin_pass = var.password
  node_admin_pass    = var.password
  vpc_id             = var.vpc_id
  subnet_id          = var.subnet_id

  master_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  analysis_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  analysis_task_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 1
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }

  tags = {
    foo = "bar"
    key = "value"
  }
}
```

### Create a stream cluster

```hcl
data "huaweicloud_availability_zones" "test" {}

variable "cluster_name" {}
variable "password" {}
variable "vpc_id" {}
variable "subnet_id" {}

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = var.cluster_name
  type               = "STREAMING"
  version            = "MRS 1.9.2"
  manager_admin_pass = var.password
  node_admin_pass    = var.password
  vpc_id             = var.vpc_id
  subnet_id          = var.subnet_id
  component_list     = ["Storm"]

  master_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }

  tags = {
    foo = "bar"
    key = "value"
  }
}
```

### Create a hybrid cluster

```hcl
data "huaweicloud_availability_zones" "test" {}

variable "cluster_name" {}
variable "password" {}
variable "vpc_id" {}
variable "subnet_id" {}

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = var.cluster_name
  version            = "MRS 1.9.2"
  type               = "MIXED"
  component_list     = ["Hadoop", "Spark", "Hive", "Tez", "Storm"]
  manager_admin_pass = var.password
  node_admin_pass    = var.password
  vpc_id             = var.vpc_id
  subnet_id          = var.subnet_id

  master_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  analysis_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  analysis_task_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 1
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  streaming_task_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 1
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }

  tags = {
    foo = "bar"
    key = "value"
  }
}
```

### Create a custom cluster

```hcl
data "huaweicloud_availability_zones" "test" {}

variable "cluster_name" {}
variable "password" {}
variable "vpc_id" {}
variable "subnet_id" {}

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = var.cluster_name
  version            = "MRS 3.1.0"
  type               = "CUSTOM"
  safe_mode          = true
  manager_admin_pass = var.password
  node_admin_pass    = var.password
  vpc_id             = var.vpc_id
  subnet_id          = var.subnet_id
  template_id        = "mgmt_control_combined_v4"
  component_list     = ["DBService", "Hadoop", "ZooKeeper", "Ranger"]

  master_nodes {
    flavor            = "c6.4xlarge.4.linux.bigdata"
    node_number       = 3
    root_volume_type  = "SAS"
    root_volume_size  = 480
    data_volume_type  = "SAS"
    data_volume_size  = 600
    data_volume_count = 1
    assigned_roles = [
      "OMSServer:1,2",
      "SlapdServer:1,2",
      "KerberosServer:1,2",
      "KerberosAdmin:1,2",
      "quorumpeer:1,2,3",
      "NameNode:2,3",
      "Zkfc:2,3",
      "JournalNode:1,2,3",
      "ResourceManager:2,3",
      "JobHistoryServer:3",
      "DBServer:1,3",
      "HttpFS:1,3",
      "TimelineServer:3",
      "RangerAdmin:1,2",
      "UserSync:2",
      "TagSync:2",
      "KerberosClient",
      "SlapdClient",
      "meta"
    ]
  }

  custom_nodes {
    group_name        = "node_group_1"
    flavor            = "c6.4xlarge.4.linux.bigdata"
    node_number       = 4
    root_volume_type  = "SAS"
    root_volume_size  = 480
    data_volume_type  = "SAS"
    data_volume_size  = 600
    data_volume_count = 1
    assigned_roles = [
      "DataNode",
      "NodeManager",
      "KerberosClient",
      "SlapdClient",
      "meta"
    ]
  }
}

```

### Create an analysis cluster and bind public IP

```hcl
data "huaweicloud_availability_zones" "test" {}

variable "cluster_name" {}
variable "password" {}
variable "vpc_id" {}
variable "subnet_id" {}
variable "public_ip" {}

resource "huaweicloud_mapreduce_cluster" "test" {
  availability_zone  = data.huaweicloud_availability_zones.test.names[0]
  name               = var.cluster_name
  version            = "MRS 1.9.2"
  type               = "ANALYSIS"
  component_list     = ["Hadoop", "Hive", "Tez"]
  manager_admin_pass = var.password
  node_admin_pass    = var.password
  vpc_id             = var.vpc_id
  subnet_id          = var.subnet_id
  public_ip          = var.public_ip

  master_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  analysis_core_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 2
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }
  analysis_task_nodes {
    flavor            = "c6.2xlarge.4.linux.bigdata"
    node_number       = 1
    root_volume_type  = "SAS"
    root_volume_size  = 300
    data_volume_type  = "SAS"
    data_volume_size  = 480
    data_volume_count = 1
  }

  tags = {
    foo = "bar"
    key = "value"
  }
}
```

## Argument Reference

The following arguments are supported:

* `region` - (Optional, String, ForceNew) The region in which to create the MapReduce cluster resource. If omitted, the
  provider-level region will be used. Changing this will create a new MapReduce cluster resource.

* `availability_zone` - (Required, String, ForceNew) Specifies the availability zone in which to create the cluster.
  Please following [reference](https://developer.huaweicloud.com/intl/en-us/endpoint?all)
  Changing this will create a new MapReduce cluster resource.

* `name` - (Required, String) Specifies the name of the MapReduce cluster. The name can contain 2 to 64
  characters, which may consist of letters, digits, underscores (_) and hyphens (-).

* `version` - (Required, String, ForceNew) Specifies the MapReduce cluster version. The valid values are `MRS 1.9.2`
  , `MRS 3.0.5` and `MRS 3.1.0`. Changing this will create a new MapReduce cluster resource.

* `component_list` - (Required, List, ForceNew) Specifies the list of component names. For the components supported by
  the cluster, please following [reference](https://support.huaweicloud.com/intl/en-us/productdesc-mrs/mrs_08_0005.html)
  Changing this will create a new MapReduce cluster resource.

* `master_nodes` - (Required, List, ForceNew) Specifies the informations about master nodes in the MapReduce cluster.
  The `nodes` object structure of the `master_nodes` is documented below.
  Changing this will create a new MapReduce cluster resource.

* `manager_admin_pass` - (Required, String, ForceNew) Specifies the administrator password, which is used to log in to
  the cluster management page. The password can contain 8 to 26 characters and cannot be the username or the username
  spelled backwards. The password must contain lowercase letters, uppercase letters, digits, spaces and the special
  characters: `!?,.:-_{}[]@$^+=/`. Changing this will create a new MapReduce cluster resource.

* `vpc_id` - (Required, String, ForceNew) Specifies the ID of the VPC which bound to the MapReduce cluster. Changing
  this will create a new MapReduce cluster resource.

* `subnet_id` - (Required, String, ForceNew) Specifies the network ID of a subnet which bound to the MapReduce cluster.
  Changing this will create a new MapReduce cluster resource.

* `type` - (Optional, String, ForceNew) Specifies the type of the MapReduce cluster. The valid values are **ANALYSIS***,
  **STREAMING** and **MIXED**, defaults to **ANALYSIS**. Changing this will create a new MapReduce cluster resource.

* `enterprise_project_id` - (Optional, String, ForceNew) Specifies a unique ID in UUID format of enterprise project.
  Changing this will create a new MapReduce cluster resource.

* `public_ip` - (Optional, String, ForceNew) Specifies the EIP address which bound to the MapReduce cluster.
The EIP must have been created and must be in the same region as the cluster.
 Changing this will create a new MapReduce cluster resource.

* `eip_id` - (Optional, String, ForceNew) Specifies the EIP ID which bound to the MapReduce cluster.
The EIP must have been created and must be in the same region as the cluster.
 Changing this will create a new MapReduce cluster resource.

* `log_collection` - (Optional, Bool, ForceNew) Specifies whether logs are collected when cluster installation fails.
  Defaults to true. If `log_collection` set true, the OBS buckets will be created and only used to collect logs that
  record MapReduce cluster creation failures. Changing this will create a new MapReduce cluster resource.

* `node_admin_pass` - (Optional, String, ForceNew) Specifies the administrator password, which is used to log in to the
  each nodes(/ECSs). The password can contain 8 to 26 characters and cannot be the username or the username spelled
  backwards. The password must contain lowercase letters, uppercase letters, digits, spaces and the special
  characters: `!?,.:-_{}[]@$^+=/`. Changing this will create a new MapReduce cluster resource. This parameter
  and `node_key_pair` are alternative.

* `node_key_pair` - (Optional, String, ForceNew) Specifies the name of a key pair, which is used to log in to the each
  nodes(/ECSs). Changing this will create a new MapReduce cluster resource.

* `safe_mode` - (Optional, Bool, ForceNew) Specifies whether the running mode of the MapReduce cluster is secure,
  defaults to **true**. The options are as follows:
  + **true**: enable Kerberos authentication.
  + **false**: disable Kerberos authentication. Changing this will create a new MapReduce cluster resource.

* `security_group_ids` - (Optional, List, ForceNew) Specifies an array of one or more security group ID to attach to the
  MapReduce cluster. If using the specified security group, the group need to open the specified port (9022) rules.

* `template_id` - (Optional, String, ForceNew) Specifies the template used for node deployment when the cluster type is
  **CUSTOM**. The options are as follows:
  + **mgmt_control_combined_v2**: template for jointly deploying the management and control nodes. The management and
  control roles are co-deployed on the Master node, and data instances are deployed in the same node group. This
  deployment mode applies to scenarios where the number of control nodes is less than 100, reducing costs.
  + **mgmt_control_separated_v2**: The management and control roles are deployed on different master nodes, and data
  instances are deployed in the same node group. This deployment mode is applicable to a cluster with 100 to 500 nodes
  and delivers better performance in high-concurrency load scenarios.
  + **mgmt_control_data_separated_v2**: The management role and control role are deployed on different Master nodes,
  and data instances are deployed in different node groups. This deployment mode is applicable to a cluster with more
  than 500 nodes. Components can be deployed separately, which can be used for a larger cluster scale.

* `analysis_core_nodes` - (Optional, List, ForceNew) Specifies the informations about analysis core nodes in the
  MapReduce cluster.
  The `nodes` object structure of the `analysis_core_nodes` is documented below.
  Changing this will create a new MapReduce cluster resource.

* `streaming_core_nodes` - (Optional, List, ForceNew) Specifies the informations about streaming core nodes in the
  MapReduce cluster.
  The `nodes` object structure of the `streaming_core_nodes` is documented below.
  Changing this will create a new MapReduce cluster resource.

* `analysis_task_nodes` - (Optional, List, ForceNew) Specifies the informations about analysis task nodes in the
  MapReduce cluster.
  The `nodes` object structure of the `analysis_task_nodes` is documented below.
  Changing this will create a new MapReduce cluster resource.

* `streaming_task_nodes` - (Optional, List, ForceNew) Specifies the informations about streaming task nodes in the
  MapReduce cluster.
  The `nodes` object structure of the `streaming_task_nodes` is documented below.
  Changing this will create a new MapReduce cluster resource.

* `custom_nodes` - (Optional, List, ForceNew) Specifies the informations about custom nodes in the MapReduce cluster.
  The `nodes` object structure of the `custom_nodes` is documented below.
  Changing this will create a new MapReduce cluster resource.

* `component_configs` - (Optional, List, ForceNew) Specifies the component configurations of the cluster.
  The [object](#component_configurations) structure is documented below.
  Changing this will create a new MapReduce cluster resource.

* `tags` - (Optional, Map) Specifies the key/value pairs to associate with the cluster.

The `nodes` block supports:

* `group_name` - (Optional, String, ForceNew) Specifies the name of nodes for the node group.

  -> This parameter is only valid and mandatory for `custom_nodes`.

* `flavor` - (Required, String, ForceNew) Specifies the instance specifications for each nodes in node group.
  Changing this will create a new MapReduce cluster resource.

* `node_number` - (Required, Int) Specifies the number of nodes for the node group.

  -> Only the core node group and task node group are allowed to be updated. The number of nodes after scaling
  cannot be less than the number of nodes originally created.

* `root_volume_type` - (Required, String, ForceNew) Specifies the system disk flavor of the nodes. Changing this will
  create a new MapReduce cluster resource.

* `root_volume_size` - (Required, Int, ForceNew) Specifies the system disk size of the nodes. Changing this will create
  a new MapReduce cluster resource.

* `data_volume_count` - (Required, Int, ForceNew) Specifies the data disk number of the nodes. The number configuration
  of each node are as follows:
  + **master_nodes**: 1.
  + **analysis_core_nodes**: minimum is one and the maximum is subject to the configuration of the corresponding flavor.
  + **streaming_core_nodes**: minimum is one and the maximum is subject to the configuration of the corresponding flavor.
  + **analysis_task_nodes**: minimum is one and the maximum is subject to the configuration of the corresponding flavor.
  + **streaming_task_nodes**: minimum is one and the maximum is subject to the configuration of the corresponding flavor.

  Changing this will create a new MapReduce cluster resource.
  
* `data_volume_type` - (Optional, String, ForceNew) Specifies the data disk flavor of the nodes.
  Required if `data_volume_count` is greater than zero. Changing this will create a new MapReduce cluster resource.
   The following disk types are supported:
  + **SATA**: common I/O disk.
  + **SAS**: high I/O disk.
  + **SSD**: ultra-high I/O disk.

* `data_volume_size` - (Optional, Int, ForceNew) Specifies the data disk size of the nodes,in GB. The value range is 10
  to 32768. Required if `data_volume_count` is greater than zero. Changing this will create a new MapReduce
  cluster resource.

* `assigned_roles` - (Optional, List, ForceNew) Specifies the roles deployed in a node group.This argument is mandatory
 when the cluster type is **CUSTOM**. Each character string represents a role expression.

  **Role expression definition:**

   + If the role is deployed on all nodes in the node group, set this parameter to role_name, for example: `DataNode`.
   + If the role is deployed on a specified subscript node in the node group: role_name:index1,index2..., indexN,
 for example: `DataNode:1,2`. The subscript starts from 1.
   + Some roles support multi-instance deployment (that is, multiple instances of the same role are deployed on a node):
  role_name[instance_count], for example: `EsNode[9]`.
  
  [For details about components](https://support.huaweicloud.com/intl/en-us/productdesc-mrs/mrs_08_0005.html)

  [Mapping between roles and components](https://support.huaweicloud.com/intl/en-us/api-mrs/mrs_02_0106.html)

  -> `DBService` is a basic component of a cluster. Components such as Hive, Hue, Oozie, Loader, and Redis, and Loader
   store their metadata in DBService, and provide the metadata backup and restoration functions by using DBService.

<a name="component_configurations"></a>
The `component_configs` block supports:

* `name` - (Required, String, ForceNew) Specifies the component name of the cluster which has installed.
  Changing this will create a new MapReduce cluster resource.

* `configs` - (Required, List, ForceNew) Specifies the configuration of component installed.
  The [object](#component_configuration) structure is documented below.

<a name="component_configuration"></a>
The `configs` block supports:

* `key` - (Required, String, ForceNew) Specifies the configuration item key of component installed.
  Changing this will create a new MapReduce cluster resource.

* `value` - (Required, String, ForceNew) Specifies the configuration item value of component installed.
  Changing this will create a new MapReduce cluster resource.

* `config_file_name` - (Required, String, ForceNew) Specifies the configuration file name of component installed.
  Changing this will create a new MapReduce cluster resource.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The cluster ID in UUID format.
* `total_node_number` - The total number of nodes deployed in the cluster.
* `master_node_ip` - The IP address of the master node.
* `private_ip` - The preferred private IP address of the master node.
* `status` - The cluster state, which include: running, frozen, abnormal and failed.
* `create_time` - The cluster creation time, in RFC-3339 format.
* `update_time` - The cluster update time, in RFC-3339 format.
* `charging_start_time` - The charging start time which is the start time of billing, in RFC-3339 format.
* `node` - all the nodes attributes: master_nodes/analysis_core_nodes/streaming_core_nodes/analysis_task_nodes
  /streaming_task_nodes.
  + `host_ips` - The host list of this nodes group in the cluster.

## Timeouts

This resource provides the following timeouts configuration options:

* `create` - Default is 60 minute.
* `update` - Default is 180 minute.
* `delete` - Default is 40 minute.

## Import

Clusters can be imported by their `id`. For example,

```bash
terraform import huaweicloud_mapreduce_cluster.test b11b407c-e604-4e8d-8bc4-92398320b847
```

Note that the imported state may not be identical to your resource definition, due to some attributes missing from the
API response, security or some other reason. The missing attributes include:
`manager_admin_pass`, `node_admin_pass`,`template_id`, `assigned_roles` and `component_configs`.
It is generally recommended running `terraform plan` after importing a cluster.
You can then decide if changes should be applied to the cluster, or the resource definition
should be updated to align with the cluster. Also you can ignore changes as below.

```hcl
resource "huaweicloud_mapreduce_cluster" "test" {
    ...

  lifecycle {
    ignore_changes = [
      manager_admin_pass, node_admin_pass,
    ]
  }
}
```

package cce

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/common"
	"github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/config"
	"github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/sdk/huaweicloud"
	"github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/sdk/huaweicloud/openstack/cce/v3/nodepools"
	"github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/sdk/huaweicloud/openstack/cce/v3/nodes"
	"github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/sdk/huaweicloud/openstack/common/tags"
	"github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/utils"
	"github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/utils/fmtp"
	"github.com/huaweicloud/terraform-provider-hcs/huaweicloudstack/utils/logp"
)

func ResourceCCENodePool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCCENodePoolCreate,
		ReadContext:   resourceCCENodePoolRead,
		UpdateContext: resourceCCENodePoolUpdate,
		DeleteContext: resourceCCENodePoolDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCCENodePoolV3Import,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"initial_node_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"flavor_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"labels": { // (k8s_tags)
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"root_volume": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"size": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"volumetype": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"hw_passthrough": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: "schema: Internal",
						},
						"extend_param": {
							Type:       schema.TypeString,
							Optional:   true,
							ForceNew:   true,
							Deprecated: "use extend_params instead",
						},
						"extend_params": {
							Type:     schema.TypeMap,
							Optional: true,
							ForceNew: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"kms_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
					}},
			},
			"data_volumes": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"size": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"volumetype": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"hw_passthrough": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: "schema: Internal",
						},
						"extend_param": {
							Type:       schema.TypeString,
							Optional:   true,
							ForceNew:   true,
							Deprecated: "use extend_params instead",
						},
						"extend_params": {
							Type:     schema.TypeMap,
							Optional: true,
							ForceNew: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"kms_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
					}},
			},
			"availability_zone": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "random",
			},
			"os": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"key_pair": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ExactlyOneOf: []string{"password", "key_pair"},
			},
			"password": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Sensitive:    true,
				ExactlyOneOf: []string{"password", "key_pair"},
			},
			"storage": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"selectors": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"type": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Default:  "evs",
									},
									"match_label_size": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"match_label_volume_type": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"match_label_metadata_encrypted": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"match_label_metadata_cmkid": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"match_label_count": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"groups": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"cce_managed": {
										Type:     schema.TypeBool,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"selector_names": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"virtual_spaces": {
										Type:     schema.TypeList,
										Required: true,
										ForceNew: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"size": {
													Type:     schema.TypeString,
													Required: true,
													ForceNew: true,
												},
												"lvm_lv_type": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"lvm_path": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"runtime_lv_type": {
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										}},
								},
							},
						},
					}},
			},
			"taints": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"effect": {
							Type:     schema.TypeString,
							Required: true,
						},
					}},
			},
			"tags": common.TagsSchema(),
			// charge info: charging_mode, period_unit, period, auto_renew
			"charging_mode": common.SchemaChargingMode(nil),
			"period_unit":   common.SchemaPeriodUnit(nil),
			"period":        common.SchemaPeriod(nil),
			"auto_renew":    common.SchemaAutoRenew(nil),

			"billing_mode": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_pods": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"preinstall": {
				Type:      schema.TypeString,
				Optional:  true,
				ForceNew:  true,
				StateFunc: utils.DecodeHashAndHexEncode,
			},
			"postinstall": {
				Type:      schema.TypeString,
				Optional:  true,
				ForceNew:  true,
				StateFunc: utils.DecodeHashAndHexEncode,
			},
			"runtime": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					"docker", "containerd",
				}, false),
			},
			"extend_param": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"scall_enable": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"min_node_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_node_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scale_down_cooldown_time": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"security_groups": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pod_security_groups": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ecs_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"current_node_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceCCENodePoolTags(d *schema.ResourceData) []tags.ResourceTag {
	tagRaw := d.Get("tags").(map[string]interface{})
	return utils.ExpandResourceTags(tagRaw)
}

func buildPodSecurityGroups(ids []interface{}) []nodepools.PodSecurityGroupSpec {
	if len(ids) == 0 {
		return nil
	}

	groups := make([]nodepools.PodSecurityGroupSpec, len(ids))
	for i, id := range ids {
		groups[i] = nodepools.PodSecurityGroupSpec{Id: id.(string)}
	}

	return groups
}

func resourceCCENodePoolCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	nodePoolClient, err := config.CceV3Client(config.GetRegion(d))
	if err != nil {
		return fmtp.DiagErrorf("Error creating HuaweiCloud CCE Node Pool client: %s", err)
	}

	// validation
	billingMode := 0
	if d.Get("charging_mode").(string) == "prePaid" {
		billingMode = 1
		if err := common.ValidatePrePaidChargeInfo(d); err != nil {
			return diag.FromErr(err)
		}
	}

	// wait for the cce cluster to become available
	clusterid := d.Get("cluster_id").(string)
	stateCluster := &resource.StateChangeConf{
		Target:     []string{"Available"},
		Refresh:    waitForClusterAvailable(nodePoolClient, clusterid),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 5 * time.Second,
	}
	_, err = stateCluster.WaitForStateContext(ctx)
	if err != nil {
		return fmtp.DiagErrorf("Error waiting for HuaweiCloud CCE cluster to be Available: %s", err)
	}

	initialNodeCount := d.Get("initial_node_count").(int)
	createOpts := nodepools.CreateOpts{
		Kind:       "NodePool",
		ApiVersion: "v3",
		Metadata: nodepools.CreateMetaData{
			Name: d.Get("name").(string),
		},
		Spec: nodepools.CreateSpec{
			Type: d.Get("type").(string),
			NodeTemplate: nodes.Spec{
				Flavor:      d.Get("flavor_id").(string),
				Az:          d.Get("availability_zone").(string),
				Os:          d.Get("os").(string),
				RootVolume:  buildResourceNodeRootVolume(d),
				DataVolumes: buildResourceNodeDataVolume(d),
				Storage:     buildResourceNodeStorage(d),
				K8sTags:     buildResourceNodeK8sTags(d),
				BillingMode: billingMode,
				Count:       1,
				NodeNicSpec: nodes.NodeNicSpec{
					PrimaryNic: nodes.PrimaryNic{
						SubnetId: d.Get("subnet_id").(string),
					},
				},
				ExtendParam: buildResourceNodeExtendParam(d),
				Taints:      buildResourceNodeTaint(d),
				UserTags:    resourceCCENodePoolTags(d),
			},
			Autoscaling: nodepools.AutoscalingSpec{
				Enable:                d.Get("scall_enable").(bool),
				MinNodeCount:          d.Get("min_node_count").(int),
				MaxNodeCount:          d.Get("max_node_count").(int),
				ScaleDownCooldownTime: d.Get("scale_down_cooldown_time").(int),
				Priority:              d.Get("priority").(int),
			},
			InitialNodeCount:     &initialNodeCount,
			PodSecurityGroups:    buildPodSecurityGroups(d.Get("pod_security_groups").([]interface{})),
			CustomSecurityGroups: utils.ExpandToStringList(d.Get("security_groups").([]interface{})),
			NodeManagement: nodepools.NodeManagementSpec{
				ServerGroupReference: d.Get("ecs_group_id").(string),
			},
		},
	}

	if v, ok := d.GetOk("runtime"); ok {
		createOpts.Spec.NodeTemplate.RunTime = &nodes.RunTimeSpec{
			Name: v.(string),
		}
	}

	logp.Printf("[DEBUG] Create Options: %#v", createOpts)
	// Add loginSpec here so it wouldn't go in the above log entry
	var loginSpec nodes.LoginSpec
	if common.HasFilledOpt(d, "key_pair") {
		loginSpec = nodes.LoginSpec{
			SshKey: d.Get("key_pair").(string),
		}
	} else if common.HasFilledOpt(d, "password") {
		password, err := utils.TryPasswordEncrypt(d.Get("password").(string))
		if err != nil {
			return diag.FromErr(err)
		}
		loginSpec = nodes.LoginSpec{
			UserPassword: nodes.UserPassword{
				Username: "root",
				Password: password,
			},
		}
	}
	createOpts.Spec.NodeTemplate.Login = loginSpec

	s, err := nodepools.Create(nodePoolClient, clusterid, createOpts).Extract()
	if err != nil {
		return fmtp.DiagErrorf("Error creating HuaweiCloud Node Pool: %s", err)
	}

	if len(s.Metadata.Id) == 0 {
		return fmtp.DiagErrorf("Error fetching CreateNodePool id")
	}

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"Synchronizing"},
		Target:       []string{""},
		Refresh:      waitForCceNodePoolActive(nodePoolClient, clusterid, s.Metadata.Id),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        120 * time.Second,
		PollInterval: 20 * time.Second,
	}
	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return fmtp.DiagErrorf("Error creating HuaweiCloud CCE Node Pool: %s", err)
	}

	logp.Printf("[DEBUG] Create node pool: %v", s)

	d.SetId(s.Metadata.Id)
	return resourceCCENodePoolRead(ctx, d, meta)
}

func resourceCCENodePoolRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	nodePoolClient, err := config.CceV3Client(config.GetRegion(d))
	if err != nil {
		return fmtp.DiagErrorf("Error creating HuaweiCloud CCE Node Pool client: %s", err)
	}
	clusterid := d.Get("cluster_id").(string)
	s, err := nodepools.Get(nodePoolClient, clusterid, d.Id()).Extract()

	if err != nil {
		return common.CheckDeletedDiag(d, err, "Error retrieving HuaweiCloud Node Pool")
	}

	// The following parameters are not returned:
	// password, subnet_id, preinstall, posteinstall, taints, initial_node_count, pod_security_groups
	mErr := multierror.Append(nil,
		d.Set("region", config.GetRegion(d)),
		d.Set("name", s.Metadata.Name),
		d.Set("flavor_id", s.Spec.NodeTemplate.Flavor),
		d.Set("availability_zone", s.Spec.NodeTemplate.Az),
		d.Set("os", s.Spec.NodeTemplate.Os),
		d.Set("billing_mode", s.Spec.NodeTemplate.BillingMode),
		d.Set("key_pair", s.Spec.NodeTemplate.Login.SshKey),
		d.Set("scall_enable", s.Spec.Autoscaling.Enable),
		d.Set("min_node_count", s.Spec.Autoscaling.MinNodeCount),
		d.Set("max_node_count", s.Spec.Autoscaling.MaxNodeCount),
		d.Set("current_node_count", s.Status.CurrentNode),
		d.Set("scale_down_cooldown_time", s.Spec.Autoscaling.ScaleDownCooldownTime),
		d.Set("priority", s.Spec.Autoscaling.Priority),
		d.Set("type", s.Spec.Type),
		d.Set("ecs_group_id", s.Spec.NodeManagement.ServerGroupReference),
		d.Set("storage", flattenStorage(s.Spec.NodeTemplate.Storage)),
		d.Set("security_groups", s.Spec.CustomSecurityGroups),
	)

	if s.Spec.NodeTemplate.BillingMode != 0 {
		mErr = multierror.Append(mErr, d.Set("charging_mode", "prePaid"))
	}

	// set extend_param
	var extendParam = s.Spec.NodeTemplate.ExtendParam
	mErr = multierror.Append(mErr, d.Set("max_pods", extendParam["maxPods"]))
	delete(extendParam, "maxPods")

	extendParamToSet := map[string]string{}
	for k, v := range extendParam {
		switch v := v.(type) {
		case string:
			extendParamToSet[k] = v
		case int:
			extendParamToSet[k] = strconv.Itoa(v)
		case int32:
			extendParamToSet[k] = strconv.FormatInt(int64(v), 10)
		case float64:
			extendParamToSet[k] = strconv.FormatFloat(v, 'f', -1, 64)
		case bool:
			extendParamToSet[k] = strconv.FormatBool(v)
		default:
			logp.Printf("[WARN] can not set %s to extend_param, the value is %v", k, v)
		}
	}

	mErr = multierror.Append(mErr, d.Set("extend_param", extendParamToSet))

	if s.Spec.NodeTemplate.RunTime != nil {
		mErr = multierror.Append(mErr, d.Set("runtime", s.Spec.NodeTemplate.RunTime.Name))
	}

	labels := map[string]string{}
	for key, val := range s.Spec.NodeTemplate.K8sTags {
		if strings.Contains(key, "cce.cloud.com") {
			continue
		}
		labels[key] = val
	}
	mErr = multierror.Append(mErr, d.Set("labels", labels))

	var volumes []map[string]interface{}
	for _, pairObject := range s.Spec.NodeTemplate.DataVolumes {
		volume := make(map[string]interface{})
		volume["size"] = pairObject.Size
		volume["volumetype"] = pairObject.VolumeType
		volume["hw_passthrough"] = pairObject.HwPassthrough
		volume["extend_params"] = pairObject.ExtendParam
		volume["extend_param"] = ""
		if pairObject.Metadata != nil {
			volume["kms_key_id"] = pairObject.Metadata.SystemCmkid
		}
		volumes = append(volumes, volume)
	}
	mErr = multierror.Append(mErr, d.Set("data_volumes", volumes))

	rootVolume := []map[string]interface{}{
		{
			"size":           s.Spec.NodeTemplate.RootVolume.Size,
			"volumetype":     s.Spec.NodeTemplate.RootVolume.VolumeType,
			"hw_passthrough": s.Spec.NodeTemplate.RootVolume.HwPassthrough,
			"extend_params":  s.Spec.NodeTemplate.RootVolume.ExtendParam,
			"extend_param":   "",
		},
	}
	if s.Spec.NodeTemplate.RootVolume.Metadata != nil {
		rootVolume[0]["kms_key_id"] = s.Spec.NodeTemplate.RootVolume.Metadata.SystemCmkid
	}
	mErr = multierror.Append(mErr, d.Set("root_volume", rootVolume))

	tagmap := utils.TagsToMap(s.Spec.NodeTemplate.UserTags)
	mErr = multierror.Append(mErr,
		d.Set("tags", tagmap),
		d.Set("status", s.Status.Phase),
	)

	if err = mErr.ErrorOrNil(); err != nil {
		return fmtp.DiagErrorf("Error setting CCE Node Pool fields: %s", err)
	}

	return nil
}

func flattenStorage(storageRaw *nodes.StorageSpec) []map[string]interface{} {
	if storageRaw == nil {
		return nil
	}

	storageSelectorsRaw := storageRaw.StorageSelectors
	storageSelectors := make([]map[string]interface{}, len(storageSelectorsRaw))
	for i, s := range storageSelectorsRaw {
		storageSelectors[i] = map[string]interface{}{
			"name": s.Name,
			"type": s.StorageType,
		}

		if s.MatchLabels != (nodes.MatchLabelsSpec{}) {
			storageSelectors[i]["match_label_size"] = s.MatchLabels.Size
			storageSelectors[i]["match_label_volume_type"] = s.MatchLabels.VolumeType
			storageSelectors[i]["match_label_metadata_encrypted"] = s.MatchLabels.MetadataEncrypted
			storageSelectors[i]["match_label_metadata_cmkid"] = s.MatchLabels.MetadataCmkid
			storageSelectors[i]["match_label_count"] = s.MatchLabels.Count
		}
	}

	storageGroupsRaw := storageRaw.StorageGroups
	storageGroups := make([]map[string]interface{}, len(storageGroupsRaw))
	for i, v := range storageGroupsRaw {
		storageGroups[i] = map[string]interface{}{
			"name":           v.Name,
			"cce_managed":    v.CceManaged,
			"selector_names": v.SelectorNames,
		}

		virtualSpaces := make([]map[string]interface{}, len(v.VirtualSpaces))
		for k, s := range v.VirtualSpaces {
			virtualSpaces[k] = map[string]interface{}{
				"name": s.Name,
				"size": s.Size,
			}

			if s.LVMConfig != nil {
				virtualSpaces[k]["lvm_lv_type"] = s.LVMConfig.LvType
				virtualSpaces[k]["lvm_path"] = s.LVMConfig.Path
			}

			if s.RuntimeConfig != nil {
				virtualSpaces[k]["runtime_lv_type"] = s.RuntimeConfig.LvType
			}
		}

		storageGroups[i]["virtual_spaces"] = virtualSpaces
	}

	return []map[string]interface{}{
		{
			"selectors": storageSelectors,
			"groups":    storageGroups,
		},
	}
}

func resourceCCENodePoolUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	nodePoolClient, err := config.CceV3Client(config.GetRegion(d))
	if err != nil {
		return fmtp.DiagErrorf("Error creating HuaweiCloud CCE client: %s", err)
	}

	initialNodeCount := d.Get("initial_node_count").(int)
	var loginSpec nodes.LoginSpec
	if common.HasFilledOpt(d, "key_pair") {
		loginSpec = nodes.LoginSpec{SshKey: d.Get("key_pair").(string)}
	} else if common.HasFilledOpt(d, "password") {
		password, err := utils.TryPasswordEncrypt(d.Get("password").(string))
		if err != nil {
			return diag.FromErr(err)
		}
		loginSpec = nodes.LoginSpec{
			UserPassword: nodes.UserPassword{
				Username: "root",
				Password: password,
			},
		}
	}

	updateOpts := nodepools.UpdateOpts{
		Kind:       "NodePool",
		ApiVersion: "v3",
		Metadata: nodepools.UpdateMetaData{
			Name: d.Get("name").(string),
		},
		Spec: nodepools.UpdateSpec{
			InitialNodeCount: &initialNodeCount,
			Autoscaling: nodepools.AutoscalingSpec{
				Enable:                d.Get("scall_enable").(bool),
				MinNodeCount:          d.Get("min_node_count").(int),
				MaxNodeCount:          d.Get("max_node_count").(int),
				ScaleDownCooldownTime: d.Get("scale_down_cooldown_time").(int),
				Priority:              d.Get("priority").(int),
			},
			NodeTemplate: nodes.Spec{
				Flavor:      d.Get("flavor_id").(string),
				Az:          d.Get("availability_zone").(string),
				Login:       loginSpec,
				RootVolume:  buildResourceNodeRootVolume(d),
				DataVolumes: buildResourceNodeDataVolume(d),
				Count:       1,
				UserTags:    resourceCCENodePoolTags(d),
				K8sTags:     buildResourceNodeK8sTags(d),
				Taints:      buildResourceNodeTaint(d),
			},
			Type: d.Get("type").(string),
		},
	}

	clusterid := d.Get("cluster_id").(string)
	_, err = nodepools.Update(nodePoolClient, clusterid, d.Id(), updateOpts).Extract()
	if err != nil {
		return fmtp.DiagErrorf("Error updating HuaweiCloud Node Node Pool: %s", err)
	}

	stateConf := &resource.StateChangeConf{
		Pending:      []string{"Synchronizing"},
		Target:       []string{""},
		Refresh:      waitForCceNodePoolActive(nodePoolClient, clusterid, d.Id()),
		Timeout:      d.Timeout(schema.TimeoutCreate),
		Delay:        60 * time.Second,
		PollInterval: 10 * time.Second,
	}
	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return fmtp.DiagErrorf("Error updating HuaweiCloud CCE Node Pool: %s", err)
	}

	return resourceCCENodePoolRead(ctx, d, meta)
}

func resourceCCENodePoolDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	config := meta.(*config.Config)
	nodePoolClient, err := config.CceV3Client(config.GetRegion(d))
	if err != nil {
		return fmtp.DiagErrorf("Error creating HuaweiCloud CCE client: %s", err)
	}
	clusterid := d.Get("cluster_id").(string)
	err = nodepools.Delete(nodePoolClient, clusterid, d.Id()).ExtractErr()
	if err != nil {
		return fmtp.DiagErrorf("Error deleting HuaweiCloud CCE Node Pool: %s", err)
	}
	stateConf := &resource.StateChangeConf{
		Pending:      []string{"Deleting"},
		Target:       []string{"Deleted"},
		Refresh:      waitForCceNodePoolDelete(nodePoolClient, clusterid, d.Id()),
		Timeout:      d.Timeout(schema.TimeoutDelete),
		Delay:        60 * time.Second,
		PollInterval: 20 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return fmtp.DiagErrorf("Error deleting HuaweiCloud CCE Node Pool: %s", err)
	}

	d.SetId("")
	return nil
}

func waitForCceNodePoolActive(cceClient *golangsdk.ServiceClient, clusterId, nodePoolId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		n, err := nodepools.Get(cceClient, clusterId, nodePoolId).Extract()
		if err != nil {
			return nil, "", err
		}
		return n, n.Status.Phase, nil
	}
}

func waitForCceNodePoolDelete(cceClient *golangsdk.ServiceClient, clusterId, nodePoolId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		logp.Printf("[DEBUG] Attempting to delete HuaweiCloud CCE Node Pool %s.\n", nodePoolId)

		r, err := nodepools.Get(cceClient, clusterId, nodePoolId).Extract()

		if err != nil {
			if _, ok := err.(golangsdk.ErrDefault404); ok {
				logp.Printf("[DEBUG] Successfully deleted HuaweiCloud CCE Node Pool %s", nodePoolId)
				return r, "Deleted", nil
			}
			return r, "Deleting", err
		}

		logp.Printf("[DEBUG] HuaweiCloud CCE Node Pool %s still available.\n", nodePoolId)
		return r, r.Status.Phase, nil
	}
}

func resourceCCENodePoolV3Import(_ context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	parts := strings.SplitN(d.Id(), "/", 2)
	if len(parts) != 2 {
		err := fmtp.Errorf("Invalid format specified for CCE Node Pool. Format must be <cluster id>/<node pool id>")
		return nil, err
	}

	clusterID := parts[0]
	nodePoolID := parts[1]

	d.SetId(nodePoolID)
	d.Set("cluster_id", clusterID)

	return []*schema.ResourceData{d}, nil
}

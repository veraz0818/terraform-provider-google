// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/PublicDelegatedPrefix.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceComputePublicDelegatedPrefix() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputePublicDelegatedPrefixCreate,
		Read:   resourceComputePublicDelegatedPrefixRead,
		Delete: resourceComputePublicDelegatedPrefixDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputePublicDelegatedPrefixImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"ip_cidr_range": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The IP address range, in CIDR format, represented by this public delegated prefix.`,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. The name must be 1-63 characters long, and
comply with RFC1035. Specifically, the name must be 1-63 characters
long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'
which means the first character must be a lowercase letter, and all
following characters must be a dash, lowercase letter, or digit,
except the last character, which cannot be a dash.`,
			},
			"parent_prefix": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description:      `The URL of parent prefix. Either PublicAdvertisedPrefix or PublicDelegatedPrefix.`,
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `A region where the prefix will reside.`,
			},
			"allocatable_prefix_length": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `The allocatable prefix length supported by this public delegated prefix. This field is optional and cannot be set for prefixes in DELEGATION mode. It cannot be set for IPv4 prefixes either, and it always defaults to 32.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"is_live_migration": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `If true, the prefix will be live migrated.`,
			},
			"mode": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: verify.ValidateEnum([]string{"DELEGATION", "EXTERNAL_IPV6_FORWARDING_RULE_CREATION", "EXTERNAL_IPV6_SUBNETWORK_CREATION", ""}),
				Description: `Specifies the mode of this IPv6 PDP. MODE must be one of: DELEGATION,
EXTERNAL_IPV6_FORWARDING_RULE_CREATION and EXTERNAL_IPV6_SUBNETWORK_CREATION. Possible values: ["DELEGATION", "EXTERNAL_IPV6_FORWARDING_RULE_CREATION", "EXTERNAL_IPV6_SUBNETWORK_CREATION"]`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputePublicDelegatedPrefixCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputePublicDelegatedPrefixDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	isLiveMigrationProp, err := expandComputePublicDelegatedPrefixIsLiveMigration(d.Get("is_live_migration"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("is_live_migration"); !tpgresource.IsEmptyValue(reflect.ValueOf(isLiveMigrationProp)) && (ok || !reflect.DeepEqual(v, isLiveMigrationProp)) {
		obj["isLiveMigration"] = isLiveMigrationProp
	}
	nameProp, err := expandComputePublicDelegatedPrefixName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	parentPrefixProp, err := expandComputePublicDelegatedPrefixParentPrefix(d.Get("parent_prefix"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("parent_prefix"); !tpgresource.IsEmptyValue(reflect.ValueOf(parentPrefixProp)) && (ok || !reflect.DeepEqual(v, parentPrefixProp)) {
		obj["parentPrefix"] = parentPrefixProp
	}
	modeProp, err := expandComputePublicDelegatedPrefixMode(d.Get("mode"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("mode"); !tpgresource.IsEmptyValue(reflect.ValueOf(modeProp)) && (ok || !reflect.DeepEqual(v, modeProp)) {
		obj["mode"] = modeProp
	}
	allocatablePrefixLengthProp, err := expandComputePublicDelegatedPrefixAllocatablePrefixLength(d.Get("allocatable_prefix_length"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("allocatable_prefix_length"); !tpgresource.IsEmptyValue(reflect.ValueOf(allocatablePrefixLengthProp)) && (ok || !reflect.DeepEqual(v, allocatablePrefixLengthProp)) {
		obj["allocatablePrefixLength"] = allocatablePrefixLengthProp
	}
	ipCidrRangeProp, err := expandComputePublicDelegatedPrefixIpCidrRange(d.Get("ip_cidr_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_cidr_range"); !tpgresource.IsEmptyValue(reflect.ValueOf(ipCidrRangeProp)) && (ok || !reflect.DeepEqual(v, ipCidrRangeProp)) {
		obj["ipCidrRange"] = ipCidrRangeProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/publicDelegatedPrefixes")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new PublicDelegatedPrefix: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PublicDelegatedPrefix: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating PublicDelegatedPrefix: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/publicDelegatedPrefixes/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating PublicDelegatedPrefix", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create PublicDelegatedPrefix: %s", err)
	}

	log.Printf("[DEBUG] Finished creating PublicDelegatedPrefix %q: %#v", d.Id(), res)

	return resourceComputePublicDelegatedPrefixRead(d, meta)
}

func resourceComputePublicDelegatedPrefixRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/publicDelegatedPrefixes/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PublicDelegatedPrefix: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ComputePublicDelegatedPrefix %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}

	if err := d.Set("description", flattenComputePublicDelegatedPrefixDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}
	if err := d.Set("is_live_migration", flattenComputePublicDelegatedPrefixIsLiveMigration(res["isLiveMigration"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}
	if err := d.Set("name", flattenComputePublicDelegatedPrefixName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}
	if err := d.Set("parent_prefix", flattenComputePublicDelegatedPrefixParentPrefix(res["parentPrefix"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}
	if err := d.Set("mode", flattenComputePublicDelegatedPrefixMode(res["mode"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}
	if err := d.Set("allocatable_prefix_length", flattenComputePublicDelegatedPrefixAllocatablePrefixLength(res["allocatablePrefixLength"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}
	if err := d.Set("ip_cidr_range", flattenComputePublicDelegatedPrefixIpCidrRange(res["ipCidrRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}
	if err := d.Set("self_link", tpgresource.ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading PublicDelegatedPrefix: %s", err)
	}

	return nil
}

func resourceComputePublicDelegatedPrefixDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for PublicDelegatedPrefix: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/publicDelegatedPrefixes/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting PublicDelegatedPrefix %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "PublicDelegatedPrefix")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting PublicDelegatedPrefix", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting PublicDelegatedPrefix %q: %#v", d.Id(), res)
	return nil
}

func resourceComputePublicDelegatedPrefixImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/publicDelegatedPrefixes/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<region>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/regions/{{region}}/publicDelegatedPrefixes/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputePublicDelegatedPrefixDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePublicDelegatedPrefixIsLiveMigration(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePublicDelegatedPrefixName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePublicDelegatedPrefixParentPrefix(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePublicDelegatedPrefixMode(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputePublicDelegatedPrefixAllocatablePrefixLength(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputePublicDelegatedPrefixIpCidrRange(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandComputePublicDelegatedPrefixDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixIsLiveMigration(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixParentPrefix(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixMode(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixAllocatablePrefixLength(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandComputePublicDelegatedPrefixIpCidrRange(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

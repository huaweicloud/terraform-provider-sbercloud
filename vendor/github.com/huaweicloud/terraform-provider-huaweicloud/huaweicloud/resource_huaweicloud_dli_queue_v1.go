// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file at
//     https://www.github.com/huaweicloud/magic-modules
//
// ----------------------------------------------------------------------------

package huaweicloud

import (
	"fmt"
	"log"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
)

func ResourceDliQueueV1() *schema.Resource {
	return &schema.Resource{
		Create: resourceDliQueueV1Create,
		Read:   resourceDliQueueV1Read,
		Delete: resourceDliQueueV1Delete,

		Schema: map[string]*schema.Schema{
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cu_count": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"description": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},

			"management_subnet_cidr": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"subnet_cidr": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"vpc_cidr": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceDliQueueV1UserInputParams(d *schema.ResourceData) map[string]interface{} {
	return map[string]interface{}{
		"terraform_resource_data": d,
		"cu_count":                d.Get("cu_count"),
		"description":             d.Get("description"),
		"management_subnet_cidr":  d.Get("management_subnet_cidr"),
		"name":                    d.Get("name"),
		"subnet_cidr":             d.Get("subnet_cidr"),
		"vpc_cidr":                d.Get("vpc_cidr"),
	}
}

func resourceDliQueueV1Create(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*config.Config)
	client, err := config.DliV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	opts := resourceDliQueueV1UserInputParams(d)

	params, err := buildDliQueueV1CreateParameters(opts, nil)
	if err != nil {
		return fmt.Errorf("Error building the request body of api(create), err=%s", err)
	}
	r, err := sendDliQueueV1CreateRequest(d, params, client)
	if err != nil {
		return fmt.Errorf("Error creating DliQueueV1, err=%s", err)
	}

	id, err := navigateValue(r, []string{"queue_name"}, nil)
	if err != nil {
		return fmt.Errorf("Error constructing id, err=%s", err)
	}
	d.SetId(convertToStr(id))

	return resourceDliQueueV1Read(d, meta)
}

func resourceDliQueueV1Read(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*config.Config)
	client, err := config.DliV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	res := make(map[string]interface{})

	v, err := fetchDliQueueV1ByList(d, client)
	if err != nil {
		return err
	}
	res["list"] = fillDliQueueV1ListRespBody(v)

	states, err := flattenDliQueueV1Options(res)
	if err != nil {
		return err
	}

	return setDliQueueV1States(d, states)
}

func resourceDliQueueV1Delete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*config.Config)
	client, err := config.DliV1Client(GetRegion(d, config))
	if err != nil {
		return fmt.Errorf("Error creating sdk client, err=%s", err)
	}

	url, err := replaceVars(d, "queues/{id}", nil)
	if err != nil {
		return err
	}
	url = client.ServiceURL(url)

	log.Printf("[DEBUG] Deleting Queue %q", d.Id())
	r := golangsdk.Result{}
	_, r.Err = client.Delete(url, &golangsdk.RequestOpts{
		OkCodes:      successHTTPCodes,
		JSONBody:     nil,
		JSONResponse: &r.Body,
		MoreHeaders:  map[string]string{"Content-Type": "application/json"},
	})
	if r.Err != nil {
		return fmt.Errorf("Error deleting Queue %q, err=%s", d.Id(), r.Err)
	}

	return nil
}

func buildDliQueueV1CreateParameters(opts map[string]interface{}, arrayIndex map[string]int) (interface{}, error) {
	params := make(map[string]interface{})

	params["charging_mode"] = 1

	v, err := navigateValue(opts, []string{"management_subnet_cidr"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		params["cidr_in_mgntsubnet"] = v
	}

	v, err = navigateValue(opts, []string{"subnet_cidr"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		params["cidr_in_subnet"] = v
	}

	v, err = navigateValue(opts, []string{"vpc_cidr"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		params["cidr_in_vpc"] = v
	}

	v, err = navigateValue(opts, []string{"cu_count"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		params["cu_count"] = v
	}

	v, err = navigateValue(opts, []string{"description"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		params["description"] = v
	}

	v, err = navigateValue(opts, []string{"name"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	if e, err := isEmptyValue(reflect.ValueOf(v)); err != nil {
		return nil, err
	} else if !e {
		params["queue_name"] = v
	}

	return params, nil
}

func sendDliQueueV1CreateRequest(d *schema.ResourceData, params interface{},
	client *golangsdk.ServiceClient) (interface{}, error) {
	url := client.ServiceURL("queues")

	r := golangsdk.Result{}
	_, r.Err = client.Post(url, params, &r.Body, &golangsdk.RequestOpts{
		OkCodes: successHTTPCodes,
	})
	if r.Err != nil {
		return nil, fmt.Errorf("Error running api(create), err=%s", r.Err)
	}
	return r.Body, nil
}

func fetchDliQueueV1ByList(d *schema.ResourceData, client *golangsdk.ServiceClient) (interface{}, error) {
	link := client.ServiceURL("queues")

	return findDliQueueV1ByList(client, link, d.Id())
}

func findDliQueueV1ByList(client *golangsdk.ServiceClient, link, resourceID string) (interface{}, error) {
	r, err := sendDliQueueV1ListRequest(client, link)
	if err != nil {
		return nil, err
	}
	for _, item := range r.([]interface{}) {
		val, ok := item.(map[string]interface{})["queue_name"]
		if ok && resourceID == convertToStr(val) {
			return item, nil
		}
	}

	return nil, fmt.Errorf("Error finding the resource by list api")
}

func sendDliQueueV1ListRequest(client *golangsdk.ServiceClient, url string) (interface{}, error) {
	r := golangsdk.Result{}
	_, r.Err = client.Get(url, &r.Body, &golangsdk.RequestOpts{
		MoreHeaders: map[string]string{"Content-Type": "application/json"}})
	if r.Err != nil {
		return nil, fmt.Errorf("Error running api(list) for resource(DliQueueV1), err=%s", r.Err)
	}

	v, err := navigateValue(r.Body, []string{"queues"}, nil)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func fillDliQueueV1ListRespBody(body interface{}) interface{} {
	result := make(map[string]interface{})
	val, ok := body.(map[string]interface{})
	if !ok {
		val = make(map[string]interface{})
	}

	if v, ok := val["create_time"]; ok {
		result["create_time"] = v
	} else {
		result["create_time"] = nil
	}

	if v, ok := val["cu_count"]; ok {
		result["cu_count"] = v
	} else {
		result["cu_count"] = nil
	}

	if v, ok := val["description"]; ok {
		result["description"] = v
	} else {
		result["description"] = nil
	}

	if v, ok := val["queue_name"]; ok {
		result["queue_name"] = v
	} else {
		result["queue_name"] = nil
	}

	return result
}

func flattenDliQueueV1Options(response map[string]interface{}) (map[string]interface{}, error) {
	opts := make(map[string]interface{})

	v, err := flattenDliQueueV1CreateTime(response, nil)
	if err != nil {
		return nil, fmt.Errorf("Error flattening Queue:create_time, err: %s", err)
	}
	opts["create_time"] = v

	v, err = navigateValue(response, []string{"list", "cu_count"}, nil)
	if err != nil {
		return nil, fmt.Errorf("Error flattening Queue:cu_count, err: %s", err)
	}
	opts["cu_count"] = v

	v, err = navigateValue(response, []string{"list", "description"}, nil)
	if err != nil {
		return nil, fmt.Errorf("Error flattening Queue:description, err: %s", err)
	}
	opts["description"] = v

	v, err = navigateValue(response, []string{"list", "queue_name"}, nil)
	if err != nil {
		return nil, fmt.Errorf("Error flattening Queue:name, err: %s", err)
	}
	opts["name"] = v

	return opts, nil
}

func flattenDliQueueV1CreateTime(d interface{}, arrayIndex map[string]int) (interface{}, error) {
	v, err := navigateValue(d, []string{"list", "create_time"}, arrayIndex)
	if err != nil {
		return nil, err
	}
	val, err := convertToInt(v)
	if err != nil {
		return nil, err
	}
	return convertSeconds2Str(val / 1000), nil
}

func setDliQueueV1States(d *schema.ResourceData, opts map[string]interface{}) error {
	for k, v := range opts {
		//lintignore:R001
		if err := d.Set(k, v); err != nil {
			return fmt.Errorf("Error setting DliQueueV1:%s, err: %s", k, err)
		}
	}
	return nil
}

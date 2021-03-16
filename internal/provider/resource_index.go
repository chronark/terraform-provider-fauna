package provider

import (
	"context"
	f "github.com/fauna/faunadb-go/v3/faunadb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIndex() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: `
It is possible to rename an index by updating its name field. Renaming an index changes its Reference, but preserves inbound References to the index. Index data is not rebuilt.
An index’s terms and values fields may not be changed. If you require such a change, the existing index must be deleted and a new one created using the new definitions for terms and/or values.
If you update the unique field, existing duplicate items are not removed from the index.`,
		CreateContext: resourceIndexCreate,
		ReadContext:   resourceIndexRead,
		UpdateContext: resourceIndexUpdate,
		DeleteContext: resourceIndexDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "The logical name of the index. Cannot be `events`, `sets`, `self`, `documents`, or `_`.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"sources": {
				Description: "An array of one or more collection names",
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"terms": {
				Description: "An array of Term objects describing the fields that should be searchable. Indexed terms can be used to search for field values, via the `Match` function. The default is an empty Array.",
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"values": {
				Description: "An array of Value objects describing the fields that should be reported in search results. The default is an empty Array. When no `values` fields are defined, search results report the indexed document’s Reference.",
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Description: "The path fo the field within a document to be indexed",
							Required:    true,
							Type:        schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"reverse": {
							Description: "Whether this field’s value should sort reversed.",
							Optional:    true,
							Type:        schema.TypeBool,
							Default:     false,
						},
					},
				},
			},
			"unique": {
				Description: "If `true`, maintains a unique constraint on combined `terms` and `values`.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"serialized": {
				Description: "If `true`, writes to this index are serialized with concurrent reads and writes.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},

			"ts": {
				Description: "A timestamp when this index was created",
				Type:        schema.TypeInt,
				Computed:    true,
			},
			"active": {
				Description: "When an index is added, it is immediately available for reads, but returns incomplete results until it is built. Fauna builds the index asynchronously by scanning over relevant documents. Upon completion, the index’s `active` field is set to `true`.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"partitions": {
				Description: "The number of sub-partitions used by this index. This value can be 8 or 1:\n`1` when unique is true.\n`8` when the index has no terms.\n`1` in all other case.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}
}

func buildSources(sourceNames []interface{}) f.Arr {
	var arr f.Arr
	for i := 0; i < len(sourceNames); i++ {
		arr = append(arr, f.Collection(sourceNames[i]))
	}
	return arr
}

func buildTerms(terms []interface{}) f.Arr {
	var arr f.Arr

	for i := 0; i < len(terms); i++ {
		arr = append(arr, f.Obj{
			"field": terms[i].(map[string]interface{})["field"].([]string),
		})
	}
	return arr
}

func resourceIndexCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*f.FaunaClient)

	name := d.Get("name").(string)
	unique := d.Get("unique").(bool)
	serialized := d.Get("serialized").(bool)
	sourceNames := d.Get("sources").([]interface{})
	terms := d.Get("terms").([]interface{})

	_, err := client.Query(f.CreateIndex(f.Obj{
		"unique":     unique,
		"serialized": serialized,
		"name":       name,
		"source":     buildSources(sourceNames),
		"terms":      buildTerms(terms),
		// "values": f.Arr{f.Obj{
		// 	"field": f.Arr{"data", "elements"},
		// }},
	}))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(name)

	return resourceIndexRead(ctx, d, meta)

}

func resourceIndexRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*f.FaunaClient)
	id := d.Id()

	res, err := client.Query(f.Get(f.Index(id)))
	if err != nil {
		return diag.FromErr(err)
	}

	var (
		name       string
		ts         int64
		unique     bool
		serialized bool
		active     bool
	)
	// name
	err = res.At(f.ObjKey("name")).Get(&name)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("name", name)
	if err != nil {
		return diag.FromErr(err)
	}

	// ts
	err = res.At(f.ObjKey("ts")).Get(&ts)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("ts", ts)
	if err != nil {
		return diag.FromErr(err)
	}

	// unique
	err = res.At(f.ObjKey("unique")).Get(&unique)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("unique", unique)
	if err != nil {
		return diag.FromErr(err)
	}
	// serialized
	err = res.At(f.ObjKey("serialized")).Get(&serialized)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("serialized", serialized)
	if err != nil {
		return diag.FromErr(err)
	}
	// active
	err = res.At(f.ObjKey("active")).Get(&active)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("active", active)
	if err != nil {
		return diag.FromErr(err)
	}
	//terms
	type Term struct {
		Field []string `fauna:"field"`
	}
	var terms []Term
	err = res.At(f.ObjKey("terms")).Get(&terms)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("terms", terms)
	if err != nil {
		return diag.FromErr(err)
	}
	//sources
	var sources []string
	err = res.At(f.ObjKey("sources")).Get(&sources)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("sources", sources)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceIndexUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*f.FaunaClient)
	id := d.Id()

	if d.HasChange("name") {
		newName := d.Get("name").(string)
		_, err := client.Query(f.Update(f.Index(id), f.Obj{"name": newName}))
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(newName)
	}
	if d.HasChange("ttl_days") {
		_, err := client.Query(f.Update(f.Index(id), f.Obj{"ttl_days": d.Get("ttl_days").(int)}))
		if err != nil {
			return diag.FromErr(err)
		}
	}
	if d.HasChange("history_days") {
		_, err := client.Query(f.Update(f.Index(id), f.Obj{"history_days": d.Get("history_days").(int)}))
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceIndexRead(ctx, d, meta)
}

func resourceIndexDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*f.FaunaClient)

	id := d.Id()

	_, err := client.Query(f.Delete(f.Index(id)))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

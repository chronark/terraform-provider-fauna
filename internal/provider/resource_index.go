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
		Description: "Manipulate indices. Please keep in mind that indices can not be modified. A change will always result in a destroy and create operation.",

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
			"source": {
				Description: "An array of one or more collection names",
				Type:        schema.TypeSet,
				Required:    true,
				Elem: &schema.Schema{
					Description: "Collection name",
					Type:        schema.TypeString,
					Required:    true,
				},
			},
			"terms": {
				Description: "An array of Term objects describing the fields that should be searchable. Indexed terms can be used to search for field values, via the `Match` function. The default is an empty Array.",
				Optional:    true,
				Type:        schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": &schema.Schema{
							Type: schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								Required: true,
							},
						},
						"binding": &schema.Schema{
							Type: schema.TypeString,
							Optional:true,
						}
					}

				},

			},
			"values": {
				Description: "An array of Value objects describing the fields that should be reported in search results. The default is an empty Array. When no `values` fields are defined, search results report the indexed document’s Reference.",
				Optional:    true,
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
			"permissions": {
				Description: "Indicates who is allowed to read the index. The default is everyone can read the index.",
			},
			"data": {
				Description: " This is user-defined metadata for the index. It is provided for the developer to store information at the index level. The default is an empty object having no data.",
			},
		},
	}
}

type Collection struct {
	name        string
	ts          int64
	historyDays int64
	ttlDays     int64
}

func resourceCollectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*f.FaunaClient)

	name := d.Get("name").(string)
	res, err := client.Query(f.CreateCollection(f.Obj{
		"name":         name,
		"history_days": d.Get("history_days").(int),
		"ttl_days":     d.Get("ttl_days").(int),
	}))
	if err != nil {
		return diag.FromErr(err)
	}

	var collection Collection
	err = res.Get(&collection)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("name", collection.name)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("ts", collection.ts)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("history_days", collection.historyDays)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("ttl_days", collection.ttlDays)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(name)

	return resourceCollectionRead(ctx, d, meta)

}

func resourceCollectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*f.FaunaClient)
	id := d.Id()

	res, err := client.Query(f.Get(f.Collection(id)))
	if err != nil {
		return diag.FromErr(err)
	}

	var collection Collection
	err = res.Get(&collection)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("ts", collection.ts)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("history_days", collection.historyDays)
	if err != nil {
		return diag.FromErr(err)
	}
	err = d.Set("ttl_days", collection.ttlDays)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceCollectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*f.FaunaClient)
	id := d.Id()

	if d.HasChange("name") {
		newName := d.Get("name").(string)
		_, err := client.Query(f.Update(f.Collection(id), f.Obj{"name": newName}))
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(newName)
	}
	if d.HasChange("ttl_days") {
		_, err := client.Query(f.Update(f.Collection(id), f.Obj{"ttl_days": d.Get("ttl_days").(int)}))
		if err != nil {
			return diag.FromErr(err)
		}
	}
	if d.HasChange("history_days") {
		_, err := client.Query(f.Update(f.Collection(id), f.Obj{"history_days": d.Get("history_days").(int)}))
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return resourceCollectionRead(ctx, d, meta)
}

func resourceCollectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*f.FaunaClient)

	id := d.Id()

	_, err := client.Query(f.Delete(f.Collection(id)))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

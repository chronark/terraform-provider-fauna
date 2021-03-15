package provider

import (
	"context"
	f "github.com/fauna/faunadb-go/v3/faunadb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCollection() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Manipulate collections, watch out when removing `ttl_days` or `history_days` from a terraform config. Currently they will not reset to the fauna defaults. You need to set them manually.",

		CreateContext: resourceCollectionCreate,
		ReadContext:   resourceCollectionRead,
		UpdateContext: resourceCollectionUpdate,
		DeleteContext: resourceCollectionDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ts": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"history_days": {
				Description: "Will not reset to the fauna defaults when you delete it. Please manually set it to `30`",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"ttl_days": {
				Description: "Will not reset to the fauna defaults when you delete it. Please manually set it to `0`",
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
			},
			"id": {
				Description: "The id of this resource.",
				Type:        schema.TypeString,
				Computed:    true,
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

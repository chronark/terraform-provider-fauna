package provider

import (
	"context"
	f "github.com/fauna/faunadb-go/v3/faunadb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func resourceCollection() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Manipulate collections",

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
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ttl_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
	var diags diag.Diagnostics
	client := meta.(*f.FaunaClient)
	name := d.Get("name").(string)

	res, err := client.Query(f.CreateCollection(f.Obj{"name": name}))
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
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func resourceCollectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*f.FaunaClient)
	name := d.Get("name").(string)

	res, err := client.Query(f.Get(f.Collection(name)))
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
	d.SetId(collection.name)

	return diags
}

func resourceCollectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*f.FaunaClient)
	oldName := d.Get("id").(string)
	newName := d.Get("name").(string)
	_, err := client.Query(f.Update(f.Collection(oldName), f.Obj{"name": newName}))
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceCollectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*f.FaunaClient)

	name := d.Get("name").(string)
	_, err := client.Query(f.Delete(f.Collection(name)))
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

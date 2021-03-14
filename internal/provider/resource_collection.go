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

func resourceCollectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*f.FaunaClient)
	name := d.Get("name").(string)

	res, err := client.Query(f.CreateCollection(f.Obj{"name": name}))
	if err != nil {
		return diag.FromErr(err)
	}

	type Collection struct {
		name        string
		ts          int64
		historyDays int64
		ttlDays     int64
	}

	var collection Collection
	res.Get(&collection)
	d.Set("ts", collection.ts)
	d.Set("history_days", collection.historyDays)
	d.Set("ttl_days", collection.ttlDays)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func resourceCollectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceCollectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceCollectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

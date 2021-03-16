package provider

import (
	"context"

	f "github.com/fauna/faunadb-go/v3/faunadb"

	"github.com/hashicorp/go-uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceKeys() *schema.Resource {
	return &schema.Resource{
		Description: "A Set reference for the available authentication keys in the current database.",
		ReadContext: dataSourceKeysRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Internal id",
				Computed:    true,
				Type:        schema.TypeString,
			},
			"database": {
				Description: "A reference to a child database. If not specified, the current database is used.",
				Optional:    true,
				Type:        schema.TypeString,
			},
			"keys": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Description: "The internal reference.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"ts": {
							Description: "Fauna timestamp",
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"database": {
							Description: "Name of the database this key belongs to. Empty if using the default database.",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"role": {
							Description: "Name of the role this key has",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"hashed_secret": {
							Description: "The hashed secret of this key",
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": {
							Description: "The name you gave this key",
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

// flattenKey extracts the id and name and returns a flat map ready for terraform.
func flattenKey(raw map[string]interface{}) (map[string]interface{}, error) {
	flatKey := make(map[string]interface{})

	// name is optional
	data := raw["data"]
	if data != nil {
		name, err := raw["data"].(f.ObjectV).At(f.ObjKey("name")).GetValue()
		if err != nil {
			return nil, err
		}
		flatKey["name"] = name
	}

	flatKey["id"] = raw["ref"].(f.RefV).ID
	flatKey["ts"] = raw["ts"]
	flatKey["role"] = raw["role"]
	flatKey["hashed_secret"] = raw["hashed_secret"]

	return flatKey, nil
}

func dataSourceKeysRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var (
		diags    diag.Diagnostics
		keysExpr f.Expr
	)

	database := d.Get("database").(string)

	if database != "" {
		database = d.Get("database").(string)
		keysExpr = f.ScopedKeys(f.Database(database))
	} else {
		keysExpr = f.Keys()
	}

	client := meta.(*f.FaunaClient)
	res, err := client.Query(f.Map(f.Paginate(keysExpr, f.Size(10000)), f.Lambda("key", f.Get(f.Var("key")))))
	if err != nil {
		return diag.FromErr(err)
	}

	keys := make([]map[string]interface{}, 0)
	err = res.At(f.ObjKey("data")).Get(&keys)
	if err != nil {
		return diag.FromErr(err)
	}

	for i := 0; i < len(keys); i++ {
		flatKey, err := flattenKey(keys[i])
		if err != nil {
			return diag.FromErr(err)
		}
		if database != "" {
			flatKey["database"] = database
		}

		keys[i] = flatKey

	}

	err = d.Set("keys", keys)
	if err != nil {
		return diag.FromErr(err)
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(id)
	return diags
}

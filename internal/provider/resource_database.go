package provider

import (
	"context"
	f "github.com/fauna/faunadb-go/v3/faunadb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDatabase() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Create, Read, Update or Delete databases",

		CreateContext: resourceDatabaseCreate,
		ReadContext:   resourceDatabaseRead,
		UpdateContext: resourceDatabaseUpdate,
		DeleteContext: resourceDatabaseDelete,

		Schema: map[string]*schema.Schema{
			"ref": {
				Description: "A Reference to the database that was created.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "The name of a database. Databases cannot be named any of the following reserved words: `events`, `set`, `self`, `documents`, or `_`.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"id": {
				Description: "The id of this resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"ts": {
				Description: "The timestamp, with microsecond resolution, associated with the creation of the database.",
				Type:        schema.TypeInt,
				Computed:    true,
			},
		},
	}
}

func copyAttributes(d *schema.ResourceData, source map[string]interface{}, mapping map[string]string) error {
	for schemaName, attribute := range mapping {
		err := d.Set(schemaName, source[attribute])
		if err != nil {
			return err
		}
	}
	return nil

}

func resourceDatabaseCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*f.FaunaClient)

	name := d.Get("name").(string)

	res, err := client.Query(f.CreateDatabase(f.Obj{
		"name": name,
	}))
	if err != nil {
		return diag.FromErr(err)
	}

	database := make(map[string]interface{})
	err = res.Get(&database)
	if err != nil {
		return diag.FromErr(err)
	}

	mapping := map[string]string{
		"name": "name",
		"ts":   "ts",
	}
	err = copyAttributes(d, database, mapping)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(database["global_id"].(f.StringV).String())

	return resourceDatabaseRead(ctx, d, meta)

}

func resourceDatabaseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	client := meta.(*f.FaunaClient)

	name := d.Get("name").(string)
	res, err := client.Query(f.Get(f.Database(name)))
	if err != nil {
		return diag.FromErr(err)
	}

	database := make(map[string]interface{})
	err = res.Get(&database)
	if err != nil {
		return diag.FromErr(err)
	}

	mapping := map[string]string{
		"name": "name",
		"ts":   "ts",
		"ref":  "name",
	}
	err = copyAttributes(d, database, mapping)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(database["global_id"].(f.StringV).String())

	return diags
}

func resourceDatabaseUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*f.FaunaClient)
	ref := d.Get("ref").(string)

	if d.HasChange("name") {
		newName := d.Get("name").(string)
		res, err := client.Query(f.Update(f.Database(ref), f.Obj{"name": newName}))
		if err != nil {
			return diag.FromErr(err)
		}

		database := make(map[string]interface{})
		err = res.Get(&database)
		if err != nil {
			return diag.FromErr(err)
		}
		mapping := map[string]string{
			"name": "name",
			"ts":   "ts",
			"ref":  "name",
		}
		err = copyAttributes(d, database, mapping)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(database["global_id"].(f.StringV).String())

	}
	return resourceDatabaseRead(ctx, d, meta)
}

func resourceDatabaseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := meta.(*f.FaunaClient)

	ref := d.Get("ref").(string)

	_, err := client.Query(f.Delete(f.Database(ref)))
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}

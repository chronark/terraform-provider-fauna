package provider

import (
	"context"
	f "github.com/fauna/faunadb-go/v3/faunadb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"time"
)

func dataSourceDatabases() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "What databases do exist right now?",

		ReadContext: dataSourceDatabasesRead,

		Schema: map[string]*schema.Schema{
			"databases": {
				// This description is used by the documentation generator and the language server.
				Description: "Existing databases",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}


func dataSourceDatabasesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := meta.(*f.FaunaClient)
	res, err := client.Query(f.Paginate(f.Databases()))
	if err != nil {
		return diag.FromErr(err)
	}
	var refs []f.RefV
	err = res.At(f.ObjKey("data")).Get(&refs)
	if err != nil {
		return diag.FromErr(err)
	}
	databases := make([]string, 0)
	for _, db := range refs {
		databases = append(databases, db.ID)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	err = d.Set("databases", databases)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

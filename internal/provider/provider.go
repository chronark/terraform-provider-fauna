package provider

import (
	"context"
	"fmt"

	f "github.com/fauna/faunadb-go/v3/faunadb"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"fauna_key": {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("FAUNA_KEY", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{},
			ResourcesMap: map[string]*schema.Resource{
				"fauna_collection": resourceCollection(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}
func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var diags diag.Diagnostics

		faunaKey := d.Get("fauna_key").(string)

		if faunaKey == "" {
			return nil, diag.FromErr(fmt.Errorf("fauna key is not set"))
		}

		client := f.NewFaunaClient(faunaKey)
		return client, diags
	}
}

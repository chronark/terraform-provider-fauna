package provider

import (
	"context"
	"fmt"
	f "github.com/fauna/faunadb-go/v3/faunadb"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"fauna_key": {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive: true,
					DefaultFunc: schema.EnvDefaultFunc("FAUNA_KEY", nil),
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"scaffolding_data_source": dataSourceScaffolding(),
				"fauna_databases":         dataSourceDatabases(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"scaffolding_resource": resourceScaffolding(),
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

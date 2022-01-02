package provider

import (
	"context"
	"mime"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceScaffolding() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceScaffoldingCreate,
		ReadContext:   resourceScaffoldingRead,
		UpdateContext: resourceScaffoldingUpdate,
		DeleteContext: resourceScaffoldingDelete,

		Schema: map[string]*schema.Schema{
			"input": {
				// This description is used by the documentation generator and the language server.
				Description: "input string",
				Required:    true,
				Type:        schema.TypeString,
			},
			"charset": {
				Description: "charset",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "utf-8",
			},
			"value": {
				Description: "value",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

func resourceScaffoldingCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	inputValue := d.Get("input").(string)
	charset := d.Get("charset").(string)
	encodedValue := mime.BEncoding.Encode(charset, inputValue)
	d.SetId(encodedValue)
	err := d.Set("value", encodedValue)
	if err != nil {
		return diag.Errorf("%s", err.Error())
	}

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Trace(ctx, "created a resource")

	return nil
}

func resourceScaffoldingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return nil
}

func resourceScaffoldingUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return nil
}

func resourceScaffoldingDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return nil
}

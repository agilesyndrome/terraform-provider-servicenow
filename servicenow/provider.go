package servicenow

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/tylerhatton/terraform-provider-servicenow/servicenow/client"
	"github.com/tylerhatton/terraform-provider-servicenow/servicenow/resources"
)

// Provider is a Terraform Provider to that manages objects in a ServiceNow instance.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"instance_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICENOW_INSTANCE_URL", nil),
				Description: "The Url of the ServiceNow instance to work with. " +
					"Value can also be sourced from the SERVICENOW_INSTANCE_URL environment variable.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICENOW_USER", nil),
				Description: "Username used to manage resources in the ServiceNow instance using Basic authentication. " +
					"Value can also be sourced from the SERVICENOW_USER environment variable.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICENOW_PASSWORD", nil),
				Description: "Password of the user to manage resources. " +
					"Value can also be sourced from the SERVICENOW_PASSWORD environment variable.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"servicenow_alias":                      resources.ResourceAlias(),
			"servicenow_application":                resources.ResourceApplication(),
			"servicenow_application_menu":           resources.ResourceApplicationMenu(),
			"servicenow_application_module":         resources.ResourceApplicationModule(),
			"servicenow_basic_auth_credential":      resources.ResourceBasicAuthCredential(),
			"servicenow_content_css":                resources.ResourceContentCSS(),
			"servicenow_css_include":                resources.ResourceCSSInclude(),
			"servicenow_css_include_relation":       resources.ResourceCSSIncludeRelation(),
			"servicenow_db_table":                   resources.ResourceDBTable(),
			"servicenow_extension_point":            resources.ResourceExtensionPoint(),
			"servicenow_http_connection":            resources.ResourceHttpConnection(),
			"servicenow_js_include":                 resources.ResourceJsInclude(),
			"servicenow_js_include_relation":        resources.ResourceJsIncludeRelation(),
			"servicenow_oauth_entity":               resources.ResourceOAuthEntity(),
			"servicenow_question_choice":            resources.ResourceQuestionChoice(),
			"servicenow_role":                       resources.ResourceRole(),
			"servicenow_rest_message":               resources.ResourceRestMessage(),
			"servicenow_rest_message_header":        resources.ResourceRestMessageHeader(),
			"servicenow_rest_method":                resources.ResourceRestMethod(),
			"servicenow_rest_method_header":         resources.ResourceRestMethodHeader(),
			"servicenow_scripted_rest_api":          resources.ResourceScriptedRestApi(),
			"servicenow_scripted_rest_resource":     resources.ResourceScriptedRestResource(),
			"servicenow_script_include":             resources.ResourceScriptInclude(),
			"servicenow_server":                     resources.ResourceServer(),
			"servicenow_service_catalog":            resources.ResourceServiceCatalog(),
			"servicenow_service_catalog_category":   resources.ResourceServiceCatalogCategory(),
			"servicenow_service_catalog_item":       resources.ResourceServiceCatalogItem(),
			"servicenow_service_catalog_variable":   resources.ResourceServiceCatalogVariable(),
			"servicenow_system_property":            resources.ResourceSystemProperty(),
			"servicenow_system_property_category":   resources.ResourceSystemPropertyCategory(),
			"servicenow_system_property_relation":   resources.ResourceSystemPropertyRelation(),
			"servicenow_ui_macro":                   resources.ResourceUIMacro(),
			"servicenow_ui_page":                    resources.ResourceUIPage(),
			"servicenow_ui_script":                  resources.ResourceUIScript(),
			"servicenow_widget":                     resources.ResourceWidget(),
			"servicenow_widget_dependency":          resources.ResourceWidgetDependency(),
			"servicenow_widget_dependency_relation": resources.ResourceWidgetDependencyRelation(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"servicenow_acl":                      resources.DataSourceACL(),
			"servicenow_application":              resources.DataSourceApplication(),
			"servicenow_application_category":     resources.DataSourceApplicationCategory(),
			"servicenow_db_table":                 resources.DataSourceDBTable(),
			"servicenow_role":                     resources.DataSourceRole(),
			"servicenow_service_catalog":          resources.DataSourceServiceCatalog(),
			"servicenow_service_catalog_category": resources.DataSourceServiceCatalogCategory(),
			"servicenow_system_property":          resources.DataSourceSystemProperty(),
			"servicenow_system_property_category": resources.DataSourceSystemPropertyCategory(),
		},
		ConfigureFunc: configure,
	}
}

func configure(data *schema.ResourceData) (interface{}, error) {
	// Create a new client to talk to the instance.
	client := client.NewClient(
		data.Get("instance_url").(string),
		data.Get("username").(string),
		data.Get("password").(string))

	return client, nil
}

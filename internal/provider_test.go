package internal

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
	istioclient "istio.io/client-go/pkg/clientset/versioned/fake"
	"k8s.io/client-go/rest"
)

func TestProvider(t *testing.T) {
	provider := Provider()

	assert.NotNil(t, provider, "Provider should not be nil")
	assert.IsType(t, &schema.Provider{}, provider, "Provider should be of type *schema.Provider")
}

func TestProviderConfigure(t *testing.T) {
	// Set environment variables for testing
	os.Setenv("KUBE_CONFIG", "test_config")
	os.Setenv("KUBECONFIG", "test_config")

	d := schema.TestResourceDataRaw(t, Provider().(*schema.Provider).Schema, map[string]interface{}{
		"config_path": "test_config",
	})

	config, err := providerConfigure(d)

	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, config, "Config should not be nil")
	assert.IsType(t, &Config{}, config, "Config should be of type *Config")

	// You can add more specific assertions here, like checking if the istioClientset is properly initialized
	// For example:
	assert.IsType(t, &istioclient.Clientset{}, config.(*Config).istioClientset, "istioClientset should be of type *istioclient.Clientset")
}

func TestProviderConfigure_EmptyConfig(t *testing.T) {
	// Test with empty config
	d := schema.TestResourceDataRaw(t, Provider().(*schema.Provider).Schema, map[string]interface{}{})

	config, err := providerConfigure(d)

	assert.Nil(t, err, "Error should be nil")
	assert.NotNil(t, config, "Config should not be nil")
	assert.IsType(t, &Config{}, config, "Config should be of type *Config")
	assert.IsType(t, &rest.Config{}, config.(*Config).istioClientset.Config, "istioClientset.Config should be of type *rest.Config")
}

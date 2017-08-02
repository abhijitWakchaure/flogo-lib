package property

import (
	"testing"
	"os"

	"github.com/stretchr/testify/assert"
)

// TestRegisterEmptyId register with empty id
func TestRegisterEmptyId(t *testing.T) {
	err := Register("", "a value")
	assert.NotNil(t, err)
}

// TestRegisterOk register ok
func TestRegisterOk(t *testing.T) {
	err := Register("id_test", "a value")
	assert.Nil(t, err)

	value := Get("id_test")
	assert.Equal(t, "a value", value)
}

// TestRegisterDuplicate register duplicate
func TestRegisterDuplicate(t *testing.T) {
	err := Register("id_test2", "a value")
	assert.Nil(t, err)

	err = Register("id_test2", "a value")
	assert.NotNil(t, err)
}

// TestRegisterEnvironmentOk register environment property
func TestRegisterEnvironmentOk(t *testing.T) {
	os.Setenv("TEST_FLOGO2", "my_test_value2")
	defer os.Unsetenv("TEST_FLOGO2")

	err := Register("id_environment", "${env.TEST_FLOGO2}")
	assert.Nil(t, err)

	value := Get("id_environment")
	assert.Equal(t, "my_test_value2", value)
}

// TestDefaultResolverOk resolves environment and property value
func TestDefaultResolverOk(t *testing.T) {
	os.Setenv("TEST_FLOGO2", "my_test_value2")
	defer os.Unsetenv("TEST_FLOGO2")
	Register("id_test", "a value")
	value := Get("${property.id_test}")
    assert.Equal(t, "a value", value)
	value = Get("${env.TEST_FLOGO2}")
	assert.Equal(t, "my_test_value2", value)
}
package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// You normally want to run this under a separate "Testing" subscription
// For lab purposes you will use your assigned subscription under the Cloud Dev/Ops program tenant
var subscriptionID string = "7c0d13e5-0d38-4a95-8775-910f0da1c263"

func TestAzureLinuxVMCreation(t *testing.T) {
	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
		// Override the default terraform variables
		Vars: map[string]interface{}{
			"labelPrefix": "rodr0304",
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of output variable
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

	// Confirm VM exists
	assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID))

	// Confirm NIC exists
	nicName := terraform.Output(t, terraformOptions, "nic_name")

	assert.True(
		t,
		azure.NetworkInterfaceExists(
			t,
			nicName,
			resourceGroupName,
			subscriptionID,
		),
	)

	// Confirm VM is running the correct Ubuntu version
	vm := azure.GetVirtualMachine(t, vmName, resourceGroupName, subscriptionID)

	assert.Equal(t, "Canonical", *vm.StorageProfile.ImageReference.Publisher)
	assert.Equal(t, "0001-com-ubuntu-server-jammy", *vm.StorageProfile.ImageReference.Offer)
	assert.Equal(t, "22_04-lts-gen2", *vm.StorageProfile.ImageReference.Sku)

	// Confirm NIC is connected to VM
	vmNics := azure.GetVirtualMachineNics(
		t,
		vmName,
		resourceGroupName,
		subscriptionID,
	)

	assert.Contains(t, vmNics, nicName)
}

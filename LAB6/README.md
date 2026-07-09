# CST8918 - Lab 6: Terraform + Terratest

## Overview

This lab demonstrates Infrastructure as Code (IaC) testing using Terraform and Terratest on Microsoft Azure.

The project provisions:

- Resource Group
- Virtual Network
- Subnet
- Public IP
- Network Security Group
- Network Interface
- Ubuntu 22.04 LTS Virtual Machine

The Terratest suite validates:

- VM creation
- Network Interface existence
- Ubuntu image version

---

# Challenges Encountered

## 1. Azure Region Capacity

### Issue

The initial deployment failed because the selected Azure region did not have enough capacity for the requested VM size.

### Solution

Changed the deployment region to one with available capacity.

---

## 2. GitHub Push Failed (Large Files)

### Issue

GitHub rejected the push because Terraform provider binaries inside the `.terraform` directory exceeded GitHub's 100 MB file size limit.

Example:

```
terraform-provider-azurerm_v3.0.2_x5
209 MB
```

### Solution

Created a `.gitignore` file.

Ignored:

```
.terraform/
terraform.tfstate
terraform.tfstate.backup
.terraform.lock.hcl
```

Removed cached files:

```bash
git rm --cached -r .terraform
git rm --cached terraform.tfstate
git rm --cached terraform.tfstate.backup
git rm --cached .terraform.lock.hcl
```

Rebuilt the commit history using:

```bash
git reset --mixed HEAD~2
```

---

## 3. Go Installation

### Issue

The system did not recognize the `go` command.

```
zsh: command not found: go
```

### Solution

Installed Go using Homebrew.

```bash
brew install go
```

---

## 4. Terratest Dependency Conflict

### Issue

Running `go mod tidy` generated an ambiguous import error.

```
ambiguous import:
github.com/gruntwork-io/terratest/modules/terraform
```

### Solution

Pinned Terratest to version **v0.46.11**, which is compatible with the lab starter code.

---

## 5. Placeholder Variables

### Issue

The starter test used placeholder values.

```
<your-college-id>
```

Terraform rejected the resource names.

### Solution

Updated:

```go
subscriptionID := "7c0d13e5-..."
```

and

```go
labelPrefix := "rodr0304"
```

---

## 6. Resources Automatically Destroyed

### Issue

The Terratest automatically destroys Azure resources after a successful test.

Therefore, Azure CLI commands returned:

```
ResourceGroupNotFound
```

### Solution

Executed:

```bash
terraform apply
```

manually before validating the resources using Azure CLI.

---

## 7. Terraform Executed from Wrong Directory

### Issue

Terraform returned:

```
Error: No configuration files
```

### Cause

The command was executed from the repository root instead of the LAB6 directory.

### Solution

```bash
cd LAB6
terraform apply
```

---

# Validation

## VM Image

```bash
az vm show \
  --resource-group rodr0304-A05-RG \
  --name rodr0304A05VM \
  --query "storageProfile.imageReference"
```

Output:

```
Publisher : Canonical
Offer     : 0001-com-ubuntu-server-jammy
SKU        : 22_04-lts-gen2
```

---

## NIC Attached

```bash
az vm nic list \
  --resource-group rodr0304-A05-RG \
  --vm-name rodr0304A05VM
```

Output:

```
Primary : true
```

---

# Test Execution

```
go test -v
```

Result:

```
PASS
ok github.com/rodr0304/CST8918
```

---

# Lessons Learned

This lab provided practical experience with:

- Terraform state management
- Azure Infrastructure as Code
- Terratest integration
- Go module management
- Git history cleanup
- GitHub large file limitations
- Azure CLI resource validation
- Automated infrastructure testing

# CST8918 - Hybrid A09
## Terraform Validation with Husky and GitHub Actions

**Student:** Diniz Martins  
**Course:** CST8918 – DevOps: Infrastructure as Code  

---

# Objective

The objective of this assignment was to create a basic CI/CD pipeline for a Terraform project using GitHub Actions and Husky.

The solution validates Terraform code before commits and during Pull Requests to ensure formatting, syntax validation, and static analysis.

---

# Project Structure

```
Lab-A09/
│
├── infrastructure/
│   ├── main.tf
│   ├── providers.tf
│   ├── variables.tf
│   ├── outputs.tf
│   └── .terraform.lock.hcl
│
├── README.md
├── package.json
└── package-lock.json

.github/
└── workflows/
    └── action-terraform-verify.yml

.husky/
└── pre-commit
```

---

# Terraform Resource

A simple Azure Resource Group was created to validate the Terraform configuration.

The project includes:

- Resource Group
- Variables
- Outputs
- AzureRM Provider

---

# Husky Pre-Commit Hook

The pre-commit hook performs the following validations before allowing a commit:

- Terraform Format Check

```bash
terraform fmt -check -recursive
```

- Terraform Initialization

```bash
terraform init -backend=false
```

- Terraform Validation

```bash
terraform validate
```

- TFLint

```bash
tflint --init
tflint
```

If any validation fails, the commit is blocked.

---

# GitHub Actions Workflow

A GitHub Actions workflow was created to automatically validate every Pull Request.

The workflow executes three independent jobs:

- Terraform Format Check
- Terraform Validate
- TFLint

The workflow only runs on Pull Requests targeting:

- main
- master

---

# Testing Performed

The following tests were completed successfully:

## Husky

✔ Introduced formatting errors in the Terraform code.

✔ Verified that Husky blocked the commit.

✔ Corrected the formatting.

✔ Successfully committed after fixing the errors.

---

## GitHub Actions

✔ Created a Pull Request.

✔ Pushed intentionally invalid Terraform formatting.

✔ Confirmed that GitHub Actions failed.

✔ Corrected the formatting.

✔ Pushed the changes again.

✔ Verified that all GitHub Actions checks passed.

---

# Challenges Encountered

Several issues were encountered during development:

- Incorrect GitHub Actions working directory.
- Duplicate workflow files created during development.
- Husky path configuration needed to be adjusted.
- Terraform files were initially committed in the wrong location.
- GitHub Actions initially searched for the `infrastructure` folder in the repository root instead of `Lab-A09/infrastructure`.
- Multiple rounds of testing were required before all workflow jobs passed successfully.

Each issue was resolved by reviewing the project structure, updating workflow paths, and validating the configuration locally before pushing changes.

---

# Technologies Used

- Terraform
- AzureRM Provider
- Git
- GitHub
- GitHub Actions
- Husky
- TFLint
- Visual Studio Code

---

# Final Result

The project successfully implements:

- Local Terraform validation using Husky.
- Automatic Pull Request validation using GitHub Actions.
- Terraform formatting verification.
- Terraform syntax validation.
- Static code analysis with TFLint.

All GitHub Actions checks completed successfully.

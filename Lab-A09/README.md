# CST8918 - Hybrid A09
## Husky + GitHub Actions for Terraform Validation

### Student
**Diniz Martins**

---

# Objective

The purpose of this assignment was to create a CI/CD pipeline for a Terraform Infrastructure as Code (IaC) project using **GitHub Actions** and **Husky**.

The project validates Terraform code before commits and during Pull Requests.

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

# Technologies Used

- Terraform
- Azure
- GitHub Actions
- Husky
- TFLint
- Git

---

# Part 1 - Husky Pre-Commit Hook

The Husky pre-commit hook was configured to execute:

- Terraform Format Check (`terraform fmt`)
- Terraform Validation (`terraform validate`)
- TFLint

This prevents incorrectly formatted Terraform code from being committed.

---

# Part 2 - GitHub Actions Workflow

A GitHub Actions workflow was created to automatically execute:

- Terraform Format Check
- Terraform Validate
- TFLint

whenever a Pull Request is opened against the **main** branch.

---

# Testing the Workflow

The assignment required intentionally breaking the Terraform code to verify that the pipeline correctly detects formatting problems.

## Step 1 – Create a Pull Request

A Pull Request was created to test the GitHub Actions workflow.

![Pull Request](Screenshots/SS1%20-%20OK.png)

---

## Step 2 – Initial Pipeline Failure

The workflow initially failed because Terraform formatting was intentionally broken.

This confirms that the GitHub Actions workflow correctly detected formatting issues.

![Pipeline Failed](Screenshots/SS5%20-%20NOK.png)

---

## Step 3 – TFLint Failure

During testing, another issue was discovered.

The TFLint job was configured to execute inside an incorrect working directory.

GitHub Actions returned the following error.

![TFLint Failure](Screenshots/SS4%20-%20NOK.png)

---

## Step 4 – Investigating the Error

The workflow logs were analyzed to identify the root cause.

The problem was caused by an incorrect Terraform working directory.

After updating the workflow, the TFLint job executed successfully.

![Workflow Logs](Screenshots/SS3%20-%20NOK.png)

---

## Step 5 – Successful Pipeline

After fixing:

- Terraform formatting
- GitHub Actions workflow
- Working directory configuration

all GitHub Actions jobs completed successfully.

✔ Terraform Format Check

✔ Terraform Validate

✔ TFLint

![Successful Pipeline](Screenshots/SS2%20-%20OK.png)

---

# Challenges

During this assignment, several practical issues were encountered:

- Configuring Husky with an existing repository.
- Updating Husky to a compatible version.
- Correctly configuring Terraform inside a subfolder.
- Configuring GitHub Actions to execute Terraform commands inside the `Lab-A09/infrastructure` directory.
- Debugging GitHub Actions logs to identify workflow failures.
- Fixing the TFLint working directory configuration.

These challenges helped improve understanding of CI/CD pipelines and automated Terraform validation.

---

# Conclusion

The project successfully demonstrates a complete Terraform validation pipeline using:

- Husky Pre-Commit Hooks
- GitHub Actions
- Terraform Format Validation
- Terraform Validation
- TFLint

The workflow automatically prevents improperly formatted or invalid Terraform code from being merged into the main branch.

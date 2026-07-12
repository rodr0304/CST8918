# CST8918 - Hybrid H09 - Azure Kubernetes Service (AKS) with Terraform


## Overview

This project demonstrates how to provision an Azure Kubernetes Service (AKS) cluster using Terraform and deploy a sample multi-tier application using Kubernetes.

The infrastructure was created as Infrastructure as Code (IaC) using Terraform, and the application was deployed using Kubernetes manifests.

---

## Technologies

- Terraform
- Microsoft Azure
- Azure Kubernetes Service (AKS)
- Kubernetes
- kubectl

---

## Infrastructure

The Terraform configuration creates:

- Azure Resource Group
- Azure Kubernetes Service (AKS)
- Latest supported Kubernetes version
- SystemAssigned Managed Identity
- Auto-scaling node pool
- Minimum: 1 node
- Maximum: 3 nodes
- VM Size: Standard_B2s

---

## Project Files

```
.
├── providers.tf
├── main.tf
├── outputs.tf
├── sample-app.yaml
├── .gitignore
├── .terraform.lock.hcl
└── README.md
```

---

## Deployment Steps

Initialize Terraform:

```bash
terraform init
```

Review the execution plan:

```bash
terraform plan
```

Deploy the infrastructure:

```bash
terraform apply
```

Generate the Kubernetes configuration file:

```bash
terraform output -raw kube_config > kubeconfig
```

Configure kubectl:

```bash
export KUBECONFIG=./kubeconfig
```

Verify the AKS cluster:

```bash
kubectl get nodes
```

Deploy the sample application:

```bash
kubectl apply -f sample-app.yaml
```

Verify the deployment:

```bash
kubectl get pods
kubectl get services
```

Access the application using the External IP assigned to the **store-front** LoadBalancer service.

---

## Sample Application

The deployed application includes:

- RabbitMQ
- Order Service (Node.js)
- Product Service (Rust)
- Store Front (Vue.js)

---

## Screenshots

### PORTAL

![AKS Cluster](screenshots/Screenshot%202026-07-12%20at%2009.57.01.png)

---

### CART

![Pods](screenshots/Screenshot%202026-07-12%20at%2009.57.18.png)

---

### Products

![Services](screenshots/Screenshot%202026-07-12%20at%2009.57.26.png)

---

### Application Running

![Application](screenshots/Screenshot%202026-07-12%20at%2009.58.02.png)

---

### Application Running

![Shopping Cart](screenshots/Screenshot%202026-07-12%20at%2010.05.21.png)

---

---

## Cleanup

To avoid Azure charges, destroy the infrastructure after completing the lab:

```bash
terraform destroy
```


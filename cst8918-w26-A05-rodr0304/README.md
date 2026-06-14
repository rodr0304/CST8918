# CST8918 Lab A05 - Terraform Web Server

## Overview

This lab demonstrates the deployment of a simple Apache web server on Microsoft Azure using Terraform Infrastructure as Code (IaC).

The solution includes:

* Azure Resource Group
* Virtual Network (VNet)
* Subnet
* Public IP Address
* Network Security Group (NSG)
* Network Interface Card (NIC)
* Ubuntu 22.04 Virtual Machine
* Apache Web Server installed using Cloud-Init

## Architecture

The architecture diagram is provided in:

* a05-architecture.png

## Deployment Verification

The deployment was verified by:

* Accessing the Apache default web page through the public IP address
* Connecting to the virtual machine using SSH
* Confirming the Apache service was running

A screenshot of the verification process is provided in:

* a05-demo.png

## Files

* main.tf — Terraform infrastructure definitions
* init.sh — Cloud-init script used to install Apache
* .gitignore — Git ignore configuration
* a05-architecture.png — Architecture diagram
* a05-demo.png — Deployment verification screenshot

## Author

Diniz Rodrigues Martins
Student ID: 041179475


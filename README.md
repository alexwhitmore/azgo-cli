# Azure CLI Written in Go

## Description

A small CLI written in Go and Cobra to create and deploy a static site to Azure.

### Commands

| Command | Description | Status |
| --------- | --------- | --------- |
| create    | Create and deploy a static site to Azure from an existing GitHub repository. | Done |
| deploy    | Deploy a static site to an existing Azure static site. | Not done, handled through GitHub actions which will likely change. |
| secrets    | Upload required environment variables to an existing Azure static site. | Done |
| delete    | Delete an Azure static site. | Not done. |
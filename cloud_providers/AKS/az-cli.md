# Azure CLI

- [Cli Onboarding Cheat Sheet](https://learn.microsoft.com/en-us/cli/azure/cheat-sheet-onboarding)
- [Azure CLI Queries](https://learn.microsoft.com/en-us/cli/azure/use-azure-cli-successfully-query?tabs=examples%2Cbash)
- [JMESPath for CLI Queries](https://jmespath.org/tutorial.html)

## Install

```bash
brew install azure-cli
# docs suggest this has some azure bits, I didn't look into it further
brew install kubelogin
# extra cli bits for aks
az aks install-cli
```

## Queries

```bash
# query can be used to reduce the info 
# this example still reponds with a JSON object,
# but only with the requested keys
az account show --query "{tenantId:tenantId,subscriptionid:id}"
```

## Login & Subscription 

```bash
# help menu
az
# login first
az login
# account has few commands
az account set --subscription <subscription_id>
az account show
# {
#   "environmentName": "AzureCloud",
#   "isDefault": true,
#   "managedByTenants": [],
#   "name": "benpetersen",
#   "state": "Enabled",
#   "tenantDefaultDomain": "microsoft.onmicrosoft.com",
#   "tenantDisplayName": "Microsoft",
#   "user": {
#     "name": "benpetersen@microsoft.com",
#     "type": "user"
#   }
# }
# 
# easy view with output flag
az account show --output table
# list, can also change output with flag
az account list --output table
#
# gets the JWT access token 
az account get-access-token
# {
#   "accessToken": "eyJ0J9.eyJ0N30.UHqMA",
#   "expiresOn": "2024-11-05 21:29:16.000000",
#   "expires_on": 1730860156,
#   "subscription": "0990639e-03f6-4b45-8fdd-143da61a6e2d",
#   "tenant": "72f988bf-86f1-41af-91ab-2d7cd011db47",
#   "tokenType": "Bearer"
# }

# subscriptions you have access to 
az account list --query "[].{name:name, id:id}" --output tsv
# management group name
az account management-group list --query "[].{name:name, id:id}" --output tsv
```

## Active Directory

```bash

# active directory 
az ad signed-in-user show
az ad user show --id "{principalName}" --query "id" --output tsv
az ad user show --id "benpetersen@microsoft.com" --query "id" --output tsv
# group
az ad group show --group "{groupName}" --query "id" --output tsv
# service principal
az ad sp list --all --query "[].{displayName:displayName, id:id}" --output tsv
az ad sp list --display-name "{displayName}"
# managed identity
az ad sp list --all --filter "servicePrincipalType eq 'ManagedIdentity'"
# user user assigned managed ids
az identity list
```

# RBAC

RBAC Roles. 

```bash
# this will give lots of role options
az role definition list --query "[].{name:name, roleType:roleType, roleName:roleName}" --output tsv
# output of a single role, by name
az role definition list --name "Key Vault Administrator"
```

Scope is also needed when dealing with RBAC.  Azure provides 
four levels of scope: 
- resource
  - `/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceSubType}/{resourceName}`
  - resource IDs have a lengthy format structured like a URL
- resource group
  - `az group list --query "[].{name:name}" --output tsv`
  - my-rg1
- subscription
  - `az account list --query "[].{name:name, id:id}" --output tsv`
  - username 12345-12345-12345-12345
- management group
  - `az account management-group list --query "[].{name:name, id:id}" --output tsv`
  - might get a 'not authorized' message, depending on your abilities
    within the subscription, etc.
```bash

```


## Fancy AI 

```bash
# auto-completion, command descriptions, and examples as you type
az interactive
# get guides based on a search!
az scenario guide "keywords of interest"
az scenario guide "virtual machines"
# uhh... what do I do now?
# based on patterns will suggest to you the next steps
az next
# more AI looking things up for you
az find "active subscription"
```

## Resource Group

```bash
# set the defaults
az config set defaults.location=westus2 defaults.group=MyResourceGroup
```

## Resource Group

```bash
# create a resource group
az group create --name my-rg1 --location eastus2
```

## AKS

NOTE: a resource group is required.

```bash
# create an aks cluster in the group
az aks create --name my-aks1 \ 
  --resource-group my-rg1 \ 
  --enable-addons azure-keyvault-secrets-provider \
  --enable-oidc-issuer \ 
  --enable-workload-identity \
  --generate-ssn-keys

az aks enable-addons \
  --addons azure-keyvault-secrets-provider \
  --name my-aks1 \
  --resource-group my-rg1

# list clusters
az aks list
# extra cli bits for aks
az aks install-cli
# get the kubeconfig
az aks get-credentials --name my-aks1 --resource-group my-rg1
# verify
kubectl config get-contexts

# check if the addons have been properly installed,
# per the above flags
kubectl get pods -n kube-system -l 'app in (secrets-store-csi-driver,secrets-store-provider-azure)'
```

## Key Vault

```bash
# new key vault
az keyvault create --name my-kv1 --resource-group my-rg1 \
  --location eastus2 --enable-rbac-authorization 

# update flags on existing key vault 
az keyvault update --name my-kv1 --resource-group my-rg1 \
  --location eastus2 --enable-rbac-authroization

# create secrets
az keyvault secret set --vault-name my-kv1 --name my-secret --value my-value
# get a list of secret w/o values in json format
az keyvault secret list --vault-name my-kv1 
# get a value
az keyvault secret get --vault-name my-kv1 --ame my-secret
```

## Service Connector

I had trouble getting this to work via the CLI, I would get 
errors rather than successfully created connectors.  When I 
created via the Azure portal, it seemd to work.

```bash
# interactive, will ask for all the extra bits
az aks connection create keyvault --new

# when all flags provided
az aks connection create keyvault \
    --connection <connection-name> \
    --resource-group <resource-group-of-cluster> \
    --name <cluster-name> \
    --target-resource-group <resource-group-of-key-vault> \
    --vault <key-vault-name> \
    --enable-csi \
    --client-type none

# list
az aks connection list --resource-group <rg-name> --name <connection-name> --output table
# show a specific one
az aks connection show --connection <name> --resource-group <rg-name> --name <i-think-this-is-aks-cluster-name-which-is-awkward>
# I would expect something more like
az aks connection show <connection-name> --resource-group <rg-name> --cluster <aks-cluster-name>
```
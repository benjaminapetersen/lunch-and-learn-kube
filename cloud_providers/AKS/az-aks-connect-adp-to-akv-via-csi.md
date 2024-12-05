# Connect your Azure identity provider to the Azure Key Vault Secrets Store CSI Driver in Azure Kubernetes Service (AKS)

- [Connect your Azure identity provider to the Azure Key Vault Secrets Store CSI Driver in Azure Kubernetes Service (AKS)](https://learn.microsoft.com/en-us/azure/aks/csi-secrets-store-identity-access?tabs=azure-portal&pivots=access-with-a-microsoft-entra-workload-identity)

## TL;DR

Nilekh said to use this, but via the center `Workload ID` tab of the choices:
- Service Connector with Workload ID
- `Workload ID`
- Managed Identity

This will take care of the appropriate path to then use Workload ID for other
things.

## Environment Variables

This is easiest if setup with ENV Vars, otherwise have to 
copy/paste all the values

```bash
export RESOURCE_GROUP="myResourceGroup"
export LOCATION="eastus"
export CLUSTER_NAME="myAKSCluster"
export SERVICE_ACCOUNT_NAMESPACE="default"
export SERVICE_ACCOUNT_NAME="workload-identity-sa"
export SUBSCRIPTION="$(az account show --query id --output tsv)"
export USER_ASSIGNED_IDENTITY_NAME="myIdentity"
export FEDERATED_IDENTITY_CREDENTIAL_NAME="myFedIdentity"
# Include these variables to access key vault secrets from a pod in the cluster.
export KEYVAULT_NAME="keyvault-workload-id"
export KEYVAULT_SECRET_NAME="my-secret"
```

## Create a Managed Identity

```bash
# need subscription id
#   0990639e-03f6-4b45-8fdd-143da61a6e2d
# need cluster name
#   akv-on-aks-with-sscsi
# get the resource group
#   benpetersen-rg-akv-on-aks
# to create and cache as a env var
#   export AKS_OIDC_ISSUER="$(az aks show --name "${CLUSTER_NAME}" \
#     --resource-group "${RESOURCE_GROUP}" \
#     --query "oidcIssuerProfile.issuerUrl" \
#     --output tsv)"
# 
# filled out
export AKS_OIDC_ISSUER="$(az aks show --name "akv-on-aks-with-sscsi" \
    --resource-group "benpetersen-rg-akv-on-aks" \
    --query "oidcIssuerProfile.issuerUrl" \
    --output tsv)"

# print it 
echo $AKS_OIDC_ISSUER
# https://eastus2.oic.prod-aks.azure.com/72f988bf-86f1-41af-91ab-2d7cd011db47/43f9a1e4-5ff7-4f9e-a995-2f168afa6576/

# to create a managed identity
#    az identity create \
#     --name "${USER_ASSIGNED_IDENTITY_NAME}" \
#     --resource-group "${RESOURCE_GROUP}" \
#     --location "${LOCATION}" \
#     --subscription "${SUBSCRIPTION}"
az identity create \
    --name "akv-on-aks-with-sscsi" \
    --resource-group "benpetersen-rg-akv-on-aks" \
    --location "eastus2" \
    --subscription "0990639e-03f6-4b45-8fdd-143da61a6e2d"
# {
#   "clientId": "fb560e94-db28-4336-abf8-fcde3b853aa3",
#   "id": "id",
#   "location": "eastus2",
#   "name": "akv-on-aks-with-sscsi",
#   "principalId": "f8b4c3ff-6a01-422b-809f-6796fc2dbe81",
#   "resourceGroup": "benpetersen-rg-akv-on-aks",  
#   "type": "Microsoft.ManagedIdentity/userAssignedIdentities"
# }

# Get the credential of the managed identity created
# interestingly, this generates a kubeconfig file
# an will merge it to your .kube/config list
# az aks get-credentials --name "${CLUSTER_NAME}" --resource-group "${RESOURCE_GROUP}"
az aks get-credentials --name "akv-on-aks-with-sscsi" --resource-group "benpetersen-rg-akv-on-aks"
# Merged "akv-on-aks-with-sscsi" as current context in /Users/benjaminpetersen/.kube/config

kubectl config get-contexts
# CURRENT   NAME                        CLUSTER                     AUTHINFO                                                      NAMESPACE
# *         akv-on-aks-with-sscsi       akv-on-aks-with-sscsi       clusterUser_benpetersen-rg-akv-on-aks_akv-on-aks-with-sscsi
#           kind-basic-pod-playground   kind-basic-pod-playground   kind-basic-pod-playground
#           kind-ss-csi-driver          kind-ss-csi-driver          kind-ss-csi-driver
```
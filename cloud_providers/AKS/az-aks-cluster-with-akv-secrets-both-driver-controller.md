# AKS with AKV and SS CSI Driver

- [AKS with AKV and SS CSI Driver](https://learn.microsoft.com/en-us/azure/aks/csi-secrets-store-driver)
  - [Connect an Identity to the AKV SS CSI Driver in AKS](https://learn.microsoft.com/en-us/azure/aks/csi-secrets-store-identity-access?tabs=azure-portal&pivots=access-with-service-connector)
    - This doesn't seem to work well, the error is unhelpful
- [Azure Home](https://ms.portal.azure.com/#home)
- [Azure Services](https://ms.portal.azure.com/#allservices/category/All)
- [AKS Clusters](https://ms.portal.azure.com/#browse/Microsoft.ContainerService%2FmanagedClusters)
- [Azure Key Vaults](https://ms.portal.azure.com/#browse/Microsoft.KeyVault%2Fvaults)
- [Azure CLI Conceptual Article List](https://learn.microsoft.com/en-us/cli/azure/reference-docs-index)
  - [Authenticate via CLI](https://learn.microsoft.com/en-us/cli/azure/authenticate-azure-cli)
  - [RBAC Role Assignments via CLI](https://learn.microsoft.com/en-us/azure/role-based-access-control/role-assignments-cli)
- [Azure CLI Success](https://learn.microsoft.com/en-us/cli/azure/use-azure-cli-successfully-tips?tabs=bash%2Ccmd2)
  - [Azure CLI Behind a Proxy](https://learn.microsoft.com/en-us/cli/azure/use-azure-cli-successfully-troubleshooting#work-behind-a-proxy)
- [Azure Workload Identity](https://learn.microsoft.com/en-us/entra/workload-id/workload-identities-overview)

The above link provides references and tutorials for this topic.

```bash
# seems this is needed as it is a client-go credential (exec) plugin implementing azure authentication
brew install kubelogin

# create a new resource group
az group create --name benpetersen-rg-akv-on-aks --location eastus2
# create an AKS cluster, with the ss csi driver provider
# as well as OIDC and workload identity
az aks create --name akv-on-aks-with-sscsi --resource-group benpetersen-rg-akv-on-aks --enable-addons azure-keyvault-secrets-provider --enable-oidc-issuer --enable-workload-identity --generate-ssh-keys
# list
az aks list
# get the cli extra bits
az aks install-cli
```

To add SS CSI to an existing cluster

```bash
az aks enable-addons --addons azure-keyvault-secrets-provider --name <cluster-name> --resource-group <rg-name>
```

Verify 

```bash
# will get a kubeconfig and inject it into ~/.kube/config
az aks get-credentials --name akv-on-aks-with-sscsi --resource-group benpetersen-rg-akv-on-aks

# now check what is running (provided that you are talking to the correct cluster, of course)
kubectl config get-contexts
kubectl get pods -n kube-system -l 'app in (secrets-store-csi-driver,secrets-store-provider-azure)'
# should be several od each pod
```

Then setup a key vault

```bash
## Create a new Azure key vault
#  Enables Azure RBAC with the flag
az keyvault create --name <keyvault-name> --resource-group myResourceGroup --location eastus2 --enable-rbac-authorization
az keyvault create --name akv-on-aks-with-sscsi --resource-group benpetersen-rg-akv-on-aks --location eastus2 --enable-rbac-authorization

## Update an existing Azure key vault
az keyvault update --name <keyvault-name> --resource-group myResourceGroup --location eastus2 --enable-rbac-authorization
az keyvault update --name akv-on-aks-with-sscsi --resource-group benpetersen-rg-akv-on-aks --location eastus2 --enable-rbac-authorization

# Create a random secret in the key vault 
az keyvault secret set --vault-name akv-on-aks-with-sscsi --name ExampleSecret --value MyAKSExampleSecret
# RBAC FAIL!
```

Fix RBAC fail, if it happens

- [Azure RBAC via CLI](https://learn.microsoft.com/en-us/azure/role-based-access-control/role-assignments-cli)

```bash
# whoami equiv
az ad signed-in-user show
# {
#   "@odata.context": "https://graph.microsoft.com/v1.0/$metadata#users/$entity",
#   "businessPhones": [],
#   "displayName": "Ben Petersen",
#   "givenName": "Ben",
#   "id": "19a81a6e-6932-4fbe-848f-c35cc5ea0997",
#   "jobTitle": "Senior Product Manager",
#   "mail": "benpetersen@microsoft.com",
#   "mobilePhone": null,
#   "officeLocation": "Home Office",
#   "preferredLanguage": null,
#   "surname": "Petersen",
#   "userPrincipalName": "benpetersen@microsoft.com"
# }
az ad user show --id "{principalName}" --query "id" --output tsv
# just the Id, if you know your principalName, or need to pass it to another cmd
az ad user show --id "benpetersen@microsoft.com" --query "id" --output tsv
# 19a81a6e-6932-4fbe-848f-c35cc5ea0997
# group
az ad group show --group "{groupName}" --query "id" --output tsv
# service principal (identities used by applications)
az ad sp list --all --query "[].{displayName:displayName, id:id}" --output tsv
az ad sp list --display-name "{displayName}"
# managed identity
az ad sp list --all --filter "servicePrincipalType eq 'ManagedIdentity'"
# just user assigned managed identities
az identity list
# this will give lots of role options
az role definition list --query "[].{name:name, roleType:roleType, roleName:roleName}" --output tsv
# output of a single role, by name
az role definition list --name "Key Vault Administrator"
```

Scope is also needed when dealing with RBAC.  Azure provides 
four levels of scope: 
- resource
- resource group
- subscription
- management group

For each of these:

```yaml
# format of resource ID is this
resource scope: /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceSubType}/{resourceName}
# just get the name of RGs you have access to
resource group scope: az group list --query "[].{name:name}" --output tsv
# benpetersen-rg-akv-on-aks                                    // I've been using this one
# MC_benpetersen-rg-akv-on-aks_akv-on-aks-with-sscsi_eastus2   // Not sure what this is, must be auto
# benpetersen-rg                                               // Made this a while ago, generic
# MC_benpetersen-kube-rg_kube-1_eastus                         // Also made a while ago
# benpetersen-kube-rg                                          // Also made a while ago 
subscription scope: az account list --query "[].{name:name, id:id}" --output tsv
# benpetersen	0990639e-03f6-4b45-8fdd-143da61a6e2d
management group scope: az account management-group list --query "[].{name:name, id:id}" --output tsv
# Not authorized to perform action 'Microsoft.Management/managementGroups/read' over scope '/providers/Microsoft.Management'
```

```bash
# MC_benpetersen-rg-akv-on-aks_akv-on-aks-with-sscsi_eastus2
# subscriptions you have access to
az account list --query "[].{name:name, id:id}" --output tsv

# management group name
az account management-group list --query "[].{name:name, id:id}" --output tsv
```

So to fix RBAC to allow my main acct to create Key Vaults:

```bash
# setting RBAC role
# its interesting my ID may have the ability to set RBAC
# even if it doesn't have the specific RBAC to do the thing
az role assignment create --assignee "{assignee}" \
    --role "{roleNameOrId}" \
    --scope "/subscriptions/{subscriptionId}"
# utilizing the subscription output for me
#     benpetersen	0990639e-03f6-4b45-8fdd-143da61a6e2d
# and also my userPrincipalName for assignee
#     benpetersen@microsoft.com
az role assignment create --assignee "benpetersen@microsoft.com" \
    --role "Key Vault Administrator" \
    --scope "/subscriptions/0990639e-03f6-4b45-8fdd-143da61a6e2d"
```

Now back to my key vaults

```bash
# I have a couple.
az keyvault list
# this one is the one I've been tinkering with 
#   {
#     "location": "eastus2",
#     "name": "akv-on-aks-with-sscsi",
#     "resourceGroup": "benpetersen-rg-akv-on-aks",
#     "type": "Microsoft.KeyVault/vaults"
#   },
az keyvault secret set --vault-name akv-on-aks-with-sscsi --name ExampleSecret --value MyAKSExampleSecret
az keyvault secret set --vault-name akv-on-aks-with-sscsi --name ninnyhammers --value "You're nowt but a ninnyhammer, Sam Gamgee!"
# list secrets, no values, but JSON objects with metadata
az keyvault secret list --vault-name akv-on-aks-with-sscsi
# retrieval of one secret, including value
az keyvault secret get --vault-name akv-on-aks-with-sscsi --name ninnyhammers
```

## Connect Identity to CSI Driver: With a Service Connection?

- [Connect an Identity to the AKV SS CSI Driver in AKS](https://learn.microsoft.com/en-us/azure/aks/csi-secrets-store-identity-access?tabs=azure-portal&pivots=access-with-service-connector)

Ok, now we need to connect an identty to SS CSI Driver running on our AKS cluster
so that it can access the secrets.  There are several options for identity-based
access to an Azure Key Vault:

- Service Connector with Workload ID
- Workload ID
- User-assigned managed identity

```bash
# recall again all the bits done above:
#
# a resource group (we've been using)
az group list --query "[].{name:name}" --output tsv
# benpetersen-rg-akv-on-aks
# 
# a cluster
az aks list --output table
# Name                   Location    ResourceGroup              Version  Fqdn
# ----                   ----------  --------------             -------  -----------------------------------------------------------------------------------------
# akv-on-aks-with-sscsi  eastus2     benpetersen-rg-akv-on-aks  1.29     akv-on-aks-benpetersen-rg-a-099063-9rfmdvqg.hcp.eastus2.azmk8s.io
#
# get a Kubeconfig
az aks get-credentials \
    --resource-group benpetersen-rg-akv-on-aks \
    --name akv-on-aks-with-sscsi
#
# check context
kubectl config get-contexts
#
# driver is running on cluster already?
kubectl get pods -n kube-system -l 'app in (secrets-store-csi-driver,secrets-store-provider-azure)'
#
# got a key vault?
az keyvault list --output table
# Location    Name                   ResourceGroup
# ----------  ---------------------  -------------------------
# eastus2     akv-on-aks-with-sscsi  benpetersen-rg-akv-on-aks
#
# any secrts in the key vault?
az keyvault secret list --vault-name akv-on-aks-with-sscsi --output table
# Name
# --------
# ExampleSecret
# Ninnyhammers
```

## Service Connector (Preview) to Wire Things Up

And now to make a service connection in AKS with Service Connector (`Preview`).

```bash
# template:
az aks connection create keyvault \
    --connection <connection-name> \
    --resource-group <resource-group-of-cluster> \
    --name <cluster-name> \
    --target-resource-group <resource-group-of-key-vault> \
    --vault <key-vault-name> \
    --enable-csi \
    --client-type none
# filled out:
# sadly, horribly, connection cannot have hyphens.
# so for this one we have to swap to underscores.
az aks connection create keyvault \
    --connection akv_on_aks_with_sscsi \ 
    --resource-group benpetersen-rg-akv-on-aks \
    --name akv-on-aks-with-sscsi \
    --target-resource-group benpetersen-rg-akv-on-aks \
    --vault akv-on-aks-with-sscsi \
    --enable-csi \
    --client-type none
# one line
# unfortunately, this doesn't work
az aks connection create keyvault --connection akv_on_aks_with_sscsi --resource-group benpetersen-rg-akv-on-aks --name akv-on-aks-with-sscsi --target-resource-group benpetersen-rg-akv-on-aks --vault akv-on-aks-with-sscsi --enable-csi --client-type none    
# cli.azure.cli.core.azclierror: (GeneralUserError) Execution failed. Operation returned an invalid status code 'Conflict'
Code: GeneralUserError
az aks connection list --resource-group benpetersen-rg-akv-on-aks --name akv-on-aks-with-sscsi --output table
# Command group 'aks connection' is in preview and under development. Reference and support levels: https://aka.ms/CLI_refstatus
# ClientType    Name                   ProvisioningState    ResourceGroup              Scope
# ------------  ---------------------  -------------------  -------------------------  -------
# none          akv_on_aks_with_sscsi  Failed               benpetersen-rg-akv-on-aks  default
```

So, that didn't work well, but it appears we can use an interactive flow:

```bash
# this will ask us questions
# it begins with the resource group
# then the cluster name
# then it appears to create a new key vault
az aks connection create keyvault --new
# Start creating a new keyvault
# Created, the resource id is: /subscriptions/0990639e-03f6-4b45-8fdd-143da61a6e2d/resourceGroups/benpetersen-rg-akv-on-aks/providers/Microsoft.KeyVault/vaults/vault-PXzy
```
I didn't really want a new keyvault, I already made one, but, if this at 
least gets me over the hump, then I am fine with continuing with the path.

Did it work? It prints a lot of noise, but:

```bash
az aks connection list --resource-group benpetersen-rg-akv-on-aks --name akv-on-aks-with-sscsi --output table
# Command group 'aks connection' is in preview and under development. Reference and support levels: https://aka.ms/CLI_refstatus
# ClientType    Name                   ProvisioningState    ResourceGroup              Scope
# ------------  ---------------------  -------------------  -------------------------  -------
# none          akv_on_aks_with_sscsi  Failed               benpetersen-rg-akv-on-aks  default
```
Seems it did not work.

How about using the [Azure Portal Web UI to Create a Connection](https://ms.portal.azure.com/#browse/Microsoft.Network%2Fconnections)?
No, this provides a bunch of options about VNets and things.  Seems incorrect.

Back to [The original doc page](https://learn.microsoft.com/en-us/azure/aks/csi-secrets-store-identity-access?tabs=azure-portal&pivots=access-with-service-connector),
there is a section about `Using the Azure Portal`.  Let's try.

- Navigate to [your AKS cluster in the portal](https://ms.portal.azure.com/#@microsoft.onmicrosoft.com/resource/subscriptions/0990639e-03f6-4b45-8fdd-143da61a6e2d/resourceGroups/benpetersen-rg-akv-on-aks/providers/Microsoft.ContainerService/managedClusters/akv-on-aks-with-sscsi/overview)
- Choose `Settings` -> `Service Connector (Preview)` in the sidebar
- Choose `Create` Connection 
- Follow the instructions and click `Create` once validation completes.

Seems to have worked.  Not sure what was different, I populated with 
what seemed to be the same data.

So lets find out:

```bash
# if you don't provide all these flags it will interactively request the values
# be careful! my aks/akv are flipped around between the success/failed connection name
az aks connection show --connection aks_to_akv_with_sscsi --resource-group benpetersen-rg-akv-on-aks --name akv-on-aks-with-sscsi
```
Results in 

```javascript
{
  "additionalProperties": {},
  "authInfo": {
    "additionalProperties": {},
    "authMode": null,
    // user assigned i guess? means I creaed it?
    "authType": "userAssignedIdentity",
    "clientId": "",
    "deleteOrUpdateBehavior": null,
    "roles": [],
    "subscriptionId": "",
    "userName": null
  },
  // supposed to be "none"
  "clientType": "none",
  "configurationInfo": {
    "action": null,
    "additionalConfigurations": null,
    "additionalConnectionStringProperties": null,
    "additionalProperties": {},
    "configurationStore": null,
    "customizedKeys": {},
    "daprProperties": {
      "additionalProperties": {},
      // i think i did indicate bidirectional in web UI
      "bindingComponentDirection": null,
      "componentType": "",
      "metadata": [],
      "runtimeVersion": null,
      "scopes": [],
      // dunno if this makes sense
      "secretStoreComponent": null,
      "version": ""
    },
    "deleteOrUpdateBehavior": null
  },
  "configurations": [
    {
      "configType": "Default",
      "description": "",
      // no ID?
      "keyVaultReferenceIdentity": null,
      "name": "AZURE_KEYVAULT_TENANTID",
      "value": "72f988bf-86f1-41af-91ab-2d7cd011db47"
    },
    {
      "configType": "Default",
      "description": "",
      // no ID again?
      "keyVaultReferenceIdentity": null,
      "name": "AZURE_KEYVAULT_NAME",
      "value": "akv-on-aks-with-sscsi"
    },
    {
      "configType": "Default",
      "description": "",
      // Hmm.
      "keyVaultReferenceIdentity": null,
      "name": "AZURE_KEYVAULT_CLIENTID",
      "value": "742444a6-0477-4342-998d-d0bc91504106"
    }
  ],
  // uri per usual
  "id": "/subscriptions/0990639e-03f6-4b45-8fdd-143da61a6e2d/resourceGroups/benpetersen-rg-akv-on-aks/providers/Microsoft.ContainerService/managedClusters/akv-on-aks-with-sscsi/providers/Microsoft.ServiceLinker/linkers/aks_to_akv_with_sscsi",
  "kubernetesResourceName": [
    // saw this secret on the cluster
    //  kubectl get secret sc-akstoakvwithsscsi-secret -o yaml
    "sc-akstoakvwithsscsi-secret"
  ],
  // fine to name
  "name": "aks_to_akv_with_sscsi",
  // yay success
  "provisioningState": "Succeeded",
  "publicNetworkSolution": null,
  // this is consistent
  "resourceGroup": "benpetersen-rg-akv-on-aks",
  "scope": "default",
  // not sure about this?
  "secretStore": null,
  // appears to be metadata
  "systemData": {
    "additionalProperties": {},
    "createdAt": "2024-11-06T18:38:07.509293+00:00",
    "createdBy": "benpetersen@microsoft.com",
    "createdByType": "User",
    "lastModifiedAt": "2024-11-06T18:38:07.509293+00:00",
    "lastModifiedBy": "benpetersen@microsoft.com",
    "lastModifiedByType": "User"
  },
  // targets a keyvault
  "targetService": {
    "additionalProperties": {},
    // id of a keyvault
    "id": "/subscriptions/0990639e-03f6-4b45-8fdd-143da61a6e2d/resourceGroups/benpetersen-rg-akv-on-aks/providers/Microsoft.KeyVault/vaults/akv-on-aks-with-sscsi",
    // same RG
    "resourceGroup": "benpetersen-rg-akv-on-aks",
    "resourceProperties": {
      "additionalProperties": {},
      // oh good, this seems what we want
      "connectAsKubernetesCsiDriver": true,
      // yes, again a key vault
      "type": "KeyVault"
    },
    "type": "AzureResource"
  },
  // imagine there are things behind the scenes here
  "type": "microsoft.servicelinker/linkers",
  "vNetSolution": null
}
```

So lets see:

```bash
# yup this secret now exists 
kubectl get secret sc-akvonakswithsscsi-secret -o yaml
```
output:
```yaml
apiVersion: v1
kind: Secret
type: Opaque
metadata:
  creationTimestamp: "2024-11-06T18:42:46Z"
  name: sc-akstoakvwithsscsi-secret
  namespace: default
  resourceVersion: "231478"
  uid: e0643727-7c16-4d83-81db-aa5b1aa136ec
data:
  # no credential here, but important bits
  AZURE_KEYVAULT_CLIENTID: NzQyNDQ0YTYtMDQ3Ny00MzQyLTk5OGQtZDBiYzkxNTA0MTA2 # | base64 --decode 742444a6-0477-4342-998d-d0bc91504106
  AZURE_KEYVAULT_NAME: YWt2LW9uLWFrcy13aXRoLXNzY3Np                         # | base64 --decode akv-on-aks-with-sscsi
  AZURE_KEYVAULT_TENANTID: NzJmOTg4YmYtODZmMS00MWFmLTkxYWItMmQ3Y2QwMTFkYjQ3 # | base64 --decode 72f988bf-86f1-41af-91ab-2d7cd011db47
```

```bash
# but, this also seems interesting!
kubectl get secret serviceconnector-secret -n sc-system -o yaml
```
output
```yaml
apiVersion: v1
kind: Secret
type: Opaque
metadata:
  annotations:
    meta.helm.sh/release-name: sc-extension
    meta.helm.sh/release-namespace: sc-system
  creationTimestamp: "2024-11-06T18:42:43Z"
  labels:
    app.kubernetes.io/managed-by: Helm
  name: serviceconnector-secret
  namespace: sc-system
  resourceVersion: "231449"
  uid: dbe0f100-2117-428d-9b58-d34a3cd444c6
data:
  LOG_CORRELATIONID: NWJhMWUxYTQtN2IzMi00YTViLWJjYTktYzJhNGIwMzI0YTg0
  LOG_IKEY: ZWZkMWFiNGUtMDg3MS00YmMzLTg0YjItNjEwZjc1ZTI0MzZk
  SC_CONNECTION: YWtzX3RvX2Frdl93aXRoX3NzY3Np
  SC_OPERATION: Q3JlYXRl
  SC_RESOURCES: eyJuYW1lc3BhY2UiOnsibmFtZSI6ImRlZmF1bHQifSwic2VjcmV0Ijp7Im5hbWUiOiJzYy1ha3N0b2FrdndpdGhzc2NzaS1zZWNyZXQiLCJuYW1lc3BhY2UiOiJkZWZhdWx0IiwiZGF0YSI6eyJBWlVSRV9LRVlWQVVMVF9URU5BTlRJRCI6IjcyZjk4OGJmLTg2ZjEtNDFhZi05MWFiLTJkN2NkMDExZGI0NyIsIkFaVVJFX0tFWVZBVUxUX05BTUUiOiJha3Ytb24tYWtzLXdpdGgtc3Njc2kiLCJBWlVSRV9LRVlWQVVMVF9DTElFTlRJRCI6Ijc0MjQ0NGE2LTA0NzctNDM0Mi05OThkLWQwYmM5MTUwNDEwNiJ9fSwic2VjcmV0UHJvdmlkZXJDbGFzcyI6bnVsbCwicG9kIjpudWxsLCJkYXByQ29tcG9uZW50IjpudWxsLCJzZXJ2aWNlQWNjb3VudCI6bnVsbH0=
  # | base64 --decode
    # {
    #   "namespace": {
    #     "name": "default"
    #   },
    #   "secret": {
    #     "name": "sc-akstoakvwithsscsi-secret",
    #     "namespace": "default",
    #     "data": {
    #       "AZURE_KEYVAULT_TENANTID": "72f988bf-86f1-41af-91ab-2d7cd011db47",
    #       "AZURE_KEYVAULT_NAME": "akv-on-aks-with-sscsi",
    #       "AZURE_KEYVAULT_CLIENTID": "742444a6-0477-4342-998d-d0bc91504106"
    #     }
    #   },
    #   "secretProviderClass": null,
    #   "pod": null,
    #   "daprComponent": null,
    #   "serviceAccount": null
    # }
```

```bash
# this also seems interesting
# MSI = Managed Service Identity, Perhaps?
kubectl get secret extensions-aad-msi-token -n kube-system -o yaml
```
output
```yaml
# at some point we do need a token, this might be it?
apiVersion: v1
kind: Secret
type: Opaque
metadata:
  creationTimestamp: "2024-11-06T18:40:48Z"
  name: extensions-aad-msi-token
  namespace: kube-system
  resourceVersion: "235860"
  uid: 68277a11-7257-4ccf-bc99-0478c0f478a6
data:
  token: eyJh...iJ9
```

## Test the Connection 

The suggested test is to utilize a github repo.

```bash
# I'll clone this in ~/github.com/Azure-Samples
git clone https://github.com/Azure-Samples/serviceconnector-aks-samples.git
# seems to be the only dir in the repo
cd serviceconnector-aks-samples/azure-keyvault-csi-provider
```

Per this repo, I cloned and filled out a `SecretProviderClass` yaml:

```yaml
# Seems we need to utilize:
#   az aks connection show --connection aks_to_akv_with_sscsi --resource-group benpetersen-rg-akv-on-aks --name akv-on-aks-with-sscsi
# to get the data we need to fill this out.
#
# look!  We know this API Version
apiVersion: secrets-store.csi.x-k8s.io/v1
# and this resource!
kind: SecretProviderClass
metadata:
  # time to make these.
  name: sc-demo-keyvault-csi
spec:
  provider: azure
  parameters:
    usePodIdentity: "false"
    useVMManagedIdentity: "true"                                        # Set to true for using managed identity
    # <AZURE_KEYVAULT_CLIENTID>
    # per the above command I found this:
    #  "keyVaultReferenceIdentity": null,
    #  "name": "AZURE_KEYVAULT_CLIENTID",
    #  "value": "742444a6-0477-4342-998d-d0bc91504106"
    userAssignedIdentityID: 742444a6-0477-4342-998d-d0bc91504106        # Set the clientID of the user-assigned managed identity to use
    # <AZURE_KEYVAULT_NAME>
    keyvaultName: akv-on-aks-with-sscsi                                 # Set to the name of your key vault
    # <KEYVAULT_SECRET_NAME>
    # pick a secret, dunno if we can get multiple or not.
    # will start with one, though array seems to indicate
    # we could get a few
    objects:  |
      array:
        - |
          objectName: ninnyhammers
          objectType: secret
    # <AZURE_KEYVAULT_TENANTID>
    # per the above command I found this:
    #  "keyVaultReferenceIdentity": null,
    #  "name": "AZURE_KEYVAULT_TENANTID",
    #  "value": "72f988bf-86f1-41af-91ab-2d7cd011db47"
    tenantId: 72f988bf-86f1-41af-91ab-2d7cd011db47                      # The tenant ID of the key vault
``` 
and
```bash
kubectl apply -f secret_provider_class.my.yaml
```

And it inidicates to deploy this POD manifest:

```yaml
# This is a sample pod definition for using SecretProviderClass and the user-assigned identity to access your key vault

kind: Pod
apiVersion: v1
metadata:
  name: sc-demo-keyvault-csi
spec:
  containers:
    # just a busybox
    - name: busybox
      image: registry.k8s.io/e2e-test-images/busybox:1.29-4
      # sleeps a long time, no log or anything
      command:
        - "/bin/sleep"
        - "10000"
      # volumeMount is what we expect
      volumeMounts:
      - name: secrets-store01-inline
        mountPath: "/mnt/secrets-store"
        readOnly: true
      resources:
        requests:
          cpu: 100m
          memory: 128Mi
        limits:
          cpu: 250m
          memory: 256Mi
  # amd a volume to match
  volumes:
    - name: secrets-store01-inline
      csi:
        # CSI Driver
        driver: secrets-store.csi.k8s.io
        readOnly: true
        volumeAttributes:
          # adn this should match names or it won't work
          secretProviderClass: "sc-demo-keyvault-csi"
```

so

```bash
kubectl apply -f pod.yaml
# after a moment of creation it seems healthy
watch "kubectl get pods"
# only way to know if it works is to exec in and look for a secret
kubectl exec -it sc-demo-keyvault-csi -- /bin/sh
#  cd /mnt/secrets-store
#  tail ninnyhammers
```

Well surprise, it works.  It did indeed:
- Utilize the CSI Driver
- Utilize the AKV Provider
- Get the `ninnyhammer` secret from the specified Key Vault
- Mount the secret into the pod

## Can we also install Secrets Sync Controller?

Would it use the same credentials?
In theory, if it uses the same provider (already setup above)
to get the secrets, then it ought be using the same credentials.

Given:
- The Raw Git url to install SCI Driver Helm chart:
  - https://raw.githubusercontent.com/kubernetes-sigs/secrets-store-csi-driver/master/charts
- And the URL to SS SC Helm Chart:
  - https://github.com/kubernetes-sigs/secrets-store-sync-controller/tree/main/charts/secrets-store-sync-controller
- We can derive a SS SC Raw Git URL:
  - https://raw.githubusercontent.com/kubernetes-sigs/secrets-store-sync-controller/master/charts

But then, even more simply:

```bash
helm repo list
# NAME                    	URL
# secrets-store-csi-driver	https://kubernetes-sigs.github.io/secrets-store-csi-driver/charts
# kubernetes-dashboard    	https://kubernetes.github.io/dashboard/
#
# lets just swap the URL a bit here and add it
helm repo add secrets-store-sync-controller https://kubernetes-sigs.github.io/secrets-store-sync-controller/charts
# "secrets-store-sync-controller" has been added to your repositories
#
helm repo list
# NAME                         	URL
# secrets-store-csi-driver     	https://kubernetes-sigs.github.io/secrets-store-csi-driver/charts
# kubernetes-dashboard         	https://kubernetes.github.io/dashboard/
# secrets-store-sync-controller	https://kubernetes-sigs.github.io/secrets-store-sync-controller/charts
```

Seems fine. So how about a similar install?

```bash
# used this previously on a kind cluster
# but on Azure we just passed a flag on cluster creation and 
# this was done for us.  So, prob need to compare quick?
helm install csi secrets-store-csi-driver/secrets-store-csi-driver -n csi --set grpcSupportedProviders="azure;gcp" # deploy the secrets store csi driver
```

So:

```bash
# find the Driver
kubectl get daemonset -A
# NAMESPACE     NAME                                     
# kube-system   aks-secrets-store-csi-driver
#
# How about its config?
kubectl get cm -n kube-system
# nothing
```

So lets take a look at the driver:

```yaml
# rant his
# kubectl get daemonset  aks-secrets-store-csi-driver -n kube-system -o yaml
apiVersion: apps/v1
# deleting some noise to make this more reasonable
kind: DaemonSet
metadata:
  annotations:
    deprecated.daemonset.template.generation: "1"    
  creationTimestamp: "2024-11-06T02:43:25Z"
  labels:
    addonmanager.kubernetes.io/mode: Reconcile
    app: secrets-store-csi-driver
    # managed.
    kubernetes.azure.com/managedby: aks
  name: aks-secrets-store-csi-driver
  namespace: kube-system  
spec:
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: secrets-store-csi-driver
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: secrets-store-csi-driver
        # its managed.
        kubernetes.azure.com/managedby: aks
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: type
                operator: NotIn
                values:
                - virtual-kubelet
              - key: kubernetes.azure.com/cluster
                operator: Exists
              - key: kubernetes.io/os
                operator: In
                values:
                - linux                
      # I believe there are 3 containers that install.
      # The driver, the registrar, and something else.
      containers:
      # Container 2
      - args:
        - --v=2
        - --csi-address=/csi/csi.sock
        - --kubelet-registration-path=/var/lib/kubelet/plugins/csi-secrets-store/csi.sock
        image: mcr.microsoft.com/oss/kubernetes-csi/csi-node-driver-registrar:v2.11.1
        imagePullPolicy: IfNotPresent
        livenessProbe:
          exec:
            # csi.sock again, in the livenessProbe. curious
            command:
            - /csi-node-driver-registrar
            - --kubelet-registration-path=/var/lib/kubelet/plugins/csi-secrets-store/csi.sock
            - --mode=kubelet-registration-probe
          failureThreshold: 3
          initialDelaySeconds: 30
          periodSeconds: 10
          successThreshold: 1
          timeoutSeconds: 15
        # this is the REGISTRAR, one of the 3 things installed
        name: node-driver-registrar
        resources:
          limits:
            memory: 100Mi
          requests:
            cpu: 10m
            memory: 20Mi
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        # mounts, to communicate with providers I imagine
        volumeMounts:
        - mountPath: /csi
          name: plugin-dir
        - mountPath: /registration
          name: registration-dir
      # Container 2
      - args:
        - -v=2
        - --log-format-json=false
        - --endpoint=$(CSI_ENDPOINT)
        - --nodeid=$(KUBE_NODE_NAME)
        - --provider-volume=/var/run/secrets-store-csi-providers
        - --additional-provider-volume-paths=/etc/kubernetes/secrets-store-csi-providers
        - --metrics-addr=:8095
        - --max-call-recv-msg-size=4194304
        env:
        # this is our csi.sock endpoint, to communicate with providers 
        - name: CSI_ENDPOINT
          value: unix:///csi/csi.sock
        - name: KUBE_NODE_NAME
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: spec.nodeName
        # this looks like our image!
        image: mcr.microsoft.com/oss/kubernetes-csi/secrets-store/driver:v1.4.5
        imagePullPolicy: IfNotPresent
        livenessProbe:
          failureThreshold: 5
          httpGet:
            path: /healthz
            port: healthz
            scheme: HTTP
          initialDelaySeconds: 30
          periodSeconds: 15
          successThreshold: 1
          timeoutSeconds: 10
        # bet this guy is the right one we care about 
        name: secrets-store
        ports:
        - containerPort: 9808
          name: healthz
          protocol: TCP
        - containerPort: 8095
          name: metrics
          protocol: TCP
        resources:
          limits:
            cpu: 200m
            memory: 500Mi
          requests:
            cpu: 50m
            memory: 100Mi
        securityContext:
          privileged: true
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        # plenty of volume mounts
        volumeMounts:
        - mountPath: /csi
          name: plugin-dir
        - mountPath: /var/lib/kubelet/pods
          mountPropagation: Bidirectional
          name: mountpoint-dir
        - mountPath: /var/run/secrets-store-csi-providers
          name: providers-dir
        - mountPath: /etc/kubernetes/secrets-store-csi-providers
          name: providers-dir-0
      # container 3
      - args:
        - --csi-address=/csi/csi.sock
        - --probe-timeout=3s
        - --http-endpoint=0.0.0.0:9808
        - -v=2
        image: mcr.microsoft.com/oss/kubernetes-csi/livenessprobe:v2.13.1
        imagePullPolicy: IfNotPresent
        # just a liveness probe
        name: liveness-probe
        resources:
          limits:
            cpu: 100m
            memory: 100Mi
          requests:
            cpu: 10m
            memory: 20Mi
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
        volumeMounts:
        - mountPath: /csi
          name: plugin-dir
      dnsPolicy: Default
      priorityClassName: system-node-critical
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      serviceAccount: aks-secrets-store-csi-driver
      serviceAccountName: aks-secrets-store-csi-driver
      terminationGracePeriodSeconds: 30
      tolerations:
      - key: CriticalAddonsOnly
        operator: Exists
      - effect: NoExecute
        operator: Exists
      - effect: NoSchedule
        operator: Exists
      # volumes, to match the volumeMounts
      volumes:
      - hostPath:
          path: /var/lib/kubelet/pods
          type: DirectoryOrCreate
        name: mountpoint-dir
      - hostPath:
          path: /var/lib/kubelet/plugins_registry/
          type: Directory
        name: registration-dir
      - hostPath:
          path: /var/lib/kubelet/plugins/csi-secrets-store/
          type: DirectoryOrCreate
        name: plugin-dir
      - hostPath:
          path: /var/run/secrets-store-csi-providers
          type: DirectoryOrCreate
        name: providers-dir
      - hostPath:
          path: /etc/kubernetes/secrets-store-csi-providers
          type: DirectoryOrCreate
        name: providers-dir-0
  updateStrategy:
    rollingUpdate:
      maxSurge: 0
      maxUnavailable: 1
    type: RollingUpdate
status:
  currentNumberScheduled: 3
  desiredNumberScheduled: 3
  numberAvailable: 3
  numberMisscheduled: 0
  numberReady: 3
  observedGeneration: 1
  updatedNumberScheduled: 3
```

Not sure if any of that matters, so I guess we can try to just install the SSSC Controller:

```bash
# helm install <name> <repo>/<chart> --namespace <ns>
helm install secrets-sync-controller secrets-store-sync-controller/secrets-store-sync-controller --namespace kube-system
# NAME: secrets-sync-controller
# LAST DEPLOYED: Wed Nov  6 14:46:11 2024
# NAMESPACE: kube-system
# STATUS: deployed
# REVISION: 1
# TEST SUITE: None
#
kubectl get deployment -n kube-system && kubectl get daemonset -n kube-system
# our deployment is happy
#   NAME                                    READY   UP-TO-DATE   AVAILABLE   AGE
#   secrets-store-sync-controller-manager   1/1     1            1           2m5s
# and our csi driver is in the same ns
#   NAME                                       DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR              AGE
#   aks-secrets-store-csi-driver               3         3         3       3            3           <none>                     17h
#   aks-secrets-store-csi-driver-windows       0         0         0       0            0           <none>                     17h
#   aks-secrets-store-provider-azure           3         3         3       3            3           kubernetes.io/os=linux     17h
#   aks-secrets-store-provider-azure-windows   0         0         0       0            0           kubernetes.io/os=windows   17h
kubectl get secretproviderclass -n default
# our SecretProviderClass is in default
#   NAMESPACE   NAME                   AGE
#   default     sc-demo-keyvault-csi   27m
kubectl get pods -n default
# as is our pod utilizing it
#   NAME                   READY   STATUS    RESTARTS   AGE
#   sc-demo-keyvault-csi   1/1     Running   0          25m
```

Seems fine?  Now to sync a secret from AKV to a Kubernetes Secret, probably in default namespace again.

```yaml
# NOTE:
# THIS IS NOT CSI DRIVER!
# THIS IS SECRET SYNC CONTROLLER!
# I AM JUST SEEING IF I CAN INSTALL AND REUSE CSI DRIVER BITS!
apiVersion: secret-sync.x-k8s.io/v1alpha1
kind: SecretSync
metadata:
 name: sse2esecret  # this is the name of the secret that will be created
spec:
 # here in the tests:
 #  https://github.com/kubernetes-sigs/secrets-store-sync-controller/blob/main/test/bats/tests/e2e_provider/service_account_token_secretsync.yaml#L6
 # it appears that 'default' is used. 
 #  serviceAccountName: ${SERVICE_ACCOUNT_NAME}
 serviceAccountName: default
 # referencing what we already have
 secretProviderClassName: sc-demo-keyvault-csi
 secretObject:
  type: Opaque
  data:
  - sourcePath: ninnyhammers # name of the object in the SecretProviderClass
    # well, in this case, the ninnyhammer secret is just text.
    # there is no key inside it. 
    # so, I am going to leave blank to see what happens?
    targetKey: value  # name of the key in the Kubernetes secret
```
Well, it wont `kubectl apply` without `targetKey`, so lets see:

```bash
az keyvault secret list --vault-name akv-on-aks-with-sscsi --output table
az keyvault secret show --vault-name akv-on-aks-with-sscsi --name ninnyhammers
# {
#   "attributes": {
#     "created": "2024-11-06T15:26:49+00:00",
#     "updated": "2024-11-06T15:26:49+00:00"
#   },
#   "contentType": null,
#   "id": "https://akv-on-aks-with-sscsi.vault.azure.net/secrets/ninnyhammers/02342d20444b46f59b0d9f249bd8de54",
#   "name": "ninnyhammers",
#   "value": "Youre nowt but a ninnyhammer, Sam Gamgee"
# }
```
Looks like `value` is our choice!
Which is fine, but when created, it didn't work.  Secrets Sync Controller `logs` indicate it 
couldn't find the `Azure` provider.

```bash
kubectl logs deploy/secrets-store-sync-controller-manager -n kube-system
# 2024-11-06T20:19:41Z	ERROR	Reconciler error	{"controller": "secretsync", "controllerGroup": "secret-sync.x-k8s.io", "controllerKind": "SecretSync", "SecretSync": {"name":"sse2esecret","namespace":"default"}, "namespace": "default", "name": "sse2esecret", "reconcileID": "ad268f7e-4049-432e-b1bf-04a29de7b1f5", "error": "provider not found: provider \"azure\""}
```

```bash
# so lets look at the installed charts on cluster
helm list --all-namespaces
# delete the current SSSC
helm uninstall secrets-sync-controller -n kube-system
# install it again, borrowing the supported providers params from 
# when we installed the CSI Driver
helm install secrets-sync-controller secrets-store-sync-controller/secrets-store-sync-controller --namespace kube-system --set grpcSupportedProviders="azure;gcp"
# NAME: secrets-sync-controller
# LAST DEPLOYED: Wed Nov  6 15:18:00 2024
# NAMESPACE: kube-system
# STATUS: deployed
# REVISION: 1
# TEST SUITE: None
```

So, I still see a Reconciler error:
`
```bash
kubectl logs deploy/secrets-store-sync-controller-manager -n kube-system
# 2024-11-06T20:40:10Z	ERROR	Reconciler error	{"controller": "secretsync", "controllerGroup": "secret-sync.x-k8s.io", "controllerKind": "SecretSync", "SecretSync": {"name":"sse2esecret","namespace":"default"}, "namespace": "default", "name": "sse2esecret", "reconcileID": "982cb35a-f17e-49b4-b6ab-006555a16ddc", "error": "provider not found: provider \"azure\""}
echo '<above.json>' | jq
# {
#   "controller": "secretsync",
#   "controllerGroup": "secret-sync.x-k8s.io",
#   "controllerKind": "SecretSync",
#   "SecretSync": {
#     "name": "sse2esecret",
#     "namespace": "default"
#   },
#   "namespace": "default",
#   "name": "sse2esecret",
#   "reconcileID": "982cb35a-f17e-49b4-b6ab-006555a16ddc",
#   "error": "provider not found: provider \"azure\""
# }
```

So, lets just see if the namespace matters:

```bash
# verify no secret exists
kubectl get secrets -n kube-system | grep sse2esecret
# create the SecretSync in kube-system 
kubectl create -f secret_sync.yaml -n kube-system
# secretsync.secret-sync.x-k8s.io/sse2esecret create
kubectl get secretsync -A
# NAMESPACE     NAME          AGE
# default       sse2esecret   42m
# kube-system   sse2esecret   13s
kubectl get secret -n kube-system | grep sse2esecret
# nope.
kubectl logs deploy/secrets-store-sync-controller-manager -n kube-system
# 2024-11-06T20:50:54Z	ERROR	Reconciler error	{"controller": "secretsync", "controllerGroup": "secret-sync.x-k8s.io", "controllerKind": "SecretSync", "SecretSync": {"name":"sse2esecret","namespace":"kube-system"}, "namespace": "kube-system", "name": "sse2esecret", "reconcileID": "c8e3ec73-c247-4c03-b10e-5bb89eb13185", "error": "SecretProviderClass.secrets-store.csi.x-k8s.io \"sc-demo-keyvault-csi\" not found"}
```
Ok, so we need the `secretproviderclass` also in the same `kube-system` namespace, to 
fix this error:

```bash
kubectl apply -f secret_provider_class.my.yaml -n kube-system
# secretproviderclass.secrets-store.csi.x-k8s.io/sc-demo-keyvault-csi created
kubectl get secret -n kube-system
# nope
kubectl logs deploy/secrets-store-sync-controller-manager -n kube-system
# 2024-11-06T20:52:57Z	ERROR	Reconciler error	{"controller": "secretsync", "controllerGroup": "secret-sync.x-k8s.io", "controllerKind": "SecretSync", "SecretSync": {"name":"sse2esecret","namespace":"kube-system"}, "namespace": "kube-system", "name": "sse2esecret", "reconcileID": "04a3ff4f-f932-4346-8d86-117afe8308b0", "error": "provider not found: provider \"azure\""}
```

So that `provider not found: provider "azure"` is at least consistent.
But the thing is, I do have these:

```bash
kubectl get secretproviderclass -A
# NAMESPACE     NAME                   AGE
# default       sc-demo-keyvault-csi   92m
# kube-system   sc-demo-keyvault-csi   110s
```
So what is wrong?
Well, going [back to the README](https://github.com/kubernetes-sigs/secrets-store-sync-controller/tree/9f4597d20eb864cfa338f693dc90f3bdf10e5128)
there is a section called `2. Configure the provider container`.

Then recall:
- Providers communicate over unix sockets via RPC calls
- So, a pod must have containers that share volumes, volumeMounts, and sockets, probably
- containers in a pod are co-located to make this stuff work
- So, we probably need to `get our provider pod` from `kube-system`,
  of which there are three `aks-secrets-store-provider-azure-<id>`, and 
  copy it into our helm chart and run an install again.
  This likely requires a `values.yaml` file.

So, lets make a `values.yaml` file:

```yaml
# Default values for secrets-store-sync-controller.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.
controllerName: secrets-store-sync-controller-manager

tokenRequestAudience: 
  - audience:  # e.g. api://TokenAudienceExample

logVerbosity: 5 

validatingAdmissionPolicies:
  applyPolicies: true
  kubernetesReleaseVersion: "1.28.0"
  allowedSecretTypes:
    - "Opaque"
    - "kubernetes.io/basic-auth"
    - "bootstrap.kubernetes.io/token"
    - "kubernetes.io/dockerconfigjson"
    - "kubernetes.io/dockercfg"
    - "kubernetes.io/ssh-auth"
    - "kubernetes.io/tls"

  deniedSecretTypes:
    - "kubernetes.io/service-account-token"

image:
  repository: registry.k8s.io/secrets-store-sync/controller # e.g. my-registry.example.com/my-repo
  pullPolicy: IfNotPresent
  tag: v0.0.1

securityContext:
  # Default values, can be overridden or extended
  allowPrivilegeEscalation: false
  capabilities:
    drop:
      - ALL

resources:
  limits:
    cpu: 500m
    memory: 128Mi
  requests:
    cpu: 10m
    memory: 64Mi

podAnnotations:
  kubectl.kubernetes.io/default-container: manager

podLabels: 
  control-plane: controller-manager
  secrets-store.io/system: "true"
  app: secrets-store-sync-controller

nodeSelector:

tolerations: 
- operator: Exists

affinity:

metricsPort: 8085

providerContainer:
# this is the e2e-provider, keeping as an example.
#  - name: provider-e2e-installer
#    image: aramase/e2e-provider:v0.0.1
#    imagePullPolicy: IfNotPresent
#    args:
#      - --endpoint=unix:///provider/e2e-provider.sock
#    resources:
#      requests:
#        cpu: 50m
#        memory: 100Mi
#      limits:
#        cpu: 50m
#        memory: 100Mi
#    securityContext:
#      allowPrivilegeEscalation: false
#      readOnlyRootFilesystem: true
#      runAsUser: 0
#      capabilities:
#        drop:
#        - ALL
#    volumeMounts:
#      - mountPath: "/provider"
#        name: providervol
# the endpoint we expect, I think.
# a unix sock, with a name of azure.sock
- args:
  - --endpoint=unix:///provider/azure.sock
  - --log-format-json=false
  - -v=2
  - --construct-pem-chain=true
  - '--custom-user-agent=''kubernetes.azure.com/managedby: aks'''
  - --healthz-port=8989
  - --healthz-path=/healthz
  - --healthz-timeout=5s
  - --write-cert-and-key-in-separate-files=true
  env:
  # some fun env vars.
  # since we provided the flag to install this addon automatically,
  # we didn't have to set all this good stuff
  - name: KUBERNETES_SERVICE_HOST
    value: akv-on-aks-benpetersen-rg-a-099063-9rfmdvqg.hcp.eastus2.azmk8s.io
  - name: KUBERNETES_PORT
    value: tcp://akv-on-aks-benpetersen-rg-a-099063-9rfmdvqg.hcp.eastus2.azmk8s.io:443
  - name: KUBERNETES_PORT_443_TCP
    value: tcp://akv-on-aks-benpetersen-rg-a-099063-9rfmdvqg.hcp.eastus2.azmk8s.io:443
  - name: KUBERNETES_PORT_443_TCP_ADDR
    value: akv-on-aks-benpetersen-rg-a-099063-9rfmdvqg.hcp.eastus2.azmk8s.io
  image: mcr.microsoft.com/oss/azure/secrets-store/provider-azure:v1.5.3
  imagePullPolicy: IfNotPresent
  livenessProbe:
    failureThreshold: 3
    httpGet:
      path: /healthz
      port: 8989
      scheme: HTTP
    initialDelaySeconds: 5
    periodSeconds: 30
    successThreshold: 1
    timeoutSeconds: 10
  name: provider-azure-installer
  ports:
  - containerPort: 8898
    hostPort: 8898
    name: metrics
    protocol: TCP
  resources:
    limits:
      cpu: 100m
      memory: 100Mi
    requests:
      cpu: 50m
      memory: 100Mi
  securityContext:
    capabilities:
      drop:
      - ALL
    readOnlyRootFilesystem: true
    runAsUser: 0
  terminationMessagePath: /dev/termination-log
  terminationMessagePolicy: File
  # this might be a snag, looking at the deployment yaml on cluster, 
  # but we will see...
  volumeMounts:
#!   - mountPath: /provider
#!     name: provider-vol
  - mountPath: /var/run/secrets-store-sync-providers
    name: providervol
# lets presume the serviceaccount token will be mounted automatically?
#!   - mountPath: /var/run/secrets/kubernetes.io/serviceaccount
#!     name: kube-api-access-4jm57
#!     readOnly: true
```
And see how it works if we apply it.
Note that this yaml file, and other outputs I have referenced, 
[exist locally in my cloned down Azure-Samples repo](~/github.com/Azure-Samples/serviceconnector-aks-samples/azure-keyvault-csi-provider/helm/secret-sync-controller-values.for.azure.yaml)

```bash
helm install -f helm/secret-sync-controller-values.for.azure.yaml secrets-sync-controller secrets-store-sync-controller/secrets-store-sync-controller
# NAME: secrets-sync-controller
# LAST DEPLOYED: Wed Nov  6 16:20:44 2024
# NAMESPACE: default
# STATUS: deployed
# REVISION: 1
# TEST SUITE: None
```
So it installed, but its not happy:

```bash
watch "kubectl get deployment secrets-store-sync-controller-manager -o wide && echo '' && kubectl get pods && echo '' && kubectl get events"
# # deployment
# # NAME                                    READY   UP-TO-DATE   AVAILABLE   AGE    CONTAINERS         IMAGES
# secrets-store-sync-controller-manager     0/1  4m5s   provider-azure-installer,manager  
# # pods
# NAME                                                     READY   STATUS    RESTARTS   AGE
# sc-demo-keyvault-csi                                     1/1     Running   0          120m
# secrets-store-sync-controller-manager-665545fb88-jj8d5   0/2     Pending   0          4m5s
# # events
# LAST SEEN   TYPE      REASON              OBJECT                                                        MESSAGE
# 4m4s        Warning   FailedScheduling    pod/secrets-store-sync-controller-manager-665545fb88-jj8d5    0/3 nodes are available: 3 node(s) didn't have free ports for the requested pod ports. preempti
# on: 0/3 nodes are available: 3 No preemption victims found for incoming pod.
# 4m5s        Normal    SuccessfulCreate    replicaset/secrets-store-sync-controller-manager-665545fb88   Created pod: secrets-store-sync-controller-manager-665545fb88-jj8d5
# 4m5s        Normal    ScalingReplicaSet   deployment/secrets-store-sync-controller-manager              Scaled up replica set secrets-store-sync-controller-manager-665545fb88 to 1
```
Noticing `3 nodes didn't have free ports for the requesting pod ports`.  That seems sometihng we can fix.

```bash
# does give a little more output in the state of things.
describe secretsync -A
# ...
# Status:
#   Conditions:
#     Message: Secret update failed because the controller could not retrieve the Secret 
#     Provider Class or the SPC is misconfigured. Check the logs or the events for more information.
#     Reason: ControllerSPCError
```
Which is consistent with 

```bash
kubectl get events
# LAST SEEN   TYPE      REASON             OBJECT                    MESSAGE
# 14m         Normal    Pulled             pod/sc-demo-keyvault-csi  Container image "registry.k8s.io/e2e-test-images/busybox:1.29-4" already present on machine
# 19m         Warning   FailedScheduling   pod/secrets-store-sync-controller-manager-665545fb88-jj8d5   0/3 nodes are available: 3 node(s) didn't have free ports for the requested pod ports. preemption: 0/3 nodes are available: 3 No preemption victims found for incoming pod.
```
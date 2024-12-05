# AZ Keyvault Create 

```bash
az keyvault create \
    --name akv-on-aks-with-sscsi \
    --resource-group benpetersen-rg-akv-on-aks \
    --location eastus2 \
    --enable-rbac-authorization
```

and the output 

```javascript
{
  // id here is in URL format, like the aks cluster
  // however, there is not also an .identity{} stanza.
  // i imagine because the vault is the anchor, other 
  // identities likely depend on keys and secrets in the 
  // vaults
  "id": "/subscriptions/0990639e-03f6-4b45-8fdd-143da61a6e2d/resourceGroups/benpetersen-rg-akv-on-aks/providers/Microsoft.KeyVault/vaults/akv-on-aks-with-sscsi",
  // location as usual
  "location": "eastus2",
  // name 
  "name": "akv-on-aks-with-sscsi",
  "properties": {
    "accessPolicies": [],
    "createMode": null,
    "enablePurgeProtection": null,
    "enableRbacAuthorization": true,
    "enableSoftDelete": true,
    "enabledForDeployment": false,
    "enabledForDiskEncryption": null,
    "enabledForTemplateDeployment": null,
    "hsmPoolResourceId": null,
    "networkAcls": null,
    "privateEndpointConnections": null,
    "provisioningState": "Succeeded",
    // public access eh?
    "publicNetworkAccess": "Enabled",
    "sku": {
      "family": "A",
      "name": "standard"
    },
    // 90 days to change mind
    "softDeleteRetentionInDays": 90,
    // intersting that this is the uri, my name isn't injected as well
    "tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47",
    "vaultUri": "https://akv-on-aks-with-sscsi.vault.azure.net/"
  },
  "resourceGroup": "benpetersen-rg-akv-on-aks",
  "systemData": {
    "createdAt": "2024-11-06T03:15:53.865000+00:00",
    "createdBy": "benpetersen@microsoft.com",
    "createdByType": "User",
    "lastModifiedAt": "2024-11-06T03:15:53.865000+00:00",
    "lastModifiedBy": "benpetersen@microsoft.com",
    "lastModifiedByType": "User"
  },
  "tags": {},
  "type": "Microsoft.KeyVault/vaults"
}
```
```bash
az role definition list --name "Key Vault Administrator"
```

```json 
[
  {
    "assignableScopes": [
      "/"
    ],
    "createdBy": null,
    "createdOn": "2020-05-19T17:52:46.234923+00:00",
    // human readable
    "description": "Perform all data plane operations on a key vault and all objects in it, including certificates, keys, and secrets. Cannot manage key vault resources or manage role assignments. Only works for key vaults that use the 'Azure role-based access control' permission model.",
    // typical id in form of url
    "id": "/subscriptions/0990639e-03f6-4b45-8fdd-143da61a6e2d/providers/Microsoft.Authorization/roleDefinitions/00482a5a-887f-4fb3-b363-3b7fe8e74483",
    // name looks like a UID though
    "name": "00482a5a-887f-4fb3-b363-3b7fe8e74483",
    // permissions.actions is what it can actually do
    "permissions": [
      {
        "actions": [
          "Microsoft.Authorization/*/read",
          "Microsoft.Insights/alertRules/*",
          "Microsoft.Resources/deployments/*",
          "Microsoft.Resources/subscriptions/resourceGroups/read",
          "Microsoft.Support/*",
          // .KeyVault is most interesting 
          "Microsoft.KeyVault/checkNameAvailability/read",
          "Microsoft.KeyVault/deletedVaults/read",
          "Microsoft.KeyVault/locations/*/read",
          "Microsoft.KeyVault/vaults/*/read",
          "Microsoft.KeyVault/operations/read"
        ],
        "condition": null,
        "conditionVersion": null,
        "dataActions": [
          "Microsoft.KeyVault/vaults/*"
        ],
        "notActions": [],
        "notDataActions": []
      }
    ],
    "roleName": "Key Vault Administrator",
    "roleType": "BuiltInRole",
    "type": "Microsoft.Authorization/roleDefinitions",
    "updatedBy": null,
    "updatedOn": "2021-11-11T20:14:30.254275+00:00"
  }
]
```
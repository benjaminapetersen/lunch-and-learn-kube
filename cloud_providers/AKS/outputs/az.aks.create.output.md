```bash
# to create 
az aks create --name akv-on-aks-with-sscsi --resource-group benpetersen-rg-akv-on-aks --enable-addons azure-keyvault-secrets-provider --enable-oidc-issuer --enable-workload-identity --generate-ssh-keys
# to show what has already been created 
az aks show --name akv-on-aks-with-sscsi --resource-group benpetersen-rg-akv-on-aks
```
output 

```javascript
{
  "aadProfile": null,

  "addonProfiles": {    
    // config for the csi driver
    "azureKeyvaultSecretsProvider": {
      "config": {
        "enableSecretRotation": "false",
        "rotationPollInterval": "2m"
      },
      "enabled": true,
      "identity": {
        // this client id is important, it is a user assigned managed identity
        // and it is needed to connect to the key vault that stores the secrets
        // we need to keep track of it 
        "clientId": "742444a6-0477-4342-998d-d0bc91504106",
        "objectId": "a18a7035-639d-43e3-b880-454ec7d2924f",
        "resourceId": "/subscriptions/0990639e-03f6-4b45-8fdd-143da61a6e2d/resourcegroups/MC_benpetersen-rg-akv-on-aks_akv-on-aks-with-sscsi_eastus2/providers/Microsoft.ManagedIdentity/userAssignedIdentities/azurekeyvaultsecretsprovider-akv-on-aks-with-sscsi"
      }
    }
  },
  "agentPoolProfiles": [
    {
      "availabilityZones": null,
      "capacityReservationGroupId": null,
      "count": 3,
      "creationData": null,
      "currentOrchestratorVersion": "1.29.9",
      "enableAutoScaling": false,
      "enableEncryptionAtHost": false,
      "enableFips": false,
      "enableNodePublicIp": false,
      "enableUltraSsd": false,
      "gpuInstanceProfile": null,
      "hostGroupId": null,
      "kubeletConfig": null,
      "kubeletDiskType": "OS",
      "linuxOsConfig": null,
      "maxCount": null,
      "maxPods": 110,
      "minCount": null,
      "mode": "System",
      "name": "nodepool1",
      "networkProfile": null,
      "nodeImageVersion": "AKSUbuntu-2204gen2containerd-202410.15.0",
      "nodeLabels": null,
      "nodePublicIpPrefixId": null,
      "nodeTaints": null,
      "orchestratorVersion": "1.29",
      "osDiskSizeGb": 128,
      "osDiskType": "Managed",
      "osSku": "Ubuntu",
      "osType": "Linux",
      "podSubnetId": null,
      "powerState": {
        "code": "Running"
      },
      "provisioningState": "Succeeded",
      "proximityPlacementGroupId": null,
      "scaleDownMode": null,
      "scaleSetEvictionPolicy": null,
      "scaleSetPriority": null,
      "securityProfile": {
        "enableSecureBoot": false,
        "enableVtpm": false
      },
      "spotMaxPrice": null,
      "tags": null,
      "type": "VirtualMachineScaleSets",
      "upgradeSettings": {
        "drainTimeoutInMinutes": null,
        "maxSurge": "10%",
        "nodeSoakDurationInMinutes": null
      },
      "vmSize": "Standard_DS2_v2",
      "vnetSubnetId": null,
      "windowsProfile": null,
      "workloadRuntime": null
    }
  ],
  "apiServerAccessProfile": null,
  "autoScalerProfile": null,
  "autoUpgradeProfile": {
    "nodeOsUpgradeChannel": "NodeImage",
    "upgradeChannel": null
  },
  "azureMonitorProfile": null,
  //
  // FQDN is probably useful for osmeting
  //
  "azurePortalFqdn": "akv-on-aks-benpetersen-rg-a-099063-9rfmdvqg.portal.hcp.eastus2.azmk8s.io",
  // kube version 
  "currentKubernetesVersion": "1.29.9",
  "disableLocalAccounts": false,
  "diskEncryptionSetId": null,
  // dns prefix might be interesting
  "dnsPrefix": "akv-on-aks-benpetersen-rg-a-099063",
  "enablePodSecurityPolicy": null,
  "enableRbac": true,
  "extendedLocation": null,
  // fqdn again 
  "fqdn": "akv-on-aks-benpetersen-rg-a-099063-9rfmdvqg.hcp.eastus2.azmk8s.io",
  "fqdnSubdomain": null,
  "httpProxyConfig": null,
  "id": "/subscriptions/0990639e-03f6-4b45-8fdd-143da61a6e2d/resourcegroups/benpetersen-rg-akv-on-aks/providers/Microsoft.ContainerService/managedClusters/akv-on-aks-with-sscsi",
  // identity is likely valuable
  "identity": {
    "delegatedResources": null,
    // principal id may be useful
    "principalId": "a0fddd15-32be-46de-8b8d-509784ef838f",
    "tenantId": "72f988bf-86f1-41af-91ab-2d7cd011db47",
    "type": "SystemAssigned",
    "userAssignedIdentities": null
  },
  "identityProfile": {
    // kubelet
    "kubeletidentity": {
      "clientId": "c6d75aaf-aceb-4938-bbc4-f8db18704f85",
      "objectId": "a21f0701-6052-4216-9b1c-dad96120f6d3",
      "resourceId": "/subscriptions/0990639e-03f6-4b45-8fdd-143da61a6e2d/resourcegroups/MC_benpetersen-rg-akv-on-aks_akv-on-aks-with-sscsi_eastus2/providers/Microsoft.ManagedIdentity/userAssignedIdentities/akv-on-aks-with-sscsi-agentpool"
    }
  },
  "ingressProfile": null,
  "kubernetesVersion": "1.29",
  // linux profile is useful
  "linuxProfile": {
    // admin username if we need it
    "adminUsername": "azureuser",
    // and a public key
    "ssh": {
      // contents of ~/.ssh/id_rsa.pub locally
      // the az cli output when we first ran indicated it was going to 
      // generate a pub/private key.  the private key likely goes up to 
      // Azure so that it can be used to validate when we talk to it
      // (via messages signed with our private key?)
      "publicKeys": [
        {
          "keyData": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC01RjYaoKmvhKjHvb5kmopRAA3keaUC5SqFT5fWLCFrds2pa8URl4XLg7Ze7717xDgv/7OdV1sGtcmVtSQQTPH0iOHCT3iMqrOUvCojE5p5cqO/nHRjfJ8sohMxVP15d/aSf1DfEJWwbeYx2kkFA3+jvcOh5uha/lVA7zQ0XSg1uQCS38V/tqaUaG6EeoJuIxxF0/GgQKZAW7GYrZj6HAtN4WVAWUmpgQQhcerEnTYh2uw/VDY/sVoNvtsSDm4//Ppf7uLOG4A98Y9+/KYeGWr0mJKRDki3QQBU80c2N7WwapcUgCegS59JAKC+qJGjCBDuWg2h8TShPjjSveVdPBX"
        }
      ]
    }
  },
  // eastus2
  "location": "eastus2",
  "maxAgentPools": 100,
  "metricsProfile": {
    "costAnalysis": {
      "enabled": false
    }
  },
  // name took a while to get to 
  "name": "akv-on-aks-with-sscsi",
  "networkProfile": {
    "dnsServiceIp": "10.0.0.10",
    // only ipv4 we did not ask for ipv6
    "ipFamilies": [
      "IPv4"
    ],
    "loadBalancerProfile": {
      "allocatedOutboundPorts": null,
      "backendPoolType": "nodeIPConfiguration",
      "effectiveOutboundIPs": [
        {
          "id": "/subscriptions/0990639e-03f6-4b45-8fdd-143da61a6e2d/resourceGroups/MC_benpetersen-rg-akv-on-aks_akv-on-aks-with-sscsi_eastus2/providers/Microsoft.Network/publicIPAddresses/a818858d-d3c1-496f-926a-94ddc2d43cf8",
          "resourceGroup": "MC_benpetersen-rg-akv-on-aks_akv-on-aks-with-sscsi_eastus2"
        }
      ],
      "enableMultipleStandardLoadBalancers": null,
      "idleTimeoutInMinutes": null,
      "managedOutboundIPs": {
        "count": 1,
        "countIpv6": null
      },
      "outboundIPs": null,
      "outboundIpPrefixes": null
    },
    "loadBalancerSku": "standard",
    "natGatewayProfile": null,
    "networkDataplane": null,
    "networkMode": null,
    "networkPlugin": "kubenet",
    "networkPluginMode": null,
    "networkPolicy": "none",
    "outboundType": "loadBalancer",
    // podCidr is likely interesting
    // should look at this again when I do 
    // kube the hard way on Azure as a refresher
    "podCidr": "10.244.0.0/16",
    // not sure why its here twice
    "podCidrs": [
      "10.244.0.0/16"
    ],
    // need to look into this and think about why services get a 
    // different CIDR than pods
    "serviceCidr": "10.0.0.0/16",
    "serviceCidrs": [
      "10.0.0.0/16"
    ]
  },
  "nodeResourceGroup": "MC_benpetersen-rg-akv-on-aks_akv-on-aks-with-sscsi_eastus2",
  // OIDC Issuer is useful 
  "oidcIssuerProfile": {
    "enabled": true,
    "issuerUrl": "https://eastus2.oic.prod-aks.azure.com/72f988bf-86f1-41af-91ab-2d7cd011db47/43f9a1e4-5ff7-4f9e-a995-2f168afa6576/"
  },
  // i thought we gave it a pod identity flag, this seems odd
  "podIdentityProfile": null,
  "powerState": {
    "code": "Running"
  },
  // nope
  "privateFqdn": null,
  "privateLinkResources": null,
  "provisioningState": "Succeeded",
  // no public network yet
  "publicNetworkAccess": null,
  // RG recorded
  "resourceGroup": "benpetersen-rg-akv-on-aks",
  "resourceUid": "672ad75419d43b0001ca1571",
  // dunno what this means but might be interesting
  "securityProfile": {
    "azureKeyVaultKms": null,
    "defender": null,
    "imageCleaner": null,
    "workloadIdentity": {
      "enabled": true
    }
  },
  // didn't pick a service mesh explicitly 
  "serviceMeshProfile": null,
  "servicePrincipalProfile": {
    "clientId": "msi",
    "secret": null
  },
  "sku": {
    "name": "Base",
    "tier": "Free"
  },
  "storageProfile": {
    "blobCsiDriver": null,
    "diskCsiDriver": {
      "enabled": true
    },
    "fileCsiDriver": {
      "enabled": true
    },
    "snapshotController": {
      "enabled": true
    }
  },
  // ha
  "supportPlan": "KubernetesOfficial",
  "systemData": null,
  "tags": null,
  // managed
  "type": "Microsoft.ContainerService/ManagedClusters",
  "upgradeSettings": null,
  "windowsProfile": null,
  "workloadAutoScalerProfile": {
    "keda": null,
    "verticalPodAutoscaler": null
  }
}
```
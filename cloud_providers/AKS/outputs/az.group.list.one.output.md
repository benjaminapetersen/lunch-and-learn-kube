```bash
az group list --query "[].{name:name}" --output tsv
# DefaultResourceGroup-CUS
# benpetersen-rg-akv-on-aks
# MC_benpetersen-rg-akv-on-aks_akv-on-aks-with-sscsi_eastus2
# benpetersen-rg
# NetworkWatcherRG
# DefaultResourceGroup-WUS2
# MC_benpetersen-kube-rg_kube-1_eastus
# MA_defaultazuremonitorworkspace-wus2_westus2_managed
# benpetersen-kube-rg
# DefaultResourceGroup-EUS
```
# Accessing Clusters

https://kubernetes.io/docs/tasks/access-application-cluster/access-cluster/


```bash
# fairly human readable but not easy to parse to get the IP
kubectl cluster-info
# see your kube config details 
kubectl config view
```

## Using kube proxy

https://kubernetes.io/docs/tasks/access-application-cluster/access-cluster/#using-kubectl-proxy

```bash
kube proxy --port=8080
curl http://localhost:8080/api/
# {\"kind\":\"APIVersions\",\"versions\":[\"v1\"],\"serverAddressByClientCIDRs\":[{\"clientCIDR\":\"0.0.0.0/0\",\"serverAddress\":\"10.0.1.149:443\"}]}
```

## Without Kube Proxy

https://kubernetes.io/docs/tasks/access-application-cluster/access-cluster/#without-kubectl-proxy

```bash
APISERVER=$(kubectl config view --minify | grep server | cut -f 2- -d ":" | tr -d " ")
# https://10.186.96.230:6443
# https://<ip>:<port>
SECRET_NAME=$(kubectl get secrets | grep ^default | cut -f1 -d ' ')
# default-token-<hash>
TOKEN=$(kubectl describe secret $SECRET_NAME | grep -E '^token' | cut -f2 -d':' | tr -d " ")

curl $APISERVER/api --header "Authorization: Bearer $TOKEN" --insecure
# {\"kind\":\"APIVersions\",\"versions\":[\"v1\"],\"serverAddressByClientCIDRs\":[{\"clientCIDR\":\"0.0.0.0/0\",\"serverAddress\":\"10.0.1.149:443\"}]}
```

## Using JSON Path 

```bash
APISERVER=$(kubectl config view --minify -o jsonpath='{.clusters[0].cluster.server}')
SECRET_NAME=$(kubectl get serviceaccount default -o jsonpath='{.secrets[0].name}')
TOKEN=$(kubectl get secret $SECRET_NAME -o jsonpath='{.data.token}' | base64 --decode)

curl $APISERVER/api --header "Authorization: Bearer $TOKEN" --insecure
# {\"kind\":\"APIVersions\",\"versions\":[\"v1\"],\"serverAddressByClientCIDRs\":[{\"clientCIDR\":\"0.0.0.0/0\",\"serverAddress\":\"10.0.1.149:443\"}]}
```

## Control Plane Endpoint

Other ways to get the control plane endpoint

```bash
# Just the IP address
kubectl config view --minify | grep server | cut -f 3- -d ":" | cut -f1 -d":" | cut -f 3- -d"/"
# <ip>
# 10.186.96.230
# IP with port, etc
kubectl config view --flatten -o json | jq .clusters[0].cluster.server -r
# https://10.186.96.230:6443
# https://<ip>:<port>
```

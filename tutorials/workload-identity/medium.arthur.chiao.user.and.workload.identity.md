# User and workload identities in Kubernetes

- Author: Arthur Chiao
- Date: June 2022

## Topics

- User Identities
- Workload Identities
- Kubectl operations
- HTTP endpoints

## Intro

```bash
# lets make a cluster with kind to mess with
kind create cluster --name tut-workload-identities
kubectl config get-contexts
kubetl config use-context 
kubectl config use-context tut-workload-identities
kubectl cluster-info
# Kubernetes control plane is running at https://127.0.0.1:60104
# CoreDNS is running at https://127.0.0.1:60104/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

# for convenience
export API_SERVER_URL=https://127.0.0.1:60104 # output of above cluster info
echo $API_SERVER_URL
```
## Accessing the Kubernetes API with curl

```bash
export API_SERVER_URL=https://127.0.0.1:60104
curl $API_SERVER_URL/api/v1/namespaces
# SSL cert problem: unable to get local issuer certificate
```
Can suppress:

```bash
# now succeeds, but rejected request
curl -k $API_SERVER_URL/api/v1/namespaces
# {
#   "kind": "Status",
#   "apiVersion": "v1",
#   "code": 403
#   "status": "Failure",
#   "message": "namespaces is forbidden: User \"system:anonymous\" cannot list resource \"namespaces\" in API group \"\" at the cluster scope",
#   "reason": "Forbidden",
# }%
```
This performed both:
- `authentication` (who are you?)
  - Nothing provided for auth, so you are considered `system:anonymous`
- `authorization` (what can you do?)
  - `system:anonymous` can't do much

The `apiserver` is modular.  One of the first modules to receive the 
request is the authentication module. 

Via the above curl:
- Since no user credentials provided
  - Kubernetes Authentication Module could not assign identity
  - Request labeled anonymous, `system:anonymous`
- Config of api server could also have
  - respose `401 unauthorized`
- Once authenticated `system:anonymous` (no credentials)
  - Kubernetes Authorization module assesses 
  - `system:anonymous` forbidden to `READ namespaces`
  - `403 forbidden` error

`NOTE`: Your request was issued from outside the cluster, but it could have also
been issued from inside the cluster, via one of the nodes of the cluster.
- Kubelet could issue a request


Modules
- Authentication module is the first gatekeeper of the system
  - Authenticates with
    - static token
    - certificate
    - externally-managed identity
  - Noteworthy features
    - support `human` and `machine` (program) users
    - support `external` and `internal` users
      - apps deployed outside the cluster
      - accounts created & namaged by Kubernetes
    - supports `standard authentication strategies`
      - static token
      - bearer token
      - X509 certificate
      - OIDC
      - etc
    - supports `multiple authentication stratgies` simultaneously
    - authentication stragies can be updated
      - add new ones
      - phase out old ones
    - optionally allow `anonymous access` to the API
- Authorization module is second
  - [Authorization happens via the RBAC system](https://learnk8s.io/rbac-kubernetes)
  - (Older ABAC is no longer maintained, I believe)

## The Kubernetes API Differentiate Internal and External Users

- Why distinction between internal and external users?
  - Internal users need a specification
    - A data model
  - External users specification exists elsewhere
- Kinds of users
  - `Kubernetes managed users`  
    - User accounts created by Kuberentes cluster itself
    - used by in-cluster apps
  - `Non-Kubernetes managed users`
    - External to the cluster users
    - Users with (provided by cluster admins):
      - static tokens 
      - certificates 
    - Users authenticated through external identity providers
      - Keystone
      - Google account
      - LDAP
      - etc

## Granting access to the cluster to external users

```bash
curl 
  --cacert ${CACERT} 
  --header "Authorization: Bearer <my-token>" 
  -X GET ${API_SERVER_URL}/api
```
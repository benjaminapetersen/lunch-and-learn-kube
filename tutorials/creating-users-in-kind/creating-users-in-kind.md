# Creating Users in Kind

- https://medium.com/@lionelvillard/creating-users-in-kind-cluster-6c5ee35db3fe

```bash
# need a cluster
#  in docker (kind is kubernetes-in-docker)
# and need a certificate
#  user key
#  certificate signing request (CSR) 
kind create cluster --name tut-kind-users
```

Now we need the csr

```bash
# this will get the CSR
kind get kubeconfig --name <cluster-name> | yq e '.users[0].user.client-certificate-data' - | base64 -D > kind.csr
# this allows us to look at it
openssl x509 -in kind.csr -noout -text
```
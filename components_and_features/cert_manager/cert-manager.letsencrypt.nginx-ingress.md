# Cert Manager

This document follows the steps I used working through [the cert manager docs](https://cert-manager.io/docs/)

## POSSIBLE GOTCHAS!!!!

- This doc is following the cert-manager tutorial which suggest installing the `NGINX Ingress Controller`.
- It appears that there is a [Contour Package](https://github.com/vmware-tanzu/community-edition/tree/main/addons/packages/contour)
  in TKG, which is also an `ingress controller`, but it is not enabled by default.
- There are docs for [Deploy and Manage Contour Ingress on TKG](https://docs.vmware.com/en/VMware-vSphere/7.0/vmware-vsphere-with-tanzu/GUID-A1288362-61F7-46D9-AB42-1A5711AB4B57.html)
- For now, I am assuming since it is not deployed by default that I can carry on with NGINX Ingress Controller and not worry about it.

## First, Helm for Nginx Ingress

The first thing that cert manager wanted was [helm installed](https://cert-manager.io/docs/tutorials/acme/nginx-ingress/)
So, since I am using a `TKG 1.6.0` cluster via a `jumpbox`, I went to [helm install instructions](https://helm.sh/docs/intro/install/)

Important to note, there appears to be a `cert-manager` installed by default on TKG 1.6.x:

```bash
kubectl get deployment -A | grep cert
# cert-manager      cert-manager       1/1     1            1           2d5h
```
But also, [the docs](https://docs.vmware.com/en/VMware-Tanzu-Kubernetes-Grid/1.6/vmware-tanzu-kubernetes-grid-16/GUID-packages-cert-manager.html)
indicate that the `cert-manager` package should be installed. 🤔
So? I guess I'll try to use the one OOTB, if it doesn't work I'll install the package. I'll make a `cert-manager` doc in
my `lunch-and-learn-vmare` repository if there are any details to worry about.

What kind of box is the `jumpbox`?

```bash
hostnamectl
#    Static hostname: 92nxWLAoIwRPE
#          Icon name: computer-vm
#            Chassis: vm
#         Machine ID: 45f42e71300845a996ee1a3414517093
#            Boot ID: f0fb3cdb15704062a86973eba33ac182
#     Virtualization: vmware
#   Operating System: Ubuntu 20.04.4 LTS
#             Kernel: Linux 5.4.0-128-generic
#       Architecture: x86-64
```
Confirmed it is `ubuntu`.  Jumpboxes always seem to be `ubuntu`, but good to check.

```bash
# install helm
# if we were on MacOS we would use `brew install helm` instead, but since Ubuntu:
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh
```

## Kubernetes NGINX Ingress Controller

[Docs for Kubernetes NGINX Ingress Controller](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/)

`Kubernetes Nginx Ingress Controller` is the nginx controller provided by kubernetes itself.
You need an ingress controller to process all of the ingress rules defined by any ingress you
create and install on a cluster.  There are many third party ingress controllers.
`Kubernetes Nginx Ingress Controller` is a perfectly fine option to use unless you have reason to go
with something else.

## Deploy NGINX Ingress Controller with Helm

Time to deploy this thing!

[Deploy NGinx Ingress Controller](https://cert-manager.io/docs/tutorials/acme/nginx-ingress/)

```bash
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install quickstart ingress-nginx/ingress-nginx
# wait for a public IP address via a LoadBalancer
# watch for status via:
kubectl --namespace default get services -o wide -w quickstart-ingress-nginx-controller
# then go ahead and check the service
kubectl get svc
# but... it won't have TLS!
```
The output of `helm install quickstart ingress-nginx/ingress-nginx` is:

```yaml
# An example Ingress that makes use of the controller:
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: example
  namespace: foo
spec:
  ingressClassName: nginx
  rules:
    - host: www.example.com
      http:
        paths:
          - pathType: Prefix
            backend:
              service:
                name: exampleService
                port:
                  number: 80
            path: /
  # This section is only required if TLS is to be enabled for the Ingress
  tls:
    - hosts:
      - www.example.com
      secretName: example-tls
# If TLS is enabled for the Ingress, a Secret containing the certificate and key must also be provided:
apiVersion: v1
kind: Secret
metadata:
  name: example-tls
  namespace: foo
data:
  tls.crt: <base64 encoded cert>
  tls.key: <base64 encoded key>
type: kubernetes.io/tls
```
Which tells us a couple of things we probably have to do next.
If we run the suggested command to look for the `quickstart-ingress-nginx-controller service` we can see:

```bash
# -w means "watch"... so we should see some updates
kubectl --namespace default get services -o wide -w quickstart-ingress-nginx-controller
# NAME                                  TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)                      AGE     SELECTOR
# quickstart-ingress-nginx-controller   LoadBalancer   100.71.61.214   10.92.104.204   80:31681/TCP,443:31838/TCP   2m11s   app.kubernetes.io/component=controller,app.kubernetes.io/instance=quickstart,app.kubernetes.io/name=ingress-nginx
```
Recall that if `EXTERNAL-IP` is `<pending>` you have to wait until the IP is provided.
from the point that it becomes ready, we can just do a usual service lookup:
```bash
kubectl get svc
# they don't all have EXTERNAL-IP!
# NAME                                            TYPE           CLUSTER-IP      EXTERNAL-IP     PORT(S)                      AGE
# kubernetes                                      ClusterIP      100.64.0.1      <none>          443/TCP                      27h
# quickstart-ingress-nginx-controller             LoadBalancer   100.71.61.214   10.92.104.204   80:31681/TCP,443:31838/TCP   3m9s
# quickstart-ingress-nginx-controller-admission   ClusterIP      100.68.116.88   <none>          443/TCP                      3m9s
```

Then we need to add the IP to a DNS zone we control, such as `www.example.com`.
I'm going to use `ingress-nginx.benjaminapetersen.me` as I have this domain.

So in [GoDaddy DNS Manager](https://dcc.godaddy.com/control/benjaminapetersen.me/dns?referrer=venturehome&plid=1) I've:

```yaml
Type: A Record
Name: ingress-nginx
Data: 10.92.104.204
TTL: 600
```

Visiting `http://ingress-nginx.benjaminapetersen.me/` will result in a 404 hosted by Nginx.
Pinging the domain will indeed resolve to the expected IP address:

```bash
ping ingress-nginx.benjaminapetersen.me
# PING ingress-nginx.benjaminapetersen.me (10.92.104.204) 56(84) bytes of data.
# 64 bytes from wdc-rdops-vm06-dhcp-104-204.eng.vmware.com (10.92.104.204): icmp_seq=1 ttl=64 time=1.09 ms
# 64 bytes from wdc-rdops-vm06-dhcp-104-204.eng.vmware.com (10.92.104.204): icmp_seq=2 ttl=64 time=0.469 ms
```

## Deploy the Kubernetes-Up-And-Running-Demo App (KUARD)

Find the repository here: [KUARD](https://github.com/kubernetes-up-and-running/kuard)
The deployment instructions continue from [the cert-manager tutorial page](https://cert-manager.io/docs/tutorials/acme/nginx-ingress/)

Two yaml documents to begin:

```yaml
# put them together as one, or separate as kuard-deployment.yaml, kuard-service.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kuard
spec:
  selector:
    matchLabels:
      app: kuard
  replicas: 1
  template:
    metadata:
      labels:
        app: kuard
    spec:
      containers:
      - image: gcr.io/kuar-demo/kuard-amd64:1
        imagePullPolicy: Always
        name: kuard
        ports:
        - containerPort: 8080
---
apiVersion: v1
kind: Service
metadata:
  name: kuard
spec:
  # no type! this is default clusterIP, an internal service
  ports:
  - port: 80
    targetPort: 8080
    protocol: TCP
    # no nodePort! this isn't meant to be reached externally, an ingress must be in front.
  selector:
    app: kuard
```
These can also be referenced directy from the github URLs:

```bash
kubectl apply -f https://raw.githubusercontent.com/cert-manager/website/master/content/docs/tutorials/acme/example/deployment.yaml
# expected output: deployment.extensions "kuard" created
kubectl apply -f https://raw.githubusercontent.com/cert-manager/website/master/content/docs/tutorials/acme/example/service.yaml
# expected output: service "kuard" created
```

Likely you used the `default` namespace, else double check

```bash
# wherever
kubectl get all -n default
```

Now, we need an `ingress`.

```bash
kubectl get ingress -A
# Nada.
```
So lets make one:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: kuard
  annotations:
    kubernetes.io/ingress.class: "nginx"
    #cert-manager.io/issuer: "letsencrypt-staging"
spec:
  tls:
  - hosts:
    # - example.example.com
    - ingress-nginx.benjaminapetersen.me
    secretName: quickstart-example-tls
  rules:
  - host: ingress-nginx.benjaminapetersen.me
    # host: example.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: kuard
            port:
              number: 80
```
This file can also be pulled directly from github:

```bash
# pull, edit, and deploy:
kubectl create --edit -f https://raw.githubusercontent.com/cert-manager/website/master/content/docs/tutorials/acme/example/ingress.yaml
# expected output: ingress.extensions "kuard" created
```
Your ingress will look like:

```bash
kubectl get ingress
# NAME    CLASS    HOSTS                                ADDRESS   PORTS     AGE
# kuard   <none>   ingress-nginx.benjaminapetersen.me             80, 443   4s
# not ready yet!
# run it again and see if you get an address
kubectl get ingress -w
# NAME    CLASS    HOSTS                                ADDRESS         PORTS     AGE
# kuard   <none>   ingress-nginx.benjaminapetersen.me   10.92.104.204   80, 443   3h58m
# aha! ready.
```

Test the ingress:

- open a browser to `http://ingress-nginx.benjaminapetersen.me`
- or run `curl -kivL -H 'Host: ingress-nginx.benjaminapetersen.me' 'http://10.92.104.204'`
  - use the subdomain you setup + the ip address reported by the ingress.

The `curl` command will dump lots of interesting information. It is:
- `-v` verbose output, Lines starting with `>` indicate heders sent by curl, `<` are headers received by curl. `*` is additional info provided by curl
- `-k` insecure. skip TLS verification. the name should match the most name in the URL and the cert should be signed by CA in the cert store.
- `-i` include the HTTP response headers in output. this should be server name, cookies, date, http version, etc.
- `-L` include location. if the response indicates the resource has moved, this will cause curl to issue a new request to the new location.


```bash
*   Trying 10.92.104.204:80...
* TCP_NODELAY set
* Connected to 10.92.104.204 (10.92.104.204) port 80 (#0)
> GET / HTTP/1.1
> Host: ingress-nginx.benjaminapetersen.me
> User-Agent: curl/7.68.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 308 Permanent Redirect
HTTP/1.1 308 Permanent Redirect
< Date: Wed, 26 Oct 2022 19:19:53 GMT
Date: Wed, 26 Oct 2022 19:19:53 GMT
< Content-Type: text/html
Content-Type: text/html
< Content-Length: 164
Content-Length: 164
< Connection: keep-alive
Connection: keep-alive
< Location: https://ingress-nginx.benjaminapetersen.me
Location: https://ingress-nginx.benjaminapetersen.me

<
* Ignoring the response-body
* Connection #0 to host 10.92.104.204 left intact
* Clear auth, redirects to port from 80 to 443Issue another request to this URL: 'https://ingress-nginx.benjaminapetersen.me/'
*   Trying 10.92.104.204:443...
* TCP_NODELAY set
* Connected to ingress-nginx.benjaminapetersen.me (10.92.104.204) port 443 (#1)
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: /etc/ssl/certs/ca-certificates.crt
  CApath: /etc/ssl/certs
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
* TLSv1.3 (IN), TLS handshake, Server hello (2):                                                    # well hello there
* TLSv1.3 (IN), TLS handshake, Encrypted Extensions (8):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Change cipher spec (1):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / TLS_AES_256_GCM_SHA384
* ALPN, server accepted to use h2
* Server certificate:
*  subject: O=Acme Co; CN=Kubernetes Ingress Controller Fake Certificate                            # this is a fun CN!
*  start date: Oct 26 15:02:33 2022 GMT
*  expire date: Oct 26 15:02:33 2023 GMT
*  issuer: O=Acme Co; CN=Kubernetes Ingress Controller Fake Certificate                             # yay Fake
*  SSL certificate verify result: unable to get local issuer certificate (20), continuing anyway.   # uh oh?
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* Using Stream ID: 1 (easy handle 0x55c32dd7a2f0)
> GET / HTTP/2
> Host: ingress-nginx.benjaminapetersen.me                                                          # this is the host we setup
> user-agent: curl/7.68.0
> accept: */*
>
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* old SSL session ID is stale, removing
* Connection state changed (MAX_CONCURRENT_STREAMS == 128)!
< HTTP/2 200
HTTP/2 200
< date: Wed, 26 Oct 2022 19:19:54 GMT
date: Wed, 26 Oct 2022 19:19:54 GMT
< content-type: text/html
content-type: text/html
< content-length: 1737
content-length: 1737
< strict-transport-security: max-age=15724800; includeSubDomains
strict-transport-security: max-age=15724800; includeSubDomains

<                                                                                                   # yay HTML pages
<!doctype html>

<html lang="en">
<head>
  <meta charset="utf-8">

  <title>KUAR Demo</title>

  <link rel="stylesheet" href="/static/css/bootstrap.min.css">
  <link rel="stylesheet" href="/static/css/styles.css">

  <script>
var pageContext = {"hostname":"kuard-7b5bffcc4f-4tc7j","addrs":["100.96.1.37"],"version":"v0.8.1-1","versionColor":"hsl(18,100%,50%)","requestDump":"GET / HTTP/1.1\r\nHost: ingress-nginx.benjaminapetersen.me\r\nAccept: */*\r\nUser-Agent: curl/7.68.0\r\nX-Forwarded-For: 100.96.0.1\r\nX-Forwarded-Host: ingress-nginx.benjaminapetersen.me\r\nX-Forwarded-Port: 443\r\nX-Forwarded-Proto: https\r\nX-Forwarded-Scheme: https\r\nX-Real-Ip: 100.96.0.1\r\nX-Request-Id: ca5e1dcafb7b6fed9e4a729363393648\r\nX-Scheme: https","requestProto":"HTTP/1.1","requestAddr":"100.96.1.35:48408"}
  </script>
</head>


<svg style="position: absolute; width: 0; height: 0; overflow: hidden;" version="1.1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
<defs>
<symbol id="icon-power" viewBox="0 0 32 32">
<title>power</title>
<path class="path1" d="M12 0l-12 16h12l-8 16 28-20h-16l12-12z"></path>
</symbol>
<symbol id="icon-notification" viewBox="0 0 32 32">
<title>notification</title>
<path class="path1" d="M16 3c-3.472 0-6.737 1.352-9.192 3.808s-3.808 5.72-3.808 9.192c0 3.472 1.352 6.737 3.808 9.192s5.72 3.808 9.192 3.808c3.472 0 6.737-1.352 9.192-3.808s3.808-5.72 3.808-9.192c0-3.472-1.352-6.737-3.808-9.192s-5.72-3.808-9.192-3.808zM16 0v0c8.837 0 16 7.163 16 16s-7.163 16-16 16c-8.837 0-16-7.163-16-16s7.163-16 16-16zM14 22h4v4h-4zM14 6h4v12h-4z"></path>
</symbol>
</defs>
</svg>

<body>
  <div id="root"></div>
  <script src="/built/bundle.js" type="text/javascript"></script>
</body>
</html>
* Connection #1 to host ingress-nginx.benjaminapetersen.me left intact
```

## And now.... Cert Manager! Deploy It

Everything up to this point has just been for the purpose of creating Ingress.  Ingress are necessary for incoming traffic
into a kubernetes cluster.  Most apps need a domain name, a URL, ingress, and a TLS certificate to be verified.

How to get it?

If running a TKG cluster, cert manager is already deployed:

```bash
kubectl get deployment -A | grep cert
# cert-manager                        cert-manager                                    1/1     1            1           32h
# cert-manager                        cert-manager-cainjector                         1/1     1            1           32h
# cert-manager                        cert-manager-webhook                            1/1     1            1           32h
```

If running on minikube or something else, `Helm` continues to be the way to go to get this running on the cluster.

- See [Deploying Cert Manager with Helm](https://cert-manager.io/docs/installation/helm/)

## Cert Manager Resources

Cert manager uses CRDs to configure and control how it operates & store its state.

- `Issuers`: Define HOW cert-manager will request TLS certificates.
  - `Issuers` are namespace scoped. be sure to create them in the same namespace as the cert you want to create.
    - ingress will require this annotation:
      - `annotation: cert-manager.io/issuer`
  - `ClusterIssuers` are cluster wide. These apply across all ingress resources on a cluster.
    - ingress will instead require this annotation:
    - `annotation: cert-manager.io/cluster-issuer`
  - [More issuer information](https://cert-manager.io/docs/concepts/issuer/#namespaces)
  - [Issues with Issuers may want to see the Troubleshooting ACME Certificates guide](https://cert-manager.io/docs/faq/acme/)
- `Certificates`: define details about the certificate you want to request
  - Certificates reference an Issuer to define how they will be used.
  - [More certificate information](https://cert-manager.io/docs/concepts/certificate/)


## LetsEncrypt Issuer!

We will create two LetsEncrypt Issuers:
- staging
  - use this while learning
  - Untrusted certificates warnings will be seen
- production
  - Strict rate limits enforced


## Nodes, do the IPs have to match the ingress?

For good measure, lets note the `nodes` and their `EXTERNAL-IP` addresses for this cluster:

```bash
kubectl get nodes -o wide # wide provides EXTERNAL-IP and many other things
# NAME                                STATUS   ROLES                  EXTERNAL-IP     ...
# tkg-mgmt-vc-control-plane-drsk8     Ready    control-plane,master   10.92.108.197   ...
# tkg-mgmt-vc-md-0-7bd8b547f4-992ck   Ready    <none>                 10.92.118.1     ...
```

Create the `staging issuer`:

```yaml
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: letsencrypt-staging
spec:
  acme:
    # The ACME server URL - staging
    server: https://acme-staging-v02.api.letsencrypt.org/directory
    # Email address used for ACME registration
    email: petersenbe@vmware.com                                                # maybe should use ben@benjaminapetersen.me to match my domain
    # Name of a secret used to store the ACME account private key
    privateKeySecretRef:
        name: letsencrypt-staging
    # Enable the HTTP-01 challenge provider
    solvers:
    - http01:
        ingress:
            class:  nginx
```

This file can also be edited directly from the github examples:

```bash
kubectl create --edit -f https://raw.githubusercontent.com/cert-manager/website/master/content/docs/tutorials/acme/example/staging-issuer.yaml
# expected output: issuer.cert-manager.io "letsencrypt-staging" created
```

Create the `production issuer`:

```yaml
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: letsencrypt-prod
spec:
  acme:
    # The ACME server URL - prod
    server: https://acme-v02.api.letsencrypt.org/directory
    # Email address used for ACME registration
    email: petersenbe@vmware.com                                                # maybe should use ben@benjaminapetersen.me to match my domain
    # Name of a secret used to store the ACME account private key
    privateKeySecretRef:
      name: letsencrypt-prod
    # Enable the HTTP-01 challenge provider
    solvers:
    - http01:
        ingress:
          # so this is interesting, implies our use of the nginx ingress controller
          # now, on a TKG environment, there is likely already an ingress.
          class: nginx
```

Again can edit directly from the github examples:

```bash
kubectl create --edit -f https://raw.githubusercontent.com/cert-manager/website/master/content/docs/tutorials/acme/example/production-issuer.yaml
# expected output: issuer.cert-manager.io "letsencrypt-prod" created
```

For what its worth, we are configuring `ACME HTTP-01` challenges.  Is this secure?
How can HTTP be used to bootstrap HTTPS?
Is there a vulnerabilty to HTTPS using this?  Good homework for later!

Check on the Issers:

```bash
kubectl get issuers
# NAME                  READY   AGE
# letsencrypt-prod      True    5m51s
# letsencrypt-staging   True    8m35s
kubectl describe issuer letsencrypt-staging
kubectl describe issuer letsencrypt-prod
# and look for:
#    Reason:                ACMEAccountRegistered
# and also the url path, something like:
#    .../acme/acct/7374163
```

## Deploying a TLS Resource

Now we can request TLS Certificates.
There are two ways:
- using `ingress-shim` which allows us to just add `annotations` to `ingress` objects
- directly create `certificate` resources

Here we will use `ingress-shim` and `annotations` to create our certificates.
Cert Manager will create or update ingress resources to validate the domain.
Once verified and issued, cert-manager will update or create the secret defined in the certificate.

NOTE: the `secret` used in the `ingress` should match the `secret` used in the `certificate`.
There is no explicit checking!! This means that a `TYPO` will result in a fallback to a self-signed certificate.
This is likely part of the reason the `ingress-shim` is ideal.

Get your ingress, and update it:

```bash
kubectl get ingress kuard
vim kuard-ingress.yaml
```

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: kuard
  annotations:
    kubernetes.io/ingress.class: "nginx"
    # uncomment the following line
    # this is the annotation that will cause cert-manager to generate our certificates
    cert-manager.io/issuer: "letsencrypt-staging"
```

Then we get to look for certs:

```bash
# watch and wait for Ready: True
kubectl get certificate -w
# NAME                     READY   SECRET                   AGE
# quickstart-example-tls   False   quickstart-example-tls   8s
kubectl get certificate -o wide
# NAME                     READY   SECRET                   ISSUER                STATUS                                         AGE
# quickstart-example-tls   False   quickstart-example-tls   letsencrypt-staging   Issuing certificate as Secret does not exist   4m58s
kubectl describe certificate quickstart-example-tl
# hmm. errors? debugging time!
```
What can we do?  Well, we can checkout the `order`, which is a `Certificate Signing Request` sent to `LetsEncrypt`:

```bash
kubectl get order
# NAME                                      STATE     AGE
# quickstart-example-tls-6gzc2-2766096929   invalid   8m35s
```
If we want to check this out:

```bash
kubectl describe order quickstart-example-tls-6gzc2-2766096929
```
We get this big output:

```yaml
# kubectl describe order quickstart-example-tls-6gzc2-2766096929
Name:         quickstart-example-tls-6gzc2-2766096929
Namespace:    default
Labels:       <none>
Annotations:  cert-manager.io/certificate-name: quickstart-example-tls
              cert-manager.io/certificate-revision: 1
              cert-manager.io/private-key-secret-name: quickstart-example-tls-8wz6t
API Version:  acme.cert-manager.io/v1
Kind:         Order
Metadata:
  Creation Timestamp:  2022-10-31T19:45:51Z
  Generation:          1
  Managed Fields: # omitted
  Owner References:
    API Version:           cert-manager.io/v1
    # omited
  Resource Version:        118979
  UID:                     4b8fdc80-95d7-4294-8f14-09d28a4118f8
Spec:
  Dns Names:
    ingress-nginx.benjaminapetersen.me
  Issuer Ref:
    Group:  cert-manager.io
    Kind:   Issuer
    Name:   letsencrypt-staging # yup we want this...
  # this is likely something pretty fun.
  Request:  LS0tLS1CRUdJTiBDRVJUSUZJQ0FURSBSRVFVRVNULS0tLS0KTUlJQ2tqQ0NBWG9DQVFBd0FEQ0NBU0l3RFFZSktvWklodmNOQVFFQkJRQURnZ0VQQURDQ0FRb0NnZ0VCQU1NegpNQWEzS3RrRW9oYzE0WEt0N1hjWjVpNkJYMnJ3aHdxb3BENTAyTFZFb2lMUVNXN2R1VnNjNVFWdFhjczlSSFB2CjVielVZMlhCUjIxSUZ6UXUwdzVwVGhjMUMxMFRUWkdEQkN4WVZWM0FROUhwaTVYSE51TERKanhKT2xZWXljVTYKUS9YZzVoNHNEMnB4UGhMMCs4THpHL3M3OUEwbGVRbklUMUxMKzNwL1U3RndlWUZMeitDbzlmczlvSGhQRHZXUQpvWjFjY2M1d0hZMUFORVdhQzd4QjB6dWZJdjBhZ3ZENzNoMkVHdWU5NnJETVZ5MjdPZm5aWnNicUlqMGlPT3oyCmRHM29RaFhrcEtsRVYwODhJcnF4SVBNbkZZcmV2VGp6SEtTbW1rcU92MWF4SjJUM0ZUaGVyMldrakNHVVlDVnMKYzB4anMyUHh1bEU3Q29DbmlYY0NBd0VBQWFCTk1Fc0dDU3FHU0liM0RRRUpEakUrTUR3d0xRWURWUjBSQkNZdwpKSUlpYVc1bmNtVnpjeTF1WjJsdWVDNWlaVzVxWVcxcGJtRndaWFJsY25ObGJpNXRaVEFMQmdOVkhROEVCQU1DCkJhQXdEUVlKS29aSWh2Y05BUUVMQlFBRGdnRUJBS0NJSlR1VUpQTjE4M3VGUkNvRklNaHU2cUlpRkFNQm5jMU8KQkpBMUIxVDBOcUROVjNBNUNJZms2am45K0xKSzA1ZXlSNlhRUkp6UmErMmFqTEZoWmZsY3pJSHMyamNzcVdFbwowcE00RzY5YmNRQmYvMzhrakRNUFd0T3JLSm0reVQ3YjJJbnE3WjJ5QjVpc0NvbUZkakJoeE1TdWZpaVNzbjdSCmo5VFVpc1Q3cTNvd2dRWFVaRkE0aVNUSWk0UGFMdXNRYnp3cHJ3WTFZZDRPOFJxK3ZYUXNqNmwwZjRSU0d3ZWcKOEFrdnV5NWlkanhMWndjTm9GWWhURHMwc2Q3OTBEQ29Cclp0dlFLeEpHUE5GTXh4RXloUzFHQWd0MFdDaG9pSApPQksrejdLQ0syNWcxYnhrUzRKVUdIem9idDB4T2w1QnljSXVEVDZXMU9kT3pmSGFoem89Ci0tLS0tRU5EIENFUlRJRklDQVRFIFJFUVVFU1QtLS0tLQo=
Status:
  # this is all also likely fun.
  Authorizations:
    Challenges:
      Token:        GzauTfdg72Uy-cRt8HUize16ul6Ss6Nn4hUHYoo_Bkg
      Type:         http-01
      URL:          https://acme-staging-v02.api.letsencrypt.org/acme/chall-v3/4165525874/-PeDkg
      Token:        GzauTfdg72Uy-cRt8HUize16ul6Ss6Nn4hUHYoo_Bkg
      Type:         dns-01
      URL:          https://acme-staging-v02.api.letsencrypt.org/acme/chall-v3/4165525874/FS1ckg
      Token:        GzauTfdg72Uy-cRt8HUize16ul6Ss6Nn4hUHYoo_Bkg
      Type:         tls-alpn-01
      URL:          https://acme-staging-v02.api.letsencrypt.org/acme/chall-v3/4165525874/UVR7VA
    Identifier:     ingress-nginx.benjaminapetersen.me
    Initial State:  pending
    URL:            https://acme-staging-v02.api.letsencrypt.org/acme/authz-v3/4165525874
    Wildcard:       false
  Failure Time:     2022-10-31T19:46:14Z
  Finalize URL:     https://acme-staging-v02.api.letsencrypt.org/acme/finalize/74408214/4893805574
  State:            invalid
  URL:              https://acme-staging-v02.api.letsencrypt.org/acme/order/74408214/4893805574
# I get this challenge is important :)
Events:
  Type    Reason   Age    From          Message
  ----    ------   ----   ----          -------
  Normal  Created  4m51s  cert-manager  Created Challenge resource "quickstart-example-tls-6gzc2-2766096929-947999990" for domain "ingress-nginx.benjaminapetersen.me"
```
Lets take a look at that `Spec.Request` as it seems interesting:

```bash
# dump the CSR to a file, it appears base64 encoded
kubectl get order quickstart-example-tls-6gzc2-2766096929 -o jsonpath="{.spec.request}" | base64 --decode > order.quickstart-example-tls.csr.yaml
```
Yup, looks like this is what it is:

```text
-----BEGIN CERTIFICATE REQUEST-----
MIICkjCCAXoCAQAwADCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMMz
MAa3KtkEohc14XKt7XcZ5i6BX2rwhwqopD502LVEoiLQSW7duVsc5QVtXcs9RHPv
5bzUY2XBR21IFzQu0w5pThc1C10TTZGDBCxYVV3AQ9Hpi5XHNuLDJjxJOlYYycU6
Q/Xg5h4sD2pxPhL0+8LzG/s79A0leQnIT1LL+3p/U7FweYFLz+Co9fs9oHhPDvWQ
oZ1ccc5wHY1ANEWaC7xB0zufIv0agvD73h2EGue96rDMVy27OfnZZsbqIj0iOOz2
dG3oQhXkpKlEV088IrqxIPMnFYrevTjzHKSmmkqOv1axJ2T3FTher2WkjCGUYCVs
c0xjs2PxulE7CoCniXcCAwEAAaBNMEsGCSqGSIb3DQEJDjE+MDwwLQYDVR0RBCYw
JIIiaW5ncmVzcy1uZ2lueC5iZW5qYW1pbmFwZXRlcnNlbi5tZTALBgNVHQ8EBAMC
BaAwDQYJKoZIhvcNAQELBQADggEBAKCIJTuUJPN183uFRCoFIMhu6qIiFAMBnc1O
BJA1B1T0NqDNV3A5CIfk6jn9+LJK05eyR6XQRJzRa+2ajLFhZflczIHs2jcsqWEo
0pM4G69bcQBf/38kjDMPWtOrKJm+yT7b2Inq7Z2yB5isComFdjBhxMSufiiSsn7R
j9TUisT7q3owgQXUZFA4iSTIi4PaLusQbzwprwY1Yd4O8Rq+vXQsj6l0f4RSGweg
8Akvuy5idjxLZwcNoFYhTDs0sd790DCoBrZtvQKxJGPNFMxxEyhS1GAgt0WChoiH
OBK+z7KCK25g1bxkS4JUGHzobt0xOl5BycIuDT6W1OdOzfHahzo=
-----END CERTIFICATE REQUEST-----
```

So time to inspect it to see what the contents have to say:

```bash
# and inspect the CSR
openssl req -in order.quickstart-example-tls.csr.yaml -text -noout
```

```yaml
Certificate Request:
    Data:
        Version: 1 (0x0)
        Subject:
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (2048 bit)
                Modulus:
                    00:c3:33:30:06:b7:2a:d9:04:a2:17:35:e1:72:ad:
                    ed:77:19:e6:2e:81:5f:6a:f0:87:0a:a8:a4:3e:74:
                    d8:b5:44:a2:22:d0:49:6e:dd:b9:5b:1c:e5:05:6d:
                    5d:cb:3d:44:73:ef:e5:bc:d4:63:65:c1:47:6d:48:
                    17:34:2e:d3:0e:69:4e:17:35:0b:5d:13:4d:91:83:
                    04:2c:58:55:5d:c0:43:d1:e9:8b:95:c7:36:e2:c3:
                    26:3c:49:3a:56:18:c9:c5:3a:43:f5:e0:e6:1e:2c:
                    0f:6a:71:3e:12:f4:fb:c2:f3:1b:fb:3b:f4:0d:25:
                    79:09:c8:4f:52:cb:fb:7a:7f:53:b1:70:79:81:4b:
                    cf:e0:a8:f5:fb:3d:a0:78:4f:0e:f5:90:a1:9d:5c:
                    71:ce:70:1d:8d:40:34:45:9a:0b:bc:41:d3:3b:9f:
                    22:fd:1a:82:f0:fb:de:1d:84:1a:e7:bd:ea:b0:cc:
                    57:2d:bb:39:f9:d9:66:c6:ea:22:3d:22:38:ec:f6:
                    74:6d:e8:42:15:e4:a4:a9:44:57:4f:3c:22:ba:b1:
                    20:f3:27:15:8a:de:bd:38:f3:1c:a4:a6:9a:4a:8e:
                    bf:56:b1:27:64:f7:15:38:5e:af:65:a4:8c:21:94:
                    60:25:6c:73:4c:63:b3:63:f1:ba:51:3b:0a:80:a7:
                    89:77
                Exponent: 65537 (0x10001)
        Attributes:
        Requested Extensions:
            # yeah we want for this. are we proving we own it? is that our error?
            # not familiar enough with the bits of this process
            X509v3 Subject Alternative Name:
                DNS:ingress-nginx.benjaminapetersen.me
            X509v3 Key Usage:
                Digital Signature, Key Encipherment
    Signature Algorithm: sha256WithRSAEncryption
         a0:88:25:3b:94:24:f3:75:f3:7b:85:44:2a:05:20:c8:6e:ea:
         a2:22:14:03:01:9d:cd:4e:04:90:35:07:54:f4:36:a0:cd:57:
         70:39:08:87:e4:ea:39:fd:f8:b2:4a:d3:97:b2:47:a5:d0:44:
         9c:d1:6b:ed:9a:8c:b1:61:65:f9:5c:cc:81:ec:da:37:2c:a9:
         61:28:d2:93:38:1b:af:5b:71:00:5f:ff:7f:24:8c:33:0f:5a:
         d3:ab:28:99:be:c9:3e:db:d8:89:ea:ed:9d:b2:07:98:ac:0a:
         89:85:76:30:61:c4:c4:ae:7e:28:92:b2:7e:d1:8f:d4:d4:8a:
         c4:fb:ab:7a:30:81:05:d4:64:50:38:89:24:c8:8b:83:da:2e:
         eb:10:6f:3c:29:af:06:35:61:de:0e:f1:1a:be:bd:74:2c:8f:
         a9:74:7f:84:52:1b:07:a0:f0:09:2f:bb:2e:62:76:3c:4b:67:
         07:0d:a0:56:21:4c:3b:34:b1:de:fd:d0:30:a8:06:b6:6d:bd:
         02:b1:24:63:cd:14:cc:71:13:28:52:d4:60:20:b7:45:82:86:
         88:87:38:12:be:cf:b2:82:2b:6e:60:d5:bc:64:4b:82:54:18:
         7c:e8:6e:dd:31:3a:5e:41:c9:c2:2e:0d:3e:96:d4:e7:4e:cd:
         f1:da:87:3a
```

So there must be `Challenge` resources for this `Order`:

```bash
kubectl get challenge
# NAME                                                STATE     DOMAIN                               AGE
# quickstart-example-tls-6gzc2-2766096929-947999990   invalid   ingress-nginx.benjaminapetersen.me   13m
```

Lets investigate this sucker:

```yaml
# kubectl describe challenge quickstart-example-tls-6gzc2-2766096929-947999990
Name:         quickstart-example-tls-6gzc2-2766096929-947999990
Namespace:    default
Labels:       <none>
Annotations:  <none>
API Version:  acme.cert-manager.io/v1
Kind:         Challenge
Metadata:
  Creation Timestamp:  2022-10-31T19:45:51Z
  Generation:  1
  Owner References:
    API Version:           acme.cert-manager.io/v1
    Block Owner Deletion:  true
    Controller:            true
    Kind:                  Order
    Name:                  quickstart-example-tls-6gzc2-2766096929
    UID:                   4b8fdc80-95d7-4294-8f14-09d28a4118f8
  Resource Version:        118978
  UID:                     53496aeb-7b56-4a2f-b170-f63b34ccd641
Spec:
  # this appears to be our staging url note the authz-v3
  Authorization URL:  https://acme-staging-v02.api.letsencrypt.org/acme/authz-v3/4165525874
  # tis indeed our domain
  Dns Name:           ingress-nginx.benjaminapetersen.me
  Issuer Ref:
    Group:  cert-manager.io
    Kind:   Issuer
    # aha, again
    Name:   letsencrypt-staging
  # dunno, assume this is generated somewhere
  Key:      GzauTfdg72Uy-cRt8HUize16ul6Ss6Nn4hUHYoo_Bkg.KTNZwvDyC67N-Vci__Gt8a6wyM2QAP3rrRDYN0sovFM
  # we did ask for this solver, I believe
  Solver:
    http01:
      Ingress:
        Class:  nginx
  Token:        GzauTfdg72Uy-cRt8HUize16ul6Ss6Nn4hUHYoo_Bkg
  Type:         HTTP-01
  # similar, but note the chall-v3
  URL:          https://acme-staging-v02.api.letsencrypt.org/acme/chall-v3/4165525874/-PeDkg
  Wildcard:     false
# and this is why we are sad. forever.
Status:
  Presented:   false
  Processing:  false
  # AHA! but this is a red herring
  # we are setting an IPv4 IP address behind our domain ingress-nginx.benjaminapetersen.me
  # so why is it complaining about IPv6?
  # because LetsEncrypt always starts with IPv6. It didn't find a valid address.
  # then it likely fell back to IPv4. Also didn't find a valid.
  # why?
  # because our IP address is PRIVATE.
  # LetsEncrypt needs a publicly routable IP address to issue a certificate.
  Reason:      Error accepting authorization: acme: authorization error for ingress-nginx.benjaminapetersen.me: 400 urn:ietf:params:acme:error:dns: no valid A records found for ingress-nginx.benjaminapetersen.me; no valid AAAA records found for ingress-nginx.benjaminapetersen.me
  State:       invalid
Events:
  Type     Reason     Age   From          Message
  ----     ------     ----  ----          -------
  Normal   Started    14m   cert-manager  Challenge scheduled for processing
  Normal   Presented  14m   cert-manager  Presented challenge using HTTP-01 challenge mechanism
  Warning  Failed     14m   cert-manager  Accepting challenge authorization failed: acme: authorization error for ingress-nginx.benjaminapetersen.me: 400 urn:ietf:params:acme:error:dns: no valid A records found for ingress-nginx.benjaminapetersen.me; no valid AAAA records found for ingress-nginx.benjaminapetersen.me

```

So, we have an `IPv6` problem as the `AAAA` record is not found.
But we know this, we only config'd an `IPv4` IP, right?  on our Ingress?
Lets see if LetsEncrypt has any helps:

- https://letsencrypt.org/docs/ipv6-support/ eh, its some info. but not a solution.
- https://letsdebug.net/ `ingress-nginx.benjaminapetersen.me` LetsDebug! This is fun! And useful.

And indeed `letsdebug.net` is useful:

```yaml
ReservedAddress
FATAL
A private, inaccessible, IANA/IETF-reserved IP address was found for ingress-nginx.benjaminapetersen.me. Let's Encrypt will always fail HTTP validation for any domain that is pointing to an address that is not routable on the internet. You should either remove this address and replace it with a public one or use the DNS validation method instead.
10.191.196.251
```
So we need a public IP lol. Welp.


##################################################################################################################
##
###### ONCE I AM DONE WITH THIS ###### - TO DO MY ACTUAL WORK: https://jira.eng.vmware.com/browse/TKG-14442
##
##################################################################################################################

This seems to be likely realted to the bug that was reported:
- LetsEncrypt custom cluster issuer not working: https://jira.eng.vmware.com/browse/TKG-14442
- [Issues with Issuers may want to see the Troubleshooting ACME Certificates guide](https://cert-manager.io/docs/faq/acme/)

# Ingress

TechWorld With Nana: https://www.youtube.com/watch?v=80Ew_fsV4rM&t=37s

External Service vs Ingress

How to deploy a pod with a UI that needs to be accessed via browser?
You could use an `external service`.
    Access the app through `http://<service-ip-address>:<app-port>`
    Example: `http://124.89.191.2:35010`
But this is a dev setup only. Or a toy.
You need HTTPS.

There is a kubernetes component called `ingress`.

You need this:
- pod my-app (accessible by internal service)
- service my-app (internal service, IP is not accessible externally)
- ingress my-app (redirects to internal service)

```yaml
# what is an external service?
# type: LoadBalancer will provide an external IP address and make it an externally accessible service
# but this isn't great.
# why not?
#   why can't we use this external service + a DNS entry in something like GoDaddy DNS to make this work?
#   this is doable, but means the IP:PORT is ALSO accessible externally.
#      you don't want anther entrypoint. This is not best practice.
apiVersion: v1
kind: Service
metadata:
    name: myapp-external-service
spec:
    selector:
        app: my-app
    type: LoadBalancer # aha! now its external EXTERNAL-IP. But this isn't great
    ports:
    - protocol: TCP
      port: 8080
      targetPort: 8080
      nodePort: 35010 # with the EXTERNAL-IP from LoadBalancer we also use nodePort
```

So how about an Ingress:

```yaml
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
    name: myapp-ingress
spec:
    rules:
    # there will be a mapping to a service configured here
    # this must be a valid domain name
    # map this domain name to:
    # - an IP address to a Node's IP address,which is an entrypoint to the cluster
    # - an external server with an IP address that then routes to the cluster
    - host: myapp.com  # requests to this host will be forwarded to an internal service
      # BUT we really want HTTPS!
      # HTTP is insecure. HOWEVER
      http: # this attribute does NOT correspond to http:// in browser. This is just protocol definition
        paths: # this is the URL path. anything after the domain name will be matched here
        - backend:
            serviceName: myapp-internal-service # this is the internal service
            servicePort: 8080
---
# an internal service instead
# how is it different than external service?
# - no NodePort
# - default type ClusterIP, NOT LoadBalancer
apiVersion: v1
kind: Service
metadata:
    name: myapp-internal-service
spec:
    selector:
        app: myapp
    # no type!
    ports:
    - protocol: TCP
      port: 8080
      targetPort: 8080
      # no nodePort!
```

So how to configure ingress in a cluster?

`Ingress` requires also an `implementation` which is an `Ingress Controller`.

Configuration is then:
- my-app pod
- my-app service
- my-app ingress
- Ingress Controller Pod `processes the ingress rules`

## What is Ingress Controller?

Evaluates all the rules you have defined in your cluster and manage the redirections.
It is the entrypoint to the cluster for all requests to the domain or subdomain rules you configure.
It evaluates all rules. There may be many ingress components in the cluster.
It will evaluate which forwarding rule applies to the request.

`There are many third-party ingress controllers`!

The one from Kubernetes itself is `K8s Nginx Ingress Controller`.

## But what about a cloud service provider?

- Google Cloud
- AWS
- Azure
- Linode
- etc

These ahve OOTB Kubernetes solutions, or may have their own virtualized load balancer service.

Maybe:
- Cloud Load Balancer (implemented by cloud provider)
  - Request from Browser will hit LoadBalancer
    - LoadBalancer will redirect to Ingress Controller
      - `advantage`: you don't have to implement loadbalancer yourself!
        - LoadBalancer is likely up and running immediately

But what about Bare Metal Environment?
- Configure some kind of entrypoint to your kubernetes cluster yourself
  - This is a lot more work
- generally this is:
  - either inside the cluster
  - outside the cluster as a separate server
- and it needs to be
  - an external proxy servier
    - takes the role of the laodbalancer entrypoint to your cluster
    - gets a public IP address
    - open ports for requests to be accepted
    - proxy server acts as entrypoint to cluster
    - this is the only one accessible externally
    - NO server in the kbuernetes cluster would have a publically accessible IP address
- once the server has routing configured, it will send requests to kubernetes cluster
  - internal kubernetes service routing takes over from here

## Ingress Controller in Minikube

- Install Ingress Controller in Minikube

```bash
# automatically configures / starts the kubernetes nginx ingress controller
# handy this is offered out of the box!
minikube addons enable ingress
# then look for it
kubectl get pod -n kube-system
# nginx-ingress-controller-12345
```

Then you can create ingress rules for the controller to evaluate

## Install Kubernetes Dashboard & use with ingress rules

https://github.com/kubernetes/dashboard

```bash
# provided by the github repo
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
```
Note that since I am running on `TKG` provided cluster I may have to:

```bash
kubectl get all -n kubernetes-dashboard
```
To swap out the dockerhub proxy:
```yaml
        #image: kubernetesui/dashboard:v2.7.0
        image: harbor-repo.vmware.com/dockerhub-proxy-cache/kubernetesui/dashboard:v2.7.0
```
and then watch for a happy `Running` pod:

```bash
kubectl get pods -n kubernetes-dashboard -w
# NAME                                         READY   STATUS              RESTARTS   AGE
# dashboard-metrics-scraper-6f669b9c9b-wrs45   0/1     ImagePullBackOff    0          20m
# kubernetes-dashboard-54dbfcd689-k8n5l        0/1     ContainerCreating   0          23s
# kubernetes-dashboard-758765f476-5jmqj        0/1     ImagePullBackOff    0          20m
# kubernetes-dashboard-54dbfcd689-k8n5l        1/1     Running             0          25s  # Aha!
```

Then lets make an ingress for kubernetes dashboard:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
    name: dashboard-ingress
    # same namespace as the service & other components
    namespace: kubernetes-dashboard
    # i think we need this annotation for nginx controller to provide the IP address...
    annotations:
        kubernetes.io/ingress.class: "nginx"
spec:
    rules:
    - host: kubernetes-dashboard.benjaminapetersen.me
      http:
        paths:
        - pathType: Prefix
          path: /
          backend:
            service:
              name: kubernetes-dashboard
              # port 443 is HTTPS... this does not work.
              # but port 80 does at least get to Nginx server, however it returns a 503
              port:
               number: 443
              # there is no service listening on 80 from the default deployment, so this might not work!
              # port:
              #   number: 80
```
at present, I see `503 Service Temporarily Unavailable` for my deployment via NGINX.
This is probably a good thing.  I'm at least getting to the Nginx LoadBalancer.
But now what?

Likely I need a TLS certificate.  Look at my `kuard` ingress in my [cert-manager ingress doc](./cert_manager/../ingress.md),
which did deploy correctly, and we can see:

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: kuard # (kubernetes up and running demo)
  annotations:
    kubernetes.io/ingress.class: "nginx"
spec:
  # aha! a TLS entry in this implies a secret that is used for the certificate.
  # I don't have that above. And setting to port: 80 doesn't seem to be correct
  # for kubernetes-dashboard, its not happy.
  tls:
  - hosts:
    - ingress-nginx.benjaminapetersen.me
    secretName: quickstart-example-tls
  rules:
  - host: ingress-nginx.benjaminapetersen.me
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

kubectl describe ingress to see about `default backend`:

```bash
kubectl describe ingress dashboard-ingress -n kubernetes-dashboard
# Name:             dashboard-ingress
# Labels:           <none>
# Namespace:        kubernetes-dashboard
# Address:          10.92.104.204
# Default backend:  default-http-backend:80 (<error: endpoints "default-http-backend" not found>)
# Rules:
#   Host                                       Path  Backends
#   ----                                       ----  --------
#   kubernetes-dashboard.benjaminapetersen.me
#                                              /   kubernetes-dashboard:443 (100.96.1.40:8443)
# Annotations:                                 kubernetes.io/ingress.class: nginx
# Events:
#   Type    Reason  Age                From                      Message
#   ----    ------  ----               ----                      -------
#   Normal  Sync    20m (x4 over 31m)  nginx-ingress-controller  Scheduled for sync
```
Can also look for endpoints, since this seems to matter for `default backend`:

```bash
kubectl get endpoints -A | grep kubernetes-dashboard
# kubernetes-dashboard                dashboard-metrics-scraper                                                                               68m
# kubernetes-dashboard                kubernetes-dashboard                            100.96.1.40:8443                                        68m
```
and take a look at the single endpoints resource:

```yaml
# kubectl get endpoints kubernetes-dashboard -n kubernetes-dashboard -o yaml
apiVersion: v1
kind: Endpoints
metadata:
  creationTimestamp: "2022-10-26T16:43:14Z"
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
  resourceVersion: "779540"
  uid: a3a559ef-b2b3-4b32-ae14-ec8d43154db4
subsets:
# this is the config for the mapping.
- addresses:
  - ip: 100.96.1.40
    nodeName: tkg-mgmt-vc-md-0-7bd8b547f4-992ck
    targetRef:
      kind: Pod
      name: kubernetes-dashboard-54dbfcd689-k8n5l
      namespace: kubernetes-dashboard
      resourceVersion: "779532"
      uid: ac7dad81-39c2-4bbc-ac74-d8a01edfdef5
  ports:
  - port: 8443
    protocol: TCP
```

## But what is default-http-backend?

You simply need a service called `default-http-backend` and a `pod` that can serve up an error page or
whatever it is that you desire to serve to your users.  Then deploy with something like this to
replace the Nginx error pages (or whatever is currently handing your error pages. Assuming Kubernetes
Nginx Ingress Controller per the above choices):

```yaml
apiVersion: v1
kind: Service
metadata:
  name: default-http-backend
spec:
  selector:
    app: default-response-app # need a pod with this label to take the request
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8080
```

## Endpoints... quick refresh

Hands-On Kubernetes: https://www.youtube.com/watch?v=k9w6yhLfp9I

And endpoint is associated with a service
Services have selectors to match to pods
  - matching pods need to be healthy

As far as I undestand right now, `Ingress` maps to a `Service` which maps to `Pods` via the `Endpoints` resource,
which tracks the individual IP addresses assigned to a pod on a node, though the service has its own IP address.






.

# Kubernetes

`TODO`: copy over notes from lunch and learn bash

## A Kubernetes tutorial

Resources used:

- http://kubernetesbyexample.com/
- https://kubernetes.io/docs/reference/kubectl/cheatsheet/
    - this is really the place to start.
- O'Reilly Programming Kubernetes by Michael Hausenblas & Stefan Schimanski

## Useful Docs

- [k8s in general](https://github.com/figo/k8s-on-board)
- [k8s the hard way](https://github.com/kelseyhightower/kubernetes-the-hard-way)
- [k8s the kind way](https://kind.sigs.k8s.io/)
- [cluster-api fundamental](https://cluster-api.sigs.k8s.io/)


### Bash completion

Autocompletion is pretty handy.

```bash
source <(kubectl completion bash)
# permanently set it up in your bash profile
echo "source <(kubectl completion bash)" >> ~/.bashrc
```

### Kube Community (github.com/kubernetes/community)

- APIs
  - LEARN THIS STUFF:
    - [API Conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md)
    - [API Changes](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api_changes.md#making-a-singular-field-plural)

### Kube Docs (kubernetes.io/docs)

- APIs
  - [Deprecated API Migration Guide](https://kubernetes.io/docs/reference/using-api/deprecation-guide/)


## API

There are several ways to query the kubernetes API like a REST endpoint. This
will return JSON object responses as you would expect from other APIs.

Both of the following approaches will handle authentication and authorization
for you.

The proxy approach

```bash
# the first, is the proxy command + curl
# in one terminal window, open up a long runnin proxy
kubectl proxy --port 8080
# in a second terminal, query endpoints through the proxy with curl
curl http://127.0.0.1:8080/apis/batch/v1
```

The `--raw` flag

```bash
# the first, is to use the --raw flag
kubectl get --raw /apis/batch/v1
```

Generally, I would go with the `--raw` flag to eliminate the need to use the proxy.

Ask the cluster what kinds of resources are available via:

```bash
# get a table listing of each resource, kind, short name, etc
kubectl api-resources
# get a list of resource versions
kubectl api-versions
```

### Contexts

Borrowed from the [Kubernetes Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)


```bash
# view all context info, verbose
kubectl config view
# view just contexts as list
kubectl config get-contexts
# view your current context
kubectl config current-context
# set the context
kubectl config use-context <context-name>
# store a particular namespace as your "permanent" namespace for a context
kubectl config set-context --current --namespace=<some.preferred.namespace>
# add a basic-auth user to your context
kubectl config set-credentials kubeuser/foo.kubernetes.com --username=<name> --password=<pass>
# concat several .kube/config files to get a larger context
export KUBECONFIG=~/.kube/config:~/.kube/kubconfig2
# view the combined contexts (large output)
kubectl config view --flatten
# view the possible new contexts as list
kubectl config get-contexts
# CURRENT     NAME                CLUSTER                 AUTHINFO                  NAMESPACE
#             ben-tkg-worker      ben-tkgm-1.3.1-worker   ben-tkgm-worker-admin     ben-ns
# *           kind-2021-june-4    kind-2021-june-4        kind-2021-june-4
#             kind-pinniped       kind-pinniped           kind-pinniped             default
#             tkg-mgmt-vc@tkg-mgmt-vc tkg-mgmt-vc         tkg-mgmt-vc-admin
```

## Non-Resources

There are several things you can query in kubernetes that are not strictly considered resources.

### Healthz

```bash
kubectl get --raw '/healthz'
kubectl get --raw '/healthz?verbose=true'
```

## Resources

Most things you query in kube are resources.
Some basics of various resources collected into sections

Listing resource types is helpful to see what is available on the cluster

### api-resources

```bash
kubectl api-resources
kubectl api-resources --namespaced=true      # All namespaced resources
kubectl api-resources --namespaced=false     # All non-namespaced resources
kubectl api-resources -o name                # All resources with simple output (only the resource name)
kubectl api-resources -o wide                # All resources with expanded (aka "wide") output
kubectl api-resources --verbs=list,get       # All resources that support the "list" and "get" request verbs
kubectl api-resources --api-group=extensions # All resources in the "extensions" API group
```

### version

```bash
kubectl version
# JSON output by default
# Client Version: version.Info{Major:"1", Minor:"21", GitVersion:"v1.21.2+vmware.1", GitCommit:"54e7e68e30dd3f9f7bb4f814c9d112f54f0fb273", GitTreeState:"clean", BuildDate:"2021-06-28T22:17:36Z", GoVersion:"go1.16.5", Compiler:"gc", Platform:"linux/amd64"}
kubectl version -o yaml
# YAML version is more human friendly
# clientVersion:
#   buildDate: "2021-06-28T22:17:36Z"
#   compiler: gc
#   gitCommit: 54e7e68e30dd3f9f7bb4f814c9d112f54f0fb273
#   gitTreeState: clean
#   gitVersion: v1.21.2+vmware.1
#   goVersion: go1.16.5
#   major: "1"
#   minor: "21"
#   platform: linux/amd64
# serverVersion:
#   buildDate: "2021-07-21T18:25:07Z"
#   compiler: gc
#   gitCommit: b6b6fd4d596e3aaacba2264d2294c26c52bd0151
#   gitTreeState: clean
#   gitVersion: v1.21.2+vmware.1
#   goVersion: go1.16.5
#   major: "1"
#   minor: "21"
#   platform: linux/amd64
```


### APIService

```bash
# view api services
kubectl get apiservice
# limit to local
get apiservice | grep -v Local
```

### Pods

A pod is the basic unit of deployment in Kubernetes. It can be thought of
as a container of containers. It is a set of containers (though often only one)
that share a network & mount namespace. Each of the containers within a pod
are scheduled on the same node.

```bash
# run an pod using an image
kubectl run nginx --image=nginx
# set number of replicase
kubectl run nginx --image=nginx --replicas=5
# dry run. print output without actually creating API objects
kubectl run nginx --image=nginx --dry-run
# just like docker, you can run it in the foreground:
kubectl run -i -t nginx
# run a pod using an image from dockerhub (or other container registry)
kubectl run foo --image=<username>/<imagename>:<tag>
# specify a port
kubectl run foo nginx --image=nginx --port=9876
# see your running pods
kubectl get pods
# see details about a pod
kubectl describe pod <pod-name>
# get the IP of a pod
kubectl describe pod <pod-name> | grep IP:
# make a request to the pod via the IP returned in the above line
curl <ip.address.from.above>:<port>/info
curl 172.1.1.1:9876/info
# since kubectl creates a deployment, you must delete the
# deployment to get rid of the pod
kubectl delete deployment nginx
# you can also create a pod from a config file (yaml)
kubectl create -f https://raw.githubusercontent.com/some/file.yaml
# this file will create a service with two containers
kubectl create -f https://raw.githubusercontent.com/mhausenblas/kbe/master/specs/pods/pod.yaml
# exec into the containers
kubectl exec twocontainers -c shell -i -t -- bash
# from inside, can: curl localhost:9876/info
# create a pod with some constraints (resource limits)
kubectl create -f https://raw.githubusercontent.com/mhausenblas/kbe/master/specs/pods/constraint-pod.yaml
kubectl describe pod constraintpod
```

### Nodes

```bash
# get nodes
kubectl get nodes
# get complete spec w/node names
kubectl get nodes -o json | jq '.items[].spec'
# list tains per each node
kubectl get nodes -o json | jq '.items[].spec.taints'
# list node names + taints
kubectl get nodes -o json | jq ".items[]|{name:.metadata.name, taints:.spec.taints}"
```

Add and remove taints

```bash
# add
kubectl taint nodes w0-control-plane-qk5b8 key1=value1:NoSchedule
# remove
kubectl taint nodes w0-control-plane-qk5b8 key1=value1:NoSchedule-
```


### Secrets

```bash
# view secrets in a namespace
kubectl get secrets -n tkg-system
# view a spectific secret
kubectl get secret pinniped-data-values -n tkg-system
# view the secret as yaml
kubectl get secret pinniped-data-values -n tkg-system -o yaml
# view the contents of the secret, decoded
kubectl get secret pinniped-data-values -n tkg-system \
    -o jsonpath={.data.values\\.yaml} | base64 -d
# create a secret from a file
kubectl create secret generic registry-creds \
   --from-file=.dockerconfigjson=/home/you/.docker/config.json \
   --type=kubernetes.io/dockerconfigjson
# create a secret from cli inputs
kubectl create secret docker-registry docker-registry-cred \
    --docker-server=projects.registry.vmware.com \
    --docker-username=username \
    --docker-password=omg$secret@stuff \
    --docker-email=username@somewhere.com
# view the secret contents, using base64 to decode values
kubectl get secret docker-registry-cred -n default \
   --output="jsonpath={.data.\.dockerconfigjson}" | base64 --decode
```


## CLI

### kubectl

```bash
# either here or above, tbd.
```


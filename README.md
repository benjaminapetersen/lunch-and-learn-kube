# A Kubernetes tutorial

Resources used:

- http://kubernetesbyexample.com/
- https://kubernetes.io/docs/reference/kubectl/cheatsheet/
- O'Reilly Programming Kubernetes by Michael Hausenblas & Stefan Schimanski

## Bash completion

Autocompletion is pretty handy.

```bash
source <(kubectl completion bash)
# permanently set it up in your bash profile
echo "source <(kubectl completion bash)" >> ~/.bashrc
```

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

## Resources

Some basics of various resources collected into sections

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

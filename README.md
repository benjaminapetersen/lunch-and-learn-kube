# A Kubernetes tutorial

Resources used:

- http://kubernetesbyexample.com/





# Pods

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

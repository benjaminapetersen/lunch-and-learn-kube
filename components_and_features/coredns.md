# Hands-On Kubernetes - CoreDNS & DNS Resolution
https://www.youtube.com/watch?v=OKnOc4I-7sA

# Kubernetes CoreDNS
	`kubectl get pods -n tube-system`
	  coredns-7890-abcd
	`kubectl get service -n kube-system`
	  kube-dns  ClusterIP 10.96.0.10.  <none> 53/UDP,53/TCP  77d
	`kubectl describe service kube-dns`
	  only one endpoint....
So from this you can exec into any pod
	`kubectl exec -it <podname> -- sh`
Then
	`ls /etc`
	`cat /etc/resolv.conf            # this file is present in ever pod`
		nameserver 10.96.0.10   # refers to kube-dns
Then
 	`curl http://customer/api/v1/customers`
        # if you have "customer" in the file, it will use instead of IP address
Then
	`nslookup customer`
	# if you have an entry for "customer" it will do a lookup
	# and will resolve to customer.default.svc.cluster.local (the long name)
	# which is equivalent to the short name
	`curl http://customer.default.svc.cluster.local/api/v1/customer`
Then
	# get all services
	`kubectl get svc`
	can pair up the IPs from services to...
Then
	`cat /etc/resolv.conf`
	# again, this is present in every single pod
	# this is how DNS resolution is possible
Then
	# you may be able to exec into other pods and look at resolv.conf
	`kubectl exec -ti dnsutils -- cat /etc/resolv.conf`

## resolve.conf
https://en.wikipedia.org/wiki/Resolv.conf
- resolv.conf is the name of a computer file used in various operating systems to configure the system's Domain Name System (DNS) resolver
- The file is a plain-text file usually created by the network administrator or by applications that manage the configuration tasks of the system. The resolvconf program is one such program on FreeBSD or other Unix machines which manages the resolv.conf file.

## CoreDNS ConfigMap

As a YAML document:

```bash
kubectl get cm coredns -n kube-system -o yaml
```
```yaml
apiVersion: v1
data:
# the Corefile within is represented as a
Corefile: |
    .:53 {
        errors
        health {
        lameduck 5s
        }
        ready
        kubernetes cluster.local in-addr.arpa ip6.arpa {
        pods insecure
        fallthrough in-addr.arpa ip6.arpa
        ttl 30
        }
        prometheus :9153
        forward . /etc/resolv.conf {
        max_concurrent 1000
        }
        cache 30
        loop
        reload
        loadbalance
    }
kind: ConfigMap
metadata:
creationTimestamp: "2022-05-23T16:49:31Z"
name: coredns
namespace: kube-system
resourceVersion: "311"
uid: 9284e9c0-5a1c-4b25-999c-bef6c6224bb7
```

As a JSON document:

```bash
kubectl get cm coredns -n kube-system -o json
```
```json
{
    "apiVersion": "v1",
    "data": {
        // The Corefile embedded within is represented as plaintext.
        // This is probably not the easiest to work with...
        "Corefile": ".:53 {\n    errors\n    health {\n       lameduck 5s\n    }\n    ready\n    kubernetes cluster.local in-addr.arpa ip6.arpa {\n       pods insecure\n       fallthrough in-addr.arpa ip6.arpa\n       ttl 30\n    }\n    prometheus :9153\n    forward . /etc/resolv.conf {\n       max_concurrent 1000\n    }\n    cache 30\n    loop\n    reload\n    loadbalance\n}\n"
    },
    "kind": "ConfigMap",
    "metadata": {
        "creationTimestamp": "2022-05-23T16:49:31Z",
        "name": "coredns",
        "namespace": "kube-system",
        "resourceVersion": "311",
        "uid": "9284e9c0-5a1c-4b25-999c-bef6c6224bb7"
    }
}
```

If you were to edit the YAML you may do something like this:

```yaml
apiVersion: v1
data:

Corefile: |
    .:53 {
        errors
        health {
        lameduck 5s
        }
        ready
        kubernetes cluster.local in-addr.arpa ip6.arpa {
        pods insecure
        fallthrough in-addr.arpa ip6.arpa
        ttl 30
        }
        prometheus :9153
        ########################
        # adding some stuff here
        # so this whole file seems to be pretty standard format
        # its possible one could just replace the whole block with something new?
        ########################
        hosts custom.hosts akeesler-cluster.dev {  # I added this line!!
            <LoadBalancer IP> akeesler-cluster.dev # I added this line!!
            fallthrough                            # I added this line!!
        }                                          # I added this line!!
        forward . /etc/resolv.conf {
        max_concurrent 1000
        }
        cache 30
        loop
        reload
        loadbalance
    }
kind: ConfigMap
metadata:
creationTimestamp: "2022-05-23T16:49:31Z"
name: coredns
namespace: kube-system
resourceVersion: "311"
uid: 9284e9c0-5a1c-4b25-999c-bef6c6224bb7
```

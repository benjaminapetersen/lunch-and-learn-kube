# Pod Security

`The Hitchhiker's Guide to Pod Security - Lachlan Evenson, Microsoft`

https://www.youtube.com/watch?v=gcz5VsvOYmI

- `Pod Security Policy` is `Deprecated`
  - https://kubernetes.io/docs/concepts/security/pod-security-policy/
- `Pod Security Admission`
  - https://kubernetes.io/docs/concepts/security/pod-security-admission/


## PSA - Pod Security Admission

Pod Security Standards Levels:
- `priviledged` - open and unrestricted
- `baseline` - covers known privilege escalations while minimizing restrictions
- `restricted` - highly restricted, hardened against known and unkonwn priviledge escalations. may cause compatibilty issues

Some `spec` fields restricted by these policy include:

```yaml
spec.securityContext.sysctls
spec.nostHetwork
spec.volumes[*].hostPath
spec.containers[*].secuirtyContext.priviledged
```

Modes

```yaml
enforce: reject pods taht violate the policy
audit: record violations as annotations in the audit logs but still allow pod
warn: send warning message back to user, but still allow pod
```

Configuring Pod Security

`policies` are applied to `namespace` via `labels`.

The labels are:

```yaml
pod-security.kubernetes.io/<MODE>: <LEVEL>              # required to enable pod security
pod-security.kubernetes.io/<MODE>-version: <VERSION>    # optional, defaults to "latest"
# <OPTIONAL> cluster-wide policies and exemptions using the AdmissionConfiguration*.
# The possible <MODES> are enforce, audit, warn
```

A specific version can be supplid for each enforcement mode.

A typical use for `warn` is to get ready for a future change where you want to enforce a different policy.
`enforce` is only applied to `pod` resource.  `warn` and `audie` are also applied to `workload` resources.



FOR TKG:

- remove the PSP overlay from tkg-packages
- add a new PSA overlay with
  - `annotations` applied to each of the `namespaces` created by `pinniped-package`


## Demo 1: Is Pod Security Enabled on Your Cluster?

Confirm Pod Security is `Enabled`.  Can't use it if its not enabled!

In order to run this, spin up a `kind` cluster and perform the following actions:

```bash
# within a kind cluster

# confirm kubernets version and access to the API server
kubectl version
kubectl get nodes

# check the API server flags
kubectl -n kube-system exec kube-apiserver-kind-control-plan -it -- kube-apiserver -h | grep 'default enabled ones'
# output should be a long list of things such as:
# AlwaysDeny
# AlywaysPullImages
# CertificateSigning
# DefaultIngressClass
# NamespaceExists
# PodSecurity <----- this one, you need this one, else you can't use PSA
# etc.
# the order of plugins in this flag does not matter

# make a namespace to tinker with
kubectl create namespace verify-pod-security

# label the namespace to enforce the restricted Pod Security Standard
kubectl label namespace verify-pod-security pod-security.kubernetes.io/enforce=restricted
kubectl get ns verify-pod-security -o  yaml

# try to run a priviledge container to see it fail
kubectl -n verify-pod-security run test --dry-run=server --image=busybox --privileged
# Error from server (Forbidden): pods "test" is forbidden: violates PodSecurity "restricted:latest": privileged (container "test" must not set securityContext.privileged=true), allowPrivilegeEscalation != false (container "test" must set securityContext.allowPrivilegeEscalation=false), unrestricted capabilities (container "test" must set securityContext.capabilities.drop=["ALL"]), runAsNonRoot != true (pod or container "test" must set securityContext.runAsNonRoot=true), seccompProfile (pod or container "test" must set securityContext.seccompProfile.type to "RuntimeDefault" or "Localhost")

# cleanup
kubectl delete ns verify-pod-security
```

## Demo 2: Priviledged Level and Workload

Create some policy
Try to create a workload that violates the policy
Bring the workload into alignment with the policy

Use a `kind` cluster again :)
Or don't, but its up to you to figure out whats up then.

```bash
# use kind

kubectl create namespace verify-pod-security

kubectl label --overwrite ns verify-pod-security \
    pod-security.kubernetes.io/enforce=restricted \
    pod-security.kubernetes.io/audit=restricted

kubectl get ns verify-pod-security
kubectl get ns verify-pod-security -o yaml
# apiVersion: v1
# kind: Namespace
# metadata:
#   creationTimestamp: "2022-08-24T18:50:02Z"
#   labels:
#     kubernetes.io/metadata.name: verify-pod-security
#     pod-security.kubernetes.io/audit: restricted
#     pod-security.kubernetes.io/enforce: restricted
#   name: verify-pod-security
#   resourceVersion: "906753"
#   uid: 71d7367f-4278-4bfa-b7ed-e8e316913cbe
# spec:
#   finalizers:
#   - kubernetes
# status:
#   phase: Active

# deploy a priviledged workload to the namespace

cat <<EOF | kubectl -n verify-pod-security apply -f -
apiVersion: v1
kind: Pod
metadata:
    name: busybox-priviledged
    namespace: verify-pod-security
spec:
    containers:
    - name: busybox
      image: busybox
      args:
      - sleep
      - "10000000"
      securityContext:
        allowPrivilegeEscalation: true
EOF
# Error from server (Forbidden): error when creating "STDIN": pods "busybox-priviledged" is forbidden: violates PodSecurity
# "restricted:latest": allowPrivilegeEscalation != false (container "busybox" must set securityContext.allowPrivilegeEscalation=false),
# unrestricted capabilities (container "busybox" must set securityContext.capabilities.drop=["ALL"]), runAsNonRoot != true (pod or
# container "busybox" must set securityContext.runAsNonRoot=true), seccompProfile (pod or container "busybox" must set
# securityContext.seccompProfile.type to "RuntimeDefault" or "Localhost")

# now, lets change the policy
kubectl label --overwrite ns verify-pod-security \
    pod-security.kubernetes.io/enforce=privileged \  # we allow privileged now, yikes!
    pod-security.kubernetes.io/warn=baseline \
    pod-security.kubernetes.io/audit=baseline

# now, re-run the exact same previous command:
cat <<EOF | kubectl -n verify-pod-security apply -f -
apiVersion: v1
kind: Pod
metadata:
    name: busybox-privileged
spec:
    containers:
    - name: busybox
      image: busybox
      args:
      - sleep
      - "10000000"
      securityContext:
        allowPrivilegeEscalation: true
EOF
# pod/busybox-priviledged created <------- WHOA!!!! it created the pod. but this is bad?

# cleanup... or don't, and do the next update of labels to see busybox get tagged
# kubectl -n verify-pod-security delete pod busybox-privileged

# now, lets enforce and audit on restricted
kubectl label --overwrite ns verify-pod-security \
    pod-security.kubernetes.io/enforce=restricted \
    pod-security.kubernetes.io/audit=restricted
# Warning: existing pods in namespace "verify-pod-security" violate the new PodSecurity enforce level "restricted:latest"
# Warning: busybox-privileged: allowPrivilegeEscalation != false, unrestricted capabilities, runAsNonRoot != true, seccompProfile
# namespace/verify-pod-security labeled


# and alter the busybox to add some additional capabilities
# the NET_BIND_SERVICE should be ok, though allowPrivilegeEscalation is prob still bad.
cat <<EOF | kubectl -n verify-pod-security apply -f -
apiVersion: v1
kind: Pod
metadata:
    name: busybox-privileged
    namespace: verify-pod-security
spec:
    containers:
    - name: busybox
      image: busybox
      args:
      - sleep
      - "10000000"
      securityContext:
        allowPrivilegeEscalation: true
        capabilities:
            add:
            - NET_BIND_SERVICE
EOF
# Error from server (Forbidden): error when creating "STDIN": pods "busybox-privileged" is forbidden: violates PodSecurity
# "restricted:latest": allowPrivilegeEscalation != false (container "busybox" must set securityContext.allowPrivilegeEscalation=false),
# unrestricted capabilities (container "busybox" must set securityContext.capabilities.drop=["ALL"]), runAsNonRoot != true
# (pod or container "busybox" must set securityContext.runAsNonRoot=true), seccompProfile (pod or container "busybox"
# must set securityContext.seccompProfile.type to "RuntimeDefault" or "Localhost")

# ok, so if we drop: ALL first, does it work?
cat <<EOF | kubectl -n verify-pod-security apply -f -
apiVersion: v1
kind: Pod
metadata:
    name: busybox-restricted-some
    namespace: verify-pod-security
spec:
    containers:
    - name: busybox
      image: busybox
      args:
      - sleep
      - "10000000"
      securityContext:
        allowPrivilegeEscalation: true
        capabilities:
            drop:
            - ALL
            add:
            - NET_BIND_SERVICE
EOF
# Error from server (Forbidden): error when creating "STDIN": pods "busybox-privileged" is forbidden: violates PodSecurity "restricted:latest": allowPrivilegeEscalation != false (container "busybox" must set securityContext.allowPrivilegeEscalation=false), runAsNonRoot != true (pod or container "busybox" must set securityContext.runAsNonRoot=true), seccompProfile (pod or container "busybox" must set securityContext.seccompProfile.type to "RuntimeDefault" or "Localhost")

# so if we update secuirtyContext a bunch more and restrict our pod down, we can get it to pass
cat <<EOF | kubectl -n verify-pod-security apply -f -
apiVersion: v1
kind: Pod
metadata:
    name: busybox-restricted-more
    namespace: verify-pod-security
spec:
    containers:
    - name: busybox
      image:  busybox
      args:
      - sleep
      - "10000000"
      securityContext:
        seccompProfile:
            type: RuntimeDefault
        runAsNonRoot: true
        allowPrivilegeEscalation: false
        capabilities:
            drop:
            - ALL
            add:
            - NET_BIND_SERVICE
EOF
# pod/busybox-privileged created

# but... its not ok yet?
kubectl -n verify-pod-security describe pod busybox-restricted

# even more restricted!
cat <<EOF | kubectl -n verify-pod-security apply -f -
apiVersion: v1
kind: Pod
metadata:
    name: busybox-restricted-even-more
    namespace: verify-pod-security
spec:
    securityContexts:
        runAsUser: 65534
    containers:
    - name: busybox
      image:  busybox
      args:
      - sleep
      - "10000000"
      securityContext:
        seccompProfile:
            type: RuntimeDefault
        runAsNonRoot: true
        allowPrivilegeEscalation: false
        capabilities:
            drop:
            - ALL
            add:
            - NET_BIND_SERVICE
EOF

# is it running now?
kubectl get pod busybox-restricted-even-more -n verify-pod-security

# check our metrics against the api server /metrics endpoint
kubectl get --raw /metrics | grep pod_security_evaluations_total
# HELP pod_security_evaluations_total [ALPHA] Number of policy evaluations that occurred, not counting ignored or exempt requests.
# TYPE pod_security_evaluations_total counter
pod_security_evaluations_total{decision="allow",mode="enforce",policy_level="privileged",policy_version="latest",request_operation="create",resource="pod",subresource=""} 66
pod_security_evaluations_total{decision="allow",mode="enforce",policy_level="privileged",policy_version="latest",request_operation="update",resource="pod",subresource=""} 1
pod_security_evaluations_total{decision="allow",mode="enforce",policy_level="restricted",policy_version="latest",request_operation="create",resource="pod",subresource=""} 2
pod_security_evaluations_total{decision="deny",mode="audit",policy_level="restricted",policy_version="latest",request_operation="create",resource="pod",subresource=""} 6
pod_security_evaluations_total{decision="deny",mode="enforce",policy_level="restricted",policy_version="latest",request_operation="create",resource="pod",subresource=""} 8
```



## What to do for Pinniped?

```bash
# for each pinniped namespace, we probably want a restricted default
# this is
# - pinniped-supervisor
# - pinniped-concierge
# - tanzu-system-auth
# - any others?
#   - kube-public: we don't own kube-public but we publish into it
#   - kube-system: we don't own kube-system, but we publish into it
# - which of these are configurable in the package?
kubectl label --overwrite ns <PINNIPED.NAMESPACE> \
    pod-security.kubernetes.io/enforce=restricted \
    pod-security.kubernetes.io/audit=restricted
```

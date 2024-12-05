# Controller Tools

https://github.com/kubernetes-sigs/controller-tools

# Controller-gen

https://github.com/kubernetes-sigs/controller-tools/blob/master/cmd/controller-gen/main.go

```bash
# in the pinniped-package dir for the pinniped-config controller
# we are borrowing controller-gen from the upper ./hack tools
# then we are:
# - reading ./controllers for annotations
# - naming the +rbac:roleName
# - defining output in 2 places
#     rbac to ./
#     everything else to stdout
"${TF_ROOT}/hack/tools/bin/controller-gen" paths=./controllers +rbac:roleName=pinniped-config-controller-manager output:rbac:dir=./ output:stdout
```
# Operator Pattern

- https://kubernetes.io/docs/concepts/extend-kubernetes/operator/
- [CNCF Operator White Paper](https://github.com/cncf/tag-app-delivery/blob/163962c4b1cd70d085107fc579e3e04c2e14d59c/operator-wg/whitepaper/Operator-WhitePaper_v1-0.md)

Operators are software extensions to Kubernetes that make use of
custom resources to manage applicatoins and their
components. Operators follow Kubernetes principles, notably the
control loop.

Operators are clients of the Kubernetes API that act as controllers
for a custom resource.

Operators might:
- deploy an application on demand
- take and resore backups of application current state
- handle upgrades of an application's code alongside
  related changes such as databse schemas or other config
- publis a service to applicatoins that don't support
  Kubernetes APIs to discover them
- simulate failure in all or part of your cluster to test
  its resilience
- choose a leader for distributed applicatoin without
  an internal member election process

There are a number of frameworks available for writing Operators:
- Operator Framework / Operator SDK
- Kubebuilder
- Carvel (alternative to controllers and operators)
- etc.

# Azure Workload Identity

- [Azure Workload Identity](https://learn.microsoft.com/en-us/entra/workload-id/workload-identities-overview)
- [Workload Identity FAQ](https://learn.microsoft.com/en-us/entra/workload-id/workload-identities-faqs)
- [What are workload identities?](https://learn.microsoft.com/en-us/entra/workload-id/workload-identities-overview)
- [How to... Configure an app to trust an external IDP](https://learn.microsoft.com/en-us/entra/workload-id/workload-identity-federation-create-trust?pivots=identity-wif-apps-methods-azp)

A workload identity is an identity you assign to a software workload to
authenticate and access other services and resources.

In Microsoft Entra, workload identities are applications, service principals, and managed identities.
- `Application`: An abstract entity or template defined by its application object.
  - The application object is a `GLOBAL` representation of the application 
    for use across all tenants.
  - The application object describes:
    - How tokens are issued
    - The resources that can be accessed
    - The actions the application can take
- `Service Principal`:
  - The service principal is a `LOCAL` representation, or an 
    `application` instance, of a global application object,
    within a specific tenant.
  - The `application` object is used as a template to:
    - Create a service principal object
      - in every tenant
      - where the application is used
    - The Service Principal Object defines
      - what the app can actually do
        - in a specific tenant
      - who can access the app
      - what resources the app can access 
- `Managed Identity`: 
  - A specific type of `Service Principal` that
    - Eliminates the need for developers to manage credentials


# CM: extension-apiserver-authentication

The `extension-apiserver-authentication` configmap holds the set of config that details which headers are used for things
such as defining UserID, UserName, Group, etc.

# https://github.com/kubernetes/kubernetes/issues/93699#issuecomment-885641517

```bash
# get the configmap that contains the configuration for the headers
kubectl get cm -n kube-system extension-apiserver-authentication -o json | jq '.data'
{
  # what is the network topology? why does this matter?
  #
  # how does the API server route traffic to aggregated API servers?
  # (note: you need trust to have identity, you don't need identity to have trust)
  #  - you need a trust anchor
  #  - you can just have a public key
  #  - you don't need x509 or a certificate authority...
  # policy is also separate
  #  - you have a policy of only trusting known good certificate authority
  #  - you have a policy of only trusting keys from good sources
  #  - policy is NOT cryptography
  # crypto gives you trust, humans do not.  crypto does the math.
  # yet humans do pass the keys to be provided as configuration to the system
  #  - its a configuration error to use a bad certificate or public key that is
  #    compromised by bad actors.
  #  - it is not a computation error.  the crypto still works, the math is still
  #    the same, you just need to provide correct config (keys and certs)
  # so
  # - there is one certificate in this bundle (so its a bundle of one)
  # - there are numerous trust anchors in kube (7+)
  # - in this case, only one CA has been provided to the system
  #   and every component is reusing the same CA.
  # - ClusterAPI uses KubeADM to create clusters.
  # so
  # - this is configuration for any component that wants to...
  # so Mo Q:
  #   if i am perfectly good at networking, what configuration do i need to setup networking to send date back and forth from a kube api server
  # A:
  "client-ca-file": "-----BEGIN CERTIFICATE-----\nMIIC6jCCAdKgAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl\ncm5ldGVzMB4XDTIxMTIwNzE3MDQwMVoXDTMxMTIwNTE3MDkwMVowFTETMBEGA1UE\nAxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANAy\ng+YwVS2EmXSORU++V4JTCOHoqm/KSWcMqsYtWcFNGLdjPkHBQfgFsHe3TGOMUUaI\nPQEBclEQeVxOpvmNAq0AyMMzei3S/Dx8Y/CSOM1ziKl+afJG2iHUcoTCcdPZEkPq\nUTqvPgVvbCMtBsvSKU7URL8URmF1Hp4qK/QB5t8tKu9vRA45xrk8DnpumZfFiqFX\nejJDKlv8u0SSFR0UYkGlLcU8T8wCrRTB7z08TclXd6cmWNCLj/elBnKQQw9ROc00\n5Q9G6+wCOYvoIPnJCYiDIkbMyHyATSszS4t9+A3XaV1sEYvgy3OXNon/XIO1HCGB\nwyJMW67Sq9+hszdJO68CAwEAAaNFMEMwDgYDVR0PAQH/BAQDAgKkMBIGA1UdEwEB\n/wQIMAYBAf8CAQAwHQYDVR0OBBYEFOGypyjgTJgmo20sx6GQAkmgJpVHMA0GCSqG\nSIb3DQEBCwUAA4IBAQCTptTgmPEbm2b1TdPpt1G48TestkUseIYKE90mCEXXqyRk\nDUDnTI2iBYgK+vq76PfQLlzCmKGaSON1tAWErvonPrkk/R2Ou53OvDQh/CPFYdxM\nM++U10AZSgJpZywa3xzoCsP3O2IJYQg3mUk46lLcs7BCnmmgf2OCgRiO7ZTnQ9/U\n1epdjeK+QCbhK/iuRiwGzyNNfswbwrvOY5lysSMC7ClSN183MCilBR1k2lNM3oIt\nJhFK3poUFbtzVcTgnPZ8OCE+eaagE3VdzakJFXomWt/pcEOTHCok8FaOS0dx2ZuL\nIC1tsRiD0bUd8rInaoJlE38RAKWl1JxamTcUu0W9\n-----END CERTIFICATE-----\n",

  # this is the CN on the certificate
  # THIS is the kubernetes API server's identity
  "requestheader-allowed-names": "[\"front-proxy-client\"]",

  # this is the certificate of the front-proxy-client
  # this is for mTLS with a CN restriction
  # (note the CA may be used for multiple things)
  "requestheader-client-ca-file": "-----BEGIN CERTIFICATE-----\nMIIC6jCCAdKgAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl\ncm5ldGVzMB4XDTIxMTIwNzE3MDQwMloXDTMxMTIwNTE3MDkwMlowFTETMBEGA1UE\nAxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKT/\nq4TJQbk4n94Ji/eUf5To9H7RAAPp/HV1md/zyQxKro7syUgTRh8FtY7pWNM3J3uJ\nQzqKdnKc15lnEaxcqwBq9xiDa3is6DgQlqVTfCfz5JVFFkfxkfe0pvgG3p1cpJmf\n1AY9k9OhCMxuCKvt7eNjDriLAswkA6CWm973yCszNeYiI9SP/sFDQYUuajEa6HrG\nLi4nxGyXeLMHSIakC4fmmwX6Mg5cHvhBqitG2wReaI7M/io+fNqMtHc3CR0+8lfh\nXxSJ/1E7FgKA/Qz3BfqTrejFxzs/h5qwXG+jI3h+7DqprKgWxg4pNU2hyqiee9ME\nrhZ5QSdxQuLqu+BA4kUCAwEAAaNFMEMwDgYDVR0PAQH/BAQDAgKkMBIGA1UdEwEB\n/wQIMAYBAf8CAQAwHQYDVR0OBBYEFH7VPEJfpNj8wXEcOtnSF+m44GnMMA0GCSqG\nSIb3DQEBCwUAA4IBAQBFALpETBMsxCfhQ5r+US0JKtdjxb/dTkDLZPAD15NZMuuc\n5BqZ9tuMtMm13jXRM6EUZsv9wz6xUnNv6vuUwWqqk2eZQc/GCCoKxP8CYrHxGnEa\nPfmcsT6Q1E1wE73xjYlcDdtMyfXfGA0TE92SCrK2E0bfGvBV7NFIwJQQ2QdiWUa4\nEdKpXqcK0iWcJVpSWB4H4Ku4g23ftw/lKk7uy4wzMhc4PFEt/Rm07yrAy3SJ2HUH\nJPx6RBhdjeW+9Y8qV43hr9NLDgITBpX02AeXq52i9AVaKFFJLDK1fwRHKi9ZCBOe\nQqdE8nOqw83xotzEd7L3I7rI+1rxc7S61E61ZCn4\n-----END CERTIFICATE-----\n",

  # all of the headers that have something to do with identity
  "requestheader-extra-headers-prefix": "[\"X-Remote-Extra-\"]",
  "requestheader-group-headers": "[\"X-Remote-Group\"]",
  "requestheader-username-headers": "[\"X-Remote-User\"]"
  # needs this, first one present on the request wins
  "requestheader-uid-headers": : "[\"X-Remote-Uid\",\"X-Remote-Shizzle\"]"
}
```

```bash
# this is a the underling call made for a who am i request
curl -k 'https://<KUBE_API_SERVER_URL>/apis/identity.concierge.pinniped.dev/v1alpha1/whoamirequests' -XPOST -d '{}' -H "Content-Type: application/json" -H "Authorization: Bearer SA_TOKEN"
```

This happens to be an aggregated API server request to the Concierge aggregated API server.

network path for response:

```bash
# this is a specific problem to aggregated API servers
# they do not pass all of the headers (uid)
# admission webhooks, authorization webhooks, etc (have user info in them encoded as part of the body of the request)
#   the user information exists fully, not partially, so this is irrelevant.
curl -> KUBE_API_SERVER_URL -[headers]> CONCIERGE_AGGREGATED_API_SERVER_URL
```

```bash
# TODO: why do we sign certificates with private keys?
# if (and only if) I have a mutual TLS connection with a server who's certificate was signed by the matching private key
# to this CA (from requestheader-client-ca-file header) file and the CN on the certifiate is front-proxy-client (from requestheader-allowed-names header)
# then blindly trust the headers configured by requestheader-* config (any of them, first one wins)
#
# So the concierge trusts the APIServer
# therefore it trusts the headers.
# it would not trust the request coming from ANYONE else
# therefore the headers would be meaningless
# more specifically, you cannot spoof these headers because YOU are not trusted.
#
# The request from the APIServer to the Concierge (but in GO code) would look something like this:
#
# -k means don't validate the certificate on the connection, encryption without authentication, no secruity :)
# CONCIERGE_AGGREGATED_API_SERVER_URL comes from the apiService (kubectl get apiservice) object
curl 'CONCIERGE_AGGREGATED_API_SERVER_URL/apis/identity.concierge.pinniped.dev/v1alpha1/whoamirequests'
  -XPOST -d '{}' -H "Content-Type: application/json"
  # -H "Authorization: Bearer SA_TOKEN" -> this gets deleted.  the api server did authentication and then REMOVED this header
  -H "X-Remote-Extra-: ..." # there are Extra headers, and a batch of them, but we will ignore for now.
  -H "X-Remote-Group: system:serviceaccounts"
  -H "X-Remote-Group: system:authenticated"
  -H "X-Remote-Group: system:serviceaccount:jungle"
  -H "X-Remote-User: system:serviceaccount:jungle:panda"  # panda SA in jungle namespace
  # -H "X-Remote-Uid" this doesn't exist yet :)
```

The concierge trusts the APIServer gave it correct information in headers, so it will build a JSON structure and return it.

This API should exist in the Kube API Server itself.  Write a KEP and move it. :)


k get apiservice v1alpha1.login.concierge.pinniped.dev -o json
# caBunndle
# - how can the kube API server talk to the concierge aggregated api server?
# via the service
# how do we get packets to the aggreated API server?
# - service are used by aggreated API servers
#   to provide networking, IP address
# - service are internal by default
# - available to the entire cluster
# - no tenancy boundary by default within a cluster
# - if that is the case, a server must be secure against arbitrary network traffic
# - it can't blindly trust headers
# - it can only trust headers over an mTLS connection with a trusted source
#
# if a pod running on the cluster with its own SA token makes a request to
# the concierge aggregated API server, say the health endpoint, with the
# token as a bearer token set as the authorization header?
# - the conciege does a tokenreview against hte API server
#   who is this? this SA token is a real entity?
# - the concierge goes to a trusted source with an token from an unknown
#   source to ask if it can be trusted.
# - now I am in a pod with a client certificate, presenting the cert to
#   the concierge over the service network.  does the approach used for
#   tokens work in this case?  can you do a token review?
#   - no.
#   - why?
#   - what is the diff betweek token based auth and certificate based auth?
#     - certificates are channel bound authentication
#       this is just the public key
#       you don't have the private key
#     - tokens are forwardable (replayable)
#       you can copy/paste these along
#       - always know who you are talking to!
#     - an actor, or the API server, cannot pass along a certificate
#       TLS is terminated.
#   - we ask the API server (from the concierge) for a CA bundle that we
#     can use to verify identities.
#   - all components have to agree on the method of certificate authentication
#
# - we have the front-proxy mechanism for the API server to assert any identyt to us in a safe way. it is 99% of hte time in front of us, cuz its the front proxy, it is in front of all aggregated API servers. this is what hte API service resource means.  the kube api server says it has APIs, but it often doesn't, it is often forwarding requests to other servers, but it has already done authentication for you.  THe API server can authenticate howver it wants, via several mechanisms. this is a black box to you.  THe API server itself supports front proxy.
# - in the case of somethign on the service networking wanting to communicate iwth the API server, for token based auth we simply forward the token to the API server and ask it to do the validation for us. we need some mechanism to do that with certificates.  Generally, you do this by configuring an aggregated api server with a CA bundle.  there may be one or many CAs in the cluster.

# how do you protect against replaying tokens?
#  add an audience to the token.
#  token says "i am service A, and this token is for service B"
#  therefore, this token is only for one use, from and to specifics.
#  it cannot be used for service C.
#  if you steal the token, you can only use it for the target (service B)
#
# this is why people prefer certificate based auth in many cases, because you cannot forward the identity, you cannot replay the certificate, you do not have the private key.
# Bearer toekns are called "bearer" beause they bare everything to the endpoint.
# but unfortunately certifcate auth is harder to use.
# browsers have poor support for it.
# etc.


# look up WebAuthN (vs steve gibson's squirrel)
# cfssl vs openssl
# - cfssl is more modern/less painful than openssl
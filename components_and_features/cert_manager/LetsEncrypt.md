# LetsEncrypt

Some debugging and troubleshooting for LetsEncrypt.

## Challenge Types

- https://letsencrypt.org/docs/challenge-types/
- `HTTP-01 challenge` `http://<YOUR_DOMAIN>/.well-known/acme-challenge/<TOKEN>`
- `DNS-01 challenge` `_acme-challenge.<YOUR_DOMAIN>`
- `TLS-SNI-01`
- `TLS-ALPN-01`


## Debugging

- `Let's Debug` https://letsdebug.net/
  - Example: https://letsdebug.net/ingress-nginx.benjaminapetersen.me/1245661
  - This is pretty great.

- `Orders` https://cert-manager.io/docs/troubleshooting/acme/#2-troubleshooting-orders

- `no valid A records found for ingress-nginx.benjaminapetersen.me; no valid AAAA records found for ingress-nginx.benjaminapetersen.me`
  - One that may be obvious, if your cluster IP addresses are private, then LetsEncrypt cannot hit  your
    A or AAAA records to validate that you own the domain (as your domain will be pointing at private IP addresses
    which are not routable on the public internet).

- `Certificate as secret does not exist`
  - https://github.com/cert-manager/cert-manager/issues/5086

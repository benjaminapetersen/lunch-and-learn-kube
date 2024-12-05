# 2022, Jan 25

We got the tests hooked up, it was missing a CA bundle in `test-cmd.sh` which meant it would not write headers because the
headers have no value if the CA bundle does not exist to provide trust.  So tweaking `test-cmd.sh` to pass in `ca.crt` instead
of the `tls.crt` is what enabled our output (`tls.crt` was not a `CA`, it was just a cert, it was therefore being checked/valudated
and found lacking)

IntelliJ
- shift+shift (double shift) to open the search  box

`make test-cmd WHAT=run_cluster_authentication_trust`

+++ command: run_cluster_authentication_trust_tests
+++ [0128 10:29:57] {
    "apiVersion": "v1",
    "data": {
        "client-ca-file": "-----BEGIN CERTIFICATE-----\nMIICpjCCAY4CCQCZBiNB23olFzANBgkqhkiG9w0BAQsFADAUMRIwEAYDVQQDDAkx\nMjcuMC4wLjEwIBcNMjAwNjE0MTk0OTM4WhgPMjI5NDAzMzAxOTQ5MzhaMBQxEjAQ\nBgNVBAMMCTEyNy4wLjAuMTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB\nAMxIjMd58IhiiyK4VjmuCWBUZksSs1CcQuo5HSpOqogVZ+vR5mdJDZ56Pw/NSM5c\nRqOB3cvjGrxYQe/lKvo9D3UmWLcRKtxdlWxCfPekioJ25/dhGOxtBQcjtp/TSqTM\ntxprwT4fvsVwiwaURFoCOivF4xjQFG0K1i3/m7CiMHODy67M1EfJDrM7Vv5XPIuJ\nVF8HhWBH2HiM25ak34XhxVTX8K97k6wO9OZ5GMqbYuVobTZrSRdiv8s95rkmik6P\njn0ePKqSz6cXNXgXqTl11WtsuoGgjOdB8j/noqTF3m3z17sSBqqG/xBFuSFoNceA\nyBDb9ohbs8oY3NIZzyMrt8MCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAFgcaqRgv\nqylx4ogL5iUr0K2e/8YzsvH7zLHG6xnr7HxpR/p0lQt3dPlppECZMGDKElbCgU8f\nxVDdZ3FOxHTJ51Vnq/U5xJo+UOMJ4sS8fEH8cfNliSsvmSKzjxpPKqbCJ7VTnkW8\nlonedCPRksnhlD1U8CF21rEjKsXcLoX5PsxlS4DX3PtO0+e8aUh9F4XyZagpejq8\n0ttXkWd3IyYrpFRGDlFDxIiKx7pf+mG6JZ/ms6jloBSwwcz/Nkn5FMxiq75bQuOH\nEV+99S2du/X2bRmD1JxCiMDw8cMacIFBr6BYXsvKOlivwfHBWk8U0f+lVi60jWje\nPpKFRd1mYuEZgw==\n-----END CERTIFICATE-----\n",
        "requestheader-allowed-names": "[]",
        "requestheader-client-ca-file": "-----BEGIN CERTIFICATE-----\nMIICpjCCAY4CCQCZBiNB23olFzANBgkqhkiG9w0BAQsFADAUMRIwEAYDVQQDDAkx\nMjcuMC4wLjEwIBcNMjAwNjE0MTk0OTM4WhgPMjI5NDAzMzAxOTQ5MzhaMBQxEjAQ\nBgNVBAMMCTEyNy4wLjAuMTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB\nAMxIjMd58IhiiyK4VjmuCWBUZksSs1CcQuo5HSpOqogVZ+vR5mdJDZ56Pw/NSM5c\nRqOB3cvjGrxYQe/lKvo9D3UmWLcRKtxdlWxCfPekioJ25/dhGOxtBQcjtp/TSqTM\ntxprwT4fvsVwiwaURFoCOivF4xjQFG0K1i3/m7CiMHODy67M1EfJDrM7Vv5XPIuJ\nVF8HhWBH2HiM25ak34XhxVTX8K97k6wO9OZ5GMqbYuVobTZrSRdiv8s95rkmik6P\njn0ePKqSz6cXNXgXqTl11WtsuoGgjOdB8j/noqTF3m3z17sSBqqG/xBFuSFoNceA\nyBDb9ohbs8oY3NIZzyMrt8MCAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAFgcaqRgv\nqylx4ogL5iUr0K2e/8YzsvH7zLHG6xnr7HxpR/p0lQt3dPlppECZMGDKElbCgU8f\nxVDdZ3FOxHTJ51Vnq/U5xJo+UOMJ4sS8fEH8cfNliSsvmSKzjxpPKqbCJ7VTnkW8\nlonedCPRksnhlD1U8CF21rEjKsXcLoX5PsxlS4DX3PtO0+e8aUh9F4XyZagpejq8\n0ttXkWd3IyYrpFRGDlFDxIiKx7pf+mG6JZ/ms6jloBSwwcz/Nkn5FMxiq75bQuOH\nEV+99S2du/X2bRmD1JxCiMDw8cMacIFBr6BYXsvKOlivwfHBWk8U0f+lVi60jWje\nPpKFRd1mYuEZgw==\n-----END CERTIFICATE-----\n",
        "requestheader-extra-headers-prefix": "[\"bulbasaur-\"]",
        "requestheader-group-headers": "[\"capybara\"]",
        "requestheader-uid-headers": "[\"snorlax\"]",
        "requestheader-username-headers": "[\"panda\"]"
    },
    "kind": "ConfigMap",
    "metadata": {
        "creationTimestamp": "2022-01-28T15:29:08Z",
        "name": "extension-apiserver-authentication",
        "namespace": "kube-system",
        "resourceVersion": "42",
        "uid": "2185c0d7-60a2-4fb8-b30b-873b5c666515"
    }
}
+++ exit code: 1
+++ error: 1
Error when running run_cluster_authentication_trust_tests
warning: Immediate deletion does not wait for confirmation that the running resource has been terminated. The resource may continue to run on the cluster indefinitely.
No resources found
warning: Immediate deletion does not wait for confirmation that the running resource has been terminated. The resource may continue to run on the cluster indefinitely.
No resources found
FAILED TESTS: run_cluster_authentication_trust_tests,




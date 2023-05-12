# Testing - Kubernetes Actions

Context is writng the tests in this file:
[post-deploy configure.go / configure_test.go](https://github.com/vmware-tanzu/tanzu-framework/blob/main/addons/pinniped/post-deploy/pkg/configure/configure_test.go)

## Creating a kubefake.Client begins with feeding it some objects as "starting state"

```golang
// at "this point" in time, whatever you expect for the particular function be be run
// you may have a need for certain objects to already exist on the cluster.  This is
// where you setup a function like this that lets you create the client with some
// known state.  The state is fake, and the client is fake, its not actually talking to
// the API, its a simple object store behind the scenes. But it makes testing work.
newKubeClient: func() *kubefake.Clientset {
    c := kubefake.NewSimpleClientset(
        tkgMetadataConfigmapManagementCluster,
        conciergeDeployment,
        supervisorDeployment,
        supervisorService,
        supervisorCertificateSecret, // This one is tricky as we delete, but don't want it gone for testing. see below.
    )
    // The post-deploy job in this case was deleting a secret and depending on another
    // controller to recreate it.  I forget the reason why.
    // Since that controller does not exist in the test environment, we can essentially
    // intersept the delete action and not act on it.
    // This way the secret continues to exist for our assertions.
    c.PrependReactor("delete", "secrets", func(action kubetesting.Action) (bool, runtime.Object, error) {
        // When we delete the secret in the implementation, we expect the cert-manager controller
        // to recreate it. Since there is no cert-manager controller running, let's just tell the
        // fake kube client to not delete it (i.e., that we "handled" the delete ourselves).
        return actionIsOnObject(action, supervisorCertificateSecret), nil, nil
    })
    return c
},
```

## We can dynamically generate objects to prepopulate clients

```golang
// a function to generate authenticators, so we can swap the audience for different use cases.
jwtAuthenticator := func(audience string) *authv1alpha1.JWTAuthenticator {
    return &authv1alpha1.JWTAuthenticator{
        ObjectMeta: metav1.ObjectMeta{
            Name: "my-cool-authenticator",
        },
        Spec: authv1alpha1.JWTAuthenticatorSpec{
            Issuer:   federationDomain.Spec.Issuer,
            Audience: audience,
            TLS: &authv1alpha1.TLSSpec{
                CertificateAuthorityData: base64.StdEncoding.EncodeToString(supervisorCertificateSecret.Data["ca.crt"]),
            },
        },
    }
}

// some audience that is wrong so we can later test that it changes, always good to make it obvious
newConciergeClient: func() *pinnipedconciergefake.Clientset {
    defaultJWTAuthenticator := jwtAuthenticator("initial-incorrect-audience")
    return pinnipedconciergefake.NewSimpleClientset(defaultJWTAuthenticator)
},
// in another case, we know a specific audience we want
wantConciergeClientActions: []kubetesting.Action{
    kubetesting.NewRootGetAction(jwtAuthenticatorGVR, jwtAuthenticator("").Name),
    kubetesting.NewRootUpdateAction(jwtAuthenticatorGVR, jwtAuthenticator("https://1.2.3.4:12345")),
},
// when we have to put a full object into the client, we can skip all the "zero values".
// anything that is initialized to nil or empty string can be ignored.
wantSupervisorClientActions: []kubetesting.Action{
    kubetesting.NewGetAction(oidcIdentityProviderGVR, upstreamIdentityProvider.Namespace, upstreamIdentityProvider.Name),
    kubetesting.NewGetAction(federationDomainGVR, federationDomain.Namespace, federationDomain.Name),
    kubetesting.NewUpdateAction(federationDomainGVR, federationDomain.Namespace, &configv1alpha1.FederationDomain{
        ObjectMeta: metav1.ObjectMeta{
            Name:      "my-federation-domain",
            Namespace: "pinniped-supervisor",
        },
        Spec: configv1alpha1.FederationDomainSpec{
            Issuer: "https://1.2.3.4:12345",
            TLS:    nil, // when nil, Pinniped will use trusted CA's compiled into the container image
        },
    }),
},
```

### Actually running the test

```golang
// create a client
conciergeClientset := test.newConciergeClient()
// add it to the list
clients := Clients{
    ConciergeClientset:   conciergeClientset,
}
// run the func to test
err := TKGAuthentication(clients)
// are the recorded actions what we expected?
require.Equal(t, test.wantConciergeClientActions, conciergeClientset.Actions())
```


## Kube actions capture the action, and the object

Kube Actions Capture both an action as well as the resource being sent over the wire. For example:

```golang
testing.CreateActionImpl{
    ActionImpl:testing.ActionImpl{
        Namespace:"pinniped-supervisor",
        Verb:"create",
        Resource:schema.GroupVersionResource{
            Group:"config.supervisor.pinniped.dev",
            Version:"v1alpha1",
            Resource:"federationdomains"},
            Subresource:""
        },
        Name:"",
        Object:(
            *v1alpha1.FederationDomain)
            // Pointer here for the actual object.
            // This is a bit unfortunate as you really would like to see the object here.
            // If it was listed, you could see the diff of a deep equals.
            (0xc000581ba0)
        }
    }
```


```golang
[]testing.Action{
	testing.GetActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"get", Resource:schema.GroupVersionResource{Group:"config.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"federationdomains"}, Subresource:""}, Name:"my-federation-domain"},
	testing.CreateActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"create", Resource:schema.GroupVersionResource{Group:"config.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"federationdomains"}, Subresource:""}, Name:"", Object:(*v1alpha1.FederationDomain)(0xc000581ba0)}}

[]testing.Action{testing.GetActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"get", Resource:schema.GroupVersionResource{Group:"config.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"federationdomains"}, Subresource:""}, Name:"my-federation-domain"}, testing.CreateActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"create", Resource:schema.GroupVersionResource{Group:"config.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"federationdomains"}, Subresource:""}, Name:"", Object:(*v1alpha1.FederationDomain)(0xc0008036c0)}}


[]testing.Action{
	testing.GetActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"get", Resource:schema.GroupVersionResource{Group:"idp.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"oidcidentityproviders"}, Subresource:""}, Name:"upstream-oidc-identity-provider"},
	testing.GetActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"get", Resource:schema.GroupVersionResource{Group:"config.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"federationdomains"}, Subresource:""}, Name:"my-cool-authenticator"},
	testing.CreateActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"create", Resource:schema.GroupVersionResource{Group:"config.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"federationdomains"}, Subresource:""}, Name:"", Object:(*v1alpha1.FederationDomain)(0xc0005031e0)}}


// latest?
// expected
[]testing.Action{
	testing.GetActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"get", Resource:schema.GroupVersionResource{Group:"idp.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"oidcidentityproviders"}, Subresource:""}, Name:"upstream-oidc-identity-provider"},
	testing.GetActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"get", Resource:schema.GroupVersionResource{Group:"config.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"federationdomains"}, Subresource:""}, Name:"my-federation-domain"},
	testing.CreateActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"create", Resource:schema.GroupVersionResource{Group:"config.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"federationdomains"}, Subresource:""}, Name:"", Object:(*v1alpha1.FederationDomain)(0xc0006836c0)}
}
// actual
[]testing.Action{
	testing.GetActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"get", Resource:schema.GroupVersionResource{Group:"idp.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"oidcidentityproviders"}, Subresource:""}, Name:"upstream-oidc-identity-provider"},
	testing.GetActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"get", Resource:schema.GroupVersionResource{Group:"config.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"federationdomains"}, Subresource:""}, Name:"my-cool-authenticator"},
	testing.CreateActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"create", Resource:schema.GroupVersionResource{Group:"config.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"federationdomains"}, Subresource:""}, Name:"", Object:(*v1alpha1.FederationDomain)(0xc0000d9860)}
}


// first pair
testing.GetActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"get", Resource:schema.GroupVersionResource{Group:"idp.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"oidcidentityproviders"}, Subresource:""}, Name:"upstream-oidc-identity-provider"},
testing.GetActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"get", Resource:schema.GroupVersionResource{Group:"idp.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"oidcidentityproviders"}, Subresource:""}, Name:"upstream-oidc-identity-provider"},
// second pair - reource doesn't match
testing.GetActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"get", Resource:schema.GroupVersionResource{Group:"config.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"federationdomains"}, Subresource:""}, Name:"my-federation-domain"},
testing.GetActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"get", Resource:schema.GroupVersionResource{Group:"config.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"federationdomains"}, Subresource:""}, Name:"my-cool-authenticator"}, // NAME doesn't match, flag was wrong, fixed to "my-federation-domain"
// third pair - resource doesn't match
testing.CreateActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"create", Resource:schema.GroupVersionResource{Group:"config.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"federationdomains"}, Subresource:""}, Name:"", Object:(*v1alpha1.FederationDomain)(0xc0006836c0)}
testing.CreateActionImpl{ActionImpl:testing.ActionImpl{Namespace:"pinniped-supervisor", Verb:"create", Resource:schema.GroupVersionResource{Group:"config.supervisor.pinniped.dev", Version:"v1alpha1", Resource:"federationdomains"}, Subresource:""}, Name:"", Object:(*v1alpha1.FederationDomain)(0xc0000d9860)} // object referece is difference


// THIS BELOW is the object that we get when we print things out:
// So,
foo := &v1alpha1.FederationDomain{
	TypeMeta:v1.TypeMeta{
		Kind:"",
		APIVersion:""
	},
	ObjectMeta:v1.ObjectMeta{
		Name:"my-federation-domain",
		GenerateName:"",
		Namespace:"pinniped-supervisor",
		SelfLink:"",
		UID:"",
		ResourceVersion:"",
		Generation:0,
		CreationTimestamp:time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		DeletionTimestamp:<nil>,
		DeletionGracePeriodSeconds:(*int64)(nil),
		Labels:map[string]string(nil),
		Annotations:map[string]string(nil),
		OwnerReferences:[]v1.OwnerReference(nil),
		Finalizers:[]string(nil),
		ZZZ_DeprecatedClusterName:"",
		ManagedFields:[]v1.ManagedFieldsEntry(nil)
	},
	Spec:v1alpha1.FederationDomainSpec{
		Issuer:"https://1.2.3.4:12345",
		TLS:(*v1alpha1.FederationDomainTLSSpec)(nil)
	},
	Status:v1alpha1.FederationDomainStatus{
		Status:"",
		Message:"",
		LastUpdateTime:<nil>,
		Secrets:v1alpha1.FederationDomainSecrets{
			JWKS:v1.LocalObjectReference{
				Name:""
			},
			TokenSigningKey:v1.LocalObjectReference{
				Name:""
			},
			StateSigningKey:v1.LocalObjectReference{
				Name:""
			},
			StateEncryptionKey:v1.LocalObjectReference{
				Name:""
			}
		}
	}
}





&configv1alpha1.FederationDomain{
	TypeMeta: metav1.TypeMeta{
		Kind:       "",
		APIVersion: "",
	},
	ObjectMeta: metav1.ObjectMeta{
		Name:            "my-federation-domain",
		GenerateName:    "",
		Namespace:       "pinniped-supervisor",
		SelfLink:        "",
		UID:             "",
		ResourceVersion: "",
		Generation:      0,
		CreationTimestamp: metav1.Time{
			Time: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		DeletionTimestamp:          nil,
		DeletionGracePeriodSeconds: nil,
		Labels:                     nil,
		Annotations:                nil,
		OwnerReferences:            nil,
		Finalizers:                 nil,
		ZZZ_DeprecatedClusterName:  "",
		ManagedFields:              nil,
	},
	Spec: configv1alpha1.FederationDomainSpec{
		Issuer: "https://1.2.3.4:12345",
		TLS: &configv1alpha1.FederationDomainTLSSpec{
			SecretName: "",
		},
	},
	Status: configv1alpha1.FederationDomainStatus{
		Status:         "",
		Message:        "",
		LastUpdateTime: nil,
		Secrets: configv1alpha1.FederationDomainSecrets{
			JWKS: corev1.LocalObjectReference{
				Name: "",
			},
			TokenSigningKey: corev1.LocalObjectReference{
				Name: "",
			},
			StateSigningKey: corev1.LocalObjectReference{
				Name: "",
			},
			StateEncryptionKey: corev1.LocalObjectReference{
				Name: "",
			},
		},
	},
}




// fooActions := supervisorClientset.Actions()
// third := fooActions[2]
// castedThird := third.(kubetesting.CreateAction).GetObject() // GenericObjectInterface from Kubernetes
// castAgain := castedThird.(*configv1alpha1.FederationDomain)
//
// // fmt.Printf("what is the fed domain? %#v", castedThird)
// // castedAgain := configv1alpha1.FederationDomain(castedThird)
// fmt.Printf("what is the fed domain? %#v", castAgain)


//
// I don't think we actually need to check on the federation domain itself, the actions may be sufficient?
// _, err = clients.SupervisorClientset.ConfigV1alpha1().FederationDomains(federationDomain.Namespace).Get(context.Background(), "my-federation-domain", metav1.GetOptions{})
// if test.isMgmtCluster == true {
// 	require.NoError(t, err)
// } else {
// 	require.Error(t, err, "federationdomains.config.supervisor.pinniped.dev \"my-federation-domain\" not found")
// }


			// Note that this does create another action on the client, which skrews the expected output if we were
			// comparing the list of actions.  This is fine now, but if we change the tests in the future, we may want
			// to instead ask the client for its list of actions, get the CreateAction, extract the object, cast it, and compare.
			// because thats lots of fun.
			// authenticator, err := clients.ConciergeClientset.AuthenticationV1alpha1().JWTAuthenticators().Get(context.Background(), "my-cool-authenticator", metav1.GetOptions{})
			// require.NoError(t, err)
			// require.Equal(t, test.wantJWTAuthenticatorAudience, authenticator.Spec.Audience)



			//
// Users:
// - TKGs only for workload clusters
// - TKGm for workload clsuters
// - TKGm for management clusters
// - ANYONE ELSE...... can they use this thing intelligently??
//
// -------------------------------
//
// if we find TKGMetadata
//  return cluster.name, cluster.type
// if we don't find TKGMetadata
//  do we have the JWTAuthenticatorAudience?
//    if true, return "", cluster.type==Workload
//    if false, fall through to searching nodes for annotation, etc.
//
// -------------------------------
//
// if --jwtauthenticator-audience && tkg-metadata CM
//    use --jwtauthenticator-audience && the cluster.Type from tkg-metadata
// if --jwtauthenticator-audience ONLY
//    use --jwtauthenticator-audience && default cluster.Type to workload
// if tkg-metadata
//    use cluster.name and cluster.type from tkg-metadata
// if no tkg-metadata
//    get cluster-name annotation from node, but ERROR as we cannot find cluster.type
//
//

```

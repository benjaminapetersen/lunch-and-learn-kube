# Readme: Basics of Google Cloud

## List configurations.

```bash
gcloud config configurations list
# NAME     IS_ACTIVE  ACCOUNT                     PROJECT                    COMPUTE_DEFAULT_ZONE  COMPUTE_DEFAULT_REGION
# bap      True       admin@benjaminapetersen.me  bap-me                                           us-east4
# default  False      petersenbe@vmware.com       tanzu-user-authentication
# tua      False      admin@benjaminapetersen.me  tanzu-user-authentication  us-east4-b            us-east4
```

## Login to different accounts.

```bash
# login.
# you can login to multiple accounts, this will trigger a browser based auth flow
gcloud auth login
```

## Create a new project and set some defaults.

```bash
# create a project and set some basics
3612  gcloud projects create bap-me
3614  gcloud config set project bap-me
3609  gcloud config set compute/region us-east4
3610  gcloud config set compute/zone us-east4-b
gcloud config set account admin@benjaminapetersen.me
```

```bash
gcloud config set project tua
gcloud config set compute/region us-east4
gcloud config set compute/zone us-east4-b
gcloud config set account petersenbe@vmware.com
```


## list and then activate a specific configuration to use.

Note the `IS_ACTIVE` column will indicate the active configuration.

```bash
gcloud config configurations list
# NAME     IS_ACTIVE  ACCOUNT                     PROJECT                    COMPUTE_DEFAULT_ZONE  COMPUTE_DEFAULT_REGION
# bap      True       admin@benjaminapetersen.me  bap-me                                           us-east4
# default  False      petersenbe@vmware.com       tanzu-user-authentication
# tua      False      admin@benjaminapetersen.me  tanzu-user-authentication  us-east4-b            us-east4
gcloud config configurations activate tua
# Activated [tua].
gcloud config configurations list
# NAME     IS_ACTIVE  ACCOUNT                     PROJECT                    COMPUTE_DEFAULT_ZONE  COMPUTE_DEFAULT_REGION
# bap      False      admin@benjaminapetersen.me  bap-me                                           us-east4
# default  False      petersenbe@vmware.com       tanzu-user-authentication
# tua      True       admin@benjaminapetersen.me  tanzu-user-authentication  us-east4-b            us-east4
```

## list projects for an account

```bash
gcloud projects list
# PROJECT_ID                                     NAME                           PROJECT_NUMBER
# bap-me                                         bap-me                         639638551181
# bapapp-test                                    bapapp-test                    444111585729
# benjaminapetersen.me:api-project-699626034585  API Project                    699626034585
# double-catfish-2514                            auto-app                       421339059971
# feisty-dolphin-2514                            chore-app                      1011657426800
# js-lnl-react-todo-app                          js-lnl-react-todo-app          261713610529
# kids-allowance-ugly-llamas                     kids-allowance                 246332160683
# minecraft-316615                               minecraft                      770898797827
# udacity-microservices-on-kube                  udacity-microservices-on-kube  497249647926
```

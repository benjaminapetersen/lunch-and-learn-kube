
```bash
# using an image that has ping available, if you can ping a service endpoint for a pod
# you should be able to deploy this image, then check the logs and see if the ping connects
kubectl create job ping-idp -n tanzu-system-auth --image <some-image-that-has-ping> -- ping <idp-address>
```
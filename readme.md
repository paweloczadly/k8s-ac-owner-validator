# k8s-ac-owner-validator

Kubernetes [validating admission controller](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#validatingadmissionwebhook). Checks if a pod contains `owner` label.

## usage

Get CA certificate from your cluster (e.g. it can be found in kubeconfig). Then update `caBundle` in **deployment.yml**.

After that, generate certificates for the admission controller and deploy it:

```
./generate_cert.sh --service owner-validator --secret owner-validator-certs
kubectl apply -f deployment.yml
```
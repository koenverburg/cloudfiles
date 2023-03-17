Get all images of all running pods
```bash
kubectl get pods -o jsonpath="{.items[*].spec.containers[*].image}"
```

Replace in place
`kubectl get pod <pod_name> -n <namespace> -o yaml | kubectl replace --force -f -`

Removal
```bash
kubectl delete -n <NS> <service> acme

kubectl delete -n default deployment acme
kubectl delete -n default pod acme
```

ReplicaSets
```bash
kubectl get rs
kubectl describe rs acme
kubectl scale --replicas=5 rs/new-replica-set
```

Namespaces
```bash
k get ns
k get po -n research
k apply -f redis.yaml -n finance
k get pods --all-namespaces
k get pods -o wide --all-namespace
```

Imperative Commands
```bash
k run nginx-pod --image=nginx:alpine
k run redis --image=redis:alpine --labels='tier=db'
k create service clusterip redis-service --tcp=6379:6379
k create deployment webapp --image=kodekloud/webapp-color --replicas=3
k run custom-nginx --image=nginx --port=8080
k create namespace dev-ns
k create deployment redis-deploy --image=redis --replicas=2 -n dev-ns
```

Create a pod named `httpd` with `httpd:alpine` and expose the port
```bash
k run httpd --image=httpd:alpine --port=80
k expose pod httpd --port=80 --name httpd
```

ConfigMaps
```bash
k get configmaps
# or
k get cm # -n <namespace>
k describe cm <name>
k create cm webapp-config-map --from-env-file=./webapp.env
```

```ini
APP_COLOR=darkblue
```

Service Accounts

```
k create serviceaccount <name>
k get serviceaccount
k get sa
k describe sa <name>
```

Service Account adding to Pod, note that your have to delete the pod if you make changes to the service account
```
spec:
  serviceAccountName: <name>
```

disabling the defaulte service account on a pod using
```
spec:
  automountServiceAccountToken: false
```


taints

```bash
kubectl taint nodes controlplane node-role.kubernetes.io/control-plane=:NoSchedule-
```

Rolling deployments
Setting there is an active deployment set to RollingUpdates, we need to update the image
```bash
k set image deployment/<name> <pod>=<new image>
```

Get enabled admission plugins
```bash
kubectl exec kube-apiserver-controlplane -n kube-system -- kube-apiserver -h | grep enable-admission-plugins
```
Add new admission plugins
```bash
kubectl exec kube-apiserver-controlplane -n kube-system -- kube-apiserver --enable-admission-plugins=NamespaceAutoProvision --service-account-issuer=https://kubernetes.default.svc.cluster.local --service-account-key-file=/etc/kubernetes/pki/sa.pub --service-account-signing-key-file=/etc/kubernetes/pki/sa.key --etcd-servers=https://127.0.0.1:2379
```

To edit the kube api server use this command, it will restart after editing
```bash
vim /etc/kubernetes/manifests/kube-apiserver.yaml
# nvim /etc/kubernetes/manifests/kube-apiserver.yaml
```

#### Validating and Mutating Admission Controllers

```yaml
apiVersion: admissionregistration.k8s.io/v1
kind: ValidatingWebhookConfiguration
metadata:
  name: "pod-policy.example.com"
webhooks:
- name: "pod-policy.example.com"
  clientConfig:
    service:
      name: "webhook-service"
      namespace: "webhook-namescape"
    caBundle: ""
  rules:
  - apiGroups:  [""]
    apiVersions: ["v1"]
    operations: ["CREATE"]
    resources: ["pods"]
    scope: "NameSpaced"
```

```bash
kubectl create secret generic webhook-server-tls --from-file=Certificate=/root/keys/webhook-server-tls.crt --from-file=Key=/root/keys/webhook-server-tls.key
```

Admission controllers mutate first and then they validate

Create a tls secret
```bash
kubectl -n <ns> create secret tls <name> \
  --cert "/root/keys/webhook.crt"  \
  --key "/root/keys/webhook.key"
```

In previous steps we have deployed demo webhook which does below

- Denies all request for pod to run as root in container if no securityContext is provided.

- If no value is set for runAsNonRoot, a default of true is applied, and the user ID defaults to 1234

- Allow to run containers as root if runAsNonRoot set explicitly to false in the securityContext


In next steps we have added some pod definitions file for each scenario. Deploy those pods with existing definitions file and validate the behaviour of our webhook


Get all images of all running pods
```bash
kubectl get pods -o jsonpath="{.items[*].spec.containers[*].image}"
```

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



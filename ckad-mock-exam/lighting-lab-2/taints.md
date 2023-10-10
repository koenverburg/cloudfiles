Add taint to the node01 of the cluster.


Specification:
key: app_type
value: alpha
effect: NoSchedule


Then create a pod called alpha, with the redis image and set to toleration node01


```bash
kubectl taint node <node> <key>=<value>:<effect>
# e.g. kubectl taint node node01 app_type=alpha:NoSchedule
```

Add this taint to a pod

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: redis
spec:
  tolerations:
    - effect: NoSchedule
      key: app_type
      value: alpha
  containers:
  - name: redis
    image: redis
```

Adding a label to a node

```bash
kubectl label node <node> <key>=<value>
# e.g. kubectl label node controlplane app_type=beta
```

```bash
kubectl get node <node> --show-labels
# e.g. kubectl get node controlplane --show-labels
```


```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: beta-apps
  labels:
    app: beta-apps
spec:
  replicas: 3
  selector:
    matchLabels:
      app: beta-apps
  template:
    metadata:
      labels:
        app: beta-apps
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: app_type
                operator: In
                values:
                  - beta
      containers:
        - name: nginx
          image: nginx
```

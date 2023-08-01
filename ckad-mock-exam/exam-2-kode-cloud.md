# exam 2 on kode cloud

## 1

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: log-volume
spec:
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteMany
  storageClassName: manual
  hostPath:
    path: /opt/volume/nginx
```

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: log-claim
spec:
  storageClassName: manual
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 200Mi
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: logger
  labels:
    run: logger
spec:
  containers:
    - image: nginx:alpine
      name: logger
      volumeMounts:
        - name: log
          mountPath: /var/www/nginx
  volumes:
    - name: log
      persistentVolumeClaim:
        claimName: log-claim
```

```yaml
````

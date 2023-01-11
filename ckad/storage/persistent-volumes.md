Setting up Persistent Volumes

View the logs of the pod

`kubectl exec webapp -- cat /log/app.log`

Create a pod with volume to persist the logs at `/var/log/webapp`

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: webapp
spec:
  volumes:
    - name: vol
      hostPath:
        path: /var/log/webapp
  containers:
    - name: webapp
      image: kodekloud/event-simulator
      volumeMounts:
        - name: vol
          mountPath: /log
```
Create a persistent volume 

```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-log
spec:
  capacity:
    storage: 100Mi
  accessModes:
    - ReadWriteMany
  hostPath:
    path: /pv/log
  persistentVolumeReclaimPolicy: Retain
```
Create a Persistent Volume Claim

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: claim-log-1
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 50Mi
```

`k describe pvc claim-log-1` -> status PENDING
`k describe pv pv-log` -> status Available

The status is pending because there is a mismatch on the access mode

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: claim-log-1
spec:
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 50Mi
```

Update the webapp to use the persistent volume we just created

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: webapp
spec:
  volumes:
    - name: vol
      persistentVolumeClaim:
        claimName: claim-log-1
  containers:
    - name: webapp
      image: kodekloud/event-simulator
      volumeMounts:
        - name: vol
          mountPath: /log
```

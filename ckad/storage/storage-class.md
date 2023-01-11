`k get sc`
`k get storageclass`

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: local-pvc
spec:
  storageClassName: local-storage 
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 500Mi
```


```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: local-storage
provisioner: kubernetes.io/no-provisioner
volumeBindingMode: WaitForFirstConsumer
```
Create a nginx pod that uses the pcv local-pvc

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  volumes:
    - name: vol
      persistentVolumeClaim:
        claimName: local-pvc
  containers:
    - name: nginx
      image: nginx:alpine
      volumeMounts:
        - name: vol
          mountPath: /var/www/html
```
Create a new Storage Class called delayed-volume-sc
```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: delayed-volume-sc
provisioner: kubernetes.io/no-provisioner
volumeBindingMode: WaitForFirstConsumer
```

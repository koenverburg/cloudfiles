# CKAD mock Exams #1

Create a Pod name `nginx`

`k run nginx --image=nginx:alpine`

Create a Namespace with name `apx`

`k create ns apx`

Create a new Deployment with named `httpd-frontend` with 3 replicas and using image `httpd:2.4-alpine`

`k create deployment httpd-frontend  --image=httpd:2.4-alpine --replicas=3`


Create a new Pod named `messaging` with the image `redis:alpine` and set the labels to `tier=msg`

`k run messaging --image=redis:alpine -l tier=msg`


There is a replica set with 4 pod but they are not running, figure out why and resolve the issue

`k get rs`

`k get rs <name>` correct the image
Editing the rs will not change the pod that are "active" right now

`k delete pod -l <label>`

`k get rs` to see that all the pods have restarted and are running again 


Expose the Redis deployment
`k expose deployment redis --port=6379 --name messaging-service`


Change the environment variable to use green 
`k get pods`
`k get pod <name> -p yaml > <name>.yaml`
change the value
`k replace -f <name>.yaml --force`

Create a configmap
`k create configmap <name> --from-literal=<KEY=value>`
`k describe cm <name>`

Create a Secret
`k create secret generic <name> --from-literal=<KEY=value>`


Update this pod <name> to run as root and with the `SYS_TIME` capability

`k get pod` to find the right pod
`k get pod <name> -o yaml > <name>.yaml`
Edit the file and under spec add securityContext and add runAsUser with the value of 0
Under cotainers, add the securityContext, capabilities, add 'add: ["SYS_TIME"]'
`k replace -f <name>.yaml --force`

Export logs from a pod

`k logs <name> -n <ns> > path/to/file.logs`

Create a Persistent Volume claim, pull up the docs of the PVC, that should have an example the first couple lines until accessMode is enough
Change the values that are needed and create it using `k create -f <naml>.yaml`

Create a Redis Deployment using redis:alpine and 1 replica
`k create deployment redis --image=redis:alpine --replicas=1`
Create a Redis Service on port 6379
`k expose deployment redis --name=redis --port=6379 --target-port=6379`
Create a NetworkPolicy (ingress type) that is called redis-access that will only allow traffic from pods with the label of "access=redis"
(Again, check the docs for an example on this)[https://kubernetes.io/docs/concepts/services-networking/network-policies/]

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: redis-access # changad
  namespace: default
spec:
  podSelector:
    matchLabels:
      app: redis # changed
  policyTypes:
    - Ingress
  ingress:
    - from:
        - podSelector:
            matchLabels:
              access: redis # changed
      ports:
        - protocol: TCP
          port: 6379
```

Create a Pod called `sega` with two containers one called tails with busybox and sleep 3000
Second container is going to be called sonic with the image of nginx and the NGINX_PORT 8080.

```yaml
# sega.yaml
apiVersion: v1
kind: Pod
metadata:
    name: sega
spec:
    containers:
        - name: tails
          image: busybox
          command:
            - sleep
            - "3600"
        - name: sonic
          image: nginx
          env:
            - name: NGINX_PORT
              value: 8080

```


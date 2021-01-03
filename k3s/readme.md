# k3s Cluster (single machine server + agent)

1. Download a Raspberry Pi image [https://ubuntu.com/download/raspberry-pi-core](here)
2. burn on on the SD card
3. After the burning process, create a file called `ssh` in the root of the boot drive of the SD Card
This is needed otherwise ssh will not be avalible out of the box.
4. Setting up K3s. Because of hardware limitation at the time of writing this I'm going to use one Raspberry Pi 3b for this project.
This single Pi will serve as both the server node as the agent node.
5. Run the installation of k3s `curl -sfL https://get.k3s.io | sh -`

> After running the installation command. I could not get get `sudo k3s kubectl get node(s)` to work.
> It will return with a `Error from server (ServiceUnavailable): the server is currently unable to handle the request`
> #### Fix
> To resolve this I had to edit my `boot/cmdline.txt` and add `cgroup_memory=1 cgroup_enable=memory`
> Followed by a reboot and all works.

6. Because I mainly run windows I install `kubectl` via chocolaty so I can manage my cluster from powershell.
7. Copy the contents of `sudo cat /etc/rancher/k3s/k3s.yaml` to `~/.kube/config` on your local machine
8. On your local machine run `kubectl get nodes -o wide` and if all went well you should see your raspberry pi listed
9. Let's install a ui to manage our cluster. By running the following command we can install the kubernetes dashboard
`kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.1.0/aio/deploy/recommended.yaml`
_please note the version number in the url_
10. Create a file called `dashboard.admin-user.yml` with the contents of

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kubernetes-dashboard
```

11. Create another file called `dashboard.admin-user-role.yml` with the contents of

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kubernetes-dashboard
```
12. Let's deploy the admin role by running the following command
`kubectl create -f dashboard.admin-user.yml -f dashboard.admin-user-role.yml`

> moar info can be found here `https://rancher.com/docs/k3s/latest/en/installation/kube-dashboard/#dashboard-rbac-configuration`

13. Obtain the Bearer Token
On unix `kubectl -n kubernetes-dashboard describe secret admin-user-token | grep ^token`
On windows `kubectl -n kubernetes-dashboard describe secret admin-user-token`
14. Open this url on your browser [http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/](kubernetes-dashboard)

15. At this point we have a cluster up and running but we dont have any pods running, Let's create a test deployment.
Here we will use a simple nginx image.
16. Create a file called `test-nginx-deployment.yml` with the contents of
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-test-deployment
  labels:
    app: nginx-test-deployment
spec:
  selector:
    matchLabels:
      app: nginx-test-deployment
  replicas: 1
  template:
    metadata:
      labels:
        app: nginx-test-deployment
    spec:
      containers:
      - name: nginx-test-deployment
        image: nginx
        ports:
          - containerPort: 80
```
17. Let's apply the deployment
`kubectl apply -f test-nginx-deployment.yml`
> Moar info on [https://kubernetes.io/docs/tasks/run-application/run-stateless-application-deployment/](deployments here)!

18. If we run a `kubectl get pods` we should see the a pod running

19. Now let's see what the pod is serving out `kubectl exec -it nginx-test-deployment-7797df98f5-nhl95 curl localhost`
You should see `Welcome to nginx!`.
> _NOTE: `nginx-test-deployment-7797df98f5-nhl95` is the name of my pod your name will differ_


# go-redis-kubernetes

Go with Redis storage implementation demo deployed with Kubernetes.

The go application has a /inc endpoint which should increase the redis store each time hit.

Follow minikube set up, once installed.
```
eval $(minikube docker-env).
docker build -t goapp:v1 .
```

This will build the go app into the conext for minikube.

```
kubectl create -f redis/
kubectl create -f goapp/
```
This will then set the services and deployment scripts for Kubernetes. If you want to check the statuses,
enable the Kube dashboard. 

## 1. why k8s
- Orchastration
    - scaling
    - rollouts
    - loadbalancing
- Declarative/Imperative

## 2. Configs and Contexts and your own cluster

- Create cluster with any two of kind, docker desktop, minikube, rancher, see config file
- How to connect to different clusters
- Instal kubie for a better experience

## 4. Pod
- Setup a pod that prints hello world and then exits
- Basic kubectl commands, logging
- Setup pod with counter program
- Port forwarding
- stern
- Exec, hello world, wait to exec and play around.
- Set it up in different namespaces

## 5. Deployments and services
- Setup deployment for mongo
- Setup service for mongo
- Setup deployment and service for persistent counter along with mongo
- ClusterIP, Loadbalancer, Nodeport
- Connect mongo from different namespace to backend using externalname

## 6. Persistent volumes and volume claims
- Setup pvc for mongo to work properly
- Use bitnami mongo, outsource work, NIH syndrome

## 7. Configmaps and secrets
- hostname should come from configmap
- Change mongo to have password, store those in a secret
- git crypt
- Sealed secrets

## 8. Resource limits 
- Cpu and Memory limits

## 9. Ingress
- Note to self: Find a way to do this locally easier 
- Setup ingress for frontend of counter app

## 10. Hpa
- Make counter scale based on usage, try it out with curl on repeat

## 11. Network policies
- Only counter app should be able to access db
- Counter app shouldn't be able to acess anything other than db
- db shouldn't be able to access anything
- exec into db pod and try to ping anything in the network or the internet

## 12. Prometheus, Grafana, Loki (Optional)
- Setup a simplistic version of it, and pipe logs to a central logging system
- Know about them, but indepth knowleddge isn't required yet.  

## 13. Liveness and readiness probes
- Setup liveness and readiness for counter app

## 14. Customization, and templating
- Kustomize
- Helm intro

# Deployment

## How this is deployed and what resources are relevant
A Helm chart was created to enable the deployment of this service. 

Helm has support to generate some basic deployment files that a service would require. 
```
helm create calc-api
```

In this case the base files did not need a lot of modifications besides the image to be used and some changes in the `Service` and `Deployment` resources to configure ports, readiness and liveness probes.

The created files include:
* Deployment - To describe the desired state of the application (replicas, pod template, ports etc.)
* Service - To group our pods together and facilitate service discovery
* Ingress - To manage external access to our app
  
A StatefulSet would not make sense in our case, given that the app is stateless and that StatefulSet are primarily used to deploy datostores (because of the guarantees they give for the order and uniqueness of the pods).

## Running the deployment 
This would require a cluster and `kubectl` to be configured with the proper access.
If a cluster is not available we can use `minikube` to spin up a local one.
```
brew install minikube
minikube start
kubectl get nodes
```

First we need to build the container using the minikube docker daemon.
```
eval $(minikube docker-env) 
make build-all -B
```
s
Once the cluster is up and the image is built we can then deploy our application.
```
helm install calc-api  ./deploy/calc-api
```
## Access from outside the cluster
The two preferred methods that could be used to expose this application externally: 
* a Service of type LoadBalancer
  * this would create an ELB in the case of AWS and would configure it to route traffic to our service
  * can support any kind of traffic
* an Ingress resource
  * This would require the existence of an Ingress Controller exposed through a Load Balancer
  * costs savings, services can reuse the some Load Balancer
  * can support the type of traffic the controller implementation supports (ie nginx)
  

## Cloud Native

"Borrowing" two definitions from the Internet.

Pivotal
>Cloud native is an approach to building and running applications that exploits the advantages of the cloud computing delivery model. It is about how  applications are created and deployed, not where. It implies that the apps live in the public cloud, as opposed to an on-premises data center.

CNCF
>The CNCF defines “cloud-native” a little more narrowly, to mean using open source software stack to be containerized, where each part of the app is packaged in it's own container, dynamically orchestrated so each part is actively scheduled and managed to optimize resource utilisation, and microservices-oriented to increase the overall agility and maintainability of applications.


The 12 Factor App has been cited as one of the best guides to Cloud Nativeness, our application has been developed following those guidelines (as already described in our docs). The facts that it is containerised and can be easily deployed to a cluster are strong indicators that it is on the right path. Moreover we seem to satisfy both definitions, given the fact that our app is primarily developed to run on a public cloud, and can be dynamically orchestrated and scale .
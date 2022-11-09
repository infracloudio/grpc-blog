# setup
1. a simple greet unary gRPC service that takes first_name and last_name as input and replies with 'Hello, <first_name> <last_name>'.
2. Dockerfile to build the image - build/Dockerfile
3. K8s config files to deploy the service in cluster - deploy/

# testing [in postman]
this section will be updated

# testing w/o ingress
Ingress should not be used to expose gRPC service outside the cluster. why?
Because generally gRPC services are used internally [primarily in a microservice architecture].
And if in production, ingress is definitely not be existing or if it exists then should be put under firewall or ips are whitelisted.
But that is not ideal.
So since production is not going to have an ingress, definitely stage environment is not going to have one as it is supposed to mimic production environment.
But we need to test (functional and load) the stage environment.
If we do testing in QA/Dev environment using ingress, then we have to employ different strategies in Dev/QA and Stage.
So in order to maintain sanity across environments, it is better to create a testing-pod that runs tests within the cluster and therefore will not require ingress.

grpcurl is the ideal tool to carry out the ops if reflection is not enabled for grpc services.
https://aristanetworks.github.io/openmgmt/examples/gnoi/grpcurl/

evans-cli is the ideal tool to carry out the testing ops if reflection is enabled for grpc services.
https://github.com/ktr0731/evans
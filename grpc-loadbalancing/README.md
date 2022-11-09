One of the challenging use cases is achieving load balancing in gRPC.

Let's start with a REST server.

Say we spin up 3 Kubernetes Pods (all of them are REST servers)

Then we create a Kubernetes Service (type: ClusterIP) that points to these REST servers.

Then we create a Kubernetes Ingress(it exposes the underlying service), in order to hit these REST endpoints over internet.

The ClusterIP service acts as a load balancer and sends request to the targeted REST servers in a round-robin(default) fashion.

That is not the way how it works with gRPC servers.

Say we spin up 3 more Kubernetes Pods (all of them are gRPC servers)

Then we create a Kubernetes Service (type: ClusterIP) that points to these gRPC servers.

Then we create a Kubernetes Ingress, in order to hit these gRPC endpoints over internet.

Whenever we hit the url using gRPC, the requests will always end up in a single gRPC server.

This happens because REST servers are based on http1.1 while gRPC is inherently based on http2.

What happens here is, when we dial the server url in gRPC client, it sets up a persistent connection with a server (like a websocket, the mechanism is very different though)

Due to which every time the requests are kinda multiplexed on to that particular server.

So the onus of load balancing falls on gRPC clients, to set the load balancing settings/strategy.

Also, it is necessary to make the Kubernetes service headless (ClusterIP: None).

Here is an example of a headless Kubernetes service
```
apiVersion: v1
kind: Service
metadata:
    name: grpc-service
spec:
    clusterIP: None
    selector:
        app: grpc-server
    ports:
    - port: 80
      targetPort: 50051
```

Here is a snippet on how to set a round-robin load balancing in gRPC client.
```
cc, err := grpc.Dial(servAddr, 
        grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), grpc.WithBlock(),
        opts,
    )
```

Setting load balancing in gRPC clients is not an ideal solution though.
### Description
"For Study"  
Repo. for finding how to add external authentication in Cilium CNI, using CiliumEnvoyConfig

### How to Setup
- Apply `cilium-yaml/test-application.yaml`
- Edit, and Build Go Code (Default Port: 50051)
- Run Auth Server
- Edit `cilium-yaml/ext_authz.yaml`
  - Especially, Auth Server IP and Port. (`target_uri` in `ext_authz` and `socket_address` in `Cluster` named `default/ext_authz`)
- Apply `cilium-yaml/ext_authz.yaml`
- Run `kubectl exec -it kubectl exec -it pod/client2-84bc4c4b59-b9ct4 -- curl echo-service-1.default.svc.cluster.local:8080`. Then you can see HTTP Header with Source/Destination.
- Modify Go code as you want

### What you have to know
- because of `dns_refresh_rate: 1s`, Envoy will check health of Auth Server every 1s. Change this if you want.

### Code From
- `cilium-yaml/test-application.yaml` https://raw.githubusercontent.com/cilium/cilium/1.17.2/examples/kubernetes/servicemesh/envoy/test-application.yaml

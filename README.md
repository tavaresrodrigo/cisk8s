#CISK8s

A controller is a reconciliation loop that reads the desired state of a resource from the Kubernetes API and takes action to bring the cluster's actual state closer to the desired state [1](https://developers.redhat.com/articles/2021/06/22/kubernetes-operators-101-part-2-how-operators-work#how_operators_reconcile_kubernetes_cluster_states). The main role of a controller is to move the current cluster state closer to the desired state. 

This API server offers the methods for Custom Controller to mitigate security issues in Kubernetes cluster Control Plane Nodes by automating the application of security recommendations that must enhance the default workload security aspects of given K8s cluster. 
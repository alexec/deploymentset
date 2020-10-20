# DeploymentSet

A deployment-set describes a deployment and the namespaces it should be installed in.

Run the controller in any namespaces, and in that same namespace create a `DeploymentSet`. 

This example creates an Argo Workflows controller in each namespace labelled with `"workflows.argoproj.io/argo": "true"`


```yaml
apiVersion: alexec.alexec/v1
kind: DeploymentSet
metadata:
  name: operator
spec:
  selector:
    matchLabels:
      "workflows.argoproj.io/argo": "true"
  template:
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: workflow-controller
    spec:
      selector:
        matchLabels:
          app: workflow-controller
      template:
        metadata:
          labels:
            app: workflow-controller
        spec:
          serviceAccountName: argo
          containers:
            - name: workflow-controller
              image: argoproj/workflow-controller:latest
              command:
                - workflow-controller
              args:
                - --configmap
                - workflow-controller-configmap
                - --executor-image
                - argoproj/argoexec:latest
          securityContext:
            runAsNonRoot: true
          nodeSelector:
            kubernetes.io/os: linux
```

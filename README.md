Pseudo Config Map
=================

A small script to run as a sidecar container in order to create a quick updating
ConfigMap. At least as on Kubernetes 1.11, updates to a ConfigMap can take an
extended period of time to appear in pods.

While the community works on a built in solution (https://github.com/kubernetes/kubernetes/issues/30189),
this sidecar is a workaround. It watches for all updates to a single ConfigMap,
and replicates changes to that ConfigMap into the directory `/config`. Create
an `emptyDir` volume mounted to `/config` in the sidecar, and then also mount 
the volume (at any path) in the container that needs quick updating 
configuration.

Usage
-----

At runtime, two environment variables must be set:

1. `NAMESPACE`: the name of the current namespace that contains the ConfigMap
2. `CONFIG_MAP_NAME`: the name of the ConfigMap to watch and replicate in the emptyDir volume

In addition, mount a volume to `/config` to capture the updated configuration.

The ServiceAccount running the sidecar must have read and watch access to 
ConfigMaps in the namespace. For example, in OpenShift you can run `oc adm 
policy add-role-to-user view -z default` to give view access of all resources
to pods in the project.


Example
-------

```
apiVersion: v1
kind: Pod
metadata:
  name: pseudoconfigmap-12-gjslv
  namespace: cpitman-pseudoconfigmap
spec:
  containers:
    - env:
        - name: CONFIG_MAP_NAME
          value: test-config-map
        - name: NAMESPACE
          valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
      image: ...
      name: pseudoconfigmap
      volumeMounts:
        - mountPath: /config
          name: config-volume
      ...
  volumes:
    - emptyDir: {}
      name: config-volume
    ...
```
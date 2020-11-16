+++
cover = "/posts/screenshot-from-2020-08-07-20-16-40.png"
date = 2020-09-22T22:00:00Z
draft = true
slug = ""
title = "Simple Canary Deployments using Kubernetes StatefulSets on OpenShift"

+++
> This blog post was originally posted on the [OpenShift blog](https://www.openshift.com/blog/simple-canary-deployments-using-kubernetes-statefulsets-on-openshift). This is just a personal mirror of that blog post I wrote.

In this blog post I want to introduce a nice little trick to make canary deployments possible only with Kubernetes StatefulSets and Services. First, let me give a bit of background on why I went down that route. At Red Hat we continuously increase the number of applications we run ourselves. The Red Hat OpenShift Monitoring and Observability Teams are responsible for running a service called Telemeter. We use this service to get some basic telemetry of OpenShift clusters for Remote Health Monitoring.

In a previous version of Telemeter, the system ingested all the data into an in-memory hashring, which was scraped by two [Prometheus](https://prometheus.io/) replicas. This setup allowed us to quickly get started on providing basic telemetry functionality for OpenShift but we soon identified a major bottleneck. Even though the hashring was able to handle more requests easily, the two Prometheus instances still scraped all of the data. We needed something more scalable to replace the Prometheus instances with.

[Thanos](https://thanos.io/) was our clear choice, as we had some positive experiences with it in the past and we are already involved with the upstream community. While we were replacing the two Prometheus instances with Thanos components, including the Thanos Querier, Thanos Receiver - [which we actually developed for this project](https://thanos.io/tip/proposals/201812_thanos-remote-receive.md/) -, Thanos Store and Thanos Compactor, we also needed to make some heavy changes to the Telemeter project, which ran in front of the Thanos stack.

Going forward, we needed a Telemeter to forward all the time series it received into Thanos Receiver via the Prometheus [remote-write protocol](https://prometheus.io/docs/prometheus/latest/storage/#remote-storage-integrations). This change would allow us to get rid of the hashring, hughly simplifying the Telemeter component itself.

Needless to say, the Telemeter service needed to undergo some heavy architectural changes!

In the process of redesigning the Telemeter service, I realized that we had some bugs that were hard to reproduce locally or in our staging environment, as the issues only manifested when the system was under heavy load. We considered approaches like traffic shadowing or synthetic traffic generation in order to reproduce the bugs, but ultimately one afternoon I went in a different direction and added Canary Deployments to Telemeter, as the other approaches were a lot more involved.  
Let’s walk you through how I added Canary Deployments in a bit more depth.

### Canary Deployments with StatefulSets

The Kubernetes documentation has a wonderful section that showcases how to go about creating [Canary Deployments with Kubernetes Deployments](https://kubernetes.io/docs/concepts/cluster-administration/manage-deployment/#canary-deployments). This was the main inspiration for accomplishing the same result with StatefulSets.

```jsonnet
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: telemeter-server
spec:
  replicas: 10
  selector:
    matchLabels:
      app.kubernetes.io/name: telemeter-server
  serviceName: telemeter-server
  template:
    metadata:
      labels:
        app.kubernetes.io/name: telemeter-server
    spec:
      ...
```

Above you can see our StatefulSet’s metadata. Most importantly the `spec.template.metadata` has labels in it `app.kubernetes.io/name: telemeter-server` that are added to each Pod created by the StatefulSet controller. Those labels are used by the Kubernetes Service to select which Pods should receive traffic.

```yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    app.kubernetes.io/name: telemeter-server
  name: telemeter-server
spec:
  clusterIP: None
  ports:
  - name: external
    port: 8443
    targetPort: external
  - name: internal
    port: 8081
    targetPort: internal
  selector:
    app.kubernetes.io/name: telemeter-server
```

The most important part of this Kubernetes Service, namely the last 2 lines. We specify that we want the Kubernetes Service to select all Pods that have a label of `app.kubernetes.io/name`.

Utilizing that feature, we can actually create arbitrary Pods and as long as they have that label the Service is going to route traffic to them.

```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: telemeter-server-canary
spec:
  replicas: 2
  selector:
    matchLabels:
      app.kubernetes.io/name: telemeter-server
      track: canary
  serviceName: telemeter-server
  template:
    metadata:
      labels:
        app.kubernetes.io/name: telemeter-server
        track: canary
    spec:
      ...
```

As you can see this is a second StatefulSet. It’s pretty similar except that we added a `-canary` suffix to the name and an additional `track: canary` label.

Because the `app.kubernetes.io/name: telemeter-server` is present on the Pods of this StatefulSet, the telemeter-server Service is going to route traffic to these two Pods as well.

With this we are able to run different versions of Telemeter next to Pods running our latest stable version. In this case our Canary Deployment has two replicas so approximately 16% of traffic will be sent to the canary because `(2/(2+10) = 2/12 ~ 16%)]`. Scaling the Pods of the Canary Deployment up and down we can change the percentage of Canary traffic sent to these Pods.

Once we have verified the new changes work as planned, we pretty much always deploy the new version to the 10 stable replicas and scale down the replicas of the Canary StatefulSet to 0.

### Generating the Canary StatefulSet

Because we specify all of our Prometheus and Kubernetes YAML with Jsonnet, we can do some nice things to keep both StatefulSets consistent in the future. Labels, arguments and other things will always be updated on the Canary StatefulSet just like the stable StatefulSet.

```jsonnet
    statefulSetCanary: self.statefulSet {
      metadata+: {
        name: super.name + '-canary',
      },
      spec+: {
        replicas: 2,
        selector+: {
          matchLabels+: {
            track: 'canary',
          },
        },
        template+: {
          metadata+: {
            labels+: {
              track: 'canary',
            },
          },
          spec+: {
            containers: [
              if c.name == 'telemeter-server' then c {
                image: '${IMAGE_CANARY}:${IMAGE_CANARY_TAG}',
                command+: ['--log-level=debug'],  // Always enable debug logging for canary deployments
              }
              for c in super.containers
            ],
          },
        },
      },
    },
```

The first line of this snippet copies the original StatefulSet and creates a statefulSetCanary, which has some additional adjustments, like the `-canary` suffix, the `track: canary` label, and always enabled debug logging.

If something changes in the original StatefulSet, the exact same change will propagate to the canary StatefulSet unless we explicitly overwrite it.

### Monitoring the Canary StatefulSet

We use the [Prometheus Operator](https://github.com/prometheus-operator/prometheus-operator) to manage and configure all of the Prometheus instances that scrape our services. Therefore, we have a ServiceMonitor, that is a custom Kubernetes object that configures Prometheus, to scrape Telemeter too:

```yaml

apiVersion: monitoring.coreos.com/v1
kind: ServiceMonitor
metadata:
  labels:
    app.kubernetes.io/name: telemeter-server
    prometheus: app-sre
  name: telemeter-server
spec:
  endpoints:
  - interval: '60s'
    port: internal
    scheme: https
  jobLabel: app.kubernetes.io/name
  namespaceSelector:
    matchNames:
    - ${NAMESPACE}
  selector:
    matchLabels:
      app.kubernetes.io/name: telemeter-server
```

Again, take a look at the last 3 lines where we specify the selector to which labels we want to match. The label is exactly the same as the label selector in the Telemeter Service, which means that the Canary Deployment’s Pods will be scraped as well, as they have the same label. Not a single thing needs to change to pick up those Pods too. Yay!
Based on these scraped metrics I created a dashboard show the same metrics once for the stable Pods and once for the canary Pods. With this it’s easy to at least spot problems from a high level point of view.

![Canary Deployments Grafana Dashboard](/posts/screenshot-from-2020-08-07-20-16-40.png)

### Conclusion

While there are many options to automate Canary Deployments via Continuous Delivery or Service Meshes, I hope to have shown you that for simple use cases the capabilities of Kubernetes, or in this case OpenShift, StatefulSets, and Deployments are enough to get you pretty far. It’s a mostly manual process for us still and we are fine with that. If we would want to automate Canary Deployments more thoroughly we certainly would take a look at the options above.
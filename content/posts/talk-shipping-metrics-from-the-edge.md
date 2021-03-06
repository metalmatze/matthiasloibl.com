+++
cover = "/posts/shippingmetricsfromtheedge.png"
date = 2019-11-23T08:00:00Z
slug = "talk-shipping-metrics-from-the-edge"
title = "Talk: Shipping Metrics from the Edge"

+++
In November 2019 KubeCon North America took place at San Diego's Convention Center.  
I attended the conference and gave this talk on Wednesday.

<!--more-->

The source code and configuration used to deploy both the kube-prometheus stack to a Raspberry Pi 4 as well as the  Kubernetes deployments can be found in my [talk repository](https://github.com/metalmatze/talks/tree/master/2019-11%20-%20KubeCon%20NA).

Most of the work was porting the individual components of kube-prometheus to ARM that didn't yet support multi-arch containers, unlike [Prometheus for example](https://quay.io/repository/prometheus/prometheus?tab=tags).  
I hope to contribute the changes I made to the upstream projects so deploying a monitoring stack on a Raspberry Pi with something like [k3s](https://k3s.io/) works out of the box in the future.

Here's the abstract for this talk:

> Computing is getting pushed to the edge, it may be your car, TV, washing machine, or your toaster. All these devices have a lot of computing power these days. While extending the cloud to the edge is getting solved with projects like KubeEdge or k3s, in this talk we want to take a closer look at how to run Prometheus on them. We want to configure Prometheus in a way that we can replicate its data to a central collecting point, that is running Thanos on Kubernetes in a replicated setup, and then make use of all the shipped metrics to efficiently query across the entire fleet.

### Video

<div class="youtube">
<iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/FrcfxkbJH20?rel=0" frameborder="0" allowfullscreen></iframe>
</div>

### Slides

<script async class="speakerdeck-embed" data-id="d57b4373f8dc49d397437a261b8c9346" data-ratio="1.77777777777778" src="//speakerdeck.com/assets/embed.js"></script>

You can find my Slides over at [Speakerdeck](https://speakerdeck.com/metalmatze/prometheus-styx).
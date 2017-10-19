+++
date = "2017-08-17T17:30:00+01:00"
title = "Prometheus Styx"
slug = "prometheus-styx"
cover = "posts/34/styx.png"
+++

Introducing [Styx](https://github.com/go-pluto/styx) for Prometheus.  
A simple tool for exporting data from Prometheus to CSV, gnuplot & matplotlib.

<!--more-->

## Video

<div class="youtube">
    <iframe width="560" height="315" src="https://www.youtube-nocookie.com/embed/ovJDD9UOxLc?rel=0" frameborder="0" allowfullscreen></iframe>
</div>

## Slides

<script async class="speakerdeck-embed" data-id="578366fea10549c7872d44c61f2b50ce" data-ratio="1.77777777777778" src="//speakerdeck.com/assets/embed.js"></script>

You can find my Slides over at [Speakerdeck](https://speakerdeck.com/metalmatze/prometheus-styx).

### Installation

```bash
go get -v -u github.com/go-pluto/styx
```
Optional dependencies are [gnuplot](http://www.gnuplot.info/) & [matplotlib](https://matplotlib.org/).

### Usage

#### CSV

```bash
# export the data for the last hour from http://localhost:9090 
styx 'go_goroutines'
styx 'sum(go_goroutines)'
styx 'go_goroutines{job="prometheus"}'
styx 'go_goroutines > 100'
# export the data for the last 6 hours from http://localhost:9090
styx --duration 6h 'sum(go_goroutines)' 
# export the data from a specific prometheus for the last hour.
styx --prometheus http://prom.example.com 'sum(go_goroutines)' 
```

#### gnuplot
```bash
# plot the data for the last hour from http://localhost:9090
styx gnuplot 'sum(go_goroutines)' > goroutines.gnuplot
# plot the data for the last 6 hours from http://localhost:9090
styx gnuplot --duration 6h 'sum(go_goroutines)' > goroutines.gnuplot 
# plot the data from a specific prometheus for the last hour.
styx gnuplot --prometheus http://prom.example.com 'sum(go_goroutines)' > goroutines.gnuplot
```

Once you have written the generated content into a file you can use this to edit and plot the graph:

```bash
gnuplot -p < test.gnuplot
```

#### matplotlib

```bash
# plot the data for the last hour from http://localhost:9090
styx matplotlib 'sum(go_goroutines)' > goroutines.py
# plot the data for the last 6 hours from http://localhost:9090
styx matplotlib --duration 6h 'sum(go_goroutines)' > goroutines.py 
# plot the data from a specific prometheus for the last hour.
styx matplotlib --prometheus http://prom.example.com 'sum(go_goroutines)' > goroutines.py
```

Once you have written the generated content into a file you can use this to edit and plot the graph:

```bash
python goroutines.py
```

---

Check it out at [GitHub](https://github.com/go-pluto/styx).

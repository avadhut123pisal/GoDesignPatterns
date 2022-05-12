# GoDesignPatterns

Various design patterns are implemeted in golang.
## Problem statement:
Detect intrusions and anomalous activities in the cloud applications.

## Overview:

- In a dynamic infrastructure platform such as Kubernetes, detecting and addressing threats is important but also challenging at the same time.

- There are many events in the systems, that are not common or should not even happen like outgoing connections to the unexpected destinations or modifications in certain folders such as /etc.

- Gaining insights into such events inside the cluster makes it possible to detect attacks and potential malicious behavior and to alert operations staff at an early stage.


## Getting Started:

- In order to detect the anomalous activities in our cloud applications, we are using the [Falco](https://falco.org) tool. 
The Falco Project is an open source runtime security tool which is now a CNCF incubating project.

- Falco ships with a default set of rules that check the kernel for unusual behavior and can be extended with custom rules.

- We will consider the example of the LM-OTEL collector and detect the unusual behavior inside the container at runtime. We have developed, a set of rules that are important for LM-OTEL collector, if any of the rules gets violated due to the unusual behavior inside the LM-OTEL collector container, events will be sent to the Logicmonitor platform in the form of the Log.

## Setup:

### Prerequisite:
- We need a working Kubernetes setup for this. 
- We also require kubectl and Helm to be installed on your client machine.

### Install Falco:

```bash
helm repo add falcosecurity https://falcosecurity.github.io/charts
helm repo update
```

```bash
helm install falco falcosecurity/falco \
--set falco.fileOutput.enabled=true \
--set falco.fileOutput.filename=/var/eventlog/events.log \
--set extraVolumes[0].name=event-log \
--set extraVolumes[0].hostPath.path=/var/log \
--set extraVolumeMounts[0].name=event-log \
--set extraVolumeMounts[0].mountPath=/var/eventlog \
-f custom-rules.yaml
```

### Install LM-OTEL Collector:

```bash
helm repo add logicmonitor https://logicmonitor.github.io/k8s-helm-charts
helm repo update
```

```bash
helm install -n hackathon \                                                
--create-namespace \
--set lm.account=qauat02 \
--set lm.access_id="xxxxxxx" \
--set lm.access_key="xxxxxxx" \
--set lm.otel_name="threat-hunters-hackathon-demo" \
--set image.repository="khyati123/logicmonitor-lmotel" \
lmotel-demo logicmonitor/lmotel
```

### Install fluent-bit:

```bash
helm repo add fluent https://fluent.github.io/helm-charts
helm repo update
```

```bash
helm install fluent-bit fluent/fluent-bit -f fluent-bit-values.yaml
```

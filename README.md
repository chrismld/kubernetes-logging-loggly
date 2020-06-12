# kubernetes-logging-loggly
A sample application to demonstrate how logging works using Loggly and Fluentd

```
$ kubectl create secret generic loggly-config --from-literal=loggly.token="TOKEN"
```
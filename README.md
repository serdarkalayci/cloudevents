# cloudevents

The kourier service is a LoadBalancer service by default. 
If your cluster does not support LoadBalancer services, you can use the [kouriersvc.yaml](serving/kouriersvc.yaml) file to update it to be a NodePort service.
Also you should update Knative's config-domain configmap to use a base domain name for your local environment.
```bash
kubectl patch configmap/config-domain \
      --namespace knative-serving \
      --type merge \
      --patch '{"data":{"localdev.me":""}}'
```
After these changes, when you apply the [service.yaml](serving/service.yaml) file, you should see a it's created with the specified domain name.
```bash
kubectl get ksvc
NAME                  URL                                              LATESTCREATED               LATESTREADY                 READY   REASON
cloudevents-serving   http://cloudevents-serving.default.localdev.me   cloudevents-serving-00002   cloudevents-serving-00002   True
```

From now on you can use the URL and the port you've set for the kourier service to send events to your service.

To send a single request to the service, you can use curl:
```bash
curl -v -X POST \
    -H "content-type: application/json"  \
    -H "ce-specversion: 1.0"  \
    -H "ce-source: curl-command"  \
    -H "ce-type: curl.demo"  \
    -H "ce-id: 123-abc"  \
    -d '{"name":"Serdar", "number1": 25, "number2": 50, "sleep": 10000, "bloat": 50, "prime": 50}' \
    http://cloudevents-serving.default.localdev.me:31080
```

To try load on Knative Serving:
```bash
hey -z 30s -c 50 -m POST  \
    -H "content-type: application/json"  \
    -H "ce-specversion: 1.0"  \
    -H "ce-source: curl-command"  \
    -H "ce-type: curl.demo"  \
    -H "ce-id: 123-abc"  \
    -d '{"name":"Serdar", "number1": 25, "number2": 50, "sleep": 10000, "bloat": 50, "prime": 50}' \
  "http://cloudevents-serving.default.localdev.me:31080"
`

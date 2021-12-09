# cloudevents

To try load on Knative Serving:

`
hey -z 30s -c 50 -m POST  \
    -H "content-type: application/json"  \
    -H "ce-specversion: 1.0"  \
    -H "ce-source: curl-command"  \
    -H "ce-type: curl.demo"  \
    -H "ce-id: 123-abc"  \
    -d '{"name":"Serdar", "number1": 25, "number2": 50, "sleep": 10000, "bloat": 50, "prime": 50}' \
  "http://cloudevents-serving.default.127.0.0.1.sslip.io"
`

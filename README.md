# cloudevents

To try load on Knative Serving:

`
hey -z 30s -c 50 -m POST  \
    -H "content-type: application/json"  \
    -H "ce-specversion: 1.0"  \
    -H "ce-source: curl-command"  \
    -H "ce-type: curl.demo"  \
    -H "ce-id: 123-abc"  \
    -d '{"name":"Dave"}' \
  "http://cloudevents-go.default.127.0.0.1.sslip.io?sleep=100&prime=10000&bloat=5"
`

# Overview

Github event gateway exposes an HTTP endpoint which is registered as a repository web-hook in Github.

The gateway receives the web-hook payload from Github, converts it to Kyma Events and forwards to the Kyma Event Service.

The generated image is used in the [Kyma connector add-ons](https://github.com/abbi-gaurav/kyma-connectors/blob/master/addons/github-connector-0.0.1/chart/github-connector/templates/event-gw.yaml) Helm chart.

## Build and push Image

```shell script
make push-image DOCKER_ACCOUNT=<Your docker account>
```

## Command Line Arguments

| Argument              | Type    | Description                                                     | Default                                                                   |
|-----------------------|---------|-----------------------------------------------------------------|---------------------------------------------------------------------------|
| **verbose**           | boolean | Log each incoming webhook payload. Useful for debugging.        | false                                                                     |
| **app-name**          | string  | Name of the Kyma application that sends the events.             | github                                                                    |
| **event-publish-url** | string  | in-cluster url of the service to which Kyma events are forwaded | http://event-publish-service.kyma-system.svc.cluster.local:8080/v1/events |
| **secret-key**        | string  | Secret Key used to verify the HMAC signature of the payload     | ""                                                                        |


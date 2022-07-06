# Kratos Project Template

consul

```docker
docker run -d --name=consul_153 \
--network=host \
-e CONSUL_BIND_INTERFACE=enp5s0f0 \
consul agent \
--server=true \
--bootstrap-expect=1 \
 -node=leader --client=0.0.0.0 -ui
```
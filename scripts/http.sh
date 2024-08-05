#!/bin/bash
for d in ./internal/inbound/http/*/openapi.yaml; do
     mkdir -p ./internal/inbound/model

    domain="$(basename "$(dirname "$d")")"

    oapi-codegen \
    -config ./variables/codegen/model.cfg.yaml \
    -o ./internal/inbound/model/$domain.gen.go \
    ./internal/inbound/http/$domain/openapi.yaml

    oapi-codegen \
    -config ./variables/codegen/server.cfg.yaml \
    -o ./internal/inbound/http/$domain/openapi_server.gen.go \
    -package $domain \
    ./internal/inbound/http/$domain/openapi.yaml
done

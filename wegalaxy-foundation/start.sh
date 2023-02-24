#!/bin/sh

envsubst < /config/config.template.yaml > /config/config.yaml
exec /wegalaxy-foundation
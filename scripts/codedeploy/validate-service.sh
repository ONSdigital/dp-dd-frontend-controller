#!/bin/bash

if [[ $(docker inspect --format="{{ .State.Running }}" dp-dd-frontend-controller) == "false" ]]; then
  exit 1;
fi

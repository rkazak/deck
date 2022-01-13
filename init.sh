#!/usr/bin/env bash

until docker ps -f health=healthy | grep -q deck_kong_1; do
    COMPOSE_PROFILES=database KONG_DATABASE=postgres docker-compose up -d
    docker ps
    sleep 15
done
#!/usr/bin/env bash
COUNTER=0
MAX_ATTEMPTS=10
until docker ps -f health=healthy | grep -q deck_kong_1 || [[ $COUNTER -eq $MAX_ATTEMPTS ]] ; do
    COMPOSE_PROFILES=database KONG_DATABASE=postgres docker-compose up -d
    docker ps
    sleep 10
    ((COUNTER++))
done
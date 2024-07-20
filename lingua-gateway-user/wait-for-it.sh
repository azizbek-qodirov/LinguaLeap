#!/usr/bin/env bash
# wait-for-it.sh

cmd="$@"

until nc -z "rabbitmq" "5672"; do
  >&2 echo "RabbitMQ is unavailable - sleeping"
  sleep 1
done

>&2 echo "RabbitMQ is up - executing command"
exec $cmd

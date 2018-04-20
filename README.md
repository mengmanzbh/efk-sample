# Elasticsearch, Fluentd and Kibana (EFK) stack sample

## To run this example:
1. Go to efk folder and run: `docker-compose up -d`
2. Go to app folder and run: `docker-compose up -d`

## Tear down:
1. Go to efk folder and run: `docker-compose down`
2. Go to app folder and run: `docker-compose down`

## Elasticsearch as a single node instace:
To run elasticsearch as a single node instance uncomment the following lines inside `efk/docker-compose.yml` file:

``` YAML
#    environment:
#      - "discovery.type=single-node"    # single node
```

This can be useful to prevent memory overhead and/or other incovenient problems.


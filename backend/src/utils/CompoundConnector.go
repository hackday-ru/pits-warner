package utils

import (
  "gopkg.in/redis.v3"
  "github.com/gocql/gocql"
)

type CompoundConnector struct  {
  redisConnector redis.Client
  cassConnector gocql.ClusterConfig
}


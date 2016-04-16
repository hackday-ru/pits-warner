package utils

import (
  "gopkg.in/redis.v3"
  "github.com/gocql/gocql"
)

type CompoundConnector struct  {
  redisConnector redis.Client
  cassConnector gocql.ClusterConfig
}

func (connector *CompoundConnector) init(redisAddr, cassAddr string)  {
  connector.cassConnector = gocql.NewCluster(cassAddr)
  connector.cassConnector.Keyspace = "geodb"

  connector.redisConnector, _ = redis.NewClient(&redis.Options{
    Addr:     redisAddr,
    Password: "", // no password set
    DB:       0,  // use default DB
  })


}

package utils

import (
  "gopkg.in/redis.v3"
  "github.com/gocql/gocql"
)

type CompoundConnector struct  {
  RedisConnector *redis.Client
  CassConnector gocql.ClusterConfig
}

func (connector *CompoundConnector) Init(redisAddr, cassAddr string)  {
  //connector.CassConnector = gocql.NewCluster(cassAddr)
  //connector.CassConnector.Keyspace = "geodb"

  connector.RedisConnector = redis.NewClient(&redis.Options{
    Addr:     redisAddr,
    Password: "", // no password set
    DB:       0,  // use default DB
  })


}

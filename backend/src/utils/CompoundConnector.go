package utils

import (
  "gopkg.in/redis.v3"
  "github.com/gocql/gocql"
  "model"
  "log"
)

type CompoundConnector struct  {
  RedisConnector *redis.Client
  CassConnector *gocql.ClusterConfig
}

func (connector *CompoundConnector) Init(redisAddr, cassAddr string)  {
  connector.CassConnector = gocql.NewCluster(cassAddr)
  connector.CassConnector.Keyspace = "geodb"
  //connector.CassConnector.Consistency = gocql.Quorum

  connector.RedisConnector = redis.NewClient(&redis.Options{
    Addr:     redisAddr,
    Password: "", // no password set
    DB:       0,  // use default DB
  })
}



func (writer CompoundConnector) Write(rec model.InputRecord)  {
  acc := toString(rec.AcX) + ":" + toString(rec.AcY) + ":" + toString(rec.AcZ)


  err := writer.RedisConnector.Set(rec.Uid.String(), acc, 0).Err()
  if err != nil {
    log.Fatal(err)
  }

  writer.RedisConnector.GeoAdd(MARKERS, &redis.GeoLocation{
    Longitude:rec.Longitude, Latitude: rec.Latitude, Name: rec.Uid.String()})

  session, _ := writer.CassConnector.CreateSession()
  defer session.Close()

  session.Query(
    "INSERT INTO geodata" +
    "(id, time, geoX, geoY, geoZ, acX, acY, acZ)" +
    "values (?, ?, ?, ?, ?, ?, ?, ?)",
    rec.Uid.String(),
    rec.Timestamp,
    rec.Longitude, rec.Latitude, rec.Altitude,
    rec.AcX, rec.AcY, rec.AcZ).Exec()
}

package utils

import (
  "gopkg.in/redis.v3"
  "github.com/gocql/gocql"
  "model"
  "log"
  "strconv"
  "github.com/satori/go.uuid"
  "fmt"
)

var  conn CompoundConnector;

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
  conn = *connector
}

func GetConn() CompoundConnector{
  return conn
}

var MARKERS = "markers"

func toString(v float64) string  {
  return strconv.FormatFloat(float64(v), 'f', 5, 64)
}

func (writer CompoundConnector) ReadByLocation(c model.Coord, radius float64) model.FindResult {
  var res = []model.Coord{}
  locs, _ := writer.RedisConnector.GeoRadius(MARKERS, c.Lat, c.Lng, &redis.GeoRadiusQuery{
    Radius: radius,
    // Can be m, km, ft, or mi. Default is km.
    Unit: "m" }).Result()
  for _, e := range(locs) {
    res = append(res[:], model.Coord{
      Lat: e.Latitude,
      Lng: e.Longitude})
  }

  return model.FindResult{res}
}



func (writer CompoundConnector) Write(rec model.InputRecord)  {
  acc := toString(rec.AcX) + ":" + toString(rec.AcY) + ":" + toString(rec.AcZ)

  uid := uuid.NewV4()

  err := writer.RedisConnector.Set(uid.String(), acc, 0).Err()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(uid.String())

  err1 :=  writer.RedisConnector.GeoAdd(MARKERS, &redis.GeoLocation{
    Longitude:rec.Longitude, Latitude: rec.Latitude, Name: uid.String()}).Err()
  if err1 == nil {
    log.Fatal(err1)
  }

  session, _ := writer.CassConnector.CreateSession()
  defer session.Close()

  if err := session.Query(
    "INSERT INTO geodata" +
    "(id, timestamp, longitude, latitude, altitude, acx, acy, acz, accuracy, bearing, speed)" +
    "values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)",
    uid.String(),
    rec.Timestamp,
    rec.Longitude, rec.Latitude, rec.Altitude,
    rec.AcX, rec.AcY, rec.AcZ,
    rec.Accuracy, rec.Bearing, rec.Speed).Exec(); err != nil {
    log.Fatal(err)
  }
}

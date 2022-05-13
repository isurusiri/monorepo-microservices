package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/isurusiri/monorepo-microservices/behaviour/router"
	"github.com/isurusiri/monorepo-microservices/behaviour/storage/mongodb"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

func init() {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file", err)
	}
}

func main() {
	var mongoURI = viper.GetString("database.mongoConnectionString")
	dialInfo, err := mgo.ParseURL(mongoURI)
	if err != nil {
		panic(err)
	}

	tlsConfig := &tls.Config{}
	tlsConfig.InsecureSkipVerify = true

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal("error creating session", err)
	}

	s := &mongodb.Storage{session.DB(viper.GetString("database.mongoDbName"))}
	defer session.Close()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.URLFormat)

	router := router.InitRouter(r, s)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", viper.GetString("server.port")), router); err != nil {
		log.Fatal("error while serve http server", err)
	}
}

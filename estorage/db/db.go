package db

import (
	"os"
	"gopkg.in/mgo.v2"
	"github.com/button-tech/utils-eth-tokens-getter/estorage/db/schema"
	"gopkg.in/mgo.v2/bson"
)

var (
	host       = os.Getenv("HOST")
	database   = os.Getenv("DB")
	username   = os.Getenv("USER")
	password   = os.Getenv("PASS")
	collection = os.Getenv("COLLECTION")
)

var info = mgo.DialInfo{
	Addrs:    []string{host},
	Database: database,
	Username: username,
	Password: password,
}

func GetEthEntries() (*schema.EthEntry, error) {
	session, err := mgo.DialWithInfo(&info)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var addrs schema.EthEntry

	c := session.DB(database).C(collection)

	err = c.Find(bson.M{"currency": "eth"}).One(&addrs)
	if err != nil {
		return nil, err
	}

	return &addrs, nil
}
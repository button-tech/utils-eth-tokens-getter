package db

import (
	"github.com/button-tech/utils-eth-tokens-getter/db/schema"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

var (
	host       = os.Getenv("HOST")
	database   = os.Getenv("DB")
	username   = os.Getenv("USER")
	password   = os.Getenv("PASS")
	addressesCollection = os.Getenv("ADDRESSES_COLLECTION")
	tokensCollection = os.Getenv("TOKENS_COLLECTION")
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

	c := session.DB(database).C(addressesCollection)

	err = c.Find(bson.M{"currency": "eth"}).One(&addrs)
	if err != nil {
		return nil, err
	}

	return &addrs, nil
}

func GetTokensEntries() ([]schema.Tokens, error) {
	session, err := mgo.DialWithInfo(&info)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	var tokens []schema.Tokens

	c := session.DB(database).C(tokensCollection)

	err = c.Find(nil).All(&tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

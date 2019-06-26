package service

import (
	"encoding/base64"
	"flag"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/crypto"
	"io/ioutil"
	"os"
)

func GetEncData(dataFile string) (encData []byte) {
	flag.Parse()
	jsonFile, err := os.Open(dataFile)
	defer jsonFile.Close()
	if err != nil {
		panic(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	outs, _ := crypto.DesEncrypt(byteValue, common.Key)
	encData = make([]byte, base64.StdEncoding.EncodedLen(len(outs)))
	base64.StdEncoding.Encode(encData, outs)

	if err != nil {
		panic(err)
	}
	return
}

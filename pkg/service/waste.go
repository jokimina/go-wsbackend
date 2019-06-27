package service

import (
	"encoding/base64"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/crypto"
	"io/ioutil"
	"os"
)

func GetEncData(dataFile string) (encData []byte) {
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

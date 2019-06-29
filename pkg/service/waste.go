package service

import (
	"encoding/base64"
	"go-wsbackend/pkg/common"
	"go-wsbackend/pkg/crypto"
)

func GetEncData() (encData []byte) {
	outs, _ := crypto.DesEncrypt(*allWasteData, common.Key)
	encData = make([]byte, base64.StdEncoding.EncodedLen(len(outs)))
	base64.StdEncoding.Encode(encData, outs)
	return
}

package plumbing

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"minigit/internal/app"
	"os"
)

func NewHashObject(data []byte, objectType app.ObjectType) string {
	typedData := append([]byte(objectType), 0x00)
	typedData = append(typedData, data...)

	h := sha1.New()
	h.Write(typedData)

	objectId := fmt.Sprintf("%x", h.Sum(nil))
	if err := ioutil.WriteFile(fmt.Sprintf("%s/objects/%s", app.GitDir, objectId), typedData, os.ModePerm); err != nil {
		panic(err)
	}
	return objectId
}

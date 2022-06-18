package plumbing

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"minigit/internal/app"
	"os"
)

func getObject(objectId string, expectedType app.ObjectType) []byte {
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/objects/%s", app.GitDir, objectId))
	if err != nil {
		panic(err)
	}
	b := bytes.Split(data, []byte{0x00})
	actualType := string(b[0])
	if actualType != string(expectedType) {
		panic(fmt.Sprintf("Hash object type is different. Expected:'%s' Actual:'%s'", expectedType, actualType))
	}
	return b[1]
}

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

func CatFile(objectId string, expectedType app.ObjectType) {
	data := getObject(objectId, expectedType)
	fmt.Printf("%s", string(data))
}

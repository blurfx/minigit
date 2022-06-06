package plumbing

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"minigit/internal/app"
)

func CatFile(objectId string, expectedType app.ObjectType) {
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/objects/%s", app.GitDir, objectId))
	if err != nil {
		panic(err)
	}
	b := bytes.Split(data, []byte{0x00})
	actualType := string(b[0])
	if actualType != string(expectedType) {
		panic(fmt.Sprintf("Hash object type is different. Expected:'%s' Actual:'%s'", expectedType, actualType))
	}
	fmt.Printf("%s", string(b[1]))
}

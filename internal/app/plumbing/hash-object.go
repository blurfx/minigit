package plumbing

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"minigit/internal/app"
	"os"
)

func NewHashObject(data []byte) string {
	h := sha1.New()
	h.Write(data)
	objectId := fmt.Sprintf("%x", h.Sum(nil))
	if err := ioutil.WriteFile(fmt.Sprintf("%s/objects/%s", app.GIT_DIR, objectId), data, os.ModePerm); err != nil {
		panic(err)
	}
	return objectId
}

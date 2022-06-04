package plumbing

import (
	"fmt"
	"io/ioutil"
	"minigit/internal/app"
)

func CatFile(objectId string) {
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/objects/%s", app.GIT_DIR, objectId))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(data))
}

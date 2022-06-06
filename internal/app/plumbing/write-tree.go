package plumbing

import (
	"fmt"
	"io/ioutil"
	"minigit/internal/app"
	"minigit/internal/app/utils"
	"sort"
	"strings"
)

type TreeObject struct {
	name       string
	objectId   string
	objectType app.ObjectType
}

func WriteTree(dirname string) string {
	entries := make([]TreeObject, 0)
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		f.Name()
		relPath := fmt.Sprintf("%s/%s", dirname, f.Name())
		if utils.Contains(strings.Split(relPath, "/"), app.GitDir) {
			continue
		}

		if f.Mode().IsRegular() {
			data, err := ioutil.ReadFile(relPath)
			if err != nil {
				panic(err)
			}
			entries = append(entries, TreeObject{
				name:       f.Name(),
				objectId:   NewHashObject(data, app.ObjectTypeBlob),
				objectType: app.ObjectTypeBlob,
			})
		} else if f.IsDir() {
			entries = append(entries, TreeObject{
				name:       f.Name(),
				objectId:   WriteTree(relPath),
				objectType: app.ObjectTypeTree,
			})
		}
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].name != entries[j].name {
			return entries[i].name < entries[j].name
		}

		if entries[i].objectId != entries[j].objectId {
			return entries[i].objectId < entries[j].objectId
		}

		return entries[i].objectType < entries[j].objectType
	})

	nodes := make([]string, len(entries))
	for i := range entries {
		nodes[i] = fmt.Sprintf("%s %s %s", entries[i].name, entries[i].objectId, entries[i].objectType)
	}
	tree := strings.Join(nodes, "\n")
	return NewHashObject([]byte(tree), app.ObjectTypeTree)
}

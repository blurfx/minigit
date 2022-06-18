package plumbing

import (
	"fmt"
	"io/ioutil"
	"minigit/internal/app"
	"minigit/internal/app/utils"
	"os"
	"sort"
	"strings"
)

type TreeObject struct {
	Name       string
	ObjectId   string
	ObjectType app.ObjectType
}

func getTreeObjects(objectId string) []*TreeObject {
	tree := getObject(objectId, app.ObjectTypeTree)
	objects := strings.Split(string(tree), "\n")
	treeObjects := make([]*TreeObject, len(objects))
	for i := range objects {
		s := strings.SplitN(objects[i], " ", 3)
		treeObjects[i] = &TreeObject{
			Name:       s[0],
			ObjectId:   s[1],
			ObjectType: app.ObjectType(s[2]),
		}
	}
	return treeObjects
}

func getTree(objectId string, basePath string) map[string]string {
	result := make(map[string]string)
	objects := getTreeObjects(objectId)
	for _, obj := range objects {
		if strings.Contains(obj.Name, "/") || obj.Name == ".." || obj.Name == "." {
			panic("invalid object name")
		}
		path := basePath + obj.Name
		if obj.ObjectType == app.ObjectTypeBlob {
			result[path] = obj.ObjectId
		} else if obj.ObjectType == app.ObjectTypeTree {
			subtree := getTree(obj.ObjectId, fmt.Sprintf("%s/", path))
			for k, v := range subtree {
				result[k] = v
			}
		} else {
			panic(fmt.Sprintf("unknown tree entry '%s'", obj.ObjectType))
		}
	}
	return result
}

func ReadTree(treeObjectId string) error {
	tree := getTree(treeObjectId, "./")
	for p, objectId := range tree {
		if err := os.MkdirAll(p, os.ModePerm); err != nil {
			panic(err)
		}
		if err := ioutil.WriteFile(p, getObject(objectId, app.ObjectTypeBlob), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
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
				Name:       f.Name(),
				ObjectId:   NewHashObject(data, app.ObjectTypeBlob),
				ObjectType: app.ObjectTypeBlob,
			})
		} else if f.IsDir() {
			entries = append(entries, TreeObject{
				Name:       f.Name(),
				ObjectId:   WriteTree(relPath),
				ObjectType: app.ObjectTypeTree,
			})
		}
	}

	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Name != entries[j].Name {
			return entries[i].Name < entries[j].Name
		}

		if entries[i].ObjectId != entries[j].ObjectId {
			return entries[i].ObjectId < entries[j].ObjectId
		}

		return entries[i].ObjectType < entries[j].ObjectType
	})

	nodes := make([]string, len(entries))
	for i := range entries {
		nodes[i] = fmt.Sprintf("%s %s %s", entries[i].Name, entries[i].ObjectId, entries[i].ObjectType)
	}
	tree := strings.Join(nodes, "\n")
	return NewHashObject([]byte(tree), app.ObjectTypeTree)
}

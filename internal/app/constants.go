package app

type ObjectType string

const (
	GitDir = ".git"

	ObjectTypeBlob = ObjectType("blob")
	ObjectTypeTree = ObjectType("tree")
)

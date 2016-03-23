package db

type List struct {
	Status     string `json:"-" URIParam:"yes" oneOf:"PENDING,READY,ACTIVE,DELETED,FAILED,UNKNOWN,SUCCESS,CONFIGURING,TERMINATED"`
	DataCenter string `json:"-" URIParam:"yes"`
}

type ListRes []CreateRes

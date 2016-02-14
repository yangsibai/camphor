package models

type HashInfo struct {
	Md5    string `json:"md5" bson:"md5"`
	Sha1   string `json:"sha1" bson:"sha1"`
	Sha256 string `json:"sha256" bson:"sha256"`
	Sha512 string `json:"sha512" bson:"sha512"`
}

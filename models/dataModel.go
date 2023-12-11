package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type SelfInfo struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	NameDisplay string             `json:"name_display"`
}

type BannerInfo struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UpperTitle string             `json:"upper_title"` // ex "hello iam ..."
	LowerTitle string             `json:"lower_title"`
}

type MyPageInfo struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PageName    string             `json:"page_name"`
	Description string             `json:"description"`
	LinkData    string             `json:"link_data"`
}

type HobbyInfo struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DataName string             `json:"data_name"` // ex content writing
	Column   int                `json:"column"`    // ex 1/2/3
}

type EducationInfo struct {
	Id         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Jurusan    string             `json:"jurusan"`
	Tempat     string             `json:"tempat"`
	Kota       string             `json:"kota"`
	TahunMasuk string             `json:"tahun_masuk"`
	TahunLulus string             `json:"tahun_lulus"`
}

type ContactInfo struct {
	Id   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name"` // ex : instagram, facebook, etc
	Data string             `json:"data"` // ex : @username, link, etc
}

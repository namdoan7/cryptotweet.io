package tag

import "time"

type Tag struct {
	Id         string    `bson:"_id,omitempty"`
	Name       string    `bson:"name,omitempty" json:"name,omitempty"`
	Desciption string    `bson:"description,omitempty" json:"description,omitempty"`
	CreatedAt  time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt  time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

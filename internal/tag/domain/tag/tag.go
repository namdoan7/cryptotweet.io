package tag

import (
	"time"

	tagpb "github.com/levinhne/cryptotweet.io/internal/common/genproto/tag"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Tag struct {
	Id         string    `bson:"_id,omitempty"`
	Name       string    `bson:"name,omitempty" json:"name,omitempty"`
	Desciption string    `bson:"description,omitempty" json:"description,omitempty"`
	Type       string    `bson:"type,omitempty" json:"type,omitempty"`
	CreatedAt  time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt  time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

func (t *Tag) ToProtoMessage() *tagpb.Tag {
	return &tagpb.Tag{
		Id:          t.Id,
		Name:        t.Name,
		Description: t.Desciption,
		Type:        t.Type,
		Status:      tagpb.Tag_ACTIVE,
		CreatedAt:   &timestamppb.Timestamp{Seconds: t.CreatedAt.Unix()},
		UpdatedAt:   &timestamppb.Timestamp{Seconds: t.CreatedAt.Unix()},
	}
}

func (t *Tag) FromProtoMessage(tpb *tagpb.Tag) *Tag {
	return &Tag{
		Id:         tpb.Id,
		Name:       tpb.Name,
		Desciption: tpb.Description,
	}
}

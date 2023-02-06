package like

import (
	"bytes"
	"encoding/gob"
)

const CreateLikeTopic = "CREATE_LIKE_MSG"
const DeleteLikeTopic = "DELETE_LIKE_MSG"

type LikeMsg struct {
	Id           string
	UserId       string
	TargetId     string
	TargetType   int64
	AssociatedId string
	Time         int64
}

func (m LikeMsg) Encode() ([]byte, error) {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	if err := e.Encode(m); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (m *LikeMsg) Decode(data []byte) error {
	b := bytes.NewBuffer(data)
	d := gob.NewDecoder(b)
	if err := d.Decode(m); err != nil {
		return err
	}
	return nil
}

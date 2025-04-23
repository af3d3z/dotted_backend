package models

type Group struct {
	Id        string `json:"id"`
	CreatorId string `json:"creatorId"`
	Name      string `json:"name"`
	Photo     []byte `json:"photo"`
}

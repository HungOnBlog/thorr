package models

type ThorrOptions struct {
	File   string `json:"file"`
	Spawns int    `json:"spawns" default:"1"`
}

type Thorr struct {
	Options ThorrOptions `json:"options"`
}

func NewThorr(Options ThorrOptions) *Thorr {
	return &Thorr{
		Options: Options,
	}
}

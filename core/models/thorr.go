package models

type ThorrOptions struct {
	File string `json:"file"`
}

type Thorr struct {
	Options ThorrOptions `json:"options"`
}

func NewThorr(Options ThorrOptions) *Thorr {
	return &Thorr{
		Options: Options,
	}
}

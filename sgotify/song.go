package sgotify

import (
	"errors"
)

type Song struct {
	Title     string
	Author    string
	Id        string
	InQueue   bool
	OnSpotify bool `json:"OnSpotify,string"`
}

type Songs []Song

func (ss Songs) Contains(ID string) bool {
	for _, s := range ss {
		if s.Id == ID {
			return true
		}
	}
	return false
}

func (ss Songs) SetOnSpotify(Id string, onSpotify bool) (err error) {
	for i, _ := range ss {
		if ss[i].Id == Id {
			ss[i].OnSpotify = onSpotify
			return nil
		}
	}
	return errors.New("Not found song with Id: " + Id)
}

func (ss Songs) SongsNotOnSpotify() (_ss Songs) {
	for _, s := range ss {
		if !s.OnSpotify {
			_ss = append(_ss, s)
		}
	}
	return
}
func (ss Songs) SongsOnSpotify() (_ss Songs) {
	for _, s := range ss {
		if s.OnSpotify {
			_ss = append(_ss, s)
		}
	}
	return
}

func (ss Songs) OnSpotify(Id string) bool {
	for _, s := range ss {
		if s.Id == Id {
			return s.OnSpotify
		}
	}
	return false
}

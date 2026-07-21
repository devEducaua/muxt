package config

import (
	kdl "github.com/sblinch/kdl-go"
)

type Head struct {
	Layout Layout `kdl:"layout"`
}

type Layout struct {
	Name string `kdl:",arg"`
	Root string `kdl:"root"`
	Windows []Window `kdl:"window,multiple"`
}

type Window struct {
	Name string `kdl:",arg"`
	Panes []Pane `kdl:"pane,multiple"`
}

type Pane struct {
	Cmd string `kdl:",arg"`
	Props map[string]interface{} `kdl:",props"`
}

func ParseLayout(contents []byte) (Layout, error) {
	var h Head;
	err := kdl.Unmarshal(contents, &h);
	if err != nil {
		return Layout{}, err;
	}
	return h.Layout, nil;
}

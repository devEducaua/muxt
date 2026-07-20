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
	Attach bool `kdl:"attach"`
	Windows []Window `kdl:"window,multiple"`
}

type Window struct {
	Name string `kdl:",arg"`
	Cmd string `kdl:"cmd"`
	// TODO: add panes later
}

func ParseLayout(contents []byte) (Layout, error) {
	var h Head;
	err := kdl.Unmarshal(contents, &h);
	if err != nil {
		return Layout{}, err;
	}
	return h.Layout, nil;
}

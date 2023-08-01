package main

import "github.com/keysight/clipsgo/pkg/clips"

const (
	Free       clips.Symbol = "FREE"
	Commercial clips.Symbol = "COMMERCIAL"
	Dubious    clips.Symbol = "DUBIOUS"
)

const (
	Algorithmic clips.Symbol = "ALGORITHMIC"
	NonLinear   clips.Symbol = "NONLINEAR"
)

// ExpertSystem is a candidate expert system
type ExpertSystem struct {
	License  clips.Symbol   `json:"license"`
	Language []clips.Symbol `json:"language"`
}

// Project represents a project that might need an expert system
type Project struct {
	Logic                   clips.Symbol
	Language                []clips.Symbol
	IsExpertSystemCandidate bool
	PreferredShell          *ExpertSystem
}

func main() {
}

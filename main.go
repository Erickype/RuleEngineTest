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

var expertSystems = []*ExpertSystem{
	{
		Name:     "Clips",
		License:  Free,
		Language: []clips.Symbol{"Python", "C"},
	},
	{
		Name:     "Prolog",
		License:  Free,
		Language: []clips.Symbol{"Prolog"},
	}, {
		Name:     "Jess",
		License:  Free,
		Language: []clips.Symbol{"Java"},
	},
	{
		Name:     "Drools",
		License:  Free,
		Language: []clips.Symbol{"Java"},
	},
	{
		Name:     "Shine",
		License:  Dubious,
		Language: []clips.Symbol{"C"},
	},
}

// ExpertSystem is a candidate expert system
type ExpertSystem struct {
	Name     string
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

func insertExpertSystems(env *clips.Environment) error {
	for _, system := range expertSystems {
		_, err := env.Insert(system.Name, &system)
		if err != nil {
			return err
		}
	}
	return nil
}

func buildEnvironment(env *clips.Environment) error {
	err := env.Build(`(
	defmessage-handler ExpertSystem add-go-support ())
		(send ?self put-languages (create$ self:Languages Go)
	)`)
	if err != nil {
		return err
	}
	return nil
}

func main() {
}

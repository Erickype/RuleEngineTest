package main

import (
	_ "embed"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

//go:embed check_values.grl
var rules []byte

func main() {
	knowledgeLibrary := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(knowledgeLibrary)

	err := ruleBuilder.BuildRuleFromResource("TutorialRules", "0.0.1", pkg.NewBytesResource(rules))
	if err != nil {
		panic(err)
	}
}

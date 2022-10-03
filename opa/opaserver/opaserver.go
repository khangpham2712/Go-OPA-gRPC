package opaserver

import (
	// "github.com/open-policy-agent/opa/cmd"
	"context"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

func QueryOPAServer(input interface{}) bool {
	r := rego.New(rego.Query("data.policy.allow"),
		rego.Load([]string{"./../policy.rego", "./../data.json"}, nil))

	if r == nil {
		log.Printf("r == nil")
	} else {
		log.Printf("r != nil")
	}

	ctx := context.Background()

	preparedQuery, err := r.PrepareForEval(ctx)
	if err != nil {
		log.Fatalf("Something went wrong (1)")
	}

	result, err := preparedQuery.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatalf("Something went wrong (2)")
	}

	log.Println(result)

	if result.Allowed() {
		return true
	}

	return false

	// if err := cmd.RootCommand.Execute(); err != nil {
	// 	log.Fatalln("Some thing went wrong:", err.Error())
	// }
}

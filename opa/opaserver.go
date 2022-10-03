package opaserver

import (
	"context"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

func QueryOPAServer(input interface{}) bool {
	r := rego.New(rego.Query("output.allow"),
		rego.Load([]string{"oparules/policy.rego", "oparules/data.json"}, nil))

	ctx := context.Background()
	preparedQuery, err := r.PrepareForEval(ctx)

	if err != nil {
		log.Fatalf(err.Error())
	}

	result, err := preparedQuery.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println("result:", result)

	if result.Allowed() {
		return true
	}

	return false
}

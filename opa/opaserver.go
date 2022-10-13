package opaserver

import (
	"context"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

var r *rego.Rego

var ctx context.Context
var preparedQuery rego.PreparedEvalQuery

func RegisterOPA() {
	r = rego.New(rego.Query("data.oparules.allow"),
		rego.Load([]string{"oparules/policy.rego", "oparules/data.json"}, nil))

	ctx = context.Background()
	var err error
	preparedQuery, err = r.PrepareForEval(ctx)
	if err != nil {
		log.Fatalln("Preparation error: " + err.Error())
	}

	log.Println("OPA prepared query registered")
}

func QueryOPAServer(input interface{}) bool {
	result, err := preparedQuery.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatalln("Evaluation error: " + err.Error())
	}

	if result.Allowed() {
		return true
	}

	return false
}

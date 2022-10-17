package opaserver

import (
	"context"
	"encoding/json"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

var ctx context.Context
var preparedQuery rego.PreparedEvalQuery

func RegisterOPAQuery() {
	ctx = context.Background()
	var err error

	preparedQuery, err = rego.New(rego.Query("data.oparules"),
		rego.Load([]string{"oparules/policy.rego", "oparules/data.json"}, nil)).PrepareForEval(ctx)
	if err != nil {
		log.Fatalln("Preparation error: " + err.Error())
	}

	log.Println("OPA prepared query registered")
}

func QueryOPAServer(input interface{}) (bool, int64) {
	result, err := preparedQuery.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatalln("OPA evaluation error: " + err.Error())
	}

	var mp map[string]interface{} = result[0].Expressions[0].Value.(map[string]interface{})
	if !mp["allow"].(bool) {
		return false, 0
	}

	var payload map[string]interface{} = mp["payload"].(map[string]interface{})
	exp, err := payload["exp"].(json.Number).Int64()
	if err != nil {
		log.Fatalln("Json number parsing error: " + err.Error())
	}

	return mp["allow"].(bool), exp
}

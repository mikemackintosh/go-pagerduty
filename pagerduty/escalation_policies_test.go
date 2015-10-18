package pagerduty_test

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/mikemackintosh/go-pagerduty/pagerduty"
)

var onCall = `{
	"escalation_policy": {
		"id": "ABCDEF",
		"name": "Testing ABCDEF",
		"escalation_rules": [
			{
				"id": "RulesABC",
				"escalation_delay_in_minutes": 30,
				"rule_object": {
					"b": "c"
				}
			}
		],
		"on_call": [
			{
				"level": 1,
				"start": "today",
				"end": "tomorrow",
				"user": {
					"id": "ABC",
					"name": "Mike Yourmom",
					"email": "mike@yourmom.com",
					"user_url": "/users/ABC"
				}
			},
			{
				"level": 2,
				"start": "today",
				"end": "tomorrow",
				"user": {
					"id": "DEF",
					"name": "Is Hot",
					"email": "is@hot.com",
					"user_url": "/users/DEF"
				}
			}
		]
	}
}`

func TestEscalationPolicy_marshal(t *testing.T) {
	i := &EscalationWrapper{
		&EscalationPolicy{
			ID:   "ABCDEF",
			Name: "Testing ABCDEF",
			Rules: []EscalationRules{
				EscalationRules{
					ID:    "RulesABC",
					Delay: 30,
					Rules: map[string]string{
						"b": "c",
					},
				},
			},
			OnCall: []EscalationOnCall{
				EscalationOnCall{
					Level: 1,
					Start: "today",
					End:   "tomorrow",
					User: User{
						ID:      "ABC",
						Name:    "Mike Yourmom",
						Email:   "mike@yourmom.com",
						UserURL: "/users/ABC",
					},
				},
				EscalationOnCall{
					Level: 2,
					Start: "today",
					End:   "tomorrow",
					User: User{
						ID:      "DEF",
						Name:    "Is Hot",
						Email:   "is@hot.com",
						UserURL: "/users/DEF",
					},
				},
			},
		},
	}

	testJSONMarshal(t, i, onCall)
}

func TestEscalationPolicy_OnCall(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/escalation_policies/ABCDEF/on_call", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, onCall)
	})

	_, _, err := client.Escalations.OnCall("ABCDEF")
	if err != nil {
		t.Errorf("Escalation.OnCall returned error: %v", err)
	}

}

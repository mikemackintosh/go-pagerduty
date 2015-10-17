package pagerduty

import (
	"errors"
	"net/http"
)

// The EscalationService struct is instantiated in the pagerduty
// struct instantiation and contains a reference back to the pagerduty
// client
type EscalationService struct {
	client *Client
}

// EscalationPolicies type
type EscalationPolicies struct {
	Policies []*EscalationPolicy `json:"escalation_policies,omitempty"`
	Limit    int                 `json:"limit,omitempty"`
	Offset   int                 `json:"offset,omitempty"`
	Total    int                 `json:"total,omitempty"`
}

// EscalationPolicy type
type EscalationPolicy struct {
	ID               string            `json:"id,omitempty"`
	Name             string            `json:"name,omitempty"`
	Rules            []EscalationRules `json:"escalation_rules,omitempty"`
	Services         Services          `json:"services,omitempty"`
	EscalationOnCall string            `json:"on_call,omitempty"`
	Loops            string            `json:"num_loops,omitempty"`
}

// EscalationRules type
type EscalationRules struct {
	Policy []*EscalationPolicy `json:"schedules,omitempty"`
	ID     int                 `json:"limit,omitempty"`
	Delay  int                 `json:"escalation_delay_in_minutes,omitempty"`
	Rules  []map[string]string `json:"rule_object,omitempty"`
}

// EscalationOnCall type
type EscalationOnCall struct {
	Level int    `json:"id,omitempty"`
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
	Users User   `json:"user,omitempty"`
}

// EscalationOptions provides optional parameters to list requests
type EscalationOptions struct {
	Query       string `url:"query,omitempty"`
	RequesterId string `url:"requester_id,omitempty"`
}

type EscalationPolicyWrapper struct {
	EscalationPolicy *EscalationPolicy `json:escalation_policy,omitempty"`
}

// List returns a list of schedules
func (s *EscalationService) List(opt *EscalationOptions) (*EscalationPolicies, *http.Response, error) {
	uri, err := addOptions("escalation_policies", opt)
	if err != nil {
		return nil, nil, err
	}

	escalationpolicies := new(EscalationPolicies)
	res, err := s.client.Get(uri, escalationpolicies)
	if err != nil {
		return nil, res, err
	}

	return escalationpolicies, res, err
}

// Get returns a single schedule by id if found
func (s *EscalationService) Get(id string) (*EscalationPolicy, *http.Response, error) {
	wrapper := new(EscalationPolicyWrapper)

	res, err := s.client.Get("escalation_policies/"+id, wrapper)
	if err != nil {
		return nil, res, err
	}

	if wrapper.EscalationPolicy == nil {
		return nil, res, errors.New("pagerduty: schedule json object nil")
	}

	return wrapper.EscalationPolicy, res, nil
}

// OnCall returns a single schedule by id if found for oncall
func (s *EscalationService) OnCall(id string) (*EscalationPolicy, *http.Response, error) {
	wrapper := new(EscalationPolicyWrapper)

	res, err := s.client.Get("escalation_policies/"+id+"/on_call", wrapper)
	if err != nil {
		return nil, res, err
	}

	if wrapper.EscalationPolicy == nil {
		return nil, res, errors.New("pagerduty: schedule json object nil")
	}

	return wrapper.EscalationPolicy, res, nil
}

package v1alpha1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// LabelSelectorOperator is the set of operators that can be used in a selector requirement.
type LabelSelectorOperator string

const (
	// AutopilotRuleResourceName is the name of the singular AutopilotRule objects
	AutopilotRuleResourceName = "autopilotrule"
	// AutopilotRuleResourceShortName is the short name for AutopilotRule objects
	AutopilotRuleResourceShortName = "ar"

	// AutopilotRuleResourcePlural is the name of the plural AutopilotRule objects
	AutopilotRuleResourcePlural = "autopilotrules"

	// LabelSelectorOpIn is operator where the key must have one of the values
	LabelSelectorOpIn LabelSelectorOperator = "In"
	// LabelSelectorOpNotIn is operator where the key must not have any of the values
	LabelSelectorOpNotIn LabelSelectorOperator = "NotIn"
	// LabelSelectorOpExists is operator where the key must exist
	LabelSelectorOpExists LabelSelectorOperator = "Exists"
	// LabelSelectorOpDoesNotExist is operator where the key must not exist
	LabelSelectorOpDoesNotExist LabelSelectorOperator = "DoesNotExist"
	// LabelSelectorOpGt is operator where the key must be greater than the values
	LabelSelectorOpGt LabelSelectorOperator = "Gt"
	// LabelSelectorOpGtEq is operator where the key must be greater than or equal to the values
	LabelSelectorOpGtEq LabelSelectorOperator = "GtEq"
	// LabelSelectorOpLt is operator where the key must be less than the values
	LabelSelectorOpLt LabelSelectorOperator = "Lt"
	// LabelSelectorOpLtEq is operator where the key must be less than or equal to the values
	LabelSelectorOpLtEq LabelSelectorOperator = "LtEq"
)

// LabelSelectorRequirement is a selector that contains values, a key, and an operator that
// relates the key and values.
type LabelSelectorRequirement struct {
	// key is the label key that the selector applies to.
	// +patchMergeKey=key
	// +patchStrategy=merge
	Key string `json:"key"`
	// operator represents a key's relationship to a set of values.
	// Valid operators are In, NotIn, Exists, DoesNotExist, Lt and Gt.
	Operator LabelSelectorOperator `json:"operator"`
	// values is an array of string values. If the operator is In or NotIn,
	// the values array must be non-empty. If the operator is Exists or DoesNotExist,
	// the values array must be empty. This array is replaced during a strategic
	// merge patch.
	// +optional
	Values []string `json:"values"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AutopilotRule represents pairing with other clusters
type AutopilotRule struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`
	Spec            AutopilotRuleSpec `json:"spec"`
	// TODO: add status
	//Status          AutopilotRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AutopilotRuleList is a list of AutopilotRule objects in Kubernetes
type AutopilotRuleList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []AutopilotRule `json:"items"`
}

// AutopilotRuleSpec is the spec to create the cluster pair
type AutopilotRuleSpec struct {
	// Weight defines the weight of the rule which allows to break the tie with other conflicting policies. A rule with
	// higher weight wins over one with lower weight.
	// (optional)
	Weight int64 `json:"weight,omitempty"`
	// PollInterval defined the interval in seconds at which the conditions for the
	// rule are queried from the monitoring provider
	PollInterval int64 `json:"pollInterval,omitempty"`
	// Enforcement specifies the enforcement type for rule. Can take values: required or preferred.
	// (optional)
	Enforcement EnforcementType `json:"enforcement,omitempty"`
	// Selector allows to select the objects that are relevant with this rule using label selection
	Selector RuleObjectSelector `json:"selector"`
	// NamespaceSelector allows to select namespaces affecting the rule by labels:w
	NamespaceSelector RuleObjectSelector `json:"namespaceSelector"`
	// Conditions are the conditions to check on the rule objects
	Conditions RuleConditions `json:"conditions"`
	// Actions are the actions to run for the rule when the conditions are met
	Actions []*RuleAction `json:"actions"`
	// ActionsCoolDownPeriod is the duration in seconds for which autopilot will not
	// re-trigger any actions once they have been executed.
	ActionsCoolDownPeriod int64 `json:"actionsCoolDownPeriod,omitempty"`
}

// RuleObjectSelector defines an object for the rule
type RuleObjectSelector struct {
	// LabelSelector selects the rule objects
	meta.LabelSelector
}

// RuleConditions defines the conditions for the rule
type RuleConditions struct {
	// Expressions are the actual rule conditions
	Expressions []*LabelSelectorRequirement `json:"expressions,omitempty"`
	// For is the duration in seconds for which the conditions must hold true
	For int64 `json:"for,omitempty"`
	// Type is the condition type
	// If not provided, the controller for the CRD will pick the default type
	Type AutopilotRuleConditionType `json:"type,omitempty"`
	// Provider is an optional provider for the above condition type
	// If not provided, the controller for the CRD will pick the default provider
	Provider string `json:"provider,omitempty"`
}

// RuleAction defines an action for the rule
type RuleAction struct {
	// ObjectName is the name of the rule
	Name string `json:"name"`
	// Params are the opaque paramters that will be used for the above action
	Params map[string]string `json:"params"`
}

// AutopilotRuleStatusType is the type for rule statuses
type AutopilotRuleStatusType string

const (
	// AutopilotRuleConditonMet is for when the conditions in rule are met
	AutopilotRuleConditonMet AutopilotRuleStatusType = "ConditionMet"
	// AutopilotRuleActionFailed is when an action for a rule has failed
	AutopilotRuleActionFailed AutopilotRuleStatusType = "ActionFailed"
	// AutopilotRuleActionTriggered is when an action for a rule has triggerred
	AutopilotRuleActionTriggered AutopilotRuleStatusType = "ActionTriggered"
	// AutopilotRuleActionSuccessful is when an action for a rule is successful
	AutopilotRuleActionSuccessful AutopilotRuleStatusType = "ActionSuccessful"
)

// AutopilotRuleConditionType defines the type of a condition in a rule
type AutopilotRuleConditionType string

const (
	// RuleConditionMetrics is a monitoring type of condition in a rule
	RuleConditionMetrics AutopilotRuleConditionType = "monitoring"
)

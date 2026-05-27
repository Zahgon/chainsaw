package v1alpha1

// CatchFinally defines actions to be executed in catch, finally and cleanup blocks.
// +kubebuilder:oneOf:={required:{command}}
// +kubebuilder:oneOf:={required:{delete}}
// +kubebuilder:oneOf:={required:{describe}}
// +kubebuilder:oneOf:={required:{events}}
// +kubebuilder:oneOf:={required:{get}}
// +kubebuilder:oneOf:={required:{podLogs}}
// +kubebuilder:oneOf:={required:{script}}
// +kubebuilder:oneOf:={required:{sleep}}
// +kubebuilder:oneOf:={required:{wait}}
type CatchFinally struct {
	// Description contains a description of the operation.
	// +optional
	Description string `json:"description,omitempty"`

	// Compiler defines the default compiler to use when evaluating expressions.
	// +optional
	Compiler *Compiler `json:"compiler,omitempty"`

	// PodLogs determines the pod logs collector to execute.
	// +optional
	PodLogs *PodLogs `json:"podLogs,omitempty"`

	// Events determines the events collector to execute.
	// +optional
	Events *Events `json:"events,omitempty"`

	// Describe determines the resource describe collector to execute.
	// +optional
	Describe *Describe `json:"describe,omitempty"`

	// Wait determines the resource wait collector to execute.
	// +optional
	Wait *Wait `json:"wait,omitempty"`

	// Get determines the resource get collector to execute.
	// +optional
	Get *Get `json:"get,omitempty"`

	// Delete represents a deletion operation.
	// +optional
	Delete *Delete `json:"delete,omitempty"`

	// Command defines a command to run.
	// +optional
	Command *Command `json:"command,omitempty"`

	// Script defines a script to run.
	// +optional
	Script *Script `json:"script,omitempty"`

	// Sleep defines zzzz.
	// +optional
	Sleep *Sleep `json:"sleep,omitempty"`
}

func (f *CatchFinally) Bindings() []Binding { _ = "STUB: not implemented"; return nil }

func (f *CatchFinally) Outputs() []Output { _ = "STUB: not implemented"; return nil }

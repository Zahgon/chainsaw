package v1alpha1

// Projection can be any type.
// +k8s:deepcopy-gen=false
// +kubebuilder:validation:XPreserveUnknownFields
// +kubebuilder:validation:Type:=""
type Projection struct {
	_value any
}

func NewProjection(value any) Projection { _ = "STUB: not implemented"; return *new(Projection) }

func (a *Projection) Value() any { _ = "STUB: not implemented"; return *new(any) }

func (a *Projection) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (a *Projection) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (in *Projection) DeepCopyInto(out *Projection) { _ = "STUB: not implemented"; return }

func (in *Projection) DeepCopy() *Projection { _ = "STUB: not implemented"; return nil }

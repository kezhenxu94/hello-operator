package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HelloOperatorSpec defines the desired state of HelloOperator
type HelloOperatorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Name string `json:"name,omitempty"`
}

// HelloOperatorStatus defines the observed state of HelloOperator
type HelloOperatorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HelloOperator is the Schema for the hellooperators API
type HelloOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelloOperatorSpec   `json:"spec,omitempty"`
	Status HelloOperatorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HelloOperatorList contains a list of HelloOperator
type HelloOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HelloOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HelloOperator{}, &HelloOperatorList{})
}

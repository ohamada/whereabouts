package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// ClusterIPReservationSpec defines the desired state of ClusterIPReservation
type ClusterIPReservationSpec struct {
  Information string `json:"information"`
}

// +kubebuilder:object:root=true

// ClusterIPReservation is the Schema for the ClusterIPReservations API
type ClusterIPReservation struct {
  metav1.TypeMeta   `json:",inline"`
  metav1.ObjectMeta `json:"metadata,omitempty"`

  Spec ClusterIPReservationSpec `json:"spec"`
}

// +kubebuilder:object:root=true

// ClusterIPReservationList contains a list of ClusterIPReservation
type ClusterIPReservationList struct {
  metav1.TypeMeta `json:",inline"`
  metav1.ListMeta `json:"metadata,omitempty"`

  Items []ClusterIPReservation `json:"items"`
}

func init() {
  SchemeBuilder.Register(&ClusterIPReservation{}, &ClusterIPReservationList{})
}

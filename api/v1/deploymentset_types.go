/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DeploymentSetSpec defines the desired state of DeploymentSet
type DeploymentSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of DeploymentSet. Edit DeploymentSet_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// DeploymentSetStatus defines the observed state of DeploymentSet
type DeploymentSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// DeploymentSet is the Schema for the deploymentsets API
type DeploymentSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentSetSpec   `json:"spec,omitempty"`
	Status DeploymentSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentSetList contains a list of DeploymentSet
type DeploymentSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeploymentSet{}, &DeploymentSetList{})
}

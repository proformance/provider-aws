/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type StandardsSubscriptionObservation struct {

	// The ARN of a resource that represents your subscription to a supported standard.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type StandardsSubscriptionParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ARN of a standard - see below.
	// +kubebuilder:validation:Required
	StandardsArn *string `json:"standardsArn" tf:"standards_arn,omitempty"`
}

// StandardsSubscriptionSpec defines the desired state of StandardsSubscription
type StandardsSubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StandardsSubscriptionParameters `json:"forProvider"`
}

// StandardsSubscriptionStatus defines the observed state of StandardsSubscription.
type StandardsSubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StandardsSubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// StandardsSubscription is the Schema for the StandardsSubscriptions API. Subscribes to a Security Hub standard.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type StandardsSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StandardsSubscriptionSpec   `json:"spec"`
	Status            StandardsSubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StandardsSubscriptionList contains a list of StandardsSubscriptions
type StandardsSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StandardsSubscription `json:"items"`
}

// Repository type metadata.
var (
	StandardsSubscription_Kind             = "StandardsSubscription"
	StandardsSubscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StandardsSubscription_Kind}.String()
	StandardsSubscription_KindAPIVersion   = StandardsSubscription_Kind + "." + CRDGroupVersion.String()
	StandardsSubscription_GroupVersionKind = CRDGroupVersion.WithKind(StandardsSubscription_Kind)
)

func init() {
	SchemeBuilder.Register(&StandardsSubscription{}, &StandardsSubscriptionList{})
}

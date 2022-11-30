/*
Copyright 2022 Max T.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import resource "github.com/crossplane/crossplane-runtime/pkg/resource"

// GetItems of this ForwardingRuleList.
func (l *ForwardingRuleList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

// GetItems of this RegionTargetTCPProxyList.
func (l *RegionTargetTCPProxyList) GetItems() []resource.Managed {
	items := make([]resource.Managed, len(l.Items))
	for i := range l.Items {
		items[i] = &l.Items[i]
	}
	return items
}

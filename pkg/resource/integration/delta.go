// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package integration

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.APIID, b.ko.Spec.APIID) {
		delta.Add("Spec.APIID", a.ko.Spec.APIID, b.ko.Spec.APIID)
	} else if a.ko.Spec.APIID != nil && b.ko.Spec.APIID != nil {
		if *a.ko.Spec.APIID != *b.ko.Spec.APIID {
			delta.Add("Spec.APIID", a.ko.Spec.APIID, b.ko.Spec.APIID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ConnectionID, b.ko.Spec.ConnectionID) {
		delta.Add("Spec.ConnectionID", a.ko.Spec.ConnectionID, b.ko.Spec.ConnectionID)
	} else if a.ko.Spec.ConnectionID != nil && b.ko.Spec.ConnectionID != nil {
		if *a.ko.Spec.ConnectionID != *b.ko.Spec.ConnectionID {
			delta.Add("Spec.ConnectionID", a.ko.Spec.ConnectionID, b.ko.Spec.ConnectionID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ConnectionType, b.ko.Spec.ConnectionType) {
		delta.Add("Spec.ConnectionType", a.ko.Spec.ConnectionType, b.ko.Spec.ConnectionType)
	} else if a.ko.Spec.ConnectionType != nil && b.ko.Spec.ConnectionType != nil {
		if *a.ko.Spec.ConnectionType != *b.ko.Spec.ConnectionType {
			delta.Add("Spec.ConnectionType", a.ko.Spec.ConnectionType, b.ko.Spec.ConnectionType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ContentHandlingStrategy, b.ko.Spec.ContentHandlingStrategy) {
		delta.Add("Spec.ContentHandlingStrategy", a.ko.Spec.ContentHandlingStrategy, b.ko.Spec.ContentHandlingStrategy)
	} else if a.ko.Spec.ContentHandlingStrategy != nil && b.ko.Spec.ContentHandlingStrategy != nil {
		if *a.ko.Spec.ContentHandlingStrategy != *b.ko.Spec.ContentHandlingStrategy {
			delta.Add("Spec.ContentHandlingStrategy", a.ko.Spec.ContentHandlingStrategy, b.ko.Spec.ContentHandlingStrategy)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.CredentialsARN, b.ko.Spec.CredentialsARN) {
		delta.Add("Spec.CredentialsARN", a.ko.Spec.CredentialsARN, b.ko.Spec.CredentialsARN)
	} else if a.ko.Spec.CredentialsARN != nil && b.ko.Spec.CredentialsARN != nil {
		if *a.ko.Spec.CredentialsARN != *b.ko.Spec.CredentialsARN {
			delta.Add("Spec.CredentialsARN", a.ko.Spec.CredentialsARN, b.ko.Spec.CredentialsARN)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Description, b.ko.Spec.Description) {
		delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
	} else if a.ko.Spec.Description != nil && b.ko.Spec.Description != nil {
		if *a.ko.Spec.Description != *b.ko.Spec.Description {
			delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.IntegrationMethod, b.ko.Spec.IntegrationMethod) {
		delta.Add("Spec.IntegrationMethod", a.ko.Spec.IntegrationMethod, b.ko.Spec.IntegrationMethod)
	} else if a.ko.Spec.IntegrationMethod != nil && b.ko.Spec.IntegrationMethod != nil {
		if *a.ko.Spec.IntegrationMethod != *b.ko.Spec.IntegrationMethod {
			delta.Add("Spec.IntegrationMethod", a.ko.Spec.IntegrationMethod, b.ko.Spec.IntegrationMethod)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.IntegrationSubtype, b.ko.Spec.IntegrationSubtype) {
		delta.Add("Spec.IntegrationSubtype", a.ko.Spec.IntegrationSubtype, b.ko.Spec.IntegrationSubtype)
	} else if a.ko.Spec.IntegrationSubtype != nil && b.ko.Spec.IntegrationSubtype != nil {
		if *a.ko.Spec.IntegrationSubtype != *b.ko.Spec.IntegrationSubtype {
			delta.Add("Spec.IntegrationSubtype", a.ko.Spec.IntegrationSubtype, b.ko.Spec.IntegrationSubtype)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.IntegrationType, b.ko.Spec.IntegrationType) {
		delta.Add("Spec.IntegrationType", a.ko.Spec.IntegrationType, b.ko.Spec.IntegrationType)
	} else if a.ko.Spec.IntegrationType != nil && b.ko.Spec.IntegrationType != nil {
		if *a.ko.Spec.IntegrationType != *b.ko.Spec.IntegrationType {
			delta.Add("Spec.IntegrationType", a.ko.Spec.IntegrationType, b.ko.Spec.IntegrationType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.IntegrationURI, b.ko.Spec.IntegrationURI) {
		delta.Add("Spec.IntegrationURI", a.ko.Spec.IntegrationURI, b.ko.Spec.IntegrationURI)
	} else if a.ko.Spec.IntegrationURI != nil && b.ko.Spec.IntegrationURI != nil {
		if *a.ko.Spec.IntegrationURI != *b.ko.Spec.IntegrationURI {
			delta.Add("Spec.IntegrationURI", a.ko.Spec.IntegrationURI, b.ko.Spec.IntegrationURI)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PassthroughBehavior, b.ko.Spec.PassthroughBehavior) {
		delta.Add("Spec.PassthroughBehavior", a.ko.Spec.PassthroughBehavior, b.ko.Spec.PassthroughBehavior)
	} else if a.ko.Spec.PassthroughBehavior != nil && b.ko.Spec.PassthroughBehavior != nil {
		if *a.ko.Spec.PassthroughBehavior != *b.ko.Spec.PassthroughBehavior {
			delta.Add("Spec.PassthroughBehavior", a.ko.Spec.PassthroughBehavior, b.ko.Spec.PassthroughBehavior)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.PayloadFormatVersion, b.ko.Spec.PayloadFormatVersion) {
		delta.Add("Spec.PayloadFormatVersion", a.ko.Spec.PayloadFormatVersion, b.ko.Spec.PayloadFormatVersion)
	} else if a.ko.Spec.PayloadFormatVersion != nil && b.ko.Spec.PayloadFormatVersion != nil {
		if *a.ko.Spec.PayloadFormatVersion != *b.ko.Spec.PayloadFormatVersion {
			delta.Add("Spec.PayloadFormatVersion", a.ko.Spec.PayloadFormatVersion, b.ko.Spec.PayloadFormatVersion)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RequestParameters, b.ko.Spec.RequestParameters) {
		delta.Add("Spec.RequestParameters", a.ko.Spec.RequestParameters, b.ko.Spec.RequestParameters)
	} else if a.ko.Spec.RequestParameters != nil && b.ko.Spec.RequestParameters != nil {
		if !ackcompare.MapStringStringPEqual(a.ko.Spec.RequestParameters, b.ko.Spec.RequestParameters) {
			delta.Add("Spec.RequestParameters", a.ko.Spec.RequestParameters, b.ko.Spec.RequestParameters)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.RequestTemplates, b.ko.Spec.RequestTemplates) {
		delta.Add("Spec.RequestTemplates", a.ko.Spec.RequestTemplates, b.ko.Spec.RequestTemplates)
	} else if a.ko.Spec.RequestTemplates != nil && b.ko.Spec.RequestTemplates != nil {
		if !ackcompare.MapStringStringPEqual(a.ko.Spec.RequestTemplates, b.ko.Spec.RequestTemplates) {
			delta.Add("Spec.RequestTemplates", a.ko.Spec.RequestTemplates, b.ko.Spec.RequestTemplates)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.ResponseParameters, b.ko.Spec.ResponseParameters) {
		delta.Add("Spec.ResponseParameters", a.ko.Spec.ResponseParameters, b.ko.Spec.ResponseParameters)
	} else if a.ko.Spec.ResponseParameters != nil && b.ko.Spec.ResponseParameters != nil {
		if !reflect.DeepEqual(a.ko.Spec.ResponseParameters, b.ko.Spec.ResponseParameters) {
			delta.Add("Spec.ResponseParameters", a.ko.Spec.ResponseParameters, b.ko.Spec.ResponseParameters)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TemplateSelectionExpression, b.ko.Spec.TemplateSelectionExpression) {
		delta.Add("Spec.TemplateSelectionExpression", a.ko.Spec.TemplateSelectionExpression, b.ko.Spec.TemplateSelectionExpression)
	} else if a.ko.Spec.TemplateSelectionExpression != nil && b.ko.Spec.TemplateSelectionExpression != nil {
		if *a.ko.Spec.TemplateSelectionExpression != *b.ko.Spec.TemplateSelectionExpression {
			delta.Add("Spec.TemplateSelectionExpression", a.ko.Spec.TemplateSelectionExpression, b.ko.Spec.TemplateSelectionExpression)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TimeoutInMillis, b.ko.Spec.TimeoutInMillis) {
		delta.Add("Spec.TimeoutInMillis", a.ko.Spec.TimeoutInMillis, b.ko.Spec.TimeoutInMillis)
	} else if a.ko.Spec.TimeoutInMillis != nil && b.ko.Spec.TimeoutInMillis != nil {
		if *a.ko.Spec.TimeoutInMillis != *b.ko.Spec.TimeoutInMillis {
			delta.Add("Spec.TimeoutInMillis", a.ko.Spec.TimeoutInMillis, b.ko.Spec.TimeoutInMillis)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.TLSConfig, b.ko.Spec.TLSConfig) {
		delta.Add("Spec.TLSConfig", a.ko.Spec.TLSConfig, b.ko.Spec.TLSConfig)
	} else if a.ko.Spec.TLSConfig != nil && b.ko.Spec.TLSConfig != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.TLSConfig.ServerNameToVerify, b.ko.Spec.TLSConfig.ServerNameToVerify) {
			delta.Add("Spec.TLSConfig.ServerNameToVerify", a.ko.Spec.TLSConfig.ServerNameToVerify, b.ko.Spec.TLSConfig.ServerNameToVerify)
		} else if a.ko.Spec.TLSConfig.ServerNameToVerify != nil && b.ko.Spec.TLSConfig.ServerNameToVerify != nil {
			if *a.ko.Spec.TLSConfig.ServerNameToVerify != *b.ko.Spec.TLSConfig.ServerNameToVerify {
				delta.Add("Spec.TLSConfig.ServerNameToVerify", a.ko.Spec.TLSConfig.ServerNameToVerify, b.ko.Spec.TLSConfig.ServerNameToVerify)
			}
		}
	}

	return delta
}

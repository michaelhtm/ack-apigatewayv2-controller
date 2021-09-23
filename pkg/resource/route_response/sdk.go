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

package route_response

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/apigatewayv2-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ApiGatewayV2{}
	_ = &svcapitypes.RouteResponse{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.GetRouteResponseOutput
	resp, err = rm.sdkapi.GetRouteResponseWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "GetRouteResponse", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "NotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.ModelSelectionExpression != nil {
		ko.Spec.ModelSelectionExpression = resp.ModelSelectionExpression
	} else {
		ko.Spec.ModelSelectionExpression = nil
	}
	if resp.ResponseModels != nil {
		f1 := map[string]*string{}
		for f1key, f1valiter := range resp.ResponseModels {
			var f1val string
			f1val = *f1valiter
			f1[f1key] = &f1val
		}
		ko.Spec.ResponseModels = f1
	} else {
		ko.Spec.ResponseModels = nil
	}
	if resp.ResponseParameters != nil {
		f2 := map[string]*svcapitypes.ParameterConstraints{}
		for f2key, f2valiter := range resp.ResponseParameters {
			f2val := &svcapitypes.ParameterConstraints{}
			if f2valiter.Required != nil {
				f2val.Required = f2valiter.Required
			}
			f2[f2key] = f2val
		}
		ko.Spec.ResponseParameters = f2
	} else {
		ko.Spec.ResponseParameters = nil
	}
	if resp.RouteResponseId != nil {
		ko.Status.RouteResponseID = resp.RouteResponseId
	} else {
		ko.Status.RouteResponseID = nil
	}
	if resp.RouteResponseKey != nil {
		ko.Spec.RouteResponseKey = resp.RouteResponseKey
	} else {
		ko.Spec.RouteResponseKey = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Status.RouteResponseID == nil || r.ko.Spec.APIID == nil || r.ko.Spec.RouteID == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.GetRouteResponseInput, error) {
	res := &svcsdk.GetRouteResponseInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.RouteID != nil {
		res.SetRouteId(*r.ko.Spec.RouteID)
	}
	if r.ko.Status.RouteResponseID != nil {
		res.SetRouteResponseId(*r.ko.Status.RouteResponseID)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateRouteResponseOutput
	_ = resp
	resp, err = rm.sdkapi.CreateRouteResponseWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateRouteResponse", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.ModelSelectionExpression != nil {
		ko.Spec.ModelSelectionExpression = resp.ModelSelectionExpression
	} else {
		ko.Spec.ModelSelectionExpression = nil
	}
	if resp.ResponseModels != nil {
		f1 := map[string]*string{}
		for f1key, f1valiter := range resp.ResponseModels {
			var f1val string
			f1val = *f1valiter
			f1[f1key] = &f1val
		}
		ko.Spec.ResponseModels = f1
	} else {
		ko.Spec.ResponseModels = nil
	}
	if resp.ResponseParameters != nil {
		f2 := map[string]*svcapitypes.ParameterConstraints{}
		for f2key, f2valiter := range resp.ResponseParameters {
			f2val := &svcapitypes.ParameterConstraints{}
			if f2valiter.Required != nil {
				f2val.Required = f2valiter.Required
			}
			f2[f2key] = f2val
		}
		ko.Spec.ResponseParameters = f2
	} else {
		ko.Spec.ResponseParameters = nil
	}
	if resp.RouteResponseId != nil {
		ko.Status.RouteResponseID = resp.RouteResponseId
	} else {
		ko.Status.RouteResponseID = nil
	}
	if resp.RouteResponseKey != nil {
		ko.Spec.RouteResponseKey = resp.RouteResponseKey
	} else {
		ko.Spec.RouteResponseKey = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateRouteResponseInput, error) {
	res := &svcsdk.CreateRouteResponseInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.ModelSelectionExpression != nil {
		res.SetModelSelectionExpression(*r.ko.Spec.ModelSelectionExpression)
	}
	if r.ko.Spec.ResponseModels != nil {
		f2 := map[string]*string{}
		for f2key, f2valiter := range r.ko.Spec.ResponseModels {
			var f2val string
			f2val = *f2valiter
			f2[f2key] = &f2val
		}
		res.SetResponseModels(f2)
	}
	if r.ko.Spec.ResponseParameters != nil {
		f3 := map[string]*svcsdk.ParameterConstraints{}
		for f3key, f3valiter := range r.ko.Spec.ResponseParameters {
			f3val := &svcsdk.ParameterConstraints{}
			if f3valiter.Required != nil {
				f3val.SetRequired(*f3valiter.Required)
			}
			f3[f3key] = f3val
		}
		res.SetResponseParameters(f3)
	}
	if r.ko.Spec.RouteID != nil {
		res.SetRouteId(*r.ko.Spec.RouteID)
	}
	if r.ko.Spec.RouteResponseKey != nil {
		res.SetRouteResponseKey(*r.ko.Spec.RouteResponseKey)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer exit(err)
	input, err := rm.newUpdateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.UpdateRouteResponseOutput
	_ = resp
	resp, err = rm.sdkapi.UpdateRouteResponseWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateRouteResponse", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.ModelSelectionExpression != nil {
		ko.Spec.ModelSelectionExpression = resp.ModelSelectionExpression
	} else {
		ko.Spec.ModelSelectionExpression = nil
	}
	if resp.ResponseModels != nil {
		f1 := map[string]*string{}
		for f1key, f1valiter := range resp.ResponseModels {
			var f1val string
			f1val = *f1valiter
			f1[f1key] = &f1val
		}
		ko.Spec.ResponseModels = f1
	} else {
		ko.Spec.ResponseModels = nil
	}
	if resp.ResponseParameters != nil {
		f2 := map[string]*svcapitypes.ParameterConstraints{}
		for f2key, f2valiter := range resp.ResponseParameters {
			f2val := &svcapitypes.ParameterConstraints{}
			if f2valiter.Required != nil {
				f2val.Required = f2valiter.Required
			}
			f2[f2key] = f2val
		}
		ko.Spec.ResponseParameters = f2
	} else {
		ko.Spec.ResponseParameters = nil
	}
	if resp.RouteResponseId != nil {
		ko.Status.RouteResponseID = resp.RouteResponseId
	} else {
		ko.Status.RouteResponseID = nil
	}
	if resp.RouteResponseKey != nil {
		ko.Spec.RouteResponseKey = resp.RouteResponseKey
	} else {
		ko.Spec.RouteResponseKey = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.UpdateRouteResponseInput, error) {
	res := &svcsdk.UpdateRouteResponseInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.ModelSelectionExpression != nil {
		res.SetModelSelectionExpression(*r.ko.Spec.ModelSelectionExpression)
	}
	if r.ko.Spec.ResponseModels != nil {
		f2 := map[string]*string{}
		for f2key, f2valiter := range r.ko.Spec.ResponseModels {
			var f2val string
			f2val = *f2valiter
			f2[f2key] = &f2val
		}
		res.SetResponseModels(f2)
	}
	if r.ko.Spec.ResponseParameters != nil {
		f3 := map[string]*svcsdk.ParameterConstraints{}
		for f3key, f3valiter := range r.ko.Spec.ResponseParameters {
			f3val := &svcsdk.ParameterConstraints{}
			if f3valiter.Required != nil {
				f3val.SetRequired(*f3valiter.Required)
			}
			f3[f3key] = f3val
		}
		res.SetResponseParameters(f3)
	}
	if r.ko.Spec.RouteID != nil {
		res.SetRouteId(*r.ko.Spec.RouteID)
	}
	if r.ko.Status.RouteResponseID != nil {
		res.SetRouteResponseId(*r.ko.Status.RouteResponseID)
	}
	if r.ko.Spec.RouteResponseKey != nil {
		res.SetRouteResponseKey(*r.ko.Spec.RouteResponseKey)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteRouteResponseOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteRouteResponseWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteRouteResponse", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteRouteResponseInput, error) {
	res := &svcsdk.DeleteRouteResponseInput{}

	if r.ko.Spec.APIID != nil {
		res.SetApiId(*r.ko.Spec.APIID)
	}
	if r.ko.Spec.RouteID != nil {
		res.SetRouteId(*r.ko.Spec.RouteID)
	}
	if r.ko.Status.RouteResponseID != nil {
		res.SetRouteResponseId(*r.ko.Status.RouteResponseID)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.RouteResponse,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}

	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}

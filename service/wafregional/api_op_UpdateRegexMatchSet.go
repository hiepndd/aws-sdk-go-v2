// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafregional

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateRegexMatchSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The RegexMatchSetId of the RegexMatchSet that you want to update. RegexMatchSetId
	// is returned by CreateRegexMatchSet and by ListRegexMatchSets.
	//
	// RegexMatchSetId is a required field
	RegexMatchSetId *string `min:"1" type:"string" required:"true"`

	// An array of RegexMatchSetUpdate objects that you want to insert into or delete
	// from a RegexMatchSet. For more information, see RegexMatchTuple.
	//
	// Updates is a required field
	Updates []RegexMatchSetUpdate `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateRegexMatchSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateRegexMatchSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateRegexMatchSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.RegexMatchSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RegexMatchSetId"))
	}
	if s.RegexMatchSetId != nil && len(*s.RegexMatchSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RegexMatchSetId", 1))
	}

	if s.Updates == nil {
		invalidParams.Add(aws.NewErrParamRequired("Updates"))
	}
	if s.Updates != nil && len(s.Updates) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Updates", 1))
	}
	if s.Updates != nil {
		for i, v := range s.Updates {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Updates", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateRegexMatchSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the UpdateRegexMatchSet request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateRegexMatchSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateRegexMatchSet = "UpdateRegexMatchSet"

// UpdateRegexMatchSetRequest returns a request value for making API operation for
// AWS WAF Regional.
//
// Inserts or deletes RegexMatchTuple objects (filters) in a RegexMatchSet.
// For each RegexMatchSetUpdate object, you specify the following values:
//
//    * Whether to insert or delete the object from the array. If you want to
//    change a RegexMatchSetUpdate object, you delete the existing object and
//    add a new one.
//
//    * The part of a web request that you want AWS WAF to inspectupdate, such
//    as a query string or the value of the User-Agent header.
//
//    * The identifier of the pattern (a regular expression) that you want AWS
//    WAF to look for. For more information, see RegexPatternSet.
//
//    * Whether to perform any conversions on the request, such as converting
//    it to lowercase, before inspecting it for the specified string.
//
// For example, you can create a RegexPatternSet that matches any requests with
// User-Agent headers that contain the string B[a@]dB[o0]t. You can then configure
// AWS WAF to reject those requests.
//
// To create and configure a RegexMatchSet, perform the following steps:
//
// Create a RegexMatchSet. For more information, see CreateRegexMatchSet.
//
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateRegexMatchSet request.
//
// Submit an UpdateRegexMatchSet request to specify the part of the request
// that you want AWS WAF to inspect (for example, the header or the URI) and
// the identifier of the RegexPatternSet that contain the regular expression
// patters you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/).
//
//    // Example sending a request using UpdateRegexMatchSetRequest.
//    req := client.UpdateRegexMatchSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/waf-regional-2016-11-28/UpdateRegexMatchSet
func (c *Client) UpdateRegexMatchSetRequest(input *UpdateRegexMatchSetInput) UpdateRegexMatchSetRequest {
	op := &aws.Operation{
		Name:       opUpdateRegexMatchSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRegexMatchSetInput{}
	}

	req := c.newRequest(op, input, &UpdateRegexMatchSetOutput{})
	return UpdateRegexMatchSetRequest{Request: req, Input: input, Copy: c.UpdateRegexMatchSetRequest}
}

// UpdateRegexMatchSetRequest is the request type for the
// UpdateRegexMatchSet API operation.
type UpdateRegexMatchSetRequest struct {
	*aws.Request
	Input *UpdateRegexMatchSetInput
	Copy  func(*UpdateRegexMatchSetInput) UpdateRegexMatchSetRequest
}

// Send marshals and sends the UpdateRegexMatchSet API request.
func (r UpdateRegexMatchSetRequest) Send(ctx context.Context) (*UpdateRegexMatchSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateRegexMatchSetResponse{
		UpdateRegexMatchSetOutput: r.Request.Data.(*UpdateRegexMatchSetOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateRegexMatchSetResponse is the response type for the
// UpdateRegexMatchSet API operation.
type UpdateRegexMatchSetResponse struct {
	*UpdateRegexMatchSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateRegexMatchSet request.
func (r *UpdateRegexMatchSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

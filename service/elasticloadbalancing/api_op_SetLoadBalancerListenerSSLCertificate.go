// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for SetLoadBalancerListenerSSLCertificate.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/SetLoadBalancerListenerSSLCertificateInput
type SetLoadBalancerListenerSSLCertificateInput struct {
	_ struct{} `type:"structure"`

	// The name of the load balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `type:"string" required:"true"`

	// The port that uses the specified SSL certificate.
	//
	// LoadBalancerPort is a required field
	LoadBalancerPort *int64 `type:"integer" required:"true"`

	// The Amazon Resource Name (ARN) of the SSL certificate.
	//
	// SSLCertificateId is a required field
	SSLCertificateId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SetLoadBalancerListenerSSLCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetLoadBalancerListenerSSLCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetLoadBalancerListenerSSLCertificateInput"}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if s.LoadBalancerPort == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerPort"))
	}

	if s.SSLCertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SSLCertificateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of SetLoadBalancerListenerSSLCertificate.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/SetLoadBalancerListenerSSLCertificateOutput
type SetLoadBalancerListenerSSLCertificateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetLoadBalancerListenerSSLCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

const opSetLoadBalancerListenerSSLCertificate = "SetLoadBalancerListenerSSLCertificate"

// SetLoadBalancerListenerSSLCertificateRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Sets the certificate that terminates the specified listener's SSL connections.
// The specified certificate replaces any prior certificate that was used on
// the same load balancer and port.
//
// For more information about updating your SSL certificate, see Replace the
// SSL Certificate for Your Load Balancer (http://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-update-ssl-cert.html)
// in the Classic Load Balancers Guide.
//
//    // Example sending a request using SetLoadBalancerListenerSSLCertificateRequest.
//    req := client.SetLoadBalancerListenerSSLCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/SetLoadBalancerListenerSSLCertificate
func (c *Client) SetLoadBalancerListenerSSLCertificateRequest(input *SetLoadBalancerListenerSSLCertificateInput) SetLoadBalancerListenerSSLCertificateRequest {
	op := &aws.Operation{
		Name:       opSetLoadBalancerListenerSSLCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetLoadBalancerListenerSSLCertificateInput{}
	}

	req := c.newRequest(op, input, &SetLoadBalancerListenerSSLCertificateOutput{})
	return SetLoadBalancerListenerSSLCertificateRequest{Request: req, Input: input, Copy: c.SetLoadBalancerListenerSSLCertificateRequest}
}

// SetLoadBalancerListenerSSLCertificateRequest is the request type for the
// SetLoadBalancerListenerSSLCertificate API operation.
type SetLoadBalancerListenerSSLCertificateRequest struct {
	*aws.Request
	Input *SetLoadBalancerListenerSSLCertificateInput
	Copy  func(*SetLoadBalancerListenerSSLCertificateInput) SetLoadBalancerListenerSSLCertificateRequest
}

// Send marshals and sends the SetLoadBalancerListenerSSLCertificate API request.
func (r SetLoadBalancerListenerSSLCertificateRequest) Send(ctx context.Context) (*SetLoadBalancerListenerSSLCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetLoadBalancerListenerSSLCertificateResponse{
		SetLoadBalancerListenerSSLCertificateOutput: r.Request.Data.(*SetLoadBalancerListenerSSLCertificateOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetLoadBalancerListenerSSLCertificateResponse is the response type for the
// SetLoadBalancerListenerSSLCertificate API operation.
type SetLoadBalancerListenerSSLCertificateResponse struct {
	*SetLoadBalancerListenerSSLCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetLoadBalancerListenerSSLCertificate request.
func (r *SetLoadBalancerListenerSSLCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
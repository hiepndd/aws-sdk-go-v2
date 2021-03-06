// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssignIpv6AddressesInput struct {
	_ struct{} `type:"structure"`

	// The number of IPv6 addresses to assign to the network interface. Amazon EC2
	// automatically selects the IPv6 addresses from the subnet range. You can't
	// use this option if specifying specific IPv6 addresses.
	Ipv6AddressCount *int64 `locationName:"ipv6AddressCount" type:"integer"`

	// One or more specific IPv6 addresses to be assigned to the network interface.
	// You can't use this option if you're specifying a number of IPv6 addresses.
	Ipv6Addresses []string `locationName:"ipv6Addresses" locationNameList:"item" type:"list"`

	// The ID of the network interface.
	//
	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string" required:"true"`
}

// String returns the string representation
func (s AssignIpv6AddressesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssignIpv6AddressesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssignIpv6AddressesInput"}

	if s.NetworkInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssignIpv6AddressesOutput struct {
	_ struct{} `type:"structure"`

	// The IPv6 addresses assigned to the network interface.
	AssignedIpv6Addresses []string `locationName:"assignedIpv6Addresses" locationNameList:"item" type:"list"`

	// The ID of the network interface.
	NetworkInterfaceId *string `locationName:"networkInterfaceId" type:"string"`
}

// String returns the string representation
func (s AssignIpv6AddressesOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssignIpv6Addresses = "AssignIpv6Addresses"

// AssignIpv6AddressesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Assigns one or more IPv6 addresses to the specified network interface. You
// can specify one or more specific IPv6 addresses, or you can specify the number
// of IPv6 addresses to be automatically assigned from within the subnet's IPv6
// CIDR block range. You can assign as many IPv6 addresses to a network interface
// as you can assign private IPv4 addresses, and the limit varies per instance
// type. For information, see IP Addresses Per Network Interface Per Instance
// Type (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI)
// in the Amazon Elastic Compute Cloud User Guide.
//
// You must specify either the IPv6 addresses or the IPv6 address count in the
// request.
//
//    // Example sending a request using AssignIpv6AddressesRequest.
//    req := client.AssignIpv6AddressesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssignIpv6Addresses
func (c *Client) AssignIpv6AddressesRequest(input *AssignIpv6AddressesInput) AssignIpv6AddressesRequest {
	op := &aws.Operation{
		Name:       opAssignIpv6Addresses,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssignIpv6AddressesInput{}
	}

	req := c.newRequest(op, input, &AssignIpv6AddressesOutput{})
	return AssignIpv6AddressesRequest{Request: req, Input: input, Copy: c.AssignIpv6AddressesRequest}
}

// AssignIpv6AddressesRequest is the request type for the
// AssignIpv6Addresses API operation.
type AssignIpv6AddressesRequest struct {
	*aws.Request
	Input *AssignIpv6AddressesInput
	Copy  func(*AssignIpv6AddressesInput) AssignIpv6AddressesRequest
}

// Send marshals and sends the AssignIpv6Addresses API request.
func (r AssignIpv6AddressesRequest) Send(ctx context.Context) (*AssignIpv6AddressesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssignIpv6AddressesResponse{
		AssignIpv6AddressesOutput: r.Request.Data.(*AssignIpv6AddressesOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssignIpv6AddressesResponse is the response type for the
// AssignIpv6Addresses API operation.
type AssignIpv6AddressesResponse struct {
	*AssignIpv6AddressesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssignIpv6Addresses request.
func (r *AssignIpv6AddressesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

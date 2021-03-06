// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateRoomMembershipInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The Amazon Chime member ID (user ID or bot ID).
	//
	// MemberId is a required field
	MemberId *string `type:"string" required:"true"`

	// The role of the member.
	Role RoomMembershipRole `type:"string" enum:"true"`

	// The room ID.
	//
	// RoomId is a required field
	RoomId *string `location:"uri" locationName:"roomId" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateRoomMembershipInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRoomMembershipInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRoomMembershipInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.MemberId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MemberId"))
	}

	if s.RoomId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoomId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRoomMembershipInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MemberId != nil {
		v := *s.MemberId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MemberId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Role) > 0 {
		v := s.Role

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Role", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RoomId != nil {
		v := *s.RoomId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "roomId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateRoomMembershipOutput struct {
	_ struct{} `type:"structure"`

	// The room membership details.
	RoomMembership *RoomMembership `type:"structure"`
}

// String returns the string representation
func (s CreateRoomMembershipOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateRoomMembershipOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.RoomMembership != nil {
		v := s.RoomMembership

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "RoomMembership", v, metadata)
	}
	return nil
}

const opCreateRoomMembership = "CreateRoomMembership"

// CreateRoomMembershipRequest returns a request value for making API operation for
// Amazon Chime.
//
// Adds a member to a chat room in an Amazon Chime Enterprise account. A member
// can be either a user or a bot. The member role designates whether the member
// is a chat room administrator or a general chat room member.
//
//    // Example sending a request using CreateRoomMembershipRequest.
//    req := client.CreateRoomMembershipRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/CreateRoomMembership
func (c *Client) CreateRoomMembershipRequest(input *CreateRoomMembershipInput) CreateRoomMembershipRequest {
	op := &aws.Operation{
		Name:       opCreateRoomMembership,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/rooms/{roomId}/memberships",
	}

	if input == nil {
		input = &CreateRoomMembershipInput{}
	}

	req := c.newRequest(op, input, &CreateRoomMembershipOutput{})
	return CreateRoomMembershipRequest{Request: req, Input: input, Copy: c.CreateRoomMembershipRequest}
}

// CreateRoomMembershipRequest is the request type for the
// CreateRoomMembership API operation.
type CreateRoomMembershipRequest struct {
	*aws.Request
	Input *CreateRoomMembershipInput
	Copy  func(*CreateRoomMembershipInput) CreateRoomMembershipRequest
}

// Send marshals and sends the CreateRoomMembership API request.
func (r CreateRoomMembershipRequest) Send(ctx context.Context) (*CreateRoomMembershipResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateRoomMembershipResponse{
		CreateRoomMembershipOutput: r.Request.Data.(*CreateRoomMembershipOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateRoomMembershipResponse is the response type for the
// CreateRoomMembership API operation.
type CreateRoomMembershipResponse struct {
	*CreateRoomMembershipOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateRoomMembership request.
func (r *CreateRoomMembershipResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

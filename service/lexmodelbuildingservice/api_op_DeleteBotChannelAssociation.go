// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

type DeleteBotChannelAssociationInput struct {
	_ struct{} `type:"structure"`

	// An alias that points to the specific version of the Amazon Lex bot to which
	// this association is being made.
	//
	// BotAlias is a required field
	BotAlias *string `location:"uri" locationName:"aliasName" min:"1" type:"string" required:"true"`

	// The name of the Amazon Lex bot.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" min:"2" type:"string" required:"true"`

	// The name of the association. The name is case sensitive.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBotChannelAssociationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBotChannelAssociationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteBotChannelAssociationInput"}

	if s.BotAlias == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotAlias"))
	}
	if s.BotAlias != nil && len(*s.BotAlias) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BotAlias", 1))
	}

	if s.BotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotName"))
	}
	if s.BotName != nil && len(*s.BotName) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("BotName", 2))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBotChannelAssociationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BotAlias != nil {
		v := *s.BotAlias

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "aliasName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BotName != nil {
		v := *s.BotName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteBotChannelAssociationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteBotChannelAssociationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteBotChannelAssociationOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteBotChannelAssociation = "DeleteBotChannelAssociation"

// DeleteBotChannelAssociationRequest returns a request value for making API operation for
// Amazon Lex Model Building Service.
//
// Deletes the association between an Amazon Lex bot and a messaging platform.
//
// This operation requires permission for the lex:DeleteBotChannelAssociation
// action.
//
//    // Example sending a request using DeleteBotChannelAssociationRequest.
//    req := client.DeleteBotChannelAssociationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lex-models-2017-04-19/DeleteBotChannelAssociation
func (c *Client) DeleteBotChannelAssociationRequest(input *DeleteBotChannelAssociationInput) DeleteBotChannelAssociationRequest {
	op := &aws.Operation{
		Name:       opDeleteBotChannelAssociation,
		HTTPMethod: "DELETE",
		HTTPPath:   "/bots/{botName}/aliases/{aliasName}/channels/{name}",
	}

	if input == nil {
		input = &DeleteBotChannelAssociationInput{}
	}

	req := c.newRequest(op, input, &DeleteBotChannelAssociationOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteBotChannelAssociationRequest{Request: req, Input: input, Copy: c.DeleteBotChannelAssociationRequest}
}

// DeleteBotChannelAssociationRequest is the request type for the
// DeleteBotChannelAssociation API operation.
type DeleteBotChannelAssociationRequest struct {
	*aws.Request
	Input *DeleteBotChannelAssociationInput
	Copy  func(*DeleteBotChannelAssociationInput) DeleteBotChannelAssociationRequest
}

// Send marshals and sends the DeleteBotChannelAssociation API request.
func (r DeleteBotChannelAssociationRequest) Send(ctx context.Context) (*DeleteBotChannelAssociationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteBotChannelAssociationResponse{
		DeleteBotChannelAssociationOutput: r.Request.Data.(*DeleteBotChannelAssociationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteBotChannelAssociationResponse is the response type for the
// DeleteBotChannelAssociation API operation.
type DeleteBotChannelAssociationResponse struct {
	*DeleteBotChannelAssociationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteBotChannelAssociation request.
func (r *DeleteBotChannelAssociationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

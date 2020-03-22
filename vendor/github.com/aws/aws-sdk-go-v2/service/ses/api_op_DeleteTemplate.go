// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to delete an email template. For more information, see
// the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html).
type DeleteTemplateInput struct {
	_ struct{} `type:"structure"`

	// The name of the template to be deleted.
	//
	// TemplateName is a required field
	TemplateName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTemplateInput"}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteTemplateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTemplate = "DeleteTemplate"

// DeleteTemplateRequest returns a request value for making API operation for
// Amazon Simple Email Service.
//
// Deletes an email template.
//
// You can execute this operation no more than once per second.
//
//    // Example sending a request using DeleteTemplateRequest.
//    req := client.DeleteTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/email-2010-12-01/DeleteTemplate
func (c *Client) DeleteTemplateRequest(input *DeleteTemplateInput) DeleteTemplateRequest {
	op := &aws.Operation{
		Name:       opDeleteTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTemplateInput{}
	}

	req := c.newRequest(op, input, &DeleteTemplateOutput{})
	return DeleteTemplateRequest{Request: req, Input: input, Copy: c.DeleteTemplateRequest}
}

// DeleteTemplateRequest is the request type for the
// DeleteTemplate API operation.
type DeleteTemplateRequest struct {
	*aws.Request
	Input *DeleteTemplateInput
	Copy  func(*DeleteTemplateInput) DeleteTemplateRequest
}

// Send marshals and sends the DeleteTemplate API request.
func (r DeleteTemplateRequest) Send(ctx context.Context) (*DeleteTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTemplateResponse{
		DeleteTemplateOutput: r.Request.Data.(*DeleteTemplateOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTemplateResponse is the response type for the
// DeleteTemplate API operation.
type DeleteTemplateResponse struct {
	*DeleteTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTemplate request.
func (r *DeleteTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
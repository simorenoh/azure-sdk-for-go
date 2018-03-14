package face

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"io"
	"net/http"
)

// Client is the an API for face detection, verification, and identification.
type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient(azureRegion AzureRegions) Client {
	return Client{New(azureRegion)}
}

// DetectWithStream detect human faces in an image and returns face locations, and optionally with faceIds, landmarks,
// and attributes.
//
// imageParameter is an image stream. imageParameter will be closed upon successful return. Callers should ensure
// closure when receiving an error.returnFaceID is a value indicating whether the operation should return faceIds
// of detected faces. returnFaceLandmarks is a value indicating whether the operation should return landmarks of
// the detected faces. returnFaceAttributes is analyze and return the one or more specified face attributes in the
// comma-separated string like "returnFaceAttributes=age,gender". Supported face attributes include age, gender,
// headPose, smile, facialHair, glasses and emotion. Note that each face attribute analysis has additional
// computational and time cost.
func (client Client) DetectWithStream(ctx context.Context, imageParameter io.ReadCloser, returnFaceID *bool, returnFaceLandmarks *bool, returnFaceAttributes []AttributeType) (result ListDetectedFace, err error) {
	req, err := client.DetectWithStreamPreparer(ctx, imageParameter, returnFaceID, returnFaceLandmarks, returnFaceAttributes)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.Client", "DetectWithStream", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetectWithStreamSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.Client", "DetectWithStream", resp, "Failure sending request")
		return
	}

	result, err = client.DetectWithStreamResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.Client", "DetectWithStream", resp, "Failure responding to request")
	}

	return
}

// DetectWithStreamPreparer prepares the DetectWithStream request.
func (client Client) DetectWithStreamPreparer(ctx context.Context, imageParameter io.ReadCloser, returnFaceID *bool, returnFaceLandmarks *bool, returnFaceAttributes []AttributeType) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	queryParameters := map[string]interface{}{}
	if returnFaceID != nil {
		queryParameters["returnFaceId"] = autorest.Encode("query", *returnFaceID)
	} else {
		queryParameters["returnFaceId"] = autorest.Encode("query", true)
	}
	if returnFaceLandmarks != nil {
		queryParameters["returnFaceLandmarks"] = autorest.Encode("query", *returnFaceLandmarks)
	} else {
		queryParameters["returnFaceLandmarks"] = autorest.Encode("query", false)
	}
	if returnFaceAttributes != nil && len(returnFaceAttributes) > 0 {
		queryParameters["returnFaceAttributes"] = autorest.Encode("query", returnFaceAttributes, ",")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/octet-stream"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/detect"),
		autorest.WithFile(imageParameter),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetectWithStreamSender sends the DetectWithStream request. The method will close the
// http.Response Body if it receives an error.
func (client Client) DetectWithStreamSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetectWithStreamResponder handles the response to the DetectWithStream request. The method always
// closes the http.Response Body.
func (client Client) DetectWithStreamResponder(resp *http.Response) (result ListDetectedFace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DetectWithURL detect human faces in an image and returns face locations, and optionally with faceIds, landmarks, and
// attributes.
//
// imageURL is a JSON document with a URL pointing to the image that is to be analyzed. returnFaceID is a value
// indicating whether the operation should return faceIds of detected faces. returnFaceLandmarks is a value
// indicating whether the operation should return landmarks of the detected faces. returnFaceAttributes is analyze
// and return the one or more specified face attributes in the comma-separated string like
// "returnFaceAttributes=age,gender". Supported face attributes include age, gender, headPose, smile, facialHair,
// glasses and emotion. Note that each face attribute analysis has additional computational and time cost.
func (client Client) DetectWithURL(ctx context.Context, imageURL ImageURL, returnFaceID *bool, returnFaceLandmarks *bool, returnFaceAttributes []AttributeType) (result ListDetectedFace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: imageURL,
			Constraints: []validation.Constraint{{Target: "imageURL.URL", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.Client", "DetectWithURL", err.Error())
	}

	req, err := client.DetectWithURLPreparer(ctx, imageURL, returnFaceID, returnFaceLandmarks, returnFaceAttributes)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.Client", "DetectWithURL", nil, "Failure preparing request")
		return
	}

	resp, err := client.DetectWithURLSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.Client", "DetectWithURL", resp, "Failure sending request")
		return
	}

	result, err = client.DetectWithURLResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.Client", "DetectWithURL", resp, "Failure responding to request")
	}

	return
}

// DetectWithURLPreparer prepares the DetectWithURL request.
func (client Client) DetectWithURLPreparer(ctx context.Context, imageURL ImageURL, returnFaceID *bool, returnFaceLandmarks *bool, returnFaceAttributes []AttributeType) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	queryParameters := map[string]interface{}{}
	if returnFaceID != nil {
		queryParameters["returnFaceId"] = autorest.Encode("query", *returnFaceID)
	} else {
		queryParameters["returnFaceId"] = autorest.Encode("query", true)
	}
	if returnFaceLandmarks != nil {
		queryParameters["returnFaceLandmarks"] = autorest.Encode("query", *returnFaceLandmarks)
	} else {
		queryParameters["returnFaceLandmarks"] = autorest.Encode("query", false)
	}
	if returnFaceAttributes != nil && len(returnFaceAttributes) > 0 {
		queryParameters["returnFaceAttributes"] = autorest.Encode("query", returnFaceAttributes, ",")
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/detect"),
		autorest.WithJSON(imageURL),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DetectWithURLSender sends the DetectWithURL request. The method will close the
// http.Response Body if it receives an error.
func (client Client) DetectWithURLSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DetectWithURLResponder handles the response to the DetectWithURL request. The method always
// closes the http.Response Body.
func (client Client) DetectWithURLResponder(resp *http.Response) (result ListDetectedFace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// FindSimilar given query face's faceId, find the similar-looking faces from a faceId array or a faceListId.
//
// body is request body for Find Similar.
func (client Client) FindSimilar(ctx context.Context, body FindSimilarRequest) (result ListSimilarFace, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.FaceID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.FaceListID", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.FaceListID", Name: validation.MaxLength, Rule: 64, Chain: nil},
						{Target: "body.FaceListID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil},
					}},
				{Target: "body.FaceIds", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.FaceIds", Name: validation.MaxItems, Rule: 1000, Chain: nil}}},
				{Target: "body.MaxNumOfCandidatesReturned", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.MaxNumOfCandidatesReturned", Name: validation.InclusiveMaximum, Rule: 1000, Chain: nil},
						{Target: "body.MaxNumOfCandidatesReturned", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("face.Client", "FindSimilar", err.Error())
	}

	req, err := client.FindSimilarPreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.Client", "FindSimilar", nil, "Failure preparing request")
		return
	}

	resp, err := client.FindSimilarSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.Client", "FindSimilar", resp, "Failure sending request")
		return
	}

	result, err = client.FindSimilarResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.Client", "FindSimilar", resp, "Failure responding to request")
	}

	return
}

// FindSimilarPreparer prepares the FindSimilar request.
func (client Client) FindSimilarPreparer(ctx context.Context, body FindSimilarRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/findsimilars"),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// FindSimilarSender sends the FindSimilar request. The method will close the
// http.Response Body if it receives an error.
func (client Client) FindSimilarSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// FindSimilarResponder handles the response to the FindSimilar request. The method always
// closes the http.Response Body.
func (client Client) FindSimilarResponder(resp *http.Response) (result ListSimilarFace, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Group divide candidate faces into groups based on face similarity.
//
// body is request body for grouping.
func (client Client) Group(ctx context.Context, body GroupRequest) (result GroupResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.FaceIds", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "body.FaceIds", Name: validation.MaxItems, Rule: 1000, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("face.Client", "Group", err.Error())
	}

	req, err := client.GroupPreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.Client", "Group", nil, "Failure preparing request")
		return
	}

	resp, err := client.GroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.Client", "Group", resp, "Failure sending request")
		return
	}

	result, err = client.GroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.Client", "Group", resp, "Failure responding to request")
	}

	return
}

// GroupPreparer prepares the Group request.
func (client Client) GroupPreparer(ctx context.Context, body GroupRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/group"),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GroupSender sends the Group request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GroupResponder handles the response to the Group request. The method always
// closes the http.Response Body.
func (client Client) GroupResponder(resp *http.Response) (result GroupResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Identify identify unknown faces from a person group.
//
// body is request body for identify operation.
func (client Client) Identify(ctx context.Context, body IdentifyRequest) (result ListIdentifyResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.PersonGroupID", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "body.PersonGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
					{Target: "body.PersonGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil},
				}},
				{Target: "body.FaceIds", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "body.FaceIds", Name: validation.MaxItems, Rule: 10, Chain: nil}}},
				{Target: "body.MaxNumOfCandidatesReturned", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.MaxNumOfCandidatesReturned", Name: validation.InclusiveMaximum, Rule: 5, Chain: nil},
						{Target: "body.MaxNumOfCandidatesReturned", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil},
					}}}}}); err != nil {
		return result, validation.NewError("face.Client", "Identify", err.Error())
	}

	req, err := client.IdentifyPreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.Client", "Identify", nil, "Failure preparing request")
		return
	}

	resp, err := client.IdentifySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.Client", "Identify", resp, "Failure sending request")
		return
	}

	result, err = client.IdentifyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.Client", "Identify", resp, "Failure responding to request")
	}

	return
}

// IdentifyPreparer prepares the Identify request.
func (client Client) IdentifyPreparer(ctx context.Context, body IdentifyRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/identify"),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// IdentifySender sends the Identify request. The method will close the
// http.Response Body if it receives an error.
func (client Client) IdentifySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// IdentifyResponder handles the response to the Identify request. The method always
// closes the http.Response Body.
func (client Client) IdentifyResponder(resp *http.Response) (result ListIdentifyResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// VerifyFaceToFace verify whether two faces belong to a same person or whether one face belongs to a person.
//
// body is request body for verify operation.
func (client Client) VerifyFaceToFace(ctx context.Context, body VerifyFaceToFaceRequest) (result VerifyResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.FaceID1", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.FaceID2", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.Client", "VerifyFaceToFace", err.Error())
	}

	req, err := client.VerifyFaceToFacePreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.Client", "VerifyFaceToFace", nil, "Failure preparing request")
		return
	}

	resp, err := client.VerifyFaceToFaceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.Client", "VerifyFaceToFace", resp, "Failure sending request")
		return
	}

	result, err = client.VerifyFaceToFaceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.Client", "VerifyFaceToFace", resp, "Failure responding to request")
	}

	return
}

// VerifyFaceToFacePreparer prepares the VerifyFaceToFace request.
func (client Client) VerifyFaceToFacePreparer(ctx context.Context, body VerifyFaceToFaceRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/verify"),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// VerifyFaceToFaceSender sends the VerifyFaceToFace request. The method will close the
// http.Response Body if it receives an error.
func (client Client) VerifyFaceToFaceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// VerifyFaceToFaceResponder handles the response to the VerifyFaceToFace request. The method always
// closes the http.Response Body.
func (client Client) VerifyFaceToFaceResponder(resp *http.Response) (result VerifyResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// VerifyFaceToPerson verify whether two faces belong to a same person. Compares a face Id with a Person Id
//
// body is request body for verifying two faces in a person group
func (client Client) VerifyFaceToPerson(ctx context.Context, body VerifyFaceToPersonRequest) (result VerifyResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body.FaceID", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "body.PersonGroupID", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "body.PersonGroupID", Name: validation.MaxLength, Rule: 64, Chain: nil},
						{Target: "body.PersonGroupID", Name: validation.Pattern, Rule: `^[a-z0-9-_]+$`, Chain: nil},
					}},
				{Target: "body.PersonID", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("face.Client", "VerifyFaceToPerson", err.Error())
	}

	req, err := client.VerifyFaceToPersonPreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.Client", "VerifyFaceToPerson", nil, "Failure preparing request")
		return
	}

	resp, err := client.VerifyFaceToPersonSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "face.Client", "VerifyFaceToPerson", resp, "Failure sending request")
		return
	}

	result, err = client.VerifyFaceToPersonResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "face.Client", "VerifyFaceToPerson", resp, "Failure responding to request")
	}

	return
}

// VerifyFaceToPersonPreparer prepares the VerifyFaceToPerson request.
func (client Client) VerifyFaceToPersonPreparer(ctx context.Context, body VerifyFaceToPersonRequest) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"AzureRegion": client.AzureRegion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("https://{AzureRegion}.api.cognitive.microsoft.com/face/v1.0", urlParameters),
		autorest.WithPath("/verify"),
		autorest.WithJSON(body))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// VerifyFaceToPersonSender sends the VerifyFaceToPerson request. The method will close the
// http.Response Body if it receives an error.
func (client Client) VerifyFaceToPersonSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// VerifyFaceToPersonResponder handles the response to the VerifyFaceToPerson request. The method always
// closes the http.Response Body.
func (client Client) VerifyFaceToPersonResponder(resp *http.Response) (result VerifyResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

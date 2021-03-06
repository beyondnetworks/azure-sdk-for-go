package migrate

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
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AssessmentOptionsClient is the move your workloads to Azure.
type AssessmentOptionsClient struct {
	BaseClient
}

// NewAssessmentOptionsClient creates an instance of the AssessmentOptionsClient client.
func NewAssessmentOptionsClient(subscriptionID string, acceptLanguage string) AssessmentOptionsClient {
	return NewAssessmentOptionsClientWithBaseURI(DefaultBaseURI, subscriptionID, acceptLanguage)
}

// NewAssessmentOptionsClientWithBaseURI creates an instance of the AssessmentOptionsClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewAssessmentOptionsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) AssessmentOptionsClient {
	return AssessmentOptionsClient{NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)}
}

// Get get the available options for the properties of an assessment.
// Parameters:
// locationName - azure region in which the project is created.
func (client AssessmentOptionsClient) Get(ctx context.Context, locationName string) (result AssessmentOptionsResultList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AssessmentOptionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, locationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.AssessmentOptionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.AssessmentOptionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.AssessmentOptionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AssessmentOptionsClient) GetPreparer(ctx context.Context, locationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Migrate/locations/{locationName}/assessmentOptions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if len(client.AcceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(client.AcceptLanguage)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AssessmentOptionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AssessmentOptionsClient) GetResponder(resp *http.Response) (result AssessmentOptionsResultList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

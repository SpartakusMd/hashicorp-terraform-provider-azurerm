package appplatform

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringSettingsUpdatePutOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// MonitoringSettingsUpdatePut ...
func (c AppPlatformClient) MonitoringSettingsUpdatePut(ctx context.Context, id SpringId, input MonitoringSettingResource) (result MonitoringSettingsUpdatePutOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       fmt.Sprintf("%s/monitoringSettings/default", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// MonitoringSettingsUpdatePutThenPoll performs MonitoringSettingsUpdatePut then polls until it's completed
func (c AppPlatformClient) MonitoringSettingsUpdatePutThenPoll(ctx context.Context, id SpringId, input MonitoringSettingResource) error {
	result, err := c.MonitoringSettingsUpdatePut(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing MonitoringSettingsUpdatePut: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after MonitoringSettingsUpdatePut: %+v", err)
	}

	return nil
}
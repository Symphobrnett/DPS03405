//
// Copyright (C) 2022 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"encoding/json"
	"fmt"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/edgexfoundry/edgex-go/internal/core/metadata/container"
	dbMock "github.com/edgexfoundry/edgex-go/internal/core/metadata/infrastructure/interfaces/mocks"
	"github.com/edgexfoundry/go-mod-bootstrap/v2/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	commonDTO "github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos/requests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var testDeviceCommandRequest = requests.AddDeviceCommandRequest{
	BaseRequest: commonDTO.BaseRequest{
		RequestId:   ExampleUUID,
		Versionable: commonDTO.NewVersionable(),
	},
	ProfileName: TestDeviceProfileName,
	DeviceCommand: dtos.DeviceCommand{
		Name:      "TestDeviceCommandNewName",
		IsHidden:  false,
		ReadWrite: common.ReadWrite_RW,
		ResourceOperations: []dtos.ResourceOperation{{
			DeviceResource: TestDeviceResourceName,
			Mappings:       testMappings,
		}},
	},
}

func buildTestUpdateDeviceCommandRequest() requests.UpdateDeviceCommandRequest {
	testDeviceCommandName := TestDeviceCommandName
	testIsHidden := false
	var testUpdateDeviceCommandReq = requests.UpdateDeviceCommandRequest{
		BaseRequest: commonDTO.BaseRequest{
			RequestId:   ExampleUUID,
			Versionable: commonDTO.NewVersionable(),
		},
		ProfileName: TestDeviceProfileName,
		DeviceCommand: dtos.UpdateDeviceCommand{
			Name:     &testDeviceCommandName,
			IsHidden: &testIsHidden,
		},
	}
	return testUpdateDeviceCommandReq
}

func TestAddDeviceProfileDeviceCommand(t *testing.T) {
	deviceProfile := dtos.ToDeviceProfileModel(buildTestDeviceProfileRequest().Profile)
	expectedRequestId := ExampleUUID

	valid := testDeviceCommandRequest
	noRequestId := testDeviceCommandRequest
	noRequestId.RequestId = ""
	duplicateName := testDeviceCommandRequest
	duplicateName.DeviceCommand.Name = TestDeviceCommandName
	notFoundDeviceResource := testDeviceCommandRequest
	notFoundDeviceResource.DeviceCommand.ResourceOperations = []dtos.ResourceOperation{{
		DeviceResource: "notFoundName",
		Mappings:       testMappings,
	}}

	dic := mockDic()
	dbClientMock := &dbMock.DBClient{}
	dbClientMock.On("DeviceProfileByName", valid.ProfileName).Return(deviceProfile, nil)
	dbClientMock.On("UpdateDeviceProfile", mock.Anything).Return(nil)
	dic.Update(di.ServiceConstructorMap{
		container.DBClientInterfaceName: func(get di.Get) interface{} {
			return dbClientMock
		},
	})
	controller := NewDeviceCommandController(dic)
	assert.NotNil(t, controller)

	tests := []struct {
		name               string
		Request            []requests.AddDeviceCommandRequest
		expectedStatusCode int
	}{
		{"Valid - AddDeviceCommandRequest", []requests.AddDeviceCommandRequest{valid}, http.StatusCreated},
		{"Valid - No requestId", []requests.AddDeviceCommandRequest{noRequestId}, http.StatusCreated},
		{"invalid - Duplicate deviceCommand name", []requests.AddDeviceCommandRequest{duplicateName}, http.StatusBadRequest},
		{"invalid - Not Exist deviceResource", []requests.AddDeviceCommandRequest{notFoundDeviceResource}, http.StatusBadRequest},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			jsonData, err := json.Marshal(testCase.Request)
			require.NoError(t, err)

			reader := strings.NewReader(string(jsonData))
			req, err := http.NewRequest(http.MethodPost, common.ApiDeviceProfileDeviceCommandRoute, reader)
			require.NoError(t, err)

			// Act
			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(controller.AddDeviceProfileDeviceCommand)
			handler.ServeHTTP(recorder, req)

			var res []commonDTO.BaseResponse
			err = json.Unmarshal(recorder.Body.Bytes(), &res)
			require.NoError(t, err)

			assert.Equal(t, http.StatusMultiStatus, recorder.Result().StatusCode, "HTTP status code not as expected")
			assert.Equal(t, common.ApiVersion, res[0].ApiVersion, "API Version not as expected")
			if res[0].RequestId != "" {
				assert.Equal(t, expectedRequestId, res[0].RequestId, "RequestID not as expected")
			}
			assert.Equal(t, testCase.expectedStatusCode, res[0].StatusCode, "BaseResponse status code not as expected")
		})
	}
}

func TestAddDeviceProfileDeviceCommand_BadRequest(t *testing.T) {
	noDeviceCommandName := testDeviceCommandRequest
	noDeviceCommandName.DeviceCommand.Name = ""
	invalidReadWrite := testDeviceCommandRequest
	invalidReadWrite.DeviceCommand.ReadWrite = "invalid"
	noResourceOperations := testDeviceCommandRequest
	noResourceOperations.DeviceCommand.ResourceOperations = []dtos.ResourceOperation{{}}

	dic := mockDic()
	controller := NewDeviceCommandController(dic)
	assert.NotNil(t, controller)

	tests := []struct {
		name    string
		Request []requests.AddDeviceCommandRequest
	}{
		{"invalid - No deviceCommand name", []requests.AddDeviceCommandRequest{noDeviceCommandName}},
		{"invalid - Invalid ReadWrite", []requests.AddDeviceCommandRequest{invalidReadWrite}},
		{"invalid - No ResourceOperations", []requests.AddDeviceCommandRequest{noResourceOperations}},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			jsonData, err := json.Marshal(testCase.Request)
			require.NoError(t, err)

			reader := strings.NewReader(string(jsonData))
			req, err := http.NewRequest(http.MethodPost, common.ApiDeviceProfileDeviceCommandRoute, reader)
			require.NoError(t, err)

			// Act
			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(controller.AddDeviceProfileDeviceCommand)
			handler.ServeHTTP(recorder, req)

			var res commonDTO.BaseResponse
			err = json.Unmarshal(recorder.Body.Bytes(), &res)
			require.NoError(t, err)

			assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode, "HTTP status code not as expected")
			assert.Equal(t, common.ApiVersion, res.ApiVersion, "API Version not as expected")
			assert.Equal(t, http.StatusBadRequest, res.StatusCode, "BaseResponse status code not as expected")
			assert.NotEmpty(t, recorder.Body.String(), "Message is empty")
		})
	}
}

func TestPatchDeviceProfileDeviceCommand(t *testing.T) {
	deviceProfile := dtos.ToDeviceProfileModel(buildTestDeviceProfileRequest().Profile)
	testReq := buildTestUpdateDeviceCommandRequest()
	expectedRequestId := ExampleUUID
	emptyString := ""
	notFound := "notFoundName"

	valid := testReq
	validWithNoReqID := testReq
	validWithNoReqID.RequestId = emptyString
	notFoundDeviceCommand := testReq
	notFoundDeviceCommand.DeviceCommand.Name = &notFound
	notFoundProfile := testReq
	notFoundProfile.ProfileName = notFound
	notFoundDBError := errors.NewCommonEdgeX(errors.KindEntityDoesNotExist, fmt.Sprintf("device profile %s does not exists", notFoundProfile.ProfileName), nil)

	noProfileName := testReq
	noProfileName.ProfileName = emptyString

	dic := mockDic()
	dbClientMock := &dbMock.DBClient{}
	dbClientMock.On("DeviceProfileByName", valid.ProfileName).Return(deviceProfile, nil)
	dbClientMock.On("UpdateDeviceProfile", mock.Anything).Return(nil)

	dbClientMock.On("DeviceProfileByName", notFound).Return(deviceProfile, notFoundDBError)
	dic.Update(di.ServiceConstructorMap{
		container.DBClientInterfaceName: func(get di.Get) interface{} {
			return dbClientMock
		},
	})
	controller := NewDeviceCommandController(dic)
	assert.NotNil(t, controller)

	tests := []struct {
		name                 string
		Request              []requests.UpdateDeviceCommandRequest
		expectedStatusCode   int
		expectedResponseCode int
	}{
		{"Valid - PatchDeviceProfileDeviceCommand", []requests.UpdateDeviceCommandRequest{valid}, http.StatusMultiStatus, http.StatusOK},
		{"Valid - PatchDeviceProfileDeviceCommand no requestId", []requests.UpdateDeviceCommandRequest{validWithNoReqID}, http.StatusMultiStatus, http.StatusOK},
		{"Invalid - Not found device command", []requests.UpdateDeviceCommandRequest{notFoundDeviceCommand}, http.StatusMultiStatus, http.StatusNotFound},
		{"Invalid - Not found device profile", []requests.UpdateDeviceCommandRequest{notFoundProfile}, http.StatusMultiStatus, http.StatusNotFound},
		{"Invalid - No profile name", []requests.UpdateDeviceCommandRequest{noProfileName}, http.StatusBadRequest, http.StatusBadRequest},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			jsonData, err := json.Marshal(testCase.Request)
			require.NoError(t, err)

			reader := strings.NewReader(string(jsonData))

			req, err := http.NewRequest(http.MethodPatch, common.ApiDeviceProfileDeviceCommandRoute, reader)
			require.NoError(t, err)

			// Act
			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(controller.PatchDeviceProfileDeviceCommand)
			handler.ServeHTTP(recorder, req)

			if testCase.expectedStatusCode == http.StatusMultiStatus {
				var res []commonDTO.BaseResponse
				err = json.Unmarshal(recorder.Body.Bytes(), &res)
				require.NoError(t, err)

				// Assert
				assert.Equal(t, http.StatusMultiStatus, recorder.Result().StatusCode, "HTTP status code not as expected")
				assert.Equal(t, common.ApiVersion, res[0].ApiVersion, "API Version not as expected")
				if res[0].RequestId != "" {
					assert.Equal(t, expectedRequestId, res[0].RequestId, "RequestID not as expected")
				}
				assert.Equal(t, testCase.expectedResponseCode, res[0].StatusCode, "BaseResponse status code not as expected")
				if testCase.expectedResponseCode == http.StatusOK {
					assert.Empty(t, res[0].Message, "Message should be empty when it is successful")
				} else {
					assert.NotEmpty(t, res[0].Message, "Response message doesn't contain the error message")
				}
			} else {
				var res commonDTO.BaseResponse
				err = json.Unmarshal(recorder.Body.Bytes(), &res)
				require.NoError(t, err)

				// Assert
				assert.Equal(t, testCase.expectedStatusCode, recorder.Result().StatusCode, "HTTP status code not as expected")
				assert.Equal(t, common.ApiVersion, res.ApiVersion, "API Version not as expected")
				assert.Equal(t, testCase.expectedResponseCode, res.StatusCode, "BaseResponse status code not as expected")
				assert.NotEmpty(t, res.Message, "Response message doesn't contain the error message")
			}
		})
	}
}

//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"math"
	"net/http"
	"strconv"

	"github.com/edgexfoundry/edgex-go/internal/pkg"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/utils"
	notificationContainer "github.com/edgexfoundry/edgex-go/internal/support/notifications/container"
	"github.com/edgexfoundry/edgex-go/internal/support/notifications/v2/application"

	"github.com/edgexfoundry/go-mod-bootstrap/v2/bootstrap/container"
	"github.com/edgexfoundry/go-mod-bootstrap/v2/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	commonDTO "github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/common"
	responseDTO "github.com/edgexfoundry/go-mod-core-contracts/v2/v2/dtos/responses"

	"github.com/gorilla/mux"
)

type TransmissionController struct {
	dic *di.Container
}

// NewTransmissionController creates and initializes an TransmissionController
func NewTransmissionController(dic *di.Container) *TransmissionController {
	return &TransmissionController{
		dic: dic,
	}
}

// TransmissionById queries transmission by ID
func (tc *TransmissionController) TransmissionById(w http.ResponseWriter, r *http.Request) {
	lc := container.LoggingClientFrom(tc.dic.Get)
	ctx := r.Context()

	// URL parameters
	vars := mux.Vars(r)
	id := vars[v2.Id]

	trans, err := application.TransmissionById(id, tc.dic)
	if err != nil {
		utils.WriteErrorResponse(w, ctx, lc, err, "")
		return
	}

	response := responseDTO.NewTransmissionResponse("", "", http.StatusOK, trans)
	utils.WriteHttpHeader(w, ctx, http.StatusOK)
	pkg.Encode(response, w, lc)
}

// TransmissionsByTimeRange allows querying of transmissions by their creation timestamp within a given time range, sorted in descending order. Results are paginated.
func (tc *TransmissionController) TransmissionsByTimeRange(w http.ResponseWriter, r *http.Request) {
	lc := container.LoggingClientFrom(tc.dic.Get)
	ctx := r.Context()
	config := notificationContainer.ConfigurationFrom(tc.dic.Get)

	// parse time range (start, end), offset, and limit from incoming request
	start, end, offset, limit, err := utils.ParseTimeRangeOffsetLimit(r, 0, math.MaxInt32, -1, config.Service.MaxResultCount)
	if err != nil {
		utils.WriteErrorResponse(w, ctx, lc, err, "")
		return
	}
	transmissions, err := application.TransmissionsByTimeRange(start, end, offset, limit, tc.dic)
	if err != nil {
		utils.WriteErrorResponse(w, ctx, lc, err, "")
		return
	}

	response := responseDTO.NewMultiTransmissionsResponse("", "", http.StatusOK, transmissions)
	utils.WriteHttpHeader(w, ctx, http.StatusOK)
	pkg.Encode(response, w, lc)
}

func (tc *TransmissionController) AllTransmissions(w http.ResponseWriter, r *http.Request) {
	lc := container.LoggingClientFrom(tc.dic.Get)
	ctx := r.Context()
	config := notificationContainer.ConfigurationFrom(tc.dic.Get)

	// parse URL query string for offset and limit
	offset, limit, _, err := utils.ParseGetAllObjectsRequestQueryString(r, 0, math.MaxUint32, -1, config.Service.MaxResultCount)
	if err != nil {
		utils.WriteErrorResponse(w, ctx, lc, err, "")
		return
	}
	transmissions, err := application.AllTransmissions(offset, limit, tc.dic)
	if err != nil {
		utils.WriteErrorResponse(w, ctx, lc, err, "")
		return
	}

	response := responseDTO.NewMultiTransmissionsResponse("", "", http.StatusOK, transmissions)
	utils.WriteHttpHeader(w, ctx, http.StatusOK)
	pkg.Encode(response, w, lc)
}

// TransmissionsByStatus allows retrieval of the transmissions associated with the specified status. Ordered by create timestamp descending.
func (tc *TransmissionController) TransmissionsByStatus(w http.ResponseWriter, r *http.Request) {
	lc := container.LoggingClientFrom(tc.dic.Get)
	ctx := r.Context()
	config := notificationContainer.ConfigurationFrom(tc.dic.Get)

	vars := mux.Vars(r)
	status := vars[v2.Status]

	// parse URL query string for offset, limit
	offset, limit, _, err := utils.ParseGetAllObjectsRequestQueryString(r, 0, math.MaxInt32, -1, config.Service.MaxResultCount)
	if err != nil {
		utils.WriteErrorResponse(w, ctx, lc, err, "")
		return
	}
	transmissions, err := application.TransmissionsByStatus(offset, limit, status, tc.dic)
	if err != nil {
		utils.WriteErrorResponse(w, ctx, lc, err, "")
		return
	}

	response := responseDTO.NewMultiTransmissionsResponse("", "", http.StatusOK, transmissions)
	utils.WriteHttpHeader(w, ctx, http.StatusOK)
	pkg.Encode(response, w, lc)
}

// DeleteProcessedTransmissionsByAge deletes the processed transmissions if the current timestamp minus their created timestamp is less than the age parameter.
func (nc *TransmissionController) DeleteProcessedTransmissionsByAge(w http.ResponseWriter, r *http.Request) {
	lc := container.LoggingClientFrom(nc.dic.Get)
	ctx := r.Context()

	vars := mux.Vars(r)
	age, parsingErr := strconv.ParseInt(vars[v2.Age], 10, 64)
	if parsingErr != nil {
		err := errors.NewCommonEdgeX(errors.KindContractInvalid, "age format parsing failed", parsingErr)
		utils.WriteErrorResponse(w, ctx, lc, err, "")
		return
	}
	err := application.DeleteProcessedTransmissionsByAge(age, nc.dic)
	if err != nil {
		utils.WriteErrorResponse(w, ctx, lc, err, "")
		return
	}

	response := commonDTO.NewBaseResponse("", "", http.StatusAccepted)
	utils.WriteHttpHeader(w, ctx, http.StatusAccepted)
	// encode and send out the response
	pkg.Encode(response, w, lc)
}

// TransmissionsBySubscriptionName allows retrieval of the transmissions associated with the specified subscription name. Ordered by create timestamp descending.
func (tc *TransmissionController) TransmissionsBySubscriptionName(w http.ResponseWriter, r *http.Request) {
	lc := container.LoggingClientFrom(tc.dic.Get)
	ctx := r.Context()
	config := notificationContainer.ConfigurationFrom(tc.dic.Get)

	vars := mux.Vars(r)
	subscriptionName := vars[v2.Name]

	// parse URL query string for offset, limit
	offset, limit, _, err := utils.ParseGetAllObjectsRequestQueryString(r, 0, math.MaxInt32, -1, config.Service.MaxResultCount)
	if err != nil {
		utils.WriteErrorResponse(w, ctx, lc, err, "")
		return
	}
	transmissions, err := application.TransmissionsBySubscriptionName(offset, limit, subscriptionName, tc.dic)
	if err != nil {
		utils.WriteErrorResponse(w, ctx, lc, err, "")
		return
	}

	response := responseDTO.NewMultiTransmissionsResponse("", "", http.StatusOK, transmissions)
	utils.WriteHttpHeader(w, ctx, http.StatusOK)
	pkg.Encode(response, w, lc)
}

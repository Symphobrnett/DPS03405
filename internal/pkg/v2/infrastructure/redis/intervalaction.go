//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package redis

import (
	"encoding/json"
	"fmt"

	"github.com/edgexfoundry/edgex-go/internal/pkg/common"
	"github.com/edgexfoundry/edgex-go/internal/pkg/v2/utils"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/errors"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/v2/models"

	"github.com/gomodule/redigo/redis"
)

const (
	IntervalActionCollection     = "ss|ia"
	IntervalActionCollectionName = IntervalActionCollection + DBKeySeparator + v2.Name
)

// intervalActionStoredKey return the intervalAction's stored key which combines the collection name and object id
func intervalActionStoredKey(id string) string {
	return CreateKey(IntervalActionCollection, id)
}

// sendAddIntervalActionCmd sends redis command for adding intervalAction
func sendAddIntervalActionCmd(conn redis.Conn, storedKey string, action models.IntervalAction) errors.EdgeX {
	m, err := json.Marshal(action)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, "unable to JSON marshal intervalAction for Redis persistence", err)
	}
	_ = conn.Send(SET, storedKey, m)
	_ = conn.Send(ZADD, IntervalActionCollection, action.Modified, storedKey)
	_ = conn.Send(HSET, IntervalActionCollectionName, action.Name, storedKey)
	return nil
}

// addIntervalAction adds a new intervalAction into DB
func addIntervalAction(conn redis.Conn, action models.IntervalAction) (models.IntervalAction, errors.EdgeX) {
	exists, edgeXerr := objectIdExists(conn, intervalActionStoredKey(action.Id))
	if edgeXerr != nil {
		return action, errors.NewCommonEdgeXWrapper(edgeXerr)
	} else if exists {
		return action, errors.NewCommonEdgeX(errors.KindDuplicateName, fmt.Sprintf("intervalAction id %s already exists", action.Id), edgeXerr)
	}

	exists, edgeXerr = objectNameExists(conn, IntervalActionCollectionName, action.Name)
	if edgeXerr != nil {
		return action, errors.NewCommonEdgeXWrapper(edgeXerr)
	} else if exists {
		return action, errors.NewCommonEdgeX(errors.KindDuplicateName, fmt.Sprintf("intervalAction name %s already exists", action.Name), edgeXerr)
	}

	ts := utils.MakeTimestamp()
	if action.Created == 0 {
		action.Created = ts
	}
	action.Modified = ts

	storedKey := intervalActionStoredKey(action.Id)
	_ = conn.Send(MULTI)
	edgeXerr = sendAddIntervalActionCmd(conn, storedKey, action)
	if edgeXerr != nil {
		return action, errors.NewCommonEdgeXWrapper(edgeXerr)
	}
	_, err := conn.Do(EXEC)
	if err != nil {
		edgeXerr = errors.NewCommonEdgeX(errors.KindDatabaseError, "intervalAction creation failed", err)
	}

	return action, edgeXerr
}

// allIntervalActions queries intervalActions by offset and limit
func allIntervalActions(conn redis.Conn, offset, limit int) (intervalActions []models.IntervalAction, edgeXerr errors.EdgeX) {
	end := offset + limit - 1
	if limit == -1 { //-1 limit means that clients want to retrieve all remaining records after offset from DB, so specifying -1 for end
		end = limit
	}
	objects, edgeXerr := getObjectsByRevRange(conn, IntervalActionCollection, offset, end)
	if edgeXerr != nil {
		return intervalActions, errors.NewCommonEdgeXWrapper(edgeXerr)
	}

	intervalActions = make([]models.IntervalAction, len(objects))
	for i, o := range objects {
		action := models.IntervalAction{}
		err := json.Unmarshal(o, &action)
		if err != nil {
			return []models.IntervalAction{}, errors.NewCommonEdgeX(errors.KindDatabaseError, "intervalAction format parsing failed from the database", err)
		}
		intervalActions[i] = action
	}
	return intervalActions, nil
}

// intervalActionByName query intervalAction by name from DB
func intervalActionByName(conn redis.Conn, name string) (action models.IntervalAction, edgeXerr errors.EdgeX) {
	edgeXerr = getObjectByHash(conn, IntervalActionCollectionName, name, &action)
	if edgeXerr != nil {
		return action, errors.NewCommonEdgeXWrapper(edgeXerr)
	}
	return
}

// sendDeleteIntervalActionCmd sends redis command for deleting intervalAction
func sendDeleteIntervalActionCmd(conn redis.Conn, storedKey string, action models.IntervalAction) {
	_ = conn.Send(DEL, storedKey)
	_ = conn.Send(ZREM, IntervalActionCollection, storedKey)
	_ = conn.Send(HDEL, IntervalActionCollectionName, action.Name)
}

// deleteIntervalActionByName deletes the intervalAction by name
func deleteIntervalActionByName(conn redis.Conn, name string) errors.EdgeX {
	action, edgeXerr := intervalActionByName(conn, name)
	if edgeXerr != nil {
		return errors.NewCommonEdgeXWrapper(edgeXerr)
	}
	storedKey := intervalActionStoredKey(action.Id)
	_ = conn.Send(MULTI)
	sendDeleteIntervalActionCmd(conn, storedKey, action)
	_, err := conn.Do(EXEC)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindDatabaseError, "intervalAction deletion failed", err)
	}
	return nil
}

// intervalActionById query intervalAction by id from DB
func intervalActionById(conn redis.Conn, id string) (action models.IntervalAction, edgeXerr errors.EdgeX) {
	edgeXerr = getObjectById(conn, intervalActionStoredKey(id), &action)
	if edgeXerr != nil {
		return action, errors.NewCommonEdgeXWrapper(edgeXerr)
	}
	return
}

// updateIntervalAction updates an intervalAction
func updateIntervalAction(conn redis.Conn, action models.IntervalAction) errors.EdgeX {
	oldAction, edgeXerr := intervalActionByName(conn, action.Name)
	if edgeXerr != nil {
		return errors.NewCommonEdgeXWrapper(edgeXerr)
	}
	action.Modified = common.MakeTimestamp()
	storedKey := intervalActionStoredKey(action.Id)

	_ = conn.Send(MULTI)
	sendDeleteIntervalActionCmd(conn, storedKey, oldAction)
	edgeXerr = sendAddIntervalActionCmd(conn, storedKey, action)
	if edgeXerr != nil {
		return errors.NewCommonEdgeXWrapper(edgeXerr)
	}
	_, err := conn.Do(EXEC)
	if err != nil {
		return errors.NewCommonEdgeX(errors.KindDatabaseError, "intervalAction update failed", err)
	}
	return nil
}

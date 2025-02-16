/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package objects

import (
	"testing"
	"time"

	"gotest.tools/assert"

	"github.com/apache/incubator-yunikorn-core/pkg/common/resources"
	"github.com/apache/incubator-yunikorn-scheduler-interface/lib/go/common"
	"github.com/apache/incubator-yunikorn-scheduler-interface/lib/go/si"
)

func TestAskToString(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatal("String panic on nil ask")
		}
	}()
	var ask *AllocationAsk
	// ignore nil check from IDE we really want to do this
	askString := ask.String()
	assert.Equal(t, askString, "ask is nil", "Unexpected string returned for nil ask")
}

func TestNewAsk(t *testing.T) {
	res := resources.NewResourceFromMap(map[string]resources.Quantity{"first": 10})
	siAsk := &si.AllocationAsk{
		AllocationKey:  "ask-1",
		ApplicationID:  "app-1",
		MaxAllocations: 1,
		ResourceAsk:    res.ToProto(),
	}
	ask := NewAllocationAsk(siAsk)
	if ask == nil {
		t.Fatal("NewAllocationAsk create failed while it should not")
	}
	askStr := ask.String()
	expected := "AllocationKey ask-1, ApplicationID app-1, Resource map[first:10], PendingRepeats 1"
	assert.Equal(t, askStr, expected, "Strings should have been equal")
}

func TestPendingAskRepeat(t *testing.T) {
	res := resources.NewResourceFromMap(map[string]resources.Quantity{"first": 10})
	ask := newAllocationAsk("alloc-1", "app-1", res)
	assert.Equal(t, ask.GetPendingAskRepeat(), int32(1), "pending ask repeat should be 1")
	if !ask.updatePendingAskRepeat(1) {
		t.Errorf("increase of pending ask with 1 failed, expected repeat 2, current repeat: %d", ask.GetPendingAskRepeat())
	}
	if !ask.updatePendingAskRepeat(-1) {
		t.Errorf("decrease of pending ask with 1 failed, expected repeat 1, current repeat: %d", ask.GetPendingAskRepeat())
	}
	if ask.updatePendingAskRepeat(-2) {
		t.Errorf("decrease of pending ask with 2 did not fail, expected repeat 1, current repeat: %d", ask.GetPendingAskRepeat())
	}
	if !ask.updatePendingAskRepeat(-1) {
		t.Errorf("decrease of pending ask with 1 failed, expected repeat 0, current repeat: %d", ask.GetPendingAskRepeat())
	}
}

// the create time should not be manipulated but we need it for reservation testing
func TestGetCreateTime(t *testing.T) {
	res := resources.NewResourceFromMap(map[string]resources.Quantity{"first": 10})
	ask := newAllocationAskRepeat("alloc-1", "app-1", res, 2)
	created := ask.GetCreateTime()
	// move time 10 seconds back
	ask.createTime = created.Add(time.Second * -10)
	createdNow := ask.GetCreateTime()
	if createdNow.Equal(created) {
		t.Fatal("create time stamp should have been modified")
	}
}

func TestPlaceHolder(t *testing.T) {
	siAsk := &si.AllocationAsk{
		AllocationKey: "ask1",
		ApplicationID: "app1",
		PartitionName: "default",
	}
	ask := NewAllocationAsk(siAsk)
	assert.Assert(t, !ask.isPlaceholder(), "standard ask should not be a placeholder")
	assert.Equal(t, ask.getTaskGroup(), "", "standard ask should not have a TaskGroupName")
	siAsk = &si.AllocationAsk{
		AllocationKey: "ask1",
		ApplicationID: "app1",
		PartitionName: "default",
		TaskGroupName: "",
		Placeholder:   true,
	}
	ask = NewAllocationAsk(siAsk)
	var nilAsk *AllocationAsk
	assert.Equal(t, ask, nilAsk, "placeholder ask created without a TaskGroupName")
	siAsk.TaskGroupName = "testgroup"
	ask = NewAllocationAsk(siAsk)
	assert.Assert(t, ask != nilAsk, "placeholder ask creation failed unexpectedly")
	assert.Assert(t, ask.isPlaceholder(), "ask should have been a placeholder")
	assert.Equal(t, ask.getTaskGroup(), "testgroup", "TaskGroupName not set as expected")
}

func TestGetTimeout(t *testing.T) {
	siAsk := &si.AllocationAsk{
		AllocationKey: "ask1",
		ApplicationID: "app1",
		PartitionName: "default",
	}
	ask := NewAllocationAsk(siAsk)
	assert.Equal(t, ask.getTimeout(), time.Duration(0), "standard ask should not have timeout")
	siAsk = &si.AllocationAsk{
		AllocationKey:                "ask1",
		ApplicationID:                "app1",
		PartitionName:                "default",
		ExecutionTimeoutMilliSeconds: 10,
	}
	ask = NewAllocationAsk(siAsk)
	assert.Equal(t, ask.getTimeout(), 10*time.Millisecond, "ask timeout not set as expected")
}

func TestGetRequiredNode(t *testing.T) {
	tag := make(map[string]string)
	// unset case
	siAsk := &si.AllocationAsk{
		AllocationKey: "ask1",
		ApplicationID: "app1",
		PartitionName: "default",
		Tags:          tag,
	}
	ask := NewAllocationAsk(siAsk)
	assert.Equal(t, ask.GetRequiredNode(), "", "required node is empty as expected")
	// set case
	tag[common.DomainYuniKorn+common.KeyRequiredNode] = "NodeName"
	siAsk = &si.AllocationAsk{
		AllocationKey: "ask1",
		ApplicationID: "app1",
		PartitionName: "default",
		Tags:          tag,
	}
	ask = NewAllocationAsk(siAsk)
	assert.Equal(t, ask.GetRequiredNode(), "NodeName", "required node should be NodeName")
}

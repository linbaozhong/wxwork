// Copyright 2021 FastWeGo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package meeting_room_test

import (
	"fmt"

	"github.com/linbaozhong/wxwork/corporation"
	"github.com/linbaozhong/wxwork/corporation/apis/oa/meeting_room"
)

func ExampleAdd() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting_room.Add(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleList() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting_room.List(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleEdit() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting_room.Edit(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleDel() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting_room.Del(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetBookingInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting_room.GetBookingInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleBook() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting_room.Book(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleCancelBook() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := meeting_room.CancelBook(ctx, payload)

	fmt.Println(resp, err)
}

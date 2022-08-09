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

package corp_group_test

import (
	"fmt"

	"github.com/linbaozhong/wxwork/corporation"
	"github.com/linbaozhong/wxwork/corporation/apis/corp_group"
)

func ExampleListAppShareInfo() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := corp_group.ListAppShareInfo(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleGetToken() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := corp_group.GetToken(ctx, payload)

	fmt.Println(resp, err)
}

func ExampleTransferSession() {
	var ctx *corporation.App

	payload := []byte("{}")
	resp, err := corp_group.TransferSession(ctx, payload)

	fmt.Println(resp, err)
}

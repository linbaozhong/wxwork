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

// Package patrol_report 政民沟通/巡查上报
package patrol_report

import (
	"bytes"

	"github.com/linbaozhong/wxwork/corporation"
)

const (
	apiGetGridInfo       = "/cgi-bin/report/patrol/get_grid_info"
	apiGetCorpStatus     = "/cgi-bin/report/patrol/get_corp_status"
	apiGetUserStatus     = "/cgi-bin/report/patrol/get_user_status"
	apiCategoryStatistic = "/cgi-bin/report/patrol/category_statistic"
	apiGetOrderList      = "/cgi-bin/report/patrol/get_order_list"
	apiGetOrderInfo      = "/cgi-bin/report/patrol/get_order_info"
)

/*
获取配置的网格及网格负责人



See: https://work.weixin.qq.com/api/doc/90000/90135/93531

GET https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_grid_info?access_token=ACCESS_TOKEN
*/
func GetGridInfo(ctx *corporation.App) (resp []byte, err error) {
	return ctx.Client.HTTPGet(apiGetGridInfo)
}

/*
获取单位巡查上报数据统计



See: https://work.weixin.qq.com/api/doc/90000/90135/93532

POST https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_corp_status?access_token=ACCESS_TOKEN
*/
func GetCorpStatus(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetCorpStatus, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取个人巡查上报数据统计



See: https://work.weixin.qq.com/api/doc/90000/90135/93533

POST https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_user_status?access_token=ACCESS_TOKEN
*/
func GetUserStatus(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetUserStatus, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取上报事件分类统计



See: https://work.weixin.qq.com/api/doc/90000/90135/93534

POST https://qyapi.weixin.qq.com/cgi-bin/report/patrol/category_statistic?access_token=ACCESS_TOKEN
*/
func CategoryStatistic(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiCategoryStatistic, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取巡查上报事件列表



See: https://work.weixin.qq.com/api/doc/90000/90135/93536

POST https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_order_list?access_token=ACCESS_TOKEN
*/
func GetOrderList(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetOrderList, bytes.NewReader(payload), "application/json;charset=utf-8")
}

/*
获取巡查上报的事件详情信息



See: https://work.weixin.qq.com/api/doc/90000/90135/93535

POST https://qyapi.weixin.qq.com/cgi-bin/report/patrol/get_order_info?access_token=ACCESS_TOKEN
*/
func GetOrderInfo(ctx *corporation.App, payload []byte) (resp []byte, err error) {
	return ctx.Client.HTTPPost(apiGetOrderInfo, bytes.NewReader(payload), "application/json;charset=utf-8")
}

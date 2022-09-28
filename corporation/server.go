// Copyright 2020 FastWeGo
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

package corporation

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	eventtype "github.com/linbaozhong/wxwork/corporation/type/type_event"
	messagetype "github.com/linbaozhong/wxwork/corporation/type/type_message"

	"github.com/linbaozhong/wxwork/util"
)

/*
响应微信请求 或 推送消息/事件 的服务器
*/
type Server struct {
	Ctx *App
}

/*
EchoStr 服务器接口校验

msg_signature=ASDFQWEXZCVAQFASDFASDFSS
&timestamp=13500001234
&nonce=123412323
&echostr=ENCRYPT_STR
*/
func (s *Server) EchoStr(writer http.ResponseWriter, request *http.Request) {
	echoStr := request.URL.Query().Get("echostr")
	strs := []string{
		request.URL.Query().Get("timestamp"),
		request.URL.Query().Get("nonce"),
		s.Ctx.Config.Token,
		echoStr,
	}
	sort.Strings(strs)

	if s.Ctx.Corporation.Logger != nil {
		s.Ctx.Corporation.Logger.Println(strs)
	}

	h := sha1.New()
	_, _ = io.WriteString(h, strings.Join(strs, ""))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	if signature == request.URL.Query().Get("msg_signature") {

		// 解密 echoStr
		_, msg, _, err := util.AESDecryptMsg(echoStr, s.Ctx.Config.EncodingAESKey)
		if err != nil {
			return
		}

		writer.Write(msg)
		if s.Ctx.Corporation.Logger != nil {
			s.Ctx.Corporation.Logger.Println("echostr ", string(msg))
		}
	}

}

//
///*
//ParseXML 解析微信推送过来的消息/事件
//
//POST /api/callback?msg_signature=ASDFQWEXZCVAQFASDFASDFSS
//&timestamp=13500001234
//&nonce=123412323
//
//<xml>
//   <ToUserName><![CDATA[toUser]]></ToUserName>
//   <AgentID><![CDATA[toAgentID]]></AgentID>
//   <Encrypt><![CDATA[msg_encrypt]]></Encrypt>
//</xml>
//*/
//func (s *Server) ParseXML(request *http.Request) (m interface{}, err error) {
//	var body []byte
//	body, err = ioutil.ReadAll(request.Body)
//	if err != nil {
//		return
//	}
//
//	if s.Ctx.Corporation.Logger != nil {
//		s.Ctx.Corporation.Logger.Println(string(body))
//	}
//
//	// 加密格式 的消息
//	encryptMsg := messagetype.EncryptMessage{}
//	err = json.Unmarshal(body, &encryptMsg)
//	if err != nil {
//		return
//	}
//
//	// 验证签名
//	strs := []string{
//		request.URL.Query().Get("timestamp"),
//		request.URL.Query().Get("nonce"),
//		s.Ctx.Config.Token,
//		encryptMsg.Encrypt,
//	}
//	sort.Strings(strs)
//
//	h := sha1.New()
//	_, _ = io.WriteString(h, strings.Join(strs, ""))
//	signature := fmt.Sprintf("%x", h.Sum(nil))
//
//	if msgSignature := request.URL.Query().Get("msg_signature"); signature != msgSignature {
//		err = fmt.Errorf("%s != %s", signature, msgSignature)
//		return
//	}
//
//	// 解密
//	var xmlMsg []byte
//	_, xmlMsg, _, err = util.AESDecryptMsg(encryptMsg.Encrypt, s.Ctx.Config.EncodingAESKey)
//	if err != nil {
//		return
//	}
//	body = xmlMsg
//
//	if s.Ctx.Corporation.Logger != nil {
//		s.Ctx.Corporation.Logger.Println("AESDecryptMsg ", string(body))
//	}
//
//	return parseMsg(body)
//}

/*
ParseJSON 解析微信推送过来的消息/事件

POST /api/callback?msg_signature=ASDFQWEXZCVAQFASDFASDFSS
&timestamp=13500001234
&nonce=123412323

{
	"tousername": "wx5823bf96d3bd56c7",
	"encrypt": "No8isRLoXqFMhLlpe7R/DA7UbJ88DKJxDhJH/UVG3o1ib0Fhzdd3qWYHH/KL1mITv5qOCp2FbyILqfI7zazrp/ARgSHR177OCrv8O9UrMHWdnOaMXaz+mLd5X5VWm5r2J3Qpm+NdTQRPhHbce88frKF3wqTaZunKW7ae87bRZUfaq5tLFnyTsf6aiy0su3SsQ06dQGKPcHfYHY3upB881008Q9t9xeAZ/uqfXpYQgSLQfaX+fk/K/FQEl4QpLk94eD1YjluFY3uLnKp40zDyxgeWRAmgTtmx1eLwediVqZ8=",
	"agentid": "218"
}
*/
func (s *Server) ParseJSON(request *http.Request) (m interface{}, err error) {
	var body []byte
	body, err = ioutil.ReadAll(request.Body)
	if err != nil {
		return
	}

	if s.Ctx.Corporation.Logger != nil {
		s.Ctx.Corporation.Logger.Println(string(body))
	}

	// 加密格式 的消息
	encryptMsg := messagetype.EncryptMessage{}
	err = json.Unmarshal(body, &encryptMsg)
	if err != nil {
		return
	}

	// 验证签名
	strs := []string{
		request.URL.Query().Get("timestamp"),
		request.URL.Query().Get("nonce"),
		s.Ctx.Config.Token,
		encryptMsg.Encrypt,
	}
	sort.Strings(strs)

	h := sha1.New()
	_, _ = io.WriteString(h, strings.Join(strs, ""))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	if msgSignature := request.URL.Query().Get("msg_signature"); signature != msgSignature {
		err = fmt.Errorf("%s != %s", signature, msgSignature)
		return
	}

	// 解密
	var xmlMsg []byte
	_, xmlMsg, _, err = util.DecryptMsg(encryptMsg.Encrypt, s.Ctx.Config.EncodingAESKey)
	if err != nil {
		return
	}
	body = xmlMsg

	if s.Ctx.Corporation.Logger != nil {
		s.Ctx.Corporation.Logger.Println("AESDecryptMsg ", string(body))
	}

	return parseMsg(body)
}

// 解析消息/事件
func parseMsg(body []byte) (m interface{}, err error) {
	message := messagetype.Message{}
	err = json.Unmarshal(body, &message)
	fmt.Println(message)
	if err != nil {
		return
	}

	switch message.MsgType {
	case messagetype.MsgTypeText:
		msg := messagetype.MessageText{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeImage:
		msg := messagetype.MessageImage{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeVoice:
		msg := messagetype.MessageVoice{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeVideo:
		msg := messagetype.MessageVideo{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeLocation:
		msg := messagetype.MessageLocation{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeLink:
		msg := messagetype.MessageLink{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case messagetype.MsgTypeEvent:
		return parseEvent(body)
	}

	return
}

// ParseEvent 解析微信推送过来的事件
func parseEvent(body []byte) (m interface{}, err error) {
	event := eventtype.Event{}
	err = json.Unmarshal(body, &event)
	if err != nil {
		return
	}
	switch event.Event {
	// 事件
	case eventtype.EventTypeChangeContact:
		msg := eventtype.EventChangeContact{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		switch msg.ChangeType {
		case eventtype.EventTypeChangeContactCreateUser:
			msg := eventtype.EventChangeContactCreateUser{}
			err = json.Unmarshal(body, &msg)
			if err != nil {
				return
			}
			return msg, nil
		case eventtype.EventTypeChangeContactUpdateUser:
			msg := eventtype.EventChangeContactUpdateUser{}
			err = json.Unmarshal(body, &msg)
			if err != nil {
				return
			}
			return msg, nil
		case eventtype.EventTypeChangeContactDeleteUser:
			msg := eventtype.EventChangeContactDeleteUser{}
			err = json.Unmarshal(body, &msg)
			if err != nil {
				return
			}
			return msg, nil
		case eventtype.EventTypeChangeContactCreateParty:
			msg := eventtype.EventChangeContactCreateParty{}
			err = json.Unmarshal(body, &msg)
			if err != nil {
				return
			}
			return msg, nil
		case eventtype.EventTypeChangeContactUpdateParty:
			msg := eventtype.EventChangeContactUpdateParty{}
			err = json.Unmarshal(body, &msg)
			if err != nil {
				return
			}
			return msg, nil
		case eventtype.EventTypeChangeContactDeleteParty:
			msg := eventtype.EventChangeContactDeleteParty{}
			err = json.Unmarshal(body, &msg)
			if err != nil {
				return
			}
			return msg, nil
		case eventtype.EventTypeChangeContactUpdateTag:
			msg := eventtype.EventChangeContactUpdateTag{}
			err = json.Unmarshal(body, &msg)
			if err != nil {
				return
			}
			return msg, nil
		}
	case eventtype.EventTypeBatchJobResult:
		msg := eventtype.EventBatchJobResult{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeApproval:
		msg := eventtype.EventApproval{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeChangeExternalContact:
		msg := eventtype.EventChangeExternalContact{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		switch msg.ChangeType {
		case eventtype.EventTypeChangeExternalContactAddExternalContact:
			msg := eventtype.EventChangeExternalContactAddExternalContact{}
			err = json.Unmarshal(body, &msg)
			if err != nil {
				return
			}
			return msg, nil
		case eventtype.EventTypeChangeExternalContactAddHalfExternalContact:
			msg := eventtype.EventChangeExternalContactAddHalfExternalContact{}
			err = json.Unmarshal(body, &msg)
			if err != nil {
				return
			}
			return msg, nil
		case eventtype.EventTypeChangeExternalContactChangeExternalChat:
			msg := eventtype.EventChangeExternalContactChangeExternalChat{}
			err = json.Unmarshal(body, &msg)
			if err != nil {
				return
			}
			return msg, nil
		case eventtype.EventTypeChangeExternalContactDelExternalContact:
			msg := eventtype.EventChangeExternalContactDelExternalContact{}
			err = json.Unmarshal(body, &msg)
			if err != nil {
				return
			}
			return msg, nil
		case eventtype.EventTypeChangeExternalContactEditExternalContact:
			msg := eventtype.EventChangeExternalContactEditExternalContact{}
			err = json.Unmarshal(body, &msg)
			if err != nil {
				return
			}
			return msg, nil
		case eventtype.EventTypeChangeExternalContactDelFollowUser:
			msg := eventtype.EventChangeExternalContactDelFollowUser{}
			err = json.Unmarshal(body, &msg)
			if err != nil {
				return
			}
			return msg, nil
		}
	case eventtype.EventTypeTaskCardClick:
		msg := eventtype.EventTaskCardClick{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuView:
		msg := eventtype.EventMenuView{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuClick:
		msg := eventtype.EventMenuClick{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuLocationSelect:
		msg := eventtype.EventMenuLocationSelect{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuPicSysPhoto:
		msg := eventtype.EventMenuPicSysPhoto{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuPicSysPhotoOrAlbum:
		msg := eventtype.EventMenuPicSysPhotoOrAlbum{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuPicWeixin:
		msg := eventtype.EventMenuPicWeixin{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuScanCodePush:
		msg := eventtype.EventMenuScanCodePush{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	case eventtype.EventTypeMenuScanCodeWaitMsg:
		msg := eventtype.EventMenuScanCodeWaitMsg{}
		err = json.Unmarshal(body, &msg)
		if err != nil {
			return
		}
		return msg, nil
	}

	return
}

// Response 响应微信消息
func (s *Server) Response(writer http.ResponseWriter, request *http.Request, reply interface{}) (err error) {

	output := []byte("") // 默认回复
	if reply != nil {
		output, err = json.Marshal(reply)
		if err != nil {
			return
		}

		// 加密
		var message messagetype.ReplyEncryptMessage
		message, err = s.encryptReplyMessage(output)
		if err != nil {
			return
		}

		output, err = json.Marshal(message)
		if err != nil {
			return
		}

	}

	_, err = writer.Write(output)

	if s.Ctx.Corporation.Logger != nil {
		s.Ctx.Corporation.Logger.Println("Response: ", string(output))
	}

	return
}

// encryptReplyMessage 加密回复消息
func (s *Server) encryptReplyMessage(rawXmlMsg []byte) (replyEncryptMessage messagetype.ReplyEncryptMessage, err error) {
	cipherText, err := util.AESEncryptMsg([]byte(util.GetRandString(16)), rawXmlMsg, s.Ctx.Config.AgentId, s.Ctx.Config.EncodingAESKey)
	if err != nil {
		return
	}
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	nonce := util.GetRandString(6)

	strs := []string{
		timestamp,
		nonce,
		s.Ctx.Config.Token,
		cipherText,
	}
	sort.Strings(strs)
	h := sha1.New()
	_, _ = io.WriteString(h, strings.Join(strs, ""))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	return messagetype.ReplyEncryptMessage{
		Encrypt:      cipherText,
		MsgSignature: signature,
		TimeStamp:    timestamp,
		Nonce:        nonce,
	}, nil
}

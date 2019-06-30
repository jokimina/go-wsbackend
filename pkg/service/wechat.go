package service

import (
	"fmt"
	"github.com/silenceper/wechat/template"
	"go-wsbackend/pkg/model"
	"go-wsbackend/pkg/util"
)

func SendWechatTemplateMessage(tpl *template.Template, bindObj *model.FeedbackBindObj)(msgID int64, err error) {
	message := &template.Message{
		ToUser:     bindObj.OpenID,
		TemplateID: "c3M4soqvdhNZQQU0zHEWV2UIuLDjplKXmXd9XlzV850",
		FormID: bindObj.FormID,
		URL: "https://www.baidu.com",
		Page: fmt.Sprintf("pages/search/search?inputVal=%s", bindObj.Name),
		//URL: "",
		Data: map[string]*template.DataItem{
			"keyword1": {Value: fmt.Sprintf("%s -> %s", bindObj.Name, model.GetWasteNameByIndex(bindObj.Cats))},
			"keyword2": {Value: util.GetHumanTimeNow()},
			"keyword3": {Value: "感谢您为祖国的绿色事业添砖加瓦"},
		},
	}
	msgID, err = tpl.Send(message, true)
	return
}

package service

import (
	"fmt"
	"github.com/silenceper/wechat/template"
	"go-wsbackend/pkg/model"
	"go-wsbackend/pkg/util"
	"log"
)

func SendWechatTemplateMessage(tpl *template.Template, bindObj *model.FeedbackBindObj)(msgID int64, err error) {
	message := &template.Message{
		ToUser:     bindObj.OpenID,
		TemplateID: bindObj.TemplateID,
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
	if err != nil {
		log.Printf("send failed ,\n%s", err.Error())
	}
	return
}

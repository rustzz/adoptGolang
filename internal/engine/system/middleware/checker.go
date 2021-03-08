package middleware

import (
	"adoptGolang/internal/engine/system/middleware/errors"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"log"
)

func (mw *Middleware) AdminPerm(
	module func(obj events.MessageNewObject) error,
) func(obj events.MessageNewObject) error {
	return func(obj events.MessageNewObject) (err error) {
		// получение списка пользователей в беседе запрашивает права администратора,
		// поэтому использую его для выяснения прав администратора у группы
		parameters := params.NewMessagesGetConversationMembersBuilder()
		parameters.PeerID(obj.Message.PeerID)
		_, err = mw.Controller.Client.API.MessagesGetConversationMembers(parameters.Params)
		if err != nil {
			mw.Controller.CustomAPI.Send(
				obj.Message.PeerID,
				new(errors.BotNotAdminError).Error(),
				[10][]byte{},
			)
			return &errors.BotNotAdminError{}
		}

		if err = module(obj); err != nil { return }
		return
	}
}

// ForAdmin : использует в себе 2-х уровневую проверку,
// т.к. для проверки нужно получить список пользователей
// и позже выяснить право администратора
func (mw *Middleware) ForAdmin(
	module func(obj events.MessageNewObject) error,
) func(obj events.MessageNewObject) error {
	return mw.AdminPerm(func(obj events.MessageNewObject) (err error) {
		parameters := params.NewMessagesGetConversationMembersBuilder()
		parameters.PeerID(obj.Message.PeerID)
		resp, err := mw.Controller.Client.API.MessagesGetConversationMembers(parameters.Params)
		if err != nil { return }
		for _, profile := range resp.Items {
			log.Println(-profile.MemberID == mw.Controller.Client.GroupID && profile.IsAdmin)
			if profile.MemberID == obj.Message.FromID && profile.IsAdmin {
				if err = module(obj); err != nil { return }
				return
			}
		}
		{
			mw.Controller.CustomAPI.Send(
				obj.Message.PeerID,
				new(errors.NotAdminError).Error(),
				[10][]byte{},
			)
		}
		return &errors.NotAdminError{}
	})
}

func (mw *Middleware) ForChat(
	module func(obj events.MessageNewObject) error,
) func(obj events.MessageNewObject) error {
	return func(obj events.MessageNewObject) (err error) {
		if obj.Message.PeerID >= 2000000000 {
			if err = module(obj); err != nil { return }
			return
		}
		mw.Controller.CustomAPI.Send(
			obj.Message.PeerID,
			new(errors.IsNotChatError).Error(),
			[10][]byte{},
		)
		return &errors.IsNotChatError{}
	}
}


func (mw *Middleware) ForPrivateChat(
	module func(obj events.MessageNewObject) error,
) func(obj events.MessageNewObject) error {
	return func(obj events.MessageNewObject) (err error) {
		if !(obj.Message.PeerID >= 2000000000) {
			if err = module(obj); err != nil { return }
			return
		}
		mw.Controller.CustomAPI.Send(
			obj.Message.PeerID,
			new(errors.IsNotPrivateChatError).Error(),
			[10][]byte{},
		)
		return &errors.IsNotPrivateChatError{}
	}
}

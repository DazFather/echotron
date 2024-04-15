package echotron

import "testing"

func TestAPIResponseBase(_ *testing.T) {
	a := APIResponseBase{}
	a.Base()
}

func TestAPIResponseUpdate(_ *testing.T) {
	a := APIResponseUpdate{}
	a.Base()
}

func TestAPIResponseUser(_ *testing.T) {
	a := APIResponseUser{}
	a.Base()
}

func TestAPIResponseMessage(_ *testing.T) {
	a := APIResponseMessage{}
	a.Base()
}

func TestAPIResponseMessageArray(_ *testing.T) {
	a := APIResponseMessageArray{}
	a.Base()
}

func TestAPIResponseMessageID(_ *testing.T) {
	a := APIResponseMessageID{}
	a.Base()
}

func TestAPIResponseCommands(_ *testing.T) {
	a := APIResponseCommands{}
	a.Base()
}

func TestAPIResponseBool(_ *testing.T) {
	a := APIResponseBool{}
	a.Base()
}

func TestAPIResponseString(_ *testing.T) {
	a := APIResponseString{}
	a.Base()
}

func TestAPIResponseChat(_ *testing.T) {
	a := APIResponseChat{}
	a.Base()
}

func TestAPIResponseInviteLink(_ *testing.T) {
	a := APIResponseInviteLink{}
	a.Base()
}

func TestAPIResponseStickerSet(_ *testing.T) {
	a := APIResponseStickerSet{}
	a.Base()
}

func TestAPIResponseUserProfile(_ *testing.T) {
	a := APIResponseUserProfile{}
	a.Base()
}

func TestAPIResponseFile(_ *testing.T) {
	a := APIResponseFile{}
	a.Base()
}

func TestAPIResponseAdministrators(_ *testing.T) {
	a := APIResponseAdministrators{}
	a.Base()
}

func TestAPIResponseChatMember(_ *testing.T) {
	a := APIResponseChatMember{}
	a.Base()
}

func TestAPIResponseInteger(_ *testing.T) {
	a := APIResponseInteger{}
	a.Base()
}

func TestAPIResponsePoll(_ *testing.T) {
	a := APIResponsePoll{}
	a.Base()
}

func TestAPIResponseGameHighScore(_ *testing.T) {
	a := APIResponseGameHighScore{}
	a.Base()
}

func TestAPIResponseWebhook(_ *testing.T) {
	a := APIResponseWebhook{}
	a.Base()
}

func TestAPIResponseSentWebAppMessage(_ *testing.T) {
	a := APIResponseSentWebAppMessage{}
	a.Base()
}

func TestAPIResponseMenuButton(_ *testing.T) {
	a := APIResponseMenuButton{}
	a.Base()
}

func TestAPIResponseChatAdministratorRights(_ *testing.T) {
	a := APIResponseChatAdministratorRights{}
	a.Base()
}

func TestAPIResponseBotDescription(_ *testing.T) {
	a := APIResponseBotDescription{}
	a.Base()
}

func TestAPIResponseBotShortDescription(_ *testing.T) {
	a := APIResponseBotShortDescription{}
	a.Base()
}

func TestAPIResponseBusinessConnection(_ *testing.T) {
	a := APIResponseBusinessConnection{}
	a.Base()
}

func TestInputMediaPhoto(_ *testing.T) {
	i := InputMediaPhoto{}
	i.media()
	i.thumbnail()
	i.groupable()
}

func TestInputMediaVideo(_ *testing.T) {
	i := InputMediaVideo{}
	i.media()
	i.thumbnail()
	i.groupable()
}

func TestInputMediaAnimation(_ *testing.T) {
	i := InputMediaAnimation{}
	i.media()
	i.thumbnail()
}

func TestInputMediaAudio(_ *testing.T) {
	i := InputMediaAudio{}
	i.media()
	i.thumbnail()
	i.groupable()
}

func TestInputMediaDocument(_ *testing.T) {
	i := InputMediaDocument{}
	i.media()
	i.thumbnail()
	i.groupable()
}

func TestReactionTypeEmoji(_ *testing.T) {
	r := ReactionTypeEmoji{}
	r.ImplementsReactionType()
}

func TestReactionTypeCustomEmoji(_ *testing.T) {
	r := ReactionTypeCustomEmoji{}
	r.ImplementsReactionType()
}

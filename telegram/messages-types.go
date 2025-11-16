/*!
 * I am Karo  😊👍
 *
 * Contact me:
 *     https://www.karo.link/
 *     https://github.com/iamkaro
 *     https://www.linkedin.com/in/iamkaro
 *
 * Go-based library for developing Telegram client applications
 * https://github.com/iamkaro/Go-based-library-for-Telegram-Client.git
 * Copyright © 2020 developed.
 */

package telegram

import "fmt"

type (
	Messages []*Message
	Message  struct {
		Id                      int64
		SenderId                *messageSenderId
		ChatId                  int64
		IsOutgoing              bool
		IsPinned                bool
		IsFromOffline           bool
		CanBeSaved              bool
		HasTimestampedMedia     bool
		IsChannelPost           bool
		IsPaidStarSuggestedPost bool
		IsPaidTonSuggestedPost  bool
		ContainsUnreadMention   bool
		Date                    int64
		EditDate                int64
		ForwardInfo             *messageForwardInfo
		ImportInfo              object
		InteractionInfo         object
		UnreadReactions         array
		FactCheck               object
		SuggestedPostInfo       object
		ReplyTo                 object
		MessageThreadId         int64
		TopicId                 object
		SelfDestructType        object
		SelfDestructIn          float64
		AutoDeleteIn            float64
		ViaBotUserId            int64
		SenderBusinessBotUserId int64
		SenderBoostCount        int64
		PaidMessageStarCount    int64
		AuthorSignature         string
		MediaAlbumId            int64
		EffectId                int64
		RestrictionInfo         object
		Content                 *messageContent
		ReplyMarkup             *messageReplyMarkup
	}
	PublicMessageLink struct{ Link, Html string }

	messageSenderId struct {
		ChatId int64
		UserId int64
	}

	messageForwardInfo struct {
		Origin        *messageForwardOrigin
		Date          int32
		FromChatId    int64
		FromMessageId int64
	}
	messageForwardOrigin struct {
		Channel    *messageForwardOriginChannel
		HiddenUser *messageForwardOriginHiddenUser
		User       *messageForwardOriginUser
	}
	messageForwardOriginChannel struct {
		ChatId          int64
		MessageId       int64
		AuthorSignature string
	}
	messageForwardOriginHiddenUser struct{ SenderName string }
	messageForwardOriginUser       struct{ SenderUserId int32 }

	messageContent struct {
		Poll             *messagePoll
		Contact          *messageContact
		Location         *messageLocation
		Text             *messageText
		Photo            *messagePhoto
		Audio            *messageAudio
		Video            *messageVideo
		VideoNote        *messageVideoNote
		VoiceNote        *messageVoiceNote
		Document         *messageDocument
		ChatAddMembers   *messageChatAddMembers
		ChatChangePhoto  *messageChatChangePhoto
		ChatChangeTitle  *messageChatChangeTitle
		ChatDeleteMember *messageChatDeleteMember
		ChatDeletePhoto  bool
		ChatJoinByLink   bool
		ChatUpgradeFrom  *messageChatUpgradeFrom
		ChatUpgradeTo    *messageChatUpgradeTo
		ScreenshotTaken  bool
		Unsupported      bool
	}

	messagePoll struct {
		Id                 int64
		Question           string
		Options            []*pollOption
		TotalVoterCount    int32
		RecentVoterUserIds []int32
		IsAnonymous        bool
		Type               *pollType
		IsClosed           bool
	}
	pollOption struct {
		Text           string
		VoterCount     int32
		VotePercentage int32
		IsChosen       bool
		IsBeingChosen  bool
	}
	pollType struct {
		Quiz    *pollTypeQuiz
		Regular *pollTypeRegular
	}
	pollTypeQuiz    struct{ CorrectOptionId int32 }
	pollTypeRegular struct{ AllowMultipleAnswers bool }

	messageContact struct {
		PhoneNumber string
		FirstName   string
		LastName    string
		UserId      int32
	}
	messageLocation struct {
		Latitude   float64
		Longitude  float64
		LivePeriod int32
		ExpiresIn  int32
	}

	messageText struct {
		Text    string
		WebPage *webPage
	}
	webPage struct {
		Url        string
		DisplayUrl string
	}

	messageVoiceNote struct {
		Duration   int32
		Waveform   string
		MimeType   string
		Voice      *File
		Caption    string
		IsListened bool
	}
	messageChatAddMembers struct{ MemberUserIds []int32 }

	messagePhoto struct {
		HasStickers   bool
		MiniThumbnail *miniThumbnail
		Photos        []*photo
		Caption       string
		IsSecret      bool
	}
	messageAudio struct {
		Duration                int32
		Title                   string
		Performer               string
		FileName                string
		MimeType                string
		AlbumCoverMiniThumbnail *miniThumbnail
		AlbumCoverThumbnail     *photo
		Audio                   *File
		Caption                 string
	}
	messageVideo struct {
		Duration          int32
		Width             int32
		Height            int32
		FileName          string
		MimeType          string
		HasStickers       bool
		SupportsStreaming bool
		MiniThumbnail     *miniThumbnail
		Thumbnail         *photo
		Video             *File
		Caption           string
		IsSecret          bool
	}
	messageVideoNote struct {
		Duration      int32
		Length        int32
		MiniThumbnail *miniThumbnail
		Thumbnail     *photo
		Video         *File
		IsViewed      bool
		IsSecret      bool
	}
	messageDocument struct {
		FileName      string
		MimeType      string
		MiniThumbnail *miniThumbnail
		Thumbnail     *photo
		Document      *File
		Caption       string
	}
	messageChatChangePhoto struct {
		HasStickers   bool
		MiniThumbnail *miniThumbnail
		Photos        []*photo
	}
	miniThumbnail struct {
		Width  int32
		Height int32
		Data   string
	}
	photo struct {
		Type   string
		Photo  *File
		Width  int32
		Height int32
	}

	messageChatChangeTitle  struct{ Title string }
	messageChatDeleteMember struct{ UserId int32 }
	messageChatUpgradeTo    struct{ SuperGroupId int32 }
	messageChatUpgradeFrom  struct {
		Title        string
		BasicGroupId int32
	}

	messageReplyMarkup struct {
		Inline [][]*messageInlineButton
		Show   [][]*messageShowButton
	}
	messageInlineButton struct {
		Text string
		Type *messageInlineButtonType
	}
	messageShowButton struct {
		Text string
		Type *messageShowButtonType
	}
	messageInlineButtonType struct {
		Buy                  bool
		Callback             *string // data
		CallbackGame         bool
		CallbackWithPassword *string // data
		CopyText             *string // text
		LoginUrl             object
		SwitchInline         object
		Url                  *string // url
		User                 *int64  // user_id
		WebApp               *string // url
	}
	messageShowButtonType struct {
		RequestChat        object
		RequestLocation    bool
		RequestPhoneNumber bool
		RequestPoll        object
		RequestUsers       object
		Text               bool
		WebApp             *string // url
	}
)

func (it *messageInlineButtonType) String() string {
	switch {
	case it.Buy:
		return "is-buy"
	case it.Callback != nil:
		return "is-callback(" + decode(*it.Callback) + ")"
	case it.CallbackGame:
		return "is-game"
	case it.CallbackWithPassword != nil:
		return "is-callback-with-Password(" + decode(*it.CallbackWithPassword) + ")"
	case it.CopyText != nil:
		return "is-copy-text(" + (*it.CopyText) + ")"
	case it.LoginUrl != nil:
		return "is-login-url( ... )"
	case it.SwitchInline != nil:
		return "is-switch-inline( ... )"
	case it.Url != nil:
		return "is-url(" + (*it.Url) + ")"
	case it.User != nil:
		return "is-user(" + fmt.Sprint(*it.User) + ")"
	case it.WebApp != nil:
		return "is-web-app(" + (*it.WebApp) + ")"
	}
	return "Type error!"
}
func (it *messageShowButtonType) String() string {
	switch {
	case it.RequestChat != nil:
		return "is-request-chat( ... )"
	case it.RequestLocation:
		return "is-request-location"
	case it.RequestPhoneNumber:
		return "is-request-phone"
	case it.RequestPoll != nil:
		return "is-request-poll( ... )"
	case it.RequestUsers != nil:
		return "is-request-users( ... )"
	case it.Text:
		return "is-text"
	case it.WebApp != nil:
		return "is-web-app(" + (*it.WebApp) + ")"
	}
	return "Type error!"
}

func getMessage(value object) *Message {
	if CheckType(value, "message") {
		return &Message{
			Id:                      value.int64("id"),
			SenderId:                getMessageSenderId(value.Object("sender_id")),
			ChatId:                  value.int64("chat_id"),
			IsOutgoing:              value.bool("is_outgoing"),
			IsPinned:                value.bool("is_pinned"),
			IsFromOffline:           value.bool("is_from_offline"),
			CanBeSaved:              value.bool("can_be_saved"),
			HasTimestampedMedia:     value.bool("has_timestamped_media"),
			IsChannelPost:           value.bool("is_channel_post"),
			IsPaidStarSuggestedPost: value.bool("is_paid_star_suggested_post"),
			IsPaidTonSuggestedPost:  value.bool("is_paid_ton_suggested_post"),
			ContainsUnreadMention:   value.bool("contains_unread_mention"),
			Date:                    value.int64("date"),
			EditDate:                value.int64("edit_date"),
			ForwardInfo:             getMessageForwardInfo(value.Object("forward_info")),
			ImportInfo:              value.Object("import_info"),
			InteractionInfo:         value.Object("interaction_info"),
			UnreadReactions:         value.Array("unread_reactions"),
			FactCheck:               value.Object("fact_check"),
			SuggestedPostInfo:       value.Object("suggested_post_info"),
			ReplyTo:                 value.Object("reply_to"),
			MessageThreadId:         value.int64("message_thread_id"),
			TopicId:                 value.Object("topic_id"),
			SelfDestructType:        value.Object("self_destruct_type"),
			SelfDestructIn:          value.float64("self_destruct_in"),
			AutoDeleteIn:            value.float64("auto_delete_in"),
			ViaBotUserId:            value.int64("via_bot_user_id"),
			SenderBusinessBotUserId: value.int64("sender_business_bot_user_id"),
			SenderBoostCount:        value.int64("sender_boost_count"),
			PaidMessageStarCount:    value.int64("paid_message_star_count"),
			AuthorSignature:         value.string("author_signature"),
			MediaAlbumId:            value.int64("media_album_id"),
			EffectId:                value.int64("effect_id"),
			RestrictionInfo:         value.Object("restriction_info"),
			Content:                 getMessageContent(value.Object("content")),
			ReplyMarkup:             getReplyMarkup(value.Object("reply_markup")),
		}
	}
	return nil
}

func getMessageSenderId(value object) *messageSenderId {
	if CheckType(value, "messageSenderUser") ||
		CheckType(value, "messageSenderChat") {
		return &messageSenderId{
			ChatId: value.int64("chat_id"),
			UserId: value.int64("user_id"),
		}
	}
	return nil
}

func getMessagePoll(value object) *messagePoll {
	if CheckType(value, "messagePoll") {
		var poll = value.Object("poll")
		var Type = poll.Object("type")
		var options = poll.Array("options")
		var messagePoll = &messagePoll{
			Id:                 poll.int64("id"),
			Question:           poll.string("question"),
			Options:            make([]*pollOption, len(options)),
			TotalVoterCount:    poll.int32("total_voter_count"),
			RecentVoterUserIds: poll.Array("recent_voter_user_ids").int32(),
			IsAnonymous:        poll.bool("is_anonymous"),
			Type:               &pollType{Quiz: nil, Regular: nil},
			IsClosed:           poll.bool("is_closed"),
		}
		switch Type.string("@type") {
		case "pollTypeQuiz":
			messagePoll.Type.Quiz =
				&pollTypeQuiz{CorrectOptionId: Type.int32("correct_option_id")}
		case "pollTypeRegular":
			messagePoll.Type.Regular =
				&pollTypeRegular{AllowMultipleAnswers: Type.bool("allow_multiple_answers")}
		}
		options.Objects(func(index int, value object) {
			messagePoll.Options[index] = &pollOption{
				Text:           value.string("text"),
				VoterCount:     value.int32("voter_count"),
				VotePercentage: value.int32("vote_percentage"),
				IsChosen:       value.bool("is_chosen"),
				IsBeingChosen:  value.bool("is_being_chosen"),
			}
		})
		return messagePoll
	}
	return nil
}

func getMessageContact(value object) *messageContact {
	if CheckType(value, "messageContact") {
		var contact = value.Object("contact")
		return &messageContact{
			PhoneNumber: contact.string("phone_number"),
			FirstName:   contact.string("first_name"),
			LastName:    contact.string("last_name"),
			UserId:      contact.int32("user_id"),
		}
	}
	return nil
}

func getMessageLocation(value object) *messageLocation {
	if CheckType(value, "messageLocation") {
		var location = value.Object("location")
		return &messageLocation{
			Latitude:   location.float64("latitude"),
			Longitude:  location.float64("longitude"),
			LivePeriod: value.int32("live_period"),
			ExpiresIn:  value.int32("expires_in"),
		}
	}
	return nil
}

func getMessageText(value object) *messageText {
	if CheckType(value, "messageText") {
		var messageText = &messageText{
			Text:    value.Object("text").string("text"),
			WebPage: nil,
		}
		var web = value.Object("web_page")
		if CheckType(web, "webPage") {
			messageText.WebPage = &webPage{
				Url:        web.string("url"),
				DisplayUrl: web.string("display_url"),
			}
		}
		return messageText
	}
	return nil
}

func getMiniThumbnail(value object) *miniThumbnail {
	if CheckType(value, "miniThumbnail") {
		return &miniThumbnail{
			Width:  value.int32("width"),
			Height: value.int32("height"),
			Data:   value.string("data"),
		}
	}
	return nil
}

func getPhoto(value object) *photo {
	if CheckType(value, "photoSize") {
		return &photo{
			Type:   value.string("type"),
			Photo:  getFile(value.Object("photo")),
			Width:  value.int32("width"),
			Height: value.int32("height"),
		}
	}
	return nil
}

func getMessagePhoto(value object) *messagePhoto {
	if CheckType(value, "messagePhoto") {
		var photoData = value.Object("photo")
		var photos = photoData.Array("sizes")
		var messagePhoto = &messagePhoto{
			HasStickers:   photoData.bool("has_stickers"),
			MiniThumbnail: getMiniThumbnail(photoData.Object("minithumbnail")),
			Photos:        make([]*photo, len(photos)),
			Caption:       value.Object("caption").string("text"),
			IsSecret:      value.bool("is_secret"),
		}
		photos.Objects(func(index int, value object) {
			messagePhoto.Photos[index] = getPhoto(value)
		})
		return messagePhoto
	}
	return nil
}

func getMessageAudio(value object) *messageAudio {
	if CheckType(value, "messageAudio") {
		var audio = value.Object("audio")
		return &messageAudio{
			Duration:                audio.int32("duration"),
			Title:                   audio.string("title"),
			Performer:               audio.string("performer"),
			FileName:                audio.string("file_name"),
			MimeType:                audio.string("mime_type"),
			AlbumCoverMiniThumbnail: getMiniThumbnail(audio.Object("album_cover_minithumbnail")),
			AlbumCoverThumbnail:     getPhoto(audio.Object("album_cover_thumbnail")),
			Audio:                   getFile(audio.Object("audio")),
			Caption:                 value.Object("caption").string("text"),
		}
	}
	return nil
}

func getMessageVideo(value object) *messageVideo {
	if CheckType(value, "messageVideo") {
		var video = value.Object("video")
		return &messageVideo{
			Duration:          video.int32("duration"),
			Width:             video.int32("width"),
			Height:            video.int32("height"),
			FileName:          video.string("file_name"),
			MimeType:          video.string("mime_type"),
			HasStickers:       video.bool("has_stickers"),
			SupportsStreaming: video.bool("supports_streaming"),
			MiniThumbnail:     getMiniThumbnail(video.Object("minithumbnail")),
			Thumbnail:         getPhoto(video.Object("thumbnail")),
			Video:             getFile(video.Object("video")),
			Caption:           value.Object("caption").string("text"),
			IsSecret:          value.bool("is_secret"),
		}
	}
	return nil
}

func getMessageVideoNote(value object) *messageVideoNote {
	if CheckType(value, "messageVideoNote") {
		var videoNote = value.Object("video_note")
		return &messageVideoNote{
			Duration:      videoNote.int32("duration"),
			Length:        videoNote.int32("length"),
			MiniThumbnail: getMiniThumbnail(videoNote.Object("minithumbnail")),
			Thumbnail:     getPhoto(videoNote.Object("thumbnail")),
			Video:         getFile(videoNote.Object("video")),
			IsViewed:      value.bool("is_viewed"),
			IsSecret:      value.bool("is_secret"),
		}
	}
	return nil
}

func getMessageVoiceNote(value object) *messageVoiceNote {
	if CheckType(value, "messageVoiceNote") {
		var voiceNote = value.Object("voice_note")
		return &messageVoiceNote{
			Duration:   voiceNote.int32("duration"),
			Waveform:   voiceNote.string("waveform"),
			MimeType:   voiceNote.string("mime_type"),
			Voice:      getFile(voiceNote.Object("voice")),
			Caption:    value.Object("caption").string("text"),
			IsListened: value.bool("is_listened"),
		}
	}
	return nil
}

func getMessageDocument(value object) *messageDocument {
	if CheckType(value, "messageDocument") {
		var document = value.Object("document")
		return &messageDocument{
			FileName:      document.string("file_name"),
			MimeType:      document.string("mime_type"),
			MiniThumbnail: getMiniThumbnail(document.Object("minithumbnail")),
			Thumbnail:     getPhoto(document.Object("thumbnail")),
			Document:      getFile(document.Object("document")),
			Caption:       value.Object("caption").string("text"),
		}
	}
	return nil
}

func getMessageChatAddMembers(value object) *messageChatAddMembers {
	if CheckType(value, "messageChatAddMembers") {
		return &messageChatAddMembers{MemberUserIds: value.Array("member_user_ids").int32()}
	}
	return nil
}

func getMessageChatChangePhoto(value object) *messageChatChangePhoto {
	if CheckType(value, "messageChatChangePhoto") {
		var photoData = value.Object("photo")
		var photos = photoData.Array("sizes")
		var messageChatChangePhoto = &messageChatChangePhoto{
			HasStickers:   photoData.bool("has_stickers"),
			MiniThumbnail: getMiniThumbnail(photoData.Object("minithumbnail")),
			Photos:        make([]*photo, len(photos)),
		}
		photos.Objects(func(index int, value object) {
			messageChatChangePhoto.Photos[index] = getPhoto(value)
		})
		return messageChatChangePhoto
	}
	return nil
}

func getMessageChatChangeTitle(value object) *messageChatChangeTitle {
	if CheckType(value, "messageChatChangeTitle") {
		return &messageChatChangeTitle{Title: value.string("title")}
	}
	return nil
}

func getMessageChatDeleteMember(value object) *messageChatDeleteMember {
	if CheckType(value, "messageChatDeleteMember") {
		return &messageChatDeleteMember{UserId: value.int32("user_id")}
	}
	return nil
}

func getMessageChatUpgradeFrom(value object) *messageChatUpgradeFrom {
	if CheckType(value, "messageChatUpgradeFrom") {
		return &messageChatUpgradeFrom{
			Title:        value.string("title"),
			BasicGroupId: value.int32("basic_group_id"),
		}
	}
	return nil
}

func getMessageChatUpgradeTo(value object) *messageChatUpgradeTo {
	if CheckType(value, "messageChatUpgradeTo") {
		return &messageChatUpgradeTo{SuperGroupId: value.int32("supergroup_id")}
	}
	return nil
}

func getMessageContent(value object) *messageContent {
	if value != nil {
		return &messageContent{
			Poll:             getMessagePoll(value),
			Contact:          getMessageContact(value),
			Location:         getMessageLocation(value),
			Text:             getMessageText(value),
			Photo:            getMessagePhoto(value),
			Audio:            getMessageAudio(value),
			Video:            getMessageVideo(value),
			VideoNote:        getMessageVideoNote(value),
			VoiceNote:        getMessageVoiceNote(value),
			Document:         getMessageDocument(value),
			ChatAddMembers:   getMessageChatAddMembers(value),
			ChatChangePhoto:  getMessageChatChangePhoto(value),
			ChatChangeTitle:  getMessageChatChangeTitle(value),
			ChatDeleteMember: getMessageChatDeleteMember(value),
			ChatDeletePhoto:  value.string("@type") == "messageChatDeletePhoto",
			ChatJoinByLink:   value.string("@type") == "messageChatJoinByLink",
			ChatUpgradeFrom:  getMessageChatUpgradeFrom(value),
			ChatUpgradeTo:    getMessageChatUpgradeTo(value),
			ScreenshotTaken:  value.string("@type") == "messageScreenshotTaken",
			Unsupported:      value.string("@type") == "messageUnsupported",
		}
	}
	return nil
}

func getMessageForwardOrigin(value object) *messageForwardOrigin {
	if value != nil {
		var messageForwardOrigin = &messageForwardOrigin{
			Channel:    nil,
			HiddenUser: nil,
			User:       nil,
		}
		switch value.string("@type") {
		case "messageForwardOriginChannel":
			messageForwardOrigin.Channel = &messageForwardOriginChannel{
				ChatId:          value.int64("chat_id"),
				MessageId:       value.int64("message_id"),
				AuthorSignature: value.string("author_signature"),
			}
		case "messageForwardOriginHiddenUser":
			messageForwardOrigin.HiddenUser = &messageForwardOriginHiddenUser{
				SenderName: value.string("sender_name"),
			}
		case "messageForwardOriginUser":
			messageForwardOrigin.User = &messageForwardOriginUser{
				SenderUserId: value.int32("sender_user_id"),
			}
		}
		return messageForwardOrigin
	}
	return nil
}

func getMessageForwardInfo(value object) *messageForwardInfo {
	if CheckType(value, "messageForwardInfo") {
		return &messageForwardInfo{
			Origin:        getMessageForwardOrigin(value.Object("origin")),
			Date:          value.int32("date"),
			FromChatId:    value.int64("from_chat_id"),
			FromMessageId: value.int64("from_message_id"),
		}
	}
	return nil
}

func getPublicMessageLink(value object) *PublicMessageLink {
	if CheckType(value, "publicMessageLink") {
		return &PublicMessageLink{
			Link: value.string("link"),
			Html: value.string("html"),
		}
	}
	return nil
}

func getReplyMarkup(value object) *messageReplyMarkup {
	if CheckType(value, "replyMarkupInlineKeyboard") {
		var (
			rows           = value.Array("rows")
			keyboard       = make([][]*messageInlineButton, len(rows))
			arr      array = nil
		)
		for i, a := range rows {
			arr = a.([]interface{})
			keyboard[i] = make([]*messageInlineButton, len(arr))
			for j, v := range arr {
				keyboard[i][j] = getInlineButton(v.(map[string]interface{}))
			}
		}
		return &messageReplyMarkup{Inline: keyboard, Show: nil}
	}
	if CheckType(value, "replyMarkupShowKeyboard") {
		var (
			rows           = value.Array("rows")
			keyboard       = make([][]*messageShowButton, len(rows))
			arr      array = nil
		)
		for i, a := range rows {
			arr = a.([]interface{})
			keyboard[i] = make([]*messageShowButton, len(arr))
			for j, v := range arr {
				keyboard[i][j] = getShowButton(v.(map[string]interface{}))
			}
		}
		return &messageReplyMarkup{Inline: nil, Show: keyboard}
	}
	return nil
}

func getInlineButton(value object) *messageInlineButton {
	if CheckType(value, "inlineKeyboardButton") {
		return &messageInlineButton{
			Text: value.string("text"),
			Type: getInlineButtonType(value.Object("type")),
		}
	}
	return nil
}
func getShowButton(value object) *messageShowButton {
	if CheckType(value, "keyboardButton") {
		return &messageShowButton{
			Text: value.string("text"),
			Type: getShowButtonType(value.Object("type")),
		}
	}
	return nil
}

func getInlineButtonType(value object) *messageInlineButtonType {
	return &messageInlineButtonType{
		Buy:          getAsType(value, "inlineKeyboardButtonTypeBuy") != nil,
		Callback:     getValue[string](value, "inlineKeyboardButtonTypeCallback", "data"),
		CallbackGame: getAsType(value, "inlineKeyboardButtonTypeCallbackGame") != nil,
		CallbackWithPassword: getValue[string](value,
			"inlineKeyboardButtonTypeCallbackWithPassword", "data"),
		CopyText:     getValue[string](value, "inlineKeyboardButtonTypeCopyText", "text"),
		LoginUrl:     getAsType(value, "inlineKeyboardButtonTypeLoginUrl"),
		SwitchInline: getAsType(value, "inlineKeyboardButtonTypeSwitchInline"),
		Url:          getValue[string](value, "inlineKeyboardButtonTypeUrl", "url"),
		User:         getValue[int64](value, "inlineKeyboardButtonTypeUser", "user_id"),
		WebApp:       getValue[string](value, "inlineKeyboardButtonTypeWebApp", "url"),
	}
}
func getShowButtonType(value object) *messageShowButtonType {
	return &messageShowButtonType{
		RequestChat:        getAsType(value, "keyboardButtonTypeRequestChat"),
		RequestLocation:    getAsType(value, "keyboardButtonTypeRequestLocation") != nil,
		RequestPhoneNumber: getAsType(value, "keyboardButtonTypeRequestPhoneNumber") != nil,
		RequestPoll:        getAsType(value, "keyboardButtonTypeRequestPoll"),
		RequestUsers:       getAsType(value, "keyboardButtonTypeRequestUsers"),
		Text:               getAsType(value, "keyboardButtonTypeText") != nil,
		WebApp:             getValue[string](value, "keyboardButtonTypeWebApp", "url"),
	}
}

func getAsType(value object, typeName string) object {
	if CheckType(value, typeName) {
		return value
	}
	return nil
}
func getValue[T string | int64](from object, typeName, field string) *T {
	if CheckType(from, typeName) {
		var out, ok = from[field].(T)
		if ok {
			return &out
		}
		log("getValue", fmt.Sprintf("Can not parse(%+v) as %T", from[field], out))
	}
	return nil
}

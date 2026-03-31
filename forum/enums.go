package forum

type AuthenticationTokenGrantType string

const (
	AuthenticationTokenGrantTypeClientCredentials AuthenticationTokenGrantType = "client_credentials"
)

type CategoriesOrder string

const (
	CategoriesOrderNatural CategoriesOrder = "natural"
	CategoriesOrderList    CategoriesOrder = "list"
)

type ChatboxDuration string

const (
	ChatboxDurationDay   ChatboxDuration = "day"
	ChatboxDurationWeek  ChatboxDuration = "week"
	ChatboxDurationMonth ChatboxDuration = "month"
)

type ChatboxPostMessageRoomID string

const (
	ChatboxPostMessageRoomID1  ChatboxPostMessageRoomID = "1"
	ChatboxPostMessageRoomID2  ChatboxPostMessageRoomID = "2"
	ChatboxPostMessageRoomID3  ChatboxPostMessageRoomID = "3"
	ChatboxPostMessageRoomID4  ChatboxPostMessageRoomID = "4"
	ChatboxPostMessageRoomID13 ChatboxPostMessageRoomID = "13"
)

type ChatboxRoomID string

const (
	ChatboxRoomID1  ChatboxRoomID = "1"
	ChatboxRoomID2  ChatboxRoomID = "2"
	ChatboxRoomID3  ChatboxRoomID = "3"
	ChatboxRoomID4  ChatboxRoomID = "4"
	ChatboxRoomID13 ChatboxRoomID = "13"
)

type ConversationsDeleteDeleteType string

const (
	ConversationsDeleteDeleteTypeDelete       ConversationsDeleteDeleteType = "delete"
	ConversationsDeleteDeleteTypeDeleteIgnore ConversationsDeleteDeleteType = "delete_ignore"
)

type ConversationsFolder string

const (
	ConversationsFolderAll                ConversationsFolder = "all"
	ConversationsFolderUnread             ConversationsFolder = "unread"
	ConversationsFolderGroups             ConversationsFolder = "groups"
	ConversationsFolderMarket             ConversationsFolder = "market"
	ConversationsFolderMarketReplacements ConversationsFolder = "market_replacements"
	ConversationsFolderStaff              ConversationsFolder = "staff"
	ConversationsFolderGiveaways          ConversationsFolder = "giveaways"
	ConversationsFolderP2p                ConversationsFolder = "p2p"
)

type ConversationsOrder string

const (
	ConversationsOrderNatural        ConversationsOrder = "natural"
	ConversationsOrderNaturalReverse ConversationsOrder = "natural_reverse"
)

type FormsCreateFormID string

const (
	FormsCreateFormID1 FormsCreateFormID = "1"
)

type ForumsOrder string

const (
	ForumsOrderNatural ForumsOrder = "natural"
	ForumsOrderList    ForumsOrder = "list"
)

type NotificationsType string

const (
	NotificationsTypeMarket   NotificationsType = "market"
	NotificationsTypeNomarket NotificationsType = "nomarket"
)

type PagesOrder string

const (
	PagesOrderNatural PagesOrder = "natural"
	PagesOrderList    PagesOrder = "list"
)

type PostsOrder string

const (
	PostsOrderNatural          PostsOrder = "natural"
	PostsOrderNaturalReverse   PostsOrder = "natural_reverse"
	PostsOrderPostLikes        PostsOrder = "post_likes"
	PostsOrderPostLikesReverse PostsOrder = "post_likes_reverse"
)

type RoomIDModel string

const (
	RoomIDModel1  RoomIDModel = "1"
	RoomIDModel2  RoomIDModel = "2"
	RoomIDModel3  RoomIDModel = "3"
	RoomIDModel4  RoomIDModel = "4"
	RoomIDModel13 RoomIDModel = "13"
)

type ThreadsClaimCurrency string

const (
	ThreadsClaimCurrencyRub ThreadsClaimCurrency = "rub"
	ThreadsClaimCurrencyUah ThreadsClaimCurrency = "uah"
	ThreadsClaimCurrencyKzt ThreadsClaimCurrency = "kzt"
	ThreadsClaimCurrencyByn ThreadsClaimCurrency = "byn"
	ThreadsClaimCurrencyUsd ThreadsClaimCurrency = "usd"
	ThreadsClaimCurrencyEur ThreadsClaimCurrency = "eur"
	ThreadsClaimCurrencyGbp ThreadsClaimCurrency = "gbp"
	ThreadsClaimCurrencyCny ThreadsClaimCurrency = "cny"
	ThreadsClaimCurrencyTry ThreadsClaimCurrency = "try"
)

type ThreadsClaimPayClaim string

const (
	ThreadsClaimPayClaimNow   ThreadsClaimPayClaim = "now"
	ThreadsClaimPayClaimLater ThreadsClaimPayClaim = "later"
)

type ThreadsClaimReplyGroup string

const (
	ThreadsClaimReplyGroup0   ThreadsClaimReplyGroup = "0"
	ThreadsClaimReplyGroup2   ThreadsClaimReplyGroup = "2"
	ThreadsClaimReplyGroup21  ThreadsClaimReplyGroup = "21"
	ThreadsClaimReplyGroup22  ThreadsClaimReplyGroup = "22"
	ThreadsClaimReplyGroup23  ThreadsClaimReplyGroup = "23"
	ThreadsClaimReplyGroup60  ThreadsClaimReplyGroup = "60"
	ThreadsClaimReplyGroup351 ThreadsClaimReplyGroup = "351"
)

type ThreadsClaimTransferType string

const (
	ThreadsClaimTransferTypeGuarantor ThreadsClaimTransferType = "guarantor"
	ThreadsClaimTransferTypeSafe      ThreadsClaimTransferType = "safe"
	ThreadsClaimTransferTypeNotsafe   ThreadsClaimTransferType = "notsafe"
)

type ThreadsCreateContestContestType string

const (
	ThreadsCreateContestContestTypeByFinishDate ThreadsCreateContestContestType = "by_finish_date"
)

type ThreadsCreateContestLengthOption string

const (
	ThreadsCreateContestLengthOptionMinutes ThreadsCreateContestLengthOption = "minutes"
	ThreadsCreateContestLengthOptionHours   ThreadsCreateContestLengthOption = "hours"
	ThreadsCreateContestLengthOptionDays    ThreadsCreateContestLengthOption = "days"
)

type ThreadsCreateContestPrizeDataUpgrade string

const (
	ThreadsCreateContestPrizeDataUpgrade1  ThreadsCreateContestPrizeDataUpgrade = "1"
	ThreadsCreateContestPrizeDataUpgrade6  ThreadsCreateContestPrizeDataUpgrade = "6"
	ThreadsCreateContestPrizeDataUpgrade12 ThreadsCreateContestPrizeDataUpgrade = "12"
	ThreadsCreateContestPrizeDataUpgrade14 ThreadsCreateContestPrizeDataUpgrade = "14"
	ThreadsCreateContestPrizeDataUpgrade17 ThreadsCreateContestPrizeDataUpgrade = "17"
	ThreadsCreateContestPrizeDataUpgrade19 ThreadsCreateContestPrizeDataUpgrade = "19"
	ThreadsCreateContestPrizeDataUpgrade20 ThreadsCreateContestPrizeDataUpgrade = "20"
	ThreadsCreateContestPrizeDataUpgrade21 ThreadsCreateContestPrizeDataUpgrade = "21"
	ThreadsCreateContestPrizeDataUpgrade22 ThreadsCreateContestPrizeDataUpgrade = "22"
)

type ThreadsCreateContestPrizeType string

const (
	ThreadsCreateContestPrizeTypeMoney    ThreadsCreateContestPrizeType = "money"
	ThreadsCreateContestPrizeTypeUpgrades ThreadsCreateContestPrizeType = "upgrades"
)

type ThreadsCreateContestReplyGroup string

const (
	ThreadsCreateContestReplyGroup0   ThreadsCreateContestReplyGroup = "0"
	ThreadsCreateContestReplyGroup2   ThreadsCreateContestReplyGroup = "2"
	ThreadsCreateContestReplyGroup21  ThreadsCreateContestReplyGroup = "21"
	ThreadsCreateContestReplyGroup22  ThreadsCreateContestReplyGroup = "22"
	ThreadsCreateContestReplyGroup23  ThreadsCreateContestReplyGroup = "23"
	ThreadsCreateContestReplyGroup60  ThreadsCreateContestReplyGroup = "60"
	ThreadsCreateContestReplyGroup351 ThreadsCreateContestReplyGroup = "351"
)

type ThreadsCreateReplyGroup string

const (
	ThreadsCreateReplyGroup0   ThreadsCreateReplyGroup = "0"
	ThreadsCreateReplyGroup2   ThreadsCreateReplyGroup = "2"
	ThreadsCreateReplyGroup21  ThreadsCreateReplyGroup = "21"
	ThreadsCreateReplyGroup22  ThreadsCreateReplyGroup = "22"
	ThreadsCreateReplyGroup23  ThreadsCreateReplyGroup = "23"
	ThreadsCreateReplyGroup60  ThreadsCreateReplyGroup = "60"
	ThreadsCreateReplyGroup351 ThreadsCreateReplyGroup = "351"
)

type ThreadsDirection string

const (
	ThreadsDirectionAsc  ThreadsDirection = "asc"
	ThreadsDirectionDesc ThreadsDirection = "desc"
)

type ThreadsEditReplyGroup string

const (
	ThreadsEditReplyGroup0   ThreadsEditReplyGroup = "0"
	ThreadsEditReplyGroup2   ThreadsEditReplyGroup = "2"
	ThreadsEditReplyGroup21  ThreadsEditReplyGroup = "21"
	ThreadsEditReplyGroup22  ThreadsEditReplyGroup = "22"
	ThreadsEditReplyGroup23  ThreadsEditReplyGroup = "23"
	ThreadsEditReplyGroup60  ThreadsEditReplyGroup = "60"
	ThreadsEditReplyGroup351 ThreadsEditReplyGroup = "351"
)

type ThreadsOrder string

const (
	ThreadsOrderPostDate       ThreadsOrder = "post_date"
	ThreadsOrderLastPostDate   ThreadsOrder = "last_post_date"
	ThreadsOrderReplyCount     ThreadsOrder = "reply_count"
	ThreadsOrderReplyCountAsc  ThreadsOrder = "reply_count_asc"
	ThreadsOrderFirstPostLikes ThreadsOrder = "first_post_likes"
	ThreadsOrderVoteCount      ThreadsOrder = "vote_count"
)

type ThreadsPeriod string

const (
	ThreadsPeriodDay   ThreadsPeriod = "day"
	ThreadsPeriodWeek  ThreadsPeriod = "week"
	ThreadsPeriodMonth ThreadsPeriod = "month"
	ThreadsPeriodYear  ThreadsPeriod = "year"
)

type ThreadsState string

const (
	ThreadsStateActive ThreadsState = "active"
	ThreadsStateClosed ThreadsState = "closed"
)

type UsersClaimState string

const (
	UsersClaimStateActive   UsersClaimState = "active"
	UsersClaimStateSolved   UsersClaimState = "solved"
	UsersClaimStateRejected UsersClaimState = "rejected"
	UsersClaimStateSettled  UsersClaimState = "settled"
)

type UsersContentType string

const (
	UsersContentTypePost               UsersContentType = "post"
	UsersContentTypePostComment        UsersContentType = "post_comment"
	UsersContentTypeProfilePost        UsersContentType = "profile_post"
	UsersContentTypeProfilePostComment UsersContentType = "profile_post_comment"
)

type UsersEditAllowInviteGroup string

const (
	UsersEditAllowInviteGroupNone     UsersEditAllowInviteGroup = "none"
	UsersEditAllowInviteGroupMembers  UsersEditAllowInviteGroup = "members"
	UsersEditAllowInviteGroupFollowed UsersEditAllowInviteGroup = "followed"
)

type UsersEditAllowPostProfile string

const (
	UsersEditAllowPostProfileNone     UsersEditAllowPostProfile = "none"
	UsersEditAllowPostProfileMembers  UsersEditAllowPostProfile = "members"
	UsersEditAllowPostProfileFollowed UsersEditAllowPostProfile = "followed"
)

type UsersEditAllowReceiveNewsFeed string

const (
	UsersEditAllowReceiveNewsFeedNone     UsersEditAllowReceiveNewsFeed = "none"
	UsersEditAllowReceiveNewsFeedMembers  UsersEditAllowReceiveNewsFeed = "members"
	UsersEditAllowReceiveNewsFeedFollowed UsersEditAllowReceiveNewsFeed = "followed"
)

type UsersEditAllowSendPersonalConversation string

const (
	UsersEditAllowSendPersonalConversationNone     UsersEditAllowSendPersonalConversation = "none"
	UsersEditAllowSendPersonalConversationMembers  UsersEditAllowSendPersonalConversation = "members"
	UsersEditAllowSendPersonalConversationFollowed UsersEditAllowSendPersonalConversation = "followed"
)

type UsersEditAllowViewProfile string

const (
	UsersEditAllowViewProfileNone     UsersEditAllowViewProfile = "none"
	UsersEditAllowViewProfileMembers  UsersEditAllowViewProfile = "members"
	UsersEditAllowViewProfileFollowed UsersEditAllowViewProfile = "followed"
)

type UsersEditGender string

const (
	UsersEditGenderEmpty  UsersEditGender = ""
	UsersEditGenderMale   UsersEditGender = "male"
	UsersEditGenderFemale UsersEditGender = "female"
)

type UsersEditLanguageID string

const (
	UsersEditLanguageID1 UsersEditLanguageID = "1"
	UsersEditLanguageID2 UsersEditLanguageID = "2"
)

type UsersEditTimezone string

const (
	UsersEditTimezonePacificMidway               UsersEditTimezone = "Pacific/Midway"
	UsersEditTimezonePacificHonolulu             UsersEditTimezone = "Pacific/Honolulu"
	UsersEditTimezonePacificMarquesas            UsersEditTimezone = "Pacific/Marquesas"
	UsersEditTimezoneAmericaAnchorage            UsersEditTimezone = "America/Anchorage"
	UsersEditTimezoneAmericaLosAngeles           UsersEditTimezone = "America/Los_Angeles"
	UsersEditTimezoneAmericaSantaIsabel          UsersEditTimezone = "America/Santa_Isabel"
	UsersEditTimezoneAmericaTijuana              UsersEditTimezone = "America/Tijuana"
	UsersEditTimezoneAmericaDenver               UsersEditTimezone = "America/Denver"
	UsersEditTimezoneAmericaChihuahua            UsersEditTimezone = "America/Chihuahua"
	UsersEditTimezoneAmericaPhoenix              UsersEditTimezone = "America/Phoenix"
	UsersEditTimezoneAmericaChicago              UsersEditTimezone = "America/Chicago"
	UsersEditTimezoneAmericaBelize               UsersEditTimezone = "America/Belize"
	UsersEditTimezoneAmericaMexicoCity           UsersEditTimezone = "America/Mexico_City"
	UsersEditTimezonePacificEaster               UsersEditTimezone = "Pacific/Easter"
	UsersEditTimezoneAmericaNewYork              UsersEditTimezone = "America/New_York"
	UsersEditTimezoneAmericaHavana               UsersEditTimezone = "America/Havana"
	UsersEditTimezoneAmericaBogota               UsersEditTimezone = "America/Bogota"
	UsersEditTimezoneAmericaCaracas              UsersEditTimezone = "America/Caracas"
	UsersEditTimezoneAmericaHalifax              UsersEditTimezone = "America/Halifax"
	UsersEditTimezoneAmericaGooseBay             UsersEditTimezone = "America/Goose_Bay"
	UsersEditTimezoneAmericaAsuncion             UsersEditTimezone = "America/Asuncion"
	UsersEditTimezoneAmericaSantiago             UsersEditTimezone = "America/Santiago"
	UsersEditTimezoneAmericaCuiaba               UsersEditTimezone = "America/Cuiaba"
	UsersEditTimezoneAmericaLaPaz                UsersEditTimezone = "America/La_Paz"
	UsersEditTimezoneAmericaStJohns              UsersEditTimezone = "America/St_Johns"
	UsersEditTimezoneAmericaArgentinaBuenosAires UsersEditTimezone = "America/Argentina/Buenos_Aires"
	UsersEditTimezoneAmericaArgentinaSanLuis     UsersEditTimezone = "America/Argentina/San_Luis"
	UsersEditTimezoneAmericaArgentinaMendoza     UsersEditTimezone = "America/Argentina/Mendoza"
	UsersEditTimezoneAtlanticStanley             UsersEditTimezone = "Atlantic/Stanley"
	UsersEditTimezoneAmericaGodthab              UsersEditTimezone = "America/Godthab"
	UsersEditTimezoneAmericaMontevideo           UsersEditTimezone = "America/Montevideo"
	UsersEditTimezoneAmericaSaoPaulo             UsersEditTimezone = "America/Sao_Paulo"
	UsersEditTimezoneAmericaMiquelon             UsersEditTimezone = "America/Miquelon"
	UsersEditTimezoneAmericaNoronha              UsersEditTimezone = "America/Noronha"
	UsersEditTimezoneAtlanticCapeVerde           UsersEditTimezone = "Atlantic/Cape_Verde"
	UsersEditTimezoneAtlanticAzores              UsersEditTimezone = "Atlantic/Azores"
	UsersEditTimezoneEuropeLondon                UsersEditTimezone = "Europe/London"
	UsersEditTimezoneAfricaCasablanca            UsersEditTimezone = "Africa/Casablanca"
	UsersEditTimezoneAtlanticReykjavik           UsersEditTimezone = "Atlantic/Reykjavik"
	UsersEditTimezoneEuropeAmsterdam             UsersEditTimezone = "Europe/Amsterdam"
	UsersEditTimezoneAfricaAlgiers               UsersEditTimezone = "Africa/Algiers"
	UsersEditTimezoneAfricaWindhoek              UsersEditTimezone = "Africa/Windhoek"
	UsersEditTimezoneAfricaTunis                 UsersEditTimezone = "Africa/Tunis"
	UsersEditTimezoneEuropeAthens                UsersEditTimezone = "Europe/Athens"
	UsersEditTimezoneAfricaJohannesburg          UsersEditTimezone = "Africa/Johannesburg"
	UsersEditTimezoneEuropeKaliningrad           UsersEditTimezone = "Europe/Kaliningrad"
	UsersEditTimezoneAsiaAmman                   UsersEditTimezone = "Asia/Amman"
	UsersEditTimezoneAsiaBeirut                  UsersEditTimezone = "Asia/Beirut"
	UsersEditTimezoneAfricaCairo                 UsersEditTimezone = "Africa/Cairo"
	UsersEditTimezoneAsiaJerusalem               UsersEditTimezone = "Asia/Jerusalem"
	UsersEditTimezoneAsiaGaza                    UsersEditTimezone = "Asia/Gaza"
	UsersEditTimezoneAsiaDamascus                UsersEditTimezone = "Asia/Damascus"
	UsersEditTimezoneEuropeMoscow                UsersEditTimezone = "Europe/Moscow"
	UsersEditTimezoneEuropeMinsk                 UsersEditTimezone = "Europe/Minsk"
	UsersEditTimezoneAfricaNairobi               UsersEditTimezone = "Africa/Nairobi"
	UsersEditTimezoneAsiaTehran                  UsersEditTimezone = "Asia/Tehran"
	UsersEditTimezoneAsiaDubai                   UsersEditTimezone = "Asia/Dubai"
	UsersEditTimezoneAsiaYerevan                 UsersEditTimezone = "Asia/Yerevan"
	UsersEditTimezoneAsiaBaku                    UsersEditTimezone = "Asia/Baku"
	UsersEditTimezoneIndianMauritius             UsersEditTimezone = "Indian/Mauritius"
	UsersEditTimezoneAsiaKabul                   UsersEditTimezone = "Asia/Kabul"
	UsersEditTimezoneAsiaYekaterinburg           UsersEditTimezone = "Asia/Yekaterinburg"
	UsersEditTimezoneAsiaTashkent                UsersEditTimezone = "Asia/Tashkent"
	UsersEditTimezoneAsiaKolkata                 UsersEditTimezone = "Asia/Kolkata"
	UsersEditTimezoneAsiaKathmandu               UsersEditTimezone = "Asia/Kathmandu"
	UsersEditTimezoneAsiaNovosibirsk             UsersEditTimezone = "Asia/Novosibirsk"
	UsersEditTimezoneAsiaDhaka                   UsersEditTimezone = "Asia/Dhaka"
	UsersEditTimezoneAsiaAlmaty                  UsersEditTimezone = "Asia/Almaty"
	UsersEditTimezoneAsiaRangoon                 UsersEditTimezone = "Asia/Rangoon"
	UsersEditTimezoneAsiaKrasnoyarsk             UsersEditTimezone = "Asia/Krasnoyarsk"
	UsersEditTimezoneAsiaBangkok                 UsersEditTimezone = "Asia/Bangkok"
	UsersEditTimezoneAsiaIrkutsk                 UsersEditTimezone = "Asia/Irkutsk"
	UsersEditTimezoneAsiaHongKong                UsersEditTimezone = "Asia/Hong_Kong"
	UsersEditTimezoneAsiaSingapore               UsersEditTimezone = "Asia/Singapore"
	UsersEditTimezoneAustraliaPerth              UsersEditTimezone = "Australia/Perth"
	UsersEditTimezoneAsiaYakutsk                 UsersEditTimezone = "Asia/Yakutsk"
	UsersEditTimezoneAsiaTokyo                   UsersEditTimezone = "Asia/Tokyo"
	UsersEditTimezoneAsiaSeoul                   UsersEditTimezone = "Asia/Seoul"
	UsersEditTimezoneAustraliaAdelaide           UsersEditTimezone = "Australia/Adelaide"
	UsersEditTimezoneAustraliaDarwin             UsersEditTimezone = "Australia/Darwin"
	UsersEditTimezoneAsiaVladivostok             UsersEditTimezone = "Asia/Vladivostok"
	UsersEditTimezoneAsiaMagadan                 UsersEditTimezone = "Asia/Magadan"
	UsersEditTimezoneAustraliaBrisbane           UsersEditTimezone = "Australia/Brisbane"
	UsersEditTimezoneAustraliaSydney             UsersEditTimezone = "Australia/Sydney"
	UsersEditTimezonePacificNoumea               UsersEditTimezone = "Pacific/Noumea"
	UsersEditTimezonePacificNorfolk              UsersEditTimezone = "Pacific/Norfolk"
	UsersEditTimezoneAsiaAnadyr                  UsersEditTimezone = "Asia/Anadyr"
	UsersEditTimezonePacificAuckland             UsersEditTimezone = "Pacific/Auckland"
	UsersEditTimezonePacificFiji                 UsersEditTimezone = "Pacific/Fiji"
	UsersEditTimezonePacificChatham              UsersEditTimezone = "Pacific/Chatham"
	UsersEditTimezonePacificTongatapu            UsersEditTimezone = "Pacific/Tongatapu"
	UsersEditTimezonePacificApia                 UsersEditTimezone = "Pacific/Apia"
	UsersEditTimezonePacificKiritimati           UsersEditTimezone = "Pacific/Kiritimati"
)

type UsersLikeType string

const (
	UsersLikeTypeLike  UsersLikeType = "like"
	UsersLikeTypeLike2 UsersLikeType = "like2"
)

type UsersOrder string

const (
	UsersOrderNatural           UsersOrder = "natural"
	UsersOrderFollowDate        UsersOrder = "follow_date"
	UsersOrderFollowDateReverse UsersOrder = "follow_date_reverse"
)

type UsersType string

const (
	UsersTypeMarket   UsersType = "market"
	UsersTypeNomarket UsersType = "nomarket"
)

package market

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// StringOrInt handles JSON fields that can be either a string or an integer.
type StringOrInt struct {
	StringValue *string
	IntValue    *int
}

func (s StringOrInt) MarshalJSON() ([]byte, error) {
	if s.IntValue != nil {
		return json.Marshal(*s.IntValue)
	}
	if s.StringValue != nil {
		return json.Marshal(*s.StringValue)
	}
	return []byte("null"), nil
}

func (s *StringOrInt) UnmarshalJSON(data []byte) error {
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	switch v := raw.(type) {
	case float64:
		i := int(v)
		s.IntValue = &i
	case string:
		s.StringValue = &v
	case nil:
		// both remain nil
	default:
		return fmt.Errorf("StringOrInt: unexpected type %T", raw)
	}
	return nil
}

func (s StringOrInt) String() string {
	if s.IntValue != nil {
		return strconv.Itoa(*s.IntValue)
	}
	if s.StringValue != nil {
		return *s.StringValue
	}
	return ""
}

type AIPriceResponse struct {
	Price      *int            `json:"price,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type AIPriceResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type AddParams struct {
	// Proxy ip or host. Required if **proxy_row** is not specified.
	ProxyIP *string `json:"proxy_ip,omitempty"`
	// Proxy password. Required if **proxy_row** is not specified.
	ProxyPass *string `json:"proxy_pass,omitempty"`
	// Proxy port. Required if **proxy_row** is not specified.
	ProxyPort *int `json:"proxy_port,omitempty"`
	// Proxy list in String format ip:port:user:pass. Each proxy must be start with new line (use \r\n separator)
	ProxyRow *string `json:"proxy_row,omitempty"`
	// Proxy username. Required if **proxy_row** is not specified.
	ProxyUser *string `json:"proxy_user,omitempty"`
}

type AddResponse struct {
	Success    *bool           `json:"success,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type AddResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type AllParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
}

type AutoBuyPriceResponse struct {
	Price      *int            `json:"price,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type AutoBuyPriceResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type BalanceExchange struct {
	From       interface{}     `json:"from,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
	To         interface{}     `json:"to,omitempty"`
}

type BalanceExchangeFrom struct {
	Field12345 *BalanceModel               `json:"12345,omitempty"`
	Balance    *BalanceExchangeFromBalance `json:"balance,omitempty"`
}

type BalanceExchangeFrom12345 struct {
	Balance     *string     `json:"balance,omitempty"`
	BalanceID   *int        `json:"balance_id,omitempty"`
	CustomTitle interface{} `json:"custom_title,omitempty"`
	FullTitle   *string     `json:"fullTitle,omitempty"`
	MerchantID  *int        `json:"merchant_id,omitempty"`
	Title       *string     `json:"title,omitempty"`
	Type_       *string     `json:"type,omitempty"`
	UserID      *int        `json:"user_id,omitempty"`
}

type BalanceExchangeFromBalance struct {
	Balance          *string `json:"balance,omitempty"`
	ConvertedBalance *int    `json:"convertedBalance,omitempty"`
	FullTitle        *string `json:"fullTitle,omitempty"`
	Title            *string `json:"title,omitempty"`
	Type_            *string `json:"type,omitempty"`
}

type BalanceExchangeSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type BalanceExchangeTo struct {
	Balance *UserModel `json:"balance,omitempty"`
}

type BalanceExchangeToBalance struct {
	ActiveItemsCount            *int                                     `json:"active_items_count,omitempty"`
	ActivityVisible             *bool                                    `json:"activity_visible,omitempty"`
	Age                         *int                                     `json:"age,omitempty"`
	Balance                     *string                                  `json:"balance,omitempty"`
	Balances                    []BalanceExchangeToBalanceBalancesItem   `json:"balances,omitempty"`
	BumpItemPeriod              *int                                     `json:"bump_item_period,omitempty"`
	CanEdit                     *bool                                    `json:"can_edit,omitempty"`
	CanFollow                   *bool                                    `json:"can_follow,omitempty"`
	CanIgnore                   *bool                                    `json:"can_ignore,omitempty"`
	CanPostProfile              *bool                                    `json:"can_post_profile,omitempty"`
	CanViewProfile              *bool                                    `json:"can_view_profile,omitempty"`
	CanViewProfilePosts         *bool                                    `json:"can_view_profile_posts,omitempty"`
	CanWarn                     *bool                                    `json:"can_warn,omitempty"`
	ContestCount                *int                                     `json:"contest_count,omitempty"`
	ConvWelcomeMessage          *string                                  `json:"conv_welcome_message,omitempty"`
	ConvertedBalance            *int                                     `json:"convertedBalance,omitempty"`
	ConvertedDeposit            *int                                     `json:"convertedDeposit,omitempty"`
	ConvertedHold               *int                                     `json:"convertedHold,omitempty"`
	Currency                    *string                                  `json:"currency,omitempty"`
	CurrencyPhrase              *string                                  `json:"currencyPhrase,omitempty"`
	CustomAccountDownloadFormat *string                                  `json:"custom_account_download_format,omitempty"`
	CustomFields                *BalanceExchangeToBalanceCustomFields    `json:"custom_fields,omitempty"`
	CustomTitle                 *string                                  `json:"custom_title,omitempty"`
	Deposit                     *int                                     `json:"deposit,omitempty"`
	Dob                         *BalanceExchangeToBalanceDob             `json:"dob,omitempty"`
	FeedbackData                interface{}                              `json:"feedback_data,omitempty"`
	Hold                        *string                                  `json:"hold,omitempty"`
	Homepage                    *string                                  `json:"homepage,omitempty"`
	IMAPData                    interface{}                              `json:"imap_data,omitempty"`
	IsAdmin                     *bool                                    `json:"is_admin,omitempty"`
	IsBanned                    *bool                                    `json:"is_banned,omitempty"`
	IsFollowed                  *bool                                    `json:"is_followed,omitempty"`
	IsIgnored                   *bool                                    `json:"is_ignored,omitempty"`
	IsModerator                 *bool                                    `json:"is_moderator,omitempty"`
	IsStaff                     *bool                                    `json:"is_staff,omitempty"`
	IsSuperAdmin                *bool                                    `json:"is_super_admin,omitempty"`
	JoinedDate                  *int                                     `json:"joined_date,omitempty"`
	LastActivity                *int                                     `json:"last_activity,omitempty"`
	Like2Count                  *int                                     `json:"like2_count,omitempty"`
	LikeCount                   *int                                     `json:"like_count,omitempty"`
	Location                    *string                                  `json:"location,omitempty"`
	MarketCustomTitle           *string                                  `json:"market_custom_title,omitempty"`
	MaxDiscountPercent          *int                                     `json:"max_discount_percent,omitempty"`
	MessageCount                *int                                     `json:"message_count,omitempty"`
	PaidMailLeft                *int                                     `json:"paid_mail_left,omitempty"`
	PublicTags                  []BalanceExchangeToBalancePublicTagsItem `json:"public_tags,omitempty"`
	RegisterDate                *int                                     `json:"register_date,omitempty"`
	Rendered                    *BalanceExchangeToBalanceRendered        `json:"rendered,omitempty"`
	RestoreCount                *int                                     `json:"restore_count,omitempty"`
	RestoreData                 interface{}                              `json:"restore_data,omitempty"`
	ShortLink                   *string                                  `json:"short_link,omitempty"`
	SoldItemsCount              *int                                     `json:"sold_items_count,omitempty"`
	Tags                        interface{}                              `json:"tags,omitempty"`
	TelegramClient              interface{}                              `json:"telegram_client,omitempty"`
	TrophyPoints                *int                                     `json:"trophy_points,omitempty"`
	UserAllowAskDiscount        *bool                                    `json:"user_allow_ask_discount,omitempty"`
	UserID                      *int                                     `json:"user_id,omitempty"`
	UserTitle                   *string                                  `json:"user_title,omitempty"`
	Username                    *string                                  `json:"username,omitempty"`
	ViewURL                     *string                                  `json:"view_url,omitempty"`
	Visible                     *bool                                    `json:"visible,omitempty"`
	WarningPoints               *int                                     `json:"warning_points,omitempty"`
}

type BalanceExchangeToBalanceBalancesItem struct {
	Balance          *string     `json:"balance,omitempty"`
	BalanceID        *int        `json:"balance_id,omitempty"`
	ConvertedBalance *float64    `json:"convertedBalance,omitempty"`
	CustomTitle      interface{} `json:"custom_title,omitempty"`
	FullTitle        *string     `json:"fullTitle,omitempty"`
	MerchantID       *int        `json:"merchant_id,omitempty"`
	Title            *string     `json:"title,omitempty"`
	Type_            *string     `json:"type,omitempty"`
	UserID           *int        `json:"user_id,omitempty"`
}

type BalanceExchangeToBalanceCustomFields struct {
	Field4                *string       `json:"_4,omitempty"`
	AllowSelfUnban        []interface{} `json:"allowSelfUnban,omitempty"`
	BanReason             *string       `json:"ban_reason,omitempty"`
	Discord               *string       `json:"discord,omitempty"`
	FavoriteAnime         *string       `json:"favoriteAnime,omitempty"`
	FavoritePorn          *string       `json:"favoritePorn,omitempty"`
	FavoriteVape          *string       `json:"favoriteVape,omitempty"`
	Github                *string       `json:"github,omitempty"`
	Jabber                *string       `json:"jabber,omitempty"`
	LztAwardUserTrophy    *string       `json:"lztAwardUserTrophy,omitempty"`
	LztLikesIncreasing    *string       `json:"lztLikesIncreasing,omitempty"`
	LztLikesZeroing       *string       `json:"lztLikesZeroing,omitempty"`
	LztSympathyIncreasing *string       `json:"lztSympathyIncreasing,omitempty"`
	LztSympathyZeroing    *string       `json:"lztSympathyZeroing,omitempty"`
	LztUnbanAmount        *string       `json:"lztUnbanAmount,omitempty"`
	MaecenasValue         *string       `json:"maecenasValue,omitempty"`
	Matrix                *string       `json:"matrix,omitempty"`
	ScamURL               *string       `json:"scamURL,omitempty"`
	Steam                 *string       `json:"steam,omitempty"`
	Telegram              *string       `json:"telegram,omitempty"`
	Vk                    *string       `json:"vk,omitempty"`
}

type BalanceExchangeToBalanceDob struct {
	Day   *int `json:"day,omitempty"`
	Month *int `json:"month,omitempty"`
	Year  *int `json:"year,omitempty"`
}

type BalanceExchangeToBalanceIMAPData struct {
	DomainZone *BalanceExchangeToBalanceIMAPDataDomainZone `json:"domain.zone,omitempty"`
}

type BalanceExchangeToBalanceIMAPDataDomainZone struct {
	Domain     *string `json:"domain,omitempty"`
	IMAPServer *string `json:"imap_server,omitempty"`
	Port       *int    `json:"port,omitempty"`
	Secure     *bool   `json:"secure,omitempty"`
}

type BalanceExchangeToBalancePublicTagsItem struct {
	BackgroundColor *string `json:"background_color,omitempty"`
	TagID           *int    `json:"tag_id,omitempty"`
	Title           *string `json:"title,omitempty"`
}

type BalanceExchangeToBalanceRendered struct {
	Avatars     *BalanceExchangeToBalanceRenderedAvatars `json:"avatars,omitempty"`
	Backgrounds interface{}                              `json:"backgrounds,omitempty"`
	Link        *string                                  `json:"link,omitempty"`
	Username    *string                                  `json:"username,omitempty"`
}

type BalanceExchangeToBalanceRenderedAvatars struct {
	L *string `json:"l,omitempty"`
	M *string `json:"m,omitempty"`
	S *string `json:"s,omitempty"`
}

type BalanceExchangeToBalanceRenderedBackgrounds struct {
	L *string `json:"l,omitempty"`
	M *string `json:"m,omitempty"`
}

type BalanceExchangeToBalanceTagsItem struct {
	Bc                   *string `json:"bc,omitempty"`
	ForOwnedAccountsOnly *bool   `json:"forOwnedAccountsOnly,omitempty"`
	IsDefault            *bool   `json:"isDefault,omitempty"`
	TagID                *int    `json:"tag_id,omitempty"`
	Title                *string `json:"title,omitempty"`
}

type BalanceExchangeToBalanceTelegramClient struct {
	TelegramAPIHash        *string `json:"telegram_api_hash,omitempty"`
	TelegramAPIID          *string `json:"telegram_api_id,omitempty"`
	TelegramAppVersion     *string `json:"telegram_app_version,omitempty"`
	TelegramDeviceModel    *string `json:"telegram_device_model,omitempty"`
	TelegramLangCode       *string `json:"telegram_lang_code,omitempty"`
	TelegramLangPack       *string `json:"telegram_lang_pack,omitempty"`
	TelegramSystemLangCode *string `json:"telegram_system_lang_code,omitempty"`
	TelegramSystemVersion  *string `json:"telegram_system_version,omitempty"`
}

type BalanceModel struct {
	Balance     *string     `json:"balance,omitempty"`
	BalanceID   *int        `json:"balance_id,omitempty"`
	CustomTitle interface{} `json:"custom_title,omitempty"`
	FullTitle   *string     `json:"fullTitle,omitempty"`
	MerchantID  *int        `json:"merchant_id,omitempty"`
	Title       *string     `json:"title,omitempty"`
	Type_       *string     `json:"type,omitempty"`
	UserID      *int        `json:"user_id,omitempty"`
}

type BatchResponse struct {
	Jobs       *BatchResponseJobs `json:"jobs,omitempty"`
	SystemInfo *RespSystemInfo    `json:"system_info,omitempty"`
}

type BatchResponseJobs struct {
	JobID *BatchResponseJobsJobID `json:"job_id,omitempty"`
}

type BatchResponseJobsJobID struct {
	JobError  *string `json:"_job_error,omitempty"`
	JobResult *string `json:"_job_result,omitempty"`
}

type BatchResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type BattleNetParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// Guarantee type.
	Eg *CategorySearchEg `json:"eg,omitempty"`
	// List of games.
	Game []int `json:"game[],omitempty"`
	// Number of days the account has been offline.
	Daybreak *int `json:"daybreak,omitempty"`
	// List of allowed countries.
	Country []string `json:"country[],omitempty"`
	// List of disallowed countries.
	NotCountry []string `json:"not_country[],omitempty"`
	// Has linked mobile.
	Tel *CategorySearchTel `json:"tel,omitempty"`
	// Can edit BattleTag.
	EditBtag *CategorySearchEditBtag `json:"edit_btag,omitempty"`
	// Can edit full name.
	ChangeableFn *CategorySearchChangeableFn `json:"changeable_fn,omitempty"`
	// Real id.
	RealID *CategorySearchRealID `json:"real_id,omitempty"`
	// Has disabled parent control.
	ParentControl *CategorySearchParentControl `json:"parent_control,omitempty"`
	// Has no bans.
	NoBans *CategorySearchNoBans `json:"no_bans,omitempty"`
	// Minimum balance.
	BalanceMin *int `json:"balance_min,omitempty"`
	// Maximum balance.
	BalanceMax *int `json:"balance_max,omitempty"`
}

type BattleNetResponse struct {
	CacheTTL        *int                         `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                        `json:"hasNextPage,omitempty"`
	Items           []BattleNetResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                         `json:"lastModified,omitempty"`
	Page            *int                         `json:"page,omitempty"`
	PerPage         *int                         `json:"perPage,omitempty"`
	SearchUrl       *string                      `json:"searchUrl,omitempty"`
	ServerTime      *int                         `json:"serverTime,omitempty"`
	StickyItems     []interface{}                `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo              `json:"system_info,omitempty"`
	TotalItems      *int                         `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}                  `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                        `json:"wasCached,omitempty"`
}

type BattleNetResponseItemsItem struct {
	AccountLinks               []interface{}                                         `json:"accountLinks,omitempty"`
	AccountLastActivity        *int                                                  `json:"account_last_activity,omitempty"`
	AllowAskDiscount           *int                                                  `json:"allow_ask_discount,omitempty"`
	BattlenetBans              *string                                               `json:"battlenetBans,omitempty"`
	BattlenetGames             interface{}                                           `json:"battlenetGames,omitempty"`
	BattlenetTransactions      []BattleNetResponseItemsItemBattlenetTransactionsItem `json:"battlenetTransactions,omitempty"`
	BattlenetBalance           *string                                               `json:"battlenet_balance,omitempty"`
	BattlenetCanChangeTag      *int                                                  `json:"battlenet_can_change_tag,omitempty"`
	BattlenetChangeFullName    *int                                                  `json:"battlenet_change_full_name,omitempty"`
	BattlenetConvertedBalance  *int                                                  `json:"battlenet_converted_balance,omitempty"`
	BattlenetCountry           *string                                               `json:"battlenet_country,omitempty"`
	BattlenetItemID            *int                                                  `json:"battlenet_item_id,omitempty"`
	BattlenetLastActivity      *int                                                  `json:"battlenet_last_activity,omitempty"`
	BattlenetMobile            *int                                                  `json:"battlenet_mobile,omitempty"`
	BattlenetParentControl     *int                                                  `json:"battlenet_parent_control,omitempty"`
	BattlenetRealIDEnabled     *int                                                  `json:"battlenet_real_id_enabled,omitempty"`
	BumpSettings               *BattleNetResponseItemsItemBumpSettings               `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                                                 `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                                                 `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                                                 `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                                                 `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                                                 `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                                                 `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                                                 `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                                                 `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                                                 `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                                                 `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                                                 `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                                                 `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                                                 `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                                                 `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                                                 `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                                                 `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                                  `json:"category_id,omitempty"`
	Description                *string                                               `json:"description,omitempty"`
	DescriptionEnHtml          *string                                               `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                                               `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                                               `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                                               `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                                               `json:"description_en,omitempty"`
	DisplayConvertedBalance    *bool                                                 `json:"displayConvertedBalance,omitempty"`
	EditDate                   *int                                                  `json:"edit_date,omitempty"`
	EmailLoginUrl              *string                                               `json:"emailLoginUrl,omitempty"`
	EmailProvider              *string                                               `json:"email_provider,omitempty"`
	EmailType                  *string                                               `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                                                  `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                                           `json:"feedback_data,omitempty"`
	Guarantee                  interface{}                                           `json:"guarantee,omitempty"`
	HasOverwatch               *bool                                                 `json:"hasOverwatch,omitempty"`
	HasPendingAutoBuy          *bool                                                 `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                  *bool                                                 `json:"isIgnored,omitempty"`
	IsSmallExf                 *bool                                                 `json:"isSmallExf,omitempty"`
	IsSticky                   *int                                                  `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                                               `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                                               `json:"item_domain,omitempty"`
	ItemID                     *int                                                  `json:"item_id,omitempty"`
	ItemOrigin                 *string                                               `json:"item_origin,omitempty"`
	ItemState                  *string                                               `json:"item_state,omitempty"`
	NoteText                   interface{}                                           `json:"note_text,omitempty"`
	Nsb                        *int                                                  `json:"nsb,omitempty"`
	Price                      *int                                                  `json:"price,omitempty"`
	PriceWithSellerFee         *float64                                              `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                                               `json:"price_currency,omitempty"`
	PublishedDate              *int                                                  `json:"published_date,omitempty"`
	RefreshedDate              *int                                                  `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                                               `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount  *int                                                  `json:"restore_items_category_count,omitempty"`
	RubPrice                   *int                                                  `json:"rub_price,omitempty"`
	Seller                     *BattleNetResponseItemsItemSeller                     `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                                                 `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount     *int                                                  `json:"sold_items_category_count,omitempty"`
	Tags                       interface{}                                           `json:"tags,omitempty"`
	Title                      *string                                               `json:"title,omitempty"`
	TitleEn                    *string                                               `json:"title_en,omitempty"`
	UpdateStatDate             *int                                                  `json:"update_stat_date,omitempty"`
	ViewCount                  *int                                                  `json:"view_count,omitempty"`
}

type BattleNetResponseItemsItemBattlenetTransactionsItem struct {
	Date           *int    `json:"date,omitempty"`
	FormattedTotal *string `json:"formattedTotal,omitempty"`
	ProductTitle   *string `json:"productTitle,omitempty"`
	Total          *string `json:"total,omitempty"`
}

type BattleNetResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type BattleNetResponseItemsItemGuarantee struct {
	Active         interface{} `json:"active,omitempty"`
	Cancelled      interface{} `json:"cancelled,omitempty"`
	Class          *string     `json:"class,omitempty"`
	Duration       *int        `json:"duration,omitempty"`
	DurationPhrase *string     `json:"durationPhrase,omitempty"`
	EndDate        interface{} `json:"endDate,omitempty"`
	RemainingTime  interface{} `json:"remainingTime,omitempty"`
}

type BattleNetResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type BattleNetResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type BulkGetParams struct {
	// Item id.
	ItemID []int `json:"item_id,omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
}

type BulkGetResponse struct {
	Items      interface{}     `json:"items,omitempty"`
	LeftItemID []int           `json:"left_item_id,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type BulkGetResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type BuyItem struct {
	Item       *BuyItemItem    `json:"item,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type BuyItemItem struct {
	AccountLink                       *string                       `json:"accountLink,omitempty"`
	AccountLinks                      []BuyItemItemAccountLinksItem `json:"accountLinks,omitempty"`
	AccountLastActivity               *int                          `json:"account_last_activity,omitempty"`
	AllowAskDiscount                  *int                          `json:"allow_ask_discount,omitempty"`
	BumpSettings                      *BuyItemItemBumpSettings      `json:"bumpSettings,omitempty"`
	Buyer                             *BuyItemItemBuyer             `json:"buyer,omitempty"`
	BuyerAvatarDate                   *int                          `json:"buyer_avatar_date,omitempty"`
	BuyerUserGroupID                  *int                          `json:"buyer_user_group_id,omitempty"`
	CanAskDiscount                    *bool                         `json:"canAskDiscount,omitempty"`
	CanChangePassword                 *bool                         `json:"canChangePassword,omitempty"`
	CanCheckGuarantee                 *bool                         `json:"canCheckGuarantee,omitempty"`
	CanReportItem                     *bool                         `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase        *bool                         `json:"canResellItemAfterPurchase,omitempty"`
	CanUpdateItemStats                *bool                         `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount                *bool                         `json:"canValidateAccount,omitempty"`
	CanViewAccountLink                *bool                         `json:"canViewAccountLink,omitempty"`
	CanViewAccountLoginAndTempEmail   *bool                         `json:"canViewAccountLoginAndTempEmail,omitempty"`
	CanViewEmailLoginData             *bool                         `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData                  *bool                         `json:"canViewLoginData,omitempty"`
	CategoryID                        *int                          `json:"category_id,omitempty"`
	CustomFields                      []interface{}                 `json:"customFields,omitempty"`
	Deposit                           *int                          `json:"deposit,omitempty"`
	Description                       *string                       `json:"description,omitempty"`
	DescriptionEnHtml                 *string                       `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain                *string                       `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml                   *string                       `json:"descriptionHtml,omitempty"`
	DescriptionPlain                  *string                       `json:"descriptionPlain,omitempty"`
	DescriptionEn                     *string                       `json:"description_en,omitempty"`
	DisplayConvertedBalance           *bool                         `json:"displayConvertedBalance,omitempty"`
	EditDate                          *int                          `json:"edit_date,omitempty"`
	EmailLoginData                    *BuyItemItemEmailLoginData    `json:"emailLoginData,omitempty"`
	EmailLoginUrl                     *string                       `json:"emailLoginUrl,omitempty"`
	EmailProvider                     *string                       `json:"email_provider,omitempty"`
	EmailType                         *string                       `json:"email_type,omitempty"`
	ExtendedGuarantee                 *int                          `json:"extended_guarantee,omitempty"`
	ExternalAuth                      []interface{}                 `json:"externalAuth,omitempty"`
	ExtraPrices                       []BuyItemItemExtraPricesItem  `json:"extraPrices,omitempty"`
	FeedbackData                      interface{}                   `json:"feedback_data,omitempty"`
	GetEmailCodeDisplayLogin          *string                       `json:"getEmailCodeDisplayLogin,omitempty"`
	Guarantee                         interface{}                   `json:"guarantee,omitempty"`
	Information                       *string                       `json:"information,omitempty"`
	InformationEn                     *string                       `json:"information_en,omitempty"`
	IsBirthdayToday                   *bool                         `json:"isBirthdayToday,omitempty"`
	IsIgnored                         *bool                         `json:"isIgnored,omitempty"`
	IsPersonalAccount                 *bool                         `json:"isPersonalAccount,omitempty"`
	IsTrusted                         *bool                         `json:"isTrusted,omitempty"`
	IsSticky                          *int                          `json:"is_sticky,omitempty"`
	ItemOriginPhrase                  *string                       `json:"itemOriginPhrase,omitempty"`
	ItemDomain                        *string                       `json:"item_domain,omitempty"`
	ItemID                            *int                          `json:"item_id,omitempty"`
	ItemOrigin                        *string                       `json:"item_origin,omitempty"`
	ItemState                         *string                       `json:"item_state,omitempty"`
	Login                             *string                       `json:"login,omitempty"`
	LoginData                         *BuyItemItemLoginData         `json:"loginData,omitempty"`
	MarketCustomTitle                 *string                       `json:"market_custom_title,omitempty"`
	MaxDiscountPercent                *int                          `json:"max_discount_percent,omitempty"`
	NeedToRequireVideoToViewLoginData *bool                         `json:"needToRequireVideoToViewLoginData,omitempty"`
	Nsb                               *int                          `json:"nsb,omitempty"`
	Price                             *int                          `json:"price,omitempty"`
	PriceWithSellerFee                *float64                      `json:"priceWithSellerFee,omitempty"`
	PriceCurrency                     *string                       `json:"price_currency,omitempty"`
	PublishedDate                     *int                          `json:"published_date,omitempty"`
	RefreshedDate                     *int                          `json:"refreshed_date,omitempty"`
	ResaleItemOrigin                  *string                       `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount         *int                          `json:"restore_items_category_count,omitempty"`
	RubPrice                          *int                          `json:"rub_price,omitempty"`
	Seller                            *BuyItemItemSeller            `json:"seller,omitempty"`
	ShowGetEmailCodeButton            *bool                         `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount            *int                          `json:"sold_items_category_count,omitempty"`
	Tags                              interface{}                   `json:"tags,omitempty"`
	TempEmail                         *string                       `json:"temp_email,omitempty"`
	Title                             *string                       `json:"title,omitempty"`
	TitleEn                           *string                       `json:"title_en,omitempty"`
	UpdateStatDate                    *int                          `json:"update_stat_date,omitempty"`
	UserAllowAskDiscount              *int                          `json:"user_allow_ask_discount,omitempty"`
	ViewCount                         *int                          `json:"view_count,omitempty"`
	VisitorIsAuthor                   *bool                         `json:"visitorIsAuthor,omitempty"`
}

type BuyItemItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type BuyItemItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type BuyItemItemBuyer struct {
	DisplayStyleGroupID *int    `json:"display_style_group_id,omitempty"`
	IsBanned            *int    `json:"is_banned,omitempty"`
	OperationDate       *int    `json:"operation_date,omitempty"`
	UniqUsernameCSS     *string `json:"uniq_username_css,omitempty"`
	UserGroupID         *int    `json:"user_group_id,omitempty"`
	UserID              *int    `json:"user_id,omitempty"`
	Username            *string `json:"username,omitempty"`
	VisitorIsBuyer      *bool   `json:"visitorIsBuyer,omitempty"`
}

type BuyItemItemEmailLoginData struct {
	EncodedOldPassword *string `json:"encodedOldPassword,omitempty"`
	EncodedPassword    *string `json:"encodedPassword,omitempty"`
	EncodedRaw         *string `json:"encodedRaw,omitempty"`
	Login              *string `json:"login,omitempty"`
	OldPassword        *string `json:"oldPassword,omitempty"`
	Password           *string `json:"password,omitempty"`
	Raw                *string `json:"raw,omitempty"`
}

type BuyItemItemExtraPricesItem struct {
	Currency *string `json:"currency,omitempty"`
	Price    *string `json:"price,omitempty"`
}

type BuyItemItemGuarantee struct {
	Active              *bool   `json:"active,omitempty"`
	Cancelled           *bool   `json:"cancelled,omitempty"`
	Class               *string `json:"class,omitempty"`
	Duration            *int    `json:"duration,omitempty"`
	DurationPhrase      *string `json:"durationPhrase,omitempty"`
	EndDate             *int    `json:"endDate,omitempty"`
	RemainingTime       *int    `json:"remainingTime,omitempty"`
	RemainingTimePhrase *string `json:"remainingTimePhrase,omitempty"`
}

type BuyItemItemLoginData struct {
	EncodedOldPassword interface{} `json:"encodedOldPassword,omitempty"`
	EncodedPassword    *string     `json:"encodedPassword,omitempty"`
	EncodedRaw         *string     `json:"encodedRaw,omitempty"`
	Login              *string     `json:"login,omitempty"`
	OldPassword        *string     `json:"oldPassword,omitempty"`
	Password           *string     `json:"password,omitempty"`
	Raw                *string     `json:"raw,omitempty"`
}

type BuyItemItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsOnline            *bool       `json:"isOnline,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	JoinedDate          *int        `json:"joined_date,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type BuyItemSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ChangePasswordParams struct {
	// Cancel change password recommendation. It will be helpful, if you don't want to change password and get login data.
	Cancel *AccountsManagingChangePasswordCancel `json:"_cancel,omitempty"`
}

type ChangePasswordResponse struct {
	Message     *string `json:"message,omitempty"`
	NewPassword *string `json:"new_password,omitempty"`
	Status      *string `json:"status,omitempty"`
}

type ChatGPTParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// List of allowed subscriptions.
	Subscription []string `json:"subscription[],omitempty"`
	// Length of subscription.
	SubscriptionLength *int `json:"subscription_length,omitempty"`
	// In what notation is time measured.
	SubscriptionPeriod *CategorySearchSubscriptionPeriod `json:"subscription_period,omitempty"`
	// Is auto renewal enabled.
	Autorenewal *CategorySearchAutorenewal `json:"autorenewal,omitempty"`
	// Has linked mobile.
	Tel *CategorySearchTel `json:"tel,omitempty"`
	// Has transactions.
	Transactions *CategorySearchTransactions `json:"transactions,omitempty"`
	// How old is the account.
	Reg *int `json:"reg,omitempty"`
	// In what notation is time measured.
	RegPeriod *CategorySearchRegPeriod `json:"reg_period,omitempty"`
	// List of allowed tiers.
	OpenaiTier []string `json:"openai_tier[],omitempty"`
	// Minimum OpenAI credit balance.
	OpenaiBalanceMin *int `json:"openai_balance_min,omitempty"`
	// Maximum OpenAI credit balance.
	OpenaiBalanceMax *int `json:"openai_balance_max,omitempty"`
}

type ChatGPTResponse struct {
	CacheTTL        *int                       `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                      `json:"hasNextPage,omitempty"`
	Items           []ChatGPTResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                       `json:"lastModified,omitempty"`
	Page            *int                       `json:"page,omitempty"`
	PerPage         *int                       `json:"perPage,omitempty"`
	SearchUrl       *string                    `json:"searchUrl,omitempty"`
	ServerTime      *int                       `json:"serverTime,omitempty"`
	StickyItems     []interface{}              `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo            `json:"system_info,omitempty"`
	TotalItems      *int                       `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}                `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                      `json:"wasCached,omitempty"`
}

type ChatGPTResponseItemsItem struct {
	AllowAskDiscount             *int                                    `json:"allow_ask_discount,omitempty"`
	BumpSettings                 *ChatGPTResponseItemsItemBumpSettings   `json:"bumpSettings,omitempty"`
	CanBumpItem                  *bool                                   `json:"canBumpItem,omitempty"`
	CanBuyItem                   *bool                                   `json:"canBuyItem,omitempty"`
	CanChangeEmailPassword       *bool                                   `json:"canChangeEmailPassword,omitempty"`
	CanChangePassword            *bool                                   `json:"canChangePassword,omitempty"`
	CanCloseItem                 *bool                                   `json:"canCloseItem,omitempty"`
	CanDeleteItem                *bool                                   `json:"canDeleteItem,omitempty"`
	CanEditItem                  *bool                                   `json:"canEditItem,omitempty"`
	CanOpenItem                  *bool                                   `json:"canOpenItem,omitempty"`
	CanReportItem                *bool                                   `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase   *bool                                   `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem                 *bool                                   `json:"canStickItem,omitempty"`
	CanUnstickItem               *bool                                   `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats           *bool                                   `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount           *bool                                   `json:"canValidateAccount,omitempty"`
	CanViewAccountLink           *bool                                   `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData        *bool                                   `json:"canViewEmailLoginData,omitempty"`
	CanViewItemViews             *bool                                   `json:"canViewItemViews,omitempty"`
	CanViewLoginData             *bool                                   `json:"canViewLoginData,omitempty"`
	CategoryID                   *int                                    `json:"category_id,omitempty"`
	ChatgptCountry               *string                                 `json:"chatgpt_country,omitempty"`
	ChatgptItemID                *int                                    `json:"chatgpt_item_id,omitempty"`
	ChatgptPhone                 *int                                    `json:"chatgpt_phone,omitempty"`
	ChatgptRegisterDate          *int                                    `json:"chatgpt_register_date,omitempty"`
	ChatgptSubscription          *string                                 `json:"chatgpt_subscription,omitempty"`
	ChatgptSubscriptionAutoRenew *int                                    `json:"chatgpt_subscription_auto_renew,omitempty"`
	ChatgptSubscriptionEnds      *int                                    `json:"chatgpt_subscription_ends,omitempty"`
	CopyFormatData               *ChatGPTResponseItemsItemCopyFormatData `json:"copyFormatData,omitempty"`
	Description                  *string                                 `json:"description,omitempty"`
	DescriptionEnHtml            *string                                 `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain           *string                                 `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml              *string                                 `json:"descriptionHtml,omitempty"`
	DescriptionPlain             *string                                 `json:"descriptionPlain,omitempty"`
	DescriptionEn                *string                                 `json:"description_en,omitempty"`
	EditDate                     *int                                    `json:"edit_date,omitempty"`
	EmailLoginUrl                *string                                 `json:"emailLoginUrl,omitempty"`
	EmailProvider                *string                                 `json:"email_provider,omitempty"`
	EmailType                    *string                                 `json:"email_type,omitempty"`
	ExtendedGuarantee            *int                                    `json:"extended_guarantee,omitempty"`
	FeedbackData                 interface{}                             `json:"feedback_data,omitempty"`
	GptSubType                   *string                                 `json:"gptSubType,omitempty"`
	Guarantee                    interface{}                             `json:"guarantee,omitempty"`
	HasPendingAutoBuy            *bool                                   `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                    *bool                                   `json:"isIgnored,omitempty"`
	IsPersonalAccount            *bool                                   `json:"isPersonalAccount,omitempty"`
	IsSticky                     *int                                    `json:"is_sticky,omitempty"`
	ItemOriginPhrase             *string                                 `json:"itemOriginPhrase,omitempty"`
	ItemDomain                   *string                                 `json:"item_domain,omitempty"`
	ItemID                       *int                                    `json:"item_id,omitempty"`
	ItemOrigin                   *string                                 `json:"item_origin,omitempty"`
	ItemState                    *string                                 `json:"item_state,omitempty"`
	NoteText                     interface{}                             `json:"note_text,omitempty"`
	Nsb                          *int                                    `json:"nsb,omitempty"`
	Price                        *int                                    `json:"price,omitempty"`
	PriceWithSellerFee           *float64                                `json:"priceWithSellerFee,omitempty"`
	PriceWithSellerFeeLabel      *string                                 `json:"priceWithSellerFeeLabel,omitempty"`
	PriceCurrency                *string                                 `json:"price_currency,omitempty"`
	PublishedDate                *int                                    `json:"published_date,omitempty"`
	RefreshedDate                *int                                    `json:"refreshed_date,omitempty"`
	ResaleItemOrigin             *string                                 `json:"resale_item_origin,omitempty"`
	RubPrice                     *int                                    `json:"rub_price,omitempty"`
	Seller                       *ChatGPTResponseItemsItemSeller         `json:"seller,omitempty"`
	ShowGetEmailCodeButton       *bool                                   `json:"showGetEmailCodeButton,omitempty"`
	Tags                         interface{}                             `json:"tags,omitempty"`
	Title                        *string                                 `json:"title,omitempty"`
	TitleEn                      *string                                 `json:"title_en,omitempty"`
	UniqueKeyExists              *bool                                   `json:"uniqueKeyExists,omitempty"`
	UpdateStatDate               *int                                    `json:"update_stat_date,omitempty"`
	ViewCount                    *int                                    `json:"view_count,omitempty"`
}

type ChatGPTResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type ChatGPTResponseItemsItemCopyFormatData struct {
	TitleLink *string `json:"title_link,omitempty"`
}

type ChatGPTResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     interface{} `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type ChatGPTResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type CheckAccount struct {
	Item                  *CheckAccountItem `json:"item,omitempty"`
	RequireVideoRecording *bool             `json:"requireVideoRecording,omitempty"`
	Status                *string           `json:"status,omitempty"`
	SystemInfo            *RespSystemInfo   `json:"system_info,omitempty"`
}

type CheckAccountItem struct {
	AccountLink                       *string                            `json:"accountLink,omitempty"`
	AccountLinks                      []CheckAccountItemAccountLinksItem `json:"accountLinks,omitempty"`
	AccountLastActivity               *int                               `json:"account_last_activity,omitempty"`
	AllowAskDiscount                  *int                               `json:"allow_ask_discount,omitempty"`
	AskDate                           interface{}                        `json:"ask_date,omitempty"`
	AskItemID                         interface{}                        `json:"ask_item_id,omitempty"`
	AskUserID                         interface{}                        `json:"ask_user_id,omitempty"`
	AvailableTempEmail                *int                               `json:"available_temp_email,omitempty"`
	BumpSettings                      *CheckAccountItemBumpSettings      `json:"bumpSettings,omitempty"`
	BuyWithoutValidation              *int                               `json:"buy_without_validation,omitempty"`
	CanAskDiscount                    *bool                              `json:"canAskDiscount,omitempty"`
	CanChangePassword                 *bool                              `json:"canChangePassword,omitempty"`
	CanCheckGuarantee                 *bool                              `json:"canCheckGuarantee,omitempty"`
	CanReportItem                     *bool                              `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase        *bool                              `json:"canResellItemAfterPurchase,omitempty"`
	CanUpdateItemStats                *bool                              `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount                *bool                              `json:"canValidateAccount,omitempty"`
	CanViewAccountLink                *bool                              `json:"canViewAccountLink,omitempty"`
	CanViewAccountLoginAndTempEmail   *bool                              `json:"canViewAccountLoginAndTempEmail,omitempty"`
	CanViewEmailLoginData             *bool                              `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData                  *bool                              `json:"canViewLoginData,omitempty"`
	CanBeResold                       *int                               `json:"can_be_resold,omitempty"`
	CategoryID                        *int                               `json:"category_id,omitempty"`
	CategoryPrefixID                  *int                               `json:"category_prefix_id,omitempty"`
	CategoryTitle                     *string                            `json:"category_title,omitempty"`
	CategoryURL                       *string                            `json:"category_url,omitempty"`
	CheckButtonEnabled                *int                               `json:"check_button_enabled,omitempty"`
	CheckerEnabled                    *int                               `json:"checker_enabled,omitempty"`
	CustomFields                      []interface{}                      `json:"customFields,omitempty"`
	Deposit                           *int                               `json:"deposit,omitempty"`
	Description                       *string                            `json:"description,omitempty"`
	DescriptionEnHtml                 *string                            `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain                *string                            `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml                   *string                            `json:"descriptionHtml,omitempty"`
	DescriptionPlain                  *string                            `json:"descriptionPlain,omitempty"`
	DescriptionEn                     *string                            `json:"description_en,omitempty"`
	DiscountAccepted                  interface{}                        `json:"discount_accepted,omitempty"`
	DiscountPrice                     interface{}                        `json:"discount_price,omitempty"`
	EditDate                          *int                               `json:"edit_date,omitempty"`
	EmailLoginUrl                     *string                            `json:"emailLoginUrl,omitempty"`
	EmailProvider                     *string                            `json:"email_provider,omitempty"`
	EmailType                         *string                            `json:"email_type,omitempty"`
	ExtendedGuarantee                 *int                               `json:"extended_guarantee,omitempty"`
	ExternalAuth                      []interface{}                      `json:"externalAuth,omitempty"`
	ExtraPrices                       []CheckAccountItemExtraPricesItem  `json:"extraPrices,omitempty"`
	FeedbackData                      interface{}                        `json:"feedback_data,omitempty"`
	Guarantee                         interface{}                        `json:"guarantee,omitempty"`
	HasGuarantee                      *int                               `json:"has_guarantee,omitempty"`
	IsBirthdayToday                   *bool                              `json:"isBirthdayToday,omitempty"`
	IsIgnored                         *bool                              `json:"isIgnored,omitempty"`
	IsPersonalAccount                 *bool                              `json:"isPersonalAccount,omitempty"`
	IsSmallExf                        *bool                              `json:"isSmallExf,omitempty"`
	IsTrusted                         *bool                              `json:"isTrusted,omitempty"`
	IsSticky                          *int                               `json:"is_sticky,omitempty"`
	ItemOriginPhrase                  *string                            `json:"itemOriginPhrase,omitempty"`
	ItemDomain                        *string                            `json:"item_domain,omitempty"`
	ItemID                            *int                               `json:"item_id,omitempty"`
	ItemOrigin                        *string                            `json:"item_origin,omitempty"`
	ItemState                         *string                            `json:"item_state,omitempty"`
	LoginType                         *string                            `json:"login_type,omitempty"`
	MarketCustomTitle                 *string                            `json:"market_custom_title,omitempty"`
	MaxDiscountPercent                *int                               `json:"max_discount_percent,omitempty"`
	Message                           interface{}                        `json:"message,omitempty"`
	MinPrice                          *int                               `json:"min_price,omitempty"`
	NeedToRequireVideoToViewLoginData *bool                              `json:"needToRequireVideoToViewLoginData,omitempty"`
	Nsb                               *int                               `json:"nsb,omitempty"`
	Price                             *int                               `json:"price,omitempty"`
	PriceWithSellerFee                *float64                           `json:"priceWithSellerFee,omitempty"`
	PriceCurrency                     *string                            `json:"price_currency,omitempty"`
	PublishedDate                     *int                               `json:"published_date,omitempty"`
	RefreshedDate                     *int                               `json:"refreshed_date,omitempty"`
	RequireEmailLoginData             *int                               `json:"require_email_login_data,omitempty"`
	RequireTempEmail                  *int                               `json:"require_temp_email,omitempty"`
	RequireVideoRecording             *int                               `json:"require_video_recording,omitempty"`
	ResaleItemOrigin                  *string                            `json:"resale_item_origin,omitempty"`
	RubPrice                          *int                               `json:"rub_price,omitempty"`
	Seller                            *CheckAccountItemSeller            `json:"seller,omitempty"`
	ShowGetEmailCodeButton            *bool                              `json:"showGetEmailCodeButton,omitempty"`
	Tags                              interface{}                        `json:"tags,omitempty"`
	Title                             *string                            `json:"title,omitempty"`
	TitleEn                           *string                            `json:"title_en,omitempty"`
	UpdateStatDate                    *int                               `json:"update_stat_date,omitempty"`
	UserAlerted                       interface{}                        `json:"user_alerted,omitempty"`
	UserAllowAskDiscount              *int                               `json:"user_allow_ask_discount,omitempty"`
	ViewCount                         *int                               `json:"view_count,omitempty"`
	VisitorIsAuthor                   *bool                              `json:"visitorIsAuthor,omitempty"`
}

type CheckAccountItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type CheckAccountItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type CheckAccountItemExtraPricesItem struct {
	Currency *string `json:"currency,omitempty"`
	Price    *string `json:"price,omitempty"`
}

type CheckAccountItemGuarantee struct {
	Active         interface{} `json:"active,omitempty"`
	Cancelled      interface{} `json:"cancelled,omitempty"`
	Class          *string     `json:"class,omitempty"`
	Duration       *int        `json:"duration,omitempty"`
	DurationPhrase *string     `json:"durationPhrase,omitempty"`
	EndDate        interface{} `json:"endDate,omitempty"`
	RemainingTime  interface{} `json:"remainingTime,omitempty"`
}

type CheckAccountItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsOnline            *bool       `json:"isOnline,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	JoinedDate          *int        `json:"joined_date,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     interface{} `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type CheckAccountSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type CheckGuaranteeResponse struct {
	Message    *string         `json:"message,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type CheckGuaranteeResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type CheckParams struct {
	// Required if a **category** is one of list of Required email login data categories. Email login data (email:password format).
	EmailLoginData *string `json:"email_login_data,omitempty"`
	// Email type.
	EmailType *AccountPublishingCheckEmailType `json:"email_type,omitempty"`
	Extra     map[string]interface{}           `json:"extra,omitempty"`
	// Required if a **category** is one of list of Required email login data categories.
	HasEmailLoginData *bool `json:"has_email_login_data,omitempty"`
	// Account login (or email).
	Login *string `json:"login,omitempty"`
	// Account login data (login:password format).
	LoginPassword *string `json:"login_password,omitempty"`
	// Account password.
	Password *string `json:"password,omitempty"`
	// Set this parameter to **true** so that the Market will take a random proxy from its pool for each of your requests. Otherwise, if this parameter is set to **false** or not set, the Market will take a ...
	RandomProxy *bool `json:"random_proxy,omitempty"`
	// Put if you are trying to resell an account.
	ResellItemID *int `json:"resell_item_id,omitempty"`
}

type ClaimsParams struct {
	// Filter claims by their type.
	Type_ *AccountsManagingType `json:"type,omitempty"`
	// Filter claims by their state.
	ClaimState *AccountsManagingClaimState `json:"claim_state,omitempty"`
}

type ClaimsResponse struct {
	Claims     []ClaimsResponseClaimsItem `json:"claims,omitempty"`
	Stats      *ClaimsResponseStats       `json:"stats,omitempty"`
	SystemInfo *RespSystemInfo            `json:"system_info,omitempty"`
}

type ClaimsResponseClaimsItem struct {
	AmountFormatted *string                         `json:"amount_formatted,omitempty"`
	Author          *ClaimsResponseClaimsItemAuthor `json:"author,omitempty"`
	ClaimDate       *int                            `json:"claim_date,omitempty"`
	ClaimState      *string                         `json:"claim_state,omitempty"`
	MessageBody     *string                         `json:"message_body,omitempty"`
	ThreadID        *int                            `json:"thread_id,omitempty"`
}

type ClaimsResponseClaimsItemAuthor struct {
	BanReason        *string                                    `json:"ban_reason,omitempty"`
	ContestCount     *int                                       `json:"contest_count,omitempty"`
	CustomTitle      *string                                    `json:"custom_title,omitempty"`
	Fields           []ClaimsResponseClaimsItemAuthorFieldsItem `json:"fields,omitempty"`
	IsBanned         *int                                       `json:"is_banned,omitempty"`
	Links            *ClaimsResponseClaimsItemAuthorLinks       `json:"links,omitempty"`
	Permissions      *ClaimsResponseClaimsItemAuthorPermissions `json:"permissions,omitempty"`
	TrophyCount      *int                                       `json:"trophy_count,omitempty"`
	UserGroupID      *int                                       `json:"user_group_id,omitempty"`
	UserID           *int                                       `json:"user_id,omitempty"`
	UserIsFollowed   *bool                                      `json:"user_is_followed,omitempty"`
	UserIsIgnored    *bool                                      `json:"user_is_ignored,omitempty"`
	UserIsValid      *bool                                      `json:"user_is_valid,omitempty"`
	UserIsVerified   *bool                                      `json:"user_is_verified,omitempty"`
	UserIsVisitor    *bool                                      `json:"user_is_visitor,omitempty"`
	UserLastSeenDate *int                                       `json:"user_last_seen_date,omitempty"`
	UserLike2Count   *int                                       `json:"user_like2_count,omitempty"`
	UserLikeCount    *int                                       `json:"user_like_count,omitempty"`
	UserMessageCount *int                                       `json:"user_message_count,omitempty"`
	UserRegisterDate *int                                       `json:"user_register_date,omitempty"`
	UserTitle        *string                                    `json:"user_title,omitempty"`
	Username         *string                                    `json:"username,omitempty"`
	UsernameHTML     *string                                    `json:"username_html,omitempty"`
}

type ClaimsResponseClaimsItemAuthorFieldsItem struct {
	Description *string `json:"description,omitempty"`
	ID          *string `json:"id,omitempty"`
	IsRequired  *bool   `json:"is_required,omitempty"`
	Position    *string `json:"position,omitempty"`
	Title       *string `json:"title,omitempty"`
}

type ClaimsResponseClaimsItemAuthorLinks struct {
	Avatar      *string `json:"avatar,omitempty"`
	AvatarBig   *string `json:"avatar_big,omitempty"`
	AvatarSmall *string `json:"avatar_small,omitempty"`
	Detail      *string `json:"detail,omitempty"`
	Followers   *string `json:"followers,omitempty"`
	Followings  *string `json:"followings,omitempty"`
	Ignore      *string `json:"ignore,omitempty"`
	Permalink   *string `json:"permalink,omitempty"`
	Timeline    *string `json:"timeline,omitempty"`
}

type ClaimsResponseClaimsItemAuthorPermissions struct {
	Edit   *bool `json:"edit,omitempty"`
	Follow *bool `json:"follow,omitempty"`
	Ignore *bool `json:"ignore,omitempty"`
}

type ClaimsResponseStats struct {
	Market   *ClaimsResponseStatsMarket   `json:"market,omitempty"`
	NoMarket *ClaimsResponseStatsNoMarket `json:"noMarket,omitempty"`
}

type ClaimsResponseStatsMarket struct {
	Rejected *int `json:"rejected,omitempty"`
	Settled  *int `json:"settled,omitempty"`
	Solved   *int `json:"solved,omitempty"`
	Total    *int `json:"total,omitempty"`
}

type ClaimsResponseStatsNoMarket struct {
	Rejected *int `json:"rejected,omitempty"`
	Settled  *int `json:"settled,omitempty"`
	Solved   *int `json:"solved,omitempty"`
	Total    *int `json:"total,omitempty"`
}

type ClaimsResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ConfirmParams struct {
	// Balance ID that will be used to purchase specified item.
	BalanceID *int `json:"balance_id,omitempty"`
	// Current price of account in your currency.
	Price *int `json:"price,omitempty"`
}

type ConfirmResponse struct {
	Item       *ConfirmResponseItem `json:"item,omitempty"`
	Status     *string              `json:"status,omitempty"`
	SystemInfo *RespSystemInfo      `json:"system_info,omitempty"`
}

type ConfirmResponseItem struct {
	LoginData *ConfirmResponseItemLoginData `json:"loginData,omitempty"`
}

type ConfirmResponseItemLoginData struct {
	AdviceToChangePassword *bool   `json:"adviceToChangePassword,omitempty"`
	EncodedOldPassword     *string `json:"encodedOldPassword,omitempty"`
	EncodedPassword        *string `json:"encodedPassword,omitempty"`
	EncodedRaw             *string `json:"encodedRaw,omitempty"`
	Login                  *string `json:"login,omitempty"`
	OldPassword            *string `json:"oldPassword,omitempty"`
	Password               *string `json:"password,omitempty"`
	Raw                    *string `json:"raw,omitempty"`
}

type ConfirmResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ConfirmationCodeModel struct {
	CodeData *ConfirmationCodeModelCodeData `json:"codeData,omitempty"`
	Item     *ItemModel                     `json:"item,omitempty"`
}

type ConfirmationCodeModelCodeData struct {
	Code      *string `json:"code,omitempty"`
	Date      *int    `json:"date,omitempty"`
	TextPlain *string `json:"textPlain,omitempty"`
}

type ConfirmationCodeModelItem struct {
	AccountLink                       *string                                     `json:"accountLink,omitempty"`
	AccountLinks                      []ConfirmationCodeModelItemAccountLinksItem `json:"accountLinks,omitempty"`
	AccountLastActivity               *int                                        `json:"account_last_activity,omitempty"`
	AiPrice                           *int                                        `json:"aiPrice,omitempty"`
	AiPriceCheckDate                  *int                                        `json:"aiPriceCheckDate,omitempty"`
	AllowAskDiscount                  *int                                        `json:"allow_ask_discount,omitempty"`
	AutoBuyPrice                      *int                                        `json:"autoBuyPrice,omitempty"`
	AutoBuyPriceCheckDate             *int                                        `json:"autoBuyPriceCheckDate,omitempty"`
	BumpSettings                      *ConfirmationCodeModelItemBumpSettings      `json:"bumpSettings,omitempty"`
	Buyer                             *ConfirmationCodeModelItemBuyer             `json:"buyer,omitempty"`
	BuyerAvatarDate                   *int                                        `json:"buyer_avatar_date,omitempty"`
	BuyerDisplayIconGroupID           *int                                        `json:"buyer_display_icon_group_id,omitempty"`
	BuyerUniqBanner                   *string                                     `json:"buyer_uniq_banner,omitempty"`
	BuyerUserGroupID                  *int                                        `json:"buyer_user_group_id,omitempty"`
	CanAskDiscount                    *bool                                       `json:"canAskDiscount,omitempty"`
	CanChangeEmailPassword            *bool                                       `json:"canChangeEmailPassword,omitempty"`
	CanChangePassword                 *bool                                       `json:"canChangePassword,omitempty"`
	CanCheckAiPrice                   *bool                                       `json:"canCheckAiPrice,omitempty"`
	CanCheckAutoBuyPrice              *bool                                       `json:"canCheckAutoBuyPrice,omitempty"`
	CanCheckGuarantee                 *bool                                       `json:"canCheckGuarantee,omitempty"`
	CanReportItem                     *bool                                       `json:"canReportItem,omitempty"`
	CanResellItem                     *bool                                       `json:"canResellItem,omitempty"`
	CanResellItemAfterPurchase        *bool                                       `json:"canResellItemAfterPurchase,omitempty"`
	CanShareItem                      *bool                                       `json:"canShareItem,omitempty"`
	CanUpdateItemStats                *bool                                       `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount                *bool                                       `json:"canValidateAccount,omitempty"`
	CanViewAccountLink                *bool                                       `json:"canViewAccountLink,omitempty"`
	CanViewAccountLoginAndTempEmail   *bool                                       `json:"canViewAccountLoginAndTempEmail,omitempty"`
	CanViewEmailLoginData             *bool                                       `json:"canViewEmailLoginData,omitempty"`
	CanViewItemViews                  *bool                                       `json:"canViewItemViews,omitempty"`
	CanViewLoginData                  *bool                                       `json:"canViewLoginData,omitempty"`
	CartPrice                         interface{}                                 `json:"cart_price,omitempty"`
	CategoryID                        *int                                        `json:"category_id,omitempty"`
	ContentID                         interface{}                                 `json:"content_id,omitempty"`
	ContentType                       interface{}                                 `json:"content_type,omitempty"`
	CopyFormatData                    *ConfirmationCodeModelItemCopyFormatData    `json:"copyFormatData,omitempty"`
	CustomFields                      *ConfirmationCodeModelItemCustomFields      `json:"customFields,omitempty"`
	DeleteDate                        *int                                        `json:"delete_date,omitempty"`
	DeleteReason                      *string                                     `json:"delete_reason,omitempty"`
	DeleteUserID                      *int                                        `json:"delete_user_id,omitempty"`
	DeleteUsername                    *string                                     `json:"delete_username,omitempty"`
	Deposit                           *int                                        `json:"deposit,omitempty"`
	Description                       *string                                     `json:"description,omitempty"`
	DescriptionEnHtml                 *string                                     `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain                *string                                     `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml                   *string                                     `json:"descriptionHtml,omitempty"`
	DescriptionPlain                  *string                                     `json:"descriptionPlain,omitempty"`
	DescriptionEn                     *string                                     `json:"description_en,omitempty"`
	EditDate                          *int                                        `json:"edit_date,omitempty"`
	EmailProvider                     *string                                     `json:"email_provider,omitempty"`
	EmailType                         *string                                     `json:"email_type,omitempty"`
	ExtendedGuarantee                 *int                                        `json:"extended_guarantee,omitempty"`
	ExternalAuth                      []interface{}                               `json:"externalAuth,omitempty"`
	ExtraPrices                       []ConfirmationCodeModelItemExtraPricesItem  `json:"extraPrices,omitempty"`
	FeedbackData                      interface{}                                 `json:"feedback_data,omitempty"`
	GetEmailCodeDisplayLogin          interface{}                                 `json:"getEmailCodeDisplayLogin,omitempty"`
	Guarantee                         interface{}                                 `json:"guarantee,omitempty"`
	ImagePreviewLinks                 []string                                    `json:"imagePreviewLinks,omitempty"`
	InCart                            interface{}                                 `json:"in_cart,omitempty"`
	Information                       *string                                     `json:"information,omitempty"`
	InformationEn                     *string                                     `json:"information_en,omitempty"`
	IsBirthdayToday                   *bool                                       `json:"isBirthdayToday,omitempty"`
	IsIgnored                         *bool                                       `json:"isIgnored,omitempty"`
	IsPersonalAccount                 *bool                                       `json:"isPersonalAccount,omitempty"`
	IsSmallExf                        *bool                                       `json:"isSmallExf,omitempty"`
	IsTrusted                         *bool                                       `json:"isTrusted,omitempty"`
	IsFave                            interface{}                                 `json:"is_fave,omitempty"`
	IsSticky                          *int                                        `json:"is_sticky,omitempty"`
	ItemOriginPhrase                  *string                                     `json:"itemOriginPhrase,omitempty"`
	ItemDomain                        *string                                     `json:"item_domain,omitempty"`
	ItemID                            *int                                        `json:"item_id,omitempty"`
	ItemOrigin                        *string                                     `json:"item_origin,omitempty"`
	ItemState                         *string                                     `json:"item_state,omitempty"`
	Login                             *string                                     `json:"login,omitempty"`
	LoginData                         *ConfirmationCodeModelItemLoginData         `json:"loginData,omitempty"`
	MarketCustomTitle                 *string                                     `json:"market_custom_title,omitempty"`
	MaxDiscountPercent                *int                                        `json:"max_discount_percent,omitempty"`
	NeedToRequireVideoToViewLoginData *bool                                       `json:"needToRequireVideoToViewLoginData,omitempty"`
	NoteText                          *string                                     `json:"note_text,omitempty"`
	Nsb                               *int                                        `json:"nsb,omitempty"`
	PendingDeletionDate               *int                                        `json:"pending_deletion_date,omitempty"`
	Price                             *int                                        `json:"price,omitempty"`
	PriceWithSellerFee                *float64                                    `json:"priceWithSellerFee,omitempty"`
	PriceWithSellerFeeLabel           *string                                     `json:"priceWithSellerFeeLabel,omitempty"`
	PriceCurrency                     *string                                     `json:"price_currency,omitempty"`
	PublishedDate                     *int                                        `json:"published_date,omitempty"`
	RefreshedDate                     *int                                        `json:"refreshed_date,omitempty"`
	ResaleItemOrigin                  *string                                     `json:"resale_item_origin,omitempty"`
	RubPrice                          *int                                        `json:"rub_price,omitempty"`
	Seller                            *ConfirmationCodeModelItemSeller            `json:"seller,omitempty"`
	ShowGetEmailCodeButton            *bool                                       `json:"showGetEmailCodeButton,omitempty"`
	Tags                              interface{}                                 `json:"tags,omitempty"`
	TempEmail                         *string                                     `json:"temp_email,omitempty"`
	Title                             *string                                     `json:"title,omitempty"`
	TitleEn                           *string                                     `json:"title_en,omitempty"`
	UniqueKeyExists                   *bool                                       `json:"uniqueKeyExists,omitempty"`
	UpdateStatDate                    *int                                        `json:"update_stat_date,omitempty"`
	UserAllowAskDiscount              *int                                        `json:"user_allow_ask_discount,omitempty"`
	ViewCount                         *int                                        `json:"view_count,omitempty"`
	VisitorIsAuthor                   *bool                                       `json:"visitorIsAuthor,omitempty"`
}

type ConfirmationCodeModelItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type ConfirmationCodeModelItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	NextAllowedBumpDate interface{} `json:"nextAllowedBumpDate,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type ConfirmationCodeModelItemBuyer struct {
	DisplayIconGroupID  *int    `json:"display_icon_group_id,omitempty"`
	DisplayStyleGroupID *int    `json:"display_style_group_id,omitempty"`
	IsBanned            *int    `json:"is_banned,omitempty"`
	OperationDate       *int    `json:"operation_date,omitempty"`
	UniqBanner          *string `json:"uniq_banner,omitempty"`
	UniqUsernameCSS     *string `json:"uniq_username_css,omitempty"`
	UserGroupID         *int    `json:"user_group_id,omitempty"`
	UserID              *int    `json:"user_id,omitempty"`
	Username            *string `json:"username,omitempty"`
	VisitorIsBuyer      *bool   `json:"visitorIsBuyer,omitempty"`
}

type ConfirmationCodeModelItemCopyFormatData struct {
	Full      *string `json:"full,omitempty"`
	LoginData *string `json:"login_data,omitempty"`
	TitleLink *string `json:"title_link,omitempty"`
}

type ConfirmationCodeModelItemCustomFields struct {
	Field4         *string       `json:"_4,omitempty"`
	AllowSelfUnban []interface{} `json:"allowSelfUnban,omitempty"`
	BanReason      *string       `json:"ban_reason,omitempty"`
	Discord        *string       `json:"discord,omitempty"`
	Github         *string       `json:"github,omitempty"`
	Jabber         *string       `json:"jabber,omitempty"`
	LztUnbanAmount *string       `json:"lztUnbanAmount,omitempty"`
	Steam          *string       `json:"steam,omitempty"`
	Telegram       *string       `json:"telegram,omitempty"`
	Vk             *string       `json:"vk,omitempty"`
}

type ConfirmationCodeModelItemExtraPricesItem struct {
	Currency   *string  `json:"currency,omitempty"`
	Price      *string  `json:"price,omitempty"`
	PriceValue *float64 `json:"priceValue,omitempty"`
}

type ConfirmationCodeModelItemGuarantee struct {
	Active                *bool   `json:"active,omitempty"`
	Cancelled             *bool   `json:"cancelled,omitempty"`
	CancelledReason       *string `json:"cancelledReason,omitempty"`
	CancelledReasonPhrase *string `json:"cancelledReasonPhrase,omitempty"`
	Class                 *string `json:"class,omitempty"`
	Duration              *int    `json:"duration,omitempty"`
	DurationPhrase        *string `json:"durationPhrase,omitempty"`
	EndDate               *int    `json:"endDate,omitempty"`
	RemainingTime         *int    `json:"remainingTime,omitempty"`
	RemainingTimePhrase   *string `json:"remainingTimePhrase,omitempty"`
}

type ConfirmationCodeModelItemLoginData struct {
	EncodedOldPassword interface{} `json:"encodedOldPassword,omitempty"`
	EncodedPassword    *string     `json:"encodedPassword,omitempty"`
	EncodedRaw         *string     `json:"encodedRaw,omitempty"`
	Login              *string     `json:"login,omitempty"`
	OldPassword        *string     `json:"oldPassword,omitempty"`
	Password           *string     `json:"password,omitempty"`
	Raw                *string     `json:"raw,omitempty"`
}

type ConfirmationCodeModelItemSeller struct {
	ActiveItemsCount      *int                                     `json:"active_items_count,omitempty"`
	AvatarDate            *int                                     `json:"avatar_date,omitempty"`
	Contacts              *ConfirmationCodeModelItemSellerContacts `json:"contacts,omitempty"`
	DisplayStyleGroupID   *int                                     `json:"display_style_group_id,omitempty"`
	EffectiveLastActivity *int                                     `json:"effective_last_activity,omitempty"`
	IsOnline              *bool                                    `json:"isOnline,omitempty"`
	IsBanned              *int                                     `json:"is_banned,omitempty"`
	JoinedDate            *int                                     `json:"joined_date,omitempty"`
	RestoreData           interface{}                              `json:"restore_data,omitempty"`
	RestorePercents       interface{}                              `json:"restore_percents,omitempty"`
	SoldItemsCount        *int                                     `json:"sold_items_count,omitempty"`
	UserID                *int                                     `json:"user_id,omitempty"`
	Username              *string                                  `json:"username,omitempty"`
}

type ConfirmationCodeModelItemSellerContacts struct {
	BanReason *string `json:"ban_reason,omitempty"`
	Telegram  *string `json:"telegram,omitempty"`
}

type CreateClaimResponse struct {
	SystemInfo *CreateClaimResponseSystemInfo `json:"system_info,omitempty"`
	Thread     *CreateClaimResponseThread     `json:"thread,omitempty"`
}

type CreateClaimResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type CreateClaimResponseThread struct {
	CreatorUserID     *int                                  `json:"creator_user_id,omitempty"`
	CreatorUsername   *string                               `json:"creator_username,omitempty"`
	FirstPost         *CreateClaimResponseThreadFirstPost   `json:"first_post,omitempty"`
	Forum             *CreateClaimResponseThreadForum       `json:"forum,omitempty"`
	ForumID           *int                                  `json:"forum_id,omitempty"`
	Links             *CreateClaimResponseThreadLinks       `json:"links,omitempty"`
	Permissions       *CreateClaimResponseThreadPermissions `json:"permissions,omitempty"`
	ThreadCreateDate  *int                                  `json:"thread_create_date,omitempty"`
	ThreadID          *int                                  `json:"thread_id,omitempty"`
	ThreadIsDeleted   *bool                                 `json:"thread_is_deleted,omitempty"`
	ThreadIsFollowed  *bool                                 `json:"thread_is_followed,omitempty"`
	ThreadIsPublished *bool                                 `json:"thread_is_published,omitempty"`
	ThreadIsSticky    *bool                                 `json:"thread_is_sticky,omitempty"`
	ThreadPostCount   *int                                  `json:"thread_post_count,omitempty"`
	ThreadPrefixes    []interface{}                         `json:"thread_prefixes,omitempty"`
	ThreadTags        interface{}                           `json:"thread_tags,omitempty"`
	ThreadTitle       *string                               `json:"thread_title,omitempty"`
	ThreadUpdateDate  *int                                  `json:"thread_update_date,omitempty"`
	ThreadViewCount   *int                                  `json:"thread_view_count,omitempty"`
	UserIsIgnored     *bool                                 `json:"user_is_ignored,omitempty"`
}

type CreateClaimResponseThreadFirstPost struct {
	LikeUsers           []CreateClaimResponseThreadFirstPostLikeUsersItem `json:"like_users,omitempty"`
	Links               *CreateClaimResponseThreadFirstPostLinks          `json:"links,omitempty"`
	Permissions         *CreateClaimResponseThreadFirstPostPermissions    `json:"permissions,omitempty"`
	PostAttachmentCount *int                                              `json:"post_attachment_count,omitempty"`
	PostBody            *string                                           `json:"post_body,omitempty"`
	PostBodyHTML        *string                                           `json:"post_body_html,omitempty"`
	PostBodyPlainText   *string                                           `json:"post_body_plain_text,omitempty"`
	PostCreateDate      *int                                              `json:"post_create_date,omitempty"`
	PostID              *int                                              `json:"post_id,omitempty"`
	PostIsDeleted       *bool                                             `json:"post_is_deleted,omitempty"`
	PostIsFirstPost     *bool                                             `json:"post_is_first_post,omitempty"`
	PostIsPublished     *bool                                             `json:"post_is_published,omitempty"`
	PostLikeCount       *int                                              `json:"post_like_count,omitempty"`
	PostUpdateDate      *int                                              `json:"post_update_date,omitempty"`
	PosterUserID        *int                                              `json:"poster_user_id,omitempty"`
	PosterUsername      *string                                           `json:"poster_username,omitempty"`
	Signature           *string                                           `json:"signature,omitempty"`
	SignatureHTML       *string                                           `json:"signature_html,omitempty"`
	SignaturePlainText  *string                                           `json:"signature_plain_text,omitempty"`
	ThreadID            *int                                              `json:"thread_id,omitempty"`
	UserIsIgnored       *bool                                             `json:"user_is_ignored,omitempty"`
}

type CreateClaimResponseThreadFirstPostLikeUsersItem struct {
	DisplayStyleGroupID *int    `json:"display_style_group_id,omitempty"`
	IsBanned            *int    `json:"is_banned,omitempty"`
	UniqUsernameCSS     *string `json:"uniq_username_css,omitempty"`
	UserID              *int    `json:"user_id,omitempty"`
	Username            *string `json:"username,omitempty"`
}

type CreateClaimResponseThreadFirstPostLinks struct {
	Attachments  *string `json:"attachments,omitempty"`
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type CreateClaimResponseThreadFirstPostPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Like             *bool `json:"like,omitempty"`
	Reply            *bool `json:"reply,omitempty"`
	Report           *bool `json:"report,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type CreateClaimResponseThreadForum struct {
	ForumDescription       *string                                           `json:"forum_description,omitempty"`
	ForumID                *int                                              `json:"forum_id,omitempty"`
	ForumIsFollowed        *bool                                             `json:"forum_is_followed,omitempty"`
	ForumPostCount         *int                                              `json:"forum_post_count,omitempty"`
	ForumPrefixes          []CreateClaimResponseThreadForumForumPrefixesItem `json:"forum_prefixes,omitempty"`
	ForumThreadCount       *int                                              `json:"forum_thread_count,omitempty"`
	ForumTitle             *string                                           `json:"forum_title,omitempty"`
	Links                  *CreateClaimResponseThreadForumLinks              `json:"links,omitempty"`
	Permissions            *CreateClaimResponseThreadForumPermissions        `json:"permissions,omitempty"`
	ThreadDefaultPrefixID  *int                                              `json:"thread_default_prefix_id,omitempty"`
	ThreadPrefixIsRequired *bool                                             `json:"thread_prefix_is_required,omitempty"`
}

type CreateClaimResponseThreadForumForumPrefixesItem struct {
	GroupPrefixes []CreateClaimResponseThreadForumForumPrefixesItemGroupPrefixesItem `json:"group_prefixes,omitempty"`
	GroupTitle    *string                                                            `json:"group_title,omitempty"`
}

type CreateClaimResponseThreadForumForumPrefixesItemGroupPrefixesItem struct {
	PrefixID    *int    `json:"prefix_id,omitempty"`
	PrefixTitle *string `json:"prefix_title,omitempty"`
}

type CreateClaimResponseThreadForumLinks struct {
	Detail        *string `json:"detail,omitempty"`
	Followers     *string `json:"followers,omitempty"`
	Permalink     *string `json:"permalink,omitempty"`
	SubCategories *string `json:"sub-categories,omitempty"`
	SubForums     *string `json:"sub-forums,omitempty"`
	Threads       *string `json:"threads,omitempty"`
}

type CreateClaimResponseThreadForumPermissions struct {
	CreateThread     *bool `json:"create_thread,omitempty"`
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	TagThread        *bool `json:"tag_thread,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type CreateClaimResponseThreadLinks struct {
	Detail            *string `json:"detail,omitempty"`
	FirstPost         *string `json:"first_post,omitempty"`
	FirstPoster       *string `json:"first_poster,omitempty"`
	FirstPosterAvatar *string `json:"first_poster_avatar,omitempty"`
	Followers         *string `json:"followers,omitempty"`
	Forum             *string `json:"forum,omitempty"`
	LastPost          *string `json:"last_post,omitempty"`
	LastPoster        *string `json:"last_poster,omitempty"`
	Permalink         *string `json:"permalink,omitempty"`
	Posts             *string `json:"posts,omitempty"`
}

type CreateClaimResponseThreadPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	Post             *bool `json:"post,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type CreateParams struct {
	// Additional information for you.
	AdditionalData *string `json:"additional_data,omitempty"`
	// Create a test invoice.
	IsTest *bool `json:"is_test,omitempty"`
	// Invoice lifetime.
	Lifetime *float64 `json:"lifetime,omitempty"`
	// Telegram User ID for which the invoice was created.
	RequiredTelegramID *int `json:"required_telegram_id,omitempty"`
	// Telegram Username (including @) for which the invoice was created (if any).
	RequiredTelegramUsername *string `json:"required_telegram_username,omitempty"`
	// Callback url.
	URLCallback *string `json:"url_callback,omitempty"`
}

type CreateResponse struct {
	AutoPaymentID *int            `json:"auto_payment_id,omitempty"`
	Message       *string         `json:"message,omitempty"`
	Status        *string         `json:"status,omitempty"`
	SystemInfo    *RespSystemInfo `json:"system_info,omitempty"`
}

type CreateResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type CurrencyResponse struct {
	CurrencyList    *CurrencyResponseCurrencyList `json:"currencyList,omitempty"`
	LastUpdate      *int                          `json:"lastUpdate,omitempty"`
	SystemInfo      *RespSystemInfo               `json:"system_info,omitempty"`
	VisitorCurrency *string                       `json:"visitorCurrency,omitempty"`
}

type CurrencyResponseCurrencyList struct {
	AED   *CurrencyResponseCurrencyListAED   `json:"AED,omitempty"`
	ARS   *CurrencyResponseCurrencyListARS   `json:"ARS,omitempty"`
	AUD   *CurrencyResponseCurrencyListAUD   `json:"AUD,omitempty"`
	BCH   *CurrencyResponseCurrencyListBCH   `json:"BCH,omitempty"`
	BGN   *CurrencyResponseCurrencyListBGN   `json:"BGN,omitempty"`
	BNB   *CurrencyResponseCurrencyListBNB   `json:"BNB,omitempty"`
	BRL   *CurrencyResponseCurrencyListBRL   `json:"BRL,omitempty"`
	BTC   *CurrencyResponseCurrencyListBTC   `json:"BTC,omitempty"`
	CAD   *CurrencyResponseCurrencyListCAD   `json:"CAD,omitempty"`
	CHF   *CurrencyResponseCurrencyListCHF   `json:"CHF,omitempty"`
	CLP   *CurrencyResponseCurrencyListCLP   `json:"CLP,omitempty"`
	CNY   *CurrencyResponseCurrencyListCNY   `json:"CNY,omitempty"`
	COP   *CurrencyResponseCurrencyListCOP   `json:"COP,omitempty"`
	CRC   *CurrencyResponseCurrencyListCRC   `json:"CRC,omitempty"`
	CZK   *CurrencyResponseCurrencyListCZK   `json:"CZK,omitempty"`
	DASH  *CurrencyResponseCurrencyListDASH  `json:"DASH,omitempty"`
	DKK   *CurrencyResponseCurrencyListDKK   `json:"DKK,omitempty"`
	DOGE  *CurrencyResponseCurrencyListDOGE  `json:"DOGE,omitempty"`
	ETH   *CurrencyResponseCurrencyListETH   `json:"ETH,omitempty"`
	EUR   *CurrencyResponseCurrencyListEUR   `json:"EUR,omitempty"`
	GBP   *CurrencyResponseCurrencyListGBP   `json:"GBP,omitempty"`
	GEL   *CurrencyResponseCurrencyListGEL   `json:"GEL,omitempty"`
	HKD   *CurrencyResponseCurrencyListHKD   `json:"HKD,omitempty"`
	HUF   *CurrencyResponseCurrencyListHUF   `json:"HUF,omitempty"`
	IDR   *CurrencyResponseCurrencyListIDR   `json:"IDR,omitempty"`
	ILS   *CurrencyResponseCurrencyListILS   `json:"ILS,omitempty"`
	INR   *CurrencyResponseCurrencyListINR   `json:"INR,omitempty"`
	JPY   *CurrencyResponseCurrencyListJPY   `json:"JPY,omitempty"`
	KRW   *CurrencyResponseCurrencyListKRW   `json:"KRW,omitempty"`
	KWD   *CurrencyResponseCurrencyListKWD   `json:"KWD,omitempty"`
	KZT   *CurrencyResponseCurrencyListKZT   `json:"KZT,omitempty"`
	LTC   *CurrencyResponseCurrencyListLTC   `json:"LTC,omitempty"`
	MATIC *CurrencyResponseCurrencyListMATIC `json:"MATIC,omitempty"`
	MXN   *CurrencyResponseCurrencyListMXN   `json:"MXN,omitempty"`
	MYR   *CurrencyResponseCurrencyListMYR   `json:"MYR,omitempty"`
	NOK   *CurrencyResponseCurrencyListNOK   `json:"NOK,omitempty"`
	NZD   *CurrencyResponseCurrencyListNZD   `json:"NZD,omitempty"`
	PEN   *CurrencyResponseCurrencyListPEN   `json:"PEN,omitempty"`
	PHP   *CurrencyResponseCurrencyListPHP   `json:"PHP,omitempty"`
	PLN   *CurrencyResponseCurrencyListPLN   `json:"PLN,omitempty"`
	QAR   *CurrencyResponseCurrencyListQAR   `json:"QAR,omitempty"`
	RON   *CurrencyResponseCurrencyListRON   `json:"RON,omitempty"`
	RSD   *CurrencyResponseCurrencyListRSD   `json:"RSD,omitempty"`
	RUB   *CurrencyResponseCurrencyListRUB   `json:"RUB,omitempty"`
	SAR   *CurrencyResponseCurrencyListSAR   `json:"SAR,omitempty"`
	SEK   *CurrencyResponseCurrencyListSEK   `json:"SEK,omitempty"`
	SGD   *CurrencyResponseCurrencyListSGD   `json:"SGD,omitempty"`
	SOL   *CurrencyResponseCurrencyListSOL   `json:"SOL,omitempty"`
	THB   *CurrencyResponseCurrencyListTHB   `json:"THB,omitempty"`
	TON   *CurrencyResponseCurrencyListTON   `json:"TON,omitempty"`
	TRX   *CurrencyResponseCurrencyListTRX   `json:"TRX,omitempty"`
	TRY   *CurrencyResponseCurrencyListTRY   `json:"TRY,omitempty"`
	TWD   *CurrencyResponseCurrencyListTWD   `json:"TWD,omitempty"`
	UAH   *CurrencyResponseCurrencyListUAH   `json:"UAH,omitempty"`
	USD   *CurrencyResponseCurrencyListUSD   `json:"USD,omitempty"`
	USDT  *CurrencyResponseCurrencyListUSDT  `json:"USDT,omitempty"`
	UYU   *CurrencyResponseCurrencyListUYU   `json:"UYU,omitempty"`
	VND   *CurrencyResponseCurrencyListVND   `json:"VND,omitempty"`
	XMR   *CurrencyResponseCurrencyListXMR   `json:"XMR,omitempty"`
	ZAR   *CurrencyResponseCurrencyListZAR   `json:"ZAR,omitempty"`
}

type CurrencyResponseCurrencyListAED struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListARS struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListAUD struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListBCH struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListBGN struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListBNB struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListBRL struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListBTC struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListCAD struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListCHF struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListCLP struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListCNY struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListCOP struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListCRC struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListCZK struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListDASH struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListDKK struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListDOGE struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListETH struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListEUR struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListGBP struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListGEL struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListHKD struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListHUF struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListIDR struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListILS struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListINR struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListJPY struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListKRW struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListKWD struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListKZT struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListLTC struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListMATIC struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListMXN struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListMYR struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListNOK struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListNZD struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListPEN struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListPHP struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListPLN struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListQAR struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListRON struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListRSD struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListRUB struct {
	FormattedRate *string `json:"formattedRate,omitempty"`
	Rate          *int    `json:"rate,omitempty"`
	Symbol        *string `json:"symbol,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListSAR struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListSEK struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListSGD struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListSOL struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListTHB struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListTON struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListTRX struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListTRY struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListTWD struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListUAH struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListUSD struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListUSDT struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListUYU struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListVND struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListXMR struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseCurrencyListZAR struct {
	FormattedRate *string  `json:"formattedRate,omitempty"`
	Rate          *float64 `json:"rate,omitempty"`
	Symbol        *string  `json:"symbol,omitempty"`
	Title         *string  `json:"title,omitempty"`
}

type CurrencyResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type DeleteParams struct {
	// Delete all proxies.
	DeleteAll *bool `json:"delete_all,omitempty"`
	// Id of an existing proxy.
	ProxyID *int `json:"proxy_id,omitempty"`
}

type DeleteResponse struct {
	Success    *bool           `json:"success,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type DeleteResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type DiscordParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// Has linked mobile.
	Tel *CategorySearchTel `json:"tel,omitempty"`
	// Has nitro.
	Nitro *CategorySearchNitro `json:"nitro,omitempty"`
	// Nitro type.
	NitroType []string `json:"nitro_type[],omitempty"`
	// Length of nitro.
	NitroLength *int `json:"nitro_length,omitempty"`
	// In what notation is time measured.
	NitroPeriod *CategorySearchNitroPeriod `json:"nitro_period,omitempty"`
	// Minimum number of boosts.
	BoostsMin *int `json:"boosts_min,omitempty"`
	// Maximum number of boosts.
	BoostsMax *int `json:"boosts_max,omitempty"`
	// Has billing.
	Billing *CategorySearchBilling `json:"billing,omitempty"`
	// Has gifts.
	Gifts *CategorySearchGifts `json:"gifts,omitempty"`
	// Has transactions.
	Transactions *CategorySearchTransactions `json:"transactions,omitempty"`
	// List of badges.
	Badge []string `json:"badge[],omitempty"`
	// List of account conditions.
	Condition []string `json:"condition[],omitempty"`
	// Minimum number of chats.
	ChatMin *int `json:"chat_min,omitempty"`
	// Maximum number of chats.
	ChatMax *int `json:"chat_max,omitempty"`
	// Minimum number of subscribers in server, where account is administrator/owner.
	MinAdminMembers *int `json:"min_admin_members,omitempty"`
	// Maximum number of subscribers in server, where account is administrator/owner.
	MaxAdminMembers *int `json:"max_admin_members,omitempty"`
	// Minimum number of servers, where account is administrator/owner.
	MinAdmin *int `json:"min_admin,omitempty"`
	// Maximum number of servers, where account is administrator/owner.
	MaxAdmin *int `json:"max_admin,omitempty"`
	// How old is the account.
	Reg *int `json:"reg,omitempty"`
	// In what notation is time measured.
	RegPeriod *CategorySearchRegPeriod `json:"reg_period,omitempty"`
	// List of languages.
	Language []string `json:"language[],omitempty"`
	// List of languages that won't be included.
	NotLanguage []string `json:"not_language[],omitempty"`
	// Has clans.
	Clans *CategorySearchClans `json:"clans,omitempty"`
	// Minimum number of clans, where account is administrator.
	MinAdminClans *int `json:"min_admin_clans,omitempty"`
	// Maximum number of clans, where account is administrator.
	MaxAdminClans *int `json:"max_admin_clans,omitempty"`
	// Minimum number of clans, where account is owner.
	MinOwnerClans *int `json:"min_owner_clans,omitempty"`
	// Maximum number of clans, where account is owner.
	MaxOwnerClans *int `json:"max_owner_clans,omitempty"`
	// List of allowed countries.
	Country []string `json:"country[],omitempty"`
	// List of disallowed countries.
	NotCountry []string `json:"not_country[],omitempty"`
	// Minimum count of servers.
	MinServers *int `json:"min_servers,omitempty"`
	// Maximum count of servers.
	MaxServers *int `json:"max_servers,omitempty"`
	// Has two-factor authentication.
	Field2fa *CategorySearch2fa `json:"2fa,omitempty"`
	// Minimum number of Nitro full credits.
	MinFullCredits *int `json:"min_full_credits,omitempty"`
	// Maximum number of Nitro full credits.
	MaxFullCredits *int `json:"max_full_credits,omitempty"`
	// Minimum number of Nitro basic credits.
	MinBasicCredits *int `json:"min_basic_credits,omitempty"`
	// Maximum number of Nitro basic credits.
	MaxBasicCredits *int `json:"max_basic_credits,omitempty"`
	// Minimum number of Discord Orbs.
	MinOrbs *int `json:"min_orbs,omitempty"`
	// Maximum number of Discord Orbs.
	MaxOrbs *int `json:"max_orbs,omitempty"`
}

type DiscordResponse struct {
	CacheTTL        *int                       `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                      `json:"hasNextPage,omitempty"`
	Items           []DiscordResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                       `json:"lastModified,omitempty"`
	Page            *int                       `json:"page,omitempty"`
	PerPage         *int                       `json:"perPage,omitempty"`
	SearchUrl       *string                    `json:"searchUrl,omitempty"`
	ServerTime      *int                       `json:"serverTime,omitempty"`
	StickyItems     []interface{}              `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo            `json:"system_info,omitempty"`
	TotalItems      *int                       `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}                `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                      `json:"wasCached,omitempty"`
}

type DiscordResponseItemsItem struct {
	AllowAskDiscount             *int                                  `json:"allow_ask_discount,omitempty"`
	BumpSettings                 *DiscordResponseItemsItemBumpSettings `json:"bumpSettings,omitempty"`
	CanBumpItem                  *bool                                 `json:"canBumpItem,omitempty"`
	CanBuyItem                   *bool                                 `json:"canBuyItem,omitempty"`
	CanChangePassword            *bool                                 `json:"canChangePassword,omitempty"`
	CanCloseItem                 *bool                                 `json:"canCloseItem,omitempty"`
	CanDeleteItem                *bool                                 `json:"canDeleteItem,omitempty"`
	CanEditItem                  *bool                                 `json:"canEditItem,omitempty"`
	CanOpenItem                  *bool                                 `json:"canOpenItem,omitempty"`
	CanReportItem                *bool                                 `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase   *bool                                 `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem                 *bool                                 `json:"canStickItem,omitempty"`
	CanUnstickItem               *bool                                 `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats           *bool                                 `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount           *bool                                 `json:"canValidateAccount,omitempty"`
	CanViewAccountLink           *bool                                 `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData        *bool                                 `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData             *bool                                 `json:"canViewLoginData,omitempty"`
	CategoryID                   *int                                  `json:"category_id,omitempty"`
	Description                  *string                               `json:"description,omitempty"`
	DescriptionEnHtml            *string                               `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain           *string                               `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml              *string                               `json:"descriptionHtml,omitempty"`
	DescriptionPlain             *string                               `json:"descriptionPlain,omitempty"`
	DescriptionEn                *string                               `json:"description_en,omitempty"`
	DiscordAccountConditionLabel *string                               `json:"discordAccountConditionLabel,omitempty"`
	DiscordLocaleTitle           *string                               `json:"discordLocaleTitle,omitempty"`
	DiscordNitroType             interface{}                           `json:"discordNitroType,omitempty"`
	DiscordAdminMembersCount     *int                                  `json:"discord_admin_members_count,omitempty"`
	DiscordAdminServers          *string                               `json:"discord_admin_servers,omitempty"`
	DiscordAdminServersCount     *int                                  `json:"discord_admin_servers_count,omitempty"`
	DiscordAvailableBoosts       *int                                  `json:"discord_available_boosts,omitempty"`
	DiscordBilling               *int                                  `json:"discord_billing,omitempty"`
	DiscordChatCount             *int                                  `json:"discord_chat_count,omitempty"`
	DiscordCondition             *string                               `json:"discord_condition,omitempty"`
	DiscordGifts                 *int                                  `json:"discord_gifts,omitempty"`
	DiscordItemID                *int                                  `json:"discord_item_id,omitempty"`
	DiscordLocale                *string                               `json:"discord_locale,omitempty"`
	DiscordNitroEndDate          *int                                  `json:"discord_nitro_end_date,omitempty"`
	DiscordRegisterDate          *int                                  `json:"discord_register_date,omitempty"`
	DiscordVerified              *int                                  `json:"discord_verified,omitempty"`
	EditDate                     *int                                  `json:"edit_date,omitempty"`
	EmailLoginUrl                *string                               `json:"emailLoginUrl,omitempty"`
	EmailProvider                *string                               `json:"email_provider,omitempty"`
	EmailType                    *string                               `json:"email_type,omitempty"`
	ExtendedGuarantee            *int                                  `json:"extended_guarantee,omitempty"`
	FeedbackData                 interface{}                           `json:"feedback_data,omitempty"`
	Guarantee                    interface{}                           `json:"guarantee,omitempty"`
	HasPendingAutoBuy            *bool                                 `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                    *bool                                 `json:"isIgnored,omitempty"`
	IsSticky                     *int                                  `json:"is_sticky,omitempty"`
	ItemOriginPhrase             *string                               `json:"itemOriginPhrase,omitempty"`
	ItemDomain                   *string                               `json:"item_domain,omitempty"`
	ItemID                       *int                                  `json:"item_id,omitempty"`
	ItemOrigin                   *string                               `json:"item_origin,omitempty"`
	ItemState                    *string                               `json:"item_state,omitempty"`
	NoteText                     interface{}                           `json:"note_text,omitempty"`
	Nsb                          *int                                  `json:"nsb,omitempty"`
	Price                        *int                                  `json:"price,omitempty"`
	PriceWithSellerFee           *float64                              `json:"priceWithSellerFee,omitempty"`
	PriceCurrency                *string                               `json:"price_currency,omitempty"`
	PublishedDate                *int                                  `json:"published_date,omitempty"`
	RefreshedDate                *int                                  `json:"refreshed_date,omitempty"`
	ResaleItemOrigin             *string                               `json:"resale_item_origin,omitempty"`
	RubPrice                     *int                                  `json:"rub_price,omitempty"`
	Seller                       *DiscordResponseItemsItemSeller       `json:"seller,omitempty"`
	ShowGetEmailCodeButton       *bool                                 `json:"showGetEmailCodeButton,omitempty"`
	Tags                         interface{}                           `json:"tags,omitempty"`
	Title                        *string                               `json:"title,omitempty"`
	TitleEn                      *string                               `json:"title_en,omitempty"`
	UpdateStatDate               *int                                  `json:"update_stat_date,omitempty"`
	ViewCount                    *int                                  `json:"view_count,omitempty"`
}

type DiscordResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type DiscordResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     interface{} `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type DiscordResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type DiscountModel struct {
	CategoryID      *int `json:"category_id,omitempty"`
	DiscountID      *int `json:"discount_id,omitempty"`
	DiscountPercent *int `json:"discount_percent,omitempty"`
	DiscountUserID  *int `json:"discount_user_id,omitempty"`
	MaxPrice        *int `json:"max_price,omitempty"`
	MinPrice        *int `json:"min_price,omitempty"`
	UserID          *int `json:"user_id,omitempty"`
}

type DiscountRequestParams struct {
	// Message to the seller.
	Message *string `json:"message,omitempty"`
}

type DownloadParams struct {
	// Format of the downloaded accounts.
	Format *AccountsListFormat `json:"format,omitempty"`
	// Custom format string for download. (Required if **format** is set to **custom**)
	CustomFormat *string `json:"custom_format,omitempty"`
	// Accounts category.
	CategoryID *AccountsListCategoryID `json:"category_id,omitempty"`
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Account status.
	Show *AccountsListShow `json:"show,omitempty"`
	// Delete reason. (Only if **show** is set to **deleted**)
	DeleteReason *string `json:"delete_reason,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Order by.
	OrderBy *AccountsListOrderBy `json:"order_by,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Username of buyer. (If **show** is **paid**)
	Username *string `json:"username,omitempty"`
	// Start date for filtering by publication date.
	PublishedStartDate *string `json:"published_startDate,omitempty"`
	// End date for filtering by publication date.
	PublishedEndDate *string `json:"published_endDate,omitempty"`
	// Enable filtering by publication date.
	FilterByPublishedDate *bool `json:"filter_by_published_date,omitempty"`
	// Start date for filtering by buyer operation date.
	PaidStartDate *string `json:"paid_startDate,omitempty"`
	// End date for filtering by buyer operation date.
	PaidEndDate *string `json:"paid_endDate,omitempty"`
	// Enable filtering by buyer operation date.
	FilterByBuyerOperationDate *bool `json:"filter_by_buyer_operation_date,omitempty"`
	// Start date for filtering by deletion date.
	DeleteStartDate *string `json:"delete_startDate,omitempty"`
	// End date for filtering by deletion date.
	DeleteEndDate *string `json:"delete_endDate,omitempty"`
	// Enable filtering by deletion date.
	FilterByDeleteDate *bool `json:"filter_by_delete_date,omitempty"`
}

type EAParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// List of games.
	Game []string `json:"game[],omitempty"`
	// List of allowed countries.
	Country []string `json:"country[],omitempty"`
	// List of disallowed countries.
	NotCountry []string `json:"not_country[],omitempty"`
	// Minimum count of games.
	Gmin *int `json:"gmin,omitempty"`
	// Maximum count of games.
	Gmax *int `json:"gmax,omitempty"`
	// Minimum rank points in Apex Legends.
	AlRankMin *int `json:"al_rank_min,omitempty"`
	// Maximum rank points in Apex Legends.
	AlRankMax *int `json:"al_rank_max,omitempty"`
	// Minimum level in Apex Legends.
	AlLevelMin *int `json:"al_level_min,omitempty"`
	// Maximum level in Apex Legends.
	AlLevelMax *int `json:"al_level_max,omitempty"`
	// Has a ban in any game.
	HasBan *CategorySearchHasBan `json:"has_ban,omitempty"`
	// Xbox connected to account.
	XboxConnected *CategorySearchXboxConnected `json:"xbox_connected,omitempty"`
	// Steam connected to account.
	SteamConnected *CategorySearchSteamConnected `json:"steam_connected,omitempty"`
	// PSN connected to account.
	PsnConnected *CategorySearchPsnConnected `json:"psn_connected,omitempty"`
	// Name of subscription.
	Subscription *CategorySearchSubscription `json:"subscription,omitempty"`
	// Length of subscription.
	SubscriptionLength *int `json:"subscription_length,omitempty"`
	// In what notation is time measured.
	SubscriptionPeriod *CategorySearchSubscriptionPeriod `json:"subscription_period,omitempty"`
	// List of minimum hours played by game.
	HoursPlayed map[string]int `json:"hours_played,omitempty"`
	// List of maximum hours played by game.
	HoursPlayedMax map[string]int `json:"hours_played_max,omitempty"`
	// Has transactions.
	Transactions *CategorySearchTransactions `json:"transactions,omitempty"`
}

type EAResponse struct {
	CacheTTL        *int                  `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                 `json:"hasNextPage,omitempty"`
	Items           []EAResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                  `json:"lastModified,omitempty"`
	Page            *int                  `json:"page,omitempty"`
	PerPage         *int                  `json:"perPage,omitempty"`
	SearchUrl       *string               `json:"searchUrl,omitempty"`
	ServerTime      *int                  `json:"serverTime,omitempty"`
	StickyItems     []interface{}         `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo       `json:"system_info,omitempty"`
	TotalItems      *int                  `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}           `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                 `json:"wasCached,omitempty"`
}

type EAResponseItemsItem struct {
	AccountLink                *string                               `json:"accountLink,omitempty"`
	AccountLinks               []EAResponseItemsItemAccountLinksItem `json:"accountLinks,omitempty"`
	AllowAskDiscount           *int                                  `json:"allow_ask_discount,omitempty"`
	BumpSettings               *EAResponseItemsItemBumpSettings      `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                                 `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                                 `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                                 `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                                 `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                                 `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                                 `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                                 `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                                 `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                                 `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                                 `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                                 `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                                 `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                                 `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                                 `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                                 `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                                 `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                  `json:"category_id,omitempty"`
	Description                *string                               `json:"description,omitempty"`
	DescriptionEnHtml          *string                               `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                               `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                               `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                               `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                               `json:"description_en,omitempty"`
	EAAlLevel                  *int                                  `json:"ea_al_level,omitempty"`
	EAAlRankScore              *int                                  `json:"ea_al_rank_score,omitempty"`
	EABans                     []interface{}                         `json:"ea_bans,omitempty"`
	EACountry                  *string                               `json:"ea_country,omitempty"`
	EAGameCount                *int                                  `json:"ea_game_count,omitempty"`
	EAGames                    *EAResponseItemsItemEAGames           `json:"ea_games,omitempty"`
	EAHasBan                   *int                                  `json:"ea_has_ban,omitempty"`
	EAID                       *int                                  `json:"ea_id,omitempty"`
	EAItemID                   *int                                  `json:"ea_item_id,omitempty"`
	EALastActivity             *int                                  `json:"ea_last_activity,omitempty"`
	EAPsnConnected             *int                                  `json:"ea_psn_connected,omitempty"`
	EASteamConnected           *int                                  `json:"ea_steam_connected,omitempty"`
	EASubscription             *string                               `json:"ea_subscription,omitempty"`
	EASubscriptionEndDate      *int                                  `json:"ea_subscription_end_date,omitempty"`
	EAUsername                 *string                               `json:"ea_username,omitempty"`
	EAXboxConnected            *int                                  `json:"ea_xbox_connected,omitempty"`
	EditDate                   *int                                  `json:"edit_date,omitempty"`
	EmailLoginUrl              *string                               `json:"emailLoginUrl,omitempty"`
	EmailProvider              *string                               `json:"email_provider,omitempty"`
	EmailType                  *string                               `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                                  `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                           `json:"feedback_data,omitempty"`
	Guarantee                  interface{}                           `json:"guarantee,omitempty"`
	HasPendingAutoBuy          *bool                                 `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                  *bool                                 `json:"isIgnored,omitempty"`
	IsSticky                   *int                                  `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                               `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                               `json:"item_domain,omitempty"`
	ItemID                     *int                                  `json:"item_id,omitempty"`
	ItemOrigin                 *string                               `json:"item_origin,omitempty"`
	ItemState                  *string                               `json:"item_state,omitempty"`
	NoteText                   interface{}                           `json:"note_text,omitempty"`
	Nsb                        *int                                  `json:"nsb,omitempty"`
	Price                      *int                                  `json:"price,omitempty"`
	PriceWithSellerFee         *float64                              `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                               `json:"price_currency,omitempty"`
	PublishedDate              *int                                  `json:"published_date,omitempty"`
	RefreshedDate              *int                                  `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                               `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount  *int                                  `json:"restore_items_category_count,omitempty"`
	RubPrice                   *int                                  `json:"rub_price,omitempty"`
	Seller                     *EAResponseItemsItemSeller            `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                                 `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount     *int                                  `json:"sold_items_category_count,omitempty"`
	Tags                       interface{}                           `json:"tags,omitempty"`
	Title                      *string                               `json:"title,omitempty"`
	TitleEn                    *string                               `json:"title_en,omitempty"`
	UpdateStatDate             *int                                  `json:"update_stat_date,omitempty"`
	ViewCount                  *int                                  `json:"view_count,omitempty"`
}

type EAResponseItemsItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type EAResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type EAResponseItemsItemEAGames struct {
	ApexLegends *EAResponseItemsItemEAGamesApexLegends `json:"apex-legends,omitempty"`
}

type EAResponseItemsItemEAGamesApexLegends struct {
	GameID       *string `json:"game_id,omitempty"`
	Img          *string `json:"img,omitempty"`
	LastActivity *int    `json:"last_activity,omitempty"`
	Title        *string `json:"title,omitempty"`
	TotalPlayed  *int    `json:"total_played,omitempty"`
}

type EAResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type EAResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type EditParams struct {
	// Allow users to ask discount for this account.
	AllowAskDiscount *bool                         `json:"allow_ask_discount,omitempty"`
	Currency         *AccountsManagingEditCurrency `json:"currency,omitempty"`
	// Account public description.
	Description *string `json:"description,omitempty"`
	// Email login data (email:password format).
	EmailLoginData *string `json:"email_login_data,omitempty"`
	// Email type.
	EmailType *AccountsManagingEditEmailType `json:"email_type,omitempty"`
	// Account private information (visible only for buyer).
	Information *string `json:"information,omitempty"`
	// Account origin. Where did you get it from.
	ItemOrigin *AccountsManagingEditItemOrigin `json:"item_origin,omitempty"`
	// Current price of account in your currency.
	Price *int `json:"price,omitempty"`
	// Using proxy id for account checking. See GET or POST /proxy to get or edit proxy list.
	ProxyID *int `json:"proxy_id,omitempty"`
	// Title of account. If **title** specified and **title_en** is empty, **title_en** will be automatically translated to English language.
	Title *string `json:"title,omitempty"`
	// English title of account. If **title_en** specified and **title** is empty, **title** will be automatically translated to Russian language.
	TitleEn *string `json:"title_en,omitempty"`
}

type EditResponse struct {
	Discounts  []EditResponseDiscountsItem `json:"discounts,omitempty"`
	SystemInfo *RespSystemInfo             `json:"system_info,omitempty"`
	Total      *int                        `json:"total,omitempty"`
}

type EditResponseDiscountsItem struct {
	CategoryID      *int `json:"category_id,omitempty"`
	DiscountID      *int `json:"discount_id,omitempty"`
	DiscountPercent *int `json:"discount_percent,omitempty"`
	DiscountUserID  *int `json:"discount_user_id,omitempty"`
	MaxPrice        *int `json:"max_price,omitempty"`
	MinPrice        *int `json:"min_price,omitempty"`
	UserID          *int `json:"user_id,omitempty"`
}

type EditResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type EpicGamesParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// Guarantee type.
	Eg *CategorySearchEg `json:"eg,omitempty"`
	// List of games.
	Game []string `json:"game[],omitempty"`
	// Can change email.
	ChangeEmail *CategorySearchChangeEmail `json:"change_email,omitempty"`
	// Has Rocket League purchases.
	RlPurchases *bool `json:"rl_purchases,omitempty"`
	// Minimum epic wallet balance.
	BalanceMin *float64 `json:"balance_min,omitempty"`
	// Maximum epic wallet balance.
	BalanceMax *float64 `json:"balance_max,omitempty"`
	// Minimum rewards balance.
	RewardsBalanceMin *float64 `json:"rewards_balance_min,omitempty"`
	// Maximum rewards balance.
	RewardsBalanceMax *float64 `json:"rewards_balance_max,omitempty"`
	// Minimum number of games.
	Gmin *int `json:"gmin,omitempty"`
	// Maximum number of games.
	Gmax *int `json:"gmax,omitempty"`
	// List of allowed countries.
	Country []string `json:"country[],omitempty"`
	// List of disallowed countries.
	NotCountry []string `json:"not_country[],omitempty"`
	// Number of days the account has been offline.
	Daybreak *int `json:"daybreak,omitempty"`
	// List of minimum hours played by game.
	HoursPlayed map[string]int `json:"hours_played,omitempty"`
	// List of maximum hours played by game.
	HoursPlayedMax map[string]int `json:"hours_played_max,omitempty"`
}

type EpicGamesResponse struct {
	CacheTTL        *int                         `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                        `json:"hasNextPage,omitempty"`
	Items           []EpicGamesResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                         `json:"lastModified,omitempty"`
	Page            *int                         `json:"page,omitempty"`
	PerPage         *int                         `json:"perPage,omitempty"`
	SearchUrl       *string                      `json:"searchUrl,omitempty"`
	ServerTime      *int                         `json:"serverTime,omitempty"`
	StickyItems     []interface{}                `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo              `json:"system_info,omitempty"`
	TotalItems      *int                         `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}                  `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                        `json:"wasCached,omitempty"`
}

type EpicGamesResponseItemsItem struct {
	AccountLinks               []interface{}                                  `json:"accountLinks,omitempty"`
	AllowAskDiscount           *int                                           `json:"allow_ask_discount,omitempty"`
	BumpSettings               *EpicGamesResponseItemsItemBumpSettings        `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                                          `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                                          `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                                          `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                                          `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                                          `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                                          `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                                          `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                                          `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                                          `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                                          `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                                          `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                                          `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                                          `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                                          `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                                          `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                                          `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                           `json:"category_id,omitempty"`
	Description                *string                                        `json:"description,omitempty"`
	DescriptionEnHtml          *string                                        `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                                        `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                                        `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                                        `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                                        `json:"description_en,omitempty"`
	EditDate                   *int                                           `json:"edit_date,omitempty"`
	EgBalance                  *string                                        `json:"egBalance,omitempty"`
	EgGameCount                *int                                           `json:"egGameCount,omitempty"`
	EgTransactions             []EpicGamesResponseItemsItemEgTransactionsItem `json:"egTransactions,omitempty"`
	EgCanUpdateDisplayName     *int                                           `json:"eg_can_update_display_name,omitempty"`
	EgChangeEmail              *int                                           `json:"eg_change_email,omitempty"`
	EgCodeRedemptionHistory    []interface{}                                  `json:"eg_code_redemption_history,omitempty"`
	EgCountry                  *string                                        `json:"eg_country,omitempty"`
	EgCoupons                  []interface{}                                  `json:"eg_coupons,omitempty"`
	EgGames                    interface{}                                    `json:"eg_games,omitempty"`
	EgItemID                   *int                                           `json:"eg_item_id,omitempty"`
	EgLastActivity             *int                                           `json:"eg_last_activity,omitempty"`
	EgNextChangeEmailDate      *int                                           `json:"eg_next_change_email_date,omitempty"`
	EgPaymentMethods           []interface{}                                  `json:"eg_payment_methods,omitempty"`
	EgRewardsBalance           *int                                           `json:"eg_rewards_balance,omitempty"`
	EgRewardsExpirationDate    *int                                           `json:"eg_rewards_expiration_date,omitempty"`
	EgRlPurchases              *int                                           `json:"eg_rl_purchases,omitempty"`
	EgUsername                 *string                                        `json:"eg_username,omitempty"`
	EmailLoginUrl              *string                                        `json:"emailLoginUrl,omitempty"`
	EmailProvider              *string                                        `json:"email_provider,omitempty"`
	EmailType                  *string                                        `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                                           `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                                    `json:"feedback_data,omitempty"`
	Guarantee                  interface{}                                    `json:"guarantee,omitempty"`
	HasPendingAutoBuy          *bool                                          `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                  *bool                                          `json:"isIgnored,omitempty"`
	IsSticky                   *int                                           `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                                        `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                                        `json:"item_domain,omitempty"`
	ItemID                     *int                                           `json:"item_id,omitempty"`
	ItemOrigin                 *string                                        `json:"item_origin,omitempty"`
	ItemState                  *string                                        `json:"item_state,omitempty"`
	NoteText                   interface{}                                    `json:"note_text,omitempty"`
	Nsb                        *int                                           `json:"nsb,omitempty"`
	Price                      *int                                           `json:"price,omitempty"`
	PriceWithSellerFee         *float64                                       `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                                        `json:"price_currency,omitempty"`
	PublishedDate              *int                                           `json:"published_date,omitempty"`
	RefreshedDate              *int                                           `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                                        `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount  *int                                           `json:"restore_items_category_count,omitempty"`
	RubPrice                   *int                                           `json:"rub_price,omitempty"`
	Seller                     *EpicGamesResponseItemsItemSeller              `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                                          `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount     *int                                           `json:"sold_items_category_count,omitempty"`
	Tags                       interface{}                                    `json:"tags,omitempty"`
	Title                      *string                                        `json:"title,omitempty"`
	TitleEn                    *string                                        `json:"title_en,omitempty"`
	UpdateStatDate             *int                                           `json:"update_stat_date,omitempty"`
	ViewCount                  *int                                           `json:"view_count,omitempty"`
}

type EpicGamesResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type EpicGamesResponseItemsItemEgTransactionsItem struct {
	Date             *int    `json:"date,omitempty"`
	OrderType        *string `json:"orderType,omitempty"`
	PresentmentTotal *string `json:"presentmentTotal,omitempty"`
	Title            *string `json:"title,omitempty"`
}

type EpicGamesResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type EpicGamesResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type EscapeFromTarkovParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// Region.
	Region *CategorySearchRegion `json:"region,omitempty"`
	// List of versions.
	Version []string `json:"version[],omitempty"`
	// How old is the account.
	Reg *int `json:"reg,omitempty"`
	// In what notation is time measured.
	RegPeriod *CategorySearchRegPeriod `json:"reg_period,omitempty"`
	// Minimum level.
	LevelMin *int `json:"level_min,omitempty"`
	// Maximum level.
	LevelMax *int `json:"level_max,omitempty"`
	// Access to pve.
	Pve *CategorySearchPve `json:"pve,omitempty"`
	// Side in current wipe.
	Side *CategorySearchSide `json:"side,omitempty"`
}

type EscapeFromTarkovResponse struct {
	CacheTTL        *int                                `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                               `json:"hasNextPage,omitempty"`
	Items           []EscapeFromTarkovResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                                `json:"lastModified,omitempty"`
	Page            *int                                `json:"page,omitempty"`
	PerPage         *int                                `json:"perPage,omitempty"`
	SearchUrl       *string                             `json:"searchUrl,omitempty"`
	ServerTime      *int                                `json:"serverTime,omitempty"`
	StickyItems     []interface{}                       `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo                     `json:"system_info,omitempty"`
	TotalItems      *int                                `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}                         `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                               `json:"wasCached,omitempty"`
}

type EscapeFromTarkovResponseItemsItem struct {
	AccountDomain              *string                                        `json:"accountDomain,omitempty"`
	AllowAskDiscount           *int                                           `json:"allow_ask_discount,omitempty"`
	BumpSettings               *EscapeFromTarkovResponseItemsItemBumpSettings `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                                          `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                                          `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                                          `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                                          `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                                          `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                                          `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                                          `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                                          `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                                          `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                                          `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                                          `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                                          `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                                          `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                                          `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                                          `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                                          `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                           `json:"category_id,omitempty"`
	Description                *string                                        `json:"description,omitempty"`
	DescriptionEnHtml          *string                                        `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                                        `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                                        `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                                        `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                                        `json:"description_en,omitempty"`
	EditDate                   *int                                           `json:"edit_date,omitempty"`
	EmailLoginUrl              *string                                        `json:"emailLoginUrl,omitempty"`
	EmailProvider              *string                                        `json:"email_provider,omitempty"`
	EmailType                  *string                                        `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                                           `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                                    `json:"feedback_data,omitempty"`
	Guarantee                  interface{}                                    `json:"guarantee,omitempty"`
	HasPendingAutoBuy          *bool                                          `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                  *bool                                          `json:"isIgnored,omitempty"`
	IsSticky                   *int                                           `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                                        `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                                        `json:"item_domain,omitempty"`
	ItemID                     *int                                           `json:"item_id,omitempty"`
	ItemOrigin                 *string                                        `json:"item_origin,omitempty"`
	ItemState                  *string                                        `json:"item_state,omitempty"`
	NoteText                   interface{}                                    `json:"note_text,omitempty"`
	Nsb                        *int                                           `json:"nsb,omitempty"`
	Price                      *int                                           `json:"price,omitempty"`
	PriceWithSellerFee         *float64                                       `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                                        `json:"price_currency,omitempty"`
	PublishedDate              *int                                           `json:"published_date,omitempty"`
	RefreshedDate              *int                                           `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                                        `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount  *int                                           `json:"restore_items_category_count,omitempty"`
	RubPrice                   *int                                           `json:"rub_price,omitempty"`
	Seller                     *EscapeFromTarkovResponseItemsItemSeller       `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                                          `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount     *int                                           `json:"sold_items_category_count,omitempty"`
	Tags                       interface{}                                    `json:"tags,omitempty"`
	TarkovGameVersionPhrase    *string                                        `json:"tarkovGameVersionPhrase,omitempty"`
	TarkovKd                   *int                                           `json:"tarkovKd,omitempty"`
	TarkovRegionPhrase         *string                                        `json:"tarkovRegionPhrase,omitempty"`
	TarkovSecuredContainer     *string                                        `json:"tarkovSecuredContainer,omitempty"`
	TarkovDeaths               *int                                           `json:"tarkov_deaths,omitempty"`
	TarkovDollars              *int                                           `json:"tarkov_dollars,omitempty"`
	TarkovEuros                *int                                           `json:"tarkov_euros,omitempty"`
	TarkovExp                  *int                                           `json:"tarkov_exp,omitempty"`
	TarkovGameVersion          *string                                        `json:"tarkov_game_version,omitempty"`
	TarkovItemID               *int                                           `json:"tarkov_item_id,omitempty"`
	TarkovKills                *int                                           `json:"tarkov_kills,omitempty"`
	TarkovLastActivity         *int                                           `json:"tarkov_last_activity,omitempty"`
	TarkovLevel                *int                                           `json:"tarkov_level,omitempty"`
	TarkovMailForwarding       *int                                           `json:"tarkov_mail_forwarding,omitempty"`
	TarkovPurchaseDate         *int                                           `json:"tarkov_purchase_date,omitempty"`
	TarkovRegion               *string                                        `json:"tarkov_region,omitempty"`
	TarkovRegisterDate         *int                                           `json:"tarkov_register_date,omitempty"`
	TarkovRubles               *int                                           `json:"tarkov_rubles,omitempty"`
	TarkovSessions             *int                                           `json:"tarkov_sessions,omitempty"`
	TarkovSide                 *string                                        `json:"tarkov_side,omitempty"`
	TarkovTotalInGame          *int                                           `json:"tarkov_total_in_game,omitempty"`
	TarkovUsername             *string                                        `json:"tarkov_username,omitempty"`
	Title                      *string                                        `json:"title,omitempty"`
	TitleEn                    *string                                        `json:"title_en,omitempty"`
	UpdateStatDate             *int                                           `json:"update_stat_date,omitempty"`
	ViewCount                  *int                                           `json:"view_count,omitempty"`
}

type EscapeFromTarkovResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type EscapeFromTarkovResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type EscapeFromTarkovResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ExternalParams struct {
	// Cookies.
	Cookies *string `json:"cookies,omitempty"`
	// Email login data (email:password format).
	EmailLoginData *string `json:"email_login_data,omitempty"`
	// Account login data (login:password format).
	Login *string `json:"login,omitempty"`
}

type ExtraModel struct {
	// Ark. Optional. Used only if you want to upload Steam account.
	Ark *bool `json:"ark,omitempty"`
	// Ark Ascended. Optional. Used only if you want to upload Steam account.
	ArkAscended *bool `json:"ark_ascended,omitempty"`
	// Check channels. Optional. Used only if you want to upload Telegram account.
	CheckChannels *bool `json:"checkChannels,omitempty"`
	// Check ban on Hypixel. Optional. Used only if you want to upload Minecraft account.
	CheckHypixelBan *bool `json:"checkHypixelBan,omitempty"`
	// Check spam. Optional. Used only if you want to upload Telegram account.
	CheckSpam *bool `json:"checkSpam,omitempty"`
	// If set, the item will be closed **item_state = closed**.
	CloseItem *bool `json:"close_item,omitempty"`
	// Code from email (in case of problems). Optional if you want to upload Supercell account.
	ConfirmationCode *string `json:"confirmationCode,omitempty"`
	// Cookie login. Optional. Used only if you want to upload TikTok account.
	CookieLogin *bool `json:"cookie_login,omitempty"`
	// Cookies. Required if you want to upload Fortnite, Epic Games, Social Club, Instagram or TikTok account.
	Cookies *string `json:"cookies,omitempty"`
	// Dota 2 MMR. Optional. Used only if you want to upload Steam account.
	Dota2Mmr *int `json:"dota2_mmr,omitempty"`
	// EA Games. Optional. Used only if you want to upload Steam account.
	EAGames *bool `json:"ea_games,omitempty"`
	// Genshin Impact Primogems count. Optional. Used only if you want to upload miHoYo account.
	GenshinCurrency *int `json:"genshin_currency,omitempty"`
	// Honkai Star Rail Stellar Jade count. Optional. Used only if you want to upload miHoYo account.
	HonkaiCurrency *int `json:"honkai_currency,omitempty"`
	// Login without cookies. Optional if you want to upload Instagram account.
	LoginWithoutCookies *bool `json:"login_without_cookies,omitempty"`
	// Steam mafile. Optional. Used only if you want to upload Steam account.
	MFAFile *string `json:"mfa_file,omitempty"`
	// Telegram 2FA Password. Optional. Used only if you want to upload Telegram account.
	Password *string `json:"password,omitempty"`
	// Proxy line format ip:port:user:pass (prioritize over proxy_id parameter).
	Proxy *string `json:"proxy,omitempty"`
	// Region. Required if you want to upload WoT account. Optional if you want to upload miHoYo or Riot account.
	Region *string `json:"region,omitempty"`
	// Service. Required if you want to upload VPN, Cinema account or gift.
	Service *string `json:"service,omitempty"`
	// Supercell system. Required if you want to upload Supercell account.
	System *string `json:"system,omitempty"`
	// Telegram client. Optional. Used only if you want to upload Telegram account.
	TelegramClient *string `json:"telegramClient,omitempty"`
	// Contents of session.json file. Optional. Used only if you want to upload Telegram account.
	TelegramJson *string `json:"telegramJson,omitempty"`
	// The quarry. Optional. Used only if you want to upload Steam account.
	TheQuarry *bool `json:"the_quarry,omitempty"`
	// Uplay Games. Optional. Used only if you want to upload Steam account.
	UplayGames *bool `json:"uplay_games,omitempty"`
	// Warframe. Optional. Used only if you want to upload Steam account.
	Warframe *bool `json:"warframe,omitempty"`
	// Zenless Zone Zero Polychrome count. Optional. Used only if you want to upload miHoYo account.
	ZenlessCurrency *int `json:"zenless_currency,omitempty"`
}

type FastBuyParams struct {
	// Balance ID that will be used to purchase specified item.
	BalanceID *int `json:"balance_id,omitempty"`
	// Current price of account in your currency.
	Price *float64 `json:"price,omitempty"`
}

type FastSellParams struct {
	// Allow users to ask discount for this account.
	AllowAskDiscount *bool `json:"allow_ask_discount,omitempty"`
	// Account public description.
	Description *string `json:"description,omitempty"`
	// Required if a **category** is one of list of Required email login data categories. Email login data (email:password format).
	EmailLoginData *string `json:"email_login_data,omitempty"`
	// Email type.
	EmailType *AccountPublishingFastSellEmailType `json:"email_type,omitempty"`
	// Guarantee type.
	ExtendedGuarantee *AccountPublishingFastSellExtendedGuarantee `json:"extended_guarantee,omitempty"`
	Extra             map[string]interface{}                      `json:"extra,omitempty"`
	// Required if a **category** is one of list of Required email login data categories.
	HasEmailLoginData *bool `json:"has_email_login_data,omitempty"`
	// Account private information (visible only for buyer).
	Information *string `json:"information,omitempty"`
	// Account login (or email).
	Login *string `json:"login,omitempty"`
	// Account login data (login:password format).
	LoginPassword *string `json:"login_password,omitempty"`
	// Account password.
	Password *string `json:"password,omitempty"`
	// Proxy id that will be used to check account.
	ProxyID     *int        `json:"proxy_id,omitempty"`
	RandomProxy interface{} `json:"random_proxy,omitempty"`
	// Title of account. If **title** specified and **title_en** is empty, **title_en** will be automatically translated to English language.
	Title *string `json:"title,omitempty"`
	// English title of account. If **title_en** specified and **title** is empty, **title** will be automatically translated to Russian language.
	TitleEn *string `json:"title_en,omitempty"`
}

type FastSellResponse struct {
	Item       *ItemModel      `json:"item,omitempty"`
	ItemLink   *string         `json:"itemLink,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type FastSellResponseItem struct {
	AccountLink                       *string                                `json:"accountLink,omitempty"`
	AccountLinks                      []FastSellResponseItemAccountLinksItem `json:"accountLinks,omitempty"`
	AccountLastActivity               *int                                   `json:"account_last_activity,omitempty"`
	AiPrice                           *int                                   `json:"aiPrice,omitempty"`
	AiPriceCheckDate                  *int                                   `json:"aiPriceCheckDate,omitempty"`
	AllowAskDiscount                  *int                                   `json:"allow_ask_discount,omitempty"`
	AutoBuyPrice                      *int                                   `json:"autoBuyPrice,omitempty"`
	AutoBuyPriceCheckDate             *int                                   `json:"autoBuyPriceCheckDate,omitempty"`
	BumpSettings                      *FastSellResponseItemBumpSettings      `json:"bumpSettings,omitempty"`
	Buyer                             *FastSellResponseItemBuyer             `json:"buyer,omitempty"`
	BuyerAvatarDate                   *int                                   `json:"buyer_avatar_date,omitempty"`
	BuyerDisplayIconGroupID           *int                                   `json:"buyer_display_icon_group_id,omitempty"`
	BuyerUniqBanner                   *string                                `json:"buyer_uniq_banner,omitempty"`
	BuyerUserGroupID                  *int                                   `json:"buyer_user_group_id,omitempty"`
	CanAskDiscount                    *bool                                  `json:"canAskDiscount,omitempty"`
	CanChangeEmailPassword            *bool                                  `json:"canChangeEmailPassword,omitempty"`
	CanChangePassword                 *bool                                  `json:"canChangePassword,omitempty"`
	CanCheckAiPrice                   *bool                                  `json:"canCheckAiPrice,omitempty"`
	CanCheckAutoBuyPrice              *bool                                  `json:"canCheckAutoBuyPrice,omitempty"`
	CanCheckGuarantee                 *bool                                  `json:"canCheckGuarantee,omitempty"`
	CanReportItem                     *bool                                  `json:"canReportItem,omitempty"`
	CanResellItem                     *bool                                  `json:"canResellItem,omitempty"`
	CanResellItemAfterPurchase        *bool                                  `json:"canResellItemAfterPurchase,omitempty"`
	CanShareItem                      *bool                                  `json:"canShareItem,omitempty"`
	CanUpdateItemStats                *bool                                  `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount                *bool                                  `json:"canValidateAccount,omitempty"`
	CanViewAccountLink                *bool                                  `json:"canViewAccountLink,omitempty"`
	CanViewAccountLoginAndTempEmail   *bool                                  `json:"canViewAccountLoginAndTempEmail,omitempty"`
	CanViewEmailLoginData             *bool                                  `json:"canViewEmailLoginData,omitempty"`
	CanViewItemViews                  *bool                                  `json:"canViewItemViews,omitempty"`
	CanViewLoginData                  *bool                                  `json:"canViewLoginData,omitempty"`
	CartPrice                         interface{}                            `json:"cart_price,omitempty"`
	CategoryID                        *int                                   `json:"category_id,omitempty"`
	ContentID                         interface{}                            `json:"content_id,omitempty"`
	ContentType                       interface{}                            `json:"content_type,omitempty"`
	CopyFormatData                    *FastSellResponseItemCopyFormatData    `json:"copyFormatData,omitempty"`
	CustomFields                      *FastSellResponseItemCustomFields      `json:"customFields,omitempty"`
	DeleteDate                        *int                                   `json:"delete_date,omitempty"`
	DeleteReason                      *string                                `json:"delete_reason,omitempty"`
	DeleteUserID                      *int                                   `json:"delete_user_id,omitempty"`
	DeleteUsername                    *string                                `json:"delete_username,omitempty"`
	Deposit                           *int                                   `json:"deposit,omitempty"`
	Description                       *string                                `json:"description,omitempty"`
	DescriptionEnHtml                 *string                                `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain                *string                                `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml                   *string                                `json:"descriptionHtml,omitempty"`
	DescriptionPlain                  *string                                `json:"descriptionPlain,omitempty"`
	DescriptionEn                     *string                                `json:"description_en,omitempty"`
	EditDate                          *int                                   `json:"edit_date,omitempty"`
	EmailProvider                     *string                                `json:"email_provider,omitempty"`
	EmailType                         *string                                `json:"email_type,omitempty"`
	ExtendedGuarantee                 *int                                   `json:"extended_guarantee,omitempty"`
	ExternalAuth                      []interface{}                          `json:"externalAuth,omitempty"`
	ExtraPrices                       []FastSellResponseItemExtraPricesItem  `json:"extraPrices,omitempty"`
	FeedbackData                      interface{}                            `json:"feedback_data,omitempty"`
	GetEmailCodeDisplayLogin          interface{}                            `json:"getEmailCodeDisplayLogin,omitempty"`
	Guarantee                         interface{}                            `json:"guarantee,omitempty"`
	ImagePreviewLinks                 []string                               `json:"imagePreviewLinks,omitempty"`
	InCart                            interface{}                            `json:"in_cart,omitempty"`
	Information                       *string                                `json:"information,omitempty"`
	InformationEn                     *string                                `json:"information_en,omitempty"`
	IsBirthdayToday                   *bool                                  `json:"isBirthdayToday,omitempty"`
	IsIgnored                         *bool                                  `json:"isIgnored,omitempty"`
	IsPersonalAccount                 *bool                                  `json:"isPersonalAccount,omitempty"`
	IsSmallExf                        *bool                                  `json:"isSmallExf,omitempty"`
	IsTrusted                         *bool                                  `json:"isTrusted,omitempty"`
	IsFave                            interface{}                            `json:"is_fave,omitempty"`
	IsSticky                          *int                                   `json:"is_sticky,omitempty"`
	ItemOriginPhrase                  *string                                `json:"itemOriginPhrase,omitempty"`
	ItemDomain                        *string                                `json:"item_domain,omitempty"`
	ItemID                            *int                                   `json:"item_id,omitempty"`
	ItemOrigin                        *string                                `json:"item_origin,omitempty"`
	ItemState                         *string                                `json:"item_state,omitempty"`
	Login                             *string                                `json:"login,omitempty"`
	LoginData                         *FastSellResponseItemLoginData         `json:"loginData,omitempty"`
	MarketCustomTitle                 *string                                `json:"market_custom_title,omitempty"`
	MaxDiscountPercent                *int                                   `json:"max_discount_percent,omitempty"`
	NeedToRequireVideoToViewLoginData *bool                                  `json:"needToRequireVideoToViewLoginData,omitempty"`
	NoteText                          *string                                `json:"note_text,omitempty"`
	Nsb                               *int                                   `json:"nsb,omitempty"`
	PendingDeletionDate               *int                                   `json:"pending_deletion_date,omitempty"`
	Price                             *int                                   `json:"price,omitempty"`
	PriceWithSellerFee                *float64                               `json:"priceWithSellerFee,omitempty"`
	PriceWithSellerFeeLabel           *string                                `json:"priceWithSellerFeeLabel,omitempty"`
	PriceCurrency                     *string                                `json:"price_currency,omitempty"`
	PublishedDate                     *int                                   `json:"published_date,omitempty"`
	RefreshedDate                     *int                                   `json:"refreshed_date,omitempty"`
	ResaleItemOrigin                  *string                                `json:"resale_item_origin,omitempty"`
	RubPrice                          *int                                   `json:"rub_price,omitempty"`
	Seller                            *FastSellResponseItemSeller            `json:"seller,omitempty"`
	ShowGetEmailCodeButton            *bool                                  `json:"showGetEmailCodeButton,omitempty"`
	Tags                              interface{}                            `json:"tags,omitempty"`
	TempEmail                         *string                                `json:"temp_email,omitempty"`
	Title                             *string                                `json:"title,omitempty"`
	TitleEn                           *string                                `json:"title_en,omitempty"`
	UniqueKeyExists                   *bool                                  `json:"uniqueKeyExists,omitempty"`
	UpdateStatDate                    *int                                   `json:"update_stat_date,omitempty"`
	UserAllowAskDiscount              *int                                   `json:"user_allow_ask_discount,omitempty"`
	ViewCount                         *int                                   `json:"view_count,omitempty"`
	VisitorIsAuthor                   *bool                                  `json:"visitorIsAuthor,omitempty"`
}

type FastSellResponseItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type FastSellResponseItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	NextAllowedBumpDate interface{} `json:"nextAllowedBumpDate,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type FastSellResponseItemBuyer struct {
	DisplayIconGroupID  *int    `json:"display_icon_group_id,omitempty"`
	DisplayStyleGroupID *int    `json:"display_style_group_id,omitempty"`
	IsBanned            *int    `json:"is_banned,omitempty"`
	OperationDate       *int    `json:"operation_date,omitempty"`
	UniqBanner          *string `json:"uniq_banner,omitempty"`
	UniqUsernameCSS     *string `json:"uniq_username_css,omitempty"`
	UserGroupID         *int    `json:"user_group_id,omitempty"`
	UserID              *int    `json:"user_id,omitempty"`
	Username            *string `json:"username,omitempty"`
	VisitorIsBuyer      *bool   `json:"visitorIsBuyer,omitempty"`
}

type FastSellResponseItemCopyFormatData struct {
	Full      *string `json:"full,omitempty"`
	LoginData *string `json:"login_data,omitempty"`
	TitleLink *string `json:"title_link,omitempty"`
}

type FastSellResponseItemCustomFields struct {
	Field4         *string       `json:"_4,omitempty"`
	AllowSelfUnban []interface{} `json:"allowSelfUnban,omitempty"`
	BanReason      *string       `json:"ban_reason,omitempty"`
	Discord        *string       `json:"discord,omitempty"`
	Github         *string       `json:"github,omitempty"`
	Jabber         *string       `json:"jabber,omitempty"`
	LztUnbanAmount *string       `json:"lztUnbanAmount,omitempty"`
	Steam          *string       `json:"steam,omitempty"`
	Telegram       *string       `json:"telegram,omitempty"`
	Vk             *string       `json:"vk,omitempty"`
}

type FastSellResponseItemExtraPricesItem struct {
	Currency   *string  `json:"currency,omitempty"`
	Price      *string  `json:"price,omitempty"`
	PriceValue *float64 `json:"priceValue,omitempty"`
}

type FastSellResponseItemGuarantee struct {
	Active                *bool   `json:"active,omitempty"`
	Cancelled             *bool   `json:"cancelled,omitempty"`
	CancelledReason       *string `json:"cancelledReason,omitempty"`
	CancelledReasonPhrase *string `json:"cancelledReasonPhrase,omitempty"`
	Class                 *string `json:"class,omitempty"`
	Duration              *int    `json:"duration,omitempty"`
	DurationPhrase        *string `json:"durationPhrase,omitempty"`
	EndDate               *int    `json:"endDate,omitempty"`
	RemainingTime         *int    `json:"remainingTime,omitempty"`
	RemainingTimePhrase   *string `json:"remainingTimePhrase,omitempty"`
}

type FastSellResponseItemLoginData struct {
	EncodedOldPassword interface{} `json:"encodedOldPassword,omitempty"`
	EncodedPassword    *string     `json:"encodedPassword,omitempty"`
	EncodedRaw         *string     `json:"encodedRaw,omitempty"`
	Login              *string     `json:"login,omitempty"`
	OldPassword        *string     `json:"oldPassword,omitempty"`
	Password           *string     `json:"password,omitempty"`
	Raw                *string     `json:"raw,omitempty"`
}

type FastSellResponseItemSeller struct {
	ActiveItemsCount      *int                                `json:"active_items_count,omitempty"`
	AvatarDate            *int                                `json:"avatar_date,omitempty"`
	Contacts              *FastSellResponseItemSellerContacts `json:"contacts,omitempty"`
	DisplayStyleGroupID   *int                                `json:"display_style_group_id,omitempty"`
	EffectiveLastActivity *int                                `json:"effective_last_activity,omitempty"`
	IsOnline              *bool                               `json:"isOnline,omitempty"`
	IsBanned              *int                                `json:"is_banned,omitempty"`
	JoinedDate            *int                                `json:"joined_date,omitempty"`
	RestoreData           interface{}                         `json:"restore_data,omitempty"`
	RestorePercents       interface{}                         `json:"restore_percents,omitempty"`
	SoldItemsCount        *int                                `json:"sold_items_count,omitempty"`
	UserID                *int                                `json:"user_id,omitempty"`
	Username              *string                             `json:"username,omitempty"`
}

type FastSellResponseItemSellerContacts struct {
	BanReason *string `json:"ban_reason,omitempty"`
	Telegram  *string `json:"telegram,omitempty"`
}

type FastSellResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type FavoritesParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Account status.
	Show *AccountsListShow `json:"show,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Order by.
	OrderBy *AccountsListOrderBy `json:"order_by,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
}

type FeeParams struct {
	// Amount you want to send in your currency.
	Amount *float64 `json:"amount,omitempty"`
}

type FeeResponse struct {
	Calculator           interface{}     `json:"calculator,omitempty"`
	CommissionPercentage *int            `json:"commission_percentage,omitempty"`
	SpentCurrentMonth    *int            `json:"spentCurrentMonth,omitempty"`
	SystemInfo           *RespSystemInfo `json:"system_info,omitempty"`
}

type FeeResponseCalculator struct {
	CommissionAmount  *int `json:"commissionAmount,omitempty"`
	InputAmount       *int `json:"inputAmount,omitempty"`
	TotalOutputAmount *int `json:"totalOutputAmount,omitempty"`
}

type FeeResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type FortniteParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Access to market temp mail.
	TempEmail *CategorySearchTempEmail `json:"temp_email,omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// Guarantee type.
	Eg *CategorySearchEg `json:"eg,omitempty"`
	// Minimum number of skins.
	Smin *int `json:"smin,omitempty"`
	// Maximum number of skins.
	Smax *int `json:"smax,omitempty"`
	// Minimum number of V-Bucks.
	Vbmin *int `json:"vbmin,omitempty"`
	// Maximum number of V-Bucks.
	Vbmax *int `json:"vbmax,omitempty"`
	// Skins.
	Skin []string `json:"skin[],omitempty"`
	// Pickaxes.
	Pickaxe []string `json:"pickaxe[],omitempty"`
	// Gliders.
	Glider []string `json:"glider[],omitempty"`
	// Dances.
	Dance []string `json:"dance[],omitempty"`
	// Can change email.
	ChangeEmail *CategorySearchChangeEmail `json:"change_email,omitempty"`
	// Platform.
	Platform []string `json:"platform[],omitempty"`
	// Minimum number of shop skins.
	SkinsShopMin *int `json:"skins_shop_min,omitempty"`
	// Maximum number of shop skins.
	SkinsShopMax *int `json:"skins_shop_max,omitempty"`
	// Minimum number of shop pickaxes.
	PickaxesShopMin *int `json:"pickaxes_shop_min,omitempty"`
	// Maximum number of shop pickaxes.
	PickaxesShopMax *int `json:"pickaxes_shop_max,omitempty"`
	// Minimum number of shop dances.
	DancesShopMin *int `json:"dances_shop_min,omitempty"`
	// Maximum number of shop dances.
	DancesShopMax *int `json:"dances_shop_max,omitempty"`
	// Minimum number of shop gliders.
	GlidersShopMin *int `json:"gliders_shop_min,omitempty"`
	// Maximum number of shop gliders.
	GlidersShopMax *int `json:"gliders_shop_max,omitempty"`
	// Minimum total cost of all skins in the shop in V-Bucks.
	SkinsShopVbmin *int `json:"skins_shop_vbmin,omitempty"`
	// Maximum total cost of all skins in the shop in V-Bucks.
	SkinsShopVbmax *int `json:"skins_shop_vbmax,omitempty"`
	// Minimum total cost of all pickaxes in the shop in V-Bucks.
	PickaxesShopVbmin *int `json:"pickaxes_shop_vbmin,omitempty"`
	// Maximum total cost of all pickaxes in the shop in V-Bucks.
	PickaxesShopVbmax *int `json:"pickaxes_shop_vbmax,omitempty"`
	// Minimum total cost of all dances in the shop in V-Bucks.
	DancesShopVbmin *int `json:"dances_shop_vbmin,omitempty"`
	// Maximum total cost of all dances in the shop in V-Bucks.
	DancesShopVbmax *int `json:"dances_shop_vbmax,omitempty"`
	// Minimum total cost of all gliders in the shop in V-Bucks.
	GlidersShopVbmin *int `json:"gliders_shop_vbmin,omitempty"`
	// Maximum total cost of all gliders in the shop in V-Bucks.
	GlidersShopVbmax *int `json:"gliders_shop_vbmax,omitempty"`
	// Has Battle Pass.
	Bp *CategorySearchBp `json:"bp,omitempty"`
	// Minimum level.
	Lmin *int `json:"lmin,omitempty"`
	// Maximum level.
	Lmax *int `json:"lmax,omitempty"`
	// Minimum level of Battle Pass.
	BpLmin *int `json:"bp_lmin,omitempty"`
	// Maximum level of Battle Pass.
	BpLmax *int `json:"bp_lmax,omitempty"`
	// How old is last transaction.
	LastTransDate *int `json:"last_trans_date,omitempty"`
	// In what notation is time measured.
	LastTransDatePeriod *CategorySearchLastTransDatePeriod `json:"last_trans_date_period,omitempty"`
	// Has no transactions.
	NoTrans *bool `json:"no_trans,omitempty"`
	// Can be linked to Xbox.
	XboxLinkable *CategorySearchXboxLinkable `json:"xbox_linkable,omitempty"`
	// Can be linked to PSN.
	PsnLinkable *CategorySearchPsnLinkable `json:"psn_linkable,omitempty"`
	// Number of days the account has been offline.
	Daybreak *int `json:"daybreak,omitempty"`
	// Has Rocket League purchases.
	RlPurchases *bool `json:"rl_purchases,omitempty"`
	// How old is the account.
	Reg *int `json:"reg,omitempty"`
	// In what notation is time measured.
	RegPeriod *CategorySearchRegPeriod `json:"reg_period,omitempty"`
	// Minimum number of available refund credits.
	RefundCreditsMin *int `json:"refund_credits_min,omitempty"`
	// Maximum number of available refund credits.
	RefundCreditsMax *int `json:"refund_credits_max,omitempty"`
	// Minimum number of pickaxes.
	PickaxeMin *int `json:"pickaxe_min,omitempty"`
	// Maximum number of pickaxes.
	PickaxeMax *int `json:"pickaxe_max,omitempty"`
	// Minimum number of dances.
	Dmin *int `json:"dmin,omitempty"`
	// Maximum number of dances.
	Dmax *int `json:"dmax,omitempty"`
	// Minimum number of gliders.
	Gmin *int `json:"gmin,omitempty"`
	// Maximum number of gliders.
	Gmax *int `json:"gmax,omitempty"`
	// List of allowed countries.
	Country []string `json:"country[],omitempty"`
	// List of disallowed countries.
	NotCountry []string `json:"not_country[],omitempty"`
}

type FortniteResponse struct {
	CacheTTL        *int                        `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                       `json:"hasNextPage,omitempty"`
	Items           []FortniteResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                        `json:"lastModified,omitempty"`
	Page            *int                        `json:"page,omitempty"`
	PerPage         *int                        `json:"perPage,omitempty"`
	SearchUrl       *string                     `json:"searchUrl,omitempty"`
	ServerTime      *int                        `json:"serverTime,omitempty"`
	StickyItems     []interface{}               `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo             `json:"system_info,omitempty"`
	TotalItems      *int                        `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}                 `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                       `json:"wasCached,omitempty"`
}

type FortniteResponseItemsItem struct {
	AccountLinks                []interface{}                                       `json:"accountLinks,omitempty"`
	AccountLastActivity         *int                                                `json:"account_last_activity,omitempty"`
	AllowAskDiscount            *int                                                `json:"allow_ask_discount,omitempty"`
	BumpSettings                *FortniteResponseItemsItemBumpSettings              `json:"bumpSettings,omitempty"`
	CanBumpItem                 *bool                                               `json:"canBumpItem,omitempty"`
	CanBuyItem                  *bool                                               `json:"canBuyItem,omitempty"`
	CanChangePassword           *bool                                               `json:"canChangePassword,omitempty"`
	CanCloseItem                *bool                                               `json:"canCloseItem,omitempty"`
	CanDeleteItem               *bool                                               `json:"canDeleteItem,omitempty"`
	CanEditItem                 *bool                                               `json:"canEditItem,omitempty"`
	CanOpenItem                 *bool                                               `json:"canOpenItem,omitempty"`
	CanReportItem               *bool                                               `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase  *bool                                               `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem                *bool                                               `json:"canStickItem,omitempty"`
	CanUnstickItem              *bool                                               `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats          *bool                                               `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount          *bool                                               `json:"canValidateAccount,omitempty"`
	CanViewAccountLink          *bool                                               `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData       *bool                                               `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData            *bool                                               `json:"canViewLoginData,omitempty"`
	CategoryID                  *int                                                `json:"category_id,omitempty"`
	Description                 *string                                             `json:"description,omitempty"`
	DescriptionEnHtml           *string                                             `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain          *string                                             `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml             *string                                             `json:"descriptionHtml,omitempty"`
	DescriptionPlain            *string                                             `json:"descriptionPlain,omitempty"`
	DescriptionEn               *string                                             `json:"description_en,omitempty"`
	Domain                      *string                                             `json:"domain,omitempty"`
	EditDate                    *int                                                `json:"edit_date,omitempty"`
	EmailLoginUrl               *string                                             `json:"emailLoginUrl,omitempty"`
	EmailProvider               *string                                             `json:"email_provider,omitempty"`
	EmailType                   *string                                             `json:"email_type,omitempty"`
	ExtendedGuarantee           *int                                                `json:"extended_guarantee,omitempty"`
	FeedbackData                interface{}                                         `json:"feedback_data,omitempty"`
	FortniteDance               []FortniteResponseItemsItemFortniteDanceItem        `json:"fortniteDance,omitempty"`
	FortniteGliders             []FortniteResponseItemsItemFortniteGlidersItem      `json:"fortniteGliders,omitempty"`
	FortnitePastSeasons         []FortniteResponseItemsItemFortnitePastSeasonsItem  `json:"fortnitePastSeasons,omitempty"`
	FortnitePickaxe             []FortniteResponseItemsItemFortnitePickaxeItem      `json:"fortnitePickaxe,omitempty"`
	FortniteSkins               []FortniteResponseItemsItemFortniteSkinsItem        `json:"fortniteSkins,omitempty"`
	FortniteTransactions        []FortniteResponseItemsItemFortniteTransactionsItem `json:"fortniteTransactions,omitempty"`
	FortniteBalance             *int                                                `json:"fortnite_balance,omitempty"`
	FortniteBookLevel           *int                                                `json:"fortnite_book_level,omitempty"`
	FortniteBooksPurchased      *int                                                `json:"fortnite_books_purchased,omitempty"`
	FortniteChangeEmail         *int                                                `json:"fortnite_change_email,omitempty"`
	FortniteDanceCount          *int                                                `json:"fortnite_dance_count,omitempty"`
	FortniteGliderCount         *int                                                `json:"fortnite_glider_count,omitempty"`
	FortniteItemID              *int                                                `json:"fortnite_item_id,omitempty"`
	FortniteLastActivity        *int                                                `json:"fortnite_last_activity,omitempty"`
	FortniteLastTransDate       *int                                                `json:"fortnite_last_trans_date,omitempty"`
	FortniteLevel               *int                                                `json:"fortnite_level,omitempty"`
	FortniteLifetimeWins        *int                                                `json:"fortnite_lifetime_wins,omitempty"`
	FortniteNextChangeEmailDate *int                                                `json:"fortnite_next_change_email_date,omitempty"`
	FortnitePickaxeCount        *int                                                `json:"fortnite_pickaxe_count,omitempty"`
	FortnitePlatform            *string                                             `json:"fortnite_platform,omitempty"`
	FortnitePsnLinkable         *int                                                `json:"fortnite_psn_linkable,omitempty"`
	FortniteRegisterDate        *int                                                `json:"fortnite_register_date,omitempty"`
	FortniteRlPurchases         *int                                                `json:"fortnite_rl_purchases,omitempty"`
	FortniteSeasonNum           *int                                                `json:"fortnite_season_num,omitempty"`
	FortniteShopDancesCount     *int                                                `json:"fortnite_shop_dances_count,omitempty"`
	FortniteShopGlidersCount    *int                                                `json:"fortnite_shop_gliders_count,omitempty"`
	FortniteShopPickaxesCount   *int                                                `json:"fortnite_shop_pickaxes_count,omitempty"`
	FortniteShopSkinsCount      *int                                                `json:"fortnite_shop_skins_count,omitempty"`
	FortniteSkinCount           *int                                                `json:"fortnite_skin_count,omitempty"`
	FortniteXboxLinkable        *int                                                `json:"fortnite_xbox_linkable,omitempty"`
	Guarantee                   interface{}                                         `json:"guarantee,omitempty"`
	HasPendingAutoBuy           *bool                                               `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                   *bool                                               `json:"isIgnored,omitempty"`
	IsSmallExf                  *bool                                               `json:"isSmallExf,omitempty"`
	IsSticky                    *int                                                `json:"is_sticky,omitempty"`
	ItemOriginPhrase            *string                                             `json:"itemOriginPhrase,omitempty"`
	ItemDomain                  *string                                             `json:"item_domain,omitempty"`
	ItemID                      *int                                                `json:"item_id,omitempty"`
	ItemOrigin                  *string                                             `json:"item_origin,omitempty"`
	ItemState                   *string                                             `json:"item_state,omitempty"`
	NoteText                    interface{}                                         `json:"note_text,omitempty"`
	Nsb                         *int                                                `json:"nsb,omitempty"`
	Price                       *int                                                `json:"price,omitempty"`
	PriceWithSellerFee          *float64                                            `json:"priceWithSellerFee,omitempty"`
	PriceCurrency               *string                                             `json:"price_currency,omitempty"`
	PublishedDate               *int                                                `json:"published_date,omitempty"`
	RefreshedDate               *int                                                `json:"refreshed_date,omitempty"`
	ResaleItemOrigin            *string                                             `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount   *int                                                `json:"restore_items_category_count,omitempty"`
	RubPrice                    *int                                                `json:"rub_price,omitempty"`
	Seller                      *FortniteResponseItemsItemSeller                    `json:"seller,omitempty"`
	ShopCounts                  *FortniteResponseItemsItemShopCounts                `json:"shopCounts,omitempty"`
	ShowGetEmailCodeButton      *bool                                               `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount      *int                                                `json:"sold_items_category_count,omitempty"`
	Tags                        interface{}                                         `json:"tags,omitempty"`
	Title                       *string                                             `json:"title,omitempty"`
	TitleEn                     *string                                             `json:"title_en,omitempty"`
	UpdateStatDate              *int                                                `json:"update_stat_date,omitempty"`
	ViewCount                   *int                                                `json:"view_count,omitempty"`
}

type FortniteResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type FortniteResponseItemsItemFortniteDanceItem struct {
	FromShop *int    `json:"from_shop,omitempty"`
	ID       *string `json:"id,omitempty"`
	Rarity   *string `json:"rarity,omitempty"`
	Title    *string `json:"title,omitempty"`
	Type_    *string `json:"type,omitempty"`
}

type FortniteResponseItemsItemFortniteGlidersItem struct {
	FromShop *int    `json:"from_shop,omitempty"`
	ID       *string `json:"id,omitempty"`
	Rarity   *string `json:"rarity,omitempty"`
	Title    *string `json:"title,omitempty"`
	Type_    *string `json:"type,omitempty"`
}

type FortniteResponseItemsItemFortnitePastSeasonsItem struct {
	BookLevel        *int  `json:"bookLevel,omitempty"`
	NumHighBracket   *int  `json:"numHighBracket,omitempty"`
	NumLowBracket    *int  `json:"numLowBracket,omitempty"`
	NumRoyalRoyales  *int  `json:"numRoyalRoyales,omitempty"`
	NumWins          *int  `json:"numWins,omitempty"`
	PurchasedVIP     *bool `json:"purchasedVIP,omitempty"`
	SeasonLevel      *int  `json:"seasonLevel,omitempty"`
	SeasonNumber     *int  `json:"seasonNumber,omitempty"`
	SeasonXp         *int  `json:"seasonXp,omitempty"`
	SurvivorPrestige *int  `json:"survivorPrestige,omitempty"`
	SurvivorTier     *int  `json:"survivorTier,omitempty"`
}

type FortniteResponseItemsItemFortnitePickaxeItem struct {
	FromShop *int    `json:"from_shop,omitempty"`
	ID       *string `json:"id,omitempty"`
	Rarity   *string `json:"rarity,omitempty"`
	Title    *string `json:"title,omitempty"`
	Type_    *string `json:"type,omitempty"`
}

type FortniteResponseItemsItemFortniteSkinsItem struct {
	FromShop *int    `json:"from_shop,omitempty"`
	ID       *string `json:"id,omitempty"`
	Rarity   *string `json:"rarity,omitempty"`
	Title    *string `json:"title,omitempty"`
	Type_    *string `json:"type,omitempty"`
}

type FortniteResponseItemsItemFortniteTransactionsItem struct {
	Date             *int    `json:"date,omitempty"`
	OrderType        *string `json:"orderType,omitempty"`
	PresentmentTotal *string `json:"presentmentTotal,omitempty"`
	Title            *string `json:"title,omitempty"`
}

type FortniteResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type FortniteResponseItemsItemShopCounts struct {
	ShopDancesCount   *int `json:"shopDancesCount,omitempty"`
	ShopGlidersCount  *int `json:"shopGlidersCount,omitempty"`
	ShopPickaxesCount *int `json:"shopPickaxesCount,omitempty"`
	ShopSkinsCount    *int `json:"shopSkinsCount,omitempty"`
}

type FortniteResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type GamesResponse struct {
	Games      []GamesResponseGamesItem `json:"games,omitempty"`
	SystemInfo *RespSystemInfo          `json:"system_info,omitempty"`
}

type GamesResponseGamesItem struct {
	Abbr       *string `json:"abbr,omitempty"`
	AppID      *string `json:"app_id,omitempty"`
	CategoryID *int    `json:"category_id,omitempty"`
	Img        *string `json:"img,omitempty"`
	Ru         *string `json:"ru,omitempty"`
	Title      *string `json:"title,omitempty"`
	URL        *string `json:"url,omitempty"`
}

type GamesResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type GetLetters2Params struct {
	// Email login data (email:password format). Required if both *email* and *password* are not provided.
	EmailPassword *string `json:"email_password,omitempty"`
	// Email. Required if *email_password* is not provided.
	Email *string `json:"email,omitempty"`
	// Password. Required if *email_password* is not provided.
	Password *string `json:"password,omitempty"`
	// Number of letters to return.
	Limit *int `json:"limit,omitempty"`
}

type GetLetters2Response struct {
	Email      *string                          `json:"email,omitempty"`
	Letters    []GetLetters2ResponseLettersItem `json:"letters,omitempty"`
	SystemInfo *RespSystemInfo                  `json:"system_info,omitempty"`
}

type GetLetters2ResponseLettersItem struct {
	Date      *int    `json:"date,omitempty"`
	From      *string `json:"from,omitempty"`
	TextHtml  *string `json:"textHtml,omitempty"`
	TextPlain *string `json:"textPlain,omitempty"`
}

type GetLetters2ResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type GetMafileResponse struct {
	MaFile     *GetMafileResponseMaFile `json:"maFile,omitempty"`
	SystemInfo *RespSystemInfo          `json:"system_info,omitempty"`
}

type GetMafileResponseMaFile struct {
	Session        *GetMafileResponseMaFileSession `json:"Session,omitempty"`
	AccountName    *string                         `json:"account_name,omitempty"`
	DeviceID       *string                         `json:"device_id,omitempty"`
	FullyEnrolled  *bool                           `json:"fully_enrolled,omitempty"`
	IdentitySecret *string                         `json:"identity_secret,omitempty"`
	RevocationCode *string                         `json:"revocation_code,omitempty"`
	Secret1        *string                         `json:"secret_1,omitempty"`
	SerialNumber   *int                            `json:"serial_number,omitempty"`
	SharedSecret   *string                         `json:"shared_secret,omitempty"`
	TokenGid       *string                         `json:"token_gid,omitempty"`
	URI            *string                         `json:"uri,omitempty"`
}

type GetMafileResponseMaFileSession struct {
	AccessToken      *string `json:"AccessToken,omitempty"`
	RefreshToken     *string `json:"RefreshToken,omitempty"`
	SessionID        *string `json:"SessionID,omitempty"`
	SteamID          *string `json:"SteamID,omitempty"`
	SteamLoginSecure *string `json:"SteamLoginSecure,omitempty"`
}

type GetMafileResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type GetParams struct {
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
}

type GetResponse struct {
	Discounts  []GetResponseDiscountsItem `json:"discounts,omitempty"`
	SystemInfo *RespSystemInfo            `json:"system_info,omitempty"`
	Total      *int                       `json:"total,omitempty"`
}

type GetResponseDiscountsItem struct {
	CategoryID      *int `json:"category_id,omitempty"`
	DiscountID      *int `json:"discount_id,omitempty"`
	DiscountPercent *int `json:"discount_percent,omitempty"`
	DiscountUserID  *int `json:"discount_user_id,omitempty"`
	MaxPrice        *int `json:"max_price,omitempty"`
	MinPrice        *int `json:"min_price,omitempty"`
	UserID          *int `json:"user_id,omitempty"`
}

type GetResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type GiftsParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Name of subscription.
	Subscription *CategorySearchSubscription `json:"subscription,omitempty"`
	// Length of subscription.
	SubscriptionLength *int `json:"subscription_length,omitempty"`
	// In what notation is time measured.
	SubscriptionPeriod *CategorySearchSubscriptionPeriod `json:"subscription_period,omitempty"`
}

type GiftsResponse struct {
	CacheTTL        *int                     `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                    `json:"hasNextPage,omitempty"`
	Items           []GiftsResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                     `json:"lastModified,omitempty"`
	Page            *int                     `json:"page,omitempty"`
	PerPage         *int                     `json:"perPage,omitempty"`
	SearchUrl       *string                  `json:"searchUrl,omitempty"`
	ServerTime      *int                     `json:"serverTime,omitempty"`
	StickyItems     []interface{}            `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo          `json:"system_info,omitempty"`
	TotalItems      *int                     `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}              `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                    `json:"wasCached,omitempty"`
}

type GiftsResponseItemsItem struct {
	AllowAskDiscount           *int                                `json:"allow_ask_discount,omitempty"`
	BumpSettings               *GiftsResponseItemsItemBumpSettings `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                               `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                               `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                               `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                               `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                               `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                               `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                               `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                               `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                               `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                               `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                               `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                               `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                               `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                               `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                               `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                               `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                `json:"category_id,omitempty"`
	Description                *string                             `json:"description,omitempty"`
	DescriptionEnHtml          *string                             `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                             `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                             `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                             `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                             `json:"description_en,omitempty"`
	EditDate                   *int                                `json:"edit_date,omitempty"`
	EmailProvider              interface{}                         `json:"email_provider,omitempty"`
	EmailType                  *string                             `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                                `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                         `json:"feedback_data,omitempty"`
	GiftsServiceName           *string                             `json:"giftsServiceName,omitempty"`
	GiftsSubscriptionName      *string                             `json:"giftsSubscriptionName,omitempty"`
	GiftsDuration              *int                                `json:"gifts_duration,omitempty"`
	GiftsItemID                *int                                `json:"gifts_item_id,omitempty"`
	GiftsService               *string                             `json:"gifts_service,omitempty"`
	GiftsType                  *string                             `json:"gifts_type,omitempty"`
	Guarantee                  interface{}                         `json:"guarantee,omitempty"`
	HasPendingAutoBuy          *bool                               `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                  *bool                               `json:"isIgnored,omitempty"`
	IsSticky                   *int                                `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                             `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                             `json:"item_domain,omitempty"`
	ItemID                     *int                                `json:"item_id,omitempty"`
	ItemOrigin                 *string                             `json:"item_origin,omitempty"`
	ItemState                  *string                             `json:"item_state,omitempty"`
	NoteText                   interface{}                         `json:"note_text,omitempty"`
	Nsb                        *int                                `json:"nsb,omitempty"`
	Price                      *int                                `json:"price,omitempty"`
	PriceWithSellerFee         *float64                            `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                             `json:"price_currency,omitempty"`
	PublishedDate              *int                                `json:"published_date,omitempty"`
	RefreshedDate              *int                                `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                             `json:"resale_item_origin,omitempty"`
	RubPrice                   *int                                `json:"rub_price,omitempty"`
	Seller                     *GiftsResponseItemsItemSeller       `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                               `json:"showGetEmailCodeButton,omitempty"`
	Tags                       interface{}                         `json:"tags,omitempty"`
	Title                      *string                             `json:"title,omitempty"`
	TitleEn                    *string                             `json:"title_en,omitempty"`
	UpdateStatDate             *int                                `json:"update_stat_date,omitempty"`
	ViewCount                  *int                                `json:"view_count,omitempty"`
}

type GiftsResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type GiftsResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     interface{} `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type GiftsResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type HistoryParams struct {
	// Type of operation.
	Type_ *PaymentsType `json:"type,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// Currency.
	Currency *PaymentsCurrency `json:"currency,omitempty"`
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Id of the operation from which the result begins.
	OperationIDLt *int `json:"operation_id_lt,omitempty"`
	// Username of user, which receive money from you.
	Receiver *string `json:"receiver,omitempty"`
	// Username of user, which sent money to you.
	Sender *string `json:"sender,omitempty"`
	// Returns payments that are done via API.
	IsAPI *bool `json:"is_api,omitempty"`
	// Start date of operation (RFC 3339 date format).
	StartDate *string `json:"startDate,omitempty"`
	// End date of operation (RFC 3339 date format).
	EndDate *string `json:"endDate,omitempty"`
	// Wallet, which used for money payouts.
	Wallet *string `json:"wallet,omitempty"`
	// Comment for money transfers.
	Comment *string `json:"comment,omitempty"`
	// Display hold operations.
	IsHold *bool `json:"is_hold,omitempty"`
	// Display payment stats for selected period (outgoing value, incoming value).
	ShowPaymentStats *bool `json:"show_payment_stats,omitempty"`
}

type HistoryResponse struct {
	FilterDatesDefault *bool                         `json:"filterDatesDefault,omitempty"`
	HasNextPage        *bool                         `json:"hasNextPage,omitempty"`
	Input              *HistoryResponseInput         `json:"input,omitempty"`
	LastOperationId    *int                          `json:"lastOperationId,omitempty"`
	NextPageHref       *string                       `json:"nextPageHref,omitempty"`
	Page               *int                          `json:"page,omitempty"`
	PageNavLink        *string                       `json:"pageNavLink,omitempty"`
	PageNavParams      *HistoryResponsePageNavParams `json:"pageNavParams,omitempty"`
	PaymentStats       interface{}                   `json:"paymentStats,omitempty"`
	Payments           map[string]interface{}        `json:"payments,omitempty"`
	PerPage            *string                       `json:"perPage,omitempty"`
	PeriodLabel        *string                       `json:"periodLabel,omitempty"`
	PeriodLabelPhrase  *string                       `json:"periodLabelPhrase,omitempty"`
	SystemInfo         *RespSystemInfo               `json:"system_info,omitempty"`
}

type HistoryResponseInput struct {
	CategoryID    *int    `json:"category_id,omitempty"`
	Comment       *string `json:"comment,omitempty"`
	Currency      *string `json:"currency,omitempty"`
	EndDate       *string `json:"endDate,omitempty"`
	IsHold        *bool   `json:"is_hold,omitempty"`
	OperationIDLt *int    `json:"operation_id_lt,omitempty"`
	Page          *int    `json:"page,omitempty"`
	PeriodLabel   *string `json:"period_label,omitempty"`
	Pmax          *string `json:"pmax,omitempty"`
	Pmin          *string `json:"pmin,omitempty"`
	Receiver      *string `json:"receiver,omitempty"`
	Sender        *string `json:"sender,omitempty"`
	StartDate     *string `json:"startDate,omitempty"`
	Type_         *string `json:"type,omitempty"`
	UserID        *int    `json:"user_id,omitempty"`
	Wallet        *string `json:"wallet,omitempty"`
}

type HistoryResponsePageNavParams struct {
	EndDate   *string `json:"endDate,omitempty"`
	StartDate *string `json:"startDate,omitempty"`
	Type_     *string `json:"type,omitempty"`
}

type HistoryResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type HytaleParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// List of allowed editions.
	Edition []string `json:"edition[],omitempty"`
	// Minimum number of profiles with game.
	ProfilesMin *int `json:"profiles_min,omitempty"`
	// Maximum number of profiles with game.
	ProfilesMax *int `json:"profiles_max,omitempty"`
}

type HytaleResponse struct {
	CacheTTL        *int                      `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                     `json:"hasNextPage,omitempty"`
	Items           []HytaleResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                      `json:"lastModified,omitempty"`
	Page            *int                      `json:"page,omitempty"`
	PerPage         *int                      `json:"perPage,omitempty"`
	SearchUrl       *string                   `json:"searchUrl,omitempty"`
	ServerTime      *int                      `json:"serverTime,omitempty"`
	StickyItems     []interface{}             `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo           `json:"system_info,omitempty"`
	TotalItems      *int                      `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}               `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                     `json:"wasCached,omitempty"`
}

type HytaleResponseItemsItem struct {
	AllowAskDiscount        *int                                   `json:"allow_ask_discount,omitempty"`
	AutoBumpPeriod          *int                                   `json:"auto_bump_period,omitempty"`
	Buyer                   interface{}                            `json:"buyer,omitempty"`
	CanBumpItem             *bool                                  `json:"canBumpItem,omitempty"`
	CanBuyItem              *bool                                  `json:"canBuyItem,omitempty"`
	CanChangeEmailPassword  *bool                                  `json:"canChangeEmailPassword,omitempty"`
	CanChangePassword       *bool                                  `json:"canChangePassword,omitempty"`
	CanCloseItem            *bool                                  `json:"canCloseItem,omitempty"`
	CanDeleteItem           *bool                                  `json:"canDeleteItem,omitempty"`
	CanEditItem             *bool                                  `json:"canEditItem,omitempty"`
	CanManagePublicTag      *bool                                  `json:"canManagePublicTag,omitempty"`
	CanNotBumpItemReason    *string                                `json:"canNotBumpItemReason,omitempty"`
	CanOpenItem             *bool                                  `json:"canOpenItem,omitempty"`
	CanReportItem           *bool                                  `json:"canReportItem,omitempty"`
	CanResellItem           *bool                                  `json:"canResellItem,omitempty"`
	CanStickItem            *bool                                  `json:"canStickItem,omitempty"`
	CanUnstickItem          *bool                                  `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats      *bool                                  `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount      *bool                                  `json:"canValidateAccount,omitempty"`
	CanViewAccountLink      *bool                                  `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData   *bool                                  `json:"canViewEmailLoginData,omitempty"`
	CanViewItemViews        *bool                                  `json:"canViewItemViews,omitempty"`
	CanViewLoginData        *bool                                  `json:"canViewLoginData,omitempty"`
	CanViewTempEmail        *bool                                  `json:"canViewTempEmail,omitempty"`
	Category                *HytaleResponseItemsItemCategory       `json:"category,omitempty"`
	CategoryID              *int                                   `json:"category_id,omitempty"`
	CopyFormatData          *HytaleResponseItemsItemCopyFormatData `json:"copyFormatData,omitempty"`
	Description             *string                                `json:"description,omitempty"`
	DescriptionEnHtml       *string                                `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain      *string                                `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml         *string                                `json:"descriptionHtml,omitempty"`
	DescriptionPlain        *string                                `json:"descriptionPlain,omitempty"`
	DescriptionEn           *string                                `json:"description_en,omitempty"`
	Discount                *bool                                  `json:"discount,omitempty"`
	EditDate                *int                                   `json:"edit_date,omitempty"`
	EmailLoginUrl           *string                                `json:"emailLoginUrl,omitempty"`
	EmailProvider           *string                                `json:"email_provider,omitempty"`
	EmailType               *string                                `json:"email_type,omitempty"`
	ExtendedGuarantee       *int                                   `json:"extended_guarantee,omitempty"`
	FeedbackData            interface{}                            `json:"feedback_data,omitempty"`
	Guarantee               interface{}                            `json:"guarantee,omitempty"`
	HasPendingAutoBuy       *bool                                  `json:"hasPendingAutoBuy,omitempty"`
	HytaleEdition           *string                                `json:"hytale_edition,omitempty"`
	HytaleItemID            *int                                   `json:"hytale_item_id,omitempty"`
	HytaleProfiles          *int                                   `json:"hytale_profiles,omitempty"`
	ImagePreviewLinks       []interface{}                          `json:"imagePreviewLinks,omitempty"`
	IsIgnored               *bool                                  `json:"isIgnored,omitempty"`
	IsPersonalAccount       *bool                                  `json:"isPersonalAccount,omitempty"`
	IsSticky                *int                                   `json:"is_sticky,omitempty"`
	ItemOriginPhrase        *string                                `json:"itemOriginPhrase,omitempty"`
	ItemDomain              *string                                `json:"item_domain,omitempty"`
	ItemID                  *int                                   `json:"item_id,omitempty"`
	ItemOrigin              *string                                `json:"item_origin,omitempty"`
	ItemState               *string                                `json:"item_state,omitempty"`
	MaxDiscountPercent      *int                                   `json:"max_discount_percent,omitempty"`
	NoteText                interface{}                            `json:"note_text,omitempty"`
	Nsb                     *int                                   `json:"nsb,omitempty"`
	PendingDeletionDate     *int                                   `json:"pending_deletion_date,omitempty"`
	Price                   *int                                   `json:"price,omitempty"`
	PriceWithSellerFee      *float64                               `json:"priceWithSellerFee,omitempty"`
	PriceWithSellerFeeLabel *string                                `json:"priceWithSellerFeeLabel,omitempty"`
	PriceCurrency           *string                                `json:"price_currency,omitempty"`
	PublicTag               interface{}                            `json:"public_tag,omitempty"`
	PublishedDate           *int                                   `json:"published_date,omitempty"`
	RefreshedDate           *int                                   `json:"refreshed_date,omitempty"`
	ResaleItemOrigin        *string                                `json:"resale_item_origin,omitempty"`
	RubPrice                *int                                   `json:"rub_price,omitempty"`
	Seller                  *HytaleResponseItemsItemSeller         `json:"seller,omitempty"`
	ShowGetEmailCodeButton  *bool                                  `json:"showGetEmailCodeButton,omitempty"`
	Tags                    interface{}                            `json:"tags,omitempty"`
	Title                   *string                                `json:"title,omitempty"`
	TitleEn                 *string                                `json:"title_en,omitempty"`
	UniqueKeyExists         *bool                                  `json:"uniqueKeyExists,omitempty"`
	UpdateStatDate          *int                                   `json:"update_stat_date,omitempty"`
	ViewCount               *int                                   `json:"view_count,omitempty"`
}

type HytaleResponseItemsItemCategory struct {
	CategoryID    *int    `json:"category_id,omitempty"`
	CategoryName  *string `json:"category_name,omitempty"`
	CategoryTitle *string `json:"category_title,omitempty"`
	CategoryURL   *string `json:"category_url,omitempty"`
}

type HytaleResponseItemsItemCopyFormatData struct {
	TitleLink *string `json:"title_link,omitempty"`
}

type HytaleResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     interface{} `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type HytaleResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ImageResponse struct {
	Base64     *string         `json:"base64,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type ImageResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type InstagramParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// Has linked mobile.
	Tel *CategorySearchTel `json:"tel,omitempty"`
	// List of allowed countries.
	Country []string `json:"country[],omitempty"`
	// List of disallowed countries.
	NotCountry []string `json:"not_country[],omitempty"`
	// Login by cookies.
	Cookies *CategorySearchCookies `json:"cookies,omitempty"`
	// Login without cookies.
	LoginWithoutCookies *CategorySearchLoginWithoutCookies `json:"login_without_cookies,omitempty"`
	// Minimum number of followers.
	FollowersMin *int `json:"followers_min,omitempty"`
	// Maximum number of followers.
	FollowersMax *int `json:"followers_max,omitempty"`
	// Minimum number of posts.
	PostMin *int `json:"post_min,omitempty"`
	// Maximum number of posts.
	PostMax *int `json:"post_max,omitempty"`
	// How old is the account.
	Reg *int `json:"reg,omitempty"`
	// In what notation is time measured.
	RegPeriod *CategorySearchRegPeriod `json:"reg_period,omitempty"`
}

type InstagramResponse struct {
	CacheTTL        *int                         `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                        `json:"hasNextPage,omitempty"`
	Items           []InstagramResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                         `json:"lastModified,omitempty"`
	Page            *int                         `json:"page,omitempty"`
	PerPage         *int                         `json:"perPage,omitempty"`
	SearchUrl       *string                      `json:"searchUrl,omitempty"`
	ServerTime      *int                         `json:"serverTime,omitempty"`
	StickyItems     []interface{}                `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo              `json:"system_info,omitempty"`
	TotalItems      *int                         `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}                  `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                        `json:"wasCached,omitempty"`
}

type InstagramResponseItemsItem struct {
	AccountLink                  *string                                      `json:"accountLink,omitempty"`
	AccountLinks                 []InstagramResponseItemsItemAccountLinksItem `json:"accountLinks,omitempty"`
	AllowAskDiscount             *int                                         `json:"allow_ask_discount,omitempty"`
	BumpSettings                 *InstagramResponseItemsItemBumpSettings      `json:"bumpSettings,omitempty"`
	CanBumpItem                  *bool                                        `json:"canBumpItem,omitempty"`
	CanBuyItem                   *bool                                        `json:"canBuyItem,omitempty"`
	CanChangePassword            *bool                                        `json:"canChangePassword,omitempty"`
	CanCloseItem                 *bool                                        `json:"canCloseItem,omitempty"`
	CanDeleteItem                *bool                                        `json:"canDeleteItem,omitempty"`
	CanEditItem                  *bool                                        `json:"canEditItem,omitempty"`
	CanOpenItem                  *bool                                        `json:"canOpenItem,omitempty"`
	CanReportItem                *bool                                        `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase   *bool                                        `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem                 *bool                                        `json:"canStickItem,omitempty"`
	CanUnstickItem               *bool                                        `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats           *bool                                        `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount           *bool                                        `json:"canValidateAccount,omitempty"`
	CanViewAccountLink           *bool                                        `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData        *bool                                        `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData             *bool                                        `json:"canViewLoginData,omitempty"`
	CategoryID                   *int                                         `json:"category_id,omitempty"`
	Description                  *string                                      `json:"description,omitempty"`
	DescriptionEnHtml            *string                                      `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain           *string                                      `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml              *string                                      `json:"descriptionHtml,omitempty"`
	DescriptionPlain             *string                                      `json:"descriptionPlain,omitempty"`
	DescriptionEn                *string                                      `json:"description_en,omitempty"`
	EditDate                     *int                                         `json:"edit_date,omitempty"`
	EmailLoginUrl                *string                                      `json:"emailLoginUrl,omitempty"`
	EmailProvider                *string                                      `json:"email_provider,omitempty"`
	EmailType                    *string                                      `json:"email_type,omitempty"`
	ExtendedGuarantee            *int                                         `json:"extended_guarantee,omitempty"`
	FeedbackData                 interface{}                                  `json:"feedback_data,omitempty"`
	Guarantee                    interface{}                                  `json:"guarantee,omitempty"`
	HasPendingAutoBuy            *bool                                        `json:"hasPendingAutoBuy,omitempty"`
	InstagramCountry             *string                                      `json:"instagram_country,omitempty"`
	InstagramFollowCount         *int                                         `json:"instagram_follow_count,omitempty"`
	InstagramFollowerCount       *int                                         `json:"instagram_follower_count,omitempty"`
	InstagramHasCookies          *int                                         `json:"instagram_has_cookies,omitempty"`
	InstagramID                  interface{}                                  `json:"instagram_id,omitempty"`
	InstagramItemID              *int                                         `json:"instagram_item_id,omitempty"`
	InstagramLoginWithoutCookies *int                                         `json:"instagram_login_without_cookies,omitempty"`
	InstagramMobile              *int                                         `json:"instagram_mobile,omitempty"`
	InstagramPostCount           *int                                         `json:"instagram_post_count,omitempty"`
	InstagramRegisterDate        *int                                         `json:"instagram_register_date,omitempty"`
	InstagramUsername            *string                                      `json:"instagram_username,omitempty"`
	IsIgnored                    *bool                                        `json:"isIgnored,omitempty"`
	IsSticky                     *int                                         `json:"is_sticky,omitempty"`
	ItemOriginPhrase             *string                                      `json:"itemOriginPhrase,omitempty"`
	ItemDomain                   *string                                      `json:"item_domain,omitempty"`
	ItemID                       *int                                         `json:"item_id,omitempty"`
	ItemOrigin                   *string                                      `json:"item_origin,omitempty"`
	ItemState                    *string                                      `json:"item_state,omitempty"`
	NoteText                     interface{}                                  `json:"note_text,omitempty"`
	Nsb                          *int                                         `json:"nsb,omitempty"`
	Price                        *int                                         `json:"price,omitempty"`
	PriceWithSellerFee           *float64                                     `json:"priceWithSellerFee,omitempty"`
	PriceCurrency                *string                                      `json:"price_currency,omitempty"`
	PublishedDate                *int                                         `json:"published_date,omitempty"`
	RefreshedDate                *int                                         `json:"refreshed_date,omitempty"`
	ResaleItemOrigin             *string                                      `json:"resale_item_origin,omitempty"`
	RubPrice                     *int                                         `json:"rub_price,omitempty"`
	Seller                       *InstagramResponseItemsItemSeller            `json:"seller,omitempty"`
	ShowGetEmailCodeButton       *bool                                        `json:"showGetEmailCodeButton,omitempty"`
	Tags                         interface{}                                  `json:"tags,omitempty"`
	Title                        *string                                      `json:"title,omitempty"`
	TitleEn                      *string                                      `json:"title_en,omitempty"`
	UpdateStatDate               *int                                         `json:"update_stat_date,omitempty"`
	ViewCount                    *int                                         `json:"view_count,omitempty"`
}

type InstagramResponseItemsItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type InstagramResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type InstagramResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     interface{} `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type InstagramResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type InvoiceModel struct {
	AdditionalData *string `json:"additional_data,omitempty"`
	Amount         *int    `json:"amount,omitempty"`
	Comment        *string `json:"comment,omitempty"`
	ExpiresAt      *int    `json:"expires_at,omitempty"`
	InvoiceDate    *int    `json:"invoice_date,omitempty"`
	InvoiceID      *int    `json:"invoice_id,omitempty"`
	IsTest         *bool   `json:"is_test,omitempty"`
	MerchantID     *int    `json:"merchant_id,omitempty"`
	PaidDate       *int    `json:"paid_date,omitempty"`
	PayerUserID    *int    `json:"payer_user_id,omitempty"`
	PaymentID      *string `json:"payment_id,omitempty"`
	ResendAttempts *int    `json:"resend_attempts,omitempty"`
	Status         *string `json:"status,omitempty"`
	URL            *string `json:"url,omitempty"`
	URLCallback    *string `json:"url_callback,omitempty"`
	URLSuccess     *string `json:"url_success,omitempty"`
	UserID         *int    `json:"user_id,omitempty"`
}

type ItemFromListModel struct {
	AllowAskDiscount           *int                           `json:"allow_ask_discount,omitempty"`
	BumpSettings               *ItemFromListModelBumpSettings `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                          `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                          `json:"canBuyItem,omitempty"`
	CanCloseItem               *bool                          `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                          `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                          `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                          `json:"canOpenItem,omitempty"`
	CanResellItemAfterPurchase *bool                          `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                          `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                          `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                          `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                          `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                          `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                          `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                          `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                           `json:"category_id,omitempty"`
	Description                *string                        `json:"description,omitempty"`
	DescriptionEn              *string                        `json:"description_en,omitempty"`
	DescriptionHTML            *string                        `json:"description_html,omitempty"`
	DescriptionHTMLEn          *string                        `json:"description_html_en,omitempty"`
	ExtendedGuarantee          *int                           `json:"extended_guarantee,omitempty"`
	Guarantee                  interface{}                    `json:"guarantee,omitempty"`
	IsIgnored                  *int                           `json:"isIgnored,omitempty"`
	IsSticky                   *int                           `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                        `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                        `json:"item_domain,omitempty"`
	ItemID                     *int                           `json:"item_id,omitempty"`
	ItemOrigin                 *string                        `json:"item_origin,omitempty"`
	ItemState                  *string                        `json:"item_state,omitempty"`
	NoteText                   *string                        `json:"note_text,omitempty"`
	Nsb                        *int                           `json:"nsb,omitempty"`
	Price                      *int                           `json:"price,omitempty"`
	PriceCurrency              *string                        `json:"price_currency,omitempty"`
	PublishedDate              *int                           `json:"published_date,omitempty"`
	RefreshedDate              *int                           `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                        `json:"resale_item_origin,omitempty"`
	RubPrice                   *int                           `json:"rub_price,omitempty"`
	Seller                     *ItemFromListModelSeller       `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                          `json:"showGetEmailCodeButton,omitempty"`
	Tags                       interface{}                    `json:"tags,omitempty"`
	Title                      *string                        `json:"title,omitempty"`
	TitleEn                    *string                        `json:"title_en,omitempty"`
	UpdateStatDate             *int                           `json:"update_stat_date,omitempty"`
	ViewCount                  *int                           `json:"view_count,omitempty"`
}

type ItemFromListModelBumpSettings struct {
	CanBumpItem         *bool   `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool   `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         *string `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    *string `json:"shortErrorPhrase,omitempty"`
}

type ItemFromListModelSeller struct {
	ActiveItemCount     *int        `json:"active_item_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type ItemListModel struct {
	HasNextPage     *bool                          `json:"hasNextPage,omitempty"`
	Items           []ItemListModelItemsItem       `json:"items,omitempty"`
	Page            *int                           `json:"page,omitempty"`
	PerPage         *int                           `json:"perPage,omitempty"`
	SearchUrl       *string                        `json:"searchUrl,omitempty"`
	StickyItems     []ItemListModelStickyItemsItem `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo                `json:"system_info,omitempty"`
	TotalItems      *int                           `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}                    `json:"totalItemsPrice,omitempty"`
}

type ItemListModelItemsItem struct {
	AllowAskDiscount           *int                                `json:"allow_ask_discount,omitempty"`
	BumpSettings               *ItemListModelItemsItemBumpSettings `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                               `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                               `json:"canBuyItem,omitempty"`
	CanCloseItem               *bool                               `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                               `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                               `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                               `json:"canOpenItem,omitempty"`
	CanResellItemAfterPurchase *bool                               `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                               `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                               `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                               `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                               `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                               `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                               `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                               `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                `json:"category_id,omitempty"`
	Description                *string                             `json:"description,omitempty"`
	DescriptionEn              *string                             `json:"description_en,omitempty"`
	DescriptionHTML            *string                             `json:"description_html,omitempty"`
	DescriptionHTMLEn          *string                             `json:"description_html_en,omitempty"`
	ExtendedGuarantee          *int                                `json:"extended_guarantee,omitempty"`
	Guarantee                  interface{}                         `json:"guarantee,omitempty"`
	IsIgnored                  *int                                `json:"isIgnored,omitempty"`
	IsSticky                   *int                                `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                             `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                             `json:"item_domain,omitempty"`
	ItemID                     *int                                `json:"item_id,omitempty"`
	ItemOrigin                 *string                             `json:"item_origin,omitempty"`
	ItemState                  *string                             `json:"item_state,omitempty"`
	NoteText                   *string                             `json:"note_text,omitempty"`
	Nsb                        *int                                `json:"nsb,omitempty"`
	Price                      *int                                `json:"price,omitempty"`
	PriceCurrency              *string                             `json:"price_currency,omitempty"`
	PublishedDate              *int                                `json:"published_date,omitempty"`
	RefreshedDate              *int                                `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                             `json:"resale_item_origin,omitempty"`
	RubPrice                   *int                                `json:"rub_price,omitempty"`
	Seller                     *ItemListModelItemsItemSeller       `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                               `json:"showGetEmailCodeButton,omitempty"`
	Tags                       interface{}                         `json:"tags,omitempty"`
	Title                      *string                             `json:"title,omitempty"`
	TitleEn                    *string                             `json:"title_en,omitempty"`
	UpdateStatDate             *int                                `json:"update_stat_date,omitempty"`
	ViewCount                  *int                                `json:"view_count,omitempty"`
}

type ItemListModelItemsItemBumpSettings struct {
	CanBumpItem         *bool   `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool   `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         *string `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    *string `json:"shortErrorPhrase,omitempty"`
}

type ItemListModelItemsItemSeller struct {
	ActiveItemCount     *int        `json:"active_item_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type ItemListModelStickyItemsItem struct {
	AllowAskDiscount           *int                                      `json:"allow_ask_discount,omitempty"`
	BumpSettings               *ItemListModelStickyItemsItemBumpSettings `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                                     `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                                     `json:"canBuyItem,omitempty"`
	CanCloseItem               *bool                                     `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                                     `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                                     `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                                     `json:"canOpenItem,omitempty"`
	CanResellItemAfterPurchase *bool                                     `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                                     `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                                     `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                                     `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                                     `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                                     `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                                     `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                                     `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                      `json:"category_id,omitempty"`
	Description                *string                                   `json:"description,omitempty"`
	DescriptionEn              *string                                   `json:"description_en,omitempty"`
	DescriptionHTML            *string                                   `json:"description_html,omitempty"`
	DescriptionHTMLEn          *string                                   `json:"description_html_en,omitempty"`
	ExtendedGuarantee          *int                                      `json:"extended_guarantee,omitempty"`
	Guarantee                  interface{}                               `json:"guarantee,omitempty"`
	IsIgnored                  *int                                      `json:"isIgnored,omitempty"`
	IsSticky                   *int                                      `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                                   `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                                   `json:"item_domain,omitempty"`
	ItemID                     *int                                      `json:"item_id,omitempty"`
	ItemOrigin                 *string                                   `json:"item_origin,omitempty"`
	ItemState                  *string                                   `json:"item_state,omitempty"`
	NoteText                   *string                                   `json:"note_text,omitempty"`
	Nsb                        *int                                      `json:"nsb,omitempty"`
	Price                      *int                                      `json:"price,omitempty"`
	PriceCurrency              *string                                   `json:"price_currency,omitempty"`
	PublishedDate              *int                                      `json:"published_date,omitempty"`
	RefreshedDate              *int                                      `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                                   `json:"resale_item_origin,omitempty"`
	RubPrice                   *int                                      `json:"rub_price,omitempty"`
	Seller                     *ItemListModelStickyItemsItemSeller       `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                                     `json:"showGetEmailCodeButton,omitempty"`
	Tags                       interface{}                               `json:"tags,omitempty"`
	Title                      *string                                   `json:"title,omitempty"`
	TitleEn                    *string                                   `json:"title_en,omitempty"`
	UpdateStatDate             *int                                      `json:"update_stat_date,omitempty"`
	ViewCount                  *int                                      `json:"view_count,omitempty"`
}

type ItemListModelStickyItemsItemBumpSettings struct {
	CanBumpItem         *bool   `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool   `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         *string `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    *string `json:"shortErrorPhrase,omitempty"`
}

type ItemListModelStickyItemsItemSeller struct {
	ActiveItemCount     *int        `json:"active_item_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type ItemListModelSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ItemModel struct {
	AccountLink                       *string                     `json:"accountLink,omitempty"`
	AccountLinks                      []ItemModelAccountLinksItem `json:"accountLinks,omitempty"`
	AccountLastActivity               *int                        `json:"account_last_activity,omitempty"`
	AiPrice                           *int                        `json:"aiPrice,omitempty"`
	AiPriceCheckDate                  *int                        `json:"aiPriceCheckDate,omitempty"`
	AllowAskDiscount                  *int                        `json:"allow_ask_discount,omitempty"`
	AutoBuyPrice                      *int                        `json:"autoBuyPrice,omitempty"`
	AutoBuyPriceCheckDate             *int                        `json:"autoBuyPriceCheckDate,omitempty"`
	BumpSettings                      *ItemModelBumpSettings      `json:"bumpSettings,omitempty"`
	Buyer                             *ItemModelBuyer             `json:"buyer,omitempty"`
	BuyerAvatarDate                   *int                        `json:"buyer_avatar_date,omitempty"`
	BuyerDisplayIconGroupID           *int                        `json:"buyer_display_icon_group_id,omitempty"`
	BuyerUniqBanner                   *string                     `json:"buyer_uniq_banner,omitempty"`
	BuyerUserGroupID                  *int                        `json:"buyer_user_group_id,omitempty"`
	CanAskDiscount                    *bool                       `json:"canAskDiscount,omitempty"`
	CanChangeEmailPassword            *bool                       `json:"canChangeEmailPassword,omitempty"`
	CanChangePassword                 *bool                       `json:"canChangePassword,omitempty"`
	CanCheckAiPrice                   *bool                       `json:"canCheckAiPrice,omitempty"`
	CanCheckAutoBuyPrice              *bool                       `json:"canCheckAutoBuyPrice,omitempty"`
	CanCheckGuarantee                 *bool                       `json:"canCheckGuarantee,omitempty"`
	CanReportItem                     *bool                       `json:"canReportItem,omitempty"`
	CanResellItem                     *bool                       `json:"canResellItem,omitempty"`
	CanResellItemAfterPurchase        *bool                       `json:"canResellItemAfterPurchase,omitempty"`
	CanShareItem                      *bool                       `json:"canShareItem,omitempty"`
	CanUpdateItemStats                *bool                       `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount                *bool                       `json:"canValidateAccount,omitempty"`
	CanViewAccountLink                *bool                       `json:"canViewAccountLink,omitempty"`
	CanViewAccountLoginAndTempEmail   *bool                       `json:"canViewAccountLoginAndTempEmail,omitempty"`
	CanViewEmailLoginData             *bool                       `json:"canViewEmailLoginData,omitempty"`
	CanViewItemViews                  *bool                       `json:"canViewItemViews,omitempty"`
	CanViewLoginData                  *bool                       `json:"canViewLoginData,omitempty"`
	CartPrice                         interface{}                 `json:"cart_price,omitempty"`
	CategoryID                        *int                        `json:"category_id,omitempty"`
	ContentID                         interface{}                 `json:"content_id,omitempty"`
	ContentType                       interface{}                 `json:"content_type,omitempty"`
	CopyFormatData                    *ItemModelCopyFormatData    `json:"copyFormatData,omitempty"`
	CustomFields                      *ItemModelCustomFields      `json:"customFields,omitempty"`
	DeleteDate                        *int                        `json:"delete_date,omitempty"`
	DeleteReason                      *string                     `json:"delete_reason,omitempty"`
	DeleteUserID                      *int                        `json:"delete_user_id,omitempty"`
	DeleteUsername                    *string                     `json:"delete_username,omitempty"`
	Deposit                           *int                        `json:"deposit,omitempty"`
	Description                       *string                     `json:"description,omitempty"`
	DescriptionEnHtml                 *string                     `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain                *string                     `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml                   *string                     `json:"descriptionHtml,omitempty"`
	DescriptionPlain                  *string                     `json:"descriptionPlain,omitempty"`
	DescriptionEn                     *string                     `json:"description_en,omitempty"`
	EditDate                          *int                        `json:"edit_date,omitempty"`
	EmailProvider                     *string                     `json:"email_provider,omitempty"`
	EmailType                         *string                     `json:"email_type,omitempty"`
	ExtendedGuarantee                 *int                        `json:"extended_guarantee,omitempty"`
	ExternalAuth                      []interface{}               `json:"externalAuth,omitempty"`
	ExtraPrices                       []ItemModelExtraPricesItem  `json:"extraPrices,omitempty"`
	FeedbackData                      interface{}                 `json:"feedback_data,omitempty"`
	GetEmailCodeDisplayLogin          interface{}                 `json:"getEmailCodeDisplayLogin,omitempty"`
	Guarantee                         interface{}                 `json:"guarantee,omitempty"`
	ImagePreviewLinks                 []string                    `json:"imagePreviewLinks,omitempty"`
	InCart                            interface{}                 `json:"in_cart,omitempty"`
	Information                       *string                     `json:"information,omitempty"`
	InformationEn                     *string                     `json:"information_en,omitempty"`
	IsBirthdayToday                   *bool                       `json:"isBirthdayToday,omitempty"`
	IsIgnored                         *bool                       `json:"isIgnored,omitempty"`
	IsPersonalAccount                 *bool                       `json:"isPersonalAccount,omitempty"`
	IsSmallExf                        *bool                       `json:"isSmallExf,omitempty"`
	IsTrusted                         *bool                       `json:"isTrusted,omitempty"`
	IsFave                            interface{}                 `json:"is_fave,omitempty"`
	IsSticky                          *int                        `json:"is_sticky,omitempty"`
	ItemOriginPhrase                  *string                     `json:"itemOriginPhrase,omitempty"`
	ItemDomain                        *string                     `json:"item_domain,omitempty"`
	ItemID                            *int                        `json:"item_id,omitempty"`
	ItemOrigin                        *string                     `json:"item_origin,omitempty"`
	ItemState                         *string                     `json:"item_state,omitempty"`
	Login                             *string                     `json:"login,omitempty"`
	LoginData                         *ItemModelLoginData         `json:"loginData,omitempty"`
	MarketCustomTitle                 *string                     `json:"market_custom_title,omitempty"`
	MaxDiscountPercent                *int                        `json:"max_discount_percent,omitempty"`
	NeedToRequireVideoToViewLoginData *bool                       `json:"needToRequireVideoToViewLoginData,omitempty"`
	NoteText                          *string                     `json:"note_text,omitempty"`
	Nsb                               *int                        `json:"nsb,omitempty"`
	PendingDeletionDate               *int                        `json:"pending_deletion_date,omitempty"`
	Price                             *int                        `json:"price,omitempty"`
	PriceWithSellerFee                *float64                    `json:"priceWithSellerFee,omitempty"`
	PriceWithSellerFeeLabel           *string                     `json:"priceWithSellerFeeLabel,omitempty"`
	PriceCurrency                     *string                     `json:"price_currency,omitempty"`
	PublishedDate                     *int                        `json:"published_date,omitempty"`
	RefreshedDate                     *int                        `json:"refreshed_date,omitempty"`
	ResaleItemOrigin                  *string                     `json:"resale_item_origin,omitempty"`
	RubPrice                          *int                        `json:"rub_price,omitempty"`
	Seller                            *ItemModelSeller            `json:"seller,omitempty"`
	ShowGetEmailCodeButton            *bool                       `json:"showGetEmailCodeButton,omitempty"`
	Tags                              interface{}                 `json:"tags,omitempty"`
	TempEmail                         *string                     `json:"temp_email,omitempty"`
	Title                             *string                     `json:"title,omitempty"`
	TitleEn                           *string                     `json:"title_en,omitempty"`
	UniqueKeyExists                   *bool                       `json:"uniqueKeyExists,omitempty"`
	UpdateStatDate                    *int                        `json:"update_stat_date,omitempty"`
	UserAllowAskDiscount              *int                        `json:"user_allow_ask_discount,omitempty"`
	ViewCount                         *int                        `json:"view_count,omitempty"`
	VisitorIsAuthor                   *bool                       `json:"visitorIsAuthor,omitempty"`
}

type ItemModelAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type ItemModelBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	NextAllowedBumpDate interface{} `json:"nextAllowedBumpDate,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type ItemModelBuyer struct {
	DisplayIconGroupID  *int    `json:"display_icon_group_id,omitempty"`
	DisplayStyleGroupID *int    `json:"display_style_group_id,omitempty"`
	IsBanned            *int    `json:"is_banned,omitempty"`
	OperationDate       *int    `json:"operation_date,omitempty"`
	UniqBanner          *string `json:"uniq_banner,omitempty"`
	UniqUsernameCSS     *string `json:"uniq_username_css,omitempty"`
	UserGroupID         *int    `json:"user_group_id,omitempty"`
	UserID              *int    `json:"user_id,omitempty"`
	Username            *string `json:"username,omitempty"`
	VisitorIsBuyer      *bool   `json:"visitorIsBuyer,omitempty"`
}

type ItemModelCopyFormatData struct {
	Full      *string `json:"full,omitempty"`
	LoginData *string `json:"login_data,omitempty"`
	TitleLink *string `json:"title_link,omitempty"`
}

type ItemModelCustomFields struct {
	Field4         *string       `json:"_4,omitempty"`
	AllowSelfUnban []interface{} `json:"allowSelfUnban,omitempty"`
	BanReason      *string       `json:"ban_reason,omitempty"`
	Discord        *string       `json:"discord,omitempty"`
	Github         *string       `json:"github,omitempty"`
	Jabber         *string       `json:"jabber,omitempty"`
	LztUnbanAmount *string       `json:"lztUnbanAmount,omitempty"`
	Steam          *string       `json:"steam,omitempty"`
	Telegram       *string       `json:"telegram,omitempty"`
	Vk             *string       `json:"vk,omitempty"`
}

type ItemModelExtraPricesItem struct {
	Currency   *string  `json:"currency,omitempty"`
	Price      *string  `json:"price,omitempty"`
	PriceValue *float64 `json:"priceValue,omitempty"`
}

type ItemModelGuarantee struct {
	Active                *bool   `json:"active,omitempty"`
	Cancelled             *bool   `json:"cancelled,omitempty"`
	CancelledReason       *string `json:"cancelledReason,omitempty"`
	CancelledReasonPhrase *string `json:"cancelledReasonPhrase,omitempty"`
	Class                 *string `json:"class,omitempty"`
	Duration              *int    `json:"duration,omitempty"`
	DurationPhrase        *string `json:"durationPhrase,omitempty"`
	EndDate               *int    `json:"endDate,omitempty"`
	RemainingTime         *int    `json:"remainingTime,omitempty"`
	RemainingTimePhrase   *string `json:"remainingTimePhrase,omitempty"`
}

type ItemModelLoginData struct {
	EncodedOldPassword interface{} `json:"encodedOldPassword,omitempty"`
	EncodedPassword    *string     `json:"encodedPassword,omitempty"`
	EncodedRaw         *string     `json:"encodedRaw,omitempty"`
	Login              *string     `json:"login,omitempty"`
	OldPassword        *string     `json:"oldPassword,omitempty"`
	Password           *string     `json:"password,omitempty"`
	Raw                *string     `json:"raw,omitempty"`
}

type ItemModelSeller struct {
	ActiveItemsCount      *int                     `json:"active_items_count,omitempty"`
	AvatarDate            *int                     `json:"avatar_date,omitempty"`
	Contacts              *ItemModelSellerContacts `json:"contacts,omitempty"`
	DisplayStyleGroupID   *int                     `json:"display_style_group_id,omitempty"`
	EffectiveLastActivity *int                     `json:"effective_last_activity,omitempty"`
	IsOnline              *bool                    `json:"isOnline,omitempty"`
	IsBanned              *int                     `json:"is_banned,omitempty"`
	JoinedDate            *int                     `json:"joined_date,omitempty"`
	RestoreData           interface{}              `json:"restore_data,omitempty"`
	RestorePercents       interface{}              `json:"restore_percents,omitempty"`
	SoldItemsCount        *int                     `json:"sold_items_count,omitempty"`
	UserID                *int                     `json:"user_id,omitempty"`
	Username              *string                  `json:"username,omitempty"`
}

type ItemModelSellerContacts struct {
	BanReason *string `json:"ban_reason,omitempty"`
	Telegram  *string `json:"telegram,omitempty"`
}

type ListParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Currency of the created invoice.
	Currency *InvoicesCurrency `json:"currency,omitempty"`
	// Status of the invoice.
	Status *InvoicesStatus `json:"status,omitempty"`
	// Invoice amount.
	Amount *float64 `json:"amount,omitempty"`
	// Merchant ID.
	MerchantID *int `json:"merchant_id,omitempty"`
}

type ListResponse struct {
	Payments   interface{}     `json:"payments,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type ListResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type MihoyoParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// Has linked email.
	Email *CategorySearchEmail `json:"email,omitempty"`
	// Has linked external accounts.
	EA *CategorySearchEA `json:"ea,omitempty"`
	// Region.
	Region []string `json:"region,omitempty"`
	// List of disallowed regions.
	NotRegion []string `json:"not_region,omitempty"`
	// List of characters.
	GenshinCharacter []int `json:"genshin_character[],omitempty"`
	// List of minimum constellations on characters.
	GenshinCharacterConstellations map[string]int `json:"genshin_character_constellations,omitempty"`
	// List of maximum constellations on characters.
	GenshinCharacterConstellationsMax map[string]int `json:"genshin_character_constellations_max,omitempty"`
	// List of weapons.
	GenshinWeapon []int `json:"genshin_weapon[],omitempty"`
	// Minimum number of characters.
	GenshinCharMin *int `json:"genshin_char_min,omitempty"`
	// Maximum number of characters.
	GenshinCharMax *int `json:"genshin_char_max,omitempty"`
	// Minimum number of legendary characters.
	GenshinLegendaryMin *int `json:"genshin_legendary_min,omitempty"`
	// Maximum number of legendary characters.
	GenshinLegendaryMax *int `json:"genshin_legendary_max,omitempty"`
	// Minimum level.
	GenshinLevelMin *int `json:"genshin_level_min,omitempty"`
	// Maximum level.
	GenshinLevelMax *int `json:"genshin_level_max,omitempty"`
	// Minimum number of legendary weapon characters.
	GenshinLegendaryWeaponMin *int `json:"genshin_legendary_weapon_min,omitempty"`
	// Maximum number of legendary weapon characters.
	GenshinLegendaryWeaponMax *int `json:"genshin_legendary_weapon_max,omitempty"`
	// Minimum number of constellations on legendary characters.
	ConstellationsMin *int `json:"constellations_min,omitempty"`
	// Maximum number of constellations on legendary characters.
	ConstellationsMax *int `json:"constellations_max,omitempty"`
	// Minimum number of achievements.
	GenshinAchievementMin *int `json:"genshin_achievement_min,omitempty"`
	// Maximum number of achievements.
	GenshinAchievementMax *int `json:"genshin_achievement_max,omitempty"`
	// Minimum number of primogems.
	GenshinCurrencyMin *int `json:"genshin_currency_min,omitempty"`
	// Maximum number of primogems.
	GenshinCurrencyMax *int `json:"genshin_currency_max,omitempty"`
	// List of characters.
	HonkaiCharacter []int `json:"honkai_character[],omitempty"`
	// List of minimum eidolons on characters.
	HonkaiCharacterEidolons map[string]int `json:"honkai_character_eidolons,omitempty"`
	// List of maximum eidolons on characters.
	HonkaiCharacterEidolonsMax map[string]int `json:"honkai_character_eidolons_max,omitempty"`
	// List of weapons.
	HonkaiWeapon []int `json:"honkai_weapon[],omitempty"`
	// Minimum number of characters.
	HonkaiCharMin *int `json:"honkai_char_min,omitempty"`
	// Maximum number of characters.
	HonkaiCharMax *int `json:"honkai_char_max,omitempty"`
	// Minimum number of legendary characters.
	HonkaiLegendaryMin *int `json:"honkai_legendary_min,omitempty"`
	// Maximum number of legendary characters.
	HonkaiLegendaryMax *int `json:"honkai_legendary_max,omitempty"`
	// Minimum level.
	HonkaiLevelMin *int `json:"honkai_level_min,omitempty"`
	// Maximum level.
	HonkaiLevelMax *int `json:"honkai_level_max,omitempty"`
	// Minimum number of legendary weapon characters.
	HonkaiLegendaryWeaponMin *int `json:"honkai_legendary_weapon_min,omitempty"`
	// Maximum number of legendary weapon characters.
	HonkaiLegendaryWeaponMax *int `json:"honkai_legendary_weapon_max,omitempty"`
	// Minimum number of constellations on Honkai: Star Rail legendary characters.
	EidolonsMin *int `json:"eidolons_min,omitempty"`
	// Maximum number of legendary Honkai: Star Rail weapon characters.
	EidolonsMax *int `json:"eidolons_max,omitempty"`
	// Minimum number of achievements.
	HonkaiAchievementMin *int `json:"honkai_achievement_min,omitempty"`
	// Maximum number of achievements.
	HonkaiAchievementMax *int `json:"honkai_achievement_max,omitempty"`
	// Minimum number of Stellar Jade.
	HonkaiCurrencyMin *int `json:"honkai_currency_min,omitempty"`
	// Maximum number of Stellar Jade.
	HonkaiCurrencyMax *int `json:"honkai_currency_max,omitempty"`
	// List of Zenless Zone Zero characters.
	ZenlessCharacter []int `json:"zenless_character[],omitempty"`
	// List of minimum cinemas on characters.
	ZenlessCharacterCinemas map[string]int `json:"zenless_character_cinemas,omitempty"`
	// List of maximum cinemas on characters.
	ZenlessCharacterCinemasMax map[string]int `json:"zenless_character_cinemas_max,omitempty"`
	// List of Zenless Zone Zero weapons.
	ZenlessWeapon []int `json:"zenless_weapon[],omitempty"`
	// Minimum number of Zenless Zone Zero legendary characters.
	ZenlessLegendaryMin *int `json:"zenless_legendary_min,omitempty"`
	// Maximum number of Zenless Zone Zero legendary characters.
	ZenlessLegendaryMax *int `json:"zenless_legendary_max,omitempty"`
	// Minimum number of cinemas on Zenless Zone Zero characters.
	CinemasMin *int `json:"cinemas_min,omitempty"`
	// Maximum number of cinemas on Zenless Zone Zero characters.
	CinemasMax *int `json:"cinemas_max,omitempty"`
	// Minimum number of legendary Zenless Zone Zero weapon characters.
	ZenlessLegendaryWeaponMin *int `json:"zenless_legendary_weapon_min,omitempty"`
	// Maximum number of legendary Zenless Zone Zero weapon characters.
	ZenlessLegendaryWeaponMax *int `json:"zenless_legendary_weapon_max,omitempty"`
	// Minimum number of Zenless Zone Zero characters.
	ZenlessCharMin *int `json:"zenless_char_min,omitempty"`
	// Maximum number of Zenless Zone Zero characters.
	ZenlessCharMax *int `json:"zenless_char_max,omitempty"`
	// Minimum Zenless Zone Zero level.
	ZenlessLevelMin *int `json:"zenless_level_min,omitempty"`
	// Maximum Zenless Zone Zero level.
	ZenlessLevelMax *int `json:"zenless_level_max,omitempty"`
	// Minimum count of Zenless Zone Zero achievements.
	ZenlessAchievementMin *int `json:"zenless_achievement_min,omitempty"`
	// Maximum count of Zenless Zone Zero achievements.
	ZenlessAchievementMax *int `json:"zenless_achievement_max,omitempty"`
	// Minimum count of Zenless Zone Zero polychrome.
	ZenlessCurrencyMin *int `json:"zenless_currency_min,omitempty"`
	// Maximum count of Zenless Zone Zero polychrome.
	ZenlessCurrencyMax *int `json:"zenless_currency_max,omitempty"`
	// Number of days the account has been offline.
	Daybreak *int `json:"daybreak,omitempty"`
}

type MihoyoResponse struct {
	CacheTTL        *int                      `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                     `json:"hasNextPage,omitempty"`
	Items           []MihoyoResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                      `json:"lastModified,omitempty"`
	Page            *int                      `json:"page,omitempty"`
	PerPage         *int                      `json:"perPage,omitempty"`
	SearchUrl       *string                   `json:"searchUrl,omitempty"`
	ServerTime      *int                      `json:"serverTime,omitempty"`
	StickyItems     []interface{}             `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo           `json:"system_info,omitempty"`
	TotalItems      *int                      `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}               `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                     `json:"wasCached,omitempty"`
}

type MihoyoResponseItemsItem struct {
	AccountLink                           *string                                        `json:"accountLink,omitempty"`
	AccountLinks                          []MihoyoResponseItemsItemAccountLinksItem      `json:"accountLinks,omitempty"`
	AllowAskDiscount                      *int                                           `json:"allow_ask_discount,omitempty"`
	BumpSettings                          *MihoyoResponseItemsItemBumpSettings           `json:"bumpSettings,omitempty"`
	CanBumpItem                           *bool                                          `json:"canBumpItem,omitempty"`
	CanBuyItem                            *bool                                          `json:"canBuyItem,omitempty"`
	CanChangePassword                     *bool                                          `json:"canChangePassword,omitempty"`
	CanCloseItem                          *bool                                          `json:"canCloseItem,omitempty"`
	CanDeleteItem                         *bool                                          `json:"canDeleteItem,omitempty"`
	CanEditItem                           *bool                                          `json:"canEditItem,omitempty"`
	CanOpenItem                           *bool                                          `json:"canOpenItem,omitempty"`
	CanReportItem                         *bool                                          `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase            *bool                                          `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem                          *bool                                          `json:"canStickItem,omitempty"`
	CanUnstickItem                        *bool                                          `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats                    *bool                                          `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount                    *bool                                          `json:"canValidateAccount,omitempty"`
	CanViewAccountLink                    *bool                                          `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData                 *bool                                          `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData                      *bool                                          `json:"canViewLoginData,omitempty"`
	CategoryID                            *int                                           `json:"category_id,omitempty"`
	Description                           *string                                        `json:"description,omitempty"`
	DescriptionEnHtml                     *string                                        `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain                    *string                                        `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml                       *string                                        `json:"descriptionHtml,omitempty"`
	DescriptionPlain                      *string                                        `json:"descriptionPlain,omitempty"`
	DescriptionEn                         *string                                        `json:"description_en,omitempty"`
	EditDate                              *int                                           `json:"edit_date,omitempty"`
	EmailLoginUrl                         *string                                        `json:"emailLoginUrl,omitempty"`
	EmailProvider                         *string                                        `json:"email_provider,omitempty"`
	EmailType                             *string                                        `json:"email_type,omitempty"`
	ExtendedGuarantee                     *int                                           `json:"extended_guarantee,omitempty"`
	FeedbackData                          interface{}                                    `json:"feedback_data,omitempty"`
	GenshinCharacters                     []MihoyoResponseItemsItemGenshinCharactersItem `json:"genshinCharacters,omitempty"`
	Guarantee                             interface{}                                    `json:"guarantee,omitempty"`
	HasPendingAutoBuy                     *bool                                          `json:"hasPendingAutoBuy,omitempty"`
	HonkaiCharacters                      []MihoyoResponseItemsItemHonkaiCharactersItem  `json:"honkaiCharacters,omitempty"`
	IsIgnored                             *bool                                          `json:"isIgnored,omitempty"`
	IsSticky                              *int                                           `json:"is_sticky,omitempty"`
	ItemOriginPhrase                      *string                                        `json:"itemOriginPhrase,omitempty"`
	ItemDomain                            *string                                        `json:"item_domain,omitempty"`
	ItemID                                *int                                           `json:"item_id,omitempty"`
	ItemOrigin                            *string                                        `json:"item_origin,omitempty"`
	ItemState                             *string                                        `json:"item_state,omitempty"`
	MihoyoLinkedAccounts                  *MihoyoResponseItemsItemMihoyoLinkedAccounts   `json:"mihoyoLinkedAccounts,omitempty"`
	MihoyoLinkedAccountsString            *string                                        `json:"mihoyoLinkedAccountsString,omitempty"`
	MihoyoRegionPhrase                    *string                                        `json:"mihoyoRegionPhrase,omitempty"`
	MihoyoEmail                           *int                                           `json:"mihoyo_email,omitempty"`
	MihoyoGenshinAbyssProcess             *string                                        `json:"mihoyo_genshin_abyss_process,omitempty"`
	MihoyoGenshinAchievementCount         *int                                           `json:"mihoyo_genshin_achievement_count,omitempty"`
	MihoyoGenshinActivityDays             *int                                           `json:"mihoyo_genshin_activity_days,omitempty"`
	MihoyoGenshinCharacterCount           *int                                           `json:"mihoyo_genshin_character_count,omitempty"`
	MihoyoGenshinConstellationsCount      *int                                           `json:"mihoyo_genshin_constellations_count,omitempty"`
	MihoyoGenshinCurrency                 *int                                           `json:"mihoyo_genshin_currency,omitempty"`
	MihoyoGenshinLegendaryCharactersCount *int                                           `json:"mihoyo_genshin_legendary_characters_count,omitempty"`
	MihoyoGenshinLegendaryWeaponsCount    *int                                           `json:"mihoyo_genshin_legendary_weapons_count,omitempty"`
	MihoyoGenshinLevel                    *int                                           `json:"mihoyo_genshin_level,omitempty"`
	MihoyoHasLinkedAccounts               *int                                           `json:"mihoyo_has_linked_accounts,omitempty"`
	MihoyoHonkaiAbyssProcess              *string                                        `json:"mihoyo_honkai_abyss_process,omitempty"`
	MihoyoHonkaiAchievementCount          *int                                           `json:"mihoyo_honkai_achievement_count,omitempty"`
	MihoyoHonkaiActivityDays              *int                                           `json:"mihoyo_honkai_activity_days,omitempty"`
	MihoyoHonkaiCharacterCount            *int                                           `json:"mihoyo_honkai_character_count,omitempty"`
	MihoyoHonkaiCurrency                  *int                                           `json:"mihoyo_honkai_currency,omitempty"`
	MihoyoHonkaiEidolonsCount             *int                                           `json:"mihoyo_honkai_eidolons_count,omitempty"`
	MihoyoHonkaiLegendaryCharactersCount  *int                                           `json:"mihoyo_honkai_legendary_characters_count,omitempty"`
	MihoyoHonkaiLegendaryWeaponsCount     *int                                           `json:"mihoyo_honkai_legendary_weapons_count,omitempty"`
	MihoyoHonkaiLevel                     *int                                           `json:"mihoyo_honkai_level,omitempty"`
	MihoyoID                              *int                                           `json:"mihoyo_id,omitempty"`
	MihoyoItemID                          *int                                           `json:"mihoyo_item_id,omitempty"`
	MihoyoLastActivity                    *int                                           `json:"mihoyo_last_activity,omitempty"`
	MihoyoRegion                          *string                                        `json:"mihoyo_region,omitempty"`
	MihoyoZenlessAbyssProcess             *string                                        `json:"mihoyo_zenless_abyss_process,omitempty"`
	MihoyoZenlessAchievementCount         *int                                           `json:"mihoyo_zenless_achievement_count,omitempty"`
	MihoyoZenlessActivityDays             *int                                           `json:"mihoyo_zenless_activity_days,omitempty"`
	MihoyoZenlessCharacterCount           *int                                           `json:"mihoyo_zenless_character_count,omitempty"`
	MihoyoZenlessCinemasCount             *int                                           `json:"mihoyo_zenless_cinemas_count,omitempty"`
	MihoyoZenlessCurrency                 *int                                           `json:"mihoyo_zenless_currency,omitempty"`
	MihoyoZenlessLegendaryCharactersCount *int                                           `json:"mihoyo_zenless_legendary_characters_count,omitempty"`
	MihoyoZenlessLegendaryWeaponsCount    *int                                           `json:"mihoyo_zenless_legendary_weapons_count,omitempty"`
	MihoyoZenlessLevel                    *int                                           `json:"mihoyo_zenless_level,omitempty"`
	NoteText                              interface{}                                    `json:"note_text,omitempty"`
	Nsb                                   *int                                           `json:"nsb,omitempty"`
	Price                                 *int                                           `json:"price,omitempty"`
	PriceWithSellerFee                    *float64                                       `json:"priceWithSellerFee,omitempty"`
	PriceCurrency                         *string                                        `json:"price_currency,omitempty"`
	PublishedDate                         *int                                           `json:"published_date,omitempty"`
	RefreshedDate                         *int                                           `json:"refreshed_date,omitempty"`
	ResaleItemOrigin                      *string                                        `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount             *int                                           `json:"restore_items_category_count,omitempty"`
	RubPrice                              *int                                           `json:"rub_price,omitempty"`
	Seller                                *MihoyoResponseItemsItemSeller                 `json:"seller,omitempty"`
	ShowGetEmailCodeButton                *bool                                          `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount                *int                                           `json:"sold_items_category_count,omitempty"`
	Tags                                  interface{}                                    `json:"tags,omitempty"`
	Title                                 *string                                        `json:"title,omitempty"`
	TitleEn                               *string                                        `json:"title_en,omitempty"`
	UpdateStatDate                        *int                                           `json:"update_stat_date,omitempty"`
	ViewCount                             *int                                           `json:"view_count,omitempty"`
	ZenlessCharacters                     []MihoyoResponseItemsItemZenlessCharactersItem `json:"zenlessCharacters,omitempty"`
}

type MihoyoResponseItemsItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type MihoyoResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type MihoyoResponseItemsItemGenshinCharactersItem struct {
	ActivedConstellationNum *int                                                          `json:"actived_constellation_num,omitempty"`
	Background              *string                                                       `json:"background,omitempty"`
	Costumes                []interface{}                                                 `json:"costumes,omitempty"`
	Element                 *string                                                       `json:"element,omitempty"`
	External                interface{}                                                   `json:"external,omitempty"`
	Fetter                  *int                                                          `json:"fetter,omitempty"`
	Icon                    *string                                                       `json:"icon,omitempty"`
	ID                      *int                                                          `json:"id,omitempty"`
	Image                   *string                                                       `json:"image,omitempty"`
	Level                   *int                                                          `json:"level,omitempty"`
	Name                    *string                                                       `json:"name,omitempty"`
	Rarity                  *int                                                          `json:"rarity,omitempty"`
	Reliquaries             []MihoyoResponseItemsItemGenshinCharactersItemReliquariesItem `json:"reliquaries,omitempty"`
	Weapon                  *MihoyoResponseItemsItemGenshinCharactersItemWeapon           `json:"weapon,omitempty"`
}

type MihoyoResponseItemsItemGenshinCharactersItemReliquariesItem struct {
	Icon    *string `json:"icon,omitempty"`
	ID      *int    `json:"id,omitempty"`
	Level   *int    `json:"level,omitempty"`
	Name    *string `json:"name,omitempty"`
	Pos     *int    `json:"pos,omitempty"`
	PosName *string `json:"pos_name,omitempty"`
	Rarity  *int    `json:"rarity,omitempty"`
}

type MihoyoResponseItemsItemGenshinCharactersItemWeapon struct {
	AffixLevel   *int    `json:"affix_level,omitempty"`
	Desc         *string `json:"desc,omitempty"`
	Icon         *string `json:"icon,omitempty"`
	ID           *int    `json:"id,omitempty"`
	Level        *int    `json:"level,omitempty"`
	Name         *string `json:"name,omitempty"`
	PromoteLevel *int    `json:"promote_level,omitempty"`
	Rarity       *int    `json:"rarity,omitempty"`
	Type_        *int    `json:"type,omitempty"`
	TypeName     *string `json:"type_name,omitempty"`
}

type MihoyoResponseItemsItemHonkaiCharactersItem struct {
	BaseType     *int                                                       `json:"base_type,omitempty"`
	Element      *string                                                    `json:"element,omitempty"`
	ElementImage *string                                                    `json:"elementImage,omitempty"`
	Equip        *MihoyoResponseItemsItemHonkaiCharactersItemEquip          `json:"equip,omitempty"`
	FigurePath   *string                                                    `json:"figure_path,omitempty"`
	Icon         *string                                                    `json:"icon,omitempty"`
	ID           *int                                                       `json:"id,omitempty"`
	Image        *string                                                    `json:"image,omitempty"`
	Level        *int                                                       `json:"level,omitempty"`
	Name         *string                                                    `json:"name,omitempty"`
	Ornaments    []MihoyoResponseItemsItemHonkaiCharactersItemOrnamentsItem `json:"ornaments,omitempty"`
	Rank         *int                                                       `json:"rank,omitempty"`
	Rarity       *int                                                       `json:"rarity,omitempty"`
	Relics       []MihoyoResponseItemsItemHonkaiCharactersItemRelicsItem    `json:"relics,omitempty"`
}

type MihoyoResponseItemsItemHonkaiCharactersItemEquip struct {
	Desc   *string `json:"desc,omitempty"`
	Icon   *string `json:"icon,omitempty"`
	ID     *int    `json:"id,omitempty"`
	Level  *int    `json:"level,omitempty"`
	Name   *string `json:"name,omitempty"`
	Rank   *int    `json:"rank,omitempty"`
	Rarity *int    `json:"rarity,omitempty"`
}

type MihoyoResponseItemsItemHonkaiCharactersItemOrnamentsItem struct {
	Desc         *string                                                                  `json:"desc,omitempty"`
	Icon         *string                                                                  `json:"icon,omitempty"`
	ID           *int                                                                     `json:"id,omitempty"`
	Level        *int                                                                     `json:"level,omitempty"`
	MainProperty *MihoyoResponseItemsItemHonkaiCharactersItemOrnamentsItemMainProperty    `json:"main_property,omitempty"`
	Name         *string                                                                  `json:"name,omitempty"`
	Pos          *int                                                                     `json:"pos,omitempty"`
	Properties   []MihoyoResponseItemsItemHonkaiCharactersItemOrnamentsItemPropertiesItem `json:"properties,omitempty"`
	Rarity       *int                                                                     `json:"rarity,omitempty"`
}

type MihoyoResponseItemsItemHonkaiCharactersItemOrnamentsItemMainProperty struct {
	PropertyType *int    `json:"property_type,omitempty"`
	Times        *int    `json:"times,omitempty"`
	Value        *string `json:"value,omitempty"`
}

type MihoyoResponseItemsItemHonkaiCharactersItemOrnamentsItemPropertiesItem struct {
	PropertyType *int    `json:"property_type,omitempty"`
	Times        *int    `json:"times,omitempty"`
	Value        *string `json:"value,omitempty"`
}

type MihoyoResponseItemsItemHonkaiCharactersItemRelicsItem struct {
	Desc         *string                                                               `json:"desc,omitempty"`
	Icon         *string                                                               `json:"icon,omitempty"`
	ID           *int                                                                  `json:"id,omitempty"`
	Level        *int                                                                  `json:"level,omitempty"`
	MainProperty *MihoyoResponseItemsItemHonkaiCharactersItemRelicsItemMainProperty    `json:"main_property,omitempty"`
	Name         *string                                                               `json:"name,omitempty"`
	Pos          *int                                                                  `json:"pos,omitempty"`
	Properties   []MihoyoResponseItemsItemHonkaiCharactersItemRelicsItemPropertiesItem `json:"properties,omitempty"`
	Rarity       *int                                                                  `json:"rarity,omitempty"`
}

type MihoyoResponseItemsItemHonkaiCharactersItemRelicsItemMainProperty struct {
	PropertyType *int    `json:"property_type,omitempty"`
	Times        *int    `json:"times,omitempty"`
	Value        *string `json:"value,omitempty"`
}

type MihoyoResponseItemsItemHonkaiCharactersItemRelicsItemPropertiesItem struct {
	PropertyType *int    `json:"property_type,omitempty"`
	Times        *int    `json:"times,omitempty"`
	Value        *string `json:"value,omitempty"`
}

type MihoyoResponseItemsItemMihoyoLinkedAccounts struct {
	Legacy *bool    `json:"legacy,omitempty"`
	Links  []string `json:"links,omitempty"`
}

type MihoyoResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type MihoyoResponseItemsItemZenlessCharactersItem struct {
	AvatarProfession *int                                                `json:"avatar_profession,omitempty"`
	CampNameMi18n    *string                                             `json:"camp_name_mi18n,omitempty"`
	ElementIcon      *string                                             `json:"elementIcon,omitempty"`
	ElementType      *int                                                `json:"element_type,omitempty"`
	FullNameMi18n    *string                                             `json:"full_name_mi18n,omitempty"`
	ID               *int                                                `json:"id,omitempty"`
	Level            *int                                                `json:"level,omitempty"`
	Name             *string                                             `json:"name,omitempty"`
	NameMi18n        *string                                             `json:"name_mi18n,omitempty"`
	ProfessionIcon   *string                                             `json:"professionIcon,omitempty"`
	Rank             *int                                                `json:"rank,omitempty"`
	Rarity           *int                                                `json:"rarity,omitempty"`
	RarityIcon       *string                                             `json:"rarityIcon,omitempty"`
	Weapon           *MihoyoResponseItemsItemZenlessCharactersItemWeapon `json:"weapon,omitempty"`
}

type MihoyoResponseItemsItemZenlessCharactersItemWeapon struct {
	Icon           *string                                                                `json:"icon,omitempty"`
	ID             *int                                                                   `json:"id,omitempty"`
	Level          *int                                                                   `json:"level,omitempty"`
	MainProperties []MihoyoResponseItemsItemZenlessCharactersItemWeaponMainPropertiesItem `json:"main_properties,omitempty"`
	Name           *string                                                                `json:"name,omitempty"`
	Profession     *int                                                                   `json:"profession,omitempty"`
	Properties     []MihoyoResponseItemsItemZenlessCharactersItemWeaponPropertiesItem     `json:"properties,omitempty"`
	Rarity         *int                                                                   `json:"rarity,omitempty"`
	RarityIcon     *string                                                                `json:"rarityIcon,omitempty"`
	Star           *int                                                                   `json:"star,omitempty"`
	StarIcon       *string                                                                `json:"starIcon,omitempty"`
	TalentContent  *string                                                                `json:"talent_content,omitempty"`
	TalentTitle    *string                                                                `json:"talent_title,omitempty"`
}

type MihoyoResponseItemsItemZenlessCharactersItemWeaponMainPropertiesItem struct {
	Base         *string `json:"base,omitempty"`
	PropertyID   *int    `json:"property_id,omitempty"`
	PropertyName *string `json:"property_name,omitempty"`
}

type MihoyoResponseItemsItemZenlessCharactersItemWeaponPropertiesItem struct {
	Base         *string `json:"base,omitempty"`
	PropertyID   *int    `json:"property_id,omitempty"`
	PropertyName *string `json:"property_name,omitempty"`
}

type MihoyoResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type MinecraftParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Name of subscription.
	Subscription *CategorySearchSubscription `json:"subscription,omitempty"`
	// Length of subscription.
	SubscriptionLength *int `json:"subscription_length,omitempty"`
	// In what notation is time measured.
	SubscriptionPeriod *CategorySearchSubscriptionPeriod `json:"subscription_period,omitempty"`
	// Is auto renewal enabled.
	Autorenewal *CategorySearchAutorenewal `json:"autorenewal,omitempty"`
	// Has java edition.
	Java *CategorySearchJava `json:"java,omitempty"`
	// Has bedrock edition.
	Bedrock *CategorySearchBedrock `json:"bedrock,omitempty"`
	// Has Minecraft Dungeons.
	Dungeons *CategorySearchDungeons `json:"dungeons,omitempty"`
	// Has Minecraft Legends.
	Legends *CategorySearchLegends `json:"legends,omitempty"`
	// Can change nickname.
	ChangeNickname *CategorySearchChangeNickname `json:"change_nickname,omitempty"`
	// List of capes.
	Capes []string `json:"capes[],omitempty"`
	// Minimum number of capes.
	CapesMin *int `json:"capes_min,omitempty"`
	// Maximum number of capes.
	CapesMax *int `json:"capes_max,omitempty"`
	// List of allowed countries.
	Country []string `json:"country[],omitempty"`
	// List of disallowed countries.
	NotCountry []string `json:"not_country[],omitempty"`
	// Has active Hypixel ban.
	HypixelBan *CategorySearchHypixelBan `json:"hypixel_ban,omitempty"`
	// Is API enabled in Hypixel Skyblock.
	HypixelSkyblockAPIEnabled *CategorySearchHypixelSkyblockAPIEnabled `json:"hypixel_skyblock_api_enabled,omitempty"`
	// Rank on hypixel.
	RankHypixel []string `json:"rank_hypixel[],omitempty"`
	// Minimum number of level hypixel.
	LevelHypixelMin *int `json:"level_hypixel_min,omitempty"`
	// Maximum number of level hypixel.
	LevelHypixelMax *int `json:"level_hypixel_max,omitempty"`
	// Minimum number of achievement hypixel.
	AchievementHypixelMin *int `json:"achievement_hypixel_min,omitempty"`
	// Maximum number of achievement hypixel.
	AchievementHypixelMax *int `json:"achievement_hypixel_max,omitempty"`
	// Minimum level on Hypixel SkyBlock.
	LevelHypixelSkyblockMin *int `json:"level_hypixel_skyblock_min,omitempty"`
	// Maximum level on Hypixel SkyBlock.
	LevelHypixelSkyblockMax *int `json:"level_hypixel_skyblock_max,omitempty"`
	// Minimum net worth on Hypixel SkyBlock.
	NetWorthHypixelSkyblockMin *int `json:"net_worth_hypixel_skyblock_min,omitempty"`
	// Maximum net worth on Hypixel SkyBlock.
	NetWorthHypixelSkyblockMax *int `json:"net_worth_hypixel_skyblock_max,omitempty"`
	// How old is the account.
	Reg *int `json:"reg,omitempty"`
	// In what notation is time measured.
	RegPeriod *CategorySearchRegPeriod `json:"reg_period,omitempty"`
	// How old is the last login account.
	LastLoginHypixel *int `json:"last_login_hypixel,omitempty"`
	// In what notation is time measured.
	LastLoginHypixelPeriod *CategorySearchLastLoginHypixelPeriod `json:"last_login_hypixel_period,omitempty"`
	// Can change details.
	CanChangeDetails *CategorySearchCanChangeDetails `json:"can_change_details,omitempty"`
	// Minimum number of characters in nickname.
	NicknameLengthMin *int `json:"nickname_length_min,omitempty"`
	// Maximum number of characters in nickname.
	NicknameLengthMax *int `json:"nickname_length_max,omitempty"`
	// Was Hypixel ban parsed by Market.
	HypixelBanParsed *CategorySearchHypixelBanParsed `json:"hypixel_ban_parsed,omitempty"`
	// Minimum number of Minecoins.
	MinecoinsMin *int `json:"minecoins_min,omitempty"`
	// Maximum number of Minecoins.
	MinecoinsMax *int `json:"minecoins_max,omitempty"`
}

type MinecraftResponse struct {
	CacheTTL        *int                         `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                        `json:"hasNextPage,omitempty"`
	Items           []MinecraftResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                         `json:"lastModified,omitempty"`
	Page            *int                         `json:"page,omitempty"`
	PerPage         *int                         `json:"perPage,omitempty"`
	SearchUrl       *string                      `json:"searchUrl,omitempty"`
	ServerTime      *int                         `json:"serverTime,omitempty"`
	StickyItems     []interface{}                `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo              `json:"system_info,omitempty"`
	TotalItems      *int                         `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}                  `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                        `json:"wasCached,omitempty"`
}

type MinecraftResponseItemsItem struct {
	AccountLink                      *string                                      `json:"accountLink,omitempty"`
	AccountLinks                     []MinecraftResponseItemsItemAccountLinksItem `json:"accountLinks,omitempty"`
	AllowAskDiscount                 *int                                         `json:"allow_ask_discount,omitempty"`
	BumpSettings                     *MinecraftResponseItemsItemBumpSettings      `json:"bumpSettings,omitempty"`
	CanBumpItem                      *bool                                        `json:"canBumpItem,omitempty"`
	CanBuyItem                       *bool                                        `json:"canBuyItem,omitempty"`
	CanChangePassword                *bool                                        `json:"canChangePassword,omitempty"`
	CanCloseItem                     *bool                                        `json:"canCloseItem,omitempty"`
	CanDeleteItem                    *bool                                        `json:"canDeleteItem,omitempty"`
	CanEditItem                      *bool                                        `json:"canEditItem,omitempty"`
	CanOpenItem                      *bool                                        `json:"canOpenItem,omitempty"`
	CanReportItem                    *bool                                        `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase       *bool                                        `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem                     *bool                                        `json:"canStickItem,omitempty"`
	CanUnstickItem                   *bool                                        `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats               *bool                                        `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount               *bool                                        `json:"canValidateAccount,omitempty"`
	CanViewAccountLink               *bool                                        `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData            *bool                                        `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData                 *bool                                        `json:"canViewLoginData,omitempty"`
	CategoryID                       *int                                         `json:"category_id,omitempty"`
	Description                      *string                                      `json:"description,omitempty"`
	DescriptionEnHtml                *string                                      `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain               *string                                      `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml                  *string                                      `json:"descriptionHtml,omitempty"`
	DescriptionPlain                 *string                                      `json:"descriptionPlain,omitempty"`
	DescriptionEn                    *string                                      `json:"description_en,omitempty"`
	EditDate                         *int                                         `json:"edit_date,omitempty"`
	EmailLoginUrl                    *string                                      `json:"emailLoginUrl,omitempty"`
	EmailProvider                    *string                                      `json:"email_provider,omitempty"`
	EmailType                        *string                                      `json:"email_type,omitempty"`
	ExtendedGuarantee                *int                                         `json:"extended_guarantee,omitempty"`
	FeedbackData                     interface{}                                  `json:"feedback_data,omitempty"`
	Guarantee                        interface{}                                  `json:"guarantee,omitempty"`
	HasPendingAutoBuy                *bool                                        `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                        *bool                                        `json:"isIgnored,omitempty"`
	IsSticky                         *int                                         `json:"is_sticky,omitempty"`
	ItemOriginPhrase                 *string                                      `json:"itemOriginPhrase,omitempty"`
	ItemDomain                       *string                                      `json:"item_domain,omitempty"`
	ItemID                           *int                                         `json:"item_id,omitempty"`
	ItemOrigin                       *string                                      `json:"item_origin,omitempty"`
	ItemState                        *string                                      `json:"item_state,omitempty"`
	MinecraftHasPaidLicense          *bool                                        `json:"minecraftHasPaidLicense,omitempty"`
	MinecraftBedrock                 *int                                         `json:"minecraft_bedrock,omitempty"`
	MinecraftCanChangeNickname       *int                                         `json:"minecraft_can_change_nickname,omitempty"`
	MinecraftCapes                   []interface{}                                `json:"minecraft_capes,omitempty"`
	MinecraftCapesCount              *int                                         `json:"minecraft_capes_count,omitempty"`
	MinecraftCountry                 *string                                      `json:"minecraft_country,omitempty"`
	MinecraftCreatedAt               *int                                         `json:"minecraft_created_at,omitempty"`
	MinecraftDungeons                *int                                         `json:"minecraft_dungeons,omitempty"`
	MinecraftEmailResetDate          *int                                         `json:"minecraft_email_reset_date,omitempty"`
	MinecraftHypixelAchievement      *int                                         `json:"minecraft_hypixel_achievement,omitempty"`
	MinecraftHypixelBan              *int                                         `json:"minecraft_hypixel_ban,omitempty"`
	MinecraftHypixelBanReason        *string                                      `json:"minecraft_hypixel_ban_reason,omitempty"`
	MinecraftHypixelLastLogin        *int                                         `json:"minecraft_hypixel_last_login,omitempty"`
	MinecraftHypixelLevel            *int                                         `json:"minecraft_hypixel_level,omitempty"`
	MinecraftHypixelRank             *string                                      `json:"minecraft_hypixel_rank,omitempty"`
	MinecraftHypixelSkyblockLevel    *int                                         `json:"minecraft_hypixel_skyblock_level,omitempty"`
	MinecraftHypixelSkyblockNetWorth *int                                         `json:"minecraft_hypixel_skyblock_net_worth,omitempty"`
	MinecraftID                      *string                                      `json:"minecraft_id,omitempty"`
	MinecraftItemID                  *int                                         `json:"minecraft_item_id,omitempty"`
	MinecraftJava                    *int                                         `json:"minecraft_java,omitempty"`
	MinecraftLegends                 *int                                         `json:"minecraft_legends,omitempty"`
	MinecraftNickname                *string                                      `json:"minecraft_nickname,omitempty"`
	MinecraftSkin                    *string                                      `json:"minecraft_skin,omitempty"`
	MinecraftSubscriptionAutoRenew   *int                                         `json:"minecraft_subscription_auto_renew,omitempty"`
	MinecraftSubscriptionEnds        *int                                         `json:"minecraft_subscription_ends,omitempty"`
	MinecraftSubscriptionName        *string                                      `json:"minecraft_subscription_name,omitempty"`
	NoteText                         interface{}                                  `json:"note_text,omitempty"`
	Nsb                              *int                                         `json:"nsb,omitempty"`
	Price                            *int                                         `json:"price,omitempty"`
	PriceWithSellerFee               *float64                                     `json:"priceWithSellerFee,omitempty"`
	PriceCurrency                    *string                                      `json:"price_currency,omitempty"`
	PublishedDate                    *int                                         `json:"published_date,omitempty"`
	RefreshedDate                    *int                                         `json:"refreshed_date,omitempty"`
	ResaleItemOrigin                 *string                                      `json:"resale_item_origin,omitempty"`
	RubPrice                         *int                                         `json:"rub_price,omitempty"`
	Seller                           *MinecraftResponseItemsItemSeller            `json:"seller,omitempty"`
	ShowGetEmailCodeButton           *bool                                        `json:"showGetEmailCodeButton,omitempty"`
	Tags                             interface{}                                  `json:"tags,omitempty"`
	Title                            *string                                      `json:"title,omitempty"`
	TitleEn                          *string                                      `json:"title_en,omitempty"`
	UpdateStatDate                   *int                                         `json:"update_stat_date,omitempty"`
	ViewCount                        *int                                         `json:"view_count,omitempty"`
}

type MinecraftResponseItemsItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type MinecraftResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type MinecraftResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     interface{} `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type MinecraftResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type NoteParams struct {
	// Text of note.
	Text *string `json:"text,omitempty"`
}

type OrdersParams struct {
	// User id.
	UserID *int `json:"user_id,omitempty"`
	// Accounts category.
	CategoryID *AccountsListCategoryID `json:"category_id,omitempty"`
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Account status.
	Show *AccountsListShow `json:"show,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// Login.
	Login *string `json:"login,omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Order by.
	OrderBy *AccountsListOrderBy `json:"order_by,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
}

type ParamsResponse struct {
	BaseParams interface{}                `json:"base_params,omitempty"`
	Category   *ParamsResponseCategory    `json:"category,omitempty"`
	Params     []ParamsResponseParamsItem `json:"params,omitempty"`
	SystemInfo *RespSystemInfo            `json:"system_info,omitempty"`
}

type ParamsResponseCategory struct {
	AccountPriceMin           *int    `json:"account_price_min,omitempty"`
	AddItemAvailable          *int    `json:"add_item_available,omitempty"`
	AvailableTempEmail        *int    `json:"available_temp_email,omitempty"`
	BuyWithoutValidation      *int    `json:"buy_without_validation,omitempty"`
	CanBeResold               *int    `json:"can_be_resold,omitempty"`
	CategoryDescriptionHTML   *string `json:"category_description_html,omitempty"`
	CategoryDescriptionHTMLEn *string `json:"category_description_html_en,omitempty"`
	CategoryH1HTMLEn          *string `json:"category_h1_html_en,omitempty"`
	CategoryID                *int    `json:"category_id,omitempty"`
	CategoryLoginURL          *string `json:"category_login_url,omitempty"`
	CategoryName              *string `json:"category_name,omitempty"`
	CategoryOrder             *int    `json:"category_order,omitempty"`
	CategoryTitle             *string `json:"category_title,omitempty"`
	CategoryURL               *string `json:"category_url,omitempty"`
	CheckButtonEnabled        *int    `json:"check_button_enabled,omitempty"`
	CheckerEnabled            *int    `json:"checker_enabled,omitempty"`
	Cookies                   *string `json:"cookies,omitempty"`
	DisplayInList             *int    `json:"display_in_list,omitempty"`
	GuestHidden               *int    `json:"guest_hidden,omitempty"`
	HasAccountLink            *int    `json:"has_account_link,omitempty"`
	HasGuarantee              *int    `json:"has_guarantee,omitempty"`
	LoginType                 *string `json:"login_type,omitempty"`
	MassUploadItemAvailable   *int    `json:"mass_upload_item_available,omitempty"`
	MaxInvalidUploadTries     *int    `json:"max_invalid_upload_tries,omitempty"`
	RecoveryLink              *string `json:"recovery_link,omitempty"`
	RequireEldForNativeAccs   *int    `json:"require_eld_for_native_accs,omitempty"`
	RequireEmailLoginData     *int    `json:"require_email_login_data,omitempty"`
	RequireTempEmail          *int    `json:"require_temp_email,omitempty"`
	RequireVideoRecording     *int    `json:"require_video_recording,omitempty"`
	ResaleDurationLimitDays   *int    `json:"resale_duration_limit_days,omitempty"`
	SubCategoryID             *int    `json:"sub_category_id,omitempty"`
	SupportEmailLoginData     *int    `json:"support_email_login_data,omitempty"`
	SupportPersonalProxy      *int    `json:"support_personal_proxy,omitempty"`
	SupportTempEmail          *int    `json:"support_temp_email,omitempty"`
	TopQueries                *string `json:"top_queries,omitempty"`
}

type ParamsResponseParamsItem struct {
	Description *string     `json:"description,omitempty"`
	Input       *string     `json:"input,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Values      interface{} `json:"values,omitempty"`
}

type ParamsResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type PayoutParams struct {
	Extra      map[string]interface{} `json:"extra,omitempty"`
	IncludeFee *bool                  `json:"include_fee,omitempty"`
}

type PayoutServicesResponse struct {
	SystemInfo *RespSystemInfo                     `json:"system_info,omitempty"`
	Systems    []PayoutServicesResponseSystemsItem `json:"systems,omitempty"`
}

type PayoutServicesResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type PayoutServicesResponseSystemsItem struct {
	Commission        *string     `json:"commission,omitempty"`
	HasWallet         *bool       `json:"has_wallet,omitempty"`
	InstantPayout     *bool       `json:"instant_payout,omitempty"`
	IsUnavailable     *bool       `json:"is_unavailable,omitempty"`
	Max               *int        `json:"max,omitempty"`
	Min               *int        `json:"min,omitempty"`
	P2p               *bool       `json:"p2p,omitempty"`
	ProblematicPayout *bool       `json:"problematic_payout,omitempty"`
	Providers         interface{} `json:"providers,omitempty"`
	System            *string     `json:"system,omitempty"`
}

type PayoutServicesResponseSystemsItemProviders struct {
	BCH   *PayoutServicesResponseSystemsItemProvidersBCH   `json:"BCH,omitempty"`
	BEP20 *PayoutServicesResponseSystemsItemProvidersBEP20 `json:"BEP20,omitempty"`
	BNB   *PayoutServicesResponseSystemsItemProvidersBNB   `json:"BNB,omitempty"`
	BTC   *PayoutServicesResponseSystemsItemProvidersBTC   `json:"BTC,omitempty"`
	DASH  *PayoutServicesResponseSystemsItemProvidersDASH  `json:"DASH,omitempty"`
	DOGE  *PayoutServicesResponseSystemsItemProvidersDOGE  `json:"DOGE,omitempty"`
	ERC20 *PayoutServicesResponseSystemsItemProvidersERC20 `json:"ERC20,omitempty"`
	ETH   *PayoutServicesResponseSystemsItemProvidersETH   `json:"ETH,omitempty"`
	LTC   *PayoutServicesResponseSystemsItemProvidersLTC   `json:"LTC,omitempty"`
	SOL   *PayoutServicesResponseSystemsItemProvidersSOL   `json:"SOL,omitempty"`
	TON   *PayoutServicesResponseSystemsItemProvidersTON   `json:"TON,omitempty"`
	TRC20 *PayoutServicesResponseSystemsItemProvidersTRC20 `json:"TRC20,omitempty"`
	TRX   *PayoutServicesResponseSystemsItemProvidersTRX   `json:"TRX,omitempty"`
	XMR   *PayoutServicesResponseSystemsItemProvidersXMR   `json:"XMR,omitempty"`
}

type PayoutServicesResponseSystemsItemProvidersBCH struct {
	IsUnavailable *bool   `json:"isUnavailable,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type PayoutServicesResponseSystemsItemProvidersBEP20 struct {
	IsUnavailable *bool   `json:"isUnavailable,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type PayoutServicesResponseSystemsItemProvidersBNB struct {
	IsUnavailable *bool   `json:"isUnavailable,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type PayoutServicesResponseSystemsItemProvidersBTC struct {
	IsUnavailable *bool   `json:"isUnavailable,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type PayoutServicesResponseSystemsItemProvidersDASH struct {
	IsUnavailable *bool   `json:"isUnavailable,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type PayoutServicesResponseSystemsItemProvidersDOGE struct {
	IsUnavailable *bool   `json:"isUnavailable,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type PayoutServicesResponseSystemsItemProvidersERC20 struct {
	IsUnavailable *bool   `json:"isUnavailable,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type PayoutServicesResponseSystemsItemProvidersETH struct {
	IsUnavailable *bool   `json:"isUnavailable,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type PayoutServicesResponseSystemsItemProvidersLTC struct {
	IsUnavailable *bool   `json:"isUnavailable,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type PayoutServicesResponseSystemsItemProvidersSOL struct {
	IsUnavailable *bool   `json:"isUnavailable,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type PayoutServicesResponseSystemsItemProvidersTON struct {
	IsUnavailable *bool   `json:"isUnavailable,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type PayoutServicesResponseSystemsItemProvidersTRC20 struct {
	IsUnavailable *bool   `json:"isUnavailable,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type PayoutServicesResponseSystemsItemProvidersTRX struct {
	IsUnavailable *bool   `json:"isUnavailable,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type PayoutServicesResponseSystemsItemProvidersXMR struct {
	IsUnavailable *bool   `json:"isUnavailable,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type RespSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type RiotParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// Minimum valorant rank.
	Rmin *int `json:"rmin,omitempty"`
	// Maximum valorant rank.
	Rmax *int `json:"rmax,omitempty"`
	// Last minimum valorant rank.
	LastRmin *int `json:"last_rmin,omitempty"`
	// Last maximum valorant rank.
	LastRmax *int `json:"last_rmax,omitempty"`
	// Previous minimum rank.
	PreviousRmin *int `json:"previous_rmin,omitempty"`
	// Previous maximum rank.
	PreviousRmax *int `json:"previous_rmax,omitempty"`
	// List of weapon skins.
	WeaponSkin []string `json:"weaponSkin[],omitempty"`
	// List of buddies.
	Buddy []string `json:"buddy[],omitempty"`
	// List of agents.
	Agent []string `json:"agent[],omitempty"`
	// List of champions.
	Champion []string `json:"champion[],omitempty"`
	// List of LoL skins.
	Skin []string `json:"skin[],omitempty"`
	// List of allowed countries.
	Country []string `json:"country[],omitempty"`
	// List of disallowed countries.
	NotCountry []string `json:"not_country[],omitempty"`
	// Number of days the account has been offline.
	Daybreak *int `json:"daybreak,omitempty"`
	// Minimum level in Valorant.
	ValorantLevelMin *int `json:"valorant_level_min,omitempty"`
	// Maximum level in Valorant.
	ValorantLevelMax *int `json:"valorant_level_max,omitempty"`
	// Minimum level in LoL.
	LolLevelMin *int `json:"lol_level_min,omitempty"`
	// Maximum level in LoL.
	LolLevelMax *int `json:"lol_level_max,omitempty"`
	// Minimum inventory value.
	InvMin *int `json:"inv_min,omitempty"`
	// Maximum inventory value.
	InvMax *int `json:"inv_max,omitempty"`
	// Minimum number of Valorant points.
	VpMin *int `json:"vp_min,omitempty"`
	// Maximum number of Valorant points.
	VpMax *int `json:"vp_max,omitempty"`
	// Minimum number of skins.
	ValorantSmin *int `json:"valorant_smin,omitempty"`
	// Maximum number of skins.
	ValorantSmax *int `json:"valorant_smax,omitempty"`
	// List of allowed rank types.
	ValorantRankType []string `json:"valorant_rank_type[],omitempty"`
	// Minimum amount of agents.
	Amin *int `json:"amin,omitempty"`
	// Maximum amount of agents.
	Amax *int `json:"amax,omitempty"`
	// List of allowed regions in Valorant.
	ValorantRegion []string `json:"valorant_region[],omitempty"`
	// List of disallowed regions in Valorant.
	ValorantNotRegion []string `json:"valorant_not_region[],omitempty"`
	// List of allowed regions in LoL.
	LolRegion []string `json:"lol_region[],omitempty"`
	// List of disallowed regions in LoL.
	LolNotRegion []string `json:"lol_not_region[],omitempty"`
	// Has any knife.
	Knife *bool `json:"knife,omitempty"`
	// Minimum number of skins in LoL.
	LolSmin *int `json:"lol_smin,omitempty"`
	// Maximum number of skins in LoL.
	LolSmax *int `json:"lol_smax,omitempty"`
	// Minimum number of champions.
	ChampionMin *int `json:"champion_min,omitempty"`
	// Maximum number of champions.
	ChampionMax *int `json:"champion_max,omitempty"`
	// Minimum win-rate.
	WinRateMin *int `json:"win_rate_min,omitempty"`
	// Maximum win-rate.
	WinRateMax *int `json:"win_rate_max,omitempty"`
	// Minimum wallet blue balance.
	BlueMin *int `json:"blue_min,omitempty"`
	// Maximum wallet blue balance.
	BlueMax *int `json:"blue_max,omitempty"`
	// Minimum wallet orange balance.
	OrangeMin *int `json:"orange_min,omitempty"`
	// Maximum wallet orange balance.
	OrangeMax *int `json:"orange_max,omitempty"`
	// Minimum wallet mythic balance.
	MythicMin *int `json:"mythic_min,omitempty"`
	// Maximum wallet mythic balance.
	MythicMax *int `json:"mythic_max,omitempty"`
	// Minimum wallet riot balance.
	RiotMin *int `json:"riot_min,omitempty"`
	// Maximum wallet riot balance.
	RiotMax *int `json:"riot_max,omitempty"`
	// Has linked email.
	Email *CategorySearchEmail `json:"email,omitempty"`
	// Has linked mobile.
	Tel *CategorySearchTel `json:"tel,omitempty"`
	// Minimum knifes in Valorant.
	ValorantKnifeMin *int `json:"valorant_knife_min,omitempty"`
	// Maximum knifes in Valorant.
	ValorantKnifeMax *int `json:"valorant_knife_max,omitempty"`
	// Minimum number of Valorant Radiant Points.
	RpMin *int `json:"rp_min,omitempty"`
	// Maximum number of Valorant Radiant Points.
	RpMax *int `json:"rp_max,omitempty"`
	// Minimum number of Valorant free agents.
	FaMin *int `json:"fa_min,omitempty"`
	// Maximum number of Valorant free agents.
	FaMax *int `json:"fa_max,omitempty"`
	// List of allowed ranks in LoL.
	LolRank []string `json:"lol_rank[],omitempty"`
}

type RiotResponse struct {
	CacheTTL        *int                    `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                   `json:"hasNextPage,omitempty"`
	Items           []RiotResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                    `json:"lastModified,omitempty"`
	Page            *int                    `json:"page,omitempty"`
	PerPage         *int                    `json:"perPage,omitempty"`
	SearchUrl       *string                 `json:"searchUrl,omitempty"`
	ServerTime      *int                    `json:"serverTime,omitempty"`
	StickyItems     []interface{}           `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo         `json:"system_info,omitempty"`
	TotalItems      *int                    `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}             `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                   `json:"wasCached,omitempty"`
}

type RiotResponseItemsItem struct {
	AccountLink                *string                                 `json:"accountLink,omitempty"`
	AccountLinks               []RiotResponseItemsItemAccountLinksItem `json:"accountLinks,omitempty"`
	AccountLastActivity        *int                                    `json:"account_last_activity,omitempty"`
	AllowAskDiscount           *int                                    `json:"allow_ask_discount,omitempty"`
	BumpSettings               *RiotResponseItemsItemBumpSettings      `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                                   `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                                   `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                                   `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                                   `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                                   `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                                   `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                                   `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                                   `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                                   `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                                   `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                                   `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                                   `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                                   `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                                   `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                                   `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                                   `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                    `json:"category_id,omitempty"`
	Description                *string                                 `json:"description,omitempty"`
	DescriptionEnHtml          *string                                 `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                                 `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                                 `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                                 `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                                 `json:"description_en,omitempty"`
	EditDate                   *int                                    `json:"edit_date,omitempty"`
	EmailLoginUrl              *string                                 `json:"emailLoginUrl,omitempty"`
	EmailProvider              *string                                 `json:"email_provider,omitempty"`
	EmailType                  *string                                 `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                                    `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                             `json:"feedback_data,omitempty"`
	Guarantee                  interface{}                             `json:"guarantee,omitempty"`
	HasPendingAutoBuy          *bool                                   `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                  *bool                                   `json:"isIgnored,omitempty"`
	IsSmallExf                 *bool                                   `json:"isSmallExf,omitempty"`
	IsSticky                   *int                                    `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                                 `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                                 `json:"item_domain,omitempty"`
	ItemID                     *int                                    `json:"item_id,omitempty"`
	ItemOrigin                 *string                                 `json:"item_origin,omitempty"`
	ItemState                  *string                                 `json:"item_state,omitempty"`
	LolInventory               interface{}                             `json:"lolInventory,omitempty"`
	LolRegionPhrase            *string                                 `json:"lolRegionPhrase,omitempty"`
	NoteText                   interface{}                             `json:"note_text,omitempty"`
	Nsb                        *int                                    `json:"nsb,omitempty"`
	Price                      *int                                    `json:"price,omitempty"`
	PriceWithSellerFee         *float64                                `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                                 `json:"price_currency,omitempty"`
	PublishedDate              *int                                    `json:"published_date,omitempty"`
	RefreshedDate              *int                                    `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                                 `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount  *int                                    `json:"restore_items_category_count,omitempty"`
	RiotAccountVerified        *int                                    `json:"riot_account_verified,omitempty"`
	RiotCountry                *string                                 `json:"riot_country,omitempty"`
	RiotEmailVerified          *int                                    `json:"riot_email_verified,omitempty"`
	RiotID                     *string                                 `json:"riot_id,omitempty"`
	RiotItemID                 *int                                    `json:"riot_item_id,omitempty"`
	RiotLastActivity           *int                                    `json:"riot_last_activity,omitempty"`
	RiotLolChampionCount       *int                                    `json:"riot_lol_champion_count,omitempty"`
	RiotLolLevel               *int                                    `json:"riot_lol_level,omitempty"`
	RiotLolRank                *string                                 `json:"riot_lol_rank,omitempty"`
	RiotLolRankWinRate         *int                                    `json:"riot_lol_rank_win_rate,omitempty"`
	RiotLolRegion              *string                                 `json:"riot_lol_region,omitempty"`
	RiotLolSkinCount           *int                                    `json:"riot_lol_skin_count,omitempty"`
	RiotLolWalletBlue          *int                                    `json:"riot_lol_wallet_blue,omitempty"`
	RiotLolWalletMythic        *int                                    `json:"riot_lol_wallet_mythic,omitempty"`
	RiotLolWalletOrange        *int                                    `json:"riot_lol_wallet_orange,omitempty"`
	RiotLolWalletRiot          *int                                    `json:"riot_lol_wallet_riot,omitempty"`
	RiotPasswordChange         *int                                    `json:"riot_password_change,omitempty"`
	RiotPhoneVerified          *int                                    `json:"riot_phone_verified,omitempty"`
	RiotUsername               *string                                 `json:"riot_username,omitempty"`
	RiotValorantAgentCount     *int                                    `json:"riot_valorant_agent_count,omitempty"`
	RiotValorantInventoryValue *int                                    `json:"riot_valorant_inventory_value,omitempty"`
	RiotValorantKnife          *int                                    `json:"riot_valorant_knife,omitempty"`
	RiotValorantLastRank       *int                                    `json:"riot_valorant_last_rank,omitempty"`
	RiotValorantLevel          *int                                    `json:"riot_valorant_level,omitempty"`
	RiotValorantPreviousRank   *int                                    `json:"riot_valorant_previous_rank,omitempty"`
	RiotValorantRank           *int                                    `json:"riot_valorant_rank,omitempty"`
	RiotValorantRankType       *string                                 `json:"riot_valorant_rank_type,omitempty"`
	RiotValorantRegion         *string                                 `json:"riot_valorant_region,omitempty"`
	RiotValorantSkinCount      *int                                    `json:"riot_valorant_skin_count,omitempty"`
	RiotValorantWalletFa       *int                                    `json:"riot_valorant_wallet_fa,omitempty"`
	RiotValorantWalletRp       *int                                    `json:"riot_valorant_wallet_rp,omitempty"`
	RiotValorantWalletVp       *int                                    `json:"riot_valorant_wallet_vp,omitempty"`
	RubPrice                   *int                                    `json:"rub_price,omitempty"`
	Seller                     *RiotResponseItemsItemSeller            `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                                   `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount     *int                                    `json:"sold_items_category_count,omitempty"`
	Tags                       interface{}                             `json:"tags,omitempty"`
	Title                      *string                                 `json:"title,omitempty"`
	TitleEn                    *string                                 `json:"title_en,omitempty"`
	UpdateStatDate             *int                                    `json:"update_stat_date,omitempty"`
	ValorantInventory          interface{}                             `json:"valorantInventory,omitempty"`
	ValorantLastRankTitle      *string                                 `json:"valorantLastRankTitle,omitempty"`
	ValorantPreviousRankTitle  *string                                 `json:"valorantPreviousRankTitle,omitempty"`
	ValorantRankImgPath        *string                                 `json:"valorantRankImgPath,omitempty"`
	ValorantRankTitle          *string                                 `json:"valorantRankTitle,omitempty"`
	ValorantRegionPhrase       *string                                 `json:"valorantRegionPhrase,omitempty"`
	ViewCount                  *int                                    `json:"view_count,omitempty"`
}

type RiotResponseItemsItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type RiotResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type RiotResponseItemsItemLolInventory struct {
	Champion []int       `json:"Champion,omitempty"`
	Skin     interface{} `json:"Skin,omitempty"`
}

type RiotResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type RiotResponseItemsItemValorantInventory struct {
	Agent       []string    `json:"Agent,omitempty"`
	Buddy       []string    `json:"Buddy,omitempty"`
	WeaponSkins interface{} `json:"WeaponSkins,omitempty"`
}

type RiotResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type RobloxParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Has verified email.
	Email *CategorySearchEmail `json:"email,omitempty"`
	// Minimum robux.
	RobuxMin *int `json:"robux_min,omitempty"`
	// Maximum robux.
	RobuxMax *int `json:"robux_max,omitempty"`
	// Minimum friends.
	FriendsMin *int `json:"friends_min,omitempty"`
	// Maximum friends.
	FriendsMax *int `json:"friends_max,omitempty"`
	// Minimum number of followers.
	FollowersMin *int `json:"followers_min,omitempty"`
	// Maximum number of followers.
	FollowersMax *int `json:"followers_max,omitempty"`
	// List of allowed countries.
	Country []string `json:"country,omitempty"`
	// List of disallowed countries.
	NotCountry []string `json:"not_country,omitempty"`
	// How old is the account.
	Reg *int `json:"reg,omitempty"`
	// In what notation is time measured.
	RegPeriod *CategorySearchRegPeriod `json:"reg_period,omitempty"`
	// Name of subscription.
	Subscription *CategorySearchSubscription `json:"subscription,omitempty"`
	// Length of subscription.
	SubscriptionLength *int `json:"subscription_length,omitempty"`
	// In what notation is time measured.
	SubscriptionPeriod *CategorySearchSubscriptionPeriod `json:"subscription_period,omitempty"`
	// Is auto renewal enabled.
	Autorenewal *CategorySearchAutorenewal `json:"autorenewal,omitempty"`
	// Xbox connected to account.
	XboxConnected *CategorySearchXboxConnected `json:"xbox_connected,omitempty"`
	// PSN connected to account.
	PsnConnected *CategorySearchPsnConnected `json:"psn_connected,omitempty"`
	// Has verified.
	Verified *CategorySearchVerified `json:"verified,omitempty"`
	// Account is age verified via documents.
	AgeVerified *CategorySearchAgeVerified `json:"age_verified,omitempty"`
	// Minimum amount of incoming robux.
	IncomingRobuxTotalMin *int `json:"incoming_robux_total_min,omitempty"`
	// Maximum amount of incoming robux.
	IncomingRobuxTotalMax *int `json:"incoming_robux_total_max,omitempty"`
	// Minimum limited items value.
	LimitedPriceMin *int `json:"limited_price_min,omitempty"`
	// Maximum limited items value.
	LimitedPriceMax *int `json:"limited_price_max,omitempty"`
	// Minimum total Robux cost of all game passes in popular Roblox games..
	GamepassMin *int `json:"gamepass_min,omitempty"`
	// Maximum total Robux cost of all game passes in popular Roblox games..
	GamepassMax *int `json:"gamepass_max,omitempty"`
	// Has game donations.
	GameDonations *CategorySearchGameDonations `json:"game_donations,omitempty"`
	// Minimum inventory value.
	InvMin *int `json:"inv_min,omitempty"`
	// Maximum inventory value.
	InvMax *int `json:"inv_max,omitempty"`
	// Minimum UGC limited items value.
	UgcLimitedPriceMin *int `json:"ugc_limited_price_min,omitempty"`
	// Maximum UGC limited items value.
	UgcLimitedPriceMax *int `json:"ugc_limited_price_max,omitempty"`
	// Minimum credit balance.
	CreditBalanceMin *int `json:"credit_balance_min,omitempty"`
	// Maximum credit balance.
	CreditBalanceMax *int `json:"credit_balance_max,omitempty"`
	// Minimum offsale items count.
	OffsaleMin *int `json:"offsale_min,omitempty"`
	// Maximum offsale items count.
	OffsaleMax *int `json:"offsale_max,omitempty"`
	// Voice chat is available.
	Voice *CategorySearchVoice `json:"voice,omitempty"`
	// List of allowed age groups.
	AgeGroup []string `json:"age_group[],omitempty"`
	// List of disallowed age groups.
	NotAgeGroup []string `json:"not_age_group[],omitempty"`
}

type RobloxResponse struct {
	CacheTTL        *int                      `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                     `json:"hasNextPage,omitempty"`
	Items           []RobloxResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                      `json:"lastModified,omitempty"`
	Page            *int                      `json:"page,omitempty"`
	PerPage         *int                      `json:"perPage,omitempty"`
	SearchUrl       *string                   `json:"searchUrl,omitempty"`
	ServerTime      *int                      `json:"serverTime,omitempty"`
	StickyItems     []interface{}             `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo           `json:"system_info,omitempty"`
	TotalItems      *int                      `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}               `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                     `json:"wasCached,omitempty"`
}

type RobloxResponseItemsItem struct {
	AccountLink                 *string                                                 `json:"accountLink,omitempty"`
	AccountLinks                []RobloxResponseItemsItemAccountLinksItem               `json:"accountLinks,omitempty"`
	AllowAskDiscount            *int                                                    `json:"allow_ask_discount,omitempty"`
	BumpSettings                *RobloxResponseItemsItemBumpSettings                    `json:"bumpSettings,omitempty"`
	CanBumpItem                 *bool                                                   `json:"canBumpItem,omitempty"`
	CanBuyItem                  *bool                                                   `json:"canBuyItem,omitempty"`
	CanChangePassword           *bool                                                   `json:"canChangePassword,omitempty"`
	CanCloseItem                *bool                                                   `json:"canCloseItem,omitempty"`
	CanDeleteItem               *bool                                                   `json:"canDeleteItem,omitempty"`
	CanEditItem                 *bool                                                   `json:"canEditItem,omitempty"`
	CanOpenItem                 *bool                                                   `json:"canOpenItem,omitempty"`
	CanReportItem               *bool                                                   `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase  *bool                                                   `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem                *bool                                                   `json:"canStickItem,omitempty"`
	CanUnstickItem              *bool                                                   `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats          *bool                                                   `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount          *bool                                                   `json:"canValidateAccount,omitempty"`
	CanViewAccountLink          *bool                                                   `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData       *bool                                                   `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData            *bool                                                   `json:"canViewLoginData,omitempty"`
	CategoryID                  *int                                                    `json:"category_id,omitempty"`
	CreditBalance               *string                                                 `json:"creditBalance,omitempty"`
	Description                 *string                                                 `json:"description,omitempty"`
	DescriptionEnHtml           *string                                                 `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain          *string                                                 `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml             *string                                                 `json:"descriptionHtml,omitempty"`
	DescriptionPlain            *string                                                 `json:"descriptionPlain,omitempty"`
	DescriptionEn               *string                                                 `json:"description_en,omitempty"`
	EditDate                    *int                                                    `json:"edit_date,omitempty"`
	EmailLoginUrl               *string                                                 `json:"emailLoginUrl,omitempty"`
	EmailProvider               *string                                                 `json:"email_provider,omitempty"`
	EmailType                   *string                                                 `json:"email_type,omitempty"`
	ExtendedGuarantee           *int                                                    `json:"extended_guarantee,omitempty"`
	FeedbackData                interface{}                                             `json:"feedback_data,omitempty"`
	Guarantee                   interface{}                                             `json:"guarantee,omitempty"`
	HasPendingAutoBuy           *bool                                                   `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                   *bool                                                   `json:"isIgnored,omitempty"`
	IsSticky                    *int                                                    `json:"is_sticky,omitempty"`
	ItemOriginPhrase            *string                                                 `json:"itemOriginPhrase,omitempty"`
	ItemDomain                  *string                                                 `json:"item_domain,omitempty"`
	ItemID                      *int                                                    `json:"item_id,omitempty"`
	ItemOrigin                  *string                                                 `json:"item_origin,omitempty"`
	ItemState                   *string                                                 `json:"item_state,omitempty"`
	NoteText                    interface{}                                             `json:"note_text,omitempty"`
	Nsb                         *int                                                    `json:"nsb,omitempty"`
	Price                       *int                                                    `json:"price,omitempty"`
	PriceWithSellerFee          *float64                                                `json:"priceWithSellerFee,omitempty"`
	PriceCurrency               *string                                                 `json:"price_currency,omitempty"`
	PublishedDate               *int                                                    `json:"published_date,omitempty"`
	RefreshedDate               *int                                                    `json:"refreshed_date,omitempty"`
	ResaleItemOrigin            *string                                                 `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount   *int                                                    `json:"restore_items_category_count,omitempty"`
	RobloxGameDonations         []RobloxResponseItemsItemRobloxGameDonationsItem        `json:"robloxGameDonations,omitempty"`
	RobloxGameDonationsDetails  []RobloxResponseItemsItemRobloxGameDonationsDetailsItem `json:"robloxGameDonationsDetails,omitempty"`
	RobloxLinkedAccounts        *string                                                 `json:"robloxLinkedAccounts,omitempty"`
	RobloxAgeVerified           *int                                                    `json:"roblox_age_verified,omitempty"`
	RobloxCountry               *string                                                 `json:"roblox_country,omitempty"`
	RobloxCreditBalance         *float64                                                `json:"roblox_credit_balance,omitempty"`
	RobloxEmailVerified         *int                                                    `json:"roblox_email_verified,omitempty"`
	RobloxFollowers             *int                                                    `json:"roblox_followers,omitempty"`
	RobloxFriends               *int                                                    `json:"roblox_friends,omitempty"`
	RobloxGamePassTotalRobux    *int                                                    `json:"roblox_game_pass_total_robux,omitempty"`
	RobloxID                    *int                                                    `json:"roblox_id,omitempty"`
	RobloxIncomingRobuxTotal    *int                                                    `json:"roblox_incoming_robux_total,omitempty"`
	RobloxInventoryPrice        *int                                                    `json:"roblox_inventory_price,omitempty"`
	RobloxItemID                *int                                                    `json:"roblox_item_id,omitempty"`
	RobloxLimitedPrice          *int                                                    `json:"roblox_limited_price,omitempty"`
	RobloxPsnConnected          *int                                                    `json:"roblox_psn_connected,omitempty"`
	RobloxRegisterDate          *int                                                    `json:"roblox_register_date,omitempty"`
	RobloxRobux                 *int                                                    `json:"roblox_robux,omitempty"`
	RobloxSubscription          *string                                                 `json:"roblox_subscription,omitempty"`
	RobloxSubscriptionAutoRenew *int                                                    `json:"roblox_subscription_auto_renew,omitempty"`
	RobloxSubscriptionEndDate   *int                                                    `json:"roblox_subscription_end_date,omitempty"`
	RobloxUgcLimitedPrice       *int                                                    `json:"roblox_ugc_limited_price,omitempty"`
	RobloxUsername              *string                                                 `json:"roblox_username,omitempty"`
	RobloxVerified              *int                                                    `json:"roblox_verified,omitempty"`
	RobloxXboxConnected         *int                                                    `json:"roblox_xbox_connected,omitempty"`
	RubPrice                    *int                                                    `json:"rub_price,omitempty"`
	Seller                      *RobloxResponseItemsItemSeller                          `json:"seller,omitempty"`
	ShowGetEmailCodeButton      *bool                                                   `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount      *int                                                    `json:"sold_items_category_count,omitempty"`
	Tags                        interface{}                                             `json:"tags,omitempty"`
	Title                       *string                                                 `json:"title,omitempty"`
	TitleEn                     *string                                                 `json:"title_en,omitempty"`
	UpdateStatDate              *int                                                    `json:"update_stat_date,omitempty"`
	ViewCount                   *int                                                    `json:"view_count,omitempty"`
}

type RobloxResponseItemsItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type RobloxResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type RobloxResponseItemsItemRobloxGameDonationsDetailsItem struct {
	Amount  *int    `json:"amount,omitempty"`
	Product *string `json:"product,omitempty"`
	Type_   *string `json:"type,omitempty"`
}

type RobloxResponseItemsItemRobloxGameDonationsItem struct {
	Amount *int    `json:"amount,omitempty"`
	ID     *int    `json:"id,omitempty"`
	Title  *string `json:"title,omitempty"`
}

type RobloxResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type RobloxResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type SaveChanges struct {
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type SaveChangesSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type SocialClubParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Number of days the account has been offline.
	Daybreak *int `json:"daybreak,omitempty"`
	// Minimum number of Social Club level.
	LevelMin *int `json:"level_min,omitempty"`
	// Maximum number of Social Club level.
	LevelMax *int `json:"level_max,omitempty"`
	// Minimum number of GTA V cash
	CashMin *int `json:"cash_min,omitempty"`
	// Maximum number of GTA V cash
	CashMax *int `json:"cash_max,omitempty"`
	// Minimum number of GTA V bank cash
	BankCashMin *int `json:"bank_cash_min,omitempty"`
	// Maximum number of GTA V bank cash
	BankCashMax *int `json:"bank_cash_max,omitempty"`
	// List of games.
	Game []string `json:"game[],omitempty"`
}

type SocialClubResponse struct {
	CacheTTL        *int                          `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                         `json:"hasNextPage,omitempty"`
	Items           []SocialClubResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                          `json:"lastModified,omitempty"`
	Page            *int                          `json:"page,omitempty"`
	PerPage         *int                          `json:"perPage,omitempty"`
	SearchUrl       *string                       `json:"searchUrl,omitempty"`
	ServerTime      *int                          `json:"serverTime,omitempty"`
	StickyItems     []interface{}                 `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo               `json:"system_info,omitempty"`
	TotalItems      *int                          `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}                   `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                         `json:"wasCached,omitempty"`
}

type SocialClubResponseItemsItem struct {
	AccountLinks               []interface{}                            `json:"accountLinks,omitempty"`
	AccountLastActivity        *int                                     `json:"account_last_activity,omitempty"`
	AllowAskDiscount           *int                                     `json:"allow_ask_discount,omitempty"`
	BumpSettings               *SocialClubResponseItemsItemBumpSettings `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                                    `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                                    `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                                    `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                                    `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                                    `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                                    `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                                    `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                                    `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                                    `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                                    `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                                    `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                                    `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                                    `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                                    `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                                    `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                                    `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                     `json:"category_id,omitempty"`
	Description                *string                                  `json:"description,omitempty"`
	DescriptionEnHtml          *string                                  `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                                  `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                                  `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                                  `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                                  `json:"description_en,omitempty"`
	EditDate                   *int                                     `json:"edit_date,omitempty"`
	EmailLoginUrl              *string                                  `json:"emailLoginUrl,omitempty"`
	EmailProvider              *string                                  `json:"email_provider,omitempty"`
	EmailType                  *string                                  `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                                     `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                              `json:"feedback_data,omitempty"`
	Guarantee                  interface{}                              `json:"guarantee,omitempty"`
	HasPendingAutoBuy          *bool                                    `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                  *bool                                    `json:"isIgnored,omitempty"`
	IsSmallExf                 *bool                                    `json:"isSmallExf,omitempty"`
	IsSticky                   *int                                     `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                                  `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                                  `json:"item_domain,omitempty"`
	ItemID                     *int                                     `json:"item_id,omitempty"`
	ItemOrigin                 *string                                  `json:"item_origin,omitempty"`
	ItemState                  *string                                  `json:"item_state,omitempty"`
	NoteText                   interface{}                              `json:"note_text,omitempty"`
	Nsb                        *int                                     `json:"nsb,omitempty"`
	Price                      *int                                     `json:"price,omitempty"`
	PriceWithSellerFee         *float64                                 `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                                  `json:"price_currency,omitempty"`
	PublishedDate              *int                                     `json:"published_date,omitempty"`
	RefreshedDate              *int                                     `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                                  `json:"resale_item_origin,omitempty"`
	RubPrice                   *int                                     `json:"rub_price,omitempty"`
	Seller                     *SocialClubResponseItemsItemSeller       `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                                    `json:"showGetEmailCodeButton,omitempty"`
	SocialclubBankCash         *int                                     `json:"socialclub_bank_cash,omitempty"`
	SocialclubCash             *int                                     `json:"socialclub_cash,omitempty"`
	SocialclubGames            interface{}                              `json:"socialclub_games,omitempty"`
	SocialclubHasGtav          *int                                     `json:"socialclub_has_gtav,omitempty"`
	SocialclubHasRdr2          *int                                     `json:"socialclub_has_rdr2,omitempty"`
	SocialclubItemID           *int                                     `json:"socialclub_item_id,omitempty"`
	SocialclubLastActivity     *int                                     `json:"socialclub_last_activity,omitempty"`
	SocialclubLevel            *int                                     `json:"socialclub_level,omitempty"`
	Tags                       interface{}                              `json:"tags,omitempty"`
	Title                      *string                                  `json:"title,omitempty"`
	TitleEn                    *string                                  `json:"title_en,omitempty"`
	UpdateStatDate             *int                                     `json:"update_stat_date,omitempty"`
	ViewCount                  *int                                     `json:"view_count,omitempty"`
}

type SocialClubResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type SocialClubResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     interface{} `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type SocialClubResponseItemsItemSocialclubGamesItem struct {
	Abbr            *string     `json:"abbr,omitempty"`
	AppID           *string     `json:"app_id,omitempty"`
	CategoryID      *int        `json:"category_id,omitempty"`
	DefaultPlatform *string     `json:"defaultPlatform,omitempty"`
	ID              *int        `json:"id,omitempty"`
	Img             *string     `json:"img,omitempty"`
	InternalGameID  *int        `json:"internal_game_id,omitempty"`
	LastSeen        *string     `json:"lastSeen,omitempty"`
	Name            *string     `json:"name,omitempty"`
	Platform        *string     `json:"platform,omitempty"`
	Ru              interface{} `json:"ru,omitempty"`
	Title           *string     `json:"title,omitempty"`
	URL             *string     `json:"url,omitempty"`
}

type SocialClubResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type StatesParams struct {
	// User ID.
	UserID StringOrInt `json:"user_id,omitempty"`
}

type StatesResponse struct {
	SystemInfo     *RespSystemInfo               `json:"system_info,omitempty"`
	UserItemStates *StatesResponseUserItemStates `json:"userItemStates,omitempty"`
}

type StatesResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type StatesResponseUserItemStates struct {
	Active            *StatesResponseUserItemStatesActive            `json:"active,omitempty"`
	AutoBump          *StatesResponseUserItemStatesAutoBump          `json:"auto_bump,omitempty"`
	Awaiting          *StatesResponseUserItemStatesAwaiting          `json:"awaiting,omitempty"`
	Closed            *StatesResponseUserItemStatesClosed            `json:"closed,omitempty"`
	ClosedInactive    *StatesResponseUserItemStatesClosedInactive    `json:"closed_inactive,omitempty"`
	Deleted           *StatesResponseUserItemStatesDeleted           `json:"deleted,omitempty"`
	DiscountRequest   *StatesResponseUserItemStatesDiscountRequest   `json:"discount_request,omitempty"`
	InBuyersFavorites *StatesResponseUserItemStatesInBuyersFavorites `json:"in_buyers_favorites,omitempty"`
	Paid              *StatesResponseUserItemStatesPaid              `json:"paid,omitempty"`
	PendingDeletion   *StatesResponseUserItemStatesPendingDeletion   `json:"pending_deletion,omitempty"`
	PreActive         *StatesResponseUserItemStatesPreActive         `json:"pre_active,omitempty"`
	PreUpload         *StatesResponseUserItemStatesPreUpload         `json:"pre_upload,omitempty"`
	Stickied          *StatesResponseUserItemStatesStickied          `json:"stickied,omitempty"`
}

type StatesResponseUserItemStatesActive struct {
	ItemCount *int    `json:"item_count,omitempty"`
	ItemState *string `json:"item_state,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type StatesResponseUserItemStatesAutoBump struct {
	ItemCount *int    `json:"item_count,omitempty"`
	ItemState *string `json:"item_state,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type StatesResponseUserItemStatesAwaiting struct {
	ItemCount *int    `json:"item_count,omitempty"`
	ItemState *string `json:"item_state,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type StatesResponseUserItemStatesClosed struct {
	ItemCount *int    `json:"item_count,omitempty"`
	ItemState *string `json:"item_state,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type StatesResponseUserItemStatesClosedInactive struct {
	ItemCount *int    `json:"item_count,omitempty"`
	ItemState *string `json:"item_state,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type StatesResponseUserItemStatesDeleted struct {
	ItemCount *int    `json:"item_count,omitempty"`
	ItemState *string `json:"item_state,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type StatesResponseUserItemStatesDiscountRequest struct {
	ItemCount *int    `json:"item_count,omitempty"`
	ItemState *string `json:"item_state,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type StatesResponseUserItemStatesInBuyersFavorites struct {
	ItemCount *int    `json:"item_count,omitempty"`
	ItemState *string `json:"item_state,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type StatesResponseUserItemStatesPaid struct {
	ItemCount *int    `json:"item_count,omitempty"`
	ItemState *string `json:"item_state,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type StatesResponseUserItemStatesPendingDeletion struct {
	ItemCount *int    `json:"item_count,omitempty"`
	ItemState *string `json:"item_state,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type StatesResponseUserItemStatesPreActive struct {
	ItemCount *int    `json:"item_count,omitempty"`
	ItemState *string `json:"item_state,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type StatesResponseUserItemStatesPreUpload struct {
	ItemCount *int    `json:"item_count,omitempty"`
	ItemState *string `json:"item_state,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type StatesResponseUserItemStatesStickied struct {
	ItemCount   *int    `json:"item_count,omitempty"`
	ItemState   *string `json:"item_state,omitempty"`
	StickyLimit *int    `json:"stickyLimit,omitempty"`
	Title       *string `json:"title,omitempty"`
}

type Status struct {
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type StatusItem struct {
	Item       *ItemModel      `json:"item,omitempty"`
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type StatusItemItem struct {
	AccountLink                       *string                          `json:"accountLink,omitempty"`
	AccountLinks                      []StatusItemItemAccountLinksItem `json:"accountLinks,omitempty"`
	AccountLastActivity               *int                             `json:"account_last_activity,omitempty"`
	AiPrice                           *int                             `json:"aiPrice,omitempty"`
	AiPriceCheckDate                  *int                             `json:"aiPriceCheckDate,omitempty"`
	AllowAskDiscount                  *int                             `json:"allow_ask_discount,omitempty"`
	AutoBuyPrice                      *int                             `json:"autoBuyPrice,omitempty"`
	AutoBuyPriceCheckDate             *int                             `json:"autoBuyPriceCheckDate,omitempty"`
	BumpSettings                      *StatusItemItemBumpSettings      `json:"bumpSettings,omitempty"`
	Buyer                             *StatusItemItemBuyer             `json:"buyer,omitempty"`
	BuyerAvatarDate                   *int                             `json:"buyer_avatar_date,omitempty"`
	BuyerDisplayIconGroupID           *int                             `json:"buyer_display_icon_group_id,omitempty"`
	BuyerUniqBanner                   *string                          `json:"buyer_uniq_banner,omitempty"`
	BuyerUserGroupID                  *int                             `json:"buyer_user_group_id,omitempty"`
	CanAskDiscount                    *bool                            `json:"canAskDiscount,omitempty"`
	CanChangeEmailPassword            *bool                            `json:"canChangeEmailPassword,omitempty"`
	CanChangePassword                 *bool                            `json:"canChangePassword,omitempty"`
	CanCheckAiPrice                   *bool                            `json:"canCheckAiPrice,omitempty"`
	CanCheckAutoBuyPrice              *bool                            `json:"canCheckAutoBuyPrice,omitempty"`
	CanCheckGuarantee                 *bool                            `json:"canCheckGuarantee,omitempty"`
	CanReportItem                     *bool                            `json:"canReportItem,omitempty"`
	CanResellItem                     *bool                            `json:"canResellItem,omitempty"`
	CanResellItemAfterPurchase        *bool                            `json:"canResellItemAfterPurchase,omitempty"`
	CanShareItem                      *bool                            `json:"canShareItem,omitempty"`
	CanUpdateItemStats                *bool                            `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount                *bool                            `json:"canValidateAccount,omitempty"`
	CanViewAccountLink                *bool                            `json:"canViewAccountLink,omitempty"`
	CanViewAccountLoginAndTempEmail   *bool                            `json:"canViewAccountLoginAndTempEmail,omitempty"`
	CanViewEmailLoginData             *bool                            `json:"canViewEmailLoginData,omitempty"`
	CanViewItemViews                  *bool                            `json:"canViewItemViews,omitempty"`
	CanViewLoginData                  *bool                            `json:"canViewLoginData,omitempty"`
	CartPrice                         interface{}                      `json:"cart_price,omitempty"`
	CategoryID                        *int                             `json:"category_id,omitempty"`
	ContentID                         interface{}                      `json:"content_id,omitempty"`
	ContentType                       interface{}                      `json:"content_type,omitempty"`
	CopyFormatData                    *StatusItemItemCopyFormatData    `json:"copyFormatData,omitempty"`
	CustomFields                      *StatusItemItemCustomFields      `json:"customFields,omitempty"`
	DeleteDate                        *int                             `json:"delete_date,omitempty"`
	DeleteReason                      *string                          `json:"delete_reason,omitempty"`
	DeleteUserID                      *int                             `json:"delete_user_id,omitempty"`
	DeleteUsername                    *string                          `json:"delete_username,omitempty"`
	Deposit                           *int                             `json:"deposit,omitempty"`
	Description                       *string                          `json:"description,omitempty"`
	DescriptionEnHtml                 *string                          `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain                *string                          `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml                   *string                          `json:"descriptionHtml,omitempty"`
	DescriptionPlain                  *string                          `json:"descriptionPlain,omitempty"`
	DescriptionEn                     *string                          `json:"description_en,omitempty"`
	EditDate                          *int                             `json:"edit_date,omitempty"`
	EmailProvider                     *string                          `json:"email_provider,omitempty"`
	EmailType                         *string                          `json:"email_type,omitempty"`
	ExtendedGuarantee                 *int                             `json:"extended_guarantee,omitempty"`
	ExternalAuth                      []interface{}                    `json:"externalAuth,omitempty"`
	ExtraPrices                       []StatusItemItemExtraPricesItem  `json:"extraPrices,omitempty"`
	FeedbackData                      interface{}                      `json:"feedback_data,omitempty"`
	GetEmailCodeDisplayLogin          interface{}                      `json:"getEmailCodeDisplayLogin,omitempty"`
	Guarantee                         interface{}                      `json:"guarantee,omitempty"`
	ImagePreviewLinks                 []string                         `json:"imagePreviewLinks,omitempty"`
	InCart                            interface{}                      `json:"in_cart,omitempty"`
	Information                       *string                          `json:"information,omitempty"`
	InformationEn                     *string                          `json:"information_en,omitempty"`
	IsBirthdayToday                   *bool                            `json:"isBirthdayToday,omitempty"`
	IsIgnored                         *bool                            `json:"isIgnored,omitempty"`
	IsPersonalAccount                 *bool                            `json:"isPersonalAccount,omitempty"`
	IsSmallExf                        *bool                            `json:"isSmallExf,omitempty"`
	IsTrusted                         *bool                            `json:"isTrusted,omitempty"`
	IsFave                            interface{}                      `json:"is_fave,omitempty"`
	IsSticky                          *int                             `json:"is_sticky,omitempty"`
	ItemOriginPhrase                  *string                          `json:"itemOriginPhrase,omitempty"`
	ItemDomain                        *string                          `json:"item_domain,omitempty"`
	ItemID                            *int                             `json:"item_id,omitempty"`
	ItemOrigin                        *string                          `json:"item_origin,omitempty"`
	ItemState                         *string                          `json:"item_state,omitempty"`
	Login                             *string                          `json:"login,omitempty"`
	LoginData                         *StatusItemItemLoginData         `json:"loginData,omitempty"`
	MarketCustomTitle                 *string                          `json:"market_custom_title,omitempty"`
	MaxDiscountPercent                *int                             `json:"max_discount_percent,omitempty"`
	NeedToRequireVideoToViewLoginData *bool                            `json:"needToRequireVideoToViewLoginData,omitempty"`
	NoteText                          *string                          `json:"note_text,omitempty"`
	Nsb                               *int                             `json:"nsb,omitempty"`
	PendingDeletionDate               *int                             `json:"pending_deletion_date,omitempty"`
	Price                             *int                             `json:"price,omitempty"`
	PriceWithSellerFee                *float64                         `json:"priceWithSellerFee,omitempty"`
	PriceWithSellerFeeLabel           *string                          `json:"priceWithSellerFeeLabel,omitempty"`
	PriceCurrency                     *string                          `json:"price_currency,omitempty"`
	PublishedDate                     *int                             `json:"published_date,omitempty"`
	RefreshedDate                     *int                             `json:"refreshed_date,omitempty"`
	ResaleItemOrigin                  *string                          `json:"resale_item_origin,omitempty"`
	RubPrice                          *int                             `json:"rub_price,omitempty"`
	Seller                            *StatusItemItemSeller            `json:"seller,omitempty"`
	ShowGetEmailCodeButton            *bool                            `json:"showGetEmailCodeButton,omitempty"`
	Tags                              interface{}                      `json:"tags,omitempty"`
	TempEmail                         *string                          `json:"temp_email,omitempty"`
	Title                             *string                          `json:"title,omitempty"`
	TitleEn                           *string                          `json:"title_en,omitempty"`
	UniqueKeyExists                   *bool                            `json:"uniqueKeyExists,omitempty"`
	UpdateStatDate                    *int                             `json:"update_stat_date,omitempty"`
	UserAllowAskDiscount              *int                             `json:"user_allow_ask_discount,omitempty"`
	ViewCount                         *int                             `json:"view_count,omitempty"`
	VisitorIsAuthor                   *bool                            `json:"visitorIsAuthor,omitempty"`
}

type StatusItemItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type StatusItemItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	NextAllowedBumpDate interface{} `json:"nextAllowedBumpDate,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type StatusItemItemBuyer struct {
	DisplayIconGroupID  *int    `json:"display_icon_group_id,omitempty"`
	DisplayStyleGroupID *int    `json:"display_style_group_id,omitempty"`
	IsBanned            *int    `json:"is_banned,omitempty"`
	OperationDate       *int    `json:"operation_date,omitempty"`
	UniqBanner          *string `json:"uniq_banner,omitempty"`
	UniqUsernameCSS     *string `json:"uniq_username_css,omitempty"`
	UserGroupID         *int    `json:"user_group_id,omitempty"`
	UserID              *int    `json:"user_id,omitempty"`
	Username            *string `json:"username,omitempty"`
	VisitorIsBuyer      *bool   `json:"visitorIsBuyer,omitempty"`
}

type StatusItemItemCopyFormatData struct {
	Full      *string `json:"full,omitempty"`
	LoginData *string `json:"login_data,omitempty"`
	TitleLink *string `json:"title_link,omitempty"`
}

type StatusItemItemCustomFields struct {
	Field4         *string       `json:"_4,omitempty"`
	AllowSelfUnban []interface{} `json:"allowSelfUnban,omitempty"`
	BanReason      *string       `json:"ban_reason,omitempty"`
	Discord        *string       `json:"discord,omitempty"`
	Github         *string       `json:"github,omitempty"`
	Jabber         *string       `json:"jabber,omitempty"`
	LztUnbanAmount *string       `json:"lztUnbanAmount,omitempty"`
	Steam          *string       `json:"steam,omitempty"`
	Telegram       *string       `json:"telegram,omitempty"`
	Vk             *string       `json:"vk,omitempty"`
}

type StatusItemItemExtraPricesItem struct {
	Currency   *string  `json:"currency,omitempty"`
	Price      *string  `json:"price,omitempty"`
	PriceValue *float64 `json:"priceValue,omitempty"`
}

type StatusItemItemGuarantee struct {
	Active                *bool   `json:"active,omitempty"`
	Cancelled             *bool   `json:"cancelled,omitempty"`
	CancelledReason       *string `json:"cancelledReason,omitempty"`
	CancelledReasonPhrase *string `json:"cancelledReasonPhrase,omitempty"`
	Class                 *string `json:"class,omitempty"`
	Duration              *int    `json:"duration,omitempty"`
	DurationPhrase        *string `json:"durationPhrase,omitempty"`
	EndDate               *int    `json:"endDate,omitempty"`
	RemainingTime         *int    `json:"remainingTime,omitempty"`
	RemainingTimePhrase   *string `json:"remainingTimePhrase,omitempty"`
}

type StatusItemItemLoginData struct {
	EncodedOldPassword interface{} `json:"encodedOldPassword,omitempty"`
	EncodedPassword    *string     `json:"encodedPassword,omitempty"`
	EncodedRaw         *string     `json:"encodedRaw,omitempty"`
	Login              *string     `json:"login,omitempty"`
	OldPassword        *string     `json:"oldPassword,omitempty"`
	Password           *string     `json:"password,omitempty"`
	Raw                *string     `json:"raw,omitempty"`
}

type StatusItemItemSeller struct {
	ActiveItemsCount      *int                          `json:"active_items_count,omitempty"`
	AvatarDate            *int                          `json:"avatar_date,omitempty"`
	Contacts              *StatusItemItemSellerContacts `json:"contacts,omitempty"`
	DisplayStyleGroupID   *int                          `json:"display_style_group_id,omitempty"`
	EffectiveLastActivity *int                          `json:"effective_last_activity,omitempty"`
	IsOnline              *bool                         `json:"isOnline,omitempty"`
	IsBanned              *int                          `json:"is_banned,omitempty"`
	JoinedDate            *int                          `json:"joined_date,omitempty"`
	RestoreData           interface{}                   `json:"restore_data,omitempty"`
	RestorePercents       interface{}                   `json:"restore_percents,omitempty"`
	SoldItemsCount        *int                          `json:"sold_items_count,omitempty"`
	UserID                *int                          `json:"user_id,omitempty"`
	Username              *string                       `json:"username,omitempty"`
}

type StatusItemItemSellerContacts struct {
	BanReason *string `json:"ban_reason,omitempty"`
	Telegram  *string `json:"telegram,omitempty"`
}

type StatusItemSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type StatusSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type SteamInventoryValueParams struct {
	// Application id.
	AppID *AccountsManagingAppID `json:"app_id,omitempty"`
	// Currency in which the inventory value will be returned
	Currency *AccountsManagingCurrency `json:"currency,omitempty"`
	// Ignore cache.
	IgnoreCache *bool `json:"ignore_cache,omitempty"`
}

type SteamInventoryValueResponse struct {
	AppId      *int                             `json:"appId,omitempty"`
	Data       *SteamInventoryValueResponseData `json:"data,omitempty"`
	Query      *string                          `json:"query,omitempty"`
	SystemInfo *RespSystemInfo                  `json:"system_info,omitempty"`
}

type SteamInventoryValueResponseData struct {
	AppId               *int                   `json:"appId,omitempty"`
	AppTitle            *string                `json:"appTitle,omitempty"`
	Currency            *string                `json:"currency,omitempty"`
	CurrencyIcon        *string                `json:"currencyIcon,omitempty"`
	ItemCount           *int                   `json:"itemCount,omitempty"`
	Items               map[string]interface{} `json:"items,omitempty"`
	Language            *string                `json:"language,omitempty"`
	MarketableItemCount *int                   `json:"marketableItemCount,omitempty"`
	SteamID             *string                `json:"steam_id,omitempty"`
	Time                *int                   `json:"time,omitempty"`
	TotalValue          *float64               `json:"totalValue,omitempty"`
}

type SteamInventoryValueResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type SteamParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// List of games.
	Game []int `json:"game[],omitempty"`
	// List of minimum hours played by game.
	HoursPlayed map[string]int `json:"hours_played,omitempty"`
	// List of maximum hours played by game.
	HoursPlayedMax map[string]int `json:"hours_played_max,omitempty"`
	// Guarantee type.
	Eg *CategorySearchEg `json:"eg,omitempty"`
	// List of VAC bans by game.
	Vac []int `json:"vac[],omitempty"`
	// Don't check game existence while checking for vac.
	VacSkipGameCheck *bool `json:"vac_skip_game_check,omitempty"`
	// Has community ban.
	Rt *CategorySearchRt `json:"rt,omitempty"`
	// Has lifetime trade ban.
	TradeBan *CategorySearchTradeBan `json:"trade_ban,omitempty"`
	// Has temporary trade limit.
	TradeLimit *CategorySearchTradeLimit `json:"trade_limit,omitempty"`
	// Number of days the account has been offline.
	Daybreak *int `json:"daybreak,omitempty"`
	// Has 5 $ limit.
	Limit *CategorySearchLimit `json:"limit,omitempty"`
	// Has .mafile (Steam Guard Authenticator).
	Mafile *CategorySearchMafile `json:"mafile,omitempty"`
	// How old is the account.
	Reg *int `json:"reg,omitempty"`
	// In what notation is time measured.
	RegPeriod *CategorySearchRegPeriod `json:"reg_period,omitempty"`
	// Minimum level.
	Lmin *int `json:"lmin,omitempty"`
	// Maximum level.
	Lmax *int `json:"lmax,omitempty"`
	// Minimum rank in CS2 Matchmaking.
	Rmin *int `json:"rmin,omitempty"`
	// Maximum rank in CS2 Matchmaking.
	Rmax *int `json:"rmax,omitempty"`
	// Minimum rank in CS2 Wingman.
	WingmanRmin *int `json:"wingman_rmin,omitempty"`
	// Maximum rank in CS2 Wingman.
	WingmanRmax *int `json:"wingman_rmax,omitempty"`
	// Has no VAC ban.
	NoVac *bool `json:"no_vac,omitempty"`
	// Has CS2 Matchmaking ban.
	MmBan *CategorySearchMmBan `json:"mm_ban,omitempty"`
	// Minimum balance.
	BalanceMin *int `json:"balance_min,omitempty"`
	// Maximum balance.
	BalanceMax *int `json:"balance_max,omitempty"`
	// Game ID to check inventory price.
	InvGame *int `json:"inv_game,omitempty"`
	// Minimum inventory price for game.
	InvMin *float64 `json:"inv_min,omitempty"`
	// Maximum inventory price for game.
	InvMax *float64 `json:"inv_max,omitempty"`
	// Minimum number of friends.
	FriendsMin *int `json:"friends_min,omitempty"`
	// Maximum number of friends.
	FriendsMax *int `json:"friends_max,omitempty"`
	// Minimum number of games.
	Gmin *int `json:"gmin,omitempty"`
	// Maximum number of games.
	Gmax *int `json:"gmax,omitempty"`
	// Minimum number of wins.
	WinCountMin *int `json:"win_count_min,omitempty"`
	// Maximum number of wins.
	WinCountMax *int `json:"win_count_max,omitempty"`
	// List of medal IDs.
	MedalID []int `json:"medal_id[],omitempty"`
	// Search for medals using "OR" instead of "AND".
	MedalOperatorOr *bool `json:"medal_operator_or,omitempty"`
	// Minimum number of medals.
	MedalMin *int `json:"medal_min,omitempty"`
	// Maximum number of medals.
	MedalMax *int `json:"medal_max,omitempty"`
	// List of gifts.
	Gift []string `json:"gift[],omitempty"`
	// Minimum number of gifts.
	GiftMin *int `json:"gift_min,omitempty"`
	// Maximum number of gifts.
	GiftMax *int `json:"gift_max,omitempty"`
	// Minimum number of recently played hours.
	RecentlyHoursMin *int `json:"recently_hours_min,omitempty"`
	// Maximum number of recently played hours.
	RecentlyHoursMax *int `json:"recently_hours_max,omitempty"`
	// List of allowed countries.
	Country []string `json:"country[],omitempty"`
	// List of disallowed countries.
	NotCountry []string `json:"not_country[],omitempty"`
	// Minimum CS2 rank.
	Cs2ProfileRankMin *int `json:"cs2_profile_rank_min,omitempty"`
	// Maximum CS2 rank.
	Cs2ProfileRankMax *int `json:"cs2_profile_rank_max,omitempty"`
	// Minimum number of Dota 2 MMR.
	SolommrMin *int `json:"solommr_min,omitempty"`
	// Maximum number of Dota 2 MMR.
	SolommrMax *int `json:"solommr_max,omitempty"`
	// Minimum number of Dota 2 games.
	D2GameCountMin *int `json:"d2_game_count_min,omitempty"`
	// Maximum number of Dota 2 games.
	D2GameCountMax *int `json:"d2_game_count_max,omitempty"`
	// Minimum number of Dota 2 wins.
	D2WinCountMin *int `json:"d2_win_count_min,omitempty"`
	// Maximum number of Dota 2 wins.
	D2WinCountMax *int `json:"d2_win_count_max,omitempty"`
	// Minimum number of Dota 2 behavior.
	D2BehaviorMin *int `json:"d2_behavior_min,omitempty"`
	// Maximum number of Dota 2 behavior.
	D2BehaviorMax *int `json:"d2_behavior_max,omitempty"`
	// Minimum FACEIT level.
	FaceitLvlMin *int `json:"faceit_lvl_min,omitempty"`
	// Maximum FACEIT level.
	FaceitLvlMax *int `json:"faceit_lvl_max,omitempty"`
	// Minimum number of Steam points.
	PointsMin *int `json:"points_min,omitempty"`
	// Maximum number of Steam points.
	PointsMax *int `json:"points_max,omitempty"`
	// Minimum number of relevant games.
	RelevantGmin *int `json:"relevant_gmin,omitempty"`
	// Maximum number of relevant games.
	RelevantGmax *int `json:"relevant_gmax,omitempty"`
	// How old is last transaction.
	LastTransDate *int `json:"last_trans_date,omitempty"`
	// In what notation is time measured.
	LastTransDatePeriod *CategorySearchLastTransDatePeriod `json:"last_trans_date_period,omitempty"`
	// How new is last transaction.
	LastTransDateLater *int `json:"last_trans_date_later,omitempty"`
	// In what notation is time measured.
	LastTransDatePeriodLater *CategorySearchLastTransDatePeriodLater `json:"last_trans_date_period_later,omitempty"`
	// Has no transactions.
	NoTrans *bool `json:"no_trans,omitempty"`
	// Has transactions.
	Trans *bool `json:"trans,omitempty"`
	// Minimum amount spent on gifts.
	GiftsPurchaseMin *float64 `json:"gifts_purchase_min,omitempty"`
	// Maximum amount spent on gifts.
	GiftsPurchaseMax *float64 `json:"gifts_purchase_max,omitempty"`
	// Minimum amount of refunds.
	RefundsPurchaseMin *float64 `json:"refunds_purchase_min,omitempty"`
	// Minimum amount of refunds.
	RefundsPurchaseMax *float64 `json:"refunds_purchase_max,omitempty"`
	// Minimum spend amount on in-game purchases.
	IngamePurchaseMin *float64 `json:"ingame_purchase_min,omitempty"`
	// Maximum spend amount on in-game purchases.
	IngamePurchaseMax *float64 `json:"ingame_purchase_max,omitempty"`
	// Minimum spend amount on all purchases.
	GamesPurchaseMin *float64 `json:"games_purchase_min,omitempty"`
	// Maximum spend amount on all purchases.
	GamesPurchaseMax *float64 `json:"games_purchase_max,omitempty"`
	// Minimum spend amount on all purchases.
	PurchaseMin *float64 `json:"purchase_min,omitempty"`
	// Maximum spend amount on all purchases.
	PurchaseMax *float64 `json:"purchase_max,omitempty"`
	// Has activated keys.
	HasActivatedKeys *CategorySearchHasActivatedKeys `json:"has_activated_keys,omitempty"`
	// Minimum Premier ELO in CS2.
	EloMin *int `json:"elo_min,omitempty"`
	// Maximum Premier ELO in CS2.
	EloMax *int `json:"elo_max,omitempty"`
	// Map for rank in CS2.
	Cs2MapRank *CategorySearchCs2MapRank `json:"cs2_map_rank,omitempty"`
	// Minimum rank in CS2 on a certain map.
	Cs2MapRmin *int `json:"cs2_map_rmin,omitempty"`
	// Maximum rank in CS2 on a certain map.
	Cs2MapRmax *int `json:"cs2_map_rmax,omitempty"`
	// Has FACEIT account.
	HasFaceit *CategorySearchHasFaceit `json:"has_faceit,omitempty"`
	// Minimum FACEIT level.
	FaceitCsgoLvlMin *int `json:"faceit_csgo_lvl_min,omitempty"`
	// Maximum FACEIT level.
	FaceitCsgoLvlMax *int `json:"faceit_csgo_lvl_max,omitempty"`
	// Minimum number of Rust deaths
	RustDeathsMin *int `json:"rust_deaths_min,omitempty"`
	// Maximum number of Rust deaths
	RustDeathsMax *int `json:"rust_deaths_max,omitempty"`
	// Minimum number of Rust kills
	RustKillsMin *int `json:"rust_kills_min,omitempty"`
	// Maximum number of Rust kills
	RustKillsMax *int `json:"rust_kills_max,omitempty"`
	// How old is last match of Dota 2.
	D2LastMatchDate *int `json:"d2_last_match_date,omitempty"`
	// In what notation is time measured.
	D2LastMatchDatePeriod *CategorySearchD2LastMatchDatePeriod `json:"d2_last_match_date_period,omitempty"`
	// Minimum number of available to collect trading cards.
	CardsMin *int `json:"cards_min,omitempty"`
	// Maximum number of available to collect trading cards.
	CardsMax *int `json:"cards_max,omitempty"`
	// Minimum number of available games with available to collect trading cards.
	CardsGamesMin *int `json:"cards_games_min,omitempty"`
	// Maximum number of available games with available to collect trading cards.
	CardsGamesMax *int `json:"cards_games_max,omitempty"`
	// Ignore inventory value if game has VAC ban.
	SkipVacInv *bool `json:"skip_vac_inv,omitempty"`
}

type SteamPreviewParams struct {
	// Type of page.
	Type_ *AccountsManagingType `json:"type,omitempty"`
}

type SteamResponse struct {
	CacheTTL        *int                     `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                    `json:"hasNextPage,omitempty"`
	Items           []SteamResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                     `json:"lastModified,omitempty"`
	Page            *int                     `json:"page,omitempty"`
	PerPage         *int                     `json:"perPage,omitempty"`
	SearchUrl       *string                  `json:"searchUrl,omitempty"`
	ServerTime      *int                     `json:"serverTime,omitempty"`
	StickyItems     []interface{}            `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo          `json:"system_info,omitempty"`
	TotalItems      *int                     `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}              `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                    `json:"wasCached,omitempty"`
}

type SteamResponseItemsItem struct {
	AccountLink                *string                                       `json:"accountLink,omitempty"`
	AccountLinks               []SteamResponseItemsItemAccountLinksItem      `json:"accountLinks,omitempty"`
	AccountLastActivity        *int                                          `json:"account_last_activity,omitempty"`
	AllowAskDiscount           *int                                          `json:"allow_ask_discount,omitempty"`
	BumpSettings               *SteamResponseItemsItemBumpSettings           `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                                         `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                                         `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                                         `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                                         `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                                         `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                                         `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                                         `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                                         `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                                         `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                                         `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                                         `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                                         `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                                         `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                                         `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                                         `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                                         `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                          `json:"category_id,omitempty"`
	ChineseAccount             *bool                                         `json:"chineseAccount,omitempty"`
	Cs2MapsRanks               []interface{}                                 `json:"cs2MapsRanks,omitempty"`
	Cs2PremierElo              interface{}                                   `json:"cs2PremierElo,omitempty"`
	Cs2RankExpired             *bool                                         `json:"cs2RankExpired,omitempty"`
	Description                *string                                       `json:"description,omitempty"`
	DescriptionEnHtml          *string                                       `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                                       `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                                       `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                                       `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                                       `json:"description_en,omitempty"`
	DisplayConvertedBalance    *bool                                         `json:"displayConvertedBalance,omitempty"`
	Dota2CalibrationWarning    *bool                                         `json:"dota2CalibrationWarning,omitempty"`
	EditDate                   *int                                          `json:"edit_date,omitempty"`
	EmailLoginUrl              *string                                       `json:"emailLoginUrl,omitempty"`
	EmailProvider              *string                                       `json:"email_provider,omitempty"`
	EmailType                  *string                                       `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                                          `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                                   `json:"feedback_data,omitempty"`
	Guarantee                  interface{}                                   `json:"guarantee,omitempty"`
	HasCs2                     *bool                                         `json:"hasCs2,omitempty"`
	HasDota2                   *bool                                         `json:"hasDota2,omitempty"`
	HasPendingAutoBuy          *bool                                         `json:"hasPendingAutoBuy,omitempty"`
	HasPossibleBanInDota2      *bool                                         `json:"hasPossibleBanInDota2,omitempty"`
	HasPubg                    *bool                                         `json:"hasPubg,omitempty"`
	HasRust                    *bool                                         `json:"hasRust,omitempty"`
	HasTf2                     *bool                                         `json:"hasTf2,omitempty"`
	InventoryValue             []interface{}                                 `json:"inventoryValue,omitempty"`
	IsIgnored                  *bool                                         `json:"isIgnored,omitempty"`
	IsSmallExf                 *bool                                         `json:"isSmallExf,omitempty"`
	IsSticky                   *int                                          `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                                       `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                                       `json:"item_domain,omitempty"`
	ItemID                     *int                                          `json:"item_id,omitempty"`
	ItemOrigin                 *string                                       `json:"item_origin,omitempty"`
	ItemState                  *string                                       `json:"item_state,omitempty"`
	NoteText                   interface{}                                   `json:"note_text,omitempty"`
	Nsb                        *int                                          `json:"nsb,omitempty"`
	Price                      *int                                          `json:"price,omitempty"`
	PriceWithSellerFee         *float64                                      `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                                       `json:"price_currency,omitempty"`
	PublishedDate              *int                                          `json:"published_date,omitempty"`
	RefreshedDate              *int                                          `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                                       `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount  *int                                          `json:"restore_items_category_count,omitempty"`
	RubPrice                   *int                                          `json:"rub_price,omitempty"`
	Seller                     *SteamResponseItemsItemSeller                 `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                                         `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount     *int                                          `json:"sold_items_category_count,omitempty"`
	SteamCs2Medals             []interface{}                                 `json:"steamCs2Medals,omitempty"`
	SteamData                  *SteamResponseItemsItemSteamData              `json:"steamData,omitempty"`
	SteamDota2WinRate          *int                                          `json:"steamDota2WinRate,omitempty"`
	SteamLifetimeTradeBan      *bool                                         `json:"steamLifetimeTradeBan,omitempty"`
	SteamRelevantGameCount     *int                                          `json:"steamRelevantGameCount,omitempty"`
	SteamTransactions          []SteamResponseItemsItemSteamTransactionsItem `json:"steamTransactions,omitempty"`
	SteamBalance               *string                                       `json:"steam_balance,omitempty"`
	SteamBans                  interface{}                                   `json:"steam_bans,omitempty"`
	SteamCardsCount            *int                                          `json:"steam_cards_count,omitempty"`
	SteamCardsGames            *int                                          `json:"steam_cards_games,omitempty"`
	SteamCommunityBan          *int                                          `json:"steam_community_ban,omitempty"`
	SteamConvertedBalance      *int                                          `json:"steam_converted_balance,omitempty"`
	SteamCountry               *string                                       `json:"steam_country,omitempty"`
	SteamCs2BanDate            *int                                          `json:"steam_cs2_ban_date,omitempty"`
	SteamCs2BanDateActive      *bool                                         `json:"steam_cs2_ban_date_active,omitempty"`
	SteamCs2BanType            *int                                          `json:"steam_cs2_ban_type,omitempty"`
	SteamCs2InvValue           *int                                          `json:"steam_cs2_inv_value,omitempty"`
	SteamCs2LastActivity       *int                                          `json:"steam_cs2_last_activity,omitempty"`
	SteamCs2LastLaunched       *int                                          `json:"steam_cs2_last_launched,omitempty"`
	SteamCs2PremierElo         *int                                          `json:"steam_cs2_premier_elo,omitempty"`
	SteamCs2ProfileRank        *int                                          `json:"steam_cs2_profile_rank,omitempty"`
	SteamCs2RankID             *int                                          `json:"steam_cs2_rank_id,omitempty"`
	SteamCs2WinCount           *int                                          `json:"steam_cs2_win_count,omitempty"`
	SteamCs2WingmanRankID      *int                                          `json:"steam_cs2_wingman_rank_id,omitempty"`
	SteamDota2Behavior         *int                                          `json:"steam_dota2_behavior,omitempty"`
	SteamDota2GameCount        *int                                          `json:"steam_dota2_game_count,omitempty"`
	SteamDota2InvValue         *int                                          `json:"steam_dota2_inv_value,omitempty"`
	SteamDota2LastMatchDate    *int                                          `json:"steam_dota2_last_match_date,omitempty"`
	SteamDota2LoseCount        *int                                          `json:"steam_dota2_lose_count,omitempty"`
	SteamDota2SoloMmr          *int                                          `json:"steam_dota2_solo_mmr,omitempty"`
	SteamDota2WinCount         *int                                          `json:"steam_dota2_win_count,omitempty"`
	SteamDstInvValue           *int                                          `json:"steam_dst_inv_value,omitempty"`
	SteamFaceitLevel           *int                                          `json:"steam_faceit_level,omitempty"`
	SteamFriendCount           *int                                          `json:"steam_friend_count,omitempty"`
	SteamFullGames             *SteamResponseItemsItemSteamFullGames         `json:"steam_full_games,omitempty"`
	SteamGameCount             *int                                          `json:"steam_game_count,omitempty"`
	SteamGiftCount             *int                                          `json:"steam_gift_count,omitempty"`
	SteamHasActivatedKeys      *int                                          `json:"steam_has_activated_keys,omitempty"`
	SteamHoursPlayedRecently   *string                                       `json:"steam_hours_played_recently,omitempty"`
	SteamInvValue              *int                                          `json:"steam_inv_value,omitempty"`
	SteamIsLimited             *int                                          `json:"steam_is_limited,omitempty"`
	SteamItemID                *int                                          `json:"steam_item_id,omitempty"`
	SteamKf2InvValue           *int                                          `json:"steam_kf2_inv_value,omitempty"`
	SteamLastActivity          *int                                          `json:"steam_last_activity,omitempty"`
	SteamLastTransactionDate   *int                                          `json:"steam_last_transaction_date,omitempty"`
	SteamLevel                 *int                                          `json:"steam_level,omitempty"`
	SteamLimitSpent            *string                                       `json:"steam_limit_spent,omitempty"`
	SteamMarket                *int                                          `json:"steam_market,omitempty"`
	SteamMarketBanEndDate      *int                                          `json:"steam_market_ban_end_date,omitempty"`
	SteamMarketRestrictions    *int                                          `json:"steam_market_restrictions,omitempty"`
	SteamMFA                   *int                                          `json:"steam_mfa,omitempty"`
	SteamPoints                *int                                          `json:"steam_points,omitempty"`
	SteamPubgInvValue          *int                                          `json:"steam_pubg_inv_value,omitempty"`
	SteamRegisterDate          *int                                          `json:"steam_register_date,omitempty"`
	SteamRustDeaths            *int                                          `json:"steam_rust_deaths,omitempty"`
	SteamRustInvValue          *int                                          `json:"steam_rust_inv_value,omitempty"`
	SteamRustKillPlayer        *int                                          `json:"steam_rust_kill_player,omitempty"`
	SteamSteamInvValue         *int                                          `json:"steam_steam_inv_value,omitempty"`
	SteamTf2InvValue           *int                                          `json:"steam_tf2_inv_value,omitempty"`
	SteamTotalGamesRub         *int                                          `json:"steam_total_games_rub,omitempty"`
	SteamTotalGiftsRub         *int                                          `json:"steam_total_gifts_rub,omitempty"`
	SteamTotalIngameRub        *int                                          `json:"steam_total_ingame_rub,omitempty"`
	SteamTotalPurchasedRub     *int                                          `json:"steam_total_purchased_rub,omitempty"`
	SteamTotalRefundsRub       *int                                          `json:"steam_total_refunds_rub,omitempty"`
	SteamUnturnedInvValue      *int                                          `json:"steam_unturned_inv_value,omitempty"`
	Tags                       interface{}                                   `json:"tags,omitempty"`
	Title                      *string                                       `json:"title,omitempty"`
	TitleEn                    *string                                       `json:"title_en,omitempty"`
	UpdateStatDate             *int                                          `json:"update_stat_date,omitempty"`
	ViewCount                  *int                                          `json:"view_count,omitempty"`
}

type SteamResponseItemsItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type SteamResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type SteamResponseItemsItemGuarantee struct {
	Active         interface{} `json:"active,omitempty"`
	Cancelled      interface{} `json:"cancelled,omitempty"`
	Class          *string     `json:"class,omitempty"`
	Duration       *int        `json:"duration,omitempty"`
	DurationPhrase *string     `json:"durationPhrase,omitempty"`
	EndDate        interface{} `json:"endDate,omitempty"`
	RemainingTime  interface{} `json:"remainingTime,omitempty"`
}

type SteamResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type SteamResponseItemsItemSteamData struct {
	SteamBanTypeID []interface{} `json:"steam_ban_type_id,omitempty"`
}

type SteamResponseItemsItemSteamFullGames struct {
	List  interface{} `json:"list,omitempty"`
	Total *int        `json:"total,omitempty"`
}

type SteamResponseItemsItemSteamTransactionsItem struct {
	Amount  *string `json:"amount,omitempty"`
	Date    *string `json:"date,omitempty"`
	Product *string `json:"product,omitempty"`
	Source  *string `json:"source,omitempty"`
	Type_   *string `json:"type,omitempty"`
}

type SteamResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type SteamSDAParams struct {
	// Confirmation id. (Required along with **nonce** if you want to confirm action).
	ID *int `json:"id,omitempty"`
	// Confirmation nonce. (Required along with **id** if you want to confirm action).
	Nonce *int `json:"nonce,omitempty"`
}

type SteamUpdateValueParams struct {
	// Update the entire Steam inventory.
	All *bool `json:"all,omitempty"`
	// Application id.
	AppID *AccountsManagingSteamUpdateValueAppID `json:"app_id,omitempty"`
	// Parse inventory when authorized (Parse trade banned items).
	Authorize *bool `json:"authorize,omitempty"`
}

type SteamUpdateValueResponse struct {
	Item       *ItemModel      `json:"item,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type SteamUpdateValueResponseItem struct {
	AccountLink                       *string                                        `json:"accountLink,omitempty"`
	AccountLinks                      []SteamUpdateValueResponseItemAccountLinksItem `json:"accountLinks,omitempty"`
	AccountLastActivity               *int                                           `json:"account_last_activity,omitempty"`
	AiPrice                           *int                                           `json:"aiPrice,omitempty"`
	AiPriceCheckDate                  *int                                           `json:"aiPriceCheckDate,omitempty"`
	AllowAskDiscount                  *int                                           `json:"allow_ask_discount,omitempty"`
	AutoBuyPrice                      *int                                           `json:"autoBuyPrice,omitempty"`
	AutoBuyPriceCheckDate             *int                                           `json:"autoBuyPriceCheckDate,omitempty"`
	BumpSettings                      *SteamUpdateValueResponseItemBumpSettings      `json:"bumpSettings,omitempty"`
	Buyer                             *SteamUpdateValueResponseItemBuyer             `json:"buyer,omitempty"`
	BuyerAvatarDate                   *int                                           `json:"buyer_avatar_date,omitempty"`
	BuyerDisplayIconGroupID           *int                                           `json:"buyer_display_icon_group_id,omitempty"`
	BuyerUniqBanner                   *string                                        `json:"buyer_uniq_banner,omitempty"`
	BuyerUserGroupID                  *int                                           `json:"buyer_user_group_id,omitempty"`
	CanAskDiscount                    *bool                                          `json:"canAskDiscount,omitempty"`
	CanChangeEmailPassword            *bool                                          `json:"canChangeEmailPassword,omitempty"`
	CanChangePassword                 *bool                                          `json:"canChangePassword,omitempty"`
	CanCheckAiPrice                   *bool                                          `json:"canCheckAiPrice,omitempty"`
	CanCheckAutoBuyPrice              *bool                                          `json:"canCheckAutoBuyPrice,omitempty"`
	CanCheckGuarantee                 *bool                                          `json:"canCheckGuarantee,omitempty"`
	CanReportItem                     *bool                                          `json:"canReportItem,omitempty"`
	CanResellItem                     *bool                                          `json:"canResellItem,omitempty"`
	CanResellItemAfterPurchase        *bool                                          `json:"canResellItemAfterPurchase,omitempty"`
	CanShareItem                      *bool                                          `json:"canShareItem,omitempty"`
	CanUpdateItemStats                *bool                                          `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount                *bool                                          `json:"canValidateAccount,omitempty"`
	CanViewAccountLink                *bool                                          `json:"canViewAccountLink,omitempty"`
	CanViewAccountLoginAndTempEmail   *bool                                          `json:"canViewAccountLoginAndTempEmail,omitempty"`
	CanViewEmailLoginData             *bool                                          `json:"canViewEmailLoginData,omitempty"`
	CanViewItemViews                  *bool                                          `json:"canViewItemViews,omitempty"`
	CanViewLoginData                  *bool                                          `json:"canViewLoginData,omitempty"`
	CartPrice                         interface{}                                    `json:"cart_price,omitempty"`
	CategoryID                        *int                                           `json:"category_id,omitempty"`
	ContentID                         interface{}                                    `json:"content_id,omitempty"`
	ContentType                       interface{}                                    `json:"content_type,omitempty"`
	CopyFormatData                    *SteamUpdateValueResponseItemCopyFormatData    `json:"copyFormatData,omitempty"`
	CustomFields                      *SteamUpdateValueResponseItemCustomFields      `json:"customFields,omitempty"`
	DeleteDate                        *int                                           `json:"delete_date,omitempty"`
	DeleteReason                      *string                                        `json:"delete_reason,omitempty"`
	DeleteUserID                      *int                                           `json:"delete_user_id,omitempty"`
	DeleteUsername                    *string                                        `json:"delete_username,omitempty"`
	Deposit                           *int                                           `json:"deposit,omitempty"`
	Description                       *string                                        `json:"description,omitempty"`
	DescriptionEnHtml                 *string                                        `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain                *string                                        `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml                   *string                                        `json:"descriptionHtml,omitempty"`
	DescriptionPlain                  *string                                        `json:"descriptionPlain,omitempty"`
	DescriptionEn                     *string                                        `json:"description_en,omitempty"`
	EditDate                          *int                                           `json:"edit_date,omitempty"`
	EmailProvider                     *string                                        `json:"email_provider,omitempty"`
	EmailType                         *string                                        `json:"email_type,omitempty"`
	ExtendedGuarantee                 *int                                           `json:"extended_guarantee,omitempty"`
	ExternalAuth                      []interface{}                                  `json:"externalAuth,omitempty"`
	ExtraPrices                       []SteamUpdateValueResponseItemExtraPricesItem  `json:"extraPrices,omitempty"`
	FeedbackData                      interface{}                                    `json:"feedback_data,omitempty"`
	GetEmailCodeDisplayLogin          interface{}                                    `json:"getEmailCodeDisplayLogin,omitempty"`
	Guarantee                         interface{}                                    `json:"guarantee,omitempty"`
	ImagePreviewLinks                 []string                                       `json:"imagePreviewLinks,omitempty"`
	InCart                            interface{}                                    `json:"in_cart,omitempty"`
	Information                       *string                                        `json:"information,omitempty"`
	InformationEn                     *string                                        `json:"information_en,omitempty"`
	IsBirthdayToday                   *bool                                          `json:"isBirthdayToday,omitempty"`
	IsIgnored                         *bool                                          `json:"isIgnored,omitempty"`
	IsPersonalAccount                 *bool                                          `json:"isPersonalAccount,omitempty"`
	IsSmallExf                        *bool                                          `json:"isSmallExf,omitempty"`
	IsTrusted                         *bool                                          `json:"isTrusted,omitempty"`
	IsFave                            interface{}                                    `json:"is_fave,omitempty"`
	IsSticky                          *int                                           `json:"is_sticky,omitempty"`
	ItemOriginPhrase                  *string                                        `json:"itemOriginPhrase,omitempty"`
	ItemDomain                        *string                                        `json:"item_domain,omitempty"`
	ItemID                            *int                                           `json:"item_id,omitempty"`
	ItemOrigin                        *string                                        `json:"item_origin,omitempty"`
	ItemState                         *string                                        `json:"item_state,omitempty"`
	Login                             *string                                        `json:"login,omitempty"`
	LoginData                         *SteamUpdateValueResponseItemLoginData         `json:"loginData,omitempty"`
	MarketCustomTitle                 *string                                        `json:"market_custom_title,omitempty"`
	MaxDiscountPercent                *int                                           `json:"max_discount_percent,omitempty"`
	NeedToRequireVideoToViewLoginData *bool                                          `json:"needToRequireVideoToViewLoginData,omitempty"`
	NoteText                          *string                                        `json:"note_text,omitempty"`
	Nsb                               *int                                           `json:"nsb,omitempty"`
	PendingDeletionDate               *int                                           `json:"pending_deletion_date,omitempty"`
	Price                             *int                                           `json:"price,omitempty"`
	PriceWithSellerFee                *float64                                       `json:"priceWithSellerFee,omitempty"`
	PriceWithSellerFeeLabel           *string                                        `json:"priceWithSellerFeeLabel,omitempty"`
	PriceCurrency                     *string                                        `json:"price_currency,omitempty"`
	PublishedDate                     *int                                           `json:"published_date,omitempty"`
	RefreshedDate                     *int                                           `json:"refreshed_date,omitempty"`
	ResaleItemOrigin                  *string                                        `json:"resale_item_origin,omitempty"`
	RubPrice                          *int                                           `json:"rub_price,omitempty"`
	Seller                            *SteamUpdateValueResponseItemSeller            `json:"seller,omitempty"`
	ShowGetEmailCodeButton            *bool                                          `json:"showGetEmailCodeButton,omitempty"`
	Tags                              interface{}                                    `json:"tags,omitempty"`
	TempEmail                         *string                                        `json:"temp_email,omitempty"`
	Title                             *string                                        `json:"title,omitempty"`
	TitleEn                           *string                                        `json:"title_en,omitempty"`
	UniqueKeyExists                   *bool                                          `json:"uniqueKeyExists,omitempty"`
	UpdateStatDate                    *int                                           `json:"update_stat_date,omitempty"`
	UserAllowAskDiscount              *int                                           `json:"user_allow_ask_discount,omitempty"`
	ViewCount                         *int                                           `json:"view_count,omitempty"`
	VisitorIsAuthor                   *bool                                          `json:"visitorIsAuthor,omitempty"`
}

type SteamUpdateValueResponseItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type SteamUpdateValueResponseItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	NextAllowedBumpDate interface{} `json:"nextAllowedBumpDate,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type SteamUpdateValueResponseItemBuyer struct {
	DisplayIconGroupID  *int    `json:"display_icon_group_id,omitempty"`
	DisplayStyleGroupID *int    `json:"display_style_group_id,omitempty"`
	IsBanned            *int    `json:"is_banned,omitempty"`
	OperationDate       *int    `json:"operation_date,omitempty"`
	UniqBanner          *string `json:"uniq_banner,omitempty"`
	UniqUsernameCSS     *string `json:"uniq_username_css,omitempty"`
	UserGroupID         *int    `json:"user_group_id,omitempty"`
	UserID              *int    `json:"user_id,omitempty"`
	Username            *string `json:"username,omitempty"`
	VisitorIsBuyer      *bool   `json:"visitorIsBuyer,omitempty"`
}

type SteamUpdateValueResponseItemCopyFormatData struct {
	Full      *string `json:"full,omitempty"`
	LoginData *string `json:"login_data,omitempty"`
	TitleLink *string `json:"title_link,omitempty"`
}

type SteamUpdateValueResponseItemCustomFields struct {
	Field4         *string       `json:"_4,omitempty"`
	AllowSelfUnban []interface{} `json:"allowSelfUnban,omitempty"`
	BanReason      *string       `json:"ban_reason,omitempty"`
	Discord        *string       `json:"discord,omitempty"`
	Github         *string       `json:"github,omitempty"`
	Jabber         *string       `json:"jabber,omitempty"`
	LztUnbanAmount *string       `json:"lztUnbanAmount,omitempty"`
	Steam          *string       `json:"steam,omitempty"`
	Telegram       *string       `json:"telegram,omitempty"`
	Vk             *string       `json:"vk,omitempty"`
}

type SteamUpdateValueResponseItemExtraPricesItem struct {
	Currency   *string  `json:"currency,omitempty"`
	Price      *string  `json:"price,omitempty"`
	PriceValue *float64 `json:"priceValue,omitempty"`
}

type SteamUpdateValueResponseItemGuarantee struct {
	Active                *bool   `json:"active,omitempty"`
	Cancelled             *bool   `json:"cancelled,omitempty"`
	CancelledReason       *string `json:"cancelledReason,omitempty"`
	CancelledReasonPhrase *string `json:"cancelledReasonPhrase,omitempty"`
	Class                 *string `json:"class,omitempty"`
	Duration              *int    `json:"duration,omitempty"`
	DurationPhrase        *string `json:"durationPhrase,omitempty"`
	EndDate               *int    `json:"endDate,omitempty"`
	RemainingTime         *int    `json:"remainingTime,omitempty"`
	RemainingTimePhrase   *string `json:"remainingTimePhrase,omitempty"`
}

type SteamUpdateValueResponseItemLoginData struct {
	EncodedOldPassword interface{} `json:"encodedOldPassword,omitempty"`
	EncodedPassword    *string     `json:"encodedPassword,omitempty"`
	EncodedRaw         *string     `json:"encodedRaw,omitempty"`
	Login              *string     `json:"login,omitempty"`
	OldPassword        *string     `json:"oldPassword,omitempty"`
	Password           *string     `json:"password,omitempty"`
	Raw                *string     `json:"raw,omitempty"`
}

type SteamUpdateValueResponseItemSeller struct {
	ActiveItemsCount      *int                                        `json:"active_items_count,omitempty"`
	AvatarDate            *int                                        `json:"avatar_date,omitempty"`
	Contacts              *SteamUpdateValueResponseItemSellerContacts `json:"contacts,omitempty"`
	DisplayStyleGroupID   *int                                        `json:"display_style_group_id,omitempty"`
	EffectiveLastActivity *int                                        `json:"effective_last_activity,omitempty"`
	IsOnline              *bool                                       `json:"isOnline,omitempty"`
	IsBanned              *int                                        `json:"is_banned,omitempty"`
	JoinedDate            *int                                        `json:"joined_date,omitempty"`
	RestoreData           interface{}                                 `json:"restore_data,omitempty"`
	RestorePercents       interface{}                                 `json:"restore_percents,omitempty"`
	SoldItemsCount        *int                                        `json:"sold_items_count,omitempty"`
	UserID                *int                                        `json:"user_id,omitempty"`
	Username              *string                                     `json:"username,omitempty"`
}

type SteamUpdateValueResponseItemSellerContacts struct {
	BanReason *string `json:"ban_reason,omitempty"`
	Telegram  *string `json:"telegram,omitempty"`
}

type SteamUpdateValueResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type SteamValueParams struct {
	// Application id.
	AppID *AccountsManagingAppID `json:"app_id,omitempty"`
	// Currency in which the inventory value will be returned
	Currency *AccountsManagingCurrency `json:"currency,omitempty"`
	// Ignore cache.
	IgnoreCache *bool `json:"ignore_cache,omitempty"`
}

type SteamValueResponse struct {
	AppId      *int                    `json:"appId,omitempty"`
	Data       *SteamValueResponseData `json:"data,omitempty"`
	Query      *string                 `json:"query,omitempty"`
	SystemInfo *RespSystemInfo         `json:"system_info,omitempty"`
}

type SteamValueResponseData struct {
	AppId               *int        `json:"appId,omitempty"`
	AppTitle            *string     `json:"appTitle,omitempty"`
	Currency            *string     `json:"currency,omitempty"`
	CurrencyIcon        *string     `json:"currencyIcon,omitempty"`
	ItemCount           *int        `json:"itemCount,omitempty"`
	Items               interface{} `json:"items,omitempty"`
	Language            *string     `json:"language,omitempty"`
	MarketableItemCount *int        `json:"marketableItemCount,omitempty"`
	SteamID             *string     `json:"steam_id,omitempty"`
	Time                *int        `json:"time,omitempty"`
	TotalValue          *float64    `json:"totalValue,omitempty"`
}

type SteamValueResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type SupercellParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Number of days the account has been offline.
	Daybreak *int `json:"daybreak,omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// Guarantee type.
	Eg *CategorySearchEg `json:"eg,omitempty"`
	// Has linked mobile.
	Tel *CategorySearchTel `json:"tel,omitempty"`
	// Minimum Brawl Stars level.
	BrawlLevelMin *int `json:"brawl_level_min,omitempty"`
	// Maximum Brawl Stars level.
	BrawlLevelMax *int `json:"brawl_level_max,omitempty"`
	// Minimum number of Brawl Stars trophies.
	BrawlCupMin *int `json:"brawl_cup_min,omitempty"`
	// Maximum number of Brawl Stars trophies.
	BrawlCupMax *int `json:"brawl_cup_max,omitempty"`
	// Minimum number of Brawl Stars wins.
	BrawlWinsMin *int `json:"brawl_wins_min,omitempty"`
	// Maximum number of Brawl Stars wins.
	BrawlWinsMax *int `json:"brawl_wins_max,omitempty"`
	// Has Brawl Pass.
	BrawlPass *CategorySearchBrawlPass `json:"brawl_pass,omitempty"`
	// List of brawlers.
	Brawler []string `json:"brawler[],omitempty"`
	// Minimum number of brawlers.
	BrawlersMin *int `json:"brawlers_min,omitempty"`
	// Maximum number of brawlers.
	BrawlersMax *int `json:"brawlers_max,omitempty"`
	// Minimum number of legendary brawlers.
	LegendaryBrawlersMin *int `json:"legendary_brawlers_min,omitempty"`
	// Maximum number of legendary brawlers.
	LegendaryBrawlersMax *int `json:"legendary_brawlers_max,omitempty"`
	// Minimum Clash Royale level.
	RoyaleLevelMin *int `json:"royale_level_min,omitempty"`
	// Maximum Clash Royale level.
	RoyaleLevelMax *int `json:"royale_level_max,omitempty"`
	// Minimum number of Clash Royale trophies.
	RoyaleCupMin *int `json:"royale_cup_min,omitempty"`
	// Maximum number of Clash Royale trophies.
	RoyaleCupMax *int `json:"royale_cup_max,omitempty"`
	// Minimum number of Clash Royale wins.
	RoyaleWinsMin *int `json:"royale_wins_min,omitempty"`
	// Maximum number of Clash Royale wins.
	RoyaleWinsMax *int `json:"royale_wins_max,omitempty"`
	// Minimum King level in Clash Royale.
	KingLevelMin *int `json:"king_level_min,omitempty"`
	// Maximum King level in Clash Royale.
	KingLevelMax *int `json:"king_level_max,omitempty"`
	// Has Royale Pass.
	RoyalePass *CategorySearchRoyalePass `json:"royale_pass,omitempty"`
	// Minimum Clash of Clans level.
	ClashLevelMin *int `json:"clash_level_min,omitempty"`
	// Maximum Clash of Clans level.
	ClashLevelMax *int `json:"clash_level_max,omitempty"`
	// Minimum number of Clash of Clans trophies.
	ClashCupMin *int `json:"clash_cup_min,omitempty"`
	// Maximum number of Clash of Clans trophies.
	ClashCupMax *int `json:"clash_cup_max,omitempty"`
	// Minimum number of Clash of Clans wins.
	ClashWinsMin *int `json:"clash_wins_min,omitempty"`
	// Maximum number of Clash of Clans wins.
	ClashWinsMax *int `json:"clash_wins_max,omitempty"`
	// Has Battle Pass.
	ClashPass *CategorySearchClashPass `json:"clash_pass,omitempty"`
	// Minimum total heroes level count in Clash of Clans.
	TotalHeroesLevelMin *int `json:"total_heroes_level_min,omitempty"`
	// Maximum total heroes level count in Clash of Clans.
	TotalHeroesLevelMax *int `json:"total_heroes_level_max,omitempty"`
	// Minimum total troops level count in Clash of Clans.
	TotalTroopsLevelMin *int `json:"total_troops_level_min,omitempty"`
	// Maximum total troops level count in Clash of Clans.
	TotalTroopsLevelMax *int `json:"total_troops_level_max,omitempty"`
	// Minimum total spells level count in Clash of Clans.
	TotalSpellsLevelMin *int `json:"total_spells_level_min,omitempty"`
	// Maximum total spells level count in Clash of Clans.
	TotalSpellsLevelMax *int `json:"total_spells_level_max,omitempty"`
	// Minimum total builder village heroes level count in Clash of Clans.
	TotalBuilderHeroesLevelMin *int `json:"total_builder_heroes_level_min,omitempty"`
	// Maximum total builder village heroes level count in Clash of Clans.
	TotalBuilderHeroesLevelMax *int `json:"total_builder_heroes_level_max,omitempty"`
	// Minimum total builder village troops level count in Clash of Clans.
	TotalBuilderTroopsLevelMin *int `json:"total_builder_troops_level_min,omitempty"`
	// Maximum total builder village troops level count in Clash of Clans.
	TotalBuilderTroopsLevelMax *int `json:"total_builder_troops_level_max,omitempty"`
	// Minimum level of town hall.
	TownHallLevelMin *int `json:"town_hall_level_min,omitempty"`
	// Maximum level of town hall.
	TownHallLevelMax *int `json:"town_hall_level_max,omitempty"`
	// Minimum level of builder hall.
	BuilderHallLevelMin *int `json:"builder_hall_level_min,omitempty"`
	// Maximum level of builder hall.
	BuilderHallLevelMax *int `json:"builder_hall_level_max,omitempty"`
	// Minimum number of builder hall cups.
	BuilderHallCupMin *int `json:"builder_hall_cup_min,omitempty"`
	// Maximum number of builder hall cups.
	BuilderHallCupMax *int `json:"builder_hall_cup_max,omitempty"`
	// Minimum account creation year (e.g. 2023).
	CreationYearMin *int `json:"creation_year_min,omitempty"`
	// Maximum account creation year (e.g. 2024).
	CreationYearMax *int `json:"creation_year_max,omitempty"`
}

type SupercellResponse struct {
	CacheTTL        *int                         `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                        `json:"hasNextPage,omitempty"`
	Items           []SupercellResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                         `json:"lastModified,omitempty"`
	Page            *int                         `json:"page,omitempty"`
	PerPage         *int                         `json:"perPage,omitempty"`
	SearchUrl       *string                      `json:"searchUrl,omitempty"`
	ServerTime      *int                         `json:"serverTime,omitempty"`
	StickyItems     []interface{}                `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo              `json:"system_info,omitempty"`
	TotalItems      *int                         `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}                  `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                        `json:"wasCached,omitempty"`
}

type SupercellResponseItemsItem struct {
	AccountLink                      *string                                      `json:"accountLink,omitempty"`
	AccountLinks                     []SupercellResponseItemsItemAccountLinksItem `json:"accountLinks,omitempty"`
	AllowAskDiscount                 *int                                         `json:"allow_ask_discount,omitempty"`
	BumpSettings                     *SupercellResponseItemsItemBumpSettings      `json:"bumpSettings,omitempty"`
	CanBumpItem                      *bool                                        `json:"canBumpItem,omitempty"`
	CanBuyItem                       *bool                                        `json:"canBuyItem,omitempty"`
	CanChangePassword                *bool                                        `json:"canChangePassword,omitempty"`
	CanCloseItem                     *bool                                        `json:"canCloseItem,omitempty"`
	CanDeleteItem                    *bool                                        `json:"canDeleteItem,omitempty"`
	CanEditItem                      *bool                                        `json:"canEditItem,omitempty"`
	CanOpenItem                      *bool                                        `json:"canOpenItem,omitempty"`
	CanReportItem                    *bool                                        `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase       *bool                                        `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem                     *bool                                        `json:"canStickItem,omitempty"`
	CanUnstickItem                   *bool                                        `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats               *bool                                        `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount               *bool                                        `json:"canValidateAccount,omitempty"`
	CanViewAccountLink               *bool                                        `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData            *bool                                        `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData                 *bool                                        `json:"canViewLoginData,omitempty"`
	CategoryID                       *int                                         `json:"category_id,omitempty"`
	Description                      *string                                      `json:"description,omitempty"`
	DescriptionEnHtml                *string                                      `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain               *string                                      `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml                  *string                                      `json:"descriptionHtml,omitempty"`
	DescriptionPlain                 *string                                      `json:"descriptionPlain,omitempty"`
	DescriptionEn                    *string                                      `json:"description_en,omitempty"`
	EditDate                         *int                                         `json:"edit_date,omitempty"`
	EmailLoginUrl                    *string                                      `json:"emailLoginUrl,omitempty"`
	EmailProvider                    *string                                      `json:"email_provider,omitempty"`
	EmailType                        *string                                      `json:"email_type,omitempty"`
	ExtendedGuarantee                *int                                         `json:"extended_guarantee,omitempty"`
	FeedbackData                     interface{}                                  `json:"feedback_data,omitempty"`
	Guarantee                        interface{}                                  `json:"guarantee,omitempty"`
	HasPendingAutoBuy                *bool                                        `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                        *bool                                        `json:"isIgnored,omitempty"`
	IsSmallExf                       *bool                                        `json:"isSmallExf,omitempty"`
	IsSticky                         *int                                         `json:"is_sticky,omitempty"`
	ItemOriginPhrase                 *string                                      `json:"itemOriginPhrase,omitempty"`
	ItemDomain                       *string                                      `json:"item_domain,omitempty"`
	ItemID                           *int                                         `json:"item_id,omitempty"`
	ItemOrigin                       *string                                      `json:"item_origin,omitempty"`
	ItemState                        *string                                      `json:"item_state,omitempty"`
	NoteText                         interface{}                                  `json:"note_text,omitempty"`
	Nsb                              *int                                         `json:"nsb,omitempty"`
	Price                            *int                                         `json:"price,omitempty"`
	PriceWithSellerFee               *float64                                     `json:"priceWithSellerFee,omitempty"`
	PriceCurrency                    *string                                      `json:"price_currency,omitempty"`
	PublishedDate                    *int                                         `json:"published_date,omitempty"`
	RefreshedDate                    *int                                         `json:"refreshed_date,omitempty"`
	ResaleItemOrigin                 *string                                      `json:"resale_item_origin,omitempty"`
	RubPrice                         *int                                         `json:"rub_price,omitempty"`
	Seller                           *SupercellResponseItemsItemSeller            `json:"seller,omitempty"`
	ShowGetEmailCodeButton           *bool                                        `json:"showGetEmailCodeButton,omitempty"`
	SupercellBrawlers                interface{}                                  `json:"supercellBrawlers,omitempty"`
	SupercellArena                   *string                                      `json:"supercell_arena,omitempty"`
	SupercellBrawlerCount            *int                                         `json:"supercell_brawler_count,omitempty"`
	SupercellBuilderHallCupCount     *int                                         `json:"supercell_builder_hall_cup_count,omitempty"`
	SupercellBuilderHallLevel        *int                                         `json:"supercell_builder_hall_level,omitempty"`
	SupercellID                      *string                                      `json:"supercell_id,omitempty"`
	SupercellItemID                  *int                                         `json:"supercell_item_id,omitempty"`
	SupercellKingLevel               *int                                         `json:"supercell_king_level,omitempty"`
	SupercellLaserBattlePass         *int                                         `json:"supercell_laser_battle_pass,omitempty"`
	SupercellLaserLevel              *int                                         `json:"supercell_laser_level,omitempty"`
	SupercellLaserTrophies           *int                                         `json:"supercell_laser_trophies,omitempty"`
	SupercellLaserVictories          *int                                         `json:"supercell_laser_victories,omitempty"`
	SupercellLastActivity            *int                                         `json:"supercell_last_activity,omitempty"`
	SupercellLegendaryBrawlerCount   *int                                         `json:"supercell_legendary_brawler_count,omitempty"`
	SupercellMagicBattlePass         *int                                         `json:"supercell_magic_battle_pass,omitempty"`
	SupercellMagicLevel              *int                                         `json:"supercell_magic_level,omitempty"`
	SupercellMagicTrophies           *int                                         `json:"supercell_magic_trophies,omitempty"`
	SupercellMagicVictories          *int                                         `json:"supercell_magic_victories,omitempty"`
	SupercellPhone                   *int                                         `json:"supercell_phone,omitempty"`
	SupercellScrollBattlePass        *int                                         `json:"supercell_scroll_battle_pass,omitempty"`
	SupercellScrollLevel             *int                                         `json:"supercell_scroll_level,omitempty"`
	SupercellScrollTrophies          *int                                         `json:"supercell_scroll_trophies,omitempty"`
	SupercellScrollVictories         *int                                         `json:"supercell_scroll_victories,omitempty"`
	SupercellSystems                 *string                                      `json:"supercell_systems,omitempty"`
	SupercellTotalBuilderHeroesLevel *int                                         `json:"supercell_total_builder_heroes_level,omitempty"`
	SupercellTotalBuilderTroopsLevel *int                                         `json:"supercell_total_builder_troops_level,omitempty"`
	SupercellTotalHeroesLevel        *int                                         `json:"supercell_total_heroes_level,omitempty"`
	SupercellTotalSpellsLevel        *int                                         `json:"supercell_total_spells_level,omitempty"`
	SupercellTotalTroopsLevel        *int                                         `json:"supercell_total_troops_level,omitempty"`
	SupercellTownHallLevel           *int                                         `json:"supercell_town_hall_level,omitempty"`
	Tags                             interface{}                                  `json:"tags,omitempty"`
	Title                            *string                                      `json:"title,omitempty"`
	TitleEn                          *string                                      `json:"title_en,omitempty"`
	UpdateStatDate                   *int                                         `json:"update_stat_date,omitempty"`
	ViewCount                        *int                                         `json:"view_count,omitempty"`
}

type SupercellResponseItemsItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type SupercellResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type SupercellResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     interface{} `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type SupercellResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type Tag struct {
	AddedTagId *int            `json:"addedTagId,omitempty"`
	DeleteTags []int           `json:"deleteTags,omitempty"`
	ItemId     *int            `json:"itemId,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
	Tag        *TagTag         `json:"tag,omitempty"`
}

type TagSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type TagTag struct {
	Bc                   *string `json:"bc,omitempty"`
	ForOwnedAccountsOnly *bool   `json:"forOwnedAccountsOnly,omitempty"`
	IsDefault            *bool   `json:"isDefault,omitempty"`
	TagID                *int    `json:"tag_id,omitempty"`
	Title                *string `json:"title,omitempty"`
}

type TelegramCodeResponse struct {
	Codes *TelegramCodeResponseCodes `json:"codes,omitempty"`
	Item  *ItemModel                 `json:"item,omitempty"`
}

type TelegramCodeResponseCodes struct {
	Code *string `json:"code,omitempty"`
	Date *int    `json:"date,omitempty"`
}

type TelegramCodeResponseItem struct {
	AccountLink                       *string                                    `json:"accountLink,omitempty"`
	AccountLinks                      []TelegramCodeResponseItemAccountLinksItem `json:"accountLinks,omitempty"`
	AccountLastActivity               *int                                       `json:"account_last_activity,omitempty"`
	AiPrice                           *int                                       `json:"aiPrice,omitempty"`
	AiPriceCheckDate                  *int                                       `json:"aiPriceCheckDate,omitempty"`
	AllowAskDiscount                  *int                                       `json:"allow_ask_discount,omitempty"`
	AutoBuyPrice                      *int                                       `json:"autoBuyPrice,omitempty"`
	AutoBuyPriceCheckDate             *int                                       `json:"autoBuyPriceCheckDate,omitempty"`
	BumpSettings                      *TelegramCodeResponseItemBumpSettings      `json:"bumpSettings,omitempty"`
	Buyer                             *TelegramCodeResponseItemBuyer             `json:"buyer,omitempty"`
	BuyerAvatarDate                   *int                                       `json:"buyer_avatar_date,omitempty"`
	BuyerDisplayIconGroupID           *int                                       `json:"buyer_display_icon_group_id,omitempty"`
	BuyerUniqBanner                   *string                                    `json:"buyer_uniq_banner,omitempty"`
	BuyerUserGroupID                  *int                                       `json:"buyer_user_group_id,omitempty"`
	CanAskDiscount                    *bool                                      `json:"canAskDiscount,omitempty"`
	CanChangeEmailPassword            *bool                                      `json:"canChangeEmailPassword,omitempty"`
	CanChangePassword                 *bool                                      `json:"canChangePassword,omitempty"`
	CanCheckAiPrice                   *bool                                      `json:"canCheckAiPrice,omitempty"`
	CanCheckAutoBuyPrice              *bool                                      `json:"canCheckAutoBuyPrice,omitempty"`
	CanCheckGuarantee                 *bool                                      `json:"canCheckGuarantee,omitempty"`
	CanReportItem                     *bool                                      `json:"canReportItem,omitempty"`
	CanResellItem                     *bool                                      `json:"canResellItem,omitempty"`
	CanResellItemAfterPurchase        *bool                                      `json:"canResellItemAfterPurchase,omitempty"`
	CanShareItem                      *bool                                      `json:"canShareItem,omitempty"`
	CanUpdateItemStats                *bool                                      `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount                *bool                                      `json:"canValidateAccount,omitempty"`
	CanViewAccountLink                *bool                                      `json:"canViewAccountLink,omitempty"`
	CanViewAccountLoginAndTempEmail   *bool                                      `json:"canViewAccountLoginAndTempEmail,omitempty"`
	CanViewEmailLoginData             *bool                                      `json:"canViewEmailLoginData,omitempty"`
	CanViewItemViews                  *bool                                      `json:"canViewItemViews,omitempty"`
	CanViewLoginData                  *bool                                      `json:"canViewLoginData,omitempty"`
	CartPrice                         interface{}                                `json:"cart_price,omitempty"`
	CategoryID                        *int                                       `json:"category_id,omitempty"`
	ContentID                         interface{}                                `json:"content_id,omitempty"`
	ContentType                       interface{}                                `json:"content_type,omitempty"`
	CopyFormatData                    *TelegramCodeResponseItemCopyFormatData    `json:"copyFormatData,omitempty"`
	CustomFields                      *TelegramCodeResponseItemCustomFields      `json:"customFields,omitempty"`
	DeleteDate                        *int                                       `json:"delete_date,omitempty"`
	DeleteReason                      *string                                    `json:"delete_reason,omitempty"`
	DeleteUserID                      *int                                       `json:"delete_user_id,omitempty"`
	DeleteUsername                    *string                                    `json:"delete_username,omitempty"`
	Deposit                           *int                                       `json:"deposit,omitempty"`
	Description                       *string                                    `json:"description,omitempty"`
	DescriptionEnHtml                 *string                                    `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain                *string                                    `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml                   *string                                    `json:"descriptionHtml,omitempty"`
	DescriptionPlain                  *string                                    `json:"descriptionPlain,omitempty"`
	DescriptionEn                     *string                                    `json:"description_en,omitempty"`
	EditDate                          *int                                       `json:"edit_date,omitempty"`
	EmailProvider                     *string                                    `json:"email_provider,omitempty"`
	EmailType                         *string                                    `json:"email_type,omitempty"`
	ExtendedGuarantee                 *int                                       `json:"extended_guarantee,omitempty"`
	ExternalAuth                      []interface{}                              `json:"externalAuth,omitempty"`
	ExtraPrices                       []TelegramCodeResponseItemExtraPricesItem  `json:"extraPrices,omitempty"`
	FeedbackData                      interface{}                                `json:"feedback_data,omitempty"`
	GetEmailCodeDisplayLogin          interface{}                                `json:"getEmailCodeDisplayLogin,omitempty"`
	Guarantee                         interface{}                                `json:"guarantee,omitempty"`
	ImagePreviewLinks                 []string                                   `json:"imagePreviewLinks,omitempty"`
	InCart                            interface{}                                `json:"in_cart,omitempty"`
	Information                       *string                                    `json:"information,omitempty"`
	InformationEn                     *string                                    `json:"information_en,omitempty"`
	IsBirthdayToday                   *bool                                      `json:"isBirthdayToday,omitempty"`
	IsIgnored                         *bool                                      `json:"isIgnored,omitempty"`
	IsPersonalAccount                 *bool                                      `json:"isPersonalAccount,omitempty"`
	IsSmallExf                        *bool                                      `json:"isSmallExf,omitempty"`
	IsTrusted                         *bool                                      `json:"isTrusted,omitempty"`
	IsFave                            interface{}                                `json:"is_fave,omitempty"`
	IsSticky                          *int                                       `json:"is_sticky,omitempty"`
	ItemOriginPhrase                  *string                                    `json:"itemOriginPhrase,omitempty"`
	ItemDomain                        *string                                    `json:"item_domain,omitempty"`
	ItemID                            *int                                       `json:"item_id,omitempty"`
	ItemOrigin                        *string                                    `json:"item_origin,omitempty"`
	ItemState                         *string                                    `json:"item_state,omitempty"`
	Login                             *string                                    `json:"login,omitempty"`
	LoginData                         *TelegramCodeResponseItemLoginData         `json:"loginData,omitempty"`
	MarketCustomTitle                 *string                                    `json:"market_custom_title,omitempty"`
	MaxDiscountPercent                *int                                       `json:"max_discount_percent,omitempty"`
	NeedToRequireVideoToViewLoginData *bool                                      `json:"needToRequireVideoToViewLoginData,omitempty"`
	NoteText                          *string                                    `json:"note_text,omitempty"`
	Nsb                               *int                                       `json:"nsb,omitempty"`
	PendingDeletionDate               *int                                       `json:"pending_deletion_date,omitempty"`
	Price                             *int                                       `json:"price,omitempty"`
	PriceWithSellerFee                *float64                                   `json:"priceWithSellerFee,omitempty"`
	PriceWithSellerFeeLabel           *string                                    `json:"priceWithSellerFeeLabel,omitempty"`
	PriceCurrency                     *string                                    `json:"price_currency,omitempty"`
	PublishedDate                     *int                                       `json:"published_date,omitempty"`
	RefreshedDate                     *int                                       `json:"refreshed_date,omitempty"`
	ResaleItemOrigin                  *string                                    `json:"resale_item_origin,omitempty"`
	RubPrice                          *int                                       `json:"rub_price,omitempty"`
	Seller                            *TelegramCodeResponseItemSeller            `json:"seller,omitempty"`
	ShowGetEmailCodeButton            *bool                                      `json:"showGetEmailCodeButton,omitempty"`
	Tags                              interface{}                                `json:"tags,omitempty"`
	TempEmail                         *string                                    `json:"temp_email,omitempty"`
	Title                             *string                                    `json:"title,omitempty"`
	TitleEn                           *string                                    `json:"title_en,omitempty"`
	UniqueKeyExists                   *bool                                      `json:"uniqueKeyExists,omitempty"`
	UpdateStatDate                    *int                                       `json:"update_stat_date,omitempty"`
	UserAllowAskDiscount              *int                                       `json:"user_allow_ask_discount,omitempty"`
	ViewCount                         *int                                       `json:"view_count,omitempty"`
	VisitorIsAuthor                   *bool                                      `json:"visitorIsAuthor,omitempty"`
}

type TelegramCodeResponseItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type TelegramCodeResponseItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	NextAllowedBumpDate interface{} `json:"nextAllowedBumpDate,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type TelegramCodeResponseItemBuyer struct {
	DisplayIconGroupID  *int    `json:"display_icon_group_id,omitempty"`
	DisplayStyleGroupID *int    `json:"display_style_group_id,omitempty"`
	IsBanned            *int    `json:"is_banned,omitempty"`
	OperationDate       *int    `json:"operation_date,omitempty"`
	UniqBanner          *string `json:"uniq_banner,omitempty"`
	UniqUsernameCSS     *string `json:"uniq_username_css,omitempty"`
	UserGroupID         *int    `json:"user_group_id,omitempty"`
	UserID              *int    `json:"user_id,omitempty"`
	Username            *string `json:"username,omitempty"`
	VisitorIsBuyer      *bool   `json:"visitorIsBuyer,omitempty"`
}

type TelegramCodeResponseItemCopyFormatData struct {
	Full      *string `json:"full,omitempty"`
	LoginData *string `json:"login_data,omitempty"`
	TitleLink *string `json:"title_link,omitempty"`
}

type TelegramCodeResponseItemCustomFields struct {
	Field4         *string       `json:"_4,omitempty"`
	AllowSelfUnban []interface{} `json:"allowSelfUnban,omitempty"`
	BanReason      *string       `json:"ban_reason,omitempty"`
	Discord        *string       `json:"discord,omitempty"`
	Github         *string       `json:"github,omitempty"`
	Jabber         *string       `json:"jabber,omitempty"`
	LztUnbanAmount *string       `json:"lztUnbanAmount,omitempty"`
	Steam          *string       `json:"steam,omitempty"`
	Telegram       *string       `json:"telegram,omitempty"`
	Vk             *string       `json:"vk,omitempty"`
}

type TelegramCodeResponseItemExtraPricesItem struct {
	Currency   *string  `json:"currency,omitempty"`
	Price      *string  `json:"price,omitempty"`
	PriceValue *float64 `json:"priceValue,omitempty"`
}

type TelegramCodeResponseItemGuarantee struct {
	Active                *bool   `json:"active,omitempty"`
	Cancelled             *bool   `json:"cancelled,omitempty"`
	CancelledReason       *string `json:"cancelledReason,omitempty"`
	CancelledReasonPhrase *string `json:"cancelledReasonPhrase,omitempty"`
	Class                 *string `json:"class,omitempty"`
	Duration              *int    `json:"duration,omitempty"`
	DurationPhrase        *string `json:"durationPhrase,omitempty"`
	EndDate               *int    `json:"endDate,omitempty"`
	RemainingTime         *int    `json:"remainingTime,omitempty"`
	RemainingTimePhrase   *string `json:"remainingTimePhrase,omitempty"`
}

type TelegramCodeResponseItemLoginData struct {
	EncodedOldPassword interface{} `json:"encodedOldPassword,omitempty"`
	EncodedPassword    *string     `json:"encodedPassword,omitempty"`
	EncodedRaw         *string     `json:"encodedRaw,omitempty"`
	Login              *string     `json:"login,omitempty"`
	OldPassword        *string     `json:"oldPassword,omitempty"`
	Password           *string     `json:"password,omitempty"`
	Raw                *string     `json:"raw,omitempty"`
}

type TelegramCodeResponseItemSeller struct {
	ActiveItemsCount      *int                                    `json:"active_items_count,omitempty"`
	AvatarDate            *int                                    `json:"avatar_date,omitempty"`
	Contacts              *TelegramCodeResponseItemSellerContacts `json:"contacts,omitempty"`
	DisplayStyleGroupID   *int                                    `json:"display_style_group_id,omitempty"`
	EffectiveLastActivity *int                                    `json:"effective_last_activity,omitempty"`
	IsOnline              *bool                                   `json:"isOnline,omitempty"`
	IsBanned              *int                                    `json:"is_banned,omitempty"`
	JoinedDate            *int                                    `json:"joined_date,omitempty"`
	RestoreData           interface{}                             `json:"restore_data,omitempty"`
	RestorePercents       interface{}                             `json:"restore_percents,omitempty"`
	SoldItemsCount        *int                                    `json:"sold_items_count,omitempty"`
	UserID                *int                                    `json:"user_id,omitempty"`
	Username              *string                                 `json:"username,omitempty"`
}

type TelegramCodeResponseItemSellerContacts struct {
	BanReason *string `json:"ban_reason,omitempty"`
	Telegram  *string `json:"telegram,omitempty"`
}

type TelegramParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Has a spam ban.
	Spam *CategorySearchSpam `json:"spam,omitempty"`
	// Has a cloud password.
	Password *CategorySearchPassword `json:"password,omitempty"`
	// Has a premium subscription.
	Premium *CategorySearchPremium `json:"premium,omitempty"`
	// When premium subscription will be active
	PremiumExpiration *int `json:"premium_expiration,omitempty"`
	// In what notation is time measured
	PremiumExpirationPeriod *CategorySearchPremiumExpirationPeriod `json:"premium_expiration_period,omitempty"`
	// List of allowed countries.
	Country []string `json:"country[],omitempty"`
	// List of disallowed countries.
	NotCountry []string `json:"not_country[],omitempty"`
	// Number of days the account has been offline.
	Daybreak *int `json:"daybreak,omitempty"`
	// Minimum number of channels.
	MinChannels *int `json:"min_channels,omitempty"`
	// Maximum number of channels.
	MaxChannels *int `json:"max_channels,omitempty"`
	// Minimum number of chats.
	MinChats *int `json:"min_chats,omitempty"`
	// Maximum number of chats.
	MaxChats *int `json:"max_chats,omitempty"`
	// Minimum number of conversations.
	MinConversations *int `json:"min_conversations,omitempty"`
	// Maximum number of conversations.
	MaxConversations *int `json:"max_conversations,omitempty"`
	// Minimum number of channels, where account is administrator/owner.
	MinAdmin *int `json:"min_admin,omitempty"`
	// Maximum number of channels, where account is administrator/owner.
	MaxAdmin *int `json:"max_admin,omitempty"`
	// Minimum number of subscribers in channel, where account is administrator/owner.
	MinAdminSub *int `json:"min_admin_sub,omitempty"`
	// Maximum number of subscribers in channel, where account is administrator/owner.
	MaxAdminSub *int `json:"max_admin_sub,omitempty"`
	// Minimum number of digits in ID.
	DigMin *int `json:"dig_min,omitempty"`
	// Maximum number of digits in ID.
	DigMax *int `json:"dig_max,omitempty"`
	// Minimum number of contacts.
	MinContacts *int `json:"min_contacts,omitempty"`
	// Maximum number of contacts.
	MaxContacts *int `json:"max_contacts,omitempty"`
	// Minimum number of Telegram Stars.
	MinStars *int `json:"min_stars,omitempty"`
	// Maximum number of Telegram Stars.
	MaxStars *int `json:"max_stars,omitempty"`
	// Birthday was X time before.
	Birthday *int `json:"birthday,omitempty"`
	// In what notation is time measured.
	BirthdayPeriod *CategorySearchBirthdayPeriod `json:"birthday_period,omitempty"`
	// Birthday was X time after.
	BirthdayAfter *int `json:"birthday_after,omitempty"`
	// In what notation is time measured.
	BirthdayAfterPeriod *CategorySearchBirthdayAfterPeriod `json:"birthday_after_period,omitempty"`
	// Minimum ID of account, will be rounded down till nearest 10k. Available if your balance is higher than 100000 RUB.
	MinID *int `json:"min_id,omitempty"`
	// Maximum ID of account, will be rounded down till nearest 10k. Available if your balance is higher than 100000 RUB.
	MaxID *int `json:"max_id,omitempty"`
	// Allow geo spam block in search with spam=no.
	AllowGeoSpamblock *bool `json:"allow_geo_spamblock,omitempty"`
	// Minimum number of Telegram gifts on account.
	MinGifts *int `json:"min_gifts,omitempty"`
	// Maximum number of Telegram gifts on account.
	MaxGifts *int `json:"max_gifts,omitempty"`
	// Minimum number of Telegram NFT gifts on account.
	MinNftGifts *int `json:"min_nft_gifts,omitempty"`
	// Maximum number of Telegram NFT gifts on account.
	MaxNftGifts *int `json:"max_nft_gifts,omitempty"`
	// Minimum value of all Stars gifts.
	MinGiftsStars *int `json:"min_gifts_stars,omitempty"`
	// Maximum value of all Stars gifts.
	MaxGiftsStars *int `json:"max_gifts_stars,omitempty"`
	// Minimum value of all Stars gifts after convert.
	MinGiftsConvertStars *int `json:"min_gifts_convert_stars,omitempty"`
	// Maximum value of all Stars gifts after convert.
	MaxGiftsConvertStars *int `json:"max_gifts_convert_stars,omitempty"`
	// List of allowed DC ID.
	DcID []int `json:"dc_id[],omitempty"`
	// List of disallowed DC ID.
	NotDcID []int `json:"not_dc_id[],omitempty"`
	// Has linked email.
	Email *CategorySearchEmail `json:"email,omitempty"`
	// Minimum number of bots.
	MinBots *int `json:"min_bots,omitempty"`
	// Maximum number of bots.
	MaxBots *int `json:"max_bots,omitempty"`
	// Minimum active users in bot.
	MinBotActiveUsers *int `json:"min_bot_active_users,omitempty"`
	// Maximum active users in bot.
	MaxBotActiveUsers *int `json:"max_bot_active_users,omitempty"`
}

type TelegramResponse struct {
	CacheTTL        *int                        `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                       `json:"hasNextPage,omitempty"`
	Items           []TelegramResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                        `json:"lastModified,omitempty"`
	Page            *int                        `json:"page,omitempty"`
	PerPage         *int                        `json:"perPage,omitempty"`
	SearchUrl       *string                     `json:"searchUrl,omitempty"`
	ServerTime      *int                        `json:"serverTime,omitempty"`
	StickyItems     []interface{}               `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo             `json:"system_info,omitempty"`
	TotalItems      *int                        `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}                 `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                       `json:"wasCached,omitempty"`
}

type TelegramResponseItemsItem struct {
	AccountLinks               []interface{}                                   `json:"accountLinks,omitempty"`
	AllowAskDiscount           *int                                            `json:"allow_ask_discount,omitempty"`
	BumpSettings               *TelegramResponseItemsItemBumpSettings          `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                                           `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                                           `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                                           `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                                           `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                                           `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                                           `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                                           `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                                           `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                                           `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                                           `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                                           `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                                           `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                                           `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                                           `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                                           `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                                           `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                            `json:"category_id,omitempty"`
	Description                *string                                         `json:"description,omitempty"`
	DescriptionEnHtml          *string                                         `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                                         `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                                         `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                                         `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                                         `json:"description_en,omitempty"`
	EditDate                   *int                                            `json:"edit_date,omitempty"`
	EmailProvider              interface{}                                     `json:"email_provider,omitempty"`
	EmailType                  *string                                         `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                                            `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                                     `json:"feedback_data,omitempty"`
	Guarantee                  interface{}                                     `json:"guarantee,omitempty"`
	HasPendingAutoBuy          *bool                                           `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                  *bool                                           `json:"isIgnored,omitempty"`
	IsSticky                   *int                                            `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                                         `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                                         `json:"item_domain,omitempty"`
	ItemID                     *int                                            `json:"item_id,omitempty"`
	ItemOrigin                 *string                                         `json:"item_origin,omitempty"`
	ItemState                  *string                                         `json:"item_state,omitempty"`
	NoteText                   interface{}                                     `json:"note_text,omitempty"`
	Nsb                        *int                                            `json:"nsb,omitempty"`
	Price                      *int                                            `json:"price,omitempty"`
	PriceWithSellerFee         *float64                                        `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                                         `json:"price_currency,omitempty"`
	PublishedDate              *int                                            `json:"published_date,omitempty"`
	RefreshedDate              *int                                            `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                                         `json:"resale_item_origin,omitempty"`
	RubPrice                   *int                                            `json:"rub_price,omitempty"`
	Seller                     *TelegramResponseItemsItemSeller                `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                                           `json:"showGetEmailCodeButton,omitempty"`
	Tags                       interface{}                                     `json:"tags,omitempty"`
	TelegramAdminCount         *int                                            `json:"telegram_admin_count,omitempty"`
	TelegramAdminSubsCount     *int                                            `json:"telegram_admin_subs_count,omitempty"`
	TelegramBirthday           *int                                            `json:"telegram_birthday,omitempty"`
	TelegramChannelsCount      *int                                            `json:"telegram_channels_count,omitempty"`
	TelegramChatsCount         *int                                            `json:"telegram_chats_count,omitempty"`
	TelegramContactsCount      *int                                            `json:"telegram_contacts_count,omitempty"`
	TelegramConversationsCount *int                                            `json:"telegram_conversations_count,omitempty"`
	TelegramCountry            *string                                         `json:"telegram_country,omitempty"`
	TelegramGroupCounters      *TelegramResponseItemsItemTelegramGroupCounters `json:"telegram_group_counters,omitempty"`
	TelegramIDCount            *int                                            `json:"telegram_id_count,omitempty"`
	TelegramItemID             *int                                            `json:"telegram_item_id,omitempty"`
	TelegramLastSeen           *int                                            `json:"telegram_last_seen,omitempty"`
	TelegramPassword           *int                                            `json:"telegram_password,omitempty"`
	TelegramPremium            *int                                            `json:"telegram_premium,omitempty"`
	TelegramPremiumExpires     *int                                            `json:"telegram_premium_expires,omitempty"`
	TelegramSpamBlock          interface{}                                     `json:"telegram_spam_block,omitempty"`
	TelegramStarsCount         *int                                            `json:"telegram_stars_count,omitempty"`
	Title                      *string                                         `json:"title,omitempty"`
	TitleEn                    *string                                         `json:"title_en,omitempty"`
	UpdateStatDate             *int                                            `json:"update_stat_date,omitempty"`
	ViewCount                  *int                                            `json:"view_count,omitempty"`
}

type TelegramResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type TelegramResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     interface{} `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type TelegramResponseItemsItemTelegramGroupCounters struct {
	Admin         *int `json:"admin,omitempty"`
	Channels      *int `json:"channels,omitempty"`
	Chats         *int `json:"chats,omitempty"`
	Conversations *int `json:"conversations,omitempty"`
}

type TelegramResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type TempEmailPasswordResponse struct {
	Item *TempEmailPasswordResponseItem `json:"item,omitempty"`
}

type TempEmailPasswordResponseItem struct {
	Account *string `json:"account,omitempty"`
}

type TikTokParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// Has linked mobile.
	Tel *CategorySearchTel `json:"tel,omitempty"`
	// How old is the account.
	Reg *int `json:"reg,omitempty"`
	// In what notation is time measured.
	RegPeriod *CategorySearchRegPeriod `json:"reg_period,omitempty"`
	// Minimum number of followers.
	FollowersMin *int `json:"followers_min,omitempty"`
	// Maximum number of followers.
	FollowersMax *int `json:"followers_max,omitempty"`
	// Minimum number of posts.
	PostMin *int `json:"post_min,omitempty"`
	// Maximum number of posts.
	PostMax *int `json:"post_max,omitempty"`
	// Minimum number of likes.
	LikeMin *int `json:"like_min,omitempty"`
	// Maximum number of likes.
	LikeMax *int `json:"like_max,omitempty"`
	// Minimum number of coins.
	CoinsMin *int `json:"coins_min,omitempty"`
	// Maximum number of coins.
	CoinsMax *int `json:"coins_max,omitempty"`
	// Login by cookies.
	CookieLogin *CategorySearchCookieLogin `json:"cookie_login,omitempty"`
	// Has verified.
	Verified *CategorySearchVerified `json:"verified,omitempty"`
	// Has linked email.
	Email *CategorySearchEmail `json:"email,omitempty"`
}

type TikTokResponse struct {
	CacheTTL        *int                      `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                     `json:"hasNextPage,omitempty"`
	Items           []TikTokResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                      `json:"lastModified,omitempty"`
	Page            *int                      `json:"page,omitempty"`
	PerPage         *int                      `json:"perPage,omitempty"`
	SearchUrl       *string                   `json:"searchUrl,omitempty"`
	ServerTime      *int                      `json:"serverTime,omitempty"`
	StickyItems     []interface{}             `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo           `json:"system_info,omitempty"`
	TotalItems      *int                      `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}               `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                     `json:"wasCached,omitempty"`
}

type TikTokResponseItemsItem struct {
	AccountLink                *string                                   `json:"accountLink,omitempty"`
	AccountLinks               []TikTokResponseItemsItemAccountLinksItem `json:"accountLinks,omitempty"`
	AllowAskDiscount           *int                                      `json:"allow_ask_discount,omitempty"`
	BumpSettings               *TikTokResponseItemsItemBumpSettings      `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                                     `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                                     `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                                     `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                                     `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                                     `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                                     `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                                     `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                                     `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                                     `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                                     `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                                     `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                                     `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                                     `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                                     `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                                     `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                                     `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                      `json:"category_id,omitempty"`
	Description                *string                                   `json:"description,omitempty"`
	DescriptionEnHtml          *string                                   `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                                   `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                                   `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                                   `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                                   `json:"description_en,omitempty"`
	EditDate                   *int                                      `json:"edit_date,omitempty"`
	EmailProvider              interface{}                               `json:"email_provider,omitempty"`
	EmailType                  *string                                   `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                                      `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                               `json:"feedback_data,omitempty"`
	Guarantee                  interface{}                               `json:"guarantee,omitempty"`
	HasPendingAutoBuy          *bool                                     `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                  *bool                                     `json:"isIgnored,omitempty"`
	IsSticky                   *int                                      `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                                   `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                                   `json:"item_domain,omitempty"`
	ItemID                     *int                                      `json:"item_id,omitempty"`
	ItemOrigin                 *string                                   `json:"item_origin,omitempty"`
	ItemState                  *string                                   `json:"item_state,omitempty"`
	NoteText                   interface{}                               `json:"note_text,omitempty"`
	Nsb                        *int                                      `json:"nsb,omitempty"`
	Price                      *int                                      `json:"price,omitempty"`
	PriceWithSellerFee         *float64                                  `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                                   `json:"price_currency,omitempty"`
	PublishedDate              *int                                      `json:"published_date,omitempty"`
	RefreshedDate              *int                                      `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                                   `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount  *int                                      `json:"restore_items_category_count,omitempty"`
	RubPrice                   *int                                      `json:"rub_price,omitempty"`
	Seller                     *TikTokResponseItemsItemSeller            `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                                     `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount     *int                                      `json:"sold_items_category_count,omitempty"`
	Tags                       interface{}                               `json:"tags,omitempty"`
	Title                      *string                                   `json:"title,omitempty"`
	TitleEn                    *string                                   `json:"title_en,omitempty"`
	TtCoins                    *int                                      `json:"tt_coins,omitempty"`
	TtCookieLogin              *int                                      `json:"tt_cookie_login,omitempty"`
	TtCountries                *string                                   `json:"tt_countries,omitempty"`
	TtCreateTime               *int                                      `json:"tt_createTime,omitempty"`
	TtFollowers                *int                                      `json:"tt_followers,omitempty"`
	TtFollowing                *int                                      `json:"tt_following,omitempty"`
	TtHasEmail                 *int                                      `json:"tt_hasEmail,omitempty"`
	TtHasLivePermission        *int                                      `json:"tt_hasLivePermission,omitempty"`
	TtHasMobile                *int                                      `json:"tt_hasMobile,omitempty"`
	TtID                       *int                                      `json:"tt_id,omitempty"`
	TtItemID                   *int                                      `json:"tt_item_id,omitempty"`
	TtLikes                    *int                                      `json:"tt_likes,omitempty"`
	TtPermalink                *string                                   `json:"tt_permalink,omitempty"`
	TtPrivateAccount           *int                                      `json:"tt_privateAccount,omitempty"`
	TtScreenName               *string                                   `json:"tt_screen_name,omitempty"`
	TtTopCountry               *string                                   `json:"tt_top_country,omitempty"`
	TtUniqueId                 *string                                   `json:"tt_uniqueId,omitempty"`
	TtVerified                 *int                                      `json:"tt_verified,omitempty"`
	TtVideos                   *int                                      `json:"tt_videos,omitempty"`
	UpdateStatDate             *int                                      `json:"update_stat_date,omitempty"`
	ViewCount                  *int                                      `json:"view_count,omitempty"`
}

type TikTokResponseItemsItemAccountLinksItem struct {
	IconClass *string `json:"iconClass,omitempty"`
	Link      *string `json:"link,omitempty"`
	Text      *string `json:"text,omitempty"`
}

type TikTokResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type TikTokResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type TikTokResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type TransferParams struct {
	// Transfer comment.
	Comment *string `json:"comment,omitempty"`
	// Hold length option.
	HoldLengthOption *PaymentsTransferHoldLengthOption `json:"hold_length_option,omitempty"`
	// Hold length value.
	HoldLengthValue *int `json:"hold_length_value,omitempty"`
	// Is the deal happening on Telegram?
	TelegramDeal *bool `json:"telegram_deal,omitempty"`
	// Telegram username of the user you are dialoguing with.
	TelegramUsername *string `json:"telegram_username,omitempty"`
	// Hold transfer or not.
	TransferHold *bool `json:"transfer_hold,omitempty"`
	// User id of receiver. If **user_id** specified, **username** is not required.
	UserID *int `json:"user_id,omitempty"`
	// Username of receiver. If **username** specified, **user_id** is not required.
	Username *string `json:"username,omitempty"`
}

type UplayParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// List of games.
	Game []string `json:"game[],omitempty"`
	// List of allowed countries.
	Country []string `json:"country[],omitempty"`
	// List of disallowed countries.
	NotCountry []string `json:"not_country[],omitempty"`
	// Number of days the account has been offline.
	Daybreak *int `json:"daybreak,omitempty"`
	// Minimum count of games.
	Gmin *int `json:"gmin,omitempty"`
	// Maximum count of games.
	Gmax *int `json:"gmax,omitempty"`
	// Name of subscription.
	Subscription *CategorySearchSubscription `json:"subscription,omitempty"`
	// Length of subscription.
	SubscriptionLength *int `json:"subscription_length,omitempty"`
	// In what notation is time measured.
	SubscriptionPeriod *CategorySearchSubscriptionPeriod `json:"subscription_period,omitempty"`
	// Minimum level in Tom Clancy's Rainbow Six Siege.
	R6LevelMin *int `json:"r6_level_min,omitempty"`
	// Maximum level in Tom Clancy's Rainbow Six Siege.
	R6LevelMax *int `json:"r6_level_max,omitempty"`
	// Minimum rank points in Tom Clancy's Rainbow Six Siege.
	R6RankMin *int `json:"r6_rank_min,omitempty"`
	// Maximum rank points in Tom Clancy's Rainbow Six Siege.
	R6RankMax *int `json:"r6_rank_max,omitempty"`
	// Minimum count of operators in Tom Clancy's Rainbow Six Siege.
	R6OperatorsMin *int `json:"r6_operators_min,omitempty"`
	// Maximum count of operators in Tom Clancy's Rainbow Six Siege.
	R6OperatorsMax *int `json:"r6_operators_max,omitempty"`
	// Is account banned in Tom Clancy's Rainbow Six Siege
	R6Ban *CategorySearchR6Ban `json:"r6_ban,omitempty"`
	// Minimum number of skins in Tom Clancy's Rainbow Six Siege.
	R6Smin *int `json:"r6_smin,omitempty"`
	// Maximum number of skins in Tom Clancy's Rainbow Six Siege.
	R6Smax *int `json:"r6_smax,omitempty"`
	// List of weapon skins in Tom Clancy's Rainbow Six Siege.
	R6Skin []string `json:"r6_skin[],omitempty"`
	// List of operators in Tom Clancy's Rainbow Six Siege.
	R6Operator []string `json:"r6_operator[],omitempty"`
	// Xbox connected to account.
	XboxConnected *CategorySearchXboxConnected `json:"xbox_connected,omitempty"`
	// PSN connected to account.
	PsnConnected *CategorySearchPsnConnected `json:"psn_connected,omitempty"`
	// Steam connected to account.
	SteamConnected *CategorySearchSteamConnected `json:"steam_connected,omitempty"`
	// Minimum balance.
	BalanceMin *float64 `json:"balance_min,omitempty"`
	// Maximum balance.
	BalanceMax *float64 `json:"balance_max,omitempty"`
	// Has transactions.
	Transactions *CategorySearchTransactions `json:"transactions,omitempty"`
	// How old is the account.
	Reg *int `json:"reg,omitempty"`
	// In what notation is time measured.
	RegPeriod *CategorySearchRegPeriod `json:"reg_period,omitempty"`
}

type UplayResponse struct {
	CacheTTL        *int                     `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                    `json:"hasNextPage,omitempty"`
	Items           []UplayResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                     `json:"lastModified,omitempty"`
	Page            *int                     `json:"page,omitempty"`
	PerPage         *int                     `json:"perPage,omitempty"`
	SearchUrl       *string                  `json:"searchUrl,omitempty"`
	ServerTime      *int                     `json:"serverTime,omitempty"`
	StickyItems     []interface{}            `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo          `json:"system_info,omitempty"`
	TotalItems      *int                     `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}              `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                    `json:"wasCached,omitempty"`
}

type UplayResponseItemsItem struct {
	AccountLastActivity        *int                                    `json:"account_last_activity,omitempty"`
	AllowAskDiscount           *int                                    `json:"allow_ask_discount,omitempty"`
	BumpSettings               *UplayResponseItemsItemBumpSettings     `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                                   `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                                   `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                                   `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                                   `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                                   `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                                   `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                                   `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                                   `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                                   `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                                   `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                                   `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                                   `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                                   `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                                   `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                                   `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                                   `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                    `json:"category_id,omitempty"`
	Description                *string                                 `json:"description,omitempty"`
	DescriptionEnHtml          *string                                 `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                                 `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                                 `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                                 `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                                 `json:"description_en,omitempty"`
	EditDate                   *int                                    `json:"edit_date,omitempty"`
	EmailLoginUrl              *string                                 `json:"emailLoginUrl,omitempty"`
	EmailProvider              *string                                 `json:"email_provider,omitempty"`
	EmailType                  *string                                 `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                                    `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                             `json:"feedback_data,omitempty"`
	Guarantee                  interface{}                             `json:"guarantee,omitempty"`
	HasPendingAutoBuy          *bool                                   `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                  *bool                                   `json:"isIgnored,omitempty"`
	IsSmallExf                 *bool                                   `json:"isSmallExf,omitempty"`
	IsSticky                   *int                                    `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                                 `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                                 `json:"item_domain,omitempty"`
	ItemID                     *int                                    `json:"item_id,omitempty"`
	ItemOrigin                 *string                                 `json:"item_origin,omitempty"`
	ItemState                  *string                                 `json:"item_state,omitempty"`
	NoteText                   interface{}                             `json:"note_text,omitempty"`
	Nsb                        *int                                    `json:"nsb,omitempty"`
	Price                      *int                                    `json:"price,omitempty"`
	PriceWithSellerFee         *float64                                `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                                 `json:"price_currency,omitempty"`
	PublishedDate              *int                                    `json:"published_date,omitempty"`
	R6Operators                []UplayResponseItemsItemR6OperatorsItem `json:"r6Operators,omitempty"`
	R6Skins                    interface{}                             `json:"r6Skins,omitempty"`
	RefreshedDate              *int                                    `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                                 `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount  *int                                    `json:"restore_items_category_count,omitempty"`
	RubPrice                   *int                                    `json:"rub_price,omitempty"`
	Seller                     *UplayResponseItemsItemSeller           `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                                   `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount     *int                                    `json:"sold_items_category_count,omitempty"`
	Tags                       interface{}                             `json:"tags,omitempty"`
	Title                      *string                                 `json:"title,omitempty"`
	TitleEn                    *string                                 `json:"title_en,omitempty"`
	UpdateStatDate             *int                                    `json:"update_stat_date,omitempty"`
	UplayLinkedAccounts        *string                                 `json:"uplayLinkedAccounts,omitempty"`
	UplayR6Rank                *string                                 `json:"uplayR6Rank,omitempty"`
	UplayCountry               *string                                 `json:"uplay_country,omitempty"`
	UplayCreatedDate           *int                                    `json:"uplay_created_date,omitempty"`
	UplayGameCount             *int                                    `json:"uplay_game_count,omitempty"`
	UplayGames                 *UplayResponseItemsItemUplayGames       `json:"uplay_games,omitempty"`
	UplayItemID                *int                                    `json:"uplay_item_id,omitempty"`
	UplayLastActivity          *int                                    `json:"uplay_last_activity,omitempty"`
	UplayPsnConnected          *int                                    `json:"uplay_psn_connected,omitempty"`
	UplayR6                    *bool                                   `json:"uplay_r6,omitempty"`
	UplayR6Ban                 *int                                    `json:"uplay_r6_ban,omitempty"`
	UplayR6BanActive           *bool                                   `json:"uplay_r6_ban_active,omitempty"`
	UplayR6ExternalWarning     *bool                                   `json:"uplay_r6_external_warning,omitempty"`
	UplayR6Level               *int                                    `json:"uplay_r6_level,omitempty"`
	UplayR6Operators           *string                                 `json:"uplay_r6_operators,omitempty"`
	UplayR6OperatorsCount      *int                                    `json:"uplay_r6_operators_count,omitempty"`
	UplayR6Skins               *string                                 `json:"uplay_r6_skins,omitempty"`
	UplayR6SkinsCount          *int                                    `json:"uplay_r6_skins_count,omitempty"`
	UplayR6SteamWarning        *bool                                   `json:"uplay_r6_steam_warning,omitempty"`
	UplaySteamConnected        *int                                    `json:"uplay_steam_connected,omitempty"`
	UplaySubscription          *string                                 `json:"uplay_subscription,omitempty"`
	UplaySubscriptionEndDate   *int                                    `json:"uplay_subscription_end_date,omitempty"`
	UplayXboxConnected         *int                                    `json:"uplay_xbox_connected,omitempty"`
	ViewCount                  *int                                    `json:"view_count,omitempty"`
}

type UplayResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type UplayResponseItemsItemR6OperatorsItem struct {
	Img  *string `json:"img,omitempty"`
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}

type UplayResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type UplayResponseItemsItemUplayGames struct {
	FfffffffFfffFfffFfffFfffffffffff *UplayResponseItemsItemUplayGamesFfffffffFfffFfffFfffFfffffffffff `json:"ffffffff-ffff-ffff-ffff-ffffffffffff,omitempty"`
}

type UplayResponseItemsItemUplayGamesFfffffffFfffFfffFfffFfffffffffff struct {
	Abbr          *string `json:"abbr,omitempty"`
	GameId        *string `json:"gameId,omitempty"`
	Img           *string `json:"img,omitempty"`
	PveTimePlayed *int    `json:"pveTimePlayed,omitempty"`
	PvpTimePlayed *int    `json:"pvpTimePlayed,omitempty"`
	Title         *string `json:"title,omitempty"`
}

type UplayResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type UserModel struct {
	ActiveItemsCount            *int                      `json:"active_items_count,omitempty"`
	ActivityVisible             *bool                     `json:"activity_visible,omitempty"`
	Age                         *int                      `json:"age,omitempty"`
	Balance                     *string                   `json:"balance,omitempty"`
	Balances                    []UserModelBalancesItem   `json:"balances,omitempty"`
	BumpItemPeriod              *int                      `json:"bump_item_period,omitempty"`
	CanEdit                     *bool                     `json:"can_edit,omitempty"`
	CanFollow                   *bool                     `json:"can_follow,omitempty"`
	CanIgnore                   *bool                     `json:"can_ignore,omitempty"`
	CanPostProfile              *bool                     `json:"can_post_profile,omitempty"`
	CanViewProfile              *bool                     `json:"can_view_profile,omitempty"`
	CanViewProfilePosts         *bool                     `json:"can_view_profile_posts,omitempty"`
	CanWarn                     *bool                     `json:"can_warn,omitempty"`
	ContestCount                *int                      `json:"contest_count,omitempty"`
	ConvWelcomeMessage          *string                   `json:"conv_welcome_message,omitempty"`
	ConvertedBalance            *int                      `json:"convertedBalance,omitempty"`
	ConvertedDeposit            *int                      `json:"convertedDeposit,omitempty"`
	ConvertedHold               *int                      `json:"convertedHold,omitempty"`
	Currency                    *string                   `json:"currency,omitempty"`
	CurrencyPhrase              *string                   `json:"currencyPhrase,omitempty"`
	CustomAccountDownloadFormat *string                   `json:"custom_account_download_format,omitempty"`
	CustomFields                *UserModelCustomFields    `json:"custom_fields,omitempty"`
	CustomTitle                 *string                   `json:"custom_title,omitempty"`
	Deposit                     *int                      `json:"deposit,omitempty"`
	Dob                         *UserModelDob             `json:"dob,omitempty"`
	FeedbackData                interface{}               `json:"feedback_data,omitempty"`
	Hold                        *string                   `json:"hold,omitempty"`
	Homepage                    *string                   `json:"homepage,omitempty"`
	IMAPData                    interface{}               `json:"imap_data,omitempty"`
	IsAdmin                     *bool                     `json:"is_admin,omitempty"`
	IsBanned                    *bool                     `json:"is_banned,omitempty"`
	IsFollowed                  *bool                     `json:"is_followed,omitempty"`
	IsIgnored                   *bool                     `json:"is_ignored,omitempty"`
	IsModerator                 *bool                     `json:"is_moderator,omitempty"`
	IsStaff                     *bool                     `json:"is_staff,omitempty"`
	IsSuperAdmin                *bool                     `json:"is_super_admin,omitempty"`
	JoinedDate                  *int                      `json:"joined_date,omitempty"`
	LastActivity                *int                      `json:"last_activity,omitempty"`
	Like2Count                  *int                      `json:"like2_count,omitempty"`
	LikeCount                   *int                      `json:"like_count,omitempty"`
	Location                    *string                   `json:"location,omitempty"`
	MarketCustomTitle           *string                   `json:"market_custom_title,omitempty"`
	MaxDiscountPercent          *int                      `json:"max_discount_percent,omitempty"`
	MessageCount                *int                      `json:"message_count,omitempty"`
	PaidMailLeft                *int                      `json:"paid_mail_left,omitempty"`
	PublicTags                  []UserModelPublicTagsItem `json:"public_tags,omitempty"`
	RegisterDate                *int                      `json:"register_date,omitempty"`
	Rendered                    *UserModelRendered        `json:"rendered,omitempty"`
	RestoreCount                *int                      `json:"restore_count,omitempty"`
	RestoreData                 interface{}               `json:"restore_data,omitempty"`
	ShortLink                   *string                   `json:"short_link,omitempty"`
	SoldItemsCount              *int                      `json:"sold_items_count,omitempty"`
	Tags                        interface{}               `json:"tags,omitempty"`
	TelegramClient              interface{}               `json:"telegram_client,omitempty"`
	TrophyPoints                *int                      `json:"trophy_points,omitempty"`
	UserAllowAskDiscount        *bool                     `json:"user_allow_ask_discount,omitempty"`
	UserID                      *int                      `json:"user_id,omitempty"`
	UserTitle                   *string                   `json:"user_title,omitempty"`
	Username                    *string                   `json:"username,omitempty"`
	ViewURL                     *string                   `json:"view_url,omitempty"`
	Visible                     *bool                     `json:"visible,omitempty"`
	WarningPoints               *int                      `json:"warning_points,omitempty"`
}

type UserModelBalancesItem struct {
	Balance          *string     `json:"balance,omitempty"`
	BalanceID        *int        `json:"balance_id,omitempty"`
	ConvertedBalance *float64    `json:"convertedBalance,omitempty"`
	CustomTitle      interface{} `json:"custom_title,omitempty"`
	FullTitle        *string     `json:"fullTitle,omitempty"`
	MerchantID       *int        `json:"merchant_id,omitempty"`
	Title            *string     `json:"title,omitempty"`
	Type_            *string     `json:"type,omitempty"`
	UserID           *int        `json:"user_id,omitempty"`
}

type UserModelCustomFields struct {
	Field4                *string       `json:"_4,omitempty"`
	AllowSelfUnban        []interface{} `json:"allowSelfUnban,omitempty"`
	BanReason             *string       `json:"ban_reason,omitempty"`
	Discord               *string       `json:"discord,omitempty"`
	FavoriteAnime         *string       `json:"favoriteAnime,omitempty"`
	FavoritePorn          *string       `json:"favoritePorn,omitempty"`
	FavoriteVape          *string       `json:"favoriteVape,omitempty"`
	Github                *string       `json:"github,omitempty"`
	Jabber                *string       `json:"jabber,omitempty"`
	LztAwardUserTrophy    *string       `json:"lztAwardUserTrophy,omitempty"`
	LztLikesIncreasing    *string       `json:"lztLikesIncreasing,omitempty"`
	LztLikesZeroing       *string       `json:"lztLikesZeroing,omitempty"`
	LztSympathyIncreasing *string       `json:"lztSympathyIncreasing,omitempty"`
	LztSympathyZeroing    *string       `json:"lztSympathyZeroing,omitempty"`
	LztUnbanAmount        *string       `json:"lztUnbanAmount,omitempty"`
	MaecenasValue         *string       `json:"maecenasValue,omitempty"`
	Matrix                *string       `json:"matrix,omitempty"`
	ScamURL               *string       `json:"scamURL,omitempty"`
	Steam                 *string       `json:"steam,omitempty"`
	Telegram              *string       `json:"telegram,omitempty"`
	Vk                    *string       `json:"vk,omitempty"`
}

type UserModelDob struct {
	Day   *int `json:"day,omitempty"`
	Month *int `json:"month,omitempty"`
	Year  *int `json:"year,omitempty"`
}

type UserModelIMAPData struct {
	DomainZone *UserModelIMAPDataDomainZone `json:"domain.zone,omitempty"`
}

type UserModelIMAPDataDomainZone struct {
	Domain     *string `json:"domain,omitempty"`
	IMAPServer *string `json:"imap_server,omitempty"`
	Port       *int    `json:"port,omitempty"`
	Secure     *bool   `json:"secure,omitempty"`
}

type UserModelPublicTagsItem struct {
	BackgroundColor *string `json:"background_color,omitempty"`
	TagID           *int    `json:"tag_id,omitempty"`
	Title           *string `json:"title,omitempty"`
}

type UserModelRendered struct {
	Avatars     *UserModelRenderedAvatars `json:"avatars,omitempty"`
	Backgrounds interface{}               `json:"backgrounds,omitempty"`
	Link        *string                   `json:"link,omitempty"`
	Username    *string                   `json:"username,omitempty"`
}

type UserModelRenderedAvatars struct {
	L *string `json:"l,omitempty"`
	M *string `json:"m,omitempty"`
	S *string `json:"s,omitempty"`
}

type UserModelRenderedBackgrounds struct {
	L *string `json:"l,omitempty"`
	M *string `json:"m,omitempty"`
}

type UserModelTagsItem struct {
	Bc                   *string `json:"bc,omitempty"`
	ForOwnedAccountsOnly *bool   `json:"forOwnedAccountsOnly,omitempty"`
	IsDefault            *bool   `json:"isDefault,omitempty"`
	TagID                *int    `json:"tag_id,omitempty"`
	Title                *string `json:"title,omitempty"`
}

type UserModelTelegramClient struct {
	TelegramAPIHash        *string `json:"telegram_api_hash,omitempty"`
	TelegramAPIID          *string `json:"telegram_api_id,omitempty"`
	TelegramAppVersion     *string `json:"telegram_app_version,omitempty"`
	TelegramDeviceModel    *string `json:"telegram_device_model,omitempty"`
	TelegramLangCode       *string `json:"telegram_lang_code,omitempty"`
	TelegramLangPack       *string `json:"telegram_lang_pack,omitempty"`
	TelegramSystemLangCode *string `json:"telegram_system_lang_code,omitempty"`
	TelegramSystemVersion  *string `json:"telegram_system_version,omitempty"`
}

type UserParams struct {
	// User id.
	UserID *int `json:"user_id,omitempty"`
	// Accounts category.
	CategoryID *AccountsListCategoryID `json:"category_id,omitempty"`
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Account status.
	Show *AccountsListShow `json:"show,omitempty"`
	// Delete reason. (Only if **show** is set to **deleted**)
	DeleteReason *string `json:"delete_reason,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// Login.
	Login *string `json:"login,omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Order by.
	OrderBy *AccountsListOrderBy `json:"order_by,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Username of buyer. (If **show** is **paid**)
	Username *string `json:"username,omitempty"`
	// Start date for filtering by publication date.
	PublishedStartDate *string `json:"published_startDate,omitempty"`
	// End date for filtering by publication date.
	PublishedEndDate *string `json:"published_endDate,omitempty"`
	// Enable filtering by publication date.
	FilterByPublishedDate *bool `json:"filter_by_published_date,omitempty"`
	// Start date for filtering by buyer operation date.
	PaidStartDate *string `json:"paid_startDate,omitempty"`
	// End date for filtering by buyer operation date.
	PaidEndDate *string `json:"paid_endDate,omitempty"`
	// Enable filtering by buyer operation date.
	FilterByBuyerOperationDate *bool `json:"filter_by_buyer_operation_date,omitempty"`
	// Start date for filtering by deletion date.
	DeleteStartDate *string `json:"delete_startDate,omitempty"`
	// End date for filtering by deletion date.
	DeleteEndDate *string `json:"delete_endDate,omitempty"`
	// Enable filtering by deletion date.
	FilterByDeleteDate *bool `json:"filter_by_delete_date,omitempty"`
}

type ViewedParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Account status.
	Show *AccountsListShow `json:"show,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Order by.
	OrderBy *AccountsListOrderBy `json:"order_by,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
}

type VpnParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// List of allowed VPN services.
	Service []string `json:"service[],omitempty"`
	// Length of subscription.
	SubscriptionLength *int `json:"subscription_length,omitempty"`
	// In what notation is time measured.
	SubscriptionPeriod *CategorySearchSubscriptionPeriod `json:"subscription_period,omitempty"`
	// Is auto renewal enabled.
	Autorenewal *CategorySearchAutorenewal `json:"autorenewal,omitempty"`
}

type VpnResponse struct {
	CacheTTL        *int                   `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                  `json:"hasNextPage,omitempty"`
	Items           []VpnResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                   `json:"lastModified,omitempty"`
	Page            *int                   `json:"page,omitempty"`
	PerPage         *int                   `json:"perPage,omitempty"`
	SearchUrl       *string                `json:"searchUrl,omitempty"`
	ServerTime      *int                   `json:"serverTime,omitempty"`
	StickyItems     []interface{}          `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo        `json:"system_info,omitempty"`
	TotalItems      *int                   `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}            `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                  `json:"wasCached,omitempty"`
}

type VpnResponseItemsItem struct {
	AllowAskDiscount           *int                              `json:"allow_ask_discount,omitempty"`
	BumpSettings               *VpnResponseItemsItemBumpSettings `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                             `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                             `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                             `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                             `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                             `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                             `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                             `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                             `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                             `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                             `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                             `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                             `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                             `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                             `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                             `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                             `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                              `json:"category_id,omitempty"`
	Description                *string                           `json:"description,omitempty"`
	DescriptionEnHtml          *string                           `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                           `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                           `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                           `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                           `json:"description_en,omitempty"`
	EditDate                   *int                              `json:"edit_date,omitempty"`
	EmailProvider              interface{}                       `json:"email_provider,omitempty"`
	EmailType                  *string                           `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                              `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                       `json:"feedback_data,omitempty"`
	Guarantee                  interface{}                       `json:"guarantee,omitempty"`
	HasPendingAutoBuy          *bool                             `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                  *bool                             `json:"isIgnored,omitempty"`
	IsSticky                   *int                              `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                           `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                           `json:"item_domain,omitempty"`
	ItemID                     *int                              `json:"item_id,omitempty"`
	ItemOrigin                 *string                           `json:"item_origin,omitempty"`
	ItemState                  *string                           `json:"item_state,omitempty"`
	NoteText                   interface{}                       `json:"note_text,omitempty"`
	Nsb                        *int                              `json:"nsb,omitempty"`
	Price                      *int                              `json:"price,omitempty"`
	PriceWithSellerFee         *float64                          `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                           `json:"price_currency,omitempty"`
	PublishedDate              *int                              `json:"published_date,omitempty"`
	RefreshedDate              *int                              `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                           `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount  *int                              `json:"restore_items_category_count,omitempty"`
	RubPrice                   *int                              `json:"rub_price,omitempty"`
	Seller                     *VpnResponseItemsItemSeller       `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                             `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount     *int                              `json:"sold_items_category_count,omitempty"`
	Tags                       interface{}                       `json:"tags,omitempty"`
	Title                      *string                           `json:"title,omitempty"`
	TitleEn                    *string                           `json:"title_en,omitempty"`
	UpdateStatDate             *int                              `json:"update_stat_date,omitempty"`
	ViewCount                  *int                              `json:"view_count,omitempty"`
	VpnProductTitle            *string                           `json:"vpnProductTitle,omitempty"`
	VpnExpireDate              *int                              `json:"vpn_expire_date,omitempty"`
	VpnItemID                  *int                              `json:"vpn_item_id,omitempty"`
	VpnRenewable               *int                              `json:"vpn_renewable,omitempty"`
	VpnService                 *string                           `json:"vpn_service,omitempty"`
}

type VpnResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type VpnResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type VpnResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type WOTParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// Has linked mobile.
	Tel *CategorySearchTel `json:"tel,omitempty"`
	// Number of days the account has been offline.
	Daybreak *int `json:"daybreak,omitempty"`
	// Minimum number of battles.
	BattlesMin *int `json:"battles_min,omitempty"`
	// Maximum number of battles.
	BattlesMax *int `json:"battles_max,omitempty"`
	// Minimum number of gold.
	GoldMin *int `json:"gold_min,omitempty"`
	// Maximum number of gold.
	GoldMax *int `json:"gold_max,omitempty"`
	// Minimum number of silver.
	SilverMin *int `json:"silver_min,omitempty"`
	// Maximum number of silver.
	SilverMax *int `json:"silver_max,omitempty"`
	// Minimum number of top tanks.
	TopMin *int `json:"top_min,omitempty"`
	// Maximum number of top tanks.
	TopMax *int `json:"top_max,omitempty"`
	// Minimum number of premium tanks.
	PremMin *int `json:"prem_min,omitempty"`
	// Maximum number of premium tanks.
	PremMax *int `json:"prem_max,omitempty"`
	// Minimum number of top premium tanks.
	TopPremMin *int `json:"top_prem_min,omitempty"`
	// Maximum number of top premium tanks.
	TopPremMax *int `json:"top_prem_max,omitempty"`
	// Minimum number of wins.
	WinPmin *int `json:"win_pmin,omitempty"`
	// Maximum number of wins.
	WinPmax *int `json:"win_pmax,omitempty"`
	// List of tanks.
	Tank []int `json:"tank[],omitempty"`
	// Region.
	Region []string `json:"region[],omitempty"`
	// Exclude region.
	NotRegion []string `json:"not_region[],omitempty"`
	// Has a premium subscription.
	Premium *CategorySearchPremium `json:"premium,omitempty"`
	// When premium subscription will be active
	PremiumExpiration *int `json:"premium_expiration,omitempty"`
	// In what notation is time measured
	PremiumExpirationPeriod *CategorySearchPremiumExpirationPeriod `json:"premium_expiration_period,omitempty"`
	// Has clan.
	Clan *CategorySearchClan `json:"clan,omitempty"`
	// List of allowed clan role.
	ClanRole []string `json:"clan_role[],omitempty"`
	// List of disallowed clan role.
	NotClanRole []string `json:"not_clan_role[],omitempty"`
	// Minimum number of gold in clan treasure.
	ClanGoldMin *int `json:"clan_gold_min,omitempty"`
	// Maximum number of gold in clan treasure.
	ClanGoldMax *int `json:"clan_gold_max,omitempty"`
	// Minimum number of credits in clan treasure.
	ClanCreditsMin *int `json:"clan_credits_min,omitempty"`
	// Maximum number of credits in clan treasure.
	ClanCreditsMax *int `json:"clan_credits_max,omitempty"`
	// Minimum number of crystal in clan treasure.
	ClanCrystalMin *int `json:"clan_crystal_min,omitempty"`
	// Maximum number of crystal in clan treasure.
	ClanCrystalMax *int `json:"clan_crystal_max,omitempty"`
	// List of allowed countries.
	Country []string `json:"country[],omitempty"`
	// List of disallowed countries.
	NotCountry []string `json:"not_country[],omitempty"`
}

type WOTResponse struct {
	CacheTTL        *int                   `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                  `json:"hasNextPage,omitempty"`
	Items           []WOTResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                   `json:"lastModified,omitempty"`
	Page            *int                   `json:"page,omitempty"`
	PerPage         *int                   `json:"perPage,omitempty"`
	SearchUrl       *string                `json:"searchUrl,omitempty"`
	ServerTime      *int                   `json:"serverTime,omitempty"`
	StickyItems     []interface{}          `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo        `json:"system_info,omitempty"`
	TotalItems      *int                   `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}            `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                  `json:"wasCached,omitempty"`
}

type WOTResponseItemsItem struct {
	AccountLinks               []interface{}                     `json:"accountLinks,omitempty"`
	AccountLastActivity        *int                              `json:"account_last_activity,omitempty"`
	AllowAskDiscount           *int                              `json:"allow_ask_discount,omitempty"`
	BumpSettings               *WOTResponseItemsItemBumpSettings `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                             `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                             `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                             `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                             `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                             `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                             `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                             `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                             `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                             `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                             `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                             `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                             `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                             `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                             `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                             `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                             `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                              `json:"category_id,omitempty"`
	Description                *string                           `json:"description,omitempty"`
	DescriptionEnHtml          *string                           `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                           `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                           `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                           `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                           `json:"description_en,omitempty"`
	EditDate                   *int                              `json:"edit_date,omitempty"`
	EmailProvider              interface{}                       `json:"email_provider,omitempty"`
	EmailType                  *string                           `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                              `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                       `json:"feedback_data,omitempty"`
	Guarantee                  interface{}                       `json:"guarantee,omitempty"`
	HasPendingAutoBuy          *bool                             `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                  *bool                             `json:"isIgnored,omitempty"`
	IsSmallExf                 *bool                             `json:"isSmallExf,omitempty"`
	IsSticky                   *int                              `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                           `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                           `json:"item_domain,omitempty"`
	ItemID                     *int                              `json:"item_id,omitempty"`
	ItemOrigin                 *string                           `json:"item_origin,omitempty"`
	ItemState                  *string                           `json:"item_state,omitempty"`
	NoteText                   interface{}                       `json:"note_text,omitempty"`
	Nsb                        *int                              `json:"nsb,omitempty"`
	Price                      *int                              `json:"price,omitempty"`
	PriceWithSellerFee         *float64                          `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                           `json:"price_currency,omitempty"`
	PublishedDate              *int                              `json:"published_date,omitempty"`
	RefreshedDate              *int                              `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                           `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount  *int                              `json:"restore_items_category_count,omitempty"`
	RubPrice                   *int                              `json:"rub_price,omitempty"`
	Seller                     *WOTResponseItemsItemSeller       `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                             `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount     *int                              `json:"sold_items_category_count,omitempty"`
	Tags                       interface{}                       `json:"tags,omitempty"`
	Title                      *string                           `json:"title,omitempty"`
	TitleEn                    *string                           `json:"title_en,omitempty"`
	UpdateStatDate             *int                              `json:"update_stat_date,omitempty"`
	ViewCount                  *int                              `json:"view_count,omitempty"`
	WotLauncherTitle           *string                           `json:"wotLauncherTitle,omitempty"`
	WotPremiumTankCount        *int                              `json:"wotPremiumTankCount,omitempty"`
	WotPremiumTanks            map[string]interface{}            `json:"wotPremiumTanks,omitempty"`
	WotRegionPhrase            *string                           `json:"wotRegionPhrase,omitempty"`
	WotTankCount               *int                              `json:"wotTankCount,omitempty"`
	WotTanks                   map[string]interface{}            `json:"wotTanks,omitempty"`
	WotTopPremiumTanks         interface{}                       `json:"wotTopPremiumTanks,omitempty"`
	WotTopTanks                interface{}                       `json:"wotTopTanks,omitempty"`
	WOTBattleCount             *int                              `json:"wot_battle_count,omitempty"`
	WOTBlitz                   *int                              `json:"wot_blitz,omitempty"`
	WOTCredits                 *int                              `json:"wot_credits,omitempty"`
	WOTGold                    *int                              `json:"wot_gold,omitempty"`
	WOTHasClan                 *bool                             `json:"wot_has_clan,omitempty"`
	WOTItemID                  *int                              `json:"wot_item_id,omitempty"`
	WOTLastActivity            *int                              `json:"wot_last_activity,omitempty"`
	WOTLossCount               *int                              `json:"wot_loss_count,omitempty"`
	WOTMobile                  *int                              `json:"wot_mobile,omitempty"`
	WOTPremium                 *int                              `json:"wot_premium,omitempty"`
	WOTPremiumExpires          *int                              `json:"wot_premium_expires,omitempty"`
	WOTPremiumTanks            *int                              `json:"wot_premium_tanks,omitempty"`
	WOTRegion                  *string                           `json:"wot_region,omitempty"`
	WOTRegisterDate            *int                              `json:"wot_register_date,omitempty"`
	WOTTopPremiumTanks         *int                              `json:"wot_top_premium_tanks,omitempty"`
	WOTTopTanks                *int                              `json:"wot_top_tanks,omitempty"`
	WOTWinCount                *int                              `json:"wot_win_count,omitempty"`
	WOTWinCountPercents        *int                              `json:"wot_win_count_percents,omitempty"`
}

type WOTResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type WOTResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type WOTResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type WarfaceParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Minimum rank.
	RankMin *int `json:"rank_min,omitempty"`
	// Maximum rank.
	RankMax *int `json:"rank_max,omitempty"`
	// Minimum bonus rank.
	BonusRankMin *int `json:"bonus_rank_min,omitempty"`
	// Maximum bonus rank.
	BonusRankMax *int `json:"bonus_rank_max,omitempty"`
	// Has linked mobile.
	Tel *CategorySearchTel `json:"tel,omitempty"`
	// Number of days the account has been offline.
	Daybreak *int `json:"daybreak,omitempty"`
	// Minimum amount of Kredits.
	KreditsMin *int `json:"kredits_min,omitempty"`
	// Maximum amount of Kredits.
	KreditsMax *int `json:"kredits_max,omitempty"`
	// Minimum total donated Kredits.
	TotalKreditsMin *int `json:"total_kredits_min,omitempty"`
	// Maximum total donated Kredits.
	TotalKreditsMax *int `json:"total_kredits_max,omitempty"`
}

type WarfaceResponse struct {
	CacheTTL        *int                       `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                      `json:"hasNextPage,omitempty"`
	Items           []WarfaceResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                       `json:"lastModified,omitempty"`
	Page            *int                       `json:"page,omitempty"`
	PerPage         *int                       `json:"perPage,omitempty"`
	SearchUrl       *string                    `json:"searchUrl,omitempty"`
	ServerTime      *int                       `json:"serverTime,omitempty"`
	StickyItems     []interface{}              `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo            `json:"system_info,omitempty"`
	TotalItems      *int                       `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}                `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                      `json:"wasCached,omitempty"`
}

type WarfaceResponseItemsItem struct {
	AccountLastActivity        *int                                    `json:"account_last_activity,omitempty"`
	AllowAskDiscount           *int                                    `json:"allow_ask_discount,omitempty"`
	BumpSettings               *WarfaceResponseItemsItemBumpSettings   `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                                   `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                                   `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                                   `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                                   `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                                   `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                                   `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                                   `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                                   `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                                   `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                                   `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                                   `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                                   `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                                   `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                                   `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                                   `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                                   `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                    `json:"category_id,omitempty"`
	Description                *string                                 `json:"description,omitempty"`
	DescriptionEnHtml          *string                                 `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                                 `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                                 `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                                 `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                                 `json:"description_en,omitempty"`
	Domain                     *string                                 `json:"domain,omitempty"`
	EditDate                   *int                                    `json:"edit_date,omitempty"`
	EmailProvider              interface{}                             `json:"email_provider,omitempty"`
	EmailType                  *string                                 `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                                    `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                             `json:"feedback_data,omitempty"`
	Guarantee                  interface{}                             `json:"guarantee,omitempty"`
	HasPendingAutoBuy          *bool                                   `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                  *bool                                   `json:"isIgnored,omitempty"`
	IsSmallExf                 *bool                                   `json:"isSmallExf,omitempty"`
	IsSticky                   *int                                    `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                                 `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                                 `json:"item_domain,omitempty"`
	ItemID                     *int                                    `json:"item_id,omitempty"`
	ItemOrigin                 *string                                 `json:"item_origin,omitempty"`
	ItemState                  *string                                 `json:"item_state,omitempty"`
	NoteText                   interface{}                             `json:"note_text,omitempty"`
	Nsb                        *int                                    `json:"nsb,omitempty"`
	Price                      *int                                    `json:"price,omitempty"`
	PriceWithSellerFee         *float64                                `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                                 `json:"price_currency,omitempty"`
	PublishedDate              *int                                    `json:"published_date,omitempty"`
	RefreshedDate              *int                                    `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                                 `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount  *int                                    `json:"restore_items_category_count,omitempty"`
	RubPrice                   *int                                    `json:"rub_price,omitempty"`
	Seller                     *WarfaceResponseItemsItemSeller         `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                                   `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount     *int                                    `json:"sold_items_category_count,omitempty"`
	Tags                       interface{}                             `json:"tags,omitempty"`
	Title                      *string                                 `json:"title,omitempty"`
	TitleEn                    *string                                 `json:"title_en,omitempty"`
	UpdateStatDate             *int                                    `json:"update_stat_date,omitempty"`
	ViewCount                  *int                                    `json:"view_count,omitempty"`
	WfActiveLoan               *int                                    `json:"wf_active_loan,omitempty"`
	WfBonusRank                *int                                    `json:"wf_bonus_rank,omitempty"`
	WfItemID                   *int                                    `json:"wf_item_id,omitempty"`
	WfLastGameDate             *int                                    `json:"wf_last_game_date,omitempty"`
	WfLoan                     *bool                                   `json:"wf_loan,omitempty"`
	WfMailMobile               *int                                    `json:"wf_mail_mobile,omitempty"`
	WfMobile                   *int                                    `json:"wf_mobile,omitempty"`
	WfPlayers                  *bool                                   `json:"wf_players,omitempty"`
	WfRank                     *int                                    `json:"wf_rank,omitempty"`
	WfServer1                  *int                                    `json:"wf_server_1,omitempty"`
	WfServer2                  *int                                    `json:"wf_server_2,omitempty"`
	WfServer3                  *int                                    `json:"wf_server_3,omitempty"`
	WfServers                  []WarfaceResponseItemsItemWfServersItem `json:"wf_servers,omitempty"`
}

type WarfaceResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type WarfaceResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type WarfaceResponseItemsItemWfServersItem struct {
	ID    *int    `json:"id,omitempty"`
	Rank  *int    `json:"rank,omitempty"`
	Title *string `json:"title,omitempty"`
}

type WarfaceResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type WotBlitzParams struct {
	// The number of the page to display results from.
	Page *int `json:"page,omitempty"`
	// Minimal price of account (Inclusive).
	Pmin *int `json:"pmin,omitempty"`
	// Maximum price of account (Inclusive).
	Pmax *int `json:"pmax,omitempty"`
	// The word or words contained in the account title.
	Title *string `json:"title,omitempty"`
	// Order by.
	OrderBy *CategorySearchOrderBy `json:"order_by,omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	TagID []int `json:"tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotTagID []int `json:"not_tag_id[],omitempty"`
	// List of tag ids (Tag list is available via **GET /me**).
	PublicTagID []int `json:"public_tag_id[],omitempty"`
	// List of tag ids that won't be included (Tag list is available via **GET /me**).
	NotPublicTagID []int `json:"not_public_tag_id[],omitempty"`
	// List of account origins.
	Origin []string `json:"origin[],omitempty"`
	// List of account origins that won't be included.
	NotOrigin []string `json:"not_origin[],omitempty"`
	// Search accounts of user.
	UserID *int `json:"user_id,omitempty"`
	// Not sold before.
	Nsb *bool `json:"nsb,omitempty"`
	// Sold before.
	Sb *bool `json:"sb,omitempty"`
	// Not sold by me before.
	NsbByMe *bool `json:"nsb_by_me,omitempty"`
	// Sold by me before.
	SbByMe *bool `json:"sb_by_me,omitempty"`
	// Currency in which the cost of the account will be searched.
	Currency *CategorySearchCurrency `json:"currency,omitempty"`
	// Has email login data.
	EmailLoginData *bool `json:"email_login_data,omitempty"`
	// Email provider.
	EmailProvider []string `json:"email_provider[],omitempty"`
	// Email provider.
	NotEmailProvider *CategorySearchNotEmailProvider `json:"not_email_provider[],omitempty"`
	// Parse same item ids.
	ParseSameItemIds *bool `json:"parse_same_item_ids,omitempty"`
	// Email type.
	EmailType []string `json:"email_type[],omitempty"`
	// Domain of native/autoreg email.
	ItemDomain *string `json:"item_domain,omitempty"`
	// Has linked mobile.
	Tel *CategorySearchTel `json:"tel,omitempty"`
	// Number of days the account has been offline.
	Daybreak *int `json:"daybreak,omitempty"`
	// Minimum number of battles.
	BattlesMin *int `json:"battles_min,omitempty"`
	// Maximum number of battles.
	BattlesMax *int `json:"battles_max,omitempty"`
	// Minimum number of gold.
	GoldMin *int `json:"gold_min,omitempty"`
	// Maximum number of gold.
	GoldMax *int `json:"gold_max,omitempty"`
	// Minimum number of silver.
	SilverMin *int `json:"silver_min,omitempty"`
	// Maximum number of silver.
	SilverMax *int `json:"silver_max,omitempty"`
	// Minimum number of top tanks.
	TopMin *int `json:"top_min,omitempty"`
	// Maximum number of top tanks.
	TopMax *int `json:"top_max,omitempty"`
	// Minimum number of premium tanks.
	PremMin *int `json:"prem_min,omitempty"`
	// Maximum number of premium tanks.
	PremMax *int `json:"prem_max,omitempty"`
	// Minimum number of top premium tanks.
	TopPremMin *int `json:"top_prem_min,omitempty"`
	// Maximum number of top premium tanks.
	TopPremMax *int `json:"top_prem_max,omitempty"`
	// Minimum number of wins.
	WinPmin *int `json:"win_pmin,omitempty"`
	// Maximum number of wins.
	WinPmax *int `json:"win_pmax,omitempty"`
	// List of tanks.
	Tank []int `json:"tank[],omitempty"`
	// Region.
	Region []string `json:"region[],omitempty"`
	// Exclude region.
	NotRegion []string `json:"not_region[],omitempty"`
	// Has a premium subscription.
	Premium *CategorySearchPremium `json:"premium,omitempty"`
	// When premium subscription will be active
	PremiumExpiration *int `json:"premium_expiration,omitempty"`
	// In what notation is time measured
	PremiumExpirationPeriod *CategorySearchPremiumExpirationPeriod `json:"premium_expiration_period,omitempty"`
	// Has clan.
	Clan *CategorySearchClan `json:"clan,omitempty"`
	// List of allowed clan role.
	ClanRole []string `json:"clan_role[],omitempty"`
	// List of disallowed clan role.
	NotClanRole []string `json:"not_clan_role[],omitempty"`
	// Minimum number of gold in clan treasure.
	ClanGoldMin *int `json:"clan_gold_min,omitempty"`
	// Maximum number of gold in clan treasure.
	ClanGoldMax *int `json:"clan_gold_max,omitempty"`
	// Minimum number of credits in clan treasure.
	ClanCreditsMin *int `json:"clan_credits_min,omitempty"`
	// Maximum number of credits in clan treasure.
	ClanCreditsMax *int `json:"clan_credits_max,omitempty"`
	// Minimum number of crystal in clan treasure.
	ClanCrystalMin *int `json:"clan_crystal_min,omitempty"`
	// Maximum number of crystal in clan treasure.
	ClanCrystalMax *int `json:"clan_crystal_max,omitempty"`
	// List of allowed countries.
	Country []string `json:"country[],omitempty"`
	// List of disallowed countries.
	NotCountry []string `json:"not_country[],omitempty"`
}

type WotBlitzResponse struct {
	CacheTTL        *int                        `json:"cacheTTL,omitempty"`
	HasNextPage     *bool                       `json:"hasNextPage,omitempty"`
	Items           []WotBlitzResponseItemsItem `json:"items,omitempty"`
	LastModified    *int                        `json:"lastModified,omitempty"`
	Page            *int                        `json:"page,omitempty"`
	PerPage         *int                        `json:"perPage,omitempty"`
	SearchUrl       *string                     `json:"searchUrl,omitempty"`
	ServerTime      *int                        `json:"serverTime,omitempty"`
	StickyItems     []interface{}               `json:"stickyItems,omitempty"`
	SystemInfo      *RespSystemInfo             `json:"system_info,omitempty"`
	TotalItems      *int                        `json:"totalItems,omitempty"`
	TotalItemsPrice interface{}                 `json:"totalItemsPrice,omitempty"`
	WasCached       *bool                       `json:"wasCached,omitempty"`
}

type WotBlitzResponseItemsItem struct {
	AccountLinks               []interface{}                          `json:"accountLinks,omitempty"`
	AccountLastActivity        *int                                   `json:"account_last_activity,omitempty"`
	AllowAskDiscount           *int                                   `json:"allow_ask_discount,omitempty"`
	BumpSettings               *WotBlitzResponseItemsItemBumpSettings `json:"bumpSettings,omitempty"`
	CanBumpItem                *bool                                  `json:"canBumpItem,omitempty"`
	CanBuyItem                 *bool                                  `json:"canBuyItem,omitempty"`
	CanChangePassword          *bool                                  `json:"canChangePassword,omitempty"`
	CanCloseItem               *bool                                  `json:"canCloseItem,omitempty"`
	CanDeleteItem              *bool                                  `json:"canDeleteItem,omitempty"`
	CanEditItem                *bool                                  `json:"canEditItem,omitempty"`
	CanOpenItem                *bool                                  `json:"canOpenItem,omitempty"`
	CanReportItem              *bool                                  `json:"canReportItem,omitempty"`
	CanResellItemAfterPurchase *bool                                  `json:"canResellItemAfterPurchase,omitempty"`
	CanStickItem               *bool                                  `json:"canStickItem,omitempty"`
	CanUnstickItem             *bool                                  `json:"canUnstickItem,omitempty"`
	CanUpdateItemStats         *bool                                  `json:"canUpdateItemStats,omitempty"`
	CanValidateAccount         *bool                                  `json:"canValidateAccount,omitempty"`
	CanViewAccountLink         *bool                                  `json:"canViewAccountLink,omitempty"`
	CanViewEmailLoginData      *bool                                  `json:"canViewEmailLoginData,omitempty"`
	CanViewLoginData           *bool                                  `json:"canViewLoginData,omitempty"`
	CategoryID                 *int                                   `json:"category_id,omitempty"`
	Description                *string                                `json:"description,omitempty"`
	DescriptionEnHtml          *string                                `json:"descriptionEnHtml,omitempty"`
	DescriptionEnPlain         *string                                `json:"descriptionEnPlain,omitempty"`
	DescriptionHtml            *string                                `json:"descriptionHtml,omitempty"`
	DescriptionPlain           *string                                `json:"descriptionPlain,omitempty"`
	DescriptionEn              *string                                `json:"description_en,omitempty"`
	EditDate                   *int                                   `json:"edit_date,omitempty"`
	EmailProvider              interface{}                            `json:"email_provider,omitempty"`
	EmailType                  *string                                `json:"email_type,omitempty"`
	ExtendedGuarantee          *int                                   `json:"extended_guarantee,omitempty"`
	FeedbackData               interface{}                            `json:"feedback_data,omitempty"`
	Guarantee                  interface{}                            `json:"guarantee,omitempty"`
	HasPendingAutoBuy          *bool                                  `json:"hasPendingAutoBuy,omitempty"`
	IsIgnored                  *bool                                  `json:"isIgnored,omitempty"`
	IsSmallExf                 *bool                                  `json:"isSmallExf,omitempty"`
	IsSticky                   *int                                   `json:"is_sticky,omitempty"`
	ItemOriginPhrase           *string                                `json:"itemOriginPhrase,omitempty"`
	ItemDomain                 *string                                `json:"item_domain,omitempty"`
	ItemID                     *int                                   `json:"item_id,omitempty"`
	ItemOrigin                 *string                                `json:"item_origin,omitempty"`
	ItemState                  *string                                `json:"item_state,omitempty"`
	NoteText                   interface{}                            `json:"note_text,omitempty"`
	Nsb                        *int                                   `json:"nsb,omitempty"`
	Price                      *int                                   `json:"price,omitempty"`
	PriceWithSellerFee         *float64                               `json:"priceWithSellerFee,omitempty"`
	PriceCurrency              *string                                `json:"price_currency,omitempty"`
	PublishedDate              *int                                   `json:"published_date,omitempty"`
	RefreshedDate              *int                                   `json:"refreshed_date,omitempty"`
	ResaleItemOrigin           *string                                `json:"resale_item_origin,omitempty"`
	RestoreItemsCategoryCount  *int                                   `json:"restore_items_category_count,omitempty"`
	RubPrice                   *int                                   `json:"rub_price,omitempty"`
	Seller                     *WotBlitzResponseItemsItemSeller       `json:"seller,omitempty"`
	ShowGetEmailCodeButton     *bool                                  `json:"showGetEmailCodeButton,omitempty"`
	SoldItemsCategoryCount     *int                                   `json:"sold_items_category_count,omitempty"`
	Tags                       interface{}                            `json:"tags,omitempty"`
	Title                      *string                                `json:"title,omitempty"`
	TitleEn                    *string                                `json:"title_en,omitempty"`
	UpdateStatDate             *int                                   `json:"update_stat_date,omitempty"`
	ViewCount                  *int                                   `json:"view_count,omitempty"`
	WotLauncherTitle           *string                                `json:"wotLauncherTitle,omitempty"`
	WotPremiumTankCount        *int                                   `json:"wotPremiumTankCount,omitempty"`
	WotPremiumTanks            map[string]interface{}                 `json:"wotPremiumTanks,omitempty"`
	WotRegionPhrase            *string                                `json:"wotRegionPhrase,omitempty"`
	WotTankCount               *int                                   `json:"wotTankCount,omitempty"`
	WotTanks                   map[string]interface{}                 `json:"wotTanks,omitempty"`
	WotTopPremiumTanks         interface{}                            `json:"wotTopPremiumTanks,omitempty"`
	WotTopTanks                interface{}                            `json:"wotTopTanks,omitempty"`
	WOTBattleCount             *int                                   `json:"wot_battle_count,omitempty"`
	WOTBlitz                   *int                                   `json:"wot_blitz,omitempty"`
	WOTCredits                 *int                                   `json:"wot_credits,omitempty"`
	WOTGold                    *int                                   `json:"wot_gold,omitempty"`
	WOTHasClan                 *bool                                  `json:"wot_has_clan,omitempty"`
	WOTItemID                  *int                                   `json:"wot_item_id,omitempty"`
	WOTLastActivity            *int                                   `json:"wot_last_activity,omitempty"`
	WOTLossCount               *int                                   `json:"wot_loss_count,omitempty"`
	WOTMobile                  *int                                   `json:"wot_mobile,omitempty"`
	WOTPremium                 *int                                   `json:"wot_premium,omitempty"`
	WOTPremiumExpires          *int                                   `json:"wot_premium_expires,omitempty"`
	WOTPremiumTanks            *int                                   `json:"wot_premium_tanks,omitempty"`
	WOTRegion                  *string                                `json:"wot_region,omitempty"`
	WOTRegisterDate            *int                                   `json:"wot_register_date,omitempty"`
	WOTTopPremiumTanks         *int                                   `json:"wot_top_premium_tanks,omitempty"`
	WOTTopTanks                *int                                   `json:"wot_top_tanks,omitempty"`
	WOTWinCount                *int                                   `json:"wot_win_count,omitempty"`
	WOTWinCountPercents        *int                                   `json:"wot_win_count_percents,omitempty"`
}

type WotBlitzResponseItemsItemBumpSettings struct {
	CanBumpItem         *bool       `json:"canBumpItem,omitempty"`
	CanBumpItemGlobally *bool       `json:"canBumpItemGlobally,omitempty"`
	ErrorPhrase         interface{} `json:"errorPhrase,omitempty"`
	ShortErrorPhrase    interface{} `json:"shortErrorPhrase,omitempty"`
}

type WotBlitzResponseItemsItemSeller struct {
	ActiveItemsCount    *int        `json:"active_items_count,omitempty"`
	AvatarDate          *int        `json:"avatar_date,omitempty"`
	DisplayStyleGroupID *int        `json:"display_style_group_id,omitempty"`
	IsBanned            *int        `json:"is_banned,omitempty"`
	RestoreData         interface{} `json:"restore_data,omitempty"`
	RestorePercents     *int        `json:"restore_percents,omitempty"`
	SoldItemsCount      *int        `json:"sold_items_count,omitempty"`
	UserID              *int        `json:"user_id,omitempty"`
	Username            *string     `json:"username,omitempty"`
}

type WotBlitzResponseSystemInfo struct {
	LogID     *int `json:"log_id,omitempty"`
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

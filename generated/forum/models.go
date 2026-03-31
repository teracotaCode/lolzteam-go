package forum

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

type AllParams struct {
	// The time in milliseconds (e.g. 1767214800) before last content date.
	Before *int `json:"before,omitempty"`
	// Id of the container forum to search for contents. Child forums of the specified forum will be included in the search.
	ForumID *int `json:"forum_id,omitempty"`
	// Number of results in a page.
	Limit *int `json:"limit,omitempty"`
	// Page number of results.
	Page *int `json:"page,omitempty"`
	// Search query. Can be skipped if **user_id** is set.
	Q *string `json:"q,omitempty"`
	// Tag to search for tagged contents.
	Tag    *string     `json:"tag,omitempty"`
	UserID StringOrInt `json:"user_id,omitempty"`
}

type AllResponse struct {
	Data       []AllResponseDataItem  `json:"data,omitempty"`
	DataTotal  *int                   `json:"data_total,omitempty"`
	Links      *AllResponseLinks      `json:"links,omitempty"`
	SystemInfo *RespSystemInfo        `json:"system_info,omitempty"`
	Users      []AllResponseUsersItem `json:"users,omitempty"`
}

type AllResponseDataItem struct {
	ContentID           interface{}                     `json:"content_id,omitempty"`
	ContentType         *string                         `json:"content_type,omitempty"`
	CreatorUserID       *int                            `json:"creator_user_id,omitempty"`
	CreatorUsername     *string                         `json:"creator_username,omitempty"`
	CreatorUsernameHTML *string                         `json:"creator_username_html,omitempty"`
	FirstPost           *AllResponseDataItemFirstPost   `json:"first_post,omitempty"`
	Forum               *AllResponseDataItemForum       `json:"forum,omitempty"`
	ForumID             *int                            `json:"forum_id,omitempty"`
	LastPost            *AllResponseDataItemLastPost    `json:"last_post,omitempty"`
	Links               *AllResponseDataItemLinks       `json:"links,omitempty"`
	NodeTitle           *string                         `json:"node_title,omitempty"`
	Permissions         *AllResponseDataItemPermissions `json:"permissions,omitempty"`
	ThreadCreateDate    *int                            `json:"thread_create_date,omitempty"`
	ThreadID            *int                            `json:"thread_id,omitempty"`
	ThreadIsClosed      *bool                           `json:"thread_is_closed,omitempty"`
	ThreadIsDeleted     *bool                           `json:"thread_is_deleted,omitempty"`
	ThreadIsFollowed    *bool                           `json:"thread_is_followed,omitempty"`
	ThreadIsPublished   *bool                           `json:"thread_is_published,omitempty"`
	ThreadIsStarred     *bool                           `json:"thread_is_starred,omitempty"`
	ThreadIsSticky      *bool                           `json:"thread_is_sticky,omitempty"`
	ThreadPostCount     *int                            `json:"thread_post_count,omitempty"`
	ThreadPrefixes      []interface{}                   `json:"thread_prefixes,omitempty"`
	ThreadTags          interface{}                     `json:"thread_tags,omitempty"`
	ThreadTitle         *string                         `json:"thread_title,omitempty"`
	ThreadUpdateDate    *int                            `json:"thread_update_date,omitempty"`
	ThreadViewCount     *int                            `json:"thread_view_count,omitempty"`
	UserIsIgnored       *bool                           `json:"user_is_ignored,omitempty"`
}

type AllResponseDataItemFirstPost struct {
	Links              *AllResponseDataItemFirstPostLinks       `json:"links,omitempty"`
	Permissions        *AllResponseDataItemFirstPostPermissions `json:"permissions,omitempty"`
	PostBody           *string                                  `json:"post_body,omitempty"`
	PostBodyHTML       *string                                  `json:"post_body_html,omitempty"`
	PostBodyPlainText  *string                                  `json:"post_body_plain_text,omitempty"`
	PostCreateDate     *int                                     `json:"post_create_date,omitempty"`
	PostID             *int                                     `json:"post_id,omitempty"`
	PostIsDeleted      *bool                                    `json:"post_is_deleted,omitempty"`
	PostIsFirstPost    *bool                                    `json:"post_is_first_post,omitempty"`
	PostIsLiked        *bool                                    `json:"post_is_liked,omitempty"`
	PostIsPublished    *bool                                    `json:"post_is_published,omitempty"`
	PostLikeCount      *int                                     `json:"post_like_count,omitempty"`
	PostUpdateDate     *int                                     `json:"post_update_date,omitempty"`
	PosterUserID       *int                                     `json:"poster_user_id,omitempty"`
	PosterUsername     *string                                  `json:"poster_username,omitempty"`
	PosterUsernameHTML *string                                  `json:"poster_username_html,omitempty"`
	Signature          *string                                  `json:"signature,omitempty"`
	SignatureHTML      *string                                  `json:"signature_html,omitempty"`
	SignaturePlainText *string                                  `json:"signature_plain_text,omitempty"`
	ThreadID           *int                                     `json:"thread_id,omitempty"`
	ThreadIsDeleted    *bool                                    `json:"thread_is_deleted,omitempty"`
	UserIsIgnored      *bool                                    `json:"user_is_ignored,omitempty"`
}

type AllResponseDataItemFirstPostLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type AllResponseDataItemFirstPostPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	Like   *bool `json:"like,omitempty"`
	Reply  *bool `json:"reply,omitempty"`
	Report *bool `json:"report,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type AllResponseDataItemForum struct {
	ForumDescription       *string                              `json:"forum_description,omitempty"`
	ForumID                *int                                 `json:"forum_id,omitempty"`
	ForumIsFollowed        *bool                                `json:"forum_is_followed,omitempty"`
	ForumPostCount         *int                                 `json:"forum_post_count,omitempty"`
	ForumPrefixes          []interface{}                        `json:"forum_prefixes,omitempty"`
	ForumThreadCount       *int                                 `json:"forum_thread_count,omitempty"`
	ForumTitle             *string                              `json:"forum_title,omitempty"`
	Links                  *AllResponseDataItemForumLinks       `json:"links,omitempty"`
	ParentNodeID           *int                                 `json:"parent_node_id,omitempty"`
	Permissions            *AllResponseDataItemForumPermissions `json:"permissions,omitempty"`
	ThreadDefaultPrefixID  *int                                 `json:"thread_default_prefix_id,omitempty"`
	ThreadPrefixIsRequired *bool                                `json:"thread_prefix_is_required,omitempty"`
}

type AllResponseDataItemForumLinks struct {
	Detail        *string `json:"detail,omitempty"`
	Followers     *string `json:"followers,omitempty"`
	Permalink     *string `json:"permalink,omitempty"`
	SubCategories *string `json:"sub-categories,omitempty"`
	SubForums     *string `json:"sub-forums,omitempty"`
	Threads       *string `json:"threads,omitempty"`
}

type AllResponseDataItemForumPermissions struct {
	CreateThread *bool `json:"create_thread,omitempty"`
	Delete       *bool `json:"delete,omitempty"`
	Edit         *bool `json:"edit,omitempty"`
	Follow       *bool `json:"follow,omitempty"`
	TagThread    *bool `json:"tag_thread,omitempty"`
	View         *bool `json:"view,omitempty"`
}

type AllResponseDataItemLastPost struct {
	Links              *AllResponseDataItemLastPostLinks       `json:"links,omitempty"`
	Permissions        *AllResponseDataItemLastPostPermissions `json:"permissions,omitempty"`
	PostBody           *string                                 `json:"post_body,omitempty"`
	PostBodyHTML       *string                                 `json:"post_body_html,omitempty"`
	PostBodyPlainText  *string                                 `json:"post_body_plain_text,omitempty"`
	PostCreateDate     *int                                    `json:"post_create_date,omitempty"`
	PostID             *int                                    `json:"post_id,omitempty"`
	PostIsDeleted      *bool                                   `json:"post_is_deleted,omitempty"`
	PostIsFirstPost    *bool                                   `json:"post_is_first_post,omitempty"`
	PostIsLiked        *bool                                   `json:"post_is_liked,omitempty"`
	PostIsPublished    *bool                                   `json:"post_is_published,omitempty"`
	PostLikeCount      *int                                    `json:"post_like_count,omitempty"`
	PostUpdateDate     *int                                    `json:"post_update_date,omitempty"`
	PosterUserID       *int                                    `json:"poster_user_id,omitempty"`
	PosterUsername     *string                                 `json:"poster_username,omitempty"`
	PosterUsernameHTML *string                                 `json:"poster_username_html,omitempty"`
	Signature          *string                                 `json:"signature,omitempty"`
	SignatureHTML      *string                                 `json:"signature_html,omitempty"`
	SignaturePlainText *string                                 `json:"signature_plain_text,omitempty"`
	ThreadID           *int                                    `json:"thread_id,omitempty"`
	ThreadIsDeleted    *bool                                   `json:"thread_is_deleted,omitempty"`
	UserIsIgnored      *bool                                   `json:"user_is_ignored,omitempty"`
}

type AllResponseDataItemLastPostLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type AllResponseDataItemLastPostPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	Like   *bool `json:"like,omitempty"`
	Reply  *bool `json:"reply,omitempty"`
	Report *bool `json:"report,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type AllResponseDataItemLinks struct {
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

type AllResponseDataItemPermissions struct {
	Bump   *AllResponseDataItemPermissionsBump `json:"bump,omitempty"`
	Delete *bool                               `json:"delete,omitempty"`
	Edit   *bool                               `json:"edit,omitempty"`
	Follow *bool                               `json:"follow,omitempty"`
	Post   *bool                               `json:"post,omitempty"`
	View   *bool                               `json:"view,omitempty"`
}

type AllResponseDataItemPermissionsBump struct {
	AvailableCount    *int        `json:"available_count,omitempty"`
	Can               *bool       `json:"can,omitempty"`
	Error             interface{} `json:"error,omitempty"`
	NextAvailableTime interface{} `json:"next_available_time,omitempty"`
}

type AllResponseLinks struct {
	Next  *string `json:"next,omitempty"`
	Page  *int    `json:"page,omitempty"`
	Pages *int    `json:"pages,omitempty"`
}

type AllResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type AllResponseUsersItem struct {
	Balance                     *string                                               `json:"balance,omitempty"`
	Banner                      *string                                               `json:"banner,omitempty"`
	Birthday                    *AllResponseUsersItemBirthday                         `json:"birthday,omitempty"`
	ContestCount                *int                                                  `json:"contest_count,omitempty"`
	ConvWelcomeMessage          *string                                               `json:"conv_welcome_message,omitempty"`
	CuratorTitles               []string                                              `json:"curator_titles,omitempty"`
	Currency                    *string                                               `json:"currency,omitempty"`
	CustomTitle                 *string                                               `json:"custom_title,omitempty"`
	DisplayBannerID             *int                                                  `json:"display_banner_id,omitempty"`
	DisplayIconGroupID          *int                                                  `json:"display_icon_group_id,omitempty"`
	EditPermissions             *AllResponseUsersItemEditPermissions                  `json:"edit_permissions,omitempty"`
	Fields                      []AllResponseUsersItemFieldsItem                      `json:"fields,omitempty"`
	Hold                        *string                                               `json:"hold,omitempty"`
	IsBanned                    *int                                                  `json:"is_banned,omitempty"`
	Links                       *AllResponseUsersItemLinks                            `json:"links,omitempty"`
	Permissions                 *AllResponseUsersItemPermissions                      `json:"permissions,omitempty"`
	SecretAnswerFirstLetter     *string                                               `json:"secret_answer_first_letter,omitempty"`
	SecretAnswerRendered        *string                                               `json:"secret_answer_rendered,omitempty"`
	SelfPermissions             *AllResponseUsersItemSelfPermissions                  `json:"self_permissions,omitempty"`
	ShortLink                   *string                                               `json:"short_link,omitempty"`
	TrophyCount                 *int                                                  `json:"trophy_count,omitempty"`
	UserDeposit                 *int                                                  `json:"user_deposit,omitempty"`
	UserEmail                   *string                                               `json:"user_email,omitempty"`
	UserExternalAuthentications []AllResponseUsersItemUserExternalAuthenticationsItem `json:"user_external_authentications,omitempty"`
	UserFollowers               *AllResponseUsersItemUserFollowers                    `json:"user_followers,omitempty"`
	UserFollowing               *AllResponseUsersItemUserFollowing                    `json:"user_following,omitempty"`
	UserGroupID                 *int                                                  `json:"user_group_id,omitempty"`
	UserGroups                  []AllResponseUsersItemUserGroupsItem                  `json:"user_groups,omitempty"`
	UserID                      *int                                                  `json:"user_id,omitempty"`
	UserIsFollowed              *bool                                                 `json:"user_is_followed,omitempty"`
	UserIsIgnored               *bool                                                 `json:"user_is_ignored,omitempty"`
	UserIsValid                 *bool                                                 `json:"user_is_valid,omitempty"`
	UserIsVerified              *bool                                                 `json:"user_is_verified,omitempty"`
	UserIsVisitor               *bool                                                 `json:"user_is_visitor,omitempty"`
	UserLastSeenDate            *int                                                  `json:"user_last_seen_date,omitempty"`
	UserLike2Count              *int                                                  `json:"user_like2_count,omitempty"`
	UserLikeCount               *int                                                  `json:"user_like_count,omitempty"`
	UserMessageCount            *int                                                  `json:"user_message_count,omitempty"`
	UserRegisterDate            *int                                                  `json:"user_register_date,omitempty"`
	UserTimezoneOffset          *int                                                  `json:"user_timezone_offset,omitempty"`
	UserTitle                   *string                                               `json:"user_title,omitempty"`
	UserUnreadConversationCount *int                                                  `json:"user_unread_conversation_count,omitempty"`
	UserUnreadNotificationCount *int                                                  `json:"user_unread_notification_count,omitempty"`
	Username                    *string                                               `json:"username,omitempty"`
	UsernameHTML                *string                                               `json:"username_html,omitempty"`
}

type AllResponseUsersItemBirthday struct {
	Age       *int                                   `json:"age,omitempty"`
	Format    *string                                `json:"format,omitempty"`
	TimeStamp *AllResponseUsersItemBirthdayTimeStamp `json:"timeStamp,omitempty"`
}

type AllResponseUsersItemBirthdayTimeStamp struct {
	Date         *string `json:"date,omitempty"`
	Timezone     *string `json:"timezone,omitempty"`
	TimezoneType *int    `json:"timezone_type,omitempty"`
}

type AllResponseUsersItemEditPermissions struct {
	Fields            *bool `json:"fields,omitempty"`
	HideUsernameLogs  *bool `json:"hide_username_logs,omitempty"`
	Password          *bool `json:"password,omitempty"`
	PrimaryGroupID    *bool `json:"primary_group_id,omitempty"`
	SecondaryGroupIds *bool `json:"secondary_group_ids,omitempty"`
	ShortLink         *bool `json:"short_link,omitempty"`
	UserDobDay        *bool `json:"user_dob_day,omitempty"`
	UserDobMonth      *bool `json:"user_dob_month,omitempty"`
	UserDobYear       *bool `json:"user_dob_year,omitempty"`
	UserEmail         *bool `json:"user_email,omitempty"`
	UserTitle         *bool `json:"user_title,omitempty"`
	Username          *bool `json:"username,omitempty"`
}

type AllResponseUsersItemFieldsItem struct {
	Choices       []AllResponseUsersItemFieldsItemChoicesItem `json:"choices,omitempty"`
	Description   *string                                     `json:"description,omitempty"`
	ID            *string                                     `json:"id,omitempty"`
	IsMultiChoice *bool                                       `json:"is_multi_choice,omitempty"`
	IsRequired    *bool                                       `json:"is_required,omitempty"`
	Position      *string                                     `json:"position,omitempty"`
	Title         *string                                     `json:"title,omitempty"`
	Value         *string                                     `json:"value,omitempty"`
	Values        interface{}                                 `json:"values,omitempty"`
}

type AllResponseUsersItemFieldsItemChoicesItem struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type AllResponseUsersItemLinks struct {
	Avatar      *string `json:"avatar,omitempty"`
	AvatarBig   *string `json:"avatar_big,omitempty"`
	AvatarSmall *string `json:"avatar_small,omitempty"`
	BackgroundL *string `json:"background_l,omitempty"`
	BackgroundM *string `json:"background_m,omitempty"`
	Detail      *string `json:"detail,omitempty"`
	Followers   *string `json:"followers,omitempty"`
	Followings  *string `json:"followings,omitempty"`
	Ignore      *string `json:"ignore,omitempty"`
	Permalink   *string `json:"permalink,omitempty"`
	Status      *string `json:"status,omitempty"`
	Timeline    *string `json:"timeline,omitempty"`
}

type AllResponseUsersItemPermissions struct {
	Edit        *bool `json:"edit,omitempty"`
	Follow      *bool `json:"follow,omitempty"`
	Ignore      *bool `json:"ignore,omitempty"`
	ProfilePost *bool `json:"profile_post,omitempty"`
}

type AllResponseUsersItemSelfPermissions struct {
	CreateConversation *bool `json:"create_conversation,omitempty"`
}

type AllResponseUsersItemUserExternalAuthenticationsItem struct {
	Provider    *string `json:"provider,omitempty"`
	ProviderKey *string `json:"provider_key,omitempty"`
}

type AllResponseUsersItemUserFollowers struct {
	Count *int                                         `json:"count,omitempty"`
	Users []AllResponseUsersItemUserFollowersUsersItem `json:"users,omitempty"`
}

type AllResponseUsersItemUserFollowersUsersItem struct {
	Avatar       *string `json:"avatar,omitempty"`
	UserID       *int    `json:"user_id,omitempty"`
	Username     *string `json:"username,omitempty"`
	UsernameHTML *string `json:"username_html,omitempty"`
}

type AllResponseUsersItemUserFollowing struct {
	Count *int                                         `json:"count,omitempty"`
	Users []AllResponseUsersItemUserFollowingUsersItem `json:"users,omitempty"`
}

type AllResponseUsersItemUserFollowingUsersItem struct {
	Avatar       *string `json:"avatar,omitempty"`
	UserID       *int    `json:"user_id,omitempty"`
	Username     *string `json:"username,omitempty"`
	UsernameHTML *string `json:"username_html,omitempty"`
}

type AllResponseUsersItemUserGroupsItem struct {
	DisplayBannerSelectable *bool   `json:"display_banner_selectable,omitempty"`
	DisplayGroupSelectable  *bool   `json:"display_group_selectable,omitempty"`
	DisplayIconSelectable   *bool   `json:"display_icon_selectable,omitempty"`
	IsPrimaryGroup          *bool   `json:"is_primary_group,omitempty"`
	UserGroupBannerCSSClass *string `json:"user_group_banner_css_class,omitempty"`
	UserGroupBannerText     *string `json:"user_group_banner_text,omitempty"`
	UserGroupBannerTextEn   *string `json:"user_group_banner_text_en,omitempty"`
	UserGroupIconClass      *string `json:"user_group_icon_class,omitempty"`
	UserGroupID             *int    `json:"user_group_id,omitempty"`
	UserGroupTitle          *string `json:"user_group_title,omitempty"`
	UserGroupTitleEn        *string `json:"user_group_title_en,omitempty"`
}

type BumpResponse struct {
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type BumpResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type CSSParams struct {
	// The names or identifiers of the CSS selectors to retrieve.
	CSS []string `json:"css,omitempty"`
}

type CSSResponse struct {
	Contents   *string         `json:"contents,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type CSSResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ClaimParams struct {
	// Allow ask hidden content.
	AllowAskHiddenContent *bool `json:"allow_ask_hidden_content,omitempty"`
	// Contacts and wallets of the responder. Specify the known data about the responder, if any. Optional if **as_is_market_deal** is 0.
	AsData *string `json:"as_data,omitempty"`
	// Funds transfer receipt. Upload a receipt for the transfer of funds, use the "View receipt" button in your wallet. May be uploaded to [Imgur](https://imgur.com/upload). Write "no" if you have not paid....
	AsFundsReceipt *string `json:"as_funds_receipt,omitempty"`
	// Market item id. Required if **as_is_market_deal** is 1.
	AsMarketItemID *int `json:"as_market_item_id,omitempty"`
	// Screenshot showing the respondent's Telegram login. If the correspondence was conducted in Telegram, upload a screenshot that will display the respondent's Telegram login against the background of you...
	AsTgLoginScreenshot *string `json:"as_tg_login_screenshot,omitempty"`
	// Allow commenting if user can't post in thread.
	CommentIgnoreGroup *bool `json:"comment_ignore_group,omitempty"`
	// Currency of Claim.
	Currency *ThreadsClaimCurrency `json:"currency,omitempty"`
	// Don't alert followers about thread creation.
	DontAlertFollowers *bool `json:"dont_alert_followers,omitempty"`
	// Hide contacts.
	HideContacts *bool `json:"hide_contacts,omitempty"`
	// Pay claim fee now or later. (Only for **transfer_type** = **notsafe**)
	PayClaim *ThreadsClaimPayClaim `json:"pay_claim,omitempty"`
	// Allow to reply only users with chosen or higher group.
	ReplyGroup *ThreadsClaimReplyGroup `json:"reply_group,omitempty"`
	// Date to schedule thread creation (format: `DD-MM-YYYY`).
	ScheduleDate *string `json:"schedule_date,omitempty"`
	// Time to schedule thread creation (format: `HH:MM`).
	ScheduleTime *string `json:"schedule_time,omitempty"`
	// Thread tags.
	Tags []string `json:"tags,omitempty"`
	// Receive forum notifications of new posts in this thread.
	WatchThread *bool `json:"watch_thread,omitempty"`
	// Receive email notifications of new posts in this thread.
	WatchThreadEmail *bool `json:"watch_thread_email,omitempty"`
	// Watch thread state.
	WatchThreadState *bool `json:"watch_thread_state,omitempty"`
}

type ClaimResponse struct {
	SystemInfo *RespSystemInfo  `json:"system_info,omitempty"`
	Thread     *RespThreadModel `json:"thread,omitempty"`
}

type ClaimResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ClaimResponseThread struct {
	Contest             *ClaimResponseThreadContest      `json:"contest,omitempty"`
	CreatorUserID       *int                             `json:"creator_user_id,omitempty"`
	CreatorUsername     *string                          `json:"creator_username,omitempty"`
	CreatorUsernameHTML *string                          `json:"creator_username_html,omitempty"`
	FirstPost           *ClaimResponseThreadFirstPost    `json:"first_post,omitempty"`
	ForumID             *int                             `json:"forum_id,omitempty"`
	LastPost            *ClaimResponseThreadLastPost     `json:"last_post,omitempty"`
	Links               *ClaimResponseThreadLinks        `json:"links,omitempty"`
	NodeTitle           *string                          `json:"node_title,omitempty"`
	Permissions         *ClaimResponseThreadPermissions  `json:"permissions,omitempty"`
	Restrictions        *ClaimResponseThreadRestrictions `json:"restrictions,omitempty"`
	ThreadCreateDate    *int                             `json:"thread_create_date,omitempty"`
	ThreadID            *int                             `json:"thread_id,omitempty"`
	ThreadIsClosed      *bool                            `json:"thread_is_closed,omitempty"`
	ThreadIsDeleted     *bool                            `json:"thread_is_deleted,omitempty"`
	ThreadIsFollowed    *bool                            `json:"thread_is_followed,omitempty"`
	ThreadIsPublished   *bool                            `json:"thread_is_published,omitempty"`
	ThreadIsStarred     *bool                            `json:"thread_is_starred,omitempty"`
	ThreadIsSticky      *bool                            `json:"thread_is_sticky,omitempty"`
	ThreadPostCount     *int                             `json:"thread_post_count,omitempty"`
	ThreadPrefixes      []interface{}                    `json:"thread_prefixes,omitempty"`
	ThreadTags          interface{}                      `json:"thread_tags,omitempty"`
	ThreadTitle         *string                          `json:"thread_title,omitempty"`
	ThreadUpdateDate    *int                             `json:"thread_update_date,omitempty"`
	ThreadViewCount     *int                             `json:"thread_view_count,omitempty"`
	UserIsIgnored       *bool                            `json:"user_is_ignored,omitempty"`
}

type ClaimResponseThreadContest struct {
	AlreadyParticipate    *bool                                  `json:"already_participate,omitempty"`
	ChanceToWin           *float64                               `json:"chance_to_win,omitempty"`
	CountWinners          *int                                   `json:"count_winners,omitempty"`
	FinishDate            *int                                   `json:"finish_date,omitempty"`
	IsFinished            *int                                   `json:"is_finished,omitempty"`
	IsMoneyPlaces         *int                                   `json:"is_money_places,omitempty"`
	NeededMembers         *int                                   `json:"needed_members,omitempty"`
	NowCountMembers       *int                                   `json:"now_count_members,omitempty"`
	Permissions           *ClaimResponseThreadContestPermissions `json:"permissions,omitempty"`
	PrizeData             *int                                   `json:"prize_data,omitempty"`
	PrizeType             *string                                `json:"prize_type,omitempty"`
	PrizeTypePhrase       *string                                `json:"prize_type_phrase,omitempty"`
	RequireLikeCount      *int                                   `json:"require_like_count,omitempty"`
	RequireTotalLikeCount *int                                   `json:"require_total_like_count,omitempty"`
	Type_                 *string                                `json:"type,omitempty"`
	Winners               []int                                  `json:"winners,omitempty"`
}

type ClaimResponseThreadContestPermissions struct {
	CanFinish           *bool   `json:"can_finish,omitempty"`
	CanParticipate      *bool   `json:"can_participate,omitempty"`
	CanParticipateError *string `json:"can_participate_error,omitempty"`
	CanViewUserList     *bool   `json:"can_view_user_list,omitempty"`
}

type ClaimResponseThreadFirstPost struct {
	Links              *ClaimResponseThreadFirstPostLinks       `json:"links,omitempty"`
	Permissions        *ClaimResponseThreadFirstPostPermissions `json:"permissions,omitempty"`
	PostBody           *string                                  `json:"post_body,omitempty"`
	PostBodyHTML       *string                                  `json:"post_body_html,omitempty"`
	PostBodyPlainText  *string                                  `json:"post_body_plain_text,omitempty"`
	PostCreateDate     *int                                     `json:"post_create_date,omitempty"`
	PostID             *int                                     `json:"post_id,omitempty"`
	PostIsDeleted      *bool                                    `json:"post_is_deleted,omitempty"`
	PostIsFirstPost    *bool                                    `json:"post_is_first_post,omitempty"`
	PostIsLiked        *bool                                    `json:"post_is_liked,omitempty"`
	PostIsPublished    *bool                                    `json:"post_is_published,omitempty"`
	PostLikeCount      *int                                     `json:"post_like_count,omitempty"`
	PostUpdateDate     *int                                     `json:"post_update_date,omitempty"`
	PosterUserID       *int                                     `json:"poster_user_id,omitempty"`
	PosterUsername     *string                                  `json:"poster_username,omitempty"`
	PosterUsernameHTML *string                                  `json:"poster_username_html,omitempty"`
	Signature          *string                                  `json:"signature,omitempty"`
	SignatureHTML      *string                                  `json:"signature_html,omitempty"`
	SignaturePlainText *string                                  `json:"signature_plain_text,omitempty"`
	ThreadID           *int                                     `json:"thread_id,omitempty"`
	ThreadIsDeleted    *bool                                    `json:"thread_is_deleted,omitempty"`
	UserIsIgnored      *bool                                    `json:"user_is_ignored,omitempty"`
}

type ClaimResponseThreadFirstPostLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type ClaimResponseThreadFirstPostPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	Like   *bool `json:"like,omitempty"`
	Reply  *bool `json:"reply,omitempty"`
	Report *bool `json:"report,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type ClaimResponseThreadLastPost struct {
	Links              *ClaimResponseThreadLastPostLinks       `json:"links,omitempty"`
	Permissions        *ClaimResponseThreadLastPostPermissions `json:"permissions,omitempty"`
	PostBody           *string                                 `json:"post_body,omitempty"`
	PostBodyHTML       *string                                 `json:"post_body_html,omitempty"`
	PostBodyPlainText  *string                                 `json:"post_body_plain_text,omitempty"`
	PostCreateDate     *int                                    `json:"post_create_date,omitempty"`
	PostID             *int                                    `json:"post_id,omitempty"`
	PostIsDeleted      *bool                                   `json:"post_is_deleted,omitempty"`
	PostIsFirstPost    *bool                                   `json:"post_is_first_post,omitempty"`
	PostIsLiked        *bool                                   `json:"post_is_liked,omitempty"`
	PostIsPublished    *bool                                   `json:"post_is_published,omitempty"`
	PostLikeCount      *int                                    `json:"post_like_count,omitempty"`
	PostUpdateDate     *int                                    `json:"post_update_date,omitempty"`
	PosterUserID       *int                                    `json:"poster_user_id,omitempty"`
	PosterUsername     *string                                 `json:"poster_username,omitempty"`
	PosterUsernameHTML *string                                 `json:"poster_username_html,omitempty"`
	Signature          *string                                 `json:"signature,omitempty"`
	SignatureHTML      *string                                 `json:"signature_html,omitempty"`
	SignaturePlainText *string                                 `json:"signature_plain_text,omitempty"`
	ThreadID           *int                                    `json:"thread_id,omitempty"`
	ThreadIsDeleted    *bool                                   `json:"thread_is_deleted,omitempty"`
	UserIsIgnored      *bool                                   `json:"user_is_ignored,omitempty"`
}

type ClaimResponseThreadLastPostLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type ClaimResponseThreadLastPostPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	Like   *bool `json:"like,omitempty"`
	Reply  *bool `json:"reply,omitempty"`
	Report *bool `json:"report,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type ClaimResponseThreadLinks struct {
	Detail            *string `json:"detail,omitempty"`
	FirstPost         *string `json:"first_post,omitempty"`
	FirstPoster       *string `json:"first_poster,omitempty"`
	FirstPosterAvatar *string `json:"first_poster_avatar,omitempty"`
	Followers         *string `json:"followers,omitempty"`
	Forum             *string `json:"forum,omitempty"`
	LastPost          *string `json:"last_post,omitempty"`
	Permalink         *string `json:"permalink,omitempty"`
	Posts             *string `json:"posts,omitempty"`
}

type ClaimResponseThreadPermissions struct {
	Bump      *ClaimResponseThreadPermissionsBump `json:"bump,omitempty"`
	Delete    *bool                               `json:"delete,omitempty"`
	Edit      *bool                               `json:"edit,omitempty"`
	EditTags  *bool                               `json:"edit_tags,omitempty"`
	EditTitle *bool                               `json:"edit_title,omitempty"`
	Follow    *bool                               `json:"follow,omitempty"`
	Post      *bool                               `json:"post,omitempty"`
	View      *bool                               `json:"view,omitempty"`
}

type ClaimResponseThreadPermissionsBump struct {
	AvailableCount    *int        `json:"available_count,omitempty"`
	Can               *bool       `json:"can,omitempty"`
	Error             interface{} `json:"error,omitempty"`
	NextAvailableTime interface{} `json:"next_available_time,omitempty"`
}

type ClaimResponseThreadRestrictions struct {
	MaxReplyCount *int `json:"max_reply_count,omitempty"`
	ReplyDelay    *int `json:"reply_delay,omitempty"`
}

type ClaimsParams struct {
	// Filter claims by their type.
	Type_ *UsersType `json:"type,omitempty"`
	// Filter claims by their state.
	ClaimState *UsersClaimState `json:"claim_state,omitempty"`
}

type ClaimsResponse struct {
	Claims     []ClaimsResponseClaimsItem `json:"claims,omitempty"`
	Stats      *ClaimsResponseStats       `json:"stats,omitempty"`
	SystemInfo *RespSystemInfo            `json:"system_info,omitempty"`
}

type ClaimsResponseClaimsItem struct {
	Amount               *int           `json:"amount,omitempty"`
	AmountFormatted      *string        `json:"amount_formatted,omitempty"`
	Author               *RespUserModel `json:"author,omitempty"`
	ClaimDate            *int           `json:"claim_date,omitempty"`
	ClaimState           *string        `json:"claim_state,omitempty"`
	MessageBody          *string        `json:"message_body,omitempty"`
	MessageBodyHTML      *string        `json:"message_body_html,omitempty"`
	MessageBodyPlainText *string        `json:"message_body_plain_text,omitempty"`
	ThreadID             *int           `json:"thread_id,omitempty"`
}

type ClaimsResponseClaimsItemAuthor struct {
	Balance                     *string                                                         `json:"balance,omitempty"`
	Banner                      *string                                                         `json:"banner,omitempty"`
	Birthday                    *ClaimsResponseClaimsItemAuthorBirthday                         `json:"birthday,omitempty"`
	ContestCount                *int                                                            `json:"contest_count,omitempty"`
	ConvWelcomeMessage          *string                                                         `json:"conv_welcome_message,omitempty"`
	CuratorTitles               []string                                                        `json:"curator_titles,omitempty"`
	Currency                    *string                                                         `json:"currency,omitempty"`
	CustomTitle                 *string                                                         `json:"custom_title,omitempty"`
	DisplayBannerID             *int                                                            `json:"display_banner_id,omitempty"`
	DisplayIconGroupID          *int                                                            `json:"display_icon_group_id,omitempty"`
	EditPermissions             *ClaimsResponseClaimsItemAuthorEditPermissions                  `json:"edit_permissions,omitempty"`
	Fields                      []ClaimsResponseClaimsItemAuthorFieldsItem                      `json:"fields,omitempty"`
	Hold                        *string                                                         `json:"hold,omitempty"`
	IsBanned                    *int                                                            `json:"is_banned,omitempty"`
	Links                       *ClaimsResponseClaimsItemAuthorLinks                            `json:"links,omitempty"`
	Permissions                 *ClaimsResponseClaimsItemAuthorPermissions                      `json:"permissions,omitempty"`
	SecretAnswerFirstLetter     *string                                                         `json:"secret_answer_first_letter,omitempty"`
	SecretAnswerRendered        *string                                                         `json:"secret_answer_rendered,omitempty"`
	SelfPermissions             *ClaimsResponseClaimsItemAuthorSelfPermissions                  `json:"self_permissions,omitempty"`
	ShortLink                   *string                                                         `json:"short_link,omitempty"`
	TrophyCount                 *int                                                            `json:"trophy_count,omitempty"`
	UserDeposit                 *int                                                            `json:"user_deposit,omitempty"`
	UserEmail                   *string                                                         `json:"user_email,omitempty"`
	UserExternalAuthentications []ClaimsResponseClaimsItemAuthorUserExternalAuthenticationsItem `json:"user_external_authentications,omitempty"`
	UserFollowers               *ClaimsResponseClaimsItemAuthorUserFollowers                    `json:"user_followers,omitempty"`
	UserFollowing               *ClaimsResponseClaimsItemAuthorUserFollowing                    `json:"user_following,omitempty"`
	UserGroupID                 *int                                                            `json:"user_group_id,omitempty"`
	UserGroups                  []ClaimsResponseClaimsItemAuthorUserGroupsItem                  `json:"user_groups,omitempty"`
	UserID                      *int                                                            `json:"user_id,omitempty"`
	UserIsFollowed              *bool                                                           `json:"user_is_followed,omitempty"`
	UserIsIgnored               *bool                                                           `json:"user_is_ignored,omitempty"`
	UserIsValid                 *bool                                                           `json:"user_is_valid,omitempty"`
	UserIsVerified              *bool                                                           `json:"user_is_verified,omitempty"`
	UserIsVisitor               *bool                                                           `json:"user_is_visitor,omitempty"`
	UserLastSeenDate            *int                                                            `json:"user_last_seen_date,omitempty"`
	UserLike2Count              *int                                                            `json:"user_like2_count,omitempty"`
	UserLikeCount               *int                                                            `json:"user_like_count,omitempty"`
	UserMessageCount            *int                                                            `json:"user_message_count,omitempty"`
	UserRegisterDate            *int                                                            `json:"user_register_date,omitempty"`
	UserTimezoneOffset          *int                                                            `json:"user_timezone_offset,omitempty"`
	UserTitle                   *string                                                         `json:"user_title,omitempty"`
	UserUnreadConversationCount *int                                                            `json:"user_unread_conversation_count,omitempty"`
	UserUnreadNotificationCount *int                                                            `json:"user_unread_notification_count,omitempty"`
	Username                    *string                                                         `json:"username,omitempty"`
	UsernameHTML                *string                                                         `json:"username_html,omitempty"`
}

type ClaimsResponseClaimsItemAuthorBirthday struct {
	Age       *int                                             `json:"age,omitempty"`
	Format    *string                                          `json:"format,omitempty"`
	TimeStamp *ClaimsResponseClaimsItemAuthorBirthdayTimeStamp `json:"timeStamp,omitempty"`
}

type ClaimsResponseClaimsItemAuthorBirthdayTimeStamp struct {
	Date         *string `json:"date,omitempty"`
	Timezone     *string `json:"timezone,omitempty"`
	TimezoneType *int    `json:"timezone_type,omitempty"`
}

type ClaimsResponseClaimsItemAuthorEditPermissions struct {
	Fields            *bool `json:"fields,omitempty"`
	HideUsernameLogs  *bool `json:"hide_username_logs,omitempty"`
	Password          *bool `json:"password,omitempty"`
	PrimaryGroupID    *bool `json:"primary_group_id,omitempty"`
	SecondaryGroupIds *bool `json:"secondary_group_ids,omitempty"`
	ShortLink         *bool `json:"short_link,omitempty"`
	UserDobDay        *bool `json:"user_dob_day,omitempty"`
	UserDobMonth      *bool `json:"user_dob_month,omitempty"`
	UserDobYear       *bool `json:"user_dob_year,omitempty"`
	UserEmail         *bool `json:"user_email,omitempty"`
	UserTitle         *bool `json:"user_title,omitempty"`
	Username          *bool `json:"username,omitempty"`
}

type ClaimsResponseClaimsItemAuthorFieldsItem struct {
	Choices       []ClaimsResponseClaimsItemAuthorFieldsItemChoicesItem `json:"choices,omitempty"`
	Description   *string                                               `json:"description,omitempty"`
	ID            *string                                               `json:"id,omitempty"`
	IsMultiChoice *bool                                                 `json:"is_multi_choice,omitempty"`
	IsRequired    *bool                                                 `json:"is_required,omitempty"`
	Position      *string                                               `json:"position,omitempty"`
	Title         *string                                               `json:"title,omitempty"`
	Value         *string                                               `json:"value,omitempty"`
	Values        interface{}                                           `json:"values,omitempty"`
}

type ClaimsResponseClaimsItemAuthorFieldsItemChoicesItem struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type ClaimsResponseClaimsItemAuthorLinks struct {
	Avatar      *string `json:"avatar,omitempty"`
	AvatarBig   *string `json:"avatar_big,omitempty"`
	AvatarSmall *string `json:"avatar_small,omitempty"`
	BackgroundL *string `json:"background_l,omitempty"`
	BackgroundM *string `json:"background_m,omitempty"`
	Detail      *string `json:"detail,omitempty"`
	Followers   *string `json:"followers,omitempty"`
	Followings  *string `json:"followings,omitempty"`
	Ignore      *string `json:"ignore,omitempty"`
	Permalink   *string `json:"permalink,omitempty"`
	Status      *string `json:"status,omitempty"`
	Timeline    *string `json:"timeline,omitempty"`
}

type ClaimsResponseClaimsItemAuthorPermissions struct {
	Edit        *bool `json:"edit,omitempty"`
	Follow      *bool `json:"follow,omitempty"`
	Ignore      *bool `json:"ignore,omitempty"`
	ProfilePost *bool `json:"profile_post,omitempty"`
}

type ClaimsResponseClaimsItemAuthorSelfPermissions struct {
	CreateConversation *bool `json:"create_conversation,omitempty"`
}

type ClaimsResponseClaimsItemAuthorUserExternalAuthenticationsItem struct {
	Provider    *string `json:"provider,omitempty"`
	ProviderKey *string `json:"provider_key,omitempty"`
}

type ClaimsResponseClaimsItemAuthorUserFollowers struct {
	Count *int                                                   `json:"count,omitempty"`
	Users []ClaimsResponseClaimsItemAuthorUserFollowersUsersItem `json:"users,omitempty"`
}

type ClaimsResponseClaimsItemAuthorUserFollowersUsersItem struct {
	Avatar       *string `json:"avatar,omitempty"`
	UserID       *int    `json:"user_id,omitempty"`
	Username     *string `json:"username,omitempty"`
	UsernameHTML *string `json:"username_html,omitempty"`
}

type ClaimsResponseClaimsItemAuthorUserFollowing struct {
	Count *int                                                   `json:"count,omitempty"`
	Users []ClaimsResponseClaimsItemAuthorUserFollowingUsersItem `json:"users,omitempty"`
}

type ClaimsResponseClaimsItemAuthorUserFollowingUsersItem struct {
	Avatar       *string `json:"avatar,omitempty"`
	UserID       *int    `json:"user_id,omitempty"`
	Username     *string `json:"username,omitempty"`
	UsernameHTML *string `json:"username_html,omitempty"`
}

type ClaimsResponseClaimsItemAuthorUserGroupsItem struct {
	DisplayBannerSelectable *bool   `json:"display_banner_selectable,omitempty"`
	DisplayGroupSelectable  *bool   `json:"display_group_selectable,omitempty"`
	DisplayIconSelectable   *bool   `json:"display_icon_selectable,omitempty"`
	IsPrimaryGroup          *bool   `json:"is_primary_group,omitempty"`
	UserGroupBannerCSSClass *string `json:"user_group_banner_css_class,omitempty"`
	UserGroupBannerText     *string `json:"user_group_banner_text,omitempty"`
	UserGroupBannerTextEn   *string `json:"user_group_banner_text_en,omitempty"`
	UserGroupIconClass      *string `json:"user_group_icon_class,omitempty"`
	UserGroupID             *int    `json:"user_group_id,omitempty"`
	UserGroupTitle          *string `json:"user_group_title,omitempty"`
	UserGroupTitleEn        *string `json:"user_group_title_en,omitempty"`
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
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ContentsParams struct {
	// Page number of contents.
	Page *int `json:"page,omitempty"`
	// Number of contents in a page.
	Limit *int `json:"limit,omitempty"`
}

type ContentsResponse struct {
	Data       []ContentsResponseDataItem `json:"data,omitempty"`
	DataTotal  *int                       `json:"data_total,omitempty"`
	Links      *ContentsResponseLinks     `json:"links,omitempty"`
	SystemInfo *RespSystemInfo            `json:"system_info,omitempty"`
	User       *RespUserModel             `json:"user,omitempty"`
}

type ContentsResponseDataItem struct {
	ContentID           interface{}                             `json:"content_id,omitempty"`
	ContentType         *string                                 `json:"content_type,omitempty"`
	LikeUsers           []ContentsResponseDataItemLikeUsersItem `json:"like_users,omitempty"`
	Links               *ContentsResponseDataItemLinks          `json:"links,omitempty"`
	Permissions         *ContentsResponseDataItemPermissions    `json:"permissions,omitempty"`
	PostAttachmentCount *int                                    `json:"post_attachment_count,omitempty"`
	PostBody            *string                                 `json:"post_body,omitempty"`
	PostBodyHTML        *string                                 `json:"post_body_html,omitempty"`
	PostBodyPlainText   *string                                 `json:"post_body_plain_text,omitempty"`
	PostCreateDate      *int                                    `json:"post_create_date,omitempty"`
	PostID              *int                                    `json:"post_id,omitempty"`
	PostIsDeleted       *bool                                   `json:"post_is_deleted,omitempty"`
	PostIsFirstPost     *bool                                   `json:"post_is_first_post,omitempty"`
	PostIsPublished     *bool                                   `json:"post_is_published,omitempty"`
	PostLikeCount       *int                                    `json:"post_like_count,omitempty"`
	PostUpdateDate      *int                                    `json:"post_update_date,omitempty"`
	PosterUserID        *int                                    `json:"poster_user_id,omitempty"`
	PosterUsername      *string                                 `json:"poster_username,omitempty"`
	PosterUsernameHTML  *string                                 `json:"poster_username_html,omitempty"`
	Signature           *string                                 `json:"signature,omitempty"`
	SignatureHTML       *string                                 `json:"signature_html,omitempty"`
	SignaturePlainText  *string                                 `json:"signature_plain_text,omitempty"`
	Thread              *ContentsResponseDataItemThread         `json:"thread,omitempty"`
	ThreadID            *int                                    `json:"thread_id,omitempty"`
	UserIsIgnored       *bool                                   `json:"user_is_ignored,omitempty"`
}

type ContentsResponseDataItemLikeUsersItem struct {
	DisplayStyleGroupID *int    `json:"display_style_group_id,omitempty"`
	IsBanned            *int    `json:"is_banned,omitempty"`
	UniqUsernameCSS     *string `json:"uniq_username_css,omitempty"`
	UserID              *int    `json:"user_id,omitempty"`
	Username            *string `json:"username,omitempty"`
}

type ContentsResponseDataItemLinks struct {
	Attachments  *string `json:"attachments,omitempty"`
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type ContentsResponseDataItemPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Like             *bool `json:"like,omitempty"`
	Reply            *bool `json:"reply,omitempty"`
	Report           *bool `json:"report,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type ContentsResponseDataItemThread struct {
	CreatorUserID       *int                                       `json:"creator_user_id,omitempty"`
	CreatorUsername     *string                                    `json:"creator_username,omitempty"`
	CreatorUsernameHTML *string                                    `json:"creator_username_html,omitempty"`
	ForumID             *int                                       `json:"forum_id,omitempty"`
	Links               *ContentsResponseDataItemThreadLinks       `json:"links,omitempty"`
	Permissions         *ContentsResponseDataItemThreadPermissions `json:"permissions,omitempty"`
	ThreadCreateDate    *int                                       `json:"thread_create_date,omitempty"`
	ThreadID            *int                                       `json:"thread_id,omitempty"`
	ThreadIsDeleted     *bool                                      `json:"thread_is_deleted,omitempty"`
	ThreadIsFollowed    *bool                                      `json:"thread_is_followed,omitempty"`
	ThreadIsPublished   *bool                                      `json:"thread_is_published,omitempty"`
	ThreadIsSticky      *bool                                      `json:"thread_is_sticky,omitempty"`
	ThreadPostCount     *int                                       `json:"thread_post_count,omitempty"`
	ThreadPrefixes      []interface{}                              `json:"thread_prefixes,omitempty"`
	ThreadTags          interface{}                                `json:"thread_tags,omitempty"`
	ThreadTitle         *string                                    `json:"thread_title,omitempty"`
	ThreadUpdateDate    *int                                       `json:"thread_update_date,omitempty"`
	ThreadViewCount     *int                                       `json:"thread_view_count,omitempty"`
	UserIsIgnored       *bool                                      `json:"user_is_ignored,omitempty"`
}

type ContentsResponseDataItemThreadLinks struct {
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

type ContentsResponseDataItemThreadPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	Post             *bool `json:"post,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type ContentsResponseLinks struct {
	Next  *string `json:"next,omitempty"`
	Page  *int    `json:"page,omitempty"`
	Pages *int    `json:"pages,omitempty"`
}

type ContentsResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ContentsResponseUser struct {
	Balance                     *string                                               `json:"balance,omitempty"`
	Banner                      *string                                               `json:"banner,omitempty"`
	Birthday                    *ContentsResponseUserBirthday                         `json:"birthday,omitempty"`
	ContestCount                *int                                                  `json:"contest_count,omitempty"`
	ConvWelcomeMessage          *string                                               `json:"conv_welcome_message,omitempty"`
	CuratorTitles               []string                                              `json:"curator_titles,omitempty"`
	Currency                    *string                                               `json:"currency,omitempty"`
	CustomTitle                 *string                                               `json:"custom_title,omitempty"`
	DisplayBannerID             *int                                                  `json:"display_banner_id,omitempty"`
	DisplayIconGroupID          *int                                                  `json:"display_icon_group_id,omitempty"`
	EditPermissions             *ContentsResponseUserEditPermissions                  `json:"edit_permissions,omitempty"`
	Fields                      []ContentsResponseUserFieldsItem                      `json:"fields,omitempty"`
	Hold                        *string                                               `json:"hold,omitempty"`
	IsBanned                    *int                                                  `json:"is_banned,omitempty"`
	Links                       *ContentsResponseUserLinks                            `json:"links,omitempty"`
	Permissions                 *ContentsResponseUserPermissions                      `json:"permissions,omitempty"`
	SecretAnswerFirstLetter     *string                                               `json:"secret_answer_first_letter,omitempty"`
	SecretAnswerRendered        *string                                               `json:"secret_answer_rendered,omitempty"`
	SelfPermissions             *ContentsResponseUserSelfPermissions                  `json:"self_permissions,omitempty"`
	ShortLink                   *string                                               `json:"short_link,omitempty"`
	TrophyCount                 *int                                                  `json:"trophy_count,omitempty"`
	UserDeposit                 *int                                                  `json:"user_deposit,omitempty"`
	UserEmail                   *string                                               `json:"user_email,omitempty"`
	UserExternalAuthentications []ContentsResponseUserUserExternalAuthenticationsItem `json:"user_external_authentications,omitempty"`
	UserFollowers               *ContentsResponseUserUserFollowers                    `json:"user_followers,omitempty"`
	UserFollowing               *ContentsResponseUserUserFollowing                    `json:"user_following,omitempty"`
	UserGroupID                 *int                                                  `json:"user_group_id,omitempty"`
	UserGroups                  []ContentsResponseUserUserGroupsItem                  `json:"user_groups,omitempty"`
	UserID                      *int                                                  `json:"user_id,omitempty"`
	UserIsFollowed              *bool                                                 `json:"user_is_followed,omitempty"`
	UserIsIgnored               *bool                                                 `json:"user_is_ignored,omitempty"`
	UserIsValid                 *bool                                                 `json:"user_is_valid,omitempty"`
	UserIsVerified              *bool                                                 `json:"user_is_verified,omitempty"`
	UserIsVisitor               *bool                                                 `json:"user_is_visitor,omitempty"`
	UserLastSeenDate            *int                                                  `json:"user_last_seen_date,omitempty"`
	UserLike2Count              *int                                                  `json:"user_like2_count,omitempty"`
	UserLikeCount               *int                                                  `json:"user_like_count,omitempty"`
	UserMessageCount            *int                                                  `json:"user_message_count,omitempty"`
	UserRegisterDate            *int                                                  `json:"user_register_date,omitempty"`
	UserTimezoneOffset          *int                                                  `json:"user_timezone_offset,omitempty"`
	UserTitle                   *string                                               `json:"user_title,omitempty"`
	UserUnreadConversationCount *int                                                  `json:"user_unread_conversation_count,omitempty"`
	UserUnreadNotificationCount *int                                                  `json:"user_unread_notification_count,omitempty"`
	Username                    *string                                               `json:"username,omitempty"`
	UsernameHTML                *string                                               `json:"username_html,omitempty"`
}

type ContentsResponseUserBirthday struct {
	Age       *int                                   `json:"age,omitempty"`
	Format    *string                                `json:"format,omitempty"`
	TimeStamp *ContentsResponseUserBirthdayTimeStamp `json:"timeStamp,omitempty"`
}

type ContentsResponseUserBirthdayTimeStamp struct {
	Date         *string `json:"date,omitempty"`
	Timezone     *string `json:"timezone,omitempty"`
	TimezoneType *int    `json:"timezone_type,omitempty"`
}

type ContentsResponseUserEditPermissions struct {
	Fields            *bool `json:"fields,omitempty"`
	HideUsernameLogs  *bool `json:"hide_username_logs,omitempty"`
	Password          *bool `json:"password,omitempty"`
	PrimaryGroupID    *bool `json:"primary_group_id,omitempty"`
	SecondaryGroupIds *bool `json:"secondary_group_ids,omitempty"`
	ShortLink         *bool `json:"short_link,omitempty"`
	UserDobDay        *bool `json:"user_dob_day,omitempty"`
	UserDobMonth      *bool `json:"user_dob_month,omitempty"`
	UserDobYear       *bool `json:"user_dob_year,omitempty"`
	UserEmail         *bool `json:"user_email,omitempty"`
	UserTitle         *bool `json:"user_title,omitempty"`
	Username          *bool `json:"username,omitempty"`
}

type ContentsResponseUserFieldsItem struct {
	Choices       []ContentsResponseUserFieldsItemChoicesItem `json:"choices,omitempty"`
	Description   *string                                     `json:"description,omitempty"`
	ID            *string                                     `json:"id,omitempty"`
	IsMultiChoice *bool                                       `json:"is_multi_choice,omitempty"`
	IsRequired    *bool                                       `json:"is_required,omitempty"`
	Position      *string                                     `json:"position,omitempty"`
	Title         *string                                     `json:"title,omitempty"`
	Value         *string                                     `json:"value,omitempty"`
	Values        interface{}                                 `json:"values,omitempty"`
}

type ContentsResponseUserFieldsItemChoicesItem struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type ContentsResponseUserLinks struct {
	Avatar      *string `json:"avatar,omitempty"`
	AvatarBig   *string `json:"avatar_big,omitempty"`
	AvatarSmall *string `json:"avatar_small,omitempty"`
	BackgroundL *string `json:"background_l,omitempty"`
	BackgroundM *string `json:"background_m,omitempty"`
	Detail      *string `json:"detail,omitempty"`
	Followers   *string `json:"followers,omitempty"`
	Followings  *string `json:"followings,omitempty"`
	Ignore      *string `json:"ignore,omitempty"`
	Permalink   *string `json:"permalink,omitempty"`
	Status      *string `json:"status,omitempty"`
	Timeline    *string `json:"timeline,omitempty"`
}

type ContentsResponseUserPermissions struct {
	Edit        *bool `json:"edit,omitempty"`
	Follow      *bool `json:"follow,omitempty"`
	Ignore      *bool `json:"ignore,omitempty"`
	ProfilePost *bool `json:"profile_post,omitempty"`
}

type ContentsResponseUserSelfPermissions struct {
	CreateConversation *bool `json:"create_conversation,omitempty"`
}

type ContentsResponseUserUserExternalAuthenticationsItem struct {
	Provider    *string `json:"provider,omitempty"`
	ProviderKey *string `json:"provider_key,omitempty"`
}

type ContentsResponseUserUserFollowers struct {
	Count *int                                         `json:"count,omitempty"`
	Users []ContentsResponseUserUserFollowersUsersItem `json:"users,omitempty"`
}

type ContentsResponseUserUserFollowersUsersItem struct {
	Avatar       *string `json:"avatar,omitempty"`
	UserID       *int    `json:"user_id,omitempty"`
	Username     *string `json:"username,omitempty"`
	UsernameHTML *string `json:"username_html,omitempty"`
}

type ContentsResponseUserUserFollowing struct {
	Count *int                                         `json:"count,omitempty"`
	Users []ContentsResponseUserUserFollowingUsersItem `json:"users,omitempty"`
}

type ContentsResponseUserUserFollowingUsersItem struct {
	Avatar       *string `json:"avatar,omitempty"`
	UserID       *int    `json:"user_id,omitempty"`
	Username     *string `json:"username,omitempty"`
	UsernameHTML *string `json:"username_html,omitempty"`
}

type ContentsResponseUserUserGroupsItem struct {
	DisplayBannerSelectable *bool   `json:"display_banner_selectable,omitempty"`
	DisplayGroupSelectable  *bool   `json:"display_group_selectable,omitempty"`
	DisplayIconSelectable   *bool   `json:"display_icon_selectable,omitempty"`
	IsPrimaryGroup          *bool   `json:"is_primary_group,omitempty"`
	UserGroupBannerCSSClass *string `json:"user_group_banner_css_class,omitempty"`
	UserGroupBannerText     *string `json:"user_group_banner_text,omitempty"`
	UserGroupBannerTextEn   *string `json:"user_group_banner_text_en,omitempty"`
	UserGroupIconClass      *string `json:"user_group_icon_class,omitempty"`
	UserGroupID             *int    `json:"user_group_id,omitempty"`
	UserGroupTitle          *string `json:"user_group_title,omitempty"`
	UserGroupTitleEn        *string `json:"user_group_title_en,omitempty"`
}

type CreateContestParams struct {
	// Allow ask hidden content.
	AllowAskHiddenContent *bool `json:"allow_ask_hidden_content,omitempty"`
	// Allow commenting if user can't post in thread.
	CommentIgnoreGroup *bool `json:"comment_ignore_group,omitempty"`
	// Winner count (prize count). Optional if **prize_type** is **money**.
	CountWinners *int `json:"count_winners,omitempty"`
	// Don't alert followers about thread creation.
	DontAlertFollowers *bool `json:"dont_alert_followers,omitempty"`
	// Hide contacts.
	HideContacts *bool `json:"hide_contacts,omitempty"`
	// Enable the distribution of money prizes by places. Optional if **prize_type** is **money**.
	IsMoneyPlaces *bool `json:"is_money_places,omitempty"`
	// Giveaway duration type. The maximum duration is 3 days. Required if **contest_type** is **by_finish_date**.
	LengthOption *ThreadsCreateContestLengthOption `json:"length_option,omitempty"`
	// Giveaway duration value. The maximum duration is 3 days. Required if **contest_type** is **by_finish_date**.
	LengthValue *int `json:"length_value,omitempty"`
	// How much money will each winner receive. Optional if **prize_type** is **money**.
	PrizeDataMoney *float64 `json:"prize_data_money,omitempty"`
	// How much money will receive each place. Required if **is_money_places** is **1**.
	PrizeDataPlaces []float64 `json:"prize_data_places,omitempty"`
	// Which upgrade will each winner receive. Required if **prize_type** is **upgrades**.
	PrizeDataUpgrade *ThreadsCreateContestPrizeDataUpgrade `json:"prize_data_upgrade,omitempty"`
	// Allow to reply only users with chosen or higher group.
	ReplyGroup *ThreadsCreateContestReplyGroup `json:"reply_group,omitempty"`
	// Date to schedule thread creation (format: `DD-MM-YYYY`).
	ScheduleDate *string `json:"schedule_date,omitempty"`
	// Time to schedule thread creation (format: `HH:MM`).
	ScheduleTime *string `json:"schedule_time,omitempty"`
	// Secret answer of your account.
	SecretAnswer *string `json:"secret_answer,omitempty"`
	// Thread tags.
	Tags []string `json:"tags,omitempty"`
	// Thread title. Can be skipped if **title_en** set.
	Title *string `json:"title,omitempty"`
	// Thread english title. Can be skipped if **title** set.
	TitleEn *string `json:"title_en,omitempty"`
	// Receive forum notifications of new posts in this thread.
	WatchThread *bool `json:"watch_thread,omitempty"`
	// Receive email notifications of new posts in this thread.
	WatchThreadEmail *bool `json:"watch_thread_email,omitempty"`
	// Watch thread state.
	WatchThreadState *bool `json:"watch_thread_state,omitempty"`
}

type CreateContestResponse struct {
	SystemInfo *RespSystemInfo  `json:"system_info,omitempty"`
	Thread     *RespThreadModel `json:"thread,omitempty"`
}

type CreateContestResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type CreateContestResponseThread struct {
	Contest             *CreateContestResponseThreadContest      `json:"contest,omitempty"`
	CreatorUserID       *int                                     `json:"creator_user_id,omitempty"`
	CreatorUsername     *string                                  `json:"creator_username,omitempty"`
	CreatorUsernameHTML *string                                  `json:"creator_username_html,omitempty"`
	FirstPost           *CreateContestResponseThreadFirstPost    `json:"first_post,omitempty"`
	ForumID             *int                                     `json:"forum_id,omitempty"`
	LastPost            *CreateContestResponseThreadLastPost     `json:"last_post,omitempty"`
	Links               *CreateContestResponseThreadLinks        `json:"links,omitempty"`
	NodeTitle           *string                                  `json:"node_title,omitempty"`
	Permissions         *CreateContestResponseThreadPermissions  `json:"permissions,omitempty"`
	Restrictions        *CreateContestResponseThreadRestrictions `json:"restrictions,omitempty"`
	ThreadCreateDate    *int                                     `json:"thread_create_date,omitempty"`
	ThreadID            *int                                     `json:"thread_id,omitempty"`
	ThreadIsClosed      *bool                                    `json:"thread_is_closed,omitempty"`
	ThreadIsDeleted     *bool                                    `json:"thread_is_deleted,omitempty"`
	ThreadIsFollowed    *bool                                    `json:"thread_is_followed,omitempty"`
	ThreadIsPublished   *bool                                    `json:"thread_is_published,omitempty"`
	ThreadIsStarred     *bool                                    `json:"thread_is_starred,omitempty"`
	ThreadIsSticky      *bool                                    `json:"thread_is_sticky,omitempty"`
	ThreadPostCount     *int                                     `json:"thread_post_count,omitempty"`
	ThreadPrefixes      []interface{}                            `json:"thread_prefixes,omitempty"`
	ThreadTags          interface{}                              `json:"thread_tags,omitempty"`
	ThreadTitle         *string                                  `json:"thread_title,omitempty"`
	ThreadUpdateDate    *int                                     `json:"thread_update_date,omitempty"`
	ThreadViewCount     *int                                     `json:"thread_view_count,omitempty"`
	UserIsIgnored       *bool                                    `json:"user_is_ignored,omitempty"`
}

type CreateContestResponseThreadContest struct {
	AlreadyParticipate    *bool                                          `json:"already_participate,omitempty"`
	ChanceToWin           *float64                                       `json:"chance_to_win,omitempty"`
	CountWinners          *int                                           `json:"count_winners,omitempty"`
	FinishDate            *int                                           `json:"finish_date,omitempty"`
	IsFinished            *int                                           `json:"is_finished,omitempty"`
	IsMoneyPlaces         *int                                           `json:"is_money_places,omitempty"`
	NeededMembers         *int                                           `json:"needed_members,omitempty"`
	NowCountMembers       *int                                           `json:"now_count_members,omitempty"`
	Permissions           *CreateContestResponseThreadContestPermissions `json:"permissions,omitempty"`
	PrizeData             *int                                           `json:"prize_data,omitempty"`
	PrizeType             *string                                        `json:"prize_type,omitempty"`
	PrizeTypePhrase       *string                                        `json:"prize_type_phrase,omitempty"`
	RequireLikeCount      *int                                           `json:"require_like_count,omitempty"`
	RequireTotalLikeCount *int                                           `json:"require_total_like_count,omitempty"`
	Type_                 *string                                        `json:"type,omitempty"`
	Winners               []int                                          `json:"winners,omitempty"`
}

type CreateContestResponseThreadContestPermissions struct {
	CanFinish           *bool   `json:"can_finish,omitempty"`
	CanParticipate      *bool   `json:"can_participate,omitempty"`
	CanParticipateError *string `json:"can_participate_error,omitempty"`
	CanViewUserList     *bool   `json:"can_view_user_list,omitempty"`
}

type CreateContestResponseThreadFirstPost struct {
	Links              *CreateContestResponseThreadFirstPostLinks       `json:"links,omitempty"`
	Permissions        *CreateContestResponseThreadFirstPostPermissions `json:"permissions,omitempty"`
	PostBody           *string                                          `json:"post_body,omitempty"`
	PostBodyHTML       *string                                          `json:"post_body_html,omitempty"`
	PostBodyPlainText  *string                                          `json:"post_body_plain_text,omitempty"`
	PostCreateDate     *int                                             `json:"post_create_date,omitempty"`
	PostID             *int                                             `json:"post_id,omitempty"`
	PostIsDeleted      *bool                                            `json:"post_is_deleted,omitempty"`
	PostIsFirstPost    *bool                                            `json:"post_is_first_post,omitempty"`
	PostIsLiked        *bool                                            `json:"post_is_liked,omitempty"`
	PostIsPublished    *bool                                            `json:"post_is_published,omitempty"`
	PostLikeCount      *int                                             `json:"post_like_count,omitempty"`
	PostUpdateDate     *int                                             `json:"post_update_date,omitempty"`
	PosterUserID       *int                                             `json:"poster_user_id,omitempty"`
	PosterUsername     *string                                          `json:"poster_username,omitempty"`
	PosterUsernameHTML *string                                          `json:"poster_username_html,omitempty"`
	Signature          *string                                          `json:"signature,omitempty"`
	SignatureHTML      *string                                          `json:"signature_html,omitempty"`
	SignaturePlainText *string                                          `json:"signature_plain_text,omitempty"`
	ThreadID           *int                                             `json:"thread_id,omitempty"`
	ThreadIsDeleted    *bool                                            `json:"thread_is_deleted,omitempty"`
	UserIsIgnored      *bool                                            `json:"user_is_ignored,omitempty"`
}

type CreateContestResponseThreadFirstPostLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type CreateContestResponseThreadFirstPostPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	Like   *bool `json:"like,omitempty"`
	Reply  *bool `json:"reply,omitempty"`
	Report *bool `json:"report,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type CreateContestResponseThreadLastPost struct {
	Links              *CreateContestResponseThreadLastPostLinks       `json:"links,omitempty"`
	Permissions        *CreateContestResponseThreadLastPostPermissions `json:"permissions,omitempty"`
	PostBody           *string                                         `json:"post_body,omitempty"`
	PostBodyHTML       *string                                         `json:"post_body_html,omitempty"`
	PostBodyPlainText  *string                                         `json:"post_body_plain_text,omitempty"`
	PostCreateDate     *int                                            `json:"post_create_date,omitempty"`
	PostID             *int                                            `json:"post_id,omitempty"`
	PostIsDeleted      *bool                                           `json:"post_is_deleted,omitempty"`
	PostIsFirstPost    *bool                                           `json:"post_is_first_post,omitempty"`
	PostIsLiked        *bool                                           `json:"post_is_liked,omitempty"`
	PostIsPublished    *bool                                           `json:"post_is_published,omitempty"`
	PostLikeCount      *int                                            `json:"post_like_count,omitempty"`
	PostUpdateDate     *int                                            `json:"post_update_date,omitempty"`
	PosterUserID       *int                                            `json:"poster_user_id,omitempty"`
	PosterUsername     *string                                         `json:"poster_username,omitempty"`
	PosterUsernameHTML *string                                         `json:"poster_username_html,omitempty"`
	Signature          *string                                         `json:"signature,omitempty"`
	SignatureHTML      *string                                         `json:"signature_html,omitempty"`
	SignaturePlainText *string                                         `json:"signature_plain_text,omitempty"`
	ThreadID           *int                                            `json:"thread_id,omitempty"`
	ThreadIsDeleted    *bool                                           `json:"thread_is_deleted,omitempty"`
	UserIsIgnored      *bool                                           `json:"user_is_ignored,omitempty"`
}

type CreateContestResponseThreadLastPostLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type CreateContestResponseThreadLastPostPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	Like   *bool `json:"like,omitempty"`
	Reply  *bool `json:"reply,omitempty"`
	Report *bool `json:"report,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type CreateContestResponseThreadLinks struct {
	Detail            *string `json:"detail,omitempty"`
	FirstPost         *string `json:"first_post,omitempty"`
	FirstPoster       *string `json:"first_poster,omitempty"`
	FirstPosterAvatar *string `json:"first_poster_avatar,omitempty"`
	Followers         *string `json:"followers,omitempty"`
	Forum             *string `json:"forum,omitempty"`
	LastPost          *string `json:"last_post,omitempty"`
	Permalink         *string `json:"permalink,omitempty"`
	Posts             *string `json:"posts,omitempty"`
}

type CreateContestResponseThreadPermissions struct {
	Bump      *CreateContestResponseThreadPermissionsBump `json:"bump,omitempty"`
	Delete    *bool                                       `json:"delete,omitempty"`
	Edit      *bool                                       `json:"edit,omitempty"`
	EditTags  *bool                                       `json:"edit_tags,omitempty"`
	EditTitle *bool                                       `json:"edit_title,omitempty"`
	Follow    *bool                                       `json:"follow,omitempty"`
	Post      *bool                                       `json:"post,omitempty"`
	View      *bool                                       `json:"view,omitempty"`
}

type CreateContestResponseThreadPermissionsBump struct {
	AvailableCount    *int        `json:"available_count,omitempty"`
	Can               *bool       `json:"can,omitempty"`
	Error             interface{} `json:"error,omitempty"`
	NextAvailableTime interface{} `json:"next_available_time,omitempty"`
}

type CreateContestResponseThreadRestrictions struct {
	MaxReplyCount *int `json:"max_reply_count,omitempty"`
	ReplyDelay    *int `json:"reply_delay,omitempty"`
}

type CreateParams struct {
	// Allow ask hidden content.
	AllowAskHiddenContent *bool `json:"allow_ask_hidden_content,omitempty"`
	// Allow commenting if user can't post in thread.
	CommentIgnoreGroup *bool `json:"comment_ignore_group,omitempty"`
	// Don't alert followers about thread creation.
	DontAlertFollowers *bool `json:"dont_alert_followers,omitempty"`
	// Hide contacts.
	HideContacts *bool `json:"hide_contacts,omitempty"`
	// Prefix ids.
	PrefixID []int `json:"prefix_id,omitempty"`
	// Allow to reply only users with chosen or higher group.
	ReplyGroup *ThreadsCreateReplyGroup `json:"reply_group,omitempty"`
	// Date to schedule thread creation (format: `DD-MM-YYYY`).
	ScheduleDate *string `json:"schedule_date,omitempty"`
	// Time to schedule thread creation (format: `HH:MM`).
	ScheduleTime *string `json:"schedule_time,omitempty"`
	// Thread tags.
	Tags []string `json:"tags,omitempty"`
	// Thread title. Can be skipped if **title_en** set.
	Title *string `json:"title,omitempty"`
	// Thread english title. Can be skipped if **title** set.
	TitleEn *string `json:"title_en,omitempty"`
	// Receive forum notifications of new posts in this thread.
	WatchThread *bool `json:"watch_thread,omitempty"`
	// Receive email notifications of new posts in this thread.
	WatchThreadEmail *bool `json:"watch_thread_email,omitempty"`
	// Watch thread state.
	WatchThreadState *bool `json:"watch_thread_state,omitempty"`
}

type CreatePostParams struct {
	// ID of the message being replied to.
	ReplyMessageID *int `json:"reply_message_id,omitempty"`
}

type CreatePostResponse struct {
	Message    *RespConversationMessageModel `json:"message,omitempty"`
	SystemInfo *RespSystemInfo               `json:"system_info,omitempty"`
}

type CreatePostResponseMessage struct {
	ConversationID       *int                                  `json:"conversation_id,omitempty"`
	CreatorUserID        *int                                  `json:"creator_user_id,omitempty"`
	CreatorUsername      *string                               `json:"creator_username,omitempty"`
	CreatorUsernameHTML  *string                               `json:"creator_username_html,omitempty"`
	Links                *CreatePostResponseMessageLinks       `json:"links,omitempty"`
	MessageBody          *string                               `json:"message_body,omitempty"`
	MessageBodyHTML      *string                               `json:"message_body_html,omitempty"`
	MessageBodyPlainText *string                               `json:"message_body_plain_text,omitempty"`
	MessageCreateDate    *int                                  `json:"message_create_date,omitempty"`
	MessageEditDate      *int                                  `json:"message_edit_date,omitempty"`
	MessageID            *int                                  `json:"message_id,omitempty"`
	MessageIsSystem      *bool                                 `json:"message_is_system,omitempty"`
	MessageIsUnread      *int                                  `json:"message_is_unread,omitempty"`
	MessageNeedTranslate *bool                                 `json:"message_need_translate,omitempty"`
	Permissions          *CreatePostResponseMessagePermissions `json:"permissions,omitempty"`
	UserIsIgnored        *bool                                 `json:"user_is_ignored,omitempty"`
}

type CreatePostResponseMessageLinks struct {
	Conversation  *string `json:"conversation,omitempty"`
	Creator       *string `json:"creator,omitempty"`
	CreatorAvatar *string `json:"creator_avatar,omitempty"`
	Detail        *string `json:"detail,omitempty"`
}

type CreatePostResponseMessagePermissions struct {
	Delete       *bool `json:"delete,omitempty"`
	Edit         *bool `json:"edit,omitempty"`
	StickUnstick *bool `json:"stick-unstick,omitempty"`
	View         *bool `json:"view,omitempty"`
}

type CreatePostResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type CreateResponse struct {
	Conversation *RespConversationModel `json:"conversation,omitempty"`
	SystemInfo   *RespSystemInfo        `json:"system_info,omitempty"`
}

type CreateResponseConversation struct {
	Alerts                   *int                                       `json:"alerts,omitempty"`
	ConversationCreateDate   *int                                       `json:"conversation_create_date,omitempty"`
	ConversationID           *int                                       `json:"conversation_id,omitempty"`
	ConversationIsDeleted    *bool                                      `json:"conversation_is_deleted,omitempty"`
	ConversationIsNew        *bool                                      `json:"conversation_is_new,omitempty"`
	ConversationIsOpen       *bool                                      `json:"conversation_is_open,omitempty"`
	ConversationLastReadDate *int                                       `json:"conversation_last_read_date,omitempty"`
	ConversationMessageCount *int                                       `json:"conversation_message_count,omitempty"`
	ConversationOnlineCount  *int                                       `json:"conversation_online_count,omitempty"`
	ConversationTitle        *string                                    `json:"conversation_title,omitempty"`
	ConversationUpdateDate   *int                                       `json:"conversation_update_date,omitempty"`
	CreatorIsIgnored         *bool                                      `json:"creator_is_ignored,omitempty"`
	CreatorUserID            *int                                       `json:"creator_user_id,omitempty"`
	CreatorUsername          *string                                    `json:"creator_username,omitempty"`
	CreatorUsernameHTML      *string                                    `json:"creator_username_html,omitempty"`
	IsGroup                  *int                                       `json:"is_group,omitempty"`
	IsStarred                *int                                       `json:"is_starred,omitempty"`
	IsUnread                 *int                                       `json:"is_unread,omitempty"`
	Links                    *CreateResponseConversationLinks           `json:"links,omitempty"`
	Permissions              *CreateResponseConversationPermissions     `json:"permissions,omitempty"`
	Recipient                *CreateResponseConversationRecipient       `json:"recipient,omitempty"`
	Recipients               []CreateResponseConversationRecipientsItem `json:"recipients,omitempty"`
}

type CreateResponseConversationLinks struct {
	Avatar    *string `json:"avatar,omitempty"`
	Detail    *string `json:"detail,omitempty"`
	Messages  *string `json:"messages,omitempty"`
	Permalink *string `json:"permalink,omitempty"`
}

type CreateResponseConversationPermissions struct {
	EditOwnPost       *bool `json:"editOwnPost,omitempty"`
	Invite            *bool `json:"invite,omitempty"`
	Kick              *bool `json:"kick,omitempty"`
	ManageInviteLinks *bool `json:"manage_invite_links,omitempty"`
	Reply             *bool `json:"reply,omitempty"`
	StickyMessages    *bool `json:"stickyMessages,omitempty"`
	UploadAvatar      *bool `json:"upload_avatar,omitempty"`
	View              *bool `json:"view,omitempty"`
}

type CreateResponseConversationRecipient struct {
	Avatar          *string `json:"avatar,omitempty"`
	ContactsChanged *bool   `json:"contacts_changed,omitempty"`
	IsOnline        *bool   `json:"is_online,omitempty"`
	LastActivity    *int    `json:"last_activity,omitempty"`
	UserID          *int    `json:"user_id,omitempty"`
	Username        *string `json:"username,omitempty"`
	UsernameHTML    *string `json:"username_html,omitempty"`
}

type CreateResponseConversationRecipientsItem struct {
	Avatar          *string `json:"avatar,omitempty"`
	ContactsChanged *bool   `json:"contacts_changed,omitempty"`
	IsOnline        *bool   `json:"is_online,omitempty"`
	LastActivity    *int    `json:"last_activity,omitempty"`
	UserID          *int    `json:"user_id,omitempty"`
	Username        *string `json:"username,omitempty"`
	UsernameHTML    *string `json:"username_html,omitempty"`
}

type CreateResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type CropParams struct {
	// Selection size.
	Crop *int `json:"crop,omitempty"`
	// The starting point of the selection by width. Default value - 0
	X *int `json:"x,omitempty"`
	// The starting point of the selection by height. Default value - 0
	Y *int `json:"y,omitempty"`
}

type CropPostParams struct {
	// Selection size.
	Crop *int `json:"crop,omitempty"`
	// The starting point of the selection by width. Default value - 0
	X *int `json:"x,omitempty"`
	// The starting point of the selection by height. Default value - 0
	Y *int `json:"y,omitempty"`
}

type CropPostResponse struct {
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type CropPostResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type CropResponse struct {
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type CropResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type DeleteParams struct {
	// Reason of the thread removal.
	Reason *string `json:"reason,omitempty"`
}

type DisableResponse struct {
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type DisableResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type EditFeedOptionsParams struct {
	// List of keywords to exclude specific threads from the feed.
	Keywords []string `json:"keywords,omitempty"`
	// Array of forum ids to exclude from the feed.
	NodeIds []int `json:"node_ids,omitempty"`
}

type EditMessageResponse struct {
	Message    *RespChatboxMessageModel `json:"message,omitempty"`
	SystemInfo *RespSystemInfo          `json:"system_info,omitempty"`
}

type EditMessageResponseMessage struct {
	CanReport   *bool                           `json:"can_report,omitempty"`
	Date        *int                            `json:"date,omitempty"`
	IsDeleted   *bool                           `json:"is_deleted,omitempty"`
	Message     *string                         `json:"message,omitempty"`
	MessageJson *string                         `json:"messageJson,omitempty"`
	MessageRaw  *string                         `json:"messageRaw,omitempty"`
	MessageID   *int                            `json:"message_id,omitempty"`
	Room        *EditMessageResponseMessageRoom `json:"room,omitempty"`
	User        *EditMessageResponseMessageUser `json:"user,omitempty"`
}

type EditMessageResponseMessageRoom struct {
	CanReport *bool   `json:"can_report,omitempty"`
	Eng       *bool   `json:"eng,omitempty"`
	Market    *bool   `json:"market,omitempty"`
	RoomID    *int    `json:"room_id,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type EditMessageResponseMessageUser struct {
	AvatarDate          *int                                      `json:"avatar_date,omitempty"`
	BackgroundDate      *int                                      `json:"background_date,omitempty"`
	ContestCount        *int                                      `json:"contest_count,omitempty"`
	CustomTitle         *string                                   `json:"custom_title,omitempty"`
	DisplayBannerID     *int                                      `json:"display_banner_id,omitempty"`
	DisplayIconGroupID  *int                                      `json:"display_icon_group_id,omitempty"`
	DisplayStyleGroupID *int                                      `json:"display_style_group_id,omitempty"`
	IsAdmin             *bool                                     `json:"is_admin,omitempty"`
	IsBanned            *bool                                     `json:"is_banned,omitempty"`
	IsModerator         *bool                                     `json:"is_moderator,omitempty"`
	IsStaff             *bool                                     `json:"is_staff,omitempty"`
	LastActivity        *int                                      `json:"last_activity,omitempty"`
	Like2Count          *int                                      `json:"like2_count,omitempty"`
	LikeCount           *int                                      `json:"like_count,omitempty"`
	MessageCount        *int                                      `json:"message_count,omitempty"`
	RegisterDate        *int                                      `json:"register_date,omitempty"`
	Rendered            *EditMessageResponseMessageUserRendered   `json:"rendered,omitempty"`
	ShortLink           *string                                   `json:"short_link,omitempty"`
	TrophyPoints        *int                                      `json:"trophy_points,omitempty"`
	UniqBanner          *EditMessageResponseMessageUserUniqBanner `json:"uniq_banner,omitempty"`
	UniqUsernameCSS     *string                                   `json:"uniq_username_css,omitempty"`
	UserID              *int                                      `json:"user_id,omitempty"`
	Username            *string                                   `json:"username,omitempty"`
}

type EditMessageResponseMessageUserRendered struct {
	Avatars  *EditMessageResponseMessageUserRenderedAvatars `json:"avatars,omitempty"`
	Link     *string                                        `json:"link,omitempty"`
	Username *string                                        `json:"username,omitempty"`
}

type EditMessageResponseMessageUserRenderedAvatars struct {
	L *string `json:"l,omitempty"`
	M *string `json:"m,omitempty"`
	S *string `json:"s,omitempty"`
}

type EditMessageResponseMessageUserUniqBanner struct {
	BannerCSS    *string `json:"banner_css,omitempty"`
	BannerIcon   *string `json:"banner_icon,omitempty"`
	BannerText   *string `json:"banner_text,omitempty"`
	UsernameIcon *string `json:"username_icon,omitempty"`
}

type EditMessageResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type EditParams struct {
	// Whether user activity is visible.
	ActivityVisible *bool `json:"activity_visible,omitempty"`
	// Alert settings.
	Alert map[string]bool `json:"alert,omitempty"`
	// Who can invite you to groups.
	AllowInviteGroup *UsersEditAllowInviteGroup `json:"allow_invite_group,omitempty"`
	// Who can post on your profile.
	AllowPostProfile *UsersEditAllowPostProfile `json:"allow_post_profile,omitempty"`
	// Who can see your news feed.
	AllowReceiveNewsFeed *UsersEditAllowReceiveNewsFeed `json:"allow_receive_news_feed,omitempty"`
	// Who can send you personal conversations.
	AllowSendPersonalConversation *UsersEditAllowSendPersonalConversation `json:"allow_send_personal_conversation,omitempty"`
	// Who can view your profile.
	AllowViewProfile *UsersEditAllowViewProfile `json:"allow_view_profile,omitempty"`
	// This message is shown when someone wants to write to you.
	ConvWelcomeMessage *string `json:"conv_welcome_message,omitempty"`
	// Id of the banner you want to display.
	DisplayBannerID *int `json:"display_banner_id,omitempty"`
	// Id of the group you want to display.
	DisplayGroupID *int `json:"display_group_id,omitempty"`
	// Id of the icon group you want to display.
	DisplayIconGroupID *int `json:"display_icon_group_id,omitempty"`
	// Custom user profile fields.
	Fields map[string]interface{} `json:"fields,omitempty"`
	// User gender.
	Gender *UsersEditGender `json:"gender,omitempty"`
	// Hide username change logs.
	HideUsernameChangeLogs *bool `json:"hide_username_change_logs,omitempty"`
	// User interface language ID.
	LanguageID *UsersEditLanguageID `json:"language_id,omitempty"`
	// Whether to receive admin emails.
	ReceiveAdminEmail *bool `json:"receive_admin_email,omitempty"`
	// Secret answer.
	SecretAnswer *string `json:"secret_answer,omitempty"`
	// Secret answer type.
	SecretAnswerType *int `json:"secret_answer_type,omitempty"`
	// Profile short link.
	ShortLink *string `json:"short_link,omitempty"`
	// Show date of birth (day and month).
	ShowDobDate *bool `json:"show_dob_date,omitempty"`
	// Show year of birth.
	ShowDobYear *bool `json:"show_dob_year,omitempty"`
	// User timezone.
	Timezone *UsersEditTimezone `json:"timezone,omitempty"`
	// Your date of birth (day).
	UserDobDay *int `json:"user_dob_day,omitempty"`
	// Your date of birth (month).
	UserDobMonth *int `json:"user_dob_month,omitempty"`
	// Your date of birth (year).
	UserDobYear *int `json:"user_dob_year,omitempty"`
	// New custom title of the user.
	UserTitle *string `json:"user_title,omitempty"`
	// New username.
	Username *string `json:"username,omitempty"`
}

type EditResponse struct {
	Message    interface{}     `json:"message,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type EditResponseMessage struct {
	Alerts                   *int                                `json:"alerts,omitempty"`
	ConversationCreateDate   *int                                `json:"conversation_create_date,omitempty"`
	ConversationID           *int                                `json:"conversation_id,omitempty"`
	ConversationIsDeleted    *bool                               `json:"conversation_is_deleted,omitempty"`
	ConversationIsNew        *bool                               `json:"conversation_is_new,omitempty"`
	ConversationIsOpen       *bool                               `json:"conversation_is_open,omitempty"`
	ConversationLastReadDate *int                                `json:"conversation_last_read_date,omitempty"`
	ConversationMessageCount *int                                `json:"conversation_message_count,omitempty"`
	ConversationOnlineCount  *int                                `json:"conversation_online_count,omitempty"`
	ConversationTitle        *string                             `json:"conversation_title,omitempty"`
	ConversationUpdateDate   *int                                `json:"conversation_update_date,omitempty"`
	CreatorIsIgnored         *bool                               `json:"creator_is_ignored,omitempty"`
	CreatorUserID            *int                                `json:"creator_user_id,omitempty"`
	CreatorUsername          *string                             `json:"creator_username,omitempty"`
	CreatorUsernameHTML      *string                             `json:"creator_username_html,omitempty"`
	IsGroup                  *int                                `json:"is_group,omitempty"`
	IsStarred                *int                                `json:"is_starred,omitempty"`
	IsUnread                 *int                                `json:"is_unread,omitempty"`
	Links                    *EditResponseMessageLinks           `json:"links,omitempty"`
	Permissions              *EditResponseMessagePermissions     `json:"permissions,omitempty"`
	Recipient                *EditResponseMessageRecipient       `json:"recipient,omitempty"`
	Recipients               []EditResponseMessageRecipientsItem `json:"recipients,omitempty"`
}

type EditResponseMessageLinks struct {
	Avatar    *string `json:"avatar,omitempty"`
	Detail    *string `json:"detail,omitempty"`
	Messages  *string `json:"messages,omitempty"`
	Permalink *string `json:"permalink,omitempty"`
}

type EditResponseMessagePermissions struct {
	EditOwnPost       *bool `json:"editOwnPost,omitempty"`
	Invite            *bool `json:"invite,omitempty"`
	Kick              *bool `json:"kick,omitempty"`
	ManageInviteLinks *bool `json:"manage_invite_links,omitempty"`
	Reply             *bool `json:"reply,omitempty"`
	StickyMessages    *bool `json:"stickyMessages,omitempty"`
	UploadAvatar      *bool `json:"upload_avatar,omitempty"`
	View              *bool `json:"view,omitempty"`
}

type EditResponseMessageRecipient struct {
	Avatar          *string `json:"avatar,omitempty"`
	ContactsChanged *bool   `json:"contacts_changed,omitempty"`
	IsOnline        *bool   `json:"is_online,omitempty"`
	LastActivity    *int    `json:"last_activity,omitempty"`
	UserID          *int    `json:"user_id,omitempty"`
	Username        *string `json:"username,omitempty"`
	UsernameHTML    *string `json:"username_html,omitempty"`
}

type EditResponseMessageRecipientsItem struct {
	Avatar          *string `json:"avatar,omitempty"`
	ContactsChanged *bool   `json:"contacts_changed,omitempty"`
	IsOnline        *bool   `json:"is_online,omitempty"`
	LastActivity    *int    `json:"last_activity,omitempty"`
	UserID          *int    `json:"user_id,omitempty"`
	Username        *string `json:"username,omitempty"`
	UsernameHTML    *string `json:"username_html,omitempty"`
}

type EditResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type EnableResponse struct {
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type EnableResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ExecuteResponse struct {
	Jobs *ExecuteResponseJobs `json:"jobs,omitempty"`
}

type ExecuteResponseJobs struct {
	JobID map[string]interface{} `json:"job_id,omitempty"`
}

type FieldsResponse struct {
	Fields     []FieldsResponseFieldsItem `json:"fields,omitempty"`
	SystemInfo *RespSystemInfo            `json:"system_info,omitempty"`
}

type FieldsResponseFieldsItem struct {
	Description *string `json:"description,omitempty"`
	ID          *string `json:"id,omitempty"`
	IsRequired  *bool   `json:"is_required,omitempty"`
	Position    *string `json:"position,omitempty"`
	Title       *string `json:"title,omitempty"`
}

type FieldsResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type FindParams struct {
	// Username to filter. Usernames start with the query will be returned.
	Username *string `json:"username,omitempty"`
	// Custom fields to filter. Example: **custom_fields[telegram]=telegramLogin**.
	CustomFields map[string]string `json:"custom_fields,omitempty"`
	// List of fields to include.
	FieldsInclude []string `json:"fields_include,omitempty"`
}

type FindResponse struct {
	Ids        []int           `json:"ids,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
	Tags       interface{}     `json:"tags,omitempty"`
}

type FindResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type FollowParams struct {
	// Whether to receive notification as email.
	Email *bool `json:"email,omitempty"`
}

type FollowedParams struct {
	// If included in the request, only the thread count is returned as **threads_total**.
	Total *bool `json:"total,omitempty"`
	// List of fields to include.
	FieldsInclude []string `json:"fields_include,omitempty"`
}

type FollowedResponse struct {
	Forums     []FollowedResponseForumsItem `json:"forums,omitempty"`
	SystemInfo *RespSystemInfo              `json:"system_info,omitempty"`
}

type FollowedResponseForumsItem struct {
	Follow                 *FollowedResponseForumsItemFollow             `json:"follow,omitempty"`
	ForumDescription       *string                                       `json:"forum_description,omitempty"`
	ForumID                *int                                          `json:"forum_id,omitempty"`
	ForumIsFollowed        *bool                                         `json:"forum_is_followed,omitempty"`
	ForumPostCount         *int                                          `json:"forum_post_count,omitempty"`
	ForumPrefixes          []FollowedResponseForumsItemForumPrefixesItem `json:"forum_prefixes,omitempty"`
	ForumThreadCount       *int                                          `json:"forum_thread_count,omitempty"`
	ForumTitle             *string                                       `json:"forum_title,omitempty"`
	Links                  *FollowedResponseForumsItemLinks              `json:"links,omitempty"`
	Permissions            *FollowedResponseForumsItemPermissions        `json:"permissions,omitempty"`
	ThreadDefaultPrefixID  *int                                          `json:"thread_default_prefix_id,omitempty"`
	ThreadPrefixIsRequired *bool                                         `json:"thread_prefix_is_required,omitempty"`
}

type FollowedResponseForumsItemFollow struct {
	Alert *bool `json:"alert,omitempty"`
	Email *bool `json:"email,omitempty"`
	Post  *bool `json:"post,omitempty"`
}

type FollowedResponseForumsItemForumPrefixesItem struct {
	GroupPrefixes []FollowedResponseForumsItemForumPrefixesItemGroupPrefixesItem `json:"group_prefixes,omitempty"`
	GroupTitle    *string                                                        `json:"group_title,omitempty"`
}

type FollowedResponseForumsItemForumPrefixesItemGroupPrefixesItem struct {
	PrefixID    *int    `json:"prefix_id,omitempty"`
	PrefixTitle *string `json:"prefix_title,omitempty"`
}

type FollowedResponseForumsItemLinks struct {
	Detail        *string `json:"detail,omitempty"`
	Followers     *string `json:"followers,omitempty"`
	Permalink     *string `json:"permalink,omitempty"`
	SubCategories *string `json:"sub-categories,omitempty"`
	SubForums     *string `json:"sub-forums,omitempty"`
	Threads       *string `json:"threads,omitempty"`
}

type FollowedResponseForumsItemPermissions struct {
	CreateThread     *bool `json:"create_thread,omitempty"`
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	TagThread        *bool `json:"tag_thread,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type FollowedResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type FollowersParams struct {
	// Ordering of followers.
	Order *UsersOrder `json:"order,omitempty"`
	// Page number of followers.
	Page *int `json:"page,omitempty"`
	// Number of followers in a page.
	Limit *int `json:"limit,omitempty"`
}

type FollowersResponse struct {
	SystemInfo *RespSystemInfo              `json:"system_info,omitempty"`
	Users      []FollowersResponseUsersItem `json:"users,omitempty"`
}

type FollowersResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type FollowersResponseUsersItem struct {
	Follow   *FollowersResponseUsersItemFollow `json:"follow,omitempty"`
	UserID   *int                              `json:"user_id,omitempty"`
	Username *string                           `json:"username,omitempty"`
}

type FollowersResponseUsersItemFollow struct {
	Alert *bool `json:"alert,omitempty"`
	Email *bool `json:"email,omitempty"`
	Post  *bool `json:"post,omitempty"`
}

type FollowingsParams struct {
	// Ordering of users.
	Order *UsersOrder `json:"order,omitempty"`
	// Page number of users.
	Page *int `json:"page,omitempty"`
	// Number of users in a page.
	Limit *int `json:"limit,omitempty"`
}

type FollowingsResponse struct {
	SystemInfo *RespSystemInfo               `json:"system_info,omitempty"`
	Users      []FollowingsResponseUsersItem `json:"users,omitempty"`
	UsersTotal *int                          `json:"users_total,omitempty"`
}

type FollowingsResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type FollowingsResponseUsersItem struct {
	ContentID          interface{}                              `json:"content_id,omitempty"`
	ContentType        *string                                  `json:"content_type,omitempty"`
	ContestCount       *int                                     `json:"contest_count,omitempty"`
	CustomFields       *FollowingsResponseUsersItemCustomFields `json:"custom_fields,omitempty"`
	CustomTitle        *string                                  `json:"custom_title,omitempty"`
	FollowDate         *int                                     `json:"follow_date,omitempty"`
	IsBanned           *int                                     `json:"is_banned,omitempty"`
	Links              *FollowingsResponseUsersItemLinks        `json:"links,omitempty"`
	Permissions        *FollowingsResponseUsersItemPermissions  `json:"permissions,omitempty"`
	ShortLink          *string                                  `json:"short_link,omitempty"`
	TrophyCount        *int                                     `json:"trophy_count,omitempty"`
	UserFollowersCount *int                                     `json:"user_followers_count,omitempty"`
	UserFollowingCount *int                                     `json:"user_following_count,omitempty"`
	UserGroupID        *int                                     `json:"user_group_id,omitempty"`
	UserID             *int                                     `json:"user_id,omitempty"`
	UserIsFollowed     *bool                                    `json:"user_is_followed,omitempty"`
	UserIsIgnored      *bool                                    `json:"user_is_ignored,omitempty"`
	UserIsValid        *bool                                    `json:"user_is_valid,omitempty"`
	UserIsVerified     *bool                                    `json:"user_is_verified,omitempty"`
	UserIsVisitor      *bool                                    `json:"user_is_visitor,omitempty"`
	UserLastSeenDate   *int                                     `json:"user_last_seen_date,omitempty"`
	UserLike2Count     *int                                     `json:"user_like2_count,omitempty"`
	UserLikeCount      *int                                     `json:"user_like_count,omitempty"`
	UserMessageCount   *int                                     `json:"user_message_count,omitempty"`
	UserRegisterDate   *int                                     `json:"user_register_date,omitempty"`
	UserTitle          *string                                  `json:"user_title,omitempty"`
	Username           *string                                  `json:"username,omitempty"`
	UsernameHTML       *string                                  `json:"username_html,omitempty"`
}

type FollowingsResponseUsersItemCustomFields struct {
	Field4                *string       `json:"_4,omitempty"`
	AllowSelfUnban        []interface{} `json:"allowSelfUnban,omitempty"`
	Discord               *string       `json:"discord,omitempty"`
	Github                *string       `json:"github,omitempty"`
	Jabber                *string       `json:"jabber,omitempty"`
	LztAwardUserTrophy    *string       `json:"lztAwardUserTrophy,omitempty"`
	LztCuratorNodeTitle   *string       `json:"lztCuratorNodeTitle,omitempty"`
	LztCuratorNodeTitleEn *string       `json:"lztCuratorNodeTitleEn,omitempty"`
	LztDeposit            *string       `json:"lztDeposit,omitempty"`
	LztInnovation20Link   *string       `json:"lztInnovation20Link,omitempty"`
	LztInnovation30Link   *string       `json:"lztInnovation30Link,omitempty"`
	LztInnovationLink     *string       `json:"lztInnovationLink,omitempty"`
	LztLikesIncreasing    *string       `json:"lztLikesIncreasing,omitempty"`
	LztLikesZeroing       *string       `json:"lztLikesZeroing,omitempty"`
	LztSympathyIncreasing *string       `json:"lztSympathyIncreasing,omitempty"`
	LztSympathyZeroing    *string       `json:"lztSympathyZeroing,omitempty"`
	MaecenasValue         *string       `json:"maecenasValue,omitempty"`
	ScamURL               *string       `json:"scamURL,omitempty"`
	Steam                 *string       `json:"steam,omitempty"`
	Telegram              *string       `json:"telegram,omitempty"`
	Vk                    *string       `json:"vk,omitempty"`
}

type FollowingsResponseUsersItemLinks struct {
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

type FollowingsResponseUsersItemPermissions struct {
	Edit        *bool `json:"edit,omitempty"`
	Follow      *bool `json:"follow,omitempty"`
	Ignore      *bool `json:"ignore,omitempty"`
	ProfilePost *bool `json:"profile_post,omitempty"`
}

type GetFeedOptionsResponse struct {
	DefaultExcludedForumsIds []int                              `json:"default_excluded_forums_ids,omitempty"`
	ExcludedForumsIds        []int                              `json:"excluded_forums_ids,omitempty"`
	Forums                   []GetFeedOptionsResponseForumsItem `json:"forums,omitempty"`
	Keywords                 *string                            `json:"keywords,omitempty"`
	SystemInfo               *RespSystemInfo                    `json:"system_info,omitempty"`
}

type GetFeedOptionsResponseForumsItem struct {
	ForumDescription *string                                      `json:"forum_description,omitempty"`
	ForumID          *int                                         `json:"forum_id,omitempty"`
	ForumIsFollowed  *bool                                        `json:"forum_is_followed,omitempty"`
	ForumTitle       *string                                      `json:"forum_title,omitempty"`
	HasChildren      *bool                                        `json:"has_children,omitempty"`
	Links            *GetFeedOptionsResponseForumsItemLinks       `json:"links,omitempty"`
	ParentNodeID     *int                                         `json:"parent_node_id,omitempty"`
	Permissions      *GetFeedOptionsResponseForumsItemPermissions `json:"permissions,omitempty"`
}

type GetFeedOptionsResponseForumsItemLinks struct {
	Detail        *string `json:"detail,omitempty"`
	Followers     *string `json:"followers,omitempty"`
	Permalink     *string `json:"permalink,omitempty"`
	SubCategories *string `json:"sub-categories,omitempty"`
	SubForums     *string `json:"sub-forums,omitempty"`
	Threads       *string `json:"threads,omitempty"`
}

type GetFeedOptionsResponseForumsItemPermissions struct {
	CreateThread *bool `json:"create_thread,omitempty"`
	Delete       *bool `json:"delete,omitempty"`
	Edit         *bool `json:"edit,omitempty"`
	Follow       *bool `json:"follow,omitempty"`
	TagThread    *bool `json:"tag_thread,omitempty"`
	View         *bool `json:"view,omitempty"`
}

type GetFeedOptionsResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type GetGetResponse struct {
	Conversation *RespConversationModel `json:"conversation,omitempty"`
	SystemInfo   *RespSystemInfo        `json:"system_info,omitempty"`
}

type GetGetResponseConversation struct {
	Alerts                   *int                                       `json:"alerts,omitempty"`
	ConversationCreateDate   *int                                       `json:"conversation_create_date,omitempty"`
	ConversationID           *int                                       `json:"conversation_id,omitempty"`
	ConversationIsDeleted    *bool                                      `json:"conversation_is_deleted,omitempty"`
	ConversationIsNew        *bool                                      `json:"conversation_is_new,omitempty"`
	ConversationIsOpen       *bool                                      `json:"conversation_is_open,omitempty"`
	ConversationLastReadDate *int                                       `json:"conversation_last_read_date,omitempty"`
	ConversationMessageCount *int                                       `json:"conversation_message_count,omitempty"`
	ConversationOnlineCount  *int                                       `json:"conversation_online_count,omitempty"`
	ConversationTitle        *string                                    `json:"conversation_title,omitempty"`
	ConversationUpdateDate   *int                                       `json:"conversation_update_date,omitempty"`
	CreatorIsIgnored         *bool                                      `json:"creator_is_ignored,omitempty"`
	CreatorUserID            *int                                       `json:"creator_user_id,omitempty"`
	CreatorUsername          *string                                    `json:"creator_username,omitempty"`
	CreatorUsernameHTML      *string                                    `json:"creator_username_html,omitempty"`
	IsGroup                  *int                                       `json:"is_group,omitempty"`
	IsStarred                *int                                       `json:"is_starred,omitempty"`
	IsUnread                 *int                                       `json:"is_unread,omitempty"`
	Links                    *GetGetResponseConversationLinks           `json:"links,omitempty"`
	Permissions              *GetGetResponseConversationPermissions     `json:"permissions,omitempty"`
	Recipient                *GetGetResponseConversationRecipient       `json:"recipient,omitempty"`
	Recipients               []GetGetResponseConversationRecipientsItem `json:"recipients,omitempty"`
}

type GetGetResponseConversationLinks struct {
	Avatar    *string `json:"avatar,omitempty"`
	Detail    *string `json:"detail,omitempty"`
	Messages  *string `json:"messages,omitempty"`
	Permalink *string `json:"permalink,omitempty"`
}

type GetGetResponseConversationPermissions struct {
	EditOwnPost       *bool `json:"editOwnPost,omitempty"`
	Invite            *bool `json:"invite,omitempty"`
	Kick              *bool `json:"kick,omitempty"`
	ManageInviteLinks *bool `json:"manage_invite_links,omitempty"`
	Reply             *bool `json:"reply,omitempty"`
	StickyMessages    *bool `json:"stickyMessages,omitempty"`
	UploadAvatar      *bool `json:"upload_avatar,omitempty"`
	View              *bool `json:"view,omitempty"`
}

type GetGetResponseConversationRecipient struct {
	Avatar          *string `json:"avatar,omitempty"`
	ContactsChanged *bool   `json:"contacts_changed,omitempty"`
	IsOnline        *bool   `json:"is_online,omitempty"`
	LastActivity    *int    `json:"last_activity,omitempty"`
	UserID          *int    `json:"user_id,omitempty"`
	Username        *string `json:"username,omitempty"`
	UsernameHTML    *string `json:"username_html,omitempty"`
}

type GetGetResponseConversationRecipientsItem struct {
	Avatar          *string `json:"avatar,omitempty"`
	ContactsChanged *bool   `json:"contacts_changed,omitempty"`
	IsOnline        *bool   `json:"is_online,omitempty"`
	LastActivity    *int    `json:"last_activity,omitempty"`
	UserID          *int    `json:"user_id,omitempty"`
	Username        *string `json:"username,omitempty"`
	UsernameHTML    *string `json:"username_html,omitempty"`
}

type GetGetResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type GetIgnoreResponse struct {
	Ignored    []GetIgnoreResponseIgnoredItem `json:"ignored,omitempty"`
	SystemInfo *RespSystemInfo                `json:"system_info,omitempty"`
}

type GetIgnoreResponseIgnoredItem struct {
	AvatarDate          *int                                  `json:"avatar_date,omitempty"`
	BackgroundDate      *int                                  `json:"background_date,omitempty"`
	ContestCount        *int                                  `json:"contest_count,omitempty"`
	CustomTitle         *string                               `json:"custom_title,omitempty"`
	DisplayBannerID     *int                                  `json:"display_banner_id,omitempty"`
	DisplayIconGroupID  *int                                  `json:"display_icon_group_id,omitempty"`
	DisplayStyleGroupID *int                                  `json:"display_style_group_id,omitempty"`
	IsBanned            *bool                                 `json:"is_banned,omitempty"`
	LastActivity        *int                                  `json:"last_activity,omitempty"`
	Like2Count          *int                                  `json:"like2_count,omitempty"`
	LikeCount           *int                                  `json:"like_count,omitempty"`
	MessageCount        *int                                  `json:"message_count,omitempty"`
	RegisterDate        *int                                  `json:"register_date,omitempty"`
	Rendered            *GetIgnoreResponseIgnoredItemRendered `json:"rendered,omitempty"`
	ShortLink           interface{}                           `json:"short_link,omitempty"`
	TrophyPoints        *int                                  `json:"trophy_points,omitempty"`
	UniqBanner          interface{}                           `json:"uniq_banner,omitempty"`
	UniqUsernameCSS     *string                               `json:"uniq_username_css,omitempty"`
	UserID              *int                                  `json:"user_id,omitempty"`
	Username            *string                               `json:"username,omitempty"`
}

type GetIgnoreResponseIgnoredItemRendered struct {
	Avatars  *GetIgnoreResponseIgnoredItemRenderedAvatars `json:"avatars,omitempty"`
	Link     *string                                      `json:"link,omitempty"`
	Username *string                                      `json:"username,omitempty"`
}

type GetIgnoreResponseIgnoredItemRenderedAvatars struct {
	L *string `json:"l,omitempty"`
	M *string `json:"m,omitempty"`
	S *string `json:"s,omitempty"`
}

type GetIgnoreResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type GetLeaderboardParams struct {
	// Duration.
	Duration *ChatboxDuration `json:"duration,omitempty"`
}

type GetLeaderboardResponse struct {
	Leaderboard []GetLeaderboardResponseLeaderboardItem `json:"leaderboard,omitempty"`
	SystemInfo  *RespSystemInfo                         `json:"system_info,omitempty"`
}

type GetLeaderboardResponseLeaderboardItem struct {
	AvatarDate          *int                                             `json:"avatar_date,omitempty"`
	BackgroundDate      *int                                             `json:"background_date,omitempty"`
	ContestCount        *int                                             `json:"contest_count,omitempty"`
	Count               *int                                             `json:"count,omitempty"`
	CustomTitle         *string                                          `json:"custom_title,omitempty"`
	DisplayBannerID     *int                                             `json:"display_banner_id,omitempty"`
	DisplayIconGroupID  *int                                             `json:"display_icon_group_id,omitempty"`
	DisplayStyleGroupID *int                                             `json:"display_style_group_id,omitempty"`
	IsBanned            *bool                                            `json:"is_banned,omitempty"`
	LastActivity        *int                                             `json:"last_activity,omitempty"`
	Like2Count          *int                                             `json:"like2_count,omitempty"`
	LikeCount           *int                                             `json:"like_count,omitempty"`
	MessageCount        *int                                             `json:"message_count,omitempty"`
	RegisterDate        *int                                             `json:"register_date,omitempty"`
	Rendered            *GetLeaderboardResponseLeaderboardItemRendered   `json:"rendered,omitempty"`
	ShortLink           interface{}                                      `json:"short_link,omitempty"`
	TrophyPoints        *int                                             `json:"trophy_points,omitempty"`
	UniqBanner          *GetLeaderboardResponseLeaderboardItemUniqBanner `json:"uniq_banner,omitempty"`
	UniqUsernameCSS     *string                                          `json:"uniq_username_css,omitempty"`
	UserID              *int                                             `json:"user_id,omitempty"`
	Username            *string                                          `json:"username,omitempty"`
}

type GetLeaderboardResponseLeaderboardItemRendered struct {
	Avatars  *GetLeaderboardResponseLeaderboardItemRenderedAvatars `json:"avatars,omitempty"`
	Link     *string                                               `json:"link,omitempty"`
	Username *string                                               `json:"username,omitempty"`
}

type GetLeaderboardResponseLeaderboardItemRenderedAvatars struct {
	L *string `json:"l,omitempty"`
	M *string `json:"m,omitempty"`
	S *string `json:"s,omitempty"`
}

type GetLeaderboardResponseLeaderboardItemUniqBanner struct {
	BannerCSS  *string `json:"banner_css,omitempty"`
	BannerIcon *string `json:"banner_icon,omitempty"`
	BannerText *string `json:"banner_text,omitempty"`
}

type GetLeaderboardResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type GetMessagesParams struct {
	// Message id to get older chat messages.
	BeforeMessageID *int `json:"before_message_id,omitempty"`
}

type GetMessagesResponse struct {
	Messages   []GetMessagesResponseMessagesItem `json:"messages,omitempty"`
	SystemInfo *RespSystemInfo                   `json:"system_info,omitempty"`
}

type GetMessagesResponseMessagesItem struct {
	CanReport   *bool                                `json:"can_report,omitempty"`
	Date        *int                                 `json:"date,omitempty"`
	IsDeleted   *bool                                `json:"is_deleted,omitempty"`
	Message     *string                              `json:"message,omitempty"`
	MessageJson *string                              `json:"messageJson,omitempty"`
	MessageRaw  *string                              `json:"messageRaw,omitempty"`
	MessageID   *int                                 `json:"message_id,omitempty"`
	Room        *GetMessagesResponseMessagesItemRoom `json:"room,omitempty"`
	User        *GetMessagesResponseMessagesItemUser `json:"user,omitempty"`
}

type GetMessagesResponseMessagesItemRoom struct {
	CanReport *bool   `json:"can_report,omitempty"`
	Eng       *bool   `json:"eng,omitempty"`
	Market    *bool   `json:"market,omitempty"`
	RoomID    *int    `json:"room_id,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type GetMessagesResponseMessagesItemUser struct {
	AvatarDate          *int                                           `json:"avatar_date,omitempty"`
	BackgroundDate      *int                                           `json:"background_date,omitempty"`
	ContestCount        *int                                           `json:"contest_count,omitempty"`
	CustomTitle         *string                                        `json:"custom_title,omitempty"`
	DisplayBannerID     *int                                           `json:"display_banner_id,omitempty"`
	DisplayIconGroupID  *int                                           `json:"display_icon_group_id,omitempty"`
	DisplayStyleGroupID *int                                           `json:"display_style_group_id,omitempty"`
	IsAdmin             *bool                                          `json:"is_admin,omitempty"`
	IsBanned            *bool                                          `json:"is_banned,omitempty"`
	IsModerator         *bool                                          `json:"is_moderator,omitempty"`
	IsStaff             *bool                                          `json:"is_staff,omitempty"`
	LastActivity        *int                                           `json:"last_activity,omitempty"`
	Like2Count          *int                                           `json:"like2_count,omitempty"`
	LikeCount           *int                                           `json:"like_count,omitempty"`
	MessageCount        *int                                           `json:"message_count,omitempty"`
	RegisterDate        *int                                           `json:"register_date,omitempty"`
	Rendered            *GetMessagesResponseMessagesItemUserRendered   `json:"rendered,omitempty"`
	ShortLink           *string                                        `json:"short_link,omitempty"`
	TrophyPoints        *int                                           `json:"trophy_points,omitempty"`
	UniqBanner          *GetMessagesResponseMessagesItemUserUniqBanner `json:"uniq_banner,omitempty"`
	UniqUsernameCSS     *string                                        `json:"uniq_username_css,omitempty"`
	UserID              *int                                           `json:"user_id,omitempty"`
	Username            *string                                        `json:"username,omitempty"`
}

type GetMessagesResponseMessagesItemUserRendered struct {
	Avatars  *GetMessagesResponseMessagesItemUserRenderedAvatars `json:"avatars,omitempty"`
	Link     *string                                             `json:"link,omitempty"`
	Username *string                                             `json:"username,omitempty"`
}

type GetMessagesResponseMessagesItemUserRenderedAvatars struct {
	L *string `json:"l,omitempty"`
	M *string `json:"m,omitempty"`
	S *string `json:"s,omitempty"`
}

type GetMessagesResponseMessagesItemUserUniqBanner struct {
	BannerCSS    *string `json:"banner_css,omitempty"`
	BannerIcon   *string `json:"banner_icon,omitempty"`
	BannerText   *string `json:"banner_text,omitempty"`
	UsernameIcon *string `json:"username_icon,omitempty"`
}

type GetMessagesResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type GetParams struct {
	// List of fields to include.
	FieldsInclude []string `json:"fields_include,omitempty"`
}

type GetResponse struct {
	Category   *GetResponseCategory `json:"category,omitempty"`
	SystemInfo *RespSystemInfo      `json:"system_info,omitempty"`
}

type GetResponseCategory struct {
	CategoryDescription *string                         `json:"category_description,omitempty"`
	CategoryID          *int                            `json:"category_id,omitempty"`
	CategoryTitle       *string                         `json:"category_title,omitempty"`
	Links               *GetResponseCategoryLinks       `json:"links,omitempty"`
	Permissions         *GetResponseCategoryPermissions `json:"permissions,omitempty"`
}

type GetResponseCategoryLinks struct {
	Detail        *string `json:"detail,omitempty"`
	Permalink     *string `json:"permalink,omitempty"`
	SubCategories *string `json:"sub-categories,omitempty"`
	SubForums     *string `json:"sub-forums,omitempty"`
}

type GetResponseCategoryPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type GetResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type GroupedResponse struct {
	Data       interface{}               `json:"data,omitempty"`
	SystemInfo *RespSystemInfo           `json:"system_info,omitempty"`
	Tabs       []GroupedResponseTabsItem `json:"tabs,omitempty"`
}

type GroupedResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type GroupedResponseTabsItem struct {
	IsDefault *bool   `json:"isDefault,omitempty"`
	IsHidden  *bool   `json:"isHidden,omitempty"`
	LinkTitle *string `json:"link_title,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type HideResponse struct {
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type HideResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type IgnoreEditParams struct {
	// Ignore user's conversations.
	IgnoreConversations *bool `json:"ignore_conversations,omitempty"`
	// Ignore user's content.
	IgnoreContent *bool `json:"ignore_content,omitempty"`
	// Restrict user from viewing your profile.
	RestrictViewProfile *bool `json:"restrict_view_profile,omitempty"`
}

type IgnoredParams struct {
	// If included in the request, only the user count is returned as **users_total**.
	Total *bool `json:"total,omitempty"`
}

type IgnoredResponse struct {
	SystemInfo *RespSystemInfo            `json:"system_info,omitempty"`
	Users      []IgnoredResponseUsersItem `json:"users,omitempty"`
}

type IgnoredResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type IgnoredResponseUsersItem struct {
	CanEdit             *bool                                 `json:"can_edit,omitempty"`
	CanFollow           *bool                                 `json:"can_follow,omitempty"`
	CanIgnore           *bool                                 `json:"can_ignore,omitempty"`
	CanPostProfile      *bool                                 `json:"can_post_profile,omitempty"`
	CanViewProfile      *bool                                 `json:"can_view_profile,omitempty"`
	CanViewProfilePosts *bool                                 `json:"can_view_profile_posts,omitempty"`
	CanWarn             *bool                                 `json:"can_warn,omitempty"`
	ContestCount        *int                                  `json:"contest_count,omitempty"`
	ConvWelcomeMessage  *string                               `json:"conv_welcome_message,omitempty"`
	ConvertedDeposit    *int                                  `json:"convertedDeposit,omitempty"`
	CustomFields        *IgnoredResponseUsersItemCustomFields `json:"custom_fields,omitempty"`
	Deposit             *int                                  `json:"deposit,omitempty"`
	Homepage            *string                               `json:"homepage,omitempty"`
	IgnoredInfo         *IgnoredResponseUsersItemIgnoredInfo  `json:"ignored_info,omitempty"`
	IsAdmin             *bool                                 `json:"is_admin,omitempty"`
	IsBanned            *bool                                 `json:"is_banned,omitempty"`
	IsFollowed          *bool                                 `json:"is_followed,omitempty"`
	IsIgnored           *bool                                 `json:"is_ignored,omitempty"`
	IsModerator         *bool                                 `json:"is_moderator,omitempty"`
	IsStaff             *bool                                 `json:"is_staff,omitempty"`
	LastActivity        *int                                  `json:"last_activity,omitempty"`
	Like2Count          *int                                  `json:"like2_count,omitempty"`
	LikeCount           *int                                  `json:"like_count,omitempty"`
	Location            *string                               `json:"location,omitempty"`
	MessageCount        *int                                  `json:"message_count,omitempty"`
	RegisterDate        *int                                  `json:"register_date,omitempty"`
	Rendered            *IgnoredResponseUsersItemRendered     `json:"rendered,omitempty"`
	ShortLink           *string                               `json:"short_link,omitempty"`
	TrophyPoints        *int                                  `json:"trophy_points,omitempty"`
	UserID              *int                                  `json:"user_id,omitempty"`
	UserTitle           *string                               `json:"user_title,omitempty"`
	Username            *string                               `json:"username,omitempty"`
	ViewURL             *string                               `json:"view_url,omitempty"`
	WarningPoints       *int                                  `json:"warning_points,omitempty"`
}

type IgnoredResponseUsersItemCustomFields struct {
	Field4                *string     `json:"_4,omitempty"`
	Discord               *string     `json:"discord,omitempty"`
	Github                *string     `json:"github,omitempty"`
	Jabber                *string     `json:"jabber,omitempty"`
	LztLikesIncreasing    interface{} `json:"lztLikesIncreasing,omitempty"`
	LztLikesZeroing       interface{} `json:"lztLikesZeroing,omitempty"`
	LztSympathyIncreasing interface{} `json:"lztSympathyIncreasing,omitempty"`
	LztSympathyZeroing    interface{} `json:"lztSympathyZeroing,omitempty"`
	Matrix                interface{} `json:"matrix,omitempty"`
	ScamURL               interface{} `json:"scamURL,omitempty"`
	Steam                 *string     `json:"steam,omitempty"`
	Telegram              interface{} `json:"telegram,omitempty"`
	Vk                    *string     `json:"vk,omitempty"`
}

type IgnoredResponseUsersItemIgnoredInfo struct {
	IgnoreContent       *int `json:"ignore_content,omitempty"`
	IgnoreConversations *int `json:"ignore_conversations,omitempty"`
	RestrictViewProfile *int `json:"restrict_view_profile,omitempty"`
}

type IgnoredResponseUsersItemRendered struct {
	Avatars     *IgnoredResponseUsersItemRenderedAvatars `json:"avatars,omitempty"`
	Backgrounds interface{}                              `json:"backgrounds,omitempty"`
	Link        *string                                  `json:"link,omitempty"`
	Username    *string                                  `json:"username,omitempty"`
}

type IgnoredResponseUsersItemRenderedAvatars struct {
	L *string `json:"l,omitempty"`
	M *string `json:"m,omitempty"`
	S *string `json:"s,omitempty"`
}

type IndexParams struct {
	// Room id.
	RoomID *ChatboxRoomID `json:"room_id,omitempty"`
}

type IndexResponse struct {
	Ban         interface{}               `json:"ban,omitempty"`
	Commands    []string                  `json:"commands,omitempty"`
	Ignore      []IndexResponseIgnoreItem `json:"ignore,omitempty"`
	Permissions *IndexResponsePermissions `json:"permissions,omitempty"`
	Rooms       []IndexResponseRoomsItem  `json:"rooms,omitempty"`
	RoomsOnline *IndexResponseRoomsOnline `json:"roomsOnline,omitempty"`
	SystemInfo  *RespSystemInfo           `json:"system_info,omitempty"`
}

type IndexResponseIgnoreItem struct {
	AvatarDate          *int                             `json:"avatar_date,omitempty"`
	BackgroundDate      *int                             `json:"background_date,omitempty"`
	ContestCount        *int                             `json:"contest_count,omitempty"`
	CustomTitle         *string                          `json:"custom_title,omitempty"`
	DisplayBannerID     *int                             `json:"display_banner_id,omitempty"`
	DisplayIconGroupID  *int                             `json:"display_icon_group_id,omitempty"`
	DisplayStyleGroupID *int                             `json:"display_style_group_id,omitempty"`
	IsAdmin             *bool                            `json:"is_admin,omitempty"`
	IsBanned            *bool                            `json:"is_banned,omitempty"`
	IsModerator         *bool                            `json:"is_moderator,omitempty"`
	IsStaff             *bool                            `json:"is_staff,omitempty"`
	LastActivity        *int                             `json:"last_activity,omitempty"`
	Like2Count          *int                             `json:"like2_count,omitempty"`
	LikeCount           *int                             `json:"like_count,omitempty"`
	MessageCount        *int                             `json:"message_count,omitempty"`
	RegisterDate        *int                             `json:"register_date,omitempty"`
	Rendered            *IndexResponseIgnoreItemRendered `json:"rendered,omitempty"`
	ShortLink           interface{}                      `json:"short_link,omitempty"`
	TrophyPoints        *int                             `json:"trophy_points,omitempty"`
	UniqBanner          interface{}                      `json:"uniq_banner,omitempty"`
	UniqUsernameCSS     *string                          `json:"uniq_username_css,omitempty"`
	UserID              *int                             `json:"user_id,omitempty"`
	Username            *string                          `json:"username,omitempty"`
}

type IndexResponseIgnoreItemRendered struct {
	Avatars  *IndexResponseIgnoreItemRenderedAvatars `json:"avatars,omitempty"`
	Link     *string                                 `json:"link,omitempty"`
	Username *string                                 `json:"username,omitempty"`
}

type IndexResponseIgnoreItemRenderedAvatars struct {
	L *string `json:"l,omitempty"`
	M *string `json:"m,omitempty"`
	S *string `json:"s,omitempty"`
}

type IndexResponsePermissions struct {
	Ban              *bool `json:"ban,omitempty"`
	DeleteAnyMessage *bool `json:"deleteAnyMessage,omitempty"`
	EditAnyMessage   *bool `json:"editAnyMessage,omitempty"`
	PostMessage      *bool `json:"postMessage,omitempty"`
	ViewAnyMessage   *bool `json:"viewAnyMessage,omitempty"`
	ViewMessages     *bool `json:"viewMessages,omitempty"`
}

type IndexResponseRoomsItem struct {
	CanReport *bool   `json:"can_report,omitempty"`
	Eng       *bool   `json:"eng,omitempty"`
	Market    *bool   `json:"market,omitempty"`
	RoomID    *int    `json:"room_id,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type IndexResponseRoomsOnline struct {
	Chat0 *int `json:"chat:0,omitempty"`
}

type IndexResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type LikesParams struct {
	// Filter by forum section.
	NodeID *int `json:"node_id,omitempty"`
	// Like type.
	LikeType *UsersLikeType `json:"like_type,omitempty"`
	// Likes type.
	Type_ *UsersType `json:"type,omitempty"`
	// Page number.
	Page *int `json:"page,omitempty"`
	// Content type.
	ContentType *UsersContentType `json:"content_type,omitempty"`
	// Get only likes from specified user.
	SearchUserID *int `json:"search_user_id,omitempty"`
	// Show weekly statistics.
	Stats *bool `json:"stats,omitempty"`
}

type LikesResponse struct {
	SystemInfo *RespSystemInfo          `json:"system_info,omitempty"`
	Users      []LikesResponseUsersItem `json:"users,omitempty"`
}

type LikesResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type LikesResponseUsersItem struct {
	UserID   *int    `json:"user_id,omitempty"`
	Username *string `json:"username,omitempty"`
}

type ListGetParams struct {
	// Page number of messages.
	Page *int `json:"page,omitempty"`
	// Number of messages in a page.
	Limit *int `json:"limit,omitempty"`
	// Ordering of messages.
	Order *ConversationsOrder `json:"order,omitempty"`
	// Date to get older messages.
	Before *int `json:"before,omitempty"`
	// Date to get newer messages.
	After *int `json:"after,omitempty"`
}

type ListGetResponse struct {
	Links         *ListGetResponseLinks         `json:"links,omitempty"`
	Messages      []ListGetResponseMessagesItem `json:"messages,omitempty"`
	MessagesTotal *int                          `json:"messages_total,omitempty"`
	SystemInfo    *RespSystemInfo               `json:"system_info,omitempty"`
}

type ListGetResponseLinks struct {
	Next  *string `json:"next,omitempty"`
	Page  *int    `json:"page,omitempty"`
	Pages *int    `json:"pages,omitempty"`
}

type ListGetResponseMessagesItem struct {
	ConversationID       *int                                    `json:"conversation_id,omitempty"`
	CreatorUserID        *int                                    `json:"creator_user_id,omitempty"`
	CreatorUsername      *string                                 `json:"creator_username,omitempty"`
	CreatorUsernameHTML  *string                                 `json:"creator_username_html,omitempty"`
	Links                *ListGetResponseMessagesItemLinks       `json:"links,omitempty"`
	MessageBody          *string                                 `json:"message_body,omitempty"`
	MessageBodyHTML      *string                                 `json:"message_body_html,omitempty"`
	MessageBodyPlainText *string                                 `json:"message_body_plain_text,omitempty"`
	MessageCreateDate    *int                                    `json:"message_create_date,omitempty"`
	MessageEditDate      *int                                    `json:"message_edit_date,omitempty"`
	MessageID            *int                                    `json:"message_id,omitempty"`
	MessageIsSystem      *bool                                   `json:"message_is_system,omitempty"`
	MessageIsUnread      *int                                    `json:"message_is_unread,omitempty"`
	MessageNeedTranslate *bool                                   `json:"message_need_translate,omitempty"`
	Permissions          *ListGetResponseMessagesItemPermissions `json:"permissions,omitempty"`
	UserIsIgnored        *bool                                   `json:"user_is_ignored,omitempty"`
}

type ListGetResponseMessagesItemLinks struct {
	Conversation  *string `json:"conversation,omitempty"`
	Creator       *string `json:"creator,omitempty"`
	CreatorAvatar *string `json:"creator_avatar,omitempty"`
	Detail        *string `json:"detail,omitempty"`
}

type ListGetResponseMessagesItemPermissions struct {
	Delete       *bool `json:"delete,omitempty"`
	Edit         *bool `json:"edit,omitempty"`
	StickUnstick *bool `json:"stick-unstick,omitempty"`
	View         *bool `json:"view,omitempty"`
}

type ListGetResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ListParams struct {
	// Filter to get only posts from the specified user.
	PostsUserID *int `json:"posts_user_id,omitempty"`
	// Page number of contents.
	Page *int `json:"page,omitempty"`
	// Number of contents in a page.
	Limit *int `json:"limit,omitempty"`
	// List of fields to include.
	FieldsInclude []string `json:"fields_include,omitempty"`
}

type ListResponse struct {
	Categories      []ListResponseCategoriesItem `json:"categories,omitempty"`
	CategoriesTotal *int                         `json:"categories_total,omitempty"`
	SystemInfo      *RespSystemInfo              `json:"system_info,omitempty"`
}

type ListResponseCategoriesItem struct {
	CategoryDescription *string                                `json:"category_description,omitempty"`
	CategoryID          *int                                   `json:"category_id,omitempty"`
	CategoryTitle       *string                                `json:"category_title,omitempty"`
	Links               *ListResponseCategoriesItemLinks       `json:"links,omitempty"`
	Permissions         *ListResponseCategoriesItemPermissions `json:"permissions,omitempty"`
}

type ListResponseCategoriesItemLinks struct {
	Detail        *string `json:"detail,omitempty"`
	Permalink     *string `json:"permalink,omitempty"`
	SubCategories *string `json:"sub-categories,omitempty"`
	SubForums     *string `json:"sub-forums,omitempty"`
}

type ListResponseCategoriesItemPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type ListResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type MoveParams struct {
	// Apply thread prefix.
	ApplyThreadPrefix *bool `json:"apply_thread_prefix,omitempty"`
	// Prefix ids. Set "0" to remove all thread prefixes.
	PrefixID []int `json:"prefix_id,omitempty"`
	// Send a notification to users who are followed to target node.
	SendAlert *bool `json:"send_alert,omitempty"`
	// Thread title.
	Title *string `json:"title,omitempty"`
	// Thread title english.
	TitleEn *string `json:"title_en,omitempty"`
}

type NavigationResponse struct {
	Elements      []NavigationResponseElementsItem `json:"elements,omitempty"`
	ElementsCount *int                             `json:"elements_count,omitempty"`
	SystemInfo    *RespSystemInfo                  `json:"system_info,omitempty"`
}

type NavigationResponseElementsItem struct {
	CategoryDescription *string                                    `json:"category_description,omitempty"`
	CategoryID          *int                                       `json:"category_id,omitempty"`
	CategoryTitle       *string                                    `json:"category_title,omitempty"`
	HasSubElements      *bool                                      `json:"has_sub_elements,omitempty"`
	Links               *NavigationResponseElementsItemLinks       `json:"links,omitempty"`
	NavigationDepth     *int                                       `json:"navigation_depth,omitempty"`
	NavigationID        *int                                       `json:"navigation_id,omitempty"`
	NavigationParentID  *int                                       `json:"navigation_parent_id,omitempty"`
	NavigationType      *string                                    `json:"navigation_type,omitempty"`
	Permissions         *NavigationResponseElementsItemPermissions `json:"permissions,omitempty"`
}

type NavigationResponseElementsItemLinks struct {
	Detail        *string `json:"detail,omitempty"`
	Permalink     *string `json:"permalink,omitempty"`
	SubCategories *string `json:"sub-categories,omitempty"`
	SubElements   *string `json:"sub-elements,omitempty"`
	SubForums     *string `json:"sub-forums,omitempty"`
}

type NavigationResponseElementsItemPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type NavigationResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type OnlineResponse struct {
	SystemInfo *RespSystemInfo           `json:"system_info,omitempty"`
	Users      []OnlineResponseUsersItem `json:"users,omitempty"`
}

type OnlineResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type OnlineResponseUsersItem struct {
	AvatarDate          *int                               `json:"avatar_date,omitempty"`
	BackgroundDate      *int                               `json:"background_date,omitempty"`
	ContestCount        *int                               `json:"contest_count,omitempty"`
	CustomTitle         *string                            `json:"custom_title,omitempty"`
	DisplayBannerID     *int                               `json:"display_banner_id,omitempty"`
	DisplayIconGroupID  *int                               `json:"display_icon_group_id,omitempty"`
	DisplayStyleGroupID *int                               `json:"display_style_group_id,omitempty"`
	IsAdmin             *bool                              `json:"is_admin,omitempty"`
	IsBanned            *bool                              `json:"is_banned,omitempty"`
	IsModerator         *bool                              `json:"is_moderator,omitempty"`
	IsStaff             *bool                              `json:"is_staff,omitempty"`
	LastActivity        *int                               `json:"last_activity,omitempty"`
	Like2Count          *int                               `json:"like2_count,omitempty"`
	LikeCount           *int                               `json:"like_count,omitempty"`
	MessageCount        *int                               `json:"message_count,omitempty"`
	RegisterDate        *int                               `json:"register_date,omitempty"`
	Rendered            *OnlineResponseUsersItemRendered   `json:"rendered,omitempty"`
	ShortLink           *string                            `json:"short_link,omitempty"`
	TrophyPoints        *int                               `json:"trophy_points,omitempty"`
	UniqBanner          *OnlineResponseUsersItemUniqBanner `json:"uniq_banner,omitempty"`
	UniqUsernameCSS     *string                            `json:"uniq_username_css,omitempty"`
	UserID              *int                               `json:"user_id,omitempty"`
	Username            *string                            `json:"username,omitempty"`
}

type OnlineResponseUsersItemRendered struct {
	Avatars  *OnlineResponseUsersItemRenderedAvatars `json:"avatars,omitempty"`
	Link     *string                                 `json:"link,omitempty"`
	Username *string                                 `json:"username,omitempty"`
}

type OnlineResponseUsersItemRenderedAvatars struct {
	L *string `json:"l,omitempty"`
	M *string `json:"m,omitempty"`
	S *string `json:"s,omitempty"`
}

type OnlineResponseUsersItemUniqBanner struct {
	BannerCSS    *string `json:"banner_css,omitempty"`
	BannerIcon   *string `json:"banner_icon,omitempty"`
	BannerText   *string `json:"banner_text,omitempty"`
	UsernameIcon *string `json:"username_icon,omitempty"`
}

type PopularResponse struct {
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
	Tags       interface{}     `json:"tags,omitempty"`
}

type PopularResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type PostMessageParams struct {
	// ID of the message being replied to.
	ReplyMessageID *int `json:"reply_message_id,omitempty"`
}

type PostMessageResponse struct {
	Message    *RespChatboxMessageModel `json:"message,omitempty"`
	SystemInfo *RespSystemInfo          `json:"system_info,omitempty"`
}

type PostMessageResponseMessage struct {
	CanReport   *bool                           `json:"can_report,omitempty"`
	Date        *int                            `json:"date,omitempty"`
	IsDeleted   *bool                           `json:"is_deleted,omitempty"`
	Message     *string                         `json:"message,omitempty"`
	MessageJson *string                         `json:"messageJson,omitempty"`
	MessageRaw  *string                         `json:"messageRaw,omitempty"`
	MessageID   *int                            `json:"message_id,omitempty"`
	Room        *PostMessageResponseMessageRoom `json:"room,omitempty"`
	User        *PostMessageResponseMessageUser `json:"user,omitempty"`
}

type PostMessageResponseMessageRoom struct {
	CanReport *bool   `json:"can_report,omitempty"`
	Eng       *bool   `json:"eng,omitempty"`
	Market    *bool   `json:"market,omitempty"`
	RoomID    *int    `json:"room_id,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type PostMessageResponseMessageUser struct {
	AvatarDate          *int                                      `json:"avatar_date,omitempty"`
	BackgroundDate      *int                                      `json:"background_date,omitempty"`
	ContestCount        *int                                      `json:"contest_count,omitempty"`
	CustomTitle         *string                                   `json:"custom_title,omitempty"`
	DisplayBannerID     *int                                      `json:"display_banner_id,omitempty"`
	DisplayIconGroupID  *int                                      `json:"display_icon_group_id,omitempty"`
	DisplayStyleGroupID *int                                      `json:"display_style_group_id,omitempty"`
	IsAdmin             *bool                                     `json:"is_admin,omitempty"`
	IsBanned            *bool                                     `json:"is_banned,omitempty"`
	IsModerator         *bool                                     `json:"is_moderator,omitempty"`
	IsStaff             *bool                                     `json:"is_staff,omitempty"`
	LastActivity        *int                                      `json:"last_activity,omitempty"`
	Like2Count          *int                                      `json:"like2_count,omitempty"`
	LikeCount           *int                                      `json:"like_count,omitempty"`
	MessageCount        *int                                      `json:"message_count,omitempty"`
	RegisterDate        *int                                      `json:"register_date,omitempty"`
	Rendered            *PostMessageResponseMessageUserRendered   `json:"rendered,omitempty"`
	ShortLink           *string                                   `json:"short_link,omitempty"`
	TrophyPoints        *int                                      `json:"trophy_points,omitempty"`
	UniqBanner          *PostMessageResponseMessageUserUniqBanner `json:"uniq_banner,omitempty"`
	UniqUsernameCSS     *string                                   `json:"uniq_username_css,omitempty"`
	UserID              *int                                      `json:"user_id,omitempty"`
	Username            *string                                   `json:"username,omitempty"`
}

type PostMessageResponseMessageUserRendered struct {
	Avatars  *PostMessageResponseMessageUserRenderedAvatars `json:"avatars,omitempty"`
	Link     *string                                        `json:"link,omitempty"`
	Username *string                                        `json:"username,omitempty"`
}

type PostMessageResponseMessageUserRenderedAvatars struct {
	L *string `json:"l,omitempty"`
	M *string `json:"m,omitempty"`
	S *string `json:"s,omitempty"`
}

type PostMessageResponseMessageUserUniqBanner struct {
	BannerCSS    *string `json:"banner_css,omitempty"`
	BannerIcon   *string `json:"banner_icon,omitempty"`
	BannerText   *string `json:"banner_text,omitempty"`
	UsernameIcon *string `json:"username_icon,omitempty"`
}

type PostMessageResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type PostsParams struct {
	// The time in milliseconds (e.g. 1767214800) before last content date.
	Before *int `json:"before,omitempty"`
	// Number of post data to be returned.
	DataLimit *int `json:"data_limit,omitempty"`
	// Id of the container forum to search for contents. Child forums of the specified forum will be included in the search.
	ForumID *int `json:"forum_id,omitempty"`
	// Number of results in a page.
	Limit *int `json:"limit,omitempty"`
	// Page number of results.
	Page *int `json:"page,omitempty"`
	// Search query. Can be skipped if **user_id** is set.
	Q *string `json:"q,omitempty"`
	// Tag to search for tagged contents.
	Tag    *string     `json:"tag,omitempty"`
	UserID StringOrInt `json:"user_id,omitempty"`
}

type PostsResponse struct {
	Data       []PostsResponseDataItem `json:"data,omitempty"`
	DataTotal  *int                    `json:"data_total,omitempty"`
	Links      *PostsResponseLinks     `json:"links,omitempty"`
	SystemInfo *RespSystemInfo         `json:"system_info,omitempty"`
}

type PostsResponseDataItem struct {
	ContentID           interface{}                       `json:"content_id,omitempty"`
	ContentType         *string                           `json:"content_type,omitempty"`
	CreatorUserID       *int                              `json:"creator_user_id,omitempty"`
	CreatorUsername     *string                           `json:"creator_username,omitempty"`
	CreatorUsernameHTML *string                           `json:"creator_username_html,omitempty"`
	FirstPost           *PostsResponseDataItemFirstPost   `json:"first_post,omitempty"`
	Forum               *PostsResponseDataItemForum       `json:"forum,omitempty"`
	ForumID             *int                              `json:"forum_id,omitempty"`
	Links               *PostsResponseDataItemLinks       `json:"links,omitempty"`
	Permissions         *PostsResponseDataItemPermissions `json:"permissions,omitempty"`
	ThreadCreateDate    *int                              `json:"thread_create_date,omitempty"`
	ThreadID            *int                              `json:"thread_id,omitempty"`
	ThreadIsDeleted     *bool                             `json:"thread_is_deleted,omitempty"`
	ThreadIsFollowed    *bool                             `json:"thread_is_followed,omitempty"`
	ThreadIsPublished   *bool                             `json:"thread_is_published,omitempty"`
	ThreadIsSticky      *bool                             `json:"thread_is_sticky,omitempty"`
	ThreadPostCount     *int                              `json:"thread_post_count,omitempty"`
	ThreadPrefixes      []interface{}                     `json:"thread_prefixes,omitempty"`
	ThreadTags          interface{}                       `json:"thread_tags,omitempty"`
	ThreadTitle         *string                           `json:"thread_title,omitempty"`
	ThreadUpdateDate    *int                              `json:"thread_update_date,omitempty"`
	ThreadViewCount     *int                              `json:"thread_view_count,omitempty"`
	UserIsIgnored       *bool                             `json:"user_is_ignored,omitempty"`
}

type PostsResponseDataItemFirstPost struct {
	Links               *PostsResponseDataItemFirstPostLinks       `json:"links,omitempty"`
	Permissions         *PostsResponseDataItemFirstPostPermissions `json:"permissions,omitempty"`
	PostAttachmentCount *int                                       `json:"post_attachment_count,omitempty"`
	PostBody            *string                                    `json:"post_body,omitempty"`
	PostBodyHTML        *string                                    `json:"post_body_html,omitempty"`
	PostBodyPlainText   *string                                    `json:"post_body_plain_text,omitempty"`
	PostCreateDate      *int                                       `json:"post_create_date,omitempty"`
	PostID              *int                                       `json:"post_id,omitempty"`
	PostIsDeleted       *bool                                      `json:"post_is_deleted,omitempty"`
	PostIsFirstPost     *bool                                      `json:"post_is_first_post,omitempty"`
	PostIsPublished     *bool                                      `json:"post_is_published,omitempty"`
	PostLikeCount       *int                                       `json:"post_like_count,omitempty"`
	PostUpdateDate      *int                                       `json:"post_update_date,omitempty"`
	PosterUserID        *int                                       `json:"poster_user_id,omitempty"`
	PosterUsername      *string                                    `json:"poster_username,omitempty"`
	PosterUsernameHTML  *string                                    `json:"poster_username_html,omitempty"`
	Signature           *string                                    `json:"signature,omitempty"`
	SignatureHTML       *string                                    `json:"signature_html,omitempty"`
	SignaturePlainText  *string                                    `json:"signature_plain_text,omitempty"`
	ThreadID            *int                                       `json:"thread_id,omitempty"`
	UserIsIgnored       *bool                                      `json:"user_is_ignored,omitempty"`
}

type PostsResponseDataItemFirstPostLinks struct {
	Attachments  *string `json:"attachments,omitempty"`
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type PostsResponseDataItemFirstPostPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Like             *bool `json:"like,omitempty"`
	Reply            *bool `json:"reply,omitempty"`
	Report           *bool `json:"report,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type PostsResponseDataItemForum struct {
	ForumDescription       *string                                `json:"forum_description,omitempty"`
	ForumID                *int                                   `json:"forum_id,omitempty"`
	ForumIsFollowed        *bool                                  `json:"forum_is_followed,omitempty"`
	ForumPostCount         *int                                   `json:"forum_post_count,omitempty"`
	ForumPrefixes          []interface{}                          `json:"forum_prefixes,omitempty"`
	ForumThreadCount       *int                                   `json:"forum_thread_count,omitempty"`
	ForumTitle             *string                                `json:"forum_title,omitempty"`
	Links                  *PostsResponseDataItemForumLinks       `json:"links,omitempty"`
	Permissions            *PostsResponseDataItemForumPermissions `json:"permissions,omitempty"`
	ThreadDefaultPrefixID  *int                                   `json:"thread_default_prefix_id,omitempty"`
	ThreadPrefixIsRequired *bool                                  `json:"thread_prefix_is_required,omitempty"`
}

type PostsResponseDataItemForumLinks struct {
	Detail        *string `json:"detail,omitempty"`
	Followers     *string `json:"followers,omitempty"`
	Permalink     *string `json:"permalink,omitempty"`
	SubCategories *string `json:"sub-categories,omitempty"`
	SubForums     *string `json:"sub-forums,omitempty"`
	Threads       *string `json:"threads,omitempty"`
}

type PostsResponseDataItemForumPermissions struct {
	CreateThread     *bool `json:"create_thread,omitempty"`
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	TagThread        *bool `json:"tag_thread,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type PostsResponseDataItemLinks struct {
	Detail            *string `json:"detail,omitempty"`
	FirstPost         *string `json:"first_post,omitempty"`
	FirstPoster       *string `json:"first_poster,omitempty"`
	FirstPosterAvatar *string `json:"first_poster_avatar,omitempty"`
	Followers         *string `json:"followers,omitempty"`
	Forum             *string `json:"forum,omitempty"`
	LastPost          *string `json:"last_post,omitempty"`
	Permalink         *string `json:"permalink,omitempty"`
	Posts             *string `json:"posts,omitempty"`
}

type PostsResponseDataItemPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	Post             *bool `json:"post,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type PostsResponseLinks struct {
	Next  *string `json:"next,omitempty"`
	Page  *int    `json:"page,omitempty"`
	Pages *int    `json:"pages,omitempty"`
}

type PostsResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ProfilePostsParams struct {
	// The time in milliseconds (e.g. 1767214800) before last content date.
	Before *int `json:"before,omitempty"`
	// Number of results in a page.
	Limit *int `json:"limit,omitempty"`
	// Page number of results.
	Page *int `json:"page,omitempty"`
	// Search query. Can be skipped if **user_id** is set.
	Q *string `json:"q,omitempty"`
	// User ID to filter profile posts.
	UserID *int `json:"user_id,omitempty"`
}

type ProfilePostsResponse struct {
	Data       []ProfilePostsResponseDataItem `json:"data,omitempty"`
	DataTotal  *int                           `json:"data_total,omitempty"`
	Links      *ProfilePostsResponseLinks     `json:"links,omitempty"`
	SystemInfo *RespSystemInfo                `json:"system_info,omitempty"`
}

type ProfilePostsResponseDataItem struct {
	ContentID          interface{}                              `json:"content_id,omitempty"`
	ContentType        *string                                  `json:"content_type,omitempty"`
	Links              *ProfilePostsResponseDataItemLinks       `json:"links,omitempty"`
	Permissions        *ProfilePostsResponseDataItemPermissions `json:"permissions,omitempty"`
	PostBody           *string                                  `json:"post_body,omitempty"`
	PostCommentCount   *int                                     `json:"post_comment_count,omitempty"`
	PostCreateDate     *int                                     `json:"post_create_date,omitempty"`
	PostIsDeleted      *bool                                    `json:"post_is_deleted,omitempty"`
	PostIsPublished    *bool                                    `json:"post_is_published,omitempty"`
	PostLikeCount      *int                                     `json:"post_like_count,omitempty"`
	PosterUserID       *int                                     `json:"poster_user_id,omitempty"`
	PosterUsername     *string                                  `json:"poster_username,omitempty"`
	PosterUsernameHTML *string                                  `json:"poster_username_html,omitempty"`
	ProfilePostID      *int                                     `json:"profile_post_id,omitempty"`
	TimelineUser       *RespUserModel                           `json:"timeline_user,omitempty"`
	TimelineUserID     *int                                     `json:"timeline_user_id,omitempty"`
	TimelineUsername   *string                                  `json:"timeline_username,omitempty"`
	UserIsIgnored      *bool                                    `json:"user_is_ignored,omitempty"`
}

type ProfilePostsResponseDataItemLinks struct {
	Comments     *string `json:"comments,omitempty"`
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Timeline     *string `json:"timeline,omitempty"`
	TimelineUser *string `json:"timeline_user,omitempty"`
}

type ProfilePostsResponseDataItemPermissions struct {
	Comment *bool `json:"comment,omitempty"`
	Delete  *bool `json:"delete,omitempty"`
	Edit    *bool `json:"edit,omitempty"`
	Like    *bool `json:"like,omitempty"`
	Report  *bool `json:"report,omitempty"`
	View    *bool `json:"view,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUser struct {
	Balance                     *string                                                                   `json:"balance,omitempty"`
	Banner                      *string                                                                   `json:"banner,omitempty"`
	Birthday                    *ProfilePostsResponseDataItemTimelineUserBirthday                         `json:"birthday,omitempty"`
	ContestCount                *int                                                                      `json:"contest_count,omitempty"`
	ConvWelcomeMessage          *string                                                                   `json:"conv_welcome_message,omitempty"`
	CuratorTitles               []string                                                                  `json:"curator_titles,omitempty"`
	Currency                    *string                                                                   `json:"currency,omitempty"`
	CustomTitle                 *string                                                                   `json:"custom_title,omitempty"`
	DisplayBannerID             *int                                                                      `json:"display_banner_id,omitempty"`
	DisplayIconGroupID          *int                                                                      `json:"display_icon_group_id,omitempty"`
	EditPermissions             *ProfilePostsResponseDataItemTimelineUserEditPermissions                  `json:"edit_permissions,omitempty"`
	Fields                      []ProfilePostsResponseDataItemTimelineUserFieldsItem                      `json:"fields,omitempty"`
	Hold                        *string                                                                   `json:"hold,omitempty"`
	IsBanned                    *int                                                                      `json:"is_banned,omitempty"`
	Links                       *ProfilePostsResponseDataItemTimelineUserLinks                            `json:"links,omitempty"`
	Permissions                 *ProfilePostsResponseDataItemTimelineUserPermissions                      `json:"permissions,omitempty"`
	SecretAnswerFirstLetter     *string                                                                   `json:"secret_answer_first_letter,omitempty"`
	SecretAnswerRendered        *string                                                                   `json:"secret_answer_rendered,omitempty"`
	SelfPermissions             *ProfilePostsResponseDataItemTimelineUserSelfPermissions                  `json:"self_permissions,omitempty"`
	ShortLink                   *string                                                                   `json:"short_link,omitempty"`
	TrophyCount                 *int                                                                      `json:"trophy_count,omitempty"`
	UserDeposit                 *int                                                                      `json:"user_deposit,omitempty"`
	UserEmail                   *string                                                                   `json:"user_email,omitempty"`
	UserExternalAuthentications []ProfilePostsResponseDataItemTimelineUserUserExternalAuthenticationsItem `json:"user_external_authentications,omitempty"`
	UserFollowers               *ProfilePostsResponseDataItemTimelineUserUserFollowers                    `json:"user_followers,omitempty"`
	UserFollowing               *ProfilePostsResponseDataItemTimelineUserUserFollowing                    `json:"user_following,omitempty"`
	UserGroupID                 *int                                                                      `json:"user_group_id,omitempty"`
	UserGroups                  []ProfilePostsResponseDataItemTimelineUserUserGroupsItem                  `json:"user_groups,omitempty"`
	UserID                      *int                                                                      `json:"user_id,omitempty"`
	UserIsFollowed              *bool                                                                     `json:"user_is_followed,omitempty"`
	UserIsIgnored               *bool                                                                     `json:"user_is_ignored,omitempty"`
	UserIsValid                 *bool                                                                     `json:"user_is_valid,omitempty"`
	UserIsVerified              *bool                                                                     `json:"user_is_verified,omitempty"`
	UserIsVisitor               *bool                                                                     `json:"user_is_visitor,omitempty"`
	UserLastSeenDate            *int                                                                      `json:"user_last_seen_date,omitempty"`
	UserLike2Count              *int                                                                      `json:"user_like2_count,omitempty"`
	UserLikeCount               *int                                                                      `json:"user_like_count,omitempty"`
	UserMessageCount            *int                                                                      `json:"user_message_count,omitempty"`
	UserRegisterDate            *int                                                                      `json:"user_register_date,omitempty"`
	UserTimezoneOffset          *int                                                                      `json:"user_timezone_offset,omitempty"`
	UserTitle                   *string                                                                   `json:"user_title,omitempty"`
	UserUnreadConversationCount *int                                                                      `json:"user_unread_conversation_count,omitempty"`
	UserUnreadNotificationCount *int                                                                      `json:"user_unread_notification_count,omitempty"`
	Username                    *string                                                                   `json:"username,omitempty"`
	UsernameHTML                *string                                                                   `json:"username_html,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUserBirthday struct {
	Age       *int                                                       `json:"age,omitempty"`
	Format    *string                                                    `json:"format,omitempty"`
	TimeStamp *ProfilePostsResponseDataItemTimelineUserBirthdayTimeStamp `json:"timeStamp,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUserBirthdayTimeStamp struct {
	Date         *string `json:"date,omitempty"`
	Timezone     *string `json:"timezone,omitempty"`
	TimezoneType *int    `json:"timezone_type,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUserEditPermissions struct {
	Fields            *bool `json:"fields,omitempty"`
	HideUsernameLogs  *bool `json:"hide_username_logs,omitempty"`
	Password          *bool `json:"password,omitempty"`
	PrimaryGroupID    *bool `json:"primary_group_id,omitempty"`
	SecondaryGroupIds *bool `json:"secondary_group_ids,omitempty"`
	ShortLink         *bool `json:"short_link,omitempty"`
	UserDobDay        *bool `json:"user_dob_day,omitempty"`
	UserDobMonth      *bool `json:"user_dob_month,omitempty"`
	UserDobYear       *bool `json:"user_dob_year,omitempty"`
	UserEmail         *bool `json:"user_email,omitempty"`
	UserTitle         *bool `json:"user_title,omitempty"`
	Username          *bool `json:"username,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUserFieldsItem struct {
	Choices       []ProfilePostsResponseDataItemTimelineUserFieldsItemChoicesItem `json:"choices,omitempty"`
	Description   *string                                                         `json:"description,omitempty"`
	ID            *string                                                         `json:"id,omitempty"`
	IsMultiChoice *bool                                                           `json:"is_multi_choice,omitempty"`
	IsRequired    *bool                                                           `json:"is_required,omitempty"`
	Position      *string                                                         `json:"position,omitempty"`
	Title         *string                                                         `json:"title,omitempty"`
	Value         *string                                                         `json:"value,omitempty"`
	Values        interface{}                                                     `json:"values,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUserFieldsItemChoicesItem struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUserLinks struct {
	Avatar      *string `json:"avatar,omitempty"`
	AvatarBig   *string `json:"avatar_big,omitempty"`
	AvatarSmall *string `json:"avatar_small,omitempty"`
	BackgroundL *string `json:"background_l,omitempty"`
	BackgroundM *string `json:"background_m,omitempty"`
	Detail      *string `json:"detail,omitempty"`
	Followers   *string `json:"followers,omitempty"`
	Followings  *string `json:"followings,omitempty"`
	Ignore      *string `json:"ignore,omitempty"`
	Permalink   *string `json:"permalink,omitempty"`
	Status      *string `json:"status,omitempty"`
	Timeline    *string `json:"timeline,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUserPermissions struct {
	Edit        *bool `json:"edit,omitempty"`
	Follow      *bool `json:"follow,omitempty"`
	Ignore      *bool `json:"ignore,omitempty"`
	ProfilePost *bool `json:"profile_post,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUserSelfPermissions struct {
	CreateConversation *bool `json:"create_conversation,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUserUserExternalAuthenticationsItem struct {
	Provider    *string `json:"provider,omitempty"`
	ProviderKey *string `json:"provider_key,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUserUserFollowers struct {
	Count *int                                                             `json:"count,omitempty"`
	Users []ProfilePostsResponseDataItemTimelineUserUserFollowersUsersItem `json:"users,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUserUserFollowersUsersItem struct {
	Avatar       *string `json:"avatar,omitempty"`
	UserID       *int    `json:"user_id,omitempty"`
	Username     *string `json:"username,omitempty"`
	UsernameHTML *string `json:"username_html,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUserUserFollowing struct {
	Count *int                                                             `json:"count,omitempty"`
	Users []ProfilePostsResponseDataItemTimelineUserUserFollowingUsersItem `json:"users,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUserUserFollowingUsersItem struct {
	Avatar       *string `json:"avatar,omitempty"`
	UserID       *int    `json:"user_id,omitempty"`
	Username     *string `json:"username,omitempty"`
	UsernameHTML *string `json:"username_html,omitempty"`
}

type ProfilePostsResponseDataItemTimelineUserUserGroupsItem struct {
	DisplayBannerSelectable *bool   `json:"display_banner_selectable,omitempty"`
	DisplayGroupSelectable  *bool   `json:"display_group_selectable,omitempty"`
	DisplayIconSelectable   *bool   `json:"display_icon_selectable,omitempty"`
	IsPrimaryGroup          *bool   `json:"is_primary_group,omitempty"`
	UserGroupBannerCSSClass *string `json:"user_group_banner_css_class,omitempty"`
	UserGroupBannerText     *string `json:"user_group_banner_text,omitempty"`
	UserGroupBannerTextEn   *string `json:"user_group_banner_text_en,omitempty"`
	UserGroupIconClass      *string `json:"user_group_icon_class,omitempty"`
	UserGroupID             *int    `json:"user_group_id,omitempty"`
	UserGroupTitle          *string `json:"user_group_title,omitempty"`
	UserGroupTitleEn        *string `json:"user_group_title_en,omitempty"`
}

type ProfilePostsResponseLinks struct {
	Next  *string `json:"next,omitempty"`
	Page  *int    `json:"page,omitempty"`
	Pages *int    `json:"pages,omitempty"`
}

type ProfilePostsResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ReadAllResponse struct {
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type ReadAllResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ReadParams struct {
	// If notification_id is omitted, it's mark all existing notifications as read.
	NotificationID *int `json:"notification_id,omitempty"`
}

type RecentParams struct {
	// Maximum number of days to search for threads.
	Days *int `json:"days,omitempty"`
	// Maximum number of result threads. The limit may get decreased if the value is too large.
	Limit *int `json:"limit,omitempty"`
	// Id of the container forum to search for threads. Child forums of the specified forum will be included in the search.
	ForumID *int `json:"forum_id,omitempty"`
	// Number of thread data to be returned. Default value is 20.
	DataLimit *int `json:"data_limit,omitempty"`
}

type RecentResponse struct {
	Data       []RecentResponseDataItem    `json:"data,omitempty"`
	SystemInfo *RespSystemInfo             `json:"system_info,omitempty"`
	Threads    []RecentResponseThreadsItem `json:"threads,omitempty"`
}

type RecentResponseDataItem struct {
	ContentID           interface{}                        `json:"content_id,omitempty"`
	ContentType         *string                            `json:"content_type,omitempty"`
	CreatorUserID       *int                               `json:"creator_user_id,omitempty"`
	CreatorUsername     *string                            `json:"creator_username,omitempty"`
	CreatorUsernameHTML *string                            `json:"creator_username_html,omitempty"`
	FirstPost           *RecentResponseDataItemFirstPost   `json:"first_post,omitempty"`
	Forum               *RecentResponseDataItemForum       `json:"forum,omitempty"`
	ForumID             *int                               `json:"forum_id,omitempty"`
	Links               *RecentResponseDataItemLinks       `json:"links,omitempty"`
	Permissions         *RecentResponseDataItemPermissions `json:"permissions,omitempty"`
	ThreadCreateDate    *int                               `json:"thread_create_date,omitempty"`
	ThreadID            *int                               `json:"thread_id,omitempty"`
	ThreadIsDeleted     *bool                              `json:"thread_is_deleted,omitempty"`
	ThreadIsFollowed    *bool                              `json:"thread_is_followed,omitempty"`
	ThreadIsPublished   *bool                              `json:"thread_is_published,omitempty"`
	ThreadIsSticky      *bool                              `json:"thread_is_sticky,omitempty"`
	ThreadPostCount     *int                               `json:"thread_post_count,omitempty"`
	ThreadPrefixes      []interface{}                      `json:"thread_prefixes,omitempty"`
	ThreadTags          interface{}                        `json:"thread_tags,omitempty"`
	ThreadTitle         *string                            `json:"thread_title,omitempty"`
	ThreadUpdateDate    *int                               `json:"thread_update_date,omitempty"`
	ThreadViewCount     *int                               `json:"thread_view_count,omitempty"`
	UserIsIgnored       *bool                              `json:"user_is_ignored,omitempty"`
}

type RecentResponseDataItemFirstPost struct {
	Links               *RecentResponseDataItemFirstPostLinks       `json:"links,omitempty"`
	Permissions         *RecentResponseDataItemFirstPostPermissions `json:"permissions,omitempty"`
	PostAttachmentCount *int                                        `json:"post_attachment_count,omitempty"`
	PostBody            *string                                     `json:"post_body,omitempty"`
	PostBodyHTML        *string                                     `json:"post_body_html,omitempty"`
	PostBodyPlainText   *string                                     `json:"post_body_plain_text,omitempty"`
	PostCreateDate      *int                                        `json:"post_create_date,omitempty"`
	PostID              *int                                        `json:"post_id,omitempty"`
	PostIsDeleted       *bool                                       `json:"post_is_deleted,omitempty"`
	PostIsFirstPost     *bool                                       `json:"post_is_first_post,omitempty"`
	PostIsPublished     *bool                                       `json:"post_is_published,omitempty"`
	PostLikeCount       *int                                        `json:"post_like_count,omitempty"`
	PostUpdateDate      *int                                        `json:"post_update_date,omitempty"`
	PosterUserID        *int                                        `json:"poster_user_id,omitempty"`
	PosterUsername      *string                                     `json:"poster_username,omitempty"`
	PosterUsernameHTML  *string                                     `json:"poster_username_html,omitempty"`
	Signature           *string                                     `json:"signature,omitempty"`
	SignatureHTML       *string                                     `json:"signature_html,omitempty"`
	SignaturePlainText  *string                                     `json:"signature_plain_text,omitempty"`
	ThreadID            *int                                        `json:"thread_id,omitempty"`
	UserIsIgnored       *bool                                       `json:"user_is_ignored,omitempty"`
}

type RecentResponseDataItemFirstPostLinks struct {
	Attachments  *string `json:"attachments,omitempty"`
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type RecentResponseDataItemFirstPostPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Like             *bool `json:"like,omitempty"`
	Reply            *bool `json:"reply,omitempty"`
	Report           *bool `json:"report,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type RecentResponseDataItemForum struct {
	ForumDescription       *string                                 `json:"forum_description,omitempty"`
	ForumID                *int                                    `json:"forum_id,omitempty"`
	ForumIsFollowed        *bool                                   `json:"forum_is_followed,omitempty"`
	ForumPostCount         *int                                    `json:"forum_post_count,omitempty"`
	ForumPrefixes          []interface{}                           `json:"forum_prefixes,omitempty"`
	ForumThreadCount       *int                                    `json:"forum_thread_count,omitempty"`
	ForumTitle             *string                                 `json:"forum_title,omitempty"`
	Links                  *RecentResponseDataItemForumLinks       `json:"links,omitempty"`
	Permissions            *RecentResponseDataItemForumPermissions `json:"permissions,omitempty"`
	ThreadDefaultPrefixID  *int                                    `json:"thread_default_prefix_id,omitempty"`
	ThreadPrefixIsRequired *bool                                   `json:"thread_prefix_is_required,omitempty"`
}

type RecentResponseDataItemForumLinks struct {
	Detail        *string `json:"detail,omitempty"`
	Followers     *string `json:"followers,omitempty"`
	Permalink     *string `json:"permalink,omitempty"`
	SubCategories *string `json:"sub-categories,omitempty"`
	SubForums     *string `json:"sub-forums,omitempty"`
	Threads       *string `json:"threads,omitempty"`
}

type RecentResponseDataItemForumPermissions struct {
	CreateThread     *bool `json:"create_thread,omitempty"`
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	TagThread        *bool `json:"tag_thread,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type RecentResponseDataItemLinks struct {
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

type RecentResponseDataItemPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	Post             *bool `json:"post,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type RecentResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type RecentResponseThreadsItem struct {
	Contest             *RecentResponseThreadsItemContest      `json:"contest,omitempty"`
	CreatorUserID       *int                                   `json:"creator_user_id,omitempty"`
	CreatorUsername     *string                                `json:"creator_username,omitempty"`
	CreatorUsernameHTML *string                                `json:"creator_username_html,omitempty"`
	FirstPost           *RecentResponseThreadsItemFirstPost    `json:"first_post,omitempty"`
	ForumID             *int                                   `json:"forum_id,omitempty"`
	LastPost            *RecentResponseThreadsItemLastPost     `json:"last_post,omitempty"`
	Links               *RecentResponseThreadsItemLinks        `json:"links,omitempty"`
	NodeTitle           *string                                `json:"node_title,omitempty"`
	Permissions         *RecentResponseThreadsItemPermissions  `json:"permissions,omitempty"`
	Restrictions        *RecentResponseThreadsItemRestrictions `json:"restrictions,omitempty"`
	ThreadCreateDate    *int                                   `json:"thread_create_date,omitempty"`
	ThreadID            *int                                   `json:"thread_id,omitempty"`
	ThreadIsClosed      *bool                                  `json:"thread_is_closed,omitempty"`
	ThreadIsDeleted     *bool                                  `json:"thread_is_deleted,omitempty"`
	ThreadIsFollowed    *bool                                  `json:"thread_is_followed,omitempty"`
	ThreadIsPublished   *bool                                  `json:"thread_is_published,omitempty"`
	ThreadIsStarred     *bool                                  `json:"thread_is_starred,omitempty"`
	ThreadIsSticky      *bool                                  `json:"thread_is_sticky,omitempty"`
	ThreadPostCount     *int                                   `json:"thread_post_count,omitempty"`
	ThreadPrefixes      []interface{}                          `json:"thread_prefixes,omitempty"`
	ThreadTags          interface{}                            `json:"thread_tags,omitempty"`
	ThreadTitle         *string                                `json:"thread_title,omitempty"`
	ThreadUpdateDate    *int                                   `json:"thread_update_date,omitempty"`
	ThreadViewCount     *int                                   `json:"thread_view_count,omitempty"`
	UserIsIgnored       *bool                                  `json:"user_is_ignored,omitempty"`
}

type RecentResponseThreadsItemContest struct {
	AlreadyParticipate    *bool                                        `json:"already_participate,omitempty"`
	ChanceToWin           *float64                                     `json:"chance_to_win,omitempty"`
	CountWinners          *int                                         `json:"count_winners,omitempty"`
	FinishDate            *int                                         `json:"finish_date,omitempty"`
	IsFinished            *int                                         `json:"is_finished,omitempty"`
	IsMoneyPlaces         *int                                         `json:"is_money_places,omitempty"`
	NeededMembers         *int                                         `json:"needed_members,omitempty"`
	NowCountMembers       *int                                         `json:"now_count_members,omitempty"`
	Permissions           *RecentResponseThreadsItemContestPermissions `json:"permissions,omitempty"`
	PrizeData             *int                                         `json:"prize_data,omitempty"`
	PrizeType             *string                                      `json:"prize_type,omitempty"`
	PrizeTypePhrase       *string                                      `json:"prize_type_phrase,omitempty"`
	RequireLikeCount      *int                                         `json:"require_like_count,omitempty"`
	RequireTotalLikeCount *int                                         `json:"require_total_like_count,omitempty"`
	Type_                 *string                                      `json:"type,omitempty"`
	Winners               []int                                        `json:"winners,omitempty"`
}

type RecentResponseThreadsItemContestPermissions struct {
	CanFinish           *bool   `json:"can_finish,omitempty"`
	CanParticipate      *bool   `json:"can_participate,omitempty"`
	CanParticipateError *string `json:"can_participate_error,omitempty"`
	CanViewUserList     *bool   `json:"can_view_user_list,omitempty"`
}

type RecentResponseThreadsItemFirstPost struct {
	Links              *RecentResponseThreadsItemFirstPostLinks       `json:"links,omitempty"`
	Permissions        *RecentResponseThreadsItemFirstPostPermissions `json:"permissions,omitempty"`
	PostBody           *string                                        `json:"post_body,omitempty"`
	PostBodyHTML       *string                                        `json:"post_body_html,omitempty"`
	PostBodyPlainText  *string                                        `json:"post_body_plain_text,omitempty"`
	PostCreateDate     *int                                           `json:"post_create_date,omitempty"`
	PostID             *int                                           `json:"post_id,omitempty"`
	PostIsDeleted      *bool                                          `json:"post_is_deleted,omitempty"`
	PostIsFirstPost    *bool                                          `json:"post_is_first_post,omitempty"`
	PostIsLiked        *bool                                          `json:"post_is_liked,omitempty"`
	PostIsPublished    *bool                                          `json:"post_is_published,omitempty"`
	PostLikeCount      *int                                           `json:"post_like_count,omitempty"`
	PostUpdateDate     *int                                           `json:"post_update_date,omitempty"`
	PosterUserID       *int                                           `json:"poster_user_id,omitempty"`
	PosterUsername     *string                                        `json:"poster_username,omitempty"`
	PosterUsernameHTML *string                                        `json:"poster_username_html,omitempty"`
	Signature          *string                                        `json:"signature,omitempty"`
	SignatureHTML      *string                                        `json:"signature_html,omitempty"`
	SignaturePlainText *string                                        `json:"signature_plain_text,omitempty"`
	ThreadID           *int                                           `json:"thread_id,omitempty"`
	ThreadIsDeleted    *bool                                          `json:"thread_is_deleted,omitempty"`
	UserIsIgnored      *bool                                          `json:"user_is_ignored,omitempty"`
}

type RecentResponseThreadsItemFirstPostLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type RecentResponseThreadsItemFirstPostPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	Like   *bool `json:"like,omitempty"`
	Reply  *bool `json:"reply,omitempty"`
	Report *bool `json:"report,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type RecentResponseThreadsItemLastPost struct {
	Links              *RecentResponseThreadsItemLastPostLinks       `json:"links,omitempty"`
	Permissions        *RecentResponseThreadsItemLastPostPermissions `json:"permissions,omitempty"`
	PostBody           *string                                       `json:"post_body,omitempty"`
	PostBodyHTML       *string                                       `json:"post_body_html,omitempty"`
	PostBodyPlainText  *string                                       `json:"post_body_plain_text,omitempty"`
	PostCreateDate     *int                                          `json:"post_create_date,omitempty"`
	PostID             *int                                          `json:"post_id,omitempty"`
	PostIsDeleted      *bool                                         `json:"post_is_deleted,omitempty"`
	PostIsFirstPost    *bool                                         `json:"post_is_first_post,omitempty"`
	PostIsLiked        *bool                                         `json:"post_is_liked,omitempty"`
	PostIsPublished    *bool                                         `json:"post_is_published,omitempty"`
	PostLikeCount      *int                                          `json:"post_like_count,omitempty"`
	PostUpdateDate     *int                                          `json:"post_update_date,omitempty"`
	PosterUserID       *int                                          `json:"poster_user_id,omitempty"`
	PosterUsername     *string                                       `json:"poster_username,omitempty"`
	PosterUsernameHTML *string                                       `json:"poster_username_html,omitempty"`
	Signature          *string                                       `json:"signature,omitempty"`
	SignatureHTML      *string                                       `json:"signature_html,omitempty"`
	SignaturePlainText *string                                       `json:"signature_plain_text,omitempty"`
	ThreadID           *int                                          `json:"thread_id,omitempty"`
	ThreadIsDeleted    *bool                                         `json:"thread_is_deleted,omitempty"`
	UserIsIgnored      *bool                                         `json:"user_is_ignored,omitempty"`
}

type RecentResponseThreadsItemLastPostLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type RecentResponseThreadsItemLastPostPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	Like   *bool `json:"like,omitempty"`
	Reply  *bool `json:"reply,omitempty"`
	Report *bool `json:"report,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type RecentResponseThreadsItemLinks struct {
	Detail            *string `json:"detail,omitempty"`
	FirstPost         *string `json:"first_post,omitempty"`
	FirstPoster       *string `json:"first_poster,omitempty"`
	FirstPosterAvatar *string `json:"first_poster_avatar,omitempty"`
	Followers         *string `json:"followers,omitempty"`
	Forum             *string `json:"forum,omitempty"`
	LastPost          *string `json:"last_post,omitempty"`
	Permalink         *string `json:"permalink,omitempty"`
	Posts             *string `json:"posts,omitempty"`
}

type RecentResponseThreadsItemPermissions struct {
	Bump      *RecentResponseThreadsItemPermissionsBump `json:"bump,omitempty"`
	Delete    *bool                                     `json:"delete,omitempty"`
	Edit      *bool                                     `json:"edit,omitempty"`
	EditTags  *bool                                     `json:"edit_tags,omitempty"`
	EditTitle *bool                                     `json:"edit_title,omitempty"`
	Follow    *bool                                     `json:"follow,omitempty"`
	Post      *bool                                     `json:"post,omitempty"`
	View      *bool                                     `json:"view,omitempty"`
}

type RecentResponseThreadsItemPermissionsBump struct {
	AvailableCount    *int        `json:"available_count,omitempty"`
	Can               *bool       `json:"can,omitempty"`
	Error             interface{} `json:"error,omitempty"`
	NextAvailableTime interface{} `json:"next_available_time,omitempty"`
}

type RecentResponseThreadsItemRestrictions struct {
	MaxReplyCount *int `json:"max_reply_count,omitempty"`
	ReplyDelay    *int `json:"reply_delay,omitempty"`
}

type ReportReasonsResponse struct {
	Reasons    []string        `json:"reasons,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type ReportReasonsResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ResetResponse struct {
	Success     *bool           `json:"success,omitempty"`
	SystemInfo  *RespSystemInfo `json:"system_info,omitempty"`
	WaitingTime *string         `json:"waiting_time,omitempty"`
}

type ResetResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type RespChatboxMessageModel struct {
	CanReport   *bool                        `json:"can_report,omitempty"`
	Date        *int                         `json:"date,omitempty"`
	IsDeleted   *bool                        `json:"is_deleted,omitempty"`
	Message     *string                      `json:"message,omitempty"`
	MessageJson *string                      `json:"messageJson,omitempty"`
	MessageRaw  *string                      `json:"messageRaw,omitempty"`
	MessageID   *int                         `json:"message_id,omitempty"`
	Room        *RespChatboxMessageModelRoom `json:"room,omitempty"`
	User        *RespChatboxMessageModelUser `json:"user,omitempty"`
}

type RespChatboxMessageModelRoom struct {
	CanReport *bool   `json:"can_report,omitempty"`
	Eng       *bool   `json:"eng,omitempty"`
	Market    *bool   `json:"market,omitempty"`
	RoomID    *int    `json:"room_id,omitempty"`
	Title     *string `json:"title,omitempty"`
}

type RespChatboxMessageModelUser struct {
	AvatarDate          *int                                   `json:"avatar_date,omitempty"`
	BackgroundDate      *int                                   `json:"background_date,omitempty"`
	ContestCount        *int                                   `json:"contest_count,omitempty"`
	CustomTitle         *string                                `json:"custom_title,omitempty"`
	DisplayBannerID     *int                                   `json:"display_banner_id,omitempty"`
	DisplayIconGroupID  *int                                   `json:"display_icon_group_id,omitempty"`
	DisplayStyleGroupID *int                                   `json:"display_style_group_id,omitempty"`
	IsAdmin             *bool                                  `json:"is_admin,omitempty"`
	IsBanned            *bool                                  `json:"is_banned,omitempty"`
	IsModerator         *bool                                  `json:"is_moderator,omitempty"`
	IsStaff             *bool                                  `json:"is_staff,omitempty"`
	LastActivity        *int                                   `json:"last_activity,omitempty"`
	Like2Count          *int                                   `json:"like2_count,omitempty"`
	LikeCount           *int                                   `json:"like_count,omitempty"`
	MessageCount        *int                                   `json:"message_count,omitempty"`
	RegisterDate        *int                                   `json:"register_date,omitempty"`
	Rendered            *RespChatboxMessageModelUserRendered   `json:"rendered,omitempty"`
	ShortLink           *string                                `json:"short_link,omitempty"`
	TrophyPoints        *int                                   `json:"trophy_points,omitempty"`
	UniqBanner          *RespChatboxMessageModelUserUniqBanner `json:"uniq_banner,omitempty"`
	UniqUsernameCSS     *string                                `json:"uniq_username_css,omitempty"`
	UserID              *int                                   `json:"user_id,omitempty"`
	Username            *string                                `json:"username,omitempty"`
}

type RespChatboxMessageModelUserRendered struct {
	Avatars  *RespChatboxMessageModelUserRenderedAvatars `json:"avatars,omitempty"`
	Link     *string                                     `json:"link,omitempty"`
	Username *string                                     `json:"username,omitempty"`
}

type RespChatboxMessageModelUserRenderedAvatars struct {
	L *string `json:"l,omitempty"`
	M *string `json:"m,omitempty"`
	S *string `json:"s,omitempty"`
}

type RespChatboxMessageModelUserUniqBanner struct {
	BannerCSS    *string `json:"banner_css,omitempty"`
	BannerIcon   *string `json:"banner_icon,omitempty"`
	BannerText   *string `json:"banner_text,omitempty"`
	UsernameIcon *string `json:"username_icon,omitempty"`
}

type RespConversationMessageModel struct {
	ConversationID       *int                                     `json:"conversation_id,omitempty"`
	CreatorUserID        *int                                     `json:"creator_user_id,omitempty"`
	CreatorUsername      *string                                  `json:"creator_username,omitempty"`
	CreatorUsernameHTML  *string                                  `json:"creator_username_html,omitempty"`
	Links                *RespConversationMessageModelLinks       `json:"links,omitempty"`
	MessageBody          *string                                  `json:"message_body,omitempty"`
	MessageBodyHTML      *string                                  `json:"message_body_html,omitempty"`
	MessageBodyPlainText *string                                  `json:"message_body_plain_text,omitempty"`
	MessageCreateDate    *int                                     `json:"message_create_date,omitempty"`
	MessageEditDate      *int                                     `json:"message_edit_date,omitempty"`
	MessageID            *int                                     `json:"message_id,omitempty"`
	MessageIsSystem      *bool                                    `json:"message_is_system,omitempty"`
	MessageIsUnread      *int                                     `json:"message_is_unread,omitempty"`
	MessageNeedTranslate *bool                                    `json:"message_need_translate,omitempty"`
	Permissions          *RespConversationMessageModelPermissions `json:"permissions,omitempty"`
	UserIsIgnored        *bool                                    `json:"user_is_ignored,omitempty"`
}

type RespConversationMessageModelLinks struct {
	Conversation  *string `json:"conversation,omitempty"`
	Creator       *string `json:"creator,omitempty"`
	CreatorAvatar *string `json:"creator_avatar,omitempty"`
	Detail        *string `json:"detail,omitempty"`
}

type RespConversationMessageModelPermissions struct {
	Delete       *bool `json:"delete,omitempty"`
	Edit         *bool `json:"edit,omitempty"`
	StickUnstick *bool `json:"stick-unstick,omitempty"`
	View         *bool `json:"view,omitempty"`
}

type RespConversationModel struct {
	Alerts                   *int                                  `json:"alerts,omitempty"`
	ConversationCreateDate   *int                                  `json:"conversation_create_date,omitempty"`
	ConversationID           *int                                  `json:"conversation_id,omitempty"`
	ConversationIsDeleted    *bool                                 `json:"conversation_is_deleted,omitempty"`
	ConversationIsNew        *bool                                 `json:"conversation_is_new,omitempty"`
	ConversationIsOpen       *bool                                 `json:"conversation_is_open,omitempty"`
	ConversationLastReadDate *int                                  `json:"conversation_last_read_date,omitempty"`
	ConversationMessageCount *int                                  `json:"conversation_message_count,omitempty"`
	ConversationOnlineCount  *int                                  `json:"conversation_online_count,omitempty"`
	ConversationTitle        *string                               `json:"conversation_title,omitempty"`
	ConversationUpdateDate   *int                                  `json:"conversation_update_date,omitempty"`
	CreatorIsIgnored         *bool                                 `json:"creator_is_ignored,omitempty"`
	CreatorUserID            *int                                  `json:"creator_user_id,omitempty"`
	CreatorUsername          *string                               `json:"creator_username,omitempty"`
	CreatorUsernameHTML      *string                               `json:"creator_username_html,omitempty"`
	IsGroup                  *int                                  `json:"is_group,omitempty"`
	IsStarred                *int                                  `json:"is_starred,omitempty"`
	IsUnread                 *int                                  `json:"is_unread,omitempty"`
	Links                    *RespConversationModelLinks           `json:"links,omitempty"`
	Permissions              *RespConversationModelPermissions     `json:"permissions,omitempty"`
	Recipient                *RespConversationModelRecipient       `json:"recipient,omitempty"`
	Recipients               []RespConversationModelRecipientsItem `json:"recipients,omitempty"`
}

type RespConversationModelLinks struct {
	Avatar    *string `json:"avatar,omitempty"`
	Detail    *string `json:"detail,omitempty"`
	Messages  *string `json:"messages,omitempty"`
	Permalink *string `json:"permalink,omitempty"`
}

type RespConversationModelPermissions struct {
	EditOwnPost       *bool `json:"editOwnPost,omitempty"`
	Invite            *bool `json:"invite,omitempty"`
	Kick              *bool `json:"kick,omitempty"`
	ManageInviteLinks *bool `json:"manage_invite_links,omitempty"`
	Reply             *bool `json:"reply,omitempty"`
	StickyMessages    *bool `json:"stickyMessages,omitempty"`
	UploadAvatar      *bool `json:"upload_avatar,omitempty"`
	View              *bool `json:"view,omitempty"`
}

type RespConversationModelRecipient struct {
	Avatar          *string `json:"avatar,omitempty"`
	ContactsChanged *bool   `json:"contacts_changed,omitempty"`
	IsOnline        *bool   `json:"is_online,omitempty"`
	LastActivity    *int    `json:"last_activity,omitempty"`
	UserID          *int    `json:"user_id,omitempty"`
	Username        *string `json:"username,omitempty"`
	UsernameHTML    *string `json:"username_html,omitempty"`
}

type RespConversationModelRecipientsItem struct {
	Avatar          *string `json:"avatar,omitempty"`
	ContactsChanged *bool   `json:"contacts_changed,omitempty"`
	IsOnline        *bool   `json:"is_online,omitempty"`
	LastActivity    *int    `json:"last_activity,omitempty"`
	UserID          *int    `json:"user_id,omitempty"`
	Username        *string `json:"username,omitempty"`
	UsernameHTML    *string `json:"username_html,omitempty"`
}

type RespLinkModel struct {
	LinkDescription *string                   `json:"link_description,omitempty"`
	LinkID          *int                      `json:"link_id,omitempty"`
	LinkTitle       *string                   `json:"link_title,omitempty"`
	Links           *RespLinkModelLinks       `json:"links,omitempty"`
	Permissions     *RespLinkModelPermissions `json:"permissions,omitempty"`
}

type RespLinkModelLinks struct {
	Detail *string `json:"detail,omitempty"`
	Target *string `json:"target,omitempty"`
}

type RespLinkModelPermissions struct {
	View *bool `json:"view,omitempty"`
}

type RespNotificationModel struct {
	ContentAction          *string                     `json:"content_action,omitempty"`
	ContentID              interface{}                 `json:"content_id,omitempty"`
	ContentType            *string                     `json:"content_type,omitempty"`
	CreatorUserID          *int                        `json:"creator_user_id,omitempty"`
	CreatorUsername        *string                     `json:"creator_username,omitempty"`
	CreatorUsernameHTML    *string                     `json:"creator_username_html,omitempty"`
	Links                  *RespNotificationModelLinks `json:"links,omitempty"`
	NotificationCreateDate *int                        `json:"notification_create_date,omitempty"`
	NotificationHTML       *string                     `json:"notification_html,omitempty"`
	NotificationID         *int                        `json:"notification_id,omitempty"`
	NotificationIsUnread   *bool                       `json:"notification_is_unread,omitempty"`
	NotificationType       *string                     `json:"notification_type,omitempty"`
}

type RespNotificationModelLinks struct {
	Content       *string `json:"content,omitempty"`
	CreatorAvatar *string `json:"creator_avatar,omitempty"`
}

type RespPostCommentModel struct {
	Links                    *RespPostCommentModelLinks       `json:"links,omitempty"`
	Permissions              *RespPostCommentModelPermissions `json:"permissions,omitempty"`
	PostCommentBody          *string                          `json:"post_comment_body,omitempty"`
	PostCommentBodyHTML      *string                          `json:"post_comment_body_html,omitempty"`
	PostCommentBodyPlainText *string                          `json:"post_comment_body_plain_text,omitempty"`
	PostCommentCreateDate    *int                             `json:"post_comment_create_date,omitempty"`
	PostCommentID            *int                             `json:"post_comment_id,omitempty"`
	PostCommentIsDeleted     *bool                            `json:"post_comment_is_deleted,omitempty"`
	PostCommentIsPublished   *bool                            `json:"post_comment_is_published,omitempty"`
	PostCommentLikeCount     *int                             `json:"post_comment_like_count,omitempty"`
	PostCommentUpdateDate    *int                             `json:"post_comment_update_date,omitempty"`
	PostID                   *int                             `json:"post_id,omitempty"`
	PosterUserID             *int                             `json:"poster_user_id,omitempty"`
	PosterUsername           *string                          `json:"poster_username,omitempty"`
	PosterUsernameHTML       *string                          `json:"poster_username_html,omitempty"`
	ThreadID                 *int                             `json:"thread_id,omitempty"`
	UserIsIgnored            *bool                            `json:"user_is_ignored,omitempty"`
}

type RespPostCommentModelLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Post         *string `json:"post,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type RespPostCommentModelPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	Like   *bool `json:"like,omitempty"`
	Reply  *bool `json:"reply,omitempty"`
	Report *bool `json:"report,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type RespPostModel struct {
	Links              *RespPostModelLinks       `json:"links,omitempty"`
	Permissions        *RespPostModelPermissions `json:"permissions,omitempty"`
	PostBody           *string                   `json:"post_body,omitempty"`
	PostBodyHTML       *string                   `json:"post_body_html,omitempty"`
	PostBodyPlainText  *string                   `json:"post_body_plain_text,omitempty"`
	PostCreateDate     *int                      `json:"post_create_date,omitempty"`
	PostID             *int                      `json:"post_id,omitempty"`
	PostIsDeleted      *bool                     `json:"post_is_deleted,omitempty"`
	PostIsFirstPost    *bool                     `json:"post_is_first_post,omitempty"`
	PostIsPublished    *bool                     `json:"post_is_published,omitempty"`
	PostLikeCount      *int                      `json:"post_like_count,omitempty"`
	PostUpdateDate     *int                      `json:"post_update_date,omitempty"`
	PosterUserID       *int                      `json:"poster_user_id,omitempty"`
	PosterUsername     *string                   `json:"poster_username,omitempty"`
	PosterUsernameHTML *string                   `json:"poster_username_html,omitempty"`
	Signature          *string                   `json:"signature,omitempty"`
	SignatureHTML      *string                   `json:"signature_html,omitempty"`
	SignaturePlainText *string                   `json:"signature_plain_text,omitempty"`
	ThreadID           *int                      `json:"thread_id,omitempty"`
	ThreadIsDeleted    *bool                     `json:"thread_is_deleted,omitempty"`
	UserIsIgnored      *bool                     `json:"user_is_ignored,omitempty"`
}

type RespPostModelLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type RespPostModelPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	Like   *bool `json:"like,omitempty"`
	Reply  *bool `json:"reply,omitempty"`
	Report *bool `json:"report,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type RespProfilePostCommentModel struct {
	CommentBody          *string                                 `json:"comment_body,omitempty"`
	CommentBodyHTML      *string                                 `json:"comment_body_html,omitempty"`
	CommentBodyPlainText *string                                 `json:"comment_body_plain_text,omitempty"`
	CommentCreateDate    *int                                    `json:"comment_create_date,omitempty"`
	CommentID            *int                                    `json:"comment_id,omitempty"`
	CommentUserID        *int                                    `json:"comment_user_id,omitempty"`
	CommentUsername      *string                                 `json:"comment_username,omitempty"`
	CommentUsernameHTML  *string                                 `json:"comment_username_html,omitempty"`
	Links                *RespProfilePostCommentModelLinks       `json:"links,omitempty"`
	Permissions          *RespProfilePostCommentModelPermissions `json:"permissions,omitempty"`
	ProfilePostID        *int                                    `json:"profile_post_id,omitempty"`
	TimelineUserID       *int                                    `json:"timeline_user_id,omitempty"`
	UserIsIgnored        *bool                                   `json:"user_is_ignored,omitempty"`
}

type RespProfilePostCommentModelLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	ProfilePost  *string `json:"profile_post,omitempty"`
	Timeline     *string `json:"timeline,omitempty"`
	TimelineUser *string `json:"timeline_user,omitempty"`
}

type RespProfilePostCommentModelPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type RespProfilePostModel struct {
	Links                  *RespProfilePostModelLinks       `json:"links,omitempty"`
	Permissions            *RespProfilePostModelPermissions `json:"permissions,omitempty"`
	PostBody               *string                          `json:"post_body,omitempty"`
	PostBodyHTML           *string                          `json:"post_body_html,omitempty"`
	PostBodyPlainText      *string                          `json:"post_body_plain_text,omitempty"`
	PostCommentCount       *int                             `json:"post_comment_count,omitempty"`
	PostCommentsIsDisabled *int                             `json:"post_comments_is_disabled,omitempty"`
	PostCreateDate         *int                             `json:"post_create_date,omitempty"`
	PostIsDeleted          *bool                            `json:"post_is_deleted,omitempty"`
	PostIsLiked            *bool                            `json:"post_is_liked,omitempty"`
	PostIsPublished        *bool                            `json:"post_is_published,omitempty"`
	PostIsSticked          *bool                            `json:"post_is_sticked,omitempty"`
	PostLikeCount          *int                             `json:"post_like_count,omitempty"`
	PosterUserID           *int                             `json:"poster_user_id,omitempty"`
	PosterUsername         *string                          `json:"poster_username,omitempty"`
	PosterUsernameHTML     *string                          `json:"poster_username_html,omitempty"`
	ProfilePostID          *int                             `json:"profile_post_id,omitempty"`
	TimelineUser           *RespUserModel                   `json:"timeline_user,omitempty"`
	TimelineUserID         *int                             `json:"timeline_user_id,omitempty"`
	TimelineUsername       *string                          `json:"timeline_username,omitempty"`
	UserIsIgnored          *bool                            `json:"user_is_ignored,omitempty"`
}

type RespProfilePostModelLinks struct {
	Comments     *string `json:"comments,omitempty"`
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Timeline     *string `json:"timeline,omitempty"`
	TimelineUser *string `json:"timeline_user,omitempty"`
}

type RespProfilePostModelPermissions struct {
	Comment *bool `json:"comment,omitempty"`
	Delete  *bool `json:"delete,omitempty"`
	Edit    *bool `json:"edit,omitempty"`
	Like    *bool `json:"like,omitempty"`
	Report  *bool `json:"report,omitempty"`
	Stick   *bool `json:"stick,omitempty"`
	View    *bool `json:"view,omitempty"`
}

type RespProfilePostModelTimelineUser struct {
	Balance                     *string                                                           `json:"balance,omitempty"`
	Banner                      *string                                                           `json:"banner,omitempty"`
	Birthday                    *RespProfilePostModelTimelineUserBirthday                         `json:"birthday,omitempty"`
	ContestCount                *int                                                              `json:"contest_count,omitempty"`
	ConvWelcomeMessage          *string                                                           `json:"conv_welcome_message,omitempty"`
	CuratorTitles               []string                                                          `json:"curator_titles,omitempty"`
	Currency                    *string                                                           `json:"currency,omitempty"`
	CustomTitle                 *string                                                           `json:"custom_title,omitempty"`
	DisplayBannerID             *int                                                              `json:"display_banner_id,omitempty"`
	DisplayIconGroupID          *int                                                              `json:"display_icon_group_id,omitempty"`
	EditPermissions             *RespProfilePostModelTimelineUserEditPermissions                  `json:"edit_permissions,omitempty"`
	Fields                      []RespProfilePostModelTimelineUserFieldsItem                      `json:"fields,omitempty"`
	Hold                        *string                                                           `json:"hold,omitempty"`
	IsBanned                    *int                                                              `json:"is_banned,omitempty"`
	Links                       *RespProfilePostModelTimelineUserLinks                            `json:"links,omitempty"`
	Permissions                 *RespProfilePostModelTimelineUserPermissions                      `json:"permissions,omitempty"`
	SecretAnswerFirstLetter     *string                                                           `json:"secret_answer_first_letter,omitempty"`
	SecretAnswerRendered        *string                                                           `json:"secret_answer_rendered,omitempty"`
	SelfPermissions             *RespProfilePostModelTimelineUserSelfPermissions                  `json:"self_permissions,omitempty"`
	ShortLink                   *string                                                           `json:"short_link,omitempty"`
	TrophyCount                 *int                                                              `json:"trophy_count,omitempty"`
	UserDeposit                 *int                                                              `json:"user_deposit,omitempty"`
	UserEmail                   *string                                                           `json:"user_email,omitempty"`
	UserExternalAuthentications []RespProfilePostModelTimelineUserUserExternalAuthenticationsItem `json:"user_external_authentications,omitempty"`
	UserFollowers               *RespProfilePostModelTimelineUserUserFollowers                    `json:"user_followers,omitempty"`
	UserFollowing               *RespProfilePostModelTimelineUserUserFollowing                    `json:"user_following,omitempty"`
	UserGroupID                 *int                                                              `json:"user_group_id,omitempty"`
	UserGroups                  []RespProfilePostModelTimelineUserUserGroupsItem                  `json:"user_groups,omitempty"`
	UserID                      *int                                                              `json:"user_id,omitempty"`
	UserIsFollowed              *bool                                                             `json:"user_is_followed,omitempty"`
	UserIsIgnored               *bool                                                             `json:"user_is_ignored,omitempty"`
	UserIsValid                 *bool                                                             `json:"user_is_valid,omitempty"`
	UserIsVerified              *bool                                                             `json:"user_is_verified,omitempty"`
	UserIsVisitor               *bool                                                             `json:"user_is_visitor,omitempty"`
	UserLastSeenDate            *int                                                              `json:"user_last_seen_date,omitempty"`
	UserLike2Count              *int                                                              `json:"user_like2_count,omitempty"`
	UserLikeCount               *int                                                              `json:"user_like_count,omitempty"`
	UserMessageCount            *int                                                              `json:"user_message_count,omitempty"`
	UserRegisterDate            *int                                                              `json:"user_register_date,omitempty"`
	UserTimezoneOffset          *int                                                              `json:"user_timezone_offset,omitempty"`
	UserTitle                   *string                                                           `json:"user_title,omitempty"`
	UserUnreadConversationCount *int                                                              `json:"user_unread_conversation_count,omitempty"`
	UserUnreadNotificationCount *int                                                              `json:"user_unread_notification_count,omitempty"`
	Username                    *string                                                           `json:"username,omitempty"`
	UsernameHTML                *string                                                           `json:"username_html,omitempty"`
}

type RespProfilePostModelTimelineUserBirthday struct {
	Age       *int                                               `json:"age,omitempty"`
	Format    *string                                            `json:"format,omitempty"`
	TimeStamp *RespProfilePostModelTimelineUserBirthdayTimeStamp `json:"timeStamp,omitempty"`
}

type RespProfilePostModelTimelineUserBirthdayTimeStamp struct {
	Date         *string `json:"date,omitempty"`
	Timezone     *string `json:"timezone,omitempty"`
	TimezoneType *int    `json:"timezone_type,omitempty"`
}

type RespProfilePostModelTimelineUserEditPermissions struct {
	Fields            *bool `json:"fields,omitempty"`
	HideUsernameLogs  *bool `json:"hide_username_logs,omitempty"`
	Password          *bool `json:"password,omitempty"`
	PrimaryGroupID    *bool `json:"primary_group_id,omitempty"`
	SecondaryGroupIds *bool `json:"secondary_group_ids,omitempty"`
	ShortLink         *bool `json:"short_link,omitempty"`
	UserDobDay        *bool `json:"user_dob_day,omitempty"`
	UserDobMonth      *bool `json:"user_dob_month,omitempty"`
	UserDobYear       *bool `json:"user_dob_year,omitempty"`
	UserEmail         *bool `json:"user_email,omitempty"`
	UserTitle         *bool `json:"user_title,omitempty"`
	Username          *bool `json:"username,omitempty"`
}

type RespProfilePostModelTimelineUserFieldsItem struct {
	Choices       []RespProfilePostModelTimelineUserFieldsItemChoicesItem `json:"choices,omitempty"`
	Description   *string                                                 `json:"description,omitempty"`
	ID            *string                                                 `json:"id,omitempty"`
	IsMultiChoice *bool                                                   `json:"is_multi_choice,omitempty"`
	IsRequired    *bool                                                   `json:"is_required,omitempty"`
	Position      *string                                                 `json:"position,omitempty"`
	Title         *string                                                 `json:"title,omitempty"`
	Value         *string                                                 `json:"value,omitempty"`
	Values        interface{}                                             `json:"values,omitempty"`
}

type RespProfilePostModelTimelineUserFieldsItemChoicesItem struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type RespProfilePostModelTimelineUserLinks struct {
	Avatar      *string `json:"avatar,omitempty"`
	AvatarBig   *string `json:"avatar_big,omitempty"`
	AvatarSmall *string `json:"avatar_small,omitempty"`
	BackgroundL *string `json:"background_l,omitempty"`
	BackgroundM *string `json:"background_m,omitempty"`
	Detail      *string `json:"detail,omitempty"`
	Followers   *string `json:"followers,omitempty"`
	Followings  *string `json:"followings,omitempty"`
	Ignore      *string `json:"ignore,omitempty"`
	Permalink   *string `json:"permalink,omitempty"`
	Status      *string `json:"status,omitempty"`
	Timeline    *string `json:"timeline,omitempty"`
}

type RespProfilePostModelTimelineUserPermissions struct {
	Edit        *bool `json:"edit,omitempty"`
	Follow      *bool `json:"follow,omitempty"`
	Ignore      *bool `json:"ignore,omitempty"`
	ProfilePost *bool `json:"profile_post,omitempty"`
}

type RespProfilePostModelTimelineUserSelfPermissions struct {
	CreateConversation *bool `json:"create_conversation,omitempty"`
}

type RespProfilePostModelTimelineUserUserExternalAuthenticationsItem struct {
	Provider    *string `json:"provider,omitempty"`
	ProviderKey *string `json:"provider_key,omitempty"`
}

type RespProfilePostModelTimelineUserUserFollowers struct {
	Count *int                                                     `json:"count,omitempty"`
	Users []RespProfilePostModelTimelineUserUserFollowersUsersItem `json:"users,omitempty"`
}

type RespProfilePostModelTimelineUserUserFollowersUsersItem struct {
	Avatar       *string `json:"avatar,omitempty"`
	UserID       *int    `json:"user_id,omitempty"`
	Username     *string `json:"username,omitempty"`
	UsernameHTML *string `json:"username_html,omitempty"`
}

type RespProfilePostModelTimelineUserUserFollowing struct {
	Count *int                                                     `json:"count,omitempty"`
	Users []RespProfilePostModelTimelineUserUserFollowingUsersItem `json:"users,omitempty"`
}

type RespProfilePostModelTimelineUserUserFollowingUsersItem struct {
	Avatar       *string `json:"avatar,omitempty"`
	UserID       *int    `json:"user_id,omitempty"`
	Username     *string `json:"username,omitempty"`
	UsernameHTML *string `json:"username_html,omitempty"`
}

type RespProfilePostModelTimelineUserUserGroupsItem struct {
	DisplayBannerSelectable *bool   `json:"display_banner_selectable,omitempty"`
	DisplayGroupSelectable  *bool   `json:"display_group_selectable,omitempty"`
	DisplayIconSelectable   *bool   `json:"display_icon_selectable,omitempty"`
	IsPrimaryGroup          *bool   `json:"is_primary_group,omitempty"`
	UserGroupBannerCSSClass *string `json:"user_group_banner_css_class,omitempty"`
	UserGroupBannerText     *string `json:"user_group_banner_text,omitempty"`
	UserGroupBannerTextEn   *string `json:"user_group_banner_text_en,omitempty"`
	UserGroupIconClass      *string `json:"user_group_icon_class,omitempty"`
	UserGroupID             *int    `json:"user_group_id,omitempty"`
	UserGroupTitle          *string `json:"user_group_title,omitempty"`
	UserGroupTitleEn        *string `json:"user_group_title_en,omitempty"`
}

type RespSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type RespThreadModel struct {
	Contest             *RespThreadModelContest      `json:"contest,omitempty"`
	CreatorUserID       *int                         `json:"creator_user_id,omitempty"`
	CreatorUsername     *string                      `json:"creator_username,omitempty"`
	CreatorUsernameHTML *string                      `json:"creator_username_html,omitempty"`
	FirstPost           *RespThreadModelFirstPost    `json:"first_post,omitempty"`
	ForumID             *int                         `json:"forum_id,omitempty"`
	LastPost            *RespThreadModelLastPost     `json:"last_post,omitempty"`
	Links               *RespThreadModelLinks        `json:"links,omitempty"`
	NodeTitle           *string                      `json:"node_title,omitempty"`
	Permissions         *RespThreadModelPermissions  `json:"permissions,omitempty"`
	Restrictions        *RespThreadModelRestrictions `json:"restrictions,omitempty"`
	ThreadCreateDate    *int                         `json:"thread_create_date,omitempty"`
	ThreadID            *int                         `json:"thread_id,omitempty"`
	ThreadIsClosed      *bool                        `json:"thread_is_closed,omitempty"`
	ThreadIsDeleted     *bool                        `json:"thread_is_deleted,omitempty"`
	ThreadIsFollowed    *bool                        `json:"thread_is_followed,omitempty"`
	ThreadIsPublished   *bool                        `json:"thread_is_published,omitempty"`
	ThreadIsStarred     *bool                        `json:"thread_is_starred,omitempty"`
	ThreadIsSticky      *bool                        `json:"thread_is_sticky,omitempty"`
	ThreadPostCount     *int                         `json:"thread_post_count,omitempty"`
	ThreadPrefixes      []interface{}                `json:"thread_prefixes,omitempty"`
	ThreadTags          interface{}                  `json:"thread_tags,omitempty"`
	ThreadTitle         *string                      `json:"thread_title,omitempty"`
	ThreadUpdateDate    *int                         `json:"thread_update_date,omitempty"`
	ThreadViewCount     *int                         `json:"thread_view_count,omitempty"`
	UserIsIgnored       *bool                        `json:"user_is_ignored,omitempty"`
}

type RespThreadModelContest struct {
	AlreadyParticipate    *bool                              `json:"already_participate,omitempty"`
	ChanceToWin           *float64                           `json:"chance_to_win,omitempty"`
	CountWinners          *int                               `json:"count_winners,omitempty"`
	FinishDate            *int                               `json:"finish_date,omitempty"`
	IsFinished            *int                               `json:"is_finished,omitempty"`
	IsMoneyPlaces         *int                               `json:"is_money_places,omitempty"`
	NeededMembers         *int                               `json:"needed_members,omitempty"`
	NowCountMembers       *int                               `json:"now_count_members,omitempty"`
	Permissions           *RespThreadModelContestPermissions `json:"permissions,omitempty"`
	PrizeData             *int                               `json:"prize_data,omitempty"`
	PrizeType             *string                            `json:"prize_type,omitempty"`
	PrizeTypePhrase       *string                            `json:"prize_type_phrase,omitempty"`
	RequireLikeCount      *int                               `json:"require_like_count,omitempty"`
	RequireTotalLikeCount *int                               `json:"require_total_like_count,omitempty"`
	Type_                 *string                            `json:"type,omitempty"`
	Winners               []int                              `json:"winners,omitempty"`
}

type RespThreadModelContestPermissions struct {
	CanFinish           *bool   `json:"can_finish,omitempty"`
	CanParticipate      *bool   `json:"can_participate,omitempty"`
	CanParticipateError *string `json:"can_participate_error,omitempty"`
	CanViewUserList     *bool   `json:"can_view_user_list,omitempty"`
}

type RespThreadModelFirstPost struct {
	Links              *RespThreadModelFirstPostLinks       `json:"links,omitempty"`
	Permissions        *RespThreadModelFirstPostPermissions `json:"permissions,omitempty"`
	PostBody           *string                              `json:"post_body,omitempty"`
	PostBodyHTML       *string                              `json:"post_body_html,omitempty"`
	PostBodyPlainText  *string                              `json:"post_body_plain_text,omitempty"`
	PostCreateDate     *int                                 `json:"post_create_date,omitempty"`
	PostID             *int                                 `json:"post_id,omitempty"`
	PostIsDeleted      *bool                                `json:"post_is_deleted,omitempty"`
	PostIsFirstPost    *bool                                `json:"post_is_first_post,omitempty"`
	PostIsLiked        *bool                                `json:"post_is_liked,omitempty"`
	PostIsPublished    *bool                                `json:"post_is_published,omitempty"`
	PostLikeCount      *int                                 `json:"post_like_count,omitempty"`
	PostUpdateDate     *int                                 `json:"post_update_date,omitempty"`
	PosterUserID       *int                                 `json:"poster_user_id,omitempty"`
	PosterUsername     *string                              `json:"poster_username,omitempty"`
	PosterUsernameHTML *string                              `json:"poster_username_html,omitempty"`
	Signature          *string                              `json:"signature,omitempty"`
	SignatureHTML      *string                              `json:"signature_html,omitempty"`
	SignaturePlainText *string                              `json:"signature_plain_text,omitempty"`
	ThreadID           *int                                 `json:"thread_id,omitempty"`
	ThreadIsDeleted    *bool                                `json:"thread_is_deleted,omitempty"`
	UserIsIgnored      *bool                                `json:"user_is_ignored,omitempty"`
}

type RespThreadModelFirstPostLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type RespThreadModelFirstPostPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	Like   *bool `json:"like,omitempty"`
	Reply  *bool `json:"reply,omitempty"`
	Report *bool `json:"report,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type RespThreadModelLastPost struct {
	Links              *RespThreadModelLastPostLinks       `json:"links,omitempty"`
	Permissions        *RespThreadModelLastPostPermissions `json:"permissions,omitempty"`
	PostBody           *string                             `json:"post_body,omitempty"`
	PostBodyHTML       *string                             `json:"post_body_html,omitempty"`
	PostBodyPlainText  *string                             `json:"post_body_plain_text,omitempty"`
	PostCreateDate     *int                                `json:"post_create_date,omitempty"`
	PostID             *int                                `json:"post_id,omitempty"`
	PostIsDeleted      *bool                               `json:"post_is_deleted,omitempty"`
	PostIsFirstPost    *bool                               `json:"post_is_first_post,omitempty"`
	PostIsLiked        *bool                               `json:"post_is_liked,omitempty"`
	PostIsPublished    *bool                               `json:"post_is_published,omitempty"`
	PostLikeCount      *int                                `json:"post_like_count,omitempty"`
	PostUpdateDate     *int                                `json:"post_update_date,omitempty"`
	PosterUserID       *int                                `json:"poster_user_id,omitempty"`
	PosterUsername     *string                             `json:"poster_username,omitempty"`
	PosterUsernameHTML *string                             `json:"poster_username_html,omitempty"`
	Signature          *string                             `json:"signature,omitempty"`
	SignatureHTML      *string                             `json:"signature_html,omitempty"`
	SignaturePlainText *string                             `json:"signature_plain_text,omitempty"`
	ThreadID           *int                                `json:"thread_id,omitempty"`
	ThreadIsDeleted    *bool                               `json:"thread_is_deleted,omitempty"`
	UserIsIgnored      *bool                               `json:"user_is_ignored,omitempty"`
}

type RespThreadModelLastPostLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type RespThreadModelLastPostPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	Like   *bool `json:"like,omitempty"`
	Reply  *bool `json:"reply,omitempty"`
	Report *bool `json:"report,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type RespThreadModelLinks struct {
	Detail            *string `json:"detail,omitempty"`
	FirstPost         *string `json:"first_post,omitempty"`
	FirstPoster       *string `json:"first_poster,omitempty"`
	FirstPosterAvatar *string `json:"first_poster_avatar,omitempty"`
	Followers         *string `json:"followers,omitempty"`
	Forum             *string `json:"forum,omitempty"`
	LastPost          *string `json:"last_post,omitempty"`
	Permalink         *string `json:"permalink,omitempty"`
	Posts             *string `json:"posts,omitempty"`
}

type RespThreadModelPermissions struct {
	Bump      *RespThreadModelPermissionsBump `json:"bump,omitempty"`
	Delete    *bool                           `json:"delete,omitempty"`
	Edit      *bool                           `json:"edit,omitempty"`
	EditTags  *bool                           `json:"edit_tags,omitempty"`
	EditTitle *bool                           `json:"edit_title,omitempty"`
	Follow    *bool                           `json:"follow,omitempty"`
	Post      *bool                           `json:"post,omitempty"`
	View      *bool                           `json:"view,omitempty"`
}

type RespThreadModelPermissionsBump struct {
	AvailableCount    *int        `json:"available_count,omitempty"`
	Can               *bool       `json:"can,omitempty"`
	Error             interface{} `json:"error,omitempty"`
	NextAvailableTime interface{} `json:"next_available_time,omitempty"`
}

type RespThreadModelRestrictions struct {
	MaxReplyCount *int `json:"max_reply_count,omitempty"`
	ReplyDelay    *int `json:"reply_delay,omitempty"`
}

type RespUserModel struct {
	Balance                     *string                                        `json:"balance,omitempty"`
	Banner                      *string                                        `json:"banner,omitempty"`
	Birthday                    *RespUserModelBirthday                         `json:"birthday,omitempty"`
	ContestCount                *int                                           `json:"contest_count,omitempty"`
	ConvWelcomeMessage          *string                                        `json:"conv_welcome_message,omitempty"`
	CuratorTitles               []string                                       `json:"curator_titles,omitempty"`
	Currency                    *string                                        `json:"currency,omitempty"`
	CustomTitle                 *string                                        `json:"custom_title,omitempty"`
	DisplayBannerID             *int                                           `json:"display_banner_id,omitempty"`
	DisplayIconGroupID          *int                                           `json:"display_icon_group_id,omitempty"`
	EditPermissions             *RespUserModelEditPermissions                  `json:"edit_permissions,omitempty"`
	Fields                      []RespUserModelFieldsItem                      `json:"fields,omitempty"`
	Hold                        *string                                        `json:"hold,omitempty"`
	IsBanned                    *int                                           `json:"is_banned,omitempty"`
	Links                       *RespUserModelLinks                            `json:"links,omitempty"`
	Permissions                 *RespUserModelPermissions                      `json:"permissions,omitempty"`
	SecretAnswerFirstLetter     *string                                        `json:"secret_answer_first_letter,omitempty"`
	SecretAnswerRendered        *string                                        `json:"secret_answer_rendered,omitempty"`
	SelfPermissions             *RespUserModelSelfPermissions                  `json:"self_permissions,omitempty"`
	ShortLink                   *string                                        `json:"short_link,omitempty"`
	TrophyCount                 *int                                           `json:"trophy_count,omitempty"`
	UserDeposit                 *int                                           `json:"user_deposit,omitempty"`
	UserEmail                   *string                                        `json:"user_email,omitempty"`
	UserExternalAuthentications []RespUserModelUserExternalAuthenticationsItem `json:"user_external_authentications,omitempty"`
	UserFollowers               *RespUserModelUserFollowers                    `json:"user_followers,omitempty"`
	UserFollowing               *RespUserModelUserFollowing                    `json:"user_following,omitempty"`
	UserGroupID                 *int                                           `json:"user_group_id,omitempty"`
	UserGroups                  []RespUserModelUserGroupsItem                  `json:"user_groups,omitempty"`
	UserID                      *int                                           `json:"user_id,omitempty"`
	UserIsFollowed              *bool                                          `json:"user_is_followed,omitempty"`
	UserIsIgnored               *bool                                          `json:"user_is_ignored,omitempty"`
	UserIsValid                 *bool                                          `json:"user_is_valid,omitempty"`
	UserIsVerified              *bool                                          `json:"user_is_verified,omitempty"`
	UserIsVisitor               *bool                                          `json:"user_is_visitor,omitempty"`
	UserLastSeenDate            *int                                           `json:"user_last_seen_date,omitempty"`
	UserLike2Count              *int                                           `json:"user_like2_count,omitempty"`
	UserLikeCount               *int                                           `json:"user_like_count,omitempty"`
	UserMessageCount            *int                                           `json:"user_message_count,omitempty"`
	UserRegisterDate            *int                                           `json:"user_register_date,omitempty"`
	UserTimezoneOffset          *int                                           `json:"user_timezone_offset,omitempty"`
	UserTitle                   *string                                        `json:"user_title,omitempty"`
	UserUnreadConversationCount *int                                           `json:"user_unread_conversation_count,omitempty"`
	UserUnreadNotificationCount *int                                           `json:"user_unread_notification_count,omitempty"`
	Username                    *string                                        `json:"username,omitempty"`
	UsernameHTML                *string                                        `json:"username_html,omitempty"`
}

type RespUserModelBirthday struct {
	Age       *int                            `json:"age,omitempty"`
	Format    *string                         `json:"format,omitempty"`
	TimeStamp *RespUserModelBirthdayTimeStamp `json:"timeStamp,omitempty"`
}

type RespUserModelBirthdayTimeStamp struct {
	Date         *string `json:"date,omitempty"`
	Timezone     *string `json:"timezone,omitempty"`
	TimezoneType *int    `json:"timezone_type,omitempty"`
}

type RespUserModelEditPermissions struct {
	Fields            *bool `json:"fields,omitempty"`
	HideUsernameLogs  *bool `json:"hide_username_logs,omitempty"`
	Password          *bool `json:"password,omitempty"`
	PrimaryGroupID    *bool `json:"primary_group_id,omitempty"`
	SecondaryGroupIds *bool `json:"secondary_group_ids,omitempty"`
	ShortLink         *bool `json:"short_link,omitempty"`
	UserDobDay        *bool `json:"user_dob_day,omitempty"`
	UserDobMonth      *bool `json:"user_dob_month,omitempty"`
	UserDobYear       *bool `json:"user_dob_year,omitempty"`
	UserEmail         *bool `json:"user_email,omitempty"`
	UserTitle         *bool `json:"user_title,omitempty"`
	Username          *bool `json:"username,omitempty"`
}

type RespUserModelFieldsItem struct {
	Choices       []RespUserModelFieldsItemChoicesItem `json:"choices,omitempty"`
	Description   *string                              `json:"description,omitempty"`
	ID            *string                              `json:"id,omitempty"`
	IsMultiChoice *bool                                `json:"is_multi_choice,omitempty"`
	IsRequired    *bool                                `json:"is_required,omitempty"`
	Position      *string                              `json:"position,omitempty"`
	Title         *string                              `json:"title,omitempty"`
	Value         *string                              `json:"value,omitempty"`
	Values        interface{}                          `json:"values,omitempty"`
}

type RespUserModelFieldsItemChoicesItem struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type RespUserModelLinks struct {
	Avatar      *string `json:"avatar,omitempty"`
	AvatarBig   *string `json:"avatar_big,omitempty"`
	AvatarSmall *string `json:"avatar_small,omitempty"`
	BackgroundL *string `json:"background_l,omitempty"`
	BackgroundM *string `json:"background_m,omitempty"`
	Detail      *string `json:"detail,omitempty"`
	Followers   *string `json:"followers,omitempty"`
	Followings  *string `json:"followings,omitempty"`
	Ignore      *string `json:"ignore,omitempty"`
	Permalink   *string `json:"permalink,omitempty"`
	Status      *string `json:"status,omitempty"`
	Timeline    *string `json:"timeline,omitempty"`
}

type RespUserModelPermissions struct {
	Edit        *bool `json:"edit,omitempty"`
	Follow      *bool `json:"follow,omitempty"`
	Ignore      *bool `json:"ignore,omitempty"`
	ProfilePost *bool `json:"profile_post,omitempty"`
}

type RespUserModelSelfPermissions struct {
	CreateConversation *bool `json:"create_conversation,omitempty"`
}

type RespUserModelUserExternalAuthenticationsItem struct {
	Provider    *string `json:"provider,omitempty"`
	ProviderKey *string `json:"provider_key,omitempty"`
}

type RespUserModelUserFollowers struct {
	Count *int                                  `json:"count,omitempty"`
	Users []RespUserModelUserFollowersUsersItem `json:"users,omitempty"`
}

type RespUserModelUserFollowersUsersItem struct {
	Avatar       *string `json:"avatar,omitempty"`
	UserID       *int    `json:"user_id,omitempty"`
	Username     *string `json:"username,omitempty"`
	UsernameHTML *string `json:"username_html,omitempty"`
}

type RespUserModelUserFollowing struct {
	Count *int                                  `json:"count,omitempty"`
	Users []RespUserModelUserFollowingUsersItem `json:"users,omitempty"`
}

type RespUserModelUserFollowingUsersItem struct {
	Avatar       *string `json:"avatar,omitempty"`
	UserID       *int    `json:"user_id,omitempty"`
	Username     *string `json:"username,omitempty"`
	UsernameHTML *string `json:"username_html,omitempty"`
}

type RespUserModelUserGroupsItem struct {
	DisplayBannerSelectable *bool   `json:"display_banner_selectable,omitempty"`
	DisplayGroupSelectable  *bool   `json:"display_group_selectable,omitempty"`
	DisplayIconSelectable   *bool   `json:"display_icon_selectable,omitempty"`
	IsPrimaryGroup          *bool   `json:"is_primary_group,omitempty"`
	UserGroupBannerCSSClass *string `json:"user_group_banner_css_class,omitempty"`
	UserGroupBannerText     *string `json:"user_group_banner_text,omitempty"`
	UserGroupBannerTextEn   *string `json:"user_group_banner_text_en,omitempty"`
	UserGroupIconClass      *string `json:"user_group_icon_class,omitempty"`
	UserGroupID             *int    `json:"user_group_id,omitempty"`
	UserGroupTitle          *string `json:"user_group_title,omitempty"`
	UserGroupTitleEn        *string `json:"user_group_title_en,omitempty"`
}

type ResultsParams struct {
	// Number of results in a page.
	Limit *int `json:"limit,omitempty"`
	// Page number of results.
	Page *int `json:"page,omitempty"`
}

type ResultsResponse struct {
	Data       []ResultsResponseDataItem `json:"data,omitempty"`
	DataTotal  *int                      `json:"data_total,omitempty"`
	SearchTags map[string]interface{}    `json:"search_tags,omitempty"`
	SystemInfo *RespSystemInfo           `json:"system_info,omitempty"`
}

type ResultsResponseDataItem struct {
	ContentID           interface{}                         `json:"content_id,omitempty"`
	ContentType         *string                             `json:"content_type,omitempty"`
	CreatorUserID       *int                                `json:"creator_user_id,omitempty"`
	CreatorUsername     *string                             `json:"creator_username,omitempty"`
	CreatorUsernameHTML *string                             `json:"creator_username_html,omitempty"`
	FirstPost           *ResultsResponseDataItemFirstPost   `json:"first_post,omitempty"`
	Forum               *ResultsResponseDataItemForum       `json:"forum,omitempty"`
	ForumID             *int                                `json:"forum_id,omitempty"`
	Links               *ResultsResponseDataItemLinks       `json:"links,omitempty"`
	Permissions         *ResultsResponseDataItemPermissions `json:"permissions,omitempty"`
	ThreadCreateDate    *int                                `json:"thread_create_date,omitempty"`
	ThreadID            *int                                `json:"thread_id,omitempty"`
	ThreadIsDeleted     *bool                               `json:"thread_is_deleted,omitempty"`
	ThreadIsFollowed    *bool                               `json:"thread_is_followed,omitempty"`
	ThreadIsPublished   *bool                               `json:"thread_is_published,omitempty"`
	ThreadIsSticky      *bool                               `json:"thread_is_sticky,omitempty"`
	ThreadPostCount     *int                                `json:"thread_post_count,omitempty"`
	ThreadPrefixes      []interface{}                       `json:"thread_prefixes,omitempty"`
	ThreadTags          interface{}                         `json:"thread_tags,omitempty"`
	ThreadTitle         *string                             `json:"thread_title,omitempty"`
	ThreadUpdateDate    *int                                `json:"thread_update_date,omitempty"`
	ThreadViewCount     *int                                `json:"thread_view_count,omitempty"`
	UserIsIgnored       *bool                               `json:"user_is_ignored,omitempty"`
}

type ResultsResponseDataItemFirstPost struct {
	Links               *ResultsResponseDataItemFirstPostLinks       `json:"links,omitempty"`
	Permissions         *ResultsResponseDataItemFirstPostPermissions `json:"permissions,omitempty"`
	PostAttachmentCount *int                                         `json:"post_attachment_count,omitempty"`
	PostBody            *string                                      `json:"post_body,omitempty"`
	PostBodyHTML        *string                                      `json:"post_body_html,omitempty"`
	PostBodyPlainText   *string                                      `json:"post_body_plain_text,omitempty"`
	PostCreateDate      *int                                         `json:"post_create_date,omitempty"`
	PostID              *int                                         `json:"post_id,omitempty"`
	PostIsDeleted       *bool                                        `json:"post_is_deleted,omitempty"`
	PostIsFirstPost     *bool                                        `json:"post_is_first_post,omitempty"`
	PostIsPublished     *bool                                        `json:"post_is_published,omitempty"`
	PostLikeCount       *int                                         `json:"post_like_count,omitempty"`
	PostUpdateDate      *int                                         `json:"post_update_date,omitempty"`
	PosterUserID        *int                                         `json:"poster_user_id,omitempty"`
	PosterUsername      *string                                      `json:"poster_username,omitempty"`
	PosterUsernameHTML  *string                                      `json:"poster_username_html,omitempty"`
	Signature           *string                                      `json:"signature,omitempty"`
	SignatureHTML       *string                                      `json:"signature_html,omitempty"`
	SignaturePlainText  *string                                      `json:"signature_plain_text,omitempty"`
	ThreadID            *int                                         `json:"thread_id,omitempty"`
	UserIsIgnored       *bool                                        `json:"user_is_ignored,omitempty"`
}

type ResultsResponseDataItemFirstPostLinks struct {
	Attachments  *string `json:"attachments,omitempty"`
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type ResultsResponseDataItemFirstPostPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Like             *bool `json:"like,omitempty"`
	Reply            *bool `json:"reply,omitempty"`
	Report           *bool `json:"report,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type ResultsResponseDataItemForum struct {
	ForumDescription       *string                                         `json:"forum_description,omitempty"`
	ForumID                *int                                            `json:"forum_id,omitempty"`
	ForumIsFollowed        *bool                                           `json:"forum_is_followed,omitempty"`
	ForumPostCount         *int                                            `json:"forum_post_count,omitempty"`
	ForumPrefixes          []ResultsResponseDataItemForumForumPrefixesItem `json:"forum_prefixes,omitempty"`
	ForumThreadCount       *int                                            `json:"forum_thread_count,omitempty"`
	ForumTitle             *string                                         `json:"forum_title,omitempty"`
	Links                  *ResultsResponseDataItemForumLinks              `json:"links,omitempty"`
	Permissions            *ResultsResponseDataItemForumPermissions        `json:"permissions,omitempty"`
	ThreadDefaultPrefixID  *int                                            `json:"thread_default_prefix_id,omitempty"`
	ThreadPrefixIsRequired *bool                                           `json:"thread_prefix_is_required,omitempty"`
}

type ResultsResponseDataItemForumForumPrefixesItem struct {
	GroupPrefixes []ResultsResponseDataItemForumForumPrefixesItemGroupPrefixesItem `json:"group_prefixes,omitempty"`
	GroupTitle    *string                                                          `json:"group_title,omitempty"`
}

type ResultsResponseDataItemForumForumPrefixesItemGroupPrefixesItem struct {
	PrefixID    *int    `json:"prefix_id,omitempty"`
	PrefixTitle *string `json:"prefix_title,omitempty"`
}

type ResultsResponseDataItemForumLinks struct {
	Detail        *string `json:"detail,omitempty"`
	Followers     *string `json:"followers,omitempty"`
	Permalink     *string `json:"permalink,omitempty"`
	SubCategories *string `json:"sub-categories,omitempty"`
	SubForums     *string `json:"sub-forums,omitempty"`
	Threads       *string `json:"threads,omitempty"`
}

type ResultsResponseDataItemForumPermissions struct {
	CreateThread     *bool `json:"create_thread,omitempty"`
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	TagThread        *bool `json:"tag_thread,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type ResultsResponseDataItemLinks struct {
	Detail            *string `json:"detail,omitempty"`
	FirstPost         *string `json:"first_post,omitempty"`
	FirstPoster       *string `json:"first_poster,omitempty"`
	FirstPosterAvatar *string `json:"first_poster_avatar,omitempty"`
	Followers         *string `json:"followers,omitempty"`
	Forum             *string `json:"forum,omitempty"`
	LastPost          *string `json:"last_post,omitempty"`
	Permalink         *string `json:"permalink,omitempty"`
	Posts             *string `json:"posts,omitempty"`
}

type ResultsResponseDataItemPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	EditTags         *bool `json:"edit_tags,omitempty"`
	EditTitle        *bool `json:"edit_title,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	Post             *bool `json:"post,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type ResultsResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type SaveChanges struct {
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type SaveChangesSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type SearchParams struct {
	// Id of conversation.
	ConversationID *int `json:"conversation_id,omitempty"`
	// Search query string.
	Q *string `json:"q,omitempty"`
	// Search for recipients.
	SearchRecipients *bool `json:"search_recipients,omitempty"`
}

type SearchResponse struct {
	Conversations []SearchResponseConversationsItem `json:"conversations,omitempty"`
	Recipients    *bool                             `json:"recipients,omitempty"`
	SystemInfo    *RespSystemInfo                   `json:"system_info,omitempty"`
}

type SearchResponseConversationsItem struct {
	Alerts                   *int                                            `json:"alerts,omitempty"`
	ConversationCreateDate   *int                                            `json:"conversation_create_date,omitempty"`
	ConversationID           *int                                            `json:"conversation_id,omitempty"`
	ConversationIsDeleted    *bool                                           `json:"conversation_is_deleted,omitempty"`
	ConversationIsNew        *bool                                           `json:"conversation_is_new,omitempty"`
	ConversationIsOpen       *bool                                           `json:"conversation_is_open,omitempty"`
	ConversationLastReadDate *int                                            `json:"conversation_last_read_date,omitempty"`
	ConversationMessageCount *int                                            `json:"conversation_message_count,omitempty"`
	ConversationOnlineCount  *int                                            `json:"conversation_online_count,omitempty"`
	ConversationTitle        *string                                         `json:"conversation_title,omitempty"`
	ConversationUpdateDate   *int                                            `json:"conversation_update_date,omitempty"`
	CreatorIsIgnored         *bool                                           `json:"creator_is_ignored,omitempty"`
	CreatorUserID            *int                                            `json:"creator_user_id,omitempty"`
	CreatorUsername          *string                                         `json:"creator_username,omitempty"`
	CreatorUsernameHTML      *string                                         `json:"creator_username_html,omitempty"`
	IsGroup                  *int                                            `json:"is_group,omitempty"`
	IsStarred                *int                                            `json:"is_starred,omitempty"`
	IsUnread                 *int                                            `json:"is_unread,omitempty"`
	Links                    *SearchResponseConversationsItemLinks           `json:"links,omitempty"`
	Permissions              *SearchResponseConversationsItemPermissions     `json:"permissions,omitempty"`
	Recipient                *SearchResponseConversationsItemRecipient       `json:"recipient,omitempty"`
	Recipients               []SearchResponseConversationsItemRecipientsItem `json:"recipients,omitempty"`
}

type SearchResponseConversationsItemLinks struct {
	Avatar    *string `json:"avatar,omitempty"`
	Detail    *string `json:"detail,omitempty"`
	Messages  *string `json:"messages,omitempty"`
	Permalink *string `json:"permalink,omitempty"`
}

type SearchResponseConversationsItemPermissions struct {
	EditOwnPost       *bool `json:"editOwnPost,omitempty"`
	Invite            *bool `json:"invite,omitempty"`
	Kick              *bool `json:"kick,omitempty"`
	ManageInviteLinks *bool `json:"manage_invite_links,omitempty"`
	Reply             *bool `json:"reply,omitempty"`
	StickyMessages    *bool `json:"stickyMessages,omitempty"`
	UploadAvatar      *bool `json:"upload_avatar,omitempty"`
	View              *bool `json:"view,omitempty"`
}

type SearchResponseConversationsItemRecipient struct {
	Avatar          *string `json:"avatar,omitempty"`
	ContactsChanged *bool   `json:"contacts_changed,omitempty"`
	IsOnline        *bool   `json:"is_online,omitempty"`
	LastActivity    *int    `json:"last_activity,omitempty"`
	UserID          *int    `json:"user_id,omitempty"`
	Username        *string `json:"username,omitempty"`
	UsernameHTML    *string `json:"username_html,omitempty"`
}

type SearchResponseConversationsItemRecipientsItem struct {
	Avatar          *string `json:"avatar,omitempty"`
	ContactsChanged *bool   `json:"contacts_changed,omitempty"`
	IsOnline        *bool   `json:"is_online,omitempty"`
	LastActivity    *int    `json:"last_activity,omitempty"`
	UserID          *int    `json:"user_id,omitempty"`
	Username        *string `json:"username,omitempty"`
	UsernameHTML    *string `json:"username_html,omitempty"`
}

type SearchResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type SecretAnswerTypesResponse struct {
	Data       []SecretAnswerTypesResponseDataItem `json:"data,omitempty"`
	SystemInfo *RespSystemInfo                     `json:"system_info,omitempty"`
}

type SecretAnswerTypesResponseDataItem struct {
	RenderedPhrase *string `json:"renderedPhrase,omitempty"`
	SaID           *int    `json:"sa_id,omitempty"`
}

type SecretAnswerTypesResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type StarResponse struct {
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type StarResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type StartResponse struct {
	Conversation *RespConversationModel `json:"conversation,omitempty"`
	SystemInfo   *RespSystemInfo        `json:"system_info,omitempty"`
}

type StartResponseConversation struct {
	Alerts                   *int                                      `json:"alerts,omitempty"`
	ConversationCreateDate   *int                                      `json:"conversation_create_date,omitempty"`
	ConversationID           *int                                      `json:"conversation_id,omitempty"`
	ConversationIsDeleted    *bool                                     `json:"conversation_is_deleted,omitempty"`
	ConversationIsNew        *bool                                     `json:"conversation_is_new,omitempty"`
	ConversationIsOpen       *bool                                     `json:"conversation_is_open,omitempty"`
	ConversationLastReadDate *int                                      `json:"conversation_last_read_date,omitempty"`
	ConversationMessageCount *int                                      `json:"conversation_message_count,omitempty"`
	ConversationOnlineCount  *int                                      `json:"conversation_online_count,omitempty"`
	ConversationTitle        *string                                   `json:"conversation_title,omitempty"`
	ConversationUpdateDate   *int                                      `json:"conversation_update_date,omitempty"`
	CreatorIsIgnored         *bool                                     `json:"creator_is_ignored,omitempty"`
	CreatorUserID            *int                                      `json:"creator_user_id,omitempty"`
	CreatorUsername          *string                                   `json:"creator_username,omitempty"`
	CreatorUsernameHTML      *string                                   `json:"creator_username_html,omitempty"`
	IsGroup                  *int                                      `json:"is_group,omitempty"`
	IsStarred                *int                                      `json:"is_starred,omitempty"`
	IsUnread                 *int                                      `json:"is_unread,omitempty"`
	Links                    *StartResponseConversationLinks           `json:"links,omitempty"`
	Permissions              *StartResponseConversationPermissions     `json:"permissions,omitempty"`
	Recipient                *StartResponseConversationRecipient       `json:"recipient,omitempty"`
	Recipients               []StartResponseConversationRecipientsItem `json:"recipients,omitempty"`
}

type StartResponseConversationLinks struct {
	Avatar    *string `json:"avatar,omitempty"`
	Detail    *string `json:"detail,omitempty"`
	Messages  *string `json:"messages,omitempty"`
	Permalink *string `json:"permalink,omitempty"`
}

type StartResponseConversationPermissions struct {
	EditOwnPost       *bool `json:"editOwnPost,omitempty"`
	Invite            *bool `json:"invite,omitempty"`
	Kick              *bool `json:"kick,omitempty"`
	ManageInviteLinks *bool `json:"manage_invite_links,omitempty"`
	Reply             *bool `json:"reply,omitempty"`
	StickyMessages    *bool `json:"stickyMessages,omitempty"`
	UploadAvatar      *bool `json:"upload_avatar,omitempty"`
	View              *bool `json:"view,omitempty"`
}

type StartResponseConversationRecipient struct {
	Avatar          *string `json:"avatar,omitempty"`
	ContactsChanged *bool   `json:"contacts_changed,omitempty"`
	IsOnline        *bool   `json:"is_online,omitempty"`
	LastActivity    *int    `json:"last_activity,omitempty"`
	UserID          *int    `json:"user_id,omitempty"`
	Username        *string `json:"username,omitempty"`
	UsernameHTML    *string `json:"username_html,omitempty"`
}

type StartResponseConversationRecipientsItem struct {
	Avatar          *string `json:"avatar,omitempty"`
	ContactsChanged *bool   `json:"contacts_changed,omitempty"`
	IsOnline        *bool   `json:"is_online,omitempty"`
	LastActivity    *int    `json:"last_activity,omitempty"`
	UserID          *int    `json:"user_id,omitempty"`
	Username        *string `json:"username,omitempty"`
	UsernameHTML    *string `json:"username_html,omitempty"`
}

type StartResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type TaggedParams struct {
	// Number of results in a page.
	Limit *int `json:"limit,omitempty"`
	// Page number of results.
	Page *int `json:"page,omitempty"`
	// Tag to search for tagged contents.
	Tag *string `json:"tag,omitempty"`
	// Array of tags to search for tagged contents.
	Tags []string `json:"tags,omitempty"`
}

type TaggedResponse struct {
	Data       []TaggedResponseDataItem `json:"data,omitempty"`
	DataTotal  *int                     `json:"data_total,omitempty"`
	SearchTags map[string]interface{}   `json:"search_tags,omitempty"`
	SystemInfo *RespSystemInfo          `json:"system_info,omitempty"`
}

type TaggedResponseDataItem struct {
	ContentID           interface{}                        `json:"content_id,omitempty"`
	ContentType         *string                            `json:"content_type,omitempty"`
	CreatorUserID       *int                               `json:"creator_user_id,omitempty"`
	CreatorUsername     *string                            `json:"creator_username,omitempty"`
	CreatorUsernameHTML *string                            `json:"creator_username_html,omitempty"`
	FirstPost           *TaggedResponseDataItemFirstPost   `json:"first_post,omitempty"`
	Forum               *TaggedResponseDataItemForum       `json:"forum,omitempty"`
	ForumID             *int                               `json:"forum_id,omitempty"`
	Links               *TaggedResponseDataItemLinks       `json:"links,omitempty"`
	Permissions         *TaggedResponseDataItemPermissions `json:"permissions,omitempty"`
	ThreadCreateDate    *int                               `json:"thread_create_date,omitempty"`
	ThreadID            *int                               `json:"thread_id,omitempty"`
	ThreadIsDeleted     *bool                              `json:"thread_is_deleted,omitempty"`
	ThreadIsFollowed    *bool                              `json:"thread_is_followed,omitempty"`
	ThreadIsPublished   *bool                              `json:"thread_is_published,omitempty"`
	ThreadIsSticky      *bool                              `json:"thread_is_sticky,omitempty"`
	ThreadPostCount     *int                               `json:"thread_post_count,omitempty"`
	ThreadPrefixes      []interface{}                      `json:"thread_prefixes,omitempty"`
	ThreadTags          interface{}                        `json:"thread_tags,omitempty"`
	ThreadTitle         *string                            `json:"thread_title,omitempty"`
	ThreadUpdateDate    *int                               `json:"thread_update_date,omitempty"`
	ThreadViewCount     *int                               `json:"thread_view_count,omitempty"`
	UserIsIgnored       *bool                              `json:"user_is_ignored,omitempty"`
}

type TaggedResponseDataItemFirstPost struct {
	Links               *TaggedResponseDataItemFirstPostLinks       `json:"links,omitempty"`
	Permissions         *TaggedResponseDataItemFirstPostPermissions `json:"permissions,omitempty"`
	PostAttachmentCount *int                                        `json:"post_attachment_count,omitempty"`
	PostBody            *string                                     `json:"post_body,omitempty"`
	PostBodyHTML        *string                                     `json:"post_body_html,omitempty"`
	PostBodyPlainText   *string                                     `json:"post_body_plain_text,omitempty"`
	PostCreateDate      *int                                        `json:"post_create_date,omitempty"`
	PostID              *int                                        `json:"post_id,omitempty"`
	PostIsDeleted       *bool                                       `json:"post_is_deleted,omitempty"`
	PostIsFirstPost     *bool                                       `json:"post_is_first_post,omitempty"`
	PostIsPublished     *bool                                       `json:"post_is_published,omitempty"`
	PostLikeCount       *int                                        `json:"post_like_count,omitempty"`
	PostUpdateDate      *int                                        `json:"post_update_date,omitempty"`
	PosterUserID        *int                                        `json:"poster_user_id,omitempty"`
	PosterUsername      *string                                     `json:"poster_username,omitempty"`
	PosterUsernameHTML  *string                                     `json:"poster_username_html,omitempty"`
	Signature           *string                                     `json:"signature,omitempty"`
	SignatureHTML       *string                                     `json:"signature_html,omitempty"`
	SignaturePlainText  *string                                     `json:"signature_plain_text,omitempty"`
	ThreadID            *int                                        `json:"thread_id,omitempty"`
	UserIsIgnored       *bool                                       `json:"user_is_ignored,omitempty"`
}

type TaggedResponseDataItemFirstPostLinks struct {
	Attachments  *string `json:"attachments,omitempty"`
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type TaggedResponseDataItemFirstPostPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Like             *bool `json:"like,omitempty"`
	Reply            *bool `json:"reply,omitempty"`
	Report           *bool `json:"report,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type TaggedResponseDataItemForum struct {
	ForumDescription       *string                                        `json:"forum_description,omitempty"`
	ForumID                *int                                           `json:"forum_id,omitempty"`
	ForumIsFollowed        *bool                                          `json:"forum_is_followed,omitempty"`
	ForumPostCount         *int                                           `json:"forum_post_count,omitempty"`
	ForumPrefixes          []TaggedResponseDataItemForumForumPrefixesItem `json:"forum_prefixes,omitempty"`
	ForumThreadCount       *int                                           `json:"forum_thread_count,omitempty"`
	ForumTitle             *string                                        `json:"forum_title,omitempty"`
	Links                  *TaggedResponseDataItemForumLinks              `json:"links,omitempty"`
	Permissions            *TaggedResponseDataItemForumPermissions        `json:"permissions,omitempty"`
	ThreadDefaultPrefixID  *int                                           `json:"thread_default_prefix_id,omitempty"`
	ThreadPrefixIsRequired *bool                                          `json:"thread_prefix_is_required,omitempty"`
}

type TaggedResponseDataItemForumForumPrefixesItem struct {
	GroupPrefixes []TaggedResponseDataItemForumForumPrefixesItemGroupPrefixesItem `json:"group_prefixes,omitempty"`
	GroupTitle    *string                                                         `json:"group_title,omitempty"`
}

type TaggedResponseDataItemForumForumPrefixesItemGroupPrefixesItem struct {
	PrefixID    *int    `json:"prefix_id,omitempty"`
	PrefixTitle *string `json:"prefix_title,omitempty"`
}

type TaggedResponseDataItemForumLinks struct {
	Detail        *string `json:"detail,omitempty"`
	Followers     *string `json:"followers,omitempty"`
	Permalink     *string `json:"permalink,omitempty"`
	SubCategories *string `json:"sub-categories,omitempty"`
	SubForums     *string `json:"sub-forums,omitempty"`
	Threads       *string `json:"threads,omitempty"`
}

type TaggedResponseDataItemForumPermissions struct {
	CreateThread     *bool `json:"create_thread,omitempty"`
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	TagThread        *bool `json:"tag_thread,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type TaggedResponseDataItemLinks struct {
	Detail            *string `json:"detail,omitempty"`
	FirstPost         *string `json:"first_post,omitempty"`
	FirstPoster       *string `json:"first_poster,omitempty"`
	FirstPosterAvatar *string `json:"first_poster_avatar,omitempty"`
	Followers         *string `json:"followers,omitempty"`
	Forum             *string `json:"forum,omitempty"`
	LastPost          *string `json:"last_post,omitempty"`
	Permalink         *string `json:"permalink,omitempty"`
	Posts             *string `json:"posts,omitempty"`
}

type TaggedResponseDataItemPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	EditTags         *bool `json:"edit_tags,omitempty"`
	EditTitle        *bool `json:"edit_title,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	Post             *bool `json:"post,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type TaggedResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type ThreadsParams struct {
	// The time in milliseconds (e.g. 1767214800) before last content date.
	Before *int `json:"before,omitempty"`
	// Number of thread data to be returned.
	DataLimit *int `json:"data_limit,omitempty"`
	// Id of the container forum to search for contents. Child forums of the specified forum will be included in the search.
	ForumID *int `json:"forum_id,omitempty"`
	// Number of results in a page.
	Limit *int `json:"limit,omitempty"`
	// Page number of results.
	Page *int `json:"page,omitempty"`
	// Search query. Can be skipped if **user_id** is set.
	Q *string `json:"q,omitempty"`
	// Tag to search for tagged contents.
	Tag    *string     `json:"tag,omitempty"`
	UserID StringOrInt `json:"user_id,omitempty"`
}

type ThreadsResponse struct {
	Data       []ThreadsResponseDataItem `json:"data,omitempty"`
	DataTotal  *int                      `json:"data_total,omitempty"`
	Links      *ThreadsResponseLinks     `json:"links,omitempty"`
	SystemInfo *RespSystemInfo           `json:"system_info,omitempty"`
}

type ThreadsResponseDataItem struct {
	ContentID           interface{}                         `json:"content_id,omitempty"`
	ContentType         *string                             `json:"content_type,omitempty"`
	CreatorUserID       *int                                `json:"creator_user_id,omitempty"`
	CreatorUsername     *string                             `json:"creator_username,omitempty"`
	CreatorUsernameHTML *string                             `json:"creator_username_html,omitempty"`
	FirstPost           *ThreadsResponseDataItemFirstPost   `json:"first_post,omitempty"`
	Forum               *ThreadsResponseDataItemForum       `json:"forum,omitempty"`
	ForumID             *int                                `json:"forum_id,omitempty"`
	Links               *ThreadsResponseDataItemLinks       `json:"links,omitempty"`
	Permissions         *ThreadsResponseDataItemPermissions `json:"permissions,omitempty"`
	ThreadCreateDate    *int                                `json:"thread_create_date,omitempty"`
	ThreadID            *int                                `json:"thread_id,omitempty"`
	ThreadIsDeleted     *bool                               `json:"thread_is_deleted,omitempty"`
	ThreadIsFollowed    *bool                               `json:"thread_is_followed,omitempty"`
	ThreadIsPublished   *bool                               `json:"thread_is_published,omitempty"`
	ThreadIsSticky      *bool                               `json:"thread_is_sticky,omitempty"`
	ThreadPostCount     *int                                `json:"thread_post_count,omitempty"`
	ThreadPrefixes      []interface{}                       `json:"thread_prefixes,omitempty"`
	ThreadTags          interface{}                         `json:"thread_tags,omitempty"`
	ThreadTitle         *string                             `json:"thread_title,omitempty"`
	ThreadUpdateDate    *int                                `json:"thread_update_date,omitempty"`
	ThreadViewCount     *int                                `json:"thread_view_count,omitempty"`
	UserIsIgnored       *bool                               `json:"user_is_ignored,omitempty"`
}

type ThreadsResponseDataItemFirstPost struct {
	Links               *ThreadsResponseDataItemFirstPostLinks       `json:"links,omitempty"`
	Permissions         *ThreadsResponseDataItemFirstPostPermissions `json:"permissions,omitempty"`
	PostAttachmentCount *int                                         `json:"post_attachment_count,omitempty"`
	PostBody            *string                                      `json:"post_body,omitempty"`
	PostBodyHTML        *string                                      `json:"post_body_html,omitempty"`
	PostBodyPlainText   *string                                      `json:"post_body_plain_text,omitempty"`
	PostCreateDate      *int                                         `json:"post_create_date,omitempty"`
	PostID              *int                                         `json:"post_id,omitempty"`
	PostIsDeleted       *bool                                        `json:"post_is_deleted,omitempty"`
	PostIsFirstPost     *bool                                        `json:"post_is_first_post,omitempty"`
	PostIsPublished     *bool                                        `json:"post_is_published,omitempty"`
	PostLikeCount       *int                                         `json:"post_like_count,omitempty"`
	PostUpdateDate      *int                                         `json:"post_update_date,omitempty"`
	PosterUserID        *int                                         `json:"poster_user_id,omitempty"`
	PosterUsername      *string                                      `json:"poster_username,omitempty"`
	PosterUsernameHTML  *string                                      `json:"poster_username_html,omitempty"`
	Signature           *string                                      `json:"signature,omitempty"`
	SignatureHTML       *string                                      `json:"signature_html,omitempty"`
	SignaturePlainText  *string                                      `json:"signature_plain_text,omitempty"`
	ThreadID            *int                                         `json:"thread_id,omitempty"`
	UserIsIgnored       *bool                                        `json:"user_is_ignored,omitempty"`
}

type ThreadsResponseDataItemFirstPostLinks struct {
	Attachments  *string `json:"attachments,omitempty"`
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type ThreadsResponseDataItemFirstPostPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Like             *bool `json:"like,omitempty"`
	Reply            *bool `json:"reply,omitempty"`
	Report           *bool `json:"report,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type ThreadsResponseDataItemForum struct {
	ForumDescription       *string                                  `json:"forum_description,omitempty"`
	ForumID                *int                                     `json:"forum_id,omitempty"`
	ForumIsFollowed        *bool                                    `json:"forum_is_followed,omitempty"`
	ForumPostCount         *int                                     `json:"forum_post_count,omitempty"`
	ForumPrefixes          []interface{}                            `json:"forum_prefixes,omitempty"`
	ForumThreadCount       *int                                     `json:"forum_thread_count,omitempty"`
	ForumTitle             *string                                  `json:"forum_title,omitempty"`
	Links                  *ThreadsResponseDataItemForumLinks       `json:"links,omitempty"`
	Permissions            *ThreadsResponseDataItemForumPermissions `json:"permissions,omitempty"`
	ThreadDefaultPrefixID  *int                                     `json:"thread_default_prefix_id,omitempty"`
	ThreadPrefixIsRequired *bool                                    `json:"thread_prefix_is_required,omitempty"`
}

type ThreadsResponseDataItemForumLinks struct {
	Detail        *string `json:"detail,omitempty"`
	Followers     *string `json:"followers,omitempty"`
	Permalink     *string `json:"permalink,omitempty"`
	SubCategories *string `json:"sub-categories,omitempty"`
	SubForums     *string `json:"sub-forums,omitempty"`
	Threads       *string `json:"threads,omitempty"`
}

type ThreadsResponseDataItemForumPermissions struct {
	CreateThread     *bool `json:"create_thread,omitempty"`
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	TagThread        *bool `json:"tag_thread,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type ThreadsResponseDataItemLinks struct {
	Detail            *string `json:"detail,omitempty"`
	FirstPost         *string `json:"first_post,omitempty"`
	FirstPoster       *string `json:"first_poster,omitempty"`
	FirstPosterAvatar *string `json:"first_poster_avatar,omitempty"`
	Followers         *string `json:"followers,omitempty"`
	Forum             *string `json:"forum,omitempty"`
	LastPost          *string `json:"last_post,omitempty"`
	Permalink         *string `json:"permalink,omitempty"`
	Posts             *string `json:"posts,omitempty"`
}

type ThreadsResponseDataItemPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	Post             *bool `json:"post,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type ThreadsResponseLinks struct {
	Next  *string `json:"next,omitempty"`
	Page  *int    `json:"page,omitempty"`
	Pages *int    `json:"pages,omitempty"`
}

type ThreadsResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type TokenParams struct {
	// Client ID.
	ClientID *string `json:"client_id,omitempty"`
	// Client secret.
	ClientSecret *string `json:"client_secret,omitempty"`
	// Authorization code.
	Code *string `json:"code,omitempty"`
	// Grant type.
	GrantType *AuthenticationTokenGrantType `json:"grant_type,omitempty"`
	// Account password.
	Password *string `json:"password,omitempty"`
	// Redirect URI.
	RedirectURI *string `json:"redirect_uri,omitempty"`
	// Refresh token.
	RefreshToken *string `json:"refresh_token,omitempty"`
	// Scope.
	Scope []string `json:"scope,omitempty"`
	// Account username/email.
	Username *string `json:"username,omitempty"`
}

type TokenResponse struct {
	// The access token issued by the authorization server
	AccessToken *string `json:"access_token,omitempty"`
	// The lifetime in seconds of the access token
	ExpiresIn *int `json:"expires_in,omitempty"`
	// The refresh token, which can be used to obtain new access tokens
	RefreshToken *string `json:"refresh_token,omitempty"`
	// The scope of the access token
	Scope *string `json:"scope,omitempty"`
	// The type of the token
	TokenType *string `json:"token_type,omitempty"`
}

type TrophiesResponse struct {
	SystemInfo *RespSystemInfo                `json:"system_info,omitempty"`
	Trophies   []TrophiesResponseTrophiesItem `json:"trophies,omitempty"`
}

type TrophiesResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type TrophiesResponseTrophiesItem struct {
	Description *string `json:"description,omitempty"`
	Title       *string `json:"title,omitempty"`
	TrophyID    *int    `json:"trophy_id,omitempty"`
	TrophyURL   *string `json:"trophy_url,omitempty"`
}

type UnreadParams struct {
	// Maximum number of result threads. The limit may get decreased if the value is too large (depending on the system configuration).
	Limit *int `json:"limit,omitempty"`
	// Id of the container forum to search for threads. Child forums of the specified forum will be included in the search.
	ForumID *int `json:"forum_id,omitempty"`
	// Number of thread data to be returned. Default value is 20.
	DataLimit *int `json:"data_limit,omitempty"`
}

type UnreadResponse struct {
	Data       []UnreadResponseDataItem    `json:"data,omitempty"`
	SystemInfo *RespSystemInfo             `json:"system_info,omitempty"`
	Threads    []UnreadResponseThreadsItem `json:"threads,omitempty"`
}

type UnreadResponseDataItem struct {
	ContentID           interface{}                        `json:"content_id,omitempty"`
	ContentType         *string                            `json:"content_type,omitempty"`
	CreatorUserID       *int                               `json:"creator_user_id,omitempty"`
	CreatorUsername     *string                            `json:"creator_username,omitempty"`
	CreatorUsernameHTML *string                            `json:"creator_username_html,omitempty"`
	FirstPost           *UnreadResponseDataItemFirstPost   `json:"first_post,omitempty"`
	Forum               *UnreadResponseDataItemForum       `json:"forum,omitempty"`
	ForumID             *int                               `json:"forum_id,omitempty"`
	Links               *UnreadResponseDataItemLinks       `json:"links,omitempty"`
	Permissions         *UnreadResponseDataItemPermissions `json:"permissions,omitempty"`
	ThreadCreateDate    *int                               `json:"thread_create_date,omitempty"`
	ThreadID            *int                               `json:"thread_id,omitempty"`
	ThreadIsDeleted     *bool                              `json:"thread_is_deleted,omitempty"`
	ThreadIsFollowed    *bool                              `json:"thread_is_followed,omitempty"`
	ThreadIsPublished   *bool                              `json:"thread_is_published,omitempty"`
	ThreadIsSticky      *bool                              `json:"thread_is_sticky,omitempty"`
	ThreadPostCount     *int                               `json:"thread_post_count,omitempty"`
	ThreadPrefixes      []interface{}                      `json:"thread_prefixes,omitempty"`
	ThreadTags          interface{}                        `json:"thread_tags,omitempty"`
	ThreadTitle         *string                            `json:"thread_title,omitempty"`
	ThreadUpdateDate    *int                               `json:"thread_update_date,omitempty"`
	ThreadViewCount     *int                               `json:"thread_view_count,omitempty"`
	UserIsIgnored       *bool                              `json:"user_is_ignored,omitempty"`
}

type UnreadResponseDataItemFirstPost struct {
	LikeUsers           []UnreadResponseDataItemFirstPostLikeUsersItem `json:"like_users,omitempty"`
	Links               *UnreadResponseDataItemFirstPostLinks          `json:"links,omitempty"`
	Permissions         *UnreadResponseDataItemFirstPostPermissions    `json:"permissions,omitempty"`
	PostAttachmentCount *int                                           `json:"post_attachment_count,omitempty"`
	PostBody            *string                                        `json:"post_body,omitempty"`
	PostBodyHTML        *string                                        `json:"post_body_html,omitempty"`
	PostBodyPlainText   *string                                        `json:"post_body_plain_text,omitempty"`
	PostCreateDate      *int                                           `json:"post_create_date,omitempty"`
	PostID              *int                                           `json:"post_id,omitempty"`
	PostIsDeleted       *bool                                          `json:"post_is_deleted,omitempty"`
	PostIsFirstPost     *bool                                          `json:"post_is_first_post,omitempty"`
	PostIsPublished     *bool                                          `json:"post_is_published,omitempty"`
	PostLikeCount       *int                                           `json:"post_like_count,omitempty"`
	PostUpdateDate      *int                                           `json:"post_update_date,omitempty"`
	PosterUserID        *int                                           `json:"poster_user_id,omitempty"`
	PosterUsername      *string                                        `json:"poster_username,omitempty"`
	PosterUsernameHTML  *string                                        `json:"poster_username_html,omitempty"`
	Signature           *string                                        `json:"signature,omitempty"`
	SignatureHTML       *string                                        `json:"signature_html,omitempty"`
	SignaturePlainText  *string                                        `json:"signature_plain_text,omitempty"`
	ThreadID            *int                                           `json:"thread_id,omitempty"`
	UserIsIgnored       *bool                                          `json:"user_is_ignored,omitempty"`
}

type UnreadResponseDataItemFirstPostLikeUsersItem struct {
	DisplayStyleGroupID *int    `json:"display_style_group_id,omitempty"`
	IsBanned            *int    `json:"is_banned,omitempty"`
	UniqUsernameCSS     *string `json:"uniq_username_css,omitempty"`
	UserID              *int    `json:"user_id,omitempty"`
	Username            *string `json:"username,omitempty"`
}

type UnreadResponseDataItemFirstPostLinks struct {
	Attachments  *string `json:"attachments,omitempty"`
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type UnreadResponseDataItemFirstPostPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Like             *bool `json:"like,omitempty"`
	Reply            *bool `json:"reply,omitempty"`
	Report           *bool `json:"report,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type UnreadResponseDataItemForum struct {
	ForumDescription       *string                                 `json:"forum_description,omitempty"`
	ForumID                *int                                    `json:"forum_id,omitempty"`
	ForumIsFollowed        *bool                                   `json:"forum_is_followed,omitempty"`
	ForumPostCount         *int                                    `json:"forum_post_count,omitempty"`
	ForumPrefixes          []interface{}                           `json:"forum_prefixes,omitempty"`
	ForumThreadCount       *int                                    `json:"forum_thread_count,omitempty"`
	ForumTitle             *string                                 `json:"forum_title,omitempty"`
	Links                  *UnreadResponseDataItemForumLinks       `json:"links,omitempty"`
	Permissions            *UnreadResponseDataItemForumPermissions `json:"permissions,omitempty"`
	ThreadDefaultPrefixID  *int                                    `json:"thread_default_prefix_id,omitempty"`
	ThreadPrefixIsRequired *bool                                   `json:"thread_prefix_is_required,omitempty"`
}

type UnreadResponseDataItemForumLinks struct {
	Detail        *string `json:"detail,omitempty"`
	Followers     *string `json:"followers,omitempty"`
	Permalink     *string `json:"permalink,omitempty"`
	SubCategories *string `json:"sub-categories,omitempty"`
	SubForums     *string `json:"sub-forums,omitempty"`
	Threads       *string `json:"threads,omitempty"`
}

type UnreadResponseDataItemForumPermissions struct {
	CreateThread     *bool `json:"create_thread,omitempty"`
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	TagThread        *bool `json:"tag_thread,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type UnreadResponseDataItemLinks struct {
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

type UnreadResponseDataItemPermissions struct {
	Delete           *bool `json:"delete,omitempty"`
	Edit             *bool `json:"edit,omitempty"`
	Follow           *bool `json:"follow,omitempty"`
	Post             *bool `json:"post,omitempty"`
	UploadAttachment *bool `json:"upload_attachment,omitempty"`
	View             *bool `json:"view,omitempty"`
}

type UnreadResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type UnreadResponseThreadsItem struct {
	Contest             *UnreadResponseThreadsItemContest      `json:"contest,omitempty"`
	CreatorUserID       *int                                   `json:"creator_user_id,omitempty"`
	CreatorUsername     *string                                `json:"creator_username,omitempty"`
	CreatorUsernameHTML *string                                `json:"creator_username_html,omitempty"`
	FirstPost           *UnreadResponseThreadsItemFirstPost    `json:"first_post,omitempty"`
	ForumID             *int                                   `json:"forum_id,omitempty"`
	LastPost            *UnreadResponseThreadsItemLastPost     `json:"last_post,omitempty"`
	Links               *UnreadResponseThreadsItemLinks        `json:"links,omitempty"`
	NodeTitle           *string                                `json:"node_title,omitempty"`
	Permissions         *UnreadResponseThreadsItemPermissions  `json:"permissions,omitempty"`
	Restrictions        *UnreadResponseThreadsItemRestrictions `json:"restrictions,omitempty"`
	ThreadCreateDate    *int                                   `json:"thread_create_date,omitempty"`
	ThreadID            *int                                   `json:"thread_id,omitempty"`
	ThreadIsClosed      *bool                                  `json:"thread_is_closed,omitempty"`
	ThreadIsDeleted     *bool                                  `json:"thread_is_deleted,omitempty"`
	ThreadIsFollowed    *bool                                  `json:"thread_is_followed,omitempty"`
	ThreadIsPublished   *bool                                  `json:"thread_is_published,omitempty"`
	ThreadIsStarred     *bool                                  `json:"thread_is_starred,omitempty"`
	ThreadIsSticky      *bool                                  `json:"thread_is_sticky,omitempty"`
	ThreadPostCount     *int                                   `json:"thread_post_count,omitempty"`
	ThreadPrefixes      []interface{}                          `json:"thread_prefixes,omitempty"`
	ThreadTags          interface{}                            `json:"thread_tags,omitempty"`
	ThreadTitle         *string                                `json:"thread_title,omitempty"`
	ThreadUpdateDate    *int                                   `json:"thread_update_date,omitempty"`
	ThreadViewCount     *int                                   `json:"thread_view_count,omitempty"`
	UserIsIgnored       *bool                                  `json:"user_is_ignored,omitempty"`
}

type UnreadResponseThreadsItemContest struct {
	AlreadyParticipate    *bool                                        `json:"already_participate,omitempty"`
	ChanceToWin           *float64                                     `json:"chance_to_win,omitempty"`
	CountWinners          *int                                         `json:"count_winners,omitempty"`
	FinishDate            *int                                         `json:"finish_date,omitempty"`
	IsFinished            *int                                         `json:"is_finished,omitempty"`
	IsMoneyPlaces         *int                                         `json:"is_money_places,omitempty"`
	NeededMembers         *int                                         `json:"needed_members,omitempty"`
	NowCountMembers       *int                                         `json:"now_count_members,omitempty"`
	Permissions           *UnreadResponseThreadsItemContestPermissions `json:"permissions,omitempty"`
	PrizeData             *int                                         `json:"prize_data,omitempty"`
	PrizeType             *string                                      `json:"prize_type,omitempty"`
	PrizeTypePhrase       *string                                      `json:"prize_type_phrase,omitempty"`
	RequireLikeCount      *int                                         `json:"require_like_count,omitempty"`
	RequireTotalLikeCount *int                                         `json:"require_total_like_count,omitempty"`
	Type_                 *string                                      `json:"type,omitempty"`
	Winners               []int                                        `json:"winners,omitempty"`
}

type UnreadResponseThreadsItemContestPermissions struct {
	CanFinish           *bool   `json:"can_finish,omitempty"`
	CanParticipate      *bool   `json:"can_participate,omitempty"`
	CanParticipateError *string `json:"can_participate_error,omitempty"`
	CanViewUserList     *bool   `json:"can_view_user_list,omitempty"`
}

type UnreadResponseThreadsItemFirstPost struct {
	Links              *UnreadResponseThreadsItemFirstPostLinks       `json:"links,omitempty"`
	Permissions        *UnreadResponseThreadsItemFirstPostPermissions `json:"permissions,omitempty"`
	PostBody           *string                                        `json:"post_body,omitempty"`
	PostBodyHTML       *string                                        `json:"post_body_html,omitempty"`
	PostBodyPlainText  *string                                        `json:"post_body_plain_text,omitempty"`
	PostCreateDate     *int                                           `json:"post_create_date,omitempty"`
	PostID             *int                                           `json:"post_id,omitempty"`
	PostIsDeleted      *bool                                          `json:"post_is_deleted,omitempty"`
	PostIsFirstPost    *bool                                          `json:"post_is_first_post,omitempty"`
	PostIsLiked        *bool                                          `json:"post_is_liked,omitempty"`
	PostIsPublished    *bool                                          `json:"post_is_published,omitempty"`
	PostLikeCount      *int                                           `json:"post_like_count,omitempty"`
	PostUpdateDate     *int                                           `json:"post_update_date,omitempty"`
	PosterUserID       *int                                           `json:"poster_user_id,omitempty"`
	PosterUsername     *string                                        `json:"poster_username,omitempty"`
	PosterUsernameHTML *string                                        `json:"poster_username_html,omitempty"`
	Signature          *string                                        `json:"signature,omitempty"`
	SignatureHTML      *string                                        `json:"signature_html,omitempty"`
	SignaturePlainText *string                                        `json:"signature_plain_text,omitempty"`
	ThreadID           *int                                           `json:"thread_id,omitempty"`
	ThreadIsDeleted    *bool                                          `json:"thread_is_deleted,omitempty"`
	UserIsIgnored      *bool                                          `json:"user_is_ignored,omitempty"`
}

type UnreadResponseThreadsItemFirstPostLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type UnreadResponseThreadsItemFirstPostPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	Like   *bool `json:"like,omitempty"`
	Reply  *bool `json:"reply,omitempty"`
	Report *bool `json:"report,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type UnreadResponseThreadsItemLastPost struct {
	Links              *UnreadResponseThreadsItemLastPostLinks       `json:"links,omitempty"`
	Permissions        *UnreadResponseThreadsItemLastPostPermissions `json:"permissions,omitempty"`
	PostBody           *string                                       `json:"post_body,omitempty"`
	PostBodyHTML       *string                                       `json:"post_body_html,omitempty"`
	PostBodyPlainText  *string                                       `json:"post_body_plain_text,omitempty"`
	PostCreateDate     *int                                          `json:"post_create_date,omitempty"`
	PostID             *int                                          `json:"post_id,omitempty"`
	PostIsDeleted      *bool                                         `json:"post_is_deleted,omitempty"`
	PostIsFirstPost    *bool                                         `json:"post_is_first_post,omitempty"`
	PostIsLiked        *bool                                         `json:"post_is_liked,omitempty"`
	PostIsPublished    *bool                                         `json:"post_is_published,omitempty"`
	PostLikeCount      *int                                          `json:"post_like_count,omitempty"`
	PostUpdateDate     *int                                          `json:"post_update_date,omitempty"`
	PosterUserID       *int                                          `json:"poster_user_id,omitempty"`
	PosterUsername     *string                                       `json:"poster_username,omitempty"`
	PosterUsernameHTML *string                                       `json:"poster_username_html,omitempty"`
	Signature          *string                                       `json:"signature,omitempty"`
	SignatureHTML      *string                                       `json:"signature_html,omitempty"`
	SignaturePlainText *string                                       `json:"signature_plain_text,omitempty"`
	ThreadID           *int                                          `json:"thread_id,omitempty"`
	ThreadIsDeleted    *bool                                         `json:"thread_is_deleted,omitempty"`
	UserIsIgnored      *bool                                         `json:"user_is_ignored,omitempty"`
}

type UnreadResponseThreadsItemLastPostLinks struct {
	Detail       *string `json:"detail,omitempty"`
	Likes        *string `json:"likes,omitempty"`
	Permalink    *string `json:"permalink,omitempty"`
	Poster       *string `json:"poster,omitempty"`
	PosterAvatar *string `json:"poster_avatar,omitempty"`
	Report       *string `json:"report,omitempty"`
	Thread       *string `json:"thread,omitempty"`
}

type UnreadResponseThreadsItemLastPostPermissions struct {
	Delete *bool `json:"delete,omitempty"`
	Edit   *bool `json:"edit,omitempty"`
	Like   *bool `json:"like,omitempty"`
	Reply  *bool `json:"reply,omitempty"`
	Report *bool `json:"report,omitempty"`
	View   *bool `json:"view,omitempty"`
}

type UnreadResponseThreadsItemLinks struct {
	Detail            *string `json:"detail,omitempty"`
	FirstPost         *string `json:"first_post,omitempty"`
	FirstPoster       *string `json:"first_poster,omitempty"`
	FirstPosterAvatar *string `json:"first_poster_avatar,omitempty"`
	Followers         *string `json:"followers,omitempty"`
	Forum             *string `json:"forum,omitempty"`
	LastPost          *string `json:"last_post,omitempty"`
	Permalink         *string `json:"permalink,omitempty"`
	Posts             *string `json:"posts,omitempty"`
}

type UnreadResponseThreadsItemPermissions struct {
	Bump      *UnreadResponseThreadsItemPermissionsBump `json:"bump,omitempty"`
	Delete    *bool                                     `json:"delete,omitempty"`
	Edit      *bool                                     `json:"edit,omitempty"`
	EditTags  *bool                                     `json:"edit_tags,omitempty"`
	EditTitle *bool                                     `json:"edit_title,omitempty"`
	Follow    *bool                                     `json:"follow,omitempty"`
	Post      *bool                                     `json:"post,omitempty"`
	View      *bool                                     `json:"view,omitempty"`
}

type UnreadResponseThreadsItemPermissionsBump struct {
	AvailableCount    *int        `json:"available_count,omitempty"`
	Can               *bool       `json:"can,omitempty"`
	Error             interface{} `json:"error,omitempty"`
	NextAvailableTime interface{} `json:"next_available_time,omitempty"`
}

type UnreadResponseThreadsItemRestrictions struct {
	MaxReplyCount *int `json:"max_reply_count,omitempty"`
	ReplyDelay    *int `json:"reply_delay,omitempty"`
}

type UnstarResponse struct {
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type UnstarResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type UpdateParams struct {
	// Allow members to delete their own messages.
	AllowDeleteOwnMessages *bool `json:"allow_delete_own_messages,omitempty"`
	// Allow members to edit their own messages.
	AllowEditMessages *bool `json:"allow_edit_messages,omitempty"`
	// Allow members to stick messages.
	AllowStickyMessages *bool `json:"allow_sticky_messages,omitempty"`
	// Make conversation history visible to new members.
	HistoryOpen *bool `json:"history_open,omitempty"`
	// Allow members to invite others.
	OpenInvite *bool `json:"open_invite,omitempty"`
	// New conversation title.
	Title *string `json:"title,omitempty"`
}

type UpdateResponse struct {
	Conversation *RespConversationModel `json:"conversation,omitempty"`
	SystemInfo   *RespSystemInfo        `json:"system_info,omitempty"`
}

type UpdateResponseConversation struct {
	Alerts                   *int                                       `json:"alerts,omitempty"`
	ConversationCreateDate   *int                                       `json:"conversation_create_date,omitempty"`
	ConversationID           *int                                       `json:"conversation_id,omitempty"`
	ConversationIsDeleted    *bool                                      `json:"conversation_is_deleted,omitempty"`
	ConversationIsNew        *bool                                      `json:"conversation_is_new,omitempty"`
	ConversationIsOpen       *bool                                      `json:"conversation_is_open,omitempty"`
	ConversationLastReadDate *int                                       `json:"conversation_last_read_date,omitempty"`
	ConversationMessageCount *int                                       `json:"conversation_message_count,omitempty"`
	ConversationOnlineCount  *int                                       `json:"conversation_online_count,omitempty"`
	ConversationTitle        *string                                    `json:"conversation_title,omitempty"`
	ConversationUpdateDate   *int                                       `json:"conversation_update_date,omitempty"`
	CreatorIsIgnored         *bool                                      `json:"creator_is_ignored,omitempty"`
	CreatorUserID            *int                                       `json:"creator_user_id,omitempty"`
	CreatorUsername          *string                                    `json:"creator_username,omitempty"`
	CreatorUsernameHTML      *string                                    `json:"creator_username_html,omitempty"`
	IsGroup                  *int                                       `json:"is_group,omitempty"`
	IsStarred                *int                                       `json:"is_starred,omitempty"`
	IsUnread                 *int                                       `json:"is_unread,omitempty"`
	Links                    *UpdateResponseConversationLinks           `json:"links,omitempty"`
	Permissions              *UpdateResponseConversationPermissions     `json:"permissions,omitempty"`
	Recipient                *UpdateResponseConversationRecipient       `json:"recipient,omitempty"`
	Recipients               []UpdateResponseConversationRecipientsItem `json:"recipients,omitempty"`
}

type UpdateResponseConversationLinks struct {
	Avatar    *string `json:"avatar,omitempty"`
	Detail    *string `json:"detail,omitempty"`
	Messages  *string `json:"messages,omitempty"`
	Permalink *string `json:"permalink,omitempty"`
}

type UpdateResponseConversationPermissions struct {
	EditOwnPost       *bool `json:"editOwnPost,omitempty"`
	Invite            *bool `json:"invite,omitempty"`
	Kick              *bool `json:"kick,omitempty"`
	ManageInviteLinks *bool `json:"manage_invite_links,omitempty"`
	Reply             *bool `json:"reply,omitempty"`
	StickyMessages    *bool `json:"stickyMessages,omitempty"`
	UploadAvatar      *bool `json:"upload_avatar,omitempty"`
	View              *bool `json:"view,omitempty"`
}

type UpdateResponseConversationRecipient struct {
	Avatar          *string `json:"avatar,omitempty"`
	ContactsChanged *bool   `json:"contacts_changed,omitempty"`
	IsOnline        *bool   `json:"is_online,omitempty"`
	LastActivity    *int    `json:"last_activity,omitempty"`
	UserID          *int    `json:"user_id,omitempty"`
	Username        *string `json:"username,omitempty"`
	UsernameHTML    *string `json:"username_html,omitempty"`
}

type UpdateResponseConversationRecipientsItem struct {
	Avatar          *string `json:"avatar,omitempty"`
	ContactsChanged *bool   `json:"contacts_changed,omitempty"`
	IsOnline        *bool   `json:"is_online,omitempty"`
	LastActivity    *int    `json:"last_activity,omitempty"`
	UserID          *int    `json:"user_id,omitempty"`
	Username        *string `json:"username,omitempty"`
	UsernameHTML    *string `json:"username_html,omitempty"`
}

type UpdateResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type UploadParams struct {
	// Selection size.
	Crop *int `json:"crop,omitempty"`
	// The starting point of the selection by width. Default value - 0
	X *int `json:"x,omitempty"`
	// The starting point of the selection by height. Default value - 0
	Y *int `json:"y,omitempty"`
}

type UploadPostParams struct {
	// Selection size.
	Crop *int `json:"crop,omitempty"`
	// The starting point of the selection by width. Default value - 0
	X *int `json:"x,omitempty"`
	// The starting point of the selection by height. Default value - 0
	Y *int `json:"y,omitempty"`
}

type UploadPostResponse struct {
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type UploadPostResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type UploadResponse struct {
	Message    *string         `json:"message,omitempty"`
	Status     *string         `json:"status,omitempty"`
	SystemInfo *RespSystemInfo `json:"system_info,omitempty"`
}

type UploadResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type UsersParams struct {
	// Search query.
	Q *string `json:"q,omitempty"`
}

type UsersResponse struct {
	SystemInfo *RespSystemInfo          `json:"system_info,omitempty"`
	Users      []UsersResponseUsersItem `json:"users,omitempty"`
}

type UsersResponseSystemInfo struct {
	Time      *int `json:"time,omitempty"`
	VisitorID *int `json:"visitor_id,omitempty"`
}

type UsersResponseUsersItem struct {
	Balance                     *string                                                 `json:"balance,omitempty"`
	Banner                      *string                                                 `json:"banner,omitempty"`
	Birthday                    *UsersResponseUsersItemBirthday                         `json:"birthday,omitempty"`
	ContestCount                *int                                                    `json:"contest_count,omitempty"`
	ConvWelcomeMessage          *string                                                 `json:"conv_welcome_message,omitempty"`
	CuratorTitles               []string                                                `json:"curator_titles,omitempty"`
	Currency                    *string                                                 `json:"currency,omitempty"`
	CustomTitle                 *string                                                 `json:"custom_title,omitempty"`
	DisplayBannerID             *int                                                    `json:"display_banner_id,omitempty"`
	DisplayIconGroupID          *int                                                    `json:"display_icon_group_id,omitempty"`
	EditPermissions             *UsersResponseUsersItemEditPermissions                  `json:"edit_permissions,omitempty"`
	Fields                      []UsersResponseUsersItemFieldsItem                      `json:"fields,omitempty"`
	Hold                        *string                                                 `json:"hold,omitempty"`
	IsBanned                    *int                                                    `json:"is_banned,omitempty"`
	Links                       *UsersResponseUsersItemLinks                            `json:"links,omitempty"`
	Permissions                 *UsersResponseUsersItemPermissions                      `json:"permissions,omitempty"`
	SecretAnswerFirstLetter     *string                                                 `json:"secret_answer_first_letter,omitempty"`
	SecretAnswerRendered        *string                                                 `json:"secret_answer_rendered,omitempty"`
	SelfPermissions             *UsersResponseUsersItemSelfPermissions                  `json:"self_permissions,omitempty"`
	ShortLink                   *string                                                 `json:"short_link,omitempty"`
	TrophyCount                 *int                                                    `json:"trophy_count,omitempty"`
	UserDeposit                 *int                                                    `json:"user_deposit,omitempty"`
	UserEmail                   *string                                                 `json:"user_email,omitempty"`
	UserExternalAuthentications []UsersResponseUsersItemUserExternalAuthenticationsItem `json:"user_external_authentications,omitempty"`
	UserFollowers               *UsersResponseUsersItemUserFollowers                    `json:"user_followers,omitempty"`
	UserFollowing               *UsersResponseUsersItemUserFollowing                    `json:"user_following,omitempty"`
	UserGroupID                 *int                                                    `json:"user_group_id,omitempty"`
	UserGroups                  []UsersResponseUsersItemUserGroupsItem                  `json:"user_groups,omitempty"`
	UserID                      *int                                                    `json:"user_id,omitempty"`
	UserIsFollowed              *bool                                                   `json:"user_is_followed,omitempty"`
	UserIsIgnored               *bool                                                   `json:"user_is_ignored,omitempty"`
	UserIsValid                 *bool                                                   `json:"user_is_valid,omitempty"`
	UserIsVerified              *bool                                                   `json:"user_is_verified,omitempty"`
	UserIsVisitor               *bool                                                   `json:"user_is_visitor,omitempty"`
	UserLastSeenDate            *int                                                    `json:"user_last_seen_date,omitempty"`
	UserLike2Count              *int                                                    `json:"user_like2_count,omitempty"`
	UserLikeCount               *int                                                    `json:"user_like_count,omitempty"`
	UserMessageCount            *int                                                    `json:"user_message_count,omitempty"`
	UserRegisterDate            *int                                                    `json:"user_register_date,omitempty"`
	UserTimezoneOffset          *int                                                    `json:"user_timezone_offset,omitempty"`
	UserTitle                   *string                                                 `json:"user_title,omitempty"`
	UserUnreadConversationCount *int                                                    `json:"user_unread_conversation_count,omitempty"`
	UserUnreadNotificationCount *int                                                    `json:"user_unread_notification_count,omitempty"`
	Username                    *string                                                 `json:"username,omitempty"`
	UsernameHTML                *string                                                 `json:"username_html,omitempty"`
}

type UsersResponseUsersItemBirthday struct {
	Age       *int                                     `json:"age,omitempty"`
	Format    *string                                  `json:"format,omitempty"`
	TimeStamp *UsersResponseUsersItemBirthdayTimeStamp `json:"timeStamp,omitempty"`
}

type UsersResponseUsersItemBirthdayTimeStamp struct {
	Date         *string `json:"date,omitempty"`
	Timezone     *string `json:"timezone,omitempty"`
	TimezoneType *int    `json:"timezone_type,omitempty"`
}

type UsersResponseUsersItemEditPermissions struct {
	Fields            *bool `json:"fields,omitempty"`
	HideUsernameLogs  *bool `json:"hide_username_logs,omitempty"`
	Password          *bool `json:"password,omitempty"`
	PrimaryGroupID    *bool `json:"primary_group_id,omitempty"`
	SecondaryGroupIds *bool `json:"secondary_group_ids,omitempty"`
	ShortLink         *bool `json:"short_link,omitempty"`
	UserDobDay        *bool `json:"user_dob_day,omitempty"`
	UserDobMonth      *bool `json:"user_dob_month,omitempty"`
	UserDobYear       *bool `json:"user_dob_year,omitempty"`
	UserEmail         *bool `json:"user_email,omitempty"`
	UserTitle         *bool `json:"user_title,omitempty"`
	Username          *bool `json:"username,omitempty"`
}

type UsersResponseUsersItemFieldsItem struct {
	Choices       []UsersResponseUsersItemFieldsItemChoicesItem `json:"choices,omitempty"`
	Description   *string                                       `json:"description,omitempty"`
	ID            *string                                       `json:"id,omitempty"`
	IsMultiChoice *bool                                         `json:"is_multi_choice,omitempty"`
	IsRequired    *bool                                         `json:"is_required,omitempty"`
	Position      *string                                       `json:"position,omitempty"`
	Title         *string                                       `json:"title,omitempty"`
	Value         *string                                       `json:"value,omitempty"`
	Values        interface{}                                   `json:"values,omitempty"`
}

type UsersResponseUsersItemFieldsItemChoicesItem struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type UsersResponseUsersItemLinks struct {
	Avatar      *string `json:"avatar,omitempty"`
	AvatarBig   *string `json:"avatar_big,omitempty"`
	AvatarSmall *string `json:"avatar_small,omitempty"`
	BackgroundL *string `json:"background_l,omitempty"`
	BackgroundM *string `json:"background_m,omitempty"`
	Detail      *string `json:"detail,omitempty"`
	Followers   *string `json:"followers,omitempty"`
	Followings  *string `json:"followings,omitempty"`
	Ignore      *string `json:"ignore,omitempty"`
	Permalink   *string `json:"permalink,omitempty"`
	Status      *string `json:"status,omitempty"`
	Timeline    *string `json:"timeline,omitempty"`
}

type UsersResponseUsersItemPermissions struct {
	Edit        *bool `json:"edit,omitempty"`
	Follow      *bool `json:"follow,omitempty"`
	Ignore      *bool `json:"ignore,omitempty"`
	ProfilePost *bool `json:"profile_post,omitempty"`
}

type UsersResponseUsersItemSelfPermissions struct {
	CreateConversation *bool `json:"create_conversation,omitempty"`
}

type UsersResponseUsersItemUserExternalAuthenticationsItem struct {
	Provider    *string `json:"provider,omitempty"`
	ProviderKey *string `json:"provider_key,omitempty"`
}

type UsersResponseUsersItemUserFollowers struct {
	Count *int                                           `json:"count,omitempty"`
	Users []UsersResponseUsersItemUserFollowersUsersItem `json:"users,omitempty"`
}

type UsersResponseUsersItemUserFollowersUsersItem struct {
	Avatar       *string `json:"avatar,omitempty"`
	UserID       *int    `json:"user_id,omitempty"`
	Username     *string `json:"username,omitempty"`
	UsernameHTML *string `json:"username_html,omitempty"`
}

type UsersResponseUsersItemUserFollowing struct {
	Count *int                                           `json:"count,omitempty"`
	Users []UsersResponseUsersItemUserFollowingUsersItem `json:"users,omitempty"`
}

type UsersResponseUsersItemUserFollowingUsersItem struct {
	Avatar       *string `json:"avatar,omitempty"`
	UserID       *int    `json:"user_id,omitempty"`
	Username     *string `json:"username,omitempty"`
	UsernameHTML *string `json:"username_html,omitempty"`
}

type UsersResponseUsersItemUserGroupsItem struct {
	DisplayBannerSelectable *bool   `json:"display_banner_selectable,omitempty"`
	DisplayGroupSelectable  *bool   `json:"display_group_selectable,omitempty"`
	DisplayIconSelectable   *bool   `json:"display_icon_selectable,omitempty"`
	IsPrimaryGroup          *bool   `json:"is_primary_group,omitempty"`
	UserGroupBannerCSSClass *string `json:"user_group_banner_css_class,omitempty"`
	UserGroupBannerText     *string `json:"user_group_banner_text,omitempty"`
	UserGroupBannerTextEn   *string `json:"user_group_banner_text_en,omitempty"`
	UserGroupIconClass      *string `json:"user_group_icon_class,omitempty"`
	UserGroupID             *int    `json:"user_group_id,omitempty"`
	UserGroupTitle          *string `json:"user_group_title,omitempty"`
	UserGroupTitleEn        *string `json:"user_group_title_en,omitempty"`
}

type VoteParams struct {
	// The id of the response to vote for. Can be skipped if **response_ids** set.
	ResponseID *int `json:"response_id,omitempty"`
	// An array of ids of responses (if the poll allows multiple choices).
	ResponseIds []int `json:"response_ids,omitempty"`
}

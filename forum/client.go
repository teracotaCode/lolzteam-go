package forum

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// Requester is the interface for making HTTP requests.
type Requester interface {
	Request(ctx context.Context, method, path string, opts RequestOptions) (json.RawMessage, error)
}

// RequestOptions holds options for a request.
type RequestOptions struct {
	Params         map[string]interface{}
	JSON           map[string]interface{}
	JSONBody       interface{}
	Files          map[string]FileUpload
	Search         bool
	ForceMultipart bool
}

// FileUpload represents a file to upload.
type FileUpload struct {
	Filename string
	Data     []byte
}

// ForumClient is the API client.
type ForumClient struct {
	r                   Requester
	Assets              *AssetsService
	Authentication      *AuthenticationService
	BatchRequests       *BatchRequestsService
	Categories          *CategoriesService
	Chatbox             *ChatboxService
	ContentTagging      *ContentTaggingService
	Conversations       *ConversationsService
	Forms               *FormsService
	Forums              *ForumsService
	LinkForums          *LinkForumsService
	Navigation          *NavigationService
	Notifications       *NotificationsService
	Pages               *PagesService
	PostComments        *PostCommentsService
	Posts               *PostsService
	ProfilePostComments *ProfilePostCommentsService
	ProfilePosts        *ProfilePostsService
	Searching           *SearchingService
	Threads             *ThreadsService
	Users               *UsersService
}

// New creates a new ForumClient.
func New(r Requester) *ForumClient {
	c := &ForumClient{r: r}
	c.Assets = &AssetsService{r: r}
	c.Authentication = &AuthenticationService{r: r}
	c.BatchRequests = &BatchRequestsService{r: r}
	c.Categories = &CategoriesService{r: r}
	c.Chatbox = &ChatboxService{r: r}
	c.ContentTagging = &ContentTaggingService{r: r}
	c.Conversations = &ConversationsService{r: r}
	c.Forms = &FormsService{r: r}
	c.Forums = &ForumsService{r: r}
	c.LinkForums = &LinkForumsService{r: r}
	c.Navigation = &NavigationService{r: r}
	c.Notifications = &NotificationsService{r: r}
	c.Pages = &PagesService{r: r}
	c.PostComments = &PostCommentsService{r: r}
	c.Posts = &PostsService{r: r}
	c.ProfilePostComments = &ProfilePostCommentsService{r: r}
	c.ProfilePosts = &ProfilePostsService{r: r}
	c.Searching = &SearchingService{r: r}
	c.Threads = &ThreadsService{r: r}
	c.Users = &UsersService{r: r}
	return c
}

// AssetsService handles Assets operations.
type AssetsService struct {
	r Requester
}

// CSS Get CSS
func (s *AssetsService) CSS(ctx context.Context, params *CSSParams) (*CSSResponse, error) {
	path := "/css"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.CSS != nil {
			opts.Params["css"] = params.CSS
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result CSSResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AuthenticationService handles Authentication operations.
type AuthenticationService struct {
	r Requester
}

// Token Get Access Token
func (s *AuthenticationService) Token(ctx context.Context, params *TokenParams) (*TokenResponse, error) {
	path := "/oauth/token"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.ClientID != nil {
			opts.JSON["client_id"] = *params.ClientID
		}
		if params.ClientSecret != nil {
			opts.JSON["client_secret"] = *params.ClientSecret
		}
		if params.Code != nil {
			opts.JSON["code"] = *params.Code
		}
		if params.GrantType != nil {
			opts.JSON["grant_type"] = *params.GrantType
		}
		if params.Password != nil {
			opts.JSON["password"] = *params.Password
		}
		if params.RedirectURI != nil {
			opts.JSON["redirect_uri"] = *params.RedirectURI
		}
		if params.RefreshToken != nil {
			opts.JSON["refresh_token"] = *params.RefreshToken
		}
		if params.Scope != nil {
			opts.JSON["scope"] = params.Scope
		}
		if params.Username != nil {
			opts.JSON["username"] = *params.Username
		}
	}
	opts.ForceMultipart = true
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result TokenResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchRequestsService handles BatchRequests operations.
type BatchRequestsService struct {
	r Requester
}

// Execute Batch
func (s *BatchRequestsService) Execute(ctx context.Context, jobs []map[string]interface{}) (*ExecuteResponse, error) {
	path := "/batch"
	opts := RequestOptions{}
	opts.JSONBody = jobs
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result ExecuteResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CategoriesService handles Categories operations.
type CategoriesService struct {
	r Requester
}

// Get Get Category
func (s *CategoriesService) Get(ctx context.Context, categoryID int) (*GetResponse, error) {
	path := "/categories/{category_id}"
	path = strings.Replace(path, "{category_id}", fmt.Sprintf("%d", categoryID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Categories
func (s *CategoriesService) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/categories"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.PostsUserID != nil {
			opts.Params["posts_user_id"] = *params.PostsUserID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ChatboxService handles Chatbox operations.
type ChatboxService struct {
	r Requester
}

// DeleteIgnore Unignore Chat User
func (s *ChatboxService) DeleteIgnore(ctx context.Context, userID StringOrInt) (*SaveChanges, error) {
	path := "/chatbox/ignore"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["user_id"] = userID
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteMessage Delete Chat Message
func (s *ChatboxService) DeleteMessage(ctx context.Context, messageID int) (*SaveChanges, error) {
	path := "/chatbox/messages"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["message_id"] = messageID
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// EditMessage Edit Chat Message
func (s *ChatboxService) EditMessage(ctx context.Context, message string, messageID int) (*EditMessageResponse, error) {
	path := "/chatbox/messages"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["message"] = message
	opts.JSON["message_id"] = messageID
	raw, err := s.r.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}
	var result EditMessageResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetIgnore Get Ignored Chat Users
func (s *ChatboxService) GetIgnore(ctx context.Context) (*GetIgnoreResponse, error) {
	path := "/chatbox/ignore"
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetIgnoreResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetLeaderboard Get Chat Leaderboard
func (s *ChatboxService) GetLeaderboard(ctx context.Context, params *GetLeaderboardParams) (*GetLeaderboardResponse, error) {
	path := "/chatbox/messages/leaderboard"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Duration != nil {
			opts.Params["duration"] = *params.Duration
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetLeaderboardResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMessages Get Chat Messages
func (s *ChatboxService) GetMessages(ctx context.Context, roomID ChatboxRoomID, params *GetMessagesParams) (*GetMessagesResponse, error) {
	path := "/chatbox/messages"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	opts.Params["room_id"] = roomID
	if params != nil {
		if params.BeforeMessageID != nil {
			opts.Params["before_message_id"] = *params.BeforeMessageID
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetMessagesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Index Get Chats
func (s *ChatboxService) Index(ctx context.Context, params *IndexParams) (*IndexResponse, error) {
	path := "/chatbox"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.RoomID != nil {
			opts.Params["room_id"] = *params.RoomID
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result IndexResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Online Get Chat Online
func (s *ChatboxService) Online(ctx context.Context, roomID ChatboxRoomID) (*OnlineResponse, error) {
	path := "/chatbox/messages/online"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	opts.Params["room_id"] = roomID
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result OnlineResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PostIgnore Ignore Chat User
func (s *ChatboxService) PostIgnore(ctx context.Context, userID StringOrInt) (*SaveChanges, error) {
	path := "/chatbox/ignore"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["user_id"] = userID
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PostMessage Create Chat Message
func (s *ChatboxService) PostMessage(ctx context.Context, message string, roomID ChatboxPostMessageRoomID, params *PostMessageParams) (*PostMessageResponse, error) {
	path := "/chatbox/messages"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["message"] = message
	opts.JSON["room_id"] = roomID
	if params != nil {
		if params.ReplyMessageID != nil {
			opts.JSON["reply_message_id"] = *params.ReplyMessageID
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result PostMessageResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Report Report Chat Message
func (s *ChatboxService) Report(ctx context.Context, messageID int, reason string) (*SaveChanges, error) {
	path := "/chatbox/messages/report"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["message_id"] = messageID
	opts.JSON["reason"] = reason
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReportReasons Get Chat Message Report Reasons
func (s *ChatboxService) ReportReasons(ctx context.Context, messageID int) (*ReportReasonsResponse, error) {
	path := "/chatbox/messages/report"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	opts.Params["message_id"] = messageID
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ReportReasonsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ContentTaggingService handles ContentTagging operations.
type ContentTaggingService struct {
	r Requester
}

// Find Get Filtered Content
func (s *ContentTaggingService) Find(ctx context.Context, tag string) (*FindResponse, error) {
	path := "/tags/find"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	opts.Params["tag"] = tag
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result FindResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get Get Tagged Content
func (s *ContentTaggingService) Get(ctx context.Context, tagID int, params *GetParams) (*GetResponse, error) {
	path := "/tags/{tag_id}"
	path = strings.Replace(path, "{tag_id}", fmt.Sprintf("%d", tagID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Tags
func (s *ContentTaggingService) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/tags/list"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.PostsUserID != nil {
			opts.Params["posts_user_id"] = *params.PostsUserID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Popular Get Popular Tags
func (s *ContentTaggingService) Popular(ctx context.Context) (*PopularResponse, error) {
	path := "/tags"
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result PopularResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ConversationsService handles Conversations operations.
type ConversationsService struct {
	r Requester
}

// Create Create Conversation
func (s *ConversationsService) Create(ctx context.Context, params *CreateParams) (*CreateResponse, error) {
	path := "/conversations"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.AllowAskHiddenContent != nil {
			opts.JSON["allow_ask_hidden_content"] = *params.AllowAskHiddenContent
		}
		if params.CommentIgnoreGroup != nil {
			opts.JSON["comment_ignore_group"] = *params.CommentIgnoreGroup
		}
		if params.DontAlertFollowers != nil {
			opts.JSON["dont_alert_followers"] = *params.DontAlertFollowers
		}
		if params.HideContacts != nil {
			opts.JSON["hide_contacts"] = *params.HideContacts
		}
		if params.PrefixID != nil {
			opts.JSON["prefix_id"] = params.PrefixID
		}
		if params.ReplyGroup != nil {
			opts.JSON["reply_group"] = *params.ReplyGroup
		}
		if params.ScheduleDate != nil {
			opts.JSON["schedule_date"] = *params.ScheduleDate
		}
		if params.ScheduleTime != nil {
			opts.JSON["schedule_time"] = *params.ScheduleTime
		}
		if params.Tags != nil {
			opts.JSON["tags"] = params.Tags
		}
		if params.Title != nil {
			opts.JSON["title"] = *params.Title
		}
		if params.TitleEn != nil {
			opts.JSON["title_en"] = *params.TitleEn
		}
		if params.WatchThread != nil {
			opts.JSON["watch_thread"] = *params.WatchThread
		}
		if params.WatchThreadEmail != nil {
			opts.JSON["watch_thread_email"] = *params.WatchThreadEmail
		}
		if params.WatchThreadState != nil {
			opts.JSON["watch_thread_state"] = *params.WatchThreadState
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result CreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePost Create Conversation Message
func (s *ConversationsService) CreatePost(ctx context.Context, conversationID int, messageBody string, params *CreatePostParams) (*CreatePostResponse, error) {
	path := "/conversations/{conversation_id}/messages"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%d", conversationID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["message_body"] = messageBody
	if params != nil {
		if params.ReplyMessageID != nil {
			opts.JSON["reply_message_id"] = *params.ReplyMessageID
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result CreatePostResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete Leave Conversation
func (s *ConversationsService) Delete(ctx context.Context, conversationID int, deleteType ConversationsDeleteDeleteType) (*SaveChanges, error) {
	path := "/conversations"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["conversation_id"] = conversationID
	opts.JSON["delete_type"] = deleteType
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDelete Delete Conversation Message
func (s *ConversationsService) DeleteDelete(ctx context.Context, conversationID int, messageID int) (*SaveChanges, error) {
	path := "/conversations/{conversation_id}/messages/{message_id}"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%d", conversationID), 1)
	path = strings.Replace(path, "{message_id}", fmt.Sprintf("%d", messageID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Disable Disable Conversation Alerts
func (s *ConversationsService) Disable(ctx context.Context, conversationID int) (*DisableResponse, error) {
	path := "/conversations/{conversation_id}/alerts"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%d", conversationID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result DisableResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Edit Edit Conversation Message
func (s *ConversationsService) Edit(ctx context.Context, conversationID int, messageID int, messageBody string) (*EditResponse, error) {
	path := "/conversations/{conversation_id}/messages/{message_id}"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%d", conversationID), 1)
	path = strings.Replace(path, "{message_id}", fmt.Sprintf("%d", messageID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["message_body"] = messageBody
	raw, err := s.r.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}
	var result EditResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Enable Enable Conversation Alerts
func (s *ConversationsService) Enable(ctx context.Context, conversationID int) (*EnableResponse, error) {
	path := "/conversations/{conversation_id}/alerts"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%d", conversationID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result EnableResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get Get Conversation Message
func (s *ConversationsService) Get(ctx context.Context, messageID int) (*GetResponse, error) {
	path := "/conversations/messages/{message_id}"
	path = strings.Replace(path, "{message_id}", fmt.Sprintf("%d", messageID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetGet Get Conversation
func (s *ConversationsService) GetGet(ctx context.Context, conversationID int) (*GetGetResponse, error) {
	path := "/conversations/{conversation_id}"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%d", conversationID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetGetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Invite Invite Users to Conversation
func (s *ConversationsService) Invite(ctx context.Context, conversationID int, recipients []string) (*SaveChanges, error) {
	path := "/conversations/{conversation_id}/invite"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%d", conversationID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["recipients"] = recipients
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Kick Kick User from Conversation
func (s *ConversationsService) Kick(ctx context.Context, conversationID int, userID int) (*SaveChanges, error) {
	path := "/conversations/{conversation_id}/kick"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%d", conversationID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["user_id"] = userID
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Conversations
func (s *ConversationsService) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/conversations"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.PostsUserID != nil {
			opts.Params["posts_user_id"] = *params.PostsUserID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ListGet Get Conversation Messages
func (s *ConversationsService) ListGet(ctx context.Context, conversationID int, params *ListGetParams) (*ListGetResponse, error) {
	path := "/conversations/{conversation_id}/messages"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%d", conversationID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.Order != nil {
			opts.Params["order"] = *params.Order
		}
		if params.Before != nil {
			opts.Params["before"] = *params.Before
		}
		if params.After != nil {
			opts.Params["after"] = *params.After
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListGetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Read Read a Conversation
func (s *ConversationsService) Read(ctx context.Context, conversationID int) (*SaveChanges, error) {
	path := "/conversations/{conversation_id}/read"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%d", conversationID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReadAll Read All Conversations
func (s *ConversationsService) ReadAll(ctx context.Context) (*ReadAllResponse, error) {
	path := "/conversations/read-all"
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result ReadAllResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Save Send Content To Saved Messages
func (s *ConversationsService) Save(ctx context.Context, link string) (*SaveChanges, error) {
	path := "/conversations/save"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["link"] = link
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Search Search Conversations Messages
func (s *ConversationsService) Search(ctx context.Context, params *SearchParams) (*SearchResponse, error) {
	path := "/conversations/search"
	opts := RequestOptions{}
	opts.Search = true
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.ConversationID != nil {
			opts.JSON["conversation_id"] = *params.ConversationID
		}
		if params.Q != nil {
			opts.JSON["q"] = *params.Q
		}
		if params.SearchRecipients != nil {
			opts.JSON["search_recipients"] = *params.SearchRecipients
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SearchResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Star Star Conversation
func (s *ConversationsService) Star(ctx context.Context, conversationID int) (*StarResponse, error) {
	path := "/conversations/{conversation_id}/star"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%d", conversationID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result StarResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Start Start Conversation
func (s *ConversationsService) Start(ctx context.Context, userID StringOrInt) (*StartResponse, error) {
	path := "/conversations/start"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["user_id"] = userID
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result StartResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Stick Stick Conversation Message
func (s *ConversationsService) Stick(ctx context.Context, conversationID int, messageID int) (*SaveChanges, error) {
	path := "/conversations/{conversation_id}/messages/{message_id}/stick"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%d", conversationID), 1)
	path = strings.Replace(path, "{message_id}", fmt.Sprintf("%d", messageID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Unstar Unstar Conversation
func (s *ConversationsService) Unstar(ctx context.Context, conversationID int) (*UnstarResponse, error) {
	path := "/conversations/{conversation_id}/star"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%d", conversationID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result UnstarResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Unstick Unstick Conversation Message
func (s *ConversationsService) Unstick(ctx context.Context, conversationID int, messageID int) (*SaveChanges, error) {
	path := "/conversations/{conversation_id}/messages/{message_id}/stick"
	path = strings.Replace(path, "{conversation_id}", fmt.Sprintf("%d", conversationID), 1)
	path = strings.Replace(path, "{message_id}", fmt.Sprintf("%d", messageID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Update Edit Conversation
func (s *ConversationsService) Update(ctx context.Context, conversationID int, params *UpdateParams) (*UpdateResponse, error) {
	path := "/conversations"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["conversation_id"] = conversationID
	if params != nil {
		if params.AllowDeleteOwnMessages != nil {
			opts.JSON["allow_delete_own_messages"] = *params.AllowDeleteOwnMessages
		}
		if params.AllowEditMessages != nil {
			opts.JSON["allow_edit_messages"] = *params.AllowEditMessages
		}
		if params.AllowStickyMessages != nil {
			opts.JSON["allow_sticky_messages"] = *params.AllowStickyMessages
		}
		if params.HistoryOpen != nil {
			opts.JSON["history_open"] = *params.HistoryOpen
		}
		if params.OpenInvite != nil {
			opts.JSON["open_invite"] = *params.OpenInvite
		}
		if params.Title != nil {
			opts.JSON["title"] = *params.Title
		}
	}
	raw, err := s.r.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}
	var result UpdateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// FormsService handles Forms operations.
type FormsService struct {
	r Requester
}

// Create Create Form
func (s *FormsService) Create(ctx context.Context, params *CreateParams) (*CreateResponse, error) {
	path := "/forms/save"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.AllowAskHiddenContent != nil {
			opts.JSON["allow_ask_hidden_content"] = *params.AllowAskHiddenContent
		}
		if params.CommentIgnoreGroup != nil {
			opts.JSON["comment_ignore_group"] = *params.CommentIgnoreGroup
		}
		if params.DontAlertFollowers != nil {
			opts.JSON["dont_alert_followers"] = *params.DontAlertFollowers
		}
		if params.HideContacts != nil {
			opts.JSON["hide_contacts"] = *params.HideContacts
		}
		if params.PrefixID != nil {
			opts.JSON["prefix_id"] = params.PrefixID
		}
		if params.ReplyGroup != nil {
			opts.JSON["reply_group"] = *params.ReplyGroup
		}
		if params.ScheduleDate != nil {
			opts.JSON["schedule_date"] = *params.ScheduleDate
		}
		if params.ScheduleTime != nil {
			opts.JSON["schedule_time"] = *params.ScheduleTime
		}
		if params.Tags != nil {
			opts.JSON["tags"] = params.Tags
		}
		if params.Title != nil {
			opts.JSON["title"] = *params.Title
		}
		if params.TitleEn != nil {
			opts.JSON["title_en"] = *params.TitleEn
		}
		if params.WatchThread != nil {
			opts.JSON["watch_thread"] = *params.WatchThread
		}
		if params.WatchThreadEmail != nil {
			opts.JSON["watch_thread_email"] = *params.WatchThreadEmail
		}
		if params.WatchThreadState != nil {
			opts.JSON["watch_thread_state"] = *params.WatchThreadState
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result CreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Forms List
func (s *FormsService) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/forms"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.PostsUserID != nil {
			opts.Params["posts_user_id"] = *params.PostsUserID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ForumsService handles Forums operations.
type ForumsService struct {
	r Requester
}

// EditFeedOptions Edit Feed Options
func (s *ForumsService) EditFeedOptions(ctx context.Context, params *EditFeedOptionsParams) (*SaveChanges, error) {
	path := "/forums/feed/options"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Keywords != nil {
			opts.JSON["keywords"] = params.Keywords
		}
		if params.NodeIds != nil {
			opts.JSON["node_ids"] = params.NodeIds
		}
	}
	raw, err := s.r.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Follow Follow Forum
func (s *ForumsService) Follow(ctx context.Context, forumID int, params *FollowParams) (*SaveChanges, error) {
	path := "/forums/{forum_id}/followers"
	path = strings.Replace(path, "{forum_id}", fmt.Sprintf("%d", forumID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Email != nil {
			opts.JSON["email"] = *params.Email
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Followed Get Followed Forums
func (s *ForumsService) Followed(ctx context.Context, params *FollowedParams) (*FollowedResponse, error) {
	path := "/forums/followed"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Total != nil {
			opts.Params["total"] = *params.Total
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result FollowedResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Followers Get Followers
func (s *ForumsService) Followers(ctx context.Context, forumID int) (*FollowersResponse, error) {
	path := "/forums/{forum_id}/followers"
	path = strings.Replace(path, "{forum_id}", fmt.Sprintf("%d", forumID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result FollowersResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get Get Forum
func (s *ForumsService) Get(ctx context.Context, forumID int) (*GetResponse, error) {
	path := "/forums/{forum_id}"
	path = strings.Replace(path, "{forum_id}", fmt.Sprintf("%d", forumID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetFeedOptions Get Feed Options
func (s *ForumsService) GetFeedOptions(ctx context.Context) (*GetFeedOptionsResponse, error) {
	path := "/forums/feed/options"
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetFeedOptionsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Grouped Get Forums Tree
func (s *ForumsService) Grouped(ctx context.Context) (*GroupedResponse, error) {
	path := "/forums/grouped"
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GroupedResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Forums
func (s *ForumsService) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/forums"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.PostsUserID != nil {
			opts.Params["posts_user_id"] = *params.PostsUserID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Unfollow Unfollow Forum
func (s *ForumsService) Unfollow(ctx context.Context, forumID int) (*SaveChanges, error) {
	path := "/forums/{forum_id}/followers"
	path = strings.Replace(path, "{forum_id}", fmt.Sprintf("%d", forumID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// LinkForumsService handles LinkForums operations.
type LinkForumsService struct {
	r Requester
}

// Get Get Link Forum
func (s *LinkForumsService) Get(ctx context.Context, linkID int) (*GetResponse, error) {
	path := "/link-forums/{link_id}"
	path = strings.Replace(path, "{link_id}", fmt.Sprintf("%d", linkID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Links Forum
func (s *LinkForumsService) List(ctx context.Context) (*ListResponse, error) {
	path := "/link-forums"
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// NavigationService handles Navigation operations.
type NavigationService struct {
	r Requester
}

// List Get Navigation
func (s *NavigationService) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/navigation"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.PostsUserID != nil {
			opts.Params["posts_user_id"] = *params.PostsUserID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// NotificationsService handles Notifications operations.
type NotificationsService struct {
	r Requester
}

// Get Get Notification
func (s *NotificationsService) Get(ctx context.Context, notificationID int) (*GetResponse, error) {
	path := "/notifications/{notification_id}/content"
	path = strings.Replace(path, "{notification_id}", fmt.Sprintf("%d", notificationID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Notifications
func (s *NotificationsService) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/notifications"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.PostsUserID != nil {
			opts.Params["posts_user_id"] = *params.PostsUserID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Read Mark Notification Read
func (s *NotificationsService) Read(ctx context.Context, params *ReadParams) (*SaveChanges, error) {
	path := "/notifications/read"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.NotificationID != nil {
			opts.JSON["notification_id"] = *params.NotificationID
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PagesService handles Pages operations.
type PagesService struct {
	r Requester
}

// Get Get Page
func (s *PagesService) Get(ctx context.Context, pageID int) (*GetResponse, error) {
	path := "/pages/{page_id}"
	path = strings.Replace(path, "{page_id}", fmt.Sprintf("%d", pageID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Pages
func (s *PagesService) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/pages"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.PostsUserID != nil {
			opts.Params["posts_user_id"] = *params.PostsUserID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PostCommentsService handles PostComments operations.
type PostCommentsService struct {
	r Requester
}

// Create Create Post Comment
func (s *PostCommentsService) Create(ctx context.Context, commentBody string, postID int) (*CreateResponse, error) {
	path := "/posts/comments"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["comment_body"] = commentBody
	opts.JSON["post_id"] = postID
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result CreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete Delete Post Comment
func (s *PostCommentsService) Delete(ctx context.Context, postCommentID int, params *DeleteParams) (*SaveChanges, error) {
	path := "/posts/comments"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["post_comment_id"] = postCommentID
	if params != nil {
		if params.Reason != nil {
			opts.JSON["reason"] = *params.Reason
		}
	}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Edit Edit Post Comment
func (s *PostCommentsService) Edit(ctx context.Context, commentBody string, postCommentID int) (*EditResponse, error) {
	path := "/posts/comments"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["comment_body"] = commentBody
	opts.JSON["post_comment_id"] = postCommentID
	raw, err := s.r.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}
	var result EditResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get Get Post Comments
func (s *PostCommentsService) Get(ctx context.Context, postID int, params *GetParams) (*GetResponse, error) {
	path := "/posts/comments"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	opts.Params["post_id"] = postID
	if params != nil {
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PostsService handles Posts operations.
type PostsService struct {
	r Requester
}

// Create Create Post
func (s *PostsService) Create(ctx context.Context, postBody string, params *CreateParams) (*CreateResponse, error) {
	path := "/posts"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["post_body"] = postBody
	if params != nil {
		if params.AllowAskHiddenContent != nil {
			opts.JSON["allow_ask_hidden_content"] = *params.AllowAskHiddenContent
		}
		if params.CommentIgnoreGroup != nil {
			opts.JSON["comment_ignore_group"] = *params.CommentIgnoreGroup
		}
		if params.DontAlertFollowers != nil {
			opts.JSON["dont_alert_followers"] = *params.DontAlertFollowers
		}
		if params.HideContacts != nil {
			opts.JSON["hide_contacts"] = *params.HideContacts
		}
		if params.PrefixID != nil {
			opts.JSON["prefix_id"] = params.PrefixID
		}
		if params.ReplyGroup != nil {
			opts.JSON["reply_group"] = *params.ReplyGroup
		}
		if params.ScheduleDate != nil {
			opts.JSON["schedule_date"] = *params.ScheduleDate
		}
		if params.ScheduleTime != nil {
			opts.JSON["schedule_time"] = *params.ScheduleTime
		}
		if params.Tags != nil {
			opts.JSON["tags"] = params.Tags
		}
		if params.Title != nil {
			opts.JSON["title"] = *params.Title
		}
		if params.TitleEn != nil {
			opts.JSON["title_en"] = *params.TitleEn
		}
		if params.WatchThread != nil {
			opts.JSON["watch_thread"] = *params.WatchThread
		}
		if params.WatchThreadEmail != nil {
			opts.JSON["watch_thread_email"] = *params.WatchThreadEmail
		}
		if params.WatchThreadState != nil {
			opts.JSON["watch_thread_state"] = *params.WatchThreadState
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result CreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete Delete Post
func (s *PostsService) Delete(ctx context.Context, postID int, params *DeleteParams) (*SaveChanges, error) {
	path := "/posts/{post_id}"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%d", postID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Reason != nil {
			opts.JSON["reason"] = *params.Reason
		}
	}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Edit Edit Post
func (s *PostsService) Edit(ctx context.Context, postID int, params *EditParams) (*EditResponse, error) {
	path := "/posts/{post_id}"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%d", postID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.ActivityVisible != nil {
			opts.JSON["activity_visible"] = *params.ActivityVisible
		}
		if params.Alert != nil {
			for k, v := range params.Alert {
				opts.JSON[fmt.Sprintf("alert[%s]", k)] = v
			}
		}
		if params.AllowInviteGroup != nil {
			opts.JSON["allow_invite_group"] = *params.AllowInviteGroup
		}
		if params.AllowPostProfile != nil {
			opts.JSON["allow_post_profile"] = *params.AllowPostProfile
		}
		if params.AllowReceiveNewsFeed != nil {
			opts.JSON["allow_receive_news_feed"] = *params.AllowReceiveNewsFeed
		}
		if params.AllowSendPersonalConversation != nil {
			opts.JSON["allow_send_personal_conversation"] = *params.AllowSendPersonalConversation
		}
		if params.AllowViewProfile != nil {
			opts.JSON["allow_view_profile"] = *params.AllowViewProfile
		}
		if params.ConvWelcomeMessage != nil {
			opts.JSON["conv_welcome_message"] = *params.ConvWelcomeMessage
		}
		if params.DisplayBannerID != nil {
			opts.JSON["display_banner_id"] = *params.DisplayBannerID
		}
		if params.DisplayGroupID != nil {
			opts.JSON["display_group_id"] = *params.DisplayGroupID
		}
		if params.DisplayIconGroupID != nil {
			opts.JSON["display_icon_group_id"] = *params.DisplayIconGroupID
		}
		if params.Fields != nil {
			for k, v := range params.Fields {
				opts.JSON[fmt.Sprintf("fields[%s]", k)] = v
			}
		}
		if params.Gender != nil {
			opts.JSON["gender"] = *params.Gender
		}
		if params.HideUsernameChangeLogs != nil {
			opts.JSON["hide_username_change_logs"] = *params.HideUsernameChangeLogs
		}
		if params.LanguageID != nil {
			opts.JSON["language_id"] = *params.LanguageID
		}
		if params.ReceiveAdminEmail != nil {
			opts.JSON["receive_admin_email"] = *params.ReceiveAdminEmail
		}
		if params.SecretAnswer != nil {
			opts.JSON["secret_answer"] = *params.SecretAnswer
		}
		if params.SecretAnswerType != nil {
			opts.JSON["secret_answer_type"] = *params.SecretAnswerType
		}
		if params.ShortLink != nil {
			opts.JSON["short_link"] = *params.ShortLink
		}
		if params.ShowDobDate != nil {
			opts.JSON["show_dob_date"] = *params.ShowDobDate
		}
		if params.ShowDobYear != nil {
			opts.JSON["show_dob_year"] = *params.ShowDobYear
		}
		if params.Timezone != nil {
			opts.JSON["timezone"] = *params.Timezone
		}
		if params.UserDobDay != nil {
			opts.JSON["user_dob_day"] = *params.UserDobDay
		}
		if params.UserDobMonth != nil {
			opts.JSON["user_dob_month"] = *params.UserDobMonth
		}
		if params.UserDobYear != nil {
			opts.JSON["user_dob_year"] = *params.UserDobYear
		}
		if params.UserTitle != nil {
			opts.JSON["user_title"] = *params.UserTitle
		}
		if params.Username != nil {
			opts.JSON["username"] = *params.Username
		}
	}
	raw, err := s.r.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}
	var result EditResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get Get Post
func (s *PostsService) Get(ctx context.Context, postID int) (*GetResponse, error) {
	path := "/posts/{post_id}"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%d", postID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Like Like Post
func (s *PostsService) Like(ctx context.Context, postID int) (*SaveChanges, error) {
	path := "/posts/{post_id}/likes"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%d", postID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Likes Get Post Likes
func (s *PostsService) Likes(ctx context.Context, postID int, params *LikesParams) (*LikesResponse, error) {
	path := "/posts/{post_id}/likes"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%d", postID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.NodeID != nil {
			opts.Params["node_id"] = *params.NodeID
		}
		if params.LikeType != nil {
			opts.Params["like_type"] = *params.LikeType
		}
		if params.Type_ != nil {
			opts.Params["type"] = *params.Type_
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.ContentType != nil {
			opts.Params["content_type"] = *params.ContentType
		}
		if params.SearchUserID != nil {
			opts.Params["search_user_id"] = *params.SearchUserID
		}
		if params.Stats != nil {
			opts.Params["stats"] = *params.Stats
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result LikesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Posts
func (s *PostsService) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/posts"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.PostsUserID != nil {
			opts.Params["posts_user_id"] = *params.PostsUserID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Report Report Post Comment
func (s *PostsService) Report(ctx context.Context, message string, postCommentID int) (*SaveChanges, error) {
	path := "/posts/comments/report"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["message"] = message
	opts.JSON["post_comment_id"] = postCommentID
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReportPost Report Post
func (s *PostsService) ReportPost(ctx context.Context, postID int, message string) (*SaveChanges, error) {
	path := "/posts/{post_id}/report"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%d", postID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["message"] = message
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReportReasons Get Post Report Reasons
func (s *PostsService) ReportReasons(ctx context.Context, postID int) (*ReportReasonsResponse, error) {
	path := "/posts/{post_id}/report"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%d", postID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ReportReasonsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Unlike Unlike Post
func (s *PostsService) Unlike(ctx context.Context, postID int) (*SaveChanges, error) {
	path := "/posts/{post_id}/likes"
	path = strings.Replace(path, "{post_id}", fmt.Sprintf("%d", postID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ProfilePostCommentsService handles ProfilePostComments operations.
type ProfilePostCommentsService struct {
	r Requester
}

// Create Create Profile Post Comment
func (s *ProfilePostCommentsService) Create(ctx context.Context, commentBody string, profilePostID int) (*CreateResponse, error) {
	path := "/profile-posts/comments"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["comment_body"] = commentBody
	opts.JSON["profile_post_id"] = profilePostID
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result CreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete Delete Profile Post Comment
func (s *ProfilePostCommentsService) Delete(ctx context.Context, commentID int) (*SaveChanges, error) {
	path := "/profile-posts/comments"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["comment_id"] = commentID
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Edit Edit Profile Post Comment
func (s *ProfilePostCommentsService) Edit(ctx context.Context, commentBody string, commentID int) (*EditResponse, error) {
	path := "/profile-posts/comments"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["comment_body"] = commentBody
	opts.JSON["comment_id"] = commentID
	raw, err := s.r.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}
	var result EditResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get Get Profile Post Comment
func (s *ProfilePostCommentsService) Get(ctx context.Context, profilePostID int, commentID int) (*GetResponse, error) {
	path := "/profile-posts/{profile_post_id}/comments/{comment_id}"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%d", profilePostID), 1)
	path = strings.Replace(path, "{comment_id}", fmt.Sprintf("%d", commentID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Profile Post Comments
func (s *ProfilePostCommentsService) List(ctx context.Context, profilePostID int, params *ListParams) (*ListResponse, error) {
	path := "/profile-posts/comments"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	opts.Params["profile_post_id"] = profilePostID
	if params != nil {
		if params.PostsUserID != nil {
			opts.Params["posts_user_id"] = *params.PostsUserID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Report Report a Profile Post Comment
func (s *ProfilePostCommentsService) Report(ctx context.Context, commentID int, message string) (*SaveChanges, error) {
	path := "/profile-posts/comments/{comment_id}/report"
	path = strings.Replace(path, "{comment_id}", fmt.Sprintf("%d", commentID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["message"] = message
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ProfilePostsService handles ProfilePosts operations.
type ProfilePostsService struct {
	r Requester
}

// Create Create Profile Post
func (s *ProfilePostsService) Create(ctx context.Context, postBody string, userID StringOrInt) (*CreateResponse, error) {
	path := "/profile-posts"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["post_body"] = postBody
	opts.JSON["user_id"] = userID
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result CreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete Delete Profile Post
func (s *ProfilePostsService) Delete(ctx context.Context, profilePostID int, params *DeleteParams) (*SaveChanges, error) {
	path := "/profile-posts/{profile_post_id}"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%d", profilePostID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Reason != nil {
			opts.JSON["reason"] = *params.Reason
		}
	}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Edit Edit Profile Post
func (s *ProfilePostsService) Edit(ctx context.Context, profilePostID int, params *EditParams) (*EditResponse, error) {
	path := "/profile-posts/{profile_post_id}"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%d", profilePostID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.ActivityVisible != nil {
			opts.JSON["activity_visible"] = *params.ActivityVisible
		}
		if params.Alert != nil {
			for k, v := range params.Alert {
				opts.JSON[fmt.Sprintf("alert[%s]", k)] = v
			}
		}
		if params.AllowInviteGroup != nil {
			opts.JSON["allow_invite_group"] = *params.AllowInviteGroup
		}
		if params.AllowPostProfile != nil {
			opts.JSON["allow_post_profile"] = *params.AllowPostProfile
		}
		if params.AllowReceiveNewsFeed != nil {
			opts.JSON["allow_receive_news_feed"] = *params.AllowReceiveNewsFeed
		}
		if params.AllowSendPersonalConversation != nil {
			opts.JSON["allow_send_personal_conversation"] = *params.AllowSendPersonalConversation
		}
		if params.AllowViewProfile != nil {
			opts.JSON["allow_view_profile"] = *params.AllowViewProfile
		}
		if params.ConvWelcomeMessage != nil {
			opts.JSON["conv_welcome_message"] = *params.ConvWelcomeMessage
		}
		if params.DisplayBannerID != nil {
			opts.JSON["display_banner_id"] = *params.DisplayBannerID
		}
		if params.DisplayGroupID != nil {
			opts.JSON["display_group_id"] = *params.DisplayGroupID
		}
		if params.DisplayIconGroupID != nil {
			opts.JSON["display_icon_group_id"] = *params.DisplayIconGroupID
		}
		if params.Fields != nil {
			for k, v := range params.Fields {
				opts.JSON[fmt.Sprintf("fields[%s]", k)] = v
			}
		}
		if params.Gender != nil {
			opts.JSON["gender"] = *params.Gender
		}
		if params.HideUsernameChangeLogs != nil {
			opts.JSON["hide_username_change_logs"] = *params.HideUsernameChangeLogs
		}
		if params.LanguageID != nil {
			opts.JSON["language_id"] = *params.LanguageID
		}
		if params.ReceiveAdminEmail != nil {
			opts.JSON["receive_admin_email"] = *params.ReceiveAdminEmail
		}
		if params.SecretAnswer != nil {
			opts.JSON["secret_answer"] = *params.SecretAnswer
		}
		if params.SecretAnswerType != nil {
			opts.JSON["secret_answer_type"] = *params.SecretAnswerType
		}
		if params.ShortLink != nil {
			opts.JSON["short_link"] = *params.ShortLink
		}
		if params.ShowDobDate != nil {
			opts.JSON["show_dob_date"] = *params.ShowDobDate
		}
		if params.ShowDobYear != nil {
			opts.JSON["show_dob_year"] = *params.ShowDobYear
		}
		if params.Timezone != nil {
			opts.JSON["timezone"] = *params.Timezone
		}
		if params.UserDobDay != nil {
			opts.JSON["user_dob_day"] = *params.UserDobDay
		}
		if params.UserDobMonth != nil {
			opts.JSON["user_dob_month"] = *params.UserDobMonth
		}
		if params.UserDobYear != nil {
			opts.JSON["user_dob_year"] = *params.UserDobYear
		}
		if params.UserTitle != nil {
			opts.JSON["user_title"] = *params.UserTitle
		}
		if params.Username != nil {
			opts.JSON["username"] = *params.Username
		}
	}
	raw, err := s.r.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}
	var result EditResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get Get Profile Post
func (s *ProfilePostsService) Get(ctx context.Context, profilePostID int) (*GetResponse, error) {
	path := "/profile-posts/{profile_post_id}"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%d", profilePostID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Like Like Profile Post
func (s *ProfilePostsService) Like(ctx context.Context, profilePostID int) (*SaveChanges, error) {
	path := "/profile-posts/{profile_post_id}/likes"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%d", profilePostID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Likes Get Profile Post Likes
func (s *ProfilePostsService) Likes(ctx context.Context, profilePostID int) (*LikesResponse, error) {
	path := "/profile-posts/{profile_post_id}/likes"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%d", profilePostID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result LikesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Profile Posts
func (s *ProfilePostsService) List(ctx context.Context, userID StringOrInt, params *ListParams) (*ListResponse, error) {
	path := "/users/{user_id}/profile-posts"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.PostsUserID != nil {
			opts.Params["posts_user_id"] = *params.PostsUserID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Report Report a Profile Post
func (s *ProfilePostsService) Report(ctx context.Context, profilePostID int, message string) (*SaveChanges, error) {
	path := "/profile-posts/{profile_post_id}/report"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%d", profilePostID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["message"] = message
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReportReasons Get Profile Post Report Reasons
func (s *ProfilePostsService) ReportReasons(ctx context.Context, profilePostID int) (*ReportReasonsResponse, error) {
	path := "/profile-posts/{profile_post_id}/report"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%d", profilePostID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ReportReasonsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Stick Stick Profile Post
func (s *ProfilePostsService) Stick(ctx context.Context, profilePostID int) (*SaveChanges, error) {
	path := "/profile-posts/{profile_post_id}/stick"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%d", profilePostID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Unlike Unlike Profile Post
func (s *ProfilePostsService) Unlike(ctx context.Context, profilePostID int) (*SaveChanges, error) {
	path := "/profile-posts/{profile_post_id}/likes"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%d", profilePostID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Unstick Unstick Profile Post
func (s *ProfilePostsService) Unstick(ctx context.Context, profilePostID int) (*SaveChanges, error) {
	path := "/profile-posts/{profile_post_id}/stick"
	path = strings.Replace(path, "{profile_post_id}", fmt.Sprintf("%d", profilePostID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SearchingService handles Searching operations.
type SearchingService struct {
	r Requester
}

// All Search
func (s *SearchingService) All(ctx context.Context, params *AllParams) (*AllResponse, error) {
	path := "/search"
	opts := RequestOptions{}
	opts.Search = true
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Before != nil {
			opts.JSON["before"] = *params.Before
		}
		if params.ForumID != nil {
			opts.JSON["forum_id"] = *params.ForumID
		}
		if params.Limit != nil {
			opts.JSON["limit"] = *params.Limit
		}
		if params.Page != nil {
			opts.JSON["page"] = *params.Page
		}
		if params.Q != nil {
			opts.JSON["q"] = *params.Q
		}
		if params.Tag != nil {
			opts.JSON["tag"] = *params.Tag
		}
		opts.JSON["user_id"] = params.UserID
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result AllResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Posts Search Post
func (s *SearchingService) Posts(ctx context.Context, params *PostsParams) (*PostsResponse, error) {
	path := "/search/posts"
	opts := RequestOptions{}
	opts.Search = true
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Before != nil {
			opts.JSON["before"] = *params.Before
		}
		if params.DataLimit != nil {
			opts.JSON["data_limit"] = *params.DataLimit
		}
		if params.ForumID != nil {
			opts.JSON["forum_id"] = *params.ForumID
		}
		if params.Limit != nil {
			opts.JSON["limit"] = *params.Limit
		}
		if params.Page != nil {
			opts.JSON["page"] = *params.Page
		}
		if params.Q != nil {
			opts.JSON["q"] = *params.Q
		}
		if params.Tag != nil {
			opts.JSON["tag"] = *params.Tag
		}
		opts.JSON["user_id"] = params.UserID
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result PostsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ProfilePosts Search Profile Posts
func (s *SearchingService) ProfilePosts(ctx context.Context, params *ProfilePostsParams) (*ProfilePostsResponse, error) {
	path := "/search/profile-posts"
	opts := RequestOptions{}
	opts.Search = true
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Before != nil {
			opts.JSON["before"] = *params.Before
		}
		if params.Limit != nil {
			opts.JSON["limit"] = *params.Limit
		}
		if params.Page != nil {
			opts.JSON["page"] = *params.Page
		}
		if params.Q != nil {
			opts.JSON["q"] = *params.Q
		}
		if params.UserID != nil {
			opts.JSON["user_id"] = *params.UserID
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result ProfilePostsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Results Get Search Results
func (s *SearchingService) Results(ctx context.Context, searchID StringOrInt, params *ResultsParams) (*ResultsResponse, error) {
	path := "/search/{search_id}/results"
	path = strings.Replace(path, "{search_id}", fmt.Sprintf("%v", searchID), 1)
	opts := RequestOptions{}
	opts.Search = true
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ResultsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Tagged Search Tagged
func (s *SearchingService) Tagged(ctx context.Context, params *TaggedParams) (*TaggedResponse, error) {
	path := "/search/tagged"
	opts := RequestOptions{}
	opts.Search = true
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Limit != nil {
			opts.JSON["limit"] = *params.Limit
		}
		if params.Page != nil {
			opts.JSON["page"] = *params.Page
		}
		if params.Tag != nil {
			opts.JSON["tag"] = *params.Tag
		}
		if params.Tags != nil {
			opts.JSON["tags"] = params.Tags
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result TaggedResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Threads Search Thread
func (s *SearchingService) Threads(ctx context.Context, params *ThreadsParams) (*ThreadsResponse, error) {
	path := "/search/threads"
	opts := RequestOptions{}
	opts.Search = true
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Before != nil {
			opts.JSON["before"] = *params.Before
		}
		if params.DataLimit != nil {
			opts.JSON["data_limit"] = *params.DataLimit
		}
		if params.ForumID != nil {
			opts.JSON["forum_id"] = *params.ForumID
		}
		if params.Limit != nil {
			opts.JSON["limit"] = *params.Limit
		}
		if params.Page != nil {
			opts.JSON["page"] = *params.Page
		}
		if params.Q != nil {
			opts.JSON["q"] = *params.Q
		}
		if params.Tag != nil {
			opts.JSON["tag"] = *params.Tag
		}
		opts.JSON["user_id"] = params.UserID
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result ThreadsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Users Search Users
func (s *SearchingService) Users(ctx context.Context, params *UsersParams) (*UsersResponse, error) {
	path := "/search/users"
	opts := RequestOptions{}
	opts.Search = true
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Q != nil {
			opts.JSON["q"] = *params.Q
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result UsersResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ThreadsService handles Threads operations.
type ThreadsService struct {
	r Requester
}

// Bump Bump Thread
func (s *ThreadsService) Bump(ctx context.Context, threadID int) (*BumpResponse, error) {
	path := "/threads/{thread_id}/bump"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result BumpResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Claim Create Claim
func (s *ThreadsService) Claim(ctx context.Context, asAmount float64, asIsMarketDeal bool, asResponder string, postBody string, transferType ThreadsClaimTransferType, params *ClaimParams) (*ClaimResponse, error) {
	path := "/claims"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["as_amount"] = asAmount
	opts.JSON["as_is_market_deal"] = asIsMarketDeal
	opts.JSON["as_responder"] = asResponder
	opts.JSON["post_body"] = postBody
	opts.JSON["transfer_type"] = transferType
	if params != nil {
		if params.AllowAskHiddenContent != nil {
			opts.JSON["allow_ask_hidden_content"] = *params.AllowAskHiddenContent
		}
		if params.AsData != nil {
			opts.JSON["as_data"] = *params.AsData
		}
		if params.AsFundsReceipt != nil {
			opts.JSON["as_funds_receipt"] = *params.AsFundsReceipt
		}
		if params.AsMarketItemID != nil {
			opts.JSON["as_market_item_id"] = *params.AsMarketItemID
		}
		if params.AsTgLoginScreenshot != nil {
			opts.JSON["as_tg_login_screenshot"] = *params.AsTgLoginScreenshot
		}
		if params.CommentIgnoreGroup != nil {
			opts.JSON["comment_ignore_group"] = *params.CommentIgnoreGroup
		}
		if params.Currency != nil {
			opts.JSON["currency"] = *params.Currency
		}
		if params.DontAlertFollowers != nil {
			opts.JSON["dont_alert_followers"] = *params.DontAlertFollowers
		}
		if params.HideContacts != nil {
			opts.JSON["hide_contacts"] = *params.HideContacts
		}
		if params.PayClaim != nil {
			opts.JSON["pay_claim"] = *params.PayClaim
		}
		if params.ReplyGroup != nil {
			opts.JSON["reply_group"] = *params.ReplyGroup
		}
		if params.ScheduleDate != nil {
			opts.JSON["schedule_date"] = *params.ScheduleDate
		}
		if params.ScheduleTime != nil {
			opts.JSON["schedule_time"] = *params.ScheduleTime
		}
		if params.Tags != nil {
			opts.JSON["tags"] = params.Tags
		}
		if params.WatchThread != nil {
			opts.JSON["watch_thread"] = *params.WatchThread
		}
		if params.WatchThreadEmail != nil {
			opts.JSON["watch_thread_email"] = *params.WatchThreadEmail
		}
		if params.WatchThreadState != nil {
			opts.JSON["watch_thread_state"] = *params.WatchThreadState
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result ClaimResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create Create Thread
func (s *ThreadsService) Create(ctx context.Context, forumID int, postBody string, params *CreateParams) (*CreateResponse, error) {
	path := "/threads"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["forum_id"] = forumID
	opts.JSON["post_body"] = postBody
	if params != nil {
		if params.AllowAskHiddenContent != nil {
			opts.JSON["allow_ask_hidden_content"] = *params.AllowAskHiddenContent
		}
		if params.CommentIgnoreGroup != nil {
			opts.JSON["comment_ignore_group"] = *params.CommentIgnoreGroup
		}
		if params.DontAlertFollowers != nil {
			opts.JSON["dont_alert_followers"] = *params.DontAlertFollowers
		}
		if params.HideContacts != nil {
			opts.JSON["hide_contacts"] = *params.HideContacts
		}
		if params.PrefixID != nil {
			opts.JSON["prefix_id"] = params.PrefixID
		}
		if params.ReplyGroup != nil {
			opts.JSON["reply_group"] = *params.ReplyGroup
		}
		if params.ScheduleDate != nil {
			opts.JSON["schedule_date"] = *params.ScheduleDate
		}
		if params.ScheduleTime != nil {
			opts.JSON["schedule_time"] = *params.ScheduleTime
		}
		if params.Tags != nil {
			opts.JSON["tags"] = params.Tags
		}
		if params.Title != nil {
			opts.JSON["title"] = *params.Title
		}
		if params.TitleEn != nil {
			opts.JSON["title_en"] = *params.TitleEn
		}
		if params.WatchThread != nil {
			opts.JSON["watch_thread"] = *params.WatchThread
		}
		if params.WatchThreadEmail != nil {
			opts.JSON["watch_thread_email"] = *params.WatchThreadEmail
		}
		if params.WatchThreadState != nil {
			opts.JSON["watch_thread_state"] = *params.WatchThreadState
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result CreateResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateContest Create Contest
func (s *ThreadsService) CreateContest(ctx context.Context, contestType ThreadsCreateContestContestType, postBody string, prizeType ThreadsCreateContestPrizeType, requireLikeCount int, requireTotalLikeCount int, params *CreateContestParams) (*CreateContestResponse, error) {
	path := "/contests"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["contest_type"] = contestType
	opts.JSON["post_body"] = postBody
	opts.JSON["prize_type"] = prizeType
	opts.JSON["require_like_count"] = requireLikeCount
	opts.JSON["require_total_like_count"] = requireTotalLikeCount
	if params != nil {
		if params.AllowAskHiddenContent != nil {
			opts.JSON["allow_ask_hidden_content"] = *params.AllowAskHiddenContent
		}
		if params.CommentIgnoreGroup != nil {
			opts.JSON["comment_ignore_group"] = *params.CommentIgnoreGroup
		}
		if params.CountWinners != nil {
			opts.JSON["count_winners"] = *params.CountWinners
		}
		if params.DontAlertFollowers != nil {
			opts.JSON["dont_alert_followers"] = *params.DontAlertFollowers
		}
		if params.HideContacts != nil {
			opts.JSON["hide_contacts"] = *params.HideContacts
		}
		if params.IsMoneyPlaces != nil {
			opts.JSON["is_money_places"] = *params.IsMoneyPlaces
		}
		if params.LengthOption != nil {
			opts.JSON["length_option"] = *params.LengthOption
		}
		if params.LengthValue != nil {
			opts.JSON["length_value"] = *params.LengthValue
		}
		if params.PrizeDataMoney != nil {
			opts.JSON["prize_data_money"] = *params.PrizeDataMoney
		}
		if params.PrizeDataPlaces != nil {
			opts.JSON["prize_data_places"] = params.PrizeDataPlaces
		}
		if params.PrizeDataUpgrade != nil {
			opts.JSON["prize_data_upgrade"] = *params.PrizeDataUpgrade
		}
		if params.ReplyGroup != nil {
			opts.JSON["reply_group"] = *params.ReplyGroup
		}
		if params.ScheduleDate != nil {
			opts.JSON["schedule_date"] = *params.ScheduleDate
		}
		if params.ScheduleTime != nil {
			opts.JSON["schedule_time"] = *params.ScheduleTime
		}
		if params.SecretAnswer != nil {
			opts.JSON["secret_answer"] = *params.SecretAnswer
		}
		if params.Tags != nil {
			opts.JSON["tags"] = params.Tags
		}
		if params.Title != nil {
			opts.JSON["title"] = *params.Title
		}
		if params.TitleEn != nil {
			opts.JSON["title_en"] = *params.TitleEn
		}
		if params.WatchThread != nil {
			opts.JSON["watch_thread"] = *params.WatchThread
		}
		if params.WatchThreadEmail != nil {
			opts.JSON["watch_thread_email"] = *params.WatchThreadEmail
		}
		if params.WatchThreadState != nil {
			opts.JSON["watch_thread_state"] = *params.WatchThreadState
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result CreateContestResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete Delete Thread
func (s *ThreadsService) Delete(ctx context.Context, threadID int, params *DeleteParams) (*SaveChanges, error) {
	path := "/threads/{thread_id}"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Reason != nil {
			opts.JSON["reason"] = *params.Reason
		}
	}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Edit Edit thread
func (s *ThreadsService) Edit(ctx context.Context, threadID int, params *EditParams) (*EditResponse, error) {
	path := "/threads/{thread_id}"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.ActivityVisible != nil {
			opts.JSON["activity_visible"] = *params.ActivityVisible
		}
		if params.Alert != nil {
			for k, v := range params.Alert {
				opts.JSON[fmt.Sprintf("alert[%s]", k)] = v
			}
		}
		if params.AllowInviteGroup != nil {
			opts.JSON["allow_invite_group"] = *params.AllowInviteGroup
		}
		if params.AllowPostProfile != nil {
			opts.JSON["allow_post_profile"] = *params.AllowPostProfile
		}
		if params.AllowReceiveNewsFeed != nil {
			opts.JSON["allow_receive_news_feed"] = *params.AllowReceiveNewsFeed
		}
		if params.AllowSendPersonalConversation != nil {
			opts.JSON["allow_send_personal_conversation"] = *params.AllowSendPersonalConversation
		}
		if params.AllowViewProfile != nil {
			opts.JSON["allow_view_profile"] = *params.AllowViewProfile
		}
		if params.ConvWelcomeMessage != nil {
			opts.JSON["conv_welcome_message"] = *params.ConvWelcomeMessage
		}
		if params.DisplayBannerID != nil {
			opts.JSON["display_banner_id"] = *params.DisplayBannerID
		}
		if params.DisplayGroupID != nil {
			opts.JSON["display_group_id"] = *params.DisplayGroupID
		}
		if params.DisplayIconGroupID != nil {
			opts.JSON["display_icon_group_id"] = *params.DisplayIconGroupID
		}
		if params.Fields != nil {
			for k, v := range params.Fields {
				opts.JSON[fmt.Sprintf("fields[%s]", k)] = v
			}
		}
		if params.Gender != nil {
			opts.JSON["gender"] = *params.Gender
		}
		if params.HideUsernameChangeLogs != nil {
			opts.JSON["hide_username_change_logs"] = *params.HideUsernameChangeLogs
		}
		if params.LanguageID != nil {
			opts.JSON["language_id"] = *params.LanguageID
		}
		if params.ReceiveAdminEmail != nil {
			opts.JSON["receive_admin_email"] = *params.ReceiveAdminEmail
		}
		if params.SecretAnswer != nil {
			opts.JSON["secret_answer"] = *params.SecretAnswer
		}
		if params.SecretAnswerType != nil {
			opts.JSON["secret_answer_type"] = *params.SecretAnswerType
		}
		if params.ShortLink != nil {
			opts.JSON["short_link"] = *params.ShortLink
		}
		if params.ShowDobDate != nil {
			opts.JSON["show_dob_date"] = *params.ShowDobDate
		}
		if params.ShowDobYear != nil {
			opts.JSON["show_dob_year"] = *params.ShowDobYear
		}
		if params.Timezone != nil {
			opts.JSON["timezone"] = *params.Timezone
		}
		if params.UserDobDay != nil {
			opts.JSON["user_dob_day"] = *params.UserDobDay
		}
		if params.UserDobMonth != nil {
			opts.JSON["user_dob_month"] = *params.UserDobMonth
		}
		if params.UserDobYear != nil {
			opts.JSON["user_dob_year"] = *params.UserDobYear
		}
		if params.UserTitle != nil {
			opts.JSON["user_title"] = *params.UserTitle
		}
		if params.Username != nil {
			opts.JSON["username"] = *params.Username
		}
	}
	raw, err := s.r.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}
	var result EditResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Finish Finish Contest
func (s *ThreadsService) Finish(ctx context.Context, threadID int) (*SaveChanges, error) {
	path := "/contests/{thread_id}/finish"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Follow Follow Thread
func (s *ThreadsService) Follow(ctx context.Context, threadID int, params *FollowParams) (*SaveChanges, error) {
	path := "/threads/{thread_id}/followers"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Email != nil {
			opts.JSON["email"] = *params.Email
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Followed Get Followed Threads
func (s *ThreadsService) Followed(ctx context.Context, params *FollowedParams) (*FollowedResponse, error) {
	path := "/threads/followed"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Total != nil {
			opts.Params["total"] = *params.Total
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result FollowedResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Followers Get Thread Followers
func (s *ThreadsService) Followers(ctx context.Context, threadID int) (*FollowersResponse, error) {
	path := "/threads/{thread_id}/followers"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result FollowersResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get Get Thread
func (s *ThreadsService) Get(ctx context.Context, threadID int, params *GetParams) (*GetResponse, error) {
	path := "/threads/{thread_id}"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetGet Get Poll
func (s *ThreadsService) GetGet(ctx context.Context, threadID int) (*GetGetResponse, error) {
	path := "/threads/{thread_id}/poll"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetGetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Hide Hide Thread
func (s *ThreadsService) Hide(ctx context.Context, threadID int) (*HideResponse, error) {
	path := "/threads/{thread_id}/hide"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result HideResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Threads
func (s *ThreadsService) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/threads"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.PostsUserID != nil {
			opts.Params["posts_user_id"] = *params.PostsUserID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Move Move Thread
func (s *ThreadsService) Move(ctx context.Context, threadID int, nodeID string, params *MoveParams) (*SaveChanges, error) {
	path := "/threads/{thread_id}/move"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["node_id"] = nodeID
	if params != nil {
		if params.ApplyThreadPrefix != nil {
			opts.JSON["apply_thread_prefix"] = *params.ApplyThreadPrefix
		}
		if params.PrefixID != nil {
			opts.JSON["prefix_id"] = params.PrefixID
		}
		if params.SendAlert != nil {
			opts.JSON["send_alert"] = *params.SendAlert
		}
		if params.Title != nil {
			opts.JSON["title"] = *params.Title
		}
		if params.TitleEn != nil {
			opts.JSON["title_en"] = *params.TitleEn
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Navigation Get Navigation Elements
func (s *ThreadsService) Navigation(ctx context.Context, threadID int) (*NavigationResponse, error) {
	path := "/threads/{thread_id}/navigation"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result NavigationResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Recent Get Recent Threads
func (s *ThreadsService) Recent(ctx context.Context, params *RecentParams) (*RecentResponse, error) {
	path := "/threads/recent"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Days != nil {
			opts.Params["days"] = *params.Days
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.ForumID != nil {
			opts.Params["forum_id"] = *params.ForumID
		}
		if params.DataLimit != nil {
			opts.Params["data_limit"] = *params.DataLimit
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result RecentResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Star Bookmark Thread
func (s *ThreadsService) Star(ctx context.Context, threadID int) (*SaveChanges, error) {
	path := "/threads/{thread_id}/star"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Unfollow Unfollow Thread
func (s *ThreadsService) Unfollow(ctx context.Context, threadID int) (*SaveChanges, error) {
	path := "/threads/{thread_id}/followers"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Unread Get Unread Threads
func (s *ThreadsService) Unread(ctx context.Context, params *UnreadParams) (*UnreadResponse, error) {
	path := "/threads/new"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.ForumID != nil {
			opts.Params["forum_id"] = *params.ForumID
		}
		if params.DataLimit != nil {
			opts.Params["data_limit"] = *params.DataLimit
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result UnreadResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Unstar Unbookmark Thread
func (s *ThreadsService) Unstar(ctx context.Context, threadID int) (*SaveChanges, error) {
	path := "/threads/{thread_id}/star"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Vote Vote Poll
func (s *ThreadsService) Vote(ctx context.Context, threadID int, params *VoteParams) (*SaveChanges, error) {
	path := "/threads/{thread_id}/poll/votes"
	path = strings.Replace(path, "{thread_id}", fmt.Sprintf("%d", threadID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.ResponseID != nil {
			opts.JSON["response_id"] = *params.ResponseID
		}
		if params.ResponseIds != nil {
			opts.JSON["response_ids"] = params.ResponseIds
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UsersService handles Users operations.
type UsersService struct {
	r Requester
}

// CancelReset Cancel Secret Answer Reset
func (s *UsersService) CancelReset(ctx context.Context) (*SaveChanges, error) {
	path := "/account/secret-answer/reset"
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Claims Get User Claims
func (s *UsersService) Claims(ctx context.Context, userID StringOrInt, params *ClaimsParams) (*ClaimsResponse, error) {
	path := "/users/{user_id}/claims"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Type_ != nil {
			opts.Params["type"] = *params.Type_
		}
		if params.ClaimState != nil {
			opts.Params["claim_state"] = *params.ClaimState
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ClaimsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Contents Get Contents
func (s *UsersService) Contents(ctx context.Context, userID StringOrInt, params *ContentsParams) (*ContentsResponse, error) {
	path := "/users/{user_id}/timeline"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ContentsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Crop Crop Avatar
func (s *UsersService) Crop(ctx context.Context, userID StringOrInt, params *CropParams) (*CropResponse, error) {
	path := "/users/{user_id}/avatar/crop"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Crop != nil {
			opts.JSON["crop"] = *params.Crop
		}
		if params.X != nil {
			opts.JSON["x"] = *params.X
		}
		if params.Y != nil {
			opts.JSON["y"] = *params.Y
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result CropResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CropPost Crop Background
func (s *UsersService) CropPost(ctx context.Context, userID StringOrInt, params *CropPostParams) (*CropPostResponse, error) {
	path := "/users/{user_id}/background/crop"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Crop != nil {
			opts.JSON["crop"] = *params.Crop
		}
		if params.X != nil {
			opts.JSON["x"] = *params.X
		}
		if params.Y != nil {
			opts.JSON["y"] = *params.Y
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result CropPostResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete Delete Avatar
func (s *UsersService) Delete(ctx context.Context, userID StringOrInt) (*SaveChanges, error) {
	path := "/users/{user_id}/avatar"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteDelete Delete Background
func (s *UsersService) DeleteDelete(ctx context.Context, userID StringOrInt) (*SaveChanges, error) {
	path := "/users/{user_id}/background"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Edit Edit User
func (s *UsersService) Edit(ctx context.Context, userID StringOrInt, params *EditParams) (*SaveChanges, error) {
	path := "/users/{user_id}"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.ActivityVisible != nil {
			opts.JSON["activity_visible"] = *params.ActivityVisible
		}
		if params.Alert != nil {
			for k, v := range params.Alert {
				opts.JSON[fmt.Sprintf("alert[%s]", k)] = v
			}
		}
		if params.AllowInviteGroup != nil {
			opts.JSON["allow_invite_group"] = *params.AllowInviteGroup
		}
		if params.AllowPostProfile != nil {
			opts.JSON["allow_post_profile"] = *params.AllowPostProfile
		}
		if params.AllowReceiveNewsFeed != nil {
			opts.JSON["allow_receive_news_feed"] = *params.AllowReceiveNewsFeed
		}
		if params.AllowSendPersonalConversation != nil {
			opts.JSON["allow_send_personal_conversation"] = *params.AllowSendPersonalConversation
		}
		if params.AllowViewProfile != nil {
			opts.JSON["allow_view_profile"] = *params.AllowViewProfile
		}
		if params.ConvWelcomeMessage != nil {
			opts.JSON["conv_welcome_message"] = *params.ConvWelcomeMessage
		}
		if params.DisplayBannerID != nil {
			opts.JSON["display_banner_id"] = *params.DisplayBannerID
		}
		if params.DisplayGroupID != nil {
			opts.JSON["display_group_id"] = *params.DisplayGroupID
		}
		if params.DisplayIconGroupID != nil {
			opts.JSON["display_icon_group_id"] = *params.DisplayIconGroupID
		}
		if params.Fields != nil {
			for k, v := range params.Fields {
				opts.JSON[fmt.Sprintf("fields[%s]", k)] = v
			}
		}
		if params.Gender != nil {
			opts.JSON["gender"] = *params.Gender
		}
		if params.HideUsernameChangeLogs != nil {
			opts.JSON["hide_username_change_logs"] = *params.HideUsernameChangeLogs
		}
		if params.LanguageID != nil {
			opts.JSON["language_id"] = *params.LanguageID
		}
		if params.ReceiveAdminEmail != nil {
			opts.JSON["receive_admin_email"] = *params.ReceiveAdminEmail
		}
		if params.SecretAnswer != nil {
			opts.JSON["secret_answer"] = *params.SecretAnswer
		}
		if params.SecretAnswerType != nil {
			opts.JSON["secret_answer_type"] = *params.SecretAnswerType
		}
		if params.ShortLink != nil {
			opts.JSON["short_link"] = *params.ShortLink
		}
		if params.ShowDobDate != nil {
			opts.JSON["show_dob_date"] = *params.ShowDobDate
		}
		if params.ShowDobYear != nil {
			opts.JSON["show_dob_year"] = *params.ShowDobYear
		}
		if params.Timezone != nil {
			opts.JSON["timezone"] = *params.Timezone
		}
		if params.UserDobDay != nil {
			opts.JSON["user_dob_day"] = *params.UserDobDay
		}
		if params.UserDobMonth != nil {
			opts.JSON["user_dob_month"] = *params.UserDobMonth
		}
		if params.UserDobYear != nil {
			opts.JSON["user_dob_year"] = *params.UserDobYear
		}
		if params.UserTitle != nil {
			opts.JSON["user_title"] = *params.UserTitle
		}
		if params.Username != nil {
			opts.JSON["username"] = *params.Username
		}
	}
	raw, err := s.r.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fields Get User Fields
func (s *UsersService) Fields(ctx context.Context) (*FieldsResponse, error) {
	path := "/users/fields"
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result FieldsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Find Find Users
func (s *UsersService) Find(ctx context.Context, params *FindParams) (*FindResponse, error) {
	path := "/users/find"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Username != nil {
			opts.Params["username"] = *params.Username
		}
		if params.CustomFields != nil {
			for k, v := range params.CustomFields {
				opts.Params[fmt.Sprintf("custom_fields[%s]", k)] = v
			}
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result FindResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Follow Follow User
func (s *UsersService) Follow(ctx context.Context, userID StringOrInt) (*SaveChanges, error) {
	path := "/users/{user_id}/followers"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Followers Get User Followers
func (s *UsersService) Followers(ctx context.Context, userID StringOrInt, params *FollowersParams) (*FollowersResponse, error) {
	path := "/users/{user_id}/followers"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Order != nil {
			opts.Params["order"] = *params.Order
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result FollowersResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Followings Get Followed Users By User
func (s *UsersService) Followings(ctx context.Context, userID StringOrInt, params *FollowingsParams) (*FollowingsResponse, error) {
	path := "/users/{user_id}/followings"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Order != nil {
			opts.Params["order"] = *params.Order
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result FollowingsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get Get User
func (s *UsersService) Get(ctx context.Context, userID StringOrInt, params *GetParams) (*GetResponse, error) {
	path := "/users/{user_id}"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Ignore Ignore User
func (s *UsersService) Ignore(ctx context.Context, userID StringOrInt) (*SaveChanges, error) {
	path := "/users/{user_id}/ignore"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// IgnoreEdit Edit Ignoring Options
func (s *UsersService) IgnoreEdit(ctx context.Context, userID StringOrInt, params *IgnoreEditParams) (*SaveChanges, error) {
	path := "/users/{user_id}/ignore"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.IgnoreConversations != nil {
			opts.JSON["ignore_conversations"] = *params.IgnoreConversations
		}
		if params.IgnoreContent != nil {
			opts.JSON["ignore_content"] = *params.IgnoreContent
		}
		if params.RestrictViewProfile != nil {
			opts.JSON["restrict_view_profile"] = *params.RestrictViewProfile
		}
	}
	raw, err := s.r.Request(ctx, "PUT", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Ignored Get Ignored Users
func (s *UsersService) Ignored(ctx context.Context, params *IgnoredParams) (*IgnoredResponse, error) {
	path := "/users/ignored"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Total != nil {
			opts.Params["total"] = *params.Total
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result IgnoredResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Likes Get User Likes
func (s *UsersService) Likes(ctx context.Context, userID StringOrInt, params *LikesParams) (*LikesResponse, error) {
	path := "/users/{user_id}/likes"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.NodeID != nil {
			opts.Params["node_id"] = *params.NodeID
		}
		if params.LikeType != nil {
			opts.Params["like_type"] = *params.LikeType
		}
		if params.Type_ != nil {
			opts.Params["type"] = *params.Type_
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.ContentType != nil {
			opts.Params["content_type"] = *params.ContentType
		}
		if params.SearchUserID != nil {
			opts.Params["search_user_id"] = *params.SearchUserID
		}
		if params.Stats != nil {
			opts.Params["stats"] = *params.Stats
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result LikesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Users
func (s *UsersService) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/users"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.PostsUserID != nil {
			opts.Params["posts_user_id"] = *params.PostsUserID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.FieldsInclude != nil {
			opts.Params["fields_include"] = params.FieldsInclude
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ListResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Reset Reset Secret Answer
func (s *UsersService) Reset(ctx context.Context) (*ResetResponse, error) {
	path := "/account/secret-answer/reset"
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result ResetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SecretAnswerTypes Get Secret Answer Types
func (s *UsersService) SecretAnswerTypes(ctx context.Context) (*SecretAnswerTypesResponse, error) {
	path := "/users/secret-answer/types"
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result SecretAnswerTypesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Trophies Get Trophies
func (s *UsersService) Trophies(ctx context.Context, userID StringOrInt) (*TrophiesResponse, error) {
	path := "/users/{user_id}/trophies"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result TrophiesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Unfollow Unfollow User
func (s *UsersService) Unfollow(ctx context.Context, userID StringOrInt) (*SaveChanges, error) {
	path := "/users/{user_id}/followers"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Unignore Unignore User
func (s *UsersService) Unignore(ctx context.Context, userID StringOrInt) (*SaveChanges, error) {
	path := "/users/{user_id}/ignore"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result SaveChanges
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Upload Upload Avatar
func (s *UsersService) Upload(ctx context.Context, userID StringOrInt, avatar []byte, params *UploadParams) (*UploadResponse, error) {
	path := "/users/{user_id}/avatar"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["avatar"] = avatar
	if params != nil {
		if params.Crop != nil {
			opts.JSON["crop"] = *params.Crop
		}
		if params.X != nil {
			opts.JSON["x"] = *params.X
		}
		if params.Y != nil {
			opts.JSON["y"] = *params.Y
		}
	}
	delete(opts.JSON, "avatar")
	opts.Files = make(map[string]FileUpload)
	opts.Files["avatar"] = FileUpload{Filename: "avatar", Data: avatar}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result UploadResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UploadPost Upload Background
func (s *UsersService) UploadPost(ctx context.Context, userID StringOrInt, background []byte, params *UploadPostParams) (*UploadPostResponse, error) {
	path := "/users/{user_id}/background"
	path = strings.Replace(path, "{user_id}", fmt.Sprintf("%v", userID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["background"] = background
	if params != nil {
		if params.Crop != nil {
			opts.JSON["crop"] = *params.Crop
		}
		if params.X != nil {
			opts.JSON["x"] = *params.X
		}
		if params.Y != nil {
			opts.JSON["y"] = *params.Y
		}
	}
	delete(opts.JSON, "background")
	opts.Files = make(map[string]FileUpload)
	opts.Files["background"] = FileUpload{Filename: "background", Data: background}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result UploadPostResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Ensure imports are used.
var _ = fmt.Sprintf
var _ = strings.Replace
var _ = json.Marshal

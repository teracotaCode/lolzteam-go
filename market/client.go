package market

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

// MarketClient is the API client.
type MarketClient struct {
	r                 Requester
	AccountPublishing *AccountPublishingService
	AccountPurchasing *AccountPurchasingService
	AccountsList      *AccountsListService
	AccountsManaging  *AccountsManagingService
	BatchRequests     *BatchRequestsService
	Cart              *CartService
	Categories        *CategoriesService
	CategorySearch    *CategorySearchService
	CustomDiscounts   *CustomDiscountsService
	IMAP              *IMAPService
	Invoices          *InvoicesService
	Payments          *PaymentsService
	Profile           *ProfileService
	Proxy             *ProxyService
}

// New creates a new MarketClient.
func New(r Requester) *MarketClient {
	c := &MarketClient{r: r}
	c.AccountPublishing = &AccountPublishingService{r: r}
	c.AccountPurchasing = &AccountPurchasingService{r: r}
	c.AccountsList = &AccountsListService{r: r}
	c.AccountsManaging = &AccountsManagingService{r: r}
	c.BatchRequests = &BatchRequestsService{r: r}
	c.Cart = &CartService{r: r}
	c.Categories = &CategoriesService{r: r}
	c.CategorySearch = &CategorySearchService{r: r}
	c.CustomDiscounts = &CustomDiscountsService{r: r}
	c.IMAP = &IMAPService{r: r}
	c.Invoices = &InvoicesService{r: r}
	c.Payments = &PaymentsService{r: r}
	c.Profile = &ProfileService{r: r}
	c.Proxy = &ProxyService{r: r}
	return c
}

// AccountPublishingService handles AccountPublishing operations.
type AccountPublishingService struct {
	r Requester
}

// Add Add Account
func (s *AccountPublishingService) Add(ctx context.Context, categoryID AccountPublishingAddCategoryID, currency AccountPublishingAddCurrency, itemOrigin AccountPublishingAddItemOrigin, price float64, params *AddParams) (*AddResponse, error) {
	path := "/item/add"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["category_id"] = categoryID
	opts.JSON["currency"] = currency
	opts.JSON["item_origin"] = itemOrigin
	opts.JSON["price"] = price
	if params != nil {
		if params.ProxyIP != nil {
			opts.JSON["proxy_ip"] = *params.ProxyIP
		}
		if params.ProxyPass != nil {
			opts.JSON["proxy_pass"] = *params.ProxyPass
		}
		if params.ProxyPort != nil {
			opts.JSON["proxy_port"] = *params.ProxyPort
		}
		if params.ProxyRow != nil {
			opts.JSON["proxy_row"] = *params.ProxyRow
		}
		if params.ProxyUser != nil {
			opts.JSON["proxy_user"] = *params.ProxyUser
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result AddResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Check Check Account Details
func (s *AccountPublishingService) Check(ctx context.Context, itemID int, params *CheckParams) (*BuyItem, error) {
	path := "/{item_id}/goods/check"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.EmailLoginData != nil {
			opts.JSON["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailType != nil {
			opts.JSON["email_type"] = *params.EmailType
		}
		if params.Extra != nil {
			for k, v := range params.Extra {
				opts.JSON[fmt.Sprintf("extra[%s]", k)] = v
			}
		}
		if params.HasEmailLoginData != nil {
			opts.JSON["has_email_login_data"] = *params.HasEmailLoginData
		}
		if params.Login != nil {
			opts.JSON["login"] = *params.Login
		}
		if params.LoginPassword != nil {
			opts.JSON["login_password"] = *params.LoginPassword
		}
		if params.Password != nil {
			opts.JSON["password"] = *params.Password
		}
		if params.RandomProxy != nil {
			opts.JSON["random_proxy"] = *params.RandomProxy
		}
		if params.ResellItemID != nil {
			opts.JSON["resell_item_id"] = *params.ResellItemID
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result BuyItem
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// External Add an External Account
func (s *AccountPublishingService) External(ctx context.Context, itemID int, type_ AccountPublishingExternalType, params *ExternalParams) (*SaveChanges, error) {
	path := "/{item_id}/external-account"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["type"] = type_
	if params != nil {
		if params.Cookies != nil {
			opts.JSON["cookies"] = *params.Cookies
		}
		if params.EmailLoginData != nil {
			opts.JSON["email_login_data"] = *params.EmailLoginData
		}
		if params.Login != nil {
			opts.JSON["login"] = *params.Login
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

// FastSell Fast Account Upload
func (s *AccountPublishingService) FastSell(ctx context.Context, categoryID AccountPublishingFastSellCategoryID, currency AccountPublishingFastSellCurrency, itemOrigin AccountPublishingFastSellItemOrigin, price float64, params *FastSellParams) (*FastSellResponse, error) {
	path := "/item/fast-sell"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["category_id"] = categoryID
	opts.JSON["currency"] = currency
	opts.JSON["item_origin"] = itemOrigin
	opts.JSON["price"] = price
	if params != nil {
		if params.AllowAskDiscount != nil {
			opts.JSON["allow_ask_discount"] = *params.AllowAskDiscount
		}
		if params.Description != nil {
			opts.JSON["description"] = *params.Description
		}
		if params.EmailLoginData != nil {
			opts.JSON["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailType != nil {
			opts.JSON["email_type"] = *params.EmailType
		}
		if params.ExtendedGuarantee != nil {
			opts.JSON["extended_guarantee"] = *params.ExtendedGuarantee
		}
		if params.Extra != nil {
			for k, v := range params.Extra {
				opts.JSON[fmt.Sprintf("extra[%s]", k)] = v
			}
		}
		if params.HasEmailLoginData != nil {
			opts.JSON["has_email_login_data"] = *params.HasEmailLoginData
		}
		if params.Information != nil {
			opts.JSON["information"] = *params.Information
		}
		if params.Login != nil {
			opts.JSON["login"] = *params.Login
		}
		if params.LoginPassword != nil {
			opts.JSON["login_password"] = *params.LoginPassword
		}
		if params.Password != nil {
			opts.JSON["password"] = *params.Password
		}
		if params.ProxyID != nil {
			opts.JSON["proxy_id"] = *params.ProxyID
		}
		opts.JSON["random_proxy"] = params.RandomProxy
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
	var result FastSellResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AccountPurchasingService handles AccountPurchasing operations.
type AccountPurchasingService struct {
	r Requester
}

// Check Check Account
func (s *AccountPurchasingService) Check(ctx context.Context, itemID int) (*CheckAccount, error) {
	path := "/{item_id}/check-account"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result CheckAccount
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Confirm Confirm Buy
func (s *AccountPurchasingService) Confirm(ctx context.Context, itemID int, params *ConfirmParams) (*ConfirmResponse, error) {
	path := "/{item_id}/confirm-buy"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.BalanceID != nil {
			opts.JSON["balance_id"] = *params.BalanceID
		}
		if params.Price != nil {
			opts.JSON["price"] = *params.Price
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result ConfirmResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DiscountCancel Cancel Discount Request
func (s *AccountPurchasingService) DiscountCancel(ctx context.Context, itemID int) (*SaveChanges, error) {
	path := "/{item_id}/discount"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
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

// DiscountRequest Discount Request
func (s *AccountPurchasingService) DiscountRequest(ctx context.Context, itemID int, discountPrice float64, params *DiscountRequestParams) (*SaveChanges, error) {
	path := "/{item_id}/discount"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["discount_price"] = discountPrice
	if params != nil {
		if params.Message != nil {
			opts.JSON["message"] = *params.Message
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

// FastBuy Fast Buy Account
func (s *AccountPurchasingService) FastBuy(ctx context.Context, itemID int, params *FastBuyParams) (*BuyItem, error) {
	path := "/{item_id}/fast-buy"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.BalanceID != nil {
			opts.JSON["balance_id"] = *params.BalanceID
		}
		if params.Price != nil {
			opts.JSON["price"] = *params.Price
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result BuyItem
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AccountsListService handles AccountsList operations.
type AccountsListService struct {
	r Requester
}

// Download Download Accounts Data
func (s *AccountsListService) Download(ctx context.Context, type_ AccountsListType, params *DownloadParams) (string, error) {
	path := "/user/{type}/download"
	path = strings.Replace(path, "{type}", fmt.Sprintf("%v", type_), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Format != nil {
			opts.Params["format"] = *params.Format
		}
		if params.CustomFormat != nil {
			opts.Params["custom_format"] = *params.CustomFormat
		}
		if params.CategoryID != nil {
			opts.Params["category_id"] = *params.CategoryID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Show != nil {
			opts.Params["show"] = *params.Show
		}
		if params.DeleteReason != nil {
			opts.Params["delete_reason"] = *params.DeleteReason
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.Username != nil {
			opts.Params["username"] = *params.Username
		}
		if params.PublishedStartDate != nil {
			opts.Params["published_startDate"] = *params.PublishedStartDate
		}
		if params.PublishedEndDate != nil {
			opts.Params["published_endDate"] = *params.PublishedEndDate
		}
		if params.FilterByPublishedDate != nil {
			opts.Params["filter_by_published_date"] = *params.FilterByPublishedDate
		}
		if params.PaidStartDate != nil {
			opts.Params["paid_startDate"] = *params.PaidStartDate
		}
		if params.PaidEndDate != nil {
			opts.Params["paid_endDate"] = *params.PaidEndDate
		}
		if params.FilterByBuyerOperationDate != nil {
			opts.Params["filter_by_buyer_operation_date"] = *params.FilterByBuyerOperationDate
		}
		if params.DeleteStartDate != nil {
			opts.Params["delete_startDate"] = *params.DeleteStartDate
		}
		if params.DeleteEndDate != nil {
			opts.Params["delete_endDate"] = *params.DeleteEndDate
		}
		if params.FilterByDeleteDate != nil {
			opts.Params["filter_by_delete_date"] = *params.FilterByDeleteDate
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

// Favorites Get All Favourites Accounts
func (s *AccountsListService) Favorites(ctx context.Context, params *FavoritesParams) (*ItemListModel, error) {
	path := "/fave"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Show != nil {
			opts.Params["show"] = *params.Show
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ItemListModel
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Orders Get All Purchased Accounts
func (s *AccountsListService) Orders(ctx context.Context, params *OrdersParams) (*ItemListModel, error) {
	path := "/user/orders"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.CategoryID != nil {
			opts.Params["category_id"] = *params.CategoryID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Show != nil {
			opts.Params["show"] = *params.Show
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Login != nil {
			opts.Params["login"] = *params.Login
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ItemListModel
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// States Get User Items States
func (s *AccountsListService) States(ctx context.Context, params *StatesParams) (*StatesResponse, error) {
	path := "/user/item-states"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		opts.Params["user_id"] = params.UserID
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result StatesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// User Get All User Accounts
func (s *AccountsListService) User(ctx context.Context, params *UserParams) (*ItemListModel, error) {
	path := "/user/items"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.CategoryID != nil {
			opts.Params["category_id"] = *params.CategoryID
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Show != nil {
			opts.Params["show"] = *params.Show
		}
		if params.DeleteReason != nil {
			opts.Params["delete_reason"] = *params.DeleteReason
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Login != nil {
			opts.Params["login"] = *params.Login
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.Username != nil {
			opts.Params["username"] = *params.Username
		}
		if params.PublishedStartDate != nil {
			opts.Params["published_startDate"] = *params.PublishedStartDate
		}
		if params.PublishedEndDate != nil {
			opts.Params["published_endDate"] = *params.PublishedEndDate
		}
		if params.FilterByPublishedDate != nil {
			opts.Params["filter_by_published_date"] = *params.FilterByPublishedDate
		}
		if params.PaidStartDate != nil {
			opts.Params["paid_startDate"] = *params.PaidStartDate
		}
		if params.PaidEndDate != nil {
			opts.Params["paid_endDate"] = *params.PaidEndDate
		}
		if params.FilterByBuyerOperationDate != nil {
			opts.Params["filter_by_buyer_operation_date"] = *params.FilterByBuyerOperationDate
		}
		if params.DeleteStartDate != nil {
			opts.Params["delete_startDate"] = *params.DeleteStartDate
		}
		if params.DeleteEndDate != nil {
			opts.Params["delete_endDate"] = *params.DeleteEndDate
		}
		if params.FilterByDeleteDate != nil {
			opts.Params["filter_by_delete_date"] = *params.FilterByDeleteDate
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ItemListModel
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Viewed Get All Viewed Accounts
func (s *AccountsListService) Viewed(ctx context.Context, params *ViewedParams) (*ItemListModel, error) {
	path := "/viewed"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Show != nil {
			opts.Params["show"] = *params.Show
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ItemListModel
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AccountsManagingService handles AccountsManaging operations.
type AccountsManagingService struct {
	r Requester
}

// AIPrice Get AI Price
func (s *AccountsManagingService) AIPrice(ctx context.Context, itemID int) (*AIPriceResponse, error) {
	path := "/{item_id}/ai-price"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result AIPriceResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AddMafile Add Mafile
func (s *AccountsManagingService) AddMafile(ctx context.Context, itemID int) (*StatusItem, error) {
	path := "/{item_id}/mafile"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result StatusItem
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AutoBump Auto Bump
func (s *AccountsManagingService) AutoBump(ctx context.Context, itemID int, hour int) (*Status, error) {
	path := "/{item_id}/auto-bump"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["hour"] = hour
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result Status
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AutoBumpDisable Disable Auto Bump
func (s *AccountsManagingService) AutoBumpDisable(ctx context.Context, itemID int) (*Status, error) {
	path := "/{item_id}/auto-bump"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result Status
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AutoBuyPrice Get Auto Buy Price
func (s *AccountsManagingService) AutoBuyPrice(ctx context.Context, itemID int) (*AutoBuyPriceResponse, error) {
	path := "/{item_id}/auto-buy-price"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result AutoBuyPriceResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BulkGet Bulk Get Accounts
func (s *AccountsManagingService) BulkGet(ctx context.Context, params *BulkGetParams) (*BulkGetResponse, error) {
	path := "/bulk/items"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.ItemID != nil {
			opts.JSON["item_id"] = params.ItemID
		}
		if params.ParseSameItemIds != nil {
			opts.JSON["parse_same_item_ids"] = *params.ParseSameItemIds
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result BulkGetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Bump Bump Account
func (s *AccountsManagingService) Bump(ctx context.Context, itemID int) (*Status, error) {
	path := "/{item_id}/bump"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result Status
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ChangePassword Change Password
func (s *AccountsManagingService) ChangePassword(ctx context.Context, itemID int, params *ChangePasswordParams) (*ChangePasswordResponse, error) {
	path := "/{item_id}/change-password"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Cancel != nil {
			opts.JSON["_cancel"] = *params.Cancel
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result ChangePasswordResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CheckGuarantee Check Guarantee
func (s *AccountsManagingService) CheckGuarantee(ctx context.Context, itemID int) (*CheckGuaranteeResponse, error) {
	path := "/{item_id}/check-guarantee"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result CheckGuaranteeResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Claims Get Claims
func (s *AccountsManagingService) Claims(ctx context.Context, params *ClaimsParams) (*ClaimsResponse, error) {
	path := "/claims"
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

// Close Close Account
func (s *AccountsManagingService) Close(ctx context.Context, itemID int) (*Status, error) {
	path := "/{item_id}/close"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result Status
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateClaim Create Claim
func (s *AccountsManagingService) CreateClaim(ctx context.Context, itemID int, postBody string) (*CreateClaimResponse, error) {
	path := "/claims"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["item_id"] = itemID
	opts.JSON["post_body"] = postBody
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result CreateClaimResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeclineVideoRecording Decline Video Recording Request
func (s *AccountsManagingService) DeclineVideoRecording(ctx context.Context, itemID int, iVoluntarilyAndWithFullAwarenessOfMyActionsWaiveAnyClaimsRegardingThisItem bool) (*SaveChanges, error) {
	path := "/{item_id}/decline-video-recording"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["i_voluntarily_and_with_full_awareness_of_my_actions_waive_any_claims_regarding_this_item"] = iVoluntarilyAndWithFullAwarenessOfMyActionsWaiveAnyClaimsRegardingThisItem
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

// Delete Delete Account
func (s *AccountsManagingService) Delete(ctx context.Context, itemID int, reason string) (*Status, error) {
	path := "/{item_id}"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["reason"] = reason
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result Status
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Edit Edit Account
func (s *AccountsManagingService) Edit(ctx context.Context, itemID int, params *EditParams) (*SaveChanges, error) {
	path := "/{item_id}/edit"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.AllowAskDiscount != nil {
			opts.JSON["allow_ask_discount"] = *params.AllowAskDiscount
		}
		if params.Currency != nil {
			opts.JSON["currency"] = *params.Currency
		}
		if params.Description != nil {
			opts.JSON["description"] = *params.Description
		}
		if params.EmailLoginData != nil {
			opts.JSON["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailType != nil {
			opts.JSON["email_type"] = *params.EmailType
		}
		if params.Information != nil {
			opts.JSON["information"] = *params.Information
		}
		if params.ItemOrigin != nil {
			opts.JSON["item_origin"] = *params.ItemOrigin
		}
		if params.Price != nil {
			opts.JSON["price"] = *params.Price
		}
		if params.ProxyID != nil {
			opts.JSON["proxy_id"] = *params.ProxyID
		}
		if params.Title != nil {
			opts.JSON["title"] = *params.Title
		}
		if params.TitleEn != nil {
			opts.JSON["title_en"] = *params.TitleEn
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

// EmailCode Get Email Confirmation Code
func (s *AccountsManagingService) EmailCode(ctx context.Context, itemID int) (*ConfirmationCodeModel, error) {
	path := "/{item_id}/email-code"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ConfirmationCodeModel
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Favorite Favorite
func (s *AccountsManagingService) Favorite(ctx context.Context, itemID int) (*Status, error) {
	path := "/{item_id}/star"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result Status
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get Get Account
func (s *AccountsManagingService) Get(ctx context.Context, itemID int, params *GetParams) (*GetResponse, error) {
	path := "/{item_id}"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
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

// GetLetters2 Get Email Letters
func (s *AccountsManagingService) GetLetters2(ctx context.Context, params *GetLetters2Params) (*GetLetters2Response, error) {
	path := "/letters2"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.EmailPassword != nil {
			opts.Params["email_password"] = *params.EmailPassword
		}
		if params.Email != nil {
			opts.Params["email"] = *params.Email
		}
		if params.Password != nil {
			opts.Params["password"] = *params.Password
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetLetters2Response
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetMafile Get Mafile
func (s *AccountsManagingService) GetMafile(ctx context.Context, itemID int) (*GetMafileResponse, error) {
	path := "/{item_id}/mafile"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GetMafileResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Image Get Account Image
func (s *AccountsManagingService) Image(ctx context.Context, itemID int, type_ AccountsManagingType) (*ImageResponse, error) {
	path := "/{item_id}/image"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	opts.Params["type"] = type_
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ImageResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Note Edit Note
func (s *AccountsManagingService) Note(ctx context.Context, itemID int, params *NoteParams) (*Status, error) {
	path := "/{item_id}/note-save"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.Text != nil {
			opts.JSON["text"] = *params.Text
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result Status
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Open Open Account
func (s *AccountsManagingService) Open(ctx context.Context, itemID int) (*Status, error) {
	path := "/{item_id}/open"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result Status
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PublicTag Add a Public Tag
func (s *AccountsManagingService) PublicTag(ctx context.Context, itemID int, tagID int) (*Tag, error) {
	path := "/{item_id}/public-tag"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["tag_id"] = tagID
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result Tag
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PublicUntag Remove a Public Tag
func (s *AccountsManagingService) PublicUntag(ctx context.Context, itemID int, tagID int) (*Tag, error) {
	path := "/{item_id}/public-tag"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["tag_id"] = tagID
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result Tag
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RefuseGuarantee Cancel Guarantee
func (s *AccountsManagingService) RefuseGuarantee(ctx context.Context, itemID int) (*SaveChanges, error) {
	path := "/{item_id}/refuse-guarantee"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
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

// RemoveMafile Remove Mafile
func (s *AccountsManagingService) RemoveMafile(ctx context.Context, itemID int) (*StatusItem, error) {
	path := "/{item_id}/mafile"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result StatusItem
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SteamInventoryValue Get Account Steam Inventory Value
func (s *AccountsManagingService) SteamInventoryValue(ctx context.Context, itemID int, params *SteamInventoryValueParams) (*SteamInventoryValueResponse, error) {
	path := "/{item_id}/inventory-value"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.AppID != nil {
			opts.Params["app_id"] = *params.AppID
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.IgnoreCache != nil {
			opts.Params["ignore_cache"] = *params.IgnoreCache
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result SteamInventoryValueResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SteamMafileCode Get Mafile Confirmation Code
func (s *AccountsManagingService) SteamMafileCode(ctx context.Context, itemID int) (*ConfirmationCodeModel, error) {
	path := "/{item_id}/guard-code"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ConfirmationCodeModel
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SteamPreview Get Steam HTML
func (s *AccountsManagingService) SteamPreview(ctx context.Context, itemID int, params *SteamPreviewParams) (string, error) {
	path := "/{item_id}/steam-preview"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Type_ != nil {
			opts.Params["type"] = *params.Type_
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return "", err
	}
	return string(raw), nil
}

// SteamSDA Confirm SDA
func (s *AccountsManagingService) SteamSDA(ctx context.Context, itemID int, params *SteamSDAParams) (*Status, error) {
	path := "/{item_id}/confirm-sda"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.ID != nil {
			opts.JSON["id"] = *params.ID
		}
		if params.Nonce != nil {
			opts.JSON["nonce"] = *params.Nonce
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result Status
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SteamUpdateValue Update Inventory Value
func (s *AccountsManagingService) SteamUpdateValue(ctx context.Context, itemID int, params *SteamUpdateValueParams) (*SteamUpdateValueResponse, error) {
	path := "/{item_id}/update-inventory"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.All != nil {
			opts.JSON["all"] = *params.All
		}
		if params.AppID != nil {
			opts.JSON["app_id"] = *params.AppID
		}
		if params.Authorize != nil {
			opts.JSON["authorize"] = *params.Authorize
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result SteamUpdateValueResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SteamValue Get Steam Inventory Value
func (s *AccountsManagingService) SteamValue(ctx context.Context, link string, params *SteamValueParams) (*SteamValueResponse, error) {
	path := "/steam-value"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	opts.Params["link"] = link
	if params != nil {
		if params.AppID != nil {
			opts.Params["app_id"] = *params.AppID
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.IgnoreCache != nil {
			opts.Params["ignore_cache"] = *params.IgnoreCache
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result SteamValueResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Stick Stick Account
func (s *AccountsManagingService) Stick(ctx context.Context, itemID int) (*SaveChanges, error) {
	path := "/{item_id}/stick"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
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

// Tag Add a Tag
func (s *AccountsManagingService) Tag(ctx context.Context, itemID int, tagID int) (*Tag, error) {
	path := "/{item_id}/tag"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["tag_id"] = tagID
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result Tag
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// TelegramCode Get Telegram Confirmation Code
func (s *AccountsManagingService) TelegramCode(ctx context.Context, itemID int) (*TelegramCodeResponse, error) {
	path := "/{item_id}/telegram-login-code"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result TelegramCodeResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// TelegramResetAuth Telegram Reset Auth
func (s *AccountsManagingService) TelegramResetAuth(ctx context.Context, itemID int) (*SaveChanges, error) {
	path := "/{item_id}/telegram-reset-authorizations"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
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

// TempEmailPassword Get Temp Email Password
func (s *AccountsManagingService) TempEmailPassword(ctx context.Context, itemID int) (*TempEmailPasswordResponse, error) {
	path := "/{item_id}/temp-email-password"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result TempEmailPasswordResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Transfer Change Account Owner
func (s *AccountsManagingService) Transfer(ctx context.Context, itemID int, secretAnswer string, username string) (*SaveChanges, error) {
	path := "/{item_id}/change-owner"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["secret_answer"] = secretAnswer
	opts.JSON["username"] = username
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

// Unfavorite Unfavorite
func (s *AccountsManagingService) Unfavorite(ctx context.Context, itemID int) (*SaveChanges, error) {
	path := "/{item_id}/star"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
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

// Unstick Unstick Account
func (s *AccountsManagingService) Unstick(ctx context.Context, itemID int) (*SaveChanges, error) {
	path := "/{item_id}/stick"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
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

// Untag Remove a Tag
func (s *AccountsManagingService) Untag(ctx context.Context, itemID int, tagID int) (*Tag, error) {
	path := "/{item_id}/tag"
	path = strings.Replace(path, "{item_id}", fmt.Sprintf("%d", itemID), 1)
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["tag_id"] = tagID
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result Tag
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BatchRequestsService handles BatchRequests operations.
type BatchRequestsService struct {
	r Requester
}

// Batch Batch
func (s *BatchRequestsService) Batch(ctx context.Context, jobs []map[string]interface{}) (*BatchResponse, error) {
	path := "/batch"
	opts := RequestOptions{}
	opts.JSONBody = jobs
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result BatchResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CartService handles Cart operations.
type CartService struct {
	r Requester
}

// Add Add Item to Cart
func (s *CartService) Add(ctx context.Context, itemID int) (*AddResponse, error) {
	path := "/cart"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["item_id"] = itemID
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result AddResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete Delete Item From Cart
func (s *CartService) Delete(ctx context.Context, params *DeleteParams) (*DeleteResponse, error) {
	path := "/cart"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.DeleteAll != nil {
			opts.JSON["delete_all"] = *params.DeleteAll
		}
		if params.ProxyID != nil {
			opts.JSON["proxy_id"] = *params.ProxyID
		}
	}
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result DeleteResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Get Get Cart Items
func (s *CartService) Get(ctx context.Context, params *GetParams) (*ItemListModel, error) {
	path := "/cart"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ItemListModel
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CategoriesService handles Categories operations.
type CategoriesService struct {
	r Requester
}

// Games Get Category Games
func (s *CategoriesService) Games(ctx context.Context, categoryName CategoriesCategoryName) (*GamesResponse, error) {
	path := "/{categoryName}/games"
	path = strings.Replace(path, "{categoryName}", fmt.Sprintf("%v", categoryName), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GamesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Categories
func (s *CategoriesService) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/category"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.Status != nil {
			opts.Params["status"] = *params.Status
		}
		if params.Amount != nil {
			opts.Params["amount"] = *params.Amount
		}
		if params.MerchantID != nil {
			opts.Params["merchant_id"] = *params.MerchantID
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

// Params Get Category Search Params
func (s *CategoriesService) Params(ctx context.Context, categoryName CategoriesCategoryName) (*ParamsResponse, error) {
	path := "/{categoryName}/params"
	path = strings.Replace(path, "{categoryName}", fmt.Sprintf("%v", categoryName), 1)
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ParamsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CategorySearchService handles CategorySearch operations.
type CategorySearchService struct {
	r Requester
}

// All Get Last Accounts
func (s *CategorySearchService) All(ctx context.Context, params *AllParams) (*ItemListModel, error) {
	path := "/"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ItemListModel
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BattleNet BattleNet
func (s *CategorySearchService) BattleNet(ctx context.Context, params *BattleNetParams) (*BattleNetResponse, error) {
	path := "/battlenet"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Eg != nil {
			opts.Params["eg"] = *params.Eg
		}
		if params.Game != nil {
			opts.Params["game[]"] = params.Game
		}
		if params.Daybreak != nil {
			opts.Params["daybreak"] = *params.Daybreak
		}
		if params.Country != nil {
			opts.Params["country[]"] = params.Country
		}
		if params.NotCountry != nil {
			opts.Params["not_country[]"] = params.NotCountry
		}
		if params.Tel != nil {
			opts.Params["tel"] = *params.Tel
		}
		if params.EditBtag != nil {
			opts.Params["edit_btag"] = *params.EditBtag
		}
		if params.ChangeableFn != nil {
			opts.Params["changeable_fn"] = *params.ChangeableFn
		}
		if params.RealID != nil {
			opts.Params["real_id"] = *params.RealID
		}
		if params.ParentControl != nil {
			opts.Params["parent_control"] = *params.ParentControl
		}
		if params.NoBans != nil {
			opts.Params["no_bans"] = *params.NoBans
		}
		if params.BalanceMin != nil {
			opts.Params["balance_min"] = *params.BalanceMin
		}
		if params.BalanceMax != nil {
			opts.Params["balance_max"] = *params.BalanceMax
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result BattleNetResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ChatGPT ChatGPT
func (s *CategorySearchService) ChatGPT(ctx context.Context, params *ChatGPTParams) (*ChatGPTResponse, error) {
	path := "/chatgpt"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Subscription != nil {
			opts.Params["subscription[]"] = params.Subscription
		}
		if params.SubscriptionLength != nil {
			opts.Params["subscription_length"] = *params.SubscriptionLength
		}
		if params.SubscriptionPeriod != nil {
			opts.Params["subscription_period"] = *params.SubscriptionPeriod
		}
		if params.Autorenewal != nil {
			opts.Params["autorenewal"] = *params.Autorenewal
		}
		if params.Tel != nil {
			opts.Params["tel"] = *params.Tel
		}
		if params.Transactions != nil {
			opts.Params["transactions"] = *params.Transactions
		}
		if params.Reg != nil {
			opts.Params["reg"] = *params.Reg
		}
		if params.RegPeriod != nil {
			opts.Params["reg_period"] = *params.RegPeriod
		}
		if params.OpenaiTier != nil {
			opts.Params["openai_tier[]"] = params.OpenaiTier
		}
		if params.OpenaiBalanceMin != nil {
			opts.Params["openai_balance_min"] = *params.OpenaiBalanceMin
		}
		if params.OpenaiBalanceMax != nil {
			opts.Params["openai_balance_max"] = *params.OpenaiBalanceMax
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result ChatGPTResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Discord Discord
func (s *CategorySearchService) Discord(ctx context.Context, params *DiscordParams) (*DiscordResponse, error) {
	path := "/discord"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Tel != nil {
			opts.Params["tel"] = *params.Tel
		}
		if params.Nitro != nil {
			opts.Params["nitro"] = *params.Nitro
		}
		if params.NitroType != nil {
			opts.Params["nitro_type[]"] = params.NitroType
		}
		if params.NitroLength != nil {
			opts.Params["nitro_length"] = *params.NitroLength
		}
		if params.NitroPeriod != nil {
			opts.Params["nitro_period"] = *params.NitroPeriod
		}
		if params.BoostsMin != nil {
			opts.Params["boosts_min"] = *params.BoostsMin
		}
		if params.BoostsMax != nil {
			opts.Params["boosts_max"] = *params.BoostsMax
		}
		if params.Billing != nil {
			opts.Params["billing"] = *params.Billing
		}
		if params.Gifts != nil {
			opts.Params["gifts"] = *params.Gifts
		}
		if params.Transactions != nil {
			opts.Params["transactions"] = *params.Transactions
		}
		if params.Badge != nil {
			opts.Params["badge[]"] = params.Badge
		}
		if params.Condition != nil {
			opts.Params["condition[]"] = params.Condition
		}
		if params.ChatMin != nil {
			opts.Params["chat_min"] = *params.ChatMin
		}
		if params.ChatMax != nil {
			opts.Params["chat_max"] = *params.ChatMax
		}
		if params.MinAdminMembers != nil {
			opts.Params["min_admin_members"] = *params.MinAdminMembers
		}
		if params.MaxAdminMembers != nil {
			opts.Params["max_admin_members"] = *params.MaxAdminMembers
		}
		if params.MinAdmin != nil {
			opts.Params["min_admin"] = *params.MinAdmin
		}
		if params.MaxAdmin != nil {
			opts.Params["max_admin"] = *params.MaxAdmin
		}
		if params.Reg != nil {
			opts.Params["reg"] = *params.Reg
		}
		if params.RegPeriod != nil {
			opts.Params["reg_period"] = *params.RegPeriod
		}
		if params.Language != nil {
			opts.Params["language[]"] = params.Language
		}
		if params.NotLanguage != nil {
			opts.Params["not_language[]"] = params.NotLanguage
		}
		if params.Clans != nil {
			opts.Params["clans"] = *params.Clans
		}
		if params.MinAdminClans != nil {
			opts.Params["min_admin_clans"] = *params.MinAdminClans
		}
		if params.MaxAdminClans != nil {
			opts.Params["max_admin_clans"] = *params.MaxAdminClans
		}
		if params.MinOwnerClans != nil {
			opts.Params["min_owner_clans"] = *params.MinOwnerClans
		}
		if params.MaxOwnerClans != nil {
			opts.Params["max_owner_clans"] = *params.MaxOwnerClans
		}
		if params.Country != nil {
			opts.Params["country[]"] = params.Country
		}
		if params.NotCountry != nil {
			opts.Params["not_country[]"] = params.NotCountry
		}
		if params.MinServers != nil {
			opts.Params["min_servers"] = *params.MinServers
		}
		if params.MaxServers != nil {
			opts.Params["max_servers"] = *params.MaxServers
		}
		if params.Field2fa != nil {
			opts.Params["2fa"] = *params.Field2fa
		}
		if params.MinFullCredits != nil {
			opts.Params["min_full_credits"] = *params.MinFullCredits
		}
		if params.MaxFullCredits != nil {
			opts.Params["max_full_credits"] = *params.MaxFullCredits
		}
		if params.MinBasicCredits != nil {
			opts.Params["min_basic_credits"] = *params.MinBasicCredits
		}
		if params.MaxBasicCredits != nil {
			opts.Params["max_basic_credits"] = *params.MaxBasicCredits
		}
		if params.MinOrbs != nil {
			opts.Params["min_orbs"] = *params.MinOrbs
		}
		if params.MaxOrbs != nil {
			opts.Params["max_orbs"] = *params.MaxOrbs
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result DiscordResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// EA EA (Origin)
func (s *CategorySearchService) EA(ctx context.Context, params *EAParams) (*EAResponse, error) {
	path := "/ea"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Game != nil {
			opts.Params["game[]"] = params.Game
		}
		if params.Country != nil {
			opts.Params["country[]"] = params.Country
		}
		if params.NotCountry != nil {
			opts.Params["not_country[]"] = params.NotCountry
		}
		if params.Gmin != nil {
			opts.Params["gmin"] = *params.Gmin
		}
		if params.Gmax != nil {
			opts.Params["gmax"] = *params.Gmax
		}
		if params.AlRankMin != nil {
			opts.Params["al_rank_min"] = *params.AlRankMin
		}
		if params.AlRankMax != nil {
			opts.Params["al_rank_max"] = *params.AlRankMax
		}
		if params.AlLevelMin != nil {
			opts.Params["al_level_min"] = *params.AlLevelMin
		}
		if params.AlLevelMax != nil {
			opts.Params["al_level_max"] = *params.AlLevelMax
		}
		if params.HasBan != nil {
			opts.Params["has_ban"] = *params.HasBan
		}
		if params.XboxConnected != nil {
			opts.Params["xbox_connected"] = *params.XboxConnected
		}
		if params.SteamConnected != nil {
			opts.Params["steam_connected"] = *params.SteamConnected
		}
		if params.PsnConnected != nil {
			opts.Params["psn_connected"] = *params.PsnConnected
		}
		if params.Subscription != nil {
			opts.Params["subscription"] = *params.Subscription
		}
		if params.SubscriptionLength != nil {
			opts.Params["subscription_length"] = *params.SubscriptionLength
		}
		if params.SubscriptionPeriod != nil {
			opts.Params["subscription_period"] = *params.SubscriptionPeriod
		}
		if params.HoursPlayed != nil {
			for k, v := range params.HoursPlayed {
				opts.Params[fmt.Sprintf("hours_played[%s]", k)] = v
			}
		}
		if params.HoursPlayedMax != nil {
			for k, v := range params.HoursPlayedMax {
				opts.Params[fmt.Sprintf("hours_played_max[%s]", k)] = v
			}
		}
		if params.Transactions != nil {
			opts.Params["transactions"] = *params.Transactions
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result EAResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// EpicGames Epic Games
func (s *CategorySearchService) EpicGames(ctx context.Context, params *EpicGamesParams) (*EpicGamesResponse, error) {
	path := "/epicgames"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Eg != nil {
			opts.Params["eg"] = *params.Eg
		}
		if params.Game != nil {
			opts.Params["game[]"] = params.Game
		}
		if params.ChangeEmail != nil {
			opts.Params["change_email"] = *params.ChangeEmail
		}
		if params.RlPurchases != nil {
			opts.Params["rl_purchases"] = *params.RlPurchases
		}
		if params.BalanceMin != nil {
			opts.Params["balance_min"] = *params.BalanceMin
		}
		if params.BalanceMax != nil {
			opts.Params["balance_max"] = *params.BalanceMax
		}
		if params.RewardsBalanceMin != nil {
			opts.Params["rewards_balance_min"] = *params.RewardsBalanceMin
		}
		if params.RewardsBalanceMax != nil {
			opts.Params["rewards_balance_max"] = *params.RewardsBalanceMax
		}
		if params.Gmin != nil {
			opts.Params["gmin"] = *params.Gmin
		}
		if params.Gmax != nil {
			opts.Params["gmax"] = *params.Gmax
		}
		if params.Country != nil {
			opts.Params["country[]"] = params.Country
		}
		if params.NotCountry != nil {
			opts.Params["not_country[]"] = params.NotCountry
		}
		if params.Daybreak != nil {
			opts.Params["daybreak"] = *params.Daybreak
		}
		if params.HoursPlayed != nil {
			for k, v := range params.HoursPlayed {
				opts.Params[fmt.Sprintf("hours_played[%s]", k)] = v
			}
		}
		if params.HoursPlayedMax != nil {
			for k, v := range params.HoursPlayedMax {
				opts.Params[fmt.Sprintf("hours_played_max[%s]", k)] = v
			}
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result EpicGamesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// EscapeFromTarkov Escape from Tarkov
func (s *CategorySearchService) EscapeFromTarkov(ctx context.Context, params *EscapeFromTarkovParams) (*EscapeFromTarkovResponse, error) {
	path := "/escape-from-tarkov"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Region != nil {
			opts.Params["region"] = *params.Region
		}
		if params.Version != nil {
			opts.Params["version[]"] = params.Version
		}
		if params.Reg != nil {
			opts.Params["reg"] = *params.Reg
		}
		if params.RegPeriod != nil {
			opts.Params["reg_period"] = *params.RegPeriod
		}
		if params.LevelMin != nil {
			opts.Params["level_min"] = *params.LevelMin
		}
		if params.LevelMax != nil {
			opts.Params["level_max"] = *params.LevelMax
		}
		if params.Pve != nil {
			opts.Params["pve"] = *params.Pve
		}
		if params.Side != nil {
			opts.Params["side"] = *params.Side
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result EscapeFromTarkovResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fortnite Fortnite
func (s *CategorySearchService) Fortnite(ctx context.Context, params *FortniteParams) (*FortniteResponse, error) {
	path := "/fortnite"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.TempEmail != nil {
			opts.Params["temp_email"] = *params.TempEmail
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Eg != nil {
			opts.Params["eg"] = *params.Eg
		}
		if params.Smin != nil {
			opts.Params["smin"] = *params.Smin
		}
		if params.Smax != nil {
			opts.Params["smax"] = *params.Smax
		}
		if params.Vbmin != nil {
			opts.Params["vbmin"] = *params.Vbmin
		}
		if params.Vbmax != nil {
			opts.Params["vbmax"] = *params.Vbmax
		}
		if params.Skin != nil {
			opts.Params["skin[]"] = params.Skin
		}
		if params.Pickaxe != nil {
			opts.Params["pickaxe[]"] = params.Pickaxe
		}
		if params.Glider != nil {
			opts.Params["glider[]"] = params.Glider
		}
		if params.Dance != nil {
			opts.Params["dance[]"] = params.Dance
		}
		if params.ChangeEmail != nil {
			opts.Params["change_email"] = *params.ChangeEmail
		}
		if params.Platform != nil {
			opts.Params["platform[]"] = params.Platform
		}
		if params.SkinsShopMin != nil {
			opts.Params["skins_shop_min"] = *params.SkinsShopMin
		}
		if params.SkinsShopMax != nil {
			opts.Params["skins_shop_max"] = *params.SkinsShopMax
		}
		if params.PickaxesShopMin != nil {
			opts.Params["pickaxes_shop_min"] = *params.PickaxesShopMin
		}
		if params.PickaxesShopMax != nil {
			opts.Params["pickaxes_shop_max"] = *params.PickaxesShopMax
		}
		if params.DancesShopMin != nil {
			opts.Params["dances_shop_min"] = *params.DancesShopMin
		}
		if params.DancesShopMax != nil {
			opts.Params["dances_shop_max"] = *params.DancesShopMax
		}
		if params.GlidersShopMin != nil {
			opts.Params["gliders_shop_min"] = *params.GlidersShopMin
		}
		if params.GlidersShopMax != nil {
			opts.Params["gliders_shop_max"] = *params.GlidersShopMax
		}
		if params.SkinsShopVbmin != nil {
			opts.Params["skins_shop_vbmin"] = *params.SkinsShopVbmin
		}
		if params.SkinsShopVbmax != nil {
			opts.Params["skins_shop_vbmax"] = *params.SkinsShopVbmax
		}
		if params.PickaxesShopVbmin != nil {
			opts.Params["pickaxes_shop_vbmin"] = *params.PickaxesShopVbmin
		}
		if params.PickaxesShopVbmax != nil {
			opts.Params["pickaxes_shop_vbmax"] = *params.PickaxesShopVbmax
		}
		if params.DancesShopVbmin != nil {
			opts.Params["dances_shop_vbmin"] = *params.DancesShopVbmin
		}
		if params.DancesShopVbmax != nil {
			opts.Params["dances_shop_vbmax"] = *params.DancesShopVbmax
		}
		if params.GlidersShopVbmin != nil {
			opts.Params["gliders_shop_vbmin"] = *params.GlidersShopVbmin
		}
		if params.GlidersShopVbmax != nil {
			opts.Params["gliders_shop_vbmax"] = *params.GlidersShopVbmax
		}
		if params.Bp != nil {
			opts.Params["bp"] = *params.Bp
		}
		if params.Lmin != nil {
			opts.Params["lmin"] = *params.Lmin
		}
		if params.Lmax != nil {
			opts.Params["lmax"] = *params.Lmax
		}
		if params.BpLmin != nil {
			opts.Params["bp_lmin"] = *params.BpLmin
		}
		if params.BpLmax != nil {
			opts.Params["bp_lmax"] = *params.BpLmax
		}
		if params.LastTransDate != nil {
			opts.Params["last_trans_date"] = *params.LastTransDate
		}
		if params.LastTransDatePeriod != nil {
			opts.Params["last_trans_date_period"] = *params.LastTransDatePeriod
		}
		if params.NoTrans != nil {
			opts.Params["no_trans"] = *params.NoTrans
		}
		if params.XboxLinkable != nil {
			opts.Params["xbox_linkable"] = *params.XboxLinkable
		}
		if params.PsnLinkable != nil {
			opts.Params["psn_linkable"] = *params.PsnLinkable
		}
		if params.Daybreak != nil {
			opts.Params["daybreak"] = *params.Daybreak
		}
		if params.RlPurchases != nil {
			opts.Params["rl_purchases"] = *params.RlPurchases
		}
		if params.Reg != nil {
			opts.Params["reg"] = *params.Reg
		}
		if params.RegPeriod != nil {
			opts.Params["reg_period"] = *params.RegPeriod
		}
		if params.RefundCreditsMin != nil {
			opts.Params["refund_credits_min"] = *params.RefundCreditsMin
		}
		if params.RefundCreditsMax != nil {
			opts.Params["refund_credits_max"] = *params.RefundCreditsMax
		}
		if params.PickaxeMin != nil {
			opts.Params["pickaxe_min"] = *params.PickaxeMin
		}
		if params.PickaxeMax != nil {
			opts.Params["pickaxe_max"] = *params.PickaxeMax
		}
		if params.Dmin != nil {
			opts.Params["dmin"] = *params.Dmin
		}
		if params.Dmax != nil {
			opts.Params["dmax"] = *params.Dmax
		}
		if params.Gmin != nil {
			opts.Params["gmin"] = *params.Gmin
		}
		if params.Gmax != nil {
			opts.Params["gmax"] = *params.Gmax
		}
		if params.Country != nil {
			opts.Params["country[]"] = params.Country
		}
		if params.NotCountry != nil {
			opts.Params["not_country[]"] = params.NotCountry
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result FortniteResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Gifts Gifts
func (s *CategorySearchService) Gifts(ctx context.Context, params *GiftsParams) (*GiftsResponse, error) {
	path := "/gifts"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.Subscription != nil {
			opts.Params["subscription"] = *params.Subscription
		}
		if params.SubscriptionLength != nil {
			opts.Params["subscription_length"] = *params.SubscriptionLength
		}
		if params.SubscriptionPeriod != nil {
			opts.Params["subscription_period"] = *params.SubscriptionPeriod
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result GiftsResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Hytale Hytale
func (s *CategorySearchService) Hytale(ctx context.Context, params *HytaleParams) (*HytaleResponse, error) {
	path := "/hytale"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.Edition != nil {
			opts.Params["edition[]"] = params.Edition
		}
		if params.ProfilesMin != nil {
			opts.Params["profiles_min"] = *params.ProfilesMin
		}
		if params.ProfilesMax != nil {
			opts.Params["profiles_max"] = *params.ProfilesMax
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result HytaleResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Instagram Instagram
func (s *CategorySearchService) Instagram(ctx context.Context, params *InstagramParams) (*InstagramResponse, error) {
	path := "/instagram"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Tel != nil {
			opts.Params["tel"] = *params.Tel
		}
		if params.Country != nil {
			opts.Params["country[]"] = params.Country
		}
		if params.NotCountry != nil {
			opts.Params["not_country[]"] = params.NotCountry
		}
		if params.Cookies != nil {
			opts.Params["cookies"] = *params.Cookies
		}
		if params.LoginWithoutCookies != nil {
			opts.Params["login_without_cookies"] = *params.LoginWithoutCookies
		}
		if params.FollowersMin != nil {
			opts.Params["followers_min"] = *params.FollowersMin
		}
		if params.FollowersMax != nil {
			opts.Params["followers_max"] = *params.FollowersMax
		}
		if params.PostMin != nil {
			opts.Params["post_min"] = *params.PostMin
		}
		if params.PostMax != nil {
			opts.Params["post_max"] = *params.PostMax
		}
		if params.Reg != nil {
			opts.Params["reg"] = *params.Reg
		}
		if params.RegPeriod != nil {
			opts.Params["reg_period"] = *params.RegPeriod
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result InstagramResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Mihoyo miHoYo
func (s *CategorySearchService) Mihoyo(ctx context.Context, params *MihoyoParams) (*MihoyoResponse, error) {
	path := "/mihoyo"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Email != nil {
			opts.Params["email"] = *params.Email
		}
		if params.EA != nil {
			opts.Params["ea"] = *params.EA
		}
		if params.Region != nil {
			opts.Params["region"] = params.Region
		}
		if params.NotRegion != nil {
			opts.Params["not_region"] = params.NotRegion
		}
		if params.GenshinCharacter != nil {
			opts.Params["genshin_character[]"] = params.GenshinCharacter
		}
		if params.GenshinCharacterConstellations != nil {
			for k, v := range params.GenshinCharacterConstellations {
				opts.Params[fmt.Sprintf("genshin_character_constellations[%s]", k)] = v
			}
		}
		if params.GenshinCharacterConstellationsMax != nil {
			for k, v := range params.GenshinCharacterConstellationsMax {
				opts.Params[fmt.Sprintf("genshin_character_constellations_max[%s]", k)] = v
			}
		}
		if params.GenshinWeapon != nil {
			opts.Params["genshin_weapon[]"] = params.GenshinWeapon
		}
		if params.GenshinCharMin != nil {
			opts.Params["genshin_char_min"] = *params.GenshinCharMin
		}
		if params.GenshinCharMax != nil {
			opts.Params["genshin_char_max"] = *params.GenshinCharMax
		}
		if params.GenshinLegendaryMin != nil {
			opts.Params["genshin_legendary_min"] = *params.GenshinLegendaryMin
		}
		if params.GenshinLegendaryMax != nil {
			opts.Params["genshin_legendary_max"] = *params.GenshinLegendaryMax
		}
		if params.GenshinLevelMin != nil {
			opts.Params["genshin_level_min"] = *params.GenshinLevelMin
		}
		if params.GenshinLevelMax != nil {
			opts.Params["genshin_level_max"] = *params.GenshinLevelMax
		}
		if params.GenshinLegendaryWeaponMin != nil {
			opts.Params["genshin_legendary_weapon_min"] = *params.GenshinLegendaryWeaponMin
		}
		if params.GenshinLegendaryWeaponMax != nil {
			opts.Params["genshin_legendary_weapon_max"] = *params.GenshinLegendaryWeaponMax
		}
		if params.ConstellationsMin != nil {
			opts.Params["constellations_min"] = *params.ConstellationsMin
		}
		if params.ConstellationsMax != nil {
			opts.Params["constellations_max"] = *params.ConstellationsMax
		}
		if params.GenshinAchievementMin != nil {
			opts.Params["genshin_achievement_min"] = *params.GenshinAchievementMin
		}
		if params.GenshinAchievementMax != nil {
			opts.Params["genshin_achievement_max"] = *params.GenshinAchievementMax
		}
		if params.GenshinCurrencyMin != nil {
			opts.Params["genshin_currency_min"] = *params.GenshinCurrencyMin
		}
		if params.GenshinCurrencyMax != nil {
			opts.Params["genshin_currency_max"] = *params.GenshinCurrencyMax
		}
		if params.HonkaiCharacter != nil {
			opts.Params["honkai_character[]"] = params.HonkaiCharacter
		}
		if params.HonkaiCharacterEidolons != nil {
			for k, v := range params.HonkaiCharacterEidolons {
				opts.Params[fmt.Sprintf("honkai_character_eidolons[%s]", k)] = v
			}
		}
		if params.HonkaiCharacterEidolonsMax != nil {
			for k, v := range params.HonkaiCharacterEidolonsMax {
				opts.Params[fmt.Sprintf("honkai_character_eidolons_max[%s]", k)] = v
			}
		}
		if params.HonkaiWeapon != nil {
			opts.Params["honkai_weapon[]"] = params.HonkaiWeapon
		}
		if params.HonkaiCharMin != nil {
			opts.Params["honkai_char_min"] = *params.HonkaiCharMin
		}
		if params.HonkaiCharMax != nil {
			opts.Params["honkai_char_max"] = *params.HonkaiCharMax
		}
		if params.HonkaiLegendaryMin != nil {
			opts.Params["honkai_legendary_min"] = *params.HonkaiLegendaryMin
		}
		if params.HonkaiLegendaryMax != nil {
			opts.Params["honkai_legendary_max"] = *params.HonkaiLegendaryMax
		}
		if params.HonkaiLevelMin != nil {
			opts.Params["honkai_level_min"] = *params.HonkaiLevelMin
		}
		if params.HonkaiLevelMax != nil {
			opts.Params["honkai_level_max"] = *params.HonkaiLevelMax
		}
		if params.HonkaiLegendaryWeaponMin != nil {
			opts.Params["honkai_legendary_weapon_min"] = *params.HonkaiLegendaryWeaponMin
		}
		if params.HonkaiLegendaryWeaponMax != nil {
			opts.Params["honkai_legendary_weapon_max"] = *params.HonkaiLegendaryWeaponMax
		}
		if params.EidolonsMin != nil {
			opts.Params["eidolons_min"] = *params.EidolonsMin
		}
		if params.EidolonsMax != nil {
			opts.Params["eidolons_max"] = *params.EidolonsMax
		}
		if params.HonkaiAchievementMin != nil {
			opts.Params["honkai_achievement_min"] = *params.HonkaiAchievementMin
		}
		if params.HonkaiAchievementMax != nil {
			opts.Params["honkai_achievement_max"] = *params.HonkaiAchievementMax
		}
		if params.HonkaiCurrencyMin != nil {
			opts.Params["honkai_currency_min"] = *params.HonkaiCurrencyMin
		}
		if params.HonkaiCurrencyMax != nil {
			opts.Params["honkai_currency_max"] = *params.HonkaiCurrencyMax
		}
		if params.ZenlessCharacter != nil {
			opts.Params["zenless_character[]"] = params.ZenlessCharacter
		}
		if params.ZenlessCharacterCinemas != nil {
			for k, v := range params.ZenlessCharacterCinemas {
				opts.Params[fmt.Sprintf("zenless_character_cinemas[%s]", k)] = v
			}
		}
		if params.ZenlessCharacterCinemasMax != nil {
			for k, v := range params.ZenlessCharacterCinemasMax {
				opts.Params[fmt.Sprintf("zenless_character_cinemas_max[%s]", k)] = v
			}
		}
		if params.ZenlessWeapon != nil {
			opts.Params["zenless_weapon[]"] = params.ZenlessWeapon
		}
		if params.ZenlessLegendaryMin != nil {
			opts.Params["zenless_legendary_min"] = *params.ZenlessLegendaryMin
		}
		if params.ZenlessLegendaryMax != nil {
			opts.Params["zenless_legendary_max"] = *params.ZenlessLegendaryMax
		}
		if params.CinemasMin != nil {
			opts.Params["cinemas_min"] = *params.CinemasMin
		}
		if params.CinemasMax != nil {
			opts.Params["cinemas_max"] = *params.CinemasMax
		}
		if params.ZenlessLegendaryWeaponMin != nil {
			opts.Params["zenless_legendary_weapon_min"] = *params.ZenlessLegendaryWeaponMin
		}
		if params.ZenlessLegendaryWeaponMax != nil {
			opts.Params["zenless_legendary_weapon_max"] = *params.ZenlessLegendaryWeaponMax
		}
		if params.ZenlessCharMin != nil {
			opts.Params["zenless_char_min"] = *params.ZenlessCharMin
		}
		if params.ZenlessCharMax != nil {
			opts.Params["zenless_char_max"] = *params.ZenlessCharMax
		}
		if params.ZenlessLevelMin != nil {
			opts.Params["zenless_level_min"] = *params.ZenlessLevelMin
		}
		if params.ZenlessLevelMax != nil {
			opts.Params["zenless_level_max"] = *params.ZenlessLevelMax
		}
		if params.ZenlessAchievementMin != nil {
			opts.Params["zenless_achievement_min"] = *params.ZenlessAchievementMin
		}
		if params.ZenlessAchievementMax != nil {
			opts.Params["zenless_achievement_max"] = *params.ZenlessAchievementMax
		}
		if params.ZenlessCurrencyMin != nil {
			opts.Params["zenless_currency_min"] = *params.ZenlessCurrencyMin
		}
		if params.ZenlessCurrencyMax != nil {
			opts.Params["zenless_currency_max"] = *params.ZenlessCurrencyMax
		}
		if params.Daybreak != nil {
			opts.Params["daybreak"] = *params.Daybreak
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result MihoyoResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Minecraft Minecraft
func (s *CategorySearchService) Minecraft(ctx context.Context, params *MinecraftParams) (*MinecraftResponse, error) {
	path := "/minecraft"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.Subscription != nil {
			opts.Params["subscription"] = *params.Subscription
		}
		if params.SubscriptionLength != nil {
			opts.Params["subscription_length"] = *params.SubscriptionLength
		}
		if params.SubscriptionPeriod != nil {
			opts.Params["subscription_period"] = *params.SubscriptionPeriod
		}
		if params.Autorenewal != nil {
			opts.Params["autorenewal"] = *params.Autorenewal
		}
		if params.Java != nil {
			opts.Params["java"] = *params.Java
		}
		if params.Bedrock != nil {
			opts.Params["bedrock"] = *params.Bedrock
		}
		if params.Dungeons != nil {
			opts.Params["dungeons"] = *params.Dungeons
		}
		if params.Legends != nil {
			opts.Params["legends"] = *params.Legends
		}
		if params.ChangeNickname != nil {
			opts.Params["change_nickname"] = *params.ChangeNickname
		}
		if params.Capes != nil {
			opts.Params["capes[]"] = params.Capes
		}
		if params.CapesMin != nil {
			opts.Params["capes_min"] = *params.CapesMin
		}
		if params.CapesMax != nil {
			opts.Params["capes_max"] = *params.CapesMax
		}
		if params.Country != nil {
			opts.Params["country[]"] = params.Country
		}
		if params.NotCountry != nil {
			opts.Params["not_country[]"] = params.NotCountry
		}
		if params.HypixelBan != nil {
			opts.Params["hypixel_ban"] = *params.HypixelBan
		}
		if params.HypixelSkyblockAPIEnabled != nil {
			opts.Params["hypixel_skyblock_api_enabled"] = *params.HypixelSkyblockAPIEnabled
		}
		if params.RankHypixel != nil {
			opts.Params["rank_hypixel[]"] = params.RankHypixel
		}
		if params.LevelHypixelMin != nil {
			opts.Params["level_hypixel_min"] = *params.LevelHypixelMin
		}
		if params.LevelHypixelMax != nil {
			opts.Params["level_hypixel_max"] = *params.LevelHypixelMax
		}
		if params.AchievementHypixelMin != nil {
			opts.Params["achievement_hypixel_min"] = *params.AchievementHypixelMin
		}
		if params.AchievementHypixelMax != nil {
			opts.Params["achievement_hypixel_max"] = *params.AchievementHypixelMax
		}
		if params.LevelHypixelSkyblockMin != nil {
			opts.Params["level_hypixel_skyblock_min"] = *params.LevelHypixelSkyblockMin
		}
		if params.LevelHypixelSkyblockMax != nil {
			opts.Params["level_hypixel_skyblock_max"] = *params.LevelHypixelSkyblockMax
		}
		if params.NetWorthHypixelSkyblockMin != nil {
			opts.Params["net_worth_hypixel_skyblock_min"] = *params.NetWorthHypixelSkyblockMin
		}
		if params.NetWorthHypixelSkyblockMax != nil {
			opts.Params["net_worth_hypixel_skyblock_max"] = *params.NetWorthHypixelSkyblockMax
		}
		if params.Reg != nil {
			opts.Params["reg"] = *params.Reg
		}
		if params.RegPeriod != nil {
			opts.Params["reg_period"] = *params.RegPeriod
		}
		if params.LastLoginHypixel != nil {
			opts.Params["last_login_hypixel"] = *params.LastLoginHypixel
		}
		if params.LastLoginHypixelPeriod != nil {
			opts.Params["last_login_hypixel_period"] = *params.LastLoginHypixelPeriod
		}
		if params.CanChangeDetails != nil {
			opts.Params["can_change_details"] = *params.CanChangeDetails
		}
		if params.NicknameLengthMin != nil {
			opts.Params["nickname_length_min"] = *params.NicknameLengthMin
		}
		if params.NicknameLengthMax != nil {
			opts.Params["nickname_length_max"] = *params.NicknameLengthMax
		}
		if params.HypixelBanParsed != nil {
			opts.Params["hypixel_ban_parsed"] = *params.HypixelBanParsed
		}
		if params.MinecoinsMin != nil {
			opts.Params["minecoins_min"] = *params.MinecoinsMin
		}
		if params.MinecoinsMax != nil {
			opts.Params["minecoins_max"] = *params.MinecoinsMax
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result MinecraftResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Riot Riot
func (s *CategorySearchService) Riot(ctx context.Context, params *RiotParams) (*RiotResponse, error) {
	path := "/riot"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Rmin != nil {
			opts.Params["rmin"] = *params.Rmin
		}
		if params.Rmax != nil {
			opts.Params["rmax"] = *params.Rmax
		}
		if params.LastRmin != nil {
			opts.Params["last_rmin"] = *params.LastRmin
		}
		if params.LastRmax != nil {
			opts.Params["last_rmax"] = *params.LastRmax
		}
		if params.PreviousRmin != nil {
			opts.Params["previous_rmin"] = *params.PreviousRmin
		}
		if params.PreviousRmax != nil {
			opts.Params["previous_rmax"] = *params.PreviousRmax
		}
		if params.WeaponSkin != nil {
			opts.Params["weaponSkin[]"] = params.WeaponSkin
		}
		if params.Buddy != nil {
			opts.Params["buddy[]"] = params.Buddy
		}
		if params.Agent != nil {
			opts.Params["agent[]"] = params.Agent
		}
		if params.Champion != nil {
			opts.Params["champion[]"] = params.Champion
		}
		if params.Skin != nil {
			opts.Params["skin[]"] = params.Skin
		}
		if params.Country != nil {
			opts.Params["country[]"] = params.Country
		}
		if params.NotCountry != nil {
			opts.Params["not_country[]"] = params.NotCountry
		}
		if params.Daybreak != nil {
			opts.Params["daybreak"] = *params.Daybreak
		}
		if params.ValorantLevelMin != nil {
			opts.Params["valorant_level_min"] = *params.ValorantLevelMin
		}
		if params.ValorantLevelMax != nil {
			opts.Params["valorant_level_max"] = *params.ValorantLevelMax
		}
		if params.LolLevelMin != nil {
			opts.Params["lol_level_min"] = *params.LolLevelMin
		}
		if params.LolLevelMax != nil {
			opts.Params["lol_level_max"] = *params.LolLevelMax
		}
		if params.InvMin != nil {
			opts.Params["inv_min"] = *params.InvMin
		}
		if params.InvMax != nil {
			opts.Params["inv_max"] = *params.InvMax
		}
		if params.VpMin != nil {
			opts.Params["vp_min"] = *params.VpMin
		}
		if params.VpMax != nil {
			opts.Params["vp_max"] = *params.VpMax
		}
		if params.ValorantSmin != nil {
			opts.Params["valorant_smin"] = *params.ValorantSmin
		}
		if params.ValorantSmax != nil {
			opts.Params["valorant_smax"] = *params.ValorantSmax
		}
		if params.ValorantRankType != nil {
			opts.Params["valorant_rank_type[]"] = params.ValorantRankType
		}
		if params.Amin != nil {
			opts.Params["amin"] = *params.Amin
		}
		if params.Amax != nil {
			opts.Params["amax"] = *params.Amax
		}
		if params.ValorantRegion != nil {
			opts.Params["valorant_region[]"] = params.ValorantRegion
		}
		if params.ValorantNotRegion != nil {
			opts.Params["valorant_not_region[]"] = params.ValorantNotRegion
		}
		if params.LolRegion != nil {
			opts.Params["lol_region[]"] = params.LolRegion
		}
		if params.LolNotRegion != nil {
			opts.Params["lol_not_region[]"] = params.LolNotRegion
		}
		if params.Knife != nil {
			opts.Params["knife"] = *params.Knife
		}
		if params.LolSmin != nil {
			opts.Params["lol_smin"] = *params.LolSmin
		}
		if params.LolSmax != nil {
			opts.Params["lol_smax"] = *params.LolSmax
		}
		if params.ChampionMin != nil {
			opts.Params["champion_min"] = *params.ChampionMin
		}
		if params.ChampionMax != nil {
			opts.Params["champion_max"] = *params.ChampionMax
		}
		if params.WinRateMin != nil {
			opts.Params["win_rate_min"] = *params.WinRateMin
		}
		if params.WinRateMax != nil {
			opts.Params["win_rate_max"] = *params.WinRateMax
		}
		if params.BlueMin != nil {
			opts.Params["blue_min"] = *params.BlueMin
		}
		if params.BlueMax != nil {
			opts.Params["blue_max"] = *params.BlueMax
		}
		if params.OrangeMin != nil {
			opts.Params["orange_min"] = *params.OrangeMin
		}
		if params.OrangeMax != nil {
			opts.Params["orange_max"] = *params.OrangeMax
		}
		if params.MythicMin != nil {
			opts.Params["mythic_min"] = *params.MythicMin
		}
		if params.MythicMax != nil {
			opts.Params["mythic_max"] = *params.MythicMax
		}
		if params.RiotMin != nil {
			opts.Params["riot_min"] = *params.RiotMin
		}
		if params.RiotMax != nil {
			opts.Params["riot_max"] = *params.RiotMax
		}
		if params.Email != nil {
			opts.Params["email"] = *params.Email
		}
		if params.Tel != nil {
			opts.Params["tel"] = *params.Tel
		}
		if params.ValorantKnifeMin != nil {
			opts.Params["valorant_knife_min"] = *params.ValorantKnifeMin
		}
		if params.ValorantKnifeMax != nil {
			opts.Params["valorant_knife_max"] = *params.ValorantKnifeMax
		}
		if params.RpMin != nil {
			opts.Params["rp_min"] = *params.RpMin
		}
		if params.RpMax != nil {
			opts.Params["rp_max"] = *params.RpMax
		}
		if params.FaMin != nil {
			opts.Params["fa_min"] = *params.FaMin
		}
		if params.FaMax != nil {
			opts.Params["fa_max"] = *params.FaMax
		}
		if params.LolRank != nil {
			opts.Params["lol_rank[]"] = params.LolRank
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result RiotResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Roblox Roblox
func (s *CategorySearchService) Roblox(ctx context.Context, params *RobloxParams) (*RobloxResponse, error) {
	path := "/roblox"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.Email != nil {
			opts.Params["email"] = *params.Email
		}
		if params.RobuxMin != nil {
			opts.Params["robux_min"] = *params.RobuxMin
		}
		if params.RobuxMax != nil {
			opts.Params["robux_max"] = *params.RobuxMax
		}
		if params.FriendsMin != nil {
			opts.Params["friends_min"] = *params.FriendsMin
		}
		if params.FriendsMax != nil {
			opts.Params["friends_max"] = *params.FriendsMax
		}
		if params.FollowersMin != nil {
			opts.Params["followers_min"] = *params.FollowersMin
		}
		if params.FollowersMax != nil {
			opts.Params["followers_max"] = *params.FollowersMax
		}
		if params.Country != nil {
			opts.Params["country"] = params.Country
		}
		if params.NotCountry != nil {
			opts.Params["not_country"] = params.NotCountry
		}
		if params.Reg != nil {
			opts.Params["reg"] = *params.Reg
		}
		if params.RegPeriod != nil {
			opts.Params["reg_period"] = *params.RegPeriod
		}
		if params.Subscription != nil {
			opts.Params["subscription"] = *params.Subscription
		}
		if params.SubscriptionLength != nil {
			opts.Params["subscription_length"] = *params.SubscriptionLength
		}
		if params.SubscriptionPeriod != nil {
			opts.Params["subscription_period"] = *params.SubscriptionPeriod
		}
		if params.Autorenewal != nil {
			opts.Params["autorenewal"] = *params.Autorenewal
		}
		if params.XboxConnected != nil {
			opts.Params["xbox_connected"] = *params.XboxConnected
		}
		if params.PsnConnected != nil {
			opts.Params["psn_connected"] = *params.PsnConnected
		}
		if params.Verified != nil {
			opts.Params["verified"] = *params.Verified
		}
		if params.AgeVerified != nil {
			opts.Params["age_verified"] = *params.AgeVerified
		}
		if params.IncomingRobuxTotalMin != nil {
			opts.Params["incoming_robux_total_min"] = *params.IncomingRobuxTotalMin
		}
		if params.IncomingRobuxTotalMax != nil {
			opts.Params["incoming_robux_total_max"] = *params.IncomingRobuxTotalMax
		}
		if params.LimitedPriceMin != nil {
			opts.Params["limited_price_min"] = *params.LimitedPriceMin
		}
		if params.LimitedPriceMax != nil {
			opts.Params["limited_price_max"] = *params.LimitedPriceMax
		}
		if params.GamepassMin != nil {
			opts.Params["gamepass_min"] = *params.GamepassMin
		}
		if params.GamepassMax != nil {
			opts.Params["gamepass_max"] = *params.GamepassMax
		}
		if params.GameDonations != nil {
			opts.Params["game_donations"] = *params.GameDonations
		}
		if params.InvMin != nil {
			opts.Params["inv_min"] = *params.InvMin
		}
		if params.InvMax != nil {
			opts.Params["inv_max"] = *params.InvMax
		}
		if params.UgcLimitedPriceMin != nil {
			opts.Params["ugc_limited_price_min"] = *params.UgcLimitedPriceMin
		}
		if params.UgcLimitedPriceMax != nil {
			opts.Params["ugc_limited_price_max"] = *params.UgcLimitedPriceMax
		}
		if params.CreditBalanceMin != nil {
			opts.Params["credit_balance_min"] = *params.CreditBalanceMin
		}
		if params.CreditBalanceMax != nil {
			opts.Params["credit_balance_max"] = *params.CreditBalanceMax
		}
		if params.OffsaleMin != nil {
			opts.Params["offsale_min"] = *params.OffsaleMin
		}
		if params.OffsaleMax != nil {
			opts.Params["offsale_max"] = *params.OffsaleMax
		}
		if params.Voice != nil {
			opts.Params["voice"] = *params.Voice
		}
		if params.AgeGroup != nil {
			opts.Params["age_group[]"] = params.AgeGroup
		}
		if params.NotAgeGroup != nil {
			opts.Params["not_age_group[]"] = params.NotAgeGroup
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result RobloxResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SocialClub Social Club
func (s *CategorySearchService) SocialClub(ctx context.Context, params *SocialClubParams) (*SocialClubResponse, error) {
	path := "/socialclub"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.Daybreak != nil {
			opts.Params["daybreak"] = *params.Daybreak
		}
		if params.LevelMin != nil {
			opts.Params["level_min"] = *params.LevelMin
		}
		if params.LevelMax != nil {
			opts.Params["level_max"] = *params.LevelMax
		}
		if params.CashMin != nil {
			opts.Params["cash_min"] = *params.CashMin
		}
		if params.CashMax != nil {
			opts.Params["cash_max"] = *params.CashMax
		}
		if params.BankCashMin != nil {
			opts.Params["bank_cash_min"] = *params.BankCashMin
		}
		if params.BankCashMax != nil {
			opts.Params["bank_cash_max"] = *params.BankCashMax
		}
		if params.Game != nil {
			opts.Params["game[]"] = params.Game
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result SocialClubResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Steam Steam
func (s *CategorySearchService) Steam(ctx context.Context, params *SteamParams) (*SteamResponse, error) {
	path := "/steam"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Game != nil {
			opts.Params["game[]"] = params.Game
		}
		if params.HoursPlayed != nil {
			for k, v := range params.HoursPlayed {
				opts.Params[fmt.Sprintf("hours_played[%s]", k)] = v
			}
		}
		if params.HoursPlayedMax != nil {
			for k, v := range params.HoursPlayedMax {
				opts.Params[fmt.Sprintf("hours_played_max[%s]", k)] = v
			}
		}
		if params.Eg != nil {
			opts.Params["eg"] = *params.Eg
		}
		if params.Vac != nil {
			opts.Params["vac[]"] = params.Vac
		}
		if params.VacSkipGameCheck != nil {
			opts.Params["vac_skip_game_check"] = *params.VacSkipGameCheck
		}
		if params.Rt != nil {
			opts.Params["rt"] = *params.Rt
		}
		if params.TradeBan != nil {
			opts.Params["trade_ban"] = *params.TradeBan
		}
		if params.TradeLimit != nil {
			opts.Params["trade_limit"] = *params.TradeLimit
		}
		if params.Daybreak != nil {
			opts.Params["daybreak"] = *params.Daybreak
		}
		if params.Limit != nil {
			opts.Params["limit"] = *params.Limit
		}
		if params.Mafile != nil {
			opts.Params["mafile"] = *params.Mafile
		}
		if params.Reg != nil {
			opts.Params["reg"] = *params.Reg
		}
		if params.RegPeriod != nil {
			opts.Params["reg_period"] = *params.RegPeriod
		}
		if params.Lmin != nil {
			opts.Params["lmin"] = *params.Lmin
		}
		if params.Lmax != nil {
			opts.Params["lmax"] = *params.Lmax
		}
		if params.Rmin != nil {
			opts.Params["rmin"] = *params.Rmin
		}
		if params.Rmax != nil {
			opts.Params["rmax"] = *params.Rmax
		}
		if params.WingmanRmin != nil {
			opts.Params["wingman_rmin"] = *params.WingmanRmin
		}
		if params.WingmanRmax != nil {
			opts.Params["wingman_rmax"] = *params.WingmanRmax
		}
		if params.NoVac != nil {
			opts.Params["no_vac"] = *params.NoVac
		}
		if params.MmBan != nil {
			opts.Params["mm_ban"] = *params.MmBan
		}
		if params.BalanceMin != nil {
			opts.Params["balance_min"] = *params.BalanceMin
		}
		if params.BalanceMax != nil {
			opts.Params["balance_max"] = *params.BalanceMax
		}
		if params.InvGame != nil {
			opts.Params["inv_game"] = *params.InvGame
		}
		if params.InvMin != nil {
			opts.Params["inv_min"] = *params.InvMin
		}
		if params.InvMax != nil {
			opts.Params["inv_max"] = *params.InvMax
		}
		if params.FriendsMin != nil {
			opts.Params["friends_min"] = *params.FriendsMin
		}
		if params.FriendsMax != nil {
			opts.Params["friends_max"] = *params.FriendsMax
		}
		if params.Gmin != nil {
			opts.Params["gmin"] = *params.Gmin
		}
		if params.Gmax != nil {
			opts.Params["gmax"] = *params.Gmax
		}
		if params.WinCountMin != nil {
			opts.Params["win_count_min"] = *params.WinCountMin
		}
		if params.WinCountMax != nil {
			opts.Params["win_count_max"] = *params.WinCountMax
		}
		if params.MedalID != nil {
			opts.Params["medal_id[]"] = params.MedalID
		}
		if params.MedalOperatorOr != nil {
			opts.Params["medal_operator_or"] = *params.MedalOperatorOr
		}
		if params.MedalMin != nil {
			opts.Params["medal_min"] = *params.MedalMin
		}
		if params.MedalMax != nil {
			opts.Params["medal_max"] = *params.MedalMax
		}
		if params.Gift != nil {
			opts.Params["gift[]"] = params.Gift
		}
		if params.GiftMin != nil {
			opts.Params["gift_min"] = *params.GiftMin
		}
		if params.GiftMax != nil {
			opts.Params["gift_max"] = *params.GiftMax
		}
		if params.RecentlyHoursMin != nil {
			opts.Params["recently_hours_min"] = *params.RecentlyHoursMin
		}
		if params.RecentlyHoursMax != nil {
			opts.Params["recently_hours_max"] = *params.RecentlyHoursMax
		}
		if params.Country != nil {
			opts.Params["country[]"] = params.Country
		}
		if params.NotCountry != nil {
			opts.Params["not_country[]"] = params.NotCountry
		}
		if params.Cs2ProfileRankMin != nil {
			opts.Params["cs2_profile_rank_min"] = *params.Cs2ProfileRankMin
		}
		if params.Cs2ProfileRankMax != nil {
			opts.Params["cs2_profile_rank_max"] = *params.Cs2ProfileRankMax
		}
		if params.SolommrMin != nil {
			opts.Params["solommr_min"] = *params.SolommrMin
		}
		if params.SolommrMax != nil {
			opts.Params["solommr_max"] = *params.SolommrMax
		}
		if params.D2GameCountMin != nil {
			opts.Params["d2_game_count_min"] = *params.D2GameCountMin
		}
		if params.D2GameCountMax != nil {
			opts.Params["d2_game_count_max"] = *params.D2GameCountMax
		}
		if params.D2WinCountMin != nil {
			opts.Params["d2_win_count_min"] = *params.D2WinCountMin
		}
		if params.D2WinCountMax != nil {
			opts.Params["d2_win_count_max"] = *params.D2WinCountMax
		}
		if params.D2BehaviorMin != nil {
			opts.Params["d2_behavior_min"] = *params.D2BehaviorMin
		}
		if params.D2BehaviorMax != nil {
			opts.Params["d2_behavior_max"] = *params.D2BehaviorMax
		}
		if params.FaceitLvlMin != nil {
			opts.Params["faceit_lvl_min"] = *params.FaceitLvlMin
		}
		if params.FaceitLvlMax != nil {
			opts.Params["faceit_lvl_max"] = *params.FaceitLvlMax
		}
		if params.PointsMin != nil {
			opts.Params["points_min"] = *params.PointsMin
		}
		if params.PointsMax != nil {
			opts.Params["points_max"] = *params.PointsMax
		}
		if params.RelevantGmin != nil {
			opts.Params["relevant_gmin"] = *params.RelevantGmin
		}
		if params.RelevantGmax != nil {
			opts.Params["relevant_gmax"] = *params.RelevantGmax
		}
		if params.LastTransDate != nil {
			opts.Params["last_trans_date"] = *params.LastTransDate
		}
		if params.LastTransDatePeriod != nil {
			opts.Params["last_trans_date_period"] = *params.LastTransDatePeriod
		}
		if params.LastTransDateLater != nil {
			opts.Params["last_trans_date_later"] = *params.LastTransDateLater
		}
		if params.LastTransDatePeriodLater != nil {
			opts.Params["last_trans_date_period_later"] = *params.LastTransDatePeriodLater
		}
		if params.NoTrans != nil {
			opts.Params["no_trans"] = *params.NoTrans
		}
		if params.Trans != nil {
			opts.Params["trans"] = *params.Trans
		}
		if params.GiftsPurchaseMin != nil {
			opts.Params["gifts_purchase_min"] = *params.GiftsPurchaseMin
		}
		if params.GiftsPurchaseMax != nil {
			opts.Params["gifts_purchase_max"] = *params.GiftsPurchaseMax
		}
		if params.RefundsPurchaseMin != nil {
			opts.Params["refunds_purchase_min"] = *params.RefundsPurchaseMin
		}
		if params.RefundsPurchaseMax != nil {
			opts.Params["refunds_purchase_max"] = *params.RefundsPurchaseMax
		}
		if params.IngamePurchaseMin != nil {
			opts.Params["ingame_purchase_min"] = *params.IngamePurchaseMin
		}
		if params.IngamePurchaseMax != nil {
			opts.Params["ingame_purchase_max"] = *params.IngamePurchaseMax
		}
		if params.GamesPurchaseMin != nil {
			opts.Params["games_purchase_min"] = *params.GamesPurchaseMin
		}
		if params.GamesPurchaseMax != nil {
			opts.Params["games_purchase_max"] = *params.GamesPurchaseMax
		}
		if params.PurchaseMin != nil {
			opts.Params["purchase_min"] = *params.PurchaseMin
		}
		if params.PurchaseMax != nil {
			opts.Params["purchase_max"] = *params.PurchaseMax
		}
		if params.HasActivatedKeys != nil {
			opts.Params["has_activated_keys"] = *params.HasActivatedKeys
		}
		if params.EloMin != nil {
			opts.Params["elo_min"] = *params.EloMin
		}
		if params.EloMax != nil {
			opts.Params["elo_max"] = *params.EloMax
		}
		if params.Cs2MapRank != nil {
			opts.Params["cs2_map_rank"] = *params.Cs2MapRank
		}
		if params.Cs2MapRmin != nil {
			opts.Params["cs2_map_rmin"] = *params.Cs2MapRmin
		}
		if params.Cs2MapRmax != nil {
			opts.Params["cs2_map_rmax"] = *params.Cs2MapRmax
		}
		if params.HasFaceit != nil {
			opts.Params["has_faceit"] = *params.HasFaceit
		}
		if params.FaceitCsgoLvlMin != nil {
			opts.Params["faceit_csgo_lvl_min"] = *params.FaceitCsgoLvlMin
		}
		if params.FaceitCsgoLvlMax != nil {
			opts.Params["faceit_csgo_lvl_max"] = *params.FaceitCsgoLvlMax
		}
		if params.RustDeathsMin != nil {
			opts.Params["rust_deaths_min"] = *params.RustDeathsMin
		}
		if params.RustDeathsMax != nil {
			opts.Params["rust_deaths_max"] = *params.RustDeathsMax
		}
		if params.RustKillsMin != nil {
			opts.Params["rust_kills_min"] = *params.RustKillsMin
		}
		if params.RustKillsMax != nil {
			opts.Params["rust_kills_max"] = *params.RustKillsMax
		}
		if params.D2LastMatchDate != nil {
			opts.Params["d2_last_match_date"] = *params.D2LastMatchDate
		}
		if params.D2LastMatchDatePeriod != nil {
			opts.Params["d2_last_match_date_period"] = *params.D2LastMatchDatePeriod
		}
		if params.CardsMin != nil {
			opts.Params["cards_min"] = *params.CardsMin
		}
		if params.CardsMax != nil {
			opts.Params["cards_max"] = *params.CardsMax
		}
		if params.CardsGamesMin != nil {
			opts.Params["cards_games_min"] = *params.CardsGamesMin
		}
		if params.CardsGamesMax != nil {
			opts.Params["cards_games_max"] = *params.CardsGamesMax
		}
		if params.SkipVacInv != nil {
			opts.Params["skip_vac_inv"] = *params.SkipVacInv
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result SteamResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Supercell Supercell
func (s *CategorySearchService) Supercell(ctx context.Context, params *SupercellParams) (*SupercellResponse, error) {
	path := "/supercell"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.Daybreak != nil {
			opts.Params["daybreak"] = *params.Daybreak
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Eg != nil {
			opts.Params["eg"] = *params.Eg
		}
		if params.Tel != nil {
			opts.Params["tel"] = *params.Tel
		}
		if params.BrawlLevelMin != nil {
			opts.Params["brawl_level_min"] = *params.BrawlLevelMin
		}
		if params.BrawlLevelMax != nil {
			opts.Params["brawl_level_max"] = *params.BrawlLevelMax
		}
		if params.BrawlCupMin != nil {
			opts.Params["brawl_cup_min"] = *params.BrawlCupMin
		}
		if params.BrawlCupMax != nil {
			opts.Params["brawl_cup_max"] = *params.BrawlCupMax
		}
		if params.BrawlWinsMin != nil {
			opts.Params["brawl_wins_min"] = *params.BrawlWinsMin
		}
		if params.BrawlWinsMax != nil {
			opts.Params["brawl_wins_max"] = *params.BrawlWinsMax
		}
		if params.BrawlPass != nil {
			opts.Params["brawl_pass"] = *params.BrawlPass
		}
		if params.Brawler != nil {
			opts.Params["brawler[]"] = params.Brawler
		}
		if params.BrawlersMin != nil {
			opts.Params["brawlers_min"] = *params.BrawlersMin
		}
		if params.BrawlersMax != nil {
			opts.Params["brawlers_max"] = *params.BrawlersMax
		}
		if params.LegendaryBrawlersMin != nil {
			opts.Params["legendary_brawlers_min"] = *params.LegendaryBrawlersMin
		}
		if params.LegendaryBrawlersMax != nil {
			opts.Params["legendary_brawlers_max"] = *params.LegendaryBrawlersMax
		}
		if params.RoyaleLevelMin != nil {
			opts.Params["royale_level_min"] = *params.RoyaleLevelMin
		}
		if params.RoyaleLevelMax != nil {
			opts.Params["royale_level_max"] = *params.RoyaleLevelMax
		}
		if params.RoyaleCupMin != nil {
			opts.Params["royale_cup_min"] = *params.RoyaleCupMin
		}
		if params.RoyaleCupMax != nil {
			opts.Params["royale_cup_max"] = *params.RoyaleCupMax
		}
		if params.RoyaleWinsMin != nil {
			opts.Params["royale_wins_min"] = *params.RoyaleWinsMin
		}
		if params.RoyaleWinsMax != nil {
			opts.Params["royale_wins_max"] = *params.RoyaleWinsMax
		}
		if params.KingLevelMin != nil {
			opts.Params["king_level_min"] = *params.KingLevelMin
		}
		if params.KingLevelMax != nil {
			opts.Params["king_level_max"] = *params.KingLevelMax
		}
		if params.RoyalePass != nil {
			opts.Params["royale_pass"] = *params.RoyalePass
		}
		if params.ClashLevelMin != nil {
			opts.Params["clash_level_min"] = *params.ClashLevelMin
		}
		if params.ClashLevelMax != nil {
			opts.Params["clash_level_max"] = *params.ClashLevelMax
		}
		if params.ClashCupMin != nil {
			opts.Params["clash_cup_min"] = *params.ClashCupMin
		}
		if params.ClashCupMax != nil {
			opts.Params["clash_cup_max"] = *params.ClashCupMax
		}
		if params.ClashWinsMin != nil {
			opts.Params["clash_wins_min"] = *params.ClashWinsMin
		}
		if params.ClashWinsMax != nil {
			opts.Params["clash_wins_max"] = *params.ClashWinsMax
		}
		if params.ClashPass != nil {
			opts.Params["clash_pass"] = *params.ClashPass
		}
		if params.TotalHeroesLevelMin != nil {
			opts.Params["total_heroes_level_min"] = *params.TotalHeroesLevelMin
		}
		if params.TotalHeroesLevelMax != nil {
			opts.Params["total_heroes_level_max"] = *params.TotalHeroesLevelMax
		}
		if params.TotalTroopsLevelMin != nil {
			opts.Params["total_troops_level_min"] = *params.TotalTroopsLevelMin
		}
		if params.TotalTroopsLevelMax != nil {
			opts.Params["total_troops_level_max"] = *params.TotalTroopsLevelMax
		}
		if params.TotalSpellsLevelMin != nil {
			opts.Params["total_spells_level_min"] = *params.TotalSpellsLevelMin
		}
		if params.TotalSpellsLevelMax != nil {
			opts.Params["total_spells_level_max"] = *params.TotalSpellsLevelMax
		}
		if params.TotalBuilderHeroesLevelMin != nil {
			opts.Params["total_builder_heroes_level_min"] = *params.TotalBuilderHeroesLevelMin
		}
		if params.TotalBuilderHeroesLevelMax != nil {
			opts.Params["total_builder_heroes_level_max"] = *params.TotalBuilderHeroesLevelMax
		}
		if params.TotalBuilderTroopsLevelMin != nil {
			opts.Params["total_builder_troops_level_min"] = *params.TotalBuilderTroopsLevelMin
		}
		if params.TotalBuilderTroopsLevelMax != nil {
			opts.Params["total_builder_troops_level_max"] = *params.TotalBuilderTroopsLevelMax
		}
		if params.TownHallLevelMin != nil {
			opts.Params["town_hall_level_min"] = *params.TownHallLevelMin
		}
		if params.TownHallLevelMax != nil {
			opts.Params["town_hall_level_max"] = *params.TownHallLevelMax
		}
		if params.BuilderHallLevelMin != nil {
			opts.Params["builder_hall_level_min"] = *params.BuilderHallLevelMin
		}
		if params.BuilderHallLevelMax != nil {
			opts.Params["builder_hall_level_max"] = *params.BuilderHallLevelMax
		}
		if params.BuilderHallCupMin != nil {
			opts.Params["builder_hall_cup_min"] = *params.BuilderHallCupMin
		}
		if params.BuilderHallCupMax != nil {
			opts.Params["builder_hall_cup_max"] = *params.BuilderHallCupMax
		}
		if params.CreationYearMin != nil {
			opts.Params["creation_year_min"] = *params.CreationYearMin
		}
		if params.CreationYearMax != nil {
			opts.Params["creation_year_max"] = *params.CreationYearMax
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result SupercellResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Telegram Telegram
func (s *CategorySearchService) Telegram(ctx context.Context, params *TelegramParams) (*TelegramResponse, error) {
	path := "/telegram"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.Spam != nil {
			opts.Params["spam"] = *params.Spam
		}
		if params.Password != nil {
			opts.Params["password"] = *params.Password
		}
		if params.Premium != nil {
			opts.Params["premium"] = *params.Premium
		}
		if params.PremiumExpiration != nil {
			opts.Params["premium_expiration"] = *params.PremiumExpiration
		}
		if params.PremiumExpirationPeriod != nil {
			opts.Params["premium_expiration_period"] = *params.PremiumExpirationPeriod
		}
		if params.Country != nil {
			opts.Params["country[]"] = params.Country
		}
		if params.NotCountry != nil {
			opts.Params["not_country[]"] = params.NotCountry
		}
		if params.Daybreak != nil {
			opts.Params["daybreak"] = *params.Daybreak
		}
		if params.MinChannels != nil {
			opts.Params["min_channels"] = *params.MinChannels
		}
		if params.MaxChannels != nil {
			opts.Params["max_channels"] = *params.MaxChannels
		}
		if params.MinChats != nil {
			opts.Params["min_chats"] = *params.MinChats
		}
		if params.MaxChats != nil {
			opts.Params["max_chats"] = *params.MaxChats
		}
		if params.MinConversations != nil {
			opts.Params["min_conversations"] = *params.MinConversations
		}
		if params.MaxConversations != nil {
			opts.Params["max_conversations"] = *params.MaxConversations
		}
		if params.MinAdmin != nil {
			opts.Params["min_admin"] = *params.MinAdmin
		}
		if params.MaxAdmin != nil {
			opts.Params["max_admin"] = *params.MaxAdmin
		}
		if params.MinAdminSub != nil {
			opts.Params["min_admin_sub"] = *params.MinAdminSub
		}
		if params.MaxAdminSub != nil {
			opts.Params["max_admin_sub"] = *params.MaxAdminSub
		}
		if params.DigMin != nil {
			opts.Params["dig_min"] = *params.DigMin
		}
		if params.DigMax != nil {
			opts.Params["dig_max"] = *params.DigMax
		}
		if params.MinContacts != nil {
			opts.Params["min_contacts"] = *params.MinContacts
		}
		if params.MaxContacts != nil {
			opts.Params["max_contacts"] = *params.MaxContacts
		}
		if params.MinStars != nil {
			opts.Params["min_stars"] = *params.MinStars
		}
		if params.MaxStars != nil {
			opts.Params["max_stars"] = *params.MaxStars
		}
		if params.Birthday != nil {
			opts.Params["birthday"] = *params.Birthday
		}
		if params.BirthdayPeriod != nil {
			opts.Params["birthday_period"] = *params.BirthdayPeriod
		}
		if params.BirthdayAfter != nil {
			opts.Params["birthday_after"] = *params.BirthdayAfter
		}
		if params.BirthdayAfterPeriod != nil {
			opts.Params["birthday_after_period"] = *params.BirthdayAfterPeriod
		}
		if params.MinID != nil {
			opts.Params["min_id"] = *params.MinID
		}
		if params.MaxID != nil {
			opts.Params["max_id"] = *params.MaxID
		}
		if params.AllowGeoSpamblock != nil {
			opts.Params["allow_geo_spamblock"] = *params.AllowGeoSpamblock
		}
		if params.MinGifts != nil {
			opts.Params["min_gifts"] = *params.MinGifts
		}
		if params.MaxGifts != nil {
			opts.Params["max_gifts"] = *params.MaxGifts
		}
		if params.MinNftGifts != nil {
			opts.Params["min_nft_gifts"] = *params.MinNftGifts
		}
		if params.MaxNftGifts != nil {
			opts.Params["max_nft_gifts"] = *params.MaxNftGifts
		}
		if params.MinGiftsStars != nil {
			opts.Params["min_gifts_stars"] = *params.MinGiftsStars
		}
		if params.MaxGiftsStars != nil {
			opts.Params["max_gifts_stars"] = *params.MaxGiftsStars
		}
		if params.MinGiftsConvertStars != nil {
			opts.Params["min_gifts_convert_stars"] = *params.MinGiftsConvertStars
		}
		if params.MaxGiftsConvertStars != nil {
			opts.Params["max_gifts_convert_stars"] = *params.MaxGiftsConvertStars
		}
		if params.DcID != nil {
			opts.Params["dc_id[]"] = params.DcID
		}
		if params.NotDcID != nil {
			opts.Params["not_dc_id[]"] = params.NotDcID
		}
		if params.Email != nil {
			opts.Params["email"] = *params.Email
		}
		if params.MinBots != nil {
			opts.Params["min_bots"] = *params.MinBots
		}
		if params.MaxBots != nil {
			opts.Params["max_bots"] = *params.MaxBots
		}
		if params.MinBotActiveUsers != nil {
			opts.Params["min_bot_active_users"] = *params.MinBotActiveUsers
		}
		if params.MaxBotActiveUsers != nil {
			opts.Params["max_bot_active_users"] = *params.MaxBotActiveUsers
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result TelegramResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// TikTok TikTok
func (s *CategorySearchService) TikTok(ctx context.Context, params *TikTokParams) (*TikTokResponse, error) {
	path := "/tiktok"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Tel != nil {
			opts.Params["tel"] = *params.Tel
		}
		if params.Reg != nil {
			opts.Params["reg"] = *params.Reg
		}
		if params.RegPeriod != nil {
			opts.Params["reg_period"] = *params.RegPeriod
		}
		if params.FollowersMin != nil {
			opts.Params["followers_min"] = *params.FollowersMin
		}
		if params.FollowersMax != nil {
			opts.Params["followers_max"] = *params.FollowersMax
		}
		if params.PostMin != nil {
			opts.Params["post_min"] = *params.PostMin
		}
		if params.PostMax != nil {
			opts.Params["post_max"] = *params.PostMax
		}
		if params.LikeMin != nil {
			opts.Params["like_min"] = *params.LikeMin
		}
		if params.LikeMax != nil {
			opts.Params["like_max"] = *params.LikeMax
		}
		if params.CoinsMin != nil {
			opts.Params["coins_min"] = *params.CoinsMin
		}
		if params.CoinsMax != nil {
			opts.Params["coins_max"] = *params.CoinsMax
		}
		if params.CookieLogin != nil {
			opts.Params["cookie_login"] = *params.CookieLogin
		}
		if params.Verified != nil {
			opts.Params["verified"] = *params.Verified
		}
		if params.Email != nil {
			opts.Params["email"] = *params.Email
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result TikTokResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Uplay Uplay
func (s *CategorySearchService) Uplay(ctx context.Context, params *UplayParams) (*UplayResponse, error) {
	path := "/uplay"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Game != nil {
			opts.Params["game[]"] = params.Game
		}
		if params.Country != nil {
			opts.Params["country[]"] = params.Country
		}
		if params.NotCountry != nil {
			opts.Params["not_country[]"] = params.NotCountry
		}
		if params.Daybreak != nil {
			opts.Params["daybreak"] = *params.Daybreak
		}
		if params.Gmin != nil {
			opts.Params["gmin"] = *params.Gmin
		}
		if params.Gmax != nil {
			opts.Params["gmax"] = *params.Gmax
		}
		if params.Subscription != nil {
			opts.Params["subscription"] = *params.Subscription
		}
		if params.SubscriptionLength != nil {
			opts.Params["subscription_length"] = *params.SubscriptionLength
		}
		if params.SubscriptionPeriod != nil {
			opts.Params["subscription_period"] = *params.SubscriptionPeriod
		}
		if params.R6LevelMin != nil {
			opts.Params["r6_level_min"] = *params.R6LevelMin
		}
		if params.R6LevelMax != nil {
			opts.Params["r6_level_max"] = *params.R6LevelMax
		}
		if params.R6RankMin != nil {
			opts.Params["r6_rank_min"] = *params.R6RankMin
		}
		if params.R6RankMax != nil {
			opts.Params["r6_rank_max"] = *params.R6RankMax
		}
		if params.R6OperatorsMin != nil {
			opts.Params["r6_operators_min"] = *params.R6OperatorsMin
		}
		if params.R6OperatorsMax != nil {
			opts.Params["r6_operators_max"] = *params.R6OperatorsMax
		}
		if params.R6Ban != nil {
			opts.Params["r6_ban"] = *params.R6Ban
		}
		if params.R6Smin != nil {
			opts.Params["r6_smin"] = *params.R6Smin
		}
		if params.R6Smax != nil {
			opts.Params["r6_smax"] = *params.R6Smax
		}
		if params.R6Skin != nil {
			opts.Params["r6_skin[]"] = params.R6Skin
		}
		if params.R6Operator != nil {
			opts.Params["r6_operator[]"] = params.R6Operator
		}
		if params.XboxConnected != nil {
			opts.Params["xbox_connected"] = *params.XboxConnected
		}
		if params.PsnConnected != nil {
			opts.Params["psn_connected"] = *params.PsnConnected
		}
		if params.SteamConnected != nil {
			opts.Params["steam_connected"] = *params.SteamConnected
		}
		if params.BalanceMin != nil {
			opts.Params["balance_min"] = *params.BalanceMin
		}
		if params.BalanceMax != nil {
			opts.Params["balance_max"] = *params.BalanceMax
		}
		if params.Transactions != nil {
			opts.Params["transactions"] = *params.Transactions
		}
		if params.Reg != nil {
			opts.Params["reg"] = *params.Reg
		}
		if params.RegPeriod != nil {
			opts.Params["reg_period"] = *params.RegPeriod
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result UplayResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Vpn VPN
func (s *CategorySearchService) Vpn(ctx context.Context, params *VpnParams) (*VpnResponse, error) {
	path := "/vpn"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.Service != nil {
			opts.Params["service[]"] = params.Service
		}
		if params.SubscriptionLength != nil {
			opts.Params["subscription_length"] = *params.SubscriptionLength
		}
		if params.SubscriptionPeriod != nil {
			opts.Params["subscription_period"] = *params.SubscriptionPeriod
		}
		if params.Autorenewal != nil {
			opts.Params["autorenewal"] = *params.Autorenewal
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result VpnResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// WOT World of Tanks
func (s *CategorySearchService) WOT(ctx context.Context, params *WOTParams) (*WOTResponse, error) {
	path := "/world-of-tanks"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Tel != nil {
			opts.Params["tel"] = *params.Tel
		}
		if params.Daybreak != nil {
			opts.Params["daybreak"] = *params.Daybreak
		}
		if params.BattlesMin != nil {
			opts.Params["battles_min"] = *params.BattlesMin
		}
		if params.BattlesMax != nil {
			opts.Params["battles_max"] = *params.BattlesMax
		}
		if params.GoldMin != nil {
			opts.Params["gold_min"] = *params.GoldMin
		}
		if params.GoldMax != nil {
			opts.Params["gold_max"] = *params.GoldMax
		}
		if params.SilverMin != nil {
			opts.Params["silver_min"] = *params.SilverMin
		}
		if params.SilverMax != nil {
			opts.Params["silver_max"] = *params.SilverMax
		}
		if params.TopMin != nil {
			opts.Params["top_min"] = *params.TopMin
		}
		if params.TopMax != nil {
			opts.Params["top_max"] = *params.TopMax
		}
		if params.PremMin != nil {
			opts.Params["prem_min"] = *params.PremMin
		}
		if params.PremMax != nil {
			opts.Params["prem_max"] = *params.PremMax
		}
		if params.TopPremMin != nil {
			opts.Params["top_prem_min"] = *params.TopPremMin
		}
		if params.TopPremMax != nil {
			opts.Params["top_prem_max"] = *params.TopPremMax
		}
		if params.WinPmin != nil {
			opts.Params["win_pmin"] = *params.WinPmin
		}
		if params.WinPmax != nil {
			opts.Params["win_pmax"] = *params.WinPmax
		}
		if params.Tank != nil {
			opts.Params["tank[]"] = params.Tank
		}
		if params.Region != nil {
			opts.Params["region[]"] = params.Region
		}
		if params.NotRegion != nil {
			opts.Params["not_region[]"] = params.NotRegion
		}
		if params.Premium != nil {
			opts.Params["premium"] = *params.Premium
		}
		if params.PremiumExpiration != nil {
			opts.Params["premium_expiration"] = *params.PremiumExpiration
		}
		if params.PremiumExpirationPeriod != nil {
			opts.Params["premium_expiration_period"] = *params.PremiumExpirationPeriod
		}
		if params.Clan != nil {
			opts.Params["clan"] = *params.Clan
		}
		if params.ClanRole != nil {
			opts.Params["clan_role[]"] = params.ClanRole
		}
		if params.NotClanRole != nil {
			opts.Params["not_clan_role[]"] = params.NotClanRole
		}
		if params.ClanGoldMin != nil {
			opts.Params["clan_gold_min"] = *params.ClanGoldMin
		}
		if params.ClanGoldMax != nil {
			opts.Params["clan_gold_max"] = *params.ClanGoldMax
		}
		if params.ClanCreditsMin != nil {
			opts.Params["clan_credits_min"] = *params.ClanCreditsMin
		}
		if params.ClanCreditsMax != nil {
			opts.Params["clan_credits_max"] = *params.ClanCreditsMax
		}
		if params.ClanCrystalMin != nil {
			opts.Params["clan_crystal_min"] = *params.ClanCrystalMin
		}
		if params.ClanCrystalMax != nil {
			opts.Params["clan_crystal_max"] = *params.ClanCrystalMax
		}
		if params.Country != nil {
			opts.Params["country[]"] = params.Country
		}
		if params.NotCountry != nil {
			opts.Params["not_country[]"] = params.NotCountry
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result WOTResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Warface Warface
func (s *CategorySearchService) Warface(ctx context.Context, params *WarfaceParams) (*WarfaceResponse, error) {
	path := "/warface"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.RankMin != nil {
			opts.Params["rank_min"] = *params.RankMin
		}
		if params.RankMax != nil {
			opts.Params["rank_max"] = *params.RankMax
		}
		if params.BonusRankMin != nil {
			opts.Params["bonus_rank_min"] = *params.BonusRankMin
		}
		if params.BonusRankMax != nil {
			opts.Params["bonus_rank_max"] = *params.BonusRankMax
		}
		if params.Tel != nil {
			opts.Params["tel"] = *params.Tel
		}
		if params.Daybreak != nil {
			opts.Params["daybreak"] = *params.Daybreak
		}
		if params.KreditsMin != nil {
			opts.Params["kredits_min"] = *params.KreditsMin
		}
		if params.KreditsMax != nil {
			opts.Params["kredits_max"] = *params.KreditsMax
		}
		if params.TotalKreditsMin != nil {
			opts.Params["total_kredits_min"] = *params.TotalKreditsMin
		}
		if params.TotalKreditsMax != nil {
			opts.Params["total_kredits_max"] = *params.TotalKreditsMax
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result WarfaceResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// WotBlitz WoT Blitz
func (s *CategorySearchService) WotBlitz(ctx context.Context, params *WotBlitzParams) (*WotBlitzResponse, error) {
	path := "/wot-blitz"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Title != nil {
			opts.Params["title"] = *params.Title
		}
		if params.OrderBy != nil {
			opts.Params["order_by"] = *params.OrderBy
		}
		if params.TagID != nil {
			opts.Params["tag_id[]"] = params.TagID
		}
		if params.NotTagID != nil {
			opts.Params["not_tag_id[]"] = params.NotTagID
		}
		if params.PublicTagID != nil {
			opts.Params["public_tag_id[]"] = params.PublicTagID
		}
		if params.NotPublicTagID != nil {
			opts.Params["not_public_tag_id[]"] = params.NotPublicTagID
		}
		if params.Origin != nil {
			opts.Params["origin[]"] = params.Origin
		}
		if params.NotOrigin != nil {
			opts.Params["not_origin[]"] = params.NotOrigin
		}
		if params.UserID != nil {
			opts.Params["user_id"] = *params.UserID
		}
		if params.Nsb != nil {
			opts.Params["nsb"] = *params.Nsb
		}
		if params.Sb != nil {
			opts.Params["sb"] = *params.Sb
		}
		if params.NsbByMe != nil {
			opts.Params["nsb_by_me"] = *params.NsbByMe
		}
		if params.SbByMe != nil {
			opts.Params["sb_by_me"] = *params.SbByMe
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.EmailLoginData != nil {
			opts.Params["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailProvider != nil {
			opts.Params["email_provider[]"] = params.EmailProvider
		}
		if params.NotEmailProvider != nil {
			opts.Params["not_email_provider[]"] = *params.NotEmailProvider
		}
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
		}
		if params.EmailType != nil {
			opts.Params["email_type[]"] = params.EmailType
		}
		if params.ItemDomain != nil {
			opts.Params["item_domain"] = *params.ItemDomain
		}
		if params.Tel != nil {
			opts.Params["tel"] = *params.Tel
		}
		if params.Daybreak != nil {
			opts.Params["daybreak"] = *params.Daybreak
		}
		if params.BattlesMin != nil {
			opts.Params["battles_min"] = *params.BattlesMin
		}
		if params.BattlesMax != nil {
			opts.Params["battles_max"] = *params.BattlesMax
		}
		if params.GoldMin != nil {
			opts.Params["gold_min"] = *params.GoldMin
		}
		if params.GoldMax != nil {
			opts.Params["gold_max"] = *params.GoldMax
		}
		if params.SilverMin != nil {
			opts.Params["silver_min"] = *params.SilverMin
		}
		if params.SilverMax != nil {
			opts.Params["silver_max"] = *params.SilverMax
		}
		if params.TopMin != nil {
			opts.Params["top_min"] = *params.TopMin
		}
		if params.TopMax != nil {
			opts.Params["top_max"] = *params.TopMax
		}
		if params.PremMin != nil {
			opts.Params["prem_min"] = *params.PremMin
		}
		if params.PremMax != nil {
			opts.Params["prem_max"] = *params.PremMax
		}
		if params.TopPremMin != nil {
			opts.Params["top_prem_min"] = *params.TopPremMin
		}
		if params.TopPremMax != nil {
			opts.Params["top_prem_max"] = *params.TopPremMax
		}
		if params.WinPmin != nil {
			opts.Params["win_pmin"] = *params.WinPmin
		}
		if params.WinPmax != nil {
			opts.Params["win_pmax"] = *params.WinPmax
		}
		if params.Tank != nil {
			opts.Params["tank[]"] = params.Tank
		}
		if params.Region != nil {
			opts.Params["region[]"] = params.Region
		}
		if params.NotRegion != nil {
			opts.Params["not_region[]"] = params.NotRegion
		}
		if params.Premium != nil {
			opts.Params["premium"] = *params.Premium
		}
		if params.PremiumExpiration != nil {
			opts.Params["premium_expiration"] = *params.PremiumExpiration
		}
		if params.PremiumExpirationPeriod != nil {
			opts.Params["premium_expiration_period"] = *params.PremiumExpirationPeriod
		}
		if params.Clan != nil {
			opts.Params["clan"] = *params.Clan
		}
		if params.ClanRole != nil {
			opts.Params["clan_role[]"] = params.ClanRole
		}
		if params.NotClanRole != nil {
			opts.Params["not_clan_role[]"] = params.NotClanRole
		}
		if params.ClanGoldMin != nil {
			opts.Params["clan_gold_min"] = *params.ClanGoldMin
		}
		if params.ClanGoldMax != nil {
			opts.Params["clan_gold_max"] = *params.ClanGoldMax
		}
		if params.ClanCreditsMin != nil {
			opts.Params["clan_credits_min"] = *params.ClanCreditsMin
		}
		if params.ClanCreditsMax != nil {
			opts.Params["clan_credits_max"] = *params.ClanCreditsMax
		}
		if params.ClanCrystalMin != nil {
			opts.Params["clan_crystal_min"] = *params.ClanCrystalMin
		}
		if params.ClanCrystalMax != nil {
			opts.Params["clan_crystal_max"] = *params.ClanCrystalMax
		}
		if params.Country != nil {
			opts.Params["country[]"] = params.Country
		}
		if params.NotCountry != nil {
			opts.Params["not_country[]"] = params.NotCountry
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result WotBlitzResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CustomDiscountsService handles CustomDiscounts operations.
type CustomDiscountsService struct {
	r Requester
}

// Create Create Custom Discount
func (s *CustomDiscountsService) Create(ctx context.Context, categoryID CustomDiscountsCreateCategoryID, discountPercent float64, minPrice float64, userID int, params *CreateParams) (*CreateResponse, error) {
	path := "/custom-discounts"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["category_id"] = categoryID
	opts.JSON["discount_percent"] = discountPercent
	opts.JSON["min_price"] = minPrice
	opts.JSON["user_id"] = userID
	if params != nil {
		if params.AdditionalData != nil {
			opts.JSON["additional_data"] = *params.AdditionalData
		}
		if params.IsTest != nil {
			opts.JSON["is_test"] = *params.IsTest
		}
		if params.Lifetime != nil {
			opts.JSON["lifetime"] = *params.Lifetime
		}
		if params.RequiredTelegramID != nil {
			opts.JSON["required_telegram_id"] = *params.RequiredTelegramID
		}
		if params.RequiredTelegramUsername != nil {
			opts.JSON["required_telegram_username"] = *params.RequiredTelegramUsername
		}
		if params.URLCallback != nil {
			opts.JSON["url_callback"] = *params.URLCallback
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

// Delete Delete Custom Discount
func (s *CustomDiscountsService) Delete(ctx context.Context, discountID int) (*SaveChanges, error) {
	path := "/custom-discounts"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["discount_id"] = discountID
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

// Edit Edit Custom Discount
func (s *CustomDiscountsService) Edit(ctx context.Context, discountID int, params *EditParams) (*EditResponse, error) {
	path := "/custom-discounts"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["discount_id"] = discountID
	if params != nil {
		if params.AllowAskDiscount != nil {
			opts.JSON["allow_ask_discount"] = *params.AllowAskDiscount
		}
		if params.Currency != nil {
			opts.JSON["currency"] = *params.Currency
		}
		if params.Description != nil {
			opts.JSON["description"] = *params.Description
		}
		if params.EmailLoginData != nil {
			opts.JSON["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailType != nil {
			opts.JSON["email_type"] = *params.EmailType
		}
		if params.Information != nil {
			opts.JSON["information"] = *params.Information
		}
		if params.ItemOrigin != nil {
			opts.JSON["item_origin"] = *params.ItemOrigin
		}
		if params.Price != nil {
			opts.JSON["price"] = *params.Price
		}
		if params.ProxyID != nil {
			opts.JSON["proxy_id"] = *params.ProxyID
		}
		if params.Title != nil {
			opts.JSON["title"] = *params.Title
		}
		if params.TitleEn != nil {
			opts.JSON["title_en"] = *params.TitleEn
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

// Get Get Custom Discounts
func (s *CustomDiscountsService) Get(ctx context.Context) (*GetResponse, error) {
	path := "/custom-discounts"
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

// IMAPService handles IMAP operations.
type IMAPService struct {
	r Requester
}

// Create Create IMAP Configuration
func (s *IMAPService) Create(ctx context.Context, domain string, imapServer string, port int, secure bool) (*SaveChanges, error) {
	path := "/imap"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["domain"] = domain
	opts.JSON["imap_server"] = imapServer
	opts.JSON["port"] = port
	opts.JSON["secure"] = secure
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

// Delete Delete IMAP Configuration
func (s *IMAPService) Delete(ctx context.Context, domain string) (*SaveChanges, error) {
	path := "/imap"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["domain"] = domain
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

// InvoicesService handles Invoices operations.
type InvoicesService struct {
	r Requester
}

// Create Create Invoice
func (s *InvoicesService) Create(ctx context.Context, amount float64, comment string, currency InvoicesCreateCurrency, merchantID int, paymentID string, urlSuccess string, params *CreateParams) (*CreateResponse, error) {
	path := "/invoice"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["amount"] = amount
	opts.JSON["comment"] = comment
	opts.JSON["currency"] = currency
	opts.JSON["merchant_id"] = merchantID
	opts.JSON["payment_id"] = paymentID
	opts.JSON["url_success"] = urlSuccess
	if params != nil {
		if params.AdditionalData != nil {
			opts.JSON["additional_data"] = *params.AdditionalData
		}
		if params.IsTest != nil {
			opts.JSON["is_test"] = *params.IsTest
		}
		if params.Lifetime != nil {
			opts.JSON["lifetime"] = *params.Lifetime
		}
		if params.RequiredTelegramID != nil {
			opts.JSON["required_telegram_id"] = *params.RequiredTelegramID
		}
		if params.RequiredTelegramUsername != nil {
			opts.JSON["required_telegram_username"] = *params.RequiredTelegramUsername
		}
		if params.URLCallback != nil {
			opts.JSON["url_callback"] = *params.URLCallback
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

// Get Get Invoice
func (s *InvoicesService) Get(ctx context.Context, params *GetParams) (*GetResponse, error) {
	path := "/invoice"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
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

// List Get Invoice List
func (s *InvoicesService) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/invoice/list"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.Status != nil {
			opts.Params["status"] = *params.Status
		}
		if params.Amount != nil {
			opts.Params["amount"] = *params.Amount
		}
		if params.MerchantID != nil {
			opts.Params["merchant_id"] = *params.MerchantID
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

// PaymentsService handles Payments operations.
type PaymentsService struct {
	r Requester
}

// BalanceExchange Exchange Balance
func (s *PaymentsService) BalanceExchange(ctx context.Context, amount int, fromBalance string, toBalance string) (*BalanceExchange, error) {
	path := "/balance/exchange"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["amount"] = amount
	opts.JSON["from_balance"] = fromBalance
	opts.JSON["to_balance"] = toBalance
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result BalanceExchange
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Cancel Cancel Transfer
func (s *PaymentsService) Cancel(ctx context.Context, paymentID int) (*Status, error) {
	path := "/balance/transfer/cancel"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["payment_id"] = paymentID
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result Status
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Create Create Auto Payment
func (s *PaymentsService) Create(ctx context.Context, amount float64, day PaymentsCreateDay, usernameReceiver string, params *CreateParams) (*CreateResponse, error) {
	path := "/auto-payment"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["amount"] = amount
	opts.JSON["day"] = day
	opts.JSON["username_receiver"] = usernameReceiver
	if params != nil {
		if params.AdditionalData != nil {
			opts.JSON["additional_data"] = *params.AdditionalData
		}
		if params.IsTest != nil {
			opts.JSON["is_test"] = *params.IsTest
		}
		if params.Lifetime != nil {
			opts.JSON["lifetime"] = *params.Lifetime
		}
		if params.RequiredTelegramID != nil {
			opts.JSON["required_telegram_id"] = *params.RequiredTelegramID
		}
		if params.RequiredTelegramUsername != nil {
			opts.JSON["required_telegram_username"] = *params.RequiredTelegramUsername
		}
		if params.URLCallback != nil {
			opts.JSON["url_callback"] = *params.URLCallback
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

// Currency Get Currency
func (s *PaymentsService) Currency(ctx context.Context) (*CurrencyResponse, error) {
	path := "/currency"
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result CurrencyResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Delete Delete Auto Payment
func (s *PaymentsService) Delete(ctx context.Context, autoPaymentID int) (*Status, error) {
	path := "/auto-payment"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["auto_payment_id"] = autoPaymentID
	raw, err := s.r.Request(ctx, "DELETE", path, opts)
	if err != nil {
		return nil, err
	}
	var result Status
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Fee Check Transfer Fee
func (s *PaymentsService) Fee(ctx context.Context, params *FeeParams) (*FeeResponse, error) {
	path := "/balance/transfer/fee"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Amount != nil {
			opts.Params["amount"] = *params.Amount
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result FeeResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// History Payments History
func (s *PaymentsService) History(ctx context.Context, params *HistoryParams) (*HistoryResponse, error) {
	path := "/user/payments"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.Type_ != nil {
			opts.Params["type"] = *params.Type_
		}
		if params.Pmin != nil {
			opts.Params["pmin"] = *params.Pmin
		}
		if params.Pmax != nil {
			opts.Params["pmax"] = *params.Pmax
		}
		if params.Currency != nil {
			opts.Params["currency"] = *params.Currency
		}
		if params.Page != nil {
			opts.Params["page"] = *params.Page
		}
		if params.OperationIDLt != nil {
			opts.Params["operation_id_lt"] = *params.OperationIDLt
		}
		if params.Receiver != nil {
			opts.Params["receiver"] = *params.Receiver
		}
		if params.Sender != nil {
			opts.Params["sender"] = *params.Sender
		}
		if params.IsAPI != nil {
			opts.Params["is_api"] = *params.IsAPI
		}
		if params.StartDate != nil {
			opts.Params["startDate"] = *params.StartDate
		}
		if params.EndDate != nil {
			opts.Params["endDate"] = *params.EndDate
		}
		if params.Wallet != nil {
			opts.Params["wallet"] = *params.Wallet
		}
		if params.Comment != nil {
			opts.Params["comment"] = *params.Comment
		}
		if params.IsHold != nil {
			opts.Params["is_hold"] = *params.IsHold
		}
		if params.ShowPaymentStats != nil {
			opts.Params["show_payment_stats"] = *params.ShowPaymentStats
		}
	}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result HistoryResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// List Get Auto Payments
func (s *PaymentsService) List(ctx context.Context) (*ListResponse, error) {
	path := "/auto-payments"
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

// ListGet Get List Of Balances
func (s *PaymentsService) ListGet(ctx context.Context) (*BalanceExchange, error) {
	path := "/balance/exchange"
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result BalanceExchange
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Payout Create Payout
func (s *PaymentsService) Payout(ctx context.Context, amount float64, currency PaymentsPayoutCurrency, paymentSystem string, wallet string, params *PayoutParams) (*SaveChanges, error) {
	path := "/balance/payout"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["amount"] = amount
	opts.JSON["currency"] = currency
	opts.JSON["payment_system"] = paymentSystem
	opts.JSON["wallet"] = wallet
	if params != nil {
		if params.Extra != nil {
			for k, v := range params.Extra {
				opts.JSON[fmt.Sprintf("extra[%s]", k)] = v
			}
		}
		if params.IncludeFee != nil {
			opts.JSON["include_fee"] = *params.IncludeFee
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

// PayoutServices Get Payout Services
func (s *PaymentsService) PayoutServices(ctx context.Context) (*PayoutServicesResponse, error) {
	path := "/balance/payout/services"
	opts := RequestOptions{}
	raw, err := s.r.Request(ctx, "GET", path, opts)
	if err != nil {
		return nil, err
	}
	var result PayoutServicesResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Transfer Transfer Money
func (s *PaymentsService) Transfer(ctx context.Context, amount int, currency PaymentsTransferCurrency, params *TransferParams) (*Status, error) {
	path := "/balance/transfer"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	opts.JSON["amount"] = amount
	opts.JSON["currency"] = currency
	if params != nil {
		if params.Comment != nil {
			opts.JSON["comment"] = *params.Comment
		}
		if params.HoldLengthOption != nil {
			opts.JSON["hold_length_option"] = *params.HoldLengthOption
		}
		if params.HoldLengthValue != nil {
			opts.JSON["hold_length_value"] = *params.HoldLengthValue
		}
		if params.TelegramDeal != nil {
			opts.JSON["telegram_deal"] = *params.TelegramDeal
		}
		if params.TelegramUsername != nil {
			opts.JSON["telegram_username"] = *params.TelegramUsername
		}
		if params.TransferHold != nil {
			opts.JSON["transfer_hold"] = *params.TransferHold
		}
		if params.UserID != nil {
			opts.JSON["user_id"] = *params.UserID
		}
		if params.Username != nil {
			opts.JSON["username"] = *params.Username
		}
	}
	raw, err := s.r.Request(ctx, "POST", path, opts)
	if err != nil {
		return nil, err
	}
	var result Status
	if err := json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ProfileService handles Profile operations.
type ProfileService struct {
	r Requester
}

// Edit Edit Market Settings
func (s *ProfileService) Edit(ctx context.Context, params *EditParams) (*SaveChanges, error) {
	path := "/me"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.AllowAskDiscount != nil {
			opts.JSON["allow_ask_discount"] = *params.AllowAskDiscount
		}
		if params.Currency != nil {
			opts.JSON["currency"] = *params.Currency
		}
		if params.Description != nil {
			opts.JSON["description"] = *params.Description
		}
		if params.EmailLoginData != nil {
			opts.JSON["email_login_data"] = *params.EmailLoginData
		}
		if params.EmailType != nil {
			opts.JSON["email_type"] = *params.EmailType
		}
		if params.Information != nil {
			opts.JSON["information"] = *params.Information
		}
		if params.ItemOrigin != nil {
			opts.JSON["item_origin"] = *params.ItemOrigin
		}
		if params.Price != nil {
			opts.JSON["price"] = *params.Price
		}
		if params.ProxyID != nil {
			opts.JSON["proxy_id"] = *params.ProxyID
		}
		if params.Title != nil {
			opts.JSON["title"] = *params.Title
		}
		if params.TitleEn != nil {
			opts.JSON["title_en"] = *params.TitleEn
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

// Get Get Profile
func (s *ProfileService) Get(ctx context.Context, params *GetParams) (*GetResponse, error) {
	path := "/me"
	opts := RequestOptions{}
	opts.Params = make(map[string]interface{})
	if params != nil {
		if params.ParseSameItemIds != nil {
			opts.Params["parse_same_item_ids"] = *params.ParseSameItemIds
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

// ProxyService handles Proxy operations.
type ProxyService struct {
	r Requester
}

// Add Add Proxy
func (s *ProxyService) Add(ctx context.Context, params *AddParams) (*SaveChanges, error) {
	path := "/proxy"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.ProxyIP != nil {
			opts.JSON["proxy_ip"] = *params.ProxyIP
		}
		if params.ProxyPass != nil {
			opts.JSON["proxy_pass"] = *params.ProxyPass
		}
		if params.ProxyPort != nil {
			opts.JSON["proxy_port"] = *params.ProxyPort
		}
		if params.ProxyRow != nil {
			opts.JSON["proxy_row"] = *params.ProxyRow
		}
		if params.ProxyUser != nil {
			opts.JSON["proxy_user"] = *params.ProxyUser
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

// Delete Delete Proxy
func (s *ProxyService) Delete(ctx context.Context, params *DeleteParams) (*SaveChanges, error) {
	path := "/proxy"
	opts := RequestOptions{}
	opts.JSON = make(map[string]interface{})
	if params != nil {
		if params.DeleteAll != nil {
			opts.JSON["delete_all"] = *params.DeleteAll
		}
		if params.ProxyID != nil {
			opts.JSON["proxy_id"] = *params.ProxyID
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

// Get Get Proxy
func (s *ProxyService) Get(ctx context.Context) (*GetResponse, error) {
	path := "/proxy"
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

// Ensure imports are used.
var _ = fmt.Sprintf
var _ = strings.Replace
var _ = json.Marshal

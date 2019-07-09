package _36kr

type JsonOverviewResponse struct {
	Code        int64 `json:"code"`
	Timestamp   int64 `json:"timestamp"`
	TimestampRt int64 `json:"timestamp_rt"`
	Data        Data  `json:"data"`
}

type Data struct {
	TotalCount int64  `json:"total_count"`
	Page       int64  `json:"page"`
	PageSize   int64  `json:"page_size"`
	Items      []Item `json:"items"`
}

type Item struct {
	ID               int64        `json:"id"`
	Title            string       `json:"title"`
	ProjectID        interface{}  `json:"project_id"`
	ViewsCount       string       `json:"views_count"`
	MobileViewsCount string       `json:"mobile_views_count"`
	AppViewsCount    string       `json:"app_views_count"`
	MonographicID    interface{}  `json:"monographic_id"`
	DomainID         string       `json:"domain_id"`
	GoodsID          string       `json:"goods_id"`
	IsTovc           string       `json:"is_tovc"`
	IsFree           string       `json:"is_free"`
	EntityFlag       string       `json:"entity_flag"`
	Cover            string       `json:"cover"`
	TemplateInfo     TemplateInfo `json:"template_info"`
	PublishedAt      string       `json:"published_at"`
	ColumnName       ColumnName   `json:"column_name"`
	ColumnID         interface{}  `json:"column_id"`
	UserID           string       `json:"user_id"`
	UserName         string       `json:"user_name"`
	UserAvatarURL    string       `json:"user_avatar_url"`
	Highlight        Highlight    `json:"highlight"`
	Type             Type         `json:"_type"`
	Score            interface{}  `json:"_score"`
}

type Highlight struct {
	Content      []string `json:"content"`
	ContentLight []string `json:"content_light"`
	Title        []string `json:"title"`
}

type TemplateInfo struct {
	TemplateType        TemplateType `json:"template_type"`
	TemplateTitle       string       `json:"template_title"`
	TemplateTitleIsSame *bool        `json:"template_title_isSame,omitempty"`
	TemplateCover       []string     `json:"template_cover"`
}

type ColumnName string

const (
	Empty  ColumnName = ""
	付费文章专栏 ColumnName = "付费文章专栏"
	内容     ColumnName = "内容"
)

type TemplateType string

const (
	SmallImage TemplateType = "small_image"
)

type Type string

const (
	Post Type = "post"
)

type ID struct {
	Integer *int64
	String  *string
}

type Payload struct {
	Page       int    `json:"page"`
	PerPage    int    `json:"per_page"`
	Sort       string `json:"sort"`
	EntityType string `json:"entity_type"`
	Keyword    string `json:"keyword"`
}

type JsonContentResponse struct {
	ArticleDetail ArticleDetail `json:"articleDetail"`
	Channel       []Channel     `json:"channel"`
	UserInfo      interface{}   `json:"userInfo"`
}

type ArticleDetail struct {
	IsPureReading         bool                  `json:"isPureReading"`
	RenderPureReading     bool                  `json:"renderPureReading"`
	ShowPurereading       bool                  `json:"showPurereading"`
	ScrollToCommentBox    bool                  `json:"scrollToCommentBox"`
	LoginModalShow        bool                  `json:"loginModalShow"`
	CompanyFormMsg        string                `json:"companyFormMsg"`
	ShowtextFormTrue      bool                  `json:"showtextFormTrue"`
	ArticleDetailData     ArticleDetailData     `json:"articleDetailData"`
	AuthorLatestData      AuthorLatestData      `json:"authorLatestData"`
	NextArticleDetailData NextArticleDetailData `json:"nextArticleDetailData"`
	PostCreate            string                `json:"post_create"`
	OrganArticleData      OrganArticleData      `json:"organArticleData"`
	Enterprise            interface{}           `json:"enterprise"`
}

type ArticleDetailData struct {
	Code int64                 `json:"code"`
	Msg  string                `json:"msg"`
	Data ArticleDetailDataData `json:"data"`
}

type ArticleDetailDataData struct {
	ID                      int64         `json:"id"`
	ProjectID               int64         `json:"project_id"`
	GoodsID                 int64         `json:"goods_id"`
	DomainID                int64         `json:"domain_id"`
	ColumnID                int64         `json:"column_id"`
	MonographicID           int64         `json:"monographic_id"`
	RelatedCompanyID        int64         `json:"related_company_id"`
	RelatedCompanyType      string        `json:"related_company_type"`
	RelatedCompanyName      string        `json:"related_company_name"`
	CloseComment            int64         `json:"close_comment"`
	State                   State         `json:"state"`
	Title                   string        `json:"title"`
	CatchTitle              string        `json:"catch_title"`
	Summary                 string        `json:"summary"`
	Content                 string        `json:"content"`
	Cover                   string        `json:"cover"`
	SourceType              string        `json:"source_type"`
	SourceUrls              string        `json:"source_urls"`
	RelatedPostIDS          string        `json:"related_post_ids"`
	ExtractionTags          string        `json:"extraction_tags"`
	Extra                   interface{}   `json:"extra"`
	UserID                  int64         `json:"user_id"`
	PublishedAt             string        `json:"published_at"`
	CreatedAt               string        `json:"created_at"`
	UpdatedAt               string        `json:"updated_at"`
	Counters                Counters      `json:"counters"`
	RelatedCompanyCounters  interface{}   `json:"related_company_counters"`
	RelatedPosts            interface{}   `json:"related_posts"`
	IsFree                  int64         `json:"is_free"`
	HasRightsGoods          bool          `json:"has_rights_goods"`
	IsTovc                  int64         `json:"is_tovc"`
	ImageSource             []interface{} `json:"image_source"`
	CompanyInfo             interface{}   `json:"company_info"`
	CompanyContactInfo      interface{}   `json:"company_contact_info"`
	CompanyFundInfo         interface{}   `json:"company_fund_info"`
	ShareData               ShareData     `json:"share_data"`
	TitleMobile             interface{}   `json:"title_mobile"`
	CoverMobile             interface{}   `json:"cover_mobile"`
	Audios                  []interface{} `json:"audios"`
	Column                  Column        `json:"column"`
	DBCounters              []DBCounter   `json:"db_counters"`
	Monographic             Monographic   `json:"monographic"`
	User                    PurpleUser    `json:"user"`
	Motifs                  []Motif       `json:"motifs"`
	MotifBanner             interface{}   `json:"motif_banner"`
	HasBuyGoods             bool          `json:"has_buy_goods"`
	HasRightsResearchReport bool          `json:"has_rights_research_report"`
	HasRightCompany         bool          `json:"has_right_company"`
}

type Column struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	BgColor      string `json:"bg_color"`
	Type         string `json:"type"`
}

type Counters struct {
	ViewCount int64 `json:"view_count"`
	PV        int64 `json:"pv"`
	PVMobile  int64 `json:"pv_mobile"`
	PVApp     int64 `json:"pv_app"`
	Favorite  int64 `json:"favorite"`
	Comment   int64 `json:"comment"`
	Like      int64 `json:"like"`
}

type DBCounter struct {
	ID         int64  `json:"id"`
	EntityType string `json:"entity_type"`
	EntityID   int64  `json:"entity_id"`
	CountType  string `json:"count_type"`
	Value      int64  `json:"value"`
}

type Monographic struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	Cover    string `json:"cover"`
	CoverWeb string `json:"cover_web"`
}

type Motif struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	SourceName  string `json:"source_name"`
}

type ShareData struct {
	Default Default `json:"default"`
	Weibo   Weibo   `json:"weibo"`
	Youdao  Youdao  `json:"youdao"`
}

type Default struct {
	Title       string `json:"title"`
	Cover       string `json:"cover"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type Weibo struct {
	Title string `json:"title"`
	Cover string `json:"cover"`
}

type Youdao struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type PurpleUser struct {
	ID               int64        `json:"id"`
	Name             string       `json:"name"`
	AvatarURL        string       `json:"avatar_url"`
	TovcAvatarURL    interface{}  `json:"tovc_avatar_url"`
	Introduction     string       `json:"introduction"`
	TovcBriefIntro   string       `json:"tovc_brief_intro"`
	TovcIntro        string       `json:"tovc_intro"`
	TovcTitle        string       `json:"tovc_title"`
	TovcLevel        interface{}  `json:"tovc_level"`
	RoleID           int64        `json:"role_id"`
	Title            string       `json:"title"`
	DepartmentBelong string       `json:"department_belong"`
	QrCodeURL        string       `json:"qr_code_url"`
	Roles            []PurpleRole `json:"roles"`
	Role             PurpleRole   `json:"role"`
}

type PurpleRole struct {
	ID                 int64       `json:"id"`
	Name               string      `json:"name"`
	Description        string      `json:"description"`
	UserIDServer       int64       `json:"user_id_server"`
	AvailableDomainIDS string      `json:"available_domain_ids"`
	Server             interface{} `json:"server"`
}

type AuthorLatestData struct {
	Code int64                `json:"code"`
	Msg  string               `json:"msg"`
	Data AuthorLatestDataData `json:"data"`
}

type AuthorLatestDataData struct {
	TotalCount int64  `json:"total_count"`
	Page       int64  `json:"page"`
	PageSize   int64  `json:"page_size"`
	Items      []Item `json:"items"`
}

type NextArticleDetailData struct {
	Code int64                     `json:"code"`
	Msg  string                    `json:"msg"`
	Data NextArticleDetailDataData `json:"data"`
}

type NextArticleDetailDataData struct {
	ID                      int64         `json:"id"`
	ProjectID               int64         `json:"project_id"`
	GoodsID                 int64         `json:"goods_id"`
	DomainID                int64         `json:"domain_id"`
	ColumnID                int64         `json:"column_id"`
	MonographicID           interface{}   `json:"monographic_id"`
	RelatedCompanyID        int64         `json:"related_company_id"`
	RelatedCompanyType      string        `json:"related_company_type"`
	RelatedCompanyName      string        `json:"related_company_name"`
	CloseComment            int64         `json:"close_comment"`
	State                   State         `json:"state"`
	Title                   string        `json:"title"`
	CatchTitle              string        `json:"catch_title"`
	Summary                 string        `json:"summary"`
	Content                 string        `json:"content"`
	Cover                   string        `json:"cover"`
	SourceType              string        `json:"source_type"`
	SourceUrls              string        `json:"source_urls"`
	RelatedPostIDS          string        `json:"related_post_ids"`
	ExtractionTags          string        `json:"extraction_tags"`
	Extra                   interface{}   `json:"extra"`
	UserID                  int64         `json:"user_id"`
	PublishedAt             string        `json:"published_at"`
	CreatedAt               string        `json:"created_at"`
	UpdatedAt               string        `json:"updated_at"`
	Counters                Counters      `json:"counters"`
	RelatedCompanyCounters  interface{}   `json:"related_company_counters"`
	RelatedPosts            interface{}   `json:"related_posts"`
	IsFree                  int64         `json:"is_free"`
	HasRightsGoods          bool          `json:"has_rights_goods"`
	IsTovc                  int64         `json:"is_tovc"`
	ImageSource             string        `json:"image_source"`
	CompanyInfo             interface{}   `json:"company_info"`
	CompanyContactInfo      interface{}   `json:"company_contact_info"`
	CompanyFundInfo         interface{}   `json:"company_fund_info"`
	TitleMobile             interface{}   `json:"title_mobile"`
	CoverMobile             interface{}   `json:"cover_mobile"`
	Audios                  []interface{} `json:"audios"`
	Column                  Column        `json:"column"`
	DBCounters              []interface{} `json:"db_counters"`
	User                    FluffyUser    `json:"user"`
	HasRightsResearchReport bool          `json:"has_rights_research_report"`
}

type FluffyUser struct {
	ID               int64       `json:"id"`
	Name             string      `json:"name"`
	AvatarURL        string      `json:"avatar_url"`
	TovcAvatarURL    interface{} `json:"tovc_avatar_url"`
	Introduction     string      `json:"introduction"`
	TovcBriefIntro   string      `json:"tovc_brief_intro"`
	TovcIntro        string      `json:"tovc_intro"`
	TovcTitle        string      `json:"tovc_title"`
	TovcLevel        interface{} `json:"tovc_level"`
	RoleID           int64       `json:"role_id"`
	Title            string      `json:"title"`
	DepartmentBelong string      `json:"department_belong"`
	QrCodeURL        string      `json:"qr_code_url"`
	Role             FluffyRole  `json:"role"`
}

type FluffyRole struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	UserIDServer int64  `json:"user_id_server"`
}

type OrganArticleData struct {
	Code int64                `json:"code"`
	Data OrganArticleDataData `json:"data"`
	Msg  string               `json:"msg"`
}

type OrganArticleDataData struct {
	OrganizationList []interface{} `json:"organizationList"`
}

type Channel struct {
	ID           int64       `json:"id"`
	Type         Type        `json:"type"`
	ProjectID    int64       `json:"project_id"`
	Key          string      `json:"key"`
	State        State       `json:"state"`
	IsSystem     int64       `json:"is_system"`
	API          interface{} `json:"api"`
	Name         string      `json:"name"`
	Mark         Mark        `json:"mark"`
	Description  Description `json:"description"`
	AdInfo       string      `json:"ad_info"`
	OrderNum     int64       `json:"order_num"`
	UserID       int64       `json:"user_id"`
	UserIDEdited int64       `json:"user_id_edited"`
	PublishedAt  *string     `json:"published_at"`
	CreatedAt    *string     `json:"created_at"`
	UpdatedAt    string      `json:"updated_at"`
}

type State string
const (
	Published State = "published"
)

type Description string
const (
	品牌市场部频道 Description = "品牌市场部频道"
)

type Mark string
const (
	Hot Mark = "hot"
	None Mark = "none"
)

const (
	Web Type = "web"
)

type ProjectID struct {
	Integer *int64
	String  *string
}


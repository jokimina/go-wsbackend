package _36kr

type JsonResponse struct {
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
	ProjectID        *ID          `json:"project_id"`
	ViewsCount       string       `json:"views_count"`
	MobileViewsCount string       `json:"mobile_views_count"`
	AppViewsCount    string       `json:"app_views_count"`
	MonographicID    *ID          `json:"monographic_id"`
	DomainID         string       `json:"domain_id"`
	GoodsID          string       `json:"goods_id"`
	IsTovc           string       `json:"is_tovc"`
	IsFree           string       `json:"is_free"`
	EntityFlag       string       `json:"entity_flag"`
	Cover            string       `json:"cover"`
	TemplateInfo     TemplateInfo `json:"template_info"`
	PublishedAt      string       `json:"published_at"`
	ColumnName       ColumnName   `json:"column_name"`
	ColumnID         *ID          `json:"column_id"`
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
	Empty ColumnName = ""
	付费文章专栏 ColumnName = "付费文章专栏"
	内容 ColumnName = "内容"
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
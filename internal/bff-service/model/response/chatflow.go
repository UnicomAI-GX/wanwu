package response

type GetDraftIntelligenceListResponse struct {
	Data DraftIntelligenceListData `json:"data"`
	Code int32                     `json:"code"`
	Msg  string                    `json:"msg"`
}
type DraftIntelligenceListData struct {
	Intelligences []*IntelligenceData `json:"intelligences"`
	Total         int32               `json:"total"`
	HasMore       bool                `json:"has_more"`
	NextCursorID  string              `json:"next_cursor_id"`
}

type IntelligenceData struct {
	BasicInfo      IntelligenceBasicInfo      `json:"basic_info"`
	Type           int64                      `json:"type"`
	PublishInfo    IntelligencePublishInfo    `json:"publish_info"`
	PermissionInfo IntelligencePermissionInfo `json:"permission_info"`
	OwnerInfo      *CozeUser                  `json:"owner_info"`
	FavoriteInfo   *FavoriteInfo              `json:"favorite_info"`
}

type IntelligenceBasicInfo struct {
	ID             int64   `json:"id,string"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	IconURI        string  `json:"icon_uri"`
	IconURL        string  `json:"icon_url"`
	SpaceID        int64   `json:"space_id,string"`
	OwnerID        int64   `json:"owner_id,string"`
	CreateTime     int64   `json:"create_time,string"`
	UpdateTime     int64   `json:"update_time,string"`
	Status         int64   `json:"status"`
	PublishTime    int64   `json:"publish_time,string"`
	EnterpriseID   *string `json:"enterprise_id,omitempty"`
	OrganizationID *int64  `json:"organization_id,omitempty"`
}

type IntelligencePublishInfo struct {
	PublishTime  string           `json:"publish_time"`
	HasPublished bool             `json:"has_published"`
	Connectors   []*ConnectorInfo `json:"connectors"`
}

type ConnectorInfo struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Icon            string `json:"icon"`
	ConnectorStatus int64  `json:"connector_status"`
	ShareLink       string `json:"share_link,omitempty"`
}

type IntelligencePermissionInfo struct {
	InCollaboration bool `json:"in_collaboration"`
	CanDelete       bool `json:"can_delete"`
	CanView         bool `json:"can_view"`
}

type CozeUser struct {
	UserID         int64      `json:"user_id,string"`
	Nickname       string     `json:"nickname"`
	AvatarURL      string     `json:"avatar_url"`
	UserUniqueName string     `json:"user_unique_name"`
	UserLabel      *UserLabel `json:"user_label"`
}

type UserLabel struct {
	LabelID   string `json:"label_id"`
	LabelName string `json:"label_name"`
	IconURI   string `json:"icon_uri"`
	IconURL   string `json:"icon_url"`
	JumpLink  string `json:"jump_link"`
}

type FavoriteInfo struct {
	IsFav   bool   `json:"is_fav"`
	FavTime string `json:"fav_time"`
}

type CozeCreateProjectConversationDefResponse struct {
	UniqueID string `json:"unique_id"`
	SpaceID  string `json:"space_id"`
	Code     int64  `json:"code"`
	Msg      string `json:"msg"`
}

type GetDraftIntelligenceInfoResponse struct {
	Data *GetDraftIntelligenceInfoData `json:"data"`
	Code int32                         `json:"code"`
	Msg  string                        `json:"msg"`
}

type GetDraftIntelligenceInfoData struct {
	IntelligenceType int64                    `json:"intelligence_type"`
	BasicInfo        *IntelligenceBasicInfo   `json:"basic_info"`
	PublishInfo      *IntelligencePublishInfo `json:"publish_info,omitempty"`
	OwnerInfo        *CozeUser                `json:"owner_info,omitempty"`
}

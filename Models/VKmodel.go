package Models

type Response struct {
	Count int    `json:"count"`
	Items []Item `json:"items"`
}

type Item struct {
	CanDelete     int          `json:"can_delete"`
	CanPin        int          `json:"can_pin"`
	IsPinned      int          `json:"is_pinned"`
	Comments      Comments     `json:"comments"`
	MarkedAsAds   int          `json:"marked_as_ads"`
	ShortTextRate float64      `json:"short_text_rate"`
	Hash          string       `json:"hash"`
	Type          string       `json:"type"`
	Attachments   []Attachment `json:"attachments"`
	Date          int64        `json:"date"`
	FromID        int64        `json:"from_id"`
	ID            int64        `json:"id"`
	IsFavorite    bool         `json:"is_favorite"`
	OwnerID       int64        `json:"owner_id"`
	PostSource    PostSource   `json:"post_source"`
	PostType      string       `json:"post_type"`
	Reposts       Reposts      `json:"reposts"`
	Text          string       `json:"text"`
	Views         Views        `json:"views"`
}

type Comments struct {
	CanPost       int  `json:"can_post"`
	CanClose      int  `json:"can_close"`
	Count         int  `json:"count"`
	GroupsCanPost bool `json:"groups_can_post"`
}

type Attachment struct {
	Type  string `json:"type"`
	Photo Photo  `json:"photo"`
	Link  Link   `json:"link"`
	Video Video  `json:"video"`
	Audio Audio  `json:"audio"`
}

type Photo struct {
	AlbumID      int64  `json:"album_id"`
	Date         int64  `json:"date"`
	ID           int64  `json:"id"`
	OwnerID      int64  `json:"owner_id"`
	AccessKey    string `json:"access_key"`
	Sizes        []Size `json:"sizes"`
	Text         string `json:"text"`
	UserID       int64  `json:"user_id"`
	WebViewToken string `json:"web_view_token"`
	HasTags      bool   `json:"has_tags"`
}

type Size struct {
	Height int    `json:"height"`
	Type   string `json:"type"`
	Width  int    `json:"width"`
	URL    string `json:"url"`
}

type Link struct {
	URL         string `json:"url"`
	Description string `json:"description"`
	IsFavorite  bool   `json:"is_favorite"`
	Title       string `json:"title"`
	Target      string `json:"target"`
}

type Video struct {
	ResponseType   string  `json:"response_type"`
	AccessKey      string  `json:"access_key"`
	CanComment     int     `json:"can_comment"`
	CanEdit        int     `json:"can_edit"`
	CanDelete      int     `json:"can_delete"`
	CanLike        int     `json:"can_like"`
	CanRepost      int     `json:"can_repost"`
	CanAddToFaves  int     `json:"can_add_to_faves"`
	CanAdd         int     `json:"can_add"`
	CanAttachLink  int     `json:"can_attach_link"`
	CanEditPrivacy int     `json:"can_edit_privacy"`
	Comments       int     `json:"comments"`
	Date           int64   `json:"date"`
	Description    string  `json:"description"`
	Duration       int     `json:"duration"`
	Image          []Image `json:"image"`
	ID             int64   `json:"id"`
	OwnerID        int64   `json:"owner_id"`
	IsAuthor       bool    `json:"is_author"`
	Title          string  `json:"title"`
	IsFavorite     bool    `json:"is_favorite"`
	TrackCode      string  `json:"track_code"`
	Type           string  `json:"type"`
	Views          int     `json:"views"`
	Platform       string  `json:"platform"`
	CanDislike     int     `json:"can_dislike"`
}

type Image struct {
	URL         string `json:"url"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	WithPadding int    `json:"with_padding"`
}

type Audio struct {
	Artist              string       `json:"artist"`
	ID                  int64        `json:"id"`
	OwnerID             int64        `json:"owner_id"`
	Title               string       `json:"title"`
	Duration            int          `json:"duration"`
	IsExplicit          bool         `json:"is_explicit"`
	IsFocusTrack        bool         `json:"is_focus_track"`
	TrackCode           string       `json:"track_code"`
	URL                 string       `json:"url"`
	Date                int64        `json:"date"`
	LyricsID            int          `json:"lyrics_id"`
	MainArtists         []MainArtist `json:"main_artists"`
	ShortVideosAllowed  bool         `json:"short_videos_allowed"`
	StoriesAllowed      bool         `json:"stories_allowed"`
	StoriesCoverAllowed bool         `json:"stories_cover_allowed"`
}

type MainArtist struct {
	Name       string `json:"name"`
	Domain     string `json:"domain"`
	ID         string `json:"id"`
	IsFollowed bool   `json:"is_followed"`
	CanFollow  bool   `json:"can_follow"`
}

type PostSource struct {
	Type string `json:"type"`
}

type Reposts struct {
	Count        int `json:"count"`
	WallCount    int `json:"wall_count"`
	MailCount    int `json:"mail_count"`
	UserReposted int `json:"user_reposted"`
}

type Views struct {
	Count int `json:"count"`
}

type VKResponse struct {
	Response Response `json:"response"`
}

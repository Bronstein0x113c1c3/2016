package instructure

import "time"

type Announcement_Info_Instructure struct {
	ID                      int         `json:"id"`
	Title                   string      `json:"title"`
	LastReplyAt             time.Time   `json:"last_reply_at"`
	CreatedAt               time.Time   `json:"created_at"`
	DelayedPostAt           interface{} `json:"delayed_post_at"`
	PostedAt                time.Time   `json:"posted_at"`
	AssignmentID            interface{} `json:"assignment_id"`
	RootTopicID             interface{} `json:"root_topic_id"`
	Position                int         `json:"position"`
	PodcastHasStudentPosts  bool        `json:"podcast_has_student_posts"`
	DiscussionType          string      `json:"discussion_type"`
	LockAt                  interface{} `json:"lock_at"`
	AllowRating             bool        `json:"allow_rating"`
	OnlyGradersCanRate      bool        `json:"only_graders_can_rate"`
	SortByRating            bool        `json:"sort_by_rating"`
	IsSectionSpecific       bool        `json:"is_section_specific"`
	AnonymousState          interface{} `json:"anonymous_state"`
	UserName                string      `json:"user_name"`
	DiscussionSubentryCount int         `json:"discussion_subentry_count"`
	Permissions             struct {
		Attach bool `json:"attach"`
		Update bool `json:"update"`
		Reply  bool `json:"reply"`
		Delete bool `json:"delete"`
	} `json:"permissions"`
	RequireInitialPost interface{}   `json:"require_initial_post"`
	UserCanSeePosts    bool          `json:"user_can_see_posts"`
	PodcastURL         interface{}   `json:"podcast_url"`
	ReadState          string        `json:"read_state"`
	UnreadCount        int           `json:"unread_count"`
	Subscribed         bool          `json:"subscribed"`
	Attachments        []interface{} `json:"attachments"`
	Published          bool          `json:"published"`
	CanUnpublish       bool          `json:"can_unpublish"`
	Locked             bool          `json:"locked"`
	CanLock            bool          `json:"can_lock"`
	CommentsDisabled   bool          `json:"comments_disabled"`
	Author             struct {
		ID             int         `json:"id"`
		AnonymousID    string      `json:"anonymous_id"`
		DisplayName    string      `json:"display_name"`
		AvatarImageURL string      `json:"avatar_image_url"`
		HTMLURL        string      `json:"html_url"`
		Pronouns       interface{} `json:"pronouns"`
	} `json:"author"`
	HTMLURL            string        `json:"html_url"`
	URL                string        `json:"url"`
	Pinned             bool          `json:"pinned"`
	GroupCategoryID    interface{}   `json:"group_category_id"`
	CanGroup           bool          `json:"can_group"`
	TopicChildren      []interface{} `json:"topic_children"`
	GroupTopicChildren []interface{} `json:"group_topic_children"`
	ContextCode        string        `json:"context_code"`
	LockedForUser      bool          `json:"locked_for_user"`
	LockInfo           struct {
		CanView     bool   `json:"can_view"`
		AssetString string `json:"asset_string"`
	} `json:"lock_info"`
	LockExplanation  string      `json:"lock_explanation"`
	Message          string      `json:"message"`
	SubscriptionHold string      `json:"subscription_hold"`
	TodoDate         interface{} `json:"todo_date"`
	IsAnnouncement   bool        `json:"is_announcement"`
}

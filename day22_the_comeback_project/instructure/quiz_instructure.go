package instructure

import "time"

type Quiz_Info_Instructure struct {
	ID                      int         `json:"id"`
	Title                   string      `json:"title"`
	HTMLURL                 string      `json:"html_url"`
	MobileURL               string      `json:"mobile_url"`
	Description             string      `json:"description"`
	QuizType                string      `json:"quiz_type"`
	TimeLimit               interface{} `json:"time_limit"`
	TimerAutosubmitDisabled bool        `json:"timer_autosubmit_disabled"`
	ShuffleAnswers          bool        `json:"shuffle_answers"`
	ShowCorrectAnswers      bool        `json:"show_correct_answers"`
	ScoringPolicy           string      `json:"scoring_policy"`
	AllowedAttempts         int         `json:"allowed_attempts"`
	OneQuestionAtATime      bool        `json:"one_question_at_a_time"`
	QuestionCount           int         `json:"question_count"`
	PointsPossible          float64     `json:"points_possible"`
	CantGoBack              bool        `json:"cant_go_back"`
	IPFilter                interface{} `json:"ip_filter"`
	DueAt                   interface{} `json:"due_at"`
	LockAt                  time.Time   `json:"lock_at"`
	UnlockAt                interface{} `json:"unlock_at"`
	Published               bool        `json:"published"`
	LockedForUser           bool        `json:"locked_for_user"`
	LockInfo                struct {
		LockAt      time.Time `json:"lock_at"`
		CanView     bool      `json:"can_view"`
		AssetString string    `json:"asset_string"`
	} `json:"lock_info,omitempty"`
	LockExplanation      string      `json:"lock_explanation,omitempty"`
	HideResults          interface{} `json:"hide_results"`
	ShowCorrectAnswersAt time.Time   `json:"show_correct_answers_at"`
	HideCorrectAnswersAt interface{} `json:"hide_correct_answers_at"`
	AllDates             []struct {
		DueAt    interface{} `json:"due_at"`
		UnlockAt interface{} `json:"unlock_at"`
		LockAt   time.Time   `json:"lock_at"`
		Base     bool        `json:"base"`
	} `json:"all_dates"`
	CanUpdate                        bool        `json:"can_update"`
	RequireLockdownBrowser           bool        `json:"require_lockdown_browser"`
	RequireLockdownBrowserForResults bool        `json:"require_lockdown_browser_for_results"`
	RequireLockdownBrowserMonitor    bool        `json:"require_lockdown_browser_monitor"`
	LockdownBrowserMonitorData       interface{} `json:"lockdown_browser_monitor_data"`
	Permissions                      struct {
		Manage           bool `json:"manage"`
		Read             bool `json:"read"`
		Create           bool `json:"create"`
		Update           bool `json:"update"`
		Submit           bool `json:"submit"`
		Preview          bool `json:"preview"`
		Delete           bool `json:"delete"`
		ReadStatistics   bool `json:"read_statistics"`
		Grade            bool `json:"grade"`
		ReviewGrades     bool `json:"review_grades"`
		ViewAnswerAudits bool `json:"view_answer_audits"`
	} `json:"permissions"`
	QuizReportsURL                string      `json:"quiz_reports_url"`
	QuizStatisticsURL             string      `json:"quiz_statistics_url"`
	ImportantDates                bool        `json:"important_dates"`
	QuizSubmissionVersionsHTMLURL string      `json:"quiz_submission_versions_html_url"`
	AssignmentID                  int         `json:"assignment_id"`
	OneTimeResults                bool        `json:"one_time_results"`
	AssignmentGroupID             int         `json:"assignment_group_id"`
	ShowCorrectAnswersLastAttempt bool        `json:"show_correct_answers_last_attempt"`
	VersionNumber                 int         `json:"version_number"`
	HasAccessCode                 bool        `json:"has_access_code"`
	PostToSis                     bool        `json:"post_to_sis"`
	MigrationID                   interface{} `json:"migration_id"`
	InPacedCourse                 bool        `json:"in_paced_course"`
}

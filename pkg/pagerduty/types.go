package pagerduty

type APIObject struct {
	ID      string `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	Summary string `json:"summary,omitempty"`
	Self    string `json:"self,omitempty"`
	HTMLURL string `json:"html_url,omitempty"`
}

type ContactMethod struct {
	*APIObject
}

type NotificationRule struct {
	*APIObject
}

type Team struct {
	*APIObject
}

type User struct {
	*APIObject
	Name              string `json:"name"`
	Email             string `json:"email"`
	Timezone          string `json:"timezone,omitempty"`
	Color             string `json:"color,omitempty"`
	Role              string `json:"role,omitempty"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	Description       string `json:"description,omitempty"`
	InvitationSent    bool
	ContactMethods    []ContactMethod    `json:"contact_methods"`
	NotificationRules []NotificationRule `json:"notification_rules"`
	JobTitle          string             `json:"job_title,omitempty"`
	Teams             []Team
}

type OnCalls struct {
	OnCalls []*OnCall `json:"oncalls"`
}

type OnCall struct {
	EscalationLevel  int       `json:"escalation_level"`
	EscalationPolicy APIObject `json:"escalation_policy"`
	Schedule         APIObject `json:"schedule"`
	User             *User     `json:"user"`
	Start            string    `json:"start"`
	End              string    `json:"end"`
}

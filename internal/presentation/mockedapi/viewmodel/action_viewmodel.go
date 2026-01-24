package viewmodel

type ActionViewModel struct {
	Priority     int                `json:"priority"`
	Condition    *ConditionViewModel `json:"condition"`
	Type         string             `json:"type"`
	FieldToSet   string             `json:"fieldToSet"`
	CookieName   string             `json:"cookieName"`
	CookieValue  string             `json:"cookieValue"`
	RemoveCookie bool               `json:"removeCookie"`
	Result       string             `json:"result"`
}

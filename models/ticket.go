package models

type Ticket struct {
	Id                 string         `json:"_id"`
	Url                string         `json:"url"`
	ExternalId         string         `json:"external_id"`
	CreatedAt          string         `json:"created_at"`
	Type               string         `json:"type"`
    Subject            string         `json:"subject"`
	Description        string         `json:"description"`
	Priority           string         `json:"priority"`
	Status             string         `json:"status"`
	SubmitterId        int            `json:"submitter_id"`
	AssigneeId         int            `json:"assignee_id"`
	OrganizationId     int 	          `json:"organization_id"`
	Tags               []string       `json:"tags"`
	HasIncidents       bool           `json:"has_incidents"`
	DueAt              string         `json:"due_at"`
	Via                string         `json:"via"`
}
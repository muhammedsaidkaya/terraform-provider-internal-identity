package provider

type User struct {
	UserId         string `json:"userId"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	SlackUsername  string `json:"slackUsername"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	JiraUsername   string `json:"jiraUsername"`
	GithubUsername string `json:"githubUsername"`
}

package models

//Gitlab config
type GitlabHook struct {
	ObjectKind string `json:"object_kind"`
	EventName  string `json:"event_name"`
	Ref        string`json:"ref"`
	GitlabProject *GitlabProject `json:"project"`
}

type GitlabProject struct {
	Name string `json:"name"`
	Description string `json:"description"`
	WebUrl string `json:"web_url"`
	AvatarUrl string `json:"avatar_url"`
	GitSshUrl string `json:"git_ssh_url"`
	GitHttpUrl string `json:"git_http_url"`
}
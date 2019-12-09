package github

//CreateRepoRequest models request to Github API to create a repository
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"home_page"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

//CreateRepoResponse models response of Github API to create a repository
type CreateRepoResponse struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name"`
	Owner       RepoOwner       `json:"owner"`
	Permissions RepoPermissions `json:"permissions"`
}

//RepoOwner models Github repository owner
type RepoOwner struct {
	ID      int64  `json:"id"`
	URL     string `json:"url"`
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

//RepoPermissions models Github repository permissions
type RepoPermissions struct {
	IsAdmin bool `json:"admin"`
	HasPush bool `json:"push"`
	HasPull bool `json:"pull"`
}

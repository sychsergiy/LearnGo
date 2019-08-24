package github

type CreateIssueRequestData struct {
	Title     string
	Body      string
	Milestone string
	Labels    string
	Assignees []string
}

func CreateIssue(reqBody CreateIssueRequestData) {

}

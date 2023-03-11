package jira

import (
	"context"
	"fmt"
	"io"

	"github.com/andygrunwald/go-jira"
	"github.com/ivaaaan/mira/task"
)

// I guess it's default for all projects
// Perhaps should be moved to the config
const epicNameField = "customfield_10011"

type jiraProvider struct {
	jiraClient *jira.Client
	projectKey string
}

func NewJiraProvider(URL, username, password, projectKey string) (*jiraProvider, error) {
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}

	client, err := jira.NewClient(tp.Client(), URL)
	if err != nil {
		return nil, fmt.Errorf("cannot init jira client: %v", err)
	}

	return &jiraProvider{
		jiraClient: client,
		projectKey: projectKey,
	}, nil
}

func (p *jiraProvider) createIssues(ctx context.Context, t *task.Task, issueType, parent string) error {

	fields := &jira.IssueFields{
		Summary:     t.Title,
		Description: t.Description,
		Type: jira.IssueType{
			Name: issueType,
		},
		Unknowns: make(map[string]interface{}),
		Project: jira.Project{
			Key: p.projectKey,
		},
	}

	if issueType == "Epic" {
		fields.Unknowns[epicNameField] = t.Title
	}

	if parent != "" {
		fields.Parent = &jira.Parent{
			ID: parent,
		}
	}

	issue, rsp, err := p.jiraClient.Issue.Create(&jira.Issue{
		Fields: fields,
	})

	if err != nil {
		b, err := io.ReadAll(rsp.Body)
		if err != nil {
			return fmt.Errorf("cannot read response body: %v", err)
		}

		return fmt.Errorf("cannot create jira issue: %v, response body: %v", err, string(b))
	}

	for _, c := range t.Children {
		err := p.createIssues(ctx, c, "Task", issue.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *jiraProvider) Push(ctx context.Context, t *task.Task) error {
	return p.createIssues(ctx, t, "Epic", "")
}

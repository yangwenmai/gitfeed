package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/modood/wpm"
	"github.com/tidwall/gjson"
)

var (
	width  = "32px"
	height = "32px"
	size   = ""
	debug  = false
)

// NewsFeed newsfeed
type NewsFeed struct {
	ID        string  `json:"id"`
	Type      string  `json:"type"`
	Actor     Actor   `json:"actor"`
	Repo      Repo    `json:"repo"`
	Payload   Payload `json:"payload"`
	Public    bool    `json:"public"`
	CreatedAt string  `json:"created_at"`
}

// Actor actor
type Actor struct {
	ID           int32  `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

// Repo repo
type Repo struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Payload payload
type Payload struct {
	Action       string      `json:"action"`
	Ref          string      `json:"ref"`
	RefType      string      `json:"ref_type"`
	MasterBranch string      `json:"master_branch"`
	Description  string      `json:"description"`
	PusherType   string      `json:"pusher_type"`
	Size         int32       `json:"size"`
	Forkee       Forkee      `json:"forkee"`
	PullRequest  PullRequest `json:"pull_request"`
	Comment      Comment     `json:"comment"`
	Issue        Issue       `json:"issue"`
	Member       Member      `json:"member"`
}

// Forkee forkee
type Forkee struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Owner       Owner  `json:"owner"`
	HTMLURL     string `json:"html_url"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Fork        bool   `json:"fork"`
	ForksURL    string `json:"forks_url"`
}

// PullRequest pr
type PullRequest struct {
	Number   string `json:"number"`
	State    string `json:"state"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	HTMLURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`
	Body     string `json:"body"`
}

// Comment comment
type Comment struct {
	ID                int32  `json:"id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	IssueURL          string `json:"issue_url"`
	User              User   `json:"user"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	AuthorAssociation string `json:"author_association"`
	Body              string `json:"body"`
}

// User user
type User struct {
	ID                int32  `json:"id"`
	Login             string `json:"login"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	AvatarURL         string `json:"avatar_url"`
	Type              string `json:"type"`
	SiteAdmin         string `json:"site_admin"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
}

// Issue issue
type Issue struct {
	ID                int32       `json:"id"`
	Number            string      `json:"number"`
	Title             string      `json:"title"`
	URL               string      `json:"url"`
	RepositoryURL     string      `json:"repository_url"`
	LabelsURL         string      `json:"labels_url"`
	CommentsURL       string      `json:"comments_url"`
	EventsURL         string      `json:"events_url"`
	HTMLURL           string      `json:"html_url"`
	Labels            []string    `json:"labels"`
	State             string      `json:"state"`
	Locked            bool        `json:"locked"`
	Assignee          string      `json:"assignee"`
	Assignees         []string    `json:"assignees"`
	Milestone         string      `json:"milestone"`
	Comments          int32       `json:"comments"`
	User              User        `json:"user"`
	CreatedAt         string      `json:"created_at"`
	UpdatedAt         string      `json:"updated_at"`
	ClosedAt          string      `json:"closed_at"`
	AuthorAssociation string      `json:"author_association"`
	Body              string      `json:"body"`
	PullRequest       PullRequest `json:"pull_request"`
}

// Member member
type Member struct {
	Login string `json:"login"`
}

// Owner owner
type Owner struct {
	ID           int32  `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
	Type         string `json:"type"`
	SiteAdmin    string `json:"site_admin"`
}

// PRReviewEvent review PR
func PRReviewEvent(item NewsFeed) (string, string) {
	user, repo, avatar := getFeedBaseInfoAndPrintAvatar(item)
	number := item.Payload.PullRequest.Number
	return avatar, fmt.Sprintf("%s reviewed pull request %s on %s\n\n \a at %v", user, number, repo, item.CreatedAt)
}

// PREvent open PR, close PR
func PREvent(item NewsFeed) (string, string) {
	user, repo, avatar := getFeedBaseInfoAndPrintAvatar(item)
	state := item.Payload.PullRequest.State
	number := item.Payload.PullRequest.Number
	title := item.Payload.PullRequest.Title
	body := item.Payload.PullRequest.Body

	if state == "open" {
		return avatar, fmt.Sprintf("%s opened pull request %s on %s \n %s \n %s \a at %v\n", user, number, repo, title, body, item.CreatedAt)
	}
	return avatar, fmt.Sprintf("%s closed pull request %s on %s \n %s \a at %v\n", user, number, repo, title, item.CreatedAt)
}

// comment on issue, PR
func issueCommentEvent(item NewsFeed) (string, string) {
	user, repo, avatar := getFeedBaseInfoAndPrintAvatar(item)
	number := item.Payload.Issue.Number
	body := item.Payload.Comment.Body

	group := ""
	if item.Payload.Issue.PullRequest.Body != "" {
		group = "pull request"
	} else {
		group = "issue"
	}

	return avatar, fmt.Sprintf("%s commented on %s %s on %s \n %s \a at %v\n", user, group, number, repo, body, item.CreatedAt)
}

// open issue, close issue
func issuesEvent(item NewsFeed) (string, string) {
	user, repo, avatar := getFeedBaseInfoAndPrintAvatar(item)
	state := item.Payload.Action
	number := item.Payload.Issue.Number
	title := item.Payload.Issue.Title

	return avatar, fmt.Sprintf("%s %s issue %s on %s \n %s \a at %v\n\n", user, state, number, repo, title, item.CreatedAt)
}

// comment on a commit
func commitCommentEvent(item NewsFeed) (string, string) {
	user, repo, avatar := getFeedBaseInfoAndPrintAvatar(item)
	body := item.Payload.Comment.Body
	return avatar, fmt.Sprintf("%s commented on %s \n %s \a at %v\n", user, repo, body, item.CreatedAt)
}

// # starred by following
func watchEvent(item NewsFeed) (string, string) {
	user, repo, avatar := getFeedBaseInfoAndPrintAvatar(item)
	return avatar, fmt.Sprintf("%s starred %s \a at %v \n\n", user, repo, item.CreatedAt)
}

// # forked by following
func forkEvent(item NewsFeed) (string, string) {
	user, repo, avatar := getFeedBaseInfoAndPrintAvatar(item)
	return avatar, fmt.Sprintf("%s forked %s \a at %v\n\n", user, repo, item.CreatedAt)
}

// # delete branch
func deleteEvent(item NewsFeed) (string, string) {
	user, repo, avatar := getFeedBaseInfoAndPrintAvatar(item)
	branch := item.Payload.Ref
	return avatar, fmt.Sprintf("%s deleted branch %s at \a at %v%s\n\n", user, branch, repo, item.CreatedAt)
}

// # push commits
func pushEvent(item NewsFeed) (string, string) {
	user, repo, avatar := getFeedBaseInfoAndPrintAvatar(item)
	size := item.Payload.Size
	branch := item.Payload.Ref
	return avatar, fmt.Sprintf("%s pushed %d new commit(s) to %s at %s \a at %v\n\n", user, size, branch, repo, item.CreatedAt)
}

// # create repo, branch
func createEvent(item NewsFeed) (string, string) {
	user, repo, avatar := getFeedBaseInfoAndPrintAvatar(item)
	group := item.Payload.RefType
	branch := item.Payload.Ref
	if group == "repository" {
		return avatar, fmt.Sprintf("%s created %s %s \a at %v\n\n", user, group, repo, item.CreatedAt)
	}
	return avatar, fmt.Sprintf("%s created %s %s at %s \a at %v\n\n", user, group, branch, repo, item.CreatedAt)
}

// # make public repo
func publicEvent(item NewsFeed) (string, string) {
	user, repo, avatar := getFeedBaseInfoAndPrintAvatar(item)
	return avatar, fmt.Sprintf("%s made %s public \a at %v\n\n", user, repo, item.CreatedAt)
}

// # add collab
func memberEvent(item NewsFeed) (string, string) {
	action := item.Payload.Action
	collab := item.Payload.Member.Login
	user, repo, avatar := getFeedBaseInfoAndPrintAvatar(item)
	return avatar, fmt.Sprintf("%s %s %s as a collaborator to %s \a at %v\n\n", user, action, collab, repo, item.CreatedAt)
}

// getFeedBaseInfoAndPrintAvatar 获取feed基本信息
func getFeedBaseInfoAndPrintAvatar(item NewsFeed) (string, string, string) {
	user := item.Actor.Login
	repo := item.Repo.Name
	avatarURL := item.Actor.AvatarURL
	return user, repo, loadAvatar(avatarURL)
}

func loadAvatar(avatarURL string) (avatar string) {
	if len(avatarURL) > 0 {
		res, err := http.Get(avatarURL)
		if err != nil {
			log.Fatalf("http get %s, err:%v", avatarURL, err)
			return ""
		}
		if res.Body != nil {
			defer res.Body.Close()
		}
		avatar = display(res.Body)
	}
	return avatar
}

func getReceivedEvents(user, pageNo, include, exclude string) {
	url := "https://api.github.com/users/" + user + "/received_events?page=" + pageNo

	startTime := time.Now()
	_, data, _ := GetJSON(url)
	cost("request Github API: /users/:user/received_events", startTime)
	// TODO: optimize
	r := gjson.Parse(data)

	startTime = time.Now()
	for _, it := range r.Array() {
		st := time.Now()
		event := it.Get("type").String()
		item := NewsFeed{}
		json.Unmarshal([]byte(it.String()), &item)
		// item.CreatedAt = it.Get("created_at").String()

		var content, avatar string
		switch event {
		case "PullRequestReviewCommentEvent":
			avatar, content = PRReviewEvent(item)
		case "PullRequestEvent":
			avatar, content = PREvent(item)
		case "IssueCommentEvent":
			avatar, content = issueCommentEvent(item)
		case "IssuesEvent":
			avatar, content = issuesEvent(item)
		case "CommitCommentEvent":
			avatar, content = commitCommentEvent(item)
		case "WatchEvent":
			avatar, content = watchEvent(item)
		case "ForkEvent":
			avatar, content = forkEvent(item)
		case "DeleteEvent":
			avatar, content = deleteEvent(item)
		case "PushEvent":
			avatar, content = pushEvent(item)
		case "CreateEvent":
			avatar, content = createEvent(item)
		case "PublicEvent":
			avatar, content = publicEvent(item)
		case "MemberEvent":
			avatar, content = memberEvent(item)
		default: // do nothing
		}

		output(avatar, content, include, exclude)
		cost("every one", st)
	}
	cost("every page", startTime)
}

// ReceivedEvents get received events
func ReceivedEvents(user string, maxPage int, debug bool, include, exclude string) {
	for page := 1; page <= maxPage; page++ {
		getReceivedEvents(user, fmt.Sprintf("%d", page), include, exclude)
	}
}

// display 控制台打印图片
func display(r io.Reader) string {
	data, err := ioutil.ReadAll(r)
	if err != nil {
	}

	width, height := widthAndHeight()

	var buf bytes.Buffer
	buf.WriteString("\033]1337;File=inline=1")

	if width != "" {
		buf.WriteString(";width=")
		buf.WriteString(width)
	}
	if height != "" {
		buf.WriteString(";height=")
		buf.WriteString(height)
	}
	buf.WriteString(":")
	buf.WriteString(base64.StdEncoding.EncodeToString(data))
	buf.WriteString("\a")
	return buf.String()
}

func widthAndHeight() (w, h string) {
	if width != "" {
		w = width
	}
	if height != "" {
		h = height
	}
	if size != "" {
		sp := strings.SplitN(size, ",", -1)
		if len(sp) == 2 {
			w = sp[0]
			h = sp[1]
		}
	}
	return
}

func output(avatar, content, include, exclude string) {
	if (len(exclude) > 0 && wpm.WildcardPatternMatch(content, "*"+exclude+"*")) ||
		(len(include) > 0 && !wpm.WildcardPatternMatch(content, "*"+include+"*")) {
		return
	}
	fmt.Print(avatar, content)
}

func cost(fmtStr string, startTime time.Time) {
	if debug {
		defer fmt.Printf("%s , cost (%v)\n", fmtStr, time.Now().Sub(startTime))
	}
}

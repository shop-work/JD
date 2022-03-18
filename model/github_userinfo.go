/*******
* @Author:qingmeng
* @Description:
* @File:github_userinfo
* @Date2022/3/18
 */

package model

type GitHubUserinfo struct {
	Bio               string  `json:"bio"`
	Blog              string  `json:"blog"`
	Company           string  `json:"company"`
	CreatedAt         string  `json:"created_at"`
	Email             string  `json:"email"`
	EventsUrl         string  `json:"events_url"`
	Followers         int     `json:"followers"`
	FollowersUrl      string  `json:"followers_url"`
	Following         int     `json:"following"`
	followingUrl      string  `json:"following_url"`
	gistsUrl          string  `json:"gists_url"`
	gravatarId        string  `json:"gravatar_id"`
	htmlUrl           string  `json:"html_url"`
	id                float32 `json:"id"`
	location          string  `json:"location"`
	Login             string  `json:"login"` //github账号
	Name              string  `json:"name"`  //github昵称
	nodeId            string  `json:"node_id"`
	organizationsUrl  string  `json:"organizations_url"`
	publicGists       int     `json:"public_gists"`
	publicRepos       int     `json:"public_repos"`
	receivedEventsUrl string  `json:"received_events_url"`
	reposUrl          string  `json:"repos_url"`
	siteAdmin         bool    `json:"site_admin"`
	starredUrl        string  `json:"starred_url"`
	subscriptionsUrl  string  `json:"subscriptions_url"`
	twitterUsername   string  `json:"twitter_username"`
	Type              string  `json:"type"`
	updatedAt         string  `json:"updated_at"`
	url               string  `json:"url"`
}

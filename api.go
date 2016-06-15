package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GetTotalPosts -- determines the total posts on a blog
func GetTotalPosts(blog, key string) (int, error) {
	u := buildBlogURL(blog, key)

	fmt.Printf("Executing request:[%s]\n", u)
	response, err := http.Get(u)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	b := struct {
		Response struct {
			Blog Blog `json:"blog"`
		} `json:"response"`
	}{}
	json.NewDecoder(response.Body).Decode(&b)

	return b.Response.Blog.Posts, nil
}

// GetPosts -- returns the posts for the blog in the range (limit, limit+offset)
func GetPosts(blog, key string, limit, offset int) (*[]Post, error) {
	u := buildPostURL(blog, key, limit, offset)

	fmt.Printf("Executing request:[%s]\n", u)
	response, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	r := struct {
		Response struct {
			Posts []Post `json:"posts"`
		} `json:"response"`
	}{}

	json.NewDecoder(response.Body).Decode(&r)

	return &r.Response.Posts, nil
}

func buildBlogURL(blog, key string) string {
	u := url.URL{
		Scheme: "https",
		Host:   "api.tumblr.com",
		Path:   fmt.Sprintf("v2/blog/%s.tumblr.com/info", blog),
	}

	query := u.Query()
	query.Set("api_key", key)
	u.RawQuery = query.Encode()

	return u.String()
}

func buildPostURL(blog, key string, limit, offset int) string {
	u := url.URL{
		Scheme: "https",
		Host:   "api.tumblr.com",
		Path:   fmt.Sprintf("v2/blog/%s.tumblr.com/posts", blog),
	}

	query := u.Query()
	query.Set("api_key", key)
	query.Set("filter", "text")
	query.Set("reblog_info", "false")
	query.Set("limit", strconv.Itoa(limit))
	query.Set("offset", strconv.Itoa(offset))

	u.RawQuery = query.Encode()

	return u.String()
}

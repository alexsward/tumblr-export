package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	typeText     = "text"
	typePhoto    = "photo"
	pngExtension = "png"
	jpgExtension = "jpg"
	gifExtension = "gif"
)

var regex *regexp.Regexp

func init() {
	regex, _ = regexp.Compile("[^A-Za-z0-9]+")
}

// HandlePosts handles all of the posts and saves them
func HandlePosts(posts *[]Post, path string) error {
	for _, post := range *posts {
		if typeText == post.Type {
			HandleTextPost(&post, path)
		} else if typePhoto == post.Type {
			HandleImagePost(&post, path)
		} else {
			continue
		}
	}
	return nil
}

// HandleTextPost -- handles a text post from tumblr
func HandleTextPost(post *Post, path string) error {
	filename := getFilename(post.Date.Time, "txt", post.Title, path)
	err := ioutil.WriteFile(filename, []byte(post.Body), 0666)
	if err != nil {
		fmt.Printf("Error writing file:[%s] - %s", filename, err)
		return err
	}

	return nil
}

// HandleImagePost -- handles an image post from tumblr
func HandleImagePost(post *Post, path string) error {
	for i, photo := range post.Photos {
		title := fmt.Sprintf("%s_%d", post.Title, i)
		extension := jpgExtension
		if strings.HasSuffix(photo.Original.URL, pngExtension) {
			extension = pngExtension
		} else if strings.HasSuffix(photo.Original.URL, gifExtension) {
			continue
		}

		filename := getFilename(post.Date.Time, extension, title, path)

		r, err := http.Get(photo.Original.URL)
		if err != nil {
			fmt.Printf("Error downloading image:[%s] -- %s", photo.Original.URL, err)
			continue
		}
		defer r.Body.Close()

		img, _, err := image.Decode(r.Body)
		if err != nil {
			fmt.Printf("Error decoding image:[%s] -- %s", photo.Original.URL, err)
			continue
		}

		f, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Error creating file:[%s] - %s", filename, err)
			continue
		}
		defer f.Close()

		if extension == jpgExtension {
			jpeg.Encode(f, img, &jpeg.Options{
				Quality: 100,
			})
		} else if extension == pngExtension {
			png.Encode(f, img)
		}
	}

	return nil
}

func getFilename(t time.Time, extension, title, path string) string {
	date := fmt.Sprintf("%4d%02d%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	name := regex.ReplaceAllString(title, "")

	filename := fmt.Sprintf("%s_%s.%s", date, name, extension)
	return fmt.Sprintf("%s/%s", path, filename)
}

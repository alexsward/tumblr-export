package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type config struct {
	Blog   string
	APIKey string
	Output string
}

func main() {
	conf, err := parse()
	if err != nil {
		fmt.Printf("Error parsing arguments: %s\n", err)
		os.Exit(-1)
	}

	fmt.Printf("Exporting tumblr posts. blog:[%s]\n", conf.Blog)

	count, err := GetTotalPosts(conf.Blog, conf.APIKey)
	if err != nil {
		fmt.Printf("Error getting total posts %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total posts: [%v]\n", count)

	for i := 1; i <= count; i += 20 {
		fmt.Printf("Getting posts (%v-%v)\n", i, i+20)

		posts, err := GetPosts(conf.Blog, conf.APIKey, 20, i)
		if err != nil {
			fmt.Printf("Error getting posts: %s\n", err)
			os.Exit(2)
		}

		err = HandlePosts(posts, conf.Output)
		if err != nil {
			fmt.Printf("Error handling posts: %s", err)
		}
	}
}

func parse() (*config, error) {
	c := &config{}

	flags := flag.NewFlagSet("tumblr importer", flag.ExitOnError)
	flags.StringVar(&c.Blog, "blog", "", "tumblr blog name without tumblr.com - blogname.tumblr.com")
	flags.StringVar(&c.APIKey, "key", "", "API key for tumblr")
	flags.StringVar(&c.Output, "output", "", "output directory for the saved files")

	flags.Parse(os.Args[1:])

	if c.Blog == "" {
		flags.PrintDefaults()
		return nil, errors.New("blog cannot be empty")
	}

	if c.APIKey == "" {
		flags.PrintDefaults()
		return nil, errors.New("api key cannot be empty")
	}

	return c, nil
}

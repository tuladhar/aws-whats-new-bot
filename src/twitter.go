/*
 * aws-whats-new-bot
 * Copyright (c) 2020 - Puru Tuladhar (ptuladhar3@gmail.com)
 * See LICENSE file.
 */
package main

import (
    "os"
    "log"
    "fmt"

	"github.com/mmcdole/gofeed"	

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Twitter struct {
    client *twitter.Client
}

func (t *Twitter) InitializeClient(creds *Credentials) {
    log.Println("Initializing Twitter client...")
    client, err := t.getClient(creds)
    if err != nil {
        log.Fatalln("Unable to create Twitter Client:", err)
    }
    t.client = client
}

func (t *Twitter) getClient(creds *Credentials) (*twitter.Client, error) {
    config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
    token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

    httpClient := config.Client(oauth1.NoContext, token)
    client := twitter.NewClient(httpClient)

    return client, nil
}

func (t *Twitter) VerifyCredentials() {
    log.Println("Verifying Twitter credentials...")
    verifyParams := &twitter.AccountVerifyParams{
        SkipStatus:   twitter.Bool(true),
        IncludeEmail: twitter.Bool(true),
    }

    user, _, err := t.client.Accounts.VerifyCredentials(verifyParams)
    if err != nil {
        log.Fatalln("Unable to verify credentials:", err)
    }

    if _, ok := os.LookupEnv("DEBUG"); ok {
        log.Printf("%+v", user)
    }
}

func (t *Twitter) Tweet(announcement *gofeed.Item) {
    var (
        title = announcement.Title
        link = announcement.Link
    )

    if len(title) > 250 {
        title = fmt.Sprintf("%s...", title[:250])
    } else {
        title = fmt.Sprintf("%s.", title)
    }
    
    status := fmt.Sprintf("%s\n%s", title, link)
    log.Println("Tweeting announcement:", status)

    _, _, err := t.client.Statuses.Update(status, nil)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("... Tweeted!")
}

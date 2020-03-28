package user

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"social-dashboard/api/constants"
	"social-dashboard/api/models/user"
	"social-dashboard/api/utils"
	"sort"
	"strings"

	"github.com/labstack/echo"
	m "github.com/veer66/mapkha"
)

type Account struct {
	OwnerName  string `json:"owner_name"`
	NumMessage int    `json:"number_of_message"`
}

type MessageCount struct {
	Date  string `json:"date"`
	Count int    `json:"number_of_message"`
}

func CreateUsers() error {
	fileURL := "https://s3-ap-southeast-1.amazonaws.com/wisesight-public/dev-test/rawdata.csv"

	if err := downloadFile("./rawdata.csv", fileURL); err != nil {
		return err
	}

	file, err := ioutil.ReadFile("./rawdata.csv")
	if err != nil {
		return err
	}

	_, err = user.Save(file)
	if err != nil {
		return err
	}
	return err
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func GetMessagePerDay(c echo.Context) error {
	daily := make(map[string]int)
	for _, user := range user.GetUsers() {
		daily[user.Date.Format(constants.Date)]++
	}
	return c.JSON(http.StatusOK, printMessagePerDay(daily))
}

func printMessagePerDay(data map[string]int) []MessageCount {
	var messageCounts []MessageCount
	for k, v := range data {
		messageCounts = append(messageCounts, MessageCount{Date: k, Count: v})
	}
	sort.Slice(messageCounts, func(i, j int) bool {
		return messageCounts[i].Date < messageCounts[j].Date
	})
	return messageCounts
}

func Get10AccountByMessage(c echo.Context) error {
	account := make(map[string]int)
	for _, user := range user.GetUsers() {
		if strings.Compare(user.OwnerName, "Unknown") != 0 {
			account[user.OwnerName]++
		}
	}
	return c.JSON(http.StatusOK, print10AccountByMessage(account))
}

func print10AccountByMessage(data map[string]int) []Account {
	var accounts []Account
	for k, v := range data {
		accounts = append(accounts, Account{k, v})
	}
	sort.Slice(accounts, func(i, j int) bool {
		return accounts[i].NumMessage > accounts[j].NumMessage
	})
	return accounts[:10]
}

func Get10MessageByEngagement(c echo.Context) error {
	users := user.GetUsers()
	sort.Slice(users, func(i, j int) bool {
		return users[i].Engagement > users[j].Engagement
	})
	newUsers := []user.User{}
	for _, u := range users {
		if strings.Compare(u.OwnerName, "Unknown") != 0 {
			newUsers = append(newUsers, user.User{Message: u.Message, Engagement: u.Engagement})
		}
		if len(newUsers) == 10 {
			break
		}
	}
	return c.JSON(http.StatusOK, newUsers)
}

func GetWordClouds(c echo.Context) error {
	users := user.GetUsers()
	word := make(map[string]int)
	dict, err := m.LoadDefaultDict()
	if err != nil {
		log.Println("load dic error : ", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error",
		})
	}
	wordcut := m.NewWordcut(dict)
	reg, err := regexp.Compile("[^\u0E00-\u0E7Fa-zA-Z0-9]+")
	if err != nil {
		log.Println("regexp compile error : ", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error",
		})
	}
	for _, u := range users {
		wcuts := wordcut.Segment(reg.ReplaceAllString(u.Message, ""))
		for _, wcut := range wcuts {
			word[wcut]++
		}
	}
	err = utils.GererateWordClouds(word, "./result/word_clouds.png", 200, 1000, 100)
	if err != nil {
		log.Println("generate word clouds error : ", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "please see word clouds file at result/word_clouds.png",
	})
}

func GetHashtagClouds(c echo.Context) error {
	users := user.GetUsers()
	hashtag := make(map[string]int)

	r, err := regexp.Compile("#[A-Za-z0-9]+")
	if err != nil {
		log.Println("compile regexp error : ", err)
		return err
	}
	for _, u := range users {
		if strings.Compare("facebook", u.Channel) == 0 || strings.Compare("twitter", u.Channel) == 0 ||
			strings.Compare("instagram", u.Channel) == 0 {
			messages := strings.Fields(u.Message)
			for _, message := range messages {
				if str := r.FindString(message); len(str) != 0 {
					hashtag[str]++
				}
			}
		}
	}

	err = utils.GererateWordClouds(hashtag, "./result/hashtag_clouds.png", 100, 100, 60)
	if err != nil {
		log.Println("generate word clouds error : ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "please see hashtag clouds file at result/hashtag_clouds.png",
	})
}

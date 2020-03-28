package user

import (
	"bytes"
	"fmt"
	"io"
	"social-dashboard/api/constants"
	"strconv"
	"strings"
	"time"
)

type User struct {
	ID         string     `json:"id,omitempty"`
	Type       string     `json:"type,omitempty"`
	Message    string     `json:"message,omitempty"`
	Date       *time.Time `json:"date,omitempty"`
	Engagement int        `json:"engagement,omitempty"`
	Channel    string     `json:"channel,omitempty"`
	OwnerID    string     `json:"owner_id,omitempty"`
	OwnerName  string     `json:"owner_name,omitempty"`
}

var Users []User

func Save(file []byte) ([]User, error) {
	var isFirstLine bool
	var user User
	var count int
	buf := bytes.NewBuffer(file)
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		datas := strings.Split(line, ",")
		if isFirstLine {
			for i, data := range datas {
				switch count {
				case 0:
					user.ID = data
					count++
				case 1:
					user.Type = data
					count++
				case 2:
					user.Message = data
					count++
				case 3:
					t, err := time.Parse(constants.DateWithTime, data)
					if err == nil {
						user.Date = &t
						count++
					} else {
						user.Message = fmt.Sprintf("%s %s", user.Message, data)
					}
				case 4:
					engagement, err := strconv.Atoi(data)
					if err != nil {
						return nil, err
					}
					user.Engagement = engagement
					count++
				case 5:
					user.Channel = data
					count++
				case 6:
					user.OwnerID = data
					count++
				case 7:
					user.OwnerName = strings.TrimRight(strings.Join(datas[i:], " "), "\n")
					Users = append(Users, user)
					count = 0
					goto checkpoint
				}
			}
		checkpoint:
		} else {
			isFirstLine = true
		}
	}
	return Users, nil
}

func GetUsers() []User {
	return Users
}

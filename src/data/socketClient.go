package data

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/timotew/etc/src/common"
	"github.com/timotew/etc/src/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// SocketClient web socket client
	SocketClient struct {
		// socket is the web socket for this client.
		socket *websocket.Conn
		// send is a channel on which messages are sent.
		send chan []byte
		// room is the room this client is chatting in.
		room *SocketRoom
		// connected use
		user string
	}
	TestSession struct {
		Time      int    `json:"seconds"`
		SessionID string `json:"session"`
	}

	QuestionOptions struct {
		Choice  models.Choice   `json:"choice"`
		Options []models.Option `json:"options"`
	}

	QuestionAnswer struct {
		QuestionID string `json:"questionId"`
		OptionID   string `json:"optionId"`
	}
)

func (c *SocketClient) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			smsg := string(msg)
			duo := strings.Split(smsg, "|")
			switch duo[0] {
			case "loadOptions":
				id := duo[1]
				isHex := bson.IsObjectIdHex(id)
				if isHex {
					session := common.GetSession().Copy()
					context := session.DB(common.AppConfig.Database)
					col := context.C("options")
					col1 := context.C("choices")
					defer session.Close()
					repo := &OptionRepository{C: col}
					repo1 := &ChoiceRepository{C: col1}
					options, _ := repo.GetByQuestion(id)
					choice, _ := repo1.GetByQuestionID(id, c.user)
					d := &QuestionOptions{
						Options: options,
						Choice:  choice,
					}
					j, err := json.Marshal(d)
					h := []byte(duo[1] + "|")
					h = append(h, j...)
					if err == nil {
						c.send <- h
					}
				}
				break
			case "updateTime":
				var session TestSession
				err := json.Unmarshal([]byte(duo[1]), &session)
				if err != nil {
					return
				}
				current := time.Duration(session.Time) * time.Second
				isHex := bson.IsObjectIdHex(session.SessionID)
				if isHex {
					dbSession := common.GetSession().Copy()
					context := dbSession.DB(common.AppConfig.Database)
					col := context.C("testSessions")
					defer dbSession.Close()
					repo := &TestSessionRepository{C: col}
					status := &models.TestSession{
						ID:     bson.ObjectIdHex(session.SessionID),
						Time:   current,
						Active: true,
					}
					err := repo.Update(status)
					if err != nil {
						fmt.Println(err)
						return
					}

				}
				break
			case "markQuestion":
				var report QuestionAnswer
				err := json.Unmarshal([]byte(duo[1]), &report)
				if err != nil {
					fmt.Println(err)
					return
				}
				session := common.GetSession().Copy()
				context := session.DB(common.AppConfig.Database)
				col := context.C("options")
				repo1 := &OptionRepository{C: col}
				option, _ := repo1.GetByID(report.OptionID)
				colTwo := context.C("choices")
				repoTwo := &ChoiceRepository{C: colTwo}
				oldChoice, errTwo := repoTwo.GetByQuestionOption(report.QuestionID, report.OptionID)

				if oldChoice.Correct {

				}
				if errTwo == mgo.ErrNotFound {
					// before this
					if option.Correct {

					}
					return
				}

			}
			//c.room.Forward <- msg

		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *SocketClient) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

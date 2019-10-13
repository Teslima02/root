package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"runtime"

	gomail "gopkg.in/gomail.v2"
)

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResource struct {
		Data appError `json:"data"`
	}
	configuration struct {
		Server, MongoDBHost, DBUser, DBPwd, Database, MailServer, MailUser, MailPwd, FBAccountKitAPIVersion, FBAppID, FBAppSecret, AKEndpointBaseURL, AKTokenExchangeBaseURL string
		MailPort, LogLevel                                                                                                                                                   int
	}
)

// DisplayAppError disaply HTTP error
func DisplayAppError(w http.ResponseWriter, handlerError error, message string, code int) {
	errObj := appError{
		Error:      handlerError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	//log.Printf("AppError]: %s\n", handlerError)
	Error.Printf("AppError]: %s\n", handlerError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResource{Data: errObj}); err == nil {
		w.Write(j)
	}
}

// SendMail send email to users
func SendMail(receivers []string, subject, body string) error {
	d := gomail.NewDialer(AppConfig.MailServer, AppConfig.MailPort, AppConfig.MailUser, AppConfig.MailPwd)
	s, err := d.Dial()
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	for _, r := range receivers {
		m.SetHeader("From", AppConfig.MailUser)
		m.SetAddressHeader("To", r, r)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", fmt.Sprintf("Hello %s!"+body, r))

		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", r, err)
		}
		m.Reset()
	}

	return err
}

// TODO: ERROR
// func GetDBErrorCode(errorMessage string) string{
// 	m := map[string]int{
// 		"rsc": 3711,
// 		"r":   2138,
// 		"gri": 1908,
// 		"adg": 912,
// 	}
// 	strings.Index()
// 	strings.ContainsAny()

// }

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// Initialize AppConfig
func initConfig() {
	loadAppConfig()
}

// Reads config.json and decode into AppConfig
func loadAppConfig() {
	_, filename, _, _ := runtime.Caller(1)
	filepath := path.Join(path.Dir(filename), "../common/config.json")

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}

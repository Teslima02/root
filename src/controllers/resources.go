package controllers

import (
	"github.com/timotew/etc/src/models"
)

// Models for JSON resources
type (
	TmpUserResource struct {
		Data TmpUserModel `json:"data"`
	}

	TmpUserModel struct {
		MatricNo string `json:"matricNo"`
		FullName string `json:"fullName"`
	}

	/**
	{"id":"612074265910092","phone":{"number":"+2348035387514","country_prefix":"234","national_number":"8035387514"},"application":{"id":"2051440904975153"}}*/

	FBKitCompleteStageTwoResource struct {
		ID    string                  `json:"id"`
		Phone FBKitCompletePhoneModel `json:"phone"`
		Email FBKitCompleteEmailModel `json:"email"`
	}

	FBKitCompleteEmailModel struct {
		Address string `json:"address"`
	}

	FBKitCompletePhoneModel struct {
		NationalNumber string `json:"national_number"`
		CountryPrefix  string `json:"country_prefix"`
		Number         string `json:"number"`
	}
	FBKitCompleteStageOneModel struct {
		ID                 string `json:"id"`
		AccessToken        string `json:"access_token"`
		RefreshIntervalSec int    `json:"token_refresh_interval_sec"`
	}

	FBKitCompleteModel struct {
		Code string `json:"code"`
	}

	FBKitCompleteResource struct {
		Data FBKitCompleteModel `json:"data"`
	}

	FBKitPartialResource struct {
		Data FBKitPartialModel `json:"data"`
	}
	FBKitPartialModel struct {
		LoginMode    string `json:"loginMode"`
		PhoneNumber  string `json:"phoneNumber"`
		EmailAddress string `json:"emailAddress"`
		Token        string `json:"token"`
		Status       string `json:"status"`
	}

	// For Post - /user/register

	UserResource struct {
		Data models.User `json:"data"`
	}

	// UsersResource getting multiple users
	UsersResource struct {
		Data []models.User `json:"data"`
	}

	FindUsernameResource struct {
		Text       string   `json:"username"`
		InviteList []string `json:"inviteList"`
	}

	/*
	*END resource for all routes
	 */

	// For Post - /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	// Response for authorized user Post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	// Model for authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// Model for authorized user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}

	//ClassModel for classes modeling
	ClassModel struct {
		MasterID    string `json:"masterId"`
		ClassName   string `json:"title"`
		Description string `json:"description"`
		Alias       string `json:"code"`
		Published   bool   `json:"published"`
	}

	//ClassesResource fro mutilple classes
	ClassesResource struct {
		Data []models.Class `json:"data"`
	}
	//TestsResource for multiple
	TestsResource struct {
		Data []models.Test `json:"data"`
	}

	TestSessionResource struct {
		Data TestSessionModel `json:"data"`
	}

	//TestModel for test
	TestModel struct {
		Title                  string   `json:"title"`
		Description            string   `json:"description"`
		ClassID                string   `json:"classroomId"`
		QuestionsPerAssessment int      `json:"questionsPerAssessment"`
		DueDate                string   `json:"dueDate"`
		AccessDate             string   `json:"accessDate"`
		Duration               string   `json:"duration"`
		Tags                   []string `json:"tags"`
	}

	TestSessionModel struct {
		TestID string `json:"testId"`
		UserID string `json:"userId"`
		Time   string `json:"time"`
		Active bool   `json:"active"`
	}

	//QuestionResource  jd
	QuestionResource struct {
		Data QuestionModel `json:"data"`
	}

	//QuestionModel for modeling questions
	QuestionModel struct {
		TestID string `json:"testId"`
		Body   string `json:"body"`
		Point  int    `json:"point"`
	}
	//QuestionsResource for multiple questions
	QuestionsResource struct {
		Data []models.Question `json:"data"`
	}

	//OptionResource for requests
	OptionResource struct {
		Data OptionModel `json:"data"`
	}

	//OptionModel for selected options
	OptionModel struct {
		QuestionID string `json:"questionId"`
		Correct    bool   `json:"correct"`
		Body       string `json:"body"`
	}

	//OptionsResource for getting multiple tptios
	OptionsResource struct {
		Data []models.Option `json:"data"`
	}

	//ChoiceResource for resquests
	ChoiceResource struct {
		Data ChoiceModel `json:"data"`
	}

	//ChoiceModel for model of choices
	ChoiceModel struct {
		OptionID   string `json:"optionId"`
		Correct    bool   `json:"correct"`
		QuestionID string `json:"questionId"`
	}

	//ChoicesResource for submitting multiple choices
	ChoicesResource struct {
		Data []models.Choice `json:"data"`
	}

	//ResultResource for requests
	ResultResource struct {
		Data ResultModel `json:"data"`
	}

	//ResultModel for result filter
	ResultModel struct {
		TestID string `json:"testId"`
		UserID string `json:"userId"`
		Points int    `json:"points"`
		Status string `json:"status"`
	}

	//ResultsResource for multiple results
	ResultsResource struct {
		Data []models.Result `json:"data"`
	}

	//ClassResource for requests
	ClassResource struct {
		Data ClassModel `json:"data"`
	}

	ManyClassResource struct {
		Data []ClassModel `json:"data"`
	}
	//TestResource for requests
	TestResource struct {
		Data TestModel `json:"data"`
	}
)

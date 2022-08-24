//
//  Copyright 2022, ysicing
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.
//

package zentao

import (
	"fmt"

	"github.com/imroc/req/v3"
)

type ProductPlansService struct {
	client *Client
}

type ProductPlanMeta struct {
	Branch string `json:"branch,omitempty"`
	Title  string `json:"title"`
	Parent int    `json:"parent,omitempty"`
	Desc   string `json:"desc,omitempty"`
	Begin  string `json:"begin,omitempty"`
	End    string `json:"end,omitempty"`
}

type ProductPlanCreateMsg struct {
	ID           int           `json:"id"`
	Product      int           `json:"product"`
	Branch       int           `json:"branch"`
	Parent       int           `json:"parent"`
	Title        string        `json:"title"`
	Status       string        `json:"status"`
	Desc         string        `json:"desc"`
	Begin        string        `json:"begin"`
	End          string        `json:"end"`
	Order        string        `json:"order"`
	ClosedReason string        `json:"closedReason"`
	Deleted      bool          `json:"deleted"`
	Stories      []interface{} `json:"stories"`
	Bugs         []interface{} `json:"bugs"`
}

type ProductPlanListMsg struct {
	Page  int `json:"page"`
	Total int `json:"total"`
	Limit int `json:"limit"`
	Plans []struct {
		ID           int    `json:"id"`
		Product      int    `json:"product"`
		Branch       int    `json:"branch"`
		Parent       int    `json:"parent"`
		Title        string `json:"title"`
		Status       string `json:"status"`
		Desc         string `json:"desc"`
		Begin        string `json:"begin"`
		End          string `json:"end"`
		Order        string `json:"order"`
		ClosedReason string `json:"closedReason"`
		Deleted      bool   `json:"deleted"`
		Stories      int    `json:"stories"`
		Bugs         int    `json:"bugs"`
		Hour         int    `json:"hour"`
		Project      int    `json:"project"`
		ProjectID    string `json:"projectID"`
		Expired      bool   `json:"expired"`
	} `json:"plans"`
}

type ProductPlanGetMsg struct {
	ID      int    `json:"id"`
	Product int    `json:"product"`
	Branch  int    `json:"branch"`
	Parent  int    `json:"parent"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Begin   string `json:"begin"`
	End     string `json:"end"`
	Order   string `json:"order"`
	Deleted bool   `json:"deleted"`
	Stories []struct {
		Story          int         `json:"story"`
		Plan           string      `json:"plan"`
		Order          int         `json:"order"`
		ID             int         `json:"id"`
		Parent         int         `json:"parent"`
		Product        int         `json:"product"`
		Branch         int         `json:"branch"`
		Module         int         `json:"module"`
		Source         string      `json:"source"`
		Sourcenote     string      `json:"sourceNote"`
		Frombug        int         `json:"fromBug"`
		Title          string      `json:"title"`
		Keywords       string      `json:"keywords"`
		Type           string      `json:"type"`
		Category       string      `json:"category"`
		Pri            int         `json:"pri"`
		Estimate       int         `json:"estimate"`
		Status         string      `json:"status"`
		Substatus      string      `json:"subStatus"`
		Color          string      `json:"color"`
		Stage          string      `json:"stage"`
		Stagedby       string      `json:"stagedBy"`
		Mailto         interface{} `json:"mailto"`
		Openedby       string      `json:"openedBy"`
		Openeddate     string      `json:"openedDate"`
		Assignedto     string      `json:"assignedTo"`
		Assigneddate   string      `json:"assignedDate"`
		Lasteditedby   string      `json:"lastEditedBy"`
		Lastediteddate string      `json:"lastEditedDate"`
		Reviewedby     string      `json:"reviewedBy"`
		Revieweddate   string      `json:"reviewedDate"`
		Closedby       string      `json:"closedBy"`
		Closeddate     string      `json:"closedDate"`
		Closedreason   string      `json:"closedReason"`
		Tobug          int         `json:"toBug"`
		Childstories   string      `json:"childStories"`
		Linkstories    string      `json:"linkStories"`
		Duplicatestory int         `json:"duplicateStory"`
		Version        int         `json:"version"`
		Urchanged      string      `json:"URChanged"`
		Deleted        string      `json:"deleted"`
	} `json:"stories"`
	Bugs []struct {
		ID             int    `json:"id"`
		Project        int    `json:"project"`
		Product        int    `json:"product"`
		Branch         int    `json:"branch"`
		Module         int    `json:"module"`
		Execution      int    `json:"execution"`
		Plan           int    `json:"plan"`
		Story          int    `json:"story"`
		Storyversion   int    `json:"storyVersion"`
		Task           int    `json:"task"`
		Totask         int    `json:"toTask"`
		Tostory        int    `json:"toStory"`
		Title          string `json:"title"`
		Keywords       string `json:"keywords"`
		Severity       int    `json:"severity"`
		Pri            int    `json:"pri"`
		Type           string `json:"type"`
		Os             string `json:"os"`
		Browser        string `json:"browser"`
		Hardware       string `json:"hardware"`
		Found          string `json:"found"`
		Steps          string `json:"steps"`
		Status         string `json:"status"`
		Substatus      string `json:"subStatus"`
		Color          string `json:"color"`
		Confirmed      int    `json:"confirmed"`
		Activatedcount int    `json:"activatedCount"`
		Activateddate  string `json:"activatedDate"`
		Mailto         string `json:"mailto"`
		Openedby       string `json:"openedBy"`
		Openeddate     string `json:"openedDate"`
		Openedbuild    string `json:"openedBuild"`
		Assignedto     string `json:"assignedTo"`
		Assigneddate   string `json:"assignedDate"`
		Deadline       string `json:"deadline"`
		Resolvedby     string `json:"resolvedBy"`
		Resolution     string `json:"resolution"`
		Resolvedbuild  string `json:"resolvedBuild"`
		Resolveddate   string `json:"resolvedDate"`
		Closedby       string `json:"closedBy"`
		Closeddate     string `json:"closedDate"`
		Duplicatebug   int    `json:"duplicateBug"`
		Linkbug        string `json:"linkBug"`
		Case           int    `json:"case"`
		Caseversion    int    `json:"caseVersion"`
		Result         int    `json:"result"`
		Repo           int    `json:"repo"`
		Entry          string `json:"entry"`
		Lines          string `json:"lines"`
		V1             string `json:"v1"`
		V2             string `json:"v2"`
		Repotype       string `json:"repoType"`
		Testtask       int    `json:"testtask"`
		Lasteditedby   string `json:"lastEditedBy"`
		Lastediteddate string `json:"lastEditedDate"`
		Deleted        string `json:"deleted"`
	} `json:"bugs"`
}

type ProductPlanUpdateMsg struct {
	ID      int    `json:"id"`
	Product int    `json:"product"`
	Branch  int    `json:"branch"`
	Parent  int    `json:"parent"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Begin   string `json:"begin"`
	End     string `json:"end"`
	Order   string `json:"order"`
	Deleted bool   `json:"deleted"`
	Stories []struct {
		Story          int         `json:"story"`
		Plan           string      `json:"plan"`
		Order          int         `json:"order"`
		ID             int         `json:"id"`
		Parent         int         `json:"parent"`
		Product        int         `json:"product"`
		Branch         int         `json:"branch"`
		Module         int         `json:"module"`
		Source         string      `json:"source"`
		Sourcenote     string      `json:"sourceNote"`
		Frombug        int         `json:"fromBug"`
		Title          string      `json:"title"`
		Keywords       string      `json:"keywords"`
		Type           string      `json:"type"`
		Category       string      `json:"category"`
		Pri            int         `json:"pri"`
		Estimate       int         `json:"estimate"`
		Status         string      `json:"status"`
		Substatus      string      `json:"subStatus"`
		Color          string      `json:"color"`
		Stage          string      `json:"stage"`
		Stagedby       string      `json:"stagedBy"`
		Mailto         interface{} `json:"mailto"`
		Openedby       string      `json:"openedBy"`
		Openeddate     string      `json:"openedDate"`
		Assignedto     string      `json:"assignedTo"`
		Assigneddate   string      `json:"assignedDate"`
		Lasteditedby   string      `json:"lastEditedBy"`
		Lastediteddate string      `json:"lastEditedDate"`
		Reviewedby     string      `json:"reviewedBy"`
		Revieweddate   string      `json:"reviewedDate"`
		Closedby       string      `json:"closedBy"`
		Closeddate     string      `json:"closedDate"`
		Closedreason   string      `json:"closedReason"`
		Tobug          int         `json:"toBug"`
		Childstories   string      `json:"childStories"`
		Linkstories    string      `json:"linkStories"`
		Duplicatestory int         `json:"duplicateStory"`
		Version        int         `json:"version"`
		Urchanged      string      `json:"URChanged"`
		Deleted        string      `json:"deleted"`
	} `json:"stories"`
	Bugs []struct {
		ID             int    `json:"id"`
		Project        int    `json:"project"`
		Product        int    `json:"product"`
		Branch         int    `json:"branch"`
		Module         int    `json:"module"`
		Execution      int    `json:"execution"`
		Plan           int    `json:"plan"`
		Story          int    `json:"story"`
		Storyversion   int    `json:"storyVersion"`
		Task           int    `json:"task"`
		Totask         int    `json:"toTask"`
		Tostory        int    `json:"toStory"`
		Title          string `json:"title"`
		Keywords       string `json:"keywords"`
		Severity       int    `json:"severity"`
		Pri            int    `json:"pri"`
		Type           string `json:"type"`
		Os             string `json:"os"`
		Browser        string `json:"browser"`
		Hardware       string `json:"hardware"`
		Found          string `json:"found"`
		Steps          string `json:"steps"`
		Status         string `json:"status"`
		Substatus      string `json:"subStatus"`
		Color          string `json:"color"`
		Confirmed      int    `json:"confirmed"`
		Activatedcount int    `json:"activatedCount"`
		Activateddate  string `json:"activatedDate"`
		Mailto         string `json:"mailto"`
		Openedby       string `json:"openedBy"`
		Openeddate     string `json:"openedDate"`
		Openedbuild    string `json:"openedBuild"`
		Assignedto     string `json:"assignedTo"`
		Assigneddate   string `json:"assignedDate"`
		Deadline       string `json:"deadline"`
		Resolvedby     string `json:"resolvedBy"`
		Resolution     string `json:"resolution"`
		Resolvedbuild  string `json:"resolvedBuild"`
		Resolveddate   string `json:"resolvedDate"`
		Closedby       string `json:"closedBy"`
		Closeddate     string `json:"closedDate"`
		Duplicatebug   int    `json:"duplicateBug"`
		Linkbug        string `json:"linkBug"`
		Case           int    `json:"case"`
		Caseversion    int    `json:"caseVersion"`
		Result         int    `json:"result"`
		Repo           int    `json:"repo"`
		Entry          string `json:"entry"`
		Lines          string `json:"lines"`
		V1             string `json:"v1"`
		V2             string `json:"v2"`
		Repotype       string `json:"repoType"`
		Testtask       int    `json:"testtask"`
		Lasteditedby   string `json:"lastEditedBy"`
		Lastediteddate string `json:"lastEditedDate"`
		Deleted        string `json:"deleted"`
	} `json:"bugs"`
}

type StoriesMeta struct {
	Stories []string `json:"stories"`
}

type BugMeta struct {
	Bugs []string `json:"bugs"`
}

type LinkStoriesMsg struct {
	ID      int    `json:"id"`
	Product int    `json:"product"`
	Branch  int    `json:"branch"`
	Parent  int    `json:"parent"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Begin   string `json:"begin"`
	End     string `json:"end"`
	Order   string `json:"order"`
	Deleted bool   `json:"deleted"`
	Stories []struct {
		Story          int    `json:"story"`
		Plan           string `json:"plan"`
		Order          int    `json:"order"`
		ID             int    `json:"id"`
		Parent         int    `json:"parent"`
		Product        int    `json:"product"`
		Branch         int    `json:"branch"`
		Module         int    `json:"module"`
		Source         string `json:"source"`
		Sourcenote     string `json:"sourceNote"`
		Frombug        int    `json:"fromBug"`
		Title          string `json:"title"`
		Keywords       string `json:"keywords"`
		Type           string `json:"type"`
		Category       string `json:"category"`
		Pri            int    `json:"pri"`
		Estimate       int    `json:"estimate"`
		Status         string `json:"status"`
		Substatus      string `json:"subStatus"`
		Color          string `json:"color"`
		Stage          string `json:"stage"`
		Stagedby       string `json:"stagedBy"`
		Mailto         string `json:"mailto"`
		Openedby       string `json:"openedBy"`
		Openeddate     string `json:"openedDate"`
		Assignedto     string `json:"assignedTo"`
		Assigneddate   string `json:"assignedDate"`
		Lasteditedby   string `json:"lastEditedBy"`
		Lastediteddate string `json:"lastEditedDate"`
		Reviewedby     string `json:"reviewedBy"`
		Revieweddate   string `json:"reviewedDate"`
		Closedby       string `json:"closedBy"`
		Closeddate     string `json:"closedDate"`
		Closedreason   string `json:"closedReason"`
		Tobug          int    `json:"toBug"`
		Childstories   string `json:"childStories"`
		Linkstories    string `json:"linkStories"`
		Duplicatestory int    `json:"duplicateStory"`
		Version        int    `json:"version"`
		Urchanged      string `json:"URChanged"`
		Deleted        string `json:"deleted"`
	} `json:"stories"`
	Bugs []interface{} `json:"bugs"`
}

// Create 创建产品计划
func (s *ProductPlansService) Create(id int, plan ProductPlanMeta) (*ProductPlanCreateMsg, *req.Response, error) {
	var u ProductPlanCreateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&plan).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/products/%d/plans", id)))
	return &u, resp, err
}

// List 获取产品计划列表
func (s *ProductPlansService) List(id int) (*ProductPlanListMsg, *req.Response, error) {
	var u ProductPlanListMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/products/%d/plans", id)))
	return &u, resp, err
}

// GetByID 获取计划详情
func (s *ProductPlansService) GetByID(id int) (*ProductPlanGetMsg, *req.Response, error) {
	var u ProductPlanGetMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Get(s.client.RequestURL(fmt.Sprintf("/productplans/%d", id)))
	return &u, resp, err
}

// UpdateByID 修改计划
func (s *ProductPlansService) UpdateByID(id int, plan ProductPlanMeta) (*ProductPlanUpdateMsg, *req.Response, error) {
	var u ProductPlanUpdateMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&plan).
		SetResult(&u).
		Put(s.client.RequestURL(fmt.Sprintf("/productplans/%d", id)))
	return &u, resp, err
}

// DeleteByID 删除计划
func (s *ProductPlansService) DeleteByID(id int) (*CustomResp, *req.Response, error) {
	var u CustomResp
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetResult(&u).
		Delete(s.client.RequestURL(fmt.Sprintf("/productplans/%d", id)))
	return &u, resp, err
}

// LinkStories 产品计划关联需求
func (s *ProductPlansService) LinkStories(id int, story StoriesMeta) (*LinkStoriesMsg, *req.Response, error) {
	var u LinkStoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/productplans/%d/linkstories", id)))
	return &u, resp, err
}

// LinkStories 产品计划关联需求
func (s *ProductPlansService) UnLinkStories(id int, story StoriesMeta) (*LinkStoriesMsg, *req.Response, error) {
	var u LinkStoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/productplans/%d/unlinkstories", id)))
	return &u, resp, err
}

// LinkBugs 产品计划关联Bug
func (s *ProductPlansService) LinkBugs(id int, story StoriesMeta) (*LinkStoriesMsg, *req.Response, error) {
	var u LinkStoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/productplans/%d/linkBugs", id)))
	return &u, resp, err
}

// UnLinkBugs 产品计划取消关联Bug
func (s *ProductPlansService) UnLinkBugs(id int, story StoriesMeta) (*LinkStoriesMsg, *req.Response, error) {
	var u LinkStoriesMsg
	resp, err := s.client.client.R().
		SetHeader("Token", s.client.token).
		SetBody(&story).
		SetResult(&u).
		Post(s.client.RequestURL(fmt.Sprintf("/productplans/%d/unlinkbugs", id)))
	return &u, resp, err
}

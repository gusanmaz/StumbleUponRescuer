package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type UserInfo struct {
	Timestamp int           `json:"_timestamp"`
	SubErrors []interface{} `json:"_sub_errors"`
	Success   bool          `json:"_success"`
	Error     interface{}   `json:"_error"`
	Reason    []interface{} `json:"_reason"`
	User      struct {
		Userid         int    `json:"userid"`
		Username       string `json:"username"`
		Name           string `json:"name"`
		Email          string `json:"email"`
		PictureSmall   string `json:"picture_small"`
		PictureMed     string `json:"picture_med"`
		PictureLarge   string `json:"picture_large"`
		Gender         string `json:"gender"`
		Age            int    `json:"age"`
		Location       string `json:"location"`
		LikesCount     string `json:"likes_count"`
		FollowingCount int    `json:"following_count"`
		FollowersCount int    `json:"followers_count"`
		ChannelsCount  string `json:"channels_count"`
		StumblesCount  string `json:"stumbles_count"`
		ShareFrequency int    `json:"share_frequency"`
		Desc           string `json:"desc"`
		Thumbnail      string `json:"thumbnail"`
		Picture        string `json:"picture"`
		PictureFull    string `json:"picture_full"`
		Source         string `json:"source"`
		Dna            struct {
			Total  int           `json:"_total"`
			Values []interface{} `json:"values"`
		} `json:"dna"`
		IsFollowed                  bool   `json:"is_followed"`
		DismissedOnboardingTutorial bool   `json:"dismissed_onboarding_tutorial"`
		InterestsCount              int    `json:"interests_count"`
		OwnedListsCount             int    `json:"owned_lists_count"`
		FollowedListsCount          int    `json:"followed_lists_count"`
		DiscoveriesCount            string `json:"discoveries_count"`
		IsExpert                    bool   `json:"is_expert"`
	} `json:"user"`
}

type Likes struct {
	Timestamp int           `json:"_timestamp"`
	SubErrors []interface{} `json:"_sub_errors"`
	Success   bool          `json:"_success"`
	Error     interface{}   `json:"_error"`
	Reason    []interface{} `json:"_reason"`
	Likes     struct {
		Offset string  `json:"_offset"`
		Limit  string  `json:"_limit"`
		Total  int     `json:"_total"`
		Values []ALike `json:"values"`
	} `json:"likes"`
}

type ALike struct {
	URL           string        `json:"url"`
	Urlid         string        `json:"urlid"`
	Title         string        `json:"title"`
	Views         string        `json:"views"`
	Likes         string        `json:"likes"`
	PictureSmall  string        `json:"picture_small"`
	PictureMed    string        `json:"picture_med"`
	PictureLarge  string        `json:"picture_large"`
	PictureNarrow string        `json:"picture_narrow"`
	PictureWide   string        `json:"picture_wide"`
	Tags          []interface{} `json:"tags"`
	Interests     struct {
		Name   interface{} `json:"name"`
		Total  int         `json:"_total"`
		Values []struct {
			ID        string `json:"id"`
			Name      string `json:"name"`
			Frequency int    `json:"frequency"`
			PicThumb  string `json:"pic_thumb"`
			PicMed    string `json:"pic_med"`
			IsTag     bool   `json:"is_tag"`
			DnaColor  string `json:"dna_color"`
			Followers int    `json:"followers"`
		} `json:"values"`
	} `json:"interests"`
	Rating          interface{} `json:"rating"`
	NotificationBar bool        `json:"notification_bar"`
	Channel         interface{} `json:"channel"`
	Discoverer      interface{} `json:"discoverer"`
	DiscovererID    int         `json:"discoverer_id"`
	Friend          interface{} `json:"friend"`
	Sponsored       bool        `json:"sponsored"`
	Thumbnail       string      `json:"thumbnail"`
	ThumbnailB      string      `json:"thumbnail_b"`
	Event           interface{} `json:"event"`
	ListPosition    int         `json:"list_position"`
	Created         string      `json:"created"`
	Context         interface{} `json:"context"`
	Framebreak      interface{} `json:"framebreak"`
	Sandbox         bool        `json:"sandbox"`
	NumComments     interface{} `json:"numComments"`
	TrackingCode    interface{} `json:"tracking_code"`
	ThumbupDate     string      `json:"thumbup_date"`
	GroupKey        string      `json:"group_key"`
}

type AllLikes struct {
	Likes struct {
		Values []ALike `json:"values"`
	} `json:"likes"`
}

func GetLikes(url string) (Likes, error) {
	nilLikes := Likes{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nilLikes, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nilLikes, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nilLikes, err
	}
	var suLikes Likes
	json.Unmarshal(bodyBytes, &suLikes)
	return suLikes, nil
}

func GetUserInfo(url string) (UserInfo, error) {
	nilUserInfo := UserInfo{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nilUserInfo, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nilUserInfo, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nilUserInfo, err
	}

	var suUserInfo UserInfo
	json.Unmarshal(bodyBytes, &suUserInfo)
	return suUserInfo, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments!")
		fmt.Println("Example Usage: surescue su_username")
		os.Exit(1)
	}
	userName := os.Args[1]
	userURL := fmt.Sprintf("https://www.stumbleupon.com/api/v2_0/user/%s?version=2", userName)
	suUserData, err := GetUserInfo(userURL)

	if err != nil {
		fmt.Println("User could not be found!")
		os.Exit(1)
	}
	userID := suUserData.User.Userid
	likeCnt, _ := strconv.Atoi(suUserData.User.LikesCount)
	fmt.Println("User ID: ", userID)
	fmt.Println("Like Count: ", likeCnt)

	likesPerReq := 500
	numOfReqs := likeCnt/likesPerReq + 1
	allLikes := make([]ALike, 0)

	for i := 0; i < numOfReqs; i++ {
		baseUrl := "https://www.stumbleupon.com/api/v2_0/history/"
		likesUrl := fmt.Sprintf("%s%d/likes/all?offset=%d&limit=%d&userid=%d&listId=likes", baseUrl, userID, i*likesPerReq, likesPerReq, userID)
		suLikes, err := GetLikes(likesUrl)
		if err == nil {
			allLikes = append(allLikes, suLikes.Likes.Values...)
		}
	}

	allLikesWrapper := AllLikes{}
	allLikesWrapper.Likes.Values = allLikes

	allLikesWrapperJSON, _ := json.Marshal(allLikesWrapper)
	fileName := userName + ".json"
	err = ioutil.WriteFile(fileName, allLikesWrapperJSON, 0644)
}

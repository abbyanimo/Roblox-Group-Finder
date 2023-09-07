package main

// Made By coolkosmos On Discord ID: 991425843911987252. Take Any Code From This If You Wish! Really Shitty Code Tho So GL!
import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	SYNTAX        = "Syntax: go run %s <script>"
	GROUP_API_URL = "https://groups.roproxy.com/v2/groups?groupIds="
)

type GroupInfo struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Owner              Owner  `json:"owner"`
	Created            string `json:"created"`
	HasVerifiedBadge   bool   `json:"hasVerifiedBadge"`
	IsLocked           bool   `json:"isLocked"`
	publicEntryAllowed bool   `json:"publicEntryAllowed"`
	memberCount        bool   `json:"memberCount"`
}

type Owner struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

type GroupDataResponse struct {
	Data []GroupInfo `json:"data"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		time.Sleep(1500 * time.Millisecond)
		selectedRange := getDefaultRange()
		groupIDs := generateRandomGroupIDs(selectedRange, 100)
		var groupIDsStrs []string
		for _, groupID := range groupIDs {
			groupIDsStrs = append(groupIDsStrs, strconv.Itoa(groupID))
		}
		groupIDsStr := strings.Join(groupIDsStrs, ",")

		// Fetch group data
		response, err := http.Get(GROUP_API_URL + groupIDsStr)
		if err != nil {
			fmt.Println("Failed to fetch group information:", err)
			continue
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			fmt.Println("Failed to fetch group information (status code not OK)")
			continue
		}

		var groupDataResponse GroupDataResponse
		if err := json.NewDecoder(response.Body).Decode(&groupDataResponse); err != nil {
			fmt.Println("Failed to decode JSON:", err)
			continue
		}

		for _, groupInfo := range groupDataResponse.Data {
			groupID := groupInfo.ID
			groupName := groupInfo.Name
			ownerID := groupInfo.Owner.ID

			if ownerID == 0 {
				detailedGroupData, err := getGroupData2(groupID)
				if err != nil {
					fmt.Println("Failed to fetch detailed group information:", err)
					continue
				}

				if detailedGroupData != nil && (detailedGroupData.IsLocked || !detailedGroupData.publicEntryAllowed) {
					fmt.Printf("Group ID: %d is locked\n", groupID)
				} else {
					fmt.Printf("Group ID: %d has no owner\n", groupID)
					webhookURL := "" // Your Webhook Here :D
					embed := map[string]interface{}{
						"title":       "Click Here!",
						"description": fmt.Sprintf("Group Name: %s\nMember Count: %d", groupName, detailedGroupData.memberCount),
						"url":         "https://www.roblox.com/groups/" + strconv.Itoa(groupID),
						"color":       2856686,
						"author": map[string]interface{}{
							"name": "Normal Group Found",
						},
						"footer": map[string]interface{}{
							"text": "Kosmos Go Language Finder",
						},
					}
					payload := map[string]interface{}{
						"content": "Link: https://www.roblox.com/groups/" + strconv.Itoa(groupID),
						"embeds":  []interface{}{embed},
					}
					jsonPayload, err := json.Marshal(payload)
					if err != nil {
						fmt.Println("Failed to marshal JSON payload:", err)
						continue
					}
					_, err = http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
					if err != nil {
						fmt.Println("Failed to send Discord webhook:", err)
						continue
					}
				}
			} else {
				fmt.Printf("Group ID: %d Owner: Yes\n", groupID)
			}
		}
	}
}

func getDefaultRange() []int {
	rand.Seed(time.Now().UnixNano())
	ranges := [][]int{
		{2000000, 3000000},
		{3000000, 4000000},
		{4000000, 5000000},
		{5000000, 6000000},
		{6000000, 7000000},
		{7000000, 8000000},
		{8000000, 9000000},
		{9000000, 10000000},
		{10000000, 11000000},
		{11000000, 12000000},
		{12000000, 13000000},
		{13000000, 14000000},
		{14000000, 15000000},
		{15000000, 16000000},
		{16000000, 16500000},
		{16500000, 17000000},
		{17000000, 17500000},
		{17500000, 18500000},
		{30000000, 30100000},
		{30000000, 30200000},
		{30000000, 30300000},
		{30000000, 30400000},
		{30000000, 30500000},
		{30000000, 30600000},
		{30000000, 30700000},
		{33000000, 33500000},
	}
	randomIndex := rand.Intn(len(ranges))
	return ranges[randomIndex]
}

func generateRandomGroupIDs(selectedRange []int, count int) []int {
	groupIDs := make([]int, count)
	for i := 0; i < count; i++ {
		groupIDs[i] = rand.Intn(selectedRange[1]-selectedRange[0]+1) + selectedRange[0]
	}
	return groupIDs
}

func getGroupData(groupID int) (*GroupInfo, error) {
	response, err := http.Get(fmt.Sprintf("http://groups.roproxy.com/v2/groups?groupIds=%d", groupID))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch group information (status code not OK)")
	}

	var groupDataResponse GroupDataResponse
	if err := json.NewDecoder(response.Body).Decode(&groupDataResponse); err != nil {
		return nil, err
	}

	if len(groupDataResponse.Data) == 0 {
		return nil, fmt.Errorf("No data found for group ID %d", groupID)
	}

	return &groupDataResponse.Data[0], nil
}

func getGroupData2(groupID int) (*GroupInfo, error) {
	response, err := http.Get(fmt.Sprintf("http://groups.roproxy.com/v1/groups/%d", groupID))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch group information (status code not OK)")
	}

	var groupInfo GroupInfo
	if err := json.NewDecoder(response.Body).Decode(&groupInfo); err != nil {
		return nil, err
	}

	return &groupInfo, nil
}

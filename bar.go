package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type DowntimeUsersInput struct {
	Start        int64
	End          int64
	DowntimeName string
}

type DowntimeUsersOutput struct {
	DowntimeUsers []DowntimeUser
	Result        string
}

type DowntimeUser struct {
	Name  string
	Value time.Duration
}

type DowntimesInput struct {
	Start int64
	End   int64
	Data  string
}

type DowntimesOutput struct {
	Downtimes []Downtime
	Result    string
}

type Downtime struct {
	Name  string
	Value time.Duration
}

func getDowntimes(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data DowntimesInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		logError("Bar", "Error parsing data from page: "+err.Error())
		var responseData DowntimesOutput
		responseData.Result = "nok"
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}

	startTime := time.Unix(data.Start/1000, 0)
	endTime := time.Unix(data.End/1000, 0)
	logInfo("Bar", "User asked data from "+startTime.String()+" to "+endTime.String() + " for " + data.Data)
	var downtimes []Downtime
	var firstData Downtime
	firstData.Name = "Smoking"
	firstData.Value = (1*time.Hour + 27*time.Minute) / 1000000000
	downtimes = append(downtimes, firstData)
	var secondData Downtime
	secondData.Name = "Resting"
	secondData.Value = 330 * time.Minute / 1000000000
	downtimes = append(downtimes, secondData)
	var thirdData Downtime
	thirdData.Name = "Sleeping"
	thirdData.Value = 67 * time.Minute / 1000000000
	downtimes = append(downtimes, thirdData)
	logInfo("Bar", "Sending data")
	var responseData DowntimesOutput
	responseData.Result = "ok"
	responseData.Downtimes = downtimes
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	return
}

func getDowntimeUsers(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data DowntimeUsersInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		logError("Bar", "Error parsing data from page: "+err.Error())
		var responseData DowntimeUsersOutput
		responseData.Result = "nok"
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}
	logInfo("Bar", "User asked for "+data.DowntimeName)
	var downtimeUsers []DowntimeUser
	var firstData DowntimeUser
	firstData.Name = "Petr Jahoda"
	firstData.Value = (2*time.Hour + 27*time.Minute) / 1000000000
	downtimeUsers = append(downtimeUsers, firstData)
	var secondData DowntimeUser
	secondData.Name = "Karel Gott"
	secondData.Value = 98 * time.Minute / 1000000000
	downtimeUsers = append(downtimeUsers, secondData)
	var thirdData DowntimeUser
	thirdData.Name = "James Bond"
	thirdData.Value = 33 * time.Minute / 1000000000
	downtimeUsers = append(downtimeUsers, thirdData)
	logInfo("Bar", "Sending data")
	var responseData DowntimeUsersOutput
	responseData.Result = "ok"
	responseData.DowntimeUsers = downtimeUsers
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	return
}


func getUsersDowntimes(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var data DowntimesInput
	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		logError("Bar", "Error parsing data from page: "+err.Error())
		var responseData DowntimesOutput
		responseData.Result = "nok"
		writer.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writer).Encode(responseData)
		return
	}

	startTime := time.Unix(data.Start/1000, 0)
	endTime := time.Unix(data.End/1000, 0)
	logInfo("Bar", "User asked data from "+startTime.String()+" to "+endTime.String() + " for " + data.Data)
	var downtimes []Downtime
	var firstData Downtime
	firstData.Name = "Hiking"
	firstData.Value = (3*time.Hour + 27*time.Minute) / 1000000000
	downtimes = append(downtimes, firstData)
	var secondData Downtime
	secondData.Name = "Killing"
	secondData.Value = 42 * time.Minute / 1000000000
	downtimes = append(downtimes, secondData)
	var thirdData Downtime
	thirdData.Name = "Hoping"
	thirdData.Value = 42 * time.Minute / 1000000000
	downtimes = append(downtimes, thirdData)
	logInfo("Bar", "Sending data")
	var responseData DowntimesOutput
	responseData.Result = "ok"
	responseData.Downtimes = downtimes
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(responseData)
	return
}

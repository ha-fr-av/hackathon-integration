package main

import "fmt"

const (
	jwt = "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ6ZXJvLXRlc3RpbmcuYXZpdmEuY28udWsiLCJleHAiOjE3MTIxODMxNTMsImlhdCI6MTcxMjEzOTk1MywibmJmIjoxNzEyMTM5OTUzLCJzdWIiOiJAZW5jOk4xb3Npc1psWlpHdFhpUkZFM0VlYUo5TzJPcS84YWRoR1hPUU1MTVhOWHc3elVHRm1pNUxXcEpWY1BHMmkrM1Y4dkpmZDJXYkYzM0pHYWlqTDRSYzVYbkZCUXE2TVJ6aTZvWUU1U0J1WHJQd0ovY2tReWJsSjF4aDI2ZUluRDg3VHk1L3Y2cnUwRlNwTEpBWGtlWUpuelUrYUh3RlIwdFY4ZWFoZUE3NnVSRzZLYlFjVUhPUzRPQzZNQ1lnWnZTb1J4R04xdEV6ZmpBSVA1M1lxK2hLZTJQOGpFUlRrbFpDY2hPSjBBT2tFOXpLSEZXV1NHSk5WQTI5Y2p2QUQzMkJ4UElTNnJVRHNpTTVsYjA4ZFoxTThSZURCY0lkWGtaNWtRaTd6RE01VDJhQjZnWVV6M1RETjlFRTQrcXRqU2JHR0tSdDkzQTNsU2FXVXhFUzg4T0t1OE1VZTdKTnp4cFhmK2szazhxRnlWbmVTRXRGMjhXWTJneW16a3NmWENBbjlTbUJRQkpqMkVHdTlUbjg2VFZxV3JzeVlZaHhaK1hDUnBTakZqME44TEMyLzl3UW82aU84UnlsYkc5ZW5YeFZGdFJ5Z1p6NE5qeFR0SGVJRzBSVzFDZEZwNnVwT0hPOExjd1A1TWdsYkpBQmt4djlBc1hWamRVd2hvdnFXUTNLSlBFdHZ1dHdvamhqeTY4T1lSOHVRL25xTC9nRFo4bjRpTStUZkNCcmdoZmZzVFl5VXNPc1Q5dUl3cEpkYWM2NTh0TUR4WDRJUWc0VTM2cGlBYzFCVW9YWC93ZDZDaEJHN0xVdWtMbVhWZERONHFDRksvb3paaHByUmZudDdTTHpFWnFCb0NGa2g4K0FiYTB2M256RHV0SHgvODRqKzQweTNxUVh3MnR2MUJRPSJ9.muk3flJElqkQl8t3PHV-uhFcasdIkG6BFrUrUedanzIbC3BFDnRCzPgB2AV81Gg3WCJRBuJmYREY2P7dLfjcOzmf4fJpMssnwhripQT2bRMWmZPIMXgNHALi57s0Cyw093tYehEGtXU4uNviwlrjRp-p3TYgOWKDD29VpDNh3FRu0zpqqOg9Ep67Kp8_WJ-9w17tIcSEznjAXi_2ybm6IePOVZYdYk4CN8w9PCUP7jALhFMSQQ9CNFdlHXkENFJ10dIv8pifyGhCA71qqNZc15cQoz6RrYR5BfyxcVaIEbHiDPc5u4VJs6Y8NCwfrGHfxoVSJ76nWSQ1ENReNiGkUU4xWQqWiiF1o5kVpodadt0jkgaBxVWDh9YsFGhElmRDqxmAa6ZYRyemV3OjXRFfiJTnbwbyeteZc1cv7_RFKvn80VEJ3NdDxx_1BoHMbzkZJXtB6MLy36P8SAFD8ySN7oAmCwENuGI5tXeP-M_Zfqjdw2TL7g_Nsoim5AdKZ6tRWCzWUE7NmIqTaiyVV5I7K5t8mAGrzhVY-zKKQd5NqzEjfFPmw7HkQ0NMbsFWW6jIusJMhkl-d2Ud6KEV6Te7jhFxArwduREdMCNcfLo0fUkSt3zLBZ2vdp1kNy90ULdJoGkgTDCn8eRwF5S_NBs3kBi48SLOSV5hbTm9ZqNrv0Y"
	dob = "1989-03-23"
	// hardcoded items for initial dev
	policyId  = "335d447a-929f-45c5-8b9b-56e3fde540f9"
	insuredId = "2d1c5f12-dad9-4cd4-aa74-b8eaeb56d078"
	domain    = "https://m3yktt0wkj.execute-api.eu-west-1.amazonaws.com"
)

var payload = map[string]any{
	"coverDetails": map[string]any{
		"typeOfUse":         "sdpc",
		"annualMileage":     10000,
		"locationOvernight": "4",
		"locationDaytime":   "VKD1",
	},
	"driver": map[string]any{
		"owner":            true,
		"registeredKeeper": true,
	},
	"vehicle": map[string]any{
		"tracker":            false,
		"alarmImmobilizer":   "#F",
		"dateOfPurchase":     "2021-11-21T00:00:00.000Z",
		"value":              30000,
		"yearOfManufacture":  "2020",
		"registrationNumber": "AP65GVW",
	},
	"effectiveFrom":   "2024-03-11T11:03:19.226Z",
	"transactionType": "edit-vehicle",
}

func arrange(e map[string]any) (ActionParams, error) {
	headers := map[string]string{
		"Authorization":    fmt.Sprintf("Bearer %s", jwt),
		"policyholder-dob": dob,
		"Host":             "localhost",
	}

	return ActionParams{
		Endpoint: getEndpoint(domain, policyId, insuredId),
		Headers:  headers,
		Payload:  payload,
	}, nil
}

/**
* build the endpoint to get an mta quote
**/
func getEndpoint(domain, policyId, insuredId string) string {

	//  "https://m3yktt0wkj.execute-api.eu-west-1.amazonaws.com/prod/mta/:policyId/change-vehicle/:insuredId"
	return fmt.Sprintf("%s/prod/mta/%s/change-vehicle/%s", domain, policyId, insuredId)

}

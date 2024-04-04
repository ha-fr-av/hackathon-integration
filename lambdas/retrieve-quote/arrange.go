package main

import (
	"fmt"
)

// const (
// 	jwt = "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJ6ZXJvLXRlc3RpbmcuYXZpdmEuY28udWsiLCJleHAiOjE3MTIyNjQwOTQsImlhdCI6MTcxMjIyMDg5NCwibmJmIjoxNzEyMjIwODk0LCJzdWIiOiJAZW5jOkt3NFJNSmtlRDRCbXNBazJzZ2txamZseDI5bXlCYXUrSk8zbDNYMFd6YStTSjgwS21RVnhLbktnd2xPcVBtaDltWEZMNE9rRkxTR2ZxUVdnNU5KR2dLYWlJYTV3eU5OYWEvbHdNZDBoYXNvRkxrUWdsT2F5TlhDOWZlaDM2c1I0V0R2cTdxMlN2clVHTWZtQVJrZnY3VnJITk0xWUdnS3ZTMXphblBWU3FpOUZJSVNhVjhxM05WWEZYYkhZcmVrRzA4bGJDaXdFcnQ3d1BYd1dObWFRSGFqOFlhWlVmTHhyN2JreTQzT0x2UllZNTEza0VIK2NJRnQ5MzAvUUFQajNuc1FTM0p3aG5qQ2QrenUzQ21HbnNZTjhwVDlaVTZwTDJSTUZycVp5Q2NXeXJMR09xd0pwQWtsb0JpajU3MmNKMmV2MFRjU1JBdHNaVzVkbFQwNXdFVkRwQ29hZ0I1THFBZ0RLY0NneWx6cDdpQ0RJZUlpNmRnK3pVdG5yYkhNSUtYZVhqM05ieWhTcDlYVDRLc2VsVEpMWlNrVC8yUnl6dkNDSEx3RGQwTktGMkpHaS9FalhUNFBtSkJqWFJiK1FJNi9aQ2M1NnEyM011QWZpOS9HMEZ4MUpLUlU4T0lkTEJpNWhQOS9jaXp6M1pIVVhEbHRTK0MxVGVnZzBDRjhjaVZWanY1U2xwZ051RVBKeEFBcTVKNHlyd3lSb2N1WWVPVlVIbm95TENyTVF2SHdRRzRMSGFjSW1TYW9zVmM2WFM2dFl3azB0cHdYeGIrWk9aMEQ5NERHUTNBY1M2MzVackp5enZ2VDlRWVZPK3NEVmN2Z09xdUxyY01mQjFVL0JsWjRabXZPbUY2amRHd2dTY1BnUHp1MDE0SmFQOVBCajlJOUNDd1I1TVEwPSJ9.wLBE-DLUCQOhAwzTbw9HSAoKQg2G9YuZDk1sLd0X0TCSKxKwUiM31JEBv4M-Y7t-NKDcxy2qUJqDvze3T8YlBXGBgdREj3uSLqOCfFHwPdjvjYopletap42hfUskcUg-vHm7AGBYrTOWFTFZ80mstwj0O9LqEZ0Qj_9gndwURVOBmbaL-9R0wcr-Vjh_2PGf-ml9d_0DNXYyWnT8IN884wEcax5hjZ4g8olZ___9JIv8IcHOho591sktKVIOuSW05RzZTD2yjs40inw6hZi7h9ICjZI-WCwOPVXmtz9VLAyVHhI--vJffOn2BhjnfCvNcnmZ7TevdwPrKeARGh-Ae6K9oTAyRpzVenvcaqDszd5_eCqD6XpXFZOqS_HtcJB5XKrGvpRf-c7gK0bY5U6C1OoTdWkiGw4ddDFDGtZjJ2027tP-hb_RLSa0fndfhfsK9vOw_CyM1Q2nGPQ3u00bNohglQPrPSDvRfu4BNZhI8y3Q1kAkBkutwkpK3kHQypAqsa_YUlZQGoqnXE-_v6uVDnacRudVLKEQvbE63KEBqbnCJc36wFiEyh3JkYw8e894i-mdZ4i86qn8c0TFP5qBcsBK9Hs0cPQO4JE9QwM_-PGZQUp-vjkUz8Wb55T1G7H9-6V6vM83hFZayI5t2nxzPZzgTzsN5Vvt4A4ZQCmAs4"
// 	dob = "1989-03-23"
// 	// hardcoded items for initial dev
// 	policyId  = "335d447a-929f-45c5-8b9b-56e3fde540f9"
// 	insuredId = "2d1c5f12-dad9-4cd4-aa74-b8eaeb56d078"
// 	domain    = "https://m3yktt0wkj.execute-api.eu-west-1.amazonaws.com"
// )

func arrange(e event) (ActionParams, error) {
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %s", e.UserInfo.JWT)
	headers["policyholder-dob"] = e.UserInfo.Dob
	headers["Host"] = "localhost"

	return ActionParams{
		Endpoint: getEndpoint(e.Host.Domain, e.Payload.QuoteID),
		Headers:  headers,
		QuoteID:  e.Payload.QuoteID,
	}, nil
}

/**
* build the endpoint to get an mta quote
**/
func getEndpoint(domain, quoteId string) string {
	return fmt.Sprintf("%s/prod/quote/%s", domain, quoteId)

}

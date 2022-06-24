package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/securityhub"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	fmt.Println()
	myFigure := figure.NewFigure("AWS    SECOPS", "isometric1", true)
	myFigure.Print()
	fmt.Printf("\n\n")
	awsRegion, awsKey, awsSecret := "", "", ""
	awsRegion = os.Getenv("AWS_REGION")
	awsKey = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecret = os.Getenv("AWS_SECRET_ACCESS_KEY")
	var awsSession *session.Session
	if awsRegion != "" && awsKey != "" && awsSecret != "" {
		//Read region and keys from  env vars (mainly for docker container mode)
		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String(awsRegion),
			Credentials: credentials.NewStaticCredentials(awsKey, awsSecret, "")},
		)
		if err != nil {
			fmt.Println("Got error initializing session:")
			fmt.Println(err.Error())
			return
		}
		awsSession = sess
	} else {
		//Read region and keys from  ~/.aws/config and  ~/.aws/credentials
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		awsSession = sess
	}
	lowVuln := 0
	mediumVuln := 0
	highVuln := 0
	//Creating new security hub service session
	svc := securityhub.New(awsSession)
	initialToken := ""
	orderField := "Title"
	sortOrder := "ASC"
	maxResultsPerCall := int64(100)
	filter := securityhub.SortCriterion{Field: &orderField, SortOrder: &sortOrder}
	var filterList []*securityhub.SortCriterion = []*securityhub.SortCriterion{&filter} //Order results by title asc
	for {
		input := &securityhub.GetFindingsInput{NextToken: &initialToken,
			MaxResults:   &maxResultsPerCall,
			SortCriteria: filterList}
		resp, err := svc.GetFindings(input)
		if err != nil {
			fmt.Println("Got error calling GetFindings:")
			fmt.Println(err.Error())
			return
		}
		for _, finding := range resp.Findings {
			switch gravity := *finding.Severity.Original; gravity {
			case "LOW":
				lowVuln += 1
				colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, gravity)
				fmt.Println("GRAVITY: ", colored)
			case "MEDIUM":
				mediumVuln += 1
				colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 33, gravity)
				fmt.Println("GRAVITY: ", colored)
			case "HIGH":
				highVuln += 1
				colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, gravity)
				fmt.Println("GRAVITY: ", colored)
			default:
				lowVuln += 1
				colored := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 32, "LOW")
				fmt.Println("GRAVITY: ", colored)
			}
			fmt.Printf("[%v]\n%v%v ===> %v\n\n", *finding.Title, *finding.Description,
				*finding.Remediation.Recommendation.Text, *finding.Remediation.Recommendation.Url)
		}
		if resp.NextToken != nil {
			initialToken = *resp.NextToken
		} else {
			initialToken = ""
		}
		if initialToken == "" {
			break
		}
	}
	fmt.Printf("\n\nVULNERABILITY RECAP: FOUND %v LOW, %v MEDIUM AND %v HIGH\n\n", lowVuln, mediumVuln, highVuln)
}

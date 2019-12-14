package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice"
)

func main() {
	ses := session.Must(session.NewSession())
	svc := pinpointsmsvoice.New(ses)
	svc := pinpointsmsvoice.New(ses, aws.NewConfig().WithRegion("us-east-1"))
}
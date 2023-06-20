package app_main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AppMainStackProps struct {
	awscdk.StackProps
}

func NewAppMainStack(scope constructs.Construct, id string, props *AppMainStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	awssns.NewTopic(stack, jsii.String("MyTopic"), &awssns.TopicProps{
		DisplayName: jsii.String("MyTopic"),
	})

	return stack
}

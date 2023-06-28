package pipeline_main

import (
	"cdk-go-lab/config"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewSNSNoticeStack(scope constructs.Construct, id string, props *awscdk.StackProps) awscdk.Stack {

	stack := awscdk.NewStack(scope, &id, props)

	topic := awssns.NewTopic(stack, jsii.String("NoticeTopic"), &awssns.TopicProps{
		DisplayName: jsii.String("NoticeTopic"),
	})

	awssns.NewSubscription(stack, jsii.String("NoticeSubscription"), &awssns.SubscriptionProps{
		Topic:    topic,
		Protocol: awssns.SubscriptionProtocol_EMAIL,
		Endpoint: jsii.String(config.ManualApprovalEmail),
	})

	return stack
}

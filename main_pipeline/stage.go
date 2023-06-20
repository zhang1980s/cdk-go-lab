package main_pipeline

import (
	"cdk-go-lab/main_app"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type MainAppStageProps struct {
	awscdk.StageProps
}

func NewMyAppStage(scope constructs.Construct, id string, props *MainAppStageProps) awscdk.Stage {
	var sprops awscdk.StageProps

	if props != nil {
		sprops = props.StageProps
	}

	stage := awscdk.NewStage(scope, &id, &sprops)

	main_app.NewMyAppStack(stage, "MyAppStack", nil)

	return stage
}

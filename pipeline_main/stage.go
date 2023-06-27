package pipeline_main

import (
	"cdk-go-lab/app_main"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type PipelineMainStageProps struct {
	awscdk.StageProps
}

type PipelineManualApprovalStageProps struct {
	awscdk.StageProps
}

func NewPipelineMainStage(scope constructs.Construct, id string, props *PipelineMainStageProps) awscdk.Stage {
	var sprops awscdk.StageProps

	if props != nil {
		sprops = props.StageProps
	}

	stage := awscdk.NewStage(scope, &id, &sprops)

	app_main.NewAppMainStack(stage, "MyAppMainStack", nil)

	return stage
}

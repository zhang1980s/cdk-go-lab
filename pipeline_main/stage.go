package pipeline_main

import (
	"cdk-go-lab/app_main"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/pipelines"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
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

func NewManualApprovalStage(scope constructs.Construct, id string, props *PipelineMainStageProps) awscdk.Stage {

	var sprops awscdk.StageProps

	if props != nil {
		sprops = props.StageProps
	}

	stage := awscdk.NewStage(scope, &id, &sprops)

	pipelines.NewManualApprovalStep(jsii.String("ManualApproval"), nil)

	return stage
}

package main

import (
	// "github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	sns "github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	sfn "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	tasks "github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks"
	awslambdago "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type HackathonIntegrationStackProps struct {
	awscdk.StackProps
}

func NewHackathonIntegrationStack(scope constructs.Construct, id string, props *HackathonIntegrationStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	pqfunc := awslambdago.NewGoFunction(stack, jsii.String("processQuoteLambdaFunction"), &awslambdago.GoFunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Architecture: awslambda.Architecture_ARM_64(),
		Entry:        jsii.String("../lambdas/process-quote"),
		Timeout:      awscdk.Duration_Minutes(jsii.Number(1)),
		Tracing:      awslambda.Tracing_ACTIVE,
	})

	pqtask := tasks.NewLambdaInvoke(stack, jsii.String("processQuoteLambdaTask"), &tasks.LambdaInvokeProps{
		LambdaFunction: pqfunc,
		TaskTimeout:    sfn.Timeout_Duration(awscdk.Duration_Seconds(jsii.Number(65))),
	})

	vqfunc := awslambdago.NewGoFunction(stack, jsii.String("verifyQuoteLambdaFunction"), &awslambdago.GoFunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Architecture: awslambda.Architecture_ARM_64(),
		Entry:        jsii.String("../lambdas/retrieve-quote"),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(15)),
		Tracing:      awslambda.Tracing_ACTIVE,
	})

	vqtask := tasks.NewLambdaInvoke(stack, jsii.String("verifyQuoteTask"), &tasks.LambdaInvokeProps{
		LambdaFunction: vqfunc,
		TaskTimeout:    sfn.Timeout_Duration(awscdk.Duration_Seconds(jsii.Number(20))),
		InputPath:      jsii.String("$.Payload.data"),
	})

	pqchoice := sfn.NewChoice(stack, jsii.String("Quote Processed Successfully?"), nil)
	pqcondition := sfn.Condition_IsNotString(jsii.String("$.Payload.error"))

	vqchoice := sfn.NewChoice(stack, jsii.String("Quote Verified Successfully?"), nil)
	vqcondition := sfn.Condition_IsNotString(jsii.String("$.Payload.error"))

	topic := sns.NewTopic(stack, jsii.String("HelloTopic"), nil)

	publishMessage := tasks.NewSnsPublish(stack, jsii.String("Alert slack on Error"), &tasks.SnsPublishProps{
		Topic:   topic,
		Message: sfn.TaskInput_FromObject(&map[string]any{"errormessage.$": "$.Payload.error", "executionId.$": "$$.Execution.Id"}),
		// Message: sfn.TaskInput_FromJsonPathAt(jsii.String("States.Format('A job submitted through Step Functions failed for document id {}', $.Payload.hello)")),

		ResultPath: jsii.String("$.sns"),
	})

	successState := sfn.NewPass(stack, jsii.String("SuccessState"), nil)

	v := vqtask.Next(vqchoice.When(vqcondition, successState, nil).Otherwise(publishMessage))

	definition := pqtask.Next(pqchoice.When(pqcondition, v, nil).Otherwise(publishMessage))

	sfn.NewStateMachine(stack, jsii.String("TestStateMachine"), &sfn.StateMachineProps{
		Definition: definition,
		Timeout:    awscdk.Duration_Minutes(jsii.Number(10)),
		Comment:    jsii.String("Test State Machine"),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewHackathonIntegrationStack(app, "HackathonIntegrationStack", &HackathonIntegrationStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}

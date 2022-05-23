package main

import (
	"cdk.tf/go/stack/generated/hashicorp/aws"
	"cdk.tf/go/stack/generated/hashicorp/aws/eks"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// The code that defines your stack goes here
	aws.NewAwsProvider(stack, jsii.String("aws"), &aws.AwsProviderConfig{
		Region: jsii.String("us-east-1"),
	})

	tags := make(map[string]string)
	tags["Name"] = "mkorejo-private"

	cluster_config := eks.EksClusterConfig{
		Name:                   jsii.String("mkorejo-private"),
		EnabledClusterLogTypes: jsii.Strings("api", "audit", "authenticator", "controllerManager", "scheduler"),
		RoleArn:                jsii.String(""),
		Version:                jsii.String("1.22"),
		EncryptionConfig: &eks.EksClusterEncryptionConfig{
			Provider: &eks.EksClusterEncryptionConfigProvider{
				KeyArn: jsii.String("arn:aws:kms:us-east-1:853973692277:key/a30af577-f854-4777-b8d9-f2f3606e66a8"),
			},
			Resources: jsii.Strings("secrets"),
		},
		VpcConfig: &eks.EksClusterVpcConfig{
			EndpointPrivateAccess: jsii.Bool(true),
			EndpointPublicAccess:  jsii.Bool(false),
			SecurityGroupIds:      jsii.Strings("sg-066a69fbf6d95f5e2"),
			SubnetIds:             jsii.Strings("subnet-009f9205773eaf01e", "subnet-0328f4a06d9e4ca64"),
		},
	}

	ng1_config := eks.EksNodeGroupConfig{
		ClusterName:   jsii.String("mkorejo-private"),
		NodeRoleArn:   jsii.String(""),
		NodeGroupName: jsii.String("ng1"),
		AmiType:       jsii.String("AL2_x86_64"),
		InstanceTypes: jsii.Strings("t3.medium"),
		SubnetIds:     jsii.Strings("subnet-009f9205773eaf01e", "subnet-0328f4a06d9e4ca64"),
		RemoteAccess: &eks.EksNodeGroupRemoteAccess{
			Ec2SshKey: jsii.String("muradkorejo"),
		},
		ScalingConfig: &eks.EksNodeGroupScalingConfig{
			DesiredSize: jsii.Number(2),
			MaxSize:     jsii.Number(5),
			MinSize:     jsii.Number(2),
		},
	}

	eks.NewEksCluster(stack, jsii.String("eks_cluster"), &cluster_config)
	eks.NewEksNodeGroup(stack, jsii.String("eks_ng1"), &ng1_config)

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "learn-cdktf-go")

	app.Synth()
}

package outputs

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubernetes/mongodbkubernetes"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/commons/kubernetes"
	"github.com/plantoncloud/stack-job-runner-golang-sdk/pkg/automationapi/autoapistackoutput"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

const (
	Namespace               = "namespace"
	Service                 = "service"
	KubePortForwardCommand  = "port-forward-command"
	KubeEndpoint            = "kube-endpoint"
	IngressExternalHostname = "ingress-external-hostname"
	IngressInternalHostname = "ingress-internal-hostname"
	RootUsername            = "root-username"
	RootPasswordSecretName  = "root-password-secret-name"
	RootPasswordSecretKey   = "root-password-secret-key"
)

func PulumiOutputsToStackOutputsConverter(pulumiOutputs auto.OutputMap,
	input *mongodbkubernetes.MongodbKubernetesStackInput) *mongodbkubernetes.MongodbKubernetesStackOutputs {
	return &mongodbkubernetes.MongodbKubernetesStackOutputs{
		Namespace:          autoapistackoutput.GetVal(pulumiOutputs, Namespace),
		KubeEndpoint:       autoapistackoutput.GetVal(pulumiOutputs, KubeEndpoint),
		Service:            autoapistackoutput.GetVal(pulumiOutputs, Service),
		PortForwardCommand: autoapistackoutput.GetVal(pulumiOutputs, KubePortForwardCommand),
		ExternalHostname:   autoapistackoutput.GetVal(pulumiOutputs, IngressExternalHostname),
		InternalHostname:   autoapistackoutput.GetVal(pulumiOutputs, IngressInternalHostname),
		Username:           autoapistackoutput.GetVal(pulumiOutputs, RootUsername),
		PasswordSecret: &kubernetes.KubernernetesSecretKey{
			Name: autoapistackoutput.GetVal(pulumiOutputs, RootPasswordSecretName),
			Key:  autoapistackoutput.GetVal(pulumiOutputs, RootPasswordSecretKey),
		},
	}
}

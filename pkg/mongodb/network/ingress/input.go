package ingress

import (
	mongodbcontextstate "github.com/plantoncloud/mongodb-kubernetes-pulumi-blueprint/pkg/mongodb/contextstate"
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/kubecluster/enums/kubernetesworkloadingresstype"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type input struct {
	ingressType kubernetesworkloadingresstype.KubernetesWorkloadIngressType
}

func extractInput(ctx *pulumi.Context) *input {
	var ctxState = ctx.Value(mongodbcontextstate.Key).(mongodbcontextstate.ContextState)

	return &input{
		ingressType: ctxState.Spec.IngressType,
	}
}

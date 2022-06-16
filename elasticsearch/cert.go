package elasticsearch

import (
	"temporal-deployment-pulumi-go/shared"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewElasticsearchCert creates a new elastic search cert
func NewElasticsearchCert(ctx *pulumi.Context, args shared.Args) (*corev1.Secret, error) {
	return corev1.NewSecret(ctx, "elasticsearch-master-certs", &corev1.SecretArgs{})
}

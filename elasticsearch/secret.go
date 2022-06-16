package elasticsearch

import (
	"temporal-deployment-pulumi-go/shared"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// NewElasticsearchSecret creates a new elastic search secret
func NewElasticsearchSecret(ctx *pulumi.Context, args shared.Args) (*corev1.Secret, error) {
	var releaseName string = args.GetValue("release_name").(string)

	return corev1.NewSecret(ctx, releaseName+"-credentials", &corev1.SecretArgs{
		Type: pulumi.StringPtr("Opaque"),
		Data: pulumi.ToStringMap(
			map[string]string{
				"username": args.GetValue("username").(string), // "ZWxhc3RpYw==",
				"password": args.GetValue("password").(string), // "TXVkYXJAMTIz"
			},
		),
		Metadata: metav1.ObjectMetaPtr(&metav1.ObjectMetaArgs{
			Name: pulumi.StringPtr(releaseName + "-credentials"),
			Labels: pulumi.ToStringMap(
				map[string]string{
					"app":      releaseName,
					"chart":    releaseName,
					"heritage": releaseName,
					"release":  releaseName,
				},
			),
		}),
	})
}

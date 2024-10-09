// Code generated by "internal/generate/listpages/main.go -ListOps=DescribeListenerCertificates -InputPaginator=Marker -OutputPaginator=NextMarker -- list_listener_certificates_pages_gen.go"; DO NOT EDIT.

package elbv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
)

func describeListenerCertificatesPages(ctx context.Context, conn *elasticloadbalancingv2.Client, input *elasticloadbalancingv2.DescribeListenerCertificatesInput, fn func(*elasticloadbalancingv2.DescribeListenerCertificatesOutput, bool) bool) error {
	for {
		output, err := conn.DescribeListenerCertificates(ctx, input)
		if err != nil {
			return err
		}

		lastPage := aws.ToString(output.NextMarker) == ""
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.Marker = output.NextMarker
	}
	return nil
}

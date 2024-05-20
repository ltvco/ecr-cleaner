package clean

import (
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"testing"
)

func TestCheckImageNotInUse(t *testing.T) {
	type args struct {
		all    []*ImageInfo
		detail types.ImageDetail
	}
	digest := "sha256:b31dd6ba7d28a1559be39a88c292a1a8948491b118dafd3e8139065afe55690a"
	tests := []struct {
		name string
		args args
		want bool
	}{{
		name: "Match",
		args: args{
			all:    []*ImageInfo{{Digest: "b31dd6ba7d28a1559be39a88c292a1a8948491b118dafd3e8139065afe55690a"}},
			detail: types.ImageDetail{ImageDigest: &digest},
		},
		want: false,
	},
		{
			name: "NoMatch",
			args: args{
				all:    []*ImageInfo{{Digest: "13b7e62e8df80264dbb747995705a986aa530415763a6c58f84a3ca8af9a5bcd"}},
				detail: types.ImageDetail{ImageDigest: &digest},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckImageNotInUse(tt.args.all, tt.args.detail); got != tt.want {
				t.Errorf("CheckImageNotInUse() = %v, want %v", got, tt.want)
			}
		})
	}
}

//
//func TestCleanRepos(t *testing.T) {
//	type args struct {
//		untaggedOnly  bool
//		keepLastCount int
//		profile       string
//		region        string
//		dryRun        bool
//		verbose       bool
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := CleanRepos(tt.args.untaggedOnly, tt.args.keepLastCount, tt.args.profile, tt.args.region, tt.args.dryRun, tt.args.verbose); got != tt.want {
//				t.Errorf("CleanRepos() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestDeleteECRImage(t *testing.T) {
//	type args struct {
//		ctx    context.Context
//		client *ecr.Client
//		image  types.ImageDetail
//		dryRun bool
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := DeleteECRImage(tt.args.ctx, tt.args.client, tt.args.image, tt.args.dryRun); (err != nil) != tt.wantErr {
//				t.Errorf("DeleteECRImage() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestGetECRImage(t *testing.T) {
//	type args struct {
//		ctx    context.Context
//		client *ecr.Client
//		image  types.ImageDetail
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := GetECRImage(tt.args.ctx, tt.args.client, tt.args.image); (err != nil) != tt.wantErr {
//				t.Errorf("GetECRImage() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestGetECRImages(t *testing.T) {
//	type args struct {
//		client *ecr.Client
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    []*ecr.DescribeImagesOutput
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := GetECRImages(tt.args.client)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetECRImages() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetECRImages() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestGetPrometheusImagesFromProfile(t *testing.T) {
//	tests := []struct {
//		name    string
//		want    []*ImageInfo
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := GetPrometheusImagesFromProfile()
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetPrometheusImagesFromProfile() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetPrometheusImagesFromProfile() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestGetUnique(t *testing.T) {
//	type args struct {
//		all []*ImageInfo
//	}
//	tests := []struct {
//		name string
//		args args
//		want []*ImageInfo
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := GetUnique(tt.args.all); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetUnique() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

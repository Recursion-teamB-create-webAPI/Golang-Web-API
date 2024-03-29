package utils

import (
	"reflect"
	"testing"

	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/constants"
	"github.com/Recursion-teamB-create-webAPI/Golang-Web-API.git/pkg/structs"
)

func TestGetEnvData(t *testing.T) {
	type args struct {
		beforeLevel int
	}
	tests := []struct {
		name string
		args args
		want structs.Env
	}{
		{
			name: "Successful get env data",
			args: args{beforeLevel: constants.BeforeLevel3},
			want: structs.Env{
				SearchEngineId: "The values in the environment file were used for the test.",
				KeyFileName:    "The values in the environment file were used for the test.",
				CsePath:        "The values in the environment file were used for the test.",
				PortNumber:     "The values in the environment file were used for the test.",
				DatabaseName:   "The values in the environment file were used for the test.",
				MysqlUri:       "The values in the environment file were used for the test.",
			},
		},
		{
			name: "Failed get env data",
			args: args{beforeLevel: constants.BeforeLevel1},
			want: structs.Env{
				SearchEngineId: "",
				KeyFileName:    "",
				CsePath:        "",
				PortNumber:     "",
				DatabaseName:   "",
				MysqlUri:       "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnvData(tt.args.beforeLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnvData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInitImagesJson(t *testing.T) {
	type args struct {
		beforeLevel int
	}
	tests := []struct {
		name string
		args args
		want structs.InitImageItems
	}{
		{
			name: "Successful get initImages.json",
			args: args{beforeLevel: constants.BeforeLevel3},
			want: structs.InitImageItems{
				ImageItems: []structs.Items{
					{
						Item: "cat",
						ImageData: structs.ImageArray{
							Images: [constants.SearchResultNumber]string{
								"https://i.natgeofe.com/n/548467d8-c5f1-4551-9f58-6817a8d2c45e/NationalGeographic_2572187_square.jpg",
								"https://media.4-paws.org/5/b/4/b/5b4b5a91dd9443fa1785ee7fca66850e06dcc7f9/VIER%20PFOTEN_2019-12-13_209-2890x2000-1920x1329.jpg",
								"https://i.natgeofe.com/n/548467d8-c5f1-4551-9f58-6817a8d2c45e/NationalGeographic_2572187_16x9.jpg",
								"https://upload.wikimedia.org/wikipedia/commons/thumb/1/15/Cat_August_2010-4.jpg/1200px-Cat_August_2010-4.jpg",
								"https://i.natgeofe.com/n/548467d8-c5f1-4551-9f58-6817a8d2c45e/NationalGeographic_2572187_3x4.jpg",
								"https://bestfriends.org/sites/default/files/styles/hero_mobile/public/hero-dash/Asana3808_Dashboard_Standard.jpg?h=ebad9ecf\u0026itok=cWevo33k",
								"https://i.natgeofe.com/n/4cebbf38-5df4-4ed0-864a-4ebeb64d33a4/NationalGeographic_1468962_4x3.jpg",
								"https://www.wfla.com/wp-content/uploads/sites/71/2023/05/GettyImages-1389862392.jpg?w=2560\u0026h=1440\u0026crop=1",
								"https://i.natgeofe.com/n/548467d8-c5f1-4551-9f58-6817a8d2c45e/NationalGeographic_2572187_2x3.jpg",
								"https://www.bluecross.org.uk/sites/default/files/d8/2020-12/BX146296_237A1405.JPG",
							},
						},
					},
					{
						Item: "bird",
						ImageData: structs.ImageArray{
							Images: [constants.SearchResultNumber]string{
								"https://www.birdlife.org/wp-content/uploads/2022/09/Pelican_portrait_high-res-1024x497.jpg",
								"https://www.birds.cornell.edu/home/wp-content/uploads/2023/09/334289821-Baltimore_Oriole-Matthew_Plante.jpg",
								"https://media.cnn.com/api/v1/images/stellar/prod/231102091639-american-birds-renamed-restricted.jpg?c=16x9\u0026q=h_833,w_1480,c_fill",
								"https://www.allaboutbirds.org/news/wp-content/uploads/2009/04/Allens_Hummingbird-Bob_Gunderson-BS-top-1280x756.jpg",
								"https://natureconservancy-h.assetsadobe.com/is/image/content/dam/tnc/nature/en/photos/a/m/AmericanGoldfinch_MattWilliams_4000x2200.jpg?crop=0%2C0%2C4000%2C2200\u0026wid=4000\u0026hei=2200\u0026scl=1.0",
								"https://www.birdlife.org/wp-content/uploads/2022/09/Pelican_portrait_high-res-scaled.jpg",
								"https://www.nps.gov/grsm/images/pinewarbler.jpg?maxwidth=1300\u0026maxheight=1300\u0026autorotate=false",
								"https://rockies.audubon.org/sites/default/files/styles/hero_mobile/public/aud_lazuli-bunting_populus-angustifolia_200508-01959_nape_jr_photo-evan-barrientos_mobile.jpg?itok=AJfB-Xx_",
								"https://thumbnails.cbc.ca/maven_legacy/thumbnails/349/23/bird.jpg",
								"https://i.natgeofe.com/n/6f9b6d9e-5797-4867-a859-7b0c2a66cd3b/02-bird-of-paradise-A012_C010_1029SF_0001575_3x2.jpg",
							},
						},
					},
					{
						Item: "apple",
						ImageData: structs.ImageArray{
							Images: [constants.SearchResultNumber]string{
								"https://www.apple.com/newsroom/images/2023/09/apple-introduces-the-advanced-new-apple-watch-series-9/article/Apple-Watch-S9-hero-230912_Full-Bleed-Image.jpg.large.jpg",
								"https://developer.apple.com/wwdc23/hero/endframes/p3-startframe-large_2x.jpg",
								"https://upload.wikimedia.org/wikipedia/commons/f/fa/Apple_logo_black.svg",
								"https://pbs.twimg.com/profile_images/1717013664954499072/2dcJ0Unw_400x400.png",
								"https://www.tailorbrands.com/wp-content/uploads/2021/01/apple-evolution-thumbnail.jpg",
								"https://static01.nyt.com/images/2023/06/05/multimedia/05APPLE-VR-gcmq/05APPLE-VR-gcmq-articleLarge.jpg?quality=75\u0026auto=webp\u0026disable=upscale",
								"https://post.medicalnewstoday.com/wp-content/uploads/sites/3/2022/07/what_to_know_apples_green_red_1296x728_header-1024x575.jpg",
								"https://store.storeimages.cdn-apple.com/1/as-images.apple.com/is/MT2Y3ref_VW_34FR+watch-case-40-aluminum-midnight-nc-se_VW_34FR+watch-face-40-aluminum-midnight-se_VW_34FR?wid=752\u0026hei=720\u0026bgc=fafafa\u0026trim=1\u0026fmt=p-jpg\u0026qlt=80\u0026.v=Z0VkY0NhREJ6WjFzb3N5VEYrKzFoRStGZUJWLzNFUFVydllxZFp0d1M4NTlEbzMrd1Z5SUpEd0hiT0ZLRlZGNGRCU0luK254NGZZeFNSZCtCaXAxdGg4OFMza08xNVcyVm9vUnNISFZ1UEdiaytyMkV2UXJqeE9wOUlGWnU0cExJQ3VXcFZIUHhBMlU5OU5QT1pyZjQ3WlBtR0ZPUkFPYlc1NC9nRzhiVTZUYlg5SUJPR0VaZnV5YVlTck5WQzFJWitOTEs5T0laM0FBYmtOdWx0aUJtTm5YU0ZMdUpkZktWZmlmcG5VMHJzOD0",
								"https://lookaside.fbsbx.com/lookaside/crawler/media/?media_id=100064606995009",
								"https://www.apple.com/newsroom/images/live-action/wwdc-2023/standard/privacy/Apple-WWDC23-privacy-logo-230605_big.jpg.large.jpg",
							},
						},
					},
					{
						Item: "golang",
						ImageData: structs.ImageArray{
							Images: [constants.SearchResultNumber]string{
								"https://www.freecodecamp.org/news/content/images/2021/10/golang.png",
								"https://ik.imagekit.io/ably/ghost/prod/2021/02/guide-to-pubsub-in-golang-2.jpg?tr=w-1728,q-50",
								"https://mobisoftinfotech.com/resources/wp-content/uploads/2022/02/og-hire-golang-developers.png",
								"https://miro.medium.com/v2/resize:fit:2000/0*KvlaCAJFzT86-D7J.png",
								"https://i.stack.imgur.com/7QmTd.jpg",
								"https://miro.medium.com/v2/resize:fit:1400/0*SoqCeEz9EctJBXKw.png",
								"https://static-00.iconduck.com/assets.00/golang-icon-398x512-eygvdisi.png",
								"https://miro.medium.com/v2/resize:fit:1123/1*2gxN2Z_5FicXvYuKiNpdlA.png",
								"https://www.altoros.com/blog/wp-content/uploads/2015/03/golang-internals-part-1-main-concepts-and-project-structure.png",
								"https://www.freecodecamp.org/news/content/images/2022/10/golang.png",
							},
						},
					},
					{
						Item: "coffee",
						ImageData: structs.ImageArray{
							Images: [constants.SearchResultNumber]string{
								"https://neurosciencenews.com/files/2023/06/coffee-brain-caffeine-neuroscincces.jpg",
								"https://upload.wikimedia.org/wikipedia/commons/e/e4/Latte_and_dark_coffee.jpg",
								"https://www.eatright.org/-/media/images/eatright-articles/eatright-article-feature-images/benefitsofcoffee_600x450.jpg?as=0\u0026w=967\u0026rev=6c8a9cd4a94d4cac8af8543054fd7b07\u0026hash=4C95EA3A031A783C6DFA3ED04AB4FED4",
								"https://www.rush.edu/sites/default/files/media-images/Coffee_OpenGraph.png",
								"https://i0.wp.com/images-prod.healthline.com/hlcmsresource/images/AN_images/butter-coffee-1296x728-feature.jpg?w=1155\u0026h=1528",
								"https://www.bhg.com/thmb/Dw9Sxcivh_bczUpo91Egapy-lGc=/7952x0/filters:no_upscale():strip_icc()/feshly-brewed--latte-coffee-on-a-white-table-1434303312-4d24a51e3bc748d68553a836499d0100.jpg",
								"https://upload.wikimedia.org/wikipedia/commons/thumb/e/e4/Latte_and_dark_coffee.jpg/640px-Latte_and_dark_coffee.jpg",
								"https://dynaimage.cdn.cnn.com/cnn/c_fill,g_auto,w_1200,h_675,ar_16:9/https%3A%2F%2Fcdn.cnn.com%2Fcnnnext%2Fdam%2Fassets%2F150929101049-black-coffee-stock.jpg",
								"https://static.scientificamerican.com/sciam/cache/file/4A9B64B5-4625-4635-848AF1CD534EBC1A_source.jpg?w=1200",
								"https://upload.wikimedia.org/wikipedia/commons/thumb/c/c5/Roasted_coffee_beans.jpg/640px-Roasted_coffee_beans.jpg",
							},
						},
					},
				},
			},
		},
		{
			name: "Failed get initImages.json",
			args: args{beforeLevel: constants.BeforeLevel1},
			want: structs.InitImageItems{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInitImagesJson(tt.args.beforeLevel); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("GetInitImagesJson() = %v, want %v", got, &tt.want)
			}
		})
	}
}

func TestGetGoogleCustomSearchApiResponse(t *testing.T) {
	type args struct {
		env         structs.Env
		keyword     string
		beforeLevel int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Successful get Google Custom Search Api Response",
			args: args{
				env: structs.Env{
					SearchEngineId: GetEnvData(constants.BeforeLevel3).SearchEngineId,
					KeyFileName:    GetEnvData(constants.BeforeLevel3).KeyFileName,
					CsePath:        GetEnvData(constants.BeforeLevel3).CsePath,
				},
				keyword:     "iphone",
				beforeLevel: constants.BeforeLevel3,
			},
			want: true,
		},
		{
			name: "Failed get Google Custom Search Api Response1",
			args: args{
				env: structs.Env{
					SearchEngineId: GetEnvData(constants.BeforeLevel3).SearchEngineId,
					KeyFileName:    "key-filename.json",
					CsePath:        GetEnvData(constants.BeforeLevel3).CsePath,
				},
				keyword:     "iphone",
				beforeLevel: constants.BeforeLevel3,
			},
			want: false,
		},
		{
			name: "Failed get Google Custom Search Api Response2",
			args: args{
				env: structs.Env{
					SearchEngineId: GetEnvData(constants.BeforeLevel3).SearchEngineId,
					KeyFileName:    GetEnvData(constants.BeforeLevel3).KeyFileName,
					CsePath:        GetEnvData(constants.BeforeLevel3).CsePath,
				},
				keyword:     "iphone",
				beforeLevel: constants.BeforeLevel0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetGoogleCustomSearchApiResponse(tt.args.env, tt.args.keyword, tt.args.beforeLevel)
			if (got != nil) != tt.want {
				t.Errorf("GetGoogleCustomSearchApiResponse() got1 = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWalkTargetPath(t *testing.T) {
	type args struct {
		targetFile  string
		beforeLevel int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Successful get target path",
			args: args{
				targetFile:  ".env",
				beforeLevel: constants.BeforeLevel3,
			},
			want: "Full env path",
		},
		{
			name: "Failed get target path",
			args: args{
				targetFile:  "hoge.go",
				beforeLevel: constants.BeforeLevel3,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWalkTargetPath(tt.args.targetFile, tt.args.beforeLevel); got != tt.want {
				t.Errorf("GetWalkTargetPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

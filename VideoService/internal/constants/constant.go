package constants

const VideoFeedCount = 30

const (
	MinioEndPoint    = "106.54.208.133:18000"
	MinioAccessKeyID = "admin"
	//MinioAccessKeyID     = "zMAQ2yPom14RvvcZagYl"
	//MinioSecretAccessKey = "eTtObGyyuJ84b80y1d0BIF4ebVcEU81oTLZJwVJ3"
	MinioSecretAccessKey  = "Lhj000922"
	MiniouseSSL           = false
	MinioVideoBucketName  = "videobucket"
	MinioImgBucketName    = "imagebucket"
	MinioURLConvertScheme = "http"
	MinioURLConvertHost   = MinioEndPoint
)

const (
	VideosTableName   = "video"
	FavoriteTableName = "favorite"
)

const (
	RedisAddr     = "127.0.0.1:6379"
	RedisPassword = "Lhj000922"
)

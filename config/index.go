package config

import (
	"github.com/PurpleSec/logx"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const (
	accessKeyID     = "ygmt0IforwxOWgh7ijvZ"
	secretAccessKey = "tp1KrfHzWpoMGQ9iHXz9plOXpxjXtwfVDBNYbKXJ"
	endpoint        = "127.0.0.1:9000"
	useSSL          = false

	// FileSavePath 文件保存路径
	FileSavePath = "F:\\tmp\\"
)

var (
	Log         = logx.Console(logx.Info)
	lock        = &sync.Mutex{}
	minioClient *minio.Client

	FolderType   = &FileType{""}
	ImageType    = &FileType{"image"}
	VideoType    = &FileType{"video"}
	DocumentType = &FileType{"document"}
	OtherType    = &FileType{"other"}

	TypeMap = &TypeMapFactory{}
)

type FileType struct {
	Type string
}

type TypeMapFactory struct {
	types map[string]*FileType
}

func (f *TypeMapFactory) GetFileType(fileType string) *FileType {
	if t, ok := f.types[fileType]; ok {
		return t
	}
	return OtherType
}

// init 使用享元模式, 保证每种文件类型只有一个实例
func init() {
	factory := &TypeMapFactory{
		types: make(map[string]*FileType),
	}
	factory.types[""] = FolderType
	factory.types["image/jpeg"] = ImageType
	factory.types["image/png"] = ImageType
	factory.types["image/gif"] = ImageType
	factory.types["image/bmp"] = ImageType
	factory.types["image/webp"] = ImageType
	factory.types["image/tiff"] = ImageType
	factory.types[".jpg"] = ImageType
	factory.types[".png"] = ImageType
	factory.types[".gif"] = ImageType
	factory.types[".bmp"] = ImageType
	factory.types[".webp"] = ImageType
	factory.types[".tiff"] = ImageType

	factory.types["video/mp4"] = VideoType
	factory.types["video/avi"] = VideoType
	factory.types["video/mpeg"] = VideoType
	factory.types["video/quicktime"] = VideoType
	factory.types[".mp4"] = VideoType
	factory.types[".avi"] = VideoType
	factory.types[".mpeg"] = VideoType
	factory.types[".mov"] = VideoType

	factory.types["application/msword"] = DocumentType
	factory.types["application/vnd.openxmlformats-officedocument.wordprocessingml.document"] = DocumentType
	factory.types["application/vnd.ms-excel"] = DocumentType
	factory.types["application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"] = DocumentType
	factory.types["application/vnd.ms-powerpoint"] = DocumentType
	factory.types["application/vnd.openxmlformats-officedocument.presentationml.presentation"] = DocumentType
	factory.types["application/pdf"] = DocumentType
	factory.types[".doc"] = DocumentType
	factory.types[".docx"] = DocumentType
	factory.types[".xls"] = DocumentType
	factory.types[".xlsx"] = DocumentType
	factory.types[".ppt"] = DocumentType
	factory.types[".pptx"] = DocumentType
	factory.types[".pdf"] = DocumentType
	factory.types[".log"] = OtherType
	factory.types[".txt"] = OtherType

	factory.types["application/zip"] = OtherType
	factory.types["application/x-rar-compressed"] = OtherType
	factory.types["application/x-7z-compressed"] = OtherType
	factory.types["application/x-tar"] = OtherType
	factory.types["application/x-gzip"] = OtherType
	factory.types["application/x-bzip2"] = OtherType
	factory.types["application/x-bzip"] = OtherType
	factory.types["application/x-ace-compressed"] = OtherType
	factory.types["application/x-apple-diskimage"] = OtherType
	factory.types[".zip"] = OtherType
	factory.types[".rar"] = OtherType
	factory.types[".7z"] = OtherType
	factory.types[".tar"] = OtherType
	factory.types[".gz"] = OtherType
	factory.types[".bz2"] = OtherType
	factory.types[".ace"] = OtherType
	factory.types[".dmg"] = OtherType
	factory.types[".iso"] = OtherType
	factory.types[".ini"] = OtherType
	factory.types[".conf"] = OtherType
	factory.types[".cfg"] = OtherType
	factory.types[".json"] = OtherType
	factory.types[".xml"] = OtherType
	factory.types[".html"] = OtherType
	factory.types[".htm"] = OtherType
	factory.types[".php"] = OtherType
	factory.types[".js"] = OtherType
	factory.types[".css"] = OtherType
	factory.types[".md"] = OtherType
	factory.types[".markdown"] = OtherType
	factory.types[".csv"] = OtherType
	factory.types[".sql"] = OtherType
	factory.types[".db"] = OtherType
	factory.types[".dbf"] = OtherType
	factory.types[".dbx"] = OtherType
	factory.types[".java"] = OtherType
	factory.types[".class"] = OtherType
	factory.types[".py"] = OtherType

	TypeMap = factory
}

// Minio 获取minio客户端 单例模式
func Minio() (*minio.Client, error) {
	var err error = nil
	if minioClient == nil {
		lock.Lock()
		defer lock.Unlock()
		if minioClient == nil {
			minioClient, err = minio.New(endpoint, &minio.Options{
				Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
				Secure: useSSL,
			})
			if err != nil {
				Log.Info("Minio create error :%v\n", err)
			}
		}
	}
	Log.Info("minioClient :%v\n", minioClient)
	return minioClient, err
}

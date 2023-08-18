package pubutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/teamssix/cf/pkg/util/global"
)

type TimestampCache struct {
	Timestamp     int64
	TimestampType string
}

type OSSCache struct {
	AccessKeyId  string
	BucketACL    string
	BucketURL    string
	Name         string
	ObjectNumber string
	ObjectSize   string
	Region       string
	SN           string
}

type ECSCache struct {
	AccessKeyId          string
	InstanceId           string
	InstanceName         string
	OSName               string
	OSType               string
	PrivateIpAddress     string
	PublicIpAddress      string
	RegionId             string
	SN                   string
	Status               string
	CloudAssistantStatus string
}

type RDSCache struct {
	AccessKeyId      string
	DBInstanceId     string
	DBInstanceStatus string
	Engine           string
	EngineVersion    string
	RegionId         string
	SN               string
}

type TakeoverConsoleCache struct {
	AccessKeyId            string
	ConsoleAccessKeyId     string
	ConsoleAccessKeySecret string
	CreateTime             string
	LoginUrl               string
	Password               string
	PrimaryAccountId       string
	Provider               string
	UserId                 string
	UserName               string
}

type ImageShareCache struct {
	AccessKeyId    string
	ImageId        string
	InstanceId     string
	Provider       string
	Region         string
	ShareAccountId string
	Status         string
	Time           string
}

type RDSAccountsCache struct {
	AccessKeyId  string
	DBInstanceId string
	Provider     string
	Region       string
	UserName     string
	Password     string
	Engine       string
	CreateTime   string
}

type RDSPublicCache struct {
	AccessKeyId       string
	DBInstanceId      string
	Provider          string
	Region            string
	IPAddress         string
	ConnectionAddress string
	Port              string
	Engine            string
	CreateTime        string
}

type RDSWhiteListCache struct {
	AccessKeyId  string
	DBInstanceId string
	Provider     string
	Region       string
	IPArrayName  string
	IPType       string
	WhiteList    string
	IPList       string
	Engine       string
	CreateTime   string
}

type Provider struct {
	CN string
	EN string
}

func GetUserDir() string {
	home, _ := os.UserHomeDir()
	return home
}

func GetConfigFilePath() string {
	home, _ := GetCFHomeDir()
	CreateFolder(home)
	configFilePath := filepath.Join(home, "cache.db")
	return configFilePath
}

func GetCFHomeDir() (string, error) {
	return filepath.Join(GetUserDir(), global.AppDirName), nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func CreateFolder(folder string) {
	if !FileExists(folder) {
		log.Tracef("创建 %s 目录 (Create %s directory): ", folder, folder)
		_ = os.MkdirAll(folder, 0700)
	}
}

func FormatFileSize(fileSize int64) (size string) {
	if fileSize < 1024 {
		return fmt.Sprintf("%.2f B", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2f KB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f MB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f GB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f TB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else {
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

func ReadFile(filePath string) (bool, string) {
	if FileExists(filePath) {
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		content, err := ioutil.ReadAll(file)
		return true, string(content)
	} else {
		return false, ""
	}
}

func StringClean(str string) string {
	str = strings.TrimSpace(str)
	str = strings.Replace(str, "\n", "", -1)
	return str
}

func MaskAK(ak string) string {
	if len(ak) > 7 {
		prefix := ak[:2]
		suffix := ak[len(ak)-6:]
		return prefix + strings.Repeat("*", 18) + suffix
	} else {
		return ak
	}
}

func IN(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

func CurrentTime() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	return formattedTime
}

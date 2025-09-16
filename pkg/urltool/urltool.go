package urltool

import (
	"errors"
	"net/url"
	"path"
)

// GetBasePath 获取URL路径的最后一部分
func GetBasePath(targetUrl string) (string, error) {
	myUrl, err := url.Parse(targetUrl)
	if len(myUrl.Host) == 0 {
		return "", errors.New("invalid URL")
	}
	// 检查路径是否为空或只有根路径
	if myUrl.Path == "" {
		return "", errors.New("URL has no path component")
	}
	return path.Base(myUrl.Path), err
}

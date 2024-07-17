package common

import (
	"encoding/base64"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/kms"
)

type Kms struct {
	RegionId        string `yaml:"region_id" json:"region_id"`
	Domain          string `yaml:"domain" json:"domain"`
	AccessKeyId     string `yaml:"access_key_id" json:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret" json:"access_key_secret"`
	KeyId           string `yaml:"key_id" json:"key_id"`
}

func KmsDecrypt(km Kms, str string) (string, error) {
	client, err := kms.NewClientWithAccessKey(km.RegionId, km.AccessKeyId, km.AccessKeySecret)
	if err != nil {
		// Handle exceptions
		return "", err
	}
	request := kms.CreateDecryptRequest()
	request.Scheme = "https"
	request.CiphertextBlob = str
	request.Domain = km.Domain
	request.SetHTTPSInsecure(true)
	res, err := client.Decrypt(request)
	if err != nil || !res.IsSuccess() {
		return "", err
	}
	resDe, _ := base64.StdEncoding.DecodeString(res.Plaintext)
	return string(resDe), nil
}

package wg

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"wg/config"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetEndPoint() error {
	cmd := "wg show " + config.C.Endpoint + " endpoints"
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out)

	if len(output) == 0 {
		return errors.New("no endpoint")
	}
	endpointData := strings.Split(output, " ")
	if len(endpointData) != 2 {
		return errors.New("wrong endpoint info")
	}
	endpoint := endpointData[1]

	endPointInfo := strings.Split(endpoint, ":")

	// make http request with resty to increase port number
	// and update endpoint info
	resp, err := resty.New().R().Post("http://" + endPointInfo[0] + ":8080/")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New("http request failed")
	}

	logrus.Infof("write new port: %s", resp.Body())
	// replace endpoint info
	newEndpoint := strings.Replace(endpoint, endPointInfo[1], resp.String(), -1)
	err = replaceFileData(config.C.Path, []byte(endpoint), []byte(newEndpoint))
	if err != nil {
		return err
	}

	// restart service with wg-quick down bmh
	cmd = "wg-quick down " + config.C.Endpoint
	exec.Command("sh", "-c", cmd).Run()

	time.Sleep(time.Second * 2)

	cmd = "wg-quick up " + config.C.Endpoint
	exec.Command("sh", "-c", cmd).Run()

	return nil
}

// getPortNumber get port number from config file, if not found returns error
func GetPortNumber() (uint, error) {
	v, err := parseConf()
	if err != nil {
		return 0, err
	}

	return v.GetUint("interface.listenport"), nil
}

func parseConf() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(config.C.Path)
	v.SetConfigType("ini")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}

func IncrPortNumber() (uint, error) {
	strTpl := "ListenPort = %d"

	curPort, err := GetPortNumber()
	if err != nil {
		return 0, err
	}
	newPort := curPort + 1

	err = replaceFileData(config.C.Path, []byte(fmt.Sprintf(strTpl, curPort)), []byte(fmt.Sprintf(strTpl, newPort)))
	if err != nil {
		return 0, err
	}
	return newPort, nil
}

// use systemctl to restart wg-quick@wg0 service
func RestartService() error {
	cmd := "systemctl restart wg-quick@wg0"
	return exec.Command("sh", "-c", cmd).Run()
}

func replaceFileData(file string, old, new []byte) error {
	// get config file contents
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	newContent := bytes.Replace(content, old, new, -1)
	err = os.WriteFile(config.C.Path, newContent, 0644)
	if err != nil {
		return err
	}
	return nil
}
package ikuai

import (
	"testing"

	"github.com/tossp/ikuai/action"
)

func getClient() *IKuai {
	i := NewIKuai("http://192.168.9.1", "root", "123456", true, false)

	return i
}

func TestLogin(t *testing.T) {

	login, err := getClient().Login()
	if err != nil {
		t.Error(err)
	}

	t.Log(login)
}

func TestIKuai_ShowIPGroup(t *testing.T) {
	client := getClient()

	login, err := client.Login()
	if err != nil {
		t.Error(err)
	}

	res := &action.ShowIPGroupResult{}

	ip, err := client.Run(login, action.NewIPGroupShowAction(), res)
	if err != nil {
		t.Error(err)
	}

	t.Log(ip)
}

func TestIKuai_ShowMonitorLan(t *testing.T) {
	client := getClient()

	_, err := client.Login()
	if err != nil {
		t.Error(err)
	}

	result, err := client.ShowMonitorLan()
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", result)
}

func TestIKuai_ShowSysStat(t *testing.T) {
	client := getClient()

	_, err := client.Login()
	if err != nil {
		t.Error(err)
	}

	result, err := client.ShowSysStat()

	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", result)
}

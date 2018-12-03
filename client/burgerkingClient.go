package client

import (
	"burgerking/utils"
	"fmt"
	"net/url"
)

type BurgerkingClient struct {
	Client
	resp string
}

func (bc *BurgerkingClient) StartSurvey(id string) (err error) {
	bc.Client = MakeClient()
	content, err := bc.Get("https://tellburgerking.com.cn/")
	if err != nil {
		fmt.Println("bc.Get err %s", err)
		return err
	}
	f, err := ParseForm(content)
	if err != nil {
		fmt.Println(err)
		return err
	}
	bc.SetIP(f.URL, utils.RandomIP())
	postForm := make(url.Values)
	postForm.Add("JavaScriptEnabled", "1")
	postForm.Add("FIP", "True")
	postForm.Add("AcceptCookies", "Y")
	postForm.Add("NextButton", "继续")
	content, err = bc.Post(f.URL, postForm)
	if err != nil {
		fmt.Println("bc.Post err %s", err)
		return err
	}

	f, err = ParseForm(content)
	if err != nil {
		fmt.Println(err)
		return err
	}
	postForm = make(url.Values)
	postForm.Add("JavaScriptEnabled", "1")
	postForm.Add("FIP", "True")
	postForm.Add("CN1", id[:3])
	postForm.Add("CN2", id[3:6])
	postForm.Add("CN3", id[6:9])
	postForm.Add("CN4", id[9:12])
	postForm.Add("CN5", id[12:15])
	postForm.Add("CN6", id[15:16])
	//bc.SetIP(f.URL, utils.RandomIP())
	content, err = bc.Post(f.URL, postForm)
	if err != nil {
		fmt.Println("bc.Post err %s", err)
		return err
	}

	bc.resp = content
	return
}

func (bc *BurgerkingClient) DoSurvey() (string, error) {
	content := bc.resp
	f, err := ParseForm(content)
	if err != nil {
		fmt.Println(err)
		return content, err
	}
	for f.ValCode == "" {
		fmt.Printf("working on %s \n", f.URL)
		values := f.FormatForm()
		content, err := bc.Post(f.URL, values)
		if err != nil {
			fmt.Println("bc.Post err %s", err)
			return content, err
		}
		f, err = ParseForm(content)
		if err != nil {
			fmt.Println(err)
			return content, err
		}
	}
	return f.ValCode, nil
}

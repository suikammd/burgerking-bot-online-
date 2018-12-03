package client

import (
	"net/http"
	"strings"
	"net/url"
	"io/ioutil"
	"net/http/cookiejar"
	"log"
	"os"
	"errors"
)

type Client struct {
	client http.Client
	UA string
	lastURL string
	logger  *log.Logger
}

func (c *Client) Get(url string) (string, error){
	if url == c.lastURL {
		err := errors.New("Inf loop.")
		return "", err
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.logger.Printf("http request err %s", err)
		return "", err
	}
	request.Header.Set("User-Agent", c.UA)
	request.Header.Set("Referer", c.lastURL)
	request.Header.Set("Origin", "https://tellburgerking.com.cn")

	response, err := c.client.Do(request)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	c.lastURL = url
	return string(body), nil
}

func (c *Client) Post(url string, form url.Values) (string, error) {
	if url == c.lastURL {
		err := errors.New("Inf loop.")
		return "", err
	}
	request, err := http.NewRequest("POST", url, strings.NewReader(form.Encode()))
	if err != nil {
		c.logger.Printf("http request err %s", err)
		return "", err
	}
	request.Header.Set("User-Agent", c.UA)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Referer", c.lastURL)
	request.Header.Set("Origin", "https://tellburgerking.com.cn")

	response, err := c.client.Do(request)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	c.lastURL = url
	return string(body), nil
}

func (c *Client) SetIP(URLStr string, ip string) error{
	u, err := url.Parse(URLStr)
	if err != nil {
		c.logger.Printf("url.Parse err %s", err)
		return err
	}
	cookies := c.client.Jar.Cookies(u)
	for i, cookie := range cookies {
		if cookie.Name == "T"{
			values, _ := url.ParseQuery(cookie.Value)
			values["RA"] = []string{ip}
			cookie.Value = values.Encode()
			cookies[i] = cookie
			return nil
		}
	}
	return nil
}

func MakeClient() Client {
	cookieJar, _ := cookiejar.New(nil)
	return Client{
		UA: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36",
		client: http.Client{
			Jar: cookieJar,
		},
		logger: log.New(os.Stdout, "[HTTP Client]", log.LstdFlags),
	}
}


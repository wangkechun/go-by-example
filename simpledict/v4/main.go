package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wangkechun/go-by-example/simpledict/v4/tools"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

const (
	SUCCESS = 0
)

type FanYiGoRequest struct {
	Word       string `json:"word"`
	DeviceType string `json:"deviceType"`
	DeviceId   string `json:"deviceId"`
}

type FanYiGoRespond struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data struct {
		Dict struct {
			Base struct {
				Explain struct {
					WordMeans []struct {
						Part  string   `json:"part"`
						Means []string `json:"means"`
					} `json:"word_means"`
					CaseWord []string `json:"case_word"`
				} `json:"explain"`
				ExplainZh string `json:"explain_zh"`
				Exchange  struct {
					WordIng   []string      `json:"word_ing"`
					WordDone  []string      `json:"word_done"`
					WordPast  []string      `json:"word_past"`
					WordThird []string      `json:"word_third"`
					WordPl    []interface{} `json:"word_pl"`
				} `json:"exchange"`
				PronZh string `json:"pron_zh"`
				PronAm struct {
					Text      string `json:"text"`
					Pronounce string `json:"pronounce"`
				} `json:"pron_am"`
				PronEn struct {
					Text      string `json:"text"`
					Pronounce string `json:"pronounce"`
				} `json:"pron_en"`
				Word    string `json:"word"`
				Chinese bool   `json:"chinese"`
			} `json:"base"`
			PracticalSentences []interface{} `json:"practical_sentences"`
			TrueSentences      []struct {
				Tag       string        `json:"tag"`
				Sentences []interface{} `json:"sentences"`
			} `json:"true_sentences"`
			EeMeans      []interface{} `json:"ee_means"`
			RootsAffixes []interface{} `json:"roots_affixes"`
			Phrases      []interface{} `json:"phrases"`
			Synonym      []struct {
				Part      string `json:"part"`
				WordMeans []struct {
					Mean  string   `json:"mean"`
					Words []string `json:"words"`
				} `json:"word_means"`
			} `json:"synonym"`
			Antonym []struct {
				Part      string `json:"part"`
				WordMeans []struct {
					Mean  string   `json:"mean"`
					Words []string `json:"words"`
				} `json:"word_means"`
			} `json:"antonym"`
			IndustryDict   []interface{} `json:"industry_dict"`
			SynonymExplain []interface{} `json:"synonym_explain"`
			Collins        []struct {
				Entries []struct {
					Define    string `json:"define"`
					Trans     string `json:"trans"`
					Part      string `json:"part"`
					Sentences []struct {
						Origin string `json:"origin"`
						Trans  string `json:"trans"`
					} `json:"sentences"`
				} `json:"entries"`
			} `json:"collins"`
		} `json:"dict"`
	} `json:"data"`
}

type CaiYunDictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
	UserID    string `json:"user_id"`
}

type CaiYunDictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
		KnownInLaguages int `json:"known_in_laguages"`
		Description     struct {
			Source string      `json:"source"`
			Target interface{} `json:"target"`
		} `json:"description"`
		ID   string `json:"id"`
		Item struct {
			Source string `json:"source"`
			Target string `json:"target"`
		} `json:"item"`
		ImageURL  string `json:"image_url"`
		IsSubject string `json:"is_subject"`
		Sitelink  string `json:"sitelink"`
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func FanYiGoQuery(word string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer func() {
		wg.Done()
		fmt.Println("-------------------------------------------")
		mutex.Unlock()
	}()
	client := &http.Client{}
	request := FanYiGoRequest{Word: word, DeviceId: "a132e3a425ba0e88e7f89e4a81333a7e", DeviceType: "web"}
	urlCodeStr := tools.ConvertStruct2UrlCode(request)

	var data = bytes.NewReader([]byte(urlCodeStr))
	req, err := http.NewRequest("POST", "https://open.qingxun.com/apiConsole/dict/search", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "acw_tc=2f624a2216892361198977971e07bfd1bd6953db1c523c88d8504a2ed651e0; UM_distinctid=1894e51a16243a-0d01c6ecf85be3-1b525634-1fa400-1894e51a1639f2; CNZZDATA1281143543=2007567745-1689234597-https%253A%252F%252Fwww.baidu.com%252F%7C1689234597; Hm_lvt_4155e261820146fa5d82a7049a03a46a=1689236120; promotionSource=https://open.qingxun.com/#baiduqx03-apifw-fyapi-fyapi1?bd_vid=11127669222185401009; DeviceId=a132e3a425ba0e88e7f89e4a81333a7e; Qx-OpenPlatform-Token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJGWUctSldUIiwiZXhwIjoxNjk3MDEyMTk5LCJpYXQiOjE2ODkyMzYxOTksInRva2VuIjoiQkFEODI3OEMxMEU4RTRGNURGRjhDMTc0MjAyNkQ3QzgwRTUwRERENTZFNjY0MjZBNTZBM0I0NDk5NEU2NThDMzkxQzdDMUEwQkNDRTEwNzM4MjY0QjQ2QjIxMDhDN0EwNUQzOTRERjk2QTUzMDUyNTEzRUREMDFGQzdEMzcyMDEzQTlFMjQ1MzkyMDE5RTE2QjIyRTc2Qjc2NjlDQUE5MjJDRDA5NTFDQjdGQ0U5Q0FGRDdGREZGMTQ2Q0MwNjA0QzI4MDE4QzI2QkY1QkUzQjkyNjBCRTk4MEYyMUNBODMwMUVCREQxNDYyNkVBNjkzNkEzM0ZDNzc2NEVGQThEQ0IyNDQ2QTdEN0I5NzZBMkY4ODMzNTQ5NEZFMjU5OTQ3NDVGQzkxRTgwNzdGNkMzOTBFM0I4MjU1NDZCRjVEMDRBNEU2RDMxMkI4OTQwMkI1NkNGRjc1NEEzQzI3OUJCODMyNjk2MzE5REI4QUQzMUFFMzE5REVDQTg2RUZFNEMyRUZDQUM1NjZEODdCRENFREUzNTc4NThDMEZCRkQ1NDc0NTFCNDVEQTQ0RUZCMzlGMkZENzJBNzgwMzE2NUY1QSJ9.dtup4cEBaHb2HPOE0H7Q96G2XSEeDnrzUpaKn6k4S1A; Hm_lpvt_4155e261820146fa5d82a7049a03a46a=1689236204")
	req.Header.Set("Origin", "https://open.qingxun.com")
	req.Header.Set("Referer", "https://open.qingxun.com/service/dictionariesAPI.html")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("sec-ch-ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var bodyRespond FanYiGoRespond
	json.Unmarshal(bodyText, &bodyRespond)

	mutex.Lock()
	fmt.Println("翻译狗翻译结果如下：")
	if bodyRespond.Code != SUCCESS {
		fmt.Println("暂无翻译QAQ")
		return
	}
	base := bodyRespond.Data.Dict.Base
	fmt.Printf("%s UK: [%s] US: [%s]\n", word, base.PronEn.Text, base.PronAm.Text)
	for _, part := range base.Explain.WordMeans {
		fmt.Printf("%s ", part.Part)
		for _, mean := range part.Means {
			fmt.Printf("%s; ", mean)
		}
		fmt.Println()
	}
}

func CaiYunQuery(word string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer func() {
		wg.Done()
		fmt.Println("-------------------------------------------")
		mutex.Unlock()
	}()
	client := &http.Client{}
	request := CaiYunDictRequest{TransType: "en2zh", Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	var data = bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("DNT", "1")
	req.Header.Set("os-version", "")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	req.Header.Set("app-name", "xy")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("device-id", "")
	req.Header.Set("os-type", "web")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("Sec-Fetch-Site", "cross-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cookie", "_ym_uid=16456948721020430059; _ym_d=1645694872")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode:", resp.StatusCode, "body", string(bodyText))
	}
	var dictResponse CaiYunDictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	mutex.Lock()
	fmt.Println("彩云小译翻译结果如下：")
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.En, "US:", dictResponse.Dictionary.Prons.EnUs)
	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
	example: simpleDict hello
			`)
		os.Exit(1)
	}
	word := os.Args[1]

	var wg = sync.WaitGroup{} // 同步，保证两个协程执行完毕才终止程序
	var mutex = sync.Mutex{}  // 互斥，保证两个翻译引擎的翻译结果不交叉打印，即等一个翻译结果打印完了才打印另外一个

	wg.Add(2)
	go FanYiGoQuery(word, &wg, &mutex)
	go CaiYunQuery(word, &wg, &mutex)

	wg.Wait()
}

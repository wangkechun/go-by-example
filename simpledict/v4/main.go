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
)

type BaiduDictRequest struct {
	From            string `json:"from"`
	To              string `json:"to"`
	Query           string `json:"query"`
	TransType       string `json:"transtype"`
	Token           string `json:"token"`
	SimpleMeansFlag string `json:"simple_means_flag"`
	Sign            string `json:"sign"`
	Domain          string `json:"domain"`
	Ts              string `json:"ts"`
}

type BaiduDictRespond struct {
	TransResult struct {
		Data []struct {
			Dst string `json:"dst"`
			Src string `json:"src"`
		} `json:"data"`
		From     string `json:"from"`
		To       string `json:"to"`
		Status   int    `json:"status"`
		Type     int    `json:"type"`
		Phonetic []struct {
			SrcStr string `json:"src_str"`
			TrgStr string `json:"trg_str"`
		} `json:"phonetic"`
	} `json:"trans_result"`
	DictResult struct {
		Edict struct {
			Item []struct {
				TrGroup []struct {
					Tr          []string `json:"tr"`
					Example     []string `json:"example"`
					SimilarWord []string `json:"similar_word"`
				} `json:"tr_group"`
				Pos string `json:"pos"`
			} `json:"item"`
			Word string `json:"word"`
		} `json:"edict"`
		From        string `json:"from"`
		SimpleMeans struct {
			WordName  string   `json:"word_name"`
			From      string   `json:"from"`
			WordMeans []string `json:"word_means"`
			Exchange  struct {
				WordThird []string `json:"word_third"`
				WordIng   []string `json:"word_ing"`
				WordDone  []string `json:"word_done"`
				WordPast  []string `json:"word_past"`
			} `json:"exchange"`
			Tags struct {
				Core  []string `json:"core"`
				Other []string `json:"other"`
			} `json:"tags"`
			MemorySkill string `json:"memory_skill"`
			Symbols     []struct {
				PhEn  string `json:"ph_en"`
				PhAm  string `json:"ph_am"`
				Parts []struct {
					Part  string   `json:"part"`
					Means []string `json:"means"`
				} `json:"parts"`
				PhOther string `json:"ph_other"`
			} `json:"symbols"`
		} `json:"simple_means"`
		Common struct {
			Text string `json:"text"`
			From string `json:"from"`
		} `json:"common"`
		Collins struct {
			Entry []struct {
				EntryID string `json:"entry_id"`
				Type    string `json:"type"`
				Value   []struct {
					MeanType []struct {
						InfoType string `json:"info_type"`
						InfoID   string `json:"info_id"`
						Example  []struct {
							ExampleID string `json:"example_id"`
							TtsSize   string `json:"tts_size"`
							Tran      string `json:"tran"`
							Ex        string `json:"ex"`
							TtsMp3    string `json:"tts_mp3"`
						} `json:"example"`
					} `json:"mean_type"`
					Gramarinfo []struct {
						Tran  string `json:"tran"`
						Type  string `json:"type"`
						Label string `json:"label"`
					} `json:"gramarinfo"`
					Tran   string `json:"tran"`
					Def    string `json:"def"`
					MeanID string `json:"mean_id"`
					Posp   []struct {
						Label string `json:"label"`
					} `json:"posp"`
				} `json:"value"`
			} `json:"entry"`
			WordName      string `json:"word_name"`
			Frequence     string `json:"frequence"`
			WordEmphasize string `json:"word_emphasize"`
			WordID        string `json:"word_id"`
		} `json:"collins"`
		Lang   string `json:"lang"`
		Oxford struct {
			Entry []struct {
				Tag  string `json:"tag"`
				Name string `json:"name"`
				Data []struct {
					Tag  string `json:"tag"`
					Data []struct {
						Tag   string `json:"tag"`
						P     string `json:"p"`
						PText string `json:"p_text"`
					} `json:"data"`
				} `json:"data"`
			} `json:"entry"`
			Unbox string `json:"unbox"`
		} `json:"oxford"`
		Usecase struct {
			Idiom []struct {
				P    string `json:"p"`
				Tag  string `json:"tag"`
				Data []struct {
					Tag  string `json:"tag"`
					Data []struct {
						EnText string `json:"enText,omitempty"`
						Tag    string `json:"tag"`
						ChText string `json:"chText,omitempty"`
						Data   []struct {
							EnText string `json:"enText"`
							Tag    string `json:"tag"`
							ChText string `json:"chText"`
							Before []struct {
								Tag  string `json:"tag"`
								Data []struct {
									RText string `json:"r_text"`
									R     string `json:"r"`
									Tag   string `json:"tag"`
								} `json:"data"`
							} `json:"before,omitempty"`
						} `json:"data,omitempty"`
						N string `json:"n,omitempty"`
					} `json:"data"`
				} `json:"data"`
			} `json:"idiom"`
		} `json:"usecase"`
		BaiduPhrase []struct {
			Tit   []string `json:"tit"`
			Trans []string `json:"trans"`
		} `json:"baidu_phrase"`
		Rootsaffixes []struct {
			Indicate string `json:"indicate"`
			Type     string `json:"type"`
			Meanings []struct {
				Td    string `json:"td"`
				Words []struct {
					H             string `json:"h"`
					PartsOfSpeech []struct {
						P   string `json:"p"`
						Dec string `json:"dec"`
						D   string `json:"d"`
					} `json:"partsOfSpeech"`
				} `json:"words"`
			} `json:"meanings"`
			Title string `json:"title"`
		} `json:"rootsaffixes"`
		Sanyms []struct {
			Tit  string `json:"tit"`
			Type string `json:"type"`
			Data []struct {
				P string   `json:"p"`
				D []string `json:"d"`
			} `json:"data"`
		} `json:"sanyms"`
		QueryExplainVideo struct {
			ID           int    `json:"id"`
			UserID       string `json:"user_id"`
			UserName     string `json:"user_name"`
			UserPic      string `json:"user_pic"`
			Query        string `json:"query"`
			Direction    string `json:"direction"`
			Type         string `json:"type"`
			Tag          string `json:"tag"`
			Detail       string `json:"detail"`
			Status       string `json:"status"`
			SearchType   string `json:"search_type"`
			FeedURL      string `json:"feed_url"`
			Likes        string `json:"likes"`
			Plays        string `json:"plays"`
			CreatedAt    string `json:"created_at"`
			UpdatedAt    string `json:"updated_at"`
			DuplicateID  string `json:"duplicate_id"`
			RejectReason string `json:"reject_reason"`
			CoverURL     string `json:"coverUrl"`
			VideoURL     string `json:"videoUrl"`
			ThumbURL     string `json:"thumbUrl"`
			VideoTime    string `json:"videoTime"`
			VideoType    string `json:"videoType"`
		} `json:"queryExplainVideo"`
	} `json:"dict_result"`
	LijuResult struct {
		Double string   `json:"double"`
		Tag    []string `json:"tag"`
		Single string   `json:"single"`
	} `json:"liju_result"`
	Logid int `json:"logid"`
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

func BaiduQuery(word string) {
	client := &http.Client{}
	request := BaiduDictRequest{From: "en", To: "zh", Query: word, TransType: "realtime",
		Token: "a007abed6d9b945ed01cf208df2874dd", SimpleMeansFlag: "3", Sign: "229916.483629", Domain: "common", Ts: "1689218191036"}
	urlCodeStr := tools.ConvertStruct2UrlCode(request)

	var data = bytes.NewReader([]byte(urlCodeStr))
	req, err := http.NewRequest("POST", "https://fanyi.baidu.com/v2transapi?from=en&to=zh", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Acs-Token", "1689218179799_1689218191049_oWf74w9iiFtu0d1P2auAv4FjVKLJJV2d6VimFYMeFwzXox9LrsX4cIIS73SfnGj/IV90+/Rpy0lX2irlz5QgI1V/xIh7xe8/sW9pYbgo2up1HaLiolWOp7QESS2URs1WyZ5eeeUebZPVAzJdS1jTQm7ZAJMr4/8JZPeNNfDQKGEhX9IjKFZtHpoOryxo9w0DIGnpJYeVuAoFPDFz3RBh6t60LjIF7U65c4i3VYbL95pToPbunNkTArRtz8IvuZmpMOUChMb5PfXY4IBCV98S4XzI6n/edYkzS1nSXFmlJJa9tXFCSmNUppQbx/7tnfkNdbrcx93VfTgOpD1OUyYj/aSrleYLPuoivPPtPk2pkTddzmsHB8XwDbFsNq1Npnx4TWwTIpnEfijq5L22zgbb+5XHWek2VqVmpaeC/io+TQoOABYo38UgvwYTBfokYY5gfjT/GbJldL/zBmPDSqKrD5+N5f4M74AlGmYkWxK7wgsw4+S6V4+pnT9mJCqZAUFH")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Cookie", "BIDUPSID=D803FB9891CC21AEB6D1A4B3402C7B91; PSTM=1688961652; BAIDUID=D803FB9891CC21AE6B6452F739AFF584:FG=1; BDORZ=B490B5EBF6F3CD402E515D22BCDA1598; BA_HECTOR=a50g2ka0a02l8h8k00242l8f1iat6bt1o; delPer=0; ZFY=eFo7Beo:B:AXRLKIetzlpqLzYBADhzPvaCSqan1i6yZ1I:C; BAIDUID_BFESS=D803FB9891CC21AE6B6452F739AFF584:FG=1; PSINO=3; H_PS_PSSID=36542_38642_38831_39027_39023_38942_38880_38958_38955_39009_38960_38916_38816_39064_39089_38636_26350_39095; Hm_lvt_64ecd82404c51e03dc91cb9e8c025574=1689218179; Hm_lpvt_64ecd82404c51e03dc91cb9e8c025574=1689218179; REALTIME_TRANS_SWITCH=1; FANYI_WORD_SWITCH=1; HISTORY_SWITCH=1; SOUND_SPD_SWITCH=1; SOUND_PREFER_SWITCH=1; ab_sr=1.0.1_OWQ3NzFjOTE5YzJlYjRkMWIxMmRhMmU0NzNhMGYzNzY4MjQ0ODg1MThhZDRiYjk0ZDk4Y2ZiMzQ3ZTk3YjlkNTg3OWJjZjBmOGU2ZDM3ZGRlOWFhY2E5OTMyNWIzMWJlNGMwN2YyNTAzYTZlNjI5NzkzNGY0MzQ4MTllNmYyMmU1MzZmZDJlYWY0NDk1YmRjZjY4YTZkMGVmNTgwMWVhOA==")
	req.Header.Set("Origin", "https://fanyi.baidu.com")
	req.Header.Set("Referer", "https://fanyi.baidu.com/?aldtype=16047")
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
	var bodyRespond BaiduDictRespond
	json.Unmarshal(bodyText, &bodyRespond)
	fmt.Println("百度翻译结果如下：")
	var symbol = bodyRespond.DictResult.SimpleMeans.Symbols[0]
	fmt.Printf("%s UK: [%s] US: [%s]\n", word, symbol.PhEn, symbol.PhAm)
	for _, part := range symbol.Parts {
		fmt.Printf("%s ", part.Part)
		for _, mean := range part.Means {
			fmt.Printf("%s; ", mean)
		}
		fmt.Println()
	}
}

func CaiYunQuery(word string) {
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
	//CaiYunQuery(word)
	BaiduQuery(word)

	fmt.Println()

	CaiYunQuery(word)
}

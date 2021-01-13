package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	c := colly.NewCollector()
	c.SetRequestTimeout(15 *time.Second)
	str := []string{"叶子", "纸制牌", "指甲油瓶"}
	lens := len(str)
	i := 0
	sc := "NLTjoik1AqL_U8AyN4-oApHP400W1yGm2pdKy16zY30kG5avfen32EeOlHl8t2jgcopShxnQsRvMoeIHr1XbeOLPqOBZ8B870fXeWZ2iXeQ0Xz7Lg8cWJe-_tgzra3T15iI9gSFWsgOfPFA1wDqN7z4Ffo7cYEN-1NxF6gS2bpL0ltwfSFf0D02KKMmNO13CQFBkMYBlgCl2RzaOYdNOGfbueC6EiSQzRnmoNg_I_ZtT6EuGRgtwFTNxyF_uUOqw3hm-MVVsK-JXTKy1J30IUWVptXwPj_fOH0VTsbv_Rlpd9F9XI9_ED7sZJ4R1ikaxaKpwUq4m04-Mzh7seJv3shcKFYv2-KvKf9m28k44IeABIp4mUr5dvL0G8S7TVNuDLT4YjWEEwohe5XjknMkPOtYdKUfF75BPdyoAzoklx-wQ9bjB_ugsR2yqjiw7U56i7UFLiQxGGxtq2Elt9cKV_B8hGcmHJUUlcRGAE7DRCydhs_O-4tdLU5gYaPRS8ADCHSr7VAF_HBYWmtrIGg48bE-aa-YDeFwrtvRIEDvfI1iWSLftfIKJ9WYrLULd0sN7J-929jpRuN5mRbef76g0DBI02qgnCOkkFS9W2OWZuuC9b6dfDefDZtgs0-04reRqSKIBBdCwk8o7iRC3L_PI5loiB9r0BDVz3IyB-cjBYr3HbYDbgUKliNkYozDLM6SbVyLgmR3qi-0luO2HzpM84SKX0o9p-BJtMkVrEpADBXLW9ESmjQFrY9tYj2O-JKdE_YK5oeMdf0uhJ-yeOjU5Pn9y381PMCEHnS-1txeJ04KNX37vcFSgU-mFh3Py3yAT5iLsMBLfonVEa4Co_0HvT0qZLYBGA8WoDATXhhsEyj_1uPoQ085zQPSH4GCu7L8GfI2S0KrcnlYP_nGzQKIJXrfH49VFhnccRckjGRfvOlp3Um1qlvPOAhtjo_bq5koGdtwtP6Txg103R8aj95u3XESsAhTu8S2aHIv2LRYWyA_ejARageYd-dql-3y_8C4tU_CPaGcvTjsvigrdiGyubAE8U-pEpTR3YjxFk5eASs3wac3Q6OvWB9eFuZZ-eVi8FGI46JaKN1bvOr4PDg-e19EDCRGVlfkRF7nNwzQLQZXGqIBZyqilRMsOZvePeyjWlGuykE82rEqsSqFSxcyKj9BN1y5gJd0ecI-hGP57a76K3hQn3V8055vkO10pEegRhOoDP8LrxgbBOglxCWiW1MTdodTX9MV1HVqcT8J5o2vfDbQv_-84k1fSucH2-pegyRhunV4F0LKZ6KgJBjGMEaSv5E9Gf9XiU1Ne_GlmaD0wKQkPB-WIqFdJdPilaoTH9Jao82YqMRn3R8Z7B-csP0v3ots5CjHI5gzqdzTJv1FI9urYx4Ko0r3Tgy_R8LJhTegAR8vXkwL6MpGBAp38AMJT3l-DUrTDORM8HAWGFmPYdA1Eqnmc4R1f10NH2O3-4WBdsXPyvIaF2SGXw0dqlsJYFfruI07PgXQW8V7sw1BkCcuUMDlOC2mOLMCpnwoAt8-jDXXgzVSFF-RBN3c7M2x8JmVr0FRLBEUQCbYsq08RaDFSf_fe2vxZyr6YEorb0zOyLL499gEe91n2MFEyrzZx3R6Dg646CwIVg0rZiONhT70DziRaSFB7jTMH2z-rXD_2HseVsNHZ_svgfzq-DQcXFEJXG5wdOSjKH7ONKQQ-d5ZhZsZHOWH7l6Q1_1Oh2e5xLFA"
	//var ts string
	c.OnHTML(".image", func(e *colly.HTMLElement){
		s := e.ChildAttr("a", "href")
		_, err := os.Stat("images/" + str[i] + "/" + filepath.Base(s))
		if os.IsNotExist(err){
			resp, err := http.Get(s)
			if err != nil{
				fmt.Println(err.Error())
				return
			}
			buf, err := ioutil.ReadAll(resp.Body)
			if err != nil{
				fmt.Println(err.Error())
				return
			}
			ioutil.WriteFile("images/" + str[i] + "/" + filepath.Base(s), buf, 0644)
			fmt.Println("One Ok!")
			resp.Body.Close()
		} else {
			fmt.Println("One Already Download.")
		}
	})
	c.OnRequest(func(e * colly.Request){
		fmt.Println(e.URL)
	})
	for i = 0; i < lens;i++{
		//fmt.Println("输入新的代表网址：")
		//fmt.Scanf("%s", &ts)
		//u, _ := url.Parse(ts)
		//fmt.Println(u.Query())
		//sc = u.Query().Get("sc")
		for p := 1; p <= 10; p++{
			err := c.Visit("https://www.webcrawler.com/serp?qc=images&q=" + str[i] + "&page=" + strconv.Itoa(p) + "&sc=" + sc)
			if err != nil{
				log.Fatal(err.Error())
			}
		}
	}
}

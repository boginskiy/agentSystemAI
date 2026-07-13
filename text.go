package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	url := "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"
	method := "POST"

	payload := strings.NewReader(`{
									"model": "GigaChat",
									"messages": [
										{
										"role": "system",
										"content": "Ты — профессиональный переводчик на английский язык. Переведи точно сообщение пользователя."
										},
										{
										"role": "user",
										"content": "Привет мир!"
										}],
									"stream": false,
									"update_interval": 0
								}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJjdHkiOiJqd3QiLCJlbmMiOiJBMjU2Q0JDLUhTNTEyIiwiYWxnIjoiUlNBLU9BRVAtMjU2In0.ucu-CtyRE_-4h5DyopdViZvCC02RuHTPnAQ8D3PZ0ZbWhKL3ahiRhfxfPHOgrZYvUpaKsT7IaYVOd-3Srofa7fuDjWC-9VQzoSFDtqmkh4FWAsdAFMnuUTCJMpvwgPf_yv4Nwl_dcBkgY7I0NL-M0aR7Lv1jTQgOYUopqtdcUryWXFpPv9pPUt3BOM_cuOc7htDa6pG4D2FnrYtjWwyQZnxsOiF0HY2_Z3iUOeApEpC4VjBRUrvqHDzDmYdQFyPGV8Ve12LCJVmAikUP6ELBe2xngvcHJveNlHwC83p81a5Ozh4c5PGdEA5JfpmddG2JLOjwMk3T-9_yx_OPPSBjvQ.Yp70l9YqGUHPwIfa9ZT7vg.uBbOZ7oL2i0M6cvB25_4vjVTc8EFnhKi39qX9tF7yJpiPk13XNRIIEPBfWgIKHmzT9H1GKLIRKazNCcEn138uvfUs3Sh3L7U-wz-XMhhuGQidSLCQz7y9rJy6C4YUWTxzPulHWEfPQ86dDWqLEM9CfJZrlG4PJcaxa_BZow4Y3I5CEUCiC8HJd8B1oCDjGQkaRN18Oyv2z67vk0b06TeNhwLW5wJpp3ErHqS4HxXFNvv_ZOCnaKuq3rgfGujE0Bk-0ZOlQbetT8Ls8gJjh4B8ijfiAz8v8yeNqURMsdu-LNotWGhKzhSZfu--uT5Bi2fh5uHvdOKAQv7iRa47e_3m3UDkDsjFRch40K2KlhlDqGm7W1dfMKrYc7fly0CgoUx2vLpoIfQtIeMN-Cb1oE5WKyPBv9r_XWdoY1kGs3DjSatWPZDoGrcW8P8xDxg0uR50u1kJLvHHA48ttO2O1slJh3WWT3N38LXP11MwciinfZ-jrDgKGz1n5uy1_18TLm--yoJVXsyrbqAW2FluOYlQLHWSSUiEHtPDn9RNM-WMOoEr-pplJMHJgn0hhhGgWOgHfgwPZN3Zt4kPly9Uri28maTpMMt34rDqxn9Ja7Mi-iNnLWJDD7tidXeeMVYkteFdIAXYgMnADcPFOLp7Lvtt8jilPmqOSoHrrpPsx_NObuqrNOh6Wx05YnioYU3XPkU0w9tiDn1TH4RkfHwDkCiS_dV9RUow9LhA70YDGmrL3I.J8xH-00brTvKP3fT-QxqzWXx1O3u1nk3-8URbdsvdIM")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

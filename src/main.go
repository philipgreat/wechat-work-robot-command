// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "encoding/json"
import "net/http"
import "strings"
import "log"
import "flag"
import "io/ioutil"
type Text struct {
    Content    	string `json:"content"`
   
    // many more fields…
}

type Message struct {
    MessageType    	string `json:"msgtype"`
    Text 		Text `json:"text"`
	
    // many more fields…
}

func readFromTextFile(fileName string) (string, error) {

	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		return "", err
	}
	str := string(b) // convert content to a 'string'
	return str, nil

}
func readContent(messageOrFile string) (string, error) {
	if strings.HasPrefix(messageOrFile, "@") {
		return readFromTextFile(messageOrFile[1:])
	}
	return messageOrFile, nil
}

func buildMessage(messageType string,bodyText string) string{


	text := Text{Content: bodyText}

	message:= &Message{
		MessageType: messageType,
		Text: text,
	}
	requestBody, _ := json.Marshal(message)
	return string(requestBody)
}

func sendRequest(robotKey string,body string) string {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", robotKey)
	resp, err := http.Post(url, "application/json",
	strings.NewReader(body))
	if err != nil {
			log.Fatal(err)
	}
	var res map[string]interface{}

    json.NewDecoder(resp.Body).Decode(&res)

    fmt.Println(res["json"])
	return ""
}
func main() {
	robotKey := flag.String("robotkey", "---", "robot key from robot property")
	content := flag.String("content", "hello", "a string to show your content")
	flag.Parse()
	fmt.Println("robotKey:", *robotKey)
	fmt.Println("content:", *content)

	if *robotKey == "---" {
		fmt.Println("Please use rotbot key flag to replace your robot key")
		return
	}


	body, err := readContent(*content)
	if err != nil {
		log.Panic(err)
	}

	requestBody  := buildMessage("text",body)
	sendRequest(*robotKey,requestBody)
    fmt.Println(string(requestBody))
	
}
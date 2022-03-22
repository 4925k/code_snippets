package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	MISP       = "https://202.51.1.168:443"
	APITOKEN   = "QnX9bagdXfbdWbmQJAxTzO7v2GTdLt8LikzUD75H"
	CONFIGFILE = "config.cfg"
)

func main() {
	flagFile := flag.String("file", "invalid", "give file name that contains the ips")
	flag.Parse()

	if *flagFile == "invalid" {
		log.Fatal("give me a file using --file=<file-name>")
	}

	f, err := os.Open(*flagFile)
	if err != nil {
		log.Fatalf("cannot open file: %s", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ip := scanner.Text()
		fmt.Printf("Query for %s:\n", ip)
		resp, err := MispQuery(ip)
		if err != nil {
			fmt.Println(resp)
		}

		data, _ := formatJSON(resp)
		fmt.Println(string(data))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// MispQuery takes the URL/IP to query, performs query, and returns it
func MispQuery(ip string) ([]byte, error) {

	// specifying that we do not want to verify the SSL certificate
	// since its mostly self signed
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}

	// creatint the request instance
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/attributes/restSearch/json", MISP), nil)
	if err != nil {
		log.Println("ERROR could not create a request type for MISP instance due to error", err)
		return []byte{}, err
	}
	req.Header.Add("Authorization", APITOKEN) // The API key
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	// Pass the insecure certificate configuraiton in Client type
	c := &http.Client{Transport: tr, Timeout: 3 * time.Second}

	value := fmt.Sprintf(`{"value":"%s"}`, ip)

	buf := bytes.NewBufferString(value)
	req.Body = ioutil.NopCloser(bytes.NewBuffer(buf.Bytes()))

	// Actually Make the request
	res, err := c.Do(req)

	if err != nil {
		log.Printf("ERROR could not request MISP %T %v", err, err)
		return []byte{0}, err
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("ERROR could not read MISP response", err)
		return []byte{0}, err
	}
	return content, nil
}

func formatJSON(data []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, data, "", "    ")
	if err == nil {
		return out.Bytes(), err
	}
	return data, nil
}

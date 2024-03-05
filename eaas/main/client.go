// (c) Copyright 2022 Hewlett Packard Enterprise Development LP
// @author gabe@hpe.com - 2022-10-13
package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

func main() {

    //url := "https://chopchop.ftc.hpecorp.net/black-hole/event-horizon/enter"
    //method := "POST"
    //payload := strings.NewReader("{\n  \"rootElement\": \"HP_DeviceIndication\",\n  \"namespace\": \"http://schemas.hp.com/wbem/wscim/1/cim-schema/2/HP_DeviceIndication\"\n}")

    url := "https://eaas-fut1-sbs.corp.hpecorp.net/eaas-2.1/analyzeData"
    method := "POST"
    payload := strings.NewReader("Value0=%3Ca%20%2F%3E&Value1=--template-name&Value1=cecV5XmlTemplate&Value1=--use-rule&Value1=mapToSI_dev")

    client := &http.Client{}
    req, err := http.NewRequest(method, url, payload)

    if err != nil {
        fmt.Println(err)
        return
    }
    //req.Header.Add("Content-Type", "application/json")
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("X-HP-SBS-ApplicationId", "eaas-demo")
    req.Header.Add("X-HP-SBS-ApplicationKey", "AppKey_REDACTED__EDIT_ME!")
    req.Header.Add("X-HP-SBS-TestRequest", "false")
    req.Header.Add("X-HP-SBS-IsHPIApplication", "false")
    req.Header.Add("X-HP-SBS-SessionContext", "aaaaaaaa-bbbb-4ccc-dddd-eeeeeeeeeeee")

    res, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    dst := &bytes.Buffer{}

    if err := json.Indent(dst, body, "", "  "); err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(dst.String())

}

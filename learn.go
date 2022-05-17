package main

import "fmt"

type person struct {
	greet     func()
	firstName string
	lastName  string
}

func main() {

	rohan := person{
		greet:     func() { fmt.Println("Hello rohan How are you!") },
		firstName: "Bhupendra",
	}
	bhupendra := person{
		greet: func() { fmt.Println("Hello Bhupendra How are you!") },
	}
	rohan.greet()
	bhupendra.greet()
	fmt.Println(rohan.firstName)

	// values := map[string]string{"name": "John Doe", "occupation": "gardener"}
	// json_data, err := json.Marshal(values)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// resp, err := http.Post("https://httpbin.org/post", "application/json",
	// 	bytes.NewBuffer(json_data))

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var res map[string]interface{}

	// json.NewDecoder(resp.Body).Decode(&res)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	// fmt.Println("response Status:", resp.Status)

	// fmt.Println(res)
}

// code for chrun go api
// jsonValues, _ := json.Marshal(formdataMapping)
// req, err := http.NewRequest("POST", "https://analytics.churnzero.net/i", bytes.NewBuffer(jsonValues))
// req.Header.Set("Accept", "application/json")
// req.Header.Set("Content-Type", "application/json")
// client := &http.Client{}
// resp, err := client.Do(req)
// if err != nil {
// 	panic(err)
// }
// defer resp.Body.Close()
// fmt.Println("response Status:", resp.Status)
// fmt.Println("response Headers:", resp.Header)
// body, _ := ioutil.ReadAll(resp.Body)
// fmt.Println("response Body:", string(body))

// client := &http.Client{}
// req, err := http.NewRequest("POST", "https://analytics.churnzero.net/i", nil)
// if err != nil {
// 	fmt.Print(err.Error())
// }
// req.Header.Add("Accept", "application/json")
// req.Header.Add("Content-Type", "application/json")

// res, err := client.Do(req)
// if err != nil {
// 	fmt.Print(err.Error())
// }
// defer res.Body.Close()
// body, err := ioutil.ReadAll(res.Body)
// if err != nil {
// 	fmt.Print(err.Error())
// }
// var dataObj data
// json.Unmarshal(body, &dataObj)
// fmt.Printf("API Response as struct %+v\n", dataObj)
// fmt.Println("response Status:", body.Status())
// fmt.Println("response Headers:", res.HttpResponse().Header)
// fmt.Println("response Body:", res.RawText())

//hitesh

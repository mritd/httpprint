/*
 * Copyright 2018 mritd <mritd1234@gmail.com>.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", printDetail)
	http.ListenAndServe(":8080", nil)
}

func printDetail(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("Proto: %s\n", r.Proto)
	fmt.Printf("RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Printf("Request URL: %s\n", r.URL)
	fmt.Println("Request Header:")
	for k, v := range r.Header {
		fmt.Printf("  » %-20s%s\n", k, v)
	}

	fmt.Println("Request From:")
	for k, v := range r.PostForm {
		fmt.Printf("  » %-20s%s\n", k, v)
	}

	w.Write([]byte("success"))
}

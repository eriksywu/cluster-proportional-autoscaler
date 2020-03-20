/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package autoscaler

import (
	"fmt"
	"net/http"

	"github.com/golang/glog"
)

func (s *AutoScaler) startHealthz() {
	http.HandleFunc("/healthz", s.healthFn)
	glog.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *AutoScaler) healthFn(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get("CheckConfigMap") == "" {
		return
	}
	if _, err := s.syncConfigWithServer(); err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("Encountered error checking config map: %v", err)))
		return
	}
}

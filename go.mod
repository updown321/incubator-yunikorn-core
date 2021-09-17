//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

module github.com/apache/incubator-yunikorn-core

go 1.15

require (
	github.com/HdrHistogram/hdrhistogram-go v1.0.1 // indirect
	github.com/apache/incubator-yunikorn-scheduler-interface v0.11.1-0.20210825113556-88f21e44dc01
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/google/btree v1.0.1
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.7.3
	github.com/looplab/fsm v0.1.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v0.9.4
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.4.1
	github.com/prometheus/procfs v0.0.8 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible
	go.uber.org/atomic v1.5.1 // indirect
	go.uber.org/multierr v1.4.0 // indirect
	go.uber.org/zap v1.13.0
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e
	golang.org/x/sys v0.0.0-20200413165638-669c56c373c4 // indirect
	golang.org/x/tools v0.0.0-20200415000939-92398ad77b89 // indirect
	google.golang.org/grpc v1.26.0
	gopkg.in/yaml.v2 v2.2.8
	gotest.tools v2.2.0+incompatible
	honnef.co/go/tools v0.0.1-2020.1.3 // indirect
	k8s.io/api v0.16.13
	k8s.io/apimachinery v0.16.13
	k8s.io/client-go v0.16.13
	k8s.io/klog v1.0.0
	k8s.io/kubernetes v1.16.13
)

replace (
	k8s.io/api => k8s.io/api v0.16.13
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.16.13
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.13
	k8s.io/apiserver => k8s.io/apiserver v0.16.13
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.16.13
	k8s.io/client-go => k8s.io/client-go v0.16.13
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.16.13
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.16.13
	k8s.io/code-generator => k8s.io/code-generator v0.16.13
	k8s.io/component-base => k8s.io/component-base v0.16.13
	k8s.io/cri-api => k8s.io/cri-api v0.16.13
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.16.13
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.16.13
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.16.13
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.16.13
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.16.13
	k8s.io/kubectl => k8s.io/kubectl v0.16.13
	k8s.io/kubelet => k8s.io/kubelet v0.16.13
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.16.13
	k8s.io/metrics => k8s.io/metrics v0.16.13
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.16.13
	vbom.ml/util => github.com/fvbommel/util v0.0.0-20160121211510-db5cfe13f5cc
)

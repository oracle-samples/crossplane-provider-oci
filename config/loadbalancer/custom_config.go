/*
 * Copyright (c) 2023 Oracle and/or its affiliates
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

package loadbalancer

import (
	"context"
	"github.com/oracle/provider-oci/config/common"
)

func GetBackendSetIDFunc(ctx context.Context, externalName string, parameters map[string]any,
	providerConfig map[string]any) (string, error) {
	attrsMap := make(map[string]string)
	attrsMap["loadBalancers"] = ParamLoadBalancerID
	attrsMap["backendSets"] = "name"
	return common.GetIDFromParamsFunc(parameters, attrsMap)
}

func GetBackendIDFunc(ctx context.Context, externalName string, parameters map[string]any,
	providerConfig map[string]any) (string, error) {
	attrsMap := make(map[string]string)
	attrsMap["loadBalancers"] = ParamLoadBalancerID
	attrsMap["backendSets"] = "backendset_name"
	attrsMap["backends"] = "name"
	return common.GetIDFromParamsFunc(parameters, attrsMap)
}

func GetLBHostnameIDFunc(ctx context.Context, externalName string, parameters map[string]any,
	providerConfig map[string]any) (string, error) {
	attrsMap := make(map[string]string)
	attrsMap["loadBalancers"] = ParamLoadBalancerID
	attrsMap["hostnames"] = "name"
	return common.GetIDFromParamsFunc(parameters, attrsMap)
}

func GetRuleSetIDFunc(ctx context.Context, externalName string, parameters map[string]any,
	providerConfig map[string]any) (string, error) {
	attrsMap := make(map[string]string)
	attrsMap["loadBalancers"] = ParamLoadBalancerID
	attrsMap["ruleSets"] = "name"
	return common.GetIDFromParamsFunc(parameters, attrsMap)
}

func GetCertificateIDFunc(ctx context.Context, externalName string, parameters map[string]any,
	providerConfig map[string]any) (string, error) {
	attrsMap := make(map[string]string)
	attrsMap["loadBalancers"] = ParamLoadBalancerID
	attrsMap["certificates"] = "certificate_name"
	return common.GetIDFromParamsFunc(parameters, attrsMap)
}

func GetRoutingPolicyIDFunc(ctx context.Context, externalName string, parameters map[string]any,
	providerConfig map[string]any) (string, error) {
	attrsMap := make(map[string]string)
	attrsMap["loadBalancers"] = ParamLoadBalancerID
	attrsMap["routingPolicies"] = "name"
	return common.GetIDFromParamsFunc(parameters, attrsMap)
}

func GetPathRouteSetIDFunc(ctx context.Context, externalName string, parameters map[string]any,
	providerConfig map[string]any) (string, error) {
	attrsMap := make(map[string]string)
	attrsMap["loadBalancers"] = ParamLoadBalancerID
	attrsMap["pathRouteSets"] = "name"
	return common.GetIDFromParamsFunc(parameters, attrsMap)
}

func GetListenerIDFunc(ctx context.Context, externalName string, parameters map[string]any,
	providerConfig map[string]any) (string, error) {
	attrsMap := make(map[string]string)
	attrsMap["loadBalancers"] = ParamLoadBalancerID
	attrsMap["listeners"] = "name"
	return common.GetIDFromParamsFunc(parameters, attrsMap)
}

func GetSSLCipherSuiteIDFunc(ctx context.Context, externalName string, parameters map[string]any,
	providerConfig map[string]any) (string, error) {
	attrsMap := make(map[string]string)
	attrsMap["loadBalancers"] = ParamLoadBalancerID
	attrsMap["sslCipherSuites"] = "name"
	return common.GetIDFromParamsFunc(parameters, attrsMap)
}

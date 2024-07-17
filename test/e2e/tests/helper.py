# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
#	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

"""Helper functions for apigatewayv2 tests
"""
from acktest.k8s import resource

from e2e import CRD_GROUP, CRD_VERSION, load_apigatewayv2_resource

API_RESOURCE_PLURAL = 'apis'
INTEGRATION_RESOURCE_PLURAL = 'integrations'
AUTHORIZER_RESOURCE_PLURAL = 'authorizers'
ROUTE_RESOURCE_PLURAL = 'routes'
STAGE_RESOURCE_PLURAL = 'stages'
VPC_LINK_RESOURCE_PLURAL = 'vpclinks'
API_MAPPING_RESOURCE_PLURAL = 'apimappings'
DOMAIN_NAME_RESOURCE_PLURAL = 'domainnames'


def api_ref_and_data(api_resource_name: str, replacement_values: dict, file_name: str = "httpapi"):
    ref = resource.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, API_RESOURCE_PLURAL,
        api_resource_name, namespace="default",
    )

    resource_data = load_apigatewayv2_resource(
        file_name,
        additional_replacements=replacement_values,
    )
    return ref, resource_data


def import_api_ref_and_data(api_resource_name: str, replacement_values: dict, file_name: str = "import_api"):
    ref = resource.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, API_RESOURCE_PLURAL,
        api_resource_name, namespace="default",
    )

    resource_data = load_apigatewayv2_resource(
        file_name,
        additional_replacements=replacement_values,
    )
    return ref, resource_data


def integration_ref_and_data(integration_resource_name: str, replacement_values: dict, file_name: str = "integration"):
    ref = resource.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, INTEGRATION_RESOURCE_PLURAL,
        integration_resource_name, namespace="default",
    )

    resource_data = load_apigatewayv2_resource(
        file_name,
        additional_replacements=replacement_values,
    )
    return ref, resource_data


def authorizer_ref_and_data(authorizer_resource_name: str, replacement_values: dict, file_name: str = "authorizer"):
    ref = resource.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, AUTHORIZER_RESOURCE_PLURAL,
        authorizer_resource_name, namespace="default",
    )

    resource_data = load_apigatewayv2_resource(
        file_name,
        additional_replacements=replacement_values,
    )
    return ref, resource_data


def route_ref_and_data(route_resource_name: str, replacement_values: dict, file_name: str = "route"):
    ref = resource.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, ROUTE_RESOURCE_PLURAL,
        route_resource_name, namespace="default",
    )

    resource_data = load_apigatewayv2_resource(
        file_name,
        additional_replacements=replacement_values,
    )
    return ref, resource_data


def stage_ref_and_data(stage_resource_name: str, replacement_values: dict, file_name: str = "stage"):
    ref = resource.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, STAGE_RESOURCE_PLURAL,
        stage_resource_name, namespace="default",
    )

    resource_data = load_apigatewayv2_resource(
        file_name,
        additional_replacements=replacement_values,
    )
    return ref, resource_data


def vpc_link_ref_and_data(vpc_link_resource_name: str, replacement_values: dict, file_name: str = "vpc-link"):
    ref = resource.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, VPC_LINK_RESOURCE_PLURAL,
        vpc_link_resource_name, namespace="default",
    )

    resource_data = load_apigatewayv2_resource(
        file_name,
        additional_replacements=replacement_values,
    )
    return ref, resource_data

def api_mapping_ref_and_data(api_mapping_resource_name: str, replacement_values: dict, file_name: str = "api-mapping"):
    ref = resource.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, API_MAPPING_RESOURCE_PLURAL,
        api_mapping_resource_name, namespace="default",
    )

    resource_data = load_apigatewayv2_resource(
        file_name,
        additional_replacements=replacement_values,
    )
    return ref, resource_data

def domain_name_ref_and_data(domain_name_resource_name: str, replacement_values: dict, file_name: str = "domain-name"):
    ref = resource.CustomResourceReference(
        CRD_GROUP, CRD_VERSION, DOMAIN_NAME_RESOURCE_PLURAL,
        domain_name_resource_name, namespace="default",
    )

    resource_data = load_apigatewayv2_resource(
        file_name,
        additional_replacements=replacement_values,
    )
    return ref, resource_data


class ApiGatewayValidator:

    def __init__(self, apigatewayv2_client):
        self.apigatewayv2_client = apigatewayv2_client

    def assert_api_is_present(self, api_id: str):
        aws_res = self.apigatewayv2_client.get_api(ApiId=api_id)
        assert aws_res is not None

    def assert_integration_is_present(self, api_id: str, integration_id: str):
        aws_res = self.apigatewayv2_client.get_integration(ApiId=api_id, IntegrationId=integration_id)
        assert aws_res is not None

    def assert_authorizer_is_present(self, api_id: str, authorizer_id: str):
        aws_res = self.apigatewayv2_client.get_authorizer(ApiId=api_id, AuthorizerId=authorizer_id)
        assert aws_res is not None

    def assert_route_is_present(self, api_id: str, route_id: str):
        aws_res = self.apigatewayv2_client.get_route(ApiId=api_id, RouteId=route_id)
        assert aws_res is not None

    def assert_stage_is_present(self, api_id: str, stage_name: str):
        aws_res = self.apigatewayv2_client.get_stage(ApiId=api_id, StageName=stage_name)
        assert aws_res is not None

    def assert_vpc_link_is_present(self, vpc_link_id: str):
        aws_res = self.apigatewayv2_client.get_vpc_link(VpcLinkId=vpc_link_id)
        assert aws_res is not None

    def assert_api_is_deleted(self, api_id: str):
        res_found = False
        try:
            self.apigatewayv2_client.get_api(ApiId=api_id)
            res_found = True
        except self.apigatewayv2_client.exceptions.NotFoundException:
            pass

        assert res_found is False

    def assert_integration_is_deleted(self, api_id: str, integration_id: str):
        res_found = False
        try:
            self.apigatewayv2_client.get_integration(ApiId=api_id, IntegrationId=integration_id)
            res_found = True
        except self.apigatewayv2_client.exceptions.NotFoundException:
            pass

        assert res_found is False

    def assert_authorizer_is_deleted(self, api_id: str, authorizer_id: str):
        res_found = False
        try:
            self.apigatewayv2_client.get_authorizer(ApiId=api_id, AuthorizerId=authorizer_id)
            res_found = True
        except self.apigatewayv2_client.exceptions.NotFoundException:
            pass

        assert res_found is False

    def assert_route_is_deleted(self, api_id: str, route_id: str):
        res_found = False
        try:
            self.apigatewayv2_client.get_route(ApiId=api_id, RouteId=route_id)
            res_found = True
        except self.apigatewayv2_client.exceptions.NotFoundException:
            pass

        assert res_found is False

    def assert_stage_is_deleted(self, api_id: str, stage_name: str):
        res_found = False
        try:
            self.apigatewayv2_client.get_stage(ApiId=api_id, StageName=stage_name)
            res_found = True
        except self.apigatewayv2_client.exceptions.NotFoundException:
            pass

        assert res_found is False

    def assert_vpc_link_is_deleted(self, vpc_link_id: str):
        res_found = False
        try:
            self.apigatewayv2_client.get_vpc_link(VpcLinkId=vpc_link_id)
            res_found = True
        except self.apigatewayv2_client.exceptions.NotFoundException:
            pass

        assert res_found is False

    def assert_api_name(self, api_id, expected_api_name):
        aws_res = self.apigatewayv2_client.get_api(ApiId=api_id)
        assert aws_res is not None
        assert aws_res['Name'] == expected_api_name

    def assert_integration_uri(self, api_id, integration_id, expected_uri):
        aws_res = self.apigatewayv2_client.get_integration(ApiId=api_id, IntegrationId=integration_id)
        assert aws_res is not None
        assert aws_res['IntegrationUri'] == expected_uri

    def assert_authorizer_name(self, api_id, authorizer_id, expected_authorizer_name):
        aws_res = self.apigatewayv2_client.get_authorizer(ApiId=api_id, AuthorizerId=authorizer_id)
        assert aws_res is not None
        assert aws_res['Name'] == expected_authorizer_name

    def assert_route_key(self, api_id, route_id, expected_route_key):
        aws_res = self.apigatewayv2_client.get_route(ApiId=api_id, RouteId=route_id)
        assert aws_res is not None
        assert aws_res['RouteKey'] == expected_route_key

    def assert_stage_description(self, api_id, stage_name, expected_description):
        aws_res = self.apigatewayv2_client.get_stage(ApiId=api_id, StageName=stage_name)
        assert aws_res is not None
        assert aws_res['Description'] == expected_description

    def assert_vpc_link_name(self, vpc_link_id, expected_vpc_link_name):
        aws_res = self.apigatewayv2_client.get_vpc_link(VpcLinkId=vpc_link_id)
        assert aws_res is not None
        assert aws_res['Name'] == expected_vpc_link_name

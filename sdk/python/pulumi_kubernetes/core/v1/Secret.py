# *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
from typing import Optional

import pulumi
import pulumi.runtime
from pulumi import Input, ResourceOptions

from ... import tables, version


class Secret(pulumi.CustomResource):
    """
    Secret holds secret data of a certain type. The total bytes of the values in the Data field must
    be less than MaxSecretSize bytes.
    """

    apiVersion: pulumi.Output[str]
    """
    APIVersion defines the versioned schema of this representation of an object. Servers should
    convert recognized schemas to the latest internal value, and may reject unrecognized values.
    More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
    """

    kind: pulumi.Output[str]
    """
    Kind is a string value representing the REST resource this object represents. Servers may infer
    this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More
    info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
    """

    data: pulumi.Output[dict]
    """
    Data contains the secret data. Each key must consist of alphanumeric characters, '-', '_' or
    '.'. The serialized form of the secret data is a base64 encoded string, representing the
    arbitrary (possibly non-string) data value here. Described in
    https://tools.ietf.org/html/rfc4648#section-4
    """

    metadata: pulumi.Output[dict]
    """
    Standard object's metadata. More info:
    https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
    """

    string_data: pulumi.Output[dict]
    """
    stringData allows specifying non-binary secret data in string form. It is provided as a
    write-only convenience method. All keys and values are merged into the data field on write,
    overwriting any existing values. It is never output when reading from the API.
    """

    type: pulumi.Output[str]
    """
    Used to facilitate programmatic handling of secret data.
    """

    def __init__(self, resource_name, opts=None, data=None, metadata=None, string_data=None, type=None, __name__=None, __opts__=None):
        """
        Create a Secret resource with the given unique name, arguments, and options.

        :param str resource_name: The _unique_ name of the resource.
        :param pulumi.ResourceOptions opts: A bag of options that control this resource's behavior.
        :param pulumi.Input[dict] data: Data contains the secret data. Each key must consist of alphanumeric characters, '-',
               '_' or '.'. The serialized form of the secret data is a base64 encoded string,
               representing the arbitrary (possibly non-string) data value here. Described in
               https://tools.ietf.org/html/rfc4648#section-4
        :param pulumi.Input[dict] metadata: Standard object's metadata. More info:
               https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        :param pulumi.Input[dict] string_data: stringData allows specifying non-binary secret data in string form. It is provided as
               a write-only convenience method. All keys and values are merged into the data field
               on write, overwriting any existing values. It is never output when reading from the
               API.
        :param pulumi.Input[str] type: Used to facilitate programmatic handling of secret data.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        __props__['apiVersion'] = 'v1'
        __props__['kind'] = 'Secret'
        __props__['data'] = data
        __props__['metadata'] = metadata
        __props__['stringData'] = string_data
        __props__['type'] = type

        __props__['status'] = None

        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(version=version.get_version()))

        super(Secret, self).__init__(
            "kubernetes:core/v1:Secret",
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None):
        """
        Get the state of an existing `Secret` resource, as identified by `id`.
        The ID is of the form `[namespace]/[name]`; if `[namespace]` is omitted,
        then (per Kubernetes convention) the ID becomes `default/[name]`.

        Pulumi will keep track of this resource using `resource_name` as the Pulumi ID.

        :param str resource_name: _Unique_ name used to register this resource with Pulumi.
        :param pulumi.Input[str] id: An ID for the Kubernetes resource to retrieve.
               Takes the form `[namespace]/[name]` or `[name]`.
        :param Optional[pulumi.ResourceOptions] opts: A bag of options that control this
               resource's behavior.
        """
        opts = ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))
        return Secret(resource_name, opts)

    def translate_output_property(self, prop: str) -> str:
        return tables._CASING_FORWARD_TABLE.get(prop) or prop

    def translate_input_property(self, prop: str) -> str:
        return tables._CASING_BACKWARD_TABLE.get(prop) or prop

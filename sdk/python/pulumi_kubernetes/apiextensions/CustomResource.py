import pulumi
import pulumi.runtime

from ... import tables

class CustomResource(pulumi.CustomResource):
    """
    CustomResource represents an instance of a CustomResourceDefinition (CRD). For example, the
    CoreOS Prometheus operator exposes a CRD `monitoring.coreos.com/ServiceMonitor`; to
    instantiate this as a Pulumi resource, one could call `new CustomResource`, passing the
    `ServiceMonitor` resource definition as an argument.
    """
    def __init__(self, __name__, __opts__=None, apiVersion=None, kind=None, metadata=None, properties=None):
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, str):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not apiVersion:
            raise TypeError('Missing required property apiVersion')
        __props__['apiVersion'] = apiVersion
        if not kind:
            raise TypeError('Missing required property kind')
        __props__['kind'] = kind
        __props__['metadata'] = metadata

        for k, v in properties.items():
            __props__[k] = v

        super(CustomResourceDefinition, self).__init__(
            "kubernetes:%s:%s" % (apiVersion, kind),
            __name__,
            __props__,
            __opts__)

    def translate_output_property(self, prop: str) -> str:
        return tables._CASING_FORWARD_TABLE.get(prop) or prop

    def translate_input_property(self, prop: str) -> str:
        return tables._CASING_BACKWARD_TABLE.get(prop) or prop

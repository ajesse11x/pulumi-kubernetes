// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Core.V1
{
    /// <summary>
    /// ConfigMap holds configuration data for pods to consume.
    /// </summary>
    public partial class ConfigMap : KubernetesResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers
        /// should convert recognized schemas to the latest internal value, and may reject
        /// unrecognized values. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// BinaryData contains the binary data. Each key must consist of alphanumeric characters,
        /// '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range.
        /// The keys stored in BinaryData must not overlap with the ones in the Data field, this is
        /// enforced during validation process. Using this field will require 1.10+ apiserver and
        /// kubelet.
        /// </summary>
        [Output("binaryData")]
        public Output<ImmutableDictionary<string, string>> BinaryData { get; private set; } = null!;

        /// <summary>
        /// Data contains the configuration data. Each key must consist of alphanumeric characters,
        /// '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The
        /// keys stored in Data must not overlap with the keys in the BinaryData field, this is
        /// enforced during validation process.
        /// </summary>
        [Output("data")]
        public Output<ImmutableDictionary<string, string>> Data { get; private set; } = null!;

        /// <summary>
        /// Immutable, if set to true, ensures that data stored in the ConfigMap cannot be updated
        /// (only object metadata can be modified). If not set to true, the field can be modified at
        /// any time. Defaulted to nil. This is an alpha field enabled by ImmutableEphemeralVolumes
        /// feature gate.
        /// </summary>
        [Output("immutable")]
        public Output<bool> Immutable { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers
        /// may infer this from the endpoint the client submits requests to. Cannot be updated. In
        /// CamelCase. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard object's metadata. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
        /// </summary>
        [Output("metadata")]
        public Output<Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;


        /// <summary>
        /// Create a ConfigMap resource with the given unique name, arguments, and options.
        /// </summary>
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConfigMap(string name, Types.Inputs.Core.V1.ConfigMapArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:core/v1:ConfigMap", name, SetAPIKindAndVersion(args), options)
        {
        }

        internal ConfigMap(string name, ImmutableDictionary<string, object?> dictionary, CustomResourceOptions? options = null)
            : base("kubernetes:core/v1:ConfigMap", name, dictionary, options)
        {
        }

        private static ResourceArgs SetAPIKindAndVersion(Types.Inputs.Core.V1.ConfigMapArgs? args)
        {
            args ??= new Types.Inputs.Core.V1.ConfigMapArgs();
            args.ApiVersion = "v1";
            args.Kind = "ConfigMap";
            return args;
        }

        /// <summary>
        /// Get an existing ConfigMap resource's state with the given name and ID.
        /// </summary>
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConfigMap Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConfigMap(name, default(Types.Inputs.Core.V1.ConfigMapArgs),
                CustomResourceOptions.Merge(options, new CustomResourceOptions {Id = id}));
        }

    }
}

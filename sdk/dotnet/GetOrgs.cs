// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Zitadel
{
    public static class GetOrgs
    {
        /// <summary>
        /// Datasource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.
        /// </summary>
        public static Task<GetOrgsResult> InvokeAsync(GetOrgsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrgsResult>("zitadel:index/getOrgs:getOrgs", args ?? new GetOrgsArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing an organization in ZITADEL, which is the highest level after the instance and contains several other resource including policies if the configuration differs to the default policies on the instance.
        /// </summary>
        public static Output<GetOrgsResult> Invoke(GetOrgsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrgsResult>("zitadel:index/getOrgs:getOrgs", args ?? new GetOrgsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrgsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A domain of the org.
        /// </summary>
        [Input("domain")]
        public string? Domain { get; set; }

        /// <summary>
        /// Method for querying orgs by domain, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
        /// </summary>
        [Input("domainMethod")]
        public string? DomainMethod { get; set; }

        /// <summary>
        /// Name of the org.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Method for querying orgs by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
        /// </summary>
        [Input("nameMethod")]
        public string? NameMethod { get; set; }

        /// <summary>
        /// State of the org, supported values: ORG*STATE*UNSPECIFIED, ORG*STATE*ACTIVE, ORG*STATE*INACTIVE, ORG*STATE*REMOVED
        /// </summary>
        [Input("state")]
        public string? State { get; set; }

        public GetOrgsArgs()
        {
        }
        public static new GetOrgsArgs Empty => new GetOrgsArgs();
    }

    public sealed class GetOrgsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A domain of the org.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Method for querying orgs by domain, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
        /// </summary>
        [Input("domainMethod")]
        public Input<string>? DomainMethod { get; set; }

        /// <summary>
        /// Name of the org.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Method for querying orgs by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
        /// </summary>
        [Input("nameMethod")]
        public Input<string>? NameMethod { get; set; }

        /// <summary>
        /// State of the org, supported values: ORG*STATE*UNSPECIFIED, ORG*STATE*ACTIVE, ORG*STATE*INACTIVE, ORG*STATE*REMOVED
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public GetOrgsInvokeArgs()
        {
        }
        public static new GetOrgsInvokeArgs Empty => new GetOrgsInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrgsResult
    {
        /// <summary>
        /// A domain of the org.
        /// </summary>
        public readonly string? Domain;
        /// <summary>
        /// Method for querying orgs by domain, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
        /// </summary>
        public readonly string? DomainMethod;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A set of all organization IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// Name of the org.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Method for querying orgs by name, supported values: TEXT*QUERY*METHOD*EQUALS, TEXT*QUERY*METHOD*EQUALS*IGNORE*CASE, TEXT*QUERY*METHOD*STARTS*WITH, TEXT*QUERY*METHOD*STARTS*WITH*IGNORE*CASE, TEXT*QUERY*METHOD*CONTAINS, TEXT*QUERY*METHOD*CONTAINS*IGNORE*CASE, TEXT*QUERY*METHOD*ENDS*WITH, TEXT*QUERY*METHOD*ENDS*WITH*IGNORE*CASE
        /// </summary>
        public readonly string? NameMethod;
        /// <summary>
        /// Primary domain of the org
        /// </summary>
        public readonly string PrimaryDomain;
        /// <summary>
        /// State of the org, supported values: ORG*STATE*UNSPECIFIED, ORG*STATE*ACTIVE, ORG*STATE*INACTIVE, ORG*STATE*REMOVED
        /// </summary>
        public readonly string? State;

        [OutputConstructor]
        private GetOrgsResult(
            string? domain,

            string? domainMethod,

            string id,

            ImmutableArray<string> ids,

            string? name,

            string? nameMethod,

            string primaryDomain,

            string? state)
        {
            Domain = domain;
            DomainMethod = domainMethod;
            Id = id;
            Ids = ids;
            Name = name;
            NameMethod = nameMethod;
            PrimaryDomain = primaryDomain;
            State = state;
        }
    }
}

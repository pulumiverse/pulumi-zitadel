// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace scoretechnologies.Zitadel
{
    public static class Config
    {
        [global::System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("zitadel");

        private static readonly __Value<string?> _domain = new __Value<string?>(() => __config.Get("domain"));
        /// <summary>
        /// Domain used to connect to the ZITADEL instance
        /// </summary>
        public static string? Domain
        {
            get => _domain.Get();
            set => _domain.Set(value);
        }

        private static readonly __Value<bool?> _insecure = new __Value<bool?>(() => __config.GetBoolean("insecure"));
        /// <summary>
        /// Use insecure connection
        /// </summary>
        public static bool? Insecure
        {
            get => _insecure.Get();
            set => _insecure.Set(value);
        }

        private static readonly __Value<string?> _jwtProfileFile = new __Value<string?>(() => __config.Get("jwtProfileFile"));
        /// <summary>
        /// Path to the file containing credentials to connect to ZITADEL. Either 'jwt_profile_file' or 'jwt_profile_json' is
        /// required
        /// </summary>
        public static string? JwtProfileFile
        {
            get => _jwtProfileFile.Get();
            set => _jwtProfileFile.Set(value);
        }

        private static readonly __Value<string?> _jwtProfileJson = new __Value<string?>(() => __config.Get("jwtProfileJson"));
        /// <summary>
        /// JSON value of credentials to connect to ZITADEL. Either 'jwt_profile_file' or 'jwt_profile_json' is required
        /// </summary>
        public static string? JwtProfileJson
        {
            get => _jwtProfileJson.Get();
            set => _jwtProfileJson.Set(value);
        }

        private static readonly __Value<string?> _port = new __Value<string?>(() => __config.Get("port"));
        /// <summary>
        /// Used port if not the default ports 80 or 443 are configured
        /// </summary>
        public static string? Port
        {
            get => _port.Get();
            set => _port.Set(value);
        }

        private static readonly __Value<string?> _token = new __Value<string?>(() => __config.Get("token"));
        /// <summary>
        /// Path to the file containing credentials to connect to ZITADEL
        /// </summary>
        public static string? Token
        {
            get => _token.Get();
            set => _token.Set(value);
        }

    }
}

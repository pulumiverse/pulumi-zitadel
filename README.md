# Zitadel Resource Provider

The Zitadel Resource Provider lets you manage [Zitadel](https://zitadel.com/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @scoretechnologies/zitadel
```

or `yarn`:

```bash
yarn add @scoretechnologies/zitadel
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_zitadel
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/scoretechnologies/pulumi-zitadel/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package scoretechnologies.Zitadel
```

## Configuration

The following configuration points are available for the `zitadel` provider:

- `zitadel:domain` - domain used to connect to the ZITADEL instance
- `zitadel:insecure` - use insecure connection
- `zitadel:jwtProfileFile` - path to the file containing credentials to connect to ZITADEL. Either `jwtProfileFile` or `jwtProfileJson`
- `zitadel:jwtProfileJson` - JSON value of credentials to connect to ZITADEL. Either `jwtProfileFile` or `jwtProfileJson` is required
- `zitadel:port` - used port if not the default ports 80 or 443 are configured
- `zitadel:token` - path to the file containing credentials to connect to ZITADEL

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/zitadel/api-docs/).

# Zitadel Resource Provider

The Zitadel Resource Provider lets you manage [Zitadel](https://zitadel.com/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @vavsab/pulumi-zitadel
```

or `yarn`:

```bash
yarn add @vavsab/pulumi-zitadel
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi_zitadel
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/vavsab/pulumi-zitadel/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.Zitadel
```

## Configuration

The following configuration points are available for the `zitadel` provider:

- `zitadel:apiKey` (environment: `FOO_API_KEY`) - the API key for `zitadel`
- `zitadel:region` (environment: `FOO_REGION`) - the region in which to deploy resources

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/zitadel/api-docs/).

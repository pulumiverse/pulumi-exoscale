# Exoscale Resource Provider

The Exoscale Resource Provider lets you manage [Exoscale](https://www.exoscale.com/) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumiverse/exoscale
```

or `yarn`:

```bash
yarn add @pulumiverse/exoscale
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumiverse_exoscale
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/pulumiverse/pulumi-exoscale/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumiverse.Exoscale
```

## Configuration

The following configuration points are available for the `exoscale` provider:

- `exoscale:key` (environment: `EXOSCALE_API_KEY`) - Exoscale account API key
- `exoscale:secret` (environment: `EXOSCALE_API_SECRET`) - Exoscale account API secret
- `exoscale:timeout` - Global async operations waiting time in seconds (default: 300)

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/exoscale/api-docs/).

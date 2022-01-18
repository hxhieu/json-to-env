# Download

Download from the [Releases](https://github.com/hxhieu/json-to-env/releases/latest)

OR

```
go install github.com/hxhieu/json-to-env@latest
```


# Usage

```
Convert the JSON format to .env format

Usage:
  json-to-env [flags]

Flags:
  -h, --help               help for json-to-env
  -o, --output string      The output file (default ".env")
  -s, --separator string   The nested fields separator (default "__")
```

# Example

```
// appsettings.json

{
  "Logging": {
    "LogLevel": {
      "Default": "Debug",
      "Microsoft": "Information",
      "System": "Information",
      "Microsoft.EntityFrameworkCore.Database.Command": "Warning"
    },
    "UseDeveloperExceptionPage": "true"
  }
}

```

Output as 

```
# .env

Logging__UseDeveloperExceptionPage="true"
Logging__LogLevel__Default="Debug"
Logging__LogLevel__Microsoft="Information"
Logging__LogLevel__System="Information"
Logging__LogLevel__Microsoft_EntityFrameworkCore_Database_Command="Warning"
```

# Known issues
- JSON unmarshall fields ordering are not preserved
- JSON comments `//` and `/* */` are not supported so remove them in prior
- `:` as separator will cause weird output
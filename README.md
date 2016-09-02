# Dummy HTTP Server

| Request Path      | Response Status Code |
|-------------------|----------------------|
| /NNN              | NNN                  |
| _everything else_ | 200                  |

## Usage

`docker run -i -t -p 8080:8080 jasonrm/dummy-server`

## Options

```
Usage of ./dummy-server:
  -port int
        port number (default 8080)
```

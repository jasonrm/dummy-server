# Dummy HTTP Server

| Request Path      | Response Status Code |
|-------------------|----------------------|
| /                 | 200                  |
| /XXX              | XXX                  |
| _everything else_ | 000                  |

## Usage

`docker run -i -t -p 8080:8080 jasonrm/dummy-server`

## Options

```
Usage of ./dummy-server:
  -port int
        port number (default 8080)
```

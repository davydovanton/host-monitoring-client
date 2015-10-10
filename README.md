# Host-monitoring client
Simple implementation JSON http client witch returns system statistic (CPU and memory)

## Usage
Start server
```
$ bin/main
```

and after that:
```
$ curl -i localhost:8071/stats 
```

## Response
- `used_percent_per_core` - returns arrays with percent usage for each processor

```
{
  "memory": {
    "total": 17179869184,
    "available": 6500904960,
    "used": 16786092032,
    "used_percent": 62.15975284576416,
    "free": 393777152,
    "active": 8078532608,
    "inactive": 3784556544,
    "buffers": 0,
    "cached": 6107127808,
    "wired": 1887318016,
    "shared": 0
  },
  "cpu": {
    "used_percent_per_core": [4,1.9607843137254901,2.9702970297029703,1]
  }
}
```

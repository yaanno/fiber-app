Fiber and jQuery
======================

Inspired by [Jack Herrington's experiment](https://github.com/jherr/rscs-vs-php)

Ingredients:
- Go & Fiber (optional: Air)
- jQuery (for now)

# Running

Grab the `data` server from Jack's repo and run the server:
  
```bash
% cd data
% binserve
```

This will serve the pokemon data on port 8080. You can use [binserve](https://github.com/mufeedvh/binserve) or any other fast static server.

Then run the following in this repo (assuming you have Go and optionally Air installed):

## Air

```bash
air
```

## Without Air

```bash
go run app.go
```


Serves on: `http://locahost:3000/`

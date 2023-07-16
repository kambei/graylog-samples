# Graylog Samples

## Run Graylog

```bash
docker-compose up -d
```

Go to `http://localhost:9000` and login with `admin/admin`

## Setting Grayglog for receive UDP GELF packets:

- Go to `System > Inputs`
- Click on `Select Input` and choose `GELF UDP`
- Click on `Launch new input`
- Set `Title` to `LOGs GELF UDP`
- Click on `Save`

## Find LOGS:

- On input `LOGs GELF UDP` click on `Show received messages`

## Launch Samples:

- Sample Go with `http://localhost:8080/sample/test` endpoint that receive
```
{
    "message": "Hello Graylog from sample-go!",
}
```

- Sample Node.js with `http://localhost:3000/sample/test` endpoint that receive
```
{
    "message": "Hello Graylog from sample-node!",
}
```

- Sample Quarkus with `http://localhost:8080/sample/test` endpoint that receive a `String` message in body and logs it with `log.info()`.

- Sample Spring Boot with `http://localhost:8080/sample/test` endpoint that receive a `String` message in body and logs it with `log.info()`.

---

<br>
<br>

Enjoy it!
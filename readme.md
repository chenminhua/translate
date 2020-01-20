## use docker
docker run -d -p 8080:8080 chenmhgo/go-translate

## use translate.randomaccess.world (online demo)
```sh
# single sentence translate
curl "https://translate.randomaccess.world/trans?tl=zh&q=hello%20world"

# upload a file, translate to Chinese
curl https://translate.randomaccess.world/upload\?tl\=zh -F "file=@translate/en.txt"
```


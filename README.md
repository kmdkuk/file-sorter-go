# file-sorter-go
![CI](https://github.com/kmdkuk/file-sorter-go/workflows/CI/badge.svg)

テキストファイルの行をいい感じにソートしてくれるプログラムが書きたくなった．

## Example

### before

```shell script
$ cat example.txt
tmp/
log/
.git
containers
Dockerfile*
docker-compose*
vendor/
id_rsa.sweets.enc
.env*
.circleci
.travis.yml
docker_images
config/google_key.json
lib/csv_from_spread/
docs/
.bundle/
.circleci/
.github/
.idea/
.vscode/
```

### after

```shell script
$ file-soreter -i example.txt
.bundle/
.circleci
.circleci/
.env*
.git
.github/
.idea/
.travis.yml
.vscode/
Dockerfile*
config/google_key.json
containers
docker-compose*
docker_images
docs/
id_rsa.sweets.enc
lib/csv_from_spread/
log/
tmp/
vendor/
```

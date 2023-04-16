# gitcralwer

## Usage

* generate git commits json first
```sh
echo "[" > /tmp/test.data
git log develop..feature-pet-migration  --pretty='format:{"short_hash":"%h", "message": "%s", "author": "%aN"},' >> /tmp/test.data
file=$(cat /tmp/test.data)
file2=${file:0:${#file}-1}
echo $file2 > test.data
echo "]" >> test.data
```
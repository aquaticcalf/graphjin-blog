## graphjin blog

### what i did so far

- installed graphjin

```sh
go install github.com/dosco/graphjin@latest
```

- had to alias graphjin, because it aint working 😠

```sh
alias graphjin='$(go env GOPATH)/bin/graphjin'
```

- created a new "blog" app

```sh
graphjin new blog
```

- installed postgresql

```sh
sudo pacman -S postgresql
```

- installed atlas 

```sh
curl -sSf https://atlasgo.sh | sh
```
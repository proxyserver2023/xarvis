# go-bootstrap

```bash
export PROJECT_NAME=go-bootstrap
git clone https://github.com/alamin-mahamud/go-bootstrap.git $PROJECT_NAME
cd $PROJECT_NAME
go mod download
go mod tidy 
go mod vendor
go build -o $PROJECT_NAME
./$PROJECT_NAME
```
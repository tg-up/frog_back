### После установки дайте привелегии для bash-скрипта
```bash
   sudo u+x docgenerate.bash
```
### Чтобы сгенерировать документацию, запустите bash-скрипт
```bash
    ./docgenerate.bash
```
### Далее установите swaggo/swag
```bash
    go install github.com/swaggo/swag/cmd/swag@latest
```
### Если env_go не установлен в bashrc, до добавьте следующее -
```bash
    export PATH=$(go env GOPATH)/bin:$PATH
```
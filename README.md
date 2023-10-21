# Automation-Whitelist Security Group

## Architecture

![sg](./public/sg.png)

## Desc

- Security IngressRule은 최대 50~60개만 입력가능
- 일일이 넣는건 너무 비효율적...
- Security groups에서 화이트리스트나 IngressRule 자동화 Lambda
- 현재는 환경변수를 사용해서 진행, 추후 여러 AWS Service와 연동 할수 있음

## run

```sh

    ## init
    make run

    ## destroy
    make destroy
```

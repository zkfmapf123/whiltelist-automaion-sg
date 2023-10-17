# Automation-Whitelist Security Group

## Architecture

![sg](./public/sg.png)

## Problems

- 화이트리스트를 허용하는 ip를 SG에 추가하는 업무
- SG는 기본적으로 Ingress만 허용함
- 허용되는 Ip가 추가된다면 SG는 새로 계속 만들어져 야 함
- 불필요한 Ip Ingress는 삭제되어야 함
- 이 모든게 Lambda + Api Gateway를 사용하여 쉽게 관리가 되어야 함

## Todo

- Golang + aws sdk
- infra use Terraform

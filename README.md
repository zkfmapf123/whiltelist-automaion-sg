# Automation-whitelist Security Group

## Problems

- 화이트리스트를 허용하는 ip를 SG에 추가해야 함
- SG는 기본적으로 50개의 Ingress 만 허용함
- 허용되는 ip가 추가될수록 SG는 새로만들어져야 함
- 결국 자동화된 SG WhiteList를 만들자 use Lambda

## Todo

- [ ] Lambda + Terraform
- [ ] API Gateway를 활용하여 ip값을 받으면 바로 연동
- [ ] RetrieveSG
- [ ] MakeSG
- [ ] InjectSecurityRules
- [ ] test code

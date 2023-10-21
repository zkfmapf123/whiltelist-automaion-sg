clean:
	@rm -rf hello && rm -rf function.zip

binary:
	@GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o hello hello.go

zip:
	@zip function.zip hello

## 환경변수 설정
register-environment:
	@aws lambda update-function-configuration \
	 --function-name whitelist-function \
	 --environment "Variables={VPC_ID=vpc-06151804b151e2c54,SG_PREFIX=whitelist,PORT=27931,IPRange=10.0.0.1#10.0.0.2#10.0.0.3#10.0.0.4#10.0.0.5#10.0.0.6#10.0.0.7#10.0.0.8#10.0.0.9#10.0.0.10}" \

## 람다코드 업데이트
pre-update:
	@aws lambda update-function-code \
	 --function-name whitelist-function \
	 --zip-file fileb://function.zip \

invoke:
	@aws lambda invoke --function-name whitelist-function --cli-binary-format raw-in-base64-out --payload '{"key": "value"}' out

update:
	make clean && \
	make binary && \
	make zip && \
	make register-environment && \
	make pre-update
	make invoke

## 전체 Cycle 자동화
run:
	@cd infra && terraform init && terraform apply --auto-approve
	make update
	make invoke	

destroy:
	@cd infra && terraform destroy --auto-approve
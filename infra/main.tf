###############################################################
## Lambda IAM
###############################################################
resource "aws_iam_role" "whitelist-lambda-role" {
  name = "whitelist-lambda-role"
  assume_role_policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Principal" : {
            "Service" : "lambda.amazonaws.com"
          },
          "Action" : "sts:AssumeRole"
        }
      ]
    }
  )
}

resource "aws_iam_policy" "whitelist-lambda-policy" {
  name        = "whitelist-lambda-policy"
  path        = "/"
  description = "default labmda policy"

  policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Action" : [
            "logs:CreateLogGroup",
            "logs:CreateLogStream",
            "logs:PutLogEvents"
          ],
          "Resource" : "*"
        },
        {
          "Effect" : "Allow",
          "Action" : [
            "ec2:*"
          ],
          "Resource" : "*"
        }
      ]
  })
}

resource "aws_iam_policy_attachment" "attach" {
  name       = "whitelist-lambda-attachment"
  roles      = [aws_iam_role.whitelist-lambda-role.name]
  policy_arn = aws_iam_policy.whitelist-lambda-policy.arn
}

###############################################################
## Lambda
###############################################################

resource "aws_lambda_function" "test_lambda" {
  # If the file is not in the current working directory you will need to include a
  # path.module in the filename.
  filename      = "folder/function.zip"
  function_name = "whitelist-function"
  role          = aws_iam_role.whitelist-lambda-role.arn
  handler       = "hello"
  timeout       = 10 ## timeout  시간
  runtime       = "go1.x"

  environment {
    variables = {
      Name = "whitelist-lambda"
    }
  }

  lifecycle {
    ignore_changes = [filename, environment]
  }
}

output "lambda_info" {
  value = aws_lambda_function.test_lambda
}

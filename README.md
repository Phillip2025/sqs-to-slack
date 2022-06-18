# sqs-to-slack

AWS Application to send SQS messages to Slack

* [AWS Serverless Application Repository](https://aws.amazon.com/serverless/serverlessrepo/)

## How to use

Just add the next code to your AWS SAM or Cloud Formation template.yaml
```
  sqstoslack:
    Type: AWS::Serverless::Application
    Properties:
      Location:
        ApplicationId: arn:aws:serverlessrepo:eu-west-3:937158136137:applications/sqs-to-slack
        SemanticVersion: 0.0.2
        
      
      
```

### Parameters



codeBucket="drew-ecs-demo"
aws s3 cp infrastructure/vpc.yaml s3://${codeBucket}/infrastructure/vpc.yaml
aws s3 cp infrastructure/security-groups.yaml s3://${codeBucket}/infrastructure/security-groups.yaml
aws s3 cp infrastructure/load-balancers.yaml s3://${codeBucket}/infrastructure/load-balancers.yaml
aws s3 cp infrastructure/ecs-cluster.yaml s3://${codeBucket}/infrastructure/ecs-cluster.yaml
aws s3 cp infrastructure/lifecyclehook.yaml s3://${codeBucket}/infrastructure/lifecyclehook.yaml

aws s3 cp services/account-service/service.yaml s3://${codeBucket}/services/account-service/service.yaml
aws s3 cp services/product-service/service.yaml s3://${codeBucket}/services/product-service/service.yaml
aws s3 cp services/website-service/service.yaml s3://${codeBucket}/services/website-servi


aws cloudformation deploy --template-file master.yaml --stack-name ecs-refarch-drew --capabilities CAPABILITY_IAM CAPABILITY_NAMED_IAM CAPABILITY_AUTO_EXPAND  --parameter-overrides BucketName="${codeBucket}"

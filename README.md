# 通过流水线部署方式创建CDK Application演示
---

## 环境
---
CDK语言： GO 1.18.3 
CICD流水线： AWS Code Pipeline
###
参考文档：
---
AWS 官方演示代码
https://github.com/aws-samples/aws-cdk-deploy-pipeline-using-aws-codepipeline

AWS SA开源代码
https://github.com/cowcoa/aws-cdk-go-examples
### 试验
---
- [x] 单账号，单region，单stack部署
- [ ] 单账号，单region，手工批准部署
- [ ] 单账号，单region，多stack部署
- [ ] 单账号，单region，克隆分支部署及更新
- [ ] 单账号，单region，区分开发生产环境
- [ ] 单账号，多region，Region A开发环境，Region B生产环境
- [ ] 单账号，多region，Region A开发环境，Region B，C，D生产环境
- [ ] 多账号，单region，账号 开发环境，账号 B生产环境
- [ ] 多账号，多region，账号 A Region A开发环境，账号B Region B，C，D生产环境
- [ ] 单Repository，多CDK Application
- [ ] CDK测试代码
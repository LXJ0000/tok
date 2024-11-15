# tok
## 目录结构介绍
- backend 后端
  - gateway 网关层：api
  - basicService 基础服务：account + rbac(Role-based access control)
  - coreService 核心服务：comment + favorite + video + collection + relation
  - imService 即时通讯服务：openim
- frontend 前端

## 常见问题
```
1. kratos new 指定包名
eg. 
kratos new gateway
修改 go.mod github.com/LXJ0000/tok/backend/gateway
find . -type f -name '*.go' -exec sed -i 's|module gateway|module github.com/LXJ0000/tok/backend/gateway|g' {} +
find . -type f -name '*.go' -exec sed -i 's|"gateway/|"github.com/LXJ0000/tok/backend/gateway/|g' {} +

```
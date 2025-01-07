# Chaos Plus Backend



## Dlock 基于DB的分布式锁

 - 没有redis的时候 
 - 不需要高性能分布式锁场景, 比如 获取 workerid, 获取失败则 panic
 - 基于 created_by 可重入
 - 基于 expired_at 可过期, 过期则自动解锁

## WorkerID 分布式/集群下工作节点ID

 - 使用Dlock锁分配WorkerID
 - 租期内多次定时续期, 续期失败则 panic

## GUID 分布式全局唯一ID

> 基于 github.com/yitter/idgenerator-go 和 WorkerID


## Tenant 多租户, 数据隔离用


## Settings 配置

- 系统全局配置

## SMS 短信管理

- 三方平台对接
- 模版管理

## Email 邮件管理


## Auth 认证

- 用户
- 个人/企业实名认证
- Oauth2 三方登录
- SSO 单点登录
- 二次认证
- 验证码
- 

## Resource 资源

- 树结构, 组织架构

## Permission 权限

- 平面结构
- 接口权限
- 数据权限
- 菜单权限
- 按钮权限

## Role 角色

- 树结构, 组织架构

## Files 文件管理

- 基于闭包表设计的树结构文件管理
- 物理存储层去重
- 版本管理
- 回收站
- 公开分享和加密分享;
- 复制,粘贴,剪切,移动;
- 跳过,覆盖,自动重命名;
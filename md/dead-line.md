
```plantuml
@startgantt
printscale daily zoom 4
sunday are closed
[项目环境部署] lasts 2 days
[数据库设计] lasts 1 days
[接口设计] lasts 1 days

[数据库和接口评审] lasts 1 days
[管理后台UI开发] lasts 4 days

[管理后台前后端联调] lasts 1 days
[管理后台开发] lasts 3 days
[推理服务开发] lasts 5 days

[后台管理UI开发] lasts 6 days
[用户中心开发] lasts 3 days

[首页UI开发] lasts 2 days
[用户中心UI开发] lasts 2 days

[聊天开发] lasts 4 days


[后端联调] lasts 1 days
[前后端联调] lasts 2 days
[提测] lasts 1 days


Project starts 2023-05-22
[数据库设计] starts 2023-05-22
[接口设计] starts 2023-05-22
[项目环境部署] starts 2023-05-22
[数据库和接口评审] starts at [数据库设计]'s end
[数据库和接口评审] starts at [接口设计]'s end
[首页UI开发] starts 2023-05-24
[管理后台UI开发] starts 2023-05-22
[后台管理UI开发] starts 2023-05-22
[用户中心UI开发] starts 2023-05-26

[聊天开发] starts 2023-05-23
[管理后台开发] starts 2023-05-23
[用户中心开发] starts 2023-05-23
[AI服务开发] starts 2023-05-22

[管理后台前后端联调] starts at [管理后台开发]'s end
[管理后台前后端联调] starts at [管理后台UI开发]'s end

[后端联调] starts at [推理服务开发]'s end
[后端联调] starts at [聊天api开发]'s end
[前后端联调] starts at [用户中心UI开发]'s end
[前后端联调] starts at [用户中心开发]'s end
[前后端联调] starts at [后台管理UI开发]'s end
[前后端联调] starts at [后端联调]'s end
[提测] starts at [前后端联调]'s end
@endgantt
```
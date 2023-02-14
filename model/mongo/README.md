- 索引
  - 创建索引
    - 普通索引:`db.Test.createIndex({"created_time":-1})`{-1为倒序排列}
    - 唯一索引:`db.Test.createIndex({"created_time":-1},{unique:true})`
    - 后台创建:`db.Test.createIndex({"created_time":-1},{background:1})`
    - 指定名称:`db.Test.createIndex({"created_time":-1},{name:"创建时间索引"})`
    - dropDups:在创建唯一索引时是否删除重复记录,默认false
  - 查询索引: `db.Test.getIndexes()`
  - 查询分析:explain
    - `db.Test.find({created_time:"2023-02-14 12:00:00"}).explain()`
- 聚合查询():
  - `db.User.aggregate([{$match:{age:{$gte:18,$lte:30}}},{$group:{_id:"$nick_name",spent:{$sum:"$spent"},earn:{$sum:"$earn"}}}])`
  `{ "_id" : "张三", "spent" : NumberLong(0), "earn" : NumberLong(6) }
   { "_id" : "李四", "spent" : NumberLong(11), "earn" : NumberLong(29) }`
  - 解释:User表有nick_name(昵称),spent(花费),eran(赚钱)和age(年龄)字段,查询每个用户所有开销和赚取总额;
  - mysql语句: `select nick_name,SUM(spent) AS spent,SUM(eran) AS eran from User where age >= 18 AND age <= 30;`
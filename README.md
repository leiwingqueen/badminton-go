# badminton-go
羽球八人转

## 接口设计

### 1.获取比赛列表

/match/list

```json
[
  {
    "matchId":1,
    "name":"八人转"
    "startTime":xxx,//比赛开始时间
    "createTime":xxx,
    "updateTime":xxx,
    "status":1,//1-报名中,2-进行中,3-已结束
  },...
]
```

### 2.获得比赛详情

/match/detail

参数：

matchId-比赛ID

```json
{
  "matchId":1,
  "name":"八人转"
  "startTime":xxx,//比赛开始时间
  "createTime":xxx,
  "updateTime":xxx,
  "status":1,//1-报名中,2-进行中,3-已结束
  "rank"://排行
  [
  	{
  		"uid":1,
  		"openId":xxx,
  		"nickName":xxx,
  		"win":1 //胜利场次
  		"lose":10 //输的场次
  		"points":10 //净胜分
		},...
  ],
	"matchDetail":
	[
    {
      "roundId":1//场次
      //下面是4个对战选手
      "p1":
      {
        "uid":1,
  			"openId":xxx,
  			"nickName":xxx,
      },
      "p2":
      {
        "uid":1,
  			"openId":xxx,
  			"nickName":xxx,
      },
      "p3":
      {
        "uid":1,
  			"openId":xxx,
  			"nickName":xxx,
      },
      "p4":
      {
        "uid":1,
  			"openId":xxx,
  			"nickName":xxx,
      },
			//比分
			"result":"1:10"
    }
  ]
}
```

创建比赛

生成比赛

报名

提交比分

# 数据库设计

```sql
CREATE TABLE `tb_match`
(
    `matchId`       bigint(20) NOT NULL AUTO_INCREMENT COMMENT '比赛ID',
    `name`      		varchar(256) NOT NULL DEFAULT '' COMMENT '名称',
    `startAt`     timestamp NOT NULL DEFAULT '0' COMMENT '比赛开始',
    `createAt`    timestamp NOT NULL DEFAULT '0' COMMENT '开始时间',
    `updateAt`    timestamp NOT NULL DEFAULT '0' COMMENT '更新时间',
    `status`        tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
    PRIMARY KEY (`matchId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='比赛列表';

CREATE TABLE `tb_user`
(
    `uid`       		bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `openId`      	varchar(256) NOT NULL DEFAULT '' COMMENT 'openId',
    `name`      		varchar(256) NOT NULL DEFAULT '' COMMENT '昵称',
    `createTime`    timestamp NOT NULL DEFAULT '0' COMMENT '开始时间',
    `updateTime`    timestamp NOT NULL DEFAULT '0' COMMENT '更新时间',
    PRIMARY KEY (`uid`),
  	unique key uniq_openId(openId)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

CREATE TABLE `tb_match_user`
(
  	`matchId`       bigint(20) NOT NULL COMMENT '比赛ID',
    `uid`       		bigint(20) NOT NULL COMMENT '用户ID',
    key idx_matchId(matchId)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='比赛人员';


CREATE TABLE `tb_round`
(
  	`matchId`       bigint(20) NOT NULL COMMENT '比赛ID',
  	`roundId`       bigint(20) NOT NULL COMMENT '场次ID',
    `detail`        varchar(256) NOT NULL COMMENT '参赛者逗号分隔',
    `result`        varchar(128) NOT NULL COMMENT '比分',
    unique KEY uniq_round(matchId,roundId)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='比赛明细';
```


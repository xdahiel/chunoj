# Chun OJ API 管理文档

## 登录

* 请求接口: login

* 请求方法: post

* 请求参数

  | 参数名   | 含义     | 备注             |
  | -------- | -------- | ---------------- |
  | username | 用户名   | 必填项，长度6-20 |
  | password | 用户密码 | 必填项，长度6-20 |

* 返回参数

  

  | 参数名 | 含义   | 备注 |
  | ------ | ------ | ---- |
  | code   | 状态码 |      |
  | msg    | 信息   |      |
  | data   | 数据   |      |

* 响应数据

  ```json
  {
      code: "200",
      data: "",
      msg: "success"
  }
  ```

  

## 注册

注册界面

/problem

## 题目

题目

## 比赛

/contest

比赛

## 用户

/user

用户

## 后台

/admin

后台



>setup.sql
>
>```sql
>create table users (
>    id serial not null primary key ,
>    username varchar(20) not null ,
>    password varchar(255) not null,
>    isActive boolean not null ,
>    email varchar(255)
>);
>```


## Trie Server API

使用Trie树完成关键词维护和文本搜索。

### 接口说明

#### 1. 增加关键词

- 路由: `POST` /append
- Content-Type: application/json
- BODY
  
  ```json
  {
      "key": "关键词",
      "meta": JSON // 可选，任意的一个 JSON
  }
  ```
- 成功返回
  
  ```json
  {
     "code": 200,
     "msg": "OK"
  }
  ```

#### 2. 删除关键词

- 路由: `DELETE` /del/_关键词_
- 成功返回
  
  ```json
  {
     "code": 200,
     "msg": "OK"
  }
  ```

#### 3. 查询文本中出现的关键词

- 路由: `POST` /search
- Content-Type: application/json
- BODY
  
  ```json
  {
      "text": "任意文本"
  }
  ```
- 成功返回
  
  ```json
  {
     "code": 200,
     "msg": "OK",
     "result": {
        "关键词1": {
           "count": 1, // 出现次数
           "meta": JSON // append关键词时候给定的meta，如果没有给返回 null
        },
        "关键词2" {
           "count": 2,
           "meta": JSON
        }
     }
  }
  ```



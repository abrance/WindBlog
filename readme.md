### 这是什么?

    这是一个搭建个人网站的项目

### 有什么功能?

    发布文档, 展示文档内容


### 产品功能
##### 主体功能

    列出目录
    新增目录
    删除空目录
    tag 修改属性: 改名, 修改 nice 值
    
    展示目录中的文档
    展示 markdown 文档
    上传 markdown 文档
    删除文档
    
    为 markdown 文档新增 tag
    

##### 额外功能

    归档文档
    修改文档


### 概念

##### 文档 file

    元数据
    id:
    url: 
    create_time:
    update_time:
    is_archived:
    archive_id: 归档 id


##### tag

    id: 
    name:
    is_dir: 如果是目录, 会展示到首页中
    nice: 喜爱度

### 数据存储方式

##### 文档存储

    docs 中的某目录下, markdown 目前在 md 目录下

##### tag 信息存储

    保存于 data.json 中的 tag

##### tag-文档信息存储

    保存于 data.json 中的 tagDocBinding

### 程序部署目录结构

~~~
- WindBlog
  - bin(程序目录)
    - db (使用 json 的文件存储)
      - some db file
      - some db file ... 
    - config.ini (配置文件)
    - log (日志)
    - log.json (日志配置文件)
    - server (程序二进制)
    - doc (文档存储目录)
      - md (markdown 文档存储目录)
~~~
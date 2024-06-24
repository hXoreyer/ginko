app: # 应用基本配置
  env: local # 环境名称
  port: 8888 # 服务监听端口号
  app_name: gink-demo # 应用名称
  app_url: http://localhost # 应用域名
  lang: zh-cn

log:
  level: info # 日志等级
  root_dir: ./storage/logs # 日志根目录
  filename: app.log # 日志文件名称
  format: # 写入格式 可选json
  show_line: true # 是否显示调用行
  max_backups: 3 # 旧文件的最大个数
  max_size: 500 # 日志文件最大大小（MB）
  max_age: 28 # 旧文件的最大保留天数
  compress: true # 是否压缩

database:
  driver: mysql # 数据库驱动
  host: 127.0.0.1 # 域名
  port: 3306 # 端口号
  database: go-test # 数据库名称
  username: root # 用户名
  password: root # 密码
  charset: utf8mb4 # 编码格式
  max_idle_conns: 10 # 空闲连接池中连接的最大数量
  max_open_conns: 100 # 打开数据库连接的最大数量
  log_mode: info # 日志级别
  enable_file_log_writer: true # 是否启用日志文件
  log_filename: sql.log # 日志文件名称

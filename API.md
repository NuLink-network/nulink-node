# node api documentation

## 创建用户

创建用户记录

### 请求路径

/account/create

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数            | 类型     | 说明   |
| -------------- | -------- | ------- |
|  name          |  string  |  账户名称 |
|  account_id    |  string  |  账户ID(UUID v4) |
|  ethereum_addr |  string  |  以太坊地址 |
|  encrypted_pk  |  string  |  加密的公钥 |
|  verify_pk     |  string  |  验证公钥 |

### 响应参数

| 参数      | 类型      | 说明   |
| --------- | -------- | ----- |
|  code     |  int     |       |
|  msg      |  int     |       |
|  data     |  object  |       |

## 获取用户信息

通过账户 ID 获取用户信息

### 请求路径

/account/get

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数            | 类型     | 说明   |
| -------------- | -------- | ------- |
|  account_id    |  string  |  账户ID(UUID v4) |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  int     |  响应码  |
|  msg      |  int     |  响应信息 |
|  data     |  object  |  响应数据 |

#### data 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  name          |  string  |  账户名称 |
|  account_id    |  string  |  账户ID(UUID v4) |
|  ethereum_addr |  string  |  以太坊地址 |
|  encrypted_pk  |  string  |  加密的公钥 |
|  verify_pk     |  string  |  验证公钥 |
|  status        |  number  |  账户状态 |
|  created_at    |  number  |  账户创建时间戳 |

## 判断用户是否存在

判断用户是否存在

### 请求路径

/account/isexist

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数            | 类型     | 必填    | 说明   |
| -------------- | -------- | ------ | ------- |
|  name          |  string  | 是     | 账户名称 |
|  account_id    |  string  | 是     | 账户ID (UUID v4) |
|  ethereum_addr |  string  | 是     | 以太坊地址 |
|  encrypted_pk  |  string  | 是     | 加密的公钥 |
|  verify_pk     |  string  | 是     | 验证公钥 |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  int     |  响应码  |
|  msg      |  int     |  响应信息 |
|  data     |  object  |  响应数据 |

#### data 结构

| 参数      | 类型      | 说明     |       
| --------- | -------- | ------- |  
| is_exist  |  bool  |    账户是否存在   |      

## 上传文件

上传文件

### 请求路径

/file/upload

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填    | 说明   |
| ------------ | -------- | ------ | ------- |
|  files        |  [][File](#File-结构)  | 是     | 文件列表 |
|  account_id  |  string  | 是     | 账户 ID (UUID V4) |
|  policy_id   |  number  | 否     | 策略 ID |
|  signature   |  string  | 是     | 签名 |

#### File 结构

| 参数       | 类型     | 必填    | 说明   |
| --------- | -------- | ------ | ------- |
|  id     |  string  | 是     | 文件 ID (UUID V4) |
|  md5     |  string  | 是     | 文件 MD5 |
|  name     |  string  | 是     | 文件名称 |
|  suffix     |  string  | 否     | 文件后缀 |
|  category     |  string  | 是     | 文件类型 |
|  address  |  string  | 是     | 文件地址 |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  int     |  响应码  |
|  msg      |  int     |  响应信息 |
|  data     |  object  |  响应数据 |

## 创建策略并上传文件

创建策略并上传文件

### 请求路径

/file/create-policy-and-upload

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填    | 说明   |
| ------------ | -------- | ------ | ------- |
|  files        |  [][File](#File-结构)  | 是     | 文件列表 |
|  account_id  |  string  | 是     | 账户ID (UUID V4) |
|  policy_label_id   |  string  | 是     | 策略 label ID (UUID V4)|
|  policy_label   |  string  | 是     | 策略标签 |
|  encrypted_pk   |  string  | 是     | 加密的公钥 |
|  signature   |  string  | 是     | 签名 |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  number     |  响应码  |
|  msg      |  string     |  响应信息 |
|  data     |  object  |  响应数据 |

## 删除文件

删除文件

### 请求路径

/file/delete

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填    | 说明   |
| ------------ | -------- | ------ | ------- |
|  file_ids    |  []string  | 是     | 文件 ID 列表 |
|  account_id  |  string  | 是     | 账户 ID (UUID V4) |
|  signature   |  string  | 是     | 签名 |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  number     |  响应码  |
|  msg      |  string     |  响应信息 |
|  data     |  object  |  响应数据 |

## 文件列表

返回自己的文件信息列表

### 请求路径

/file/list

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填  | 默认值 | 说明   |
| ------------ | -------- | ------ | ---- |------- |
|  account_id  |  string  |  是    |      |  账户ID (UUID V4) |
|  file_name   |  string  |  否    |      |  文件名称, 支持模糊匹配|
|  paginate    |  [Paginate](#Paginate-结构) |  否  |      | 分页 |

#### Paginate 结构

| 参数       | 类型     | 必填  | 默认值 | 说明   |
| --------- | -------- | ---- | ---- | ------- |
|  page      |  number  | 否   |  1   |  页码 |
|  page_size |  number  | 否   |  10  |  每页的数据量， 最小值: 1，最大值: 100 |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  number     |  响应码  |
|  msg      |  string     |  响应信息 |
|  data     |  object  |  响应数据 |

#### data 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  list     |  array.object     |  文件列表数据  |
|  total     |  number     |  列表总数  |

#### list 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  file_id          |  string  |  文件 ID |
|  file_name          |  string  |  文件名称 |
|  owner          |  string  |  文件拥有者 |
|  owner_id          |  string  |  文件拥有者账户 ID |
|  address            |  string  |  文件地址 |
|  thumbnail          |  string  |  文件缩略图 |
|  created_at          |  number  |  文件上传时间戳 |

## 其他人的文件列表

返回符合条件的其他人上传的文件信息列表 (不包含自己上传的文件)

### 请求路径

/file/others-list

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填  | 默认值 | 说明   |
| ------------ | -------- | ------ | ---- |------- |
|  account_id  |  string  |  是    |      |  账户ID (UUID V4) |
|  file_name   |  string  |  否    |      |  文件名称, 支持模糊匹配|
|  category   |  string  |  否    |      |  文件类型|
|  format   |  string  |  否    |      |  文件格式|
|  desc   |  bool  |  否    |  false    |  是否按照上传时间倒序 |
|  paginate    |  [Paginate](#Paginate-结构) |  否  |      | 分页信息 |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  number     |  响应码  |
|  msg      |  string     |  响应信息 |
|  data     |  object  |  响应数据 |

#### data 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  list     |  array.object     |  列表数据  |
|  total     |  number     |  列表总数  |

#### list 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  file_id          |  string  |  文件 ID |
|  file_name          |  string  |  文件名称 |
|  owner          |  string  |  文件拥有者 |
|  owner_id          |  string  |  文件拥有者账户 ID |
|  address            |  string  |  文件地址 |
|  thumbnail          |  string  |  文件缩略图 |
|  created_at          |  number  |  文件上传时间戳 |

## 文件详情

返回文件的详细信息包括文件信息，申请信息，策略信息，文件拥有者 VerifyPK

1. 未申请使用：仅返回文件信息
2. 审核未通过：返回文件信息和申请信息
3. 文件申请通过但已过期：返回文件信息，申请信息，策略信息
4. 文件申请通过且未过期：返回文件信息，申请信息，策略信息，文件拥有者 VerifyPK 且返回下载相关信息(文件ipfs地址，策略加密公钥，策略藏宝图ipfs地址)

### 请求路径

/file/detail

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填    | 说明   |
| ------------ | -------- | ------ | ------- |
|  file_id    |  string  | 是     | 文件 ID (UUID V4)|
|  consumer_id  |  string  | 是     | 文件使用者的 ID (UUID V4)，即当前用户的 ID |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  number     |  响应码  |
|  msg      |  string     |  响应信息 |
|  data     |  object  |  响应数据 |

#### data 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  文件信息    |
|  file_id    |   string   | 文件 ID   |
|  file_name     |   string   |  文件名  |
|  thumbnail     |   string   |  文件缩略图  |
|  creator     |   string   |  文件的拥有者 (策略创建者)  |
|  creator_id     |   string   | 文件的拥有者 ID (策略创建者 ID)   |
|  file_created_at     |   number   | 文件上传时间戳   |
|  申请信息   |
|  apply_id     |   number   |  申请记录 ID  |
|  status     |   number   |  申请状态，0: 未申请，1: 申请中，2: 已通过, 3: 已拒绝  |
|  apply_start_at     |   string   |  申请开始时间戳(策略的开始时间戳)  |
|  apply_end_at     |   string   |  申请结束时间戳(策略的结束时间戳)  |
|  apply_created_at     |   string   |  提交申请时间戳  |
|  策略信息     |
|  policy_id     |   number   |  策略 ID  |
|  hrac     |   string   |  策略 hrac  |
|  consumer     |   string   |  策略的使用者(申请人，文件的使用者)  |
|  consumer_id     |   string   | 策略的使用者 ID   |
|  gas     |   string   |  策略 gas |
|  tx_hash     |   string   |  策略 tx hash |
|  policy_created_at     |   string   |  策略的创建时间戳  |
|  下载信息     |
|  file_ipfs_address     |   string   |  文件 ipfs 地址 |
|  policy_encrypted_pk     |   string   | 策略加密的公钥   |
|  encrypted_treasure_map_ipfs_address     |   string   |  策略的藏宝图地址  |
|  alice_verify_pk     |   string   |  文件拥有者的 Verify 公钥  |

## 撤销策略

撤销策略并删除文件和策略的关联关系及策略对应的所有文件的使用申请

### 请求路径

/policy/revoke

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填  | 默认值 | 说明   |
| ------------ | -------- | ------ | ---- |------- |
|  account_id  |  string  | 是     | 账户ID (UUID V4) |
|  policy_id   |  number  | 是     | 策略 ID |
|  signature   |  string  | 是     | 签名 |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  number     |  响应码  |
|  msg      |  string     |  响应信息 |
|  data     |  object  |  响应数据 |

## 获取策略关联的文件信息列表

获取策略关联的文件信息

### 请求路径

/policy/file-detail-list

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填  | 默认值 | 说明   |
| ------------ | -------- | ------ | ---- |------- |
|  creator_id  |  string  | 否     | 策略的创建者账户ID (UUID V4) |
|  consumer_id  |  string  | 否     | 策略的使用者账户ID (UUID V4) |
|  policy_id  |  number  | 是     | 策略ID  |
|  paginate    |  [Paginate](#Paginate-结构) |  否  |      | 分页 |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  number     |  响应码  |
|  msg      |  string     |  响应信息 |
|  data     |  object  |  响应数据 |

#### data 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  list     |  array.object     |  列表数据  |
|  total     |  number     |  列表总数  |

#### list 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
| file_id | string | 文件 ID | 
| file_name | string | 文件名称 | 
| owner | string | 文件拥有者 | 
| owner_id | string | 文件拥有者账户 ID |
| address | string | 文件地址 | 
| thumbnail | string | 文件缩略图 |
| created_at | number | 文件上传时间戳 |
| policy_id  |  number | 策略ID  |
| policy_hrac          |  string  |  策略 hrac |
| policy_start_at   |  number  |  策略开始时间戳 |
| policy_end_at     |  number  |  策略结束时间戳 |

## 策略信息列表

获取策略信息列表

### 请求路径

/policy/list

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填  | 默认值 | 说明   |
| ------------ | -------- | ------ | ---- |------- |
|  policy_id  |  number  | 否     | 策略ID  |
|  policy_label_id  |  number  | 否     | 策略ID  |
|  creator_id  |  string  | 否     | 策略的创建者账户ID (UUID V4) |
|  consumer_id  |  string  | 否     | 策略的使用者账户ID (UUID V4) |
|  paginate    |  [Paginate](#Paginate-结构) |  否  |      | 分页 |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  number     |  响应码  |
|  msg      |  string     |  响应信息 |
|  data     |  array.object  |  响应数据 |

#### data 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  list     |  array.object     |  列表数据  |
|  total     |  number     |  列表总数  |

#### list 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  hrac        |  string  |  hrac |
|  policy_id        |  number  | 策略 ID  |
|  creator        |  string  |  策略创建者 |
|  creator_id       |  string  |  策略创建者 ID |
|  consumer      |  string  | 策略使用者  |
|  consumer_id       |  string  | 策略使用者 ID  |
|  gas     |  string  |  gas |
|  tx_hash     |  string  | 交易 hash  |
|  start_at   |  number  |  策略开始时间戳 |
|  end_at     |  number  |  策略结束时间戳 |
|  created_at      |  number  | 策略创建时间戳  |

## 申请文件使用

申请文件使用

### 请求路径

/apply/file

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填  | 默认值 | 说明   |
| ------------ | -------- | ------ | ---- |------- |
|  file_ids  |  []string  |   是   |   |  文件 ID 列表 |
|  proposer_id  |  string  |   是   |   |  申请人的账户 ID |
|  start_at  |  number  |   是   |   |  开始时间戳 |
|  end_at  |  number  |   是   |   |  结束时间戳 |
|  signature   |  string  | 是     | 签名 |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  number     |  响应码  |
|  msg      |  string     |  响应信息 |
|  data     |  object  |  响应数据 |

## 申请文件使用列表

申请文件使用信息列表
1. 文件申请不存在所有信息为空。
2. 申请未通过仅返回申请信息策略信息为空

### 请求路径

/apply/list

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填  | 默认值 | 说明   |
| ------------ | -------- | ------ | ---- |------- |
|  apply_id |  number  |  否   |   |申请记录 ID，如果传递申请记录 ID 其他参数可以不传 |
|  file_id  |  string  |   否   |   |  文件 ID |
|  proposer_id  |  string  |   否   |   |  申请人的账户 ID (如果不指定申请记录 ID 申请人的账户 ID 和文件拥有者的账户 ID 二选一) |
|  file_owner_id  |  string  |   否   |   |  文件拥有者的账户 ID (如果不指定申请记录 ID 申请人的账户 ID 和文件拥有者的账户 ID 二选一) |
|  status  |  number  |   否   |  0 (不区分状态) |  申请状态，1: 申请中，2: 已通过, 3: 已拒绝|
|  paginate    |  [Paginate](#Paginate-结构) |  否  |      | 分页 |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  number     |  响应码  |
|  msg      |  string     |  响应信息 |
|  data     |  array.object  |  响应数据 |

#### data 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  list     |  array.object     |  列表数据  |
|  total     |  number     |  列表总数  |

#### list 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  申请信息          |
|  file_id          |  string  |  文件 ID |
|  apply_id          |  number  |  申请记录 ID |
|  proposer            |  string  |  申请人 |
|  proposer_id          |  string  |  申请人账户 ID |
|  file_owner          |  string  |  文件拥有者 |
|  file_owner_id          |  string  |  文件拥有者账户 ID |
|  status          |  number  |  申请状态，1: 申请中，2: 已通过, 3: 已拒绝 |
|  start_at          |  number  |  使用开始时间戳 |
|  end_at          |  number  |  使用结束时间戳 |
|  created_at          |  number  |  申请时间戳 |
|  策略信息  |
|  policy_id        |  number  | 策略 ID  |
|  policy_label_id        |  string  | 策略 label ID  |
|  hrac        |  string  | 策略 hrac |

## 撤销文件使用申请

撤销文件使用申请

### 请求路径

/apply/revoke

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填  | 默认值 | 说明   |
| ------------ | -------- | ------ | ---- |------- |
|  proposer_id  |  string  |   是   |   |  申请人的账户 ID |
|  apply_ids  |  []number  |   是   |   |  申请记录 ID 列表 |
|  signature   |  string  | 是     | 签名 |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  number     |  响应码  |
|  msg      |  string     |  响应信息 |
|  data     |  object  |  响应数据 |

## 批准文件使用申请

发布文件对应的策略并批准文件使用申请

### 请求路径

/apply/approve

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填  | 默认值 | 说明   |
| ------------ | -------- | ------ | ---- |------- |
|  account_id  |  string  |   是   |   |  审批人的账户 ID |
|  apply_id  |  number  |   是   |   |  申请记录 ID |
|  policy  |  Policy  |   是   |   |  申请记录 ID |
|  signature   |  string  | 是     | 签名 |

#### Policy 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  hrac          |  string  |  hrac |
|  gas          |  string  |  gas |
|  tx_hash          |  string  |  交易 Hash |
|  encrypted_pk          |  string  | encrypted public key  |
|  encrypted_address          |  string  | encrypted ipfs address  |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  number     |  响应码  |
|  msg      |  string     |  响应信息 |
|  data     |  object  |  响应数据 |

## 拒绝文件使用申请

拒绝文件使用申请

### 请求路径

/apply/reject

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填  | 默认值 | 说明   |
| ------------ | -------- | ------ | ---- |------- |
|  account_id  |  string  |   是   |   |  审批人的账户 ID |
|  apply_id  |  number  |   是   |   |  申请记录 ID |
|  signature   |  string  | 是     | 签名 |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  number     |  响应码  |
|  msg      |  string     |  响应信息 |
|  data     |  object  |  响应数据 |

## 策略 label 列表

策略 label 信息列表

### 请求路径

/label/list

### 请求方法

POST

### 数据类型

application/json

### 请求参数

| 参数          |  类型     | 必填  | 默认值 | 说明   |
| ------------ | -------- | ------ | ---- |------- |
|  account_id  |  string  |   否   |   |  label 创建者的账户 ID |
|  paginate    |  [Paginate](#Paginate-结构) |  否  |      | 分页 |

### 响应参数

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  code     |  number     |  响应码  |
|  msg      |  string     |  响应信息 |
|  data     |  array.object  |  响应数据 |

#### data 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  list     |  array.object     |  列表数据  |
|  total     |  number     |  列表总数  |

#### list 结构

| 参数      | 类型      | 说明     |
| --------- | -------- | ------- |
|  label          |  string  |  label |
|  label_id          |  number  |  label ID |
|  creator          |  string  |  创建者 |
|  creator_id          |  string  |  创建者 ID |
|  created_at          |  number  |  创建时间戳 |

## 常见响应码

| Code      | MSG      |
| --------- | -------- |
|  2000     |  Success        |
|  4000     |  Invalid Parameter       |
|  5000     |  Internal Server Error       |


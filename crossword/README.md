# Crossword

## Usage:
```shell
crossword config.json
```

## Configure file, json type:
```json
{
  "source": "data.txt",
  "output": "output.sql",
  "column": 4,
  "format": [
    "update users set name='#2#', phone=#3# where uid='#1#';",
	"",
	"insert into orders (bookid, uid) values (#4#, #1#)",
	""
  ]
}
```
* **source**: 需要处理的数据文件，表格形式文本，以制表符（\t)分隔。
* **output**: 生成文件。
* **column**: 数据字段数。
* **format**: 每条数据对应的样式模板。字符串组形式，多行输出，允许空串，表示空白行。

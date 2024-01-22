goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/cloud_disk" -table="user"  -dir="./apps/user/model" --style=goZero

goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/cloud_disk" -table="file"  -dir="./apps/fs/model" --style=goZero

goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/cloud_disk" -table="user_file"  -dir="./apps/fs/model" --style=goZero

goctl model mysql datasource -url="root:123456@tcp(127.0.0.1:3306)/cloud_disk" -table="user_directory"  -dir="./apps/fs/model" --style=goZero
Đây là project book management, có database (mySQL)
sử dụng gorilla/mux và gorm v1(gorm/jinzhu)
Viết routes trước để hình dung ra được route, vì user sẽ sử dụng front-end hay postman để tác động tới code
Tệp app.go sẽ giúp kết nối tới cơ sở dữ liệu mySql
Tiếp theo là tệp utils
Và cuối cùng là controller để chỉ tới model books
user interacts with the routes and the routes send control to the controllers where we have all our logic 
and controllers then have to perform some operations with database


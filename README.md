# TransServer
TransServer is golang software that use http way to insert or update data based on condition for mysql-alike database.

For environment:
  Server(master)------Client(Nodes)
  
  Nodes have about thousand, Depend on master server performance, The http package design for concurrency handle request.
  
Use case:
  The Client can use curl do the http upload for little data quickly, Immediate get insert okay or fail result.
  
This program use mysql drive: "github.com/go-sql-driver/mysql"
  
Please let me known if you have any idea or question. 

Email: xiaojiemi@gmail.com

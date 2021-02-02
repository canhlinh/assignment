### Đề bài: Xây dựng ứng dụng web


Viết một ứng dụng web server cơ bản, cho phép người dùng có thể hiển thị nội
dung một trang web. Trang Web bao gồm: Title và body.
Ứng dụng phải đáp ứng các yêu cầu sau:
1. Thiết kế các Chức năng:
- Xem nội dung web (/view): trả về một trang web bao gồm Title và Body
- Thay đổi title của trang web (/edit): truyền lên title mới, ứng dụng lưu title
mới và trả về một trang web bao gồm title mới và body
- Thay đổi body của trang web (/save): truyền lên body mới, ứng dụng lưu
body mới trả về một trang web bao gồm title và body mới
2. Cấu hình ứng dụng:
- chạy trên cổng 8090
- Có thể sử dụng bất kỳ thư viện/Gói phần mềm nào đã được học
- Xử lý các url tương ứng với các chức năng ở trên.   
http://localhost:8090/view   
http://localhost:8090/edit   
http://localhost:8090/save   
3. Yêu cầu phần mềm:
- Phải có cấu trúc dữ liệu cho trang web
- Ttle, body, Footer được lưu dưới dạng file, ứng dụng sẽ đọc file tương
ứng và trả về trang web với nội dung lấy từ cấu trúc dữ liệu đã thiết kế
- Phải có xử lý lỗi, bao gồm: lỗi ứng dụng khi xử lý, lỗi truy cập url không tồn
tại, lỗi truyền dữ liệu sai formats



### Cài đặt project:

Tải framework:
```
make setup
```

Chạy web server
```
make run
```
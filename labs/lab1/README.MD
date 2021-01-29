Cài đặt một hàm giải quyết trò chơi `Wheel of fortune`.
Hàm có tên là wheel_of_fortune sẽ nhận vào hai tham số bắt buộc và một tham
số tuỳ chọn, trong đó: Tham số thứ nhất là một mảng danh sách các từ có thể (từ điển). Ví dụ:
[green, red, blue, ball, is, are, house, the].   

Tham số thứ hai là một mảng gồm các từ tương ứng trong chuỗi ký tự đầu
ra của hàm nhưng là các từ bị ẩn đi một số chữ cái, trong đó mỗi chữ cái
chưa biết sẽ được đại diện bằng một dấu cách (ở đây sẽ dùng ký tự ␣ để
biểu diễn cho dấu cách space). Ví dụ: [t␣␣, ␣␣␣l, i␣, bl␣␣]   

Tham số thứ ba là tham số tuỳ chọn, là chuỗi ký tự gợi ý cho kết quả đầu
ra, trong đó một số chữ cái ở vị trí tương ứng trong kết quả đầu ra được
tiết lộ để gợi ý. Ví dụ: th␣␣␣a␣l␣i␣␣bl␣e. Trường hợp có 2 từ
được tìm thấy thì chọn từ đúng với cụm từ được gợi ý trong chuỗi ký tự
tham số thứ 3. Ví dụ: bell và ball thì từ ball sẽ được chọn do ở tham số thứ
3 có chứa ký tự a.
Kết quả trả về của hàm là một chuỗi gồm các từ trong từ điển tương ứng với các
từ gợi ý trong tham số đầu vào thứ hai và thứ ba (nếu có). Ví dụ: the ball is
blue   

Để test chương trình, trong hàm main() sẽ có lời gọi tới hàm wheel_of_fortune,
trong đó giá trị của các tham số được hardcoded trong hàm main.
Bạn có nhớ trình tạo hình ảnh bạn đã viết trước đó?
Hãy viết một cái khác, nhưng lần này nó sẽ trả về việc cài đặt của image.Image
thay vì một slide dữ liệu.

Define your own Image type, implement the necessary methods, and call
pic.ShowImage.   
Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
ColorModel should return color.RGBAModel.    
At should return a color; the value v in the last picture generator corresponds to
color.RGBA{v, v, 255, 255} in this one.

Định nghĩa kiểu Image của riêng bạn, cài đặt những phương thức cần thiết, và
gọi pic.ShowImage
“Bounds” nên trả về một image.Rectangle, như image.Rect(0, 0, w, h).
ColorModel nên trả về color.RGBAModel   
“At nên trả về một màu, giá trị v trong trình tạo hình cuối cùng tương ứng với
color.RGBA(v, v, 255, 255)

```
package main
import golang.org/x/tour/pic
type Image struct{}

func main() {
    m := Image{}
    pic.ShowImage(m)
}
```

Tham khảo:

```
package main
import golang.org/x/tour/pic
func Pic(dx, dy int) [][]uint8 {
}

func main() {
    pic.Show(Pic)
}
```
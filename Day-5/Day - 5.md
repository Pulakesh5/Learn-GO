# Day - 5

- **Created On:** May 25, 2024 10:47 AM
- **Last Updated:** May 25, 2024 1:05 PM

---

Today is Saturday, so it is revision time.

Here is the solution for `rot13Reader` exercise. This question asks to implement a `Read()` method for `rot13Reader` and read one byte at a time and substitute it with its rotate13 cipher code.

- ```go
   package main

   import (
   	"bufio"
   	"fmt"
   	"io"
   	"os"
   	"strings"
   )

   type rot13Reader struct {
   	r io.Reader
   }

   func (rot13 *rot13Reader) Read(p []byte) (n int, err error) {
   	r := bufio.NewReader(rot13.r)
   	for index := range p {
   		rune, _, err := r.ReadRune()
   		if err != nil {
   			return 0, err
   		}
   		if (rune >= 'A' && rune <= 'M') || (rune >= 'a' && rune <= 'm') {
   			p[index] = byte(rune + 13)
   			fmt.Printf("%c", p[index])
   		} else if (rune >= 'N' && rune <= 'Z') || (rune >= 'n' && rune <= 'z') {
   			p[index] = byte(rune - 13)
   			fmt.Printf("%c", p[index])
   		} else {
   			p[index] = byte(rune)
   			fmt.Printf("%c", p[index])
   		}
   	}
   	return len(p), nil
   }
   func main() {
   	s := strings.NewReader("Lbh penpxrq gur pbqr!")

   	r := rot13Reader{s}
   	io.Copy(os.Stdout, &r)
   }
   /*
   You cracked the code!
   */
  ```

---

This was a fun exercise to work with the `image` package.

Here is the custom implementation of `Image` struct :

- ```go
  package main

  import (
  	"image"
  	"image/color"
  	"golang.org/x/tour/pic"
  )

  type Image struct {
  	w, h  int
  	Image *image.RGBA
  }

  func (i Image) At(x, y int) color.Color {
  	return i.Image.At(x, y)
  }
  func (i Image) Bounds() image.Rectangle {
  	return image.Rect(0, 0, i.w, i.h)
  }
  func (i Image) ColorModel() color.Model {
  	return i.Image.ColorModel()
  }
  func main() {
  	m := Image{w: 100, h: 100}
  	m.Image = image.NewRGBA(image.Rect(0, 0, m.w, m.h))
  	for x := 0; x < m.w; x++ {
  		for y := 0; y < m.h; y++ {
  			m.Image.Set(x, y, color.RGBA{uint8(x), uint8(y), uint8(x * y), 255})
  		}
  	}
  	pic.ShowImage(m)
  }

  ```

<!-- 
![A beautiful RGB image created in GO](Day%20-%205%20de2672e818d841d487097b4f00d47de2/Untitled.png)
-->
<div style="text-align: center;">
<img src="Day%20-%205%20de2672e818d841d487097b4f00d47de2/Untitled.png" alt="A beautiful RGB image created in GO" width="600" height="450">
A beautiful RGB image created in GO
</div>


---

// Just as a proof of concept for now.
// Open a png file, change the brightness of all pixels by a factor, and save it again.


package main

import (
    "fmt"
    "image"
    "image/png"
    "image/color"
    "os"
    "time"
)

func min(a, b float64) float64 {
    // as if I’m going to include the entirety of math just for this, lel
    if a < b {
        return a
    } else {
        return b
    }
}

func saturate(p uint32, f float64) uint8 {
    p2 := float64(p) * f
    p2 = min(p2, 255)
    return uint8(p2)
}

func main() {
    file, _ := os.Open("test.png")
    defer file.Close()
    img, _ := png.Decode(file)
    b := img.Bounds()
    imgSet := image.NewRGBA(b)
    fmt.Println(img.At(0, 0))
    start := time.Now()
    for x:=0; x<b.Max.X; x++ {
        for y:=0; y<b.Max.Y; y++ {
            old := img.At(x, y)
            r, g, b, a := old.RGBA()
            r, g, b = r/256, g/256, b/256
            fac := 0.8
            //r2, g2, b2 := min(255, float64(r)*1.2), min(255, float64(g)*1.2), min(255, float64(b)*1.2)
            r2, g2, b2 := saturate(r, fac), saturate(g, fac), saturate(b, fac)
            /*if x*y%50000==0 {
                fmt.Printf("old: %d, %d, %d. new: %d, %d, %d\n", r, g, b, r2, g2, b2)
            }*/
            imgSet.Set(x, y, color.RGBA{r2, g2, b2, uint8(a)})
        }
    }
    fmt.Println(time.Now().Sub(start))
    outFile, _ := os.Create("test2.png")
    defer outFile.Close()
    png.Encode(outFile, imgSet)
}


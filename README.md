# golang_scraping

Go programlama dilinde seçtiğimiz bir web sitesinden veri kazıma işlemi yaptık
ve bu işlem sonucunda çektiğimiz verileri bir json dosyasına kaydettik.
Go ile windows form uygulamasına değinmek adına çektiğimiz verileri
listelemek için listbox widget kullandık.

Go dilinde windows form işlemleri için aşağıdaki linki inceleyebilirsiniz.
```
https://github.com/polatyener-dev/windows-app
```
Web sitesinden veri alma işlemleri için de aşağıdaki kod yardımı ile kütüphanemizi projemize dahil ediyoruz.
```
$ go get github.com/PuerkitoBio/goquery
```
Ve son olarak projemizi çalıştırdığımızda cmd ekranının gelmemesi için derleme işleminde aşağıdaki kodu kullanabiliriz.
```
go build -ldflags="-H windowsgui"
```

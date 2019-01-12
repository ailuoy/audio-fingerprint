# Audio-Fingerprint

This project is to analyze the similarity between two audios.

## Building the Application
>Depends on [fpcalc](https://github.com/acoustid/chromaprint) linux extension

### Install fpcalc
```
sudo wget https://github.com/acoustid/chromaprint/releases/download/v1.4.3/chromaprint-fpcalc-1.4.3-linux-x86_64.tar.gz
sudo tar -zxvf chromaprint-fpcalc-1.4.3-linux-x86_64.tar.gz
sudo mv /usr/local/chromaprint-fpcalc-1.4.3-linux-x86_64/fpcalc /usr/local/bin
```

```
fpcalc -version
fpcalc version 1.4.3
```

## Getting Started

```
first, _ := fingerprint.GetFingerPrintInfo("first.mp3")
second, _ := fingerprint.GetFingerPrintInfo("second.mp3")
similarity, _, _ := fingerprint.Correlate(first.Fingerprint, second.Fingerprint)
fmt.Println(similarity)
```

## Benchmark
```php
go test -test.bench="BenchmarkSongFpCompare" -benchmem -benchtime=30s -cpuprofile profile_cpu.out

go tool pprof -svg profile_cpu.out > profile_cpu.svg
```

## References
### [go-fingerprint/gochroma](https://github.com/go-fingerprint/gochroma)
>Gochroma can only match exactly

### [chromaprint](https://github.com/acoustid/chromaprint)

## Thanks
- [nay-kang](https://github.com/nay-kang) Provide python script


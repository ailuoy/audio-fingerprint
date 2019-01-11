# Audio-Fingerprint

This project is to analyze the similarity between two audios.

## Getting Started

Start by not writing any code.

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


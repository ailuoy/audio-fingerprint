package fingerprint

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const (
	sampleTime = 500
	span       = 150
	step       = 1
	minOverlap = 20
	threshold  = 0.5
)

type Fingerprint struct {
	Duration    int
	Fingerprint []int32
}

func GetFingerPrintInfo(fileName string) (fingerprint Fingerprint, err error) {
	_, err = os.Stat(fileName)
	if err != nil {
		err = errors.New(fmt.Sprintf("File does not exist %s", fileName))
		return
	}
	content, fileErr := readFileFPRaw(fileName)
	if err != nil {
		err = fileErr
		return
	}
	duration, fileErr := getDuration(content)
	if err != nil {
		err = fileErr
		return
	}
	fp, _ := getFingerprint(content)
	fingerprint = Fingerprint{
		Duration:    duration,
		Fingerprint: fp,
	}

	return
}

func readFileFPRaw(fileName string) (content []byte, err error) {
	cmd := exec.Command("/bin/bash", "-c", `fpcalc -raw `+fileName)
	stdout, fileErr := cmd.StdoutPipe()
	if fileErr != nil {
		err = errors.New(fmt.Sprintf("Error:can not obtain stdout pipe for command:%s\n", fileErr))
		return
	}

	if fileErr := cmd.Start(); fileErr != nil {
		err = errors.New(fmt.Sprintf("Error:The command is err,%s", fileErr.Error()))
		return
	}

	bytes, FileErr := ioutil.ReadAll(stdout)
	if err != nil {
		err = errors.New(fmt.Sprintf("ReadAll Stdout:%s", FileErr.Error()))
		return
	}

	if FileErr := cmd.Wait(); FileErr != nil {
		err = errors.New(fmt.Sprintf("wait:%s", FileErr.Error()))
		return
	}
	str := fmt.Sprintf("%s", bytes)
	content = []byte(str)

	return
}

func getDuration(contents []byte) (duration int, err error) {
	strDurationRe := `DURATION=(\d+)\nFINGERPRINT`
	reg := regexp.MustCompile(strDurationRe)
	match := reg.FindSubmatch(contents)
	duration, err = strconv.Atoi(string(match[1]))
	return
}

func getFingerprint(contents []byte) (fingerprint []int32, err error) {
	strDurationRe := `FINGERPRINT=(.*)`
	reg := regexp.MustCompile(strDurationRe)
	match := reg.FindSubmatch(contents)
	splitStr := strings.Split(string(match[1]), ",")
	for _, v := range splitStr {
		temp, _ := strconv.Atoi(v)
		fingerprint = append(fingerprint, int32(temp))
	}
	return
}

func Correlate(source []int32, target []int32) (maxCorrIndex float32, maxCorrOffset int, err error) {
	corr, err := compare(source, target, span, step)
	if err != nil {
		return
	}
	maxCorrIndex, maxCorrOffset = getMaxCorr(corr)
	return
}

func getMaxCorr(corr []float32) (float32, int) {
	maxCorrIndex := maxIndex(corr)
	maxCorrOffset := -span + maxCorrIndex*step
	return corr[maxCorrIndex], maxCorrOffset
}

func maxIndex(corr []float32) (maxIndex int) {
	maxValue := corr[0]
	for i, value := range corr {
		if value > maxValue {
			maxValue = value
			maxIndex = i
		}
	}
	return
}

func compare(source []int32, target []int32, span int, step int) (corrXy []float32, err error) {
	sourceLen := len(source)
	targetLen := len(target)
	if span > Min(sourceLen, targetLen) {
		err = errors.New("span >= sample size: %i >= %i")
		return
	}

	var box []int
	for i := -span; i < span+1; i += step {
		box = append(box, i)
	}

	for _, offset := range box {
		rate, _ := crossCorrelation(source, target, offset)
		corrXy = append(corrXy, rate)
	}
	return

}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func crossCorrelation(source []int32, target []int32, offset int) (number float32, err error) {
	sourceLen := len(source)
	targetLen := len(target)
	if offset > 0 {
		source = source[offset:]
		if sourceLen > targetLen {
			target = target[:targetLen]
		} else {
			target = target[:sourceLen]
		}
	} else if offset < 0 {
		offset = -offset
		if offset > targetLen {
			target = []int32{}
		} else {
			target = target[offset:]
		}
		if targetLen < sourceLen {
			source = source[:targetLen]
		}
	}

	if Min(len(source), len(target)) < minOverlap {
		return
	}
	number, err = correlation(source, target)
	return
}

func correlation(source []int32, target []int32) (rate float32, err error) {
	sourceLen := len(source)
	targetLen := len(target)
	if sourceLen == 0 || targetLen == 0 {
		err = errors.New("souce pr target len 0")
		return
	}
	if sourceLen > targetLen {
		source = source[:targetLen]
	} else if sourceLen < targetLen {
		target = target[:sourceLen]
	}
	var covariance float32
	covariance = 0
	for i := 0; i < len(source); i++ {
		binStr := bin(source[i] ^ target[i])
		covariance += (32 - float32(strings.Count(binStr, "1")))
	}
	covariance = covariance / float32(sourceLen)
	rate = covariance / 32
	return
}

func bin(number int32) (binStr string) {
	return "0b" + strconv.FormatInt(int64(number), 2)
}

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

// начало решения

// asLegacyDate преобразует время в легаси-дату
func asLegacyDate(t time.Time) string {
	nano := t.UnixNano()
	sec, rest := nano/1e9, nano%1e9
	res := fmt.Sprintf("%d.%09d", sec, rest)
	res = strings.TrimRight(res, "0")
	if res[len(res)-1] == '.' {
		res += "0"
	}
	return res
}

// parseLegacyDate преобразует легаси-дату во время.
// Возвращает ошибку, если легаси-дата некорректная.
func parseLegacyDate(d string) (time.Time, error) {
	parts := strings.Split(d, ".")
	if len(parts) != 2 {
		return time.Time{}, errors.New("wrong format")
	}
	var sec, rest int64
	var err error
	if sec, err = strconv.ParseInt(parts[0], 10, 64); err != nil {
		return time.Time{}, err
	}
	if 0 < len(parts[1]) && len(parts[1]) < 9 {
		parts[1] += strings.Repeat("0", 9-len(parts[1]))
	}
	if rest, err = strconv.ParseInt(parts[1], 10, 64); err != nil {
		return time.Time{}, err
	}
	return time.Unix(sec, rest), nil
}

// конец решения

func Test_asLegacyDate(t *testing.T) {
	samples := map[time.Time]string{
		time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC): "3600.123456789",
		time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC):         "3600.0",
		time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC):         "0.0",
	}
	for src, want := range samples {
		got := asLegacyDate(src)
		if got != want {
			t.Fatalf("%v: got %v, want %v", src, got, want)
		}
	}
}

func Test_parseLegacyDate(t *testing.T) {
	samples := map[string]time.Time{
		"3600.123456789": time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC),
		"3600.0":         time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC),
		"0.0":            time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		"1.123456789":    time.Date(1970, 1, 1, 0, 0, 1, 123456789, time.UTC),
		//"1653403522.951000": time.Date(2022, 5, 24, 14, 45, 22, 700, time.UTC), // 2022-05-24 14:45:22.951
	}
	for src, want := range samples {
		got, err := parseLegacyDate(src)
		if err != nil {
			t.Fatalf("%v: unexpected error", src)
		}
		if !got.Equal(want) {
			t.Fatalf("%v: got %v, want %v", src, got, want)
		}
	}
}

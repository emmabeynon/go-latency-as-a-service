package main

import (
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Latency as a service", func() {
	Describe("GET /latency", func() {
		Context("with no duration param", func() {
			It("returns 200 status code", func() {
				writer := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", "/latency", nil)
				latencyServer(writer, request)
				Expect(writer.Code).To(Equal(200))
			})

			It("returns 'OK' after 500 ms", func() {
				writer := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", "/latency", nil)
				timeBeforeRequest := time.Now().UnixNano() / int64(time.Millisecond)
				latencyServer(writer, request)
				timeAfterRequest := time.Now().UnixNano() / int64(time.Millisecond)
				latency := timeAfterRequest - timeBeforeRequest
				Expect(writer.Body.Bytes()).To(ContainSubstring("OK"))
				Expect(latency).To(BeNumerically("~", 500, 20))
			})
		})

		Context("with a valid duration param", func() {
			It("returns 200 status code", func() {
				writer := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", "/latency?duration=100ms", nil)
				latencyServer(writer, request)
				Expect(writer.Code).To(Equal(200))
			})

			It("returns OK after 100 ms", func() {
				writer := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", "/latency?duration=100ms", nil)
				timeBeforeRequest := time.Now().UnixNano() / int64(time.Millisecond)
				latencyServer(writer, request)
				timeAfterRequest := time.Now().UnixNano() / int64(time.Millisecond)
				latency := timeAfterRequest - timeBeforeRequest
				Expect(writer.Body.Bytes()).To(ContainSubstring("OK"))
				Expect(latency).To(BeNumerically("~", 100, 10))
			})
		})

		Context("with an invalid duration param", func() {
			It("returns 400 status code", func() {
				writer := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", "/latency?duration=onehundred", nil)
				latencyServer(writer, request)
				Expect(writer.Code).To(Equal(400))
			})

			It("returns 'Error: Invalid duration parameter'", func() {
				writer := httptest.NewRecorder()
				request, _ := http.NewRequest("GET", "/latency?duration=onehundred", nil)
				latencyServer(writer, request)
				Expect(writer.Body.Bytes()).To(ContainSubstring("Error: Invalid duration parameter"))
			})
		})
	})
})
